
// define input variables
@aaa int
@bbb bool
@ccc [bool]


// open a page
p1 = PAGE "http://www.google.fr"
CLICK "input[name=btnK]" FROM p1
a = 23 
b = a + 50
c = 70 + a
y = 23
y2 = 2 + 3
y3=2+3
y4=-2+-3
x1 = ++y
x2=-y+ (++y)
x3 = 22 + ++ 45  

bb = true && false
bc = true OR false
bd = true || false

be = 2 + 3 == 5 // same as (2+3)==5
bf = 5 == (2 + 3 ) // without parenthesis, whould fail as : (5==2)+3

RETURN a, bbb 