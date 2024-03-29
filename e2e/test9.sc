@url string
@css string
p1 = PAGE url
p2 = PAGE (url + "/login")
SELECT FROM PAGE url ALL css + ","+ css AS r WHERE true  LIMIT 2 + 3

RETURN url