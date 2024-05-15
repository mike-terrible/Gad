// 
// proc.rs
//
use std::fs::File;
use std::io::Write;
use crate::uuu::parse::vars::*;
use crate::uuu::parse::words::*;
use crate::uuu::parse::declare::*;

pub fn gen_proc(mut f :&File, gen :&str, ident :i32, p :Vec<String> ) -> i32 {
  let mut id = ident; let nv = p.len(); let mut narg = 0;
  set_in_proc(true);
  let mut i = 0;
  let mut t = &p[i];
  to(f,id);
  match gen {
  "-go" => { wr(f, b"func "); },
  "-python" => { wr(f, b"def "); },
  "-mojo" => { wr(f, b"fn "); },
  "-rust" => { wr(f, b"unsafe fn "); },
  _ => { },
  };
  i += 1; if i>= nv { return id; }; let xn = &p[i];
  wr(f,xn.as_bytes()); wr(f,b"(");
  loop { i += 1; if i >= nv { wr(f,b") "); break; };
    let it_is = &p[i];
    if cmp(it_is, &RETURN) {
      i += 1; if i>= nv { break; }; let act = &p[i];
      let ztype = ontype(gen, act.to_string());
      if gen == "-python" { wr(f,b") :\n"); id += 2; return id; };
      let nz = ztype.len();
      if nz > 0 {
         match gen {
         "-go" => {
            wr(f,b") "); wr(f,ztype.as_bytes()); wr(f, b" {\n"); id += 2; return id;
         },
         "-mojo" => {
           wr(f,b") -> "); wr(f,ztype.as_bytes()); wr(f,b" :\n"); 
           id += 2; return id;
         },
         "-rust" => {
           wr(f, b") -> ");
           if ztype == "&str" { wr(f, b"String"); } else { wr(f, ztype.as_bytes() ); };
           wr(f, b" {\n");
           id += 2; 
           return id; 
         },
         _ => { },
         };
      }; // nz > 0
    }; // RETURN
    if cmp(it_is, &IS) {
      match gen {
      "-rust" | "-go" => { wr(f,b") {\n"); }, 
      "-mojo" | "-python" => { wr(f,b") :\n"); },
      _ => { },
      };
      id += 2; 
      return id;
    };
    if cmp(it_is, &WITH) {
      i += 1; if i >= nv { return id; }; let var = &p[i]; narg += 1;
      i += 1; if i >= nv { return id; }; if narg >1 { wr(f,b","); };
      wr(f,var.as_bytes());
      let like = &p[i];
      if cmp(like, &AKA) {
        i += 1; if i >= nv { return id; }; 
        let mut xtype = &p[i];
        let ztype = ontype(gen,xtype.to_string());
        match gen {
        "-go" => { wr(f,b" "); wr(f, ztype.as_bytes() ); }, 
        "-mojo" | "-rust" => { wr(f,b" :"); wr(f,ztype.as_bytes()); },
        _ => { },
        }; // match 
      };  // AKA    
    };  // WITH
  }; // loop
  return id;
}

