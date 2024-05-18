//
// amen.go
//
package main

func GenInit() {
  To(Ident);
  InArray = false;
}


func GenLoop() {
  Wr("\n"); Ident -= 2; To(Ident);
  switch Mode {
  case "-rust","-go" : Wr("};\n"); 
  case "-mojo","-python" : Wr("pass\n");
  default:
  }
}

func GenDone() { GenLoop(); }

func GenAmen() {
  InProc = false; 
  Wr("\n"); Ident -= 2; To(Ident);
  switch Mode {
  case "-rust", "-go": { Wr("}\n"); }
  case "-mojo","-python": { Wr("pass\n"); }
  default:
  };
}

