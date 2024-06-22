#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

static int giveArray(MyRT* rt, int nv,char* p[]) {
  int i = 1; i++; 
  char* t = rt->getV(i,p,nv);
  //rt->to("\n");
  //rt->to(rt->ident); //Wr("// giveArray\n"); To(Ident);  
  //rt->to(t);
  i++; 
  t = rt->getV(i,p,nv); 
  if(!rt->cmp(t,With)) return 0;
  i++; 
  t = rt->getV(i,p,nv);
  //rt->to("[");
  //rt->to(t);
  //rt->to("] =");
  i += 1; 
  i += 1; 
  t = rt->getV(i,p,nv);
  rt->to(t); 
  if(t[0] == '\"') rt->to("\""); 
  //if(rt->gen == RUST) rt->to(";\n");
  //else rt->to("\n"); 
  return 0;
}

int MyRT::goGive(MyRT* rt,char* p[],int nv) {
  int i = 0;
  i++; char* t = rt->getV(i,p,nv); 
  if(rt->cmp(t,Array)) {
    //if(rt->gen == MOJO) return mojoGiveArray(rt,nv,p);
    return giveArray(rt,nv,p);
  };
  char* varName = t;
  //
  i++,t = rt->getV(i,p,nv); // from
  if(rt->cmp(t,Eval)) return rt->fromCalc(varName, i,nv,p);
  //
  //rt->to(rt->ident);
  //if((rt->gen == GO) || ( rt->gen == MOJO)) rt->to("var "); 
  //if(rt->gen == RUST) rt->to("let mut "); 
  //rt->to(varName);
  //rt->to(" = ");
  //
  if(rt->cmp(t,Array)) {
    i++; 
    char* aname = rt->getV(i,p,nv);
    //rt->to(aname); rt->to("[");
    i++; i++; t = rt->getV(i,p,nv);
    //rt->to(t); 
    //rt->to("]"); 
    //rt->to("\n");  
    return 0;
  };
  //
  i++; t = rt->getV(i,p,nv);  //rt->to(t),rt->to("("); 
  int np = 0;
  while(++i < nv) { t = rt->getV(i,p,nv); 
    if(rt->cmp(t,With)) { i++; t = rt->getV(i,p,nv); 
      np++; //if(np>1) rt->to(",");
      //rt->to(t); if(t[0]=='"') rt->to("\"");
    };
  };
  return 0;
}
 
