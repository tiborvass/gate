package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path"

	"github.com/tsavola/wag"
	"github.com/tsavola/wag/dewag"
	"github.com/tsavola/wag/sections"

	"github.com/tsavola/gate/run"
)

type readWriteCloser struct {
	io.Reader
	io.WriteCloser
}

const (
	dumpText       = true
	dumpStacktrace = true
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	var (
		executor      = path.Join(dir, "bin/executor")
		loader        = path.Join(dir, "bin/loader")
		loaderSymbols = loader + ".symbols"
		stackSize     = 16 * 1024 * 1024
		stdio         = false
		addr          = ""
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] wasm\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.StringVar(&executor, "executor", executor, "filename")
	flag.StringVar(&loader, "loader", loader, "filename")
	flag.StringVar(&loaderSymbols, "loader-symbols", loaderSymbols, "filename")
	flag.IntVar(&stackSize, "stack-size", stackSize, "stack size")
	flag.BoolVar(&stdio, "stdio", stdio, "use stdio (conflicts with -addr)")
	flag.StringVar(&addr, "addr", addr, "socket path (conflicts with -stdio)")

	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
		os.Exit(2)
	}

	wasm, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}

	env, err := run.NewEnvironment(executor, loader, loaderSymbols)
	if err != nil {
		log.Fatal(err)
	}

	var ns sections.NameSection

	m := wag.Module{
		UnknownSectionLoader: sections.UnknownLoaders{"name": ns.Load}.Load,
	}

	err = m.Load(bufio.NewReader(wasm), env, new(bytes.Buffer), nil, run.RODataAddr, nil)
	if err != nil {
		return
	}

	_, memorySize := m.MemoryLimits()

	payload, err := run.NewPayload(&m, memorySize, int32(stackSize))
	if err != nil {
		return
	}
	defer payload.Close()

	if dumpText {
		dewag.PrintTo(os.Stderr, m.Text(), m.FunctionMap(), &ns)
	}

	var conn io.ReadWriteCloser

	if stdio {
		conn = readWriteCloser{os.Stdin, os.Stdout}
	} else if addr != "" {
		os.Remove(addr)
		l, err := net.Listen("unix", addr)
		if err != nil {
			log.Fatal(err)
		}
		conn, err = l.Accept()
		l.Close()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		conn, err = os.OpenFile(os.DevNull, os.O_RDWR, 0)
		if err != nil {
			log.Fatal(err)
		}
	}

	defer conn.Close()

	exit, trap, err := run.Run(env, payload, conn, os.Stderr)
	if err != nil {
		log.Print(err)
	} else if trap != 0 {
		log.Printf("trap: %s", trap)
	} else if exit != 0 {
		log.Printf("exit: %d", exit)
	}

	if dumpStacktrace {
		err := payload.DumpStacktrace(os.Stderr, m.FunctionMap(), m.CallMap(), m.FunctionSignatures(), &ns)
		if err != nil {
			log.Print(err)
		}
	}
}
