
// define input variables
@aaa int
@bbb bool
@ccc [bool]


// open a page
PAGE "http://www.google.fr"
CLICK "input[name=btnK]"
a = 23 
b = a + 50
c = 70 + a
y = 23
x1 = ++y
x2=--y+ (++y) // this works
// x3 = 22 + ++ 45  <<< this will not work ??
RETURN a, bbb 