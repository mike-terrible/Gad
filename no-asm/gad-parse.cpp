#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

void MyRT::parseIt(char* t) {
  char* aa[256]; 
  int i = 0;
  int j = 0;
  int k = 0;
  char q = 0;
  char lex[1024];
  while(t[i] != 0) {
    q = t[i]; i++;
    if(q == 13) q = ' '; if(q == 10) q = ' ';
    if(q == '"') {
      lex[j]=q; j++; 
      q = t[i];
      for(;;) {
        if(q == 0) break;
        if(q == '"') { i++; break; };
        lex[j] = q; j++; i++; q = t[i];
      };
      if(j>0) {
        lex[j]=0; aa[k++] = strdup(lex); j = 0;
      };
      continue;
    };
    if(q != ' ') { lex[j] = q; j++; continue; };
    lex[j] = 0; 
    if(j>0) {
      aa[k++] = strdup(lex); j=0;
    };
  };
  if(j>0) {
    lex[j]=0;
    aa[k++] = strdup(lex); j=0;  
  }; 
  goParse(aa,k);
  i = 0;
  while(i<k) delete aa[i++];

}

