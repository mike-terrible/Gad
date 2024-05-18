
//comments.go

package main

var IsComment bool = false;

func CheckComment(a string) { 
  if Cmp(a, END_COMMENT ) {
    IsComment = false; 
    Wr("\n");
    return
  }
  if Cmp(a, BEGIN_COMMENT ) {
    IsComment = true; return
  }
}

func GenComment(a string) {
 switch Mode {
 case "-rust","-go" : { Wr("// ");  Wr(a); Wr("\n"); }
 case "-mojo", "-python" :  { Wr("# "); Wr(a); Wr("\n"); }
 default: 
 };
}

