// asm_then.go
package main

import "fmt"

func Asm32Then() {
  var cur = Nev - 1
  Wr("  movl "); Wr(Result); Wr(",%eax\n");
  //Wr("  and $1,%eax\n");
  Wr("  neg %eax\n");
  Wr("  jnc "); Wr(" else"); Wr( fmt.Sprintf("%d",Evals[cur]) ); Wr("\n");
  Thens[cur] = true;
  Elses[cur] = false;
}

func AsmThen() {
  var cur = Nev - 1
  Wr("  movq "); Wr(Result); Wr("(%rip),%rax\n");
  //Wr("  and $1,%eax\n");
  Wr("  neg %rax\n");
  Wr("  jnc "); Wr(" else"); Wr( fmt.Sprintf("%d",Evals[cur]) ); Wr("\n");
  Thens[cur] = true;
  Elses[cur] = false;
}
