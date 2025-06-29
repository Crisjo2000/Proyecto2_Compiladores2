package repl

import (
	"fmt"
	"log"
	compiler "main/parser"
	parser "main/parser"
	"main/value"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type Value interface {
	// métodos comunes que necesites
}

type IntValue struct {
	Value int
}

// Visitor personalizado para recorrer el árbol de sintaxis

// Constructor del visitor
type ReplVisitor struct {
	parser.BaseVlangVisitor
	ScopeTrace  *ScopeTrace
	CallStack   *CallStack
	Console     *Console
	ErrorTable  *ErrorTable
	StructNames []string
}
type ControlFlow int

const (
	FlowNormal ControlFlow = iota
	FlowBreak
	FlowContinue
	FlowReturn
)

var Bandera_Break_If = false
var Bandera_Break = false
var Bandera_Continue = false

// Recibe DclVisitor e inicializa el ReplVisitor
func NewVisitor(dclVisitor *DclVisitor) *ReplVisitor {
	return &ReplVisitor{
		ScopeTrace:  dclVisitor.ScopeTrace,
		ErrorTable:  dclVisitor.ErrorTable,
		StructNames: dclVisitor.StructNames,
		CallStack:   NewCallStack(),
		Console:     NewConsole(),
	}
}

// tambien podemos crear un ReplVisitor sin DclVisitor
// Esto es util para hacer tests de una vez
func NewReplVisitor(errorTable *ErrorTable) *ReplVisitor {
	return &ReplVisitor{
		ScopeTrace:  NewScopeTrace(),
		ErrorTable:  errorTable,
		StructNames: []string{},
		CallStack:   NewCallStack(),
		Console:     NewConsole(),
	}
}

// GetReplContext devuelve el contexto del REPL
// que contiene la consola, el ScopeTrace, el CallStack y la ErrorTable
func (v *ReplVisitor) GetReplContext() *ReplContext {
	return &ReplContext{
		Console:    v.Console,
		ScopeTrace: v.ScopeTrace,
		CallStack:  v.CallStack,
		ErrorTable: v.ErrorTable,
	}
}

// ValidType verifica si el tipo es valido en el ScopeTrace
// Esto se usa para validar tipos de variables, funciones, etc.
func (v *ReplVisitor) ValidType(_type string) bool {
	return v.ScopeTrace.GlobalScope.ValidType(_type)
}

// Visit es el metodo que se llama para visitar un nodo del arbol de sintaxis
// Dependiendo del tipo de nodo, se llama al metodo correspondiente
// Si es un ErrorNodeImpl, se lanza un error
func (v *ReplVisitor) Visit(tree antlr.ParseTree) interface{} {
	if tree == nil {
		fmt.Println("⚠️ Árbol nulo recibido.")
		return nil
	}
	fmt.Println("Enrutado------------------")
	switch node := tree.(type) {
	case *antlr.ErrorNodeImpl:
		log.Fatal(node.GetText())

	case *parser.ProgramaContext:
		return v.VisitPrograma(node)

	case *parser.PrintlnStmtContext:
		return v.VisitPrintlnStmt(node)

	case *parser.PrintStmtContext:
		return v.VisitPrintStmt(node)
	case *parser.ValorexpresionContext:
		// Visitamos el hijo de valorexpresion
		return v.VisitValorexpresion(node)

	// Declaraciones de variables
	case *parser.IncrementoContext:
		// Visitamos el hijo de valorexpresion
		return v.VisitIncremento(node)

	case *parser.Incremento2Context:
		return v.VisitIncremento2(node)

	case *parser.DecrementoContext:
		// Visitamos el hijo de valorexpresion
		return v.VisitDecremento(node)

	case *compiler.DeclAssignContext:
		fmt.Println("🔍 Visitando DeclAssign:", node.GetText())
		return v.VisitValueDeclAssign(node)
	case *parser.DeclAssignTypeContext:
		fmt.Println("🔍 Visitando DeclAssignType:", node.GetText())
		return v.VisitDeclAssignType(node)

	// tiramos todos los case de los nodos literales
	case *parser.ValorEnteroContext:
		fmt.Println("🔍 Visitando ValorEntero:", node.GetText())
		return v.VisitValorEntero(node)
	case *parser.ValorFloatContext:
		return v.VisitValorFloat(node)
	case *parser.ValorDecimalContext:
		return v.VisitValorDecimal(node)
	case *parser.ValorCadenaContext:
		return v.VisitValorCadena(node)
	case *parser.ValorBooleanoContext:
		return v.VisitValorBooleano(node)
	case *parser.ValorCaracterContext:
		return v.VisitValorCaracter(node)
	case *parser.IdContext:
		return v.VisitId(node)

	// Expresiones binarias
	case *parser.BinaryExpContext:
		return v.VisitBinaryExp(node)
	/// Agregamos las declaraciones
	// asignaciones posibles
	case *parser.StmtContext:
		return v.VisitStmt(node)
	// parte de los If
	case *parser.IfStmtContext:
		fmt.Println("**-* VISITANDO IF STMT:", node.GetText())
		return v.VisitIfStmt(node)
	case *parser.IfChainContext:
		return v.VisitIfChain(node)
	case *parser.ElseStmtContext:
		return v.VisitElseStmt(node)
	case *parser.DirectAssignContext:
		return v.VisitDirectAssign(node)
	case *parser.UnarioContext:
		return v.VisitUnario(node)
	case *parser.ParentesisexpreContext:
		return v.VisitParentesisexpre(node)
	case *parser.Decremento2Context:
		fmt.Println("🔔 Visitando nodo decremento2")
		return v.VisitDecremento2(node)
	case *parser.WhileStmtContext:
		return v.VisitWhileStmt(node)
	case *parser.ClasicForContext:
		return v.VisitClasicFor(node)
	case *parser.ForStmtContext:
		return v.VisitForStmt(node)

	case *parser.SwitchStmtContext:
		return v.VisitSwitchStmt(node)

	case *parser.CaseClauseContext:
		return v.VisitCaseClause(node)

	case *parser.DefaultClauseContext:
		return v.VisitDefaultClause(node)
	case *parser.BreakStmtContext:
		return v.VisitBreakStmt(node)
	case *parser.ContinueStmtContext:
		return v.VisitContinueStmt(node)

	case *parser.ReturnStmtContext:
		return v.VisitReturnStmt(node)
	case *parser.FuncDeclContext:
		return v.VisitFuncDecl(node)
	case *parser.FuncCallContext:
		return v.VisitFuncCall(node)
	case *parser.IdPatternContext:
		return v.VisitIdPattern(node)
	case *parser.DeclSlicesContext:
		return v.VisitDeclSlices(node)
	case *parser.IntTypeContext:
		return v.VisitIntType(node)
	case *parser.Fn_atoiContext:
		return v.VisitFn_atoi(node)
	case *parser.Fn_parsefloatContext:

		return v.VisitFn_parsefloat(node)
	case *parser.Fn_typeOFContext:
		return v.VisitFn_typeOF(node)
	case *parser.IndexAccesoContext:
		return v.VisitIndexAcceso(node)

	case *parser.Fn_indexOfContext:
		fmt.Println("*-*-* entrando al case de index")
		return v.VisitFn_indexOf(node)
	case *parser.Fn_lenxContext:
		return v.VisitFn_lenx(node)

	case *parser.Fn_joinContext:
		return v.VisitFn_join(node)

	case *parser.StringTypeContext:
		return v.VisitStringType(node)

	case *parser.StructDeclContext:
		return v.VisitStructDecl(node)
	case *parser.FuncionexpreContext:
		return v.VisitFuncionexpre(node)

	default:
		fmt.Printf("⚠️ Tipo inesperado en Visit(): %T\n", tree)
		return tree.Accept(v) // fallback por si acaso
	}

	return nil
}

// VisitPrograma es el metodo que se llama para visitar el nodo Programa
// Este nodo es el nodo raiz del arbol de sintaxis
// En este metodo recorremos todos los statements del programa
func (v *ReplVisitor) VisitPrograma(ctx *parser.ProgramaContext) interface{} {

	fmt.Println("🔍 Visitando el programa...")
	// Vamos a recorrer todos los statements y los visitamos
	for _, stmt := range ctx.AllStmt() {
		fmt.Println("---🔍 Visitando statement*:(text)", stmt.GetText())
		v.Visit(stmt)
	}

	return nil
}

/*

Ahora recorremos todos los stmts que tenemos en el programa

*/

func (v *ReplVisitor) VisitStmt(ctx *parser.StmtContext) interface{} {

	//fmt.Println("🔍 Visitando statement*:", ctx.GetText())
	fmt.Println("Entrando a instrucciones disponibles ", ctx.GetText())

	if ctx == nil || ctx.GetChildCount() == 0 {
		fmt.Println("⚠️ Stmt vacío o nulo.")
		return nil
	}

	node := ctx.GetChild(0)

	// Aquí van todos los ifs como los que hicimos antes:
	switch node := node.(type) {
	case *parser.PrintlnStmtContext:
		fmt.Println("🔔 Visitando nodo println")
		v.VisitPrintlnStmt(node)
	case *parser.PrintStmtContext:
		fmt.Println("🔔 Visitando nodo print")
		v.VisitPrintStmt(node)
	case *parser.DeclAssignContext:
		fmt.Println("🔔 Visitando nodo declAssign")
		v.VisitValueDeclAssign(node)
	case *parser.DeclAssignTypeContext:
		fmt.Println("🔔 Visitando nodo declAssignType")
		v.VisitDeclAssignType(node)
	case *parser.DirectAssignContext:
		fmt.Println("🔔 Visitando nodo directAssign")
		v.VisitDirectAssign(node)
	case *parser.IfStmtContext:
		fmt.Println("----🔔 Visitando nodo ifStmt")
		v.VisitIfStmt(node)
	case *parser.WhileStmtContext:
		fmt.Println("🔔 Visitando nodo whileStmt")
		v.VisitWhileStmt(node)
	case *parser.ForStmtContext:
		fmt.Println("🔔 Visitando nodo forStmt")
		v.VisitForStmt(node)

	case *parser.IncrementoContext:
		fmt.Println("🔔 Visitando nodo incremento")
		v.VisitIncremento(node)

	case *parser.DecrementoContext:
		fmt.Println("🔔 Visitando nodo decremento")
		v.VisitDecremento(node)

	case *parser.Decremento2Context:
		fmt.Println("🔔 Visitando nodo decremento")
		v.VisitDecremento2(node)

	case *parser.UnarioContext:
		v.VisitUnario(node)

	case *parser.ParentesisexpreContext:
		v.VisitParentesisexpre(node)

	case *parser.Incremento2Context:
		v.VisitIncremento2(node)

	case *parser.ClasicForContext:
		v.VisitClasicFor(node)

	case *parser.SwitchStmtContext:
		v.VisitSwitchStmt(node)

	case *parser.CaseClauseContext:
		v.VisitCaseClause(node)

	case *parser.DefaultClauseContext:
		v.VisitDefaultClause(node)

	// ---- Setencias de Transferencia-------
	case *parser.BreakStmtContext:
		v.VisitBreakStmt(node)
	case *parser.ContinueStmtContext:
		v.VisitContinueStmt(node)
	case *parser.ReturnStmtContext:
		fmt.Println("🔔 Visitando nodo returnStmt")
		return v.VisitReturnStmt(node)
	case *parser.FuncDeclContext:
		fmt.Println("🔔 Visitando nodo FuncDecl")
		v.VisitFuncDecl(node)
	case *parser.FuncCallContext:
		fmt.Println("🔔 Visitando nodo FuncCall")
		return v.VisitFuncCall(node)
	case *parser.IdPatternContext:
		fmt.Println("🔔 Visitando nodo IdPattern")
		return v.VisitIdPattern(node)
	case *parser.DeclSlicesContext:
		return v.VisitDeclSlices(node)
	case *parser.IntTypeContext:
		v.VisitIntType(node)
	case *parser.Fn_atoiContext:
		v.VisitFn_atoi(node)
	case *parser.Fn_parsefloatContext:
		v.VisitFn_parsefloat(node)
	case *parser.Fn_typeOFContext:
		v.VisitFn_typeOF(node)
	case *parser.Fn_indexOfContext:
		v.VisitFn_indexOf(node)
	case *parser.IndexAccesoContext:
		v.VisitIndexAcceso(node)
	case *parser.Fn_lenxContext:
		v.VisitFn_lenx(node)
	case *parser.Fn_joinContext:
		v.VisitFn_join(node)
	case *parser.StringTypeContext:
		v.VisitStringType(node)
	case *parser.FuncionexpreContext:
		v.VisitFuncionexpre(node)

	case *parser.StructDeclContext:
		fmt.Println("🔔 Visitando nodo Struct")
		return v.VisitStructDecl(node)

	default:
		fmt.Printf("⚠️ Tipo no reconocido dentro de stmt->: %T\n", node)
	}

	return nil
}

func (v *ReplVisitor) VisitDirectAssign(ctx *parser.DirectAssignContext) interface{} {
	// Vamos a visitar la asignacion
	// Primero vamos a ver si es una asignacion de variable
	// lo que queremos hacer para una asignacion
	// es ver si tenemos el ID
	fmt.Println(" // Parte de asignacion ")
	varName := ctx.Id_pattern().GetText()

	// Obtener el valor de la expresión (que puede ser una llamada a función)
	exprResult := v.Visit(ctx.Expresion())
	if exprResult == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Expresión inválida en asignación")
		return nil
	}

	varValue, ok := exprResult.(value.IVOR)
	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "La expresión no devuelve un valor válido")
		return nil
	}

	variable := v.ScopeTrace.GetVariable(varName)
	if variable == nil {
		fmt.Println("ERROR NO EXISTE LA VARIABLE")
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Variable "+varName+" no encontrada")

		// Añadir la variable automáticamente
		varType := varValue.Type()
		if _, msg := v.ScopeTrace.AddVariable(varName, varType, varValue, true, false, ctx.GetStart()); msg != "" {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
		}
		return nil
	}

	// Asignar el valor
	ok, msg := variable.Assign(varValue, true)
	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
	}

	return nil

}

