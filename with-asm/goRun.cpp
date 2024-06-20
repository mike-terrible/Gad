#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

int MyRT::goRun(MyRT* rt,char* p[],int nv) { 
  rt -> inProc = true;
  int i = 0;
  char* t = p[i]; 
  i++; rt->xmain = rt->getV(i,p,nv); 
  i++; t = rt->getV(i,p,nv); 
  { rt->to(rt->ident);
    if(rt->gen == ASM) {
      rt->to(" .global main\n");
      rt->to(" .text\n");
      rt->to("main: push %rax\n");
      rt->to(rt->xmain),rt->to(": xor %rax,%rax\n");
      strcpy(rt->curProc,rt->xmain);
      return 0;
    };
    if(rt->gen == RUST) rt->to("fn main() {\n"); 
    if(rt->gen == GO) rt->to("func main() {\n"); 
    if(rt->gen == MOJO) rt->to("fn main() :\n"); 
    if(rt->gen == PYTHON) rt->to("def main() :\n"); 
    rt->setIdent(rt->ident+2),rt->to(rt->ident);
    if(rt->gen == RUST) rt->to("unsafe { ");
    rt->to(rt->xmain),rt->to("()");
    if(rt->gen == RUST) rt->to("; }");
    rt->to("\n");
    rt->setIdent(rt->ident-2),rt->to(rt->ident);
    if((rt->gen == GO)||(rt->gen == RUST)) rt->to("}\n"); 
    rt->to(rt->ident);
    if(rt->gen == RUST) rt->to("unsafe fn "),rt->to(rt->xmain),rt->to("() {\n"); 
    if(rt->gen == GO) rt->to("func "),rt->to(rt->xmain),rt->to("() {\n"); 
    if(rt->gen == MOJO) rt->to("fn "),rt->to(rt->xmain),rt->to("() :\n"); 
    if(rt->gen == PYTHON) rt->to("def "),rt->to(rt->xmain),rt->to("() :\n"); 
    rt->setIdent(rt->ident+2);
  };
  return 0;
}
