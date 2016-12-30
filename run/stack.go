package run

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"sort"

	"github.com/tsavola/wag/sections"
	"github.com/tsavola/wag/types"
)

type callSite struct {
	index       uint64
	stackOffset int
}

func findCaller(funcMap []byte, retAddr uint32) (num int, initial, ok bool) {
	count := len(funcMap) / 4
	if count == 0 {
		return
	}

	firstFuncAddr := binary.LittleEndian.Uint32(funcMap[:4])
	if retAddr > 0 && retAddr < firstFuncAddr {
		initial = true
		ok = true
		return
	}

	num = sort.Search(count, func(i int) bool {
		i++
		if i == count {
			return true
		} else {
			return retAddr <= binary.LittleEndian.Uint32(funcMap[i*4:(i+1)*4])
		}
	})

	if num < count {
		ok = true
	}
	return
}

func getCallSites(callMap []byte) (callSites map[int]callSite) {
	callSites = make(map[int]callSite)

	for i := 0; len(callMap) > 0; i++ {
		entry := binary.LittleEndian.Uint64(callMap[:8])
		callMap = callMap[8:]

		addr := int(uint32(entry))
		stackOffset := int(entry >> 32)

		callSites[addr] = callSite{uint64(i), stackOffset}
	}

	return
}

func writeStacktraceTo(w io.Writer, stack, funcMap, callMap []byte, funcSigs []types.Function, ns *sections.NameSection) (err error) {
	unused := binary.LittleEndian.Uint64(stack)
	if unused == 0 {
		err = errors.New("no stack")
		return
	}
	if unused > uint64(len(stack)) || (unused&7) != 0 {
		err = errors.New("corrupted stack")
		return
	}
	stack = stack[unused:]

	callSites := getCallSites(callMap)

	depth := 1

	for ; len(stack) > 0; depth++ {
		absoluteRetAddr := binary.LittleEndian.Uint64(stack[:8])

		retAddr := absoluteRetAddr - textAddr
		if retAddr > 0x7ffffffe {
			fmt.Fprintf(w, "#%d  <absolute return address 0x%x is not in text section>\n", depth, absoluteRetAddr)
			return
		}

		funcNum, start, ok := findCaller(funcMap, uint32(retAddr))
		if !ok {
			fmt.Fprintf(w, "#%d  <function not found for return address 0x%x>\n", depth, retAddr)
			return
		}

		site, found := callSites[int(retAddr)]
		if !found {
			fmt.Fprintf(w, "#%d  <unknown return address 0x%x>\n", depth, retAddr)
			return
		}

		if start {
			if site.stackOffset != 0 {
				fmt.Fprintf(w, "#%d  <start function call site stack offset is not zero>\n", depth)
			}
			if len(stack) != 8 {
				fmt.Fprintf(w, "#%d  <start function return address is not stored at start of stack>\n", depth)
			}
			return
		}

		if site.stackOffset < 8 || (site.stackOffset&7) != 0 {
			fmt.Fprintf(w, "#%d  <invalid stack offset %d>\n", depth, site.stackOffset)
			return
		}

		stack = stack[site.stackOffset:]

		var name string
		var localNames []string

		if ns != nil && funcNum < len(ns.FunctionNames) {
			name = ns.FunctionNames[funcNum].FunName
			localNames = ns.FunctionNames[funcNum].LocalNames
		} else {
			name = fmt.Sprintf("func-%d", funcNum)
		}

		var sigStr string

		if funcNum < len(funcSigs) {
			sigStr = funcSigs[funcNum].StringWithNames(localNames)
		}

		fmt.Fprintf(w, "#%d  %s%s\n", depth, name, sigStr)
	}

	if len(stack) != 0 {
		fmt.Fprintf(w, "#%d  <%d bytes of untraced stack>\n", depth, len(stack))
	}
	return
}