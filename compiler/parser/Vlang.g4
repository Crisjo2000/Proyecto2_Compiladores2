grammar Vlang; 


// === Axioma principal ===

programa: (stmt)+ EOF;

/*
Desde el codigo vamos a recorrer Stmt y vamos 
a ver de que tipo es cada uno entonces, no podemos 
darle una etiqueta a cada uno 
 */

stmt:
	 decl_stmt                
	 | assign_stmt
     | if_stmt
    // | while_stmt  
     | for_stmt
     | func_call 
	 | func_dcl
     | struct_dcl
     | strct_instancia
     | transfer_stmt       
     | println 
     | print
     | incredecre
     | forwhile_stmt
     | forClasic_stmt
     | switch_stmt
     | slice_stmt
     ;



if_stmt: if_chain (ELSE_KW if_chain)* else_stmt? # IfStmt;

if_chain: IF_KW expresion  LCOR stmt* RCOR # IfChain;

else_stmt: ELSE_KW LCOR stmt* RCOR # ElseStmt;

println: 'println'  LPAREN expresion (COMMA expresion)* RPAREN # PrintlnStmt;
print: 'print'  LPAREN expresion (COMMA expresion)* RPAREN # PrintStmt;

//while_stmt: WHILE_KW expresion LCOR stmt* RCOR # WhileStmt;

// WHILE
forwhile_stmt: FOR_KW expresion LCOR stmt* RCOR # WhileStmt;

forClasic_stmt: FOR_KW assign_stmt PYC expresion PYC incredecre LCOR stmt* RCOR # clasicFor;

// FOR NORMAL
for_stmt:
	FOR_KW ID IN_KW (expresion | range) LCOR stmt* RCOR # ForStmt;

range: expresion DOT DOT DOT expresion # NumericRange;


// === SWITCH  ===
switch_stmt: 
     SWITCH_KW ID LCOR  (case_clause)+ (default_clause)?RCOR # SwitchStmt;

case_clause:
    CASE_KW (expresion (COMMA expresion)*) COLON 
        stmt* # CaseClause;

default_clause:
    'default' COLON 
        stmt* # DefaultClause;

// === SLICE ===
slice_stmt :ID ASSIGN LBRACK RBRACK var_type LCOR (expresion (COMMA expresion)*) RCOR #DeclSlices ;

//list_expresiones : expresion;


// ---------------------- FUNCIONES------------------------------------------------




/// Llamadas a funciones catch
// id ( 1,2,3,4)
func_call: id_pattern LPAREN (parametros)? RPAREN # FuncCall
    
;


func_dcl: 
    // id(parametros | vacio )  { } 
    //    visitor -> recorrer esos stmts
    // fn id (){}
    // fn id (x int, n string) tipo_dato {}
    // fn id () tipo_dato {}
    FN_KW ID LPAREN (arg_list)? RPAREN (var_type)? LCOR stmt* RCOR # FuncDecl
    ;


/*struct 
<NombreStruct> {
<Tipo> <NombreAtributo>;
...
} */
struct_dcl:
    // struct ID { }
    /**
    TODO: 
    Agregar la posilidad de inicializar
    los atributos sin asignacion
    
     */
    'struct' ID LCOR (assign_stmt)* RCOR # StructDecl
    ;


transfer_stmt:
	RETURN_KW expresion?	# ReturnStmt
	| BREAK_KW		# BreakStmt
	| CONTINUE_KW	# ContinueStmt;


assign_stmt:
	id_pattern ASSIGN expresion  	# DirectAssign

    ; 

decl_stmt: 
    MUT ID ASSIGN expresion # DeclAssign
    | MUT id_pattern var_type ASSIGN expresion # DeclAssignType
//nuevo
   
    ; 

    //nuevo
list_expresiones: expresion (COMMA expresion)* # ListExpressions;

id_pattern: ID (DOT ID)* # IdPattern;

// === Reglas de expresiones ===
expresion
    : valor                                                #valorexpresion        
    | func_call                                            #funcionexpre
    | LPAREN expresion RPAREN                              #parentesisexpre
    | LBRACK expresion RBRACK                              #corchetesexpre
    | op=(NOT | MINUS) expresion                           #unario



	| left = expresion op = (MUL | DIV | MOD) right = expresion	# BinaryExp// a * b, a / b, a % b
	| left = expresion op = (PLUS | MINUS) right = expresion    # BinaryExp // a + b, a - b
	| left = expresion op = ( LT | LE | GT | GE) right = expresion	# BinaryExp // a < b, a <= b, a > b, a >= b
	| left = expresion op = (EQ | NEQ ) right = expresion	# BinaryExp // a == b, a != b
	| left = expresion op = AND right = expresion		    # BinaryExp // a && b
	| left = expresion op = OR right = expresion 			# BinaryExp // a || b

    | ID                                                   #id              
    | incredecre                                           #incredecr      
    | ID DOT ID                                            #expdotexp1             
    | ID DOT expresion                                     #expdotexp      
    | ID ASSIGN expresion                                  #asignacionfor
    | id_pattern LBRACK expresion RBRACK                           #indexAcceso

    // ===== Funciones =====
    |ATOI_KW LPAREN expresion RPAREN                       #fn_atoi
    |PARSEFLOAT_KW LPAREN expresion RPAREN                 #fn_parsefloat
    |TYPEOF_KW LPAREN expresion RPAREN                     #fn_typeOF
    | INDEX_KW LPAREN expresion COMMA expresion RPAREN     #fn_indexOf
    | LEN LPAREN expresion COMMA expresion RPAREN          #fn_lenx
    | JOIN_KW LPAREN expresion COMMA expresion RPAREN      #fn_join

    
    ;

// === Parámetros en llamadas ===
parametros: func_param (COMMA func_param)* # ParamList;
func_param : id_pattern #funcParam;


// ===== Argumentos en las declaraciones === catch
// -----> ArgList
// valor type 
arg_list: func_arg (COMMA func_arg)* # ArgList;
func_arg: ID var_type # FuncArg; // 
              
 
// === Tipos de valores simples ===
valores : valor ;

// === Subcontextos para d ===
valor
    : ENTERO    #valorEntero
    | DECIMAL   #valorDecimal
    | CADENA    #valorCadena
    | BOOLEANO  #valorBooleano
    | CARACTER  #valorCaracter
    | FLOAT     #valorFloat
    ;

// ===== STRUCTS ==== catch
// * Structs

/*  struct id { 
    tipo id : variosTipos;
    ...
} 
*/

// Instancia de Struct
strct_instancia: 'struct' ID LCOR struct_prop* RCOR # StructInstancia;

/*
miInstancia = Persona{
 Nombre: "Alice",
 Edad: 25,
 EsEstudiante: false,
}
 */

struct_prop:
//nuevo:
var_type ID (COLON)? (expresion)? (COMMA)?	# StructAttr

	//var_type ID ASSIGN ID LCOR (ID COLON expresion COMMA) RCOR 	# StructAttr
;
// tipos de primitivos que puede ser un struct catch
var_type
    : 'int'     # IntType
    | 'float'   # FloatType
    | 'string'  # StringType
    | 'bool'    # BoolType
    | 'char'    # CharType
    | 'void'    # VoidType
    ;

struct_vector: LBRACK ID RBRACK LPAREN RPAREN # StructVector;

// === Incremento / Decremento ===
incredecre
    : ID INC    #incremento
    | ID DEC    #decremento
    | ID INC2 expresion #incremento2
    | ID DEC2 expresion #decremento2
    ;

// === Tokens de palabras clave ===
LEN     : 'len' ;
CAP     : 'cap' ;
APPEND  : 'append' ;
MUT     : 'mut' ;
// === Literales ===
BOOLEANO : 'true' | 'false' ;
ENTERO   : [0-9]+ ;
FLOAT    : [0-9]+ '.' [0-9]* | [0-9]* '.' [0-9]+ ;
DECIMAL  : [0-9]+ '.' [0-9]+ ;
CADENA   : '"' (~["\\] | '\\' .)* '"' ;
CARACTER : '\'' . '\'' ;




IF_KW : 'if' ;
ELSE_KW : 'else' ;
WHILE_KW : 'while' ;
FOR_KW : 'for' ;
IN_KW : 'in' ;
RETURN_KW : 'return' ;
SWITCH_KW : 'switch';
CASE_KW  : 'case';
DEFAULT  : 'default';

BREAK_KW : 'break' ;
CONTINUE_KW : 'continue' ;
FN_KW : 'fn' ;
MAIN_KW : 'main' ;

ATOI_KW : 'Atoi';
PARSEFLOAT_KW :'parse_float';
TYPEOF_KW : 'TypeOf';

INDEX_KW : 'indexOf';
JOIN_KW :'join';
//LONGITUD_KW :'len' ;

// === Operadores ===
PLUS    : '+' ;
MINUS   : '-' ;
MUL     : '*' ;
DIV     : '/' ;
MOD     : '%' ;
NOT     : '!' ;
OR      : '||' ;
AND     : '&&' ;
EQ      : '==' ;
NEQ     : '!=' ;
LT      : '<' ;
LE      : '<=' ;
GT      : '>' ;
GE      : '>=' ;
ASSIGN  : '=' ;
INC     : '++' ;
DEC     : '--' ;
INC2    : '+=' ;
DEC2    : '-=' ;
COLON  : ':' ;
PYC     : ';' ;

// === Símbolos ===
LPAREN  : '(' ;
RPAREN  : ')' ;
LBRACK  : '[' ;
LCOR  : '{' ;
RBRACK  : ']' ;
RCOR  : '}' ;
DOT     : '.' ;
COMMA   : ',' ;

// === Identificadores ===
ID : [a-zA-Z_][a-zA-Z0-9_]* ;

// === Espacios y comentarios ===
WS : [ \t\r\n]+ -> skip ;
LINE_COMMENT  : '//' ~[\r\n]* -> skip ;
BLOCK_COMMENT : '/*' .*? '*/' -> skip ;

/*

Gramatica -> Expresiones -> Declaraciones -> Asignaciones -> Println ->
-> Sentencias de Control -> Funciones -> Sentencias de Transferencia -> 
-> Structs -> Slices

 */