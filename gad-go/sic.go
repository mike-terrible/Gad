
package main
// 
// sic.go
//

import "strings"


func GenSic(nv int, p [256]string )  {
  To(GetIdent()); var i = 0; 
  for { i += 1;
    if i >= nv {
      if Mode == "-rust" { Wr(";"); };
      Wr("\n"); return;
    };
    var t = p[i];
    Wr(t);
    if strings.HasPrefix(t,"\"") { Wr("\""); }; 
    Wr(" ");
  };
}
