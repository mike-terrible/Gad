//asm32_op2real.go
package main

import "fmt"
import "strings"

/*
var Nfcom = 0

func Fcom() ( zgo string, znext string ) {
  zgo = fmt.Sprintf("fcom%d_ok", Nfcom);
  znext = fmt.Sprintf("fcom%d", Nfcom);
  Nfcom += 1;
  return;
}
*/

func asm32fcom(cc string,qfrom string,qto string) {
  Wr("# asm32fcom ",cc," ",qfrom,",",qto," result: ",Result,"\n");
  Wr("  ffree %st(0)\n");
  Wr("  ffree %st(1)\n");
  Wr("  xor %edx,%edx\n");
  Wr("  fldl ",qfrom,"\n",
     "  fldl ",qto,"\n");
  Wr("  fcom\n");
  Wr("  fstsw %ax\n");
  Wr("  sahf\n");
  Wr("  ffree %st(0)\n");
  Wr("  ffree %st(1)\n");
  //var lb,lbe = Fcom();
  switch cc {
  case " < ":  Wr("  setb %dl\n");
  case " <= ": Wr("  setbe %dl\n");
  case " > ":  Wr("  seta %dl\n");
  case " >= ": Wr("  setae %dl\n");
  case " == ": Wr("  sete %dl\n");
  case " != ": Wr("  setne %dl\n");
  default: {
     Wr("# ill op ",cc,"\n"); 
     return; 
  }};
  Wr("  movb %dl,",Result,"\n");
  Wr("  movl %edx,",Result,"\n");
}


func asm32op2real(cc string,from string,to string) {
  Wr("# asm32op2real ",cc," ",from,",",to," result: ",Result,"\n");
  Wr("  ffree %st(0)\n");
  Wr("  ffree %st(1)\n");
  Wr("  fldl ",from,"\n",
     "  fldl ",to,"\n");
  Wr("  ",cc," %st(1),%st(0)\n");
  Wr("  fstl ",Result,"\n");
  Wr("  ffree %st(0)\n");
  Wr("  ffree %st(1)\n");
}

func Asm32real2(xop string,xto string,xfrom string) {
  var from = ""; var to = "";
  var dt = TypeOfLiteral(xfrom);
  switch dt { 
  case DTYPE_REAL,DTYPE_NUM: {
    var lt = GetLit(); 
    from = Lit(lt);
    Da(from); Da(":\n");
    Da("  .double "); Da(xfrom); Da("\n");
  }
  default: {
    if strings.HasPrefix(xfrom,"gad_") {
      from = xfrom; 
    } else { 
      from = fmt.Sprintf("%s.%s",CurProc,xfrom);
    };
  }};
  //
  dt = TypeOfLiteral(xto);
  if (dt == DTYPE_NUM) || (dt == DTYPE_REAL ) {
    var lt = GetLit(); to = Lit(lt);
    Da(to); Da(":\n");
    Da("  .double "); Da(xto); Da("\n");
  } else {
    if !strings.HasPrefix(xto,"gad_") { to = fmt.Sprintf("%s.%s",CurProc,xto); 
    } else { to = xto; };
  };

  //
  switch xop {
  case " = "," != ", " < ", " <= ", " > ", " >= ": asm32fcom(xop,from,to);
  case " + ": asm32op2real("fadd",from,to);
  case " - ": asm32op2real("fsub",from,to);
  case " * ": asm32op2real("fmul",from,to);
  case " / ": asm32op2real("fdiv",from,to);
  default: Wr("# unsupported operation ",xop,"\n");
  };
}

