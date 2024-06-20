#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

int MyRT::goWith(MyRT *rt,char* p[],int nv) {
  int i = 0; char* t = NULL; char buf[128];
  for(;;) {
    i++; 
    if(i>=nv) break;
    t = p[i];
    switch(rt->gen) {
    case GO: case RUST:
      rt->to(rt->ident);
      rt->to(rt->curVar);
      sprintf(buf,"%d",rt->inInit);
      rt->to("["); rt->to(buf); rt->to("]");
      rt->to(" = ");
      rt->to(t); 
      if(t[0]=='\"') rt->to("\"");
      if(rt->gen == RUST) rt->to(";");
      rt->to("\n");
      rt->inInit += 1;
      break;
    case MOJO: case PYTHON:
      rt->to(rt->ident);
      rt->to(rt->curVar);
      rt->to(".append(");
      rt->to(t);
      if(t[0]=='\"') rt->to("\"");
      rt->to(")\n");
      break;
    default: break;
    };
  };
  return 0;
}

void MyRT::goArray(char* p[],int nv, char* varV, char* vsize, char* vtype, char* be ) {
  if(varV == NULL) return;  to(ident);
  switch(gen) {
  case GO: {
    to("var ");
    to(varV); to(" ");
    to("["); 
    if(strcmp(vsize,"?")!=0) to(vsize);
    to("]");
    to(onType(vtype));
    if(be == NULL) { inArray = false; to(";\n"); return; };
    strcpy(curVar,varV);
    to("\n");
    inArray = true; 
    if(cmp(be,Aka)) { inInit = 0; }
    else { 
      inInit = 1;
      to(ident);
      to(varV);
      to("[0] = "); 
      to(be); if(be[0] == '\"') to("\""); 
      to("\n");
    }; 
    return;
  }
  case RUST: {
    if(inProc) { to("let mut "); to(varV); } else {  
      to("static mut "); to(varV);  
    };
    to(": ");
    to("[");
    to(onType(vtype));
    to(";");
    if(strcmp(vsize,"?")!=0) to(vsize); 
    to("]");
    to(" = ["); to(onValue(vtype)); to(";"); to(vsize); 
    to("];\n");
    if(be == NULL) { inArray = false; return; };
    strcpy(curVar,varV);
    inArray = true;  
    if(cmp(be,Aka)) {
      inInit = 0;
    } else {
      inInit = 1;
      to(ident);
      to(varV);
      to("[0] = "); 
      to(be); if (be[0] == '\"') to("\""); 
      to(";");
    };
    to("\n");
    return;
  }
  case PYTHON: {
    to(varV);
    to(" = [ ]\n");
    if(be == NULL) { inArray = false; to("\n"); return; };
    strcpy(curVar,varV);
    inArray = true; 
    if(cmp(be,Aka)) { inInit = 0; } else {
      inInit = 1;
      to(ident);
      to(curVar);
      to(".append("); 
      to(be); if(be[0] == '\"') { to("\""); };
    };    
    to("\n");
    return;
  }
  case MOJO: {
    to("var ");
    to(varV);
    to(" = List[");
    to(onType(vtype));
    to("]()");
    strcpy(curVar, varV);
    inArray = true; 
    if(be == NULL) { inArray = false; to("\n"); return; };
    if(cmp(be,Aka)) { inInit = 0; } else {
      inInit = 1;
      to(ident);
      to(curVar);
      to(".append("); 
      to(be); if(be[0] =='\"') { to("\""); };
    }; 
    to("\n");
    return;
   }
   default: break; 
   }; // switch
}


void MyRT::goVar(char* var, char* vtype, char* val) {
  if(var == NULL) throw 2; to(ident);
  if((gen == GO) || (gen == MOJO)) to("var "),to(var);
  if(gen == RUST) {
    if(inProc) to("let mut "),to(var);
    else to("static mut "),to(var);
  };
  if(gen == PYTHON) to(var);
  if(vtype != NULL) {
    if(gen == GO) to(" "),to(onType(vtype));
    if(gen == RUST) to(":"),to(onType(vtype));
    if(gen == MOJO) to(":"),to(onType(vtype));
  };
  if(val != NULL) {
    if(cmp(val,On)) {
      if(gen == RUST) to(" = true;\n");
      if(gen == GO) to(" = true\n");
      if((gen == MOJO) || (gen == PYTHON)) to(" = True\n");
      return;
    };
    if(cmp(val,Off)) {
      if(gen == RUST) to(" = false;\n");
      if(gen == GO) to(" = false\n");
      if((gen == MOJO) || (gen == PYTHON)) to(" = False\n");
      return;
    };
    to(" = "),to(val);
    if(val[0]=='\"') { 
      to("\"");
    };
    if(gen == RUST) to(";\n"); else to("\n");
  } else {
    switch(gen) {
    case RUST: to(";\n"); break;
    case GO: to("\n"); break;
    default: break;
    };  
  };
}

