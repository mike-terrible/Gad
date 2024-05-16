
// parse.go

package main

import "os"
import "bufio"
import "fmt"
import "strings"

var Mode string = "-go";
var Out *bufio.Writer;
var Ident int = 0;
var InProc = false;

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
  var zname = "./out.go"
  Mode = mode;
  switch mode {
  case "-go": { zname = "./out.go"; }
  case "-mojo": { zname = "./out.mojo"; }
  case "-rust": { zname = "./out.rs"; }
  case "-python": { zname = "./out.py"; }
  default: { }
  }
  fmt.Println("output: ",zname);
  fmt.Println("mode: ",Mode)
  out,err := os.Create(zname)
  Out = bufio.NewWriter(out)
  //
  if Mode == "-go" {
    Wr("package main\n");
    Ident = 2;
  };
  //
  fs := bufio.NewScanner(rf)
  fs.Split(bufio.ScanLines)
  var a string = "";
  for fs.Scan() {
    a = fs.Text()
    fmt.Println(a)
    CheckComment(a)
    if IsComment { 
      GenComment(a); continue; 
    }
    pt,nv := Lexer(a)
    fmt.Println("nv=",nv);
    if nv == 0 { continue; }
    var s = pt[0];
    fmt.Println(s)
    if Cmp(s,LOOP) { GenLoop(); continue; }
    if Cmp(s,DONE) { GenDone(); continue; }
    if Cmp(s,RETURN) { GenReturn(nv,pt); continue; }
    if Cmp(s,WHEN) { GenWhen(nv,pt); continue; }
    if Cmp(s,SIC) { GenSic(nv,pt); continue; }
    if Cmp(s,ELSE) { GenElse(); continue; }
    if Cmp(s,THEN) { GenThen(); continue;  }
    if Cmp(s,IF) { GenIf(nv,pt); continue; }
    if Cmp(s,GIVE) { GenGive(nv,pt); continue; }
    if Cmp(s,JOB) { GenJob(nv,pt); continue; }
    if Cmp(s,SHOW) { GenShow(nv, pt); continue; }
    if Cmp(s,MESS) { GenMess(nv, pt); continue; }
    if Cmp(s,RUN) { GenRun(nv, pt); continue; }
    if Cmp(s,AMEN) { GenAmen(); continue; }
    if Cmp(s,DECLARE) { GenDeclare(nv, pt); continue; }
    if Cmp(s,IS) { GenIs(nv, pt); continue; }
    if Cmp(s,PROC) { GenProc(nv , pt); continue; };
  }
  rf.Close()
  Out.Flush()
  out.Close()
}