/*
Con este metodo subimos el valor de la expresion
*/
func (v *ReplVisitor) VisitValorexpresion(ctx *parser.ValorexpresionContext) interface{} {
	// solo vamos a subir al hijo
	fmt.Println("🔍 Visitando Valorexpresion:", ctx.GetText())

	/*
		Todo esta bien, pero, tenemos que verificar el tipo de objeto que
		es el ctx y mandarlo al visit de nuevo para que suba pero
		con una forma de expresion con valores atomicos

		volvemos a subir el valor a la expresion

	*/
	valor := ctx.Valor()

	return v.Visit(valor)
}

// print normal
func (v *ReplVisitor) VisitPrintStmt(ctx *parser.PrintStmtContext) interface{} {

	fmt.Println("🔍 Visitando PrintStmt22:", ctx.GetText())
	fmt.Println("---------------------- Print", ctx.GetText())

	// Total de expresiones
	exprs := ctx.AllExpresion()

	// Lista de resultados para imprimir
	resultados := []string{}

	for _, expr := range exprs {
		result := v.Visit(expr)
		if result == nil {
			v.ErrorTable.NewSemanticError(expr.GetStart(), "Expresión vacía dentro de print")
			continue
		}

		val, ok := result.(value.IVOR)
		if !ok {
			v.ErrorTable.NewSemanticError(expr.GetStart(), "La expresión no devuelve un valor válido")
			continue
		}

		// Convertir el valor a string y agregar a resultados
		text := fmt.Sprintf("%v", val.Value())
		resultados = append(resultados, text)
	}

	// Unir con espacios y enviar a la consola
	finalOutput := strings.Join(resultados, " ")
	v.Console.Println(finalOutput)

	return nil
}

