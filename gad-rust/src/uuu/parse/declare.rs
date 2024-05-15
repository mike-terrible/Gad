// 
// declare.rs
//
use std::fs::File;
use std::io::Write;
use crate::uuu::parse::vars::*;
use crate::uuu::parse::words::*;

pub fn gen_declare(mut f :&File, gen :&str, ident :i32, p :Vec<String> ) -> i32 {
  let mut id = ident; let nv = p.len();
  let mut i = 0;
  let mut var :&str = ""; 
  let mut like :&str = ""; 
  let mut vtype :&str = "";
  let mut be :&str = "";
  let mut val :&str = "";
  i += 1; if i < nv { var =  &p[i]; };
  i += 1; if i < nv { like = &p[i]; };
  i += 1; if i < nv { vtype = &p[i] };
  i += 1; if i < nv { be = &p[i]; };
  i += 1; if i< nv { val = &p[i]; };
  id = goVar(f, gen, id, var, vtype, val);
  return id;
}

fn goVar(mut f :&File, gen :&str, ident :i32, var :&str, vtype :&str, val :&str ) -> i32 {
  let mut id = ident;
  if var.len() == 0 { return id; };
  to(f,id);
  match gen {
  "-go" | "-mojo" => { wr(f,b"var "); wr(f,var.as_bytes()); },
  "-rust" => {
    if get_in_proc() {
      wr(f,b"let mut "); wr(f,var.as_bytes());
    } else { wr(f,b"static mut "); wr(f,var.as_bytes()); }; 
  },
  "-python" => { wr(f, var.as_bytes()); },
  _ => { },
  };
  if vtype.len() > 0 {
    let ztype = ontype(gen,vtype.to_string());
    match gen {
    "-go" => { 
      wr(f,b" "); wr(f,ztype.as_bytes());
    },
    "-rust" | "-mojo" => {
      wr(f,b":"); wr(f,ztype.as_bytes());
    },
    
    _ => { },
    };  
  };  
  if val.len() > 0 {
    if cmp(val,&ON) {
      match gen {
      "-rust" => { wr(f, b" = true;\n"); },
      "-go" => { wr(f, b" = true\n"); },
      "-mojo" | "-python" => { wr(f, b" = True\n"); },
      _ => { },
      };
      return id;
    };
    if cmp(val,&OFF) {
      match gen {
      "-rust" => { wr(f, b" = false;\n"); },
      "-go" => { wr(f, b" = false\n"); },
      "-mojo" | "-python" => { wr(f, b" = False\n"); },
      _ => { },
      };
      return id;
    };
    wr(f,b" = "); wr(f,val.as_bytes());
    if val.starts_with("\"") {  wr(f,b"\""); }
    if gen == "-rust" { wr(f,b";"); };
    wr(f,b"\n");
  };
  return id;
}   

pub fn ontype(gen :&str,vtype :String) -> &str {
  if cmp(&vtype,&STR) {
    match gen {
    "-go" => { return "string"; },
    "-rust" => { return "&str"; },
    "-mojo" => { return "String"; },
    _ => { return ""; },
    };
  };
  if cmp(&vtype,&NUM) {
    match gen {
    "-rust" => { return "i64"; },
    "-go" => { return "int"; },
    "-mojo" => { return "Int"; },
     _ => { return ""; },
    };
  };
  if cmp(&vtype,&REAL) {
    match gen {
    "-rust" => { return "f64"; },
    "-go" => { return "float64"; },
    "-mojo" => { return "Float64"; },
    _ => { return ""; },
    };
  };
  if cmp(&vtype,&LIGHT) {
    match gen {
    "-go" | "-rust" => { return "bool"; },
    "-mojo" => { return "Bool"; },
    _ => { },
    }; 
  };
  return "";
}

