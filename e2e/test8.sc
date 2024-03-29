
// objects
@a  {toto : int, titi : bool , tutu : [int]}
@b {combo : { titi:int,toto:bool }, tutu: [{tata:int, tyty : string}] }
@c string
@d int
@bb bin

@tt1 { a:string, b:int, a: int} // duplicated a should be silently ignored
@tt2 { a:string, b:int, a: string} // duplicated a should be silently ignored

x1a = { s:c,t:d }
x1b = { t:d, s:c} 
x1 = x1a
z = x1a == x1b // should be true !

x2 = {c,d, bb}
x3 = { s:c, d}
x4 = { x1, x2, x3}


y1 = x1.s
y2 = x2.d
y3 = x3.d
y4 = x4.x2

y5 = x2.bb
y6 = x2.bb[3]

// time stamp
zz = NOW

RETURN zz, a, y4 , x4, y5, y6