
// gogad.go main
//
package main

import "os"
import "fmt"

const GO = "-go";
const RUST = "-rust";
const MOJO = "-mojo";
const PYTHON = "-python";
const ASM = "-asm";
var NeedBoolOf = false;

func main() {
  fmt.Println("gad compiler ", Ver);
  var z = os.Args;
  var n = len(z)
  var fn = ""
  var  mode = GO
  if n > 1 { fn = os.Args[1]; }
  if n > 2 { mode = os.Args[2]; }
  if n > 1 { Parser(fn,mode); return; }
  fmt.Println("USAGE: gad fname [-go | -rust | -mojo | -python | -asm ]")
}

