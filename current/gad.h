
//
// gad.h !!!!!!!!!!!!!!!!!!!!!!!!!!!1
//
//
//#include <vector>
//using std::vector;
//
namespace Gad {
  typedef enum { ANY, COMMENT } State;

  typedef enum { GO , RUST , MOJO , PYTHON } CodeGen;
  class MyRT;  
  using Fn = int(*)(MyRT*,char* [],int);  
 
  struct It { 
    MyRT* mrt;
    const char** verb;
    It(MyRT*, const char* [], Fn);  
    Fn go;
  };
  // types 
  extern const char* Str[];  
  extern const char* Num[];  
  extern const char* Real[]; 
  extern const char* Light[]; 
  // commands
  extern const char* Init[];
  extern const char* Array[];
  extern const char* Comment[]; 
  extern const char* EndComment[]; 
  extern const char* On[];
  extern const char* Off[];
  extern const char* With[];
  extern const char* Return[];
  extern const char* When[];
  extern const char* Repeat[];
  extern const char* Sic[];
  extern const char* Else[];
  extern const char* Then[];
  extern const char* If[];
  extern const char* Show[];
  extern const char* Skrepa[];
  extern const char* Give[];
  extern const char* Job[];
  extern const char* Amen[];
  extern const char* Delo[];
  extern const char* Declare[];
  extern const char* Aka[];
  extern const char* Is[];
  extern const char* Pora[];
  extern const char* Done[];
  extern const char* Loop[];

  struct MyRT {
    static const char* ver;  
    FILE *fi,*out;
    int ident;
    //
    bool inArray;
    int inInit;
    char curVar[256];
    //
    int inProc;
    CodeGen gen;
    bool ok;
    State st;
    const char* atom;
    char* fname;
    char* xmain;
    MyRT(FILE*);

    int need(const char*);
    int setGen(char*);
    char* seekNotBlank(char*);
    void done();
  
    void to(const char*,const char*);
    void to(const char*);
    void to(int);
  
    void parseIt(char*);
    int goParse(char* [],int);
    
    int cmp(const char*,const char*);
    int cmp(const char*, const char* []);

    void goVar(char*,char*,char*);
    
    void goArray(char* [], int,  char* ,  char*,  char*,  char*  );
    static int goWith(MyRT*,char* [],int);
    static int goInit(MyRT*, char* [],int);

    static int goReturn(MyRT*,char* [],int);
    
    static int goWhen(MyRT*, char* [],int);
    
    static int goSic(MyRT*, char* [],int);
    
    static int goElse(MyRT*, char* [],int);
    
    static int goThen(MyRT*, char* [],int);
    static int goIf(MyRT*, char* [],int);
    static int goGive(MyRT*,char* [],int);
    
    static int goJob(MyRT*,char* [],int);
    
    static int goShow(MyRT*, char* [],int);
    static int goSkrepa(MyRT*, char* [],int);
    
    static int goDeclare(MyRT*, char* [],int);
    
    static int goDone(MyRT*,char* [],int);
    static int goLoop(MyRT*,char* [],int);
    
    static int goAmen(MyRT*,char* [],int);

    static int goIs(MyRT*,char* [],int);    
    static int goDelo(MyRT*,char* [],int);
    static int goPora(MyRT*,char* [],int);

    char* getV(int,char* [],int);
    char* onValue(char*);
    char* onType(char*);
    void setIdent(int);

    MyRT* at(const char*);
    MyRT* seek(const char* []);
  
  };
  
  
}

