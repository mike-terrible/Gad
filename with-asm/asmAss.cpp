// asmAss.cpp

#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

void Gad::asmAss(MyRT* rt,char* xto,char* xfrom) {
  rt->to("# asmAss "),rt->to(xfrom),rt->to(","),rt->to(xto),rt->to("\n");
  char from[512]; char to[512];
  bool literal = false;
  DType dt = typeOfLiteral(xfrom);
  long dl = 0L;
  switch(dt) {
  case REAL:
    dl = valReal(xfrom);
    sprintf(from,"$%ld",dl);
    rt->to("# movq "); rt->to(xfrom); rt->to(",%rsi\n");
    rt->to("  movq "); rt->to(from); rt->to(",%rsi\n");
    break;
  case NUM:
    dl = valNum(xfrom);
    sprintf(from,"$%ld",dl);
    rt->to("# movq "); rt->to(xfrom); rt->to(",%rsi\n");
    rt->to("  movq "); rt->to(from); rt->to(",%rsi\n");  
    break;
  case UNDEF:
    if(memcmp(xfrom,"gad_",4)!=0) sprintf(from,"%s.%s",rt->curProc,xfrom);
    else strcpy(from,xfrom);
    rt->to("  lea "),rt->to(from),rt->to("(%rip),%rsi\n");
    rt->to("  movq (%rsi),%rsi\n");
    break;
  default: break;
  };
  if(memcmp("gad_",xto,4)!=0) sprintf(to,"%s.%s",rt->curProc,xto);
  else strcpy(to,xto);
  rt->to("  lea "); rt->to(to); rt->to("(%rip),%rdi\n");
  rt->to("  movq %rsi,(%rdi)\n");
}

