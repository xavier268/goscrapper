%{
    package parser

    // each object has a value and a type.
    type value struct {
        v string    // a string in go that produce the value of the object
        t string    // a string representing the gotype of the object
        c int       // the code returned by lexer is stored here. Always set by the lexer, even for variables (set as IDENTIFIER). A valid go type, without spaces.
    }

    var lx *myLexer // shorthand for lx

    // options for SELECT ONE, ANY, ALL
    type selopt struct {
        from value      // can be *rod.Page or *rod.Element
        css value       // css selector
        loopv string    // loop variable identifier
        where []value   // list of where conditions, applied on loopv
        limit value
        cases []casopt
    }

    // cases for select ANY
    type casopt struct {
        def bool
        e1 value
        e2 value
    }
   
%}


%token <value>   // all token use value
            MULTI DIV MOD PLUS MINUS PLUSPLUS 
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
    selopt selopt
    casopt casopt
    casopts []casopt
}

%type <value> expression expressionUnary expressionAtom
%type <value> ope2 ope1
%type <value> typeDefinition INTTYPE BOOLTYPE BINTYPE
%type <value> IDENTIFIER  NUMBER STRING BOOL NOW

%type <list> returnList 

%type <values> expressionList keytypeList
%type <value> keytype asClause

%type <mvalue>  keyExpressionList

%type <casopt> case
%type <casopts> cases
%type <selopt> selectOptions 


// definition des precedences et des associativités
// les opérateurs definis en dernier ont la précedence la plus élevée.
%nonassoc LTE LT GTE GT CONTAINS ASSIGN HREF 
%left EQ NEQ 
%left PAGE 
%left TEXT
%left OR 
%left AND
%left NOT
%left MOD
%left PLUS 
%left MINUS
%left MULTI DIV
%left PLUSPLUS 
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
    : INTTYPE{$$.v="int"}
    | STRINGTYPE{$$.v="string"}
    | BOOLTYPE {$$.v ="bool"}
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
    | CLICK expression FROM expression{ /* todo - css from page */}
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
    | SELECT FROM expression ALL expression asClause selectOptions {opt:= $7; opt.from=$3; opt.css=$5;opt.loopv=$6.v; lx.selectAll(opt)} // do not wait, but compatible with dynamic pages
    | SELECT FROM expression ONE expression asClause {lx.selectOne($3,$5,$6);}// one exactly, and wait for it
    // below, FROM should be a page or an rod.Element, identifier will be set to the expression specified for the matched css
    | SELECT FROM expression asClause ANY cases {lx.addLines("{// select TODO" );} // one exactly, and wait for it
    | SELECT FROM expression ANY asClause  cases {lx.addLines("{// select TODO" );} // one exactly, and wait for it - alternative syntax
    ;

asClause // pre-declares the select loop variable, so it is available in the where clause or case clause
    : AS IDENTIFIER {         
        $$ = $2; 
        if typ,ok := lx.vars[$2.v] ; ok {
            lx.errorf("variable %s was already declared (type : %s), cannot be redeclared as SELECT variable", $2.v, typ)
        }
        lx.vars[$2.v] = "*rod.Element";
        }
    ;


selectOptions
    : {$$ = selopt{}}
    | selectOptions WHERE expression {$$ = $1; $$.where = append($$.where, $3)}
    | selectOptions LIMIT expression {$$=$1; $$.limit = $3} // ovewrite previous limits
    ;

cases // at least one case is required
    : case{$$ = []casopt{}}
    | cases case{$$ = append($1, $2)}
    ;

case // when one of the css expr is found, loop variable is set to the matched rodElement, 
     // and the statemensta re executed. Staements have access to the loop variable.
     // NB : Notice the SEMICOLON after the statements group !
    : CASE expression COLON statements SEMICOLON {/*todo*/} // $2 is a css, the loop variable is set to the $4 expression.
    | DEFAULT COLON statements SEMICOLON {/*todo*/} // set loop variable to $2
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
    | LOWER{ $$ = $1 }
    | UPPER{ $$ = $1 }
    | NOT{ $$ = $1 }
    | PAGE {$$ = $1}
    | TEXT {$$ = $1}
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

    | IDENTIFIER { $$ = lx.vGetVar($1.v) }
    | STRING { $$ =value{v:$1.v, t:"string"} }
    | NUMBER { $$ =value{v:$1.v, t:"int"} }
    | BOOL { $$ =value{v:$1.v, t:"bool"} }
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