package arm

import (
	"fmt"
	"log"
	compiler "main/parser"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// Estructura para variables en el visitor ARM
// (simplificada para el generador de código)
type ArmVariable struct {
	Name  string
	Type  string
	Value interface{} // Puede ser int, float64, string, bool, nil
}

type ArmScope struct {
	Parent    *ArmScope
	Variables map[string]*ArmVariable
}

func NewArmScope(parent *ArmScope) *ArmScope {
	return &ArmScope{
		Parent:    parent,
		Variables: make(map[string]*ArmVariable),
	}
}

type ArmVisitor struct {
	compiler.BaseVlangVisitor               // ReplVisitor para manejar el contexto del REPL
	Generator                 *ArmGenerator // Generador de instrucciones ARM
	Scope                     *ArmScope     // Scope actual
}

func (v *ArmVisitor) Visit(tree antlr.ParseTree) interface{} {
	if tree == nil {
		fmt.Println("⚠️ Árbol nulo recibido.")
		return nil
	}
	fmt.Println("Enrutado------------------")
	switch node := tree.(type) {
	case *antlr.ErrorNodeImpl:
		log.Fatal(node.GetText())
	case *compiler.ProgramaContext:
		return v.VisitPrograma(node)
	case *compiler.StmtContext:
		return v.VisitStmt(node)
	case *compiler.PrintlnStmtContext:
		return v.VisitPrintlnStmt(node)
	case *compiler.PrintStmtContext:
		return v.VisitPrintStmt(node)
	case *compiler.ValorEnteroContext:
		return v.VisitValorEntero(node)
	case *compiler.ValorFloatContext:
		return v.VisitValorFloat(node)
	case *compiler.ValorCadenaContext:
		return v.VisitValorCadena(node)
	case *compiler.ValorBooleanoContext:
		return v.VisitValorBoolean(node)
	case *compiler.BinaryExpContext:
		return v.VisitBinaryExp(node)
	case *compiler.FuncCallContext:
		return v.VisitFuncCall(node)
	case *compiler.ValorexpresionContext:
		return v.VisitValorExpresion(node)
	case *compiler.DeclAssignContext:
		return v.VisitDeclAssign(node)
	case *compiler.IdContext:
		return v.VisitId(node)
	case *compiler.DirectAssignContext:
		return v.VisitDirectAssign(node)
	case *compiler.IfStmtContext:
		return v.VisitIfStmt(node)
	case *compiler.IfChainContext:
		return v.VisitIfChain(node, v.Generator.GetLabel())
	case *compiler.ElseStmtContext:
		return v.VisitElseStmt(node)
	case *compiler.FuncDeclContext:
		return v.VisitFuncDecl(node)
	case *compiler.DeclAssignTypeContext:
		return v.VisitDeclAssignType(node)
	case *compiler.Decl_stmtContext:
		return v.VisitDeclStmt(node)
	default:
		fmt.Printf("⚠️ Tipo inesperado en Visit(): %T\n", tree)
		return tree.Accept(v)
	}
	return nil
}

func (v *ArmVisitor) VisitStmt(ctx *compiler.StmtContext) interface{} {
	fmt.Println("🔍 Visitando Stmt:", ctx.GetText())

	if ctx == nil || ctx.GetChildCount() == 0 {
		fmt.Println("⚠️ Stmt vacío o nulo.")
		return nil
	}

	// Visitar todos los hijos del StmtContext
	for i := 0; i < ctx.GetChildCount(); i++ {
		child := ctx.GetChild(i)
		if parseTree, ok := child.(antlr.ParseTree); ok {
			v.Visit(parseTree)
		}
	}

	return nil
}

// Recibe DclVisitor e inicializa el ReplVisitor
func NewVisitorArm() *ArmVisitor {
	return &ArmVisitor{
		Generator: &ArmGenerator{
			instructions: []string{},
			stdLib:       NewStandardLibrary(), // Inicializa la biblioteca estándar
		},
		Scope: NewArmScope(nil), // Scope global
	}
}

func (v *ArmVisitor) VisitPrograma(ctx *compiler.ProgramaContext) interface{} {
	fmt.Println("🔍 Visitando el programa en el ARMVISITOR")

	// Inicializar stack y heap al inicio del programa
	v.Generator.Comment("Inicialización de stack y heap")
	v.Generator.instructions = append(v.Generator.instructions, "MOV SP, #0x40000") // Stack pointer a una dirección segura
	v.Generator.instructions = append(v.Generator.instructions, "ADR X10, heap")    // Heap pointer en X10

	// Recorrer todos los statements del programa
	for _, stmt := range ctx.AllStmt() {
		fmt.Println("---🔍 Visitando statement*:(text)", stmt.GetText())
		v.Visit(stmt)
	}

	// Finalizar el programa correctamente (solo una vez)
	v.Generator.Comment("Fin del programa - exit syscall")
	v.Generator.instructions = append(v.Generator.instructions, "MOV X0, #0")
	v.Generator.instructions = append(v.Generator.instructions, "MOV X8, #93")
	v.Generator.instructions = append(v.Generator.instructions, "SVC #0")

	return nil
}

func (v *ArmVisitor) VisitPrintlnStmt(ctx *compiler.PrintlnStmtContext) interface{} {
	v.Generator.Comment("Print statement")

	// Evaluar todas las expresiones y guardar los tipos
	types := make([]StackObject, 0, len(ctx.AllExpresion()))
	for _, exp := range ctx.AllExpresion() {
		v.Visit(exp)
		types = append(types, v.Generator.TopObject())
	}

	// Imprimir los valores en orden (de izquierda a derecha)
	for i := len(types) - 1; i >= 0; i-- {
		typ := types[i].Type
		switch typ {
		case IntType:
			v.Generator.PopObject(X0)
			v.Generator.PrintInteger(X0)
		case StringType:
			v.Generator.PopObject(X0) // Recupera la dirección base del string
			v.Generator.PrintString(X0)
		case FloatType:
			v.Generator.PopObject(D0)
			v.Generator.PrintFloat()
		case BoolType:
			v.Generator.PopObject(X0)
			v.Generator.PrintBool()
		default:
			fmt.Println("⚠️ Tipo no soportado para impresión:", typ)
		}
	}

	return nil
}

func (v *ArmVisitor) VisitValorEntero(ctx *compiler.ValorEnteroContext) interface{} {
	fmt.Println("🔍 Visitando valor entero:", ctx.GetText())
	value := ctx.GetText()
	fmt.Println("Visiting IntLiteral:", value)

	// Generar código para cargar el valor
	v.Generator.Comment("Constant: " + value)
	intValue, err := strconv.Atoi(value)
	if err != nil {
		fmt.Printf("Error convirtiendo valor entero: %v\n", err)
		return nil
	}
	v.Generator.Mov(X0, intValue) // MOV X0, #valor
	v.Generator.Push(X0)          // STR X0, [SP, #-8]!

	obj := StackObject{
		Type:   IntType, // O el tipo correcto según la expresión
		Length: 8,
		Depth:  v.Generator.depth,
		Id:     nil,
	}
	v.Generator.PushObject(obj)

	return nil
}

func (v *ArmVisitor) VisitValorFloat(ctx *compiler.ValorFloatContext) interface{} {
	value := ctx.FLOAT().GetText()
	v.Generator.Comment("Constant: " + value)

	// Convertir el string a float64
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic("Error al convertir string a float: " + err.Error())
	}

	// Crear el objeto tipo float y empujarlo a la pila
	floatObject := v.Generator.FloatObject()
	v.Generator.PushConstant(floatObject, f)

	return nil
}

