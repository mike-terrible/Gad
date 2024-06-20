#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

int MyRT::goWith(MyRT *rt,char* p[],int nv) {
   return 0;
}

void MyRT::goArray(char* p[],int nv, char* varV, char* vsize, char* vtype, char* be ) {
  if(varV == NULL) return;  
}

void MyRT::genAsmFmt(const char* fmt,char* var) {
  da("\n");
  da(curProc),da("."),da(var),da(".cnv:\n");
  da("  .asciz \""),da(fmt),da("\"\n");
}

void MyRT::goVar(char* var, char* vtype, char* val) {
  if(var == NULL) throw 2; 
  DType dt = UNDEF; da(curProc),da("."),da(var); da(":\n"); 
  if(val != NULL) {
    if(cmp(vtype,Str)) {  da("  .asciz "); da(val); da("\""); genAsmFmt("%s",var),dt = STRING; };
    if(cmp(vtype,Num)) { da("  .quad "); da(val); genAsmFmt("%d",var); dt = NUM;  };
    if(cmp(vtype,Real)) { da("  .double "); da(val); genAsmFmt("%g",var); dt = REAL;  };
    if(cmp(vtype,Light)) {  to(" .quad ");
      if(cmp(val,On)) to("1"); if(cmp(val,Off)) to("0");
      genAsmFmt("%d",var),dt = LIGHT;      
    };
    varNew(var,false,0,dt); return;
  };  
  if(cmp(vtype,Str)) { da("  .space 256,0\n"); genAsmFmt("%s",var); dt = STRING; };  
  if(cmp(vtype,Num)) { da("  .quad 0\n"); genAsmFmt("%d",var); dt = NUM; };  
  if(cmp(vtype,Real)) { da("  .double 0.0\n"); genAsmFmt("%g",var); dt = REAL; };
  if(cmp(vtype,Light)) { da("  .quad 0\n"); genAsmFmt("%d",var); dt = LIGHT; };
  varNew(var,false,0,dt); 
  return;    
}

