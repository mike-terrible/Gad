// 
// is_.go
//
package main

func AsmGenIs(nv int, p [256]string ) { }

func GenIs(nv int, p [256]string ) {
  switch Mode {
  case ASM: { AsmGenIs(nv,p); return; }
  case GO,RUST: { Wr("{\n"); }
  };
  SetIdent(GetIdent() + 2);
}

