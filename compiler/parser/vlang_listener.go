// Code generated from parser/Vlang.g4 by ANTLR 4.13.1. DO NOT EDIT.

package compiler // Vlang
import "github.com/antlr4-go/antlr/v4"

// VlangListener is a complete listener for a parse tree produced by VlangParser.
type VlangListener interface {
	antlr.ParseTreeListener

	// EnterPrograma is called when entering the programa production.
	EnterPrograma(c *ProgramaContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterBlockStmt is called when entering the BlockStmt production.
	EnterBlockStmt(c *BlockStmtContext)

	// EnterIfStmt is called when entering the IfStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterIfChain is called when entering the IfChain production.
	EnterIfChain(c *IfChainContext)

	// EnterElseStmt is called when entering the ElseStmt production.
	EnterElseStmt(c *ElseStmtContext)

	// EnterPrintlnStmt is called when entering the PrintlnStmt production.
	EnterPrintlnStmt(c *PrintlnStmtContext)

	// EnterPrintStmt is called when entering the PrintStmt production.
	EnterPrintStmt(c *PrintStmtContext)

	// EnterWhileStmt is called when entering the WhileStmt production.
	EnterWhileStmt(c *WhileStmtContext)

	// EnterClasicFor is called when entering the ClasicFor production.
	EnterClasicFor(c *ClasicForContext)

	// EnterForStmt is called when entering the ForStmt production.
	EnterForStmt(c *ForStmtContext)

	// EnterNumericRange is called when entering the NumericRange production.
	EnterNumericRange(c *NumericRangeContext)

	// EnterSwitchStmt is called when entering the SwitchStmt production.
	EnterSwitchStmt(c *SwitchStmtContext)

	// EnterCaseClause is called when entering the CaseClause production.
	EnterCaseClause(c *CaseClauseContext)

	// EnterDefaultClause is called when entering the DefaultClause production.
	EnterDefaultClause(c *DefaultClauseContext)

	// EnterDeclSlices is called when entering the DeclSlices production.
	EnterDeclSlices(c *DeclSlicesContext)

	// EnterFuncCall is called when entering the FuncCall production.
	EnterFuncCall(c *FuncCallContext)

	// EnterFuncDecl is called when entering the FuncDecl production.
	EnterFuncDecl(c *FuncDeclContext)

	// EnterStructDecl is called when entering the StructDecl production.
	EnterStructDecl(c *StructDeclContext)

	// EnterReturnStmt is called when entering the ReturnStmt production.
	EnterReturnStmt(c *ReturnStmtContext)

	// EnterBreakStmt is called when entering the BreakStmt production.
	EnterBreakStmt(c *BreakStmtContext)

	// EnterContinueStmt is called when entering the ContinueStmt production.
	EnterContinueStmt(c *ContinueStmtContext)

	// EnterDirectAssign is called when entering the DirectAssign production.
	EnterDirectAssign(c *DirectAssignContext)

	// EnterDeclAssignInferred is called when entering the DeclAssignInferred production.
	EnterDeclAssignInferred(c *DeclAssignInferredContext)

	// EnterDeclAssignExplicit is called when entering the DeclAssignExplicit production.
	EnterDeclAssignExplicit(c *DeclAssignExplicitContext)

	// EnterListExpressions is called when entering the ListExpressions production.
	EnterListExpressions(c *ListExpressionsContext)

	// EnterIdPattern is called when entering the IdPattern production.
	EnterIdPattern(c *IdPatternContext)

	// EnterUnarioExpresion is called when entering the unarioExpresion production.
	EnterUnarioExpresion(c *UnarioExpresionContext)

	// EnterIndexAcceso is called when entering the indexAcceso production.
	EnterIndexAcceso(c *IndexAccesoContext)

	// EnterValorExpresion is called when entering the valorExpresion production.
	EnterValorExpresion(c *ValorExpresionContext)

	// EnterFuncionExpresion is called when entering the funcionExpresion production.
	EnterFuncionExpresion(c *FuncionExpresionContext)

	// EnterParentesisExpresion is called when entering the parentesisExpresion production.
	EnterParentesisExpresion(c *ParentesisExpresionContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// EnterIncreDecreExp is called when entering the increDecreExp production.
	EnterIncreDecreExp(c *IncreDecreExpContext)

	// EnterBinaryExp is called when entering the BinaryExp production.
	EnterBinaryExp(c *BinaryExpContext)

	// EnterParamList is called when entering the ParamList production.
	EnterParamList(c *ParamListContext)

	// EnterFuncParam is called when entering the funcParam production.
	EnterFuncParam(c *FuncParamContext)

	// EnterArgList is called when entering the ArgList production.
	EnterArgList(c *ArgListContext)

	// EnterFuncArg is called when entering the FuncArg production.
	EnterFuncArg(c *FuncArgContext)

	// EnterValorEntero is called when entering the valorEntero production.
	EnterValorEntero(c *ValorEnteroContext)

	// EnterValorFloat is called when entering the valorFloat production.
	EnterValorFloat(c *ValorFloatContext)

	// EnterValorCadena is called when entering the valorCadena production.
	EnterValorCadena(c *ValorCadenaContext)

	// EnterValorBooleano is called when entering the valorBooleano production.
	EnterValorBooleano(c *ValorBooleanoContext)

	// EnterStructInstancia is called when entering the StructInstancia production.
	EnterStructInstancia(c *StructInstanciaContext)

	// EnterStructAttr is called when entering the StructAttr production.
	EnterStructAttr(c *StructAttrContext)

	// EnterIntType is called when entering the IntType production.
	EnterIntType(c *IntTypeContext)

	// EnterFloatType is called when entering the FloatType production.
	EnterFloatType(c *FloatTypeContext)

	// EnterStringType is called when entering the StringType production.
	EnterStringType(c *StringTypeContext)

	// EnterBoolType is called when entering the BoolType production.
	EnterBoolType(c *BoolTypeContext)

	// EnterCharType is called when entering the CharType production.
	EnterCharType(c *CharTypeContext)

	// EnterVoidType is called when entering the VoidType production.
	EnterVoidType(c *VoidTypeContext)

	// EnterIncremento is called when entering the incremento production.
	EnterIncremento(c *IncrementoContext)

	// EnterDecremento is called when entering the decremento production.
	EnterDecremento(c *DecrementoContext)

	// ExitPrograma is called when exiting the programa production.
	ExitPrograma(c *ProgramaContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitBlockStmt is called when exiting the BlockStmt production.
	ExitBlockStmt(c *BlockStmtContext)

	// ExitIfStmt is called when exiting the IfStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitIfChain is called when exiting the IfChain production.
	ExitIfChain(c *IfChainContext)

	// ExitElseStmt is called when exiting the ElseStmt production.
	ExitElseStmt(c *ElseStmtContext)

	// ExitPrintlnStmt is called when exiting the PrintlnStmt production.
	ExitPrintlnStmt(c *PrintlnStmtContext)

	// ExitPrintStmt is called when exiting the PrintStmt production.
	ExitPrintStmt(c *PrintStmtContext)

	// ExitWhileStmt is called when exiting the WhileStmt production.
	ExitWhileStmt(c *WhileStmtContext)

	// ExitClasicFor is called when exiting the ClasicFor production.
	ExitClasicFor(c *ClasicForContext)

	// ExitForStmt is called when exiting the ForStmt production.
	ExitForStmt(c *ForStmtContext)

	// ExitNumericRange is called when exiting the NumericRange production.
	ExitNumericRange(c *NumericRangeContext)

	// ExitSwitchStmt is called when exiting the SwitchStmt production.
	ExitSwitchStmt(c *SwitchStmtContext)

	// ExitCaseClause is called when exiting the CaseClause production.
	ExitCaseClause(c *CaseClauseContext)

	// ExitDefaultClause is called when exiting the DefaultClause production.
	ExitDefaultClause(c *DefaultClauseContext)

	// ExitDeclSlices is called when exiting the DeclSlices production.
	ExitDeclSlices(c *DeclSlicesContext)

	// ExitFuncCall is called when exiting the FuncCall production.
	ExitFuncCall(c *FuncCallContext)

	// ExitFuncDecl is called when exiting the FuncDecl production.
	ExitFuncDecl(c *FuncDeclContext)

	// ExitStructDecl is called when exiting the StructDecl production.
	ExitStructDecl(c *StructDeclContext)

	// ExitReturnStmt is called when exiting the ReturnStmt production.
	ExitReturnStmt(c *ReturnStmtContext)

	// ExitBreakStmt is called when exiting the BreakStmt production.
	ExitBreakStmt(c *BreakStmtContext)

	// ExitContinueStmt is called when exiting the ContinueStmt production.
	ExitContinueStmt(c *ContinueStmtContext)

	// ExitDirectAssign is called when exiting the DirectAssign production.
	ExitDirectAssign(c *DirectAssignContext)

	// ExitDeclAssignInferred is called when exiting the DeclAssignInferred production.
	ExitDeclAssignInferred(c *DeclAssignInferredContext)

	// ExitDeclAssignExplicit is called when exiting the DeclAssignExplicit production.
	ExitDeclAssignExplicit(c *DeclAssignExplicitContext)

	// ExitListExpressions is called when exiting the ListExpressions production.
	ExitListExpressions(c *ListExpressionsContext)

	// ExitIdPattern is called when exiting the IdPattern production.
	ExitIdPattern(c *IdPatternContext)

	// ExitUnarioExpresion is called when exiting the unarioExpresion production.
	ExitUnarioExpresion(c *UnarioExpresionContext)

	// ExitIndexAcceso is called when exiting the indexAcceso production.
	ExitIndexAcceso(c *IndexAccesoContext)

	// ExitValorExpresion is called when exiting the valorExpresion production.
	ExitValorExpresion(c *ValorExpresionContext)

	// ExitFuncionExpresion is called when exiting the funcionExpresion production.
	ExitFuncionExpresion(c *FuncionExpresionContext)

	// ExitParentesisExpresion is called when exiting the parentesisExpresion production.
	ExitParentesisExpresion(c *ParentesisExpresionContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)

	// ExitIncreDecreExp is called when exiting the increDecreExp production.
	ExitIncreDecreExp(c *IncreDecreExpContext)

	// ExitBinaryExp is called when exiting the BinaryExp production.
	ExitBinaryExp(c *BinaryExpContext)

	// ExitParamList is called when exiting the ParamList production.
	ExitParamList(c *ParamListContext)

	// ExitFuncParam is called when exiting the funcParam production.
	ExitFuncParam(c *FuncParamContext)

	// ExitArgList is called when exiting the ArgList production.
	ExitArgList(c *ArgListContext)

	// ExitFuncArg is called when exiting the FuncArg production.
	ExitFuncArg(c *FuncArgContext)

	// ExitValorEntero is called when exiting the valorEntero production.
	ExitValorEntero(c *ValorEnteroContext)

	// ExitValorFloat is called when exiting the valorFloat production.
	ExitValorFloat(c *ValorFloatContext)

	// ExitValorCadena is called when exiting the valorCadena production.
	ExitValorCadena(c *ValorCadenaContext)

	// ExitValorBooleano is called when exiting the valorBooleano production.
	ExitValorBooleano(c *ValorBooleanoContext)

	// ExitStructInstancia is called when exiting the StructInstancia production.
	ExitStructInstancia(c *StructInstanciaContext)

	// ExitStructAttr is called when exiting the StructAttr production.
	ExitStructAttr(c *StructAttrContext)

	// ExitIntType is called when exiting the IntType production.
	ExitIntType(c *IntTypeContext)

	// ExitFloatType is called when exiting the FloatType production.
	ExitFloatType(c *FloatTypeContext)

	// ExitStringType is called when exiting the StringType production.
	ExitStringType(c *StringTypeContext)

	// ExitBoolType is called when exiting the BoolType production.
	ExitBoolType(c *BoolTypeContext)

	// ExitCharType is called when exiting the CharType production.
	ExitCharType(c *CharTypeContext)

	// ExitVoidType is called when exiting the VoidType production.
	ExitVoidType(c *VoidTypeContext)

	// ExitIncremento is called when exiting the incremento production.
	ExitIncremento(c *IncrementoContext)

	// ExitDecremento is called when exiting the decremento production.
	ExitDecremento(c *DecrementoContext)
}
