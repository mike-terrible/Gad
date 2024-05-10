#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

int MyRT::goDeclare(MyRT* rt,char* p[],int nv) { 
  int i = 0;
  i++; char* var = rt->getV(i,p,nv); 
  i++; char* like = rt->getV(i,p,nv); 
  i++; char* vtype = rt->getV(i,p,nv);
  i++; char* be = rt->getV(i,p,nv);
  i++; char* val = rt->getV(i,p,nv);
  rt->goVar(var,vtype,val);
  return 0;
}
