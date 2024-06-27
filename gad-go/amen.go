//
// amen.go
//
package main

func GenInit() {
  Wr("\n")
  InArray = false;
}


func RustLoop() {
  To(GetIdent()); Wr("}\n");
}

func GoLoop() {
  To(GetIdent()); Wr("}\n");
}

func GenLoop() {
  DbgTrace("GenLoop");
  if Mode == ASM { AsmLoop(); return; };
  SetIdent(GetIdent() - 2); To(GetIdent());
  switch Mode {
  case RUST: { RustLoop(); }
  case GO: { GoLoop(); }
  case MOJO: { MojoLoop(); }
  case PYTHON: { PyLoop(); }
  }
}

func GenDone() { 
  if Mode == ASM { AsmDone(); return; }
  GenLoop(); 
}

func AsmAmen() {
  InProc = false;
  Wr("  pop %rax\n"); Wr("  ret\n");
}

func GenAmen() {
  if Mode == ASM { AsmAmen(); return; }
  InProc = false; 
  To(GetIdent() - 2);
  switch Mode {
  case RUST, GO: { Wr("}\n"); }
  case MOJO,PYTHON: { Wr("pass\n"); }
  default:
  };
}

