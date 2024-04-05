// Autogenerated file. DO NOT EDIT.
// Version: 0.3.5
// Date: n/a
// Built : n/a
// (c) Xavier Gandillot 2024

package e2epack
// Generated from test3.sc

import (
	"context"
	"github.com/go-rod/rod"
	"github.com/xavier268/goscrapper/rt"
)


type Input_test3 struct {
	iii []string
	pp string
}


type Output_test3 struct {
	b string
	c int
}


// // test simple loop
// 
// @iii [string]
// @pp string
// 
// page = PAGE pp
// b = "css code" + "blablabla"
// 
// FOR i  IN iii// this starts a loop ... 
//     CLICK "." FROM page
//     c = 45
//     RETURN b,c
func Do_test3(_ctx context.Context,_in Input_test3) (_out []Output_test3, _err error) {
var iii []string = _in.iii ; _ = iii
var pp string = _in.pp ; _ = pp
// call to incOut
 _out = append(_out, Output_test3{})
var page *rod.Page= rt.GetPage(_ctx,pp);_=page
defer rt.ClosePage(page)
var b string= (("css code") + ("blablabla"));_=b
for _, i := range iii { 
 _ = i
var c int= 45;_=c
//call to saveOut
_out[len(_out)-1].b=b
_out[len(_out)-1].c=c
if _err = _ctx.Err() ; _err != nil { return _out,_err}
// call to incOut
 _out = append(_out, Output_test3{})
}
return _out[:len(_out) -1], _err
}
