//
// amen.rs
//
use std::fs::File;
use std::io::Write;

use crate::uuu::parse::vars::*;

pub fn gen_loop(mut f :&File, gen :&str, ident :i32 ) -> i32 {
  
  wr(f,b"\n"); let ii :i32 = ident - 2; to(f,ii);
  match gen {
  "-rust" | "-go" => { let _ = wr(f, b"};\n"); },
  "-mojo" | "-python" => { let _ = wr(f, b"pass\n"); },
  _ => { },
  }; 
  return ii;
}

pub fn gen_done(mut f :&File, gen :&str, ident :i32 ) -> i32 { return gen_loop( f, gen, ident ); }

pub fn gen_amen(mut f :&File, gen :&str, ident :i32 ) -> i32 {
  
  set_in_proc(false); 
  wr(f,b"\n"); let ii :i32 = ident - 2; to(f,ii);
  match gen {
  "-rust" | "-go" => { let _ = wr(f, b"}\n"); },
  "-mojo" | "-python" => { let _ = wr(f, b"pass\n"); },
  _ => { },
  };
  
  return ii;
}


