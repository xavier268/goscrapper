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
    tok tok         // token read from lexer
    node Node       // default for statements and expression
    nodes Nodes     // default for lists of expressions or statements
    nodemap NodeMap // default set of Node, with string keys, using valid id syntax.
}

%type<node> litteral litteralArray expression statement variable keyValue key
%type<nodes> expressionList program body statements returnStatements 
%type<nodemap> keyValueSet litteralObject

%token <tok>  

BOOL
NUMBER
STRING 
IDENTIFIER

ASSIGN SEMICOLON CLICK INPUT IN 
PRINT RAW SLOW LEFT RIGHT MIDDLE 
RETURN COMMA FOR
SELECT AS FROM
WHERE LIMIT
LPAREN RPAREN
LBRACKET RBRACKET
LBRACE RBRACE
DOT LEN 
PLUS MINUS PLUSPLUS MINUSMINUS MULTI DIV MOD ABS
NOT AND OR XOR
EQ NEQ LT LTE GT GTE
CONTAINS
FIND PATH WITH JOIN PAGE
COLON TEXT ATTR OF
DISTINCT 
AT /* @ */
DOTDOT /* .. */
QUESTION /*?*/
BANG /*!*/



// definition des precedences et des associativités
// les opérateurs definis en dernier ont la précedence la plus élevée.
%nonassoc ASSIGN FOR
%left OR XOR
%left AND
%left NOT
%left IN RANGE
%left PAGE PRINT ATTR TEXT 
%left LT LTE GT GTE EQ NEQ
%left PLUS 
%left MINUS
%left MULTI DIV MOD
%left PLUSPLUS 
%left LBRACKET  
%left DOT 

%%

/* NB : prefer left recursivity, easier to implement */


program
    : beforeProgram  body {$$ = $2 ; lx.root = $$}
    ;

beforeProgram
    : {lx = yylex.(*myLexer)} // initialize useful things here ...
    ;

body
    : statements returnStatements { $$ = append($1, $2 ...) }
    | returnStatements { $$ = $1 }
    ;

statements
    : statement  { $$ = Nodes{$1}}
    | statements  statement { $$ = append($1, $2)}
    ;

statement // statements are always followed by a semi-colon !
    : IDENTIFIER ASSIGN expression SEMICOLON{ $$ = lx.newNodeAssign($1,  $3)}
    | CLICK  atomExpression /*element*/  clickOptions0 SEMICOLON{ /*todo*/} // click on element    
    | INPUT atomExpression /*text*/ IN atomExpression /*element*/ SEMICOLON {/*todo*/} // input text in element

    // debug only !
    | PRINT expressionList SEMICOLON {$$ = nodePrint{ nodes:$2, raw:false}} // print %#v content of expression in expressionList
    | PRINT RAW expressionList SEMICOLON {$$ = nodePrint{ nodes:$3, raw : true}} // print %v content of expression in expressionList
    | SLOW SEMICOLON {/*todo*/} // wait for a few seconds
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

returnStatements
    : RETURN returnList0 SEMICOLON {/*todo*/} 
    | RETURN DISTINCT returnList SEMICOLON {/*todo*/} 
    | loopStatement body  { /*todo*/ }
    ;

returnList0
    : {/*todo*/} // no return arguments
    | returnList  {/*todo*/} 
    ;

returnList
    : atomExpression  {/*todo*/} 
    | returnList COMMA atomExpression  {/*todo*/} 

loopStatement
    : FOR loopVariable IN expression SEMICOLON  {/*todo*/} 
    | SELECT expression /*css*/ AS loopVariable FROM expression selectOptions0 SEMICOLON {/*todo*/} 
    ;

loopVariable
    : IDENTIFIER  {/*todo*/} 
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
    : IDENTIFIER {$$ = lx.newNodeVariable($1, false)} // get normal variable
    | AT IDENTIFIER {$$ = lx.newNodeVariable($2, true)} // get input variable
    ;




// ==============
// expressions
//===============

expression  
    : expression ope2Bool expression1 {/*todo*/} 
    | expression1 {/*todo*/} 
    ;

expression1
    : expression1 ope2 expression2
    | expression2
    ;

expression2
    : ope1 expression2{/*todo*/} 
    | expression3{/*todo*/} 
    ;

expression3
    : atomExpression{/*todo*/} 
    | accessExpression{/*todo*/} 
    ;

atomExpression
    : litteral {/*todo*/} 
    | variable {/*todo*/}    
    | LPAREN expression RPAREN {/*todo*/} 
    ;

litteral 
    : STRING  {$$ = lx.newNodeLitteral($1)} 
    | NUMBER {$$ = lx.newNodeLitteral($1)} 
    | BOOL {$$ = lx.newNodeLitteral($1)} 
    | litteralArray {$$ = $1} 
    | litteralObject {$$=$1} 
    ;

accessExpression
    : atomExpression LBRACKET expression RBRACKET  {/*todo*/} 
    | atomExpression DOT key {/*todo*/} 
    ;

litteralArray 
    : LBRACKET expressionList RBRACKET {$$ = $2}
    | LBRACKET RBRACKET { $$ = Nodes(nil)}
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

ope1
    : PLUS{/*todo*/} 
    | MINUS {/*todo*/} 
    | PLUSPLUS {/*todo*/} 
    | MINUSMINUS {/*todo*/} 
    | ABS {/*todo*/} 
    | LEN  {/*todo*/} 
    | NOT  {/*todo*/} 
    ;


ope2
    : PLUS {/*todo*/} 
    | MINUS{/*todo*/} 
    | MULTI {/*todo*/} 
    | DIV {/*todo*/} 
    | MOD {/*todo*/} 
    | RANGE  {/*todo*/} 
    | EQ{/*todo*/} 
    | NEQ{/*todo*/} 
    | GT{/*todo*/} 
    | GTE{/*todo*/} 
    | LT{/*todo*/} 
    | LTE{/*todo*/} 
    ;

ope2Bool // lazily implemented
    : AND{/*todo*/} 
    | OR{/*todo*/} 
    ;

key
    : IDENTIFIER {$$ = lx.newNodeKey($1)}

%%