#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;


MyRT::MyRT(FILE* f,char* fn) { fi = f; setIdent(0); inProc = false;
  //inFname = infn;
  //outFname = outfn;
  strcpy(infn,fn); strcpy(curProc,"modmain");
}

int MyRT::need(const char* fn) {
  out = fopen(fn,"w"),printf("\n создаём %s\n",fn);
  const char* about = "";
  switch(gen) {
  case ASM: about = "ASM"; break;
  case GO: about = "GO"; break;
  case MOJO: about = "MOJO"; break;
  case RUST: about = "RUST"; break;
  case PYTHON: about = "PYTHON"; break;
  default: break;
  };
  printf("\n codegen %s\n",about);
  return 0;
}

int MyRT::setGen(char* opt) {
  gen = GO; 
  if(opt != NULL) {
    if(cmp(opt,"-mojo")) gen = MOJO;
    else if(cmp(opt,"-asm")) gen = ASM;
    else if(cmp(opt,"-rust")) gen = RUST;
    else if(cmp(opt,"-python")) gen = PYTHON;
  };
  strcpy(outfn,infn);
  char* p = strstr(outfn,".гад");
  switch(gen) {
  case ASM: strcpy(p,".s"); return need(outfn);
  case MOJO: strcpy(p,".mojo"); return need(outfn);  
  case PYTHON: strcpy(p,".py"); return need(outfn); 
  case RUST: strcpy(p,".rs"); return need(outfn);
  case GO: strcpy(p,".go"); return need(outfn);
  default: return need("./out.go");
  }; 
  
}

void MyRT::setIdent(int ii) {
  ident = ii;
}

char* MyRT::getV(int idx,char* p[],int nv) {
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



