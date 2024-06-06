#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

int MyRT::goInit(MyRT* rt,char* [],int nv) {
  rt->to(rt->ident);
  rt->inArray = false;
  return 0;
}


int MyRT::goDone(MyRT* rt,char* p[],int nv) {
  return goLoop(rt,p,nv);
}

int MyRT::goLoop(MyRT* rt,char* p[],int nv) {
  rt->to("\n"); rt->setIdent(rt->ident-2),rt->to(rt->ident);
  if((rt->gen == RUST) || (rt->gen == GO)) rt->to("};\n"); 
  if((rt->gen == MOJO) || (rt->gen == PYTHON)) rt->to("pass\n");
  return 0;
}

int MyRT::goAmen(MyRT* rt,char* p[],int nv) {
  rt->inProc = false; 
  rt->to("\n"); 
  if(rt->gen == ASM) {
    rt->to("  ret\n");
    return 0;
  };
  rt->setIdent(rt->ident-2),rt->to(rt->ident);
  if((rt->gen == RUST) || (rt->gen == GO)) rt->to("}\n"); 
  if((rt->gen == MOJO) || (rt->gen == PYTHON)) rt->to("pass\n");
  return 0;
}

