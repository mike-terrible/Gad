#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

int MyRT::goSic(MyRT* rt,char* p[],int nv) { 
  rt->to(rt->ident); int i = 0;
  for(;;) {
    i++; 
    if(i>=nv) {
      if(rt->gen == RUST) rt->to(";"); 
      rt->to("\n"); return 0;
    };
    char* t = rt->getV(i,p,nv); 
    rt->to(t); 
    if(t[0]=='"') rt->to("\""); rt->to(" "); 
  };
  return 0;
}

