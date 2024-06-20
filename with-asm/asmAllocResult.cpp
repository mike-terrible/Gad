
// asmAllocResult.cpp

#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

char Gad::result[255];

void Gad::asmAllocResult(MyRT* rt) {
  sprintf(Gad::result,"gad_%d",Gad::zj);
  rt->da(Gad::result),rt->da(":\n");
  rt->da("  .quad 0\n");;
  Gad::zj ++;
}




