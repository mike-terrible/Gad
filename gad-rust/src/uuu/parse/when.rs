// 
// when.rs
//
use std::fs::File;
use std::io::Write;
use crate::uuu::parse::vars;
use crate::uuu::parse::words;

pub fn gen_when(mut f :&File, gen :&str, ident :i32, p :Vec<String> ) -> i32 {
  let mut id :i32 = ident;
  vars::to(f,id);
  match gen {
  "-go" => { vars::wr(f,b"for"); },
  "-rust" | "-mojo" | "-python" => { vars::wr(f,b"while"); },
  _ => { },
  };
  id += 2; let mut i = 0; let nv = p.len();
  loop {
    i += 1; if i >= nv { break; };
    let t = &p[i];  
    let be = words::cmp(t,&["repeat","повтор"]);
    if be {
      match gen {
      "-go" | "-rust" => {  vars::wr(f,b" {\n"); return id; }, 
      "-mojo" | "-python" => { vars::wr(f,b" :\n"); return id; },
      _ => { },
      };
      return id;
    };
    vars::wr(f,b" "); 
    vars::wr(f,t.as_bytes());
    if t.starts_with("\"") { vars::wr( f, b"\"" ); };
  };
  return id;
}

