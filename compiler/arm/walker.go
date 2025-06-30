package arm

import (
	"fmt"
	"main/parser"

	"github.com/antlr4-go/antlr/v4"
)

func Walk(v *ArmVisitor, tree antlr.ParseTree) interface{} {
	if tree == nil {
		return nil
	}

	switch node := tree.(type) {
	case *antlr.ErrorNodeImpl:
		fmt.Printf("Error de sintaxis cerca de: '%s'\n", node.GetText())
		return nil
	case *antlr.TerminalNodeImpl:
		return nil // Ignorar nodos terminales
	case *parser.ProgramaContext:
		return v.VisitPrograma(node)
	case *parser.StmtContext:
		return v.VisitStmt(node)
	case *parser.PrintlnStmtContext:
		return v.VisitPrintlnStmt(node)
	case *parser.PrintStmtContext:
		return v.VisitPrintStmt(node)
	case *parser.ValorEnteroContext:
		return v.VisitValorEntero(node)
	case *parser.ValorFloatContext:
		return v.VisitValorFloat(node)
	case *parser.ValorCadenaContext:
		return v.VisitValorCadena(node)
	case *parser.ValorBooleanoContext:
		return v.VisitValorBoolean(node)
	case *parser.BinaryExpContext:
		return v.VisitBinaryExp(node)
	case *parser.ParentesisExpresionContext:
		return v.VisitParentesisExpresion(node)
	case *parser.FuncCallContext:
		return v.VisitFuncCall(node)
	case *parser.ValorExpresionContext:
		return v.VisitValorExpresion(node)
	case *parser.DeclAssignInferredContext:
		return v.VisitDeclAssignInferred(node)
	case *parser.DeclAssignExplicitContext:
		return v.VisitDeclAssignExplicit(node)
	case *parser.IdContext:
		return v.VisitId(node)
	case *parser.DirectAssignContext:
		return v.VisitDirectAssign(node)
	case *parser.IfStmtContext:
		return v.VisitIfStmt(node)
	case *parser.IfChainContext:
		// Este caso es especial, necesita argumentos extra
		// Se llamará directamente desde VisitIfStmt
		return nil
	case *parser.ElseStmtContext:
		return v.VisitElseStmt(node)
	case *parser.BlockStmtContext:
		return v.VisitBlockStmt(node)
	case *parser.FuncDeclContext:
		return v.VisitFuncDecl(node)
	case *parser.SwitchStmtContext:
		return v.VisitSwitchStmt(node)
	case *parser.CaseClauseContext:
		// Caso especial, se llama desde VisitSwitchStmt
		return nil
	case *parser.DefaultClauseContext:
		return v.VisitDefaultClause(node)
	case *parser.WhileStmtContext:
		return v.VisitWhileStmt(node)
	case *parser.ClasicForContext:
		return v.VisitClasicFor(node)
	case *parser.IncrementoContext:
		return v.VisitIncremento(node)
	case *parser.DecrementoContext:
		return v.VisitDecremento(node)
	default:
		fmt.Printf("⚠️ Tipo inesperado en Walk(): %T\n", tree)
		return nil
	}
}
