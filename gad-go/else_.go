
// else_.go
//
package main

func GenElse()  {
  Ident -= 2; 
  switch Mode {
  case "-go","-rust": { To(Ident); Wr("} else {\n");  }
  case "-mojo","-python": { To(Ident); Wr("else:\n"); }
  default:
  }; 
  Ident += 2;
}

