//
// comment.rs
//
use std::fs::File;
use std::io::Write;
use crate::uuu::parse::vars::*;

pub fn gen_comment(mut f :&File, gen :&str, a :&str) {
 match gen {
 "-rust" | "-go" => {
   wr(f, b"// "); 
   wr(f,a.as_bytes());
 },
 "-mojo" | "python" => {
   wr(f, b"# ");
   wr(f,a.as_bytes());  
 },
 _ => { },
 };
}