func (v *ReplVisitor) VisitPrintlnStmt(ctx *parser.PrintlnStmtContext) interface{} {
	fmt.Println("🔍 Visitando PrintlnStmt22:", ctx.GetText())

	// Obtener todas las expresiones
	exprs := ctx.AllExpresion()
	if len(exprs) == 0 {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Println sin argumentos")
		return nil
	}

	// Construir el resultado concatenando todos los argumentos
	var result strings.Builder
	for i, expr := range exprs {
		if i > 0 {
			result.WriteString(" ")
		}

		val := v.Visit(expr)
		if val == nil {
			continue
		}

		if ivor, ok := val.(value.IVOR); ok {
			result.WriteString(fmt.Sprintf("%v", ivor.Value()))
		} else {
			result.WriteString(fmt.Sprintf("%v", val))
		}
	}

	// Imprimir el resultado
	output := result.String()
	v.Console.Println(output)
	fmt.Println("🔔 Imprimiendo valor2:", output)
	return nil
}

// VisitIfStmt(ctx *IfStmtContext) interface{}
func (v *ReplVisitor) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {

	fmt.Println(("----- parte donde se ejefuta el If "))
	fmt.Println("🔍 Visitando IfStmt:", ctx.GetText())
	executed := false

	for _, ifChain := range ctx.AllIf_chain() {
		condition := v.Visit(ifChain).(value.IVOR)

		// Verificar que la condición sea booleana
		if condition.Type() != value.IVOR_BOOL {
			v.ErrorTable.NewSemanticError(ifChain.GetStart(), "La condición del if debe ser booleana")
			continue
		}

		// Si la condición es verdadera, ejecutar el bloque
		if condition.Value().(bool) {
			executed = true
			break
		}

	}

	// 2. Si ninguna condición se cumplió y hay else, ejecutarlo
	if !executed && ctx.Else_stmt() != nil {
		v.Visit(ctx.Else_stmt())
	}

	return nil
}

func (v *ReplVisitor) VisitIfChain(ctx *parser.IfChainContext) interface{} {
	fmt.Println("🔍 Visitando IfChain:", ctx.GetText())
	// Evaluar la expresión condicional
	exprResult := v.Visit(ctx.Expresion())

	// Verificar que sea un IVOR
	ivor, ok := exprResult.(value.IVOR)
	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "La condición del if debe evaluar a un valor booleano")
		return &value.BoolValue{InternalValue: false}
	}

	// Verificar que sea de tipo booleano
	if ivor.Type() != value.IVOR_BOOL {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Se esperaba un valor booleano en la condición")
		return &value.BoolValue{InternalValue: false}
	}

	// Si la condición es verdadera, ejecutar el bloque
	if ivor.Value().(bool) {
		v.ScopeTrace.PushScope("if")
		for _, stmt := range ctx.AllStmt() {
			v.Visit(stmt)

			fmt.Println("------------------------- comp braek if 5555")
			if stmt.GetText() == BreakItem {
				v.ScopeTrace.PopScope()
				fmt.Println("----------------------- BREAK DEL IF")
				Bandera_Break_If = true
				return &value.BoolValue{InternalValue: false}

			}
			if stmt.GetText() == ContinueItem {
				v.ScopeTrace.PopScope()
				fmt.Println("----------------------- BREAK DEL IF")
				Bandera_Continue = true
				return &value.BoolValue{InternalValue: false}

			}
		}
		v.ScopeTrace.PopScope()
		return &value.BoolValue{InternalValue: true}
	}

	return &value.BoolValue{InternalValue: false}
}

func (v *ReplVisitor) VisitElseStmt(ctx *parser.ElseStmtContext) interface{} {
	v.ScopeTrace.PushScope("else")
	for _, stmt := range ctx.AllStmt() {
		v.Visit(stmt)
	}
	v.ScopeTrace.PopScope()
	return nil
}

// -------DECLARACION DE VARIABLES SIN TIPO
func (v *ReplVisitor) VisitValueDeclAssign(ctx *parser.DeclAssignContext) interface{} {
	varName := ctx.ID().GetText()
	fmt.Println("🔍 Declarando variable mut:", varName)

	// Obtener el valor de la expresión
	var varValue value.IVOR
	if expr := ctx.Expresion(); expr != nil {
		result := v.Visit(expr)
		if result == nil {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "Expresión inválida en declaración")
			return nil
		}

		// Verificar si es una llamada a función
		if fnCall, isFunc := result.(value.IVOR); isFunc {
			varValue = fnCall
		} else {
			v.ErrorTable.NewSemanticError(ctx.GetStart(),
				"La expresión no devuelve un valor válido")
			return nil
		}
	} else {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Declaración sin expresión")
		return nil
	}

	// Determinar el tipo automáticamente
	varType := varValue.Type()
	if varType == "" {
		v.ErrorTable.NewSemanticError(ctx.GetStart(),
			"No se puede inferir el tipo de la expresión")
		return nil
	}

	// Crear la variable
	if _, msg := v.ScopeTrace.AddVariable(varName, varType, varValue, false, false, ctx.GetStart()); msg != "" {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
	}

	return nil
}

// Declaracion de variable con tipo
func (v *ReplVisitor) VisitDeclAssignType(ctx *parser.DeclAssignTypeContext) interface{} {
	fmt.Println("**** Declarando var tipo mut ***")
	/*
	* 1. Obtener el nombre de la variable
	* 2. Obtene el tipo de la variable
	* 3. Obtener el valor de la variable
	 */
	fmt.Println("9-*-*-*-*-**-**- ", ctx.Id_pattern().GetText(), " -", ctx.Var_type().GetText(), " - ", ctx.Expresion().GetText())
	varName := ctx.Id_pattern().GetText()

	varValue := v.Visit(ctx.Expresion()).(value.IVOR)
	varType := varValue.Type()

	fmt.Println("-*-*-*-///*/-*-*-9-**- ", &varValue)

	if varValue == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Expresión vacía dentro de declaración de variable")
		return nil
	}
	/*
		switch varType {
		case "int":
			varValue = value.IntValue{0}
		case "float":
			varValue = value.FloatValue{0.0}

		case "string":
			varValue = value.StringValue{""}
		case "bool":
			varValue = value.BoolValue{InternalValue: false}
		default:
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "Tipo desconocido: "+varType)
			return nil
		}

		/*
			// Verificar si el tipo es válido
			if !v.ValidType(varType) {
				v.ErrorTable.NewSemanticError(ctx.GetStart(), "Tipo '"+varType+"' no válido")
				return nil
			}*/
	// Verificar si el tipo de valor coincide con el tipo declarado
	/*if !varValue.Type(varType) {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "El valor '"+varValue.Value()+"' no es del tipo '"+varType+"'")
		return nil
	}*/
	// Agregar la variable al ScopeTrace
	variable, msg := v.ScopeTrace.AddVariable(varName, varType, varValue, false, false, ctx.GetStart())
	//variable, msg := v.ScopeTrace.AddVariable(varName, varType, varValue, false, false, ctx.GetStart())
	fmt.Println("Variable agregada:", varName, "de tipo:", varType, "con valor:", varValue.Value())
	// Variable already exists
	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
		return nil
	}

	return nil
}

/*


Devolvemos el valor de la expresion en este caso es el tope,
es un entero asi que solamente tomaremos lo que nos da el usuario
y lo convertiremos a un entero

*/

// el hijo de valorexpresioncontext -> valorEnter | valorBooleano |

func (v *ReplVisitor) VisitValorEntero(ctx *parser.ValorEnteroContext) interface{} {

	fmt.Println("🔍 Visitando ValorEntero:", ctx.GetText())

	intVal, _ := strconv.Atoi(ctx.GetText())

	return &value.IntValue{
		InternalValue: intVal,
	}

}

// Manejar Floats
func (v *ReplVisitor) VisitValorFloat(ctx *parser.ValorFloatContext) interface{} {

	floatVal, _ := strconv.ParseFloat(ctx.GetText(), 64)

	return &value.FloatValue{
		InternalValue: floatVal,
	}

}
func (v *ReplVisitor) VisitIdPattern(ctx *parser.IdPatternContext) interface{} {
	// hay que ir a buscar el Id a el entorno
	return ctx.GetText()
}

func (v *ReplVisitor) VisitCorchetesexpre(ctx *parser.CorchetesexpreContext) interface{} {

	return v.Visit(ctx.Expresion())
}

// Suma Re

func (v *ReplVisitor) VisitId(ctx *parser.IdContext) interface{} {
	id := ctx.ID().GetText()
	// TODO: buscar variable en entorno
	fmt.Println("Accediendo a variable:", id)
	// debemos buscar la variable en el ScopeTrace
	variable := v.ScopeTrace.GetVariable(id)
	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Variable '"+id+"' no encontrada")
		return nil
	} else {
		// Si la variable existe, devolvemos su valor
		fmt.Println("Variable encontrada:", id, "con valor:", variable.Value)
		return variable.Value
	}
}

func (v *ReplVisitor) VisitIncredecr(ctx *parser.IncredecrContext) interface{} {
	return v.Visit(ctx.Incredecre())
}

func (v *ReplVisitor) VisitExpdotexp1(ctx *parser.Expdotexp1Context) interface{} {
	left := ctx.ID(0).GetText()
	right := ctx.ID(1).GetText()
	// TODO: acceder a campo estructurado: ID.ID
	fmt.Println("Accediendo a campo estructurado:", left, ".", right)
	return nil
}

func (v *ReplVisitor) VisitExpdotexp(ctx *parser.ExpdotexpContext) interface{} {
	id := ctx.ID().GetText()
	expr := v.Visit(ctx.Expresion())
	// TODO: ID.EXPR, posiblemente array o campo dinámico
	fmt.Println("Accediendo a campo dinámico:", id, "con expresión:", expr)
	return nil
}

func (v *ReplVisitor) VisitAsignacionfor(ctx *parser.AsignacionforContext) interface{} {
	id := ctx.ID().GetText()
	expr := v.Visit(ctx.Expresion())
	fmt.Println("Asignando valor a variable:", id, "con expresión:", expr)
	// TODO: asignar expr a id
	return nil
}

func (v *ReplVisitor) VisitValorDecimal(ctx *parser.ValorDecimalContext) interface{} {
	text := ctx.DECIMAL().GetText()
	val, err := strconv.ParseFloat(text, 64)
	if err != nil {
		// Manejo de error
		// mandamos el error a ErrorTable
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Expresión vacía dentro de paréntesis")
		log.Printf("Error al convertir '%s' a decimal: %v", text, err)
		return nil
	}
	return &value.FloatValue{
		InternalValue: val,
	}
}

func (v *ReplVisitor) VisitValorCadena(ctx *parser.ValorCadenaContext) interface{} {
	stringVal := ctx.GetText()[1 : len(ctx.GetText())-1]

	stringVal = strings.ReplaceAll(stringVal, "\\\"", "\"")
	stringVal = strings.ReplaceAll(stringVal, "\\\\", "\\")
	stringVal = strings.ReplaceAll(stringVal, "\\n", "\n")
	stringVal = strings.ReplaceAll(stringVal, "\\r", "\r")

	// si es un caracter -> solamente un char
	if len(stringVal) == 1 {
		return &value.CharacterValue{
			InternalValue: stringVal,
		}
	}

	return &value.StringValue{
		InternalValue: stringVal,
	}
}

func (v *ReplVisitor) VisitValorBooleano(ctx *parser.ValorBooleanoContext) interface{} {
	text := ctx.BOOLEANO().GetText()
	// vemos si es true o false y devolvemos eso
	// si es true entonces devolvemos un valor booleano
	if text == "true" {
		return &value.BoolValue{
			InternalValue: true,
		}
	}
	if text == "false" {
		return &value.BoolValue{
			InternalValue: false,
		}
	}
	return nil
}

func (v *ReplVisitor) VisitValorCaracter(ctx *parser.ValorCaracterContext) interface{} {
	text := ctx.CARACTER().GetText()
	// Suponiendo que vienen entre comillas simples: 'a'
	if len(text) >= 3 {
		char := text[1 : len(text)-1]
		return &value.CharacterValue{
			InternalValue: char,
		}
	}
	return nil
}

/*
Ahora haremos la seccion de las expresiones binarias completa
en esta seccion haremos uso de un strategy, el cual
explicare en la clase del dia LUNES
*/
func (v *ReplVisitor) VisitBinaryExp(ctx *parser.BinaryExpContext) interface{} {

	op := ctx.GetOp().GetText()
	left := v.Visit(ctx.GetLeft()).(value.IVOR)
	// Verificamos si hay una estrategia de retorno anticipado para este operador
	earlyCheck, ok := EarlyReturnStrats[op]

	if ok {
		ok, _, result := earlyCheck.Validate(left)

		if ok {
			return result
		}
	}

	right := v.Visit(ctx.GetRight()).(value.IVOR)

	/*
			Aqui es donde sucede la magia, buscamos el operador
			y haciendo uso de un strategy pattern, validamos
			y ejecutamos la operacion de manera dinamica
		Si el operador no existe, lanzamos un error
		Si el operador existe, validamos los tipos de datos
	*/
	strat, ok := BinaryStrats[op]

	if !ok {
		log.Fatal("Binary operator not found")
	}

	ok, msg, result := strat.Validate(left, right)

	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetOp(), msg)
		return value.DefaultNilValue
	}

	return result
}

//llamadas a funciones

func (v *ReplVisitor) VisitFuncCall(ctx *parser.FuncCallContext) interface{} {
	funcName := ctx.Id_pattern().GetText()
	fmt.Println("🔍 Llamando a función:", funcName)

	// Buscar la función
	functionObj, errMsg := v.ScopeTrace.GetFunction(funcName)
	if functionObj == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), errMsg)
		return nil
	}

	function, ok := functionObj.(*Function)
	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "El identificador '"+funcName+"' no es una función válida")
		return nil
	}

	// Ejecutar la función
	v.ScopeTrace.PushScope("func:" + funcName)
	defer v.ScopeTrace.PopScope()

	var returnValue value.IVOR = &value.IntValue{InternalValue: 0} // Valor por defecto

	for _, stmt := range function.Body {
		// Obtener el nodo concreto del statement
		concreteNode := stmt.GetChild(0)

		// Verificar si es un return
		if retCtx, ok := concreteNode.(*parser.ReturnStmtContext); ok {
			if retCtx.Expresion() != nil {
				retVal := v.Visit(retCtx.Expresion())
				if retVal == nil {
					v.ErrorTable.NewSemanticError(retCtx.GetStart(), "Expresión de retorno inválida")
					continue
				}

				if ivor, ok := retVal.(value.IVOR); ok {
					returnValue = ivor
				} else {
					v.ErrorTable.NewSemanticError(retCtx.GetStart(),
						"El valor de retorno no implementa IVOR")
				}
			}
			break
		}

		// Ejecutar otros tipos de statements
		v.Visit(stmt)
	}

	fmt.Println("🔍 Función retornó:", returnValue)
	return returnValue
}

func (v *ReplVisitor) VisitIncremento(ctx *compiler.IncrementoContext) interface{} {
	varname := ctx.ID().GetText()
	variable := v.ScopeTrace.GetVariable(varname)
	if variable == nil {

		fmt.Println("NO existe la variable")
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Variable "+varname+" no encontrada")
	} else {
		fmt.Println(variable.Value, variable.Type, " &/&%%&//((&/(&(/)/)%%$###$%?==))((/&/&%%$$$$###))")
		if variable.Type == "Int" {
			fmt.Println("se puede sumar")

			valorActual := variable.Value.Value().(int)
			valorNew := valorActual + 1

			// Crea el nuevo valor como IVOR
			valorActual1 := value.IntValue{InternalValue: valorNew}

			// Asigna el nuevo valor
			ok, msg := variable.Assign(&valorActual1, true) // ¡OJO! necesitas pasar un *puntero*
			if !ok {
				v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
			}

		} else if variable.Type == "Float" {
			fmt.Println("se puede sumar")

			valorActual := variable.Value.Value().(float64)
			valorNew := valorActual + 1.0

			valorActual1 := value.FloatValue{InternalValue: valorNew}

			// Asigna el nuevo valor
			ok, msg := variable.Assign(&valorActual1, true) // ¡OJO! necesitas pasar un *puntero*
			if !ok {
				v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
			}

		} else {
			fmt.Println("Error semantico")
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "Tipo "+variable.Type+" tipo de dato incompatible con esta operación")

		}
	}

	return nil
}

func (v *ReplVisitor) VisitIncremento2(ctx *compiler.Incremento2Context) interface{} {

	varname := ctx.ID().GetText()
	variable := v.ScopeTrace.GetVariable(varname)
	varValor := v.Visit(ctx.Expresion()).(value.IVOR)
	if variable == nil {

		fmt.Println("NO existe la variable")
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Variable "+varname+" no encontrada")
	} else {
		fmt.Println(variable.Value, variable.Type, " &/&%%&//((&/(&(/)/)%%$###$%?==))((/&/&%%$$$$###))")
		if variable.Type == "Int" {
			fmt.Println("se puede sumar")

			valorActual := variable.Value.Value().(int)
			valorNew := valorActual + varValor.Value().(int)

			// Crea el nuevo valor como IVOR
			valorActual1 := value.IntValue{InternalValue: valorNew}

			// Asigna el nuevo valor
			ok, msg := variable.Assign(&valorActual1, true) // ¡OJO! necesitas pasar un *puntero*
			if !ok {
				v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
			}

		} else if variable.Type == "Float" {
			fmt.Println("se puede sumar")

			valorActual := variable.Value.Value().(float64)
			valorNew := valorActual + varValor.Value().(float64)

			valorActual1 := value.FloatValue{InternalValue: valorNew}

			// Asigna el nuevo valor
			ok, msg := variable.Assign(&valorActual1, true) // ¡OJO! necesitas pasar un *puntero*
			if !ok {
				v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
			}

		} else {
			fmt.Println("Error semantico")
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "Tipo "+variable.Type+" tipo de dato incompatible con esta operación")

		}
	}

	return nil
}

// ------>> Decremento
func (v *ReplVisitor) VisitDecremento(ctx *compiler.DecrementoContext) interface{} {
	varname := ctx.ID().GetText()
	variable := v.ScopeTrace.GetVariable(varname)
	if variable == nil {

		fmt.Println("NO existe la variable")
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Variable "+varname+" no encontrada")
	} else {
		fmt.Println(variable.Value, variable.Type, " &/&%%&//((&/(&(/)/)%%$###$%?==))((/&/&%%$$$$###))")
		if variable.Type == "Int" {
			fmt.Println("se puede sumar")

			valorActual := variable.Value.Value().(int)
			valorNew := valorActual - 1

			// Crea el nuevo valor como IVOR
			valorActual1 := value.IntValue{InternalValue: valorNew}

			// Asigna el nuevo valor
			ok, msg := variable.Assign(&valorActual1, true) // ¡OJO! necesitas pasar un *puntero*
			if !ok {
				v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
			}

		} else if variable.Type == "Float" {
			fmt.Println("se puede sumar")

			valorActual := variable.Value.Value().(float64)
			valorNew := valorActual - 1.0

			valorActual1 := value.FloatValue{InternalValue: valorNew}

			// Asigna el nuevo valor
			ok, msg := variable.Assign(&valorActual1, true) // ¡OJO! necesitas pasar un *puntero*
			if !ok {
				v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
			}

		} else {
			fmt.Println("Error semantico")
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "Tipo "+variable.Type+" tipo de dato incompatible con esta operación")

		}
	}

	return nil
}

// decremento x-= exp
func (v *ReplVisitor) VisitDecremento2(ctx *compiler.Decremento2Context) interface{} {

	varname := ctx.ID().GetText()
	variable := v.ScopeTrace.GetVariable(varname)
	varValor := v.Visit(ctx.Expresion()).(value.IVOR)
	if variable == nil {

		fmt.Println("NO existe la variable")
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Variable "+varname+" no encontrada")
	} else {
		fmt.Println(variable.Value, variable.Type, " &/&%%&//((&/(&(/)/)%%$###$%?==))((/&/&%%$$$$###))")
		if variable.Type == "Int" {
			fmt.Println("se puede sumar")

			valorActual := variable.Value.Value().(int)
			valorNew := valorActual - varValor.Value().(int)

			// Crea el nuevo valor como IVOR
			valorActual1 := value.IntValue{InternalValue: valorNew}

			// Asigna el nuevo valor
			ok, msg := variable.Assign(&valorActual1, true) // ¡OJO! necesitas pasar un *puntero*
			if !ok {
				v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
			}

		} else if variable.Type == "Float" {
			fmt.Println("se puede sumar")

			valorActual := variable.Value.Value().(float64)
			valorNew := valorActual - varValor.Value().(float64)

			valorActual1 := value.FloatValue{InternalValue: valorNew}

			// Asigna el nuevo valor
			ok, msg := variable.Assign(&valorActual1, true) // ¡OJO! necesitas pasar un *puntero*
			if !ok {
				v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
			}

		} else {
			fmt.Println("Error semantico")
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "Tipo "+variable.Type+" tipo de dato incompatible con esta operación")

		}
	}
	return nil
}

// Operacion Unario
func (v *ReplVisitor) VisitUnario(ctx *compiler.UnarioContext) interface{} {

	fmt.Println("----- Entrando a unario")

	// obtener el ope y  valor
	op := ctx.GetOp().GetText()
	exprValue := v.Visit(ctx.Expresion())

	varValue, ok := exprValue.(value.IVOR)
	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Expresión inválida en operación unaria")
		return nil
	}

	fmt.Println("Operador unario:", op, "Valor:", varValue)

	// verificamos que operador se va usar
	strategy, exists := UnaryStrats[op]
	if !exists {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Operador unario '"+op+"' no soportado")
		return nil
	}

	// operamos
	success, msg, result := strategy.Validate(varValue)
	if !success {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
		return nil
	}

	return result

}

func (v *ReplVisitor) VisitParentesisexpre(ctx *compiler.ParentesisexpreContext) interface{} {

	fmt.Println("Entrando a Parentesis")
	return v.Visit(ctx.Expresion())

}

