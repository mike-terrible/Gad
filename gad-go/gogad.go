
// gogad.go main
//
package main

import "os"
import "fmt"

func main() {
  fmt.Println("gad compiler rel 2.01");
  var z = os.Args;
  var n = len(z)
  var fn = ""
  var  mode = "-go"
  if n > 1 { fn = os.Args[1]; }
  if n > 2 { mode = os.Args[2]; }
  if n > 1 { Parser(fn,mode); return; }
  fmt.Println("USAGE: gad fname [-go | -rust | -mojo | -python")
}

