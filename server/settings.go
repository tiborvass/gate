package server

import (
	"io"

	"github.com/tsavola/gate/run"
	"github.com/tsavola/wag/wasm"
)

type Settings struct {
	MemorySizeLimit wasm.MemorySize
	StackSize       int32
	Env             *run.Environment
	Services        func(io.Reader, io.Writer, func(io.Reader, io.Writer) (uint64, bool)) run.ServiceRegistry
	Log             Logger
	Debug           io.Writer
}
