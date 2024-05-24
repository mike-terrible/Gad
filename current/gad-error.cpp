// gad-error.cpp
#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

void MyRT::gadError(char* s,char* p[],int nv) {
  printf("\n!!!word: ");
  for(int i = 0; i < nv; i++) {
    char* z = p[i];
    if(cmp(s,z)) printf(" !{ %s }",s);
    else printf(" {%s}",z); 
  };
  printf("\n");
  throw 666;
}

