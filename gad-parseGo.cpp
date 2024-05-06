
// gad-parseGo.cpp
//
#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

void MyRT::goParse(char* p[],int nv) {
  //
  vector<const char*> With { "with","для" };
  vector<const char*> Return { "result" , "exit", "ход"  };
  vector<const char*> When { "when", "когда" };
  vector<const char*> Repeat { "repeat","повтор" };
  vector<const char*> Sic {"sic","lay","here","вот" };
  vector<const char*> Else { "?-", "else", "иначе", "погасло" };
  vector<const char*> Then { "?+","then","ли", "тогда" };
  vector<const char*> If { "?!","if","если","горит" };
  vector<const char*> Show { "show", "показать" };
  vector<const char*> Skrepa { "mess","скрепа", "скрижаль", "грамота" };
  vector<const char*> Give { "give", "дать" };
  vector<const char*> Job { "job", "начать"};
  vector<const char*> Amen { "amen", "done", "loop", "аминь", "весть" , "опять" };
  vector<const char*> Delo { "proc","procedure", "десница","дело" };
  vector<const char*> Declare { "dcl","declare","пусть"};
  vector<const char*> Is { "is" , "суть" };
  vector<const char*> Pora {"execute","start","пора" };
  //
  int i = 0;
  while(i<nv) printf(" {%s}",p[i++]);
  printf("\n");
  i = 0;
  char* t = getV(i,p,nv); if(t == NULL) return;
  if(cmp(t,Return)) {
    i++; t = getV(i,p,nv);
    if(t==NULL) {
      to(ident),to("return");
      return;
    };
    to(ident),to("return "),to(t);
    if(gen == RUST) to(";");;
    to("\n");
    return;
  };
  if(cmp(t,When)) {
    to(ident); 
    if(gen == GO) to("for");
    if((gen == RUST) || (gen == MOJO) || ( gen ==  PYTHON)) to("while"); 
    setIdent(ident+2);
    for(;;) {
      i++; if(i>=nv) break;
      t = getV(i,p,nv); if(t == NULL) return;
      if(cmp(t,Repeat)) {
        if((gen == GO) || (gen == RUST)) to(" {\n");
        if((gen == MOJO) || (gen == PYTHON)) to(" :\n");
        return;
      };
      to(" ");
      to(t); if(t[0]=='"') to("\""); 
    };
    return;
  };
  
  if(cmp(t,Sic)) {
    to(ident);
    for(;;) {
      i++; 
      t = getV(i,p,nv); 
      if(t == NULL) { 
        if(gen == RUST) to(";");
        to("\n"); 
        return; 
      };
      to(t); if(t[0]=='"') to("\""); to(" "); 
    };
    return;
  };
  if(cmp(t,Else)) {
    setIdent(ident-2);
    if((gen == GO) || (gen == RUST)) to(ident),to("} else {\n"); 
    if((gen == MOJO) || (gen == PYTHON)) to(ident),to("else:\n"); 
    setIdent(ident+2);
    return;
  };
  
  if(cmp(t,Then)) {
    if((gen == GO) || (gen == RUST)) to(" {\n");
    if((gen ==  MOJO) || (gen == PYTHON)) to(" :\n"); 
    setIdent(ident+2);
    return;
  };
  
  if(cmp(t,If)) {
    i++; t = getV(i,p,nv); if(t == NULL) return;
    to(ident);
    to("if "),to(t);
    while( i < nv ) {
      i++; t = getV(i,p,nv); if(t == NULL) return;
      if(cmp(t,Then)) {
        if((gen == RUST) || (gen == GO)) to(" {\n");
        if((gen == MOJO) || (gen == PYTHON)) to(" :\n"); 
        setIdent(ident+2);
        return;
      }; 
      to(" "),to(t);
    };
    return;
  };
  
  if(cmp(t,Give)) {
    i++;
    t = getV(i,p,nv); if(t==NULL) return;
    to(ident);
    if((gen == GO) || ( gen == MOJO)) to("var "); 
    if(gen == RUST) to("let mut "); 
    to(t); to(" = ");
    i++;
    t = getV(i,p,nv);
    /*
    if(t!=NULL) {
      if(cmp(t,"из")) { };
    };
    */
    i++;
    t = getV(i,p,nv);
    to(t),to("(");
    int np = 0;
    for(;;) {
      i++; if(i>=nv) break;
      t = getV(i,p,nv); if(t == NULL) break;
      if(cmp(t,With)) {
        i++; t = getV(i,p,nv); if(t == NULL) break;
        np++;
        if(np>1) to(",");
        to(t); if(t[0]=='"') to("\"");
      };
    };
    if(gen == RUST) to(");\n"); else to(")\n"); 
    return;
  };
  
  if(cmp(t,Job)) {
    i++;
    t = getV(i,p,nv);
    to("\n"),to(ident),to(t),to("(");
    int np = 0;
    for(;;) {
      i++; if(i>=nv) break;
      t = getV(i,p,nv); if(t == NULL) break;
      if(cmp(t,With)) {
        i++; t = getV(i,p,nv); if(t == NULL) break;
        np++;
        if(np>1) to(",");
        to(t); if(t[0]=='"') to("\"");
      };
    };
    if(gen == RUST) to(");\n");
    else to(")\n");
    return;
  };
  
  if(cmp(t,Show)) {
    int np = 0;
    while(++i < nv) {
      t = getV(i,p,nv); if(t == NULL) break;
      if(cmp(t,With)) {
        i++;
        t = getV(i,p,nv); if(t == NULL) break;
        np++;
        if(gen == RUST) {
          to(ident),to("print!(\"{ } \",");
          to(t); if(t[0]=='"') to("\""); to(");\n");
          continue;
        };
        if(gen == GO) {
          to(ident),to("print(");
          to(t); if(t[0]=='"') to("\""); to(",\" \");\n");
          continue;
        };
        if((gen == PYTHON)||(gen == MOJO)) {
          to(ident),to("print(");
          to(t); if(t[0]=='"') to("\"");
          to(",end =\" \")\n");
        }; 
      };
    };
    return;
  };
  
  if(cmp(t, Skrepa)) {
    i++;
    t = getV(i,p,nv); if(t == NULL) return;
    to(ident); 
    if(gen == RUST) to("println!("),to(t),to("\");\n"); 
    if(gen == GO) to("println("),to(t),to("\")\n"); 
    if((gen == MOJO) || (gen == PYTHON)) to("print("),to(t),to("\")\n"); 
    return;
  };
  
  if(cmp(t,Pora)) {
    i++; 
    xmain = getV(i,p,nv); if(xmain == NULL) return;
    i++;
    t = getV(i,p,nv); 
    if(t == NULL) {
      to(ident);
      if(gen == RUST) to("fn "),to(xmain),to("() {\n"); 
      if(gen == GO) to("func "),to(xmain),to("() {\n"); 
      if(gen == MOJO) to("fn "),to(xmain),to("() raises :\n"); 
      if(gen == PYTHON) to("def "),to(xmain),to("() :\n"); 
      setIdent(ident+2);
    } else {
      to(ident);
      if(gen == RUST) to("fn main() {\n"); 
      if(gen == GO) to("func main() {\n"); 
      if(gen == MOJO) to("fn main() raises :\n"); 
      if(gen == PYTHON) to("def main() :\n"); 
      setIdent(ident+2),to(ident),to(xmain),to("()\n");
      setIdent(ident-2),to(ident);
      if((gen == GO)||(gen == RUST)) to("}\n"); 
      to(ident);
      if(gen == RUST) to("fn "),to(xmain),to("() {\n"); 
      if(gen == GO) to("func "),to(xmain),to("() {\n"); 
      if(gen == MOJO) to("fn "),to(xmain),to("() raises :\n"); 
      if(gen == PYTHON) to("def "),to(xmain),to("() :\n"); 
      setIdent(ident+2);
    };
    return;
  };
  
  if(cmp(t,Amen)) {
    to("\n");
    setIdent(ident-2),to(ident);
    if((gen == RUST) || (gen == GO)) to("}\n"); 
    if((gen == MOJO) || (gen == PYTHON)) to("pass\n"); 
    return;
  };
  
  if(cmp(t,Declare)) {
    i++; char* var = getV(i,p,nv); 
    i++; char* like = getV(i,p,nv); 
    i++; char* vtype = getV(i,p,nv);
    i++; char* be = getV(i,p,nv);
    i++; char* val = getV(i,p,nv);
    goVar(var,vtype,val); return;
  };
  
  if(cmp(t,Is)) {
    if((gen == GO) || (gen == RUST)) to("{\n"); 
    setIdent(ident+2); 
    return;
  };
  if(cmp(t,Delo)) {
    to(ident);
    if(gen == GO) to("func "); if(gen == PYTHON) to("def "); 
    if((gen == MOJO) || (gen == RUST)) to("fn "); 
    i++;
    char* xn = getV(i,p,nv); if(xn == NULL) return;
    to(xn),to("(");
    int narg = 0;
    for(;;) { i++; if(i>=nv) { to(") "); break; };
      char* itIs = getV(i,p,nv);
      if(cmp(itIs,"ход")) {
        i++; char *act = getV(i,p,nv);
        if(act != NULL) {
          char* ztype = onType(act);
          if(gen == PYTHON) { to(") :\n"); setIdent(ident + 2); return; };
          int nz = strlen(ztype);
          if(nz > 0) {
            if(gen == GO) { to(") "),to(ztype),to(" {\n"); setIdent(ident + 2); return; };
            if(gen == MOJO) { to(") -> "),to(ztype),to(" :\n"); setIdent(ident + 2); return; };
            if(gen == RUST) { to(") -> "),to(ztype),to(" {\n"),setIdent(ident + 2); return; };
          };
        };
      };
      if(cmp(itIs,"суть")) {
        if((gen == RUST) || (gen == GO))  to(") {\n"); 
        if((gen == MOJO) || (gen == PYTHON)) to(") :\n"); 
        setIdent(ident+2); return;
      };
      if(cmp(itIs,"для")) { narg ++;
        i++; char* var = getV(i,p,nv);
        i++; if(narg>1) to(",");
        to(var);
        char* like = getV(i,p,nv); 
        if(like != NULL) if(cmp(like,"как")) {
          i++; char* xtype = getV(i,p,nv);
          if(xtype != NULL) {
            if(gen == GO) to(" "),to(onType(xtype)); 
            if((gen == MOJO) || (gen == RUST)) to(" :"),to(onType(xtype)); 
          };  
        };
        continue;
      };
    };
  };
}

