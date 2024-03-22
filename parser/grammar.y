%{
    package parser

    import (
       // runtime used by scrapper
        _ "github.com/xavier268/goscrapper/rt"
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
%token FOR RETURN WAITFOR OPTIONS IGNORE HEADLESS TIMEOUT DISTINCT FILTER CURRENT SORT LIMIT LET COLLECT ASC DESC NONE NULL TRUE FALSE USE
// Group operators
%token INTO KEEP WITH COUNT ALL ANY AGGREGATE
// Wait operaor
%token EVENT
// Unary operators
%token LIKE NOT IN WHILE
// litterals
%token BOOL PARAM IDENTIFIER IGNOREID STRING NUMBER NAMESPACESEPARATOR
// actions
%token SELECT CLICK DOCUMENT CONTAINS 
// define DO after DOCUMENT ?
%token DO


// definition des precedences et des associativités
// les opérateurs definis en dernier ont la précedence la plus élevée.
%nonassoc LTE LT GTE GT EQ NEQ IN
%left OR XOR
%left AND
%left NOT
%left PLUS MINUS
%left MULTI DIV
%left MOD





%%

/* NB : preferer recursivité gauche plus facile à implémenter et respecte associativité */

program
    : head body
    | body
    ;


// head defines options
head 
    : head options 
    | options
    ;

options
    : HEADLESS
    | IGNORE stringList
    ;

stringList
    : STRING
    | stringList COMMA STRING
    ;

// program body contains statements, followed by either RETURN or 

body
    : statements returnExpression
    | returnExpression
    ;

statements
    : statement
    | statements statement
    ;

statement 
    : IDENTIFIER ASSIGN expression
    | DOCUMENT stringExpression
    | SELECT stringExpression
    | CLICK  stringExpression
    ;

returnExpression
    : RETURN expression
    | FOR IDENTIFIER IN stringExpression body
    ;


// expressions ...

expression
    : stringExpression
    | numExpression
    | boolExpression
    | IDENTIFIER // variable   
    ;

stringExpression
    : string
    | stringExpression opeString2String string
    | LPAREN stringExpression RPAREN
    ;

string
    : STRING     // litteral
    ;

numExpression
    : number
    | numExpression opeNum2Num number
    | numExpression MINUS numExpression
    | MINUS numExpression
    | LPAREN numExpression RPAREN
    ;

number
    : NUMBER // litteral
    ;

boolExpression
    : bool
    | boolExpression opeBool2Bool bool
    | NOT boolExpression
    | stringExpression opeCompareString2Bool stringExpression
    | numExpression opeCompareNum2Bool numExpression
    | LPAREN boolExpression RPAREN
    ;

bool    
    : BOOL // litteral
    ;

opeBool2Bool
    : AND
    | OR
    | XOR
    ;

opeNum2Num // except MINUS because unary
    : PLUS
    | DIV
    | MULTI
    | MOD
    ;


opeString2String
    : PLUS
    ;

opeCompareNum2Bool
    : GT
    | GTE
    | LT
    | LTE
    | EQ
    | NEQ
    ;

opeCompareString2Bool
    : GT
    | GTE
    | LT
    | LTE
    | EQ
    | NEQ
    | CONTAINS
    ;

%%