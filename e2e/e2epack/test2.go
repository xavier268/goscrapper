// Autogenerated file. DO NOT EDIT.
// Version: 0.2.1
// Date: n/a
// Built : n/a
// (c) Xavier Gandillot 2024

package e2epack
// Generated from C:\Users\xavie\Desktop\goscrapper\e2e\test2.sc

import (
)


type Input_test2 struct {
	a int
	b int
	a2 []int
	a1 [][]bool
	c bool
}


type Output_test2 struct {
	c bool
	d int
	a2 []int
}


// @a int
// @b int
// @a2 [int]
// @a1 [[bool]]
// @c bool
// 
// d = a+b-a*b/a
// RETURN c , d,a2
func Do_test2(_in Input_test2) (_out []Output_test2, _err error) {
var a int = _in.a ; _ = a
var b int = _in.b ; _ = b
var a2 []int = _in.a2 ; _ = a2
var a1 [][]bool = _in.a1 ; _ = a1
var c bool = _in.c ; _ = c
// call to incOut
 _out = append(_out, Output_test2{})
{
var d int= (((((((( a ) + ( b ))) - ( a ))) * ( b ))) / ( a ));_=d
//call to saveOut
_out[len(_out)-1].d=d
_out[len(_out)-1].a2=a2
_out[len(_out)-1].c=c
// call to incOut
 _out = append(_out, Output_test2{})
}
return _out[:len(_out) -1], _err
}
