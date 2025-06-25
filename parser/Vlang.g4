//gramatica completa Vlang.g4
grammar Vlang; 

// === Axioma principal ===
programa: (stmt)+ EOF;

stmt:
	 decl_stmt                
	 | assign_stmt
     | if_stmt
     | while_stmt  
     | for_stmt                 // FOR ORIGINAL: for ID in expresion/range
     | for_enhanced_stmt        // NUEVOS TIPOS DE FOR
     | switch_stmt              // NUEVO: switch statement
     | func_call 
	 | func_dcl
     | struct_dcl
     | strct_instancia
     | transfer_stmt       
     | println 
     | print
     ;

if_stmt: if_chain (ELSE_KW if_chain)* else_stmt? # IfStmt;

if_chain: IF_KW expresion  LCOR stmt* RCOR # IfChain;

else_stmt: ELSE_KW LCOR stmt* RCOR # ElseStmt;

println: 'println'  LPAREN expresion (COMMA expresion)* RPAREN # PrintlnStmt;
print: 'print'  LPAREN expresion (COMMA expresion)* RPAREN # PrintStmt;

while_stmt: WHILE_KW expresion LCOR stmt* RCOR # WhileStmt;

// === FOR ORIGINAL (mantener compatibilidad) ===
for_stmt:
	FOR_KW ID IN_KW (expresion | range) LCOR stmt* RCOR # ForStmt;

// === NUEVOS TIPOS DE FOR ===
for_enhanced_stmt:
    // Tipo 1: for condición { }
    FOR_KW expresion LCOR stmt* RCOR                                    # ForCondition
    
    // Tipo 2: for inicialización; condición; incremento { }
    | FOR_KW for_init? SEMICOLON expresion? SEMICOLON for_update? LCOR stmt* RCOR    # ForClassic
    
    // Tipo 3: for índice, valor in slice { }
    | FOR_KW ID COMMA ID IN_KW expresion LCOR stmt* RCOR               # ForIndexValue
    
    // Tipo 4: for infinito { }
    | FOR_KW LCOR stmt* RCOR                                           # ForInfinite
    
    // NUEVO: Tipo 5: for índice, valor = range expresion { } - PARA ARRAYS
    | FOR_KW ID COMMA ID ASSIGN 'range' expresion LCOR stmt* RCOR     # ForRange
    ;

// === AUXILIARES PARA FOR CLÁSICO ===
for_init:
    decl_stmt       // mut i = 0
    | assign_stmt   // i = 0
    ;

for_update:
    assign_stmt     // i = i + 1, i += 1
    | incredecre    // i++, i--
    ;

// === SWITCH STATEMENT ===
switch_stmt:
    SWITCH_KW expresion? LCOR switch_case* default_case? RCOR          # SwitchStmt
    ;

switch_case:
    CASE_KW case_values COLON stmt*                                    # SwitchCase
    ;

case_values:
    expresion (COMMA expresion)*                                       # CaseValueList
    ;

default_case:
    DEFAULT_KW COLON stmt*                                             # DefaultCase
    ;

range: expresion DOT DOT DOT expresion # NumericRange;

/// Llamadas a funciones catch
// id ( 1,2,3,4)
func_call: id_pattern LPAREN (parametros)? RPAREN # FuncCall;

func_dcl: 
    // id(parametros | vacio )  { } 
    //    visitor -> recorrer esos stmts
    // fn id (x int, n string) 
    'fn' ID LPAREN (arg_list)? RPAREN (var_type)? LCOR stmt* RCOR # FuncDecl
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
    'struct' ID LCOR (struct_field)* RCOR # StructDecl
    ;

// Nueva regla para campos de struct
struct_field:
    var_type ID # StructField
    ;

// === TRANSFER STATEMENTS (ACTUALIZADO) ===
transfer_stmt:
	RETURN_KW expresion?	                    # ReturnStmt
	| BREAK_KW ID?                              # BreakStmt      // break o break label
	| CONTINUE_KW ID?                           # ContinueStmt   // continue o continue label
    | FALLTHROUGH_KW                            # FallthroughStmt // Para switch
    ;

assign_stmt:
	id_pattern ASSIGN expresion  	# DirectAssign
    | id_pattern PLUS_ASSIGN expresion      # PlusAssign
    | id_pattern MINUS_ASSIGN expresion     # MinusAssign
    | id_pattern MUL_ASSIGN expresion       # MulAssign
    | id_pattern DIV_ASSIGN expresion       # DivAssign
    | id_pattern MOD_ASSIGN expresion       # ModAssign
    ; 

