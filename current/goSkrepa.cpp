#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

int MyRT::goSkrepa(MyRT* rt,char* p[],int nv) { return rt->goSkrepa(p,nv); } 
int MyRT::goSkrepa(char* p[],int nv) {
  int i = 0;
  i++;
  char* t = getV(i,p,nv); if(t == NULL) return 0;
  to(ident); 
  if(gen == RUST) {
    to("println!(");
    if(t[0]=='\"') to(t),to("\"");
    else to("\"{ }\","),to(t);
    to(");\n"); 
  } else if(gen == GO) {
    to("println("),to(t);
    if(t[0]=='\"') to("\"");
    to(")\n");
  } else if((gen == MOJO) || (gen == PYTHON)) {
    to("print("),to(t);
    if(t[0]=='\"') to("\"");
    to(")\n");
  };
  return 0;
}

