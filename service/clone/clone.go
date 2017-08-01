package clone

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/tsavola/gate/service"
)

const (
	Name    = "clone"
	Version = 0

	packetHeaderSize = 8
)

type clone struct {
	f func(io.Reader, io.Writer) (uint64, bool)
	m map[uint64]struct{}
}

func Register(r *service.Registry, f func(io.Reader, io.Writer) (uint64, bool)) {
	service.Register(r, Name, Version, &clone{
		f: f,
		m: make(map[uint64]struct{}),
	})
}

func (c *clone) New() service.Instance {
	return c
}

func (c *clone) Handle(buf []byte, replies chan<- []byte) {
	buf = buf[:packetHeaderSize]

	var (
		id uint64
		ok bool
	)

	id, ok = c.f(bytes.NewReader(nil), ioutil.Discard)
	if ok {
		c.m[id] = struct{}{}
		buf = append(buf, fmt.Sprintf("%016x", id)...)
	}

	replies <- buf
}

func (c *clone) Shutdown() {
}
