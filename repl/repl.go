package repl

import (
	"bufio"
	"fmt"
	"io"

	"vichr.me/go/monkey-lang-interpreter/evaluator"
	"vichr.me/go/monkey-lang-interpreter/lexer"
	"vichr.me/go/monkey-lang-interpreter/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf("%s", PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.NewLexer(line)
		p := parser.NewParser(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

		// this is just for testing the parser, we will remove it later when we implement the evaluator
		// io.WriteString(out, program.String())
		// io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []error) {
	for _, err := range errors {
		io.WriteString(out, "\t"+err.Error()+"\n")
	}
}
