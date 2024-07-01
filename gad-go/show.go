// 
// show.go
//
package main

import "strings"

func GenShow(nv int, p *Seq )  {
  var i = 0; 
  for { i += 1; if i >= nv { break; }; var t = (*p)[i];
    if Cmp(t,WITH) { 
       i += 1; if i >= nv { break; }; 
       t = (*p)[i]; 
       switch Mode {
       case ASM: AsmShow(t);
       case RUST: {
         To(GetIdent()); Wr("print!(\"{ } \",", t); 
         if strings.HasPrefix(t, "\"") {  Wr("\"");  };  
         Wr(");\n");
       }
       case GO: {
         To(GetIdent()); Wr("print(", t ); 
         if strings.HasPrefix(t,"\"") { Wr("\""); };
         Wr(",\" \");\n");
       }
       case MOJO, PYTHON: {
         To(GetIdent()); Wr("print(", t); 
         if strings.HasPrefix(t,"\"") { Wr("\""); }; 
         Wr(",end =\" \")\n");
       }}; // match
    }; // if
  }; // loop
}


