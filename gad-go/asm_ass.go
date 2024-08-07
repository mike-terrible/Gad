
// asm_ass.go

package main

import (
  "fmt"
  "strings"
)


func AsmAss(xto string ,xfrom string) {
  Wr("# asmAss ", xfrom,",", xto, "\n");
  var from string; var to string;
  var dt = TypeOfLiteral(xfrom);
  switch(dt) {
  case DTYPE_REAL: {
    from = ValReal(xfrom);
    Wr("# movq ", xfrom, ",%rsi\n",
       "  movq $", from, ",%rsi\n");
  }
  case DTYPE_NUM: {
    from = ValNum(xfrom);
    Wr("# movq ", xfrom, ",%rsi\n",
       "  movq $", from, ",%rsi\n"); 
  }
  case DTYPE_UNDEF: {
    if ! strings.HasPrefix(xfrom,"gad_") { from = fmt.Sprintf("%s.%s",CurProc,xfrom);
    } else { from = xfrom; };
    Wr("  lea ", from, "(%rip),%rsi\n",
       "  movq (%rsi),%rsi\n");
  }
  };
  if ! strings.HasPrefix(xto,"gad_") {  to = fmt.Sprintf("%s.%s",CurProc,xto);
  } else { to = xto; };
  Wr("  lea ", to, "(%rip),%rdi\n",
     "  movq %rsi,(%rdi)\n");
}

