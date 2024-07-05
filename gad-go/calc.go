// calc.go

package main

import "fmt"
import "strings"

var EvId = 0;
var Nev = 0;

var Evals []int = make([]int,255);
var Thens []bool = make([]bool,255);
var Loops []bool = make([]bool,255); 
var Elses []bool = make([]bool,255);

var Op = " + add - sub * mul / div % mod < lt > gt <= le >= ge != ne <- -> = to == eq ";

var St []string = make([]string,255);

var Zj int = 0
var Result string = "";

func Pri(op string) int {
  switch op {
  case "<","lt",">","gt","<=","le",">=","ge","!=","ne","==","eq": return 4;
  case "*","mul", "div","%","mod": return 3;
  case "+","add","-","sub": return 2;
  case "<-","to", "->" : return 1;
  };
  return 0;
}


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
  return strings.Contains(Op," " + t + " ");
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
    AsmAllocResult(); 
    St[top] = Result;
    top += 1; AsmOp1(xop,xn1); 
    return top; 
  }
  var top = nt - 1; 
  var xn1 = St[top];
  AllocResult();
  St[top] = Result; top += 1;
  To(GetIdent()); 
  switch Mode { 
  case GO,MOJO: Wr("var "); 
  case RUST: Wr("let mut ");  
  }
  Wr(Result, " = ", xn1, xop );
  eoi();
  return top;
}

func goAss(nt int) int {
  fmt.Printf(" goAss (%d) %s = %s\n",nt,St[nt-1],St[nt-2]);
  var /*xn2*/ xn1 = St[nt-1];
  var /*xn1*/ xn2 = St[nt-2]; 
  switch Mode { 
  case ASM32: { Asm32Ass(xn2,xn1); return nt-2; };
  case ASM: { AsmAss(xn2,xn1);  return nt-2; }
  };
  To(GetIdent());
  Wr(xn2); Wr(" = "); Wr(xn1);
  eoi();
  return nt-2;
}

func goOp2(xop string,nt int) int {
  switch Mode { 
  case ASM32: {
    var top = nt;
    var /*xn2*/ xn1 = "$0";  if (nt - 1) >= 0 { /*xn2*/ xn1 = St[nt - 1]; top = nt - 1; };
    var /*xn1*/ xn2 = "$0";  if (nt - 2) >= 0 { /*xn1*/ xn2 = St[nt - 2]; top = nt - 2; };
    AsmAllocResult(); 
    St[top] = Result; top += 1; 
    Asm32Op2(xop,xn2,xn1); 
    return top;
  }
  case ASM: {
    var top = nt;
    var /*xn2*/ xn1 = "$0";  if (nt - 1) >= 0 { /*xn2*/ xn1 = St[nt - 1]; top = nt - 1; };
    var /*xn1*/ xn2 = "$0";  if (nt - 2) >= 0 { /*xn1*/ xn2 = St[nt - 2]; top = nt - 2; };
    AsmAllocResult(); 
    St[top] = Result; top += 1; 
    AsmOp2(xop,xn2,xn1); 
    return top; 
  }};
  var top = nt;
  var /*xn2*/ xn1 = "0";  if (nt - 1) >= 0 { /*xn2*/ xn1 = St[nt - 1]; top = nt -1; }; 
  var /*xn1*/ xn2 = "0";  if (nt - 2) >= 0 { /*xn1*/ xn2 = St[nt - 2]; top = nt -2; };
  AllocResult();
  St[top] = Result; top += 1;
  To(GetIdent()); 
  switch Mode { 
  case GO,MOJO: Wr("var "); case RUST: Wr("let mut ");  
  };
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
  case ASM32: Asm32Then();
  case ASM: AsmThen();
  case GO: GoThen();
  case RUST: RustThen();
  case MOJO: MojoThen();
  case PYTHON: PythonThen();
  };
}

func genRepeat() {
  Loops[Nev - 1] = true;
  Elses[Nev - 1] = false;
  switch Mode {
  case ASM32: Asm32Repeat();
  case ASM: AsmRepeat();
  case GO:  GoRepeat();
  case RUST: RustRepeat();
  case MOJO: MojoRepeat();
  case PYTHON: PyRepeat();
  };
}

