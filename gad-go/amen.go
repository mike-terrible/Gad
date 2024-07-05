//
// amen.go
//
package main

func EvClear() {
}

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
  switch Mode {
  case ASM32: {
    if Nev > 0 { Asm32Loop(); }; 
    return;
  }
  case ASM: { 
    if Nev > 0 { AsmLoop(); }; 
    return; 
  };
  };
  SetIdent(GetIdent() - 2); To(GetIdent());
  switch Mode {
  case RUST: RustLoop();
  case GO:   GoLoop();
  case MOJO: MojoLoop();
  case PYTHON: PyLoop();
  };
}

func GenDone() {
  switch Mode { 
  case ASM32: { Asm32Done(); return; } 
  case ASM: { AsmDone(); return; }
  };
  GenLoop(); 
}

func Asm32Amen() {
  InProc = false;
  Wr("# amen ",CurProc,"\n");
  Wr("  pop %eax\n","  ret\n");
}

func AsmAmen() {
  InProc = false;
  Wr("# amen ",CurProc,"\n");
  Wr("  pop %rax\n","  ret\n");
}

func GenAmen() {
  switch Mode {
  case ASM32: { Asm32Amen(); return; }
  case ASM: { AsmAmen(); return;  }
  };
  InProc = false; 
  To(GetIdent() - 2);
  switch Mode {
  case RUST, GO: { Wr("}\n"); }
  case MOJO,PYTHON: { Wr("pass\n"); }
  default:
  };
}

