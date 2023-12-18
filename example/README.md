# examples for awesomeDSL

let question_without_space = variable, function call, if-else statement, builtin function

when you type question, the bot will answer you with value at right of =

for example,

with let a = 1, when you type 'a', the bot will answer you with '1'

## example1 

Best bank service bot ever
```
let bankname="bob";
let c_name="jeff";
let cos_level=2;
let hi="hello, welcome to "+bankname;
let howtostealmoney="please go ahead to ask manager";
let hello_msg=fun(x){
    if (x<5){
        "hello, welcome to "+bankname;
    }else{
        "hello, fuck off from "+bankname;
    };
};
let hello=hello_msg(cos_level);
let givemeallmoney="yeah, please wait";
let iamangry=if (cos_level<5){
    let cos_level=10;
}else{
    let cos_level=cos_level;
};
let bye="bye, have a nice day";
```
To build and run this awesome bot
![banker](banker.gif)

## example2
```
let value=100;
let my_list=["hello","hi"];
let card_remaining=1000;
let ok=if (card_remaining>value){
    "yes, please";
}else{
    "no thanks";
};
let bye="bye, see you next time!";
let whatisyourname="my name is bot";
let hi=my_list[0]+" may i help you";
let canyouhelpme=my_list[1]+" just ask me please";
let paymybill=ok;
let payfee=fun(x){
    x>100;
};
let x=payfee(value);
let shouldipayfee=if(x){
    "yes";
}else{
    "no";
};
```
To build this bot in your store
![pay](pay.gif)

Because `OPENAI_FOR_DSL` not set in this 2 examples, if ask undefined questions, the bot will simples reply it do not understand what you say.

If you wanner use openai-api to get better answers, please set `OPENAI_FOR_DSL` to `true` and `OPENAI_API` to in your environment variables.Your api looks like `sk-xxxxxxxxxxx`

See also [openai-api](https://platform.openai.com/docs/quickstart?context=python).
