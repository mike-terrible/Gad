#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

int MyRT::goShow(MyRT* rt,char* p[],int nv) { return rt->goShow(p,nv); }

int MyRT::goShow(char* p[],int nv) {
  int i = 0;
  while(++i < nv) { char* t = getV(i,p,nv); if(t == NULL) break;
    if(cmp(t,With)) { i++; t = getV(i,p,nv); if(t == NULL) break;
      if(gen == RUST) {
        to(ident),to("print!(\"{ } \",");
        to(t); if(t[0]=='"') to("\""); to(");\n");
      } else if(gen == GO) {
        to(ident),to("print("); to(t); if(t[0]=='"') to("\""); to(",\" \");\n");
      } else if((gen == PYTHON)||(gen == MOJO)) {
        to(ident),to("print("); to(t); if(t[0]=='"') to("\""); to(",end =\" \")\n");
      }; 
    }; // With
  }; // While
  return 0;
}


