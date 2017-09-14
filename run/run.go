// Copyright (c) 2016 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package run

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"syscall"

	"github.com/tsavola/wag"
	"github.com/tsavola/wag/reader"
	"github.com/tsavola/wag/sections"
	"github.com/tsavola/wag/traps"
	"github.com/tsavola/wag/types"
	"github.com/tsavola/wag/wasm"

	"github.com/tsavola/gate/internal/memfd"
)

var (
	pageSize = os.Getpagesize()
)

func roundToPage(size int) uint32 {
	mask := uint32(pageSize) - 1
	return (uint32(size) + mask) &^ mask
}

// checkCurrentGid makes sure that this process belongs to gid.
func checkCurrentGid(gid uint) (err error) {
	currentGroups, err := syscall.Getgroups()
	if err != nil {
		return
	}

	currentGroups = append(currentGroups, syscall.Getgid())

	for _, currentGid := range currentGroups {
		if uint(currentGid) == gid {
			return
		}
	}

	err = fmt.Errorf("this process does not belong to group %d", gid)
	return
}

func randAddrs() (textAddr, heapAddr, stackAddr uint64) {
	b := make([]byte, 12)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	textAddr = randAddr(minTextAddr, maxTextAddr, b[0:4])
	heapAddr = randAddr(minHeapAddr, maxHeapAddr, b[4:8])
	stackAddr = randAddr(minStackAddr, maxStackAddr, b[8:12])
	return
}

func randAddr(minAddr, maxAddr uint64, b []byte) uint64 {
	minPage := minAddr / uint64(pageSize)
	maxPage := maxAddr / uint64(pageSize)
	page := minPage + uint64(endian.Uint32(b))%(maxPage-minPage)
	return page * uint64(pageSize)
}

// imageInfo is like the info object in loader.c
type imageInfo struct {
	TextAddr       uint64
	HeapAddr       uint64
	StackAddr      uint64
	PageSize       uint32
	RODataSize     uint32
	TextSize       uint32
	MemoryOffset   uint32
	InitMemorySize uint32
	GrowMemorySize uint32
	StackSize      uint32
	MagicNumber    uint32
	Arg            int32
}

type Image struct {
	maps *os.File
	info imageInfo
}

func (image *Image) Init(ctx context.Context, rt *Runtime) (err error) {
	numFiles := 1
	err = rt.acquireFiles(ctx, numFiles)
	if err != nil {
		return
	}
	defer func() {
		rt.releaseFiles(numFiles)
	}()

	err = image.init(ctx, rt)
	if err != nil {
		return
	}

	numFiles = 0
	return
}

func (image *Image) init(ctx context.Context, rt *Runtime) (err error) {
	mapsFd, err := memfd.Create("maps", memfd.CLOEXEC|memfd.ALLOW_SEALING)
	if err != nil {
		return
	}

	image.maps = os.NewFile(uintptr(mapsFd), "maps")
	return
}

func (image *Image) Release(rt *Runtime) (err error) {
	if image.maps == nil {
		return
	}

	err = image.maps.Close()
	image.maps = nil

	rt.releaseFiles(1)
	return
}

func (image *Image) Populate(m *wag.Module, growMemorySize wasm.MemorySize, stackSize int32,
) (err error) {
	initMemorySize, _ := m.MemoryLimits()

	if initMemorySize > growMemorySize {
		err = fmt.Errorf("initial memory size %d exceeds maximum memory size %d", initMemorySize, growMemorySize)
		return
	}

	roData := m.ROData()
	text := m.Text()
	data, memoryOffset := m.Data()

	_, err = image.maps.Write(roData)
	if err != nil {
		return
	}

	roDataSize := roundToPage(len(roData))

	_, err = image.maps.WriteAt(text, int64(roDataSize))
	if err != nil {
		return
	}

	textSize := roundToPage(len(text))

	_, err = image.maps.WriteAt(data, int64(roDataSize)+int64(textSize))
	if err != nil {
		return
	}

	globalsMemorySize := roundToPage(memoryOffset + int(growMemorySize))
	totalSize := int64(roDataSize) + int64(textSize) + int64(globalsMemorySize) + int64(stackSize)

	err = image.maps.Truncate(totalSize)
	if err != nil {
		return
	}

	_, err = memfd.Fcntl(int(image.maps.Fd()), memfd.F_ADD_SEALS, memfd.F_SEAL_SHRINK|memfd.F_SEAL_GROW)
	if err != nil {
		return
	}

	textAddr, heapAddr, stackAddr := randAddrs()

	image.info = imageInfo{
		TextAddr:       textAddr,
		HeapAddr:       heapAddr,
		StackAddr:      stackAddr,
		PageSize:       uint32(pageSize),
		RODataSize:     roDataSize,
		TextSize:       textSize,
		MemoryOffset:   uint32(memoryOffset),
		InitMemorySize: uint32(initMemorySize),
		GrowMemorySize: uint32(growMemorySize),
		StackSize:      uint32(stackSize),
		MagicNumber:    magicNumber,
		Arg:            image.info.Arg, // in case SetArg was called before this
	}
	return
}

