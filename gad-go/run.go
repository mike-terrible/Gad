// 
// run.go
//
package main

func GenRun(nv int, p [256]string )  {
  var i = 0; 
  InProc = true;
  i += 1; if i >= nv { return; };
  var xmain = p[i]; 
  i += 1; if i >= nv { return; };
  To(Ident);
  switch Mode {
  case "-rust": { Wr("fn main() {\n"); }
  case "-go": { Wr("func main() {\n"); }
  case "-mojo": { Wr("fn main() :\n"); }
  case "-python": { Wr("def main() :\n"); }
  default:
  };
  Ident += 2; To(Ident);
  if Mode == "-rust" { Wr("unsafe { "); };
  Wr(xmain); Wr("()");
  if Mode == "-rust" { Wr("; }"); };
  Wr("\n");
  Ident -= 2; To(Ident);
  switch Mode {
  case "-go","-rust": { Wr("}\n"); }
  default:
  };
  To(Ident);
  switch Mode {
  case "-rust": { Wr("unsafe fn "); Wr(xmain); Wr("() {\n");  } 
  case "-go": { Wr("func "); Wr(xmain); Wr("() {\n"); } 
  case "-mojo": { Wr("fn "); Wr(xmain); Wr("() :\n"); } 
  case "-python": { Wr("def "); Wr(xmain); Wr("() :\n"); }
  default:
  }; 
  //
  Ident += 2; 
}


