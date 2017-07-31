NM		?= nm
OBJCOPY		?= objcopy
STRIP		?= strip

CPPFLAGS	+= -DNDEBUG -DPIE
CFLAGS		+= -std=gnu99 -Os -g -Wall -Wextra -Wno-unused-parameter -fPIE -fomit-frame-pointer -fno-stack-protector -Wno-main
LDFLAGS		+= -static -nostartfiles -nostdlib -Wl,-z,noexecstack -Wl,-Ttext-segment=0x200000000 -Wl,--build-id=none

../../lib/loader: loader.o runtime.o
	mkdir -p $(dir $@)
	$(CC) $(CFLAGS) $(LDFLAGS) -o $@ loader.o runtime.o
	$(NM) $@ | grep " __gate_" > $(dir $@)runtime.map
	$(OBJCOPY) --remove-section=.comment --remove-section=.eh_frame $@
	$(STRIP) $@

loader.o: loader.c ../defs.h Makefile
	$(CC) $(CPPFLAGS) -DLOADER_TEXT_ADDR=0x200000000 -DLOADER_TEXT_UNMAP_LEN=4096 $(CFLAGS) -c -o $@ loader.c

runtime.o: nopad
	$(CC) $(CPPFLAGS) -DRUNTIME_PAD=$$(python -c "print(0x200001000 - 0x$$(nm nopad | grep ' runtime_start_with_syscall$$' | cut -d' ' -f1))") -c -o $@ runtime.S

nopad: loader.o runtime-nopad.o
	$(CC) $(CFLAGS) $(LDFLAGS) -o $@ loader.o runtime-nopad.o

runtime-nopad.o: runtime.S ../defs.h Makefile
	$(CC) $(CPPFLAGS) -DRUNTIME_PAD=0 -c -o $@ runtime.S

clean:
	rm -f *.o nopad

.PHONY: clean
