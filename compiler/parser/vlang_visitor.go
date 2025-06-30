// Code generated from parser/Vlang.g4 by ANTLR 4.13.1. DO NOT EDIT.

package compiler // Vlang
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by VlangParser.
type VlangVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by VlangParser#programa.
	VisitPrograma(ctx *ProgramaContext) interface{}

	// Visit a parse tree produced by VlangParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by VlangParser#BlockStmt.
	VisitBlockStmt(ctx *BlockStmtContext) interface{}

	// Visit a parse tree produced by VlangParser#IfStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by VlangParser#IfChain.
	VisitIfChain(ctx *IfChainContext) interface{}

	// Visit a parse tree produced by VlangParser#ElseStmt.
	VisitElseStmt(ctx *ElseStmtContext) interface{}

	// Visit a parse tree produced by VlangParser#PrintlnStmt.
	VisitPrintlnStmt(ctx *PrintlnStmtContext) interface{}

	// Visit a parse tree produced by VlangParser#PrintStmt.
	VisitPrintStmt(ctx *PrintStmtContext) interface{}

	// Visit a parse tree produced by VlangParser#WhileStmt.
	VisitWhileStmt(ctx *WhileStmtContext) interface{}

	// Visit a parse tree produced by VlangParser#ClasicFor.
	VisitClasicFor(ctx *ClasicForContext) interface{}

	// Visit a parse tree produced by VlangParser#ForStmt.
	VisitForStmt(ctx *ForStmtContext) interface{}

	// Visit a parse tree produced by VlangParser#NumericRange.
	VisitNumericRange(ctx *NumericRangeContext) interface{}

	// Visit a parse tree produced by VlangParser#SwitchStmt.
	VisitSwitchStmt(ctx *SwitchStmtContext) interface{}

	// Visit a parse tree produced by VlangParser#CaseClause.
	VisitCaseClause(ctx *CaseClauseContext) interface{}

	// Visit a parse tree produced by VlangParser#DefaultClause.
	VisitDefaultClause(ctx *DefaultClauseContext) interface{}

	// Visit a parse tree produced by VlangParser#DeclSlices.
	VisitDeclSlices(ctx *DeclSlicesContext) interface{}

	// Visit a parse tree produced by VlangParser#FuncCall.
	VisitFuncCall(ctx *FuncCallContext) interface{}

	// Visit a parse tree produced by VlangParser#FuncDecl.
	VisitFuncDecl(ctx *FuncDeclContext) interface{}

	// Visit a parse tree produced by VlangParser#StructDecl.
	VisitStructDecl(ctx *StructDeclContext) interface{}

	// Visit a parse tree produced by VlangParser#ReturnStmt.
	VisitReturnStmt(ctx *ReturnStmtContext) interface{}

	// Visit a parse tree produced by VlangParser#BreakStmt.
	VisitBreakStmt(ctx *BreakStmtContext) interface{}

	// Visit a parse tree produced by VlangParser#ContinueStmt.
	VisitContinueStmt(ctx *ContinueStmtContext) interface{}

	// Visit a parse tree produced by VlangParser#DirectAssign.
	VisitDirectAssign(ctx *DirectAssignContext) interface{}

	// Visit a parse tree produced by VlangParser#DeclAssignInferred.
	VisitDeclAssignInferred(ctx *DeclAssignInferredContext) interface{}

	// Visit a parse tree produced by VlangParser#DeclAssignExplicit.
	VisitDeclAssignExplicit(ctx *DeclAssignExplicitContext) interface{}

	// Visit a parse tree produced by VlangParser#ListExpressions.
	VisitListExpressions(ctx *ListExpressionsContext) interface{}

	// Visit a parse tree produced by VlangParser#IdPattern.
	VisitIdPattern(ctx *IdPatternContext) interface{}

	// Visit a parse tree produced by VlangParser#unarioExpresion.
	VisitUnarioExpresion(ctx *UnarioExpresionContext) interface{}

	// Visit a parse tree produced by VlangParser#indexAcceso.
	VisitIndexAcceso(ctx *IndexAccesoContext) interface{}

	// Visit a parse tree produced by VlangParser#valorExpresion.
	VisitValorExpresion(ctx *ValorExpresionContext) interface{}

	// Visit a parse tree produced by VlangParser#funcionExpresion.
	VisitFuncionExpresion(ctx *FuncionExpresionContext) interface{}

	// Visit a parse tree produced by VlangParser#parentesisExpresion.
	VisitParentesisExpresion(ctx *ParentesisExpresionContext) interface{}

	// Visit a parse tree produced by VlangParser#id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by VlangParser#increDecreExp.
	VisitIncreDecreExp(ctx *IncreDecreExpContext) interface{}

	// Visit a parse tree produced by VlangParser#BinaryExp.
	VisitBinaryExp(ctx *BinaryExpContext) interface{}

	// Visit a parse tree produced by VlangParser#ParamList.
	VisitParamList(ctx *ParamListContext) interface{}

	// Visit a parse tree produced by VlangParser#funcParam.
	VisitFuncParam(ctx *FuncParamContext) interface{}

	// Visit a parse tree produced by VlangParser#ArgList.
	VisitArgList(ctx *ArgListContext) interface{}

	// Visit a parse tree produced by VlangParser#FuncArg.
	VisitFuncArg(ctx *FuncArgContext) interface{}

	// Visit a parse tree produced by VlangParser#valorEntero.
	VisitValorEntero(ctx *ValorEnteroContext) interface{}

	// Visit a parse tree produced by VlangParser#valorFloat.
	VisitValorFloat(ctx *ValorFloatContext) interface{}

	// Visit a parse tree produced by VlangParser#valorCadena.
	VisitValorCadena(ctx *ValorCadenaContext) interface{}

	// Visit a parse tree produced by VlangParser#valorBooleano.
	VisitValorBooleano(ctx *ValorBooleanoContext) interface{}

	// Visit a parse tree produced by VlangParser#StructInstancia.
	VisitStructInstancia(ctx *StructInstanciaContext) interface{}

	// Visit a parse tree produced by VlangParser#StructAttr.
	VisitStructAttr(ctx *StructAttrContext) interface{}

	// Visit a parse tree produced by VlangParser#IntType.
	VisitIntType(ctx *IntTypeContext) interface{}

	// Visit a parse tree produced by VlangParser#FloatType.
	VisitFloatType(ctx *FloatTypeContext) interface{}

	// Visit a parse tree produced by VlangParser#StringType.
	VisitStringType(ctx *StringTypeContext) interface{}

	// Visit a parse tree produced by VlangParser#BoolType.
	VisitBoolType(ctx *BoolTypeContext) interface{}

	// Visit a parse tree produced by VlangParser#CharType.
	VisitCharType(ctx *CharTypeContext) interface{}

	// Visit a parse tree produced by VlangParser#VoidType.
	VisitVoidType(ctx *VoidTypeContext) interface{}

	// Visit a parse tree produced by VlangParser#incremento.
	VisitIncremento(ctx *IncrementoContext) interface{}

	// Visit a parse tree produced by VlangParser#decremento.
	VisitDecremento(ctx *DecrementoContext) interface{}
}
