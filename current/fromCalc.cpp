#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

struct Buf { char b[255]; };

const char* Op[] =  { "+","add", "-","sub", "*","mul", "/","div", "%","mod",
 "<","lt",">","gt", "<=", "le", ">=", "ge","!=", "ne", "<-", "to",
  "==", "eq",nullptr 
};
static Buf St [256];
static int zj = 0;

static int isOp(MyRT* rt,char* t) { return rt->cmp(t,Op); }

static void eoi(MyRT* rt) {
  switch (rt->gen) {
  case GO: case RUST : rt->to(";\n");  break; 
  case MOJO: case PYTHON: rt->to("\n"); break;
  default: break;
  }
}

static char result[255];

static void asmAllocResult(MyRT* rt) {
  sprintf(result,"gad_%d",zj);
  rt->to("  .data\n");
  rt->to(result),rt->to(":\n");
  rt->to("  .quad 0\n");
  rt->to("  .text\n  nop\n");
}

static void asmAss(MyRT* rt,char* xto,char* xfrom) {
  rt->onDebug("asmAss");
  char from[128]; char to[128];
  char q = xfrom[0];
  bool literal = false;
  if((q>='0') && (q<='9')) {
   literal = true;
   from[0]='$';
   strcpy(&from[1],xfrom);
   rt->to("  mov "),rt->to(from),rt->to(",%rsi\n");
  } else {
    strcpy(from,xfrom);
    rt->to("  lea ");
    if(memcmp("gad_",from,4)!=0) rt->to(rt->curProc),rt->to(".");
    rt->to(from),rt->to("(%rip),%rsi\n");
  };
  strcpy(to,xto);
  rt->to("  lea ");
  if(memcmp("gad_",to,4)!=0) rt->to(rt->curProc),rt->to(".");
  rt->to(to),rt->to("(%rip),%rdi\n");
  if(literal) { rt->to(" mov %rsi,%rbx\n"); } 
  else { rt->to("  mov (%rsi),%rbx\n"); };
  rt->to("  mov %rbx,(%rdi)\n");
}

static void asmOp1(MyRT* rt,const char* xop) {
  if(strcmp(xop," + 1") == 0) {
    rt->to("  inc %rbx\n"); return;
  };
}

static void allocResult(MyRT* rt) {
  if(rt->gen == ASM) { asmAllocResult(rt); return; };
  sprintf(result,"gad_%d",zj);
  rt->to("\n"); 
  rt->to(rt->ident);
  if((rt->gen == MOJO)||(rt->gen == GO)) rt->to("var "),rt->to(result),rt->to(" ");
  if(rt->gen == RUST) rt->to("let "),rt->to(result);
  zj++;
}

static int goOp1(MyRT* rt,const char* xop,int nt)  {
  int top = nt;
  top --; if(top<0) return 0;; 
  char xn1[255];  strcpy(xn1, St[top].b ); 
  allocResult(rt);
  strcpy(St[top].b ,result); top ++;
  if(rt->gen == ASM) {
    char q = xn1[0];
    if((q>='0') && (q<='9')) {
      char buf[128]; buf[0]='$';
      strcpy(&buf[1],xn1);
      rt->to("  mov "); rt->to(buf); rt->to(",%rbx\n"); 
    } else {
      rt->to("  lea ");
      if(memcmp("gad_",xn1,4)!=0) rt->to(rt->curProc),rt->to(".");
      rt->to(xn1),rt->to("(%rip),%rdi\n");
      rt->to("  mov (%rdi),%rbx\n");
    };
    asmOp1(rt,xop);
    rt->to("  mov %rbx,(%rdi)\n");
  } else {
    rt->to(" = ");
    rt->to(xn1); rt->to(xop); 
    eoi(rt);
  };
  return top;
}



static int goAss(MyRT* rt,int nt) {
  int top = nt;
  top --; if(top<0) return 0; char xn2[255]; strcpy(xn2, St[top].b);
  top --; if(top<0) return 0; char xn1[255]; strcpy(xn1, St[top].b);
  if(rt->gen == ASM) {
     asmAss(rt,xn2,xn1);
     return top;
  };
  rt->to("\n"); 
  rt->to(rt->ident);
  rt->to(xn2);
  rt->to(" = ");
  rt->to(xn1); 
  if(rt->gen == RUST) rt->to(";");
  rt->to("\n");
  return top;
}

static void asmOp2(MyRT* rt, const char* xop, char* xto, char* xfrom) {
  //
  char from[128]; char to[128];
  char q = xfrom[0];
  if((q >= '0') && (q <= '9')) {
    from[0]='$';  strcpy(&from[1],xfrom);
    rt->to("  mov "),rt->to(from),rt->to(",%rsi\n"); 
  } else {
    rt->to("  lea "); 
    if(memcmp("gad_",from,4)!=0) rt->to(rt->curProc),rt->to("."); 
    rt->to(from); rt->to("(%rip),%rsi\n");
    rt->to("  mov (%rsi),%rsi\n");
  };
  //
  q = xto[0];
  if((q >= '0') && (q <= '9')) {
    to[0]='$'; strcpy(&to[1],xto);
    rt->to("  mov "),rt->to(to),rt->to(",%rdi\n"); 
  } else {
    strcpy(to,xto);
    rt->to("  lea "); 
    if(memcmp("gad_",to,4)!=0) rt->to(rt->curProc),rt->to(".");
    rt->to(to); rt->to("(%rip),%rdi\n");
    rt->to("  mov (%rdi),%rdi\n");
  };
  //
  if(strcmp(xop," + ")==0) { rt->to("  add %rsi,%rdi\n"); }
  else if(strcmp(xop," - ")==0) { rt->to("  sub %rsi,%rdi\n"); }
  else if(strcmp(xop," * ")==0) { rt->to("  mul %rsi,%rdi\n"); }
  //
  rt->to("  lea "); rt->to(result); rt->to("(%rip),%rsi\n");
  rt->to("  mov %rdi,(%rsi)\n");
}


static int goOp2(MyRT* rt, const char* xop,int nt) {
  int top = nt;
  top --; if(top<0) return 0; char xn2[255]; strcpy(xn2, St[top].b);
  top --; if(top<0) return 0; char xn1[255]; strcpy(xn1, St[top].b);
  allocResult(rt);
  strcpy(St[top].b, result); top ++;
  if(rt->gen == ASM) {
    asmOp2(rt,xop,xn1,xn2);
    return top;
  };
  rt->to(" = ");
  rt->to(xn1); rt->to(xop); rt->to(xn2); 
  eoi(rt);
  return top;
}


int MyRT::fromCalc(char* varName, int iStart,int nv,char* p[]) {
  int i = iStart;
  int top = 0;
  for(;;) { i ++; if(i >= nv) break; 
    char* t = p[i];
    if(cmp(t,Repeat)) break; if(cmp(t,Then)) break;
    if(isOp(this,t)) {
      to(ident);
      if(cmp(t,"inc") || cmp(t,"++")) top = goOp1(this," + 1",top);
      else if(cmp(t, "<-") || cmp(t, "to")) top = goAss(this,top);
      else if(cmp(t, "<=") || cmp(t, "le")) top = goOp2(this, " <= ", top);
      else if(cmp(t, "<") || cmp(t,"lt")) top = goOp2(this, " < ",top);
      else if(cmp(t, ">=") || cmp(t,"ge")) top = goOp2(this, " >= ",top);
      else if(cmp(t, ">") || cmp(t,"gt")) top = goOp2(this, " > ", top);
      else if(cmp(t, "!=") || cmp(t,"ne")) top = goOp2(this, " != ",top);
      else if(cmp(t, "==") || cmp(t,"eq")) top = goOp2(this, " == ", top);
      else if(cmp(t,"add") || cmp(t,"+")) top = goOp2(this," + ",top); 
      else if(cmp(t,"sub") || cmp(t,"-")) top = goOp2(this," - ",top); 
      else if(cmp(t,"mul") || cmp(t,"*")) top = goOp2(this," * ",top); 
      else if(cmp(t,"div") || cmp(t,"/")) top = goOp2(this," / ",top); 
      else if(cmp(t,"mod") || cmp(t,"%")) top = goOp2(this," % ",top); 
    } else {
      strcpy(St[top].b, t);
      if(St[top].b[0] == '\"') {
        int nt = strlen(St[top].b);
        St[top].b[nt] = '\"';
      };
      top ++;
    };
  }; // for
  if(gen != ASM) {
    to("\n");
    to(ident);
  };
  if(strcmp(varName, "?")!=0) {
    if(gen == ASM) asmAss(this,varName,St[0].b);
    else {
      to(varName); to(" = "); to(St[0].b); 
      eoi(this);
    }; 
   }

  return 0;
}

int MyRT::goEval(MyRT* rt,char* p[],int nv) {
  int i = 0; char* t; char any[2];
  strcpy(any,"?");
  while(++i < nv) {
    t = p[i];
    if(rt->cmp(t,Repeat)) {
      rt->to("\n");
      rt->to(rt->ident);
      if(rt->gen == GO) rt->to("for {\n");
      if(rt->gen == RUST) rt->to("loop {\n");
      if((rt->gen == MOJO) || ( rt->gen ==  PYTHON)) rt->to("while True:\n"); 
      rt->setIdent(rt->ident+2);
      rt->fromCalc(any,0,nv,p);
      rt->to("\n"),rt->to(rt->ident);
      if((rt->gen == MOJO) || (rt->gen == PYTHON)) rt->to("if not "); 
      else rt->to("if !");
      rt->to(St[0].b);
      if((rt->gen == MOJO) || (rt->gen == PYTHON)) rt->to(":\n");
      if((rt->gen == GO) || (rt->gen == RUST)) rt->to(" {\n");
      rt->setIdent(rt->ident + 2); 
      rt->to(rt->ident);
      rt->to("break\n");
      rt->setIdent(rt->ident - 2);
      if(rt->gen == RUST ) rt->to(rt->ident),rt->to("};\n");
      if(rt->gen == GO ) rt->to(rt->ident),rt->to("};\n");
      return 0;
    };
    if(rt->cmp(t,Then)) {
      rt->fromCalc(any,0,nv,p); 
      rt->to("if "); 
      rt->to(St[0].b); 
      if((rt->gen == MOJO) || (rt->gen == PYTHON)) rt->to(":\n");
      else rt->to(" {\n");
      rt->setIdent(rt->ident + 2);
      return 0;
    };
    //
    return rt->fromCalc(any,0,nv,p);
    //
  };
  return 0;
}

