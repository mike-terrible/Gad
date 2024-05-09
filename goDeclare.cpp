#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

int MyRT::goDeclare(MyRT* rt,char* p[],int nv) { return rt->goDeclare(p,nv); }
int MyRT::goDeclare(char* p[],int nv) {
  int i = 0;
  i++; char* var = getV(i,p,nv); 
  i++; char* like = getV(i,p,nv); 
  i++; char* vtype = getV(i,p,nv);
  i++; char* be = getV(i,p,nv);
  i++; char* val = getV(i,p,nv);
  goVar(var,vtype,val);
  return 0;
}