func (v *ReplVisitor) VisitWhileStmt(ctx *parser.WhileStmtContext) interface{} {
	fmt.Println("entrando al foor while *-*-*-****-*-*-*")

	//varname := ctx.().GetText()
	//variable := v.ScopeTrace.GetVariable(varname)
	//varValor := v.Visit(ctx.Expresion()).(value.IVOR)

	//fmt.Println(varValor, " 8888888888888888888888888888888888")
	// Crear un nuevo scope para el while
	v.ScopeTrace.PushScope("while")
	defer v.ScopeTrace.PopScope()

	for {
		// 1. Evaluar la condición
		condition := v.Visit(ctx.Expresion())
		condValue, ok := condition.(value.IVOR)

		// Verificar que la condición sea válida
		if !ok {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "La condición del while debe evaluar a un valor")
			break
		}

		// Verificar que sea booleana
		if condValue.Type() != value.IVOR_BOOL {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "La condición del while debe ser booleana")
			break
		}

		// 2. Si la condición es falsa, salir del bucle
		if !condValue.Value().(bool) {
			break
		}

		// 3. Ejecutar el cuerpo del while
		for _, stmt := range ctx.AllStmt() {
			v.Visit(stmt)
			if Bandera_Break == true {
				Bandera_Break = false
				Bandera_Break_If = false
				return nil
			}
			if Bandera_Continue {
				Bandera_Continue = false
				break
			}

		}
	}

	return nil
}

func (v *ReplVisitor) VisitClasicFor(ctx *compiler.ClasicForContext) interface{} {
	fmt.Println("🔍 Visitando ForStmt clasico:", ctx.GetText())

	// 1. Ejecutar la asignación inicial (i = 0)
	v.Visit(ctx.Assign_stmt())

	// Crear un nuevo scope para el for
	v.ScopeTrace.PushScope("for")
	defer v.ScopeTrace.PopScope()

	loopCount := 0
	maxLoops := 100000 // Límite para evitar bucles infinitos

	for loopCount < maxLoops {
		loopCount++

		// 2. Evaluar la condición (i < 5)
		condition := v.Visit(ctx.Expresion())
		condValue, ok := condition.(value.IVOR)

		if !ok {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "La condición del for debe evaluar a un valor")
			break
		}

		// Verificar que sea booleana
		if condValue.Type() != value.IVOR_BOOL {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "La condición del for debe ser booleana")
			break
		}

		// Si la condición es falsa, salir del bucle
		if !condValue.Value().(bool) {
			break
		}

		// 3. Ejecutar el cuerpo del for
		for _, stmt := range ctx.AllStmt() {
			v.Visit(stmt)
			// Manejar señales de break
			fmt.Println("-FORRRR-----------------*-*-*-*- ", " *** ", stmt.GetText(), " *-* ", BreakItem)
			if stmt.GetText() == BreakItem {
				fmt.Println("-----------------------  ENTRANDO AL BREAK")
				return nil // Salir del bucle for
			} else if Bandera_Break_If {
				fmt.Println("ENTRANDO AL NULO DEL IF ")
				Bandera_Break_If = false
				Bandera_Break = false
				return nil

			}
			if Bandera_Continue {
				fmt.Println("-------------------- entrando a continue")
				Bandera_Continue = false
				break

			}
		}

		// 4. Ejecutar el incremento (i++)
		v.Visit(ctx.Incredecre())
	}

	if loopCount >= maxLoops {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Bucle for excedió el máximo de iteraciones")
	}

	return nil
}

// ---------------------------- FOR RANGE -----------------------------------
func (v *ReplVisitor) VisitForStmt(ctx *compiler.ForStmtContext) interface{} {
	return nil
}

// -------------------------- SWITCH FUN --------------------------------------------
func (v *ReplVisitor) VisitSwitchStmt(ctx *parser.SwitchStmtContext) interface{} {
	fmt.Println("🔍 Visitando SwitchStmt:", ctx.GetText())

	// 1. Obtener el valor a comparar (variable del switch)
	varName := ctx.ID().GetText()
	variable := v.ScopeTrace.GetVariable(varName)

	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Variable '"+varName+"' no definida")
		return nil
	}

	switchValue := variable.Value
	matched := false

	// Crear un nuevo scope para el switch
	v.ScopeTrace.PushScope("switch")
	defer v.ScopeTrace.PopScope()

	// 2. Procesar todos los casos
	for _, caseItem := range ctx.AllCase_clause() {
		caseCtx := caseItem.(*parser.CaseClauseContext)

		// Visitar el caso para obtener todas las expresiones a comparar
		caseExprValues := v.VisitCaseClause(caseCtx).([]value.IVOR)

		// Comparar con cada expresión del caso
		for _, caseExprValue := range caseExprValues {
			if equals(switchValue, caseExprValue) {
				matched = true
				// Ejecutar las sentencias del caso
				for _, stmt := range caseCtx.AllStmt() {
					v.Visit(stmt)
					if Bandera_Break == true {
						Bandera_Break = false
						Bandera_Break_If = false
						return nil

					}

				}
				break // Salir del bucle de comparación
			}
		}

		if matched {
			break // Salir después de ejecutar el primer caso que coincide
		}
	}

	// 3. Procesar el caso default si no hubo coincidencias
	if !matched && ctx.Default_clause() != nil {
		defaultCtx := ctx.Default_clause().(*parser.DefaultClauseContext)
		for _, stmt := range defaultCtx.AllStmt() {
			v.Visit(stmt)
		}
	}

	return nil
}

func (v *ReplVisitor) VisitCaseClause(ctx *parser.CaseClauseContext) interface{} {
	// Recoger todas las expresiones del caso (pueden ser múltiples separadas por coma)
	var caseExprValues []value.IVOR

	// La primera expresión es obligatoria
	firstExpr := v.Visit(ctx.Expresion(0)).(value.IVOR)
	caseExprValues = append(caseExprValues, firstExpr)

	// Expresiones adicionales (si las hay)
	for i := 1; i < len(ctx.AllExpresion()); i++ {
		expr := v.Visit(ctx.Expresion(i)).(value.IVOR)
		caseExprValues = append(caseExprValues, expr)
	}

	return caseExprValues
}

func (v *ReplVisitor) VisitDefaultClause(ctx *parser.DefaultClauseContext) interface{} {
	// El default no necesita lógica especial, las sentencias se ejecutan en VisitSwitchStmt
	return nil
}

// Función auxiliar para comparar valores
func equals(a, b value.IVOR) bool {
	if a.Type() != b.Type() {
		return false
	}

	switch aVal := a.(type) {
	case *value.IntValue:
		return aVal.InternalValue == b.(*value.IntValue).InternalValue
	case *value.FloatValue:
		return aVal.InternalValue == b.(*value.FloatValue).InternalValue
	case *value.StringValue:
		return aVal.InternalValue == b.(*value.StringValue).InternalValue
	case *value.BoolValue:
		return aVal.InternalValue == b.(*value.BoolValue).InternalValue
	default:
		return false
	}
}

// =========================== Sentencias de transferencia ============================

func (v *ReplVisitor) VisitBreakStmt(ctx *compiler.BreakStmtContext) interface{} {
	fmt.Println("--------- BREAK DETECTADO ----------- ", ctx.GetText())
	Bandera_Break = true
	return nil // Usamos la constante definida en el paquete
}

func (v *ReplVisitor) VisitContinueStmt(ctx *compiler.ContinueStmtContext) interface{} {
	Bandera_Continue = true
	return ContinueItem
}

func (v *ReplVisitor) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {
	if ctx.Expresion() != nil {
		return v.Visit(ctx.Expresion())
	}
	return &value.NilValue{} // Retorno por defecto si no hay expresión
}

