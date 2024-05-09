#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

int MyRT::goIs(MyRT* rt,char* p[],int nv) { return rt->goIs(p,nv); }
int MyRT::goIs(char* p[],int nv) {
  if((gen == GO) || (gen == RUST)) to("{\n"); 
  setIdent(ident+2); return 0;
}
