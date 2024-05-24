
// gad-parseGo.cpp
//
#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;


int MyRT::goParse(char* pp[],int nv) {
  char* p[256];
  int i = 0;
  while(i<nv) {
    char* tv = pp[i];
    char* v = aliasV(tv);
    if(v != nullptr) p[i] = v; else p[i] = tv; 
    printf(" {%s}",p[i]);
    i++;
  };  
  printf("\n");
  if(nv==0) return 0;
  i = 0;
  char* t = getV(i,p,nv); if(t == NULL) return 0; 
  static Gad::It ai[]  =  { 
    It(nullptr,Loop,MyRT::goLoop),
    It(nullptr,Done,MyRT::goDone),
    It(nullptr,Return,MyRT::goReturn),
    It(nullptr,When,MyRT::goWhen), 
    It(nullptr,Sic,MyRT::goSic), 
    It(nullptr,Else,MyRT::goElse),
    It(nullptr,Then,MyRT::goThen), 
    It(nullptr,If,MyRT::goIf), 
    It(nullptr,Give,MyRT::goGive), 
    It(nullptr,Job,MyRT::goJob),
    It(nullptr,Show,MyRT::goShow), 
    It(nullptr,Mess,MyRT::goMess),  
    It(nullptr,Run,MyRT::goRun), 
    It(nullptr,Amen,MyRT::goAmen), 
    It(nullptr,Declare,MyRT::goDeclare), //
    It(nullptr,Is,MyRT::goIs), 
    It(nullptr,Proc,MyRT::goProc),
    It(nullptr,Init,MyRT::goInit),
    It(nullptr,With,MyRT::goWith),
    It(nullptr,Alias,MyRT::goAlias),
    It(nullptr,nullptr,nullptr)
  };
  int j = 0;
  while(ai[j].verb != nullptr) {
    ai[j].mrt = this; 
    if(cmp(t,ai[j].verb)) try {
      ai[j].go(this,p,nv);
      return 0;
    } catch(int brk) {
      return brk;
    };
    j++;
  };
  gadError(t,p,nv);
  return 0;
}

