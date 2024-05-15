// 
// job.rs
//
use std::fs::File;
use std::io::Write;
use crate::uuu::parse::vars::*;
use crate::uuu::parse::words::*;

pub fn gen_job(mut f :&File, gen :&str, ident :i32, p :Vec<String> ) -> i32 {

  let nv = p.len();
  let mut id = ident;
  let mut i = 0;
  i += 1; if i >= nv { return id; };
  let mut t = &p[i];
  wr(f,b"\n"); to(f,id); wr(f,t.as_bytes()); wr(f,b"(");
  let mut np = 0; 
  loop { i += 1; if i >= nv { break; }; t = &p[i];
    if cmp(t,&WITH) { i += 1; if i >= nv { break; }; 
      t = &p[i]; np += 1; if np > 1 { wr(f,b","); };
      wr(f, t.as_bytes());
      if t.starts_with("\"") { wr(f,b"\""); };
    }; 
  };
  if gen == "-rust" {
    wr(f, b");\n");
  } else {
    wr(f,b")\n");
  };
  return id;

}


