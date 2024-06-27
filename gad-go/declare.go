// 
// declare.go
//
package main

import "fmt"
import "strings"

func GenDeclare(nv int, p [256]string)  {
  var i int = 0;
  var varV string = ""; 
  var vtype string = "";
  var val string = "";
  var vsize string = "";
  var like string = "";
  var be string = "";
  i += 1; if i < nv { varV =  p[i]; };
  i += 1; 
  if i < nv { like = p[i];  
    if Cmp(like,ARRAY) { 
      i += 1; if i < nv { vsize = p[i]; };
      i += 1; if i < nv { vtype = p[i]; }; 
      i +=1;  if i < nv { be = p[i]; };
      goArray(nv,p, varV, vsize, vtype, be ); 
      return;
    };
  };
  i += 1; if i < nv { vtype = p[i] };
  i += 1; if i < nv { /* be = p[i]; */};
  i += 1; if i< nv { val = p[i]; };
  goVar(varV, vtype, val);
}

func GenWith(nv int, p [256]string) {
  var i = 0; var t = ""; var buf = ""; 
  for {
    i += 1; 
    if i >=nv { break; }
    t = p[i];
    switch(Mode) {
    case GO, RUST:
    { To(GetIdent());
      Wr(CurVar);
      buf = fmt.Sprintf("%d",InInit);
      Wr("["); Wr(buf); Wr("]"); Wr(" = "); Wr(t);
      if strings.HasPrefix(t, "\"") { Wr("\""); };
      if Mode == "-rust" { Wr(";"); };
      Wr("\n");
      InInit += 1;
    }  
    case MOJO,PYTHON:
    { To(GetIdent());
      Wr(CurVar); Wr(".append("); Wr(t);
      if strings.HasPrefix(t,"\"") { Wr("\""); };
      Wr(")\n");
    }
    default: 
    };
  };
}

func goArray(nv int, p [256]string, varV string, vsize string, vtype string, be string) {
  if len(varV) == 0 { return; }; To(GetIdent());
  if Mode == GO {
    Wr("var "); Wr(varV); Wr(" "); Wr("["); 
    if  vsize != "?"  { Wr(vsize); };
    Wr("]"); Wr(OnType(vtype));
    if(be == "") { InArray = false; Wr(";\n"); return; };
    CurVar = varV;
    Wr("\n");
    InArray = true; 
    if Cmp(be,AKA) {
      InInit = 0;
    } else { 
      InInit = 1;
      To(GetIdent());
      Wr(varV); Wr("[0] = "); Wr(be);
      if strings.HasPrefix(be,"\"")  { Wr("\""); };
      Wr("\n");
    }; 
    return;
  };
  if Mode == RUST {
    if(InProc)  { Wr("let mut "); Wr(varV); } else {  
      Wr("static mut "); Wr(varV);  
    };
    Wr(": "); Wr("["); Wr(OnType(vtype)); Wr(";");
    if  vsize != "?" { Wr(vsize); };
    Wr("]"); Wr(" = ["); Wr(OnValue(vtype)); Wr(";"); Wr(vsize); Wr("];\n");
    if be == "" { InArray = false; return; };
    CurVar = varV;
    InArray = true;  
    if Cmp(be,AKA) { InInit = 0; } else {
      InInit = 1; To(GetIdent());
      Wr(varV); Wr("[0] = "); Wr(be);
      if strings.HasPrefix(be, "\"")  { Wr("\""); };
      Wr(";");
    };
    Wr("\n");
    return;
  };
  if(Mode == PYTHON) {
    Wr(varV); Wr(" = [ ]\n");
    if(be == "") { InArray = false; Wr("\n"); return; };
    CurVar = varV;
    InArray = true; 
    if Cmp(be,AKA) { InInit = 0; } else {
      InInit = 1
      To(GetIdent());
      Wr(CurVar); Wr(".append("); Wr( be);
      if strings.HasPrefix(be,"\"") { Wr("\""); };
    }; 
    Wr("\n");
    return;
  };
  if(Mode == MOJO) {
    Wr("var "); Wr(varV); Wr(" = List["); Wr(OnType(vtype)); Wr("]()");
    CurVar = varV;
    InArray = true; 
    if be == "" { InArray = false; Wr("\n"); return; };
    if Cmp(be,AKA) { InInit = 0; } else {
      InInit = 1
      To(GetIdent());
      Wr(CurVar);Wr(".append("); Wr( be ); 
      if strings.HasPrefix(be,"\"") { Wr("\""); };
    }; 
    Wr("\n");
    return;
  };
}

func goVar(varV string, vtype string, val string )  {
  if len(varV) == 0 { return; };
  if Mode == ASM { AsmGoVar(varV,vtype,val);  return; }
  To(GetIdent());
  switch Mode {
  case  GO,MOJO: { Wr("var "); Wr(varV); }
  case RUST: {
    if InProc {
      Wr("let mut "); Wr(varV);
    } else { Wr("static mut "); Wr(varV ); }; 
  }
  case PYTHON: { /* Wr(varV); */ }
  default:
  };
  if len(vtype) > 0 {
    var ztype = OnType(vtype);
    switch Mode {
    case GO: { Wr(" "); Wr( ztype ); }
    case RUST,MOJO: { Wr(":"); Wr(ztype); }
    };
  }; 
  if len(val) > 0 {
    if Cmp(val,ON) {
      switch Mode {
      case RUST: { Wr(" = true;\n"); }
      case GO: { Wr(" = true\n"); }
      case MOJO,PYTHON: { Wr(" = True\n"); }
      };
      return ;
    };
    if Cmp(val,OFF) {
      switch Mode {
      case RUST: { Wr(" = false;\n"); }
      case GO: { Wr(" = false\n"); }
      case MOJO,PYTHON: { Wr(" = False\n"); }
      };
      return;
    };
    Wr(" = "); Wr(val);
    if strings.HasPrefix(val, "\"") {  Wr("\""); }
    if Mode == RUST { Wr(";"); };
    Wr("\n");
  } else {
    switch Mode {
    case RUST: { Wr(";\n"); }
    default: { Wr("\n"); }
    };

  }
}   

func OnValue(vtype string) string {
  if Cmp(vtype,STR) { return "\"\"";  }
  if Cmp(vtype,NUM) { return "0";  }
  if Cmp(vtype,REAL) { return "0.0";  }
  if Cmp(vtype,LIGHT) { return "false";  };  
  return "";
}

func OnType(vtype string) string {
  if Cmp(vtype,STR) {
    switch Mode {
    case GO: { return "string"; }
    case RUST: { return "&str"; }
    case MOJO: { return "String"; }
    default: { return ""; }
    };
  };
  if Cmp(vtype,NUM) {
    switch Mode {
    case RUST: { return "i64"; }
    case GO:  { return "int"; }
    case MOJO: { return "Int"; }
    default: { return ""; }
    };
  };
  if Cmp(vtype,REAL) {
    switch Mode {
    case RUST: { return "f64"; }
    case GO:  { return "float64"; }
    case MOJO: { return "Float64"; }
    default: { return ""; }
    };
  };
  if Cmp(vtype,LIGHT) {
    switch Mode {
    case GO,RUST: { return "bool"; }
    case MOJO: { return "Bool"; }
    default:
    }; 
  };
  return "";
}

