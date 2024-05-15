// 
// is_.rs
//
use std::fs::File;
use std::io::Write;
use crate::uuu::parse::vars::*;
use crate::uuu::parse::words::*;

pub fn gen_is(mut f :&File, gen :&str, ident :i32, p :Vec<String> ) -> i32 {
  let mut id = ident; let nv = p.len();
  match gen {
  "-go" | "-rust" => {
    wr(f, b"{\n");
  },
  _ => { },
  };  
  id += 2;
  return id;
}

