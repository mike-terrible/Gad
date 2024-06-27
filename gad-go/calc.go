// calc.go

package main

import "fmt"
import "strings"

var EvId = 0;
var Nev = 0;

var Evals [256]int;
var Thens [256]bool;
var Loops [256]bool;
var Elses [256]bool;

var Op = " + add - sub * mul / div % mod < lt > gt <= le >= ge != ne <- to == eq ";

var St [256]string;

var Zj int = 0
var Result string = "";

func AsmAllocResult() {
  Result = fmt.Sprintf("gad_%d",Zj);
  Da(Result); Da(":\n"); Da("  .quad 0\n");
  Zj += 1;
}

func AllocResult() string {
  Result = fmt.Sprintf("gad_%d",Zj);
  Zj += 1;
  return Result;
}

func isOp(t string) bool {
  if strings.Contains(Op," " + t + " ") { return true; }
  return false;
}

func eoi() {
  switch Mode {
  case GO,RUST : { Wr(";\n"); }
  case MOJO,PYTHON: { Wr("\n"); }
  }
}

func goOp1(xop string,nt int) int { 
  if Mode == ASM { 
    var top = nt - 1;
    var xn1 = St[top];
    AsmAllocResult(); St[top] = Result; top += 1; AsmOp1(xop,xn1); 
    return top; 
  }
  var top = nt - 1;
  var xn1 = St[top];
  AllocResult();
  St[top] = Result; top += 1;
  To(GetIdent()); 
  switch Mode { 
  case GO,MOJO: { Wr("var "); } 
  case RUST: { Wr("let mut "); }  
  }
  Wr(Result); Wr(" = "); Wr(xn1); Wr(xop);
  eoi();
  return top;
}

func goAss(nt int) int {
  fmt.Printf(" goAss (%d) %s = %s\n",nt,St[nt-1],St[nt-2]);
  var xn2 = St[nt-1];
  var xn1 = St[nt-2]; 
  if Mode == ASM { AsmAss(xn2,xn1);  return nt-2; }
  To(GetIdent());
  Wr(St[nt-1]); Wr(" = "); Wr(St[nt-2]);
  eoi();
  return nt-2;
}

func goOp2(xop string,nt int) int { 
  if Mode == ASM {
    var top = nt;
    var xn2 = "$0";  if (nt - 1) >= 0 { xn2 = St[nt - 1]; top = nt - 1; };
    var xn1 = "$0";  if (nt - 2) >= 0 { xn1 = St[nt - 2]; top = nt - 2; };
    AsmAllocResult(); 
    St[top] = Result; top += 1;
    AsmOp2(xop,xn2,xn1); 
    return top; 
  };
  var top = nt;
  var xn2 = "0";  if (nt - 1) >= 0 { xn2 = St[nt - 1]; top = nt -1; }; 
  var xn1 = "0";  if (nt - 2) >= 0 { xn1 = St[nt - 2]; top = nt -2; };
  AllocResult();
  St[top] = Result; top += 1;
  To(GetIdent()); 
  switch Mode { 
  case GO,MOJO: { Wr("var "); } 
  case RUST: { Wr("let mut "); }  
  }
  if strings.Contains(" == ; != ; > ; < ; >= ; <= ",xop) {
    NeedBoolOf = true;
    Wr(Result); Wr(" = ");
    if Mode == GO {
       Wr("BoolOf("); Wr(xn2); Wr(xop); Wr(xn1); Wr(")");
    };
    if (Mode == RUST) || (Mode == PYTHON) || (Mode == MOJO) {
       Wr("bool_of("); Wr(xn2); Wr(xop); Wr(xn1); Wr(")");
    };
  } else {
    Wr(Result); Wr(" = "); Wr(xn2); Wr(xop); Wr(xn1);
  };   
  eoi();
  return top
}

func genThen() {
  switch Mode {
  case ASM: { AsmThen(); }
  case GO: { GoThen(); }
  case RUST: { RustThen(); }
  case MOJO: { MojoThen(); }
  case PYTHON: { PythonThen(); }
  };
}

func genRepeat() {
  Loops[Nev - 1] = true;
  Elses[Nev - 1] = false;
  switch Mode {
  case ASM: { AsmRepeat(); }
  case GO: { GoRepeat(); }
  case RUST: { RustRepeat(); }
  case MOJO: { MojoRepeat(); }
  case PYTHON: { PyRepeat(); }
  };
}

func InitRepeat() {
  switch Mode {
  case ASM: {
    Wr("\n");
    var zp = fmt.Sprintf("ev%d:\n",Evals[ Nev - 1]);
    Wr(zp);   
  } 
  case RUST: {
    To(GetIdent());
    Wr("loop {");
    To(GetIdent() + 2); 
  }
  case GO: {
    To(GetIdent());
    Wr("for {");
    To(GetIdent() + 2); 
  }
  case PYTHON,MOJO: {
    To(GetIdent());
    Wr("while True:"); 
    To(GetIdent() + 2); 
  }}
}

func FromEvil(varName string, iStart int, nv int, p [256]string) {
  var npp int = 0;
  var pp  [256]string;
  var ops [256]string;
  var aa  [256]string;
  var cop int = 0; 
  var carg int = 0;
  var i int = iStart;
  for { i += 1; if i>=nv { break; }
    var t string =  p[i];
    if t == ")" {
      if cop>0 { cop -= 1;
        var xop = ops[cop];
        if isOp(xop) {
          var a1 = ""; var a2 = "";
          if carg > 0 {
            carg -= 1;  a1 = aa[carg]; 
            if a1 != "(" { pp[npp] = a1; npp += 1; };
          };
          if carg > 0 {
            carg -= 1; a2 = aa[carg]; 
            if(a2 != "(") { pp[npp] = a2; npp += 1; };
          };
          pp[npp] = xop; npp += 1;
        };
      };
      continue;
    };
    if isOp(t) { ops[cop] = t; cop += 1; } else {
      aa[carg] = t; carg += 1;
    };
  }; // for
  if npp>0 { Calc(varName,-1,npp,pp); }
}


func Calc(varName string, start int, nv int, p [256]string ) {
  FromCalc(varName, start, nv, p);
  return;
}

/********************************************************/

func AsmRepeat() {
  Wr("  lea "); Wr(Result); Wr(",%rsi\n");
  Wr("  mov (%rsi),%rax\n");
  Wr("  dec %rax\n");
  Wr("  jnz "); Wr(fmt.Sprintf("leave_%d\n",Evals[Nev - 1])); Wr("\n");
}

func GoRepeat() {
  To(GetIdent());
  Wr("if "); Wr(Result); Wr(" != 1 { break; }\n");
  To(GetIdent()); 
}

func RustRepeat() { 
  To(GetIdent());
  Wr("if "); Wr(Result); Wr(" != 1 { break; }\n");
  To(GetIdent()); 
}

/********************************************************/

func AsmThen() {
  var cur = Nev - 1
  Wr("  lea "); Wr(Result); Wr("(%rip),%rax\n");
  Wr("  dec %rax\n");
  Wr("  jnz "); Wr(" else"); Wr( fmt.Sprintf("%d",Evals[cur]) ); Wr("\n");
  Thens[cur] = true;
  Elses[cur] = false;
}

func GoThen() {
  To(GetIdent());
  Wr("if ");  Wr(Result); Wr(" == 1 {\n");
  SetIdent(GetIdent() + 2);
}

func RustThen() { GoThen(); }

func PythonThen() {
  To(GetIdent());
  Wr("if  "); Wr(Result); Wr(" == 1:\n");
  SetIdent(GetIdent() + 2); 
}

func MojoThen() { PythonThen(); }

/********************************************************/

