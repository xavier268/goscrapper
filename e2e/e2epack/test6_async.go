// Autogenerated file. DO NOT EDIT.
// Version: 0.3.8
// Date: n/a
// Built : n/a
// (c) Xavier Gandillot 2024

package e2epack
// Generated from test6.sc

import (
	"context"
)


type Input_test6_async struct {
	i int
	ai []int
	aai [][]int
}


type Output_test6_async struct {
	x1 int
	x2 int
}


// // testing arrays
// 
// @i int
// @ai [int]
// @aai [[int]]
// 
// // x0 = [] // <== this should fail
// 
// x1 = ai[i]
// x2 = ai[ 0]
// x3 = aai[1][2]
// x4 = aai[ai[3]][ai[ai[1]]]
// 
// x5 = [1,2,3]
// x6 = ["un","deux"]
// x7 = [x5, x5]
// 
// x8 = x5 + 4
// x9 = x7 + [ 4,5,6]
// 
// x10 = x5 ++ [ 4, 5 ,6 ]
// x11 = x7 ++ x7
// x12 = x11 ++ x11 ++ x11
// 
// RETURN x1,x2
func DoAsync_test6_async(_ctx context.Context,_ch chan<- Output_test6_async,  _in Input_test6_async) (_err error) {
var _out Output_test6_async
var i int = _in.i ; _ = i
var ai []int = _in.ai ; _ = ai
var aai [][]int = _in.aai ; _ = aai
// call to incOut
 _out = Output_test6_async{}
var x1 int= (ai)[i];_=x1
var x2 int= (ai)[0];_=x2
var x3 int= ((aai)[1])[2];_=x3
var x4 int= ((aai)[(ai)[3]])[(ai)[(ai)[1]]];_=x4
var x5 []int= []int{1,2,3};_=x5
var x6 []string= []string{"un","deux"};_=x6
var x7 [][]int= [][]int{x5,x5};_=x7
var x8 []int= append(x5,4);_=x8
var x9 [][]int= append(x7,[]int{4,5,6});_=x9
var x10 []int= append(x5,[]int{4,5,6}...);_=x10
var x11 [][]int= append(x7,x7...);_=x11
var x12 [][]int= append(append(x11,x11...),x11...);_=x12
//call to saveOut
_out.x1=x1
_out.x2=x2
select {case <- _ctx.Done():return _ctx.Err();case _ch <- _out:}
 _err = _ctx.Err()
if _err != nil { return _err}
// call to incOut
 _out = Output_test6_async{}
return _err
}
