#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

int MyRT::goReturn(MyRT* rt,char* p[],int nv) {
  int i = 0;
  i++; 
  if(i == nv) { rt->to("  ret\n"); return 0; };
  char* t = rt->getV(i,p,nv);
  // return t;
  //
  return 0;
}

