// awesomeDSL is a simple programming language, which helps you build your own DSL.
// support basic data types, such as integer, boolean, string, array
// have builtin functions, such as len, first, last, rest, push
// support user defined function
// support if-else statement
package main

import (
	"awesomeDSL/gpt"
	"awesomeDSL/tui"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Stderr.WriteString("Usage:./program <filename>\n")
		os.Exit(1)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		os.Stderr.WriteString("Open file error\n")
		os.Exit(1)
	}

	defer file.Close()
	gpt.Init()
	tui.LoadScript(file)
	tui.GUI()
}
