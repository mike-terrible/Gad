#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

  int Gad::nA = 0;
  Pair Gad::AliasTab[256];
  
  char* Gad::aliasV(char* xxx) {
    int i = 0;
    while(i < nA) {
      char* t = AliasTab[i].k;
      if(t == nullptr) return nullptr;
      char* v = AliasTab[i].v;
      int n = strlen(t);
      if( strstr(t,xxx) != nullptr ) return AliasTab[i].v;
      i++;
    };
    return nullptr;
  }

  void Gad::allocAlias(char* k,char* v) {
    AliasTab[nA].k = strdup(k);
    AliasTab[nA].v = strdup(v);
    printf("\n alias [%s] as [%s]\n",k,v);
    nA++;
    AliasTab[nA].k = nullptr;
    int i = 0;
    while(i <= nA) {
      printf("<%s> : <%s>\n",AliasTab[i].k,AliasTab[i].v);
      i++;
    };
  }


int MyRT::goAlias(MyRT* rt,char* p[],int nv) {
  int i = 0;
  i += 1;  
  char* k = rt->getV(i,p,nv);
  i += 1; 
  char* v = rt->getV(i,p,nv);
  printf("\n alias %s -> %s\n",k,v);
  allocAlias(k,v);
  return 0;
}
