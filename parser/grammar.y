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
    

expression // is something that evaluates to a value.
   : unaryOperator expression
   | expression OR expression
   | expression AND expression
   | predicate
   ;

unaryOperator
   : NOT
   ;

predicate
    : predicate compareOperator predicate
    | predicate IN predicate
    | expressionAtom
    ;

compareOperator
    : EQ
    | NEQ
    | GT
    | GTE
    | LT
    | LTE
    ;


expressionAtom
    : expressionAtom MULTI expressionAtom
    | expressionAtom DIV expressionAtom
    | expressionAtom MOD expressionAtom
    | expressionAtom PLUS expressionAtom
    | expressionAtom MINUS expressionAtom
    | functionCallExpression
    | litteral
    | LPAREN expression RPAREN
    ;

litteral
    : NUMBER
    | BOOL
    | STRING
    | arrayLitteral
    | LBRACE paramList RBRACE
    | NONE
    ;

arrayLitteral   
    : LBRACKET paramList RBRACKET
    ;


%%
