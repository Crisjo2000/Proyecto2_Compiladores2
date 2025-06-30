package arm

import (
	"fmt"
	compiler "main/parser"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type ArmVisitor struct {
	Generator   *ArmGenerator
	SymbolTable *SymbolTable
}

func NewArmVisitor() *ArmVisitor {
	return &ArmVisitor{
		Generator: &ArmGenerator{
			instructions: []string{},
			stdLib:       NewStandardLibrary(),
		},
		SymbolTable: NewSymbolTable(),
	}
}

// VisitPrograma es el punto de entrada para la generación de código.
func (v *ArmVisitor) VisitPrograma(ctx *compiler.ProgramaContext) interface{} {
	v.Generator.instructions = append(v.Generator.instructions, ".text", ".global main")

	// Primero, visita todas las declaraciones de funciones para registrarlas
	for _, stmt := range ctx.AllStmt() {
		if fnDecl, ok := stmt.GetChild(0).(*compiler.FuncDeclContext); ok {
			Walk(v, fnDecl)
		}
	}

	// Comprobar si existe la función 'main'
	mainFunc := v.SymbolTable.Resolve("main")
	if mainFunc == nil || mainFunc.Kind != FunctionKind {
		fmt.Println("Error: no se encontró la función 'main'.")
	} else {
		startCode := []string{
			".global _start",
			"_start:",
			"   bl main",
			"   mov x0, #0",
			"   mov x8, #93",
			"   svc #0",
		}
		v.Generator.instructions = append(startCode, v.Generator.instructions...)
	}

	return nil
}

func (v *ArmVisitor) VisitStmt(ctx *compiler.StmtContext) interface{} {
	// Visitar todos los hijos del StmtContext
	for i := 0; i < ctx.GetChildCount(); i++ {
		child := ctx.GetChild(i)
		if parseTree, ok := child.(antlr.ParseTree); ok {
			Walk(v, parseTree)
		}
	}
	return nil
}

func (v *ArmVisitor) VisitPrintlnStmt(ctx *compiler.PrintlnStmtContext) interface{} {
	// cod agregado 28/06
	v.Generator.Comment("Print statement")

	for _, exp := range ctx.AllExpresion() {
		Walk(v, exp)
	}

	// Obtener el objeto de la pila
	var value StackObject
	obj := v.Generator.TopObject()

	// Pop con el registro correcto según el tipo
	switch obj.Type {
	case IntType:
		value = v.Generator.PopObject(X0)
		v.Generator.PrintInteger(X0)
	case StringType:
		value = v.Generator.PopObject(X0)
		v.Generator.PrintString(X0)
	case FloatType:
		value = v.Generator.PopObject(D0)
		v.Generator.PrintFloat()
	case BoolType:
		value = v.Generator.PopObject(X0)
		v.Generator.PrintBool()
	default:
		fmt.Println("⚠️ Tipo no soportado para impresión:", obj.Type)
	}

	fmt.Println(value)
	return nil

	// FIN agre

	// cod comentado 28/06
	/*fmt.Println("🔍 Visitando println statement:", ctx.GetText())

	v.Generator.Comment("Print statement")
	v.Generator.Comment("Visiting expresions")

	// Evaluar cada expresión (por si acaso hay más de una, aunque normalmente es una sola)
	for _, exp := range ctx.AllExpresion() {
		fmt.Println("---🔍 Visitando expresión:", exp.GetText())
		v.Visit(exp)
	}

	v.Generator.Comment("Popping value to print")

	// Verifica el tipo del objeto en la cima de la pila
	isDouble := v.Generator.TopObject().Type == FloatType

	// Pop con el registro correcto según el tipo
	var value StackObject
	if isDouble {
		value = v.Generator.PopObject(D0)
	} else {
		value = v.Generator.PopObject(X0)
	}

	// Generar la instrucción de impresión según el tipo
	switch value.Type {
	case IntType:
		v.Generator.PrintInteger(X0)
	case StringType:
		v.Generator.PrintString(X0)
	case FloatType:
		v.Generator.PrintFloat()
	case BoolType:
		v.Generator.PrintBool()

	default:
		fmt.Println("⚠️ Tipo no soportado para impresión:", value.Type)
	}

	return nil*/
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
	fmt.Println("🔍 Visitando Id:", ctx.GetText())
	id := ctx.GetText()

	// 1. Verificar que la variable exista en la tabla de símbolos
	symbol := v.SymbolTable.Resolve(id)
	if symbol == nil {
		fmt.Printf("error semántico: la variable '%s' no ha sido declarada\n", id)
		// Si no se encuentra, no podemos generar código para cargarla.
		// Empujamos un objeto "venenoso" o nulo para indicar el error.
		// Esto es una simplificación; un sistema de errores real detendría la compilación.
		v.Generator.PushObject(StackObject{Type: -1}) // Usamos un tipo inválido
		return nil
	}

	// 2. Si existe, buscar su valor en la pila del generador
	offset, obj, found := v.Generator.FindObject(id)
	if !found {
		// Error de lógica interna si el símbolo existe pero el objeto no.
		fmt.Printf("Error interno: No se encontró el objeto en la pila para '%s'\n", id)
		v.Generator.PushObject(StackObject{Type: -1})
		return nil
	}

	// 3. Generar código para cargar el valor de la variable en un registro
	v.Generator.Comment("Load variable " + id)
	if obj.Type == FloatType {
		v.Generator.Fldr(D0, SP, offset) // Cargar valor desde la pila
		v.Generator.Push(D0)             // Ponerlo en la cima de la pila
	} else {
		v.Generator.Ldr(X0, SP, offset) // Cargar valor desde la pila
		v.Generator.Push(X0)            // Ponerlo en la cima de la pila
	}
	v.Generator.PushObject(obj) // Empujar el objeto virtual para mantener la consistencia

	return nil
}

func (v *ArmVisitor) VisitValorExpresion(ctx *compiler.ValorExpresionContext) interface{} {
	fmt.Println("🔍 Visitando ValorExpresion:", ctx.GetText())
	valor := ctx.Valor()

	return Walk(v, valor)
}

func (v *ArmVisitor) VisitParentesisExpresion(ctx *compiler.ParentesisExpresionContext) interface{} {
	// Simplemente visita la expresión interna
	return Walk(v, ctx.Expresion())
}

func (v *ArmVisitor) VisitBinaryExp(ctx *compiler.BinaryExpContext) interface{} {

	//--------------------------------Aritméticos-----------------------------------
	if ctx.GetOp().GetText() == "+" || ctx.GetOp().GetText() == "-" ||
		ctx.GetOp().GetText() == "*" || ctx.GetOp().GetText() == "/" {
		v.Generator.Comment("Aritmetic operand")
		var operation = ctx.GetOp().GetText()
		v.Generator.Comment("Visiting left expression: ")
		Walk(v, ctx.Expresion(0))
		v.Generator.Comment("Visiting right expression: ")
		Walk(v, ctx.Expresion(1))

		v.Generator.Comment("popping operands")
		var isRightFloat = v.Generator.TopObject().Type == FloatType
		var left StackObject
		if isRightFloat {
			v.Generator.PopObject(D0)
		} else {
			v.Generator.PopObject(X0)
		}

		var isLeftFloat = v.Generator.TopObject().Type == FloatType
		if isLeftFloat {
			left = v.Generator.PopObject(D1)
		} else {
			left = v.Generator.PopObject(X1)
		}

		if isLeftFloat || isRightFloat {
			v.Generator.Comment("Performing float operation")

			if !isLeftFloat {
				v.Generator.Scvtf(D1, X1)

			}
			if !isRightFloat {
				v.Generator.Scvtf(D0, X0)
			}

			switch operation {
			case "+":
				v.Generator.Fadd(D0, D0, D1)
			case "-":
				v.Generator.Fsub(D0, D0, D1)
			case "*":
				v.Generator.Fmul(D0, D0, D1)
			case "/":
				v.Generator.Fdiv(D0, D0, D1)
			}

			v.Generator.Push(D0)
			if isLeftFloat {
				v.Generator.PushObject(v.Generator.FloatObject())
			} else {
				v.Generator.PushObject(v.Generator.IntObject())
			}
			return nil
		}

		switch operation {
		case "+":
			v.Generator.Add(X0, X0, X1)
		case "-":
			v.Generator.Sub(X0, X0, X1)
		case "*":
			v.Generator.Mul(X0, X0, X1)
		case "/":
			v.Generator.Div(X0, X0, X1)
		}
		v.Generator.Comment("Pushing result")
		v.Generator.Push(X0)
		v.Generator.PushObject(v.Generator.CloneObject(left))

		//-------------------------------Relacionales-----------------------------------
	} else if ctx.GetOp().GetText() == "==" || ctx.GetOp().GetText() == "!=" || ctx.GetOp().GetText() == "<" || ctx.GetOp().GetText() == ">" || ctx.GetOp().GetText() == "<=" || ctx.GetOp().GetText() == ">=" {
		Walk(v, ctx.Expresion(0))
		Walk(v, ctx.Expresion(1))

		right := v.Generator.TopObject()
		left := v.Generator.stack[len(v.Generator.stack)-2]
		op := ctx.GetOp().GetText()

		// Validación de tipos
		isRelational := op == ">" || op == "<" || op == ">=" || op == "<="
		if isRelational && (left.Type == StringType || left.Type == BoolType || right.Type == StringType || right.Type == BoolType) {
			fmt.Printf("Error de tipo: El operador '%s' no se puede aplicar a tipos string o bool.\n", op)
			return nil
		}
		if left.Type != right.Type && (left.Type == StringType || left.Type == BoolType || right.Type == StringType || right.Type == BoolType) {
			fmt.Printf("Error de tipo: No se puede comparar tipo %v con %v usando '%s'.\n", left.Type, right.Type, op)
			return nil
		}

		// Generación de código
		if left.Type == FloatType || right.Type == FloatType {
			// Comparación de flotantes
			v.Generator.PopObject(D0) // right
			v.Generator.PopObject(D1) // left
			if left.Type == IntType {
				v.Generator.Scvtf(D1, X1)
			}
			if right.Type == IntType {
				v.Generator.Scvtf(D0, X0)
			}
			v.Generator.Fcmp(D1, D0)
		} else {
			// Comparación de enteros, booleanos o strings (punteros)
			v.Generator.PopObject(X1) // right
			v.Generator.PopObject(X0) // left
			v.Generator.Cmp(X0, X1)
		}

		trueLabel := v.Generator.GetLabel()
		endLabel := v.Generator.GetLabel()
		switch op {
		case "==":
			v.Generator.Beq(trueLabel)
		case "!=":
			v.Generator.Bne(trueLabel)
		case "<":
			v.Generator.Blt(trueLabel)
		case ">":
			v.Generator.Bgt(trueLabel)
		case "<=":
			v.Generator.Ble(trueLabel)
		case ">=":
			v.Generator.Bge(trueLabel)
		}
		v.Generator.Mov(X0, 0)
		v.Generator.B(endLabel)
		v.Generator.SetLabel(trueLabel)
		v.Generator.Mov(X0, 1)
		v.Generator.SetLabel(endLabel)
		v.Generator.Push(X0)
		v.Generator.PushObject(v.Generator.BoolObject())
	} else if ctx.GetOp().GetText() == "&&" || ctx.GetOp().GetText() == "||" {
		// Lógicos (a implementar)
	}
	return nil
}

func (v *ArmVisitor) VisitDeclAssignInferred(ctx *compiler.DeclAssignInferredContext) interface{} {
	varName := ctx.ID().GetText()
	v.Generator.Comment("Inferred Declaration and Assignment for " + varName)

	// Visitar la expresión para obtener el valor
	Walk(v, ctx.Expresion())

	// Obtener el tipo del valor de la pila y almacenarlo
	val := v.Generator.PopObject(X0) // El valor está en X0 (o D0)

	// Almacenar el símbolo
	v.SymbolTable.Define(varName, val.Type, FunctionKind) // Asumiendo que `val.Type` es correcto

	// Mover el objeto en la pila
	v.Generator.PushObject(StackObject{
		Type: val.Type,
		Id:   &varName,
	})

	return nil
}

func (v *ArmVisitor) VisitDeclAssignExplicit(ctx *compiler.DeclAssignExplicitContext) interface{} {
	varName := ctx.ID().GetText()
	typeName := ctx.TYPE().GetText()

	v.Generator.Comment("Explicit Declaration and Assignment for " + varName)

	// 1. Visitar la expresión para poner su valor en la cima de la pila
	Walk(v, ctx.Expresion())

	// 2. Obtener el valor de la pila (y su tipo real)
	valueOnStack := v.Generator.PopObject(X0)

	// 3. Determinar el tipo de la variable según la declaración
	declaredType, ok := typeStringToStackObjectType(typeName)
	if !ok {
		// Manejar error: tipo desconocido
		fmt.Printf("Error: tipo desconocido '%s' para la variable '%s'\n", typeName, varName)
		return nil
	}

	// 4. Verificar si los tipos son compatibles (simplificado por ahora)
	if valueOnStack.Type != declaredType {
		// Aquí se podría manejar coerción de tipos (ej. int a float)
		// Por ahora, solo informamos.
		fmt.Printf("Advertencia: asignando tipo %v a variable de tipo %v\n", valueOnStack.Type, declaredType)
	}

	// 5. Definir el símbolo en la tabla de símbolos con el tipo declarado
	v.SymbolTable.Define(varName, declaredType, VariableKind)

	// 6. Crear y empujar el objeto a la pila con el tipo declarado
	//    El valor ya está en el registro correcto (X0 o D0) desde PopObject
	if declaredType == FloatType {
		v.Generator.Push(D0)
	} else {
		v.Generator.Push(X0)
	}

	v.Generator.PushObject(StackObject{
		Type:   declaredType,
		Length: 8, // Asumimos 8 bytes por ahora
		Depth:  v.Generator.depth,
		Id:     &varName,
	})

	return nil
}

func (v *ArmVisitor) VisitDirectAssign(ctx *compiler.DirectAssignContext) interface{} {
	varName := ctx.ID().GetText()

	// 1. Verificar que la variable exista
	symbol := v.SymbolTable.Resolve(varName)
	if symbol == nil {
		fmt.Printf("Error semántico: la variable '%s' no ha sido declarada antes de la asignación.\n", varName)
		return nil
	}

	// 2. Encontrar la posición de la variable en la pila
	offset, _, found := v.Generator.FindObject(varName)
	if !found {
		fmt.Printf("Error interno: no se encontró el objeto en la pila para '%s'\n", varName)
		return nil
	}

	v.Generator.Comment("Direct assignment to " + varName)

	// 3. Evaluar la expresión del lado derecho
	Walk(v, ctx.Expresion())

	// 4. Sacar el valor de la expresión de la pila a un registro
	valueOnStack := v.Generator.PopObject(X0) // O D0 si es float

	// 5. Verificar compatibilidad de tipos (simplificado)
	if symbol.Type != valueOnStack.Type {
		fmt.Printf("Advertencia: asignando tipo %v a variable de tipo %v\n", valueOnStack.Type, symbol.Type)
	}

	// 6. Almacenar el nuevo valor desde el registro a la posición de la variable en la pila
	if symbol.Type == FloatType {
		v.Generator.Fstr(D0, SP, offset)
	} else {
		v.Generator.Str(X0, SP, offset)
	}

	return nil
}

// codigo para hacer los if  if
func (v *ArmVisitor) VisitIfStmt(ctx *compiler.IfStmtContext) interface{} {
	v.Generator.Comment("If statement")

	endIfLabel := v.Generator.NewLabel("endif")

	// Procesar la cadena de if/else if/else
	v.VisitIfChain(ctx.IfChain(), endIfLabel)

	v.Generator.AddLabel(endIfLabel)
	v.Generator.Comment("End of if statement")

	return nil
}

func (v *ArmVisitor) VisitIfChain(ctx *compiler.IfChainContext, endLabel string) interface{} {
	// Cada 'if' o 'else if' tiene una condición y un bloque.
	// Si la condición es falsa, saltamos a la siguiente cláusula.
	nextClauseLabel := v.Generator.NewLabel("next_clause")

	// 1. Evaluar la condición
	v.Generator.Comment("if condition")
	Walk(v, ctx.Expresion())
	v.Generator.PopObject(X0)
	v.Generator.Cmp(X0, 0)           // Comparar el resultado booleano (0 o 1) con 0
	v.Generator.Beq(nextClauseLabel) // Si es igual a 0 (falso), saltar

	// 2. Si la condición es verdadera, ejecutar el bloque
	v.Generator.Comment("if block")
	Walk(v, ctx.BlockStmt())
	v.Generator.B(endLabel) // Salir de toda la estructura if

	// 3. Etiqueta para la siguiente cláusula
	v.Generator.AddLabel(nextClauseLabel)

	// 4. Procesar el 'else if' o 'else' si existe
	if ctx.ElseStmt() != nil {
		Walk(v, ctx.ElseStmt())
	}

	return nil
}

func (v *ArmVisitor) VisitElseStmt(ctx *compiler.ElseStmtContext) interface{} {
	if ctx.IfChain() != nil {
		// Es un 'else if'
		// El label de finalización se pasa recursivamente
		endLabel := v.Generator.instructions[len(v.Generator.instructions)-1] // Hack para obtener el último label
		v.VisitIfChain(ctx.IfChain(), endLabel)
	} else {
		// Es un 'else' final
		v.Generator.Comment("else block")
		Walk(v, ctx.BlockStmt())
	}
	return nil
}

// Stubs para metodos no implementados
func (v *ArmVisitor) VisitPrintStmt(ctx *compiler.PrintStmtContext) interface{} {
	return nil
}
func (v *ArmVisitor) VisitFuncCall(ctx *compiler.FuncCallContext) interface{} {
	funcName := ctx.ID().GetText()

	// Caso especial para la función `print`
	if funcName == "print" {
		// Invertir el orden de evaluación de argumentos para la pila
		args := ctx.AllExpresion()
		for i := len(args) - 1; i >= 0; i-- {
			Walk(v, args[i])
		}

		// Procesar cada argumento
		for range args {
			obj := v.Generator.TopObject()
			switch obj.Type {
			case IntType:
				v.Generator.PopObject(X0)
				v.Generator.PrintInteger(X0)
			case StringType:
				v.Generator.PopObject(X0)
				v.Generator.PrintString(X0)
			case FloatType:
				v.Generator.PopObject(D0)
				v.Generator.PrintFloat()
			case BoolType:
				v.Generator.PopObject(X0)
				v.Generator.PrintBool()
			default:
				// Pop de un tipo desconocido para limpiar la pila
				v.Generator.PopObject(X0)
			}
		}
		// Imprimir nueva línea al final de `print`
		v.Generator.PrintString("\\n") // Asumiendo que PrintString puede manejar secuencias de escape
		return nil
	}

	// Lógica para funciones definidas por el usuario
	symbol := v.SymbolTable.Resolve(funcName)
	if symbol == nil || symbol.Kind != FunctionKind {
		fmt.Printf("Error: llamada a función no definida '%s'\n", funcName)
		return nil
	}

	v.Generator.Comment("Calling function " + funcName)
	v.Generator.Bl(funcName)

	return nil
}

func (v *ArmVisitor) VisitFuncDecl(ctx *compiler.FuncDeclContext) interface{} {
	funcName := ctx.ID().GetText()
	v.Generator.Comment("Function declaration: " + funcName)
	v.Generator.AddLabel(funcName)

	// Crear una entrada en la tabla de símbolos para la función
	symbol := &Symbol{
		Name:       funcName,
		Kind:       FunctionKind,
		Type:       -1, // El tipo de retorno se determinará más adelante
		Parameters: []*Symbol{},
	}

	// Procesar parámetros
	if ctx.FuncParams() != nil {
		for _, param := range ctx.FuncParams().AllID() {
			// Esto es una simplificación. Necesitaríamos el tipo del parámetro.
			// Por ahora, asumimos que todos los parámetros son de un tipo genérico.
			paramSymbol := &Symbol{
				Name: param.GetText(),
				Kind: VariableKind,
				Type: IntType, // Simplificación: asumir int
			}
			symbol.Parameters = append(symbol.Parameters, paramSymbol)
		}
	}

	// Procesar tipo de retorno (si existe)
	if ctx.TYPE() != nil {
		returnType, ok := typeStringToStackObjectType(ctx.TYPE().GetText())
		if ok {
			symbol.Type = returnType
		}
	}

	// Registrar la función en la tabla de símbolos
	v.SymbolTable.Define(funcName, symbol.Type, FunctionKind)
	v.SymbolTable.symbols[funcName] = symbol // Asegurarse de que toda la info se guarde

	// Guardar el contexto de la función para la visita del cuerpo
	v.SymbolTable.EnterScope()
	for _, p := range symbol.Parameters {
		v.SymbolTable.Define(p.Name, p.Type, VariableKind)
	}

	// Visitar el cuerpo de la función
	Walk(v, ctx.BlockStmt())

	v.SymbolTable.ExitScope()

	// Añadir prologo y epílogo
	v.Generator.prologue(funcName)
	// El código del cuerpo se inserta aquí
	v.Generator.epilogue()

	return nil
}

func (v *ArmVisitor) VisitBlockStmt(ctx *compiler.BlockStmtContext) interface{} {
	v.Generator.Comment("Entering a new block/scope")
	v.SymbolTable.EnterScope()

	// Visitar todas las sentencias dentro del bloque
	for _, stmt := range ctx.AllStmt() {
		Walk(v, stmt)
	}

	v.SymbolTable.ExitScope()
	v.Generator.Comment("Exiting block/scope")
	return nil
}

// typeStringToStackObjectType es una función helper para convertir el string de tipo a nuestro enum
func typeStringToStackObjectType(typeName string) (StackObjectType, bool) {
	switch typeName {
	case "int":
		return IntType, true
	case "float":
		return FloatType, true
	case "string":
		return StringType, true
	case "bool":
		return BoolType, true
	case "void":
		return VoidType, true
	default:
		return -1, false
	}
}

func (v *ArmVisitor) VisitSwitchStmt(ctx *compiler.SwitchStmtContext) interface{} {
	v.Generator.Comment("Switch statement")
	endLabel := v.Generator.NewLabel("end_switch")
	nextCaseLabel := ""

	// Evaluar la expresión del switch
	Walk(v, ctx.ID())
	switchExprObj := v.Generator.PopObject(X0)
	v.Generator.MovReg(X9, X0)

	// Generar etiquetas y comparaciones para cada case
	for i, caseClause := range ctx.AllCase_clause() {
		caseLabels[i] = v.Generator.GetLabel()
		Walk(v, caseClause.(*compiler.CaseClauseContext).Expresion(0))
		caseValObj := v.Generator.PopObject(X1)

		if switchExprObj.Type != caseValObj.Type {
			fmt.Printf("Error de tipo: no se puede comparar la expresión del switch (tipo %v) con el caso (tipo %v)\n", switchExprObj.Type, caseValObj.Type)
			continue // Saltar este caso, pero no detener el switch
		}

		v.Generator.Cmp(X9, X1)
		v.Generator.Beq(caseLabels[i])
	}

	// Saltar al default si existe, si no, al final
	if ctx.Default_clause() != nil {
		v.Generator.B(defaultLabel)
	} else {
		v.Generator.B(endLabel)
	}

	// Generar el código para los bloques de cada case
	for i, caseClause := range ctx.AllCase_clause() {
		v.Generator.SetLabel(caseLabels[i])
		v.VisitCaseClause(caseClause.(*compiler.CaseClauseContext), X9, endLabel)
	}

	// Generar el código para el bloque default
	if ctx.Default_clause() != nil {
		v.Generator.SetLabel(defaultLabel)
		v.VisitDefaultClause(ctx.Default_clause().(*compiler.DefaultClauseContext))
		// El default también tiene un break implícito
		v.Generator.B(endLabel)
	}

	// Etiqueta final
	v.Generator.SetLabel(endLabel)
	return nil
}

func (v *ArmVisitor) VisitCaseClause(ctx *compiler.CaseClauseContext, switchExprReg string, endLabel string) interface{} {
	nextCase := v.Generator.NewLabel("next_case")
	v.Generator.Comment("case")

	// Walk(v, ctx.Expresion())
	v.Generator.PopObject(X1)
	v.Generator.Cmp(X0, X1)
	v.Generator.Bne(nextCase) // No es igual, prueba el siguiente caso

	for _, stmt := range ctx.AllStmt() {
		Walk(v, stmt)
	}

	v.Generator.B(endLabel)
	v.Generator.AddLabel(nextCase)

}

func (v *ArmVisitor) VisitDefaultClause(ctx *compiler.DefaultClauseContext) interface{} {
	v.Generator.Comment("default case")
	for _, stmt := range ctx.AllStmt() {
		Walk(v, stmt)
	}
}

func (v *ArmVisitor) VisitWhileStmt(ctx *compiler.WhileStmtContext) interface{} {
	startLabel := v.Generator.NewLabel("while_start")
	endLabel := v.Generator.NewLabel("while_end")

	v.Generator.AddLabel(startLabel)
	v.Generator.Comment("While condition")
	Walk(v, ctx.Expresion())
	v.Generator.PopObject(X0)
	v.Generator.Cmp(X0, 0)
	v.Generator.Beq(endLabel) // Si la condición es falsa, salir del bucle

	v.Generator.Comment("While block")
	Walk(v, ctx.BlockStmt())
	v.Generator.B(startLabel) // Volver al inicio del bucle

	v.Generator.AddLabel(endLabel)
	v.Generator.Comment("End of while statement")

	return nil
}
func (v *ArmVisitor) VisitClasicFor(ctx *compiler.ClasicForContext) interface{} {
	v.SymbolTable.EnterScope()

	// 1. Inicialización
	v.Generator.Comment("For loop initialization")
	if ctx.DeclAssignExplicit() != nil {
		Walk(v, ctx.DeclAssignExplicit())
	} else if ctx.DirectAssign() != nil {
		Walk(v, ctx.DirectAssign())
	}

	startLabel := v.Generator.NewLabel("for_start")
	endLabel := v.Generator.NewLabel("for_end")

	v.Generator.AddLabel(startLabel)

	// 2. Condición
	v.Generator.Comment("For loop condition")
	Walk(v, ctx.Expresion(0)) // La primera expresión es la condición
	v.Generator.PopObject(X0)
	v.Generator.Cmp(X0, 0)
	v.Generator.Beq(endLabel)

	// 3. Cuerpo del bucle
	v.Generator.Comment("For loop block")
	Walk(v, ctx.BlockStmt())

	// 4. Actualización
	v.Generator.Comment("For loop update")
	Walk(v, ctx.Expresion(1)) // La segunda expresión es la actualización
	v.Generator.PopObject(X0) // Descartar el resultado de la actualización de la pila

	v.Generator.B(startLabel)
	v.Generator.AddLabel(endLabel)

	v.SymbolTable.ExitScope()
	return nil
}

func (v *ArmVisitor) VisitIncremento(ctx *compiler.IncrementoContext) interface{} {
	varName := ctx.ID().GetText()
	offset, obj, found := v.Generator.FindObject(varName)
	if !found {
		fmt.Printf("Error: variable '%s' no encontrada para incremento\n", varName)
		return nil
	}

	v.Generator.Ldr(X0, SP, offset)
	v.Generator.Add(X0, X0, 1)
	v.Generator.Str(X0, SP, offset)
	v.Generator.Push(X0) // Dejar el valor actualizado en la pila
	v.Generator.PushObject(obj)
	return nil
}

func (v *ArmVisitor) VisitDecremento(ctx *compiler.DecrementoContext) interface{} {
	varName := ctx.ID().GetText()
	offset, obj, found := v.Generator.FindObject(varName)
	if !found {
		fmt.Printf("Error: variable '%s' no encontrada para decremento\n", varName)
		return nil
	}

	v.Generator.Ldr(X0, SP, offset)
	v.Generator.Sub(X0, X0, 1)
	v.Generator.Str(X0, SP, offset)
	v.Generator.Push(X0)
	v.Generator.PushObject(obj)
	return nil
}
