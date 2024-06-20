// gad-vars.cpp

#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

int Gad::NVar = 0;
Var Gad::Vars[256];

Var* Gad::varGet(char* xn) {
  int nv = NVar,i = nv;
  while(i > 0) {
    i--;
    if(strcmp(Vars[i].xname,xn)==0) return Vars + i;
  };
  return NULL;
}

Var* Gad::varNew(char* xn,bool isA,int asize,DType xtype) {
  int i = NVar; 
  NVar++;
  if(NVar>255) return NULL;
  strcpy(Vars[i].xname,xn);
  Vars[i].isArray = isA;
  Vars[i].asize = asize;
  Vars[i].dtype = xtype;
  return Vars + i;
}

void Gad::varDump(void) {
  printf("\n xref: \n");
  int i = 0;
  while(i < NVar) {
    printf("%s :",Vars[i].xname);
    DType dt = Vars[i].dtype;
    switch(dt) {
    case UNDEF: printf(" %s\n","Undef"); break;
    case LIGHT: printf(" %s\n","Light"); break;
    case NUM: printf(" %s\n","Num"); break;
    case REAL: printf(" %s\n","Real"); break;
    case STRING: printf(" %s\n","String"); break;
    default: break;
    };
    i++;
  };
  printf("\n");
}

 