func (v *ArmVisitor) VisitValorCadena(ctx *compiler.ValorCadenaContext) interface{} {
	value := ctx.CADENA().GetText()
	trimmed := strings.Trim(value, `"`)

	v.Generator.Comment("String constant: " + trimmed)

	// Crear objeto tipo string
	stringObj := v.Generator.StringObject()

	// Empujar el string a la pila
	v.Generator.PushConstant(stringObj, trimmed)

	return nil
}
func (v *ArmVisitor) VisitValorBoolean(ctx *compiler.ValorBooleanoContext) interface{} {
	value, err := strconv.ParseBool(ctx.BOOLEANO().GetText())

	if err != nil {
		fmt.Printf("Error al convertir valor booleano: %v\n", err)
		return nil
	}

	v.Generator.Comment("Boolean constant: " + ctx.BOOLEANO().GetText())
	boolObject := v.Generator.BoolObject()
	v.Generator.PushConstant(boolObject, value)

	return nil
}

func (v *ArmVisitor) VisitId(ctx *compiler.IdContext) interface{} {
	fmt.Println("🔍 Visitando ID:", ctx.GetText())
	varName := ctx.ID().GetText()

	// Buscar el objeto en la pila
	offset, obj := v.Generator.GetObject(varName)

	v.Generator.Mov(X0, offset)
	v.Generator.Add(X0, SP, X0) // Calcula la dirección del objeto en la pila
	v.Generator.Ldr(X0, X0, 0)  // Carga el valor del objeto en X0
	v.Generator.Push(X0)        // Pushea el valor a la pila

	newObject := v.Generator.CloneObject(obj)
	newObject.Id = nil
	v.Generator.PushObject(newObject) // Pushea el objeto a la pila

	return nil
}

