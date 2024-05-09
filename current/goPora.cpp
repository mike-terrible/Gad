#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

int MyRT::goPora(MyRT* rt,char* p[],int nv) { return rt->goPora(p,nv); }
int MyRT::goPora(char* p[],int nv) {
  inProc = true;
  int i = 0;
  char* t = p[i]; 
  i++; xmain = getV(i,p,nv); if(xmain == NULL) return 0;
  i++; t = getV(i,p,nv); 
  if(t == NULL) { to(ident);
    if(gen == RUST) to("unsafe fn "),to(xmain),to("() {\n"); 
    if(gen == GO) to("func "),to(xmain),to("() {\n"); 
    if(gen == MOJO) to("fn "),to(xmain),to("() raises :\n"); 
    if(gen == PYTHON) to("def "),to(xmain),to("() :\n"); 
    setIdent(ident+2);
  } 
  else { to(ident);
    if(gen == RUST) to("fn main() {\n"); 
    if(gen == GO) to("func main() {\n"); 
    if(gen == MOJO) to("fn main() :\n"); 
    if(gen == PYTHON) to("def main() :\n"); 
    setIdent(ident+2),to(ident);
    if(gen == RUST) to("unsafe { ");
    to(xmain),to("()");
    if(gen == RUST) to("; }");
    to("\n");
    setIdent(ident-2),to(ident);
    if((gen == GO)||(gen == RUST)) to("}\n"); 
    to(ident);
    if(gen == RUST) to("unsafe fn "),to(xmain),to("() {\n"); 
    if(gen == GO) to("func "),to(xmain),to("() {\n"); 
    if(gen == MOJO) to("fn "),to(xmain),to("() :\n"); 
    if(gen == PYTHON) to("def "),to(xmain),to("() :\n"); 
    setIdent(ident+2);
  };
  return 0;
}
