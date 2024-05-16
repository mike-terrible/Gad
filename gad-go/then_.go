// 
// then_.gp
//
package main

func GenThen() { 
  switch Mode {
  case "-go","-rust":  { Wr(" {\n");  }
  case "-mojo","-python": { Wr(" :\n"); }
  default:
  };  
  Ident += 2;
}


