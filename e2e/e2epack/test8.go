// Autogenerated file. DO NOT EDIT.
// Version: 0.3.2
// Date: n/a
// Built : n/a
// (c) Xavier Gandillot 2024

package e2epack
// Generated from C:\Users\xavie\Desktop\goscrapper\e2e\test8.sc

import (
	"time"
)


type Input_test8 struct {
	a struct{toto int;titi bool;tutu []int}
	b struct{combo struct{titi int;toto bool};tutu []struct{tata int;tyty string}}
	c string
	d int
	bb []byte
}


type Output_test8 struct {
	a struct{toto int;titi bool;tutu []int}
	y4 struct{d int;bb []byte;c string}
	x4 struct{x1 struct{s string;t int};x2 struct{d int;bb []byte;c string};x3 struct{s string;d int}}
	y5 []byte
	y6 byte
	zz time.Time
}


// 
// // objects
// @a  {toto : int, titi : bool , tutu : [int]}
// @b {combo : { titi:int,toto:bool }, tutu: [{tata:int, tyty : string}] }
// @c string
// @d int
// @bb bin
// 
// x1 = { s:c,t:d }
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
// RETURN a, y4 , x4, y5, y6,zz
func Do_test8(_in Input_test8) (_out []Output_test8, _err error) {
var a struct{toto int;titi bool;tutu []int} = _in.a ; _ = a
var b struct{combo struct{titi int;toto bool};tutu []struct{tata int;tyty string}} = _in.b ; _ = b
var c string = _in.c ; _ = c
var d int = _in.d ; _ = d
var bb []byte = _in.bb ; _ = bb
// call to incOut
 _out = append(_out, Output_test8{})
var x1 struct{s string;t int}= struct{s string;t int}{ c , d };_=x1
var x2 struct{d int;bb []byte;c string}= struct{d int;bb []byte;c string}{ d , bb , c };_=x2
var x3 struct{s string;d int}= struct{s string;d int}{ c , d };_=x3
var x4 struct{x1 struct{s string;t int};x2 struct{d int;bb []byte;c string};x3 struct{s string;d int}}= struct{x1 struct{s string;t int};x2 struct{d int;bb []byte;c string};x3 struct{s string;d int}}{ x1 , x2 , x3 };_=x4
var y1 string=  x1 .s;_=y1
var y2 int=  x2 .d;_=y2
var y3 int=  x3 .d;_=y3
var y4 struct{d int;bb []byte;c string}=  x4 .x2;_=y4
var y5 []byte=  x2 .bb;_=y5
var y6 byte= ( x2 .bb)[3];_=y6
var zz time.Time= time.Now();_=zz
//call to saveOut
_out[len(_out)-1].a=a
_out[len(_out)-1].x4=x4
_out[len(_out)-1].y4=y4
_out[len(_out)-1].y5=y5
_out[len(_out)-1].y6=y6
_out[len(_out)-1].zz=zz
// call to incOut
 _out = append(_out, Output_test8{})
return _out[:len(_out) -1], _err
}