func (v *ArmVisitor) VisitValorExpresion(ctx *compiler.ValorexpresionContext) interface{} {
	fmt.Println("Visitando valor expresión:", ctx.GetText())
	valor := ctx.Valor()

	return v.Visit(valor)
}

func (v *ArmVisitor) VisitBinaryExp(ctx *compiler.BinaryExpContext) interface{} {
	operation := ctx.GetOp().GetText()

	// Visitar operandos
	v.Visit(ctx.Expresion(0))
	v.Visit(ctx.Expresion(1))

	// Obtener tipos de operandos
	rightObj := v.Generator.PopObject(X1)
	leftObj := v.Generator.PopObject(X0)

	rightType := rightObj.Type
	leftType := leftObj.Type

	// Suma y concatenación
	if operation == "+" {
		if leftType == IntType && rightType == IntType {
			v.Generator.Add(X0, X0, X1)
			v.Generator.Push(X0)
			v.Generator.PushObject(v.Generator.IntObject())
			return nil
		}
		if (leftType == IntType && rightType == FloatType) || (leftType == FloatType && rightType == IntType) || (leftType == FloatType && rightType == FloatType) {
			// Conversión implícita a float
			if leftType == IntType {
				v.Generator.Scvtf(D0, X0)
			} else {
				v.Generator.Fmov(D0, X0)
			}
			if rightType == IntType {
				v.Generator.Scvtf(D1, X1)
			} else {
				v.Generator.Fmov(D1, X1)
			}
			v.Generator.Fadd(D0, D0, D1)
			v.Generator.Push(D0)
			v.Generator.PushObject(v.Generator.FloatObject())
			return nil
		}
		if leftType == StringType && rightType == StringType {
			// Concatenación de strings (simplificada: solo comentario)
			v.Generator.Comment("Concatenación de strings no implementada (placeholder)")
			v.Generator.PushObject(v.Generator.StringObject())
			return nil
		}
		fmt.Println("Error: combinación de tipos inválida para suma/concatenación")
		return nil
	}

	// Resta
	if operation == "-" {
		if leftType == IntType && rightType == IntType {
			v.Generator.Sub(X0, X0, X1)
			v.Generator.Push(X0)
			v.Generator.PushObject(v.Generator.IntObject())
			return nil
		}
		if (leftType == IntType && rightType == FloatType) || (leftType == FloatType && rightType == IntType) || (leftType == FloatType && rightType == FloatType) {
			if leftType == IntType {
				v.Generator.Scvtf(D0, X0)
			} else {
				v.Generator.Fmov(D0, X0)
			}
			if rightType == IntType {
				v.Generator.Scvtf(D1, X1)
			} else {
				v.Generator.Fmov(D1, X1)
			}
			v.Generator.Fsub(D0, D0, D1)
			v.Generator.Push(D0)
			v.Generator.PushObject(v.Generator.FloatObject())
			return nil
		}
		fmt.Println("Error: combinación de tipos inválida para resta")
		return nil
	}

	// Multiplicación
	if operation == "*" {
		if leftType == IntType && rightType == IntType {
			v.Generator.Mul(X0, X0, X1)
			v.Generator.Push(X0)
			v.Generator.PushObject(v.Generator.IntObject())
			return nil
		}
		if (leftType == IntType && rightType == FloatType) || (leftType == FloatType && rightType == IntType) || (leftType == FloatType && rightType == FloatType) {
			if leftType == IntType {
				v.Generator.Scvtf(D0, X0)
			} else {
				v.Generator.Fmov(D0, X0)
			}
			if rightType == IntType {
				v.Generator.Scvtf(D1, X1)
			} else {
				v.Generator.Fmov(D1, X1)
			}
			v.Generator.Fmul(D0, D0, D1)
			v.Generator.Push(D0)
			v.Generator.PushObject(v.Generator.FloatObject())
			return nil
		}
		fmt.Println("Error: combinación de tipos inválida para multiplicación")
		return nil
	}

	// División
	if operation == "/" {
		if leftType == IntType && rightType == IntType {
			// Verificar división por cero
			v.Generator.Comment("Verificar división por cero (no implementado)")
			v.Generator.Div(X0, X0, X1)
			v.Generator.Push(X0)
			v.Generator.PushObject(v.Generator.IntObject())
			return nil
		}
		if (leftType == IntType && rightType == FloatType) || (leftType == FloatType && rightType == IntType) || (leftType == FloatType && rightType == FloatType) {
			if leftType == IntType {
				v.Generator.Scvtf(D0, X0)
			} else {
				v.Generator.Fmov(D0, X0)
			}
			if rightType == IntType {
				v.Generator.Scvtf(D1, X1)
			} else {
				v.Generator.Fmov(D1, X1)
			}
			v.Generator.Fdiv(D0, D0, D1)
			v.Generator.Push(D0)
			v.Generator.PushObject(v.Generator.FloatObject())
			return nil
		}
		fmt.Println("Error: combinación de tipos inválida para división")
		return nil
	}

	// Módulo
	if operation == "%" {
		if leftType == IntType && rightType == IntType {
			// Verificar división por cero
			v.Generator.Comment("Verificar módulo por cero (no implementado)")
			v.Generator.Comment("Módulo no implementado (placeholder)")
			v.Generator.Push(X0)
			v.Generator.PushObject(v.Generator.IntObject())
			return nil
		}
		fmt.Println("Error: combinación de tipos inválida para módulo")
		return nil
	}

	// Operaciones relacionales
	if operation == "==" || operation == "!=" || operation == "<" || operation == ">" || operation == "<=" || operation == ">=" {
		// int y float64
		if (leftType == IntType || leftType == FloatType) && (rightType == IntType || rightType == FloatType) {
			// Convertir ambos a float si alguno es float
			if leftType == IntType {
				v.Generator.Scvtf(D0, X0)
			} else {
				v.Generator.Fmov(D0, X0)
			}
			if rightType == IntType {
				v.Generator.Scvtf(D1, X1)
			} else {
				v.Generator.Fmov(D1, X1)
			}
			// Comparar D0 y D1
			v.Generator.instructions = append(v.Generator.instructions, "FCMP D0, D1")
			trueLabel := v.Generator.GetLabel()
			endLabel := v.Generator.GetLabel()
			cmpInstr := map[string]string{"==": "BEQ", "!=": "BNE", "<": "BLT", ">": "BGT", "<=": "BLE", ">=": "BGE"}[operation]
			v.Generator.Mov(X0, 0)
			v.Generator.instructions = append(v.Generator.instructions, cmpInstr+" "+trueLabel)
			v.Generator.B(endLabel)
			v.Generator.SetLabel(trueLabel)
			v.Generator.Mov(X0, 1)
			v.Generator.SetLabel(endLabel)
			v.Generator.Push(X0)
			v.Generator.PushObject(v.Generator.BoolObject())
			return nil
		}
		// string: comparar direcciones (por ahora)
		if leftType == StringType && rightType == StringType {
			v.Generator.Cmp(X0, X1)
			trueLabel := v.Generator.GetLabel()
			endLabel := v.Generator.GetLabel()
			if operation == "==" {
				v.Generator.Mov(X0, 0)
				v.Generator.instructions = append(v.Generator.instructions, "BEQ "+trueLabel)
			} else if operation == "!=" {
				v.Generator.Mov(X0, 0)
				v.Generator.instructions = append(v.Generator.instructions, "BNE "+trueLabel)
			}
			v.Generator.B(endLabel)
			v.Generator.SetLabel(trueLabel)
			v.Generator.Mov(X0, 1)
			v.Generator.SetLabel(endLabel)
			v.Generator.Push(X0)
			v.Generator.PushObject(v.Generator.BoolObject())
			return nil
		}
		// bool
		if leftType == BoolType && rightType == BoolType {
			v.Generator.Cmp(X0, X1)
			trueLabel := v.Generator.GetLabel()
			endLabel := v.Generator.GetLabel()
			if operation == "==" {
				v.Generator.Mov(X0, 0)
				v.Generator.instructions = append(v.Generator.instructions, "BEQ "+trueLabel)
			} else if operation == "!=" {
				v.Generator.Mov(X0, 0)
				v.Generator.instructions = append(v.Generator.instructions, "BNE "+trueLabel)
			}
			v.Generator.B(endLabel)
			v.Generator.SetLabel(trueLabel)
			v.Generator.Mov(X0, 1)
			v.Generator.SetLabel(endLabel)
			v.Generator.Push(X0)
			v.Generator.PushObject(v.Generator.BoolObject())
			return nil
		}
		fmt.Println("Error: combinación de tipos inválida para operación relacional")
		return nil
	}

	// Operaciones lógicas
	if operation == "&&" || operation == "||" {
		if leftType == BoolType && rightType == BoolType {
			v.Generator.instructions = append(v.Generator.instructions, "AND X0, X0, X1")
			v.Generator.Push(X0)
			v.Generator.PushObject(v.Generator.BoolObject())
			return nil
		}
		fmt.Println("Error: combinación de tipos inválida para operación lógica")
		return nil
	}

	fmt.Println("Error: operación aritmética no soportada o combinación de tipos inválida")
	return nil
}

func (v *ArmVisitor) VisitDeclAssign(ctx *compiler.DeclAssignContext) interface{} {
	varName := ctx.ID().GetText()
	varType := "int"          // Por defecto, puedes mejorar para detectar el tipo real
	var value interface{} = 0 // Valor por defecto para int
	var stackType StackObjectType = IntType

	if ctx.Expresion() != nil {
		v.Visit(ctx.Expresion())
		// Extraer el tipo real de la expresión
		obj := v.Generator.TopObject()
		stackType = obj.Type
		v.Generator.PopObject(X0)
		v.Generator.Push(X0)
		if stackType == FloatType {
			varType = "float64"
			value = 0.0
		} else if stackType == StringType {
			varType = "string"
			value = ""
		} else if stackType == BoolType {
			varType = "bool"
			value = false
		}
	} else {
		// Reservar espacio en la pila ARM para el valor por defecto
		if varType == "int" {
			v.Generator.Mov(X0, 0)
			v.Generator.Push(X0)
			stackType = IntType
		} else if varType == "float64" {
			v.Generator.Comment("Valor por defecto float64 (0.0)")
			v.Generator.PushConstant(v.Generator.FloatObject(), 0.0)
			stackType = FloatType
		} else if varType == "string" {
			v.Generator.Comment("Valor por defecto string (\"\")")
			v.Generator.PushConstant(v.Generator.StringObject(), "")
			stackType = StringType
		} else if varType == "bool" {
			v.Generator.Mov(X0, 0)
			v.Generator.Push(X0)
			stackType = BoolType
		}
	}

	err := v.DeclareVariable(varName, varType, value)
	if err != nil {
		fmt.Println(err)
	}
	// Registrar el objeto en la pila ARM
	varObj := StackObject{
		Type:   stackType,
		Length: 8,
		Depth:  v.Generator.depth,
		Id:     &varName,
	}
	v.Generator.PushObject(varObj)
	v.Generator.TagObject(varName)
	return nil
}

