package main

import "strings"

var Ab strings.Builder;

var wasBuf = "";
var wasNL string = "";

func Wr(p string) {
  var n = len(p);
  if wasNL == "\n" {
    if p[0:1] == "\n" {
      var n = len(p);
      if n>0 { Out.WriteString(p[1:]); wasNL = p[n-1:n]; return; };
    };
  };
  Out.WriteString(p);
  wasBuf = p; 
  if n>0 { wasNL = p[n-1:n]; }
}


func Da(b string) {
  Ab.WriteString(b); 
}

func To(n int) {
  var i = 0;
  SetIdent(n);
  var bb strings.Builder
  bb.WriteString("\n");
  for i < n { bb.WriteString(" "); i += 1; }
  var z = bb.String();
  if z != wasBuf { Wr(z); wasBuf = z; }
}