decl_stmt: 
    MUT ID ASSIGN expresion # DeclAssign
    | MUT ID var_type ASSIGN expresion # DeclAssignType
    // NUEVO: Declaración sin inicialización
    | MUT ID var_type # DeclType
    // NUEVO: Declaración de arrays
    | ID ASSIGN array_literal # ArrayAssign
    ; 

// NUEVA REGLA: Para literales de arrays
array_literal:
    LBRACK RBRACK var_type LCOR (expresion (COMMA expresion)*)? RCOR # ArrayLiteral
    ;

id_pattern: ID (DOT ID)* # IdPattern;

// === Reglas de expresiones ===
expresion
    : valor                                                #valorexpresion        
    | func_call                                            #funcionexpre
    | assign_stmt                                          #asignacionExpr
    | LPAREN expresion RPAREN                              #parentesisexpre
    | LBRACK expresion (COMMA expresion)* RBRACK           #arrayexpre
    | op=(NOT | MINUS) expresion                           #unario

// ------------- >  Reglas de Expresiones Aritmeticas
// por facilidad la llamaremos a nuestra regla #BinaryExpr
/*¿Que puede tener una expresion binaria? |

una expresion binaria puede tener dos operandos y un operador.
el operando sera -> expresion 
los operadores -> (*| + | )

Entonces, nos ahorraremos trabajo en el BinaryExpr,
porque lo definiremos en el visit de BinaryExpr
 */
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
    ;

/*
param_list: func_param (COMMA func_param)* # ParamList;
func_param: ID? ID COLON INOUT_KW? type # FuncParam;
id   id : 
// external names -> num: value, num2: value2
arg_list: func_arg (COMMA func_arg)* # ArgList;
func_arg: (ID COLON)? (ANPERSAND)? (id_pattern | expr) # FuncArg; // 
             id : 
 
 */

// PARAMETROS -> Cuando se hace la llamada
// Argumentos -> Cuando se declara la funcion catch

// === Parámetros en llamadas ===
parametros: func_param (COMMA func_param)* # ParamList;
func_param : expresion #funcParam;

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
    | STRING_INTERPOLATION #valorStringInterpolation
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
	ID COLON expresion COMMA? 	# StructAttr
;

// tipos de primitivos que puede ser un struct catch
var_type
    : 'int'     # IntType
    | 'float'   # FloatType
    | 'f64'     # F64Type
    | 'string'  # StringType
    | 'bool'    # BoolType
    | 'char'    # CharType
    | 'void'    # VoidType
    | LBRACK RBRACK var_type # ArrayType
    ;

struct_vector: LBRACK ID RBRACK LPAREN RPAREN # StructVector;

// === Incremento / Decremento ===
incredecre
    : ID INC    #incremento
    | ID DEC    #decremento
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
STRING_INTERPOLATION : '"' (~["\\$] | '\\' . | '$' ~[{] | '$' '{' (~[}])* '}')* '"' ;

// === Identificadores ===
IF_KW : 'if' ;
ELSE_KW : 'else' ;
WHILE_KW : 'while' ;
FOR_KW : 'for' ;
IN_KW : 'in' ;
RETURN_KW : 'return' ;
BREAK_KW : 'break' ;
CONTINUE_KW : 'continue' ;

// === NUEVOS TOKENS PARA SWITCH ===
SWITCH_KW : 'switch' ;
CASE_KW : 'case' ;
DEFAULT_KW : 'default' ;
FALLTHROUGH_KW : 'fallthrough' ;

ID : [a-zA-Z_][a-zA-Z0-9_]* ;

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
PLUS_ASSIGN : '+=' ;
MINUS_ASSIGN : '-=' ;
MUL_ASSIGN : '*=' ;
DIV_ASSIGN : '/=' ;
MOD_ASSIGN : '%=' ;
COLON  : ':' ;
SEMICOLON : ';' ;

// === Símbolos ===
LPAREN  : '(' ;
RPAREN  : ')' ;
LBRACK  : '[' ;
LCOR  : '{' ;
RBRACK  : ']' ;
RCOR  : '}' ;
DOT     : '.' ;
COMMA   : ',' ;

// === Espacios y comentarios ===
WS : [ \t\r\n]+ -> skip ;
LINE_COMMENT  : '//' ~[\r\n]* -> skip ;
BLOCK_COMMENT : '/*' .*? '*/' -> skip ;