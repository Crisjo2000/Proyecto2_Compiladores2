// Code generated from parser/Vlang.g4 by ANTLR 4.13.1. DO NOT EDIT.

package compiler // Vlang
import "github.com/antlr4-go/antlr/v4"

// BaseVlangListener is a complete listener for a parse tree produced by VlangParser.
type BaseVlangListener struct{}

var _ VlangListener = &BaseVlangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVlangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVlangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVlangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVlangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPrograma is called when production programa is entered.
func (s *BaseVlangListener) EnterPrograma(ctx *ProgramaContext) {}

// ExitPrograma is called when production programa is exited.
func (s *BaseVlangListener) ExitPrograma(ctx *ProgramaContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseVlangListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseVlangListener) ExitStmt(ctx *StmtContext) {}

// EnterBlockStmt is called when production BlockStmt is entered.
func (s *BaseVlangListener) EnterBlockStmt(ctx *BlockStmtContext) {}

// ExitBlockStmt is called when production BlockStmt is exited.
func (s *BaseVlangListener) ExitBlockStmt(ctx *BlockStmtContext) {}

// EnterIfStmt is called when production IfStmt is entered.
func (s *BaseVlangListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production IfStmt is exited.
func (s *BaseVlangListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterIfChain is called when production IfChain is entered.
func (s *BaseVlangListener) EnterIfChain(ctx *IfChainContext) {}

// ExitIfChain is called when production IfChain is exited.
func (s *BaseVlangListener) ExitIfChain(ctx *IfChainContext) {}

// EnterElseStmt is called when production ElseStmt is entered.
func (s *BaseVlangListener) EnterElseStmt(ctx *ElseStmtContext) {}

// ExitElseStmt is called when production ElseStmt is exited.
func (s *BaseVlangListener) ExitElseStmt(ctx *ElseStmtContext) {}

// EnterPrintlnStmt is called when production PrintlnStmt is entered.
func (s *BaseVlangListener) EnterPrintlnStmt(ctx *PrintlnStmtContext) {}

// ExitPrintlnStmt is called when production PrintlnStmt is exited.
func (s *BaseVlangListener) ExitPrintlnStmt(ctx *PrintlnStmtContext) {}

// EnterPrintStmt is called when production PrintStmt is entered.
func (s *BaseVlangListener) EnterPrintStmt(ctx *PrintStmtContext) {}

// ExitPrintStmt is called when production PrintStmt is exited.
func (s *BaseVlangListener) ExitPrintStmt(ctx *PrintStmtContext) {}

// EnterWhileStmt is called when production WhileStmt is entered.
func (s *BaseVlangListener) EnterWhileStmt(ctx *WhileStmtContext) {}

// ExitWhileStmt is called when production WhileStmt is exited.
func (s *BaseVlangListener) ExitWhileStmt(ctx *WhileStmtContext) {}

// EnterClasicFor is called when production ClasicFor is entered.
func (s *BaseVlangListener) EnterClasicFor(ctx *ClasicForContext) {}

// ExitClasicFor is called when production ClasicFor is exited.
func (s *BaseVlangListener) ExitClasicFor(ctx *ClasicForContext) {}

// EnterForStmt is called when production ForStmt is entered.
func (s *BaseVlangListener) EnterForStmt(ctx *ForStmtContext) {}

// ExitForStmt is called when production ForStmt is exited.
func (s *BaseVlangListener) ExitForStmt(ctx *ForStmtContext) {}

// EnterNumericRange is called when production NumericRange is entered.
func (s *BaseVlangListener) EnterNumericRange(ctx *NumericRangeContext) {}

// ExitNumericRange is called when production NumericRange is exited.
func (s *BaseVlangListener) ExitNumericRange(ctx *NumericRangeContext) {}

// EnterSwitchStmt is called when production SwitchStmt is entered.
func (s *BaseVlangListener) EnterSwitchStmt(ctx *SwitchStmtContext) {}

// ExitSwitchStmt is called when production SwitchStmt is exited.
func (s *BaseVlangListener) ExitSwitchStmt(ctx *SwitchStmtContext) {}

// EnterCaseClause is called when production CaseClause is entered.
func (s *BaseVlangListener) EnterCaseClause(ctx *CaseClauseContext) {}

// ExitCaseClause is called when production CaseClause is exited.
func (s *BaseVlangListener) ExitCaseClause(ctx *CaseClauseContext) {}

// EnterDefaultClause is called when production DefaultClause is entered.
func (s *BaseVlangListener) EnterDefaultClause(ctx *DefaultClauseContext) {}

// ExitDefaultClause is called when production DefaultClause is exited.
func (s *BaseVlangListener) ExitDefaultClause(ctx *DefaultClauseContext) {}

// EnterDeclSlices is called when production DeclSlices is entered.
func (s *BaseVlangListener) EnterDeclSlices(ctx *DeclSlicesContext) {}

// ExitDeclSlices is called when production DeclSlices is exited.
func (s *BaseVlangListener) ExitDeclSlices(ctx *DeclSlicesContext) {}

// EnterFuncCall is called when production FuncCall is entered.
func (s *BaseVlangListener) EnterFuncCall(ctx *FuncCallContext) {}

// ExitFuncCall is called when production FuncCall is exited.
func (s *BaseVlangListener) ExitFuncCall(ctx *FuncCallContext) {}

// EnterFuncDecl is called when production FuncDecl is entered.
func (s *BaseVlangListener) EnterFuncDecl(ctx *FuncDeclContext) {}

// ExitFuncDecl is called when production FuncDecl is exited.
func (s *BaseVlangListener) ExitFuncDecl(ctx *FuncDeclContext) {}

// EnterStructDecl is called when production StructDecl is entered.
func (s *BaseVlangListener) EnterStructDecl(ctx *StructDeclContext) {}

// ExitStructDecl is called when production StructDecl is exited.
func (s *BaseVlangListener) ExitStructDecl(ctx *StructDeclContext) {}

// EnterReturnStmt is called when production ReturnStmt is entered.
func (s *BaseVlangListener) EnterReturnStmt(ctx *ReturnStmtContext) {}

// ExitReturnStmt is called when production ReturnStmt is exited.
func (s *BaseVlangListener) ExitReturnStmt(ctx *ReturnStmtContext) {}

// EnterBreakStmt is called when production BreakStmt is entered.
func (s *BaseVlangListener) EnterBreakStmt(ctx *BreakStmtContext) {}

// ExitBreakStmt is called when production BreakStmt is exited.
func (s *BaseVlangListener) ExitBreakStmt(ctx *BreakStmtContext) {}

// EnterContinueStmt is called when production ContinueStmt is entered.
func (s *BaseVlangListener) EnterContinueStmt(ctx *ContinueStmtContext) {}

// ExitContinueStmt is called when production ContinueStmt is exited.
func (s *BaseVlangListener) ExitContinueStmt(ctx *ContinueStmtContext) {}

// EnterDirectAssign is called when production DirectAssign is entered.
func (s *BaseVlangListener) EnterDirectAssign(ctx *DirectAssignContext) {}

// ExitDirectAssign is called when production DirectAssign is exited.
func (s *BaseVlangListener) ExitDirectAssign(ctx *DirectAssignContext) {}

// EnterDeclAssignInferred is called when production DeclAssignInferred is entered.
func (s *BaseVlangListener) EnterDeclAssignInferred(ctx *DeclAssignInferredContext) {}

// ExitDeclAssignInferred is called when production DeclAssignInferred is exited.
func (s *BaseVlangListener) ExitDeclAssignInferred(ctx *DeclAssignInferredContext) {}

// EnterDeclAssignExplicit is called when production DeclAssignExplicit is entered.
func (s *BaseVlangListener) EnterDeclAssignExplicit(ctx *DeclAssignExplicitContext) {}

// ExitDeclAssignExplicit is called when production DeclAssignExplicit is exited.
func (s *BaseVlangListener) ExitDeclAssignExplicit(ctx *DeclAssignExplicitContext) {}

// EnterListExpressions is called when production ListExpressions is entered.
func (s *BaseVlangListener) EnterListExpressions(ctx *ListExpressionsContext) {}

// ExitListExpressions is called when production ListExpressions is exited.
func (s *BaseVlangListener) ExitListExpressions(ctx *ListExpressionsContext) {}

// EnterIdPattern is called when production IdPattern is entered.
func (s *BaseVlangListener) EnterIdPattern(ctx *IdPatternContext) {}

// ExitIdPattern is called when production IdPattern is exited.
func (s *BaseVlangListener) ExitIdPattern(ctx *IdPatternContext) {}

// EnterUnarioExpresion is called when production unarioExpresion is entered.
func (s *BaseVlangListener) EnterUnarioExpresion(ctx *UnarioExpresionContext) {}

// ExitUnarioExpresion is called when production unarioExpresion is exited.
func (s *BaseVlangListener) ExitUnarioExpresion(ctx *UnarioExpresionContext) {}

// EnterIndexAcceso is called when production indexAcceso is entered.
func (s *BaseVlangListener) EnterIndexAcceso(ctx *IndexAccesoContext) {}

// ExitIndexAcceso is called when production indexAcceso is exited.
func (s *BaseVlangListener) ExitIndexAcceso(ctx *IndexAccesoContext) {}

// EnterValorExpresion is called when production valorExpresion is entered.
func (s *BaseVlangListener) EnterValorExpresion(ctx *ValorExpresionContext) {}

// ExitValorExpresion is called when production valorExpresion is exited.
func (s *BaseVlangListener) ExitValorExpresion(ctx *ValorExpresionContext) {}

// EnterFuncionExpresion is called when production funcionExpresion is entered.
func (s *BaseVlangListener) EnterFuncionExpresion(ctx *FuncionExpresionContext) {}

// ExitFuncionExpresion is called when production funcionExpresion is exited.
func (s *BaseVlangListener) ExitFuncionExpresion(ctx *FuncionExpresionContext) {}

// EnterParentesisExpresion is called when production parentesisExpresion is entered.
func (s *BaseVlangListener) EnterParentesisExpresion(ctx *ParentesisExpresionContext) {}

// ExitParentesisExpresion is called when production parentesisExpresion is exited.
func (s *BaseVlangListener) ExitParentesisExpresion(ctx *ParentesisExpresionContext) {}

// EnterId is called when production id is entered.
func (s *BaseVlangListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BaseVlangListener) ExitId(ctx *IdContext) {}

// EnterIncreDecreExp is called when production increDecreExp is entered.
func (s *BaseVlangListener) EnterIncreDecreExp(ctx *IncreDecreExpContext) {}

// ExitIncreDecreExp is called when production increDecreExp is exited.
func (s *BaseVlangListener) ExitIncreDecreExp(ctx *IncreDecreExpContext) {}

// EnterBinaryExp is called when production BinaryExp is entered.
func (s *BaseVlangListener) EnterBinaryExp(ctx *BinaryExpContext) {}

// ExitBinaryExp is called when production BinaryExp is exited.
func (s *BaseVlangListener) ExitBinaryExp(ctx *BinaryExpContext) {}

// EnterParamList is called when production ParamList is entered.
func (s *BaseVlangListener) EnterParamList(ctx *ParamListContext) {}

// ExitParamList is called when production ParamList is exited.
func (s *BaseVlangListener) ExitParamList(ctx *ParamListContext) {}

// EnterFuncParam is called when production funcParam is entered.
func (s *BaseVlangListener) EnterFuncParam(ctx *FuncParamContext) {}

// ExitFuncParam is called when production funcParam is exited.
func (s *BaseVlangListener) ExitFuncParam(ctx *FuncParamContext) {}

// EnterArgList is called when production ArgList is entered.
func (s *BaseVlangListener) EnterArgList(ctx *ArgListContext) {}

// ExitArgList is called when production ArgList is exited.
func (s *BaseVlangListener) ExitArgList(ctx *ArgListContext) {}

// EnterFuncArg is called when production FuncArg is entered.
func (s *BaseVlangListener) EnterFuncArg(ctx *FuncArgContext) {}

// ExitFuncArg is called when production FuncArg is exited.
func (s *BaseVlangListener) ExitFuncArg(ctx *FuncArgContext) {}

// EnterValorEntero is called when production valorEntero is entered.
func (s *BaseVlangListener) EnterValorEntero(ctx *ValorEnteroContext) {}

// ExitValorEntero is called when production valorEntero is exited.
func (s *BaseVlangListener) ExitValorEntero(ctx *ValorEnteroContext) {}

// EnterValorFloat is called when production valorFloat is entered.
func (s *BaseVlangListener) EnterValorFloat(ctx *ValorFloatContext) {}

// ExitValorFloat is called when production valorFloat is exited.
func (s *BaseVlangListener) ExitValorFloat(ctx *ValorFloatContext) {}

// EnterValorCadena is called when production valorCadena is entered.
func (s *BaseVlangListener) EnterValorCadena(ctx *ValorCadenaContext) {}

// ExitValorCadena is called when production valorCadena is exited.
func (s *BaseVlangListener) ExitValorCadena(ctx *ValorCadenaContext) {}

// EnterValorBooleano is called when production valorBooleano is entered.
func (s *BaseVlangListener) EnterValorBooleano(ctx *ValorBooleanoContext) {}

// ExitValorBooleano is called when production valorBooleano is exited.
func (s *BaseVlangListener) ExitValorBooleano(ctx *ValorBooleanoContext) {}

// EnterStructInstancia is called when production StructInstancia is entered.
func (s *BaseVlangListener) EnterStructInstancia(ctx *StructInstanciaContext) {}

// ExitStructInstancia is called when production StructInstancia is exited.
func (s *BaseVlangListener) ExitStructInstancia(ctx *StructInstanciaContext) {}

// EnterStructAttr is called when production StructAttr is entered.
func (s *BaseVlangListener) EnterStructAttr(ctx *StructAttrContext) {}

// ExitStructAttr is called when production StructAttr is exited.
func (s *BaseVlangListener) ExitStructAttr(ctx *StructAttrContext) {}

// EnterIntType is called when production IntType is entered.
func (s *BaseVlangListener) EnterIntType(ctx *IntTypeContext) {}

// ExitIntType is called when production IntType is exited.
func (s *BaseVlangListener) ExitIntType(ctx *IntTypeContext) {}

// EnterFloatType is called when production FloatType is entered.
func (s *BaseVlangListener) EnterFloatType(ctx *FloatTypeContext) {}

// ExitFloatType is called when production FloatType is exited.
func (s *BaseVlangListener) ExitFloatType(ctx *FloatTypeContext) {}

// EnterStringType is called when production StringType is entered.
func (s *BaseVlangListener) EnterStringType(ctx *StringTypeContext) {}

// ExitStringType is called when production StringType is exited.
func (s *BaseVlangListener) ExitStringType(ctx *StringTypeContext) {}

// EnterBoolType is called when production BoolType is entered.
func (s *BaseVlangListener) EnterBoolType(ctx *BoolTypeContext) {}

// ExitBoolType is called when production BoolType is exited.
func (s *BaseVlangListener) ExitBoolType(ctx *BoolTypeContext) {}

// EnterCharType is called when production CharType is entered.
func (s *BaseVlangListener) EnterCharType(ctx *CharTypeContext) {}

// ExitCharType is called when production CharType is exited.
func (s *BaseVlangListener) ExitCharType(ctx *CharTypeContext) {}

// EnterVoidType is called when production VoidType is entered.
func (s *BaseVlangListener) EnterVoidType(ctx *VoidTypeContext) {}

// ExitVoidType is called when production VoidType is exited.
func (s *BaseVlangListener) ExitVoidType(ctx *VoidTypeContext) {}

// EnterIncremento is called when production incremento is entered.
func (s *BaseVlangListener) EnterIncremento(ctx *IncrementoContext) {}

// ExitIncremento is called when production incremento is exited.
func (s *BaseVlangListener) ExitIncremento(ctx *IncrementoContext) {}

// EnterDecremento is called when production decremento is entered.
func (s *BaseVlangListener) EnterDecremento(ctx *DecrementoContext) {}

// ExitDecremento is called when production decremento is exited.
func (s *BaseVlangListener) ExitDecremento(ctx *DecrementoContext) {}
