#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

int MyRT::goIf(MyRT* rt,char* p[],int nv) {
  int i = 0;
  i++; char* t = rt->getV(i,p,nv); if(t == NULL) return 0;
  rt->to(rt->ident),rt->to("if"); 
  while(t!=NULL) { 
    rt->to(" "),rt->to(t); if(t[0] == '\"') rt->to("\"");
    i++; t = rt->getV(i,p,nv); if(t == NULL) return 0;
    if(rt->cmp(t,Then)) {
      if((rt->gen == RUST) || (rt->gen == GO)) rt->to(" {\n");
      if((rt->gen == MOJO) || (rt->gen == PYTHON)) rt->to(" :\n"); 
      rt->setIdent(rt->ident+2);
      return 0;
    }; 
  };
  return 0;
}

