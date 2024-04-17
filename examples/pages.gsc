// Opening and closing pages.

p1 = PAGE "http://www.google.com" ;
p2 = PAGE "http://www.wikipedia.fr" ;

PRINT "There are " , LEN PAGES , " tabs opened" ;

PRINT "Closing  tab : " , p1;
CLOSE p1 ;

PRINT "There is " , LEN PAGES , " tab opened" ;

PRINT "Opening 2 more pages" ;
p3 = PAGE "http://amazon.fr" ;
p4 = PAGE "http://github.com" ;

FOR pp IN PAGES ;
    PRINT pp ;
    RETURN LAST ;