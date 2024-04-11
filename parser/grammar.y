%{
    package parser

    import (
        "fmt"
    )

    // keep the compiler happy
    var _ = fmt.Println

    type tok struct {
        v string    // token cvalue
        t string    // token type
        c int       // lexer/parser constant code
    } 

   

    var lx *myLexer // shorthand for lx

%}
      
%union {            
    tok tok         // token read from lexer, implements Node.
    node Node       // default for statements and expression
    nodes Nodes     // default for lists of expressions or statements, implements Node.
    nodemap NodeMap // default set of Node, with string keys, using valid id syntax, implements Node.
    nodeWithBody NodeWithBody // a node that incorporates a set of nodes
    }

%type<node> litteral litteralArray atomExpression expression expression1 expression2 expression3 atomExpression ope0
%type<node> statement variable  keyValue key program returnStatement accessExpression
%type<nodes> expressionList body statements returnList returnList0
%type<nodemap> keyValueSet litteralObject
%type<nodeWithBody> loopStatement


%type<tok> ope1 ope2 ope2Bool loopVariable printOption0

%token <tok>  

BOOL
NUMBER
STRING 
IDENTIFIER

ASSIGN SEMICOLON CLICK INPUT IN 
PRINT RAW SLOW LEFT RIGHT MIDDLE 
RETURN COMMA FOR
SELECT AS FROM TO STEP
WHERE LIMIT
LPAREN RPAREN
LBRACKET RBRACKET
LBRACE RBRACE
DOT LEN 
PLUS MINUS PLUSPLUS MINUSMINUS MULTI DIV MOD ABS
NOT AND OR XOR NAND
EQ NEQ LT LTE GT GTE
CONTAINS
FIND PATH WITH JOIN PAGE
COLON TEXT ATTR OF
DISTINCT 
AT /* @ */
DOTDOT /* .. */
QUESTION /*?*/

NOW VERSION FILE_SEPARATOR



// definition des precedences et des associativités
// les opérateurs definis en dernier ont la précedence la plus élevée.
%nonassoc ASSIGN FOR NOW VERSION
%left OR XOR
%left AND NAND
%left NOT
%left IN RANGE
%left PAGE PRINT ATTR TEXT 
%left LT LTE GT GTE EQ NEQ
%left PLUS 
%left MINUS
%left MULTI DIV MOD
%left PLUSPLUS 
%left LBRACKET  DOT 

%%

/* NB : prefer left recursivity, easier to implement */


program
    : beforeProgram  body {$$ = nodeProgram{req : $2, invars : lx.ParamsList()}; lx.root = $$}
    ;

beforeProgram
    : {lx = yylex.(*myLexer)} // initialize useful things here ...
    ;

body
    : statements returnStatement { $$ = append($1, $2) }
    | returnStatement { $$ = Nodes{$1}}
    ;

statements
    : statement  { $$ = Nodes{$1}}
    | statements  statement { $$ = append($1, $2)}
    ;

statement // statements are ALWAYS followed by a semi-colon
    : IDENTIFIER ASSIGN expression SEMICOLON{ $$ = lx.newNodeAssign($1,  $3)}
    | CLICK  expression /*element*/  clickOptions0 SEMICOLON{ /*todo*/} // click on element - make sure you select it first !  
    | INPUT expression /*text*/ IN expression /*element*/ SEMICOLON {/*todo*/} // input text in element - make sure you select it first !

    // debug only !
    | PRINT printOption0 expressionList SEMICOLON {$$ = nodePrint{ nodes:$3, raw:($2.c == RAW)}} // print %v content of expression in expressionList
    | SLOW SEMICOLON {$$ = nodeSlow{m:nil}} // wait for a short delay, using SLOW_DELAY from runtime. STop waiting if context is cancelled.
    | SLOW expression SEMICOLON {$$ = nodeSlow{m:$1}} // wait for specified millis, falling back on SLOW_DELAY if millis <=0. STop waiting if context is cancelled.
    ;

printOption0
    : {$$ = tok{}} // print using %v
    | RAW {$$ = $1} // print using %#v
    ;

clickOptions0
    : {/*todo*/} 
    | clickOptions  {/*todo*/} 
    ;

clickOptions
    : clickOption   {/*todo*/} 
    | clickOptions clickOption   {/*todo*/} 
    ;

clickOption
    : LEFT  {/*todo*/} 
    | RIGHT  {/*todo*/} 
    | MIDDLE  {/*todo*/} 
    | atomExpression  {/*todo*/} // number of clicks
    ;

returnStatement
    : RETURN returnList0 SEMICOLON {$$ = nodeReturn{$2}} 
    | loopStatement body  { $$ = $1.appendBody($2)  }
    ;

returnList0
    : {$$ = Nodes{}} // no return arguments provided
    | returnList 
    ;

returnList
    : expression  {$$ = Nodes{$1}} 
    | returnList COMMA expression  {$$ = append($1, $3)} 

loopStatement
    : FOR  SEMICOLON  {$$ = lx.newNodeForLoop(nil, nil, nil, nil)} //infinite loop

    | FOR loopVariable FROM expression TO expression STEP expression SEMICOLON  
        {$$ = lx.newNodeForLoop($2, $4, $6, $8)} // numerical range
    | FOR  FROM expression TO expression STEP expression SEMICOLON SEMICOLON 
        {$$ = lx.newNodeForLoop(nil, $3, $5, $7)} // numerical range, no loop variable

    | FOR loopVariable FROM expression TO expression SEMICOLON  
        {$$ = lx.newNodeForLoop($2, $4, $6, nil)} // numerical range // numerical range
    | FOR FROM expression TO expression SEMICOLON  
        {$$ = lx.newNodeForLoop(nil, $3, $5, nil)} //     numerical range
    
    | FOR loopVariable IN expression SEMICOLON  {/*todo*/} //loop over array
    | FOR IN expression SEMICOLON  {/*todo*/} //loop over array, no loop variable

    | SELECT expression /*css*/ AS loopVariable FROM expression /*Elementer*/ selectOptions0 SEMICOLON 
        {/*todo*/} // select css elements
    | SELECT expression /*css*/ FROM expression /*Elementer*/ selectOptions0 SEMICOLON 
        {/*todo*/} // select css elements, no loop variable
    ;

loopVariable
    : IDENTIFIER  
    ;

selectOptions0
    :  {/*todo*/} 
    | selectOptions {/*todo*/}
    ;

selectOptions
    : selectOption {/*todo*/} 
    | selectOptions selectOption {/*todo*/} 
    ;

selectOption 
    : WHERE expression {/*todo*/} 
    | LIMIT expression {/*todo*/} 
    ;

variable
    : IDENTIFIER {$$ = lx.newNodeVariable($1, false, true)}     // get normal variable, check already declared
    | AT IDENTIFIER {$$ = lx.newNodeVariable($2, true, false)}  // get input variable
    ;

// ==============
// expressions
//===============

expression  // logical binary op
    : expression ope2Bool expression1 {$$ = lx.newNodeOpe2Bool($1, $2, $3)} 
    | expression1
    ;

expression1 // binary ops
    : expression1 ope2 expression2 {$$ = lx.newNodeOpe2($1,$2, $3)}
    | expression2
    ;

expression2 // unary ops
    : ope1 expression2{$$ = lx.newNodeOpe1($1, $2)} 
    | expression3
    ;

expression3 // manage access and compound litteral expressions
    : atomExpression
    | accessExpression
    | litteralArray {$$ = $1} 
    | litteralObject {$$ = $1} 
    ;

atomExpression
    : litteral 
    | variable    
    | ope0 
    | LPAREN expression RPAREN {$$ = $2} 
    ;

litteral 
    : STRING  {$$ = lx.newNodeLitteral($1)} 
    | NUMBER {$$ = lx.newNodeLitteral($1)} 
    | BOOL {$$ = lx.newNodeLitteral($1)}     
    ;

accessExpression
    : atomExpression LBRACKET expression RBRACKET { $$ = nodeArrayAccess{a:$1, i:$3}} 
    | atomExpression DOT key {$$ = nodeMapAccess{m:$1,k:$3}} 
    ;

litteralArray 
    : LBRACKET expressionList RBRACKET {$$ = $2}
    | LBRACKET RBRACKET { $$ = Nodes(nil)} // empty array is ok.
    ;

expressionList
    : expression {$$ = Nodes{$1}}
    | expressionList COMMA expression  {$$ = append($1, $3)}
    ;

litteralObject
    : LBRACE RBRACE {$$ = lx.newNodeMap(nil, nil)} // ok to be empty
    | LBRACE keyValueSet RBRACE {$$ = $2}
    ;

keyValueSet
    : keyValue {$$ = lx.newNodeMap(nil, $1)}
    | keyValueSet COMMA keyValue {$$ = lx.newNodeMap($1, $3)}
    ;

keyValue
    : key COLON expression {$$ = lx.newNodeKeyValue($1, $3)}
    ;

ope0 // no argument operator, ie, read-only system values.
    : NOW { $$ = nodeOpe0($1)}// time stamp
    | VERSION { $$ = nodeOpe0($1)}// this version
    | FILE_SEPARATOR { $$ = nodeOpe0($1)} // file separator for current system
    ;

ope1 // unary operators. Action depends on argument type.
    : PLUS
    | MINUS 
    | PLUSPLUS  
    | MINUSMINUS  
    | ABS 
    | LEN  
    | NOT 
    
    | PAGE // PAGE url -> page object
    | TEXT // TEXT ele -> text content of element
    ;


ope2 // binary operators. Action performed depends of argument types.
    : PLUS 
    | MINUS
    | MULTI  
    | DIV
    | MOD 
    | EQ
    | NEQ
    | GT 
    | GTE 
    | LT 
    | LTE

    | ATTR // el ATTR at -> value of at attribute in el
    ;

ope2Bool //  binary booleans
    : AND 
    | OR
    | NAND
    | XOR
    ;

key
    : IDENTIFIER {$$ = lx.newNodeKey($1)}

%%