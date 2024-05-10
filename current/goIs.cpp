#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

int MyRT::goIs(MyRT* rt,char* p[],int nv) { 
  if((rt->gen == GO) || (rt->gen == RUST)) rt->to("{\n"); 
  rt->setIdent(rt->ident+2); return 0;
}
