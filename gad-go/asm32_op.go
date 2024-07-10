
// asm32_op.go

package main

import "fmt"
import "strings"

var MathOp bool = false;

func asm32Op1(xop string) {
  if xop == " + 1" { Wr("  inc %ebx\n");  };
}

func Asm32Op1(xop string,xn string) {

}


func Asm32SetBit(setv string) {
  Wr("# Asm32SetBit ",setv,"\n");
  Wr("  xor %eax,%eax\n");
  Wr("  movl (%esi),%ebx\n");
  Wr("  movl (%edi),%ecx\n");
  Wr("  cmp %ebx,%ecx\n");
  Wr(setv);
  Wr("  movl %eax,",Result,"\n");
}


func Asm32cmd2(cc string,from string,to string) {
  MathOp = true;
  Wr("# Asm32cmd2 ",cc," ",from,",",to,",result: ",Result,"\n");
  Wr("  movl ",from,",%ebx\n");
  Wr("  movl ",to,",%ecx\n");
  Wr("  ", cc," %ebx,%ecx\n"); 
  Wr("  movl %ecx,",Result,"\n");
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
  switch dt { 
  case DTYPE_NUM: {
    var lt = GetLit(); 
    from = Lit(lt);
    Da(from); Da(":\n");
    Da("  .long "); Da(xfrom); Da("\n");
    Da("  .long 0\n");
    Wr("# dtype == DTYPE_NUM,from = ",from,"\n");
    Wr("  lea ",from,",%esi\n");
  }
  default: {
    if strings.HasPrefix(xfrom,"gad_") {
      from = xfrom; 
      Wr("# gad internal from = ",from,"\n");
      Wr("  lea ",from,",%esi\n");
    } else { 
      from = fmt.Sprintf("%s.%s",CurProc,xfrom);
      Wr("# from = ",from,"\n");
      Wr("  lea ",from,",%esi\n");  
    };
  }};
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
  MathOp = false;
  //
  switch xop {
  case " + ": Asm32cmd2("addl",from,to);
  case " - ": Asm32cmd2("subl",from,to); 
  case " * ": Asm32cmd2("imul",from,to); 
  case " / ": Asm32cmd2("idivl",from,to); 
  case " == ": Asm32SetBit("  sete %al\n");
  case " < ":  Asm32SetBit("  setb %al\n");
  case " <= ": Asm32SetBit("  setbe %al\n");
  case " > ":  Asm32SetBit("  seta %al\n");
  case " >= ": Asm32SetBit("  setae %al\n"); 
  case " != ": Asm32SetBit("  setne %al\n"); 
  };
  //
}


