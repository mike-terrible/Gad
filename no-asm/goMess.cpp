#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

static int lit = 0;

static int getLit() { int v = lit; lit ++; return v; }

int MyRT::goMess(MyRT* rt,char* p[],int nv) {  
  int i = 0;
  i++;
  char* t = rt->getV(i,p,nv);
  if(rt->gen == ASM) {
    int lt = getLit();
    if(t[0]=='\"') {
      char b[128];
      int lt = getLit();
      sprintf(b,"gad_c%d", lt);
      rt->to("  .data\n");
      rt->to(b),rt->to(":\n");
      rt->to("  .asciz "),rt->to(t),rt->to("\""),rt->to("\n");
      rt->to("  .text\n");
      rt->to("  lea "),rt->to(b),rt->to("(%rip),%rdi\n");
      rt->to("  sub %rax,%rax\n");
      rt->to("  call puts\n");
      return 0;
    };
    char bvar[512]; char bfmt[512];
    sprintf(bvar,"%s.%s",rt->curProc,t);
    sprintf(bfmt,"%s.%s.cnv",rt->curProc,t);
    rt->to("  xor %rax,%rax\n"); 
    rt->to("  lea "),rt->to(bfmt); rt->to("(%rip),%rdi\n");
    rt->to("  lea "),rt->to(bvar); rt->to("(%rip),%rsi\n"); 
    Var *d = varGet(t);
    if(d!=NULL) if(d->dtype == NUM) {
      rt->to("  mov (%rsi),%rsi\n");
    };
    rt->to("  call printf\n");
    return 0;
  }; 
  rt->to("\n"),rt->to(rt->ident); 
  if(rt->gen == RUST) {
    rt->to("println!(");
    if(t[0]=='\"') rt->to(t),rt->to("\"");
    else rt->to("\"{ }\","),rt->to(t);
    rt->to(");\n"); return 0;
  };
  if(rt->gen == GO) {
    rt->to("println("),rt->to(t);
    if(t[0]=='\"') rt->to("\"");
    rt->to(")\n");
    return 0;
  };
  if((rt->gen == MOJO) || (rt->gen == PYTHON)) {
    rt->to("print("),rt->to(t);
    if(t[0]=='\"') rt->to("\"");
    rt->to(")\n");
  };
  return 0;
}

