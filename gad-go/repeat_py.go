// py_repeat.go

package main

func PyRepeat() { 
  To(GetIdent());
  Wr("if ", Result, " != 1:\n");
  To(GetIdent() + 2); Wr("break\n");
  To(GetIdent() - 2 );
}

func MojoRepeat() {
  PyRepeat()
}


func PyLoop() { To(GetIdent()); Wr("pass\n"); }

func MojoLoop() { PyLoop(); } 

