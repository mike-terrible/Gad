#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

int MyRT::goProc(MyRT* rt,char* p[],int nv) { 
  rt->inProc = true;
  int i = 0;
  char* t = p[i];
  /*
  rt->to(rt->ident); 
  if(rt->gen == GO) rt->to("func "); 
  if(rt->gen == PYTHON) rt->to("def "); 
  if(rt->gen == MOJO) rt->to("fn "); 
  if(rt->gen == RUST) rt->to("unsafe fn ");
  */ 
  i++; char* xn = rt->getV(i,p,nv);
  strcpy(rt->curProc,xn);
  if(rt->gen == ASM) {
    rt->to("\n");
    rt->to(xn),rt->to(":"),rt->to("\n");
    return 0;
  };
  /*
  rt->to(xn),rt->to("("); int narg = 0;
  for(;;) { i++; if(i>=nv) { rt->to(") "); break; };
    char* itIs = rt->getV(i,p,nv);
    if(rt->cmp(itIs,Return)) { i++; char *act = rt->getV(i,p,nv);
      { char* ztype = rt->onType(act); 
        if(rt->gen == PYTHON) { rt->to(") :\n"); rt->setIdent(rt->ident + 2); return 0; };
        int nz = strlen(ztype);
        if(nz > 0) {
          if(rt->gen == GO) { 
            rt->to(") "),rt->to(ztype),rt->to(" {\n"); rt->setIdent(rt->ident + 2); return 0; 
          };
          if(rt->gen == MOJO) { rt->to(") -> "),rt->to(ztype),rt->to(" :\n"); 
            rt->setIdent(rt->ident + 2); return 0; 
          };
          if(rt->gen == RUST) { rt->to(") -> ");
            if(rt->cmp(ztype,"&str")) rt->to("String");
            else rt->to(ztype);
            rt->to(" {\n");
            rt->setIdent(rt->ident + 2); 
            return 0; 
          };
         };
      };
    }; 
    if(rt->cmp(itIs,Is)) {
      if((rt->gen == RUST) || (rt->gen == GO))  rt->to(") {\n"); 
      if((rt->gen == MOJO) || (rt->gen == PYTHON)) rt->to(") :\n"); 
      rt->setIdent(rt->ident+2); return 0;
    };
    if(rt->cmp(itIs,With)) { narg ++;
      i++; char* var = rt->getV(i,p,nv);
      i++; if(narg>1) rt->to(",");
      rt->to(var);
      char* like = rt->getV(i,p,nv); 
      if(rt->cmp(like,Aka)) {
        i++; char* xtype = rt->getV(i,p,nv);
        {
          if(rt->gen == GO) rt->to(" "),rt->to(rt->onType(xtype)); 
          if((rt->gen == MOJO) || (rt->gen == RUST)) rt->to(" :"),rt->to(rt->onType(xtype)); 
        };  
      };
    }; // with
  }; // for 
  */ 
  return 0;
}
