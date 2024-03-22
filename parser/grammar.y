%{
    package parser





%}


%token      MULTI DIV MOD PLUS MINUS PLUSPLUS MINUSMINUS
            LTE GTE LT GT EQ NEQ
            COLON SEMICOLON DOT COMMA LBRACKET RBRACKET LPAREN RPAREN LBRACE RBRACE
            AND OR XOR
            DOTDOT ASSIGN QUESTION REGEXMATCH REGEXNOTMATCH

// Keywords
            FOR RETURN WAITFOR OPTIONS IGNORE HEADLESS TIMEOUT 
            DISTINCT FILTER CURRENT SORT LIMIT LET COLLECT 
            ASC DESC NONE NULL TRUE FALSE USE
            INTO KEEP WITH COUNT ALL ANY AGGREGATE
            EVENT
            LIKE NOT IN WHILE
            BOOL AT IDENTIFIER IGNOREID STRING NUMBER NAMESPACESEPARATOR
// actions
            SELECT CLICK DOCUMENT PAGE CONTAINS

%union {            // ceci redéclare yySymType !
    value string
}

%type <value> expression stringExpression stringList string 
%type <value> number numExpression
%type <value> bool boolExpression
%type <value> IDENTIFIER NUMBER STRING BOOL



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
    : head body { yylex.(*myLexer).finalize() }
    | body      { yylex.(*myLexer).finalize() }
    ;


// head defines options
head 
    : head options 
    | options
    ;

options
    : IGNORE stringList { 
                            yylex.(*myLexer).addLines( "rt.Ignore(" + $2 + ")")
                        }
    | AT IDENTIFIER  IDENTIFIER { // @ paramName paramType
                            yylex.(*myLexer).setParam ($2,$3)
                        }
    ;

stringList
    : STRING                   
    | stringList COMMA STRING   { $$ = $1 +  "," + $3}
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
    : IDENTIFIER ASSIGN expression { /* todo */}
    | PAGE stringExpression { /* todo */}
    | SELECT stringExpression { /* todo */}
    | CLICK  stringExpression { /* todo */}
    ;

returnExpression
    : RETURN expression { /* todo */}
    | FOR IDENTIFIER IN stringExpression body { /* todo */}
    ;


// expressions ...

expression
    : stringExpression
    | numExpression
    | boolExpression
    | IDENTIFIER { /* todo */}
    ;

stringExpression
    : string
    | stringExpression opeString2String string { /* todo */}
    | LPAREN stringExpression RPAREN { $$ = $2}
    ;

string
    : STRING     // litteral
    ;

numExpression
    : number { $$ = $1}
    // to complete ...  { /* todo */}
    | numExpression PLUS numExpression { $$ = "("+$1+")+("+ $3 + ")"}
    | numExpression MINUS numExpression { $$ = "("+$1+")-("+ $3 + ")"}
    | MINUS numExpression { $$ = "-(" + $2 + ")"}
    | LPAREN numExpression RPAREN { $$ = "(" + $2 + ")"}
    ;

number
    : NUMBER // litteral
    ;

boolExpression
    : bool
    | boolExpression opeBool2Bool bool { /* todo */}
    | NOT boolExpression { $$ = "!("+$2+")"}
    | stringExpression opeCompareString2Bool stringExpression { /* todo */}
    | numExpression opeCompareNum2Bool numExpression { /* todo */}
    | LPAREN boolExpression RPAREN { /* todo */}
    ;

bool    
    : BOOL // litteral
    ;

opeBool2Bool
    : AND
    | OR
    | XOR
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