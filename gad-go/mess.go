// 
// mess.go
//
package main

import "strings"

func GenMess(nv int, p *Seq)  {
  To(GetIdent());
  var i = 0; i += 1; if i >= nv { return; };
  var t = (*p)[i];
  switch Mode {
  case ASM32: Asm32Mess(t); 
  case ASM: AsmMess(t); 
  case RUST: {
    Wr("println!(");
    if strings.HasPrefix(t,"\"") { Wr(t, "\""); } else {  Wr("\"{ }\",", t ); };
    Wr(");\n"); 
  }
  case GO: {
    Wr("println(", t);
    if strings.HasPrefix(t,"\"") { Wr("\""); };
    Wr(")\n");
  }
  case MOJO,PYTHON: {
    Wr("print(", t);
    if strings.HasPrefix(t, "\"") { Wr("\""); };
    Wr(")\n")
  }};
}


