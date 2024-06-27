
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
 if Mode == ASM {
   Wr("\n"); Wr("# "); Wr(a); Wr("\n");
 };
 To(GetIdent());
 switch Mode {
 case RUST,GO : { Wr("// "); Wr(a); Wr("\n"); }
 case MOJO,PYTHON :  { Wr("# "); Wr(a); Wr("\n"); }
 };
}

