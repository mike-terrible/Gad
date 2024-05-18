
// lex.go

package main

import "fmt"
import "strings"

func Lexer(pp string) ( ret [256]string , nret int  ) {
  var baby string = pp;
  ret[0]=""
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
      if quot {
        blank = false; quot = false;
      } else {
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
        ret[nret] = buf; nret += 1; buf = "";
      };
      i += 1; 
      continue;
    };
    buf = strings.Join([]string{ buf ,a },"");
    i += 1;
  };
  if len(buf) > 0 {
    ret[nret] = buf;
    nret += 1;
  };
  // output
  i = 0;
  for i < nret {
    var be = ret[i]; 
    fmt.Print("<! {")
    fmt.Print(be)
    fmt.Print("} !>");
    i += 1;
  };
  fmt.Println("\n");
  return;
}