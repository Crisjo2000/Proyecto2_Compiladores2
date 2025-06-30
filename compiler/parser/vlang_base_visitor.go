// Code generated from parser/Vlang.g4 by ANTLR 4.13.1. DO NOT EDIT.

package compiler // Vlang
import "github.com/antlr4-go/antlr/v4"

type BaseVlangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseVlangVisitor) VisitPrograma(ctx *ProgramaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitBlockStmt(ctx *BlockStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitIfStmt(ctx *IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitIfChain(ctx *IfChainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitElseStmt(ctx *ElseStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitPrintlnStmt(ctx *PrintlnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitPrintStmt(ctx *PrintStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitWhileStmt(ctx *WhileStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitClasicFor(ctx *ClasicForContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitForStmt(ctx *ForStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitNumericRange(ctx *NumericRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitSwitchStmt(ctx *SwitchStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitCaseClause(ctx *CaseClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitDefaultClause(ctx *DefaultClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitDeclSlices(ctx *DeclSlicesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitFuncCall(ctx *FuncCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitFuncDecl(ctx *FuncDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitStructDecl(ctx *StructDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitBreakStmt(ctx *BreakStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitContinueStmt(ctx *ContinueStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitDirectAssign(ctx *DirectAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitDeclAssignInferred(ctx *DeclAssignInferredContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitDeclAssignExplicit(ctx *DeclAssignExplicitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitListExpressions(ctx *ListExpressionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitIdPattern(ctx *IdPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitUnarioExpresion(ctx *UnarioExpresionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitIndexAcceso(ctx *IndexAccesoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitValorExpresion(ctx *ValorExpresionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitFuncionExpresion(ctx *FuncionExpresionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitParentesisExpresion(ctx *ParentesisExpresionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitId(ctx *IdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitIncreDecreExp(ctx *IncreDecreExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitBinaryExp(ctx *BinaryExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitParamList(ctx *ParamListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitFuncParam(ctx *FuncParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitArgList(ctx *ArgListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitFuncArg(ctx *FuncArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitValorEntero(ctx *ValorEnteroContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitValorFloat(ctx *ValorFloatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitValorCadena(ctx *ValorCadenaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitValorBooleano(ctx *ValorBooleanoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitStructInstancia(ctx *StructInstanciaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitStructAttr(ctx *StructAttrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitIntType(ctx *IntTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitFloatType(ctx *FloatTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitStringType(ctx *StringTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitBoolType(ctx *BoolTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitCharType(ctx *CharTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitVoidType(ctx *VoidTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitIncremento(ctx *IncrementoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitDecremento(ctx *DecrementoContext) interface{} {
	return v.VisitChildren(ctx)
}
