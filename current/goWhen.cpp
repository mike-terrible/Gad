#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

int MyRT::goWhen(MyRT* rt,char* p[],int nv) { return rt->goWhen(p,nv); }
int MyRT::goWhen(char* p[],int nv) {
  int i = 0;
  to(ident); 
  if(gen == GO) to("for");
  if((gen == RUST) || (gen == MOJO) || ( gen ==  PYTHON)) to("while"); 
  setIdent(ident+2);
  while(++i < nv) { char* t = getV(i,p,nv); if(t == NULL) return 0;
    if(cmp(t,Repeat)) {
      if((gen == GO) || (gen == RUST)) to(" {\n");
      if((gen == MOJO) || (gen == PYTHON)) to(" :\n");
      return 0;
    };
    to(" ");
    to(t); if(t[0]=='"') to("\""); 
  };
  return 0;
}

