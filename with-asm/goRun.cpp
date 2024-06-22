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
  rt->to(rt->ident);
  rt->to(" .global main\n");
  rt->to(" .text\n");
  rt->to("main: push %rax\n");
  rt->to(rt->xmain),rt->to(": xor %rax,%rax\n");
  strcpy(rt->curProc,rt->xmain);
  return 0;
}
