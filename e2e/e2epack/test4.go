// Autogenerated file. DO NOT EDIT.
// Version: 0.3.4
// Date: n/a
// Built : n/a
// (c) Xavier Gandillot 2024

package e2epack
// Generated from C:\Users\xavie\Desktop\goscrapper\e2e\test4.sc

import (
	"context"
	"github.com/go-rod/rod"
	"github.com/xavier268/goscrapper/rt"
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
// // z = 2 // uncommenting this should fail because loop variable z would be redeclared 
// 
// 
// c = a  + a
// page = PAGE "google.com"
// 
// SELECT FROM page ALL a AS z WHERE 2+4==1 LIMIT 2+66 WHERE z == z
//     d = a + a
//     RETURN c,a,d
func Do_test4(_ctx context.Context,_in Input_test4) (_out []Output_test4, _err error) {
var a string = _in.a ; _ = a
var b []string = _in.b ; _ = b
// call to incOut
 _out = append(_out, Output_test4{})
var c string= ((a) + (a));_=c
var page *rod.Page= rt.GetPage("google.com");_=page
defer rt.ClosePage(page)
_it002:=rt.NewSelectAllIterator(_ctx, page,a,((2) + (66))); 
for z, _ok002 := _it002.Next(); _ok002;z, _ok002 = _it002.Next(){_=z;
if (((((2) + (4))) == (1))) {continue;}
if (((z) == (z))) {continue;}
var d string= ((a) + (a));_=d
//call to saveOut
_out[len(_out)-1].a=a
_out[len(_out)-1].c=c
_out[len(_out)-1].d=d
if _err := _ctx.Err(); _err != nil { return _out, _err ;}
// call to incOut
 _out = append(_out, Output_test4{})
}
return _out[:len(_out) -1], _err
}
