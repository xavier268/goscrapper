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


type Input_test7 struct {
	a []byte
	a1 [][]byte
}


type Output_test7 struct {
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
func Do_test7(_ctx context.Context,_in Input_test7) (_out []Output_test7, _err error) {
var a []byte = _in.a ; _ = a
var a1 [][]byte = _in.a1 ; _ = a1
// call to incOut
 _out = append(_out, Output_test7{})
var x0 [][]byte= a1;_=x0
var x1 [][]byte= append(a1,a);_=x1
var x2 [][]byte= append(a1,a1...);_=x2
var x3 []byte= (a1)[0];_=x3
var x4 byte= (a)[0];_=x4
//call to saveOut
_out[len(_out)-1].a=a
_out[len(_out)-1].a1=a1
if _err = _ctx.Err() ; _err != nil { return _out,_err}
// call to incOut
 _out = append(_out, Output_test7{})
return _out[:len(_out) -1], _err
}
