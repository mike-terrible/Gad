// 
// give.rs
//
use std::fs::File;
use std::io::Write;
use crate::uuu::parse::vars::*;
use crate::uuu::parse::words::*;

pub fn gen_give (mut f :&File, gen :&str, ident :i32, p :Vec<String> ) -> i32 {
  let mut id = ident; 
  let nv = p.len(); 
  let mut i = 0;
  i += 1; 
  if i>= nv { return id; }; 
  let mut t = &p[i];
  to(f,id);
  match gen {
  "-go" | "-mojo" => { wr(f,b"var "); },
  "-rust" => { wr(f,b"let mut "); },
  _ => { },
  };
  wr(f,t.as_bytes()); // var name
  wr(f,b" = ");
  i += 1; if i >= nv { return id; }; // {from}  
  i += 1; if i >= nv { return id; }; 
  t = &p[i]; // proc name
  wr(f,t.as_bytes()); let mut np = 0; // ( 
  loop { 
    i += 1; if i >= nv { break; };
    t = &p[i];
    if cmp(t,&WITH) {
      i += 1; if i>= nv { break; }; 
      let tt = &p[i];
      np += 1; if np >=1 { wr(f,b","); };
      wr(f,tt.as_bytes()); 
      if tt.starts_with("\"") { wr(f,b"\""); };
    };    
  };
  if gen == "-rust" { wr(f,b");\n"); } else  { wr(f,b")\n"); };
  return id;
}


