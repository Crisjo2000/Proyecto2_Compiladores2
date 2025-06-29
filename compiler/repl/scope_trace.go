package repl

import (
	"bytes"
	"fmt"
	"main/value"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type ScopeTrace struct {
	GlobalScope  *BaseScope
	CurrentScope *BaseScope

	Reporte string
}

func (s *ScopeTrace) PushScope(name string) *BaseScope {

	newScope := NewLocalScope(name)
	s.CurrentScope.AddChild(newScope)
	s.CurrentScope = newScope

	return s.CurrentScope
}

func (s *ScopeTrace) PopScope() {
	s.CurrentScope = s.CurrentScope.Parent()
}

func (s *ScopeTrace) Reset() {
	s.CurrentScope = s.GlobalScope
}

func (s *ScopeTrace) AddVariable(name string, varType string, value value.IVOR, isConst bool, allowNil bool, token antlr.Token) (*Variable, string) {
	return s.CurrentScope.AddVariable(name, varType, value, isConst, allowNil, token)
}

func (s *ScopeTrace) GetVariable(name string) *Variable {
	return s.CurrentScope.GetVariable(name)
}

func (s *ScopeTrace) AddFunction(name string, function value.IVOR) (bool, string) {
	return s.CurrentScope.AddFunction(name, function)
}

func (s *ScopeTrace) GetFunction(name string) (value.IVOR, string) {
	return s.CurrentScope.GetFunction(name)
}

func (s *ScopeTrace) IsMutatingEnvironment() bool {
	return s.CurrentScope.IsMutatingScope()
}

func NewScopeTrace() *ScopeTrace {
	globalScope := NewGlobalScope()
	return &ScopeTrace{
		GlobalScope:  globalScope,
		CurrentScope: globalScope,
	}
}

func NewVectorScope() *BaseScope {
	var scope = &BaseScope{
		name:      "vector",
		variables: make(map[string]*Variable),
		children:  make([]*BaseScope, 0),
		functions: make(map[string]value.IVOR),
		parent:    nil,
	}

	return scope
}

func NewStructScope() *BaseScope {

	newGlobal := NewGlobalScope()

	return &BaseScope{
		name:      "struct",
		variables: make(map[string]*Variable),
		children:  make([]*BaseScope, 0),
		functions: make(map[string]value.IVOR),
		structs:   make(map[string]*Struct),
		parent:    newGlobal,
		isStruct:  true,
	}
}

type ReportTable struct {
	GlobalScope ReportScope
}

type ReportScope struct {
	Name        string
	Vars        []ReportSymbol
	Funcs       []ReportSymbol
	Structs     []ReportSymbol
	ChildScopes []ReportScope
}

type ReportSymbol struct {
	Name   string
	Type   string
	Line   int
	Column int
}

// nuevooooooo
type HTMLReport struct {
	Reporte string
}

func (s *ScopeTrace) Report() ReportTable {
	return ReportTable{
		GlobalScope: s.CurrentScope.Report(),
	}
}

func (s *BaseScope) Report() ReportScope {

	reportScope := ReportScope{
		Name:        s.name,
		Vars:        make([]ReportSymbol, 0),
		Funcs:       make([]ReportSymbol, 0),
		Structs:     make([]ReportSymbol, 0),
		ChildScopes: make([]ReportScope, 0),
	}

	for _, v := range s.variables {

		token := v.Token
		line := 0
		column := 0

		if token != nil {
			line = token.GetLine()
			column = token.GetColumn()
		}

		s.reportito += "<tr><td>" + v.Name + "</td><td>" + v.Type + "</td><td>" + v.Value.Type() + "</td><td>" + s.name + "</td></tr>"
		reportScope.Vars = append(reportScope.Vars, ReportSymbol{
			Name:   v.Name,
			Type:   v.Type,
			Line:   line,
			Column: column,
		})

	}

	for _, f := range s.functions {
		switch function := f.(type) {
		case *BuiltInFunction:
			reportScope.Funcs = append(reportScope.Funcs, ReportSymbol{
				Name:   function.Name,
				Type:   "Embebida: " + function.Name,
				Line:   0,
				Column: 0,
			})
		case *Function:

			line := 0
			column := 0

			if function.Token != nil {
				line = function.Token.GetLine()
				column = function.Token.GetColumn()
			}

			reportScope.Funcs = append(reportScope.Funcs, ReportSymbol{
				Name:   function.Name,
				Type:   function.ReturnType,
				Line:   line,
				Column: column,
			})
		}
	}

	for _, v := range s.structs {
		reportScope.Structs = append(reportScope.Structs, ReportSymbol{
			Name:   v.Name,
			Type:   v.Name,
			Line:   v.Token.GetLine(),
			Column: v.Token.GetColumn(),
		})
	}

	for _, v := range s.children {
		reportScope.ChildScopes = append(reportScope.ChildScopes, v.Report())
	}

	return reportScope
}

func (s *ScopeTrace) GenerateHTMLReport() string {
	var buf bytes.Buffer

	buf.WriteString(`<!DOCTYPE html>
    <html>
    <head>
        <title>Tabla de Símbolos - V-Lang Cherry</title>
        <style>
            body { font-family: Arial, sans-serif; margin: 20px; background-color: #f5f5f5; }
            h1 { color: #2c3e50; text-align: center; }
            .report-container { 
                background-color: white; 
                border-radius: 8px; 
                box-shadow: 0 2px 4px rgba(0,0,0,0.1); 
                padding: 20px; 
                margin-bottom: 20px; 
            }
            .scope-container { 
                margin-left: 20px; 
                margin-bottom: 15px;
                border-left: 3px solid #3498db;
                padding-left: 10px;
            }
            .scope-title { 
                color: #e74c3c; 
                margin-top: 15px; 
                margin-bottom: 5px; 
                font-size: 1.1em;
            }
            table {
                width: 100%; 
                border-collapse: collapse; 
                margin-top: 10px;
                margin-bottom: 15px;
            }
            th, td { 
                border: 1px solid #ddd; 
                padding: 8px; 
                text-align: left; 
            }
            th { 
                background-color: #3498db; 
                color: white; 
            }
            tr:nth-child(even) { background-color: #f2f2f2; }
        </style>
    </head>
    <body>
        <h1>Tabla de Símbolos</h1>
        <div class="report-container">`)

	s.writeScopeHTML(s.GlobalScope, &buf, 0)

	buf.WriteString(`</div>
    </body>
    </html>`)

	return buf.String()
}

func (s *ScopeTrace) writeScopeHTML(scope *BaseScope, buf *bytes.Buffer, depth int) {
	// Función para verificar si el scope tiene contenido
	hasContent := func(s *BaseScope) bool {
		return len(s.variables) > 0 || len(s.functions) > 0 || len(s.structs) > 0
	}

	// Si el scope actual no tiene contenido, no lo mostramos
	if !hasContent(scope) {
		return
	}

	// Escribir el encabezado del scope
	indent := strings.Repeat("&nbsp;&nbsp;&nbsp;&nbsp;", depth)
	buf.WriteString(fmt.Sprintf(`<div class="scope-container">
        <div class="scope-title">%sÁmbito: %s</div>`, indent, scope.name))

	// Mostrar variables si existen
	if len(scope.variables) > 0 {
		buf.WriteString(`<table>
            <tr>
                <th>Nombre</th>
                <th>Tipo</th>
                <th>Línea</th>
                <th>Columna</th>
            </tr>`)

		for _, v := range scope.variables {
			line, column := 0, 0
			if v.Token != nil {
				line = v.Token.GetLine()
				column = v.Token.GetColumn()
			}
			buf.WriteString(fmt.Sprintf(`
            <tr>
                <td>%s%s</td>
                <td>%s</td>
                <td>%d</td>
                <td>%d</td>
            </tr>`, indent, v.Name, v.Type, line, column))
		}
		buf.WriteString(`</table>`)
	}

	// Mostrar funciones si existen (similar a variables)
	if len(scope.functions) > 0 {
		buf.WriteString(`<h4>Funciones:</h4>
        <table>
            <tr>
                <th>Nombre</th>
                <th>Tipo de Retorno</th>
                <th>Línea</th>
                <th>Columna</th>
            </tr>`)

		for name := range scope.functions {
			line, column := 0, 0
			buf.WriteString(fmt.Sprintf(`
            <tr>
                <td>%s%s</td>
                <td>%s</td>
                <td>%d</td>
                <td>%d</td>
            </tr>`, indent, name, "function", line, column))
		}
		buf.WriteString(`</table>`)
	}

	// Mostrar structs si existen (similar a variables)
	if len(scope.structs) > 0 {
		buf.WriteString(`<h4>Structs:</h4>
        <table>
            <tr>
                <th>Nombre</th>
                <th>Línea</th>
                <th>Columna</th>
            </tr>`)

		for _, st := range scope.structs {
			buf.WriteString(fmt.Sprintf(`
            <tr>
                <td>%s%s</td>
                <td>%d</td>
                <td>%d</td>
            </tr>`, indent, st.Name, st.Token.GetLine(), st.Token.GetColumn()))
		}
		buf.WriteString(`</table>`)
	}

	buf.WriteString(`</div>`)

	// Procesar hijos recursivamente (solo si tienen contenido)
	for _, child := range scope.children {
		if hasContent(child) {
			s.writeScopeHTML(child, buf, depth+1)
		}
	}
}
