# goscrapper

Domain specific language for web-scrapping.

[![Go Reference](https://pkg.go.dev/badge/github.com/xavier268/goscrapper.svg)](https://pkg.go.dev/github.com/xavier268/goscrapper) [![Go Report Card](https://goreportcard.com/badge/github.com/xavier268/goscrapper)](https://goreportcard.com/report/github.com/xavier268/goscrapper)


- [goscrapper](#goscrapper)
- [Introduction](#introduction)
- [Gsc language reference](#gsc-language-reference)
  - [Request structure](#request-structure)
  - [Litteral types](#litteral-types)
    - [String litterals](#string-litterals)
  - [Variables and scope](#variables-and-scope)
    - [Compile time checks on variables](#compile-time-checks-on-variables)
    - [Runtime checks on variables](#runtime-checks-on-variables)
    - [Example](#example)
  - [Expressions and operators](#expressions-and-operators)
    - [Constants](#constants)
    - [Numerical operators](#numerical-operators)
    - [String operators](#string-operators)
    - [Array and element operators](#array-and-element-operators)
    - [Boolean operators](#boolean-operators)
    - [Pages and elements](#pages-and-elements)
    - [Time](#time)
  - [Statements](#statements)
    - [Assignement](#assignement)
    - [IF THEN ELSE](#if-then-else)
    - [RETURN statement](#return-statement)
    - [FOR](#for)
    - [SELECT](#select)
  - [Reserved keywords](#reserved-keywords)
- [Interpreter settings](#interpreter-settings)
  - [Using context](#using-context)
  - [Sync or Async modes](#sync-or-async-modes)
- [GSC command line tool](#gsc-command-line-tool)
  
# Introduction

Developing and fine-tuning web-scrappers can be a tedious, error prone task. A lot of repetetive boiler plate code is often needed, and some subtle error checking can easily be forgotten, renderering a hand written scrapper prone to unexpected failures. That where Goscrapper (**gsc** in short) gets in !

**gsc** is a domain specific language designed to write easy to understand, easy to debug, but resilient and reasonnably efficient web scrapping requests.

For instance, the following request is designed to capture up to 3 divs from google, and return an array with the text content of these divs :

````
	page = PAGE "http://www.google.fr" ;        // load google page
	SELECT "div" AS x FROM page LIMIT 3 ;       // select the divs we are interested in
    	RETURN TEXT x ;                         // return the text content for each selected div
                                                // browser is closed, everything is left in a clean state.

````
This request is available among others in the examples folder.

To experiment with this request (and debug it), a small command line utility, named **gsc**, is available :

````
$> .\bin\gsc-0.4.7.exe .\examples\google.gsc
Runtime log : initializing browser ...
[["GmailImages
Connexion
.../...
Paramètres"], ["GmailImages
Connexion"], ["GmailImages
Connexion"]]
````

As expected, you got an array of 3 elements, each containing the text content of the first 3 div found on the google home page.

Once you are confident your requests behaves as expected, you obviously want to incorporate it into a golang program. That would be the way to do it in your go code :

````go

import "github.com/xavier268/goscrapper/parser"

// .../...

// get the request source
req := `page = PAGE "http://www.google.fr" ;        // load google page
	    SELECT "div" AS x FROM page LIMIT 3 ;       // select the divs we are interested in
    	RETURN TEXT x ;                             // return the text content for each selected div
                                                    // browser is closed, everything is left in a clean state.`

// compile it, check there is no error.
comp, err := parser.Compile("google test", req)     // the request name is used for information only ...
if err != nil {                                     // check compile errors ...
   // .../...

// create an interpreter to execute the compiled request
int := NewInterpreter(context.Background())         // if the provided context is cancenned, the requests stops immediately

// provide named parameters for the requests, if needed
int = int.With(map[string]any{"password" : "myVerySecretPassword"})

// Execute, and get the result
res, err := int.Eval(comp)

// The same compiled request (comp) can then be reused with a new Interpreter and different input parameters
res2, err := NewInterpreter(context.Background()).With(map[string]any{"password" : "anotherParametr"}).Eval(comp)

````

Note that convenient *function Eval(requestName, requestSource) (any, error)* provides a shortcut when neither customization nor efficiency is critical. The above code becomes :

````go 

// compile, and evaluate in one step.
res, err := parser.Eval("myRequest", req)

````

[top](#goscrapper)

# Gsc language reference

## Request structure

A request is a list of statements. The last statement should be a single RETURN statement.
Statements are followed by a mandatory semi-colon (;).

Both block  ( /* a block comment */ ) and line comments ( // a line comment ) can appear anywhere in the request. 
Line breaks and spaces are not significant.

All symbols are case-sensitive, reserved keywords are uppercase (eg : SELECT). Some of the keywords can be written using the usual symbols ( eg : PLUS can also be written + ).

## Litteral types

Gsc can directly create litterals for the following types from the request source :

* nil, written as *NIL* or *nil*.
* boolean are written as *false* or *true*.
* numbers are the usual signed integers,
* strings,
* arrays,
* objects.

In addition, the langage itself can produce :

* time.Time (eg : time stamps),
* Page ( html page)
* Element (html Element)
* Hash (a go array, not a slice, of size md5.Size)
  
### String litterals 

String litterals follow a special syntax, to facilitate escaping. 

String litteral are either written between double quotes (") or between single quotes ('). 

In a string litteral, no caractere is ever escaped (not even "\n"), except for the same quote used to delimit the litteral string. 
Only single quotes need to be escaped in single quote litterals, and only double quotes in double quotes litterals. 

To escape a sequence of one or more quote inside a string, just add one more to the sequence. 
To represent a single quote, write 2. To escape a group of 2 quotes, write 3. 

For instance, *'In this single-quoted string, ''internal'' single quotes need to be escaped but not "double" quotes'.* 

[top](#goscrapper)

## Variables and scope

A variable name starts with a lower or upper case letter (A-Za-z), followed by zero or more letters and digits (A-Za_z0-9). No other character is allowed. A variable name may not be a gsc [keyword](#reserved-keywords).

Variable values depends on the scope. A new scope is used inside each loop. When retrieving a variable value, the interpreter attemps to return the inner most scope. Prefixing a variable with $ forces the global scope.

A variable prefixed by @ is an input parameter. Reading from a @ variable is actually the only way to declare an input parameter. An input parameter may never be assigned to. No local or glbal variable can be read from or assigned to with the same name as a named parameter.

There is no formal global/local variable declaration, but a variable must *have a chance* to be assigned to before it can be read. 
* A compile time error will happen if a request reads from a variable that had no prior *chance* of being assigned before.
* A runtime error will happen if :
  * assignement is conditionnal, actual runtime assignement may not occur before variable is used, and a runtime error will happen,
  * an input parameter is used, but no value was given to it when launching request execution, a runtime error will happen.

### Compile time checks on variables

When a source request is compiled, a variable on the left hand side :
* must have a legal name,
* must not be a known named parameter,
* is registered as a declared.

At compile time, a variable on the right hand side :
* must have a legal name,
* if it is prefixed with @ (param),
  *  is rejected if already declared as a global/local variable,
  *  is declared as an named input parameter,
* else,
  *  is rejected if not already seen on a left hand side.

### Runtime checks on variables

At runtime, a variable on the left hand side :
* may not be a known input parameter,
* can be assigned multiple times, with different values,
* is assigned in the current scope, unless the global specifier ($) is used,

At runtime, a variable on the right hand side :
* returns the specified named parameter if it is an input parameter, (@ is implicit if it was already declared as input parameter)
* returns its current assigned value, reading from the inner most scope, or,
* returns its global scoped value if prefixed with $,
* if no value can be found, return an error.

### Example 

See  : [variables.gsc](/examples/variables.gsc)

[top](#goscrapper)

## Expressions and operators

Expressions can use parenthesis () to enforce precedence.

Object members are accessed by appending a dot (.) and a key to the object. The key starts with a letter, and contains letter and digits. There are no quotes around a key. Keys are never evaluated. Accessing a non existant key returns nil.

Array elements are accessed using the usual bracket notation. Bounds are checked at runtime.

[top](#goscrapper)

### Constants

- NIL // same as nil
  
Ansi codes :

- RED
- GREEN
- YELLOW
- BLUE
- CYAN
- MAGENTA
- NORMAL

System constants :

- VERSION // Version string  : vx.y.z
- FILE_SEPARATOR // / or \, depending on os

[top](#goscrapper)

### Numerical operators

Returning a number :

- int + int
- int - int
- int * int
- int / int
- int % int // modulo

- ++ int  // int + 1
- -- int  // int - 1
- ABS int // absolute value of int
  
- LEN array // length
- LEN string // length

[top](#goscrapper)

### String operators

Returning a string :

- string + string // concatenate strings
  
- any FORMAT fmt  // format any value, using the format fmt
- RAW any // return a detailed golang string representation , see fmt.Sprintf("%#v",any)
- GO any // return a golang representation, see fmt.Sprint(any)
- JSON any // return a json representation of any
- GSC any // return a GSC representation of any
- NL // new line
    
[top](#goscrapper)

### Array and element operators

Returning an array :

- array + element // append element
- array ++ array // merge both arrays into one

[top](#goscrapper)

### Boolean operators

Returning a bool value :

- a <= b  // a and b same type, and comparable
- a < b // a and b same type, and comparable
- a > b // a and b same type, and comparable
- a >= b // a and b same type, and comparable
- a == b  // works with any value
- a != b // works with any value
  
- string CONTAINS substring // true or false
- array CONTAINS element // true or false

- bool && bool // AND
- bool || bool // OR
- ! bool // NOT
- bool != bool // XOR

[top](#goscrapper)

### Pages and elements

Expressions to create or manipulate DOM elements :

- PAGE url ; // creates a new page (tab), loading specified url. The returned expression is of type *rt.Page.
- TEXT elemnt // return the TEXT content of element as a string value
- element ATTR att // return string value of attribute att in element

[top](#goscrapper)

### Time 

- NOW // returns a time.Time object, use as a time stamp

[top](#goscrapper)

## Statements

Statements are always followed by a semi -colon.

Statements can be grouped between parenthesis.

[top](#goscrapper)

### Assignement

a = b ;     // local/global -> local
$a = b ;    // local/global -> global
a = $b ;    // global -> local
a = @c ;    // input param -> local

Assign an expression to a variable. See [variables and scope](#variables-and-scope) for more detail.

[top](#goscrapper)

### IF THEN ELSE

The traditionnal IF THEN ELSE contruct is available. IF THEN ELSE can be nested. Beware of dangling ELSE, prefer using parenthesis to group statement sequences. Expression must evaluate to a boolean value.

- IF expression THEN statement ;
- IF expression THEN  statement ELSE statement ;

[top](#goscrapper)

### RETURN statement

The last request statement must be a RETURN statement. There can be only one RETURN statement per request.

- RETURN ; // no argument, just return error status.
- RETURN a, b , c ; // return a comma separated list of expressions as arguments.

The return expression arguments are evaluated for each innermost loop and gathered into an array. If the interpreter is in [Async Mode](#sync-or-async-modes), this array is sent immediately and forgotten. If interpreter is in [Sync Mode](#sync-or-async-modes), results are aggregated into a large array, and will be returned together at the end of the request.

RETURN can be limited to only distinct values, or only to the last computed value :

- RETURN DISTINCT a,b ; // will only return [a,b] pairs which are disctincts
- RETURN LAST a,b ; // will only retrun a single [a,b] value, the last one computed.

See [examples](./examples/loopOverArray.gsc). 

[top](#goscrapper)

### FOR

Integer loops can move in either direction.

- FOR ; // start an infinite loop until the RETURN statement
- FOR  FROM intExpression TO intExpression ; // loop over ints
- FOR  FROM intExpression TO intExpression STEP intExpression ; // loop over ints, with steps

A loop variable can be specified, that will be instantiated within each loop. The loopVariable follows the same rules and scoping as a local variable assignement.

- FOR loopVariable FROM intExpression TO intExpression ; // loop over ints
- FOR loopVariable FROM intExpression TO intExpression STEP intExpression ; // loop over ints, with steps

It is also possible to loop over arrays :

- FOR loopVariable IN array ; // loop over array elements.
- FOR IN array ; // same, without loopVariable assignement.

[top](#goscrapper)

### SELECT

Selecting elements from the DOM tree should be done using a SELECT loop. Selection uses css by default, using the XPATh qaulifier uses xpath intead. 
SELECT can loop either over an entire page, or over a DOM element.

- SELECT css AS loopVariable FROM pageOrElement ;
- SELECT XPATH xpath AS loopVariable FROM pageOrElement ;

You do not ghave to declare a loopVariable :

- SELECT css FROM pageOrElement ;
- SELECT XPATH xpath FROM pageOrElement ;
  
You may limit the selected space, using LIMIT or WHERE clauses. WHERE clauses can use the loopVariable.
For instance :

- SELECT css AS loopVariable FROM pageOrElement WHERE (TEXT loopVariable) WHERE ( loopVariable ATTR href == "/") LIMIT 5 ;

SELECT is never blocking. If page loads dynamically, selects will keep track of elements already seen, and try to load more new unseen elements.

[top](#goscrapper)

## Reserved keywords

[top](#goscrapper)

# Interpreter settings

[top](#goscrapper)

## Using context

A context.Context is passed when creating an Interpreter. 

This context is check by all non trivial operations, and gsc ensures that context cancelation is handled almost immediately by any running request.

[top](#goscrapper)

## Sync or Async modes

A request running in its own interpreter is thread safe. Multiple requests can run on different interpreters, using the same compiled request tree.

Interpreters can either run in **Async** or in **Sync** Mode. 

In **Sync** Mode, the full result set is returned to the calling thread when the requests finishes execution, together with an error status. This is the default mode.

In **Async** Mode, the interpreter is passed a channel when this mode is set. During request executions, RETURN results are sent though this channel as soon as they are evaluated. When the request terminates, only the error status is return to the calling thread.
In Async Mode, if the channel is blocking because its capacity was reached, the request execution will block (but context cancelation is still monitored).


[top](#goscrapper)

# GSC command line tool

[top](#goscrapper)
