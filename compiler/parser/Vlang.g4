grammar Vlang;

// === Reglas del Parser ===

programa: (stmt)+ EOF;

stmt:
    (  decl_stmt
    | assign_stmt
    | if_stmt
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
    | block_stmt
    ) PUNTO_COMA?
    ;

block_stmt: LLA_IZQ stmt* LLA_DER # BlockStmt;

if_stmt: if_chain (ELSE if_chain)* else_stmt? # IfStmt;

if_chain: IF expresion LLA_IZQ stmt* LLA_DER # IfChain;

else_stmt: ELSE LLA_IZQ stmt* LLA_DER # ElseStmt;

println: PRINTLN PAR_IZQ expresion (COMA expresion)* PAR_DER # PrintlnStmt;
print: PRINT PAR_IZQ expresion (COMA expresion)* PAR_DER # PrintStmt;

forwhile_stmt: FOR expresion LLA_IZQ stmt* LLA_DER # WhileStmt;

forClasic_stmt: FOR assign_stmt PUNTO_COMA expresion PUNTO_COMA incredecre LLA_IZQ stmt* LLA_DER # ClasicFor;

for_stmt:
    FOR ID 'in' (expresion | range) LLA_IZQ stmt* LLA_DER # ForStmt;

range: expresion '...' expresion # NumericRange;

switch_stmt:
    SWITCH ID LLA_IZQ (case_clause)+ (default_clause)? LLA_DER # SwitchStmt;

case_clause:
    CASE (expresion (COMA expresion)*) D_PUNTO
        stmt* # CaseClause;

default_clause:
    DEFAULT D_PUNTO
        stmt* # DefaultClause;

slice_stmt: ID ASIGN '[' ']' var_type LLA_IZQ (expresion (COMA expresion)*) LLA_DER #DeclSlices;

func_call: id_pattern PAR_IZQ (parametros)? PAR_DER # FuncCall;

func_dcl:
    FN ID PAR_IZQ (arg_list)? PAR_DER (var_type)? LLA_IZQ stmt* LLA_DER # FuncDecl;

struct_dcl:
    'struct' ID LLA_IZQ (assign_stmt)* LLA_DER # StructDecl;

transfer_stmt:
    RETURN expresion?    # ReturnStmt
    | BREAK               # BreakStmt
    | 'continue'          # ContinueStmt
    ;

assign_stmt:
    id_pattern ASIGN expresion      # DirectAssign
    ;

decl_stmt:
    MUT ID ASIGN expresion                # DeclAssignInferred
    | MUT ID var_type (ASIGN expresion)?  # DeclAssignExplicit
    ;

list_expresiones: expresion (COMA expresion)* # ListExpressions;

id_pattern: ID ('.' ID)* # IdPattern;

expresion
    : valor                                             # valorExpresion
    | func_call                                         # funcionExpresion
    | PAR_IZQ expresion PAR_DER                           # parentesisExpresion
    | op=('!' | '-') expresion                            # unarioExpresion
    | left=expresion op=('*' | '/' | '%') right=expresion # BinaryExp
    | left=expresion op=('+' | '-') right=expresion      # BinaryExp
    | left=expresion op=('<' | '<=' | '>' | '>=') right=expresion # BinaryExp
    | left=expresion op=('==' | '!=') right=expresion     # BinaryExp
    | left=expresion op='&&' right=expresion              # BinaryExp
    | left=expresion op='||' right=expresion              # BinaryExp
    | ID                                                # id
    | incredecre                                        # increDecreExp
    | id_pattern '[' expresion ']'                      # indexAcceso
    ;

parametros: func_param (COMA func_param)* # ParamList;
func_param : id_pattern #funcParam;

arg_list: func_arg (COMA func_arg)* # ArgList;
func_arg: ID var_type # FuncArg;

valor
    : ENTERO    # valorEntero
    | FLOAT     # valorFloat
    | CADENA    # valorCadena
    | BOOLEANO  # valorBooleano
    ;

strct_instancia: 'struct' ID LLA_IZQ struct_prop* LLA_DER # StructInstancia;

struct_prop:
    var_type ID D_PUNTO? (expresion)? (COMA)?    # StructAttr
    ;

var_type
    : 'int'     # IntType
    | 'float'   # FloatType
    | 'string'  # StringType
    | 'bool'    # BoolType
    | 'char'    # CharType
    | 'void'    # VoidType
    ;

incredecre
    : ID INC    # incremento
    | ID DEC    # decremento
    ;


// === Reglas del Lexer ===

// Palabras Clave
MUT: 'mut';
VAR: 'var';
PRINT: 'print';
PRINTLN: 'println';
IF: 'if';
ELSE: 'else';
SWITCH: 'switch';
CASE: 'case';
DEFAULT: 'default';
BREAK: 'break';
FOR: 'for';
FN: 'fn';
RETURN: 'return';

// Literales e Identificadores
ID:         [a-zA-Z_][a-zA-Z_0-9]*;
FLOAT:      [0-9]+ '.' [0-9]+;
ENTERO:     [0-9]+;
CADENA:     '"' (ESC | ~["\\])* '"';
BOOLEANO:   'true' | 'false';

// Operadores y Puntuación
ASIGN:      '=';
SUM:        '+';
RESTA:      '-';
MUL:        '*';
DIV:        '/';
MAYOR:      '>';
MENOR:      '<';
MAYOR_IGUAL:'>=';
MENOR_IGUAL:'<=';
IGUAL:      '==';
DIFERENTE:  '!=';
AND:        '&&';
OR:         '||';
INC:        '++';
DEC:        '--';
PUNTO_COMA: ';';
COMA:       ',';
D_PUNTO:    ':';
PAR_IZQ:    '(';
PAR_DER:    ')';
LLA_IZQ:    '{';
LLA_DER:    '}';

// Comentarios y Espacios en Blanco
COMENTARIO_LINEA: '#' ~[\r\n]* -> skip;
COMENTARIO_BLOQUE: '/*' .*? '*/' -> skip;
WS:         [ \t\r\n]+ -> skip;

// Fragmentos
fragment ESC: '\\' (["\\/bfnrt] | 'u' [0-9a-fA-F]{4});