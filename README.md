Это настоящий язык программирования Gad , который может компилироваться в go или mojo или в python3 или в rust  
Для сборки можно использовать команду g++ ./*.cpp -O3 -s -o ./gad. Правда можно и clang++ ./*.cpp -O3 -s -o ./gad 
Почему Gad? - так python ведь змий. А змий как раз гад ползучий. Внутренний язык программирования Gad отличается от python тем , что количество пробелов от начала каждой строки вообще ни имеют никакого значения. Для так называемыхх ключевых слов используется кирилица. Для идентификаторов кирилица используется если целевым языком является go, rust или python. Но для mojo идентификаторы обязательно должны быть латинскими буквами.
Следует заметить , что gad несколько больше чем препроцессор. На самом деле замышлялся для абстрагирования от некоторых неудобств языка mojo. 
Запускается обычно как gad имяФайла [-go | -mojo | -python | -rust ]  
В зависимости от выбраного целевого языка ( go, mojo , python или rust ) код генерируется в имяФайла.[go | mojo | py | rs]  
Программа на языке гад выглядит примерно вот так:  
`донос`  
`  пример программы на языке гад`   
`зри`    
  
`пусть voldemar как строка будет "дурак"` 

`дело parade для vasya как число для fireplace как свет ход число`   
` пусть vova как строка будет "к торжественному маршу"`   
`  показать для vova для "\n"`   
`  показать для vasya`   
`  показать для " год дали васе\n"`   
`  горит fireplace ли`   
`    грамота "давайте вместе подуем на свечу"`    
`  погасло`    
`    скрижаль "свеча горела на столе свеча горела"`    
`  весть`   
`  ход 1`    
`аминь`    
    
`пора privet поехали`  
`  грамота "бог в помощь!"`   
`  пусть vasya как число будет 1`    
`  пусть fire как свеча будет зажечь`   
`  дать pedro из parade для vasya для fire`     
`  показать для "педро = " для pedro для "\n"`  
`  пусть one как цел будет 1`    
`  когда one < 10 повтор`   
`    показать для one для " "`   
`    вот one = one + 1`    
`  опять`     
`  скрижаль "Конец Работы"`     
`аминь`    
исходники компилятора без поддержки ассемблера находятся в папке no-asm  
исходники комилятора с поддержкой генерацми кода на ассемблере GNU As для архитектуры x86_64 находятся в папке with-asm 
[Реализация компилятора на языке rust](https://github.com/mike-terrible/Gad/tree/main/gad-rust)    

Есть даже кое-какая документаця по языку программирования Гад.  [руководство по программировнию на языке Гад](https://github.com/mike-terrible/Gad/wiki)   
  
Следует иметь ввиду что полная версия компилятора сейчас пока только на языке Go.  Это go.gad версия 2.06 - поддерживает ключи -go -rust -mojo -python -asm -asm32 . Собирается хоть по go build в соответствующей папке , но там есть более комактный вариант сборки по source ./make.src

[полная версия компилятора на языке Go](https://github.com/mike-terrible/Gad/tree/main/gad-go)
