# TODO list

## Arrays

allow to access elements of arrays using [] selectors.

## objects

manage 

## loops

implements more for loops (with counter) ?
implements select loop with counter ?

# functions ( eg : concat(x,y,c, ...))

define builtin functions architecture ? (takes a value as input, returns a value, and have NO SIDE EFFECT)
execution is often defered.

## commands (eg CLICK, ENTER, EXITS, ? ...)

commands have a side effect.
execution in the order where they are specified.

## a switch statement on various potential css outcome

tested successively without waiting, until the first one triggers 
````
// syntax could be :

a = SWITCH
    CASE "css1" : expression1
    CASE "css2" : expression2
    DEFAULT : expression3
    TIMEOUT number : expression4

// where expressions should all have the same type, which becomes the type of a, typically a string message ...
// DEFAULT and TIMEOUT are optionals.

````

### CSS strings

by defaults, will be queried for an array of matching elements ?
could be extended to provide
* number of matching elements ?
* text of selector