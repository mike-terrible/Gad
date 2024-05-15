//
// words.rs
//
pub static WITH :[&str;4] = [ "with","для", "по", "ще" ]; 
pub static BEGIN_COMMENT :[&str;3] = ["(*","донос","исполать"];
pub static END_COMMENT :[&str;2] = [ "*)", "зри" ]; 
pub static LOOP :[&str;3] = ["loop", "опять","далі"];
pub static DONE :[&str;6] = ["?!)","done", "весть", "авось", "невже","погляд"];
pub static RETURN :[&str;5] = ["result" , "exit", "ход", "дать", "здобич"];
pub static WHEN :[&str;3] = ["when", "когда" , "коли"];
pub static SIC :[&str;5] = ["sic","lay","here","вот", "вже"];
pub static ELSE :[&str;6] = ["?-", "else", "иначе", "погасло", "нема", "або" ];
pub static THEN :[&str;6] = ["?+","then","ли", "тогда", "так" ,"є" ];
pub static IF :[&str;5] = ["?!","if","если","горит", "чи"];
pub static GIVE :[&str;5] = ["give", "!!!", "дать", "дати", "кошик"];
pub static JOB :[&str;5] = ["@@","job", "call", "начать", "почати"];
pub static SHOW :[&str;3] = ["show", "показать", "вистава"];
pub static MESS :[&str;5] = [ "mess", "скрепа", "скрижаль", "грамота", "новина" ];
pub static RUN :[&str;7] = [ "???","execute","exec","start","run","пора","час" ];
pub static AMEN :[&str;4] = ["amen","end","аминь","все" ];
pub static DECLARE :[&str;6] = [ "dcl","declare","пусть","да", "хай" , "ну" ];
//
pub static ON :[&str;3] = [ "on","зажечь", "запалити" ];
pub static OFF :[&str;3] = [ "off", "потушить", "загасити" ];
pub static STR :[&str;3] = [ "string","строка", "текст" ];
pub static NUM :[&str;5] = [ "integer","int","num","число","цел" ];
pub static REAL :[&str;3] = [ "real","вещ", "дій" ];
pub static LIGHT :[&str;5] = [ "light","свеча","свет","свічка", "світло" ];
pub static IS :[&str;4] = [ "is" , "суть", "докладно", "саме"  ];
pub static AKA :[&str;7 ] = [ ":","as","aka", "как", "пока" ,"таке", "це" ]; 
//
pub static PROC :[&str;5] = [ "proc","procedure", "десница","дело","справа" ];

pub fn cmp(a :&str, pp :&[&str]) -> bool {
  for t in pp {
    if a.starts_with(t) {  return true; };
  };
  return false;
}



