%{
    package parser

    import (
       // "fmt"
        )

   
	
    // Déclaration d'une interface pour gérer les différents types de noeuds de l'AST
    type Expr interface {}

%}


// Arithmetic operators
%token MULTI DIV MOD PLUS MINUS PLUSPLUS MINUSMINUS
// Comparison operators
%token LTE GTE LT GT EQ NEQ
// Punctuation
%token COLON SEMICOLON DOT COMMA LBRACKET RBRACKET LPAREN RPAREN LBRACE RBRACE
// Bool operators
%token AND OR XOR
// Other operators
%token DOTDOT ASSIGN QUESTION REGEXMATCH REGEXNOTMATCH
// Keywords
%token FOR RETURN WAITFOR OPTIONS TIMEOUT DISTINCT FILTER CURRENT SORT LIMIT LET COLLECT ASC DESC NONE NULL TRUE FALSE USE
// Group operators
%token INTO KEEP WITH COUNT ALL ANY AGGREGATE
// Wait operaor
%token EVENT
// Unary operators
%token LIKE NOT IN DO WHILE
// litterals
%token BOOL PARAM IDENTIFIER IGNOREID STRING NUMBER NAMESPACESEPARATOR

// definition des precedences et des associativités
// les opérateurs definis en dernier ont la précedence la plus élevée.
%left LTE LT GTE GT EQ NEQ IN
%left OR
%left AND
%left PLUS MINUS
%left MULTI DIV
%left MOD





%%

/* NB : preferer recursivité gauche plus facile à implémenter et respecte associativité */

body
    : bodyStatements bodyExpression
    | bodyExpression
    ;

bodyStatements
    : bodyStatement
    | bodyStatements bodyStatement 
    ;

bodyStatement // ne renvoie pas de valeur
    : variableDeclaration
    | functionCallExpression
    ;

variableDeclaration
    : IDENTIFIER ASSIGN expression
    ;

functionCallExpression
    : IDENTIFIER LPAREN paramList RPAREN
    ;

paramList 
    : /* empty list */
    | nonemptyParamList
    ;

nonemptyParamList
    : expression
    |  nonemptyParamList COMMA expression
    ;


// A body expression collects return values
bodyExpression  
    :  returnExpression
    | forExpression
    ;

returnExpression   
    : RETURN expression
    | RETURN DISTINCT expression
    ;

forExpression
    : FOR loopVariable IN forExpressionSource DO bodyExpression
    ;

forExpressionSource
    : functionCallExpression
    | sourceVariable
    | arrayLitteral
    ;

sourceVariable
    : IDENTIFIER
    ;

loopVariable
    : IDENTIFIER
    ;

variable
    : IDENTIFIER
    ;
    
boolOperator
    : AND
    | OR
    ;

compareOperator
    : EQ
    | NEQ
    | GT
    | GTE
    | LT
    | LTE
    ;

boolOperator
    : EQ
    | NEQ
    | AND
    | OR
    | XOR
    ;

arithOperator
    : PLUS
    | MINUS
    | MULTI
    | MOD 
    | DIV
    ;

expression // is something that evaluates to a value.
   : boolExpression
   | numExpression
   | arrayExpression
   | stringExpression
   | objectExpression
   | variable
   | NONE
   ;

boolExpression
    : BOOL
    | boolExpression boolOperator boolExpression
    | numExpression compareOperator numExpression
    | arrayExpression compareOperator arrayExpression
    | objectExpression compareOperator objectExpression
    | LPAREN boolExpression RPAREN
    | functionCallExpression
    ;

numExpression 
    : NUMBER
    | numExpression arithOperator numExpression
    | MINUS numExpression
    | expression PLUSPLUS
    | LPAREN numExpression RPAREN
    | functionCallExpression
    ;

arrayExpression
    : LBRACKET paramList RBRACKET
    | LPAREN arrayExpression RPAREN
    | functionCallExpression
    ;

stringExpression
    : STRING
    | stringExpression compareOperator stringExpression
    ;

objectExpression
    : objectLitteral
    | LBRACE objectKeyValueList RBRACE
    ;

objectKeyValueList 
    : // can be empty */
    | nonemptyKeyValueList
    ;

nonemptyKeyValueList
    : keyValue
    | nonemptyKeyValueList COMMA keyValue
    ;

keyValue
    : IDENTIFIER COLON expression
    | IDENTIFIER
    ;

litteral
    : NUMBER
    | BOOL
    | STRING
    | arrayLitteral
    | objectLitteral
    ;


arrayLitteral 
    : LBRACKET litteral RBRACKET
    ;


objectLitteral
    : LBRACE litteralKVList RBRACE
    ;

litteralKVList
    : /* can be empty */
    | nonemptylitteralKVList litteralKVList
    ;

nonemptylitteralKVList
    : litteralKV
    | nonemptylitteralKVList COMMA litteralKV
    ;

litteralKV
    : IDENTIFIER COLON litteral
    ;





%%
