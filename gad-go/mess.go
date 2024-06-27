// 
// mess.go
//
package main

import "strings"

func GenMess(nv int, p [256]string)  {
  To(GetIdent());
  var i = 0; 
  i += 1; if i >= nv { return; };
  var t = p[i];
  switch Mode {
  case ASM: {
    AsmMess(t); return;
  }
  case RUST: {
    Wr("println!(");
    if strings.HasPrefix(t,"\"") { 
      Wr(t); Wr("\""); 
    } else { 
      Wr("\"{ }\","); Wr(t ); 
    };
    Wr(");\n"); return;
  }
  case GO: {
    Wr("println("); Wr(t);
    if strings.HasPrefix(t,"\"") { Wr("\""); };
    Wr(")\n");
    return;
  }
  case MOJO,PYTHON: {
    Wr("print("); Wr(t);
    if strings.HasPrefix(t, "\"") { Wr("\""); };
    Wr(")\n")
  }
  default:
  };
}


