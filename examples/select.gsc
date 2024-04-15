
	// Select specific elements in a page

	page = PAGE "http://www.wikipedia.fr" ;
	SELECT "div" AS loop FROM page LIMIT 5;    
		PRINT "**** looping ...*****" , NL, "Captured text : ",  TEXT loop ;
    	RETURN LAST "done" ;
	