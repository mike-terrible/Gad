// 
// show.go
//
package main

import "strings"

func GenShow(nv int, p [256]string )  {
  var i = 0; 
  for { i += 1; if i >= nv { break; }; var t = p[i];
    if Cmp(t,WITH) { 
       i += 1; if i >= nv { break; }; 
       t = p[i]; 
       switch Mode {
       case RUST: {
         To(GetIdent()); Wr("print!(\"{ } \","); Wr(t); 
         if strings.HasPrefix(t, "\"") {  Wr("\"");  };  
         Wr(");\n");
       }
       case GO: {
         To(GetIdent()); Wr("print("); Wr(t); 
         if strings.HasPrefix(t,"\"") { Wr("\""); };
         Wr(",\" \");\n");
       }
       case MOJO, PYTHON: {
         To(GetIdent()); Wr("print("); Wr(t); 
         if strings.HasPrefix(t,"\"") { Wr("\""); }; 
         Wr(",end =\" \")\n");
       }
       default:  
       }; // match
    }; // if
  }; // loop
}


