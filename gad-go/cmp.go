package main

import "strings"

func Cmp(a string, d []string) bool {
  var i = 0
  var n = len(d)
  for i < n {
    ii := strings.Index(a,d[i])
    if ii > -1 { return true; }
    i += 1
  }
  return false
}

