
// asmOp2Real.cpp

#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

void Gad::asmCmpsd(const char* cmpsdval,MyRT* rt,char* xto,char* xfrom) {
  char from[512]; char to[512];
  rt->to("# asmCmpsd ");
  rt->to(cmpsdval),rt->to(xfrom),rt->to(","),rt->to(xto);
  rt->to("\n");
  bool literal = false;
  DType dt = typeOfLiteral(xfrom);
  if((dt == NUM) || (dt == REAL) ) {
     long dl = 0L;
     if(dt == NUM) dl = valReal(xfrom);
     if(dt == REAL) dl = valReal(xfrom);
     sprintf(from,"$%ld",dl);
     rt->to("# movq "); rt->to(xfrom); rt->to(",%rax\n");
     rt->to("  movq "); rt->to(from); rt->to(",%rax\n");
     rt->to("  movq %rax,%xmm8\n"); 
  } else {
    if(memcmp(xfrom,"gad_",4)!=0) sprintf(from,"%s.%s",rt->curProc,xfrom);
    else  strcpy(from,xfrom); 
    rt->to("  lea "); rt->to(from); rt->to(",%rsi\n");
    rt->to("  movq (%rsi),%xmm8\n");     
  };
  dt = typeOfLiteral(xto);
  if((dt == NUM) || (dt == REAL) ) {
    long dl = 0L;
    if(dt == NUM) dl = valReal(xto);
    if(dt == REAL) dl = valReal(xto);
    sprintf(to,"$%ld",dl); 
    rt->to("# movq "); rt->to(xto); rt->to(",%rax\n");
    rt->to("  movq "); rt->to(to); rt->to(",%rax\n");
    rt->to("  movq %rax,%xmm9\n");
  } else {
    if(memcmp(xto,"gad_",4)!=0) sprintf(to,"%s.%s",rt->curProc,xto);
    else  strcpy(to,xto); 
    rt->to("  lea "); rt->to(to); rt->to(",%rdi\n");
    rt->to("  movq (%rdi),%xmm9\n");         
  };
  rt->to("  "),rt->to(cmpsdval),rt->to(" %xmm8,%xmm9\n");
  rt->to("  lea "); rt->to(result); rt->to("(%rip),%rsi\n");
  rt->to("  movq %xmm9,%rbx\n");
  rt->to("  andq $1,%rbx\n");
  rt->to("  movq %rbx,(%rsi)\n");
}

void Gad::asmOp2Real(const char* opcode,MyRT* rt, char* xto,char* xfrom) {
  char from[512]; char to[512];
  rt->to("# asmOp2Real ");
  rt->to(opcode),rt->to(" ");  rt->to(xfrom); rt->to(","); rt->to(xto);
  rt->to("\n");
  bool literal = false;
  DType dt = typeOfLiteral(xfrom);
  if((dt == NUM) || (dt == REAL) ) {
     long dl = 0L;
     if(dt == NUM) dl = valReal(xfrom);
     if(dt == REAL) dl = valReal(xfrom);
     sprintf(from,"$%ld",dl);
     rt->to("# movq "); rt->to(xfrom); rt->to(",%rax\n");
     rt->to("  movq "); rt->to(from); rt->to(",%rax\n");
     rt->to("  movq %rax,%xmm8\n"); 
  } else {
    if(memcmp(xfrom,"gad_",4)!=0) sprintf(from,"%s.%s",rt->curProc,xfrom);
    else  strcpy(from,xfrom); 
    rt->to("  lea "); rt->to(from); rt->to(",%rsi\n");
    rt->to("  movq (%rsi),%xmm8\n");     
  };
  dt = typeOfLiteral(xto);
  if((dt == NUM) || (dt == REAL) ) {
    long dl = 0L;
    if(dt == NUM) dl = valReal(xto);
    if(dt == REAL) dl = valReal(xto);
    sprintf(to,"$%ld",dl); 
    rt->to("# movq "); rt->to(xto); rt->to(",%rax\n");
    rt->to("  movq "); rt->to(to); rt->to(",%rax\n");
    rt->to("  movq %rax,%xmm9\n");
  } else {
    if(memcmp(xto,"gad_",4)!=0) sprintf(to,"%s.%s",rt->curProc,xto);
    else  strcpy(to,xto); 
    rt->to("  lea "); rt->to(to); rt->to(",%rdi\n");
    rt->to("  movq (%rdi),%xmm9\n");         
  };
  rt->to("  "),rt->to(opcode),rt->to(" %xmm8,%xmm9\n");
  rt->to("  lea "); rt->to(result); rt->to("(%rip),%rsi\n");
  rt->to("  movq %xmm9,(%rsi)\n");
}

