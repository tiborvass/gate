// Copyright (c) 2018 Timo Savola. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#![feature(lang_items)]
#![no_main]
#![no_std]

extern {
    fn __gate_debug_write(data: *const u8, size: usize);
    fn __gate_exit(status: i32) -> !;
}

fn gate_exit(status: i32) -> ! {
    unsafe {
        __gate_exit(status)
    }
}

fn gate_debug(s: &str) {
    unsafe {
        __gate_debug_write(s.as_ptr(), s.len())
    }
}

#[no_mangle]
pub fn main() {
    gate_debug("rusty world\n")
}

#[lang = "panic_fmt"]
fn panic_fmt() -> ! {
    gate_exit(1)
}
