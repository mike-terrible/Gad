// 
// job.rs
//
package main

import "strings"

func GenJob(nv int, p [256]string ) {
  var i = 0;
  i += 1; if i >= nv { return; };
  var t = p[i];
  Wr("\n"); To(Ident); Wr(t); Wr("(");
  var np = 0; 
  for { i += 1; if i >= nv { break; }; t = p[i];
    if Cmp(t,WITH) { i += 1; if i >= nv { break; }; 
      t = p[i]; np += 1; if np > 1 { Wr(","); };
      Wr(t);
      if strings.HasPrefix(t,"\"") { Wr("\""); };
    }; 
  };
  if Mode == "-rust" {
    Wr(");\n");
  } else {
    Wr(")\n");
  };
}


