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
                            yylex.(*myLexer).addLines( "rt.Ignore(" + $2.value + ")")
                        }
    | AT IDENTIFIER  IDENTIFIER { // @ paramName paramType
                            yylex.(*myLexer).setParam ($2.value,$3.value)
                        }
    ;

stringList
    : STRING                    { $$.value =  $1.value }
    | stringList COMMA STRING   { $$.value = $1.value +  "," + $3.value}
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
    | PAGE stringExpression
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