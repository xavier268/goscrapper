// Variables and Their Scopes
// Demonstrates how input parameters and variable scoping work in Goscrapper.

// To call this request from gsc, providing value for inp, do :
// $> gsc.exe -p "{inp:33}" .\examples\variables.gsc

// Initialize a global variable incorrectly (will cause an error because 'ii' is not yet declared as input)
// ii = 3; // ERROR: 'ii' is declared and assigned to,

// Attempt to declare 'ii' as an input parameter after assignment (incorrect usage)
// wrong = @ii; // ERROR: 'ii' being already a known variable, cannot be declared as an input parameter

// Correct usage of input parameters and global variables
$a = 100 + @inp; // 'a' is forced to global scope; 'inp' is registered as an input parameter; at runtime 'a' = 133
                 // Note: 'inp' must be set before executing the request, or a runtime error will occur:
                 //       NewInterpreter(context.Background()).With(map[string]any{"inp": 33})

// Incorrect assignment to an input parameter
// inp = 23; // ERROR: assigning to an input parameter is not allowed

// Incorrect variable usage: 'inp' is not declared as a normal variable
// b = a + inp; // ERROR: 'inp' was never declared as a "normal" variable

// Correct assignment using an input parameter
b = a + @inp; // OK: 'inp' is used correctly as an input parameter; at runtime, 'b' = 166

// Loop structure demonstrating local variable scope
FOR i FROM 1 TO 5; // 'i' is declared as a local (loop) variable
    e = b + i + @inp; // 'b' and 'i' are read from local scope, 'inp' is read from input parameter, 'e' is assigned to local scope
                      // at runtime, 'e' = 199 + i, with 'i' from 1 to 5
    RETURN e; // Returns the value of 'e' for each iteration, resulting in [[200], [201], [202], [203], [204]]
