#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

int MyRT::cmp(const char* a,const char* b[] ) {
  int i = 0; const char* bb;
  while((bb = b[i++]) != nullptr) if(!memcmp(a,bb,strlen(bb))) return 1;
  return 0;
}


int MyRT::cmp(const char* a, const char* b) {
  int n = strlen(b);
  int rc = memcmp(a,b,n);
  if(rc == 0) return 1;
  return 0;
}


