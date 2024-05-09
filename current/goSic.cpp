#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

int MyRT::goSic(MyRT* rt,char* p[],int nv) { return rt->goSic(p,nv); }
int MyRT::goSic(char* p[],int nv) {
  to(ident); int i = 0;
  for(;;) {
    i++; 
    char* t = getV(i,p,nv); 
    if(t == NULL) { 
      if(gen == RUST) to(";"); 
      to("\n"); return 0; 
    };
    to(t); if(t[0]=='"') to("\""); to(" "); 
  };
  return 0;
}

