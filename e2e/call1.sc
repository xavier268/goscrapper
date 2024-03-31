@ a int
@ b int
url = "https://www.wikipedia.fr"

p = PAGE url

// ensure page is loaded
SELECT FROM p ONE "div" AS found
PRINT "Page was correctly loaded for "+ url

c = a + b

// capture five divs
SELECT FROM p ALL "div" AS divel LIMIT 5
PRINT "captured :"
PRINT divel

t = TEXT divel

RETURN c, url, t