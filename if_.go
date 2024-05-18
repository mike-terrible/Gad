// if_.go
//
package main

import "strings"

func GenIf(nv int, p [256]string )  {
  var i = 0; 
  To(Ident); Wr("if");
  for {
    i += 1; if i >= nv { break; };
    var t = p[i];
    if Cmp(t, THEN) {
      switch Mode {
      case "-rust","-go": { Wr(" {\n");  Ident += 2; return; }
      case "-mojo","-python": { Wr(" :\n"); Ident += 2; return; }
      default: 
      };
      return;
    };
    Wr(" "); Wr(t); 
    if strings.HasPrefix(t,"\"") { Wr("\""); };    
  }
}

