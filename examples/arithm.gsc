// Basic Loop Arithmetic

$a = 0; // Initialize a global variable to accumulate the sum

// Loop from 1 to 5 inclusive
FOR i FROM 1 TO 5;    
    $a = $a + i; // Update global sum with the current loop index i
    a = i;       // Set local variable 'a' to the current loop index i

    // Return the current loop index and the cumulative sum at each iteration
    RETURN a, $a;