
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
    It(nullptr,Skrepa,MyRT::goSkrepa), 
    It(nullptr,Pora,MyRT::goPora),
    It(nullptr,Amen,MyRT::goAmen), 
    It(nullptr,Declare,MyRT::goDeclare), 
    It(nullptr,Is,MyRT::goIs), 
    It(nullptr,Delo,MyRT::goDelo),
    It(nullptr,nullptr,nullptr)
  };
  int j = 0;
  while(ai[j].verb != nullptr) {
    ai[j].mrt = this; 
    if(cmp(t,ai[j].verb)) return ai[j].go(this,p,nv);
    j++;
  };
  return 0;
}

