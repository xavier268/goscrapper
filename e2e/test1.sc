// A go scrapper test function

@toto int

IGNORE "*.jpg", "*.JPEG"

PAGE "http://www.google.fr"
    SELECT "input[name=q]"	
    CLICK "input[name=btnK]"
    RETURN "input[name=btnK]"