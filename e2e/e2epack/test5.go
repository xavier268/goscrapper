// Autogenerated file. DO NOT EDIT.
// Version: 0.3.7
// Date: n/a
// Built : n/a
// (c) Xavier Gandillot 2024

package e2epack
// Generated from test5.sc

import (
	"context"
)


type Input_test5 struct {
	a int
	b []int
	bb []string
}


type Output_test5 struct {
	i int
	ii string
	x int
	y int
	z int
}


// 
// @a int 
// @b [int]
// @bb [string]
// 
// FOR i IN b
//     x = 23
// 
//     FOR ii IN bb
//         y = x + 20
//         z = 30 + x
// 
// 
// RETURN i, ii, x, y,z
// 
// 
// 
// 
// 
func Do_test5(_ctx context.Context,_in Input_test5) (_out []Output_test5, _err error) {
var a int = _in.a ; _ = a
var b []int = _in.b ; _ = b
var bb []string = _in.bb ; _ = bb
// call to incOut
 _out = append(_out, Output_test5{})
for _, i := range b { 
 _ = i
var x int= 23;_=x
for _, ii := range bb { 
 _ = ii
var y int= ((x) + (20));_=y
var z int= ((30) + (x));_=z
//call to saveOut
_out[len(_out)-1].i=i
_out[len(_out)-1].ii=ii
_out[len(_out)-1].x=x
_out[len(_out)-1].y=y
_out[len(_out)-1].z=z
 _err = _ctx.Err()
if _err != nil { return _out,_err}
// call to incOut
 _out = append(_out, Output_test5{})
}
}
return _out[:len(_out) -1], _err
}
