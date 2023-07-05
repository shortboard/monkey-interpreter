package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/shortboard/monkey-interpreter/evaluator"
	"github.com/shortboard/monkey-interpreter/lexer"
	"github.com/shortboard/monkey-interpreter/object"
	"github.com/shortboard/monkey-interpreter/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
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
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
