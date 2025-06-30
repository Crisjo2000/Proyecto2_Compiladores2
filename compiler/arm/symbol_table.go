package arm

import "fmt"

// SymbolKind diferencia entre tipos de símbolos
type SymbolKind int

const (
	VariableKind SymbolKind = iota
	FunctionKind
)

// Symbol representa una entrada en la tabla de símbolos (variable, función, etc.)
type Symbol struct {
	Name   string
	Kind   SymbolKind
	Type   StackObjectType // Para variables: tipo de la variable. Para funciones: tipo de retorno.
	Scope  *Scope
	Params []*Symbol // Solo para funciones: lista de parámetros
}

// Scope representa un ámbito en el código (global, función, bloque if, etc.)
type Scope struct {
	Outer   *Scope
	Symbols map[string]*Symbol
}

// NewScope crea un nuevo ámbito anidado dentro de un ámbito existente.
func NewScope(outer *Scope) *Scope {
	return &Scope{
		Outer:   outer,
		Symbols: make(map[string]*Symbol),
	}
}

// SymbolTable gestiona todos los ámbitos y símbolos del programa.
type SymbolTable struct {
	GlobalScope  *Scope
	CurrentScope *Scope
}

// NewSymbolTable crea una nueva tabla de símbolos con un ámbito global.
func NewSymbolTable() *SymbolTable {
	global := NewScope(nil)
	return &SymbolTable{
		GlobalScope:  global,
		CurrentScope: global,
	}
}

// EnterScope entra en un nuevo ámbito.
func (st *SymbolTable) EnterScope() {
	st.CurrentScope = NewScope(st.CurrentScope)
}

// ExitScope sale del ámbito actual y regresa al ámbito exterior.
func (st *SymbolTable) ExitScope() {
	if st.CurrentScope.Outer != nil {
		st.CurrentScope = st.CurrentScope.Outer
	}
}

// Define un nuevo símbolo en el ámbito actual.
// Devuelve un error si el símbolo ya está definido en este ámbito.
func (st *SymbolTable) Define(name string, symbolType StackObjectType) (*Symbol, error) {
	if _, exists := st.CurrentScope.Symbols[name]; exists {
		return nil, fmt.Errorf("error semántico: la variable '%s' ya ha sido declarada en este ámbito", name)
	}
	symbol := &Symbol{Name: name, Type: symbolType, Scope: st.CurrentScope, Kind: VariableKind}
	st.CurrentScope.Symbols[name] = symbol
	return symbol, nil
}

// DefineFunction define un símbolo de función en el ámbito actual.
func (st *SymbolTable) DefineFunction(name string, returnType StackObjectType, params []*Symbol) (*Symbol, error) {
	if _, exists := st.CurrentScope.Symbols[name]; exists {
		return nil, fmt.Errorf("error semántico: el símbolo '%s' ya ha sido declarado en este ámbito", name)
	}
	symbol := &Symbol{
		Name:   name,
		Kind:   FunctionKind,
		Type:   returnType,
		Scope:  st.CurrentScope,
		Params: params,
	}
	st.CurrentScope.Symbols[name] = symbol
	return symbol, nil
}

// Resolve busca un símbolo por nombre, comenzando desde el ámbito actual
// y subiendo hacia los ámbitos exteriores.
// Devuelve el símbolo si se encuentra, o nil si no existe.
func (st *SymbolTable) Resolve(name string) *Symbol {
	scope := st.CurrentScope
	for scope != nil {
		if symbol, found := scope.Symbols[name]; found {
			return symbol
		}
		scope = scope.Outer
	}
	return nil
}
