
// gad-Debug.cpp

#include <string.h>
#include <stdio.h>
#include <stdlib.h>

#include "gad.h"

using namespace Gad;

void MyRT::onDebug(const char* s) {
  to("\n"),to(ident);
  switch(gen) {
  case RUST: case GO: { to("//!! "),to(s); break; }
  default: { to("#!! "),to(s); }
  };
  to("\n");
}

void MyRT::onDebug(char* s) {
  onDebug(s);
}