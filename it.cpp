
// it.cpp
//
#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

using namespace Gad;

It::It(MyRT* ru,const char* v[],Fn fn){
  go = fn;
  verb = v;
  mrt = ru;
}
/*
int MyRT::GoReturn::go(char* u[],int n) {  return self->goReturn(u,n); }
int MyRT::GoWhen::go(char* u[],int n) {  return self->goWhen(u,n); }
int MyRT::GoSic::go(char* u[],int n) {  return self->goSic(u,n); }
int MyRT::GoElse::go(char* u[],int n) {  return self->goElse(u,n); }
int MyRT::GoThen::go(char* u[],int n) {  return self->goThen(u,n); }
int MyRT::GoIf::go(char* u[],int n) {  return self->goIf(u,n); }
int MyRT::GoGive::go(char* u[],int n) {  return self->goGive(u,n); }
int MyRT::GoJob::go(char* u[],int n) {  return self->goJob(u,n); }
int MyRT::GoShow::go(char* u[],int n) {  return self->goShow(u,n); }
int MyRT::GoSkrepa::go(char* u[],int n) {  return self->goSkrepa(u,n); }
int MyRT::GoPora::go(char* u[],int n) {  return self->goPora(u,n); }
int MyRT::GoAmen::go(char* u[],int n) {  return self->goAmen(u,n); }
int MyRT::GoDeclare::go(char* u[],int n) {  return self->goDeclare(u,n); }
int MyRT::GoIs::go(char* u[],int n) {  return self->goIs(u,n); }
int MyRT::GoDelo::go(char* u[],int n) {  return self->goDelo(u,n); }
*/

