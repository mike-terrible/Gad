package main

func Wr(b string) {
  Out.WriteString(b);
}

func To(n int ) {
  var i = 0;
  for i < n { Wr(" "); i += 1; }
}

