# Goscrapper

A domain-specific language for web scraping.

[![Go Reference](https://pkg.go.dev/badge/github.com/xavier268/goscrapper.svg)](https://pkg.go.dev/github.com/xavier268/goscrapper) [![Go Report Card](https://goreportcard.com/badge/github.com/xavier268/goscrapper)](https://goreportcard.com/report/github.com/xavier268/goscrapper)

## Table of Contents
- [Goscrapper](#goscrapper)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [GSC Language Reference](#gsc-language-reference)
    - [Request Structure](#request-structure)
    - [Literal Types](#literal-types)
      - [String Literals](#string-literals)
    - [Variables and Scope](#variables-and-scope)
      - [Compile-Time Checks on Variables](#compile-time-checks-on-variables)
      - [Runtime Checks on Variables](#runtime-checks-on-variables)
      - [Example](#example)
    - [Expressions and Operators](#expressions-and-operators)
      - [Constants](#constants)
      - [Numerical Operators](#numerical-operators)
      - [String Operators](#string-operators)
      - [Array and Element Operators](#array-and-element-operators)
      - [Boolean Operators](#boolean-operators)
      - [Pages and Elements](#pages-and-elements)
      - [Time](#time)
    - [Statements](#statements)
      - [Assignment](#assignment)
      - [IF THEN ELSE](#if-then-else)
      - [RETURN Statement](#return-statement)
      - [FOR](#for)
      - [SELECT](#select)
  - [Reserved Keywords](#reserved-keywords)
  - [Interpreter Settings](#interpreter-settings)
    - [Using context](#using-context)
    - [Sync or Async Modes](#sync-or-async-modes)
  - [GSC Command Line Tool](#gsc-command-line-tool)

## Introduction

Developing and fine-tuning web scrapers can be a tedious and error-prone task. A lot of repetitive boilerplate code is often needed, and some subtle error checking can easily be overlooked, rendering a handwritten scraper prone to unexpected failures. This is where Goscrapper (**gsc**) comes in!

**gsc** is a domain-specific language designed to write easy-to-understand, easy-to-debug, but resilient and reasonably efficient web scraping requests.

For instance, the following request is designed to capture up to three `<div>` elements from Google and return an array with the text content of these `<div>` elements:

````
page = PAGE "http://www.google.fr";        // Load Google page
SELECT "div" AS x FROM page LIMIT 3;       // Select the divs we are interested in
RETURN TEXT x;                             // Return the text content for each selected div
                                           // Browser is closed, everything is left in a clean state.
````

This request is available among others in the examples folder.

To experiment with this request (and debug it), a small command line utility, named **gsc**, is available:

````
$> .\bin\gsc-0.4.7.exe .\examples\google.gsc
Runtime log: initializing browser...
[["GmailImages
Connexion
.../...
Paramètres"], ["GmailImages
Connexion"], ["GmailImages
Connexion"]]
````

As expected, you got an array of three elements, each containing the text content of the first three `<div>` found on the Google homepage.

Once you are confident your requests behave as expected, you obviously want to incorporate them into a Golang program. Here's how to do it in your Go code:

```go
import "github.com/xavier268/goscrapper/parser"

// get the request source
req := `page = PAGE "http://www.google.fr";        // Load Google page
SELECT "div" AS x FROM page LIMIT 3;               // Select the divs we are interested in
RETURN TEXT x;                                     // Return the text content for each selected div
                                                   // Browser is closed, everything is left in a clean state.`

// Compile it, check there are no errors.
comp, err := parser.Compile("google test", req)     // The request name is used for information only...
if err != nil {                                     // Check compile errors...
    // handle error
}

// Create an interpreter to execute the compiled request
int := NewInterpreter(context.Background())         // If the provided context is cancelled, the request stops immediately

// Provide named parameters for the requests, if needed
int = int.With(map[string]any{"password": "myVerySecretPassword"})

// Execute, and get the result
res, err := int.Eval(comp)

// The same compiled request (comp) can then be reused with a new Interpreter and different input parameters
res2, err := NewInterpreter(context.Background()).With(map[string]any{"password": "anotherParameter"}).Eval(comp)
```

Note that the convenient function `Eval(requestName, requestSource) (any, error)` provides a shortcut when neither customization nor efficiency is critical. The above code becomes:

```go
// Compile, and evaluate in one step.
res, err := parser.Eval("myRequest", req)
```

[Back to top](#goscrapper)

## GSC Language Reference

### Request Structure

A request is a list of statements. The last statement should be a single RETURN statement.
Statements are followed by a mandatory semicolon (`;`).

Both block (`/* a block comment */`) and line comments (`// a line comment`) can appear anywhere in the request. Line breaks and spaces are not significant.

All

 symbols are case-sensitive, reserved keywords are uppercase (e.g., `SELECT`). Some of the keywords can be written using usual symbols (e.g., `PLUS` can also be written as `+`).

### Literal Types

Gsc can directly create literals for the following types from the request source:

- `nil`, written as `NIL` or `nil`.
- Booleans are written as `false` or `true`.
- Numbers are the usual signed integers.
- Strings, arrays, objects.

In addition, the language itself can produce:

- `time.Time` (e.g., timestamps),
- `Page` (HTML page),
- `Element` (HTML element),
- `Hash` (a Go array, not a slice, of size `md5.Size`).

#### String Literals

String literals follow a special syntax to facilitate escaping.

String literals are either written between double quotes (`"`) or single quotes (`'`).

In a string literal, no character is ever escaped (not even "\n"), except for the same quote used to delimit the literal string.
Only single quotes need to be escaped in single-quote literals, and only double quotes in double-quote literals.

To escape a sequence of one or more quotes inside a string, just add one more to the sequence.
For instance, `'In this single-quoted string, ''internal'' single quotes need to be escaped but not "double" quotes'.`

[Back to top](#goscrapper)

### Variables and Scope

A variable name starts with a lower or upper case letter (`A-Z`, `a-z`), followed by zero or more letters and digits (`A-Za-z0-9`). No other character is allowed. A variable name may not be a [reserved keyword](#reserved-keywords).

Variable values depend on the scope. A new scope is used inside each loop. When retrieving a variable value, the interpreter attempts to return the innermost scope. Prefixing a variable with `$` forces the global scope.

A variable prefixed by `@` is an input parameter. Reading from an `@` variable is the only way to declare an input parameter. An input parameter may never be assigned to. No local or global variable can be read from or assigned to with the same name as a named parameter.

There is no formal global/local variable declaration, but a variable must *have a chance* to be assigned to before it can be read. 
* A compile-time error will occur if a request reads from a variable that had no prior *chance* of being assigned.
* A runtime error will occur if:
  * Assignment is conditional, actual runtime assignment may not occur before the variable is used, and a runtime error will occur,
  * An input parameter is used, but no value was given to it when launching request execution, a runtime error will occur.

#### Compile-Time Checks on Variables

When a source request is compiled, a variable on the left-hand side:
* Must have a legal name,
* Must not be a known named parameter,
* Is registered as declared.

At compile time, a variable on the right-hand side:
* Must have a legal name,
* If it is prefixed with `@` (parameter),
  * Is rejected if already declared as a global/local variable,
  * Is declared as a named input parameter,
* Else,
  * Is rejected if not already seen on a left-hand side.

#### Runtime Checks on Variables

At runtime, a variable on the left-hand side:
* May not be a known input parameter,
* Can be assigned multiple times, with different values,
* Is assigned in the current scope, unless the global specifier (`$`) is used,

At runtime, a variable on the right-hand side:
* Returns the specified named parameter if it is an input parameter, (`@` is implicit if it was already declared as input parameter),
* Returns its current assigned value, reading from the innermost scope, or,
* Returns its global scoped value if prefixed with `$`,
* If no value can be found, returns an error.

#### Example

See: [variables.gsc](/examples/variables.gsc)

[Back to top](#goscrapper)

### Expressions and Operators

Expressions can use parentheses (`()`) to enforce precedence.

Object members are accessed by appending a dot (`.`) and a key to the object. The key starts with a letter, and contains letters and digits. There are no quotes around a key. Keys are never evaluated. Accessing a non-existent key returns `nil`.

Array elements are accessed using the usual bracket notation. Bounds are checked at runtime.

#### Constants

- `NIL` // same as `nil`
  
Ansi codes:

- `RED`
- `GREEN`
- `YELLOW`
- `BLUE`
- `CYAN`
- `MAGENTA`
- `NORMAL`

System constants:

- `VERSION` // Version string: `vx.y.z`
- `FILE_SEPARATOR` // `/` or `\`, depending on OS

#### Numerical Operators

Returning a number:

- `int + int`
- `int - int`
- `int * int`
- `int / int`
- `int % int` // modulo

- `++int`  // `int + 1`
- `--int`  // `int - 1`
- `ABS int` // absolute value of `int`
  
- `LEN array` // length of array
- `LEN string` // length of string

#### String Operators

Returning a string:

- `string + string` // concatenate strings
  
- `any FORMAT fmt`  // format any value, using the format `fmt`
- `RAW any` // return a detailed Golang string representation, see `fmt.Sprintf("%#v", any)`
- `GO any` // return a Golang representation, see `fmt.Sprint(any)`
- `JSON any` // return a JSON representation of any
- `GSC any` // return a GSC representation of any
- `NL` // new line
    
#### Array and Element Operators

Returning an array:

- `array + element` // append element to array
- `array ++ array` // merge both arrays into one

#### Boolean Operators

Returning a bool value:

- `a <= b`  // `a` and `b` same type, and comparable
- `a < b` // `a` and `b` same type, and comparable
- `a > b` // `a` and `b` same type, and comparable
- `a >= b` // `a` and `b` same type, and comparable
- `a == b`  // works with any value
- `a != b` // works with any value
  
- `string CONTAINS substring` // true or false
- `array CONTAINS element` // true or false

- `bool && bool` // AND
- `bool || bool` // OR
- `!bool` // NOT
- `bool != bool` // XOR

#### Pages and Elements

Expressions to create or manipulate DOM elements:

- `PAGE url` ; // creates a new page (tab), loading specified URL. The returned expression is of type `*rt.Page`.
- `TEXT element` // return the TEXT content of element as a string value
- `element ATTR att` // return string value of attribute `att` in element

#### Time 

- `NOW` // returns a `time.Time` object, use as a timestamp

[Back to top](#goscrapper)

### Statements

Statements are always followed by a semicolon (`;`).

Statements can be grouped between parentheses.

#### Assignment

- `a = b;`     // local/global -> local
- `$a = b;`    // local/global -> global
- `a = $b;`    // global -> local
- `a = @c;`    // input param -> local

Assign an expression to a variable. See [Variables and Scope](#variables-and-scope) for more detail.

#### IF THEN ELSE

The traditional IF THEN ELSE construct is available. IF THEN ELSE can be nested. Beware of dangling ELSE, prefer using parentheses to group statement sequences. The expression must evaluate to a boolean value.

- `IF expression THEN statement;`
- `IF expression THEN statement ELSE statement;`

#### RETURN Statement

The last request statement must be a RETURN statement. There can be only one RETURN statement per request.

- `RETURN;` // no argument, just return error status.
- `RETURN a, b, c;` // return a comma-separated list of expressions as arguments.

The return expression arguments are evaluated for each innermost loop and gathered into an array. If the interpreter is in [Async Mode](#sync-or-async-modes), this array is sent immediately and forgotten. If the interpreter is in [Sync Mode](#sync-or-async-modes), results are aggregated into a large array, and will be returned together at the end of the request.

RETURN can be limited to only distinct values, or only to the last computed value:

- `RETURN DISTINCT a, b;` // will only return `[a, b]` pairs which are distinct
- `RETURN LAST a, b;` // will only return a single `[a, b]` value, the last one computed.

See [examples](./examples/loopOverArray.gsc). 

#### FOR

Integer loops can move in either direction.

- `FOR;` // start an infinite loop until the RETURN statement
- `FOR FROM intExpression TO intExpression;` // loop over ints
- `FOR FROM intExpression TO intExpression STEP intExpression;` // loop over ints, with steps

A loop variable can be specified, that will be instantiated within each loop. The loopVariable follows the same rules and scoping as a local variable assignment.

- `FOR loopVariable FROM intExpression TO int

Expression;` // loop over ints
- `FOR loopVariable FROM intExpression TO intExpression STEP intExpression;` // loop over ints, with steps

It is also possible to loop over arrays:

- `FOR loopVariable IN array;` // loop over array elements.
- `FOR IN array;` // same, without loopVariable assignment.

#### SELECT

Selecting elements from the DOM tree should be done using a SELECT loop. Selection uses CSS by default, using the XPath qualifier uses XPath instead. SELECT can loop either over an entire page, or over a DOM element.

- `SELECT css AS loopVariable FROM pageOrElement;`
- `SELECT XPATH xpath AS loopVariable FROM pageOrElement;`

You do not have to declare a loopVariable:

- `SELECT css FROM pageOrElement;`
- `SELECT XPATH xpath FROM pageOrElement;`

You may limit the selected space, using LIMIT or WHERE clauses. WHERE clauses can use the loopVariable. For instance:

- `SELECT css AS loopVariable FROM pageOrElement WHERE (TEXT loopVariable) WHERE (loopVariable ATTR href == "/") LIMIT 5;`

SELECT is never blocking. If the page loads dynamically, selects will keep track of elements already seen, and try to load more new unseen elements.

[Back to top](#goscrapper)

## Reserved Keywords

Reserved keywords are words that are part of the syntax of the language itself and cannot be used as identifiers (e.g., variable names, function names, etc.).

[Back to top](#goscrapper)

## Interpreter Settings

### Using context

A `context.Context` is passed when creating an Interpreter.

This context is checked by all non-trivial operations, and gsc ensures that context cancellation is handled almost immediately by any running request.

### Sync or Async Modes

A request running in its own interpreter is thread-safe. Multiple requests can run on different interpreters, using the same compiled request tree.

Interpreters can either run in **Async** or **Sync** Mode.

In **Sync** Mode, the full result set is returned to the calling thread when the request finishes execution, together with an error status. This is the default mode.

In **Async** Mode, the interpreter is passed a channel when this mode is set. During request executions, RETURN results are sent through this channel as soon as they are evaluated. When the request terminates, only the error status is returned to the calling thread. In Async Mode, if the channel is blocking because its capacity was reached, the request execution will block (but context cancellation is still monitored).

[Back to top](#goscrapper)

## GSC Command Line Tool

The `gsc` command line tool is provided to help developers test and debug their gsc scripts easily from the command line.

[Back to top](#goscrapper)
