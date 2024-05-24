#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

MyRT* MyRT::at(const char* a) { atom = a; return this; }

MyRT* MyRT::seek(const char* b[]) {
  ok = false;
  int i = 0; const char* bb;
  while((bb = b[i++]) != nullptr) 
   if(!memcmp(atom,bb,strlen(bb))) { ok = true; return this; };
  ok = false;
  return this;
}

int MyRT::cmp(const char* a,const char* b[] ) {
  int i = 0; const char* bb;
  //while((bb = b[i++]) != nullptr) if(!memcmp(a,bb,strlen(bb))) return 1;
  while((bb = b[i++]) != nullptr) if(strstr(a,bb) != nullptr) return 1;
  return 0;
}


int MyRT::cmp(const char* a, const char* b) {
  int n = strlen(b);
  if(strstr(a,b) != nullptr) return 1;
  //int rc = memcmp(a,b,n);
  //if(rc == 0) return 1;
  return 0;
}


