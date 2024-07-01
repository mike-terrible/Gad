// asm_op2real.go

package main

import "fmt"
import "strings"

func AsmCmpsd(cmpsdval string,xto string,xfrom string) {
  var from string; var to string;
  Wr("# asmCmpsd ", cmpsdval, " ", xfrom, ",", xto, "\n");
  //var literal bool = false;
  var dt = TypeOfLiteral(xfrom);
  if (dt == DTYPE_NUM) || (dt == DTYPE_REAL) {
    if dt == DTYPE_NUM { from = ValNum(xfrom); };
    if dt == DTYPE_REAL { from = ValReal(xfrom); };
    Wr("# movq ",xfrom,",%rax\n",
       "  movq $",from,",%rax\n",
       "  movq %rax,%xmm8\n" ); 
  } else {
    if xfrom != "gad_" { from = fmt.Sprintf("%s.%s",CurProc,xfrom); 
    } else  { from = xfrom; }; 
    Wr("  lea "); Wr(from); Wr(",%rsi\n"); 
    Wr("  movq (%rsi),%xmm8\n");
  };
  dt = TypeOfLiteral(xto);
  if (dt == DTYPE_NUM) || (dt == DTYPE_REAL) {
    if dt == DTYPE_NUM { to = ValReal(xto); }
    if dt == DTYPE_REAL { to = ValReal(xto); };
    Wr("# movq ", xto, ",%rax\n",
       "  movq $", to, ",%rax\n",
       "  movq %rax,%xmm9\n");
  } else {
    if !strings.HasPrefix(xto,"gad_") { to = fmt.Sprintf("%s.%s",CurProc,xto); 
    } else  { to = xto; }; 
    Wr("  lea "); Wr(to); Wr(",%rdi\n");
    Wr("  movq (%rdi),%xmm9\n"); 
  };
  Wr("  "); Wr(cmpsdval); Wr(" %xmm8,%xmm9\n");
  Wr("  lea "); Wr(Result); Wr("(%rip),%rsi\n");
  Wr("  movq %xmm9,%rbx\n");
  Wr("  andq $1,%rbx\n");
  Wr("  movq %rbx,(%rsi)\n");
}

func AsmOp2Real(opcode string, xto string,xfrom string) {
  var from string; var to string;
  Wr("# asmOp2Real ", opcode, " ", xfrom, ",", xto, "\n");
  var dt = TypeOfLiteral(xfrom);
  switch dt {
  case DTYPE_NUM,DTYPE_REAL: {
     var dl string;
     if dt == DTYPE_NUM  { dl = ValNum(xfrom); }
     if dt == DTYPE_REAL { dl = ValReal(xfrom); }
     from = fmt.Sprintf("$%s",dl);
     Wr("# movq ", xfrom, ",%rax\n", "  movq ", from, ",%rax\n",
        "  movq %rax,%xmm8\n"); 
  } 
  default: {
    if !strings.HasPrefix(xfrom, "gad_") { from = fmt.Sprintf("%s.%s",CurProc,xfrom); 
    } else  { from = xfrom; }; 
    Wr("  lea ", from, ",%rsi\n",
       "  movq (%rsi),%xmm8\n");
  }};
  dt = TypeOfLiteral(xto);
  switch dt  { 
  case DTYPE_NUM, DTYPE_REAL: {
    var dl string;
    if dt == DTYPE_NUM  { dl = ValNum(xto); }
    if dt == DTYPE_REAL { dl = ValReal(xto); }
    to = fmt.Sprintf("$%s",dl); 
    Wr("# movq ",xto,",%rax\n",
       "  movq ", to,",%rax\n",
       "  movq %rax,%xmm9\n");
  } 
  default: {
    if !strings.HasPrefix(xto,"gad_") { to = fmt.Sprintf("%s.%s",CurProc,xto);
    } else  { to = xto; }; 
    Wr("  lea ",to,",%rdi\n",
       "  movq (%rdi),%xmm9\n");         
  }};
  Wr("  ",opcode," %xmm8,%xmm9\n",
     "  lea ",Result,"(%rip),%rsi\n",
     "  movq %xmm9,(%rsi)\n");
}



