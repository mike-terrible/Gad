// 
// job.rs
//
package main

import "strings"

func Asm32GenJob(nv int, p *Seq ) {

}

func AsmGenJob(nv int, p *Seq ) {

}

func GenJob(nv int, p *Seq ) {
  if Mode == ASM32 { Asm32GenJob(nv, p ); return; }
  if Mode == ASM { AsmGenJob(nv, p ); return; }
  var i = 0;
  i += 1; if i >= nv { return; };
  var t = (*p)[i];
  To(GetIdent()); Wr(t, "(");
  var np = 0; 
  for { i += 1; if i >= nv { break; }; t = (*p)[i];
    if Cmp(t,WITH) { i += 1; if i >= nv { break; }; 
      t = (*p)[i]; np += 1; if np > 1 { Wr(","); };
      Wr(t);
      if strings.HasPrefix(t,"\"") { Wr("\""); };
    }; 
  };
  if Mode == RUST { Wr(");\n"); } else { Wr(")\n"); };
}


