// Package repl is a read-eval-print-loop for our language.
package repl

import (
	"awesomeDSL/evaluator"
	"awesomeDSL/lexer"
	"awesomeDSL/object"
	"awesomeDSL/parser"
	"bufio"
	"fmt"
	"io"
)

const PROMPT = "➡️ "
const WELCOME = `   __    _    _  ____  ___  _____  __  __  ____  ____   ___  __
  /__\  ( \/\/ )( ___)/ __)(  _  )(  \/  )( ___)(  _ \ / __)(  )
 /(__)\  )    (  )__) \__ \ )(_)(  )    (  )__)  )(_) )\__ \ )(__
(__)(__)(__/\__)(____)(___/(_____)(_/\/\_)(____)(____/ (___/(____)
`

// Start the repl loop
func Start(in io.Reader, out io.Writer) {
	fmt.Fprint(out, WELCOME)
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Woops! Something bad happens!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
