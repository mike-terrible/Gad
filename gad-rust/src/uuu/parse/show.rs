// 
// show.rs
//
use std::fs::File;
use std::io::Write;
use crate::uuu::parse::vars::*;
use crate::uuu::parse::words::*;

pub fn gen_show(mut f :&File, gen :&str, ident :i32, p :Vec<String> ) -> i32 {
  let mut id = ident; let mut i = 0; let nv = p.len();
  loop { i += 1; if i >= nv { break; }; let mut t = &p[i];
    if cmp(t,&WITH) { 
       i += 1; if i >= nv { break; }; 
       t = &p[i]; 
       match gen {
       "-rust" => {
         to(f,id); wr(f,b"print!(\"{ } \",");
         wr(f,t.as_bytes()); if t.starts_with("\"") { wr(f, b"\"");  };  wr(f,b");\n");
       },
       "-go" => {
         to(f,id); wr(f,b"print("); wr(f,t.as_bytes()); if t.starts_with("\"") { wr(f,b"\""); };
         wr(f,b",\" \");\n");
       },
       "-mojop" | "python" => {
         to(f,id); wr(f,b"print("); wr(f,t.as_bytes()); 
         if t.starts_with("\"") { wr(f,b"\""); }; 
         wr(f,b",end =\" \")\n");
       },
       _ =>  { },
       }; // match
    }; // if
  }; // loop
  return id; 
}


