// 
// is_.go
//
package main

func Asm32Is(nv int, p *Seq ) { }
func AsmIs(nv int, p *Seq ) { }

func GenIs(nv int, p *Seq ) {
  switch Mode {
  case ASM32: Asm32Is(nv,p);
  case ASM: AsmIs(nv,p); 
  case GO,RUST: { Wr("{\n"); SetIdent(GetIdent() + 2); }
  };
}

