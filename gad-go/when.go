
// when.go
//
package main

//import "strings"

func GenWhen(nv int,p [256]string) {
  GenEval(nv,p);
  /*
  To( Ident );
  switch Mode {
  case GO:  Wr("for")
  case RUST, MOJO, PYTHON: Wr("while"); 
  default:
  };
  Ident += 2; var i = 0;  
  for {
    i += 1; if i >= nv { break; };
    var t = p[i];  
    var be = Cmp(t,REPEAT);
    if be {
      switch Mode {
      case GO,RUST: { Wr(" {\n"); return;  } 
      case MOJO,PYTHON: { Wr(" :\n"); return; }
      default:
      };
      return;
    };
    Wr(" "); Wr(t); 
    if strings.HasPrefix(t, "\"") { Wr( "\"" ); }
  };
  */
}


