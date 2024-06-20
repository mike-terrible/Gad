// asmVar.cpp

#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;



int Gad::NVar = 0;
Var Gad::Vars[256];

long Gad::valReal(const char* a) {
  SurReal d;
  d.real = strtod(a,NULL);
  return d.num;
}

long Gad::valNum(const char* a) {
  SurReal d;
  d.num = atol(a);
  return d.num;
}

DType Gad::typeOfLiteral(const char* t) {
  if(t[0]=='\"') return STRING;
  if((t[0]>='0') && (t[0]<='9')) {
    if(strstr(t,"e")!=NULL) return REAL;
    if(strstr(t,"E")!=NULL) return REAL;
    if(strstr(t,".")!=NULL) return REAL;
    return NUM;
  };
  return UNDEF;
}

DType Gad::asmTypeOf(const char* t) {
  Var* v = varGet(t); if(v != NULL) return v->dtype;
  return typeOfLiteral(t);
}

Var* Gad::varGet(const char* xn) {
  int i = NVar, nv = NVar;
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
    printf("%s.",Vars[i].pname);
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

 
