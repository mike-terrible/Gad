// asm_op2real.go

package main

import "fmt"
import "strings"

func AsmCmpsd(cmpsdval string,xto string,xfrom string) {
  var from string; var to string;
  Wr("# asmCmpsd "); Wr(cmpsdval); Wr(xfrom); Wr(","); Wr(xto); Wr("\n");
  //var literal bool = false;
  var dt = TypeOfLiteral(xfrom);
  if (dt == DTYPE_NUM) || (dt == DTYPE_REAL) {
    var dl string;
    if dt == DTYPE_NUM { dl = ValNum(xfrom); };
    if dt == DTYPE_REAL { dl = ValReal(xfrom); };
    from = dl
    Wr("# movq ",xfrom,",%rax\n");
    Wr("  movq $",from,",%rax\n");
    Wr("  movq %rax,%xmm8\n"); 
  } else {
    if xfrom != "gad_"  { from = fmt.Sprintf("%s.%s",CurProc,xfrom); 
    } else  { from = xfrom; }; 
    Wr("  lea "); Wr(from); Wr(",%rsi\n"); 
    Wr("  movq (%rsi),%xmm8\n");
  };
  dt = TypeOfLiteral(xto);
  if (dt == DTYPE_NUM) || (dt == DTYPE_REAL) {
    var dl string;
    if dt == DTYPE_NUM { dl = ValReal(xto); }
    if dt == DTYPE_REAL { dl = ValReal(xto); };
    to = fmt.Sprintf("$%s",dl); 
    Wr("# movq "); Wr(xto); Wr(",%rax\n");
    Wr("  movq "); Wr(to); Wr(",%rax\n");
    Wr("  movq %rax,%xmm9\n");
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
  Wr("# asmOp2Real "); Wr(opcode); Wr(" ");  Wr(xfrom); Wr(","); Wr(xto);  Wr("\n");
  var dt = TypeOfLiteral(xfrom);
  if (dt == DTYPE_NUM) || (dt == DTYPE_REAL) {
     var dl string;
     if dt == DTYPE_NUM  { dl = ValNum(xfrom); }
     if dt == DTYPE_REAL { dl = ValReal(xfrom); }
     from = fmt.Sprintf("$%s",dl);
     Wr("# movq "); Wr(xfrom); Wr(",%rax\n"); Wr("  movq "); Wr(from); Wr(",%rax\n");
     Wr("  movq %rax,%xmm8\n"); 
  } else {
    if !strings.HasPrefix(xfrom, "gad_") { from = fmt.Sprintf("%s.%s",CurProc,xfrom); 
    } else  { from = xfrom; }; 
    Wr("  lea "); Wr(from); Wr(",%rsi\n");
    Wr("  movq (%rsi),%xmm8\n");
  };
  dt = TypeOfLiteral(xto);
  if (dt == DTYPE_NUM) || (dt == DTYPE_REAL) {
    var dl string;
    if dt == DTYPE_NUM  { dl = ValReal(xto); }
    if dt == DTYPE_REAL { dl = ValReal(xto); }
    to = fmt.Sprintf("$%s",dl); 
    Wr("# movq "); Wr(xto); Wr(",%rax\n");
    Wr("  movq "); Wr(to); Wr(",%rax\n");
    Wr("  movq %rax,%xmm9\n");
  } else {
    if !strings.HasPrefix(xto,"gad_") { to = fmt.Sprintf("%s.%s",CurProc,xto);
    } else  { to = xto; }; 
    Wr("  lea "); Wr(to); Wr(",%rdi\n");
    Wr("  movq (%rdi),%xmm9\n");         
  };
  Wr("  "); Wr(opcode); Wr(" %xmm8,%xmm9\n");
  Wr("  lea "); Wr(Result); Wr("(%rip),%rsi\n");
  Wr("  movq %xmm9,(%rsi)\n");
}



