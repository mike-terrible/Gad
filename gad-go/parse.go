
// parse.go

package main

import "os"
import "bufio"
import "fmt"
import "strings"

var Mode string = GO;
var Out *bufio.Writer;
var ident int = 0;
var InProc = false;
var InArray = false;
var InInit int = 0;
var CurVar string = "";
var CurProc string = "modmain";

func GetIdent() int {
  return ident;
}

func SetIdent(v int) {
  ident = v;
}

func Parser(fn string, mode string) {
  var dname = fn;
  if ! strings.HasSuffix(fn, ".гад") {
    dname = fmt.Sprintf("%s.гад",fn)
  }; 
  fmt.Println("input: ",dname); 
  rf,err := os.Open(dname)
  if err != nil { fmt.Println(err); return; }
  var zname = strings.ReplaceAll(dname,".гад",".out");
  Mode = mode;
  switch Mode {
  case GO: zname = strings.ReplaceAll(zname,".out",".go"); 
  case MOJO: zname = strings.ReplaceAll(zname,".out",".mojo"); 
  case RUST: zname = strings.ReplaceAll(zname,".out",".rs"); 
  case PYTHON: zname =strings.ReplaceAll(zname,".out",".py");
  case ASM: zname =strings.ReplaceAll(zname,".out",".s"); 
  case ASM32: zname =strings.ReplaceAll(zname,".out","-32.s");
  };
  fmt.Println("output: ",zname);
  fmt.Println("mode: ",Mode)
  out,err := os.Create(zname)
  Out = bufio.NewWriter(out)
  //
  switch Mode { 
  case ASM,ASM32: {
    Wr("  .text\n"); Wr("  .global main\n");
  }};
  if Mode == GO {
    Wr("package main\n");
    SetIdent(2);
  };
  //
  fs := bufio.NewScanner(rf)
  fs.Split(bufio.ScanLines)
  var a string = "";
  for fs.Scan() {
    a = fs.Text()
    fmt.Print(a,"\n")
    CheckComment(a)
    if IsComment { GenComment(a); continue; }
    if !IsLine { continue; }
    pt,nv := Lexer(a)
    if nv == 0 { continue; }
    var s = pt[0];
    switch {
    case Cmp(s,LOOP): GenLoop();
    case Cmp(s,DONE): GenDone(); 
    case Cmp(s,RETURN): GenReturn(nv,&pt);
    case Cmp(s,WHEN): GenEval(nv,&pt);
    case Cmp(s,SIC): GenSic(nv,&pt);  
    case Cmp(s,ELSE): GenElse();
    case Cmp(s,IF): GenEval(nv,&pt);
    case Cmp(s,GIVE): GenGive(nv,&pt);
    case Cmp(s,JOB): GenJob(nv,&pt); 
    case Cmp(s,SHOW): GenShow(nv,&pt);
    case Cmp(s,MESS): GenMess(nv, &pt);
    case Cmp(s,RUN): GenRun(nv, &pt);
    case Cmp(s,AMEN): GenAmen();
    case Cmp(s,DECLARE): GenDeclare(nv, &pt); 
    case Cmp(s,WITH): GenWith(nv , &pt);
    case Cmp(s,IS): GenIs(nv, &pt); 
    case Cmp(s,PROC): GenProc(nv , &pt); 
    case Cmp(s,INIT): GenInit(); 
    case Cmp(s,ALIAS): GenAlias(nv, &pt);
    case Cmp(s,EVAL): GenEval(nv,&pt);
    default: GadError(s, &pt, nv); 
    };
  };
  if NeedBoolOf {
    switch Mode {
    case RUST: {
      Wr("\n",
         "fn bool_of(v :bool) -> i64 {\n",
         "  if v { return 1; }\n",
         "  return 0;\n",
         "}\n");
    }
    case GO: {
      Wr("\n",
         "  func BoolOf(v bool) int {\n",
         "    if v { return 1; }\n",
         "    return 0;\n",
         "  }\n");
    }
    case PYTHON: {
      Wr("\n",
         "def bool_of(v) :\n",
         "  if v :\n",
         "    return 1\n",
         "  return 0\n" );
    }
    case MOJO: {
      Wr("\n",
         "fn bool_of(v: Bool) -> Int:\n",
         "  if v:\n",
         "    return 1\n",
         "  else:\n",
         "    return 0\n",
         "pass\n");
    }};
  };
  if Mode == PYTHON { Wr("main()\n"); };
  if Mode == ASM32 {
    Wr("gad.nl: mov %esp,%ebp\n",
       "  lea gad.nlcnv,%edi\n",
       "  push %edi\n",
       "  call printf\n",
       "  mov %ebp,%esp\n",
       "  ret\n");
     //Da("gad_true:  .byte 0,0,0,0,0,0,0,1\n");
     //Da("gad_false: .byte 0,0,0,0,0,0,0,0\n");
     var z = Ab.String();
     Wr(" .data\n",
       strings.ReplaceAll(z,"\n\n","\n"), 
       "gad.nlcnv: .byte 10,0\n",
       " .section .note.GNU-stack,\"\",@progbits\n", 
       " .end\n");
     VarDump();
  };
  if Mode == ASM {
    if NeedNL {
       Wr("gad.nl: xor %rax,%rax\n",
       " lea gad.nlcnv(%rip),%rdi\n",
       " call puts\n",
       " ret\n");
     };
     var z = Ab.String();
     Wr(" .data\n",
       strings.ReplaceAll(z,"\n\n","\n"), 
       "gad.nlcnv: .byte 0,0,0,0,0,0,0,0\n",
       " .section .note.GNU-stack,\"\",@progbits\n", 
       " .end\n");
     VarDump();
  };
  rf.Close()
  Out.Flush()
  out.Close()
}

