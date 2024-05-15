
// lex.rs
//
pub unsafe fn lexer(pp :&str) -> Vec<String> {
  let mut mm :String = pp.to_string();
  let baby = mm.as_mut_vec();
  let mut ret :Vec<String> = Vec::new();
  let n = baby.len();
  let mut i = 0;
  let mut buf :Vec<u8> = Vec::new();
  let mut blank :bool = false;
  let mut quot :bool = false;
  while i < n {
    let mut a = baby[i];
    if a == b'\r' { a = b' '; };
    if a == b'\n' { a = b' '; };
    if a == b'\"' {
      if quot {
        blank = false; quot = false;
      } else {
        buf.push(a);
        quot = true; blank = false; 
      };
      i += 1; continue;
    };
    if a == b' ' {
      if quot {
        buf.push(a);
        i += 1; continue;
      };
      if blank {
        i += 1; continue;
      };
      if buf.len() > 0 {
        blank = false;
        ret.push(String::from_utf8_unchecked(buf));
        buf = Vec::new();
      };
      i += 1; 
      continue;
    };
    buf.push(a);
    i += 1;
  };
  if buf.len() > 0 {
    ret.push(String::from_utf8_unchecked(buf));
  };
  // output
  let np = ret.len();
  println!("!parsed {}",np);
  i = 0;
  while i < np {
    let be = &ret[i]; 
    print!("<! { } !>",be);
    i += 1;
  };
  print!("\n");
  return ret;
}
