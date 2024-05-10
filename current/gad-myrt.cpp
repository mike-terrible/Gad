#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

MyRT::MyRT(FILE* f) { fi = f; setIdent(0); inProc = false;
}

int MyRT::need(const char* fn) {
  out = fopen(fn,"w"),printf("\n создаём %s\n",fn);
  return 0;
}

int MyRT::setGen(char* opt) {
  gen = GO; 
  if(opt != NULL) {
    if(cmp(opt,"-mojo")) gen = MOJO;
    else if(cmp(opt,"-rust")) gen = RUST;
    else if(cmp(opt,"-python")) gen = PYTHON;
  };
  switch(gen) {
  case MOJO: return need("./out.mojo"); 
  case PYTHON: return need("./out.py"); 
  case RUST: return need("./out.rs"); 
  case GO: default: return need("./out.go");
  }; 
  
}

void MyRT::setIdent(int ii) {
  ident = ii;
}

char* MyRT::getV(int idx,char* p[],int nv) {
  //if(idx>=nv)  return NULL;
  if(idx >= nv) {
    printf("\n nv = %d,idx = %d\n",nv,idx);
    throw 665;
  };
  return p[idx];
}

char* MyRT::seekNotBlank(char* b) {
  int i = 0;
  while(b[i] != 0) { if(b[i] == ' ') { i++; continue; }; return &b[i]; };
  return NULL;
}

void MyRT::done() {
  if(out != NULL) fclose(out),out = NULL;
}



