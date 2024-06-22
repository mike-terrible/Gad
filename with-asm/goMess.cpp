#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

static int lit = 0;

static int getLit() { int v = lit; lit ++; return v; }

void asmMess(MyRT* rt,char* t) {
  rt->to("\n");
  rt->to("# asmMess "); rt->to(t); if(t[0]=='\"') rt->to("\"");
  rt->to("\n");
  char b[512]; char bval[512];
  int lt = 0;
  DType dt = typeOfLiteral(t);
  if(dt == STRING) {
    lt = getLit();
    sprintf(b,"gad_c%d", lt);
    rt->da(b),rt->da(":\n");
    rt->da("  .asciz "),rt->da(t),rt->da("\"\n");
    rt->to("  lea "),rt->to(b),rt->to("(%rip),%rdi\n");
    rt->to("  sub %rax,%rax\n");
    rt->to("  call puts\n");
    return;
  };  
  if(dt == NUM) {
    lt = getLit();
    sprintf(b,"gad_cnv%d",lt); 
    rt->da(b),rt->da(":\n");
    rt->da("  .asciz \""); rt->da("%ld"); 
    rt->da("\"\n");
    rt->to("  lea "),rt->to(b),rt->to("(%rip),$rdi\n");
    rt->to("  movq $"),rt->to(t),rt->to(",%si\n");
    rt->to("  sub %rax,%rax\n");
    rt->to("  call printf\n");
    return;
  };  
  if(dt == REAL) {
    lt = getLit();
    sprintf(b,"gad_cnv%d",lt); 
    rt->da(b),rt->da(":\n");
    rt->da("  .asciz \"%lg\"\n");
    long dd = valReal(t);
    sprintf(bval,"$%ld",dd);
    rt->to("  movq "),rt->to(bval),rt->to("%rax");
    rt->to("  movq $rax,%xmm0\n");
    rt->to("  movq %1,%rax\n");
    rt->to("  lea "),rt->to(b),rt->to("(%rip),%si\n");
    rt->to("  call printf\n");
    return;
  };
  if(memcmp(t,"gad_",4)==0) {
    return;
  };
  Var *v = varGet(t); if(v == NULL) return;
  dt = v->dtype;
  if(dt == STRING) {
    rt->to("  xor %rax,%rax\n");
    rt->to("  lea "); rt->to(rt->curProc),rt->to("."),rt->to(t); rt->to("(%rip),%rsi\n");
    rt->to("  lea "),rt->to(rt->curProc),rt->to("."),rt->to(t),rt->to(".cnv(%rip),%rdi\n");
    rt->to("  call printf\n");
    rt->onDebug("end of asmMess");
    return;
  };
  if(dt == NUM) {
    rt->to("  lea "); rt->to(rt->curProc),rt->to("."),rt->to(t); rt->to("(%rip),%rsi\n");
    rt->to("  lea "),rt->to(rt->curProc),rt->to("."),rt->to(t),rt->to(".cnv(%rip),%rdi\n");
    rt->to("  movq (%rsi),%rsi\n");
    rt->to("  xor %rax,%rax\n");
    rt->to("  call printf\n");
    rt->onDebug("end of asmMess");
    return;
  };
  if(dt == REAL) {
    rt->to("  movq $1,%rax\n");
    rt->to("  lea "),rt->to(rt->curProc),rt->to("."),rt->to(t),rt->to(",%rdi\n");
    rt->to("  movq (%rdi),%xmm0\n");
    rt->to("  lea "),rt->to(rt->curProc),rt->to("."),rt->to(t),rt->to(".cnv(%rip),%rdi\n");
    rt->to("  call printf\n");
    rt->onDebug("end of asmMess");
  };
}

int MyRT::goMess(MyRT* rt,char* p[],int nv) {  
  rt->onDebug("goMess");
  int i = 0;
  i++;
  char* t = rt->getV(i,p,nv);
  asmMess(rt,t);
  return 0;
}