func AsmDone() {
  var cur = Nev - 1
  Wr("\n");
  if !Elses[cur]  {
    Wr("else"); Wr(fmt.Sprintf("%d",Evals[Nev - 1]) ); Wr(": nop\n");
  };
  Wr("done"); Wr(fmt.Sprintf("%d",Evals[Nev -1]) ); Wr(": nop\n"); 
  Nev -= 1;
}

func AsmLoop() {
  Wr("\n");
  var z = fmt.Sprintf("%d",Evals[Nev-1]);
  Wr("  jmp ev"); Wr(z); Wr("\n"); 
  Wr("leave_"); Wr(z); Wr(": nop\n");
  Nev -= 1;
}

func AsmElse() { 
  var cur = Nev - 1
  var lb = Evals[cur];
  Elses[cur] = true;
  var z = fmt.Sprintf("%d",lb);
  Wr("\n  jmp done"); Wr(z); Wr("\n"); 
  Wr("else"); Wr(z); Wr(": nop\n");
}

func GoElse() {
  To(GetIdent() - 2); 
  Wr("} else {\n"); To(GetIdent() + 2);  
}

func GenRepeat() {
  switch Mode {
  case ASM: { AsmRepeat(); }
  case GO: { GoRepeat();  }
  case RUST: { RustRepeat(); }
  case PYTHON: { PyRepeat(); }
  case MOJO: { MojoRepeat(); }
  }
}

func RustElse() { GoElse(); }

func PythonElse() { MojoElse(); } 

func MojoElse() {
  To(GetIdent() -2);
  Wr("else:\n"); To(GetIdent() + 2); 
}

/********************************************************/

func FromCalc(varName string, iStart int,nv int, p [256]string) int {
  DbgTrace("FromCalc");
  var i = iStart;
  var top = 0;
  for { i += 1; if i >= nv { break; } 
    var t = p[i]; 
    if t == "(" {  
      FromEvil(varName, i, nv, p); return 0;
    };
    var is_op = isOp(t);
    if is_op {
      switch(t) { 
      case "<-","to": { top = goAss(top); }
      case "inc","++": { top = goOp1(" + 1",top); }
      case "<=","le": { top = goOp2(" <= ", top); }
      case "<","lt": { top = goOp2(" < ",top);  }
      case ">=","ge": { top = goOp2(" >= ",top); }
      case ">","gt": { top = goOp2(" > ", top); }
      case "!=","ne": { top = goOp2(" != ",top); }
      case "==","eq": { top = goOp2(" == ", top); }
      case "add","+": { top = goOp2(" + ",top); }
      case "sub", "-": { top = goOp2(" - ",top); } 
      case "mul","*": { top = goOp2(" * ",top); }
      case "div","/": { top = goOp2(" / ",top); } 
      case "mod","%": { top = goOp2(" % ",top); }
      };
    } else {
      St[top] = t; top += 1;
    };
  }; // for
  if varName != "?" { 
    if Mode == ASM { AsmAss(varName,St[0]); return 0;  };
    To(GetIdent());
    Wr(varName); Wr(" = "); Wr( St[0] ); 
    eoi();
  };
  DbgTrace(")FromCalc");
  return 0;
}

func GenEval(nv int, p [256]string) {
  DbgTrace("Eval(");
  var pu [256]string; var i int; var j int;
  j = 0; i = 0;
  for {
    i += 1; if i >= nv { break; }
    var cc = p[i];
    if Cmp(cc , REPEAT ) {
      EvId += 1;
      Evals[Nev] = EvId; 
      Loops[Nev] = true;
      Elses[Nev] = false;     
      Nev += 1;
      InitRepeat();
      FromCalc("?",-1, j, pu);
      GenRepeat();
      DbgTrace(")Eval");
      return;
    };
    if Cmp(cc , THEN ) {
      EvId += 1;
      Evals[Nev] = EvId; Nev += 1;
      FromCalc("?", -1, j, pu);
      genThen();
      DbgTrace(")Eval");
      return;
    };
    pu[j] = cc; j += 1;
  };
  if j > 0 { FromCalc("?",-1,j,pu); };
  DbgTrace(")Eval");
}
