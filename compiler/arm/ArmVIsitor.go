package arm

import (
	"fmt"
	"log"
	compiler "main/parser"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type ArmVisitor struct {
	compiler.BaseVlangVisitor               // ReplVisitor para manejar el contexto del REPL
	Generator                 *ArmGenerator // Generador de instrucciones ARM
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
	}
}

func (v *ArmVisitor) VisitPrograma(ctx *compiler.ProgramaContext) interface{} {

	fmt.Println("🔍 Visitando el programa en el ARMVISITOR")
	// Vamos a recorrer todos los statements y los visitamos
	for _, stmt := range ctx.AllStmt() {
		fmt.Println("---🔍 Visitando statement*:(text)", stmt.GetText())
		v.Visit(stmt)
	}

	return nil
}

func (v *ArmVisitor) VisitPrintlnStmt(ctx *compiler.PrintlnStmtContext) interface{} {
	fmt.Println("🔍 Visitando println statement:", ctx.GetText())

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

	//--------------------------------Aritméticos-----------------------------------
	if ctx.GetOp().GetText() == "+" || ctx.GetOp().GetText() == "-" ||
		ctx.GetOp().GetText() == "*" || ctx.GetOp().GetText() == "/" {
		v.Generator.Comment("Aritmetic operand")
		var operation = ctx.GetOp().GetText()
		v.Generator.Comment("Visiting left expression: ")
		v.Visit(ctx.Expresion(0))
		v.Generator.Comment("Visiting right expression: ")
		v.Visit(ctx.Expresion(1))

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
	} else if ctx.GetOp().GetText() == "==" || ctx.GetOp().GetText() == "!=" ||
		ctx.GetOp().GetText() == "<" || ctx.GetOp().GetText() == ">" ||
		ctx.GetOp().GetText() == "<=" || ctx.GetOp().GetText() == ">=" {

		v.Generator.Comment("Relational operation")
		var operation = ctx.GetOp().GetText()
		v.Generator.Comment("Visiting left expression: ")
		v.Visit(ctx.Expresion(0)) // Visitar expresión izquierda
		v.Generator.Comment("Visiting right expression: ")
		v.Visit(ctx.Expresion(1)) // Visitar expresión derecha
		v.Generator.Comment("Popping operands")

		var isRightFloat = v.Generator.TopObject().Type == FloatType
		if isRightFloat {
			v.Generator.PopObject(D1)
		} else {
			v.Generator.PopObject(X1)
		}

		var isLeftFloat = v.Generator.TopObject().Type == FloatType
		if isLeftFloat {
			v.Generator.PopObject(D0)
		} else {
			v.Generator.PopObject(X0)
		}

		if isLeftFloat || isRightFloat {
			return nil
		}

		v.Generator.Cmp(X1, X0)
		v.Generator.Comment("Setting condition flags")

		var truelabel = v.Generator.GetLabel()
		var endlabel = v.Generator.GetLabel()

		switch operation {
		case "<":
			v.Generator.Blt(truelabel)
		case "<=":
			v.Generator.Ble(truelabel)
		case ">":
			v.Generator.Bgt(truelabel)
		case ">=":
			v.Generator.Bge(truelabel)
		case "==":
			v.Generator.Beq(truelabel)
		case "!=":
			v.Generator.Bne(truelabel)
		default:
			fmt.Println("⚠️ Operación relacional no soportada:", operation)
			return nil
		}

		v.Generator.Mov(X0, 0)
		v.Generator.Push(X0)
		v.Generator.B(endlabel)
		v.Generator.SetLabel(truelabel)
		v.Generator.Mov(X0, 1)
		v.Generator.Push(X0)
		v.Generator.SetLabel(endlabel)

		v.Generator.PushObject(v.Generator.BoolObject())
		//-------------------------------Lógicos-----------------------------------
	} else if ctx.GetOp().GetText() == "&&" || ctx.GetOp().GetText() == "||" {
		v.Generator.Comment("Logical operation")
		var operation = ctx.GetOp().GetText()

		v.Generator.Comment("Visiting left expression: ")
		v.Visit(ctx.Expresion(0)) // Visitar expresión izquierda
		v.Generator.Comment("Visiting right expression: ")
		v.Visit(ctx.Expresion(1)) // Visitar expresión derecha

		v.Generator.Comment("Popping operands")
		v.Generator.PopObject(X1)
		//left := v.Generator.PopObject(X0)

		switch operation {
		case "&&":
		}
	}

	return nil
}

func (v *ArmVisitor) VisitDeclAssign(ctx *compiler.DeclAssignContext) interface{} {
	var varName = ctx.ID().GetText()
	v.Generator.Comment("Variable declaration: " + varName)

	v.Visit(ctx.Expresion())

	topObject := v.Generator.TopObject()

	obj := StackObject{
		Type:   topObject.Type, // O el tipo correcto según la expresión
		Length: 8,
		Depth:  v.Generator.depth, // o simplemente usar método
		Id:     nil,
	}

	v.Generator.PushObject(obj) // Pushea el objeto a la pila)
	v.Generator.TagObject(varName)

	return nil
}

func (v *ArmVisitor) VisitDirectAssign(ctx *compiler.DirectAssignContext) interface{} {
	assigne := ctx.Expresion()

	if assigne == assigne.(*compiler.ValorexpresionContext) {

		var varname = ctx.Id_pattern().GetText()

		v.Generator.Comment("Direct assignment to variable: " + varname)
		v.Visit(ctx.Expresion())
		var valueObject = v.Generator.PopObject(X0)

		var offset, varObject = v.Generator.GetObject(varname)

		v.Generator.Mov(X1, offset)
		v.Generator.Add(X1, SP, X1)
		v.Generator.Str(X0, X1, 0)

		varObject.Type = valueObject.Type

		v.Generator.Push(X0)
		v.Generator.PushObject(varObject)
	}
	return nil
}
