
/*
  rust implementation for gad compiler 
*/
use std::env;
//use std::fs;

mod uuu;

fn main() {
  let mut fname :String = "?".to_string();
  let mut gen :String = "-go".to_string();
  let mut i = 0;
  let mut has_fname = false;
  println!("gad compiler rel 2.0");
  for ar in env::args() {
    println!("i = {},arg = {}",i,ar);
    match i {
      1 => { fname = ar;
             has_fname = true;
           },
      2 => { gen = ar; },
      _ => { },
    };
    i += 1
  }
  if has_fname {
    println!("fname = { },gen = { }",fname,gen);
    uuu::parse::parser(&fname,&gen)
  }
  else {
    println!("USAGE: gad file-name [-go | -rust | -mojo | -python ]");
  };
}
