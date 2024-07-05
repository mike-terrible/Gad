package main

import "strings"


func GenReturn( nv int,p *Seq ) {
  var i = 0; 
  //
  i += 1;
  if i == nv { To(GetIdent()); Wr("return");
    if Mode == RUST { Wr(";"); };
    Wr("\n");
    return;
  };
  var t = (*p)[i]; To(GetIdent()); 
  Wr("return ",t); 
  if strings.HasPrefix(t,"\"") {
    Wr("\"");
    if Mode == RUST { Wr(".to_string()"); };
  };  
  if Mode == RUST { Wr(";"); };
  Wr("\n"); 
} 


