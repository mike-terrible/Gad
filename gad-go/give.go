
// give.go
//
package main

import "strings"


func mojoGiveArray(nv int, p *Seq ) {
  var i = 1; i += 1; if i >= nv { return; };
  var t = (*p)[i]
  To(GetIdent());
  Wr(t, ".__setitem__("); 
  i += 1; if i >= nv { return; };
  t = (*p)[i]; 
  if !Cmp(t,WITH)  { return; };
  i += 1; if i >= nv { return; };
  t = (*p)[i]
  Wr(t,",");
  i += 1; if i >= nv { return; };
  i += 1; if i >= nv { return; };
  t = (*p)[i];
  Wr(t);
  if strings.HasPrefix(t,"\"") { Wr("\""); }; 
  Wr(")\n");
}

func giveArray(nv int, p *Seq ) {
  var i = 1; i += 1; if i >= nv { return; };
  var t = (*p)[i]
  To(GetIdent()); 
  Wr(t);
  i += 1; if i >= nv { return; };
  t = (*p)[i]; 
  if !Cmp(t,WITH)  { return; };
  i += 1; if i >= nv { return; };
  t = (*p)[i]
  i += 1; if i >= nv { return; };
  i += 1; if i >= nv { return; };
  var b = (*p)[i]; Wr("[", t, "] =",b);
  if strings.HasPrefix(b,"\"") { Wr("\""); };
  switch Mode { 
  case RUST: Wr(";\n"); 
  default: Wr("\n"); 
  }; 
}


func GenGive(nv int, p *Seq )  { 
  var i = 0;
  i += 1; 
  if i>= nv { return; }; 
  var t = (*p)[i];
  if Cmp(t,ARRAY) { 
    if Mode == MOJO { mojoGiveArray(nv,p); return; };
    giveArray(nv,p); return; 
  };
  var varName = t;
  i += 1; if i >= nv { return; }; // {from}  
  t = (*p)[i]; // from | array
  if Cmp(t,EVAL) { FromCalc(varName, i,nv,p); return; };
  To(GetIdent());
  switch Mode { case GO,MOJO: Wr("var "); case RUST: Wr("let mut "); };
  Wr(varName," = ");
  if Cmp(t,ARRAY) { 
    if Mode == MOJO {
      i += 1; if i >= nv { return; }
      var aname = (*p)[i];
      i += 1; i += 1; t = (*p)[i];
      Wr(aname,".__getitem__(",t,")\n");
      return;
    };
    i += 1; if i >= nv { return; }
    var aname = (*p)[i];
    i += 1; i += 1; t = (*p)[i];
    Wr(aname,"[",t,"]");
    if Mode == RUST { Wr(";"); };
    Wr("\n");
    return; 
  } // ARRAY
  //
  i += 1; if i >= nv { return; }; 
  t = (*p)[i]; // proc name
  Wr(t); var np = 0; // ( 
  for { 
    i += 1; if i >= nv { break; };
    t = (*p)[i];
    if Cmp(t,WITH) {
      i += 1; if i>= nv { break; }; 
      var tt = (*p)[i];
      np += 1; 
      if np >=1 { Wr(","); };
      Wr(tt); 
      if strings.HasPrefix(tt,"\"") { Wr("\""); };
    };    
  };
  if Mode == RUST { Wr(");\n"); } else  { Wr(")\n"); };
}

