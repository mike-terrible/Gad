#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

int MyRT::goDelo(MyRT* rt,char* p[],int nv) { return rt->goDelo(p,nv); }
int MyRT::goDelo(char* p[],int nv) {
  inProc = true;
  int i = 0;
  char* t = p[i];
  to(ident); 
  if(gen == GO) to("func "); 
  if(gen == PYTHON) to("def "); 
  if(gen == MOJO) to("fn "); 
  if(gen == RUST) to("unsafe fn "); 
  i++; char* xn = getV(i,p,nv); if(xn == NULL) return 0;
  to(xn),to("("); int narg = 0;
  for(;;) { i++; if(i>=nv) { to(") "); break; };
    char* itIs = getV(i,p,nv);
    if(cmp(itIs,Return)) { i++; char *act = getV(i,p,nv);
      if(act != NULL) { char* ztype = onType(act); 
        if(gen == PYTHON) { to(") :\n"); setIdent(ident + 2); return 0; };
        int nz = strlen(ztype);
        if(nz > 0) {
          if(gen == GO) { to(") "),to(ztype),to(" {\n"); setIdent(ident + 2); return 0; };
          if(gen == MOJO) { to(") -> "),to(ztype),to(" :\n"); setIdent(ident + 2); return 0; };
          if(gen == RUST) { to(") -> ");
            if(cmp(ztype,"&str")) to("String");
            else to(ztype);
            to(" {\n"),setIdent(ident + 2); return 0; 
          };
         };
      };
    }; 
    if(cmp(itIs,Is)) {
      if((gen == RUST) || (gen == GO))  to(") {\n"); 
      if((gen == MOJO) || (gen == PYTHON)) to(") :\n"); 
      setIdent(ident+2); return 0;
    };
    if(cmp(itIs,With)) { narg ++;
      i++; char* var = getV(i,p,nv);
      i++; if(narg>1) to(",");
      to(var);
      char* like = getV(i,p,nv); 
      if(like != NULL) if(cmp(like,Aka)) {
        i++; char* xtype = getV(i,p,nv);
        if(xtype != NULL) {
          if(gen == GO) to(" "),to(onType(xtype)); 
          if((gen == MOJO) || (gen == RUST)) to(" :"),to(onType(xtype)); 
        };  
      };
    }; // with
  }; // for  
  return 0;
}
