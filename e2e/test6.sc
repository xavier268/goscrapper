// testing arrays

@i int
@ai [int]
@aai [[int]]

// x0 = [] // <== this should fail

x1 = ai[i]
x2 = ai[ 0]
x3 = aai[1][2]
x4 = aai[ai[3]][ai[ai[1]]]

x5 = [1,2,3]
x6 = ["un","deux"]
x7 = [x5, x5]

x8 = x5 + 4
x9 = x7 + [ 4,5,6]

x10 = x5 ++ [ 4, 5 ,6 ]
x11 = x7 ++ x7
x12 = x11 ++ x11 ++ x11

RETURN x1,x2