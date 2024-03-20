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
%token AND OR
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


%type <Exp> expression term factor

%%

expression:
    expression PLUS term
    | term
    ;

term:
    term MULTI factor
    | factor
    ;

factor:
    NUMBER {}
    ;

%%


// ========== Lexer implementation ======================

