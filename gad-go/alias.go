
// alias.go
//
package main


var Alias = make(map[string]string);

func GenAlias( nv int, p *Seq ) {
  var i = 0;
  i += 1; if i >= nv { return; }
  var k = (*p)[i];
  i += 1; if i >= nv { return; }
  var v = (*p)[i];
  Alias[k] = v;

}


