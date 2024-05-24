#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

struct Buf { char b[255]; };

const char* Op[] =  { "+","add", "-","sub", "*","mul", "/","div", "%","mod", nullptr };
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

static int goOp1(MyRT* rt,const char* xop,int nt)  {
  int top = nt;
  top --; char xn1[255];  strcpy(xn1, St[top].b );
  char result[255];; sprintf(result,"gad_%d",zj);
  zj ++;
  strcpy( St[top].b ,result); top ++;
  rt->to("\n");
  rt->to(rt->ident); 
  switch (rt->gen) { 
  case GO: MOJO: rt->to("var "); break; 
  case RUST: rt->to("let mut "); break;
  default: break;  
  };
  rt->to(result); rt->to(" = ");
  rt->to(xn1); rt->to(xop); 
  eoi(rt);
  return top;
}

static int goOp2(MyRT* rt, const char* xop,int nt) {
  int top = nt;
  top --; char xn2[255]; strcpy(xn2, St[top].b);
  top --; char xn1[255]; strcpy(xn1, St[top].b);
  char result[255]; sprintf(result,"gad_%d",zj);
  zj ++;
  strcpy(St[top].b, result); top ++;
  rt->to("\n");
  rt->to(rt->ident); 
  switch (rt->gen) { 
  case GO: case MOJO: rt->to("var "); break; 
  case RUST: rt->to("let mut "); break;
  default: break;  
  };
  rt->to(result); rt->to(" = ");
  rt->to(xn1); rt->to(xop); rt->to(xn2); 
  eoi(rt);
  return top;
}


int MyRT::fromCalc(char* varName, int iStart,int nv,char* p[]) {
  int i = iStart;
  int top = 0;
  for(;;) { i ++; if(i >= nv) break; 
    char* t = p[i];
    if(isOp(this,t)) {
      to(ident);
      if(cmp(t,"inc") || cmp(t,"++")) top = goOp1(this," + 1",top); 
      else if(cmp(t,"add") || cmp(t,"+")) top = goOp2(this," + ",top); 
      else if(cmp(t,"sub") || cmp(t,"-")) top = goOp2(this," - ",top); 
      else if(cmp(t,"mul") || cmp(t,"*")) top = goOp2(this," * ",top); 
      else if(cmp(t,"div") || cmp(t,"/")) top = goOp2(this," / ",top); 
      else if(cmp(t,"mod") || cmp(t,"%")) top = goOp2(this," % ",top); 
    } else {
      strcpy(St[top].b, t); top ++;
    };
  }; // for
  to("\n");
  to(ident);
  if(strcmp(varName, "?")!=0) {
    to(varName); to(" = "); to(St[0].b); 
    eoi(this); 
   }

  return 0;
}
