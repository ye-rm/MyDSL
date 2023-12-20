// awesomeDSL is a simple DSL, which helps you build your own bot.
// support basic data types, such as integer, boolean, string, array
// have builtin functions, such as len, first, last, rest, push
// support user defined function
// support if-else statement
// support user defined function
// Example for a simple program:
//
//	let hello = puts("hello world");
//	let a = 1;
//	let b = 2;
//	let add = fn(x, y) {
//		x + y;
//	};
//	let result = add(a, b);
//	let tellmeresult= if (result > 0) {
//		puts("result is positive");
//	} else {
//		puts("result is negative");
//	};
package main

import (
	"awesomeDSL/gpt"
	"awesomeDSL/tui"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		_, _ = os.Stderr.WriteString("Usage:./program <filename>\n")
		os.Exit(1)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		_, _ = os.Stderr.WriteString("Open file error\n")
		os.Exit(1)
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	gpt.Init()
	if !tui.LoadScript(file) {
		os.Exit(1)
	}
	tui.GUI()
}
