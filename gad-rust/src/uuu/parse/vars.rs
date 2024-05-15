
// vars.rs
//
use std::fs::File;
use std::io::Write;

pub static mut XMAIN :&str = "";
pub static mut IN_PROC :bool = false;

pub fn wr(mut f :&File, b :&[u8] ) {
  let _ = f.write(b);
}

pub fn to(mut f :&File, n :i32) {
  let mut i = 0;
  while i < n { let _ = f.write(b" "); i += 1; };
}


pub fn set_in_proc(b :bool) {
 unsafe {
   IN_PROC = b;
 };
}

pub fn get_in_proc() -> bool {
  let b  :bool;
  unsafe {
    b = IN_PROC;
  };
  return b;
}