func (v *ArmVisitor) VisitDirectAssign(ctx *compiler.DirectAssignContext) interface{} {
	varName := ctx.Id_pattern().GetText()
	v.Visit(ctx.Expresion())
	// El nuevo valor está en la cima de la pila ARM
	// PopObject saca el valor y lo pone en X0
	offset, _ := v.Generator.GetObject(varName)
	obj := v.Generator.TopObject()
	v.Generator.PopObject(X0)
	v.Generator.Mov(X1, offset)
	v.Generator.Add(X1, SP, X1)
	if obj.Type == FloatType {
		v.Generator.Str(X0, X1, 0) // Para float, X0 contiene el valor binario
	} else if obj.Type == StringType {
		v.Generator.Str(X0, X1, 0) // Para string, X0 contiene la dirección
	} else {
		v.Generator.Str(X0, X1, 0) // Para int y bool
	}

	variable, found := v.FindVariable(varName)
	if !found {
		fmt.Printf("Error: variable '%s' no declarada en este ámbito\n", varName)
		return nil
	}
	// Validar tipo (aquí solo int, float64, string, bool)
	if variable.Type != "int" && variable.Type != "float64" && variable.Type != "string" && variable.Type != "bool" {
		fmt.Printf("Error: tipo incompatible en asignación a '%s'\n", varName)
		return nil
	}
	return nil
}

// codigo para hacer los if  if
func (v *ArmVisitor) VisitIfStmt(ctx *compiler.IfStmtContext) interface{} {
	fmt.Println("🔍 Visitando IfStmt")

	endLabel := v.Generator.GetLabel()

	for i, chain := range ctx.AllIf_chain() {
		fmt.Println(strconv.Itoa(i))
		nextLabel := v.Generator.GetLabel()
		realChain := chain.(*compiler.IfChainContext) // ⚠️ type assertion
		v.VisitIfChain(realChain, nextLabel)
		v.Generator.B(endLabel)
		v.Generator.SetLabel(nextLabel)
	}

	if ctx.Else_stmt() != nil {
		v.VisitElseStmt(ctx.Else_stmt().(*compiler.ElseStmtContext))

	}

	v.Generator.SetLabel(endLabel)
	return nil
}

func (v *ArmVisitor) VisitIfChain(ctx *compiler.IfChainContext, elseLabel string) interface{} {
	fmt.Println("🔍 Visitando IfChain")

	// Evaluar la expresión condicional
	v.Visit(ctx.Expresion())

	// Sacar el valor booleano
	obj := v.Generator.PopObject(X0)
	if obj.Type != BoolType {
		fmt.Println("⚠️ Condición no booleana en If")
	}

	// Si es falso, saltar
	v.Generator.Cmp(X0, "0")
	v.Generator.Beq(elseLabel)

	// Visitar statements del bloque if
	for _, stmt := range ctx.AllStmt() {
		v.Visit(stmt)
	}
	return nil
}
func (v *ArmVisitor) VisitElseStmt(ctx *compiler.ElseStmtContext) interface{} {
	//v.ScopeTrace.PushScope("else")
	for _, stmt := range ctx.AllStmt() {
		v.Visit(stmt)
	}
	//v.ScopeTrace.PopScope()
	return nil
}

