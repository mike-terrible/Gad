// 
// proc.go
//
package main

func GenProc(nv int, p [256]string ) {
  var narg = 0;
  InProc = true;
  var i = 0;
  To(GetIdent());
  switch Mode {
  case GO: Wr("func ");
  case PYTHON: Wr("def ");
  case MOJO:   Wr("fn ");
  case RUST:   Wr("unsafe fn "); 
  };
  i += 1; if i>= nv { return; }; var xn = p[i];
  Wr(xn, "(");
  for { i += 1; if i >= nv { Wr(") "); break; };
    var it_is = p[i];
    switch {
    case Cmp(it_is, RETURN): {
      i += 1; if i>= nv { break; }; var act = p[i];
      var ztype = OnType(act);
      if Mode == PYTHON { Wr(") :\n"); SetIdent(GetIdent() + 2); return; };
      var nz = len(ztype);
      if nz > 0 {
         switch Mode {
         case GO: {
            Wr(") ", ztype, " {\n"); SetIdent( GetIdent() + 2); return;
         }
         case MOJO: {
           Wr(") -> ", ztype, " :\n"); 
           SetIdent( GetIdent() + 2 ); return;
         }
         case RUST:  {
           Wr(") -> ");
           if ztype == "&str" { Wr("String"); } else { Wr(ztype ); };
           Wr(" {\n");
           SetIdent( GetIdent() + 2 ); 
           return; 
         }};  
      }; // nz > 0
    } // RETURN
    case Cmp(it_is, IS): {
      switch Mode {
      case RUST , GO: { Wr(") {\n"); } 
      case MOJO , PYTHON: { Wr(") :\n"); }
      };
      SetIdent(GetIdent() + 2 );
      return;
    }
    case Cmp(it_is, WITH): {
      i += 1; if i >= nv { return; }; var varV = p[i]; narg += 1;
      i += 1; if i >= nv { return; }; 
      if narg >1 { Wr(","); };
      Wr(varV);
      var like = p[i];
      if Cmp(like, AKA) {
        i += 1; if i >= nv { return; }; 
        var xtype = p[i];
        var ztype = OnType(xtype);
        switch Mode {
        case GO: Wr(" ", ztype ); 
        case MOJO,RUST: Wr(" :", ztype); 
        };
      };  // AKA    
    }}  // WITH
  }; // loop
}

