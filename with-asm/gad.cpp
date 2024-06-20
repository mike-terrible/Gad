#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include "gad.h"

using namespace Gad;

static MyRT *rt = NULL;

void gad(char* fn, char* opt) {
  FILE *f;  
  char* li = NULL;
  size_t nr = 0;
  int rd;
  char fname[1024];
  fname[0]=0;
  if(strstr(fn,"./") == NULL) strcpy(fname,"./");
  strcat(fname,fn);
  if(strstr(fname,".гад") == NULL) {
    strcat(fname,".гад"); 
  } else {
  };  
  f = fopen(fname,"r"); if(f == NULL) throw 1;
  rt = new MyRT(f,fname);
  rt->setGen(opt); 
  if(rt->gen != ASM) {
    printf("\n\n only -asm option allowed!\n\n");
    if(rt->data != NULL) fclose(rt->data);
    fclose(f);
    return;
  };
  switch(rt->gen) {
  case GO:
    rt->to("package main\n");
    rt->setIdent(2);
    break;
  default: break;  
  };
  rt -> st = ANY;
  char* t = NULL;
  while((rd = getline(&li,&nr,f)) != -1) {
    printf("--> %s",li);
    if(rt -> st == ANY) {
      t = rt->seekNotBlank(li); if(t == NULL) continue;
      rt->at(t)->seek(Gad::Comment); 
      if(rt->ok) {
        rt -> st = COMMENT; continue;
      };
      rt->parseIt(t);
    }
    else if(rt -> st == COMMENT) {
      t = rt->seekNotBlank(li); if(t == NULL) continue;
      if(rt->cmp(t, Gad::EndComment)) { rt -> st = ANY;  } 
      else {
        if(rt->gen == ASM) rt->to("#",li);  
        else {
          rt->to(rt->ident);
          if((rt->gen == GO) || (rt->gen == RUST)) rt->to("//",li);
          if((rt->gen == MOJO) || (rt->gen == PYTHON)) rt->to("#",li);
        };  
      };
      continue;
    };
  };
  // asm epilogur adding data section
  rt->to("  .data\n");
  fclose(rt->data); rt->data = NULL;
  rt->data = fopen("./data.s","r");
  while((rd = getline(&li,&nr,rt->data)) != -1) {
    rt->to(li);
  };
  fclose(rt->data); rt->data = NULL;
  unlink("./data.s");
  rt->to("  .section .note.GNU-stack,\"\",@progbits\n");
  rt->to("  .end\n");
  //  
  fclose(f);
  varDump();
}

int main(int argc,char** argv) {
  if(argv[1] == NULL) {
    printf("\n%s\ngad имя-файла[.гад] [ -asm ]\n по умолчанию -asm\n",MyRT::ver);
    return 0;
  };
  try {
    gad(argv[1],argv[2]);
  } catch(int ex) {
    printf("\nerr[%d]\n",ex);
  }
  return 0;
}
