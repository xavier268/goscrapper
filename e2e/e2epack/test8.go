// Autogenerated file. DO NOT EDIT.
// Version: 0.3.4
// Date: n/a
// Built : n/a
// (c) Xavier Gandillot 2024

package e2epack
// Generated from C:\Users\xavie\Desktop\goscrapper\e2e\test8.sc

import (
	"context"
	"time"
)


type Input_test8 struct {
	a struct{titi bool;toto int;tutu []int}
	b struct{combo struct{titi int;toto bool};tutu []struct{tata int;tyty string}}
	c string
	d int
	bb []byte
	tt1 struct{a int;b int}
	tt2 struct{a string;b int}
}


type Output_test8 struct {
	zz time.Time
	a struct{titi bool;toto int;tutu []int}
	y4 struct{bb []byte;c string;d int}
	x4 struct{x1 struct{s string;t int};x2 struct{bb []byte;c string;d int};x3 struct{d int;s string}}
	y5 []byte
	y6 byte
}


// 
// // objects
// @a  {toto : int, titi : bool , tutu : [int]}
// @b {combo : { titi:int,toto:bool }, tutu: [{tata:int, tyty : string}] }
// @c string
// @d int
// @bb bin
// 
// @tt1 { a:string, b:int, a: int} // duplicated a should be silently ignored
// @tt2 { a:string, b:int, a: string} // duplicated a should be silently ignored
// 
// x1a = { s:c,t:d }
// x1b = { t:d, s:c} 
// x1 = x1a
// z = x1a == x1b // should be true !
// 
// x2 = {c,d, bb}
// x3 = { s:c, d}
// x4 = { x1, x2, x3}
// 
// 
// y1 = x1.s
// y2 = x2.d
// y3 = x3.d
// y4 = x4.x2
// 
// y5 = x2.bb
// y6 = x2.bb[3]
// 
// // time stamp
// zz = NOW
// 
// RETURN zz, a, y4 , x4, y5, y6
func Do_test8(_ctx context.Context,_in Input_test8) (_out []Output_test8, _err error) {
var a struct{titi bool;toto int;tutu []int} = _in.a ; _ = a
var b struct{combo struct{titi int;toto bool};tutu []struct{tata int;tyty string}} = _in.b ; _ = b
var c string = _in.c ; _ = c
var d int = _in.d ; _ = d
var bb []byte = _in.bb ; _ = bb
var tt1 struct{a int;b int} = _in.tt1 ; _ = tt1
var tt2 struct{a string;b int} = _in.tt2 ; _ = tt2
// call to incOut
 _out = append(_out, Output_test8{})
var x1a struct{s string;t int}= struct{s string;t int}{c,d};_=x1a
var x1b struct{s string;t int}= struct{s string;t int}{c,d};_=x1b
var x1 struct{s string;t int}= x1a;_=x1
var z bool= ((x1a) == (x1b));_=z
var x2 struct{bb []byte;c string;d int}= struct{bb []byte;c string;d int}{bb,c,d};_=x2
var x3 struct{d int;s string}= struct{d int;s string}{d,c};_=x3
var x4 struct{x1 struct{s string;t int};x2 struct{bb []byte;c string;d int};x3 struct{d int;s string}}= struct{x1 struct{s string;t int};x2 struct{bb []byte;c string;d int};x3 struct{d int;s string}}{x1,x2,x3};_=x4
var y1 string= x1.s;_=y1
var y2 int= x2.d;_=y2
var y3 int= x3.d;_=y3
var y4 struct{bb []byte;c string;d int}= x4.x2;_=y4
var y5 []byte= x2.bb;_=y5
var y6 byte= (x2.bb)[3];_=y6
var zz time.Time= time.Now();_=zz
//call to saveOut
_out[len(_out)-1].a=a
_out[len(_out)-1].x4=x4
_out[len(_out)-1].y4=y4
_out[len(_out)-1].y5=y5
_out[len(_out)-1].y6=y6
_out[len(_out)-1].zz=zz
if _err = _ctx.Err() ; _err != nil { return _out,_err}
// call to incOut
 _out = append(_out, Output_test8{})
return _out[:len(_out) -1], _err
}
