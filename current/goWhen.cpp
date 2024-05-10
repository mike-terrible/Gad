#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

int MyRT::goWhen(MyRT* rt,char* p[],int nv) {
  int i = 0;
  rt->to(rt->ident); 
  if(rt->gen == GO) rt->to("for");
  else if((rt->gen == RUST) || (rt->gen == MOJO) || ( rt->gen ==  PYTHON)) rt->to("while"); 
  rt->setIdent(rt->ident+2);
  while(++i < nv) { char* t = rt->getV(i,p,nv); 
    if(rt->cmp(t,Repeat)) {
      if((rt->gen == GO) || (rt->gen == RUST)) rt->to(" {\n");
      if((rt->gen == MOJO) || (rt->gen == PYTHON)) rt->to(" :\n");
      return 0;
    };
    rt->to(" ");
    rt->to(t); if(t[0]=='"') rt->to("\""); 
  };
  return 0;
}

