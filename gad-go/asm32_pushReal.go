//
// asm32_pushReal.s
//
package main

/*
import (
  "fmt"
  "strings"
)
*/

func Asm32pushReal(result string) {
  Wr("  lea ",result,",%esi\n");
  Wr("  movl (%esi),%eax\n");
  Wr("  add $4,%esi\n");
  Wr("  movl (%esi),%ebx\n");
  Wr("  push %ebx\n");
  Wr("  push %eax\n");
}