// 
// mess.rs
//
use std::fs::File;
use std::io::Write;
use crate::uuu::parse::vars::*;
use crate::uuu::parse::words::*;

pub fn gen_mess(mut f :&File, gen :&str, ident :i32, p :Vec<String> ) -> i32 {
  let mut id = ident; 
  to(&f,id);
  let mut i = 0; let nv = p.len();
  i += 1; if i >= nv { return id; };
  let mut t = &p[i];
  match gen {
  "-rust" => {
    wr(f,b"println!(");
    if t.starts_with("\"") { 
      wr(f,t.as_bytes()); wr(f,b"\""); 
    } else { 
      wr(f,b"\"{ }\","); wr(f,t.as_bytes() ); 
    };
    wr(f,b");\n"); return id;
  },
  "-go" => {
    wr(f,b"println("); wr(f,t.as_bytes());
    if t.starts_with("\"") { wr(f,b"\""); };
    wr(f,b")\n");
    return id
  },
  "-mojo" | "python" => {
    wr(f,b"print("); wr(f,t.as_bytes());
    if t.starts_with("\"") { wr(f, b"\""); };
    wr(f,b")\n")
  },
  _ => { },
  };
  return id;
}


