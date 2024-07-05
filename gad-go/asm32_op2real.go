
//asm32_op2real.go

package main

import "fmt"
import "strings"

func Asm32real2(xop string,xto string,xfrom string) {
  var from = ""; var to = "";
  var dt = TypeOfLiteral(xfrom);
  if dt == DTYPE_REAL {
    var lt = GetLit(); var from = Lit(lt);
    Da(from); Da(":\n");
    Da("  .double "); Da(xfrom); Da("\n");
  } else {
    if !strings.HasPrefix(xfrom,"gad_") { from = fmt.Sprintf("%s.%s",CurProc,xfrom); 
    } else { from = xfrom; }
  };
  Wr("  lea ",from, ",%esi\n");
  //
  dt = TypeOfLiteral(xto);
  if(dt == DTYPE_NUM) {
    var lt = GetLit(); var to = Lit(lt);
    Da(to); Da(":\n");
    Da("  .double "); Da(xto); Da("\n");
    Wr("  lea ", to, ",%edi\n"); 
  } else {
    if !strings.HasPrefix(xto,"gad_") { to = fmt.Sprintf("%s.%s",CurProc,xto); 
    } else { to = xto; };
    Wr("  lea ", to, ",%edi\n");
  };
  //
  switch xop {
  case " + ": {
    Wr(" fldl (%esi)\n"," fldl (%edi)\n");
    Wr(" faddp\n"," fstl ",Result,"\n");
  }
  case  " - ": {
    Wr(" fldl (%esi)\n"," fldl (%edi)\n");
    Wr(" fsubp\n"," fstl ",Result,"\n");
  }
  case " * ": {
    Wr(" fldl (%esi)\n"," fldl (%edi)\n");
    Wr(" fmulp\n"," fstl ",Result,"\n");
  }
  default: Wr("# unsupported operation ",xop,"\n");
  };
}

