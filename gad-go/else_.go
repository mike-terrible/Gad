
// else_.go
//
package main

func GenElse()  {
  if Mode == ASM32 { Asm32Else(); return; };
  if Mode == ASM { AsmElse(); return; };
  SetIdent( GetIdent() - 2 );
  switch Mode {
  case GO,RUST: { To(GetIdent()); Wr("} else {\n"); } 
  case MOJO,PYTHON: { To(GetIdent()); Wr("else:\n");  }
  };
  SetIdent( GetIdent() + 2 );
}

