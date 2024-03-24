// A go scrapper test function

@toto int

@bbb bool

PAGE "http://www.google.fr"
SELECT ANY "input[name=q]"	
    CLICK "input[name=btnK]"
    a = 23 
    RETURN a