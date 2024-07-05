// calc.go

package main

import "fmt"
import "strings"


func asmOp1(xop string) {
  if xop == " + 1" {
    Wr("  inc %rbx\n"); 
  };
}


func AsmOp1(xop string,xn string) {
  if strings.Contains("0123456789",xn[0:0]) {
    var buf = "$" + xn;
    Wr("  movq "); Wr(buf); Wr(",%rbx\n"); 
  } else {
     Wr("  lea ");
     if !strings.HasPrefix(xn,"gad_") { Wr(CurProc,"."); };
     Wr(xn,",%rdi\n",
        "  movq (%rdi),%rbx\n");
  };
  asmOp1(xop);
  Wr("  mov %rbx,(%rdi)\n");
}

func AsmSetBit(setv string) {
  Wr("  xor %rax,%rax\n");
  Wr("  cmp %rsi,%rdi\n");
  Wr(setv);
  Wr("  mov %rax,%rdi\n");
}

func AsmOp2(xop string, xto string, xfrom string) {
  //
  var dt = AsmTypeOf(xfrom); var dtx = AsmTypeOf(xto);
  if (dt == DTYPE_REAL)||(dtx == DTYPE_REAL) {
     switch xop {
     case " + ":  { AsmOp2Real("addsd ",xto,xfrom); }
     case " - ": { AsmOp2Real("subsd ",xto,xfrom); }
     case " * ": { AsmOp2Real("mulsd ",xto,xfrom); }
     case " / ": { AsmOp2Real("divsd ",xto,xfrom); }
     /****************************************************************/
     case " == ": { AsmCmpsd(/* cmpsd $0,*/ "cmpeqsd ", xto, xfrom); }
     case " < ": { AsmCmpsd(/* cmpsd $1, */ "cmpltsd ", xto, xfrom); }
     case " <= ": { AsmCmpsd(/* cmpsd $2,*/ "cmplesd ", xto, xfrom); }
     case " != ": { AsmCmpsd(/* cmpsd $4,*/ "cmpneqsd ", xto, xfrom); }
     case " >= ": { AsmCmpsd(/* cmpsd $5,*/ "cmpnltsd " , xto, xfrom); }
     case " > ": { AsmCmpsd( /* cmpsd $6,*/ "cmpnlesd " , xto, xfrom); }
     };
     return;
  };
  Wr("# asmOp2 ",xop," ", xfrom, ",", xto, "\n");
  var from = ""; var to = "";
  dt = TypeOfLiteral(xfrom);
  if dt == DTYPE_NUM {
    from = fmt.Sprintf("$%s",xfrom);
    Wr("  movq ", from, ",%rsi\n"); 
  } else {
    if !strings.HasPrefix(xfrom,"gad_") { from = fmt.Sprintf("%s.%s",CurProc,xfrom); 
    } else { from = xfrom; }
    Wr("  lea ", from, ",%rsi\n","  movq (%rsi),%rsi\n");
  };
  //
  dt = TypeOfLiteral(xto);
  if(dt == DTYPE_NUM) {
    to = fmt.Sprintf("$%s",xto);
    Wr("  movq ", to, ",%rdi\n"); 
  } else {
    if !strings.HasPrefix(xto,"gad_") { to = fmt.Sprintf("%s.%s",CurProc,xto); 
    } else { to = xto; };
    Wr("  lea ", to, ",%rdi\n", "  movq (%rdi),%rdi\n");
  };
  //
  switch xop {
  case " + ": Wr("  add %rsi,%rdi\n"); 
  case " - ": Wr("  sub %rsi,%rdi\n"); 
  case " * ": Wr("  imul %rsi,%rdi\n"); 
  case " / ": Wr("  idiv %rsi,%rdi\n"); 
  case " == ": AsmSetBit("  sete %al\n");
  case " < ": AsmSetBit("  setb %al\n");
  case " <= ": AsmSetBit("  setbe %al\n");
  case " > ": AsmSetBit("  seta %al\n");
  case " >= ": AsmSetBit("  setae %al\n"); 
  case " != ": AsmSetBit("  setne %al\n"); 
  };
  //
  Wr("  lea ", Result, "(%rip),%rsi\n", "  movq %rdi,(%rsi)\n");
}

