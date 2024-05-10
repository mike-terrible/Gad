#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include "gad.h"

namespace Gad { 
  const char* Declare[] = { "dcl","declare","пусть","да", "хай" , "ну" , nullptr };
  const char* Aka[] = { ":","as","aka","как","пока","таке", "це" , nullptr };
  
  const char* If[] = { "?!","if","если","горит", "чи", nullptr };
  const char* Then[] = { "?+","then","ли", "тогда", "є" , nullptr };
  const char* Else[] = { "?-", "else", "иначе", "погасло", "нема" ,nullptr };
  const char* Done[] = { "?!)","done", "весть", "взгляд", "погляд" , nullptr };
  
  const char* When[] = { "when", "когда" , "коли" , nullptr };
  const char* Repeat[] = { "repeat","повтор", nullptr };
  const char* Loop[] = { "loop", "опять","далі", nullptr };
  
  const char* Job[] = { "@@","job", "call", "начать", "почати" ,nullptr };
  const char* Give[] = { "give", "!!!", "дать", "дати", "кошик" , nullptr };
  const char* Skrepa[] = { "mess","скрепа", "скрижаль", "грамота", "новина" , nullptr };
  const char* Pora[] = {"???","execute","exec","start","run","пора","час", nullptr};
  const char* Return[] = { "result" , "exit", "ход", "дать", "здобич" , nullptr  };
  const char* Delo[] = { "proc","procedure", "десница","дело","справа",nullptr };
  const char* Show[] = { "show", "показать", "вистава" , nullptr };
  const char* Sic[] = {"sic","lay","here","вот", "вже",  nullptr };
  const char* Is[] = { "is" , "суть", "докладно", nullptr };
  const char* Amen[] = { "amen", "done", "аминь", "все", nullptr };
  const char* With[] = { "with","для", "по", "ще" , nullptr };
  
  const char* On[] = { "on","зажечь", "запалити" ,nullptr};
  const char* Off[] = { "off", "потушить", "загасити" , nullptr };

  const char* Comment[] = { "исполать","донос","(*" , nullptr };
  const char* EndComment[] = { "зри","*)" , nullptr };
  
   
}
