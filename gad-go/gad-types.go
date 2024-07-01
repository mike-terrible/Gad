
// gad-types.go

package main

const (
  DTYPE_UNDEF = iota 
  DTYPE_LIGHT = iota
  DTYPE_NUM = iota
  DTYPE_REAL = iota
  DTYPE_STRING = iota
)

type Seq []string

type Var struct {
  xname string; 
  pname string;
  isArray bool; 
  asize int; 
  dtype int;
}

