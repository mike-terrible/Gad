package main

import (
  "fmt"
  "strings"
)

func HaveReal(xx string) string {
  var lt = GetLit();
  var name = Lit(lt);
  Da(name); Da(": .double "); Da(xx); Da("\n");
  return name;
}

func HaveNum(xx string) string {
  var lt = GetLit();
  var name = Lit(lt);
  Da(name); Da(": .long "); Da(xx); Da("\n");
  return name;
}

func Asm32Ass(xto string ,xfrom string) {
  Wr("# asm32Ass ", xfrom,",", xto, "\n");
  var from string; var to string;
  var dt = TypeOfLiteral(xfrom);
  switch(dt) {
  case DTYPE_REAL: {
    from = HaveReal(xfrom);
    /*
    Wr("# movq ", xfrom, ",%esi\n",
       "  lea ", from, ",%esi\n");
    */
  }
  case DTYPE_NUM: {
    from = HaveNum(xfrom);
    /*
    Wr("# movq ", xfrom, ",%esi\n",
       "  lea ", from, ", %esi\n"); 
    */
  }
  case DTYPE_UNDEF: {
    if ! strings.HasPrefix(xfrom,"gad_") { from = fmt.Sprintf("%s.%s",CurProc,xfrom);
    } else { from = xfrom; };
    //Wr("  lea ", from, ",%esi\n");
  }
  };
  if ! strings.HasPrefix(xto,"gad_") {  to = fmt.Sprintf("%s.%s",CurProc,xto);
  } else { to = xto; };
  Wr("  movl ",from,",%eax\n");
  Wr("  movl %eax,",to,"\n");
  Wr("  movl ",from,"+4,%eax\n");
  Wr("  movl %eax,",to,"+4\n");
  /*
  Wr("  movl ",from,",%eax\n");
  Wr("  movl ",from,"+4,%ebx\n");
  Wr("  movl %eax,",to,"\n");
  Wr("  movl %ebx,",to,"+4\n");
  */
  /*
  Wr("  lea ", from, ",%esi\n");
  Wr("  lea ", to, ",%edi\n");
  Wr("  mov $8,%ecx\n");
  Wr("  rep movsb\n");
  */
}



