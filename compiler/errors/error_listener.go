package errors

import (
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type SyntaxErrorListener struct {
	*antlr.DefaultErrorListener
	errors []string
}

func NewSyntaxErrorListener() *SyntaxErrorListener {
	return &SyntaxErrorListener{errors: make([]string, 0)}
}

func (e *SyntaxErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	errorMsg := fmt.Sprintf("Error de sintaxis en línea %d:%d - %s", line, column, msg)
	e.errors = append(e.errors, errorMsg)
}

func (e *SyntaxErrorListener) HasErrors() bool {
	return len(e.errors) > 0
}

func (e *SyntaxErrorListener) GetErrors() string {
	return strings.Join(e.errors, "\n")
}

type LexicalErrorListener struct {
	*antlr.DefaultErrorListener
	errors []string
}

func NewLexicalErrorListener() *LexicalErrorListener {
	return &LexicalErrorListener{errors: make([]string, 0)}
}

func (e *LexicalErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	errorMsg := fmt.Sprintf("Error léxico en línea %d:%d - %s", line, column, msg)
	e.errors = append(e.errors, errorMsg)
}

func (e *LexicalErrorListener) HasErrors() bool {
	return len(e.errors) > 0
}

func (e *LexicalErrorListener) GetErrors() string {
	return strings.Join(e.errors, "\n")
}
