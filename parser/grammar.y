%{
    package parser

    // each object has a value and a type.
    type value struct {
        v string    // a string in go that produce the value of the object
        t string    // a string representing the gotype of the object
        c int       // the code returned by lexer is stored here. Always set by the lexer, even for variables (set as IDENTIFIER). A valid go type, without spaces.
    }

    var lx *myLexer // shorthand for lx

   
%}


%token <value>   // all token use value
            MULTI DIV MOD PLUS MINUS PLUSPLUS MINUSMINUS
            LTE GTE LT GT EQ NEQ
            COLON SEMICOLON DOT COMMA LBRACKET RBRACKET LPAREN RPAREN LBRACE RBRACE
            AND OR NOT
            DOTDOT ASSIGN QUESTION REGEXMATCH REGEXNOTMATCH
            LOWER UPPER FORMAT
            NOW TEXT HREF

            // Keywords
            INTTYPE BOOLTYPE STRINGTYPE BINTYPE // native types
            FOR RETURN WAITFOR OPTIONS IGNORE HEADLESS TIMEOUT 
            TRUE FALSE 
            EVENT
            LIKE IN WHILE
            BOOL AT IDENTIFIER IGNOREID STRING NUMBER NAMESPACESEPARATOR
            // selects
            SELECT ALL ANY ONE AS FROM WHERE LIMIT DISTINCT SORT ASC DESC DEFAULT CASE
            // html
            CLICK DOCUMENT PAGE CONTAINS 
            
%union {            
    value value
    list []string
    values []value
    mvalue map[string]value
}

%type <value> expression expressionUnary expressionAtom
%type <value> ope2 ope1
%type <value> typeDefinition INTTYPE BOOLTYPE BINTYPE
%type <value> IDENTIFIER  NUMBER STRING BOOL NOW

%type <list> returnList 

%type <values> expressionList keytypeList
%type <value> keytype

%type <mvalue>  keyExpressionList



// definition des precedences et des associativités
// les opérateurs definis en dernier ont la précedence la plus élevée.
%nonassoc LTE LT GTE GT EQ NEQ CONTAINS ASSIGN TEXT HREF 
%left PAGE
%left OR 
%left AND
%left NOT
%left MOD
%left PLUS 
%left MINUS
%left MULTI DIV
%left PLUSPLUS MINUSMINUS
%left LBRACKET DOT

%%

/* NB : prefer left recursivity, easier to implement */


program
    : beforeProgram head beforeBody body { lx.finalize() }
    | beforeProgram beforeBody body      { lx.finalize() }
    ;

beforeProgram
    : {lx = yylex.(*myLexer)}
    ;

beforeBody // run before any body
    : { lx.incOut() }
    ;

// head defines options
head 
    : head options 
    | options
    ;

options
    : AT IDENTIFIER  typeDefinition  { lx.declInputParam($2.v, $3.v) }   // declare input parameter
    ;

typeDefinition // contains the type as its value.v The value.t member has no meaning.
    : INTTYPE
    | STRINGTYPE
    | BOOLTYPE
    | BINTYPE { $$.v = "[]byte"} // translate "bin" into "[]byte, never use 'bin' anywhere"
    | LBRACKET  typeDefinition RBRACKET { $$.v = "[]" + $2.v}
    | LBRACE keytypeList RBRACE { $$.v = lx.objectType($2)}
    ;

keytype // value with value.v = key & value.t = type
    : IDENTIFIER COLON typeDefinition { $$ = value{v:$1.v,t:$3.v}}
    ;
keytypeList // as a list of values, containing the description of object.
    : keytype { $$ = []value{$1}}
    | keytypeList COMMA keytype{ $$ = append($1, $3)}
    ;

// program body contains statements, followed by either RETURN or 

body
    : statements returnStatements {  }
    | returnStatements { }
    ;

statements
    : statement { /* todo */}
    | statements statement { /* todo */}
    ;

statement 
    : IDENTIFIER ASSIGN expression { lx.vSetVar($1.v, $3)}
    | CLICK  expression { /* todo */}
    ;

returnStatements
    : RETURN returnList { lx.declOutputParams($2) ; lx.saveOut() ;  }
    | loopClause body afterLoop { /* */ }
    ;

afterLoop
    : { lx.addLines("}")}

returnList
    : IDENTIFIER { $$ = append($$, $1.v) }
    | returnList COMMA IDENTIFIER { $$ = append($1, $3.v) }
    ;

loopClause
    : FOR IDENTIFIER IN expression  {lx.forNameInExpression($2.v, $4)}
    // | SELECT  ALL expression  { lx.selectExpression($3) } // loop on all css matching exprerssion
    // | SELECT IDENTIFIER expression // SELECT with a counter.

    // below, FROM should be a page or an rod.Element, identifier will be set to a rod.Element
    | SELECT FROM IDENTIFIER ALL expression AS IDENTIFIER selectOptions {lx.addLines("{// select TODO" );} // do not wait
    | SELECT FROM IDENTIFIER ONE expression AS IDENTIFIER {lx.addLines("{// select TODO" );}// one exactly, and wait for it
    // below, FROM should be a page or an rod.Element, identifier will be set to the expression specified for the matched css
    | SELECT FROM IDENTIFIER AS IDENTIFIER ANY cases {lx.addLines("{// select TODO" );} // one exactly, and wait for it
    | SELECT FROM IDENTIFIER ANY AS IDENTIFIER  cases {lx.addLines("{// select TODO" );} // one exactly, and wait for it - alternative syntax
    ;

selectOptions
    : {/* options are optionnal */}
    | selectOptions WHERE expression {/*todo*/}
    | selectOptions LIMIT expression {/*todo*/}
    | selectOptions SORT ASC {/*todo*/}
    | selectOptions SORT DESC {/*todo*/}
    | selectOptions DISTINCT {/*todo*/}

cases // at least one case is required
    : case{/*todo*/}
    | cases case{/*todo*/}
    ;

case // when one of the css expr is found, _element is set to the matched element, loop variable is set to the expression after the COLON
     // the expression after the colon have no access to access the matched element.
    :  CASE expression COLON expression{/*todo*/} // $2 is a css, the loop variable is set to the $4 expression.
    | DEFAULT  expression {/*todo*/} // set variable to (*rodElement)nil
    ;



// ==================

ope2 // binary operators
    : PLUS{ $$ = $1 } // concat strings, append to array
    | PLUSPLUS {$$ = $1} // aggregate arrays
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
    | PLUSPLUS {$$ = $1 }
    | MINUSMINUS {$$ = $1 }
    | LOWER{ $$ = $1 }
    | UPPER{ $$ = $1 }
    | NOT{ $$ = $1 }
    | PAGE {$$ = $1}
    ;

expression // never empty, type is controlled semantically, not syntaxically
        // unary expression are always evalueated first, before binary operations.
    : expressionUnary { $$=$1 }   
    | expression ope2 expressionUnary { $$ = lx.vOpe2($2.c, $1, $3) } 
    ;

expressionUnary // never empty
    : expressionAtom { $$ = $1 }  
    | ope1 expressionAtom { $$ = lx.vOpe1($1.c, $2) }    
    ;

expressionAtom // never empty
    : LPAREN expression RPAREN  { $$ = lx.vParen($2) }    
    | expressionAtom LBRACKET expression RBRACKET {$$ = lx.vGetElementOf($1, $3)}
    | expressionAtom DOT IDENTIFIER { $$ = lx.vAccessObject($1, $3.v)}
    | LBRACKET expressionList RBRACKET { $$ = lx.vMakeArray($2)}
    | LBRACE keyExpressionList RBRACE {$$ = lx.vMakeObject($2)} 

    | TEXT expressionAtom {/*todo*/} // will retrieve the text for the provided cssselector or element
    | HREF expressionAtom {/* */}

    | IDENTIFIER { $$ = lx.vGetVar($1.v) }
    | STRING { $$ = $1 }
    | NUMBER { $$ = $1 }
    | BOOL { $$ = $1 }
    | NOW { $$ = value{v:"time.Now()", t: "time.Time"} ; lx.imports["time"] = true}
    ;

expressionList  
    : expression { $$ = []value{$1}}
    | expressionList COMMA expression { $$ = append($1, $3) } // no type checks yet
    ;

keyExpressionList // its a map[string]value, mapping the key to the expression value
    : IDENTIFIER COLON expression { $$ = map[string]value{$1.v : $3}}   
    | keyExpressionList COMMA IDENTIFIER COLON expression { $$ = $1 ; $$[$3.v]=$5}
    // implicit kety name from identifier
    | IDENTIFIER { $$ = map[string]value{$1.v:lx.vGetVar($1.v)}}
    | keyExpressionList COMMA IDENTIFIER { $$ = $1;$$[$3.v] = lx.vGetVar($3.v)}
    ;

%%