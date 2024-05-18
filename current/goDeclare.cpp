#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

int MyRT::goDeclare(MyRT* rt,char* p[],int nv) { 
  int i = 0;
  char* var = NULL; char* like = NULL; char* vtype = NULL; char* be = NULL;
  char* vsize = NULL;
  char* val = NULL;
  i++; var = rt->getV(i,p,nv); 
  i++; like = rt->getV(i,p,nv);
  if(rt->cmp(like,Array)) { 
    i++; vsize = rt->getV(i,p,nv);
    i++; vtype = rt->getV(i,p,nv); 
    i++; be = rt->getV(i,p,nv);
    rt->goArray(p, nv, var, vsize, vtype, be ); 
    return 0;
  }; 
  i++; vtype = rt->getV(i,p,nv);
  i++; be = rt->getV(i,p,nv);
  i++; val = rt->getV(i,p,nv);
  rt->goVar(var,vtype,val);
  return 0;
}


