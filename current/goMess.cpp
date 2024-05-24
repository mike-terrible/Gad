#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

int MyRT::goMess(MyRT* rt,char* p[],int nv) {  
  int i = 0;
  i++;
  char* t = rt->getV(i,p,nv); 
  rt->to(rt->ident); 
  if(rt->gen == RUST) {
    rt->to("println!(");
    if(t[0]=='\"') rt->to(t),rt->to("\"");
    else rt->to("\"{ }\","),rt->to(t);
    rt->to(");\n"); return 0;
  };
  if(rt->gen == GO) {
    rt->to("println("),rt->to(t);
    if(t[0]=='\"') rt->to("\"");
    rt->to(")\n");
    return 0;
  };
  if((rt->gen == MOJO) || (rt->gen == PYTHON)) {
    rt->to("print("),rt->to(t);
    if(t[0]=='\"') rt->to("\"");
    rt->to(")\n");
  };
  return 0;
}

