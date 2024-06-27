// if_.go
//
package main

//import "strings"

func GenIf(nv int, p [256]string )  {
  //var i = 0;
  GenEval(nv,p);
  /*
  i += 1; if i >= nv { return; }; 
  var t = p[i]; 
  To(Ident); Wr("if "); Wr( t);
  if strings.HasPrefix(t,"\"") { Wr("\""); };
  for { 
    i += 1; if i >= nv { return; };  
    t = p[i]; 
    if Cmp(t,THEN) {
      switch Mode {
      case RUST,GO: { Wr(" {\n");  }
      case MOJO,PYTHON: { Wr(" :\n"); }
      default: 
      }; 
      Ident += 2; return;
    };
    Wr(" "); Wr( t); if strings.HasPrefix(t,"\"") { Wr("\""); };
  };
  */
}


