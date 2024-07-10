// asm_repeat.go
package main

import "fmt"

func Asm32Repeat() {
  Wr("# Asm32Repeat\n");
  Wr("  cmpl $1,"); Wr(Result); Wr("\n");
  Wr("  jnz "); Wr(fmt.Sprintf("leave_%d\n",Evals[Nev - 1])); Wr("\n");
}

func AsmRepeat() {
  Wr("# AsmRepeat\n");
  Wr("  lea "); Wr(Result); Wr(",%rsi\n");
  Wr("  mov (%rsi),%rax\n");
  Wr("  dec %rax\n");
  Wr("  jnz "); Wr(fmt.Sprintf("leave_%d\n",Evals[Nev - 1])); Wr("\n");
}
