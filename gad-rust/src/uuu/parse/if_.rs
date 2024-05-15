// 
// if_.rs
//
use std::fs::File;
use std::io::Write;
use crate::uuu::parse::vars::*;
use crate::uuu::parse::words::*;

pub fn gen_if(mut f :&File, gen :&str, ident :i32, p :Vec<String> ) -> i32 {
  let mut id = ident;  let mut i = 0; let nv = p.len();
  to(f,id); wr(f,b"if");
  loop {
    i += 1; if i >= nv { break; };
    let t = &p[i];
    if cmp(t, &THEN) {
      match gen {
      "-rust" | "-go" => { wr(f, b" {\n");  id += 2; return id; },
      "-mojo" | "-python" => { wr(f, b" :\n"); id += 2; return  id;},
      _ => { },
      };
      return id;
    };
    wr(f,b" "); wr(f,t.as_bytes()); if t.starts_with("\"") { wr(f,b"\""); };    
  };
  return id;
}

