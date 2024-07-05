
// asm32_op.go

package main

import "fmt"
import "strings"

func asm32Op1(xop string) {
  if xop == " + 1" { Wr("  inc %ebx\n");  };
}

func Asm32Op1(xop string,xn string) {

}

func Asm32SetBit(setv string) {
  Wr("  xor %eax,%eax\n");
  Wr("  cmp %ebx,%ecx\n");
  Wr(setv);
  Wr("  mov %eax,%ecx\n");
}

func Asm32cmd2(cc string) {
  Wr("  movl (%esi),%ebx\n");
  Wr("  movl (%edi),%ecx\n");
  Wr("  ", cc," %ebx,%ecx\n"); 
}
 
func Asm32Op2(xop string, xto string, xfrom string) {
  var dt = AsmTypeOf(xfrom); var dtx = AsmTypeOf(xto);
  if ( dt == DTYPE_REAL ) || ( dtx == DTYPE_REAL ) {
    Asm32real2(xop,xto,xfrom);
    return;
  };
  Wr("# asm32Op2 ",xop," ", xfrom, ",", xto, "\n");
  var from = ""; var to = "";
  dt = TypeOfLiteral(xfrom);
  if dt == DTYPE_NUM {
    var lt = GetLit(); var from = Lit(lt);
    Da(from); Da(":\n");
    Da("  .long "); Da(xfrom); Da("\n");
    Da("  .long 0\n");
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
    Da("  .long "); Da(xto); Da("\n");
    Da("  .long 0\n");
    Wr("  lea ", to, ",%edi\n"); 
  } else {
    if !strings.HasPrefix(xto,"gad_") { to = fmt.Sprintf("%s.%s",CurProc,xto); 
    } else { to = xto; };
    Wr("  lea ", to, ",%edi\n");
  };
  //
  //
  switch xop {
  case " + ": Asm32cmd2("addl");
  case " - ": Asm32cmd2("subl"); 
  case " * ": Asm32cmd2("imul"); 
  case " / ": Asm32cmd2("idiv"); 
  case " == ": AsmSetBit("  sete %al\n");
  case " < ": AsmSetBit("  setb %al\n");
  case " <= ": AsmSetBit("  setbe %al\n");
  case " > ": AsmSetBit("  seta %al\n");
  case " >= ": AsmSetBit("  setae %al\n"); 
  case " != ": AsmSetBit("  setne %al\n"); 
  };
  //
  Wr("  lea ", Result, ",%edi\n", 
   "  movl %ecx,(%edi)\n");
}


