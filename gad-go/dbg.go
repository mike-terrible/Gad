
// dbg.go
//
package main

var Dbg = false

func OnDebugNoNL(s string) {
  if !Dbg { return; }
  To(GetIdent());
  switch Mode {
  case RUST,GO: { Wr("//!! ");  }
  default: { Wr("#!! "); }
  };
  Wr(s);
}

func OnDebug(s string) {
  if !Dbg { return; }
  OnDebugNoNL(s);
  Wr("\n");
}


func DbgTrace(s string) {
  if !Dbg { return; };
  To(GetIdent());
  switch Mode {
  case RUST,GO: { Wr("//!! " + s); }
  default: { Wr("#!! " +  s); }
  };
}


