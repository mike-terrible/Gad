#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

int MyRT::goShow(MyRT* rt,char* p[],int nv) { 
  int i = 0;
  while(++i < nv) { char* t = rt->getV(i,p,nv); 
    if(rt->cmp(t,With)) { i++; if(i>=nv) break;
      t = rt->getV(i,p,nv); 
      if(rt->gen == RUST) {
        rt->to(rt->ident),rt->to("print!(\"{ } \",");
        rt->to(t); if(t[0]=='"') rt->to("\""); rt->to(");\n");
      } else if(rt->gen == GO) {
        rt->to(rt->ident),rt->to("print("); rt->to(t); if(t[0]=='"') rt->to("\""); 
        rt->to(",\" \");\n");
      } else if((rt->gen == PYTHON)||(rt->gen == MOJO)) {
        rt->to(rt->ident),rt->to("print("); rt->to(t); if(t[0]=='"') rt->to("\""); rt->to(",end =\" \")\n");
      }; 
    }; // With
  }; // While
  return 0;
}


