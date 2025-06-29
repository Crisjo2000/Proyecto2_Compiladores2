package repl

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

const (
	LexicalError  = "Error léxico"
	SyntaxError   = "Error sintáctico"
	SemanticError = "Error semántico" // <-- Analisis Semantico
	// Errores en tiempo de ejecucion
	RuntimeError = "Error en tiempo de ejecución" // <-- Ejecucion
)

// Definimos una tabla de errores
type Error struct {
	Line   int
	Column int
	Msg    string
	Type   string
}

// La tabla de Err
type ErrorTable struct {
	Errors []Error
}

func (et *ErrorTable) AddError(line int, column int, msg string, errorType string) {
	et.Errors = append(et.Errors, Error{line, column, msg, errorType})
}

func (et *ErrorTable) NewLexicalError(line int, column int, msg string) {
	et.AddError(line, column, msg, LexicalError)
}

func (et *ErrorTable) NewSyntaxError(line int, column int, msg string) {
	et.AddError(line, column, msg, SyntaxError)
}

func (et *ErrorTable) NewSemanticError(token antlr.Token, msg string) {
	et.AddError(token.GetLine(), token.GetColumn(), msg, SemanticError)
}

func (et *ErrorTable) NewRuntimeError(line int, column int, msg string) {
	et.AddError(line, column, msg, RuntimeError)
}

func NewErrorTable() *ErrorTable {
	return &ErrorTable{}
}

func (et *ErrorTable) GenerateHTMLReport() string {
	if len(et.Errors) == 0 {
		return `
        <!DOCTYPE html>
        <html>
        <head>
            <title>Reporte de Errores - VLang</title>
            <style>
                body { font-family: Arial, sans-serif; margin: 20px; }
                .success { color: green; padding: 20px; font-size: 18px; }
            </style>
        </head>
        <body>
            <div class="success">✅ No se encontraron errores</div>
        </body>
        </html>
        `
	}

	html := `
    <!DOCTYPE html>
    <html>
    <head>
        <title>Reporte de Errores - VLang</title>
        <style>
            body { font-family: Arial, sans-serif; margin: 20px; }
            h1 { color: #e74c3c; text-align: center; }
            table { width: 100%; border-collapse: collapse; margin-top: 20px; }
            th, td { border: 1px solid #ddd; padding: 8px; text-align: left; }
            th { background-color: #f2f2f2; }
            .lexical { background-color: #ffebee; }
            .syntax { background-color: #fff8e1; }
            .semantic { background-color: #e3f2fd; }
            .runtime { background-color: #fce4ec; }
            .error-count { margin: 20px 0; font-weight: bold; }
        </style>
    </head>
    <body>
        <h1>Reporte de Errores</h1>
        <div class="error-count">Total de errores: ` + fmt.Sprintf("%d", len(et.Errors)) + `</div>
        <table>
            <tr>
                <th>Tipo</th>
                <th>Línea</th>
                <th>Columna</th>
                <th>Descripción</th>
            </tr>
    `

	for _, err := range et.Errors {
		rowClass := ""
		switch err.Type {
		case LexicalError:
			rowClass = "lexical"
		case SyntaxError:
			rowClass = "syntax"
		case SemanticError:
			rowClass = "semantic"
		case RuntimeError:
			rowClass = "runtime"
		}

		html += fmt.Sprintf(`
            <tr class="%s">
                <td>%s</td>
                <td>%d</td>
                <td>%d</td>
                <td>%s</td>
            </tr>
        `, rowClass, err.Type, err.Line, err.Column, err.Msg)
	}

	html += `
        </table>
    </body>
    </html>
    `

	return html
}
