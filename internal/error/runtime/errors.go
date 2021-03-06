// Code generated by internal/cmd/runtime-errors.  You can edit it a little bit.

package runtime

var ExecutorErrors = [35]Error{
	10: {"ERR_CONT_EXEC_EXECUTOR", "runtime container", "failed to execute executor"},
	11: {"ERR_EXEC_PRCTL_NOT_DUMPABLE", "runtime executor", "prctl: failed to set not dumpable"},
	12: {"ERR_EXEC_SETRLIMIT_DATA", "runtime executor", "setrlimit: failed to set DATA limit"},
	13: {"ERR_EXEC_FCNTL_GETFD", "runtime executor", "fcntl: failed to get file descriptor flags"},
	14: {"ERR_EXEC_FCNTL_CLOEXEC", "runtime executor", "fcntl: failed to add close-on-exec flag"},
	15: {"ERR_EXEC_SIGNAL_HANDLER", "runtime executor", "signal handler registration failed"},
	16: {"ERR_EXEC_SIGPROCMASK", "runtime executor", "sigprocmask: failed to set mask"},
	17: {"ERR_EXEC_KILL", "runtime executor", "kill call failed"},
	18: {"ERR_EXEC_WAITPID", "runtime executor", "waitpid call failed"},
	19: {"ERR_EXEC_PPOLL", "runtime executor", "ppoll call failed"},
	20: {"ERR_EXEC_RECVMSG", "runtime executor", "recvmsg call failed"},
	21: {"ERR_EXEC_SEND", "runtime executor", "send call failed"},
	22: {"ERR_EXEC_VFORK", "runtime executor", "vfork call failed"},
	23: {"ERR_EXEC_MSG_CTRUNC", "runtime executor", "received truncated control message"},
	24: {"ERR_EXEC_CMSG_LEVEL", "runtime executor", "unexpected control message: not at socket level"},
	25: {"ERR_EXEC_CMSG_TYPE", "runtime executor", "unexpected control message type: no file descriptors"},
	26: {"ERR_EXEC_CMSG_LEN", "runtime executor", "unexpected control message length"},
	27: {"ERR_EXEC_CMSG_NXTHDR", "runtime executor", "multiple control message headers per recvmsg"},
	28: {"ERR_EXEC_SENDBUF_OVERFLOW_CMSG", "runtime executor", "send buffer overflow on recvmsg"},
	29: {"ERR_EXEC_SENDBUF_OVERFLOW_REAP", "runtime executor", "send buffer overflow on waitpid"},
	30: {"ERR_EXEC_KILLBUF_OVERFLOW", "runtime executor", "kill buffer overflow"},
	31: {"ERR_EXEC_DEADBUF_OVERFLOW", "runtime executor", "dead pid buffer overflow"},
	32: {"ERR_EXEC_KILLMSG_PID", "runtime executor", "received kill message with invalid pid"},
	34: {"ERR_EXEC_PRLIMIT", "runtime executor", "prlimit call failed"},
}

var ProcessErrors = [42]Error{
	2:  {"ERR_RT_RECV", "process runtime", "preadv2 call failed"},
	3:  {"ERR_RT_SEND", "process runtime", "write call failed"},
	4:  {"ERR_RT_DEBUG", "process runtime", "debug write call failed"},
	5:  {"ERR_RT_MPROTECT", "process runtime", "mprotect call failed"},
	11: {"ERR_EXECHILD_DUP2", "process executor", "child: dup2 call failed"},
	12: {"ERR_EXECHILD_NICE", "process executor", "child: nice call failed"},
	13: {"ERR_EXECHILD_SETRLIMIT_NOFILE", "process executor", "child: setrlimit: failed to set NOFILE limit"},
	14: {"ERR_EXECHILD_SETRLIMIT_NPROC", "process executor", "child: setrlimit: failed to set NPROC limit"},
	15: {"ERR_EXECHILD_PRCTL_TSC_SIGSEGV", "process executor", "child: prctl: failed to set PR_TSC_SIGSEGV"},
	16: {"ERR_EXECHILD_PERSONALITY_ADDR_NO_RANDOMIZE", "process executor", "child: failed to change personality to ADDR_NO_RANDOMIZE"},
	20: {"ERR_EXECHILD_EXEC_LOADER", "process executor", "child: failed to execute loader"},
	21: {"ERR_LOAD_PRCTL_NOT_DUMPABLE", "process loader", "prctl: set not dumpable"},
	22: {"ERR_LOAD_PERSONALITY_DEFAULT", "process loader", "failed to set default personality"},
	23: {"ERR_LOAD_READ_INFO", "process loader", "failed to read image info from input fd"},
	24: {"ERR_LOAD_MAGIC_1", "process loader", "magic number #1 mismatch"},
	25: {"ERR_LOAD_MAGIC_2", "process loader", "magic number #2 mismatch"},
	26: {"ERR_LOAD_MMAP_VECTOR", "process loader", "failed to allocate import vector via mmap"},
	27: {"ERR_LOAD_MPROTECT_VECTOR", "process loader", "mprotect: import vector write-protection failed"},
	28: {"ERR_LOAD_MMAP_TEXT", "process loader", "failed to mmap text section of image"},
	29: {"ERR_LOAD_MMAP_STACK", "process loader", "failed to mmap stack section of image"},
	30: {"ERR_LOAD_MMAP_HEAP", "process loader", "failed to mmap globals/memory section of image"},
	31: {"ERR_LOAD_CLOSE_IMAGE", "process loader", "failed to close image fd"},
	32: {"ERR_LOAD_MUNMAP_STACK", "process loader", "failed to munmap initial stack"},
	33: {"ERR_LOAD_SIGACTION", "process loader", "sigaction call failed"},
	34: {"ERR_LOAD_MUNMAP_LOADER", "process loader", "failed to munmap loader .text and .rodata"},
	35: {"ERR_LOAD_SECCOMP", "process loader", "seccomp call failed"},
	36: {"ERR_LOAD_ARGC", "process loader", "loader executed with arguments"},
	39: {"ERR_LOAD_FCNTL_OUTPUT", "process loader", "failed to set output file flags"},
	40: {"ERR_LOAD_MPROTECT_HEAP", "process loader", "mprotect: globals/memory protection failed"},
}

var ErrorsInitialized struct{}
