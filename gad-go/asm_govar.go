// asm_govar.go

package main

import "fmt"
import "strings"
import "strconv"
//import "unsafe"
import "math"

const (
  DTYPE_UNDEF = iota 
  DTYPE_LIGHT = iota
  DTYPE_NUM = iota
  DTYPE_REAL = iota
  DTYPE_STRING = iota
)

type Var struct {
  xname string; 
  pname string;
  isArray bool; 
  asize int; 
  dtype int;
}

var NVar int = 0
var Vars [256]Var

func ValReal(a string) uint64 {
  var d float64;
  d,_  = strconv.ParseFloat(a,64);
  return math.Float64bits(d);
}

func ValNum(a string) uint64 {
  var d uint64;
  d,_ = strconv.ParseUint(a,10,64);
  return d;
}

func TypeOfLiteral(t string) int {
  if strings.HasPrefix(t,"\"") { return DTYPE_STRING; }
  if strings.Contains("0123456789",t[0:1]) {
    if strings.Contains(t,"e") { return DTYPE_REAL; }
    if strings.Contains(t,"E") { return DTYPE_REAL; }
    if strings.Contains(t,".") { return DTYPE_REAL; }
    return DTYPE_NUM;
  };
  return DTYPE_UNDEF;
}

func AsmTypeOf(t string) int {
  var v *Var = VarGet(t); 
  if(v != nil) { return (*v).dtype; }
  return TypeOfLiteral(t);
}


func VarGet(xn string) *Var {
  var i = NVar; 
  for i > 0 {
    i -= 1;
    if Vars[i].xname == xn { return &(Vars[i]); }
  };
  return nil;
}


func VarNew(xn string, isA bool,asize int, xtype int) *Var {
  var i int = NVar; 
  NVar += 1;
  if NVar>255 { return nil; }
  Vars[i].xname = xn;
  Vars[i].isArray = isA;
  Vars[i].asize = asize;
  Vars[i].dtype = xtype;
  return &(Vars[i]);
}

func VarDump() {
  fmt.Print("\n xref: \n");
  var i = 0;
  for i < NVar {
    fmt.Printf("%s.",Vars[i].pname);
    fmt.Printf("%s :",Vars[i].xname);
    var dt = Vars[i].dtype;
    switch(dt) {
    case DTYPE_UNDEF:  fmt.Printf(" %s\n","Undef");
    case DTYPE_LIGHT:  fmt.Printf(" %s\n","Light");
    case DTYPE_NUM:    fmt.Printf(" %s\n","Num"); 
    case DTYPE_REAL:   fmt.Printf(" %s\n","Real");
    case DTYPE_STRING: fmt.Printf(" %s\n","String");
    };
    i += 1;
  };
  fmt.Printf("\n");
}

func GenAsmFmt(fmt string,varV string) {
  Da("\n"); Da(CurProc); Da("."); Da(varV); Da(".cnv:\n");
  Da("  .asciz \""); Da(fmt); Da("\"\n");
}

func AsmGoVar(varV string,vtype string,val string ) {
  var dt = DTYPE_UNDEF; Da(CurProc); Da("."); Da(varV); Da(":\n"); 
  if len(val) > 0 {
    if Cmp(vtype,STR) {  
       Da("  .asciz "); Da(val); Da("\""); GenAsmFmt("%s",varV); dt = DTYPE_STRING; 
    };
    if Cmp(vtype,NUM) { Da("  .quad "); Da(val); GenAsmFmt("%d",varV); dt = DTYPE_NUM;  };
    if Cmp(vtype,REAL) { Da("  .double "); Da(val); GenAsmFmt("%g",varV); dt = DTYPE_REAL;  };
    if Cmp(vtype,LIGHT) {  Wr(" .quad ");
      var one = "1";
      if Cmp(val,ON) { one = "1"; }
      if Cmp(val,OFF) { one = "0"; }
      Wr(one);
      GenAsmFmt("%d",varV); dt = DTYPE_LIGHT;      
    };
    VarNew(varV,false,0,dt); 
    return;
  };
  switch {  
  case Cmp(vtype,STR): { Da("  .space 256,0\n"); GenAsmFmt("%s",varV); dt = DTYPE_STRING; }  
  case Cmp(vtype,NUM): { Da("  .quad 0\n"); GenAsmFmt("%d",varV); dt = DTYPE_NUM; }  
  case Cmp(vtype,REAL): { Da("  .double 0.0\n"); GenAsmFmt("%g",varV); dt = DTYPE_REAL; }
  case Cmp(vtype,LIGHT): { Da("  .quad 0\n"); GenAsmFmt("%d",varV); dt = DTYPE_LIGHT; }
  };
  VarNew(varV,false,0,dt);
}

