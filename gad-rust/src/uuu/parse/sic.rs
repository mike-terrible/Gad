// 
// sic.rs
//
use std::fs::File;
use std::io::Write;
use crate::uuu::parse::vars::*;
use crate::uuu::parse::words::*;

pub fn gen_sic(mut f :&File, gen :&str, ident :i32, p :Vec<String> ) -> i32 {
  to(f,ident); let mut i = 0; let nv = p.len();
  loop { i += 1;
    if i >= nv {
      if gen == "-rust" { wr(f,b";"); };
      wr(f,b"\n"); return ident;
    };
    let t = &p[i];
    wr(f,t.as_bytes());
    if t.starts_with("\"") { wr(f,b"\""); }; 
    wr(f,b" ");
  };
  return ident;
}


