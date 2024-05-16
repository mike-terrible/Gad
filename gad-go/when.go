
// when.go
//
package main

import "strings"

func GenWhen(nv int,p [256]string) {
  To( Ident );
  switch Mode {
  case "-go":  Wr("for")
  case "-rust","-mojo","-python": Wr("while"); 
  default:
  };
  Ident += 2; var i = 0;  
  for {
    i += 1; if i >= nv { break; };
    var t = p[i];  
    var be = Cmp(t,REPEAT);
    if be {
      switch Mode {
      case "-go","-rust": { Wr(" {\n"); return;  } 
      case "-mojo","-python": { Wr(" :\n"); return; }
      default:
      };
      return;
    };
    Wr(" "); 
    Wr(t);
    if strings.HasPrefix(t, "\"") { Wr( "\"" ); }
  };
}


