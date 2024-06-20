

#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

char* MyRT::onValue(char* xtype) {
  static char vv[128];
  if(cmp(xtype,Str)) {
    strcpy(vv,"\"\""); return vv;
  };
  if(cmp(xtype,Num)) {
    strcpy(vv,"0"); return vv;
  };
  if(cmp(xtype,Real)) {
    strcpy(vv,"0.0"); return vv;
  };
  if(cmp(xtype,Light)) {
    strcpy(vv,"false"); return vv;
  };  
  vv[0]=0;
  return vv;
}

char* MyRT::onType(char* xtype) {
  static char zv[128];
  //const char* Str[] = {"string","строка",nullptr };
  if(cmp(xtype,Str)) {
    if(gen == GO) strcpy(zv,"string");
    if(gen == RUST) strcpy(zv,"&str");
    if(gen == MOJO) strcpy(zv,"String");
    return zv;
  };
  //const char* Num[] =  { "integer","int","число","цел",nullptr };
  if(cmp(xtype, Num )) {
    if(gen == RUST) strcpy(zv,"i64");
    if(gen == GO) strcpy(zv,"int");
    if(gen == MOJO) strcpy(zv,"Int");
    return zv;
  };
  //const char* Real[] = { "real","вещ", nullptr }; 
  if(cmp(xtype, Real)) {
    if(gen == RUST) strcpy(zv,"f64");
    if(gen == GO) strcpy(zv,"float64");
    if(gen == MOJO) strcpy(zv,"Float64");
    return zv;
  };

 //const char* Light[] = { "light","свеча","свет",nullptr };
 if(cmp(xtype, Light )) {
  if((gen == GO)||(gen == RUST)) strcpy(zv,"bool");
  if(gen == MOJO) strcpy(zv,"Bool");
  return zv;
 };
 zv[0]=0;
 return zv;
}

