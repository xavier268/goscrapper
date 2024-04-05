@url string
@css string
p1 = PAGE url
p2 = PAGE (url + "/login")
SELECT FROM PAGE url ALL css + ","+ css AS r WHERE true  LIMIT 2 + 3
    SELECT FROM p1 ONE css AS el1
        toto = "hello" + "world"
        titi = ( el1 == el1)
        SELECT FROM p1 ONE css AS el2
            p3 = 23+4
            tutu = (el1 == el2)
            href = el2 ATTRIBUTE "href"
            SELECT FROM p1 ANY AS el3
                CASE "html" : "an html value";
                CASE "div" : "a div was found";
                DEFAULT : "none of the above";

RETURN url