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

static void allocResult(MyRT* rt) {
  sprintf(result,"gad_%d",zj);
  rt->to("\n"); 
  rt->to(rt->ident);
  if((rt->gen == MOJO)||(rt->gen == GO)) rt->to("var "),rt->to(result),rt->to(" ");
  if(rt->gen == RUST) rt->to("let "),rt->to(result);
  if(rt->gen == PYTHON) rt->to(result),rt->to(" ");
  zj++;
}

static int goOp1(MyRT* rt,const char* xop,int nt)  {
  int top = nt;
  top --; if(top<0) return 0;; 
  char xn1[255];  strcpy(xn1, St[top].b ); 
  allocResult(rt);
  strcpy(St[top].b ,result); top ++;
  rt->to(" = ");
  rt->to(xn1); rt->to(xop); 
  eoi(rt);
  return top;
}



static int goAss(MyRT* rt,int nt) {
  int top = nt;
  top --; if(top<0) return 0; char xn2[255]; strcpy(xn2, St[top].b);
  top --; if(top<0) return 0; char xn1[255]; strcpy(xn1, St[top].b);
  rt->to("\n"); 
  rt->to(rt->ident);
  rt->to(xn2);
  rt->to(" = ");
  rt->to(xn1); 
  if(rt->gen == RUST) rt->to(";");
  rt->to("\n");
  return top;
}


static int goOp2(MyRT* rt, const char* xop,int nt) {
  int top = nt;
  top --; if(top<0) return 0; char xn2[255]; strcpy(xn2, St[top].b);
  top --; if(top<0) return 0; char xn1[255]; strcpy(xn1, St[top].b);
  allocResult(rt);
  strcpy(St[top].b, result); top ++;
  rt->to(" = ");
  if(rt->gen == GO) {
    if(strstr(" == ; != ; > ; < ; >= ; <= ",xop) != NULL) {
       rt->needBoolOf = 1;
       rt->to("BoolOf(");
       rt->to(xn2);
       rt->to(xop);
       rt->to(xn1);
       rt->to(")");
    } else {
      rt->to(xn2); rt->to(xop); rt->to(xn1);
    };
  } else if(rt->gen == RUST) {
    if(strstr(" == ; != ; > ; < ; >= ; <= ",xop) != NULL) {
       rt->needBoolOf = 1;
       rt->to("bool_of(");
       rt->to(xn2);
       rt->to(xop);
       rt->to(xn1);
       rt->to(")");
    } else {
      rt->to(xn2); rt->to(xop); rt->to(xn1);
    };
  } else if((rt->gen == PYTHON) || (rt->gen == MOJO)) {
    if(strstr(" == ; != ; > ; < ; >= ; <= ",xop) != NULL) {
      rt->needBoolOf = 1;
      rt->to("bool_of(");
      rt->to(xn2);
      rt->to(xop);
      rt->to(xn1);
      rt->to(")");
    } else {
      rt->to(xn2); rt->to(xop); rt->to(xn1);
    };
  } else {
    rt->to(xn2); rt->to(xop); rt->to(xn1); 
  };
  eoi(rt);
  return top;
}

int MyRT::fromEvil(char* varName, int iStart,int nv,char* p[]) {
  int npp  = 0;
  static char brOpen[2]; brOpen[0]='('; brOpen[1]=0;
  char* pp[128];
  char* ops[128]; char* aa[128];
  int cop = 0; int carg = 0;
  int i = iStart;;
  for(;;) {
    i++; if(i>=nv) break; 
    char* t =  p[i];
    if(t[0] == ')') {
      if(cop>0) { cop--;
        char* xop = ops[cop];
        if(isOp(this,xop)) {
          char* a1 = NULL; char* a2 = NULL;
          if(carg > 0) {
            carg--;  a1 = aa[carg]; if(a1[0]!='(') pp[npp]=a1,npp++;
          };
          if(carg > 0) {
            carg--; a2 = aa[carg]; if(a2[0]!='(') pp[npp]=a2,npp++;
          };
          pp[npp]=xop,npp++;
        };
      };
      continue;
    };
    if(isOp(this,t)) { ops[cop] = t; cop++; }
    else {
      aa[carg] = t; carg++;
    };
  }; // for
  to("\n");
  onDebug("p:"," ");
  { int j = 0;
    while(j<npp) {
      char* a = pp[j]; j++;
      if(a[0] != '(') to(a),to(" ");
    };
  };
  to("\n");
  fromCalc(varName,-1,npp,pp);
  return 0;
}

int MyRT::fromCalc(char* varName, int iStart,int nv,char* p[]) {
  int i = iStart;
  int top = 0;
  for(;;) { i ++; if(i >= nv) break; 
    char* t = p[i];
    if(cmp(t,Repeat)) break; if(cmp(t,Then)) break;
    if(cmp(t,"(")) return fromEvil(varName,i,nv,p);
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
  to("\n");
  to(ident);
  if(strcmp(varName, "?")!=0) {
    to(varName); to(" = "); to(St[0].b); 
    eoi(this);
  };

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

