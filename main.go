//go:generate go tool yacc parse.y

package main

import (
	"os"
)

func main() {
	yyErrorVerbose = true
	lexer := NewLexer(os.Stdin)
	yyParse(lexer)
	prog := lexer.Program()
	if prog == nil {
		os.Exit(1)
	}
	RunStmtBlock(prog)
}
