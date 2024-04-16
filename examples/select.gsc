// Select Specific Elements from a Web Page

// Load the Wikipedia France home page into 'page'
page = PAGE "http://www.wikipedia.fr";

// Begin a loop to iterate over the first five <div> elements found on the page
SELECT "div" AS loop FROM page LIMIT 5;    
    // Print a separator and the text content of the current <div> element
    PRINT "**** Looping... *****", NL, "Captured text: ", TEXT loop;

	// After completing the loop, return "done" to indicate successful completion
	RETURN LAST "done";
