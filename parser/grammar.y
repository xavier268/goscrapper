%{
    package parser

    // each object has a value and a type.
    type value struct {
        v string    // a string in go that produce the value of the object
        t string    // a string representing the gotype of the object
        c int       // the code returned by lexer is stored here. Always set by the lexer, even for variables (set as IDENTIFIER). A valid go type, without spaces.
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
    list []string
}

%type <value> expression expressionAtom
%type <value> ope2 ope1
%type <value> typeDefinition 
%type <value> IDENTIFIER  NUMBER STRING BOOL

%type <list> returnList 



// definition des precedences et des associativités
// les opérateurs definis en dernier ont la précedence la plus élevée.
%nonassoc LTE LT GTE GT EQ NEQ CONTAINS ASSIGN
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

/* NB : prefer left recursivity, easier to implement */


program
    : head init body { yylex.(*myLexer).finalize() }
    | init body      { yylex.(*myLexer).finalize() }
    ;

init // run before any body
    : { yylex.(*myLexer).incOut() ; yylex.(*myLexer).addLines("{")}
    ;

// head defines options
head 
    : head options 
    | options
    ;

options
    : AT IDENTIFIER  typeDefinition  { yylex.(*myLexer).declInputParam($2.v, $3.v) }   // declare input parameter
    ;

typeDefinition
    : INTTYPE
    | STRINGTYPE
    | BOOLTYPE
    | LBRACKET  typeDefinition RBRACKET { $$.v = "[]" + $2.v}
    /* to do - add objects */
    ;

// program body contains statements, followed by either RETURN or 

body
    : statements returnStatements {  yylex.(*myLexer).addLines("}")}
    | returnStatements { yylex.(*myLexer).addLines("}")}
    ;

statements
    : statement { /* todo */}
    | statements statement { /* todo */}
    ;

statement 
    : IDENTIFIER ASSIGN expression { yylex.(*myLexer).vSetVar($1.v, $3)}
    | PAGE expression { /* todo */}
    | CLICK  expression { /* todo */}
    ;

returnStatements
    : RETURN returnList { yylex.(*myLexer).declOutputParams($2) ; yylex.(*myLexer).saveOut() ;  }
    | loopClause body { /* */ }
    ;

returnList
    : IDENTIFIER { $$ = append($$, $1.v) }
    | returnList COMMA IDENTIFIER { $$ = append($1, $3.v) }
    ;

loopClause
    : FOR IDENTIFIER IN expression  {yylex.(*myLexer).forNameInExpression($2.v, $4)}
    | SELECT  expression  { yylex.(*myLexer).selectExpression($2) } // loop on all css matching exprerssion
    // | SELECT IDENTIFIER expression // SELECT with a counter.
    ;

// ==================

ope2 // binary operators
    : PLUS{ $$ = $1 }
    | MINUS{ $$ = $1 }
    | MULTI{ $$ = $1 }
    | DIV{ $$ = $1 }
    | MOD{ $$ = $1 }
    | GT{ $$ = $1 }
    | GTE{ $$ = $1 }
    | LT{ $$ = $1 }
    | LTE{ $$ = $1 }
    | EQ{ $$ = $1 }
    | NEQ{ $$ = $1 }
    | AND{ $$ = $1 }
    | OR{ $$ = $1 }
    | NOT{ $$ = $1 }
    | CONTAINS{ $$ = $1 }
    ;

ope1 // unary operators
    : MINUS{ $$ = $1 }
    | LOWER{ $$ = $1 }
    | UPPER{ $$ = $1 }
    | NOT{ $$ = $1 }
    ;

expression // never empty, type is controlled semantically, not syntaxically
    : expressionAtom { $$ = $1 }
    | expression ope2 expressionAtom { $$ = yylex.(*myLexer).vOpe2($2.c, $1, $3) }
    | ope1 expressionAtom { $$ = yylex.(*myLexer).vOpe1($1.c, $2) }    
    ;

expressionAtom // never empty
    : LPAREN expression RPAREN  { $$ = yylex.(*myLexer).vParen($2) }
    | IDENTIFIER { $$ = yylex.(*myLexer).vGetVar($1.v) }
    | IDENTIFIER LBRACKET expression RBRACKET { /* todo - expression must be a number, or maybe certain strings for an object key ? */ }
    | STRING { $$ = $1 }
    | NUMBER { $$ = $1 }
    | BOOL { $$ = $1 }
    ;


%%