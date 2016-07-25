//go:generate go tool yacc parse.y

package main

import (
	"os"
)

func main() {
	yyErrorVerbose = true
	lexer := NewLexer(os.Stdin)
	yyParse(lexer)
	ast := lexer.program
	RunStmtBlock(ast)
}