func (image *Image) SetArg(arg int32) {
	image.info.Arg = arg
}

func (image *Image) DumpGlobalsMemoryStack(w io.Writer) (err error) {
	fd := int(image.maps.Fd())

	dataMapOffset := int64(image.info.RODataSize) + int64(image.info.TextSize)

	globalsMemorySize := image.info.MemoryOffset + image.info.GrowMemorySize
	dataSize := int(globalsMemorySize) + int(image.info.StackSize)

	data, err := syscall.Mmap(fd, dataMapOffset, dataSize, syscall.PROT_READ, syscall.MAP_PRIVATE)
	if err != nil {
		panic(err)
	}
	defer syscall.Munmap(data)

	buf := data[:image.info.MemoryOffset]
	fmt.Fprintf(w, "--- GLOBALS (%d kB) ---\n", len(buf)/1024)
	for i := 0; len(buf) > 0; i += 8 {
		fmt.Fprintf(w, "%08x: %x\n", i, buf[0:8])
		buf = buf[8:]
	}

	buf = data[image.info.MemoryOffset : image.info.MemoryOffset+globalsMemorySize]
	fmt.Fprintf(w, "--- MEMORY (%d kB) ---\n", len(buf)/1024)
	for i := 0; len(buf) > 0; i += 32 {
		fmt.Fprintf(w, "%08x: %x %x %x %x\n", i, buf[0:8], buf[8:16], buf[16:24], buf[24:32])
		buf = buf[32:]
	}

	buf = data[globalsMemorySize:]
	fmt.Fprintf(w, "--- STACK (%d kB) ---\n", len(buf)/1024)
	for i := 0; len(buf) > 0; i += 32 {
		fmt.Fprintf(w, "%08x: %x %x %x %x\n", i, buf[0:8], buf[8:16], buf[16:24], buf[24:32])
		buf = buf[32:]
	}

	fmt.Fprintf(w, "---\n")
	return
}

func (image *Image) DumpStacktrace(w io.Writer, funcMap, callMap []byte, funcSigs []types.Function, ns *sections.NameSection,
) (err error) {
	fd := int(image.maps.Fd())

	offset := int64(image.info.RODataSize) + int64(image.info.TextSize) + int64(image.info.MemoryOffset) + int64(image.info.GrowMemorySize)

	size := int(image.info.StackSize)

	stack, err := syscall.Mmap(fd, offset, size, syscall.PROT_READ, syscall.MAP_PRIVATE)
	if err != nil {
		return
	}
	defer syscall.Munmap(stack)

	return writeStacktraceTo(w, image.info.TextAddr, stack, funcMap, callMap, funcSigs, ns)
}

type Process struct {
	process
	stdin  *os.File // writer
	stdout *os.File // reader
}

func (p *Process) Init(ctx context.Context, rt *Runtime, image *Image, debug io.Writer,
) (err error) {
	numFiles := 5
	if debug != nil {
		numFiles += 2
	}

	err = rt.acquireFiles(ctx, numFiles)
	if err != nil {
		return
	}
	defer func() {
		rt.releaseFiles(numFiles)
	}()

	err = p.init(ctx, rt, image, debug)
	if err != nil {
		return
	}

	numFiles = 0
	return
}

func (p *Process) init(ctx context.Context, rt *Runtime, image *Image, debug io.Writer,
) (err error) {
	var (
		stdinW         *os.File
		stdinBlockR    *os.File
		stdinNonblockR = -1
		stdoutR        *os.File
		stdoutW        *os.File
		debugR         *os.File
		debugW         *os.File
	)

	defer func() {
		if stdinW != nil {
			stdinW.Close()
		}
		if stdinBlockR != nil {
			stdinBlockR.Close()
		}
		if stdinNonblockR >= 0 {
			syscall.Close(stdinNonblockR)
		}
		if stdoutR != nil {
			stdoutR.Close()
		}
		if stdoutW != nil {
			stdoutW.Close()
		}
		if debugR != nil {
			debugR.Close()
		}
		if debugW != nil {
			debugW.Close()
		}
	}()

	stdinBlockR, stdinW, err = os.Pipe()
	if err != nil {
		return
	}

	stdinNonblockR, err = syscall.Open(fmt.Sprintf("/proc/self/fd/%d", stdinBlockR.Fd()), syscall.O_RDONLY|syscall.O_CLOEXEC|syscall.O_NONBLOCK, 0)
	if err != nil {
		return
	}

	stdoutR, stdoutW, err = os.Pipe()
	if err != nil {
		return
	}

	if debug != nil {
		debugR, debugW, err = os.Pipe()
		if err != nil {
			return
		}
	}

	err = rt.executor.execute(ctx, &p.process, &execFiles{stdinBlockR, stdinNonblockR, stdoutW, image.maps, debugW})
	if err != nil {
		return
	}

	if debug != nil {
		go copyCloseRelease(rt, debug, debugR)
	}

	p.stdin = stdinW
	p.stdout = stdoutR

	stdinW = nil
	stdinBlockR = nil
	stdinNonblockR = -1
	stdoutR = nil
	stdoutW = nil
	debugR = nil
	debugW = nil
	return
}

