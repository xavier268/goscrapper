// Autogenerated file. DO NOT EDIT.
// Version: 0.3.9
// Date: n/a
// Built : n/a
// (c) Xavier Gandillot 2024

package e2epack
// Generated from test2.sc

import (
	"context"
)


type Input_test2_async struct {
	a int
	b int
	a2 []int
	a1 [][]bool
	c bool
}


type Output_test2_async struct {
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
func DoAsync_test2_async(_ctx context.Context,_ch chan<- Output_test2_async,  _in Input_test2_async) (_err error) {
var _out Output_test2_async
var a int = _in.a ; _ = a
var b int = _in.b ; _ = b
var a2 []int = _in.a2 ; _ = a2
var a1 [][]bool = _in.a1 ; _ = a1
var c bool = _in.c ; _ = c
// call to incOut
 _out = Output_test2_async{}
var d int= ((((((((a) + (b))) - (a))) * (b))) / (a));_=d
//call to saveOut
_out.a2=a2
_out.c=c
_out.d=d
select {case <- _ctx.Done():return _ctx.Err();case _ch <- _out:}
 _err = _ctx.Err()
if _err != nil { return _err}
// call to incOut
 _out = Output_test2_async{}
return _err
}
