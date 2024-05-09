#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

int MyRT::goJob(MyRT* rt,char* p[],int nv) { return rt->goJob(p,nv); }

int MyRT::goJob(char* p[],int nv) {
  int i = 0;
  i++; char* t = getV(i,p,nv); to("\n"),to(ident),to(t),to("(");
  int np = 0;
  while(++i < nv) { t = getV(i,p,nv); if(t == NULL) break;
    if(cmp(t,With)) { i++; t = getV(i,p,nv); if(t == NULL) break;
      np++; if(np>1) to(",");
      to(t); if(t[0]=='"') to("\"");
    };
  };
  if(gen == RUST) to(");\n"); else to(")\n");
  return 0;
}


