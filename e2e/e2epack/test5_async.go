// Autogenerated file. DO NOT EDIT.
// Version: 0.3.4
// Date: n/a
// Built : n/a
// (c) Xavier Gandillot 2024

package e2epack
// Generated from C:\Users\xavie\Desktop\goscrapper\e2e\test5.sc

import (
	"context"
)


type Input_test5_async struct {
	a int
	b []int
	bb []string
}


type Output_test5_async struct {
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
func DoAsync_test5_async(_ctx context.Context,_ch chan<- Output_test5_async,  _in Input_test5_async) (_err error) {
var _out Output_test5_async
var a int = _in.a ; _ = a
var b []int = _in.b ; _ = b
var bb []string = _in.bb ; _ = bb
// call to incOut
 _out = Output_test5_async{}
for _, i := range b { 
 _ = i
var x int= 23;_=x
for _, ii := range bb { 
 _ = ii
var y int= ((x) + (20));_=y
var z int= ((30) + (x));_=z
//call to saveOut
_out.i=i
_out.ii=ii
_out.x=x
_out.y=y
_out.z=z
select {case <- _ctx.Done():return _ctx.Err();case _ch <- _out:}
if _err = _ctx.Err() ; _err != nil { return _err}
// call to incOut
 _out = Output_test5_async{}
}
}
return _err
}
