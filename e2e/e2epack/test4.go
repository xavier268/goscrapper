// Autogenerated file. DO NOT EDIT.
// Version: 0.2.1
// Date: n/a
// Built : n/a
// (c) Xavier Gandillot 2024

package e2epack
// Generated from C:\Users\xavie\Desktop\goscrapper\e2e\test4.sc

import (
)


type Input_test4 struct {
	a string
	b []string
}


type Output_test4 struct {
	c string
	a string
	d string
}


// @a string
// @b [string]
// 
// c = a  + a
// SELECT a
//     d = a + a
//     RETURN c,a,d
func Do_test4(_in Input_test4) (_out []Output_test4, _err error) {
var a string = _in.a ; _ = a
var b []string = _in.b ; _ = b
// call to incOut
 _out = append(_out, Output_test4{})
{
var c string= (( a ) + ( a ));_=c
for true { // this will implement a loop with expr ...
var d string= (( a ) + ( a ));_=d
//call to saveOut
_out[len(_out)-1].d=d
_out[len(_out)-1].a=a
_out[len(_out)-1].c=c
// call to incOut
 _out = append(_out, Output_test4{})
}
}
return _out[:len(_out) -1], _err
}
