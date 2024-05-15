// 
// run.rs
//
use std::fs::File;
use std::io::Write;
use crate::uuu::parse::vars::*;
use crate::uuu::parse::words::*;

pub fn gen_run(mut f :&File, gen :&str, ident :i32, p :Vec<String> ) -> i32 {
  let mut id = ident; let nv = p.len(); let mut i = 0; 
  set_in_proc(true);
  let mut t = &p[i];
  i += 1; if i >= nv { return id; };
  let xmain = &p[i]; 
  i += 1; if i >= nv { return id; };
  to(f,id);
  match gen {
  "-rust" => { wr(f, b"fn main() {\n"); },
  "-go" => { wr(f, b"func main() {\n"); },
  "-mojo" => { wr(f, b"fn main() :\n"); },
  "-python" => { wr(f, b"def main() :\n"); },
  _ => { },
  };
  id += 2; to(f,id);
  if gen == "-rust" { wr(f,b"unsafe { "); };
  wr(f,xmain.as_bytes()); wr(f,b"()");
  if gen == "-rust" { wr(f, b"; }"); };
  wr(f, b"\n");
  id -= 2; to(f,id);
  match gen {
  "-go"|"-rust" => { wr(f,b"}\n"); },
  _ => { },
  };
  to(f,id);
  match gen {
  "-rust" => { wr(f,b"unsafe fn "); wr(f, xmain.as_bytes()); wr(f, b"() {\n");  }, 
  "-go" => { wr(f,b"func "); wr(f, xmain.as_bytes()); wr(f, b"() {\n"); }, 
  "-mojo" => { wr(f,b"fn "); wr(f,xmain.as_bytes()); wr(f, b"() :\n"); }, 
  "-python" => { wr(f,b"def "); wr(f,xmain.as_bytes()); wr(f,b"() :\n"); },
  _ => { },
  }; 
  //
  id += 2; return id;
}


