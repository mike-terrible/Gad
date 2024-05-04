//
// gad.h !!!!!!!!!!!!!!!!!!!!!!!!!!!1
//
namespace Gad {

  typedef enum {
    ANY,
    COMMENT
  } State;

  static const int GO = 1, RUST = 2, MOJO = 4, PYTHON = 8;

  class MyRT {
  public:
    static const char* ver;  
    FILE* fi;
    FILE* out;
    int ident;
    int gen;
    State st;
    char* curVar;
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
    void goParse(char* [],int);
    
    int cmp(const char*,const char*);
    int cmp(const char*, const char* []);
    
    void goVar(char*,char*,char*);
    char* getV(int,char* [],int);
    char* onType(char*);
    void setIdent(int);
  };

}

