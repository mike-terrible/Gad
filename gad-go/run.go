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
  To(GetIdent());
  switch Mode {
  case RUST: { Wr("fn main() {\n"); }
  case GO: { Wr("func main() {"); }
  case MOJO: { Wr("fn main() :\n"); }
  case PYTHON: { Wr("def main() :\n"); }
  case ASM: {
    Wr("main: push %rax\n");
    Wr(xmain); Wr(": xor %rax,%rax\n");
    CurProc = xmain;
    return;
  }};
  switch Mode { 
  case RUST: { 
    To(GetIdent() + 2);
    Wr("unsafe { "); Wr(xmain); Wr("();  };\n");
    To(GetIdent() - 2); Wr("}\n");
    To( GetIdent() );
    Wr("unsafe fn "); Wr(xmain); Wr("() {\n");
    To(GetIdent() + 2);
  }
  case GO: {
    To(GetIdent() + 2); Wr(xmain); Wr("()");
    To(GetIdent() - 2); Wr("}\n");
    To(GetIdent()); Wr("func "); Wr(xmain); Wr("() {\n");
    To(GetIdent() + 2);
  }
  case MOJO: {
    To(GetIdent() + 2); Wr(xmain); Wr("()\n");
    To(GetIdent() - 2); Wr("}\n");
    To(GetIdent()); Wr("fn "); Wr(xmain); Wr("() :\n");
    To(GetIdent() + 2 );
  }
  case PYTHON: {
    To(GetIdent() + 2); Wr(xmain); Wr("()\n");
    To(GetIdent() - 2); 
    Wr("def "); Wr(xmain); Wr("() :\n");
    To(GetIdent() + 2);
  }};
}


