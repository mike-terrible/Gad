//
// asm32_mess.go
//
package main

import (
  "fmt"
  "strings"
)


func Asm32Show(t string) {
  Wr("\n","# asm32Show ",t); 
  if strings.HasPrefix(t,"\"") { Wr("\""); }
  Wr("\n");
  var b = ""; 
  var lt = 0;
  var dt = TypeOfLiteral(t);
  switch dt {
  case DTYPE_STRING: {
    lt = GetLit();
    b = Lit(lt);
    Da(b); Da(": "); Da("  .asciz "); Da(t); Da("\"\n");
    Da(b); Da(".cnv: "); Da(" .asciz \"%s\"\n");
    Wr("  lea ", b, ",%esi\n")
    Wr("  lea ", b, ".cnv,%edi\n")
    Wr("  sub %eax,%eax\n")
    Wr("  mov %esp,%ebp\n")
    Wr("  pushl %esi\n")
    Wr("  pushl %edi\n")
    Wr("  call printf\n")
    Wr("  mov %ebp,%esp\n");
    return;
  }
  case DTYPE_NUM: {
    lt = GetLit();
    b = Lit(lt); 
    Da(b); Da(".cnv:\n"); Da("  .asciz \""); Da("%ld"); Da("\"\n");
    Wr("  lea ", b, ",%edi\n",
       "  mov; $", t, ",%esi\n",
       "  sub %eax,%eax\n",
       "  mov %esp,%ebp\n",
       "  pushl %esi\n",
       "  pushl %edi\n",
       "  call printf\n",
       "  mov %ebp,%esp");
    return;
  }  
  case DTYPE_REAL: {
    lt = GetLit();
    var b1 = Lit(lt);
    var b1f = fmt.Sprintf("%g\n",t);
    Da(b1); Da(": .double "); Da(b1f); 
    Da(b1); Da(".cnv:\n"); Da("  .asciz \"%lg\"\n");
    Wr("  mov %esp,%ebp\n");
    Asm32pushReal(b1);
    Wr("  lea ", b1, ".cnv,%edi\n");
    Wr("  push %edi\n");
    Wr("  call printf\n");
    Wr("  mov %ebp,%esp\n");
    return;
  }};
  if strings.HasPrefix(t,"gad_") { return; }
  var v *Var = VarGet(t); if v == nil { return; }
  dt = (*v).dtype;
  switch dt {
  case DTYPE_STRING: 
    Wr("  xor %eax,%eax\n",
       "  lea ", CurProc, ".",  t,  ",%esi\n",
       "  lea ", CurProc, ".", t,  ".cnv,%edi\n",
       "  mov %esp,%ebp\n",
       "  pushl %esi\n",
       "  pushl %edi\n",
       "  call printf\n",
       "  mov %ebp,%esp\n",
       "# end of asm32Show\n" );
  case DTYPE_NUM: 
    Wr("  lea ", CurProc, ".",  t, ",%esi\n",
       "  lea ", CurProc, ".",  t, ".cnv,%edi\n",
       "  movl (%esi),%esi\n",
       "  mov %esp,%ebp\n",
       "  pushl %esi\n",
       "  pushl %edi\n",
       "  call printf\n",
       "  mov %ebp,%esp\n",
       "# end of asm32Show\n");
  case DTYPE_REAL: {
    var memo = fmt.Sprintf("%s.%s",CurProc,t);
    Wr("  lea ",memo , ",%esi\n",
       "  lea ", memo, ".cnv,%edi\n",
       "  mov %esp,%ebp\n");
    Asm32pushReal(CurProc + "." + t);
    Wr("  pushl %edi\n",
       "  call printf\n" ,
       "  mov %ebp,%esp\n",
       "# end of asmShow\n");
    } 
  };
}

func Asm32Mess(t string) {
  Wr("\n","# asm32Mess ",t); 
  if strings.HasPrefix(t,"\"") { Wr("\""); }
  Asm32Show(t);
  Wr("\n"," call gad.nl\n");
}


