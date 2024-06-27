
// asm_ass.go

package main

import (
  "fmt"
  "strings"
)

func AsmAss(xto string ,xfrom string) {
  Wr("# asmAss "); Wr(xfrom); 
  Wr(","); 
  Wr(xto); Wr("\n");
  var from string; var to string;
  var dt = TypeOfLiteral(xfrom);
  var dl uint64 = 0;
  switch(dt) {
  case DTYPE_REAL: {
    dl = ValReal(xfrom);
    from = fmt.Sprintf("$%ld",dl);
    Wr("# movq "); Wr(xfrom); Wr(",%rsi\n");
    Wr("  movq "); Wr(from); Wr(",%rsi\n");
  }
  case DTYPE_NUM: {
    //dl = ValNum(xfrom);
    //from = fmt.Sprintf("%d",dl);
    from = xfrom;
    Wr("# movq "); Wr(xfrom); Wr(",%rsi\n");
    Wr("  movq $"); Wr(from); Wr(",%rsi\n"); 
  }
  case DTYPE_UNDEF: {
    if ! strings.HasPrefix(xfrom,"gad_") { from = fmt.Sprintf("%s.%s",CurProc,xfrom);
    } else { from = xfrom; };
    Wr("  lea "); Wr(from); Wr("(%rip),%rsi\n");
    Wr("  movq (%rsi),%rsi\n");
  }
  };
  if ! strings.HasPrefix(xto,"gad_") {  to = fmt.Sprintf("%s.%s",CurProc,xto);
  } else { to = xto; };
  Wr("  lea "); Wr(to); Wr("(%rip),%rdi\n");
  Wr("  movq %rsi,(%rdi)\n");
}

