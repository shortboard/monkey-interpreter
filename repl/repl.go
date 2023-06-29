package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/shortboard/monkey-interpreter/lexer"
	"github.com/shortboard/monkey-interpreter/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		// io.Writer.Write() takes a byte slice, so we need to convert the string
		// to a byte slice
		line := scanner.Text()
		l := lexer.New(line)

		// We're going to print out all the tokens the lexer gives us until we get to the end of the input
		// (i.e. when the lexer returns token.EOF)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
