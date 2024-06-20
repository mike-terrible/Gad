#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

int MyRT::goReturn(MyRT* rt,char* p[],int nv) {
  int i = 0;
  i++; 
  if(i == nv) {
    if(rt->gen == ASM) {
      rt->to("  ret\n");
      return 0;
    };
    rt->to(rt->ident),rt->to("return");
    if(rt->gen == RUST) rt->to(";");;
    rt->to("\n");
    return 0;
  };
  char* t = rt->getV(i,p,nv);
  rt->to(rt->ident),rt->to("return "),rt->to(t);
  if(t[0]=='\"') {
    rt->to("\"");
    if(rt->gen == RUST) rt->to(".to_string()");
  };
  if(rt->gen == RUST) rt->to(";");
  rt->to("\n");
  return 0;
}

