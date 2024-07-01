// 
// is_.go
//
package main

func AsmGenIs(nv int, p *Seq ) { }

func GenIs(nv int, p *Seq ) {
  switch Mode {
  case ASM: { AsmGenIs(nv,p); return; }
  case GO,RUST: { Wr("{\n"); }
  };
  SetIdent(GetIdent() + 2);
}