func (v *ReplVisitor) VisitFuncDecl(ctx *parser.FuncDeclContext) interface{} {
	if ctx == nil {
		v.ErrorTable.NewSemanticError(nil, "Contexto de función es nil")
		return nil
	}

	// Obtener nombre de la función de manera segura
	var funcName string
	if id := ctx.ID(); id != nil {
		funcName = id.GetText()
	} else {

		return nil
	}

	fmt.Println("🔍 Declarando función:", funcName)

	// Obtener el cuerpo de la función de manera segura
	var body []parser.IStmtContext
	if stmts := ctx.AllStmt(); stmts != nil {
		body = stmts
	}

	// Crear la función
	newFunc := &Function{
		Name: funcName,
		Body: body,
	}

	// Registrar la función (verifica que ScopeTrace no sea nil)
	if v.ScopeTrace != nil {
		ok, msg := v.ScopeTrace.AddFunction(funcName, newFunc)
		if !ok {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
		} else {
			fmt.Println("✅ Función registrada:", funcName)
		}
	} else {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "ScopeTrace no inicializado")
	}

	return nil
}

func (v *ReplVisitor) VisitFuncArg(ctx *parser.FuncArgContext) interface{} {
	// Obtener nombre del parámetro y su tipo
	paramName := ctx.ID().GetText()
	typeCtx := ctx.Var_type()

	// Determinar el tipo (simplificado - podrías necesitar más lógica aquí)
	var paramType string
	switch typeCtx.GetText() {
	case "int":
		paramType = value.IVOR_INT
	case "float":
		paramType = value.IVOR_FLOAT
	case "string":
		paramType = value.IVOR_STRING
	case "bool":
		paramType = value.IVOR_BOOL
	case "char":
		paramType = value.IVOR_CHARACTER
	default:
		paramType = typeCtx.GetText() // Para tipos personalizados como structs
	}

	return &Param{
		ExternName: paramName,
		Type:       paramType,
	}
}

// Visit a parse tree produced by VlangParser#DeclSlices.
func (v *ReplVisitor) VisitDeclSlices(ctx *parser.DeclSlicesContext) interface{} {
	// Obtener el nombre de la variable
	// Obtener el nombre de la variable
	varName := ctx.ID().GetText()
	fmt.Println("🔍 Declarando slice:", varName)

	// Obtener el tipo del slice
	typeCtx := ctx.Var_type()
	if typeCtx == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Falta especificar el tipo del slice")
		return nil
	}

	// Visitar el nodo del tipo para obtener el tipo como string
	typeResult := v.Visit(typeCtx)
	if typeResult == nil {
		v.ErrorTable.NewSemanticError(typeCtx.GetStart(), "Tipo de slice no válido")
		return nil
	}

	sliceType, ok := typeResult.(string)
	if !ok {
		v.ErrorTable.NewSemanticError(typeCtx.GetStart(), "Tipo de slice no reconocido")
		return nil
	}

	fmt.Println("Tipo de slice:", sliceType)

	// Obtener todas las expresiones de inicialización
	var expressions []parser.IExpresionContext

	expressions = ctx.AllExpresion()

	fmt.Println("Número de elementos:", len(expressions))

	// Crear el slice según el tipo
	var sliceValue value.IVOR

	switch sliceType {
	case "int":
		values := make([]int, 0, len(expressions))
		for _, expr := range expressions {
			val := v.Visit(expr)
			if val == nil {
				v.ErrorTable.NewSemanticError(expr.GetStart(), "Expresión inválida en slice")
				continue
			}

			intVal, ok := val.(*value.IntValue)
			if !ok {
				v.ErrorTable.NewSemanticError(expr.GetStart(),
					fmt.Sprintf("Se esperaba tipo int en el slice, se obtuvo %T", val))
				continue
			}
			values = append(values, intVal.InternalValue)
		}
		sliceValue = &value.SliceValue{
			InternalValue: values,
			ElementType:   "int",
		}

	case "float":
		values := make([]float64, 0, len(expressions))
		for _, expr := range expressions {
			val := v.Visit(expr)
			if val == nil {
				v.ErrorTable.NewSemanticError(expr.GetStart(), "Expresión inválida en slice")
				continue
			}

			floatVal, ok := val.(*value.FloatValue)
			if !ok {
				v.ErrorTable.NewSemanticError(expr.GetStart(),
					fmt.Sprintf("Se esperaba tipo float en el slice, se obtuvo %T", val))
				continue
			}
			values = append(values, floatVal.InternalValue)
		}
		sliceValue = &value.SliceValue{
			InternalValue: values,
			ElementType:   "float",
		}
	case "string":
		values := make([]string, 0, len(expressions))
		for _, expr := range expressions {
			val := v.Visit(expr)
			if val == nil {
				v.ErrorTable.NewSemanticError(expr.GetStart(), "Expresión inválida en slice")
				continue
			}

			strVal, ok := val.(*value.StringValue)
			if !ok {
				v.ErrorTable.NewSemanticError(expr.GetStart(),
					fmt.Sprintf("Se esperaba tipo string en el slice, se obtuvo %T", val))
				continue
			}
			values = append(values, strVal.InternalValue)
		}
		sliceValue = &value.SliceValue{
			InternalValue: values,
			ElementType:   "string",
		}

	// Añadir más tipos según sea necesario (string, bool, etc.)
	default:
		v.ErrorTable.NewSemanticError(typeCtx.GetStart(),
			fmt.Sprintf("Tipo de slice no soportado: %s", sliceType))
		return nil
	}

	// Agregar la variable al ámbito
	variable, msg := v.ScopeTrace.AddVariable(varName, "[]"+sliceType, sliceValue, false, false, ctx.GetStart())
	if variable == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), msg)
	}
	return nil

}

// Visit a parse tree produced by VlangParser#ListExpressions.
func (v *ReplVisitor) VisitListExpressions(ctx *parser.ListExpressionsContext) interface{} {
	return nil
}

func (v *ReplVisitor) VisitIntType(ctx *parser.IntTypeContext) interface{} {
	return "int"
}

func (v *ReplVisitor) VisitStringType(ctx *parser.StringTypeContext) interface{} {
	return "string"
}

// Visit a parse tree produced by VlangParser#fn_atoi.
func (v *ReplVisitor) VisitFn_atoi(ctx *parser.Fn_atoiContext) interface{} {
	// Obtener la expresión dentro del paréntesis
	expr := ctx.Expresion()
	if expr == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "atoi requiere un argumento")
		return nil
	}

	// Evaluar la expresión
	result := v.Visit(expr)
	if result == nil {
		v.ErrorTable.NewSemanticError(expr.GetStart(), "Expresión inválida en atoi")
		return nil
	}

	// Verificar que sea un string
	strVal, ok := result.(*value.StringValue)
	if !ok {
		v.ErrorTable.NewSemanticError(expr.GetStart(),
			fmt.Sprintf("atoi esperaba un string, obtuvo %T", result))
		return nil
	}

	// Convertir el string a int
	intVal, err := strconv.Atoi(strVal.InternalValue)
	if err != nil {
		v.ErrorTable.NewSemanticError(expr.GetStart(),
			fmt.Sprintf("No se pudo convertir '%s' a int: %v", strVal.InternalValue, err))
		return nil
	}

	return &value.IntValue{
		InternalValue: intVal,
	}
	return nil
}

// Visit a parse tree produced by VlangParser#fn_parsefloat.
func (v *ReplVisitor) VisitFn_parsefloat(ctx *parser.Fn_parsefloatContext) interface{} {
	// Obtener la expresión dentro del paréntesis
	expr := ctx.Expresion()
	if expr == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "parse_float requiere un argumento")
		return nil
	}

	// Evaluar la expresión
	result := v.Visit(expr)
	if result == nil {
		v.ErrorTable.NewSemanticError(expr.GetStart(), "Expresión inválida en parse_float")
		return nil
	}

	// Verificar que sea un string
	strVal, ok := result.(*value.StringValue)
	if !ok {
		v.ErrorTable.NewSemanticError(expr.GetStart(),
			fmt.Sprintf("parse_float esperaba un string, obtuvo %T", result))
		return nil
	}

	// Convertir el string a float64
	floatVal, err := strconv.ParseFloat(strVal.InternalValue, 64)
	println("floatval ------------------------", floatVal)
	if err != nil {
		v.ErrorTable.NewSemanticError(expr.GetStart(),
			fmt.Sprintf("No se pudo convertir '%s' a float64: %v", strVal.InternalValue, err))
		return nil
	}

	return &value.FloatValue{
		InternalValue: floatVal,
	}

}

