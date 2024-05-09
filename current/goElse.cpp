#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

int MyRT::goElse(MyRT* rt,char* p[],int nv) { return rt->goElse(p,nv); }
int MyRT::goElse(char* p[],int nv) {
  setIdent(ident-2); int i = 0;
  if((gen == GO) || (gen == RUST)) to(ident),to("} else {\n"); 
  if((gen == MOJO) || (gen == PYTHON)) to(ident),to("else:\n"); 
  setIdent(ident+2);
  return 0;
}
