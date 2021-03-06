# Gate

Run untrusted code from anonymous sources.  Instead of sending messages
composed of passive data, send programs which can react to their environment.
Migrate or duplicate running applications across hosts and computer
architectures.  Gate is a toolkit which aims to enable such things.

- License: [3-clause BSD](LICENSE)
- Author: Timo Savola <timo.savola@iki.fi>


## Foundations

[WebAssembly](https://webassembly.org) is the interchange format of the user
programs.  However, the APIs are different from the browsers' usual WebAssembly
environments.  See low-level [C API](C.md) or the higher-level
[Rust crate](https://github.com/tsavola/gain) for details.

The sandboxing and containerization features of the Linux kernel provide layers
of security in addition to WebAssembly.  See [Security](Security.md) for
details.

Gate *services* are akin to syscalls, but they work differently.  New services
can be added easily, and available services are discovered at run time.  See
[Service implementation](Service.md) for details.


## Building blocks

Gate appears as [Go](https://golang.org) packages and programs.  The execution
mechanism is implemented in C and assembly, and needs to be built separately
(see below).  It's highly Linux-dependent.  Currently only x86-64 is supported,
but ARM64 support is in development.

Important Go packages:

  - [**wag**](https://godoc.org/github.com/tsavola/wag):
    The WebAssembly compiler
    (implemented in a [separate repository](https://github.com/tsavola/wag)).

  - [**gate/runtime**](https://godoc.org/github.com/tsavola/gate/runtime):
    Core functionality.  Interface to the execution mechanism.

  - [**gate/image**](https://godoc.org/github.com/tsavola/gate/image):
    Executable building and management.

  - [**gate/server/webserver**](https://godoc.org/github.com/tsavola/gate/server/webserver):
    HTTP server component which executes your code on purpose.  It has a
    [RESTful API](Web.md), but some actions can be invoked also via websocket.

  - [**gate/service**](https://godoc.org/github.com/tsavola/gate/service):
    Service implementation support and built-in services.

Programs:

  - **gate**:
    Command-line client.  Uses SSH keys (ed25519) for authentication.

  - **gate-server**:
    Standalone HTTP server with the built-in and plugged-in services.

  - **gate-run**:
    Run your programs locally, with the built-in and plugged-in services.

  - **gate-runtimed**:
    For optionally preconfiguring the execution environment, e.g. as a system
    service.

See the complete [list of Go packages](https://godoc.org/github.com/tsavola/gate).


## Objectives

While code is data, most of the time data cannot be treated as code for safety
reasons.  Change that at the Internet level.  Data encapsulated in code can
describe and transform itself.

Application portability.  Migrate processes between mobile devices and servers
when circumstances change: user presence, resource availability or demand,
continuity etc.

Overhead needs to be low enough so that the system can be practical.  Low
startup latency for request processing.  Low memory overhead for high density
of continually running programs.


## Work in progress

  - [x] Linux x86-64 host support
  - [x] Planned security measures have been implemented
  - [x] HTTP server for running programs
  - [x] Client can communicate with the program it runs on the server
  - [ ] Programs can discover and communicate with their peers on a server
  - [x] Support for WebAssembly version 1
  - [x] Speculative execution security issue mitigations
  - [x] Pluggable authentication
  - [x] Load programs from IPFS
  - [x] Reconnect to program instance
  - [x] Snapshot
  - [ ] Restore (wag already has support)
  - [ ] Clone programs locally or remotely (with or without snapshotting)
  - [ ] Expose program instance at some type of internet endpoint to implement ad-hoc servers
  - [ ] Mechanism for implementing services in a programmer-friendly way
  - [ ] Useful resource control policies need more thought (cgroup configuration etc.)
  - [ ] Stable APIs
  - [ ] ARM64 host support
  - [ ] Additional security measures (such as AppArmor/SELinux profiles)
  - [ ] Android host support
  - [ ] Non-Linux host support

User program support:

  - [x] Low-level C API
  - [x] [Rust](https://github.com/tsavola/gain) support
  - [ ] Improved Rust support
  - [ ] Go support
  - [ ] Approach for splitting WebAssembly app between browser (UI) and server (state)


## Build requirements

The non-Go components can be built with `make`.  They require:

  - Linux
  - gcc or clang
  - pkg-config
  - uidmap (shadow-utils)
  - protobuf-compiler
  - libc-dev
  - libcap-dev
  - libsystemd-dev, unless CGROUP_BACKEND=none is specified for make

`make bin` builds the programs using the Go 1.11 module mechanism.
(Individual packages may be buildable with older Go versions.)

Other make targets: `check` `benchmark` `install` `install-lib` `install-bin`
`install-capabilities` `install-lib-capabilities`

The capabilities targets grant [capabilities](Capabilities.md) for the
installed container binary (lib).  That requires:

  - libcap2-bin


## See also

- [Gain crate for Rust user programs](https://github.com/tsavola/gain)
- [C API for user programs](C.md)
- [ABI for user programs](ABI.md)
- [Web server API](Web.md)
- [Service implementation](Service.md)
- [Security](Security.md)
- [Container capabilities](Capabilities.md)
- [Go packages](https://godoc.org/github.com/tsavola/gate)
- [wag](https://github.com/tsavola/wag)
- [wag-toolchain](https://github.com/tsavola/wag-toolchain)

