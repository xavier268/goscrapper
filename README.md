# goscrapper

High level tool for generating scrapping helper functions in go.

## How to use this tool 

### 1.Install the gsc tool
````
$> go install github.com/xavier268/goscrapper/gsc@latest

# verify, by calling
$> gsc
````

### 2.Write a small script, describing what should be fetched from a web page
   
*See the e2e test package for examples*

### 3. Compile your script into a go function

Example, to compile call1.sc into the mypackage package :
````
$> gsc -o mypackage call1.sc
````

There are many options to compile the script.

In particular, it can be compiled either :
* in synchroneous mode (loop as requested, and return an array of structs for each loop)
* in asynchroneous mode (loop as requested, each capture value is emitted to the provided channel, function return when finished)

### 4. Use the generated function(s) from you code
   
The function is contained in a file with the same name as the script. This files also contains the structs for input/output.
Functions are provided a context.Context to allow for extarnal cancelation or timeout.



## Langage reference


### Script structure

A script starts with the declaration of the input variables, then a list of statements, then a unique return statement.

Some statements are loop statements, that will generate a result for each loop. See FOR ... and SELECT ...
Loop statements are always nested. Results are always captured from the inner most loop.

The go compiler will eliminate redundant code generated to keep the linter happy.

Line and block comments are available, using // and /* or */. Whitespace and returns are not significant, except when they separate a symbol name (eg <= is not the same as < =).

### Keywords

All symbols are case-sensitive, reserved keywords are uppercase.

### Variables

Available types are *int*, *bool*, *string*, or *bin*. The *bin* type is internally an array of bytes. Types can be combined to create objects, between braces {}, or arrays, between brakets []. Array members must have the same type. Object members types can be different.

**Input variables** are declared at the start of the script, using @.

**Output variables** are listed in the last return statement.

Variables can be upper or lower case. Of course, only upper case variables will be available outside of the package.
Variable names should start with a letter, and continue with digits or letters. Underscore is not accepted (reserved for internal variables).

**A variable can only be assigned once**. A variable that is not declared cannot be used.

It costs very little to create as many intermediate variables as needed. Those that are not output variables will be discarded when function returns. 

Number are decimal, signed, integers.

Strings use double or single quotes. Inside a string, only quotes are escaped by adding one more quotes (see examples).

Litteral and array objects are defined using braces or brackets. Keys for objects should be valid identifier (A-Za-z, then A-Za-z0-9). Object keys should never be quoted.

Booleans are either *true* or *false*.

Objects members are accessed using the dot (.) operator, followed by a valid key name.

Array members are accessed using the bracket operatoir and an integer expression as index.


### No argument operators

* **NOW** *returns a timeStamp of type time.Time* 
* **true false** *are in **lower** case, return the corresponding bool values*

### Unary operators**

* **-** *changes sign of integer.*
* **++**  *increments an integer.*
* **LOWER UPPER** *change case of string expression.*
* **NOT** or **!** *applies to bool expressions.*
* **PAGE** *applies to a string expression, that represents an URL. It opens a new page (tab) with this url. The url can be empty, in which case a blank page is open. It returns a page object. A page object has an internal ytype, and cannot be returned.*
* **TEXT** *applies to an Element (typically obtained as the loop variable of a SELECT statement) or a Page object (obatined with PAGE).*

### Binary operators**

* **+ - * / %** *are the usual operations on integers.*
* **+** *also applies to string, for concatenation.*
* **++** *applies to arrays of same type, and merge them.*
* **&& ||** *apply to booleans, for AND and OR.*
* **CONTAINS** *applies to string expressions, returning a boolean : exp1 CONTAINS exp2.*
* **== !=** *are equality/inequality operators, retruning a boolean. Apply to any type.*
* **< > <= >=** *are usual compaison operators for integers, returning a boolean.*
* *element **ATTRIBUTE** stringExpr returns a string with the value of the requested attribute for the Element.*

### Expressions

Expressions are created by applying the above operators to other expressions.

The following are also valid expressions :
* ( expr )
* expr [indexExpression] // to get an array element
* expr . key // to get an object member value
* [ expr1, exp2, ... ] to create an array from same type expressions
* { key1 : expr1, key2 : expr2, ...} to create an object from pairs of keys and expressions.

### Statements

*(See grammar for details)*

* **variable = expression** : *Assign an expression to a variable. The variable is declared, and its type derived from the expression.*
  
* **PRINT expression** : *Prints expression to stdout.*
  
* **CLICK element** : *assume LEFT click, once*.
* **CLICK element LEFT count**
* **CLICK element RIGHT count**
* **CLICK element MIDDLE count** : *Click on selected element, using the specified button, for the count times.*
  
* **CLICK css FROM pageOrElement** : *assume LEFT click, once*.
* **CLICK element LEFT count FROM pageOrElement**
* **CLICK element RIGHT count FROM pageOrElement**
* **CLICK element MIDDLE count FROM pageOrElement** : *Same as above, but first select the element from the pageOrElement using the provided css selector.*
  
* **INPUT text IN element** : *Input Text in corresponding element. Previous text is first selected, then cleared.*
* **INPUT text IN css FROM pageOrElement** : *Same as above, but will first select the page using the css selector.*
  
* **FOR identifier IN arrayExpression** : *Usual for loop over an arrayExpression. The identifier declares the loop variable, that will sucessively take all the values from the array.*
  
* **SELECT FROM pageOrElement ALL cssExpression AS loopVariable WHERE boolExpression LIMIT integerExpression** : *This is a loop statement. Loop statements can be nested. It collects all available Element from the pageOrElement provided, using the cssExpression as the selector. The loopVariable should not be already declared. For each collected Element, it only keep those matching the (optionnal) WHERE condition, and up to the number specified in the optional LIMIT. Setting LIMIT to 0 means no limit. It is possible that no element are collected.* **The loopVariable is set to a *rod.Element.**

* **SELECT FROM pageOrElement ONE cssExpression AS loopVariable** : *This is a loop statement. Loop statements can be nested. It collects just one Element avaible for the pageOrElement provided, using the cssExpression as the selector. ONE elemnet is always collected. Loop variable is* **set to a rod.Element**. *Statement will wait until timeout or one element becomes available on page.*
  
* **SELECT FROM pageOrElement ANY AS loopVariable CASE css1 : expr1 ; CASE css2 : expr2 ; DEFAULT   : expr3 ;** : *This is a loop statement. Loop statements can be nested. It expects an element to match ANY of the css selectors. It will loop until there is a match, unless a default is provided. When a match is found,* **the loopVariable is assigned the value of the corresponding expression, NOT the matching rod.Element !.** *The expression may not reference the matched element. Note the required semicolon at the end of each expression*

* **RETURN var1, var2, ...** :
*There should be EXACTLY ONE return statement, with at least ONE variable. Return variables should have been declared and assigned to. Expressions are not allowed in return list. Only types that can be provided as input are accepted as return (int, bool, string, bin, arrays and objects). Pages, Elements, and other internal types cannot be returned.*


