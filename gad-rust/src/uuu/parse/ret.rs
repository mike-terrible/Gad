// 
// ret.rs
//
use std::fs::File;
use std::io::Write;
use crate::uuu::parse::vars::*;

pub fn gen_return(mut f :&File, gen :&str, ident :i32, p :Vec<String> ) -> i32 {
  let id =  ident;
  let mut i = 0;
  let nv = p.len();
  //
  i += 1;
  if i == nv { to(f,id); wr(f, b"return");
    if gen == "-rust" { wr(f, b";"); };
    wr(f,b"\n");  
    return id;  
  };
  let t = &p[i]; to(f,ident); wr(f, b"return "); 
  wr(f, t.as_bytes() ); 
  if t.starts_with("\"") {
    wr(f, b"\"");
    if gen == "-rust" { wr(f, b".to_string()"); };
  };  
  if gen == "-rust" { wr(f, b";"); };
  wr(f, b"\n"); 
  //
  return id;  
}
