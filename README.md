# goscrapper

Domain specific language for web-scrapping.

[![Go Reference](https://pkg.go.dev/badge/github.com/xavier268/goscrapper.svg)](https://pkg.go.dev/github.com/xavier268/goscrapper) [![Go Report Card](https://goreportcard.com/badge/github.com/xavier268/goscrapper)](https://goreportcard.com/report/github.com/xavier268/goscrapper)


- [Introduction](#Introduction)
- [Language reference](#langaguage-reference)
  - [Request structure](#request-struture)
  
## Introduction

Developping and fine-tuning web-scrappers can be a tedious, error prone task. A lot of repetetive boiler plate code is often needed, and some subtle error checking can easily be forgotten, renderering a hand written scrapper prone to unexpected failures. That where Goscrapper (**gsc** in short) gets in !

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

## Gsc language reference

### Request structure

A request is a list of statements. The last statement should be a single RETURN statement.
Statements are followed by a mandatory semi-colon (;).

Both block  ( /* a block comment */ ) and line comments ( // a line comment ) can appear anywhere in the request. 
Line breaks and spaces are not significant.

All symbols are case-sensitive, reserved keywords are uppercase (eg : SELECT). Some of the keywords can be written using the usual symbols ( eg : PLUS can also be written + ).

### Litteral types

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
  
#### string litterals 

String litterals follow a special syntax, to facilitate escaping. 

String litteral are either written between double quotes (") or between single quotes ('). 

In a string litteral, no caractere is ever escaped (not even "\n"), except for the same quote used to delimit the litteral string. 
Only single quotes need to be escaped in single quote litterals, and only double quotes in double quotes litterals. 

To escape a sequence of one or more quote inside a string, just add one more to the sequence. 
To represent a single quote, write 2. To escape a group of 2 quotes, write 3. 

For instance, *'In this single-quoted string, ''internal'' single quotes need to be escaped but not "double" quotes'.* 

### Variables and scope

A variable name starts with a lower or upper case letter (A-Za-z), followed by zero or more letters and digits (A-Za_z0-9). No other character is allowed. A variable name may not be a gsc [keyword](#reserved-keywords).

There are 3 kind of variables :
* global scope variables can be read from or assigned to.
* input parameters are globla scope variables that can only be read from.
* local scope variables only exists in the current scope (eg : a loop variable in a SELECT or FOR loop). They can be read from or assign to. Local scope variables with same names can overshadow each other, or even shadow global variables. A local variable may not have the name of an input parameter.

To remove ambiguity between these kind of variable, a variable prefixed by $ is always global, and a vraiable prefixed by @ is always an input parameter. Reading from a @ variable is actually the only way to declare an input parameter.

There is no formal global/local variable declaration, but a variable must *have a chance* to be assigned to before it can be read. ( If assignement is conditionnal, actual runtime assignement may not occur before variable is used, but that is ok, value will be nil). Compiler will refuse to compile a request that read from a variable that had no prior *chance* of being assigned before.

#### Examples 

* a = "www.google.com"  // declares a as local variable, and assign to it.
* $a = 23                // assign to the global a value
* b = a                 // read either the local a, if it exists, or the global a.
* c = @ b               // declares b as an input parameter and reads from it.

### Expressions and operators

TO DO TO DO TO DO TO DO

TO DO TO DO TO DO TO DO

TO DO TO DO TO DO TO DO

TO DO TO DO TO DO TO DO


### Statements

Statements are always followed by a semi -colon.

#### Assignement

a = b ;     // local/global -> local
$a = b ;    // local/global -> global
a = $b ;    // global -> local
a = @c ;    // input param -> local

Assign an expression to a variable. See [variables and scope](#variables-and-scope).

#### RETURN statement

#### IF construct

#### FOR loops

#### SELECT dom elements

#### DOM access






### Reserved keywords