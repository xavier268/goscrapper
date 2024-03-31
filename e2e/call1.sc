@ a int
@ b int
url = "https://www.google.fr"

p = PAGE url
c = a + b

SELECT FROM p ALL "div" AS divel LIMIT 5
t = TEXT divel

RETURN c, url, t