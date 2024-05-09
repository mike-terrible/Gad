
// gad-parseGo.cpp
//
#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

int MyRT::goParse(char* p[],int nv) {

  int i = 0;
  while(i<nv) printf(" {%s}",p[i++]);
  printf("\n");
  i = 0;
  char* t = getV(i,p,nv); if(t == NULL) return 0;

  static Gad::It** ai = nullptr;
  if(ai == nullptr) ai = new It* [] {
    new It(this,Loop,MyRT::goLoop),
    new It(this,Done,MyRT::goDone),
    new It(this,Return,MyRT::goReturn),
    new It(this,When,MyRT::goWhen), 
    new It(this,Sic,MyRT::goSic), 
    new It(this,Else,MyRT::goElse),
    new It(this,Then,MyRT::goThen), 
    new It(this,If,MyRT::goIf), 
    new It(this,Give,MyRT::goGive), 
    new It(this,Job,MyRT::goJob),
    new It(this,Show,MyRT::goShow), 
    new It(this,Skrepa,MyRT::goSkrepa), 
    new It(this,Pora,MyRT::goPora),
    new It(this,Amen,MyRT::goAmen), 
    new It(this,Declare,MyRT::goDeclare), 
    new It(this,Is,MyRT::goIs), 
    new It(this,Delo,MyRT::goDelo),
    nullptr
  };
  int j = 0;
  while(ai[j] != nullptr) {
    if(cmp(t,ai[j] -> verb)) return ai[j] -> go(this,p,nv);
    j++;
  };
  return 0;
}

