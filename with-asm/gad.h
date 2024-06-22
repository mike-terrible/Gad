
//
// gad.h !!!!!!!!!!!!!!!!!!!!!!!!!!!1
//
//
//#include <vector>
//using std::vector;
//
namespace Gad {
  typedef enum { ANY, COMMENT } State;

  typedef enum { GO , ASM, RUST , MOJO , PYTHON } CodeGen;
  class MyRT;  
  using Fn = int(*)(MyRT*,char* [],int);  
 
  union SurReal {
    long num;
    double real;
  };
 
  struct It { 
    MyRT* mrt;
    const char** verb;
    It(MyRT*, const char* [], Fn);  
    Fn go;
  };
  
  struct Pair {
    char* k; char* v;
  };
  
  typedef enum {
    UNDEF,LIGHT,NUM,REAL,STRING
  } DType;
  
  struct Var {
    char xname[128]; 
    char pname[128];
    bool isArray; 
    int asize; 
    DType dtype;
  };
  
  extern char result[255];
  extern int zj;
  extern int NVar;
  extern Var Vars[256];
  //
  extern void asmOp2(MyRT*, const char*, char*, char*);
  extern void asmOp2Real(const char*,MyRT*,char*,char*);
  extern void asmCmpsd(const char*,MyRT*,char*,char*);
  extern void asmAllocResult(MyRT*);
  extern void asmAss(MyRT*,char*,char*);
  //	
  extern int goOp2(MyRT*, const char*,int);
  //
  extern long valNum(const char*);
  extern long valReal(const char*);
  extern Var* varGet(const char*);
  extern Var* varNew(char*,bool,int,DType);
  extern void varDump(void); 
  extern DType typeOfLiteral(const char*);
  extern DType asmTypeOf(const char*);

  extern int nA;
  extern Pair AliasTab[256];
  //
  char* aliasV(char*);
  void allocAlias(char*,char*);
  //
  extern char infn[];
  extern char outfn[];
  
  // types 
  extern const char* Str[];  
  extern const char* Num[];  
  extern const char* Real[]; 
  extern const char* Light[]; 
  // commands
  extern const char* Alias[];
  extern const char* Eval[];
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
  extern const char* Mess[];
  extern const char* Give[];
  extern const char* Job[];
  extern const char* Amen[];
  extern const char* Proc[];
  extern const char* Declare[];
  extern const char* Aka[];
  extern const char* Is[];
  extern const char* Run[];
  extern const char* Done[];
  extern const char* Loop[];
  
  struct MyRT {
    static const char* ver;  
    FILE *fi,*out,*data;
    int ident;
    //
    bool inArray;
    int inInit;
    char curProc[256];
    char curVar[256];
    //
    int inProc;
    int inLine;
    CodeGen gen;
    bool ok;
    State st;
    const char* atom;
    char* xmain;

    int fromEvil(char*, int,int,char* []); 

    MyRT(FILE*,char*);

    void onDebug(const char*,const char*);
    void onDebug(const char*);

    int need(const char*);
    int setGen(char*);
    char* seekNotBlank(char*);
    void done();
  
    void da(const char*);
    void to(const char*,const char*);
    void to(const char*);
    void to(int);
  
    void parseIt(char*);
    int goParse(char* [],int);
    
    int cmp(const char*,const char*);
    int cmp(const char*, const char* []);
    int cmpAll(char* [], const char* []);

    void goVar(char*,char*,char*);
    
    void goArray(char* [], int,  char* ,  char*,  char*,  char*  );
    int fromCalc( char*, int,int,char* []);
    void gadError(char* , char* [],int);
    
    static int goAlias(MyRT*,char* [],int);
    static int goWith(MyRT*,char* [],int);
    static int goInit(MyRT*, char* [],int);

    static int goReturn(MyRT*,char* [],int);
    
    
    static int goEval(MyRT*, char* [],int);
    static int goWhen(MyRT*, char* [],int);
    
    static int goSic(MyRT*, char* [],int);
    
    static int goElse(MyRT*, char* [],int);
    
    static int goThen(MyRT*, char* [],int);
    static int goIf(MyRT*, char* [],int);
    static int goGive(MyRT*,char* [],int);
    
    static int goJob(MyRT*,char* [],int);
    
    static int goShow(MyRT*, char* [],int);
    static int goMess(MyRT*, char* [],int);
    
    static int goDeclare(MyRT*, char* [],int);
    
    static int goDone(MyRT*,char* [],int);
    static int goLoop(MyRT*,char* [],int);
    
    static int goAmen(MyRT*,char* [],int);

    static int goIs(MyRT*,char* [],int);    
    //static int goDelo(MyRT*,char* [],int);
    static int goProc(MyRT*,char* [],int);
    
    //static int goPora(MyRT*,char* [],int);
    static int goRun(MyRT*,char* [],int);

    char* getV(int,char* [],int);
    char* onValue(char*);
    char* onType(char*);
    void setIdent(int);

    void genAsmFmt(const char*,char*);

    MyRT* at(const char*);
    MyRT* seek(const char* []);
  
  };
  
  
}

