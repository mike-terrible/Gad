// 
// mess.rs
//
package main

import "strings"

func GenMess(nv int, p [256]string)  {
  To(Ident);
  var i = 0; 
  i += 1; if i >= nv { return; };
  var t = p[i];
  switch Mode {
  case "-rust": {
    Wr("println!(");
    if strings.HasPrefix(t,"\"") { 
      Wr(t); Wr("\""); 
    } else { 
      Wr("\"{ }\","); Wr(t ); 
    };
    Wr(");\n"); return ;
  }
  case "-go": {
    Wr("println("); Wr(t);
    if strings.HasPrefix(t,"\"") { Wr("\""); };
    Wr(")\n");
    return;
  }
  case "-mojo","python": {
    Wr("print("); Wr(t);
    if strings.HasPrefix(t, "\"") { Wr("\""); };
    Wr(")\n")
  }
  default:
  };
}


