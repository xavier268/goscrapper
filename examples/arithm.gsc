// Basic loop arithmetics

$a = 0 ; // global var to accumulate
FOR i FROM 1 TO 5 ;    
    $a = a + i ;    // a gives the same value as $a, since no local a is defined.
    a = i ;         // this local variable reflects the current loop value
    RETURN a , $a ; // get the current loop index and the cumultaive sum