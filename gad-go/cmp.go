package main

import "strings"


func Cmp(a string, d []string) bool {
  var n = len(d)
  var i = 0  
  for i < n {
    var v = d[i]
    if strings.HasPrefix(a,v) { return true; }
    i += 1
  }
 
  return false
}

