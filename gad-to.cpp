#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

void MyRT::to(const char* p1,const char* p2) { if(out == NULL) return;
  if(p1 != NULL) fprintf(out,"%s",p1); fprintf(out," %s\n",p2);
}

void MyRT::to(int xident) {
  int j=0;
  while(j<xident) to(" "),j++;
}

void MyRT::to(const char* p) {
  if(out == NULL) return;
  fprintf(out,"%s",p);
}

