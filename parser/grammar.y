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

%type<node>             litteral litteralArray atomExpression expression expression1 expression2 expression3 atomExpression
                        variable  keyValue key program returnStatement accessExpression 
                        ope0
                        statement nonIfStatement matchedIfStatement openIfStatement
%type<nodes>            expressionList body statements returnList returnList0
%type<nodemap>          keyValueSet litteralObject
%type<nodeWithBody>     loopStatement
%type<tok>              ope1 ope2 ope2Bool loopVariable

%token <tok>  

BOOL
NUMBER
STRING 
IDENTIFIER

ASSIGN SEMICOLON CLICK INPUT IN 
SLOW LEFT RIGHT MIDDLE 
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

IF THEN ELSE

ASSERT FAIL

PRINT FORMAT RAW GO JSON GSC NL

DOLLAR NIL



// definition des precedences et des associativités
// les opérateurs definis en dernier ont la précedence la plus élevée.
%nonassoc ASSIGN FOR NOW VERSION 

%left COMMA
%left ELSE
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
    : statement SEMICOLON { $$ = Nodes{$1}}
    | statements  statement SEMICOLON { $$ = append($1, $2)}
    ;

statement // statements 
    : nonIfStatement 
    | matchedIfStatement 
    | openIfStatement 
    ;

openIfStatement
    : IF expression THEN nonIfStatement {$$ = nodeIf{cond: $2, t: $4}}
  //  | IF expression THEN nonIfStatement ELSE nonIfStatement {$$ = nodeIf{cond: $2, t: $4, e:$6}}
    | IF expression THEN matchedIfStatement {$$ = nodeIf {cond: $2, t: $4}}
    ;

matchedIfStatement
    : IF expression THEN nonIfStatement ELSE matchedIfStatement {$$ = nodeIf {cond: $2, t: $4, e: $6}}
    | IF expression THEN nonIfStatement ELSE nonIfStatement {$$ = nodeIf {cond: $2, t: $4, e: $6}}
    ;

nonIfStatement 
    : LPAREN statements RPAREN { $$ = $2 }

    | IDENTIFIER ASSIGN expression { $$ = lx.newNodeAssign($1,  $3, false)} // assign to local scope
    | DOLLAR IDENTIFIER ASSIGN expression{ $$ = lx.newNodeAssign($2,  $4, true)} // assign to global scope
    
    | CLICK  expression /*element*/  clickOptions0 { /*todo*/} // click on element - make sure you select it first !  
    | INPUT expression /*text*/ IN expression /*element*/  {/*todo*/} // input text in element - make sure you select it first !

    | FAIL {$$ = nodeFail{}} // abort with error messagge
    | FAIL expression {$$ = nodeFail{$2}} // abort with error message
    | ASSERT expression { $$ = nodeAssert {$2}} // assume expression is true, or fail

    | PRINT { $$ = nodePrint{nil}} // just print new line
    | PRINT expressionList  {$$ = nodePrint{$2}} // print content of expressions in expressionList in GO (%v) format
    
     // debug only !
    | SLOW  {$$ = nodeSlow{m:nil}} // wait for a short delay, using SLOW_DELAY from runtime. STop waiting if context is cancelled.
    | SLOW expression  {$$ = nodeSlow{m:$1}} // wait for specified millis, falling back on SLOW_DELAY if millis <=0. STop waiting if context is cancelled.
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
    | loopStatement SEMICOLON  body  { $$ = $1.appendBody($3)  }
    ;

returnList0
    : {$$ = Nodes{}} // no return arguments provided
    | returnList 
    ;

returnList
    : expression  {$$ = Nodes{$1}} 
    | returnList COMMA expression  {$$ = append($1, $3)} 

loopStatement
    : FOR    {$$ = lx.newNodeForLoop(nil, nil, nil, nil)} //infinite loop

    | FOR loopVariable FROM expression TO expression STEP expression   
        {$$ = lx.newNodeForLoop($2, $4, $6, $8)} // numerical range
    | FOR  FROM expression TO expression STEP expression   
        {$$ = lx.newNodeForLoop(nil, $3, $5, $7)} // numerical range, no loop variable

    | FOR loopVariable FROM expression TO expression   
        {$$ = lx.newNodeForLoop($2, $4, $6, nil)} // numerical range
    | FOR FROM expression TO expression   
        {$$ = lx.newNodeForLoop(nil, $3, $5, nil)} //  numerical range, no loop variable
    
    | FOR loopVariable IN expression   {$$ = lx.newNodeForArray($2, $4)} //loop over array
    | FOR IN expression   {$$ = lx.newNodeForArray(nil, $3)} //loop over array, no loop variable

    | SELECT expression /*css*/ AS loopVariable FROM expression /*Elementer*/ selectOptions0  
        {/*todo*/} // select css elements
    | SELECT expression /*css*/ FROM expression /*Elementer*/ selectOptions0  
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
    : IDENTIFIER {$$ = lx.newNodeVariable($1, false, true, false)}     // get normal (local) variable, check already declared
    | AT IDENTIFIER {$$ = lx.newNodeVariable($2, true, false, false)}  // get input parameter
    | DOLLAR IDENTIFIER {$$ = lx.newNodeVariable($2, false, true, true)}  // get GLOBAL variable
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
    | ope0 // special variables are zero-ary operators ...
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
    | NL { $$ = nodeOpe0($1)} // new line string
    | NIL { $$ = nodeOpe0($1)} // nil constant
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

    // format argument as string
    | RAW       // using %#v
    | GO        // using %v
    | GSC       // using GSC syntax
    | JSON      // JSON formatting
    ;


ope2 // binary operators. Action performed depends of argument types.
    : PLUS      // add numbers, concatenate strings, add element to array
    | PLUSPLUS  // merge two arrays
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

    | FORMAT // any FORMAT format -> string
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