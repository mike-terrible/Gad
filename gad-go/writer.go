package main

import "strings"

var Ab strings.Builder;

var wasBuf = "";
var wasNL string = "";
var daNL string = "";
var daBuf = "";


func Wtr(p string) {
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

func Wr(p ... string) {
  //for _,a := range(p) { Wtr(a) }
  var n = len(p)
  var i = 0
  for i < n { Wtr(p[i]); i += 1; }
}

func Da(b string) {
  var n = len(b);
  if daNL == "\n" {
    if b[0:1] == "\n" {
      var n = len(b);
      if n>0 { Ab.WriteString(b[1:]); daNL = b[n-1:n]; return; };
    };
  };
  Ab.WriteString(b);
  daBuf = b; 
  if n>0 { daNL = b[n-1:n]; } 
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

