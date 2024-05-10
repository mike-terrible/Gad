#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

int MyRT::goGive(MyRT* rt,char* p[],int nv) {
  int i = 0;
  i++; char* t = rt->getV(i,p,nv); 
  rt->to(rt->ident);
  if((rt->gen == GO) || ( rt->gen == MOJO)) rt->to("var "); 
  if(rt->gen == RUST) rt->to("let mut "); 
  rt->to(t);
  rt->to(" = "),i++,t = rt->getV(i,p,nv);
  /*
    if(t!=NULL) {
      if(cmp(t,"из")) { };
    };
   */
  i++; t = rt->getV(i,p,nv),rt->to(t),rt->to("("); int np = 0;
  while(++i < nv) { t = rt->getV(i,p,nv); 
    if(rt->cmp(t,With)) { i++; t = rt->getV(i,p,nv); 
      np++; if(np>1) rt->to(",");
      rt->to(t); if(t[0]=='"') rt->to("\"");
    };
  };
  if(rt->gen == RUST) rt->to(");\n"); else rt->to(")\n"); 
  return 0;
}
 
