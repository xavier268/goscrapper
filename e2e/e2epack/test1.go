// Autogenerated file. DO NOT EDIT.
// Version: 0.2.1
// Date: n/a
// Built : n/a
// (c) Xavier Gandillot 2024

package e2epack
// Generated from C:\Users\xavie\Desktop\goscrapper\e2e\test1.sc

import (
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
// PAGE "http://www.google.fr"
// CLICK "input[name=btnK]"
// a = 23 
// b = a + 50
// c = 70 + a
// y = 23
// RETURN a, bbb 
func Do_test1(_in Input_test1) (_out []Output_test1, _err error) {
var aaa int = _in.aaa ; _ = aaa
var bbb bool = _in.bbb ; _ = bbb
var ccc []bool = _in.ccc ; _ = ccc
// call to incOut
 _out = append(_out, Output_test1{})
{
var a int= 23;_=a
var b int= (( a ) + (50));_=b
var c int= ((70) + ( a ));_=c
var y int= 23;_=y
//call to saveOut
_out[len(_out)-1].bbb=bbb
_out[len(_out)-1].a=a
// call to incOut
 _out = append(_out, Output_test1{})
}
return _out[:len(_out) -1], _err
}
