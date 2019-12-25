package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/jerilseb/monkey-interpreter/lexer"
	"github.com/jerilseb/monkey-interpreter/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)

		if scanned := scanner.Scan(); !scanned {
			return
		}

		line := scanner.Text()
		if line == "exit" {
			return
		}

		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
