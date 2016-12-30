package run_test

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/tsavola/wag"
	"github.com/tsavola/wag/dewag"
	"github.com/tsavola/wag/wasm"

	"."
)

type readWriter struct {
	io.Reader
	io.Writer
}

const (
	dumpText = true
)

func TestRun(t *testing.T) {
	const (
		memorySizeLimit = wasm.Page
		stackSize       = 4096
	)

	executorBin := os.Getenv("GATE_TEST_EXECUTOR")
	loaderBin := os.Getenv("GATE_TEST_LOADER")
	wasmPath := os.Getenv("GATE_TEST_WASM")

	env, err := run.NewEnvironment(executorBin, loaderBin, loaderBin+".symbols")
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.Open(wasmPath)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)

	foo, err := run.NewFoo()
	if err != nil {
		t.Fatal(err)
	}
	defer foo.Close()

	var m wag.Module

	err = m.Load(r, env, &foo.Text, nil, run.RODataAddr, nil)
	if err != nil {
		t.Fatalf("load error: %v", err)
	}

	if _, err := foo.ROData.Write(m.ROData()); err != nil {
		t.Fatal(err)
	}

	data, memoryOffset := m.Data()

	if _, err := foo.Data.Write(data); err != nil {
		t.Fatal(err)
	}

	if dumpText && testing.Verbose() {
		dewag.PrintTo(os.Stdout, foo.Text.Bytes(), m.FunctionMap(), nil)
	}

	initMemorySize, growMemorySize := m.MemoryLimits()
	if growMemorySize > memorySizeLimit {
		growMemorySize = memorySizeLimit
	}

	payload, err := run.NewPayload(foo, memoryOffset, initMemorySize, growMemorySize, stackSize)
	if err != nil {
		t.Fatalf("payload error: %v", err)
	}
	defer payload.Close()

	var output bytes.Buffer

	exit, trap, err := run.Run(env, payload, readWriter{new(bytes.Buffer), &output}, os.Stdout)
	t.Logf("output: %#v\n", string(output.Bytes()))
	if err != nil {
		t.Fatalf("run error: %v", err)
	} else if trap != 0 {
		t.Fatalf("run trap: %s", trap)
	} else if exit != 0 {
		t.Fatalf("run exit: %s", exit)
	}

	if name := os.Getenv("GATE_TEST_DUMP"); name != "" {
		f, err := os.Create(name)
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()

		if err := payload.DumpGlobalsMemoryStack(f); err != nil {
			t.Fatalf("dump error: %v", err)
		}
	}
}
