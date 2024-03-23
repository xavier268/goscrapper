%{
    package parser





%}


%token      MULTI DIV MOD PLUS MINUS PLUSPLUS
            LTE GTE LT GT EQ NEQ
            COLON SEMICOLON DOT COMMA LBRACKET RBRACKET LPAREN RPAREN LBRACE RBRACE
            AND OR NOT
            DOTDOT ASSIGN QUESTION REGEXMATCH REGEXNOTMATCH

// Keywords
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
    value string // a string in go that produce the value of the object
    gtype string // a string representing the gotype of the object
}

%type <value> expression stringExpression stringList string 
%type <value> number numExpression
%type <value> bool boolExpression
%type <value> IDENTIFIER NUMBER STRING BOOL



// definition des precedences et des associativités
// les opérateurs definis en dernier ont la précedence la plus élevée.
%nonassoc LTE LT GTE GT EQ NEQ IN
%left OR 
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
    : IDENTIFIER ASSIGN expression { yylex.(*myLexer).setVar($1, $3)}
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
    ;

stringExpression
    : string
    | stringExpression PLUS string { $$ = "("+$1+")+("+ $3 + ")"}
    | LPAREN stringExpression RPAREN { $$ = "("+$2+")"}
    ;

string
    : STRING     // litteral
    | IDENTIFIER
    ;

numExpression
    : number { $$ = $1}
    // to complete ...  { /* todo */}
    | numExpression MULTI numExpression { $$ = "("+$1+")*("+ $3 + ")"}
    | numExpression DIV numExpression { $$ = "("+$1+")/("+ $3 + ")"}
    | numExpression MOD numExpression { $$ = "("+$1+")%("+ $3 + ")"}
    | numExpression PLUS numExpression { $$ = "("+$1+")+("+ $3 + ")"}
    | numExpression MINUS numExpression { $$ = "("+$1+")-("+ $3 + ")"}
    | MINUS numExpression { $$ = "-(" + $2 + ")"}
    | LPAREN numExpression RPAREN { $$ = "(" + $2 + ")"}
    ;

number
    : NUMBER // litteral
    | IDENTIFIER
    ;

boolExpression
    : bool
    | boolExpression EQ bool {$$ = "("+$1+"=="+$3+")" }
    | boolExpression NEQ bool {$$ = "("+$1+"!="+$3+")" }
    | boolExpression AND bool {$$ = "("+$1+"&&"+$3+")" }
    | boolExpression OR bool {$$ = "("+$1+"||"+$3+")" }
    | NOT boolExpression { $$ = "!("+$2+")"}
    | stringExpression EQ stringExpression {$$ = "("+$1+"=="+$3+")" }
    | stringExpression NEQ stringExpression {$$ = "("+$1+"!="+$3+")" }
    | stringExpression LT stringExpression {$$ = "("+$1+"<"+$3+")" }
    | stringExpression LTE stringExpression {$$ = "("+$1+"<="+$3+")" }
    | stringExpression GT stringExpression {$$ = "("+$1+">"+$3+")" }
    | stringExpression GTE stringExpression {$$ = "("+$1+">="+$3+")" }
    | numExpression EQ numExpression {$$ = "("+$1+"=="+$3+")" }
    | numExpression NEQ numExpression {$$ = "("+$1+"!="+$3+")" }
    | numExpression LT numExpression {$$ = "("+$1+"<"+$3+")" }
    | numExpression LTE numExpression {$$ = "("+$1+"<="+$3+")" }
    | numExpression GT numExpression {$$ = "("+$1+">"+$3+")" }
    | numExpression GTE numExpression {$$ = "("+$1+">="+$3+")" }
    | LPAREN boolExpression RPAREN { $$ = "("+$2+")"}
    ;

bool    
    : BOOL // litteral
    | IDENTIFIER
    ;

%%