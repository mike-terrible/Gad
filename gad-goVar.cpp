#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

void MyRT::goVar(char* var,char* vtype,char* val) {
  if(var == NULL) throw 2;
  to(ident);
  if((gen == GO) || (gen == MOJO)) to("var "),to(var);
  if(gen == RUST) to("let mut "),to(var);
  if(gen == PYTHON) to(var);
  if(vtype != NULL) {
    if(gen == GO) to(" "),to(onType(vtype));
    if(gen == RUST) to(":"),to(onType(vtype));
    if(gen == MOJO) to(":"),to(onType(vtype));
  };
  if(val != NULL) {
    if(cmp(val,"зажечь")) {
      if(gen == RUST) to(" = true;\n");
      if(gen == GO) to(" = true\n");
      if((gen == MOJO) || (gen == PYTHON)) to(" = True\n");
      return;
    };
    if(cmp(val,"погасить")) {
      if(gen == RUST) to(" = false;\n");
      if(gen == GO) to(" = false\n");
      if((gen == MOJO) || (gen == PYTHON)) to(" = False\n");
      return;
    };
    if(val[0]=='"') to(" = "),to(val),to("\"");
    else to(" = "),to(val); 
    if(gen == RUST) to(";\n"); else to("\n");
  };
}

