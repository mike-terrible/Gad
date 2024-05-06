

#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;


char* MyRT::onType(char* xtype) {
  static char zv[128];
  const char* Str[] = {"string","строка",nullptr };
  if(cmp(xtype,Str)) {
    if(gen == GO) strcpy(zv,"string");
    if(gen == RUST) strcpy(zv,"String");
    if(gen == MOJO) strcpy(zv,"String");
    return zv;
  };
  const char* Num[] =  { "integer","число","цел",nullptr };
  if(cmp(xtype, Num )) {
    if(gen == RUST) strcpy(zv,"i64");
    if(gen == GO) strcpy(zv,"int");
    if(gen == MOJO) strcpy(zv,"Int");
    return zv;
  };
  const char* Real[] = { "real","вещ", nullptr }; 
  if(cmp(xtype, Real)) {
    if(gen == RUST) strcpy(zv,"f64");
    if(gen == GO) strcpy(zv,"double");
    if(gen == MOJO) strcpy(zv,"Float64");
    return zv;
  };

 const char* Light[] = { "light","свеча","свет",nullptr };
 if(cmp(xtype, Light )) {
  if((gen == GO)||(gen == RUST)) strcpy(zv,"bool");
  if(gen == MOJO) strcpy(zv,"Bool");
  return zv;
 };
 zv[0]=0;
 return zv;
}

