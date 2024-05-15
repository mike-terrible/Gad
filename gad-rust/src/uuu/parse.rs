
/*   
  parse.rs    
*/
use std::fs;
use std::fs::File;

mod lex;
mod comment;
mod words;
mod amen;
mod ret;
mod vars;
mod when;
mod sic;
mod else_;
mod then_;
mod if_;
mod give;
mod job;
mod show;
mod mess;
mod run;
mod declare;
mod is_;
mod proc;

pub fn parser(fname: &str, gen: &str)  {
  let mut ident :i32 = 0;
  let mut dname :String = "".to_string();
  let mut gname :String = "".to_string();
  gname.push_str("./out");
  match gen {
  "-python" => { gname.push_str(".py");  },
  "-mojo" => { gname.push_str(".mojo");  },
  "-rust" => { gname.push_str(".rs");  },
  "-go" => { gname.push_str(".go");  },
  _ => { gname.push_str(".go"); },
  };
  dname.push_str(fname);
  if fname.ends_with(".гад") {
  } else {
    dname.push_str(".гад");
  };
  vars::set_in_proc(false);
  let fout = File::create(gname).expect("error creating output file!");
  let dd = fs::read_to_string(dname).expect("file not found!");
  println!("parsing { } to { }",fname,gen);
  if gen == "-go" {
    vars::wr(&fout, b"package main\n");
    ident = 2;
  };
  let mut is_comment :bool = false;
  let li = dd.split('\n');
  for t in li {
    println!("{}",t); 
    unsafe {
      let pt = lex::lexer(t);
      if pt.len() == 0 { continue; };
      let s = &pt[0];
      if words::cmp(&s,&words::BEGIN_COMMENT) { is_comment = true; continue;  }; 
      if words::cmp(&s,&words::END_COMMENT) { 
        is_comment = false; vars::wr(&fout,b"\n"); continue; 
      };
      if is_comment { 
        comment::gen_comment(&fout,gen,t);
        continue; 
      };
      if words::cmp(s,&words::LOOP) { ident = amen::gen_loop(&fout,gen,ident); continue; };
      if words::cmp(s,&words::DONE) {
        ident = amen::gen_done(&fout,gen,ident); continue;
      };
      if words::cmp(s,&words::RETURN) {
        ident = ret::gen_return(&fout, gen, ident, pt);
        continue;
      };
      if words::cmp(s, &words::WHEN) {
        ident = when::gen_when(&fout,gen,ident, pt); continue;
      };
      if words::cmp(s, &words::SIC) {
        ident = sic::gen_sic(&fout,gen,ident, pt); continue;
      };
      if words::cmp(s,&words::ELSE) {
        ident = else_::gen_else(&fout,gen,ident, pt); continue;
      };
      if words::cmp(s,&words::THEN) {
        ident = then_::gen_then(&fout,gen,ident, pt); continue; 
      };
      if words::cmp(s,&words::IF) {
        ident = if_::gen_if(&fout,gen,ident, pt); continue;
      };
      if words::cmp(s,&words::GIVE) {
        ident = give::gen_give(&fout,gen,ident, pt); continue;
      };
      if words::cmp(s,&words::JOB) {
        ident = job::gen_job(&fout,gen,ident, pt); continue;
      };
      if words::cmp(s,&words::SHOW) {
        ident = show::gen_show(&fout,gen,ident, pt); continue;
      };
      if words::cmp(s,&words::MESS) {
        ident = mess::gen_mess(&fout,gen,ident, pt); continue;
      };
      if words::cmp(s,&words::RUN) {
        ident = run::gen_run(&fout,gen,ident, pt); continue;
      };
      if words::cmp(s,&words::AMEN) {
        ident = amen::gen_amen(&fout,gen,ident); continue;
      };
      if words::cmp(s,&words::DECLARE) {
        ident = declare::gen_declare(&fout,gen,ident, pt); continue;
      };
      if words::cmp(s,&words::IS) {
        ident = is_::gen_is(&fout,gen,ident, pt); continue;
      };
      if words::cmp(s,&words::PROC) {
        ident = proc::gen_proc(&fout,gen,ident, pt); continue;
      };
    };
    
  };
}
