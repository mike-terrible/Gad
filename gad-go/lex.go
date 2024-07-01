
// lex.go

package main

//import "fmt"
import "strings"

func Lexer(pp string) ( ret Seq , nret int  ) {
  var baby string = pp;
  ret = make([]string,0)
  var n = len(baby);
  var i = 0;
  var buf string = "";
  var blank bool = false;
  var quot bool = false;
  for i < n {
    var a = baby[i:i+1];
    if a == "\r" { a = " "; };
    if a == "\n" { a = " "; };
    if a == "\"" {
      if quot { blank = false; quot = false; } else {
        buf = strings.Join( []string { buf, a } , "" )
        quot = true; blank = false; 
      };
      i += 1; continue;
    };
    if a == " " {
      if quot {
        buf = strings.Join( []string { buf, a },"");
        i += 1; continue;
      };
      if blank {
        i += 1; continue;
      };
      if len(buf) > 0 {
        blank = false;
        ret = append(ret,buf)
        nret += 1; 
        buf = "";
      };
      i += 1; 
      continue;
    };
    buf = strings.Join([]string{ buf ,a },"");
    i += 1;
  };
  if len(buf) > 0 { ret  = append (ret, buf); nret += 1; };
  // output 
  i = 0
  var ok bool
  var k string
  var v string
  for i < nret {
    v = ret[i];
    k,ok = Alias[v];
    if(ok) { ret[i] = k };
    print("{", ret[i],"} "); i += 1;
  };
  print("\n")
  return;
}
