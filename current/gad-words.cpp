#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

namespace Gad { 
  const char* Alias[] = { "\%\%\%",nullptr };
  const char* Eval[] =  { ":=", "<-", "eval", "evil", "cal", "из", "рах", nullptr };
  const char* Array[] = { "[]","array" , "ларь", "кошик", nullptr }; 
  const char* Init[] =  { "<-","init", "так", nullptr  };
  const char* With[] =  { "?","with","для", "по", "ще", nullptr };
  const char* Comment[] =  { "(*","донос","исполать", nullptr };
  const char* EndComment[] = { "*)", "зри", nullptr };
  const char* Loop[] = { "@/","loop", "опять","далі", nullptr };
  const char* Done[] = { "?!)","done", "весть", "авось", "невже","погляд",nullptr };
  const char* Return[] =  { "result" , "exit", "ход", "дать", "здобич", nullptr };
  const char* When[] = { "|","when", "while" ,"когда" , "коли", nullptr };
  const char* Sic[] = { "^^","sic","lay","here","вот", "вже", nullptr };
  const char* Else[] = { "?-", "else", "иначе", "погасло", "нема", "або", nullptr };
  const char* Then[] = { "?+", "then","ли" ,"тогда", "так" ,"є" , nullptr};
  const char* If[] = { "?!","if","если","горит", "чи", nullptr };
  const char* Give[] = { "give", "?=","!!!", "дай", "дати" , nullptr };
  const char* Job[] = { "@@","job", "call", "начать", "почати", nullptr };
  const char* Show[] ={ "!","show", "показать", "вистава",nullptr };
  const char* Mess[] = { "!!","mess", "скрепа", "скрижаль", "грамота", "новина",nullptr };
  const char* Run[] =  { "???","execute","exec","start","run","пора","час", nullptr };
  const char* Amen[] = { "amen","end","аминь","все", nullptr };
  const char* Declare[] = { "dcl","declare","пусть","да", "хай" , "ну",nullptr };
  const char* Proc[] =  { "proc","procedure", "десница","дело","справа", nullptr };
  const char* Is[] = { "is" , "суть", "докладно", "саме",nullptr  };
  const char* Aka[] = { ":","as","aka", "как", "пока" ,"таке", "це",nullptr }; 
  const char* Repeat[] = { "/@","repeat", "повтор", nullptr };
  const char* On[] = { "on","зажечь", "запалити", nullptr };
  const char* Off[] = { "off", "потушить", "загасити",nullptr };
  
}
