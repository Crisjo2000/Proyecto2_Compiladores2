// Generated from /home/luisa/Documentos/2025/Vaqueras/Proyecto 2/OLC2_Proyecto2_G10/compiler/parser/Vlang.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link VlangParser}.
 */
public interface VlangListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link VlangParser#programa}.
	 * @param ctx the parse tree
	 */
	void enterPrograma(VlangParser.ProgramaContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#programa}.
	 * @param ctx the parse tree
	 */
	void exitPrograma(VlangParser.ProgramaContext ctx);
	/**
	 * Enter a parse tree produced by {@link VlangParser#stmt}.
	 * @param ctx the parse tree
	 */
	void enterStmt(VlangParser.StmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#stmt}.
	 * @param ctx the parse tree
	 */
	void exitStmt(VlangParser.StmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IfStmt}
	 * labeled alternative in {@link VlangParser#if_stmt}.
	 * @param ctx the parse tree
	 */
	void enterIfStmt(VlangParser.IfStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IfStmt}
	 * labeled alternative in {@link VlangParser#if_stmt}.
	 * @param ctx the parse tree
	 */
	void exitIfStmt(VlangParser.IfStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IfChain}
	 * labeled alternative in {@link VlangParser#if_chain}.
	 * @param ctx the parse tree
	 */
	void enterIfChain(VlangParser.IfChainContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IfChain}
	 * labeled alternative in {@link VlangParser#if_chain}.
	 * @param ctx the parse tree
	 */
	void exitIfChain(VlangParser.IfChainContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ElseStmt}
	 * labeled alternative in {@link VlangParser#else_stmt}.
	 * @param ctx the parse tree
	 */
	void enterElseStmt(VlangParser.ElseStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ElseStmt}
	 * labeled alternative in {@link VlangParser#else_stmt}.
	 * @param ctx the parse tree
	 */
	void exitElseStmt(VlangParser.ElseStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code PrintlnStmt}
	 * labeled alternative in {@link VlangParser#println}.
	 * @param ctx the parse tree
	 */
	void enterPrintlnStmt(VlangParser.PrintlnStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code PrintlnStmt}
	 * labeled alternative in {@link VlangParser#println}.
	 * @param ctx the parse tree
	 */
	void exitPrintlnStmt(VlangParser.PrintlnStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code PrintStmt}
	 * labeled alternative in {@link VlangParser#print}.
	 * @param ctx the parse tree
	 */
	void enterPrintStmt(VlangParser.PrintStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code PrintStmt}
	 * labeled alternative in {@link VlangParser#print}.
	 * @param ctx the parse tree
	 */
	void exitPrintStmt(VlangParser.PrintStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code WhileStmt}
	 * labeled alternative in {@link VlangParser#forwhile_stmt}.
	 * @param ctx the parse tree
	 */
	void enterWhileStmt(VlangParser.WhileStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code WhileStmt}
	 * labeled alternative in {@link VlangParser#forwhile_stmt}.
	 * @param ctx the parse tree
	 */
	void exitWhileStmt(VlangParser.WhileStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code clasicFor}
	 * labeled alternative in {@link VlangParser#forClasic_stmt}.
	 * @param ctx the parse tree
	 */
	void enterClasicFor(VlangParser.ClasicForContext ctx);
	/**
	 * Exit a parse tree produced by the {@code clasicFor}
	 * labeled alternative in {@link VlangParser#forClasic_stmt}.
	 * @param ctx the parse tree
	 */
	void exitClasicFor(VlangParser.ClasicForContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ForStmt}
	 * labeled alternative in {@link VlangParser#for_stmt}.
	 * @param ctx the parse tree
	 */
	void enterForStmt(VlangParser.ForStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ForStmt}
	 * labeled alternative in {@link VlangParser#for_stmt}.
	 * @param ctx the parse tree
	 */
	void exitForStmt(VlangParser.ForStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code NumericRange}
	 * labeled alternative in {@link VlangParser#range}.
	 * @param ctx the parse tree
	 */
	void enterNumericRange(VlangParser.NumericRangeContext ctx);
	/**
	 * Exit a parse tree produced by the {@code NumericRange}
	 * labeled alternative in {@link VlangParser#range}.
	 * @param ctx the parse tree
	 */
	void exitNumericRange(VlangParser.NumericRangeContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SwitchStmt}
	 * labeled alternative in {@link VlangParser#switch_stmt}.
	 * @param ctx the parse tree
	 */
	void enterSwitchStmt(VlangParser.SwitchStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SwitchStmt}
	 * labeled alternative in {@link VlangParser#switch_stmt}.
	 * @param ctx the parse tree
	 */
	void exitSwitchStmt(VlangParser.SwitchStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code CaseClause}
	 * labeled alternative in {@link VlangParser#case_clause}.
	 * @param ctx the parse tree
	 */
	void enterCaseClause(VlangParser.CaseClauseContext ctx);
	/**
	 * Exit a parse tree produced by the {@code CaseClause}
	 * labeled alternative in {@link VlangParser#case_clause}.
	 * @param ctx the parse tree
	 */
	void exitCaseClause(VlangParser.CaseClauseContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DefaultClause}
	 * labeled alternative in {@link VlangParser#default_clause}.
	 * @param ctx the parse tree
	 */
	void enterDefaultClause(VlangParser.DefaultClauseContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DefaultClause}
	 * labeled alternative in {@link VlangParser#default_clause}.
	 * @param ctx the parse tree
	 */
	void exitDefaultClause(VlangParser.DefaultClauseContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DeclSlices}
	 * labeled alternative in {@link VlangParser#slice_stmt}.
	 * @param ctx the parse tree
	 */
	void enterDeclSlices(VlangParser.DeclSlicesContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DeclSlices}
	 * labeled alternative in {@link VlangParser#slice_stmt}.
	 * @param ctx the parse tree
	 */
	void exitDeclSlices(VlangParser.DeclSlicesContext ctx);
	/**
	 * Enter a parse tree produced by the {@code FuncCall}
	 * labeled alternative in {@link VlangParser#func_call}.
	 * @param ctx the parse tree
	 */
	void enterFuncCall(VlangParser.FuncCallContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FuncCall}
	 * labeled alternative in {@link VlangParser#func_call}.
	 * @param ctx the parse tree
	 */
	void exitFuncCall(VlangParser.FuncCallContext ctx);
	/**
	 * Enter a parse tree produced by the {@code FuncDecl}
	 * labeled alternative in {@link VlangParser#func_dcl}.
	 * @param ctx the parse tree
	 */
	void enterFuncDecl(VlangParser.FuncDeclContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FuncDecl}
	 * labeled alternative in {@link VlangParser#func_dcl}.
	 * @param ctx the parse tree
	 */
	void exitFuncDecl(VlangParser.FuncDeclContext ctx);
	/**
	 * Enter a parse tree produced by the {@code StructDecl}
	 * labeled alternative in {@link VlangParser#struct_dcl}.
	 * @param ctx the parse tree
	 */
	void enterStructDecl(VlangParser.StructDeclContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StructDecl}
	 * labeled alternative in {@link VlangParser#struct_dcl}.
	 * @param ctx the parse tree
	 */
	void exitStructDecl(VlangParser.StructDeclContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ReturnStmt}
	 * labeled alternative in {@link VlangParser#transfer_stmt}.
	 * @param ctx the parse tree
	 */
	void enterReturnStmt(VlangParser.ReturnStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ReturnStmt}
	 * labeled alternative in {@link VlangParser#transfer_stmt}.
	 * @param ctx the parse tree
	 */
	void exitReturnStmt(VlangParser.ReturnStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code BreakStmt}
	 * labeled alternative in {@link VlangParser#transfer_stmt}.
	 * @param ctx the parse tree
	 */
	void enterBreakStmt(VlangParser.BreakStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code BreakStmt}
	 * labeled alternative in {@link VlangParser#transfer_stmt}.
	 * @param ctx the parse tree
	 */
	void exitBreakStmt(VlangParser.BreakStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ContinueStmt}
	 * labeled alternative in {@link VlangParser#transfer_stmt}.
	 * @param ctx the parse tree
	 */
	void enterContinueStmt(VlangParser.ContinueStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ContinueStmt}
	 * labeled alternative in {@link VlangParser#transfer_stmt}.
	 * @param ctx the parse tree
	 */
	void exitContinueStmt(VlangParser.ContinueStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DirectAssign}
	 * labeled alternative in {@link VlangParser#assign_stmt}.
	 * @param ctx the parse tree
	 */
	void enterDirectAssign(VlangParser.DirectAssignContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DirectAssign}
	 * labeled alternative in {@link VlangParser#assign_stmt}.
	 * @param ctx the parse tree
	 */
	void exitDirectAssign(VlangParser.DirectAssignContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DeclAssign}
	 * labeled alternative in {@link VlangParser#decl_stmt}.
	 * @param ctx the parse tree
	 */
	void enterDeclAssign(VlangParser.DeclAssignContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DeclAssign}
	 * labeled alternative in {@link VlangParser#decl_stmt}.
	 * @param ctx the parse tree
	 */
	void exitDeclAssign(VlangParser.DeclAssignContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DeclAssignType}
	 * labeled alternative in {@link VlangParser#decl_stmt}.
	 * @param ctx the parse tree
	 */
	void enterDeclAssignType(VlangParser.DeclAssignTypeContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DeclAssignType}
	 * labeled alternative in {@link VlangParser#decl_stmt}.
	 * @param ctx the parse tree
	 */
	void exitDeclAssignType(VlangParser.DeclAssignTypeContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ListExpressions}
	 * labeled alternative in {@link VlangParser#list_expresiones}.
	 * @param ctx the parse tree
	 */
	void enterListExpressions(VlangParser.ListExpressionsContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ListExpressions}
	 * labeled alternative in {@link VlangParser#list_expresiones}.
	 * @param ctx the parse tree
	 */
	void exitListExpressions(VlangParser.ListExpressionsContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IdPattern}
	 * labeled alternative in {@link VlangParser#id_pattern}.
	 * @param ctx the parse tree
	 */
	void enterIdPattern(VlangParser.IdPatternContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IdPattern}
	 * labeled alternative in {@link VlangParser#id_pattern}.
	 * @param ctx the parse tree
	 */
	void exitIdPattern(VlangParser.IdPatternContext ctx);
	/**
	 * Enter a parse tree produced by the {@code incredecr}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterIncredecr(VlangParser.IncredecrContext ctx);
	/**
	 * Exit a parse tree produced by the {@code incredecr}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitIncredecr(VlangParser.IncredecrContext ctx);
	/**
	 * Enter a parse tree produced by the {@code fn_lenx}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterFn_lenx(VlangParser.Fn_lenxContext ctx);
	/**
	 * Exit a parse tree produced by the {@code fn_lenx}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitFn_lenx(VlangParser.Fn_lenxContext ctx);
	/**
	 * Enter a parse tree produced by the {@code valorexpresion}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterValorexpresion(VlangParser.ValorexpresionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code valorexpresion}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitValorexpresion(VlangParser.ValorexpresionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code asignacionfor}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterAsignacionfor(VlangParser.AsignacionforContext ctx);
	/**
	 * Exit a parse tree produced by the {@code asignacionfor}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitAsignacionfor(VlangParser.AsignacionforContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expdotexp}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterExpdotexp(VlangParser.ExpdotexpContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expdotexp}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitExpdotexp(VlangParser.ExpdotexpContext ctx);
	/**
	 * Enter a parse tree produced by the {@code fn_typeOF}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterFn_typeOF(VlangParser.Fn_typeOFContext ctx);
	/**
	 * Exit a parse tree produced by the {@code fn_typeOF}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitFn_typeOF(VlangParser.Fn_typeOFContext ctx);
	/**
	 * Enter a parse tree produced by the {@code BinaryExp}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterBinaryExp(VlangParser.BinaryExpContext ctx);
	/**
	 * Exit a parse tree produced by the {@code BinaryExp}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitBinaryExp(VlangParser.BinaryExpContext ctx);
	/**
	 * Enter a parse tree produced by the {@code corchetesexpre}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterCorchetesexpre(VlangParser.CorchetesexpreContext ctx);
	/**
	 * Exit a parse tree produced by the {@code corchetesexpre}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitCorchetesexpre(VlangParser.CorchetesexpreContext ctx);
	/**
	 * Enter a parse tree produced by the {@code funcionexpre}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterFuncionexpre(VlangParser.FuncionexpreContext ctx);
	/**
	 * Exit a parse tree produced by the {@code funcionexpre}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitFuncionexpre(VlangParser.FuncionexpreContext ctx);
	/**
	 * Enter a parse tree produced by the {@code unario}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterUnario(VlangParser.UnarioContext ctx);
	/**
	 * Exit a parse tree produced by the {@code unario}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitUnario(VlangParser.UnarioContext ctx);
	/**
	 * Enter a parse tree produced by the {@code parentesisexpre}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterParentesisexpre(VlangParser.ParentesisexpreContext ctx);
	/**
	 * Exit a parse tree produced by the {@code parentesisexpre}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitParentesisexpre(VlangParser.ParentesisexpreContext ctx);
	/**
	 * Enter a parse tree produced by the {@code fn_parsefloat}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterFn_parsefloat(VlangParser.Fn_parsefloatContext ctx);
	/**
	 * Exit a parse tree produced by the {@code fn_parsefloat}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitFn_parsefloat(VlangParser.Fn_parsefloatContext ctx);
	/**
	 * Enter a parse tree produced by the {@code indexAcceso}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterIndexAcceso(VlangParser.IndexAccesoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code indexAcceso}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitIndexAcceso(VlangParser.IndexAccesoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code id}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterId(VlangParser.IdContext ctx);
	/**
	 * Exit a parse tree produced by the {@code id}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitId(VlangParser.IdContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expdotexp1}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterExpdotexp1(VlangParser.Expdotexp1Context ctx);
	/**
	 * Exit a parse tree produced by the {@code expdotexp1}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitExpdotexp1(VlangParser.Expdotexp1Context ctx);
	/**
	 * Enter a parse tree produced by the {@code fn_indexOf}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterFn_indexOf(VlangParser.Fn_indexOfContext ctx);
	/**
	 * Exit a parse tree produced by the {@code fn_indexOf}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitFn_indexOf(VlangParser.Fn_indexOfContext ctx);
	/**
	 * Enter a parse tree produced by the {@code fn_join}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterFn_join(VlangParser.Fn_joinContext ctx);
	/**
	 * Exit a parse tree produced by the {@code fn_join}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitFn_join(VlangParser.Fn_joinContext ctx);
	/**
	 * Enter a parse tree produced by the {@code fn_atoi}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterFn_atoi(VlangParser.Fn_atoiContext ctx);
	/**
	 * Exit a parse tree produced by the {@code fn_atoi}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitFn_atoi(VlangParser.Fn_atoiContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ParamList}
	 * labeled alternative in {@link VlangParser#parametros}.
	 * @param ctx the parse tree
	 */
	void enterParamList(VlangParser.ParamListContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ParamList}
	 * labeled alternative in {@link VlangParser#parametros}.
	 * @param ctx the parse tree
	 */
	void exitParamList(VlangParser.ParamListContext ctx);
	/**
	 * Enter a parse tree produced by the {@code funcParam}
	 * labeled alternative in {@link VlangParser#func_param}.
	 * @param ctx the parse tree
	 */
	void enterFuncParam(VlangParser.FuncParamContext ctx);
	/**
	 * Exit a parse tree produced by the {@code funcParam}
	 * labeled alternative in {@link VlangParser#func_param}.
	 * @param ctx the parse tree
	 */
	void exitFuncParam(VlangParser.FuncParamContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ArgList}
	 * labeled alternative in {@link VlangParser#arg_list}.
	 * @param ctx the parse tree
	 */
	void enterArgList(VlangParser.ArgListContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ArgList}
	 * labeled alternative in {@link VlangParser#arg_list}.
	 * @param ctx the parse tree
	 */
	void exitArgList(VlangParser.ArgListContext ctx);
	/**
	 * Enter a parse tree produced by the {@code FuncArg}
	 * labeled alternative in {@link VlangParser#func_arg}.
	 * @param ctx the parse tree
	 */
	void enterFuncArg(VlangParser.FuncArgContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FuncArg}
	 * labeled alternative in {@link VlangParser#func_arg}.
	 * @param ctx the parse tree
	 */
	void exitFuncArg(VlangParser.FuncArgContext ctx);
	/**
	 * Enter a parse tree produced by {@link VlangParser#valores}.
	 * @param ctx the parse tree
	 */
	void enterValores(VlangParser.ValoresContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#valores}.
	 * @param ctx the parse tree
	 */
	void exitValores(VlangParser.ValoresContext ctx);
	/**
	 * Enter a parse tree produced by the {@code valorEntero}
	 * labeled alternative in {@link VlangParser#valor}.
	 * @param ctx the parse tree
	 */
	void enterValorEntero(VlangParser.ValorEnteroContext ctx);
	/**
	 * Exit a parse tree produced by the {@code valorEntero}
	 * labeled alternative in {@link VlangParser#valor}.
	 * @param ctx the parse tree
	 */
	void exitValorEntero(VlangParser.ValorEnteroContext ctx);
	/**
	 * Enter a parse tree produced by the {@code valorDecimal}
	 * labeled alternative in {@link VlangParser#valor}.
	 * @param ctx the parse tree
	 */
	void enterValorDecimal(VlangParser.ValorDecimalContext ctx);
	/**
	 * Exit a parse tree produced by the {@code valorDecimal}
	 * labeled alternative in {@link VlangParser#valor}.
	 * @param ctx the parse tree
	 */
	void exitValorDecimal(VlangParser.ValorDecimalContext ctx);
	/**
	 * Enter a parse tree produced by the {@code valorCadena}
	 * labeled alternative in {@link VlangParser#valor}.
	 * @param ctx the parse tree
	 */
	void enterValorCadena(VlangParser.ValorCadenaContext ctx);
	/**
	 * Exit a parse tree produced by the {@code valorCadena}
	 * labeled alternative in {@link VlangParser#valor}.
	 * @param ctx the parse tree
	 */
	void exitValorCadena(VlangParser.ValorCadenaContext ctx);
	/**
	 * Enter a parse tree produced by the {@code valorBooleano}
	 * labeled alternative in {@link VlangParser#valor}.
	 * @param ctx the parse tree
	 */
	void enterValorBooleano(VlangParser.ValorBooleanoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code valorBooleano}
	 * labeled alternative in {@link VlangParser#valor}.
	 * @param ctx the parse tree
	 */
	void exitValorBooleano(VlangParser.ValorBooleanoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code valorCaracter}
	 * labeled alternative in {@link VlangParser#valor}.
	 * @param ctx the parse tree
	 */
	void enterValorCaracter(VlangParser.ValorCaracterContext ctx);
	/**
	 * Exit a parse tree produced by the {@code valorCaracter}
	 * labeled alternative in {@link VlangParser#valor}.
	 * @param ctx the parse tree
	 */
	void exitValorCaracter(VlangParser.ValorCaracterContext ctx);
	/**
	 * Enter a parse tree produced by the {@code valorFloat}
	 * labeled alternative in {@link VlangParser#valor}.
	 * @param ctx the parse tree
	 */
	void enterValorFloat(VlangParser.ValorFloatContext ctx);
	/**
	 * Exit a parse tree produced by the {@code valorFloat}
	 * labeled alternative in {@link VlangParser#valor}.
	 * @param ctx the parse tree
	 */
	void exitValorFloat(VlangParser.ValorFloatContext ctx);
	/**
	 * Enter a parse tree produced by the {@code StructInstancia}
	 * labeled alternative in {@link VlangParser#strct_instancia}.
	 * @param ctx the parse tree
	 */
	void enterStructInstancia(VlangParser.StructInstanciaContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StructInstancia}
	 * labeled alternative in {@link VlangParser#strct_instancia}.
	 * @param ctx the parse tree
	 */
	void exitStructInstancia(VlangParser.StructInstanciaContext ctx);
	/**
	 * Enter a parse tree produced by the {@code StructAttr}
	 * labeled alternative in {@link VlangParser#struct_prop}.
	 * @param ctx the parse tree
	 */
	void enterStructAttr(VlangParser.StructAttrContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StructAttr}
	 * labeled alternative in {@link VlangParser#struct_prop}.
	 * @param ctx the parse tree
	 */
	void exitStructAttr(VlangParser.StructAttrContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IntType}
	 * labeled alternative in {@link VlangParser#var_type}.
	 * @param ctx the parse tree
	 */
	void enterIntType(VlangParser.IntTypeContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IntType}
	 * labeled alternative in {@link VlangParser#var_type}.
	 * @param ctx the parse tree
	 */
	void exitIntType(VlangParser.IntTypeContext ctx);
	/**
	 * Enter a parse tree produced by the {@code FloatType}
	 * labeled alternative in {@link VlangParser#var_type}.
	 * @param ctx the parse tree
	 */
	void enterFloatType(VlangParser.FloatTypeContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FloatType}
	 * labeled alternative in {@link VlangParser#var_type}.
	 * @param ctx the parse tree
	 */
	void exitFloatType(VlangParser.FloatTypeContext ctx);
	/**
	 * Enter a parse tree produced by the {@code StringType}
	 * labeled alternative in {@link VlangParser#var_type}.
	 * @param ctx the parse tree
	 */
	void enterStringType(VlangParser.StringTypeContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StringType}
	 * labeled alternative in {@link VlangParser#var_type}.
	 * @param ctx the parse tree
	 */
	void exitStringType(VlangParser.StringTypeContext ctx);
	/**
	 * Enter a parse tree produced by the {@code BoolType}
	 * labeled alternative in {@link VlangParser#var_type}.
	 * @param ctx the parse tree
	 */
	void enterBoolType(VlangParser.BoolTypeContext ctx);
	/**
	 * Exit a parse tree produced by the {@code BoolType}
	 * labeled alternative in {@link VlangParser#var_type}.
	 * @param ctx the parse tree
	 */
	void exitBoolType(VlangParser.BoolTypeContext ctx);
	/**
	 * Enter a parse tree produced by the {@code CharType}
	 * labeled alternative in {@link VlangParser#var_type}.
	 * @param ctx the parse tree
	 */
	void enterCharType(VlangParser.CharTypeContext ctx);
	/**
	 * Exit a parse tree produced by the {@code CharType}
	 * labeled alternative in {@link VlangParser#var_type}.
	 * @param ctx the parse tree
	 */
	void exitCharType(VlangParser.CharTypeContext ctx);
	/**
	 * Enter a parse tree produced by the {@code VoidType}
	 * labeled alternative in {@link VlangParser#var_type}.
	 * @param ctx the parse tree
	 */
	void enterVoidType(VlangParser.VoidTypeContext ctx);
	/**
	 * Exit a parse tree produced by the {@code VoidType}
	 * labeled alternative in {@link VlangParser#var_type}.
	 * @param ctx the parse tree
	 */
	void exitVoidType(VlangParser.VoidTypeContext ctx);
	/**
	 * Enter a parse tree produced by the {@code StructVector}
	 * labeled alternative in {@link VlangParser#struct_vector}.
	 * @param ctx the parse tree
	 */
	void enterStructVector(VlangParser.StructVectorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StructVector}
	 * labeled alternative in {@link VlangParser#struct_vector}.
	 * @param ctx the parse tree
	 */
	void exitStructVector(VlangParser.StructVectorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code incremento}
	 * labeled alternative in {@link VlangParser#incredecre}.
	 * @param ctx the parse tree
	 */
	void enterIncremento(VlangParser.IncrementoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code incremento}
	 * labeled alternative in {@link VlangParser#incredecre}.
	 * @param ctx the parse tree
	 */
	void exitIncremento(VlangParser.IncrementoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code decremento}
	 * labeled alternative in {@link VlangParser#incredecre}.
	 * @param ctx the parse tree
	 */
	void enterDecremento(VlangParser.DecrementoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code decremento}
	 * labeled alternative in {@link VlangParser#incredecre}.
	 * @param ctx the parse tree
	 */
	void exitDecremento(VlangParser.DecrementoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code incremento2}
	 * labeled alternative in {@link VlangParser#incredecre}.
	 * @param ctx the parse tree
	 */
	void enterIncremento2(VlangParser.Incremento2Context ctx);
	/**
	 * Exit a parse tree produced by the {@code incremento2}
	 * labeled alternative in {@link VlangParser#incredecre}.
	 * @param ctx the parse tree
	 */
	void exitIncremento2(VlangParser.Incremento2Context ctx);
	/**
	 * Enter a parse tree produced by the {@code decremento2}
	 * labeled alternative in {@link VlangParser#incredecre}.
	 * @param ctx the parse tree
	 */
	void enterDecremento2(VlangParser.Decremento2Context ctx);
	/**
	 * Exit a parse tree produced by the {@code decremento2}
	 * labeled alternative in {@link VlangParser#incredecre}.
	 * @param ctx the parse tree
	 */
	void exitDecremento2(VlangParser.Decremento2Context ctx);
}