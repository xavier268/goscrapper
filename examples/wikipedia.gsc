
// Download and Print the Text from Wikipedia's Home Page

// Load the home page of Wikipedia into variable 'p'
p = PAGE "http://www.wikipedia.fr"; 

// Extract all text from the page 'p' and store it in variable 'a'
a = TEXT p;

// Return 'a'
RETURN a;
