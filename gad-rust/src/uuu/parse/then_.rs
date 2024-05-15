// 
// then_.rs
//
use std::fs::File;
use std::io::Write;
use crate::uuu::parse::vars::*;
use crate::uuu::parse::words::*;

pub fn gen_then(mut f :&File, gen :&str, ident :i32, _p :Vec<String> ) -> i32 {
  let mut id = ident; 
  match gen {
  "-go" | "-rust" => { wr(f,b" {\n");  },
  "-mojo" | "-python" => { wr(f,b" :\n"); },
  _ => { },
  };  
  id += 2;
  return id;
}


