
// it.cpp
//
#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

It::It(MyRT* ru,const char* v[],Fn fn){
  go = fn;
  verb = v;
  mrt = ru;
}

