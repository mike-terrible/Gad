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

func AsmMess(t string) {
  Wr("\n");
  Wr("# asmMess "); Wr(t); 
  if strings.HasPrefix(t,"\"") { Wr("\""); }
  Wr("\n");
  var b = ""; var bval = "";
  var lt = 0;
  var dt = TypeOfLiteral(t);
  if dt == DTYPE_STRING {
    lt = GetLit();
    b = fmt.Sprintf("gad_c%d", lt);
    Da(b); Da(":\n");
    Da("  .asciz "); Da(t); Da("\"\n");
    Wr("  lea "); Wr(b); Wr("(%rip),%rdi\n");
    Wr("  sub %rax,%rax\n");
    Wr("  call puts\n");
    return;
  };  
  if dt == DTYPE_NUM {
    lt = GetLit();
    b = fmt.Sprintf("gad_cnv%d",lt); 
    Da(b); Da(":\n");
    Da("  .asciz \""); Da("%ld"); 
    Da("\"\n");
    Wr("  lea "); Wr(b); Wr("(%rip),$rdi\n");
    Wr("  movq $"); Wr(t); Wr(",%rsi\n");
    Wr("  sub %rax,%rax\n");
    Wr("  call printf\n");
    Wr("  call gad.nl\n");
    return;
  };  
  if dt == DTYPE_REAL {
    lt = GetLit();
    b = fmt.Sprintf("gad_cnv%d",lt); 
    Da(b);  Da(":\n");
    Da("  .asciz \"%lg\"\n");
    var dd = ValReal(t);
    bval = fmt.Sprintf("$%ld",dd);
    Wr("  movq "); Wr(bval); Wr(",%rax");
    Wr("  movq $rax,%xmm0\n");
    Wr("  movq $1,%rax\n");
    Wr("  lea "); Wr(b); Wr("(%rip),%si\n");
    Wr("  call printf\n");
    Wr("  call gad.nl\n");
    return;
  };
  if strings.HasPrefix(t,"gad_") { return; }
  var v *Var = VarGet(t); if v == nil { return; }
  dt = (*v).dtype;
  if dt == DTYPE_STRING {
    Wr("  xor %rax,%rax\n");
    Wr("  lea "); Wr(CurProc); Wr("."); Wr(t); Wr("(%rip),%rsi\n");
    Wr("  lea "); Wr(CurProc); Wr("."); Wr(t); Wr(".cnv(%rip),%rdi\n");
    Wr("  call printf\n");
    Wr("  call gad.nl\n");
    Wr("# end of asmMess\n");
    return;
  };
  if dt == DTYPE_NUM {
    Wr("  lea "); Wr(CurProc); Wr("."); Wr(t); Wr("(%rip),%rsi\n");
    Wr("  lea "); Wr(CurProc); Wr("."); Wr(t); Wr(".cnv(%rip),%rdi\n");
    Wr("  movq (%rsi),%rsi\n");
    Wr("  xor %rax,%rax\n");
    Wr("  call printf\n");
    Wr("  call gad.nl\n");
    Wr("# end of asmMess\n");
    return;
  };
  if(dt == DTYPE_REAL) {
    Wr("  movq $1,%rax\n");
    Wr("  lea "); Wr(CurProc); Wr("."); Wr(t); Wr(",%rdi\n");
    Wr("  movq (%rdi),%xmm0\n");
    Wr("  lea "); Wr(CurProc); Wr("."); Wr(t); Wr(".cnv(%rip),%rdi\n");
    Wr("  call printf\n");
    Wr("  call gad.nl\n");
    Wr("# end of asmMess\n");
  };
}



