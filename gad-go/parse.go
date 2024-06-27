
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
  if strings.HasSuffix(fn, ".гад") {
  } else {
    var a = []string { fn , ".гад" };
    dname = strings.Join(a,"");
  }; 
  fmt.Println("input: ",dname); 
  rf,err := os.Open(dname)
  if err != nil { fmt.Println(err); return; }
  var zname = dname
  Mode = mode;
  switch mode {
  case GO: { zname = strings.ReplaceAll(zname,".гад",".go"); }
  case MOJO: { zname = strings.ReplaceAll(zname,".гад",".mojo"); }
  case RUST: { zname = strings.ReplaceAll(zname,".гад",".rs"); }
  case PYTHON: { zname =strings.ReplaceAll(zname,".гад",".py"); }
  case ASM: { zname =strings.ReplaceAll(zname,".гад",".s"); }
  default: { }
  }
  fmt.Println("output: ",zname);
  fmt.Println("mode: ",Mode)
  out,err := os.Create(zname)
  Out = bufio.NewWriter(out)
  //
  if Mode == ASM {
    Wr("  .text\n"); Wr("  .global main\n");
  };
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
    case Cmp(s,LOOP): { GenLoop(); }
    case Cmp(s,DONE): { GenDone();  }
    case Cmp(s,RETURN): { GenReturn(nv,pt); }
    case Cmp(s,WHEN): { GenWhen(nv,pt); }
    case Cmp(s,SIC): { GenSic(nv,pt);  }
    case Cmp(s,ELSE): { GenElse(); }
    case Cmp(s,THEN): { GenThen(); }
    case Cmp(s,IF): { GenIf(nv,pt); }
    case Cmp(s,GIVE): { GenGive(nv,pt); }
    case Cmp(s,JOB): { GenJob(nv,pt); }
    case Cmp(s,SHOW): { GenShow(nv, pt); }
    case Cmp(s,MESS): { GenMess(nv, pt); }
    case Cmp(s,RUN): { GenRun(nv, pt); }
    case Cmp(s,AMEN): { GenAmen(); }
    case Cmp(s,DECLARE): { GenDeclare(nv, pt);  }
    case Cmp(s,WITH): { GenWith(nv , pt); }
    case Cmp(s,IS): { GenIs(nv, pt);  }
    case Cmp(s,PROC): { GenProc(nv , pt);  }
    case Cmp(s,INIT): { GenInit();  }
    case Cmp(s,ALIAS): { GenAlias(nv, pt); }
    case Cmp(s,EVAL): { GenEval(nv,pt); }
    default: {
      GadError(s, pt, nv); 
    };
    };
  };
  if NeedBoolOf {
    switch Mode {
    case RUST: {
      Wr("\n");
      Wr("fn bool_of(v :bool) -> i64 {\n");
      Wr("  if v { return 1; }\n");
      Wr("  return 0;\n");
      Wr("}\n");
    }
    case GO: {
      Wr("\n");
      Wr("  func BoolOf(v bool) int {\n");
      Wr("    if v { return 1; }\n");
      Wr("    return 0;\n");
      Wr("  }\n");
    }
    case PYTHON: {
      Wr("\n");
      Wr("def bool_of(v) :\n");
      Wr("  if v :\n");
      Wr("    return 1\n");
      Wr("  return 0\n");
    }
    case MOJO: {
      Wr("\n");
      Wr("fn bool_of(v: Bool) -> Int:\n");
      Wr("  if v:\n");
      Wr("    return 1\n");
      Wr("  else:\n");
      Wr("    return 0\n");
      Wr("pass\n");
    }
    };
  };
  if Mode == PYTHON { Wr("main()\n"); };
  if Mode == ASM {
    Wr("gad.nl: xor %rax,%rax\n");
    Wr(" lea gad.nlcnv(%rip),%rdi\n");
    Wr(" call puts\n");
    Wr(" ret\n");
    var z = Ab.String();
    Wr(" .data\n"); 
    Wr(z); 
    Wr("gad.nlcnv: .byte 0,0,0,0,0,0,0,0\n");
    Wr(" .section .note.GNU-stack,\"\",@progbits\n"); 
    Wr(" .end\n");
    VarDump();
  }
  rf.Close()
  Out.Flush()
  out.Close()
}

