// 
// else_.rs
//
use std::fs::File;
use std::io::Write;
use crate::uuu::parse::vars::*;
use crate::uuu::parse::words::*;

pub fn gen_else(mut f :&File, gen :&str, ident :i32, _p :Vec<String> ) -> i32 {
  let mut id = ident; id -= 2; 
  match gen {
  "-go" | "-rust" => {
    to(f,id); wr(f, b"} else {\n"); 
  },
  "-mojo" | "-python" => {
    to(f,id); wr(f,b"else:\n");
  },
  _ => { },
  }; 
  id += 2;
  return id;
}



