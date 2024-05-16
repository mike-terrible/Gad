// 
// declare.rs
//
package main

import "strings"

func GenDeclare(nv int, p [256]string)  {
  var i = 0;
  var varV  = ""; 
  var vtype = "";
  var val = "";
  i += 1; if i < nv { varV =  p[i]; };
  i += 1; if i < nv { /* like = p[i]; */ };
  i += 1; if i < nv { vtype = p[i] };
  i += 1; if i < nv { /* be = p[i]; */};
  i += 1; if i< nv { val = p[i]; };
  goVar(varV, vtype, val);
}

func goVar(varV string, vtype string, val string )  {
  if len(varV) == 0 { return; };
  To(Ident);
  switch Mode {
  case "-go","-mojo": { Wr("var "); Wr(varV); }
  case "-rust": {
    if InProc {
      Wr("let mut "); Wr(varV);
    } else { Wr("static mut "); Wr(varV); }; 
  }
  case "-python": { Wr(varV); }
  default:
  };
  if len(vtype) > 0 {
    var ztype = OnType(vtype);
    switch Mode {
    case "-go": { 
      Wr(" "); Wr(ztype);
    }
    case "-rust", "-mojo": {
      Wr(":"); Wr(ztype);
    }
    default:
    };  
  };  
  if len(val) > 0 {
    if Cmp(val,ON) {
      switch Mode {
      case "-rust": { Wr(" = true;\n"); }
      case "-go": { Wr(" = true\n"); }
      case "-mojo","-python": { Wr(" = True\n"); }
      default:
      };
      return ;
    };
    if Cmp(val,OFF) {
      switch Mode {
      case "-rust": { Wr(" = false;\n"); }
      case "-go": { Wr(" = false\n"); }
      case "-mojo","-python": { Wr(" = False\n"); }
      default:
      };
      return;
    };
    Wr(" = "); Wr(val);
    if strings.HasPrefix(val, "\"") {  Wr("\""); }
    if Mode == "-rust" { Wr(";"); };
    Wr("\n");
  };
}   

func OnType(vtype string) string {
  if Cmp(vtype,STR) {
    switch Mode {
    case "-go": { return "string"; }
    case "-rust": { return "&str"; }
    case "-mojo": { return "String"; }
    default: { return ""; }
    };
  };
  if Cmp(vtype,NUM) {
    switch Mode {
    case "-rust": { return "i64"; }
    case "-go": { return "int"; }
    case "-mojo": { return "Int"; }
    default: { return ""; }
    };
  };
  if Cmp(vtype,REAL) {
    switch Mode {
    case "-rust": { return "f64"; }
    case "-go": { return "float64"; }
    case "-mojo": { return "Float64"; }
    default: { return ""; }
    };
  };
  if Cmp(vtype,LIGHT) {
    switch Mode {
    case "-go","-rust": { return "bool"; }
    case "-mojo": { return "Bool"; }
    default:
    }; 
  };
  return "";
}

