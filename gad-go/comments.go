
//comments.go

package main

var IsComment bool = false;
var IsLine = true;

func CheckComment(a string) { 
  switch {
  case Cmp(a, END_COMMENT ): {
    IsComment = false; IsLine = false;
  }
  case Cmp(a, BEGIN_COMMENT): {
    IsComment = true;  IsLine = false;
  }
  default: { IsLine = true; }
  };
}

func GenComment(a string) {
 if !IsLine { return; };
 if (Mode == ASM) || (Mode == ASM32) { Wr("\n","# ", a, "\n"); return; };
 To(GetIdent());
 switch Mode {
 case RUST,GO : Wr("// ", a ,  "\n" ); 
 case MOJO,PYTHON :  Wr("# ", a, "\n"); 
 };
}