// Implementación de VisitFuncDecl
func (v *ArmVisitor) VisitFuncDecl(ctx *compiler.FuncDeclContext) interface{} {
	if ctx.ID() == nil {
		fmt.Println("FuncDeclContext sin ID. Nodo:", ctx.GetText())
		return nil
	}
	funcName := ctx.ID().GetText()
	v.Generator.Comment("Función: " + funcName)

	if funcName == "main" {
		v.EnterScope()
		for _, stmt := range ctx.AllStmt() {
			v.Visit(stmt)
		}
		v.ExitScope()
	}
	return nil
}

// Métodos para gestión de scopes y variables en ArmVisitor

// Crear un nuevo scope (por ejemplo, al entrar a un bloque)
func (v *ArmVisitor) EnterScope() {
	v.Scope = NewArmScope(v.Scope)
}

// Salir del scope actual (al salir de un bloque)
func (v *ArmVisitor) ExitScope() {
	if v.Scope.Parent != nil {
		v.Scope = v.Scope.Parent
	}
}

// Declarar una variable en el scope actual
func (v *ArmVisitor) DeclareVariable(name string, typ string, value interface{}) error {
	if _, exists := v.Scope.Variables[name]; exists {
		return fmt.Errorf("La variable '%s' ya existe en este ámbito", name)
	}
	v.Scope.Variables[name] = &ArmVariable{Name: name, Type: typ, Value: value}
	return nil
}

// Buscar una variable en los scopes (desde el más interno al más externo)
func (v *ArmVisitor) FindVariable(name string) (*ArmVariable, bool) {
	scope := v.Scope
	for scope != nil {
		if variable, ok := scope.Variables[name]; ok {
			return variable, true
		}
		scope = scope.Parent
	}
	return nil, false
}

// Implementación de VisitDeclAssignType
func (v *ArmVisitor) VisitDeclAssignType(ctx *compiler.DeclAssignTypeContext) interface{} {
	varName := ctx.Id_pattern().GetText()
	varType := ctx.Var_type().GetText()
	var value interface{}
	var stackType StackObjectType = IntType

	if ctx.Expresion() != nil {
		v.Visit(ctx.Expresion())
		obj := v.Generator.TopObject()
		stackType = obj.Type
		v.Generator.PopObject(X0)
		v.Generator.Push(X0)
		if stackType == FloatType {
			varType = "float64"
			value = 0.0
		} else if stackType == StringType {
			varType = "string"
			value = ""
		} else if stackType == BoolType {
			varType = "bool"
			value = false
		}
	} else {
		// Asignar valor por defecto según el tipo y reservar espacio en la pila ARM
		switch varType {
		case "int":
			value = 0
			v.Generator.Mov(X0, 0)
			v.Generator.Push(X0)
			stackType = IntType
		case "float", "float64":
			value = 0.0
			v.Generator.Comment("Valor por defecto float64 (0.0)")
			v.Generator.PushConstant(v.Generator.FloatObject(), 0.0)
			stackType = FloatType
		case "string":
			value = ""
			v.Generator.Comment("Valor por defecto string (\"\")")
			v.Generator.PushConstant(v.Generator.StringObject(), "")
			stackType = StringType
		case "bool":
			value = false
			v.Generator.Mov(X0, 0)
			v.Generator.Push(X0)
			stackType = BoolType
		default:
			value = nil
		}
	}

	err := v.DeclareVariable(varName, varType, value)
	if err != nil {
		fmt.Println(err)
	}
	// Registrar el objeto en la pila ARM
	varObj := StackObject{
		Type:   stackType,
		Length: 8,
		Depth:  v.Generator.depth,
		Id:     &varName,
	}
	v.Generator.PushObject(varObj)
	v.Generator.TagObject(varName)
	return nil
}

// Agrego el método VisitDeclStmt
func (v *ArmVisitor) VisitDeclStmt(ctx *compiler.Decl_stmtContext) interface{} {
	// Procesa todas las declaraciones dentro del Decl_stmt
	for i := 0; i < ctx.GetChildCount(); i++ {
		child := ctx.GetChild(i)
		if parseTree, ok := child.(antlr.ParseTree); ok {
			v.Visit(parseTree)
		}
	}
	return nil
}