func (p *Process) Kill(rt *Runtime) {
	if p.stdin == nil {
		return
	}

	p.process.kill()
	p.stdin.Close()
	p.stdout.Close()

	p.stdin = nil
	p.stdout = nil

	rt.releaseFiles(2)
	return
}

type execFiles struct {
	stdinBlock    *os.File
	stdinNonblock int
	stdout        *os.File
	maps          *os.File // Borrowed
	debug         *os.File // Optional
}

func (files *execFiles) fds() (fds []int) {
	if files.debug == nil {
		fds = make([]int, 4)
	} else {
		fds = make([]int, 5)
	}

	fds[0] = int(files.stdinBlock.Fd())
	fds[1] = files.stdinNonblock
	fds[2] = int(files.stdout.Fd())
	fds[3] = int(files.maps.Fd())

	if files.debug != nil {
		fds[4] = int(files.debug.Fd())
	}
	return
}

func (files *execFiles) release(limiter FileLimiter) {
	numFiles := 3
	files.stdinBlock.Close()
	syscall.Close(files.stdinNonblock)
	files.stdout.Close()

	// don't close maps

	if files.debug != nil {
		numFiles++
		files.debug.Close()
	}

	limiter.release(numFiles)
}

func copyCloseRelease(rt *Runtime, w io.Writer, r *os.File) {
	defer rt.releaseFiles(1)
	defer r.Close()

	io.Copy(w, r)
}

// InitImageAndProcess is otherwise same as Image.Init() + Process.Init(), but
// avoids deadlocks by allocating all required file descriptors in a single
// step.
func InitImageAndProcess(ctx context.Context, rt *Runtime, image *Image, proc *Process, debug io.Writer,
) (err error) {
	numFiles := 6
	if debug != nil {
		numFiles += 2
	}

	err = rt.acquireFiles(ctx, numFiles)
	if err != nil {
		return
	}
	defer func() {
		rt.releaseFiles(numFiles)
	}()

	err = image.init(ctx, rt)
	if err != nil {
		return
	}

	err = proc.init(ctx, rt, image, debug)
	if err != nil {
		return
	}

	numFiles = 0
	return
}

func Load(m *wag.Module, r reader.Reader, rt *Runtime, textBuf wag.Buffer, roDataBuf []byte, startTrigger chan<- struct{},
) error {
	m.MainSymbol = MainSymbol
	return m.Load(r, rt.Environment(), textBuf, roDataBuf, RODataAddr, startTrigger)
}

func Run(ctx context.Context, rt *Runtime, proc *Process, image *Image, services ServiceRegistry,
) (exit int, trap traps.Id, err error) {
	if services == nil {
		services = noServices{}
	}

	err = binary.Write(proc.stdin, endian, &image.info)
	if err != nil {
		return
	}

	err = ioLoop(ctx, services, proc)
	if err != nil {
		return
	}

	status, err := proc.killWait()
	if err != nil {
		return
	}

	switch {
	case status.Exited():
		code := status.ExitStatus()

		switch code {
		case 0, 1:
			exit = code
			return
		}

		if n := code - 100; n >= 0 && n < int(traps.NumTraps) {
			trap = traps.Id(n)
			return
		}

		err = fmt.Errorf("process exit code: %d", code)
		return

	case status.Signaled():
		err = fmt.Errorf("process termination signal: %d", status.Signal())
		return

	default:
		err = fmt.Errorf("unknown process status: %d", status)
		return
	}
}

type Instance struct {
	Image

	proc Process
}

func (inst *Instance) Init(ctx context.Context, rt *Runtime, debug io.Writer,
) error {
	return InitImageAndProcess(ctx, rt, &inst.Image, &inst.proc, debug)
}

func (inst *Instance) Kill(rt *Runtime) (err error) {
	err = inst.Image.Release(rt)
	inst.proc.Kill(rt)
	return
}

func (inst *Instance) Run(ctx context.Context, rt *Runtime, services ServiceRegistry,
) (exit int, trap traps.Id, err error) {
	return Run(ctx, rt, &inst.proc, &inst.Image, services)
}
