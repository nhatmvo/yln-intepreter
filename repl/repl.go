package repl

import (
	"bufio"
	"fmt"
	"io"
	"yln/lexer"
	"yln/parser"
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

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			io.WriteString(out, "Oops! Something went wrong when parsing YLN program!")
			io.WriteString(out, "\n")
			for _, msg := range p.Errors() {
				io.WriteString(out, "\t"+msg+"\n")
			}
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}
