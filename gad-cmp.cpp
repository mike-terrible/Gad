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


MyRT* MyRT::seek(vector<const char*> b) {
  ok = false;
  int i = 0,n = b.size(); const char* bb;
  while(i<n) {
    bb = b[i++];
    if(!memcmp(atom,bb,strlen(bb))) { ok = true; return this; };
  };
  ok = false;
  return this;
}


bool MyRT::cmp(const char* a,vector<const char*> b) {
 const char* bb;
 int i = 0,n = b.size();
 while(i<n) {
   bb = b[i++];
   if(!memcmp(a,bb,strlen(bb))) return true;
 };
 return false;
}

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