func InitRepeat() {
  switch Mode {
  case ASM32,ASM: {
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

func FromEvil(varName string, iStart int, nv int, p *Seq) {
  var npp int = 0;
  var pp  Seq;
  var ops Seq;
  var cop int = 0; 
  var i int = iStart;
  for { i += 1; if i>=nv { break; }
    var t string =  (*p)[i];
    switch {
    case isOp(t): {
      var prt = Pri(t);
      for cop > 0 {
        cop -= 1;
        var py = Pri(ops[cop]);
        if py < prt { cop += 1; break; };
        pp = append(pp, ops[cop]); npp += 1; 
      };
      ops = append( ops, t ); cop ++;
    }
    case t == "(": {
      ops = append( ops, t ); cop += 1;
    }
    case t == ")": {
      for cop > 0 {
        cop -= 1
        if ops[cop] == "(" { break; };
        pp = append(pp, ops[cop]); npp++;
      };
    }
    default: {
      pp = append(pp, t); npp += 1;
    }
    };
  }; // for
  if npp>0 { 
    var za strings.Builder;
    var i = 0
    for i < npp {
      za.WriteString(" "); za.WriteString(pp[i]);
      i += 1;
    }
    GenComment(za.String());
    FromCalc(varName,-1,npp,&pp); 
  }
}


/********************************************************/

func Asm32Repeat() {
  var zz = fmt.Sprintf("leave_%d\n",Evals[Nev - 1])
  Wr("  lea ",Result,",%esi\n",
     "  movl (%esi),%eax\n",
     "  dec %eax\n",
     "  jnz ",zz  , "\n");
}

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

func Asm32Then() {
  var cur = Nev - 1
  Wr("  lea "); Wr(Result); Wr(",%eax\n");
  Wr("  neg %eax\n");
  Wr("  jnc "); Wr(" else"); Wr( fmt.Sprintf("%d",Evals[cur]) ); Wr("\n");
  Thens[cur] = true;
  Elses[cur] = false;
}

func AsmThen() {
  var cur = Nev - 1
  Wr("  lea "); Wr(Result); Wr("(%rip),%rax\n");
  Wr("  neg %rax\n");
  Wr("  jnc "); Wr(" else"); Wr( fmt.Sprintf("%d",Evals[cur]) ); Wr("\n");
  Thens[cur] = true;
  Elses[cur] = false;
}

func GoThen() bool {
  To(GetIdent());
  Wr("if ");  Wr(Result); Wr(" == 1 {\n");
  SetIdent(GetIdent() + 2);
  return true;
}

func RustThen() bool { return GoThen(); }

func PythonThen() bool {
  To(GetIdent());
  Wr("if  "); Wr(Result); Wr(" == 1:\n");
  SetIdent(GetIdent() + 2); 
  return true;
}

func MojoThen() bool { return PythonThen(); }

/********************************************************/

func Asm32Done() bool { return AsmDone(); }

func AsmDone() bool {
   var cur = Nev - 1;
   Wr("\n");
   if !Elses[cur]  {
      Wr("else"); Wr(fmt.Sprintf("%d",Evals[Nev - 1]) ); Wr(": nop\n");
   };
   Wr("done"); Wr(fmt.Sprintf("%d",Evals[Nev -1]) ); Wr(": nop\n"); 
   Nev -= 1;
  return true;
}

func Asm32Loop() bool{ return AsmLoop(); }

func AsmLoop() bool {
  var z = fmt.Sprintf("%d",Evals[Nev-1]);
  Wr("\n","# loop ev",z,"\n");
  Wr("  jmp ev"); Wr(z); Wr("\n"); 
  Wr("leave_"); Wr(z); Wr(": nop\n");
  Nev -= 1;
  return true;
}

func Asm32Else() bool { return AsmElse(); }

func AsmElse()  bool { 
  var cur = Nev - 1
  var lb = Evals[cur];
  Elses[cur] = true;
  var z = fmt.Sprintf("%d",lb);
  Wr("\n  jmp done"); Wr(z); Wr("\n"); 
  Wr("else"); Wr(z); Wr(": nop\n");
  return true;
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

func FromCalc(varName string, iStart int,nv int, p *Seq) int {
  DbgTrace("FromCalc");
  var i = iStart;
  var top = 0;
  for { i += 1; if i >= nv { break; } 
    var t = (*p)[i]; 
    if t == "(" {  
      FromEvil(varName, i, nv, p); return 0;
    };
    var is_op = isOp(t);
    if is_op {
      switch(t) { 
      case "<-","to","->","=": top = goAss(top); 
      case "inc","++": top = goOp1(" + 1",top);
      case "<=","le":  top = goOp2(" <= ", top);
      case "<","lt":   top = goOp2(" < ",top); 
      case ">=","ge":  top = goOp2(" >= ",top);
      case ">","gt":   top = goOp2(" > ", top);
      case "!=","ne":  top = goOp2(" != ",top);
      case "==","eq":  top = goOp2(" == ", top);
      case "add","+":  top = goOp2(" + ",top);
      case "sub", "-": top = goOp2(" - ",top); 
      case "mul","*":  top = goOp2(" * ",top);
      case "div","/":  top = goOp2(" / ",top); 
      case "mod","%":  top = goOp2(" % ",top);
      };
    } else {
      St[top] = t; top += 1;
    };
  }; // for
  if varName != "?" { 
    if Mode == ASM { Asm32Ass(varName,St[0]); return 0; };
    if Mode == ASM { AsmAss(varName,St[0]); return 0;  };
    To(GetIdent());
    Wr(varName," = ", St[0] ); 
    eoi();
  };
  DbgTrace(")FromCalc");
  return 0;
}

func GenEval(nv int, p *Seq) {
  DbgTrace("Eval(");
  var pu Seq = make([]string,0) 
  var i int; var j int;
  j = 0; i = 0;
  for {
    i += 1; if i >= nv { break; }
    var cc = (*p)[i];
    if Cmp(cc , REPEAT ) {
      EvId += 1;
      Evals[Nev] = EvId; 
      Loops[Nev] = true;
      Elses[Nev] = false;     
      Nev += 1;
      InitRepeat();
      FromCalc("?",-1, j, &pu);
      GenRepeat();
      DbgTrace(")Eval");
      return;
    };
    if Cmp(cc , THEN ) {
      EvId += 1;
      Evals[Nev] = EvId;
      Loops[Nev] = false;
      Elses[Nev] = false;
      Nev += 1;
      FromCalc("?", -1, j, &pu);
      genThen();
      DbgTrace(")Eval");
      return;
    };
    pu = append(pu, cc); j += 1;
  };
  if j > 0 { FromCalc("?",-1,j,&pu); };
  DbgTrace(")Eval");
}
