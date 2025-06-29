// Generated from /home/luisa/Documentos/2025/Vaqueras/Proyecto 2/OLC2_Proyecto2_G10/compiler/parser/Vlang.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class VlangParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		LEN=10, CAP=11, APPEND=12, MUT=13, BOOLEANO=14, ENTERO=15, FLOAT=16, DECIMAL=17, 
		CADENA=18, CARACTER=19, IF_KW=20, ELSE_KW=21, WHILE_KW=22, FOR_KW=23, 
		IN_KW=24, RETURN_KW=25, SWITCH_KW=26, CASE_KW=27, DEFAULT=28, BREAK_KW=29, 
		CONTINUE_KW=30, FN_KW=31, MAIN_KW=32, ATOI_KW=33, PARSEFLOAT_KW=34, TYPEOF_KW=35, 
		INDEX_KW=36, JOIN_KW=37, PLUS=38, MINUS=39, MUL=40, DIV=41, MOD=42, NOT=43, 
		OR=44, AND=45, EQ=46, NEQ=47, LT=48, LE=49, GT=50, GE=51, ASSIGN=52, INC=53, 
		DEC=54, INC2=55, DEC2=56, COLON=57, PYC=58, LPAREN=59, RPAREN=60, LBRACK=61, 
		LCOR=62, RBRACK=63, RCOR=64, DOT=65, COMMA=66, ID=67, WS=68, LINE_COMMENT=69, 
		BLOCK_COMMENT=70;
	public static final int
		RULE_programa = 0, RULE_stmt = 1, RULE_if_stmt = 2, RULE_if_chain = 3, 
		RULE_else_stmt = 4, RULE_println = 5, RULE_print = 6, RULE_forwhile_stmt = 7, 
		RULE_forClasic_stmt = 8, RULE_for_stmt = 9, RULE_range = 10, RULE_switch_stmt = 11, 
		RULE_case_clause = 12, RULE_default_clause = 13, RULE_slice_stmt = 14, 
		RULE_func_call = 15, RULE_func_dcl = 16, RULE_struct_dcl = 17, RULE_transfer_stmt = 18, 
		RULE_assign_stmt = 19, RULE_decl_stmt = 20, RULE_list_expresiones = 21, 
		RULE_id_pattern = 22, RULE_expresion = 23, RULE_parametros = 24, RULE_func_param = 25, 
		RULE_arg_list = 26, RULE_func_arg = 27, RULE_valores = 28, RULE_valor = 29, 
		RULE_strct_instancia = 30, RULE_struct_prop = 31, RULE_var_type = 32, 
		RULE_struct_vector = 33, RULE_incredecre = 34;
	private static String[] makeRuleNames() {
		return new String[] {
			"programa", "stmt", "if_stmt", "if_chain", "else_stmt", "println", "print", 
			"forwhile_stmt", "forClasic_stmt", "for_stmt", "range", "switch_stmt", 
			"case_clause", "default_clause", "slice_stmt", "func_call", "func_dcl", 
			"struct_dcl", "transfer_stmt", "assign_stmt", "decl_stmt", "list_expresiones", 
			"id_pattern", "expresion", "parametros", "func_param", "arg_list", "func_arg", 
			"valores", "valor", "strct_instancia", "struct_prop", "var_type", "struct_vector", 
			"incredecre"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'println'", "'print'", "'struct'", "'int'", "'float'", "'string'", 
			"'bool'", "'char'", "'void'", "'len'", "'cap'", "'append'", "'mut'", 
			null, null, null, null, null, null, "'if'", "'else'", "'while'", "'for'", 
			"'in'", "'return'", "'switch'", "'case'", "'default'", "'break'", "'continue'", 
			"'fn'", "'main'", "'Atoi'", "'parse_float'", "'TypeOf'", "'indexOf'", 
			"'join'", "'+'", "'-'", "'*'", "'/'", "'%'", "'!'", "'||'", "'&&'", "'=='", 
			"'!='", "'<'", "'<='", "'>'", "'>='", "'='", "'++'", "'--'", "'+='", 
			"'-='", "':'", "';'", "'('", "')'", "'['", "'{'", "']'", "'}'", "'.'", 
			"','"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, "LEN", "CAP", 
			"APPEND", "MUT", "BOOLEANO", "ENTERO", "FLOAT", "DECIMAL", "CADENA", 
			"CARACTER", "IF_KW", "ELSE_KW", "WHILE_KW", "FOR_KW", "IN_KW", "RETURN_KW", 
			"SWITCH_KW", "CASE_KW", "DEFAULT", "BREAK_KW", "CONTINUE_KW", "FN_KW", 
			"MAIN_KW", "ATOI_KW", "PARSEFLOAT_KW", "TYPEOF_KW", "INDEX_KW", "JOIN_KW", 
			"PLUS", "MINUS", "MUL", "DIV", "MOD", "NOT", "OR", "AND", "EQ", "NEQ", 
			"LT", "LE", "GT", "GE", "ASSIGN", "INC", "DEC", "INC2", "DEC2", "COLON", 
			"PYC", "LPAREN", "RPAREN", "LBRACK", "LCOR", "RBRACK", "RCOR", "DOT", 
			"COMMA", "ID", "WS", "LINE_COMMENT", "BLOCK_COMMENT"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Vlang.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public VlangParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProgramaContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(VlangParser.EOF, 0); }
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public ProgramaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_programa; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterPrograma(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitPrograma(this);
		}
	}

	public final ProgramaContext programa() throws RecognitionException {
		ProgramaContext _localctx = new ProgramaContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_programa);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(71); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(70);
				stmt();
				}
				}
				setState(73); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & 3868205070L) != 0) || _la==ID );
			setState(75);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StmtContext extends ParserRuleContext {
		public Decl_stmtContext decl_stmt() {
			return getRuleContext(Decl_stmtContext.class,0);
		}
		public Assign_stmtContext assign_stmt() {
			return getRuleContext(Assign_stmtContext.class,0);
		}
		public If_stmtContext if_stmt() {
			return getRuleContext(If_stmtContext.class,0);
		}
		public For_stmtContext for_stmt() {
			return getRuleContext(For_stmtContext.class,0);
		}
		public Func_callContext func_call() {
			return getRuleContext(Func_callContext.class,0);
		}
		public Func_dclContext func_dcl() {
			return getRuleContext(Func_dclContext.class,0);
		}
		public Struct_dclContext struct_dcl() {
			return getRuleContext(Struct_dclContext.class,0);
		}
		public Strct_instanciaContext strct_instancia() {
			return getRuleContext(Strct_instanciaContext.class,0);
		}
		public Transfer_stmtContext transfer_stmt() {
			return getRuleContext(Transfer_stmtContext.class,0);
		}
		public PrintlnContext println() {
			return getRuleContext(PrintlnContext.class,0);
		}
		public PrintContext print() {
			return getRuleContext(PrintContext.class,0);
		}
		public IncredecreContext incredecre() {
			return getRuleContext(IncredecreContext.class,0);
		}
		public Forwhile_stmtContext forwhile_stmt() {
			return getRuleContext(Forwhile_stmtContext.class,0);
		}
		public ForClasic_stmtContext forClasic_stmt() {
			return getRuleContext(ForClasic_stmtContext.class,0);
		}
		public Switch_stmtContext switch_stmt() {
			return getRuleContext(Switch_stmtContext.class,0);
		}
		public Slice_stmtContext slice_stmt() {
			return getRuleContext(Slice_stmtContext.class,0);
		}
		public StmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stmt; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterStmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitStmt(this);
		}
	}

	public final StmtContext stmt() throws RecognitionException {
		StmtContext _localctx = new StmtContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_stmt);
		try {
			setState(93);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(77);
				decl_stmt();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(78);
				assign_stmt();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(79);
				if_stmt();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(80);
				for_stmt();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(81);
				func_call();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(82);
				func_dcl();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(83);
				struct_dcl();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(84);
				strct_instancia();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(85);
				transfer_stmt();
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(86);
				println();
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(87);
				print();
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(88);
				incredecre();
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(89);
				forwhile_stmt();
				}
				break;
			case 14:
				enterOuterAlt(_localctx, 14);
				{
				setState(90);
				forClasic_stmt();
				}
				break;
			case 15:
				enterOuterAlt(_localctx, 15);
				{
				setState(91);
				switch_stmt();
				}
				break;
			case 16:
				enterOuterAlt(_localctx, 16);
				{
				setState(92);
				slice_stmt();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class If_stmtContext extends ParserRuleContext {
		public If_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_if_stmt; }
	 
		public If_stmtContext() { }
		public void copyFrom(If_stmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IfStmtContext extends If_stmtContext {
		public List<If_chainContext> if_chain() {
			return getRuleContexts(If_chainContext.class);
		}
		public If_chainContext if_chain(int i) {
			return getRuleContext(If_chainContext.class,i);
		}
		public List<TerminalNode> ELSE_KW() { return getTokens(VlangParser.ELSE_KW); }
		public TerminalNode ELSE_KW(int i) {
			return getToken(VlangParser.ELSE_KW, i);
		}
		public Else_stmtContext else_stmt() {
			return getRuleContext(Else_stmtContext.class,0);
		}
		public IfStmtContext(If_stmtContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterIfStmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitIfStmt(this);
		}
	}

	public final If_stmtContext if_stmt() throws RecognitionException {
		If_stmtContext _localctx = new If_stmtContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_if_stmt);
		int _la;
		try {
			int _alt;
			_localctx = new IfStmtContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(95);
			if_chain();
			setState(100);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(96);
					match(ELSE_KW);
					setState(97);
					if_chain();
					}
					} 
				}
				setState(102);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			}
			setState(104);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE_KW) {
				{
				setState(103);
				else_stmt();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class If_chainContext extends ParserRuleContext {
		public If_chainContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_if_chain; }
	 
		public If_chainContext() { }
		public void copyFrom(If_chainContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IfChainContext extends If_chainContext {
		public TerminalNode IF_KW() { return getToken(VlangParser.IF_KW, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode LCOR() { return getToken(VlangParser.LCOR, 0); }
		public TerminalNode RCOR() { return getToken(VlangParser.RCOR, 0); }
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public IfChainContext(If_chainContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterIfChain(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitIfChain(this);
		}
	}

	public final If_chainContext if_chain() throws RecognitionException {
		If_chainContext _localctx = new If_chainContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_if_chain);
		int _la;
		try {
			_localctx = new IfChainContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(106);
			match(IF_KW);
			setState(107);
			expresion(0);
			setState(108);
			match(LCOR);
			setState(112);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 3868205070L) != 0) || _la==ID) {
				{
				{
				setState(109);
				stmt();
				}
				}
				setState(114);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(115);
			match(RCOR);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Else_stmtContext extends ParserRuleContext {
		public Else_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_else_stmt; }
	 
		public Else_stmtContext() { }
		public void copyFrom(Else_stmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ElseStmtContext extends Else_stmtContext {
		public TerminalNode ELSE_KW() { return getToken(VlangParser.ELSE_KW, 0); }
		public TerminalNode LCOR() { return getToken(VlangParser.LCOR, 0); }
		public TerminalNode RCOR() { return getToken(VlangParser.RCOR, 0); }
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public ElseStmtContext(Else_stmtContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterElseStmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitElseStmt(this);
		}
	}

	public final Else_stmtContext else_stmt() throws RecognitionException {
		Else_stmtContext _localctx = new Else_stmtContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_else_stmt);
		int _la;
		try {
			_localctx = new ElseStmtContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(117);
			match(ELSE_KW);
			setState(118);
			match(LCOR);
			setState(122);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 3868205070L) != 0) || _la==ID) {
				{
				{
				setState(119);
				stmt();
				}
				}
				setState(124);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(125);
			match(RCOR);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PrintlnContext extends ParserRuleContext {
		public PrintlnContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_println; }
	 
		public PrintlnContext() { }
		public void copyFrom(PrintlnContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PrintlnStmtContext extends PrintlnContext {
		public TerminalNode LPAREN() { return getToken(VlangParser.LPAREN, 0); }
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public TerminalNode RPAREN() { return getToken(VlangParser.RPAREN, 0); }
		public List<TerminalNode> COMMA() { return getTokens(VlangParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(VlangParser.COMMA, i);
		}
		public PrintlnStmtContext(PrintlnContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterPrintlnStmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitPrintlnStmt(this);
		}
	}

	public final PrintlnContext println() throws RecognitionException {
		PrintlnContext _localctx = new PrintlnContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_println);
		int _la;
		try {
			_localctx = new PrintlnStmtContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(127);
			match(T__0);
			setState(128);
			match(LPAREN);
			setState(129);
			expresion(0);
			setState(134);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(130);
				match(COMMA);
				setState(131);
				expresion(0);
				}
				}
				setState(136);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(137);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PrintContext extends ParserRuleContext {
		public PrintContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_print; }
	 
		public PrintContext() { }
		public void copyFrom(PrintContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PrintStmtContext extends PrintContext {
		public TerminalNode LPAREN() { return getToken(VlangParser.LPAREN, 0); }
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public TerminalNode RPAREN() { return getToken(VlangParser.RPAREN, 0); }
		public List<TerminalNode> COMMA() { return getTokens(VlangParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(VlangParser.COMMA, i);
		}
		public PrintStmtContext(PrintContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterPrintStmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitPrintStmt(this);
		}
	}

	public final PrintContext print() throws RecognitionException {
		PrintContext _localctx = new PrintContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_print);
		int _la;
		try {
			_localctx = new PrintStmtContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(139);
			match(T__1);
			setState(140);
			match(LPAREN);
			setState(141);
			expresion(0);
			setState(146);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(142);
				match(COMMA);
				setState(143);
				expresion(0);
				}
				}
				setState(148);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(149);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Forwhile_stmtContext extends ParserRuleContext {
		public Forwhile_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forwhile_stmt; }
	 
		public Forwhile_stmtContext() { }
		public void copyFrom(Forwhile_stmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class WhileStmtContext extends Forwhile_stmtContext {
		public TerminalNode FOR_KW() { return getToken(VlangParser.FOR_KW, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode LCOR() { return getToken(VlangParser.LCOR, 0); }
		public TerminalNode RCOR() { return getToken(VlangParser.RCOR, 0); }
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public WhileStmtContext(Forwhile_stmtContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterWhileStmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitWhileStmt(this);
		}
	}

	public final Forwhile_stmtContext forwhile_stmt() throws RecognitionException {
		Forwhile_stmtContext _localctx = new Forwhile_stmtContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_forwhile_stmt);
		int _la;
		try {
			_localctx = new WhileStmtContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(151);
			match(FOR_KW);
			setState(152);
			expresion(0);
			setState(153);
			match(LCOR);
			setState(157);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 3868205070L) != 0) || _la==ID) {
				{
				{
				setState(154);
				stmt();
				}
				}
				setState(159);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(160);
			match(RCOR);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ForClasic_stmtContext extends ParserRuleContext {
		public ForClasic_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forClasic_stmt; }
	 
		public ForClasic_stmtContext() { }
		public void copyFrom(ForClasic_stmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ClasicForContext extends ForClasic_stmtContext {
		public TerminalNode FOR_KW() { return getToken(VlangParser.FOR_KW, 0); }
		public Assign_stmtContext assign_stmt() {
			return getRuleContext(Assign_stmtContext.class,0);
		}
		public List<TerminalNode> PYC() { return getTokens(VlangParser.PYC); }
		public TerminalNode PYC(int i) {
			return getToken(VlangParser.PYC, i);
		}
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public IncredecreContext incredecre() {
			return getRuleContext(IncredecreContext.class,0);
		}
		public TerminalNode LCOR() { return getToken(VlangParser.LCOR, 0); }
		public TerminalNode RCOR() { return getToken(VlangParser.RCOR, 0); }
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public ClasicForContext(ForClasic_stmtContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterClasicFor(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitClasicFor(this);
		}
	}

	public final ForClasic_stmtContext forClasic_stmt() throws RecognitionException {
		ForClasic_stmtContext _localctx = new ForClasic_stmtContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_forClasic_stmt);
		int _la;
		try {
			_localctx = new ClasicForContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(162);
			match(FOR_KW);
			setState(163);
			assign_stmt();
			setState(164);
			match(PYC);
			setState(165);
			expresion(0);
			setState(166);
			match(PYC);
			setState(167);
			incredecre();
			setState(168);
			match(LCOR);
			setState(172);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 3868205070L) != 0) || _la==ID) {
				{
				{
				setState(169);
				stmt();
				}
				}
				setState(174);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(175);
			match(RCOR);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class For_stmtContext extends ParserRuleContext {
		public For_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_for_stmt; }
	 
		public For_stmtContext() { }
		public void copyFrom(For_stmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ForStmtContext extends For_stmtContext {
		public TerminalNode FOR_KW() { return getToken(VlangParser.FOR_KW, 0); }
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode IN_KW() { return getToken(VlangParser.IN_KW, 0); }
		public TerminalNode LCOR() { return getToken(VlangParser.LCOR, 0); }
		public TerminalNode RCOR() { return getToken(VlangParser.RCOR, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public RangeContext range() {
			return getRuleContext(RangeContext.class,0);
		}
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public ForStmtContext(For_stmtContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterForStmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitForStmt(this);
		}
	}

	public final For_stmtContext for_stmt() throws RecognitionException {
		For_stmtContext _localctx = new For_stmtContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_for_stmt);
		int _la;
		try {
			_localctx = new ForStmtContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(177);
			match(FOR_KW);
			setState(178);
			match(ID);
			setState(179);
			match(IN_KW);
			setState(182);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				{
				setState(180);
				expresion(0);
				}
				break;
			case 2:
				{
				setState(181);
				range();
				}
				break;
			}
			setState(184);
			match(LCOR);
			setState(188);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 3868205070L) != 0) || _la==ID) {
				{
				{
				setState(185);
				stmt();
				}
				}
				setState(190);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(191);
			match(RCOR);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RangeContext extends ParserRuleContext {
		public RangeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_range; }
	 
		public RangeContext() { }
		public void copyFrom(RangeContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class NumericRangeContext extends RangeContext {
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public List<TerminalNode> DOT() { return getTokens(VlangParser.DOT); }
		public TerminalNode DOT(int i) {
			return getToken(VlangParser.DOT, i);
		}
		public NumericRangeContext(RangeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterNumericRange(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitNumericRange(this);
		}
	}

	public final RangeContext range() throws RecognitionException {
		RangeContext _localctx = new RangeContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_range);
		try {
			_localctx = new NumericRangeContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(193);
			expresion(0);
			setState(194);
			match(DOT);
			setState(195);
			match(DOT);
			setState(196);
			match(DOT);
			setState(197);
			expresion(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Switch_stmtContext extends ParserRuleContext {
		public Switch_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switch_stmt; }
	 
		public Switch_stmtContext() { }
		public void copyFrom(Switch_stmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SwitchStmtContext extends Switch_stmtContext {
		public TerminalNode SWITCH_KW() { return getToken(VlangParser.SWITCH_KW, 0); }
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode LCOR() { return getToken(VlangParser.LCOR, 0); }
		public TerminalNode RCOR() { return getToken(VlangParser.RCOR, 0); }
		public List<Case_clauseContext> case_clause() {
			return getRuleContexts(Case_clauseContext.class);
		}
		public Case_clauseContext case_clause(int i) {
			return getRuleContext(Case_clauseContext.class,i);
		}
		public Default_clauseContext default_clause() {
			return getRuleContext(Default_clauseContext.class,0);
		}
		public SwitchStmtContext(Switch_stmtContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterSwitchStmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitSwitchStmt(this);
		}
	}

	public final Switch_stmtContext switch_stmt() throws RecognitionException {
		Switch_stmtContext _localctx = new Switch_stmtContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_switch_stmt);
		int _la;
		try {
			_localctx = new SwitchStmtContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(199);
			match(SWITCH_KW);
			setState(200);
			match(ID);
			setState(201);
			match(LCOR);
			setState(203); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(202);
				case_clause();
				}
				}
				setState(205); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==CASE_KW );
			setState(208);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DEFAULT) {
				{
				setState(207);
				default_clause();
				}
			}

			setState(210);
			match(RCOR);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Case_clauseContext extends ParserRuleContext {
		public Case_clauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_case_clause; }
	 
		public Case_clauseContext() { }
		public void copyFrom(Case_clauseContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CaseClauseContext extends Case_clauseContext {
		public TerminalNode CASE_KW() { return getToken(VlangParser.CASE_KW, 0); }
		public TerminalNode COLON() { return getToken(VlangParser.COLON, 0); }
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(VlangParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(VlangParser.COMMA, i);
		}
		public CaseClauseContext(Case_clauseContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterCaseClause(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitCaseClause(this);
		}
	}

	public final Case_clauseContext case_clause() throws RecognitionException {
		Case_clauseContext _localctx = new Case_clauseContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_case_clause);
		int _la;
		try {
			_localctx = new CaseClauseContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(212);
			match(CASE_KW);
			{
			setState(213);
			expresion(0);
			setState(218);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(214);
				match(COMMA);
				setState(215);
				expresion(0);
				}
				}
				setState(220);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
			setState(221);
			match(COLON);
			setState(225);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 3868205070L) != 0) || _la==ID) {
				{
				{
				setState(222);
				stmt();
				}
				}
				setState(227);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Default_clauseContext extends ParserRuleContext {
		public Default_clauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_default_clause; }
	 
		public Default_clauseContext() { }
		public void copyFrom(Default_clauseContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DefaultClauseContext extends Default_clauseContext {
		public TerminalNode DEFAULT() { return getToken(VlangParser.DEFAULT, 0); }
		public TerminalNode COLON() { return getToken(VlangParser.COLON, 0); }
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public DefaultClauseContext(Default_clauseContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterDefaultClause(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitDefaultClause(this);
		}
	}

	public final Default_clauseContext default_clause() throws RecognitionException {
		Default_clauseContext _localctx = new Default_clauseContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_default_clause);
		int _la;
		try {
			_localctx = new DefaultClauseContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(228);
			match(DEFAULT);
			setState(229);
			match(COLON);
			setState(233);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 3868205070L) != 0) || _la==ID) {
				{
				{
				setState(230);
				stmt();
				}
				}
				setState(235);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Slice_stmtContext extends ParserRuleContext {
		public Slice_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_slice_stmt; }
	 
		public Slice_stmtContext() { }
		public void copyFrom(Slice_stmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DeclSlicesContext extends Slice_stmtContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode ASSIGN() { return getToken(VlangParser.ASSIGN, 0); }
		public TerminalNode LBRACK() { return getToken(VlangParser.LBRACK, 0); }
		public TerminalNode RBRACK() { return getToken(VlangParser.RBRACK, 0); }
		public Var_typeContext var_type() {
			return getRuleContext(Var_typeContext.class,0);
		}
		public TerminalNode LCOR() { return getToken(VlangParser.LCOR, 0); }
		public TerminalNode RCOR() { return getToken(VlangParser.RCOR, 0); }
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(VlangParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(VlangParser.COMMA, i);
		}
		public DeclSlicesContext(Slice_stmtContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterDeclSlices(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitDeclSlices(this);
		}
	}

	public final Slice_stmtContext slice_stmt() throws RecognitionException {
		Slice_stmtContext _localctx = new Slice_stmtContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_slice_stmt);
		int _la;
		try {
			_localctx = new DeclSlicesContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(236);
			match(ID);
			setState(237);
			match(ASSIGN);
			setState(238);
			match(LBRACK);
			setState(239);
			match(RBRACK);
			setState(240);
			var_type();
			setState(241);
			match(LCOR);
			{
			setState(242);
			expresion(0);
			setState(247);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(243);
				match(COMMA);
				setState(244);
				expresion(0);
				}
				}
				setState(249);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
			setState(250);
			match(RCOR);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Func_callContext extends ParserRuleContext {
		public Func_callContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_func_call; }
	 
		public Func_callContext() { }
		public void copyFrom(Func_callContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FuncCallContext extends Func_callContext {
		public Id_patternContext id_pattern() {
			return getRuleContext(Id_patternContext.class,0);
		}
		public TerminalNode LPAREN() { return getToken(VlangParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(VlangParser.RPAREN, 0); }
		public ParametrosContext parametros() {
			return getRuleContext(ParametrosContext.class,0);
		}
		public FuncCallContext(Func_callContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterFuncCall(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitFuncCall(this);
		}
	}

	public final Func_callContext func_call() throws RecognitionException {
		Func_callContext _localctx = new Func_callContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_func_call);
		int _la;
		try {
			_localctx = new FuncCallContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(252);
			id_pattern();
			setState(253);
			match(LPAREN);
			setState(255);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ID) {
				{
				setState(254);
				parametros();
				}
			}

			setState(257);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Func_dclContext extends ParserRuleContext {
		public Func_dclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_func_dcl; }
	 
		public Func_dclContext() { }
		public void copyFrom(Func_dclContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FuncDeclContext extends Func_dclContext {
		public TerminalNode FN_KW() { return getToken(VlangParser.FN_KW, 0); }
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode LPAREN() { return getToken(VlangParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(VlangParser.RPAREN, 0); }
		public TerminalNode LCOR() { return getToken(VlangParser.LCOR, 0); }
		public TerminalNode RCOR() { return getToken(VlangParser.RCOR, 0); }
		public Arg_listContext arg_list() {
			return getRuleContext(Arg_listContext.class,0);
		}
		public Var_typeContext var_type() {
			return getRuleContext(Var_typeContext.class,0);
		}
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public FuncDeclContext(Func_dclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterFuncDecl(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitFuncDecl(this);
		}
	}

	public final Func_dclContext func_dcl() throws RecognitionException {
		Func_dclContext _localctx = new Func_dclContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_func_dcl);
		int _la;
		try {
			_localctx = new FuncDeclContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(259);
			match(FN_KW);
			setState(260);
			match(ID);
			setState(261);
			match(LPAREN);
			setState(263);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ID) {
				{
				setState(262);
				arg_list();
				}
			}

			setState(265);
			match(RPAREN);
			setState(267);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1008L) != 0)) {
				{
				setState(266);
				var_type();
				}
			}

			setState(269);
			match(LCOR);
			setState(273);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 3868205070L) != 0) || _la==ID) {
				{
				{
				setState(270);
				stmt();
				}
				}
				setState(275);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(276);
			match(RCOR);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Struct_dclContext extends ParserRuleContext {
		public Struct_dclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_struct_dcl; }
	 
		public Struct_dclContext() { }
		public void copyFrom(Struct_dclContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructDeclContext extends Struct_dclContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode LCOR() { return getToken(VlangParser.LCOR, 0); }
		public TerminalNode RCOR() { return getToken(VlangParser.RCOR, 0); }
		public List<Assign_stmtContext> assign_stmt() {
			return getRuleContexts(Assign_stmtContext.class);
		}
		public Assign_stmtContext assign_stmt(int i) {
			return getRuleContext(Assign_stmtContext.class,i);
		}
		public StructDeclContext(Struct_dclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterStructDecl(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitStructDecl(this);
		}
	}

	public final Struct_dclContext struct_dcl() throws RecognitionException {
		Struct_dclContext _localctx = new Struct_dclContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_struct_dcl);
		int _la;
		try {
			_localctx = new StructDeclContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(278);
			match(T__2);
			setState(279);
			match(ID);
			setState(280);
			match(LCOR);
			setState(284);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==ID) {
				{
				{
				setState(281);
				assign_stmt();
				}
				}
				setState(286);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(287);
			match(RCOR);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Transfer_stmtContext extends ParserRuleContext {
		public Transfer_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_transfer_stmt; }
	 
		public Transfer_stmtContext() { }
		public void copyFrom(Transfer_stmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ContinueStmtContext extends Transfer_stmtContext {
		public TerminalNode CONTINUE_KW() { return getToken(VlangParser.CONTINUE_KW, 0); }
		public ContinueStmtContext(Transfer_stmtContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterContinueStmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitContinueStmt(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BreakStmtContext extends Transfer_stmtContext {
		public TerminalNode BREAK_KW() { return getToken(VlangParser.BREAK_KW, 0); }
		public BreakStmtContext(Transfer_stmtContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterBreakStmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitBreakStmt(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ReturnStmtContext extends Transfer_stmtContext {
		public TerminalNode RETURN_KW() { return getToken(VlangParser.RETURN_KW, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public ReturnStmtContext(Transfer_stmtContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterReturnStmt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitReturnStmt(this);
		}
	}

	public final Transfer_stmtContext transfer_stmt() throws RecognitionException {
		Transfer_stmtContext _localctx = new Transfer_stmtContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_transfer_stmt);
		try {
			setState(295);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case RETURN_KW:
				_localctx = new ReturnStmtContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(289);
				match(RETURN_KW);
				setState(291);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,23,_ctx) ) {
				case 1:
					{
					setState(290);
					expresion(0);
					}
					break;
				}
				}
				break;
			case BREAK_KW:
				_localctx = new BreakStmtContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(293);
				match(BREAK_KW);
				}
				break;
			case CONTINUE_KW:
				_localctx = new ContinueStmtContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(294);
				match(CONTINUE_KW);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Assign_stmtContext extends ParserRuleContext {
		public Assign_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assign_stmt; }
	 
		public Assign_stmtContext() { }
		public void copyFrom(Assign_stmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DirectAssignContext extends Assign_stmtContext {
		public Id_patternContext id_pattern() {
			return getRuleContext(Id_patternContext.class,0);
		}
		public TerminalNode ASSIGN() { return getToken(VlangParser.ASSIGN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public DirectAssignContext(Assign_stmtContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterDirectAssign(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitDirectAssign(this);
		}
	}

	public final Assign_stmtContext assign_stmt() throws RecognitionException {
		Assign_stmtContext _localctx = new Assign_stmtContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_assign_stmt);
		try {
			_localctx = new DirectAssignContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(297);
			id_pattern();
			setState(298);
			match(ASSIGN);
			setState(299);
			expresion(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Decl_stmtContext extends ParserRuleContext {
		public Decl_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_decl_stmt; }
	 
		public Decl_stmtContext() { }
		public void copyFrom(Decl_stmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DeclAssignContext extends Decl_stmtContext {
		public TerminalNode MUT() { return getToken(VlangParser.MUT, 0); }
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode ASSIGN() { return getToken(VlangParser.ASSIGN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public DeclAssignContext(Decl_stmtContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterDeclAssign(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitDeclAssign(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DeclAssignTypeContext extends Decl_stmtContext {
		public TerminalNode MUT() { return getToken(VlangParser.MUT, 0); }
		public Id_patternContext id_pattern() {
			return getRuleContext(Id_patternContext.class,0);
		}
		public Var_typeContext var_type() {
			return getRuleContext(Var_typeContext.class,0);
		}
		public TerminalNode ASSIGN() { return getToken(VlangParser.ASSIGN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public DeclAssignTypeContext(Decl_stmtContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterDeclAssignType(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitDeclAssignType(this);
		}
	}

	public final Decl_stmtContext decl_stmt() throws RecognitionException {
		Decl_stmtContext _localctx = new Decl_stmtContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_decl_stmt);
		try {
			setState(311);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
			case 1:
				_localctx = new DeclAssignContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(301);
				match(MUT);
				setState(302);
				match(ID);
				setState(303);
				match(ASSIGN);
				setState(304);
				expresion(0);
				}
				break;
			case 2:
				_localctx = new DeclAssignTypeContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(305);
				match(MUT);
				setState(306);
				id_pattern();
				setState(307);
				var_type();
				setState(308);
				match(ASSIGN);
				setState(309);
				expresion(0);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class List_expresionesContext extends ParserRuleContext {
		public List_expresionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_list_expresiones; }
	 
		public List_expresionesContext() { }
		public void copyFrom(List_expresionesContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ListExpressionsContext extends List_expresionesContext {
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(VlangParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(VlangParser.COMMA, i);
		}
		public ListExpressionsContext(List_expresionesContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterListExpressions(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitListExpressions(this);
		}
	}

	public final List_expresionesContext list_expresiones() throws RecognitionException {
		List_expresionesContext _localctx = new List_expresionesContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_list_expresiones);
		int _la;
		try {
			_localctx = new ListExpressionsContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(313);
			expresion(0);
			setState(318);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(314);
				match(COMMA);
				setState(315);
				expresion(0);
				}
				}
				setState(320);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Id_patternContext extends ParserRuleContext {
		public Id_patternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_id_pattern; }
	 
		public Id_patternContext() { }
		public void copyFrom(Id_patternContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IdPatternContext extends Id_patternContext {
		public List<TerminalNode> ID() { return getTokens(VlangParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(VlangParser.ID, i);
		}
		public List<TerminalNode> DOT() { return getTokens(VlangParser.DOT); }
		public TerminalNode DOT(int i) {
			return getToken(VlangParser.DOT, i);
		}
		public IdPatternContext(Id_patternContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterIdPattern(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitIdPattern(this);
		}
	}

	public final Id_patternContext id_pattern() throws RecognitionException {
		Id_patternContext _localctx = new Id_patternContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_id_pattern);
		int _la;
		try {
			_localctx = new IdPatternContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(321);
			match(ID);
			setState(326);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOT) {
				{
				{
				setState(322);
				match(DOT);
				setState(323);
				match(ID);
				}
				}
				setState(328);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExpresionContext extends ParserRuleContext {
		public ExpresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expresion; }
	 
		public ExpresionContext() { }
		public void copyFrom(ExpresionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IncredecrContext extends ExpresionContext {
		public IncredecreContext incredecre() {
			return getRuleContext(IncredecreContext.class,0);
		}
		public IncredecrContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterIncredecr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitIncredecr(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Fn_lenxContext extends ExpresionContext {
		public TerminalNode LEN() { return getToken(VlangParser.LEN, 0); }
		public TerminalNode LPAREN() { return getToken(VlangParser.LPAREN, 0); }
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public TerminalNode COMMA() { return getToken(VlangParser.COMMA, 0); }
		public TerminalNode RPAREN() { return getToken(VlangParser.RPAREN, 0); }
		public Fn_lenxContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterFn_lenx(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitFn_lenx(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ValorexpresionContext extends ExpresionContext {
		public ValorContext valor() {
			return getRuleContext(ValorContext.class,0);
		}
		public ValorexpresionContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterValorexpresion(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitValorexpresion(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AsignacionforContext extends ExpresionContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode ASSIGN() { return getToken(VlangParser.ASSIGN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public AsignacionforContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterAsignacionfor(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitAsignacionfor(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpdotexpContext extends ExpresionContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode DOT() { return getToken(VlangParser.DOT, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public ExpdotexpContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterExpdotexp(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitExpdotexp(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Fn_typeOFContext extends ExpresionContext {
		public TerminalNode TYPEOF_KW() { return getToken(VlangParser.TYPEOF_KW, 0); }
		public TerminalNode LPAREN() { return getToken(VlangParser.LPAREN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(VlangParser.RPAREN, 0); }
		public Fn_typeOFContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterFn_typeOF(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitFn_typeOF(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BinaryExpContext extends ExpresionContext {
		public ExpresionContext left;
		public Token op;
		public ExpresionContext right;
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public TerminalNode MUL() { return getToken(VlangParser.MUL, 0); }
		public TerminalNode DIV() { return getToken(VlangParser.DIV, 0); }
		public TerminalNode MOD() { return getToken(VlangParser.MOD, 0); }
		public TerminalNode PLUS() { return getToken(VlangParser.PLUS, 0); }
		public TerminalNode MINUS() { return getToken(VlangParser.MINUS, 0); }
		public TerminalNode LT() { return getToken(VlangParser.LT, 0); }
		public TerminalNode LE() { return getToken(VlangParser.LE, 0); }
		public TerminalNode GT() { return getToken(VlangParser.GT, 0); }
		public TerminalNode GE() { return getToken(VlangParser.GE, 0); }
		public TerminalNode EQ() { return getToken(VlangParser.EQ, 0); }
		public TerminalNode NEQ() { return getToken(VlangParser.NEQ, 0); }
		public TerminalNode AND() { return getToken(VlangParser.AND, 0); }
		public TerminalNode OR() { return getToken(VlangParser.OR, 0); }
		public BinaryExpContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterBinaryExp(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitBinaryExp(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CorchetesexpreContext extends ExpresionContext {
		public TerminalNode LBRACK() { return getToken(VlangParser.LBRACK, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode RBRACK() { return getToken(VlangParser.RBRACK, 0); }
		public CorchetesexpreContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterCorchetesexpre(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitCorchetesexpre(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FuncionexpreContext extends ExpresionContext {
		public Func_callContext func_call() {
			return getRuleContext(Func_callContext.class,0);
		}
		public FuncionexpreContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterFuncionexpre(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitFuncionexpre(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class UnarioContext extends ExpresionContext {
		public Token op;
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode NOT() { return getToken(VlangParser.NOT, 0); }
		public TerminalNode MINUS() { return getToken(VlangParser.MINUS, 0); }
		public UnarioContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterUnario(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitUnario(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ParentesisexpreContext extends ExpresionContext {
		public TerminalNode LPAREN() { return getToken(VlangParser.LPAREN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(VlangParser.RPAREN, 0); }
		public ParentesisexpreContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterParentesisexpre(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitParentesisexpre(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Fn_parsefloatContext extends ExpresionContext {
		public TerminalNode PARSEFLOAT_KW() { return getToken(VlangParser.PARSEFLOAT_KW, 0); }
		public TerminalNode LPAREN() { return getToken(VlangParser.LPAREN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(VlangParser.RPAREN, 0); }
		public Fn_parsefloatContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterFn_parsefloat(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitFn_parsefloat(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IndexAccesoContext extends ExpresionContext {
		public Id_patternContext id_pattern() {
			return getRuleContext(Id_patternContext.class,0);
		}
		public TerminalNode LBRACK() { return getToken(VlangParser.LBRACK, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode RBRACK() { return getToken(VlangParser.RBRACK, 0); }
		public IndexAccesoContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterIndexAcceso(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitIndexAcceso(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IdContext extends ExpresionContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public IdContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterId(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitId(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Expdotexp1Context extends ExpresionContext {
		public List<TerminalNode> ID() { return getTokens(VlangParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(VlangParser.ID, i);
		}
		public TerminalNode DOT() { return getToken(VlangParser.DOT, 0); }
		public Expdotexp1Context(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterExpdotexp1(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitExpdotexp1(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Fn_indexOfContext extends ExpresionContext {
		public TerminalNode INDEX_KW() { return getToken(VlangParser.INDEX_KW, 0); }
		public TerminalNode LPAREN() { return getToken(VlangParser.LPAREN, 0); }
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public TerminalNode COMMA() { return getToken(VlangParser.COMMA, 0); }
		public TerminalNode RPAREN() { return getToken(VlangParser.RPAREN, 0); }
		public Fn_indexOfContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterFn_indexOf(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitFn_indexOf(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Fn_joinContext extends ExpresionContext {
		public TerminalNode JOIN_KW() { return getToken(VlangParser.JOIN_KW, 0); }
		public TerminalNode LPAREN() { return getToken(VlangParser.LPAREN, 0); }
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public TerminalNode COMMA() { return getToken(VlangParser.COMMA, 0); }
		public TerminalNode RPAREN() { return getToken(VlangParser.RPAREN, 0); }
		public Fn_joinContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterFn_join(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitFn_join(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Fn_atoiContext extends ExpresionContext {
		public TerminalNode ATOI_KW() { return getToken(VlangParser.ATOI_KW, 0); }
		public TerminalNode LPAREN() { return getToken(VlangParser.LPAREN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(VlangParser.RPAREN, 0); }
		public Fn_atoiContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterFn_atoi(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitFn_atoi(this);
		}
	}

	public final ExpresionContext expresion() throws RecognitionException {
		return expresion(0);
	}

	private ExpresionContext expresion(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpresionContext _localctx = new ExpresionContext(_ctx, _parentState);
		ExpresionContext _prevctx = _localctx;
		int _startState = 46;
		enterRecursionRule(_localctx, 46, RULE_expresion, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(394);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,28,_ctx) ) {
			case 1:
				{
				_localctx = new ValorexpresionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(330);
				valor();
				}
				break;
			case 2:
				{
				_localctx = new FuncionexpreContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(331);
				func_call();
				}
				break;
			case 3:
				{
				_localctx = new ParentesisexpreContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(332);
				match(LPAREN);
				setState(333);
				expresion(0);
				setState(334);
				match(RPAREN);
				}
				break;
			case 4:
				{
				_localctx = new CorchetesexpreContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(336);
				match(LBRACK);
				setState(337);
				expresion(0);
				setState(338);
				match(RBRACK);
				}
				break;
			case 5:
				{
				_localctx = new UnarioContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(340);
				((UnarioContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==MINUS || _la==NOT) ) {
					((UnarioContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(341);
				expresion(19);
				}
				break;
			case 6:
				{
				_localctx = new IdContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(342);
				match(ID);
				}
				break;
			case 7:
				{
				_localctx = new IncredecrContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(343);
				incredecre();
				}
				break;
			case 8:
				{
				_localctx = new Expdotexp1Context(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(344);
				match(ID);
				setState(345);
				match(DOT);
				setState(346);
				match(ID);
				}
				break;
			case 9:
				{
				_localctx = new ExpdotexpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(347);
				match(ID);
				setState(348);
				match(DOT);
				setState(349);
				expresion(9);
				}
				break;
			case 10:
				{
				_localctx = new AsignacionforContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(350);
				match(ID);
				setState(351);
				match(ASSIGN);
				setState(352);
				expresion(8);
				}
				break;
			case 11:
				{
				_localctx = new IndexAccesoContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(353);
				id_pattern();
				setState(354);
				match(LBRACK);
				setState(355);
				expresion(0);
				setState(356);
				match(RBRACK);
				}
				break;
			case 12:
				{
				_localctx = new Fn_atoiContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(358);
				match(ATOI_KW);
				setState(359);
				match(LPAREN);
				setState(360);
				expresion(0);
				setState(361);
				match(RPAREN);
				}
				break;
			case 13:
				{
				_localctx = new Fn_parsefloatContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(363);
				match(PARSEFLOAT_KW);
				setState(364);
				match(LPAREN);
				setState(365);
				expresion(0);
				setState(366);
				match(RPAREN);
				}
				break;
			case 14:
				{
				_localctx = new Fn_typeOFContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(368);
				match(TYPEOF_KW);
				setState(369);
				match(LPAREN);
				setState(370);
				expresion(0);
				setState(371);
				match(RPAREN);
				}
				break;
			case 15:
				{
				_localctx = new Fn_indexOfContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(373);
				match(INDEX_KW);
				setState(374);
				match(LPAREN);
				setState(375);
				expresion(0);
				setState(376);
				match(COMMA);
				setState(377);
				expresion(0);
				setState(378);
				match(RPAREN);
				}
				break;
			case 16:
				{
				_localctx = new Fn_lenxContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(380);
				match(LEN);
				setState(381);
				match(LPAREN);
				setState(382);
				expresion(0);
				setState(383);
				match(COMMA);
				setState(384);
				expresion(0);
				setState(385);
				match(RPAREN);
				}
				break;
			case 17:
				{
				_localctx = new Fn_joinContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(387);
				match(JOIN_KW);
				setState(388);
				match(LPAREN);
				setState(389);
				expresion(0);
				setState(390);
				match(COMMA);
				setState(391);
				expresion(0);
				setState(392);
				match(RPAREN);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(416);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,30,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(414);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,29,_ctx) ) {
					case 1:
						{
						_localctx = new BinaryExpContext(new ExpresionContext(_parentctx, _parentState));
						((BinaryExpContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(396);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(397);
						((BinaryExpContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 7696581394432L) != 0)) ) {
							((BinaryExpContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(398);
						((BinaryExpContext)_localctx).right = expresion(19);
						}
						break;
					case 2:
						{
						_localctx = new BinaryExpContext(new ExpresionContext(_parentctx, _parentState));
						((BinaryExpContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(399);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(400);
						((BinaryExpContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==PLUS || _la==MINUS) ) {
							((BinaryExpContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(401);
						((BinaryExpContext)_localctx).right = expresion(18);
						}
						break;
					case 3:
						{
						_localctx = new BinaryExpContext(new ExpresionContext(_parentctx, _parentState));
						((BinaryExpContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(402);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(403);
						((BinaryExpContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 4222124650659840L) != 0)) ) {
							((BinaryExpContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(404);
						((BinaryExpContext)_localctx).right = expresion(17);
						}
						break;
					case 4:
						{
						_localctx = new BinaryExpContext(new ExpresionContext(_parentctx, _parentState));
						((BinaryExpContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(405);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(406);
						((BinaryExpContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==EQ || _la==NEQ) ) {
							((BinaryExpContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(407);
						((BinaryExpContext)_localctx).right = expresion(16);
						}
						break;
					case 5:
						{
						_localctx = new BinaryExpContext(new ExpresionContext(_parentctx, _parentState));
						((BinaryExpContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(408);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(409);
						((BinaryExpContext)_localctx).op = match(AND);
						setState(410);
						((BinaryExpContext)_localctx).right = expresion(15);
						}
						break;
					case 6:
						{
						_localctx = new BinaryExpContext(new ExpresionContext(_parentctx, _parentState));
						((BinaryExpContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(411);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(412);
						((BinaryExpContext)_localctx).op = match(OR);
						setState(413);
						((BinaryExpContext)_localctx).right = expresion(14);
						}
						break;
					}
					} 
				}
				setState(418);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,30,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ParametrosContext extends ParserRuleContext {
		public ParametrosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametros; }
	 
		public ParametrosContext() { }
		public void copyFrom(ParametrosContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ParamListContext extends ParametrosContext {
		public List<Func_paramContext> func_param() {
			return getRuleContexts(Func_paramContext.class);
		}
		public Func_paramContext func_param(int i) {
			return getRuleContext(Func_paramContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(VlangParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(VlangParser.COMMA, i);
		}
		public ParamListContext(ParametrosContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterParamList(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitParamList(this);
		}
	}

	public final ParametrosContext parametros() throws RecognitionException {
		ParametrosContext _localctx = new ParametrosContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_parametros);
		int _la;
		try {
			_localctx = new ParamListContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(419);
			func_param();
			setState(424);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(420);
				match(COMMA);
				setState(421);
				func_param();
				}
				}
				setState(426);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Func_paramContext extends ParserRuleContext {
		public Func_paramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_func_param; }
	 
		public Func_paramContext() { }
		public void copyFrom(Func_paramContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FuncParamContext extends Func_paramContext {
		public Id_patternContext id_pattern() {
			return getRuleContext(Id_patternContext.class,0);
		}
		public FuncParamContext(Func_paramContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterFuncParam(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitFuncParam(this);
		}
	}

	public final Func_paramContext func_param() throws RecognitionException {
		Func_paramContext _localctx = new Func_paramContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_func_param);
		try {
			_localctx = new FuncParamContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(427);
			id_pattern();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Arg_listContext extends ParserRuleContext {
		public Arg_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arg_list; }
	 
		public Arg_listContext() { }
		public void copyFrom(Arg_listContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ArgListContext extends Arg_listContext {
		public List<Func_argContext> func_arg() {
			return getRuleContexts(Func_argContext.class);
		}
		public Func_argContext func_arg(int i) {
			return getRuleContext(Func_argContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(VlangParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(VlangParser.COMMA, i);
		}
		public ArgListContext(Arg_listContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterArgList(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitArgList(this);
		}
	}

	public final Arg_listContext arg_list() throws RecognitionException {
		Arg_listContext _localctx = new Arg_listContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_arg_list);
		int _la;
		try {
			_localctx = new ArgListContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(429);
			func_arg();
			setState(434);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(430);
				match(COMMA);
				setState(431);
				func_arg();
				}
				}
				setState(436);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Func_argContext extends ParserRuleContext {
		public Func_argContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_func_arg; }
	 
		public Func_argContext() { }
		public void copyFrom(Func_argContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FuncArgContext extends Func_argContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public Var_typeContext var_type() {
			return getRuleContext(Var_typeContext.class,0);
		}
		public FuncArgContext(Func_argContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterFuncArg(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitFuncArg(this);
		}
	}

	public final Func_argContext func_arg() throws RecognitionException {
		Func_argContext _localctx = new Func_argContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_func_arg);
		try {
			_localctx = new FuncArgContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(437);
			match(ID);
			setState(438);
			var_type();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ValoresContext extends ParserRuleContext {
		public ValorContext valor() {
			return getRuleContext(ValorContext.class,0);
		}
		public ValoresContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_valores; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterValores(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitValores(this);
		}
	}

	public final ValoresContext valores() throws RecognitionException {
		ValoresContext _localctx = new ValoresContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_valores);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(440);
			valor();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ValorContext extends ParserRuleContext {
		public ValorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_valor; }
	 
		public ValorContext() { }
		public void copyFrom(ValorContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ValorDecimalContext extends ValorContext {
		public TerminalNode DECIMAL() { return getToken(VlangParser.DECIMAL, 0); }
		public ValorDecimalContext(ValorContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterValorDecimal(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitValorDecimal(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ValorEnteroContext extends ValorContext {
		public TerminalNode ENTERO() { return getToken(VlangParser.ENTERO, 0); }
		public ValorEnteroContext(ValorContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterValorEntero(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitValorEntero(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ValorBooleanoContext extends ValorContext {
		public TerminalNode BOOLEANO() { return getToken(VlangParser.BOOLEANO, 0); }
		public ValorBooleanoContext(ValorContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterValorBooleano(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitValorBooleano(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ValorCaracterContext extends ValorContext {
		public TerminalNode CARACTER() { return getToken(VlangParser.CARACTER, 0); }
		public ValorCaracterContext(ValorContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterValorCaracter(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitValorCaracter(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ValorCadenaContext extends ValorContext {
		public TerminalNode CADENA() { return getToken(VlangParser.CADENA, 0); }
		public ValorCadenaContext(ValorContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterValorCadena(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitValorCadena(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ValorFloatContext extends ValorContext {
		public TerminalNode FLOAT() { return getToken(VlangParser.FLOAT, 0); }
		public ValorFloatContext(ValorContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterValorFloat(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitValorFloat(this);
		}
	}

	public final ValorContext valor() throws RecognitionException {
		ValorContext _localctx = new ValorContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_valor);
		try {
			setState(448);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ENTERO:
				_localctx = new ValorEnteroContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(442);
				match(ENTERO);
				}
				break;
			case DECIMAL:
				_localctx = new ValorDecimalContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(443);
				match(DECIMAL);
				}
				break;
			case CADENA:
				_localctx = new ValorCadenaContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(444);
				match(CADENA);
				}
				break;
			case BOOLEANO:
				_localctx = new ValorBooleanoContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(445);
				match(BOOLEANO);
				}
				break;
			case CARACTER:
				_localctx = new ValorCaracterContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(446);
				match(CARACTER);
				}
				break;
			case FLOAT:
				_localctx = new ValorFloatContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(447);
				match(FLOAT);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Strct_instanciaContext extends ParserRuleContext {
		public Strct_instanciaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_strct_instancia; }
	 
		public Strct_instanciaContext() { }
		public void copyFrom(Strct_instanciaContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructInstanciaContext extends Strct_instanciaContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode LCOR() { return getToken(VlangParser.LCOR, 0); }
		public TerminalNode RCOR() { return getToken(VlangParser.RCOR, 0); }
		public List<Struct_propContext> struct_prop() {
			return getRuleContexts(Struct_propContext.class);
		}
		public Struct_propContext struct_prop(int i) {
			return getRuleContext(Struct_propContext.class,i);
		}
		public StructInstanciaContext(Strct_instanciaContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterStructInstancia(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitStructInstancia(this);
		}
	}

	public final Strct_instanciaContext strct_instancia() throws RecognitionException {
		Strct_instanciaContext _localctx = new Strct_instanciaContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_strct_instancia);
		int _la;
		try {
			_localctx = new StructInstanciaContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(450);
			match(T__2);
			setState(451);
			match(ID);
			setState(452);
			match(LCOR);
			setState(456);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1008L) != 0)) {
				{
				{
				setState(453);
				struct_prop();
				}
				}
				setState(458);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(459);
			match(RCOR);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Struct_propContext extends ParserRuleContext {
		public Struct_propContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_struct_prop; }
	 
		public Struct_propContext() { }
		public void copyFrom(Struct_propContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructAttrContext extends Struct_propContext {
		public Var_typeContext var_type() {
			return getRuleContext(Var_typeContext.class,0);
		}
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode COLON() { return getToken(VlangParser.COLON, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode COMMA() { return getToken(VlangParser.COMMA, 0); }
		public StructAttrContext(Struct_propContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterStructAttr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitStructAttr(this);
		}
	}

	public final Struct_propContext struct_prop() throws RecognitionException {
		Struct_propContext _localctx = new Struct_propContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_struct_prop);
		int _la;
		try {
			_localctx = new StructAttrContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(461);
			var_type();
			setState(462);
			match(ID);
			setState(464);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COLON) {
				{
				setState(463);
				match(COLON);
				}
			}

			setState(467);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (((((_la - 10)) & ~0x3f) == 0 && ((1L << (_la - 10)) & 146929947229815793L) != 0)) {
				{
				setState(466);
				expresion(0);
				}
			}

			setState(470);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMMA) {
				{
				setState(469);
				match(COMMA);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Var_typeContext extends ParserRuleContext {
		public Var_typeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_var_type; }
	 
		public Var_typeContext() { }
		public void copyFrom(Var_typeContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VoidTypeContext extends Var_typeContext {
		public VoidTypeContext(Var_typeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterVoidType(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitVoidType(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BoolTypeContext extends Var_typeContext {
		public BoolTypeContext(Var_typeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterBoolType(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitBoolType(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StringTypeContext extends Var_typeContext {
		public StringTypeContext(Var_typeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterStringType(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitStringType(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CharTypeContext extends Var_typeContext {
		public CharTypeContext(Var_typeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterCharType(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitCharType(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IntTypeContext extends Var_typeContext {
		public IntTypeContext(Var_typeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterIntType(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitIntType(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FloatTypeContext extends Var_typeContext {
		public FloatTypeContext(Var_typeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterFloatType(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitFloatType(this);
		}
	}

	public final Var_typeContext var_type() throws RecognitionException {
		Var_typeContext _localctx = new Var_typeContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_var_type);
		try {
			setState(478);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__3:
				_localctx = new IntTypeContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(472);
				match(T__3);
				}
				break;
			case T__4:
				_localctx = new FloatTypeContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(473);
				match(T__4);
				}
				break;
			case T__5:
				_localctx = new StringTypeContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(474);
				match(T__5);
				}
				break;
			case T__6:
				_localctx = new BoolTypeContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(475);
				match(T__6);
				}
				break;
			case T__7:
				_localctx = new CharTypeContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(476);
				match(T__7);
				}
				break;
			case T__8:
				_localctx = new VoidTypeContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(477);
				match(T__8);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Struct_vectorContext extends ParserRuleContext {
		public Struct_vectorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_struct_vector; }
	 
		public Struct_vectorContext() { }
		public void copyFrom(Struct_vectorContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructVectorContext extends Struct_vectorContext {
		public TerminalNode LBRACK() { return getToken(VlangParser.LBRACK, 0); }
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode RBRACK() { return getToken(VlangParser.RBRACK, 0); }
		public TerminalNode LPAREN() { return getToken(VlangParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(VlangParser.RPAREN, 0); }
		public StructVectorContext(Struct_vectorContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterStructVector(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitStructVector(this);
		}
	}

	public final Struct_vectorContext struct_vector() throws RecognitionException {
		Struct_vectorContext _localctx = new Struct_vectorContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_struct_vector);
		try {
			_localctx = new StructVectorContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(480);
			match(LBRACK);
			setState(481);
			match(ID);
			setState(482);
			match(RBRACK);
			setState(483);
			match(LPAREN);
			setState(484);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class IncredecreContext extends ParserRuleContext {
		public IncredecreContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_incredecre; }
	 
		public IncredecreContext() { }
		public void copyFrom(IncredecreContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Incremento2Context extends IncredecreContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode INC2() { return getToken(VlangParser.INC2, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public Incremento2Context(IncredecreContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterIncremento2(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitIncremento2(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Decremento2Context extends IncredecreContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode DEC2() { return getToken(VlangParser.DEC2, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public Decremento2Context(IncredecreContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterDecremento2(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitDecremento2(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IncrementoContext extends IncredecreContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode INC() { return getToken(VlangParser.INC, 0); }
		public IncrementoContext(IncredecreContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterIncremento(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitIncremento(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DecrementoContext extends IncredecreContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode DEC() { return getToken(VlangParser.DEC, 0); }
		public DecrementoContext(IncredecreContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterDecremento(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitDecremento(this);
		}
	}

	public final IncredecreContext incredecre() throws RecognitionException {
		IncredecreContext _localctx = new IncredecreContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_incredecre);
		try {
			setState(496);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,39,_ctx) ) {
			case 1:
				_localctx = new IncrementoContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(486);
				match(ID);
				setState(487);
				match(INC);
				}
				break;
			case 2:
				_localctx = new DecrementoContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(488);
				match(ID);
				setState(489);
				match(DEC);
				}
				break;
			case 3:
				_localctx = new Incremento2Context(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(490);
				match(ID);
				setState(491);
				match(INC2);
				setState(492);
				expresion(0);
				}
				break;
			case 4:
				_localctx = new Decremento2Context(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(493);
				match(ID);
				setState(494);
				match(DEC2);
				setState(495);
				expresion(0);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 23:
			return expresion_sempred((ExpresionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expresion_sempred(ExpresionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 18);
		case 1:
			return precpred(_ctx, 17);
		case 2:
			return precpred(_ctx, 16);
		case 3:
			return precpred(_ctx, 15);
		case 4:
			return precpred(_ctx, 14);
		case 5:
			return precpred(_ctx, 13);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001F\u01f3\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0002\u001b\u0007\u001b"+
		"\u0002\u001c\u0007\u001c\u0002\u001d\u0007\u001d\u0002\u001e\u0007\u001e"+
		"\u0002\u001f\u0007\u001f\u0002 \u0007 \u0002!\u0007!\u0002\"\u0007\"\u0001"+
		"\u0000\u0004\u0000H\b\u0000\u000b\u0000\f\u0000I\u0001\u0000\u0001\u0000"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0003\u0001^\b\u0001"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0005\u0002c\b\u0002\n\u0002\f\u0002"+
		"f\t\u0002\u0001\u0002\u0003\u0002i\b\u0002\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0005\u0003o\b\u0003\n\u0003\f\u0003r\t\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0005\u0004y\b"+
		"\u0004\n\u0004\f\u0004|\t\u0004\u0001\u0004\u0001\u0004\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0005\u0005\u0085\b\u0005\n"+
		"\u0005\f\u0005\u0088\t\u0005\u0001\u0005\u0001\u0005\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0005\u0006\u0091\b\u0006\n"+
		"\u0006\f\u0006\u0094\t\u0006\u0001\u0006\u0001\u0006\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0005\u0007\u009c\b\u0007\n\u0007\f\u0007"+
		"\u009f\t\u0007\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0001\b\u0001\b"+
		"\u0001\b\u0001\b\u0001\b\u0001\b\u0005\b\u00ab\b\b\n\b\f\b\u00ae\t\b\u0001"+
		"\b\u0001\b\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0003\t\u00b7\b\t\u0001"+
		"\t\u0001\t\u0005\t\u00bb\b\t\n\t\f\t\u00be\t\t\u0001\t\u0001\t\u0001\n"+
		"\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\u000b\u0001\u000b\u0001"+
		"\u000b\u0001\u000b\u0004\u000b\u00cc\b\u000b\u000b\u000b\f\u000b\u00cd"+
		"\u0001\u000b\u0003\u000b\u00d1\b\u000b\u0001\u000b\u0001\u000b\u0001\f"+
		"\u0001\f\u0001\f\u0001\f\u0005\f\u00d9\b\f\n\f\f\f\u00dc\t\f\u0001\f\u0001"+
		"\f\u0005\f\u00e0\b\f\n\f\f\f\u00e3\t\f\u0001\r\u0001\r\u0001\r\u0005\r"+
		"\u00e8\b\r\n\r\f\r\u00eb\t\r\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0005"+
		"\u000e\u00f6\b\u000e\n\u000e\f\u000e\u00f9\t\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000f\u0001\u000f\u0001\u000f\u0003\u000f\u0100\b\u000f\u0001\u000f"+
		"\u0001\u000f\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0003\u0010"+
		"\u0108\b\u0010\u0001\u0010\u0001\u0010\u0003\u0010\u010c\b\u0010\u0001"+
		"\u0010\u0001\u0010\u0005\u0010\u0110\b\u0010\n\u0010\f\u0010\u0113\t\u0010"+
		"\u0001\u0010\u0001\u0010\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011"+
		"\u0005\u0011\u011b\b\u0011\n\u0011\f\u0011\u011e\t\u0011\u0001\u0011\u0001"+
		"\u0011\u0001\u0012\u0001\u0012\u0003\u0012\u0124\b\u0012\u0001\u0012\u0001"+
		"\u0012\u0003\u0012\u0128\b\u0012\u0001\u0013\u0001\u0013\u0001\u0013\u0001"+
		"\u0013\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001"+
		"\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0003\u0014\u0138"+
		"\b\u0014\u0001\u0015\u0001\u0015\u0001\u0015\u0005\u0015\u013d\b\u0015"+
		"\n\u0015\f\u0015\u0140\t\u0015\u0001\u0016\u0001\u0016\u0001\u0016\u0005"+
		"\u0016\u0145\b\u0016\n\u0016\f\u0016\u0148\t\u0016\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0003\u0017\u018b\b\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0005\u0017"+
		"\u019f\b\u0017\n\u0017\f\u0017\u01a2\t\u0017\u0001\u0018\u0001\u0018\u0001"+
		"\u0018\u0005\u0018\u01a7\b\u0018\n\u0018\f\u0018\u01aa\t\u0018\u0001\u0019"+
		"\u0001\u0019\u0001\u001a\u0001\u001a\u0001\u001a\u0005\u001a\u01b1\b\u001a"+
		"\n\u001a\f\u001a\u01b4\t\u001a\u0001\u001b\u0001\u001b\u0001\u001b\u0001"+
		"\u001c\u0001\u001c\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0001"+
		"\u001d\u0001\u001d\u0003\u001d\u01c1\b\u001d\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0005\u001e\u01c7\b\u001e\n\u001e\f\u001e\u01ca\t\u001e"+
		"\u0001\u001e\u0001\u001e\u0001\u001f\u0001\u001f\u0001\u001f\u0003\u001f"+
		"\u01d1\b\u001f\u0001\u001f\u0003\u001f\u01d4\b\u001f\u0001\u001f\u0003"+
		"\u001f\u01d7\b\u001f\u0001 \u0001 \u0001 \u0001 \u0001 \u0001 \u0003 "+
		"\u01df\b \u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001\"\u0001\"\u0001"+
		"\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0001\"\u0003\"\u01f1"+
		"\b\"\u0001\"\u0000\u0001.#\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012"+
		"\u0014\u0016\u0018\u001a\u001c\u001e \"$&(*,.02468:<>@BD\u0000\u0005\u0002"+
		"\u0000\'\'++\u0001\u0000(*\u0001\u0000&\'\u0001\u000003\u0001\u0000./"+
		"\u0223\u0000G\u0001\u0000\u0000\u0000\u0002]\u0001\u0000\u0000\u0000\u0004"+
		"_\u0001\u0000\u0000\u0000\u0006j\u0001\u0000\u0000\u0000\bu\u0001\u0000"+
		"\u0000\u0000\n\u007f\u0001\u0000\u0000\u0000\f\u008b\u0001\u0000\u0000"+
		"\u0000\u000e\u0097\u0001\u0000\u0000\u0000\u0010\u00a2\u0001\u0000\u0000"+
		"\u0000\u0012\u00b1\u0001\u0000\u0000\u0000\u0014\u00c1\u0001\u0000\u0000"+
		"\u0000\u0016\u00c7\u0001\u0000\u0000\u0000\u0018\u00d4\u0001\u0000\u0000"+
		"\u0000\u001a\u00e4\u0001\u0000\u0000\u0000\u001c\u00ec\u0001\u0000\u0000"+
		"\u0000\u001e\u00fc\u0001\u0000\u0000\u0000 \u0103\u0001\u0000\u0000\u0000"+
		"\"\u0116\u0001\u0000\u0000\u0000$\u0127\u0001\u0000\u0000\u0000&\u0129"+
		"\u0001\u0000\u0000\u0000(\u0137\u0001\u0000\u0000\u0000*\u0139\u0001\u0000"+
		"\u0000\u0000,\u0141\u0001\u0000\u0000\u0000.\u018a\u0001\u0000\u0000\u0000"+
		"0\u01a3\u0001\u0000\u0000\u00002\u01ab\u0001\u0000\u0000\u00004\u01ad"+
		"\u0001\u0000\u0000\u00006\u01b5\u0001\u0000\u0000\u00008\u01b8\u0001\u0000"+
		"\u0000\u0000:\u01c0\u0001\u0000\u0000\u0000<\u01c2\u0001\u0000\u0000\u0000"+
		">\u01cd\u0001\u0000\u0000\u0000@\u01de\u0001\u0000\u0000\u0000B\u01e0"+
		"\u0001\u0000\u0000\u0000D\u01f0\u0001\u0000\u0000\u0000FH\u0003\u0002"+
		"\u0001\u0000GF\u0001\u0000\u0000\u0000HI\u0001\u0000\u0000\u0000IG\u0001"+
		"\u0000\u0000\u0000IJ\u0001\u0000\u0000\u0000JK\u0001\u0000\u0000\u0000"+
		"KL\u0005\u0000\u0000\u0001L\u0001\u0001\u0000\u0000\u0000M^\u0003(\u0014"+
		"\u0000N^\u0003&\u0013\u0000O^\u0003\u0004\u0002\u0000P^\u0003\u0012\t"+
		"\u0000Q^\u0003\u001e\u000f\u0000R^\u0003 \u0010\u0000S^\u0003\"\u0011"+
		"\u0000T^\u0003<\u001e\u0000U^\u0003$\u0012\u0000V^\u0003\n\u0005\u0000"+
		"W^\u0003\f\u0006\u0000X^\u0003D\"\u0000Y^\u0003\u000e\u0007\u0000Z^\u0003"+
		"\u0010\b\u0000[^\u0003\u0016\u000b\u0000\\^\u0003\u001c\u000e\u0000]M"+
		"\u0001\u0000\u0000\u0000]N\u0001\u0000\u0000\u0000]O\u0001\u0000\u0000"+
		"\u0000]P\u0001\u0000\u0000\u0000]Q\u0001\u0000\u0000\u0000]R\u0001\u0000"+
		"\u0000\u0000]S\u0001\u0000\u0000\u0000]T\u0001\u0000\u0000\u0000]U\u0001"+
		"\u0000\u0000\u0000]V\u0001\u0000\u0000\u0000]W\u0001\u0000\u0000\u0000"+
		"]X\u0001\u0000\u0000\u0000]Y\u0001\u0000\u0000\u0000]Z\u0001\u0000\u0000"+
		"\u0000][\u0001\u0000\u0000\u0000]\\\u0001\u0000\u0000\u0000^\u0003\u0001"+
		"\u0000\u0000\u0000_d\u0003\u0006\u0003\u0000`a\u0005\u0015\u0000\u0000"+
		"ac\u0003\u0006\u0003\u0000b`\u0001\u0000\u0000\u0000cf\u0001\u0000\u0000"+
		"\u0000db\u0001\u0000\u0000\u0000de\u0001\u0000\u0000\u0000eh\u0001\u0000"+
		"\u0000\u0000fd\u0001\u0000\u0000\u0000gi\u0003\b\u0004\u0000hg\u0001\u0000"+
		"\u0000\u0000hi\u0001\u0000\u0000\u0000i\u0005\u0001\u0000\u0000\u0000"+
		"jk\u0005\u0014\u0000\u0000kl\u0003.\u0017\u0000lp\u0005>\u0000\u0000m"+
		"o\u0003\u0002\u0001\u0000nm\u0001\u0000\u0000\u0000or\u0001\u0000\u0000"+
		"\u0000pn\u0001\u0000\u0000\u0000pq\u0001\u0000\u0000\u0000qs\u0001\u0000"+
		"\u0000\u0000rp\u0001\u0000\u0000\u0000st\u0005@\u0000\u0000t\u0007\u0001"+
		"\u0000\u0000\u0000uv\u0005\u0015\u0000\u0000vz\u0005>\u0000\u0000wy\u0003"+
		"\u0002\u0001\u0000xw\u0001\u0000\u0000\u0000y|\u0001\u0000\u0000\u0000"+
		"zx\u0001\u0000\u0000\u0000z{\u0001\u0000\u0000\u0000{}\u0001\u0000\u0000"+
		"\u0000|z\u0001\u0000\u0000\u0000}~\u0005@\u0000\u0000~\t\u0001\u0000\u0000"+
		"\u0000\u007f\u0080\u0005\u0001\u0000\u0000\u0080\u0081\u0005;\u0000\u0000"+
		"\u0081\u0086\u0003.\u0017\u0000\u0082\u0083\u0005B\u0000\u0000\u0083\u0085"+
		"\u0003.\u0017\u0000\u0084\u0082\u0001\u0000\u0000\u0000\u0085\u0088\u0001"+
		"\u0000\u0000\u0000\u0086\u0084\u0001\u0000\u0000\u0000\u0086\u0087\u0001"+
		"\u0000\u0000\u0000\u0087\u0089\u0001\u0000\u0000\u0000\u0088\u0086\u0001"+
		"\u0000\u0000\u0000\u0089\u008a\u0005<\u0000\u0000\u008a\u000b\u0001\u0000"+
		"\u0000\u0000\u008b\u008c\u0005\u0002\u0000\u0000\u008c\u008d\u0005;\u0000"+
		"\u0000\u008d\u0092\u0003.\u0017\u0000\u008e\u008f\u0005B\u0000\u0000\u008f"+
		"\u0091\u0003.\u0017\u0000\u0090\u008e\u0001\u0000\u0000\u0000\u0091\u0094"+
		"\u0001\u0000\u0000\u0000\u0092\u0090\u0001\u0000\u0000\u0000\u0092\u0093"+
		"\u0001\u0000\u0000\u0000\u0093\u0095\u0001\u0000\u0000\u0000\u0094\u0092"+
		"\u0001\u0000\u0000\u0000\u0095\u0096\u0005<\u0000\u0000\u0096\r\u0001"+
		"\u0000\u0000\u0000\u0097\u0098\u0005\u0017\u0000\u0000\u0098\u0099\u0003"+
		".\u0017\u0000\u0099\u009d\u0005>\u0000\u0000\u009a\u009c\u0003\u0002\u0001"+
		"\u0000\u009b\u009a\u0001\u0000\u0000\u0000\u009c\u009f\u0001\u0000\u0000"+
		"\u0000\u009d\u009b\u0001\u0000\u0000\u0000\u009d\u009e\u0001\u0000\u0000"+
		"\u0000\u009e\u00a0\u0001\u0000\u0000\u0000\u009f\u009d\u0001\u0000\u0000"+
		"\u0000\u00a0\u00a1\u0005@\u0000\u0000\u00a1\u000f\u0001\u0000\u0000\u0000"+
		"\u00a2\u00a3\u0005\u0017\u0000\u0000\u00a3\u00a4\u0003&\u0013\u0000\u00a4"+
		"\u00a5\u0005:\u0000\u0000\u00a5\u00a6\u0003.\u0017\u0000\u00a6\u00a7\u0005"+
		":\u0000\u0000\u00a7\u00a8\u0003D\"\u0000\u00a8\u00ac\u0005>\u0000\u0000"+
		"\u00a9\u00ab\u0003\u0002\u0001\u0000\u00aa\u00a9\u0001\u0000\u0000\u0000"+
		"\u00ab\u00ae\u0001\u0000\u0000\u0000\u00ac\u00aa\u0001\u0000\u0000\u0000"+
		"\u00ac\u00ad\u0001\u0000\u0000\u0000\u00ad\u00af\u0001\u0000\u0000\u0000"+
		"\u00ae\u00ac\u0001\u0000\u0000\u0000\u00af\u00b0\u0005@\u0000\u0000\u00b0"+
		"\u0011\u0001\u0000\u0000\u0000\u00b1\u00b2\u0005\u0017\u0000\u0000\u00b2"+
		"\u00b3\u0005C\u0000\u0000\u00b3\u00b6\u0005\u0018\u0000\u0000\u00b4\u00b7"+
		"\u0003.\u0017\u0000\u00b5\u00b7\u0003\u0014\n\u0000\u00b6\u00b4\u0001"+
		"\u0000\u0000\u0000\u00b6\u00b5\u0001\u0000\u0000\u0000\u00b7\u00b8\u0001"+
		"\u0000\u0000\u0000\u00b8\u00bc\u0005>\u0000\u0000\u00b9\u00bb\u0003\u0002"+
		"\u0001\u0000\u00ba\u00b9\u0001\u0000\u0000\u0000\u00bb\u00be\u0001\u0000"+
		"\u0000\u0000\u00bc\u00ba\u0001\u0000\u0000\u0000\u00bc\u00bd\u0001\u0000"+
		"\u0000\u0000\u00bd\u00bf\u0001\u0000\u0000\u0000\u00be\u00bc\u0001\u0000"+
		"\u0000\u0000\u00bf\u00c0\u0005@\u0000\u0000\u00c0\u0013\u0001\u0000\u0000"+
		"\u0000\u00c1\u00c2\u0003.\u0017\u0000\u00c2\u00c3\u0005A\u0000\u0000\u00c3"+
		"\u00c4\u0005A\u0000\u0000\u00c4\u00c5\u0005A\u0000\u0000\u00c5\u00c6\u0003"+
		".\u0017\u0000\u00c6\u0015\u0001\u0000\u0000\u0000\u00c7\u00c8\u0005\u001a"+
		"\u0000\u0000\u00c8\u00c9\u0005C\u0000\u0000\u00c9\u00cb\u0005>\u0000\u0000"+
		"\u00ca\u00cc\u0003\u0018\f\u0000\u00cb\u00ca\u0001\u0000\u0000\u0000\u00cc"+
		"\u00cd\u0001\u0000\u0000\u0000\u00cd\u00cb\u0001\u0000\u0000\u0000\u00cd"+
		"\u00ce\u0001\u0000\u0000\u0000\u00ce\u00d0\u0001\u0000\u0000\u0000\u00cf"+
		"\u00d1\u0003\u001a\r\u0000\u00d0\u00cf\u0001\u0000\u0000\u0000\u00d0\u00d1"+
		"\u0001\u0000\u0000\u0000\u00d1\u00d2\u0001\u0000\u0000\u0000\u00d2\u00d3"+
		"\u0005@\u0000\u0000\u00d3\u0017\u0001\u0000\u0000\u0000\u00d4\u00d5\u0005"+
		"\u001b\u0000\u0000\u00d5\u00da\u0003.\u0017\u0000\u00d6\u00d7\u0005B\u0000"+
		"\u0000\u00d7\u00d9\u0003.\u0017\u0000\u00d8\u00d6\u0001\u0000\u0000\u0000"+
		"\u00d9\u00dc\u0001\u0000\u0000\u0000\u00da\u00d8\u0001\u0000\u0000\u0000"+
		"\u00da\u00db\u0001\u0000\u0000\u0000\u00db\u00dd\u0001\u0000\u0000\u0000"+
		"\u00dc\u00da\u0001\u0000\u0000\u0000\u00dd\u00e1\u00059\u0000\u0000\u00de"+
		"\u00e0\u0003\u0002\u0001\u0000\u00df\u00de\u0001\u0000\u0000\u0000\u00e0"+
		"\u00e3\u0001\u0000\u0000\u0000\u00e1\u00df\u0001\u0000\u0000\u0000\u00e1"+
		"\u00e2\u0001\u0000\u0000\u0000\u00e2\u0019\u0001\u0000\u0000\u0000\u00e3"+
		"\u00e1\u0001\u0000\u0000\u0000\u00e4\u00e5\u0005\u001c\u0000\u0000\u00e5"+
		"\u00e9\u00059\u0000\u0000\u00e6\u00e8\u0003\u0002\u0001\u0000\u00e7\u00e6"+
		"\u0001\u0000\u0000\u0000\u00e8\u00eb\u0001\u0000\u0000\u0000\u00e9\u00e7"+
		"\u0001\u0000\u0000\u0000\u00e9\u00ea\u0001\u0000\u0000\u0000\u00ea\u001b"+
		"\u0001\u0000\u0000\u0000\u00eb\u00e9\u0001\u0000\u0000\u0000\u00ec\u00ed"+
		"\u0005C\u0000\u0000\u00ed\u00ee\u00054\u0000\u0000\u00ee\u00ef\u0005="+
		"\u0000\u0000\u00ef\u00f0\u0005?\u0000\u0000\u00f0\u00f1\u0003@ \u0000"+
		"\u00f1\u00f2\u0005>\u0000\u0000\u00f2\u00f7\u0003.\u0017\u0000\u00f3\u00f4"+
		"\u0005B\u0000\u0000\u00f4\u00f6\u0003.\u0017\u0000\u00f5\u00f3\u0001\u0000"+
		"\u0000\u0000\u00f6\u00f9\u0001\u0000\u0000\u0000\u00f7\u00f5\u0001\u0000"+
		"\u0000\u0000\u00f7\u00f8\u0001\u0000\u0000\u0000\u00f8\u00fa\u0001\u0000"+
		"\u0000\u0000\u00f9\u00f7\u0001\u0000\u0000\u0000\u00fa\u00fb\u0005@\u0000"+
		"\u0000\u00fb\u001d\u0001\u0000\u0000\u0000\u00fc\u00fd\u0003,\u0016\u0000"+
		"\u00fd\u00ff\u0005;\u0000\u0000\u00fe\u0100\u00030\u0018\u0000\u00ff\u00fe"+
		"\u0001\u0000\u0000\u0000\u00ff\u0100\u0001\u0000\u0000\u0000\u0100\u0101"+
		"\u0001\u0000\u0000\u0000\u0101\u0102\u0005<\u0000\u0000\u0102\u001f\u0001"+
		"\u0000\u0000\u0000\u0103\u0104\u0005\u001f\u0000\u0000\u0104\u0105\u0005"+
		"C\u0000\u0000\u0105\u0107\u0005;\u0000\u0000\u0106\u0108\u00034\u001a"+
		"\u0000\u0107\u0106\u0001\u0000\u0000\u0000\u0107\u0108\u0001\u0000\u0000"+
		"\u0000\u0108\u0109\u0001\u0000\u0000\u0000\u0109\u010b\u0005<\u0000\u0000"+
		"\u010a\u010c\u0003@ \u0000\u010b\u010a\u0001\u0000\u0000\u0000\u010b\u010c"+
		"\u0001\u0000\u0000\u0000\u010c\u010d\u0001\u0000\u0000\u0000\u010d\u0111"+
		"\u0005>\u0000\u0000\u010e\u0110\u0003\u0002\u0001\u0000\u010f\u010e\u0001"+
		"\u0000\u0000\u0000\u0110\u0113\u0001\u0000\u0000\u0000\u0111\u010f\u0001"+
		"\u0000\u0000\u0000\u0111\u0112\u0001\u0000\u0000\u0000\u0112\u0114\u0001"+
		"\u0000\u0000\u0000\u0113\u0111\u0001\u0000\u0000\u0000\u0114\u0115\u0005"+
		"@\u0000\u0000\u0115!\u0001\u0000\u0000\u0000\u0116\u0117\u0005\u0003\u0000"+
		"\u0000\u0117\u0118\u0005C\u0000\u0000\u0118\u011c\u0005>\u0000\u0000\u0119"+
		"\u011b\u0003&\u0013\u0000\u011a\u0119\u0001\u0000\u0000\u0000\u011b\u011e"+
		"\u0001\u0000\u0000\u0000\u011c\u011a\u0001\u0000\u0000\u0000\u011c\u011d"+
		"\u0001\u0000\u0000\u0000\u011d\u011f\u0001\u0000\u0000\u0000\u011e\u011c"+
		"\u0001\u0000\u0000\u0000\u011f\u0120\u0005@\u0000\u0000\u0120#\u0001\u0000"+
		"\u0000\u0000\u0121\u0123\u0005\u0019\u0000\u0000\u0122\u0124\u0003.\u0017"+
		"\u0000\u0123\u0122\u0001\u0000\u0000\u0000\u0123\u0124\u0001\u0000\u0000"+
		"\u0000\u0124\u0128\u0001\u0000\u0000\u0000\u0125\u0128\u0005\u001d\u0000"+
		"\u0000\u0126\u0128\u0005\u001e\u0000\u0000\u0127\u0121\u0001\u0000\u0000"+
		"\u0000\u0127\u0125\u0001\u0000\u0000\u0000\u0127\u0126\u0001\u0000\u0000"+
		"\u0000\u0128%\u0001\u0000\u0000\u0000\u0129\u012a\u0003,\u0016\u0000\u012a"+
		"\u012b\u00054\u0000\u0000\u012b\u012c\u0003.\u0017\u0000\u012c\'\u0001"+
		"\u0000\u0000\u0000\u012d\u012e\u0005\r\u0000\u0000\u012e\u012f\u0005C"+
		"\u0000\u0000\u012f\u0130\u00054\u0000\u0000\u0130\u0138\u0003.\u0017\u0000"+
		"\u0131\u0132\u0005\r\u0000\u0000\u0132\u0133\u0003,\u0016\u0000\u0133"+
		"\u0134\u0003@ \u0000\u0134\u0135\u00054\u0000\u0000\u0135\u0136\u0003"+
		".\u0017\u0000\u0136\u0138\u0001\u0000\u0000\u0000\u0137\u012d\u0001\u0000"+
		"\u0000\u0000\u0137\u0131\u0001\u0000\u0000\u0000\u0138)\u0001\u0000\u0000"+
		"\u0000\u0139\u013e\u0003.\u0017\u0000\u013a\u013b\u0005B\u0000\u0000\u013b"+
		"\u013d\u0003.\u0017\u0000\u013c\u013a\u0001\u0000\u0000\u0000\u013d\u0140"+
		"\u0001\u0000\u0000\u0000\u013e\u013c\u0001\u0000\u0000\u0000\u013e\u013f"+
		"\u0001\u0000\u0000\u0000\u013f+\u0001\u0000\u0000\u0000\u0140\u013e\u0001"+
		"\u0000\u0000\u0000\u0141\u0146\u0005C\u0000\u0000\u0142\u0143\u0005A\u0000"+
		"\u0000\u0143\u0145\u0005C\u0000\u0000\u0144\u0142\u0001\u0000\u0000\u0000"+
		"\u0145\u0148\u0001\u0000\u0000\u0000\u0146\u0144\u0001\u0000\u0000\u0000"+
		"\u0146\u0147\u0001\u0000\u0000\u0000\u0147-\u0001\u0000\u0000\u0000\u0148"+
		"\u0146\u0001\u0000\u0000\u0000\u0149\u014a\u0006\u0017\uffff\uffff\u0000"+
		"\u014a\u018b\u0003:\u001d\u0000\u014b\u018b\u0003\u001e\u000f\u0000\u014c"+
		"\u014d\u0005;\u0000\u0000\u014d\u014e\u0003.\u0017\u0000\u014e\u014f\u0005"+
		"<\u0000\u0000\u014f\u018b\u0001\u0000\u0000\u0000\u0150\u0151\u0005=\u0000"+
		"\u0000\u0151\u0152\u0003.\u0017\u0000\u0152\u0153\u0005?\u0000\u0000\u0153"+
		"\u018b\u0001\u0000\u0000\u0000\u0154\u0155\u0007\u0000\u0000\u0000\u0155"+
		"\u018b\u0003.\u0017\u0013\u0156\u018b\u0005C\u0000\u0000\u0157\u018b\u0003"+
		"D\"\u0000\u0158\u0159\u0005C\u0000\u0000\u0159\u015a\u0005A\u0000\u0000"+
		"\u015a\u018b\u0005C\u0000\u0000\u015b\u015c\u0005C\u0000\u0000\u015c\u015d"+
		"\u0005A\u0000\u0000\u015d\u018b\u0003.\u0017\t\u015e\u015f\u0005C\u0000"+
		"\u0000\u015f\u0160\u00054\u0000\u0000\u0160\u018b\u0003.\u0017\b\u0161"+
		"\u0162\u0003,\u0016\u0000\u0162\u0163\u0005=\u0000\u0000\u0163\u0164\u0003"+
		".\u0017\u0000\u0164\u0165\u0005?\u0000\u0000\u0165\u018b\u0001\u0000\u0000"+
		"\u0000\u0166\u0167\u0005!\u0000\u0000\u0167\u0168\u0005;\u0000\u0000\u0168"+
		"\u0169\u0003.\u0017\u0000\u0169\u016a\u0005<\u0000\u0000\u016a\u018b\u0001"+
		"\u0000\u0000\u0000\u016b\u016c\u0005\"\u0000\u0000\u016c\u016d\u0005;"+
		"\u0000\u0000\u016d\u016e\u0003.\u0017\u0000\u016e\u016f\u0005<\u0000\u0000"+
		"\u016f\u018b\u0001\u0000\u0000\u0000\u0170\u0171\u0005#\u0000\u0000\u0171"+
		"\u0172\u0005;\u0000\u0000\u0172\u0173\u0003.\u0017\u0000\u0173\u0174\u0005"+
		"<\u0000\u0000\u0174\u018b\u0001\u0000\u0000\u0000\u0175\u0176\u0005$\u0000"+
		"\u0000\u0176\u0177\u0005;\u0000\u0000\u0177\u0178\u0003.\u0017\u0000\u0178"+
		"\u0179\u0005B\u0000\u0000\u0179\u017a\u0003.\u0017\u0000\u017a\u017b\u0005"+
		"<\u0000\u0000\u017b\u018b\u0001\u0000\u0000\u0000\u017c\u017d\u0005\n"+
		"\u0000\u0000\u017d\u017e\u0005;\u0000\u0000\u017e\u017f\u0003.\u0017\u0000"+
		"\u017f\u0180\u0005B\u0000\u0000\u0180\u0181\u0003.\u0017\u0000\u0181\u0182"+
		"\u0005<\u0000\u0000\u0182\u018b\u0001\u0000\u0000\u0000\u0183\u0184\u0005"+
		"%\u0000\u0000\u0184\u0185\u0005;\u0000\u0000\u0185\u0186\u0003.\u0017"+
		"\u0000\u0186\u0187\u0005B\u0000\u0000\u0187\u0188\u0003.\u0017\u0000\u0188"+
		"\u0189\u0005<\u0000\u0000\u0189\u018b\u0001\u0000\u0000\u0000\u018a\u0149"+
		"\u0001\u0000\u0000\u0000\u018a\u014b\u0001\u0000\u0000\u0000\u018a\u014c"+
		"\u0001\u0000\u0000\u0000\u018a\u0150\u0001\u0000\u0000\u0000\u018a\u0154"+
		"\u0001\u0000\u0000\u0000\u018a\u0156\u0001\u0000\u0000\u0000\u018a\u0157"+
		"\u0001\u0000\u0000\u0000\u018a\u0158\u0001\u0000\u0000\u0000\u018a\u015b"+
		"\u0001\u0000\u0000\u0000\u018a\u015e\u0001\u0000\u0000\u0000\u018a\u0161"+
		"\u0001\u0000\u0000\u0000\u018a\u0166\u0001\u0000\u0000\u0000\u018a\u016b"+
		"\u0001\u0000\u0000\u0000\u018a\u0170\u0001\u0000\u0000\u0000\u018a\u0175"+
		"\u0001\u0000\u0000\u0000\u018a\u017c\u0001\u0000\u0000\u0000\u018a\u0183"+
		"\u0001\u0000\u0000\u0000\u018b\u01a0\u0001\u0000\u0000\u0000\u018c\u018d"+
		"\n\u0012\u0000\u0000\u018d\u018e\u0007\u0001\u0000\u0000\u018e\u019f\u0003"+
		".\u0017\u0013\u018f\u0190\n\u0011\u0000\u0000\u0190\u0191\u0007\u0002"+
		"\u0000\u0000\u0191\u019f\u0003.\u0017\u0012\u0192\u0193\n\u0010\u0000"+
		"\u0000\u0193\u0194\u0007\u0003\u0000\u0000\u0194\u019f\u0003.\u0017\u0011"+
		"\u0195\u0196\n\u000f\u0000\u0000\u0196\u0197\u0007\u0004\u0000\u0000\u0197"+
		"\u019f\u0003.\u0017\u0010\u0198\u0199\n\u000e\u0000\u0000\u0199\u019a"+
		"\u0005-\u0000\u0000\u019a\u019f\u0003.\u0017\u000f\u019b\u019c\n\r\u0000"+
		"\u0000\u019c\u019d\u0005,\u0000\u0000\u019d\u019f\u0003.\u0017\u000e\u019e"+
		"\u018c\u0001\u0000\u0000\u0000\u019e\u018f\u0001\u0000\u0000\u0000\u019e"+
		"\u0192\u0001\u0000\u0000\u0000\u019e\u0195\u0001\u0000\u0000\u0000\u019e"+
		"\u0198\u0001\u0000\u0000\u0000\u019e\u019b\u0001\u0000\u0000\u0000\u019f"+
		"\u01a2\u0001\u0000\u0000\u0000\u01a0\u019e\u0001\u0000\u0000\u0000\u01a0"+
		"\u01a1\u0001\u0000\u0000\u0000\u01a1/\u0001\u0000\u0000\u0000\u01a2\u01a0"+
		"\u0001\u0000\u0000\u0000\u01a3\u01a8\u00032\u0019\u0000\u01a4\u01a5\u0005"+
		"B\u0000\u0000\u01a5\u01a7\u00032\u0019\u0000\u01a6\u01a4\u0001\u0000\u0000"+
		"\u0000\u01a7\u01aa\u0001\u0000\u0000\u0000\u01a8\u01a6\u0001\u0000\u0000"+
		"\u0000\u01a8\u01a9\u0001\u0000\u0000\u0000\u01a91\u0001\u0000\u0000\u0000"+
		"\u01aa\u01a8\u0001\u0000\u0000\u0000\u01ab\u01ac\u0003,\u0016\u0000\u01ac"+
		"3\u0001\u0000\u0000\u0000\u01ad\u01b2\u00036\u001b\u0000\u01ae\u01af\u0005"+
		"B\u0000\u0000\u01af\u01b1\u00036\u001b\u0000\u01b0\u01ae\u0001\u0000\u0000"+
		"\u0000\u01b1\u01b4\u0001\u0000\u0000\u0000\u01b2\u01b0\u0001\u0000\u0000"+
		"\u0000\u01b2\u01b3\u0001\u0000\u0000\u0000\u01b35\u0001\u0000\u0000\u0000"+
		"\u01b4\u01b2\u0001\u0000\u0000\u0000\u01b5\u01b6\u0005C\u0000\u0000\u01b6"+
		"\u01b7\u0003@ \u0000\u01b77\u0001\u0000\u0000\u0000\u01b8\u01b9\u0003"+
		":\u001d\u0000\u01b99\u0001\u0000\u0000\u0000\u01ba\u01c1\u0005\u000f\u0000"+
		"\u0000\u01bb\u01c1\u0005\u0011\u0000\u0000\u01bc\u01c1\u0005\u0012\u0000"+
		"\u0000\u01bd\u01c1\u0005\u000e\u0000\u0000\u01be\u01c1\u0005\u0013\u0000"+
		"\u0000\u01bf\u01c1\u0005\u0010\u0000\u0000\u01c0\u01ba\u0001\u0000\u0000"+
		"\u0000\u01c0\u01bb\u0001\u0000\u0000\u0000\u01c0\u01bc\u0001\u0000\u0000"+
		"\u0000\u01c0\u01bd\u0001\u0000\u0000\u0000\u01c0\u01be\u0001\u0000\u0000"+
		"\u0000\u01c0\u01bf\u0001\u0000\u0000\u0000\u01c1;\u0001\u0000\u0000\u0000"+
		"\u01c2\u01c3\u0005\u0003\u0000\u0000\u01c3\u01c4\u0005C\u0000\u0000\u01c4"+
		"\u01c8\u0005>\u0000\u0000\u01c5\u01c7\u0003>\u001f\u0000\u01c6\u01c5\u0001"+
		"\u0000\u0000\u0000\u01c7\u01ca\u0001\u0000\u0000\u0000\u01c8\u01c6\u0001"+
		"\u0000\u0000\u0000\u01c8\u01c9\u0001\u0000\u0000\u0000\u01c9\u01cb\u0001"+
		"\u0000\u0000\u0000\u01ca\u01c8\u0001\u0000\u0000\u0000\u01cb\u01cc\u0005"+
		"@\u0000\u0000\u01cc=\u0001\u0000\u0000\u0000\u01cd\u01ce\u0003@ \u0000"+
		"\u01ce\u01d0\u0005C\u0000\u0000\u01cf\u01d1\u00059\u0000\u0000\u01d0\u01cf"+
		"\u0001\u0000\u0000\u0000\u01d0\u01d1\u0001\u0000\u0000\u0000\u01d1\u01d3"+
		"\u0001\u0000\u0000\u0000\u01d2\u01d4\u0003.\u0017\u0000\u01d3\u01d2\u0001"+
		"\u0000\u0000\u0000\u01d3\u01d4\u0001\u0000\u0000\u0000\u01d4\u01d6\u0001"+
		"\u0000\u0000\u0000\u01d5\u01d7\u0005B\u0000\u0000\u01d6\u01d5\u0001\u0000"+
		"\u0000\u0000\u01d6\u01d7\u0001\u0000\u0000\u0000\u01d7?\u0001\u0000\u0000"+
		"\u0000\u01d8\u01df\u0005\u0004\u0000\u0000\u01d9\u01df\u0005\u0005\u0000"+
		"\u0000\u01da\u01df\u0005\u0006\u0000\u0000\u01db\u01df\u0005\u0007\u0000"+
		"\u0000\u01dc\u01df\u0005\b\u0000\u0000\u01dd\u01df\u0005\t\u0000\u0000"+
		"\u01de\u01d8\u0001\u0000\u0000\u0000\u01de\u01d9\u0001\u0000\u0000\u0000"+
		"\u01de\u01da\u0001\u0000\u0000\u0000\u01de\u01db\u0001\u0000\u0000\u0000"+
		"\u01de\u01dc\u0001\u0000\u0000\u0000\u01de\u01dd\u0001\u0000\u0000\u0000"+
		"\u01dfA\u0001\u0000\u0000\u0000\u01e0\u01e1\u0005=\u0000\u0000\u01e1\u01e2"+
		"\u0005C\u0000\u0000\u01e2\u01e3\u0005?\u0000\u0000\u01e3\u01e4\u0005;"+
		"\u0000\u0000\u01e4\u01e5\u0005<\u0000\u0000\u01e5C\u0001\u0000\u0000\u0000"+
		"\u01e6\u01e7\u0005C\u0000\u0000\u01e7\u01f1\u00055\u0000\u0000\u01e8\u01e9"+
		"\u0005C\u0000\u0000\u01e9\u01f1\u00056\u0000\u0000\u01ea\u01eb\u0005C"+
		"\u0000\u0000\u01eb\u01ec\u00057\u0000\u0000\u01ec\u01f1\u0003.\u0017\u0000"+
		"\u01ed\u01ee\u0005C\u0000\u0000\u01ee\u01ef\u00058\u0000\u0000\u01ef\u01f1"+
		"\u0003.\u0017\u0000\u01f0\u01e6\u0001\u0000\u0000\u0000\u01f0\u01e8\u0001"+
		"\u0000\u0000\u0000\u01f0\u01ea\u0001\u0000\u0000\u0000\u01f0\u01ed\u0001"+
		"\u0000\u0000\u0000\u01f1E\u0001\u0000\u0000\u0000(I]dhpz\u0086\u0092\u009d"+
		"\u00ac\u00b6\u00bc\u00cd\u00d0\u00da\u00e1\u00e9\u00f7\u00ff\u0107\u010b"+
		"\u0111\u011c\u0123\u0127\u0137\u013e\u0146\u018a\u019e\u01a0\u01a8\u01b2"+
		"\u01c0\u01c8\u01d0\u01d3\u01d6\u01de\u01f0";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}