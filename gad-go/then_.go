// 
// then_.gp
//
package main

func GenThen() { 
  switch Mode {
  case GO,RUST:  { Wr(" {\n");  }
  case MOJO,PYTHON: { Wr(" :\n"); }
  default:
  };  
  To(GetIdent() + 2);
}


