// Autogenerated file. DO NOT EDIT.
// Version: 0.3.4
// Date: n/a
// Built : n/a
// (c) Xavier Gandillot 2024

package e2epack
// Generated from C:\Users\xavie\Desktop\goscrapper\e2e\test7.sc

import (
	"context"
)


type Input_test7_async struct {
	a []byte
	a1 [][]byte
}


type Output_test7_async struct {
	a []byte
	a1 [][]byte
}


// // test binary typeDefinition
// 
// @ a bin
// @ a1 [bin]
// 
// x0 = a1
// x1 = a1 + a
// x2 = a1 ++ a1
// x3 = a1[0]
// x4 = a[0]
// 
// RETURN a , a1
// 
func DoAsync_test7_async(_ctx context.Context,_ch chan<- Output_test7_async,  _in Input_test7_async) (_err error) {
var _out Output_test7_async
var a []byte = _in.a ; _ = a
var a1 [][]byte = _in.a1 ; _ = a1
// call to incOut
 _out = Output_test7_async{}
var x0 [][]byte= a1;_=x0
var x1 [][]byte= append(a1,a);_=x1
var x2 [][]byte= append(a1,a1...);_=x2
var x3 []byte= (a1)[0];_=x3
var x4 byte= (a)[0];_=x4
//call to saveOut
_out.a=a
_out.a1=a1
select {case <- _ctx.Done():return _err;case _ch <- _out:}
if _err = _ctx.Err() ; _err != nil { return _err}
// call to incOut
 _out = Output_test7_async{}
return _err
}