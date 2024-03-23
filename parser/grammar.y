%{
    package parser


import (
    "fmt"
)
    // each object has a value and a type.
    type value struct {
        v string // a string in go that produce the value of the object
        t string // a string representing the gotype of the object
        c int // the code returned by lexer is stored here - it is always set, even for variables (set as IDENTIFIER)
    }

   

   


%}


%token <value>   // all token use value
            MULTI DIV MOD PLUS MINUS PLUSPLUS
            LTE GTE LT GT EQ NEQ
            COLON SEMICOLON DOT COMMA LBRACKET RBRACKET LPAREN RPAREN LBRACE RBRACE
            AND OR NOT
            DOTDOT ASSIGN QUESTION REGEXMATCH REGEXNOTMATCH
            LOWER UPPER

            // Keywords
            INTTYPE BOOLTYPE STRINGTYPE // native types
            FOR RETURN WAITFOR OPTIONS IGNORE HEADLESS TIMEOUT 
            DISTINCT FILTER CURRENT SORT LIMIT LET COLLECT 
            ASC DESC NONE NULL TRUE FALSE USE
            INTO KEEP WITH COUNT ALL ANY AGGREGATE
            EVENT
            LIKE IN WHILE
            BOOL AT IDENTIFIER IGNOREID STRING NUMBER NAMESPACESEPARATOR
            // actions
            SELECT CLICK DOCUMENT PAGE CONTAINS

%union {            
    value value
}

%type <value> expression expressionList expressionAtom
%type <value> number string bool ope2 ope1
%type <value> IDENTIFIER NUMBER STRING BOOL



// definition des precedences et des associativités
// les opérateurs definis en dernier ont la précedence la plus élevée.
%nonassoc LTE LT GTE GT EQ NEQ CONTAINS
%left OR 
%left AND
%left NOT
%left PLUS
%left MINUS
%left MULTI 
%left DIV
%left MOD
%left LOWER UPPER





%%

/* NB : preferer recursivité gauche plus facile à implémenter et respecte associativité */


program
    : head body { yylex.(*myLexer).finalize() }
    | body      { yylex.(*myLexer).finalize() }
    ;


// head defines options
head 
    : head options 
    | options
    ;

options
    : IGNORE expressionList{ 
            // yylex.(*myLexer).imports["fmt"] = true
            line := fmt.Sprintf("rt.Ignore(%s)", $2.v)
            yylex.(*myLexer).addLines(line)
     }
    | AT IDENTIFIER  typeDefinition  { /* todo */ }
    ;

typeDefinition
    : INTTYPE
    | STRINGTYPE
    | BOOLTYPE
    ;

expressionList
    : expression      {
            if $1.t != "string" {
                yylex.(*myLexer).errorf("a string list should be made of strings only")
            }
            $$.v = $1.v
            $$.t = "string"

                     }             
    | expressionList COMMA expression   { 
            if $3.t != "string" {
                yylex.(*myLexer).errorf("a string list should be made of strings only")
            }
            $$.v = $1.v + "," + $3.v
            $$.t = "string"
    }
    ;

// program body contains statements, followed by either RETURN or 

body
    : statements returnExpression { /* todo */}
    | returnExpression { /* todo */}
    ;

statements
    : statement { /* todo */}
    | statements statement { /* todo */}
    ;

statement 
    : IDENTIFIER ASSIGN expression { /*todo*/}
    | PAGE expression { /* todo */}
    | SELECT expression { /* todo */}
    | CLICK  expression { /* todo */}
    ;

returnExpression
    : RETURN expression { /* todo */}
    | RETURN  { /* todo */ }
    | FOR IDENTIFIER IN expression body { /* todo */}
    ;




// litterals

string
    : STRING     
    ;
number
    : NUMBER 
    ;
bool    
    : BOOL 
    ;

// variable
variable
    : IDENTIFIER
    ;


ope2 // binary operators
    : PLUS
    | MINUS
    | MULTI
    | DIV
    | MOD
    | GT
    | GTE
    | LT
    | LTE
    | EQ
    | NEQ
    | AND
    | OR
    | NOT
    | CONTAINS
    ;

ope1 // unary operators
    : MINUS
    | LOWER
    | UPPER
    ;

expression // never empty, type is controlled semantically, not syntaxically
    : expressionAtom { $$ = $1 }
    | expression ope2 expressionAtom { $$ = yylex.(*myLexer).Ope2($2.c, $1, $3) }
    | ope1 expressionAtom { $$ = yylex.(*myLexer).Ope1($1.c, $2) }
    
    ;

expressionAtom // never empty
    : LPAREN expression RPAREN  { $$ = yylex.(*myLexer).Paren($2) }
    | variable { /* todo */ }
    | string { /* todo */ }
    | number { /* todo */ }
    | bool { /* todo */ }
    ;


%%