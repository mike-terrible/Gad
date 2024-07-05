//
// asm_getlit.go
//
package main

import (
  "fmt"
//  "strings"
)

var lit int = 0;

func GetLit() int { 
  var v = lit; lit += 1; return v; 
}

func Lit(n int) string {
  return fmt.Sprintf("gad_c%d",n);
}