// Visit a parse tree produced by VlangParser#fn_typeOF.
func (v *ReplVisitor) VisitFn_typeOF(ctx *parser.Fn_typeOFContext) interface{} {
	// Obtener la expresión dentro del paréntesis
	expr := ctx.Expresion()
	if expr == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "typeof requiere un argumento")
		return nil
	}

	// Evaluar la expresión
	result := v.Visit(expr)
	if result == nil {
		v.ErrorTable.NewSemanticError(expr.GetStart(), "Expresión inválida en typeof")
		return nil
	}

	// Obtener el tipo del valor
	var typeStr string
	switch val := result.(type) {
	case value.IVOR:
		typeStr = val.Type()
	default:
		typeStr = fmt.Sprintf("%T", result)
	}

	return &value.StringValue{
		InternalValue: typeStr,
	}
	return nil
}

func (v *ReplVisitor) VisitIndexAcceso(ctx *compiler.IndexAccesoContext) interface{} {

	// Obtener el nombre del slice
	sliceName := ctx.Id_pattern().GetText()

	// Obtener la variable del slice del ScopeTrace
	sliceVar := v.ScopeTrace.GetVariable(sliceName)
	if sliceVar == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Slice '"+sliceName+"' no encontrado")
		return nil
	}

	// Verificar que la variable sea efectivamente un slice
	sliceValue, ok := sliceVar.Value.(*value.SliceValue)
	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "La variable '"+sliceName+"' no es un slice")
		return nil
	}

	// Evaluar la expresión del índice
	indexExpr := v.Visit(ctx.Expresion())
	if indexExpr == nil {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Expresión de índice inválida")
		return nil
	}

	indexValue, ok := indexExpr.(*value.IntValue)
	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "El índice debe ser un entero")
		return nil
	}

	// Verificar que el índice esté dentro de los límites
	switch slice := sliceValue.InternalValue.(type) {
	case []int:
		if indexValue.InternalValue < 0 || indexValue.InternalValue >= len(slice) {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "Índice fuera de rango")
			return nil
		}
		return &value.IntValue{InternalValue: slice[indexValue.InternalValue]}

	case []float64:
		if indexValue.InternalValue < 0 || indexValue.InternalValue >= len(slice) {
			v.ErrorTable.NewSemanticError(ctx.GetStart(), "Índice fuera de rango")
			return nil
		}
		return &value.FloatValue{InternalValue: slice[indexValue.InternalValue]}

	// Agregar más casos para otros tipos de slices según sea necesario

	default:
		v.ErrorTable.NewSemanticError(ctx.GetStart(), "Tipo de slice no soportado")
		return nil
	}
	return nil
}

// Visit a parse tree produced by VlangParser#fn_indexOf.
func (v *ReplVisitor) VisitFn_indexOf(ctx *parser.Fn_indexOfContext) interface{} {
	fmt.Println("------------------------------- ENTRANDO A LA FUNCION DEL INDEX")

	// Obtener el slice (primer argumento)
	sliceExpr := ctx.Expresion(0)
	sliceValue := v.Visit(sliceExpr)

	if sliceValue == nil {
		v.ErrorTable.NewSemanticError(sliceExpr.GetStart(), "Expresión de slice inválida")
		return -1
	}

	// Obtener el valor a buscar (segundo argumento)
	searchExpr := ctx.Expresion(1)
	searchValue := v.Visit(searchExpr)

	if searchValue == nil {
		v.ErrorTable.NewSemanticError(searchExpr.GetStart(), "Expresión de búsqueda inválida")
		return -1
	}

	// Verificar que sea un slice
	slice, ok := sliceValue.(*value.SliceValue)
	if !ok {
		v.ErrorTable.NewSemanticError(sliceExpr.GetStart(), "El primer argumento debe ser un slice")
		return -1
	}

	// Buscar el elemento según el tipo de slice
	switch actualSlice := slice.InternalValue.(type) {
	case []int:
		searchInt, ok := searchValue.(*value.IntValue)
		if !ok {
			v.ErrorTable.NewSemanticError(searchExpr.GetStart(),
				"Tipo incompatible: se esperaba int para buscar en slice de int")
			return -1
		}

		for i, val := range actualSlice {
			if val == searchInt.InternalValue {
				return i
			}
		}

	case []float64:
		searchFloat, ok := searchValue.(*value.FloatValue)
		if !ok {
			v.ErrorTable.NewSemanticError(searchExpr.GetStart(),
				"Tipo incompatible: se esperaba float para buscar en slice de float")
			return -1
		}

		for i, val := range actualSlice {
			if val == searchFloat.InternalValue {
				return i
			}
		}

	// Añadir más tipos según sea necesario
	default:
		v.ErrorTable.NewSemanticError(sliceExpr.GetStart(),
			fmt.Sprintf("Tipo de slice no soportado: %T", actualSlice))
		return -1
	}

	// Si no se encontró
	return -1
}

func (v *ReplVisitor) VisitFn_lenx(ctx *parser.Fn_lenxContext) interface{} {
	// Obtener la expresión del slice
	sliceExpr := ctx.Expresion(0)
	sliceValue := v.Visit(sliceExpr)

	if sliceValue == nil {
		v.ErrorTable.NewSemanticError(sliceExpr.GetStart(), "Expresión inválida en len()")
		return 0
	}

	// Verificar que sea un slice
	slice, ok := sliceValue.(*value.SliceValue)
	if !ok {
		v.ErrorTable.NewSemanticError(sliceExpr.GetStart(),
			"len() solo puede usarse con slices")
		return 0
	}

	// Calcular la longitud según el tipo de slice
	switch actualSlice := slice.InternalValue.(type) {
	case []int:
		return len(actualSlice)
	case []float64:
		return len(actualSlice)
	case []string:
		return len(actualSlice)
	case []bool:
		return len(actualSlice)
	// Añadir más tipos según sea necesario
	default:
		v.ErrorTable.NewSemanticError(sliceExpr.GetStart(),
			fmt.Sprintf("Tipo de slice no soportado: %T", actualSlice))
		return 0
	}
}

// Visit a parse tree produced by VlangParser#fn_join.
func (v *ReplVisitor) VisitFn_join(ctx *parser.Fn_joinContext) interface{} {

	// Obtener el slice de strings
	sliceValue := v.Visit(ctx.Expresion(0))
	separatorValue := v.Visit(ctx.Expresion(1))

	// Verificar tipos
	slice, ok1 := sliceValue.(*value.SliceValue)
	separator, ok2 := separatorValue.(*value.StringValue)

	if !ok1 || !ok2 {
		v.ErrorTable.NewSemanticError(ctx.GetStart(),
			"join() requiere: ([]string, string)")
		return ""
	}

	// Verificar que sea slice de strings
	strSlice, ok := slice.InternalValue.([]string)
	if !ok {
		v.ErrorTable.NewSemanticError(ctx.GetStart(),
			"El primer argumento debe ser []string")
		return ""
	}

	// Unir los strings con el separador
	result := strings.Join(strSlice, separator.InternalValue)
	return &value.StringValue{InternalValue: result}
}

// Visit a parse tree produced by VlangParser#funcionexpre.
func (v *ReplVisitor) VisitFuncionexpre(ctx *parser.FuncionexpreContext) interface{} {

	fmt.Println("🔍 Visitando llamada a función como expresión:", ctx.GetText())
	// Una llamada a función como expresión es esencialmente lo mismo que una llamada a función normal
	return v.Visit(ctx.Func_call())
}
