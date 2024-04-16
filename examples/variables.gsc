// variables and their scopes
// inp named parameter must be set before executing the request, or a runtime error will happen :
//          NewInterpreter(context.Background()).With(map[string]any{"inp": 33})

ii = 3 ;                // ERROR : ii is declared and assigned to,
// wrong = @ ii ;       // ii being alrerady a known variable, cannot be declared as an input parameter
$a = 100 + @ inp ;      // a is forced to global scope ; inp is registered as input parameter ; at runtime a = 133
// inp = 23 ;           // ERROR : assigning to an input parameter
// b = a + inp ;        // ERROR : inp was never declared as a "normal" variable
b = a + @inp ;          // OK : inp was declared as an input parameter ; at runtime, b = 166
FOR i FROM 1 TO 5 ;     // i declared as local (loop) variable
    e = b + i + @inp ;  // b and i are read from local scope, inp is read from input parameter, e is assigned to local scope
                        // at runtime, e = 199 + i, with i from 1 to 5
    RETURN e ;          // [[200], [201], [202], [203], [204]]