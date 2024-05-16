// 
// proc.go
//
package main

func GenProc(nv int, p [256]string ) {
  var narg = 0;
  InProc = true;
  var i = 0;
  To(Ident);
  switch Mode {
  case "-go": { Wr("func "); }
  case "-python": { Wr("def "); }
  case "-mojo": { Wr("fn "); }
  case "-rust": { Wr("unsafe fn "); }
  default:
  };
  i += 1; if i>= nv { return; }; var xn = p[i];
  Wr(xn); Wr("(");
  for { i += 1; if i >= nv { Wr(") "); break; };
    var it_is = p[i];
    if Cmp(it_is, RETURN) {
      i += 1; if i>= nv { break; }; var act = p[i];
      var ztype = OnType(act);
      if Mode == "-python" { Wr(") :\n"); Ident += 2; return; };
      var nz = len(ztype);
      if nz > 0 {
         if Mode == "-go" {
            Wr(") "); Wr(ztype); Wr(" {\n"); Ident += 2; return;
         };
         if Mode == "-mojo" {
           Wr(") -> "); Wr(ztype); Wr(" :\n"); 
           Ident += 2; return;
         };
         if Mode == "-rust"  {
           Wr(") -> ");
           if ztype == "&str" { Wr("String"); } else { Wr(ztype ); };
           Wr(" {\n");
           Ident += 2; 
           return; 
         };  
      }; // nz > 0
    }; // RETURN
    if Cmp(it_is, IS) {
      switch Mode {
      case "-rust" , "-go": { Wr(") {\n"); } 
      case "-mojo" , "-python": { Wr(") :\n"); }
      default:
      };
      Ident += 2; 
      return;
    };
    if Cmp(it_is, WITH) {
      i += 1; if i >= nv { return; }; var varV = p[i]; narg += 1;
      i += 1; if i >= nv { return; }; if narg >1 { Wr(","); };
      Wr(varV);
      var like = p[i];
      if Cmp(like, AKA) {
        i += 1; if i >= nv { return; }; 
        var xtype = p[i];
        var ztype = OnType(xtype);
        if Mode == "-go" { Wr(" "); Wr(ztype ); }; 
        var mojorust = false;
        if Mode == "-mojo" { mojorust = true; };
        if Mode == "-rust" { mojorust = true; };
        if mojorust { Wr(" :"); Wr(ztype); };
      };  // AKA    
    };  // WITH
  }; // loop
}

