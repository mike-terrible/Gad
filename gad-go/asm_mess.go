//
// asm_mess.go
//
package main

import (
  "fmt"
  "strings"
)

var lit int = 0;

func GetLit() int { 
  var v = lit; lit += 1; return v; 
}

func AsmShow(t string) {
  Wr("\n","# asmShow ",t); 
  if strings.HasPrefix(t,"\"") { Wr("\""); }
  Wr("\n");
  var b = ""; var bval = "";
  var lt = 0;
  var dt = TypeOfLiteral(t);
  switch dt {
  case DTYPE_STRING: {
    lt = GetLit();
    b = fmt.Sprintf("gad_c%d", lt);
    Da(b); Da(": "); Da("  .asciz "); Da(t); Da("\"\n");
    Da(b); Da(".cnv: "); Da(" .asciz \"%s\"\n");
    Wr("  lea ", b, "(%rip),%rsi\n",
       "  lea ", b, ".cnv(%rip),%rdi\n",
       "  sub %rax,%rax\n",
       "  call printf\n");
    return;
  }
  case DTYPE_NUM: {
    lt = GetLit();
    b = fmt.Sprintf("gad_cnv%d",lt); 
    Da(b); Da(":\n");
    Da("  .asciz \""); Da("%ld"); 
    Da("\"\n");
    Wr("  lea ", b, "(%rip),$rdi\n",
       "  movq $", t, ",%rsi\n",
       "  sub %rax,%rax\n",
       "  call printf\n");
    return;
  }  
  case DTYPE_REAL: {
    lt = GetLit();
    b = fmt.Sprintf("gad_cnv%d",lt); 
    Da(b);  Da(":\n");
    Da("  .asciz \"%lg\"\n");
    var dd = ValReal(t);
    bval = fmt.Sprintf("$%ld",dd);
    Wr("  movq ",bval,",%rax\n",
       "  movq %rax,%xmm0\n",
       "  movq $1,%rax\n",
       "  lea ", b, "(%rip),%si\n",
       "  call printf\n");
    return;
  }};
  if strings.HasPrefix(t,"gad_") { return; }
  var v *Var = VarGet(t); if v == nil { return; }
  dt = (*v).dtype;
  switch dt {
  case DTYPE_STRING: {
    Wr("  xor %rax,%rax\n",
       "  lea ", CurProc, ".",  t,  "(%rip),%rsi\n",
       "  lea ", CurProc, ".", t,  ".cnv(%rip),%rdi\n",
       "  call printf\n",
       "# end of asmShow\n" )
  }
  case DTYPE_NUM: {
    Wr("  lea ", CurProc, ".",  t, "(%rip),%rsi\n",
       "  lea ", CurProc, ".",  t, ".cnv(%rip),%rdi\n",
       "  movq (%rsi),%rsi\n",
       "  xor %rax,%rax\n",
       "  call printf\n",
       "# end of asmShow\n")
  }
  case DTYPE_REAL: {
    Wr("  movq $1,%rax\n",
       "  lea ",CurProc, ".", t, ",%rdi\n",
       "  movq (%rdi),%xmm0\n",
       "  lea ", CurProc,".", t, ".cnv(%rip),%rdi\n",
       "  call printf\n" ,
       "# end of asmShow\n")
  }};
}

func AsmMess(t string) {
  Wr("\n","# asmMess ",t); 
  if strings.HasPrefix(t,"\"") { Wr("\""); }
  Wr("\n");
  var b = ""; var bval = "";
  var lt = 0;
  var dt = TypeOfLiteral(t);
  switch dt {
  case DTYPE_STRING: {
    lt = GetLit();
    b = fmt.Sprintf("gad_c%d", lt);
    Da(b); Da(":\n");
    Da("  .asciz "); Da(t); Da("\"\n");
    Wr("  lea ", b, "(%rip),%rdi\n",
       "  sub %rax,%rax\n",
       "  call puts\n");
    return;
  }
  case DTYPE_NUM: {
    lt = GetLit();
    b = fmt.Sprintf("gad_cnv%d",lt); 
    Da(b); Da(":\n");
    Da("  .asciz \""); Da("%ld"); 
    Da("\"\n");
    Wr("  lea ", b, "(%rip),$rdi\n",
       "  movq $", t, ",%rsi\n",
       "  sub %rax,%rax\n",
       "  call printf\n",
       "  call gad.nl\n");
    return;
  }  
  case DTYPE_REAL: {
    lt = GetLit();
    b = fmt.Sprintf("gad_cnv%d",lt); 
    Da(b);  Da(":\n");
    Da("  .asciz \"%lg\"\n");
    var dd = ValReal(t);
    bval = fmt.Sprintf("$%ld",dd);
    Wr("  movq ",bval,",%rax\n",
       "  movq %rax,%xmm0\n",
       "  movq $1,%rax\n",
       "  lea ", b, "(%rip),%si\n",
       "  call printf\n",
       "  call gad.nl\n");
    return;
  }};
  if strings.HasPrefix(t,"gad_") { return; }
  var v *Var = VarGet(t); if v == nil { return; }
  dt = (*v).dtype;
  switch dt {
  case DTYPE_STRING: {
    Wr("  xor %rax,%rax\n",
       "  lea ", CurProc, ".",  t,  "(%rip),%rsi\n",
       "  lea ", CurProc, ".", t,  ".cnv(%rip),%rdi\n",
       "  call printf\n",
       "  call gad.nl\n",
       "# end of asmMess\n" )
  }
  case DTYPE_NUM: {
    Wr("  lea ", CurProc, ".",  t, "(%rip),%rsi\n",
       "  lea ", CurProc, ".",  t, ".cnv(%rip),%rdi\n",
       "  movq (%rsi),%rsi\n",
       "  xor %rax,%rax\n",
       "  call printf\n",
       "  call gad.nl\n",
       "# end of asmMess\n")
  }
  case DTYPE_REAL: {
    Wr("  movq $1,%rax\n",
       "  lea ",CurProc, ".", t, ",%rdi\n",
       "  movq (%rdi),%xmm0\n",
       "  lea ", CurProc,".", t, ".cnv(%rip),%rdi\n",
       "  call printf\n" ,
       "  call gad.nl\n" ,
       "# end of asmMess\n")
  }};
}



