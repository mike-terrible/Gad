package main

import "strings"


func Cmp(a string, d []string) bool {
  
  for _, v := range d {
    if strings.HasPrefix(a,v) { return true; }
  }
 
  return false
}

