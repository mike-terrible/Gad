package main

import "strings"


func GenReturn( nv int,p [256]string ) {
  var i = 0; 
  //
  i += 1;
  if i == nv { To(GetIdent()); Wr("return");
    if Mode == "-rust" { Wr(";"); };
    Wr("\n");
    return;
  };
  var t = p[i]; To(GetIdent()); Wr("return "); 
  Wr(t); 
  if strings.HasPrefix(t,"\"") {
    Wr("\"");
    if Mode == "-rust" { Wr(".to_string()"); };
  };  
  if Mode == "-rust" { Wr(";"); };
  Wr("\n"); 
} 


