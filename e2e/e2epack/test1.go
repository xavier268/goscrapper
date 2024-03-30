// Autogenerated file. DO NOT EDIT.
// Version: 0.3.3
// Date: n/a
// Built : n/a
// (c) Xavier Gandillot 2024

package e2epack
// Generated from C:\Users\xavie\Desktop\goscrapper\e2e\test1.sc

import (
	"github.com/go-rod/rod"
	"github.com/xavier268/goscrapper/rt"
)


type Input_test1 struct {
	aaa int
	bbb bool
	ccc []bool
}


type Output_test1 struct {
	a int
	bbb bool
}


// 
// // define input variables
// @aaa int
// @bbb bool
// @ccc [bool]
// 
// 
// // open a page
// p1 = PAGE "http://www.google.fr"
// CLICK "input[name=btnK]"
// a = 23 
// b = a + 50
// c = 70 + a
// y = 23
// y2 = 2 + 3
// y3=2+3
// y4=-2+-3
// x1 = ++y
// x2=-y+ (++y)
// x3 = 22 + ++ 45  
// 
// bb = true && false
// bc = true OR false
// bd = true || false
// 
// be = 2 + 3 == 5 // same as (2+3)==5
// bf = 5 == (2 + 3 ) // without parenthesis, whould fail as : (5==2)+3
// 
// RETURN a, bbb 
func Do_test1(_in Input_test1) (_out []Output_test1, _err error) {
var aaa int = _in.aaa ; _ = aaa
var bbb bool = _in.bbb ; _ = bbb
var ccc []bool = _in.ccc ; _ = ccc
// call to incOut
 _out = append(_out, Output_test1{})
var p1 *rod.Page= rt.GetPage("http://www.google.fr");_=p1
defer rt.ClosePage(p1)
var a int= 23;_=a
var b int= ((a) + (50));_=b
var c int= ((70) + (a));_=c
var y int= 23;_=y
var y2 int= ((2) + (3));_=y2
var y3 int= ((2) + (3));_=y3
var y4 int= ((-(2)) + (-(3)));_=y4
var x1 int= (y+1);_=x1
var x2 int= ((-(y)) + (((y+1))));_=x2
var x3 int= ((22) + ((45+1)));_=x3
var bb bool= ((true) && (false));_=bb
var bc bool= ((true) || (false));_=bc
var bd bool= ((true) || (false));_=bd
var be bool= ((((2) + (3))) == (5));_=be
var bf bool= ((5) == ((((2) + (3)))));_=bf
//call to saveOut
_out[len(_out)-1].a=a
_out[len(_out)-1].bbb=bbb
// call to incOut
 _out = append(_out, Output_test1{})
return _out[:len(_out) -1], _err
}
