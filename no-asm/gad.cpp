#include <string.h>
#include <stdio.h>
#include <stdlib.h>
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
  switch(rt->gen) {
  case ASM:
    printf("\n\n -asm not allowed!\n\n");
    exit(0);
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
  if(rt->gen == MOJO) if(rt->needBoolOf) {
    rt->to("\n");
    rt->to("fn bool_of(v: Bool) -> Int:\n");
    rt->to("  if v:\n");
    rt->to("    return 1\n");
    rt->to("  else:\n");
    rt->to("    return 0\n");
    rt->to("pass\n");
  };
  if(rt->gen == PYTHON) {
    if(rt->needBoolOf) {
      rt->to("\n");
      rt->to("def bool_of(v) :\n");
      rt->to("  if v :\n");
      rt->to("    return 1\n");
      rt->to("  return 0\n");
    };
    rt->to("main()\n");
  };
  if(rt->gen == GO) if (rt->needBoolOf) {
    rt->to("\n");
    rt->to("  func BoolOf(v bool) int {\n");
    rt->to("    if v { return 1; }\n");
    rt->to("    return 0;\n");
    rt->to("  }\n");
  };
  if(rt->gen == RUST) if (rt->needBoolOf) {
    rt->to("\n");
    rt->to("fn bool_of(v :bool) -> i64 {\n");
    rt->to("  if v { return 1; }\n");
    rt->to("  return 0;\n");
    rt->to("}\n");
  };
  fclose(f);
  varDump();
}

int main(int argc,char** argv) {
  if(argv[1] == NULL) {
    printf("\n%s\ngad имя-файла[.гад] [-go | -mojo | -rust | -python | -asm ]\n по умолчанию -go\n",MyRT::ver);
    return 0;
  };
  try {
    gad(argv[1],argv[2]);
  } catch(int ex) {
    printf("\nerr[%d]\n",ex);
  }
  return 0;
}
