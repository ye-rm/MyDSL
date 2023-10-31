# 程序设计实践-#DSL

## BNF定义

```bash
type
	string,int
variables:
	let x = 0;
	
operators:
	+,-,=,*,%,/
state:
	state <state_name>{
		do something
	}
match:
	match <variable>{
	value_1=>state_1;
	value_2=>state_2;
	}
catch:
	catch <variable>;
respond:
	respond <variable>;
goto:
	goto <state_name>;
have:
	have "<string>"
```



### example

```bash
let cash = 0;

state hello{
	respond hello;
	goto tasks;
}

state tasks{
	let input;
	catch input;
	match input{
		have "lookup" => lookup;
		have "save"=> save_money;
		have "draw"=> draw_money;
		have "exit"=> exit;
	}
}

state lookup{
	#send lookup query and respond
	#respond amount;
}

state save_money{
	catch input
	cash=cash+input;
}

fun add(x,y){return 2+3};
let x=add(2,3);
# others
```
