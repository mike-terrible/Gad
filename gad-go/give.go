
// give.go
//
package main

import "strings"

func GenGive(nv int, p [256]string )  { 
  var i = 0;
  i += 1; 
  if i>= nv { return; }; 
  var t = p[i];
  To(Ident);
  switch Mode {
  case "-go","-mojo": { Wr("var "); }
  case "-rust": { Wr("let mut "); }
  default:
  };
  Wr(t); // var name
  Wr(" = ");
  i += 1; if i >= nv { return; }; // {from}  
  i += 1; if i >= nv { return; }; 
  t = p[i]; // proc name
  Wr(t); var np = 0; // ( 
  for { 
    i += 1; if i >= nv { break; };
    t = p[i];
    if Cmp(t,WITH) {
      i += 1; if i>= nv { break; }; 
      var tt = p[i];
      np += 1; 
      if np >=1 { Wr(","); };
      Wr(tt); 
      if strings.HasPrefix(tt,"\"") { Wr("\""); };
    };    
  };
  if Mode == "-rust" { Wr(");\n"); } else  { Wr(")\n"); };
}

