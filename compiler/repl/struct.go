package repl

import (
	// parser "compiler/parser"
	compiler "main/parser"

	"github.com/antlr4-go/antlr/v4"
)

type Struct struct {
	Name string
	// Fields []parser.IStruct_propContext
	Fields []compiler.IStruct_propContext
	Token  antlr.Token
}
