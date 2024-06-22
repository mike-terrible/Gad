
// asmOp2.cpp

#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

static void asmSetBit(MyRT* rt,const char* setv) {
  rt->to("  xor %rax,%rax\n");
  rt->to("  cmp %rsi,%rdi\n");
  rt->to(setv);
  rt->to("  mov %rax,%rdi\n");
}

void Gad::asmOp2(MyRT* rt, const char* xop, char* xto, char* xfrom) {
  //
  DType dt = asmTypeOf(xfrom); DType dtx = asmTypeOf(xto);
  if((dt == REAL)||(dtx == REAL)) {
     if(rt->cmp(xop," + ")) asmOp2Real("addsd ",rt,xto,xfrom);
     else if(rt->cmp(xop," - ")) asmOp2Real("subsd ",rt,xto,xfrom);
     else if(rt->cmp(xop," * ")) asmOp2Real("mulsd ",rt,xto,xfrom);
     else if(rt->cmp(xop," / ")) asmOp2Real("divsd ",rt,xto,xfrom);
     /****************************************************************/
     else if(rt->cmp(xop," == ")) asmCmpsd(/* cmpsd $0,*/ "cmpeqsd ",rt,xto,xfrom);
     else if(rt->cmp(xop," < ")) asmCmpsd(/* cmpsd $1, */ "cmpltsd ",rt,xto,xfrom);
     else if(rt->cmp(xop," <= ")) asmCmpsd(/* cmpsd $2,*/ "cmplesd ",rt,xto,xfrom);
     else if(rt->cmp(xop," != ")) asmCmpsd(/* cmpsd $4,*/ "cmpneqsd ",rt,xto,xfrom);
     else if(rt->cmp(xop," >= ")) asmCmpsd(/* cmpsd $5,*/ "cmpnltsd " ,rt,xto,xfrom);
     else if(rt->cmp(xop," > ")) asmCmpsd( /* cmpsd $6,*/ "cmpnlesd " , rt, xto, xfrom);
     return;
  };
  rt->to("# asmOp2 "),rt->to(xop),rt->to(" ");
  rt->to(xfrom),rt->to(","),rt->to(xto),rt->to("\n");
  char from[512]; char to[512];
  dt = typeOfLiteral(xfrom);
  if(dt==NUM) {
    sprintf(from,"$%s",xfrom);
    rt->to("  movq "),rt->to(from),rt->to(",%rsi\n"); 
  } else {
    if(memcmp("gad_",xfrom,4)!=0) sprintf(from,"%s.%s",rt->curProc,xfrom); 
    else strcpy(from,xfrom);
    rt->to("  lea ");
    rt->to(from); rt->to("(%rip),%rsi\n");
    rt->to("  movq (%rsi),%rsi\n");
  };
  //
  dt = typeOfLiteral(xto);
  if(dt==NUM) {
    sprintf(to,"$%s",xto);
    rt->to("  movq "),rt->to(to),rt->to(",%rdi\n"); 
  } else {
    if(memcmp("gad_",xfrom,4)!=0) sprintf(to,"%s.%s",rt->curProc,xto); 
    else strcpy(to,xto);
    rt->to("  lea "); 
    rt->to(to); rt->to("(%rip),%rdi\n");
    rt->to("  movq (%rdi),%rdi\n");
  };
  //
  if(memcmp(xop," + ",3)==0) rt->to("  add %rsi,%rdi\n");
  else if(memcmp(xop," - ",3)==0) rt->to("  sub %rsi,%rdi\n");
  else if(memcmp(xop," * ",3)==0) rt->to("  imul %rsi,%rdi\n");
  else if(memcmp(xop," / ",3)==0) rt->to("  idiv %rsi,%rdi\n"); 
  else if(memcmp(xop," == ",4)==0) asmSetBit(rt,"  sete %al\n");
  else if(memcmp(xop," < ",3)==0) asmSetBit(rt,"  setb %al\n");
  else if(memcmp(xop," <= ",4)==0) asmSetBit(rt,"  setbe %al\n");
  else if(memcmp(xop," > ",3)==0) asmSetBit(rt,"  seta %al\n");
  else if(memcmp(xop," >= ",4)==0) asmSetBit(rt,"  setae %al\n");
  else if(memcmp(xop," != ",4)==0) asmSetBit(rt,"  setne %al\n");
  //
  rt->to("  lea "); rt->to(result); rt->to("(%rip),%rsi\n");
  rt->to("  movq %rdi,(%rsi)\n");
}


