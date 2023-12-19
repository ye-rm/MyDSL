# Documentation.

## usage and syntax of the awesomeDSL.

`let question_without_space` = variable, function call, if-else statement, builtin function

when you type question, the bot will answer you with value at right of `=`

for example, 

with `let a = 1`, when you type 'a', the bot will answer you with '1'

```
//builtin function
let hello = puts("hello world");
//variables
let a = 1;
let b = 2;
let c = true;
let d = "hello world";
//define function 
let add = fn(x, y) {
	x + y;
};
//call function
let result = add(a, b);
//if-else statement
let tellmeresult= if (result > 0) {
	"result is positive";
} else {
	"result is negative";
};
//array
let arr = [1, 2, 3];
//array builtin function
let first = first(arr);
let last = last(arr);
let rest = rest(arr);
//array push
let arr2 = push(arr, 4);
//array len
let len = len(arr2);
//array index
let index = arr2[0];
//string concat
let len = "hello"+" world";
```

## Three main packages of the awesomeDSL

for doc in package, this only show the exported functions and types.
if you want to see the details, please check the source code and comments.

[lexer](lexer.md) package laxer is responsible for the lexical analysis of the input file.

[parser](parser.md) parser takes the tokens from the lexer and builds an AST.

[evaluator](evaluator.md) evaluator takes the AST and evaluates it.

## other packages
[token](token.md) package token defines the token type.

[ast](ast.md) package ast defines the node in abstract syntax tree.

[object](object.md) evaluator views everything as an object.

[tui](tui.md) package tui defines the text user interface.

[gpt](gpt.md) package gpt suppport the genetic openai-api.

## unused packages
[repl](repl.md) package repl defines the read-eval-print-loop. I use tui instead of repl. But keep it as an example.

## example
check the [example folder](../example/) for the example of the awesomeDSL.

## how to use
### build
make sure you have go installed. Then run
`go build main.go`

### run
`./main yourscript.txt`

### test
`go test ./lexer ./parser ./evaluator`

### godoc locally
make sure you have `godoc` (not `go doc`) installed. Then run
`godoc -http=:6060`

## enable OpenAI API
1. get your api key from [openai](https://platform.openai.com/docs/quickstart?context=python). Your api should be in the format of `sk-xxxxxx`.
2. set environment variable `OPENAI_API_KEY` to your api key.
3. set environment variable `OPENAI_FOR_DSL` to `true`.
see also end of [example](../example/).