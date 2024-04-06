// Autogenerated file. DO NOT EDIT.
// Version: 0.3.8
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


type Input_test3_async struct {
	iii []string
	pp string
}


type Output_test3_async struct {
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
func DoAsync_test3_async(_ctx context.Context,_ch chan<- Output_test3_async,  _in Input_test3_async) (_err error) {
var _out Output_test3_async
var iii []string = _in.iii ; _ = iii
var pp string = _in.pp ; _ = pp
// call to incOut
 _out = Output_test3_async{}
var page *rod.Page= rt.GetPage(_ctx,pp);_=page
defer rt.ClosePage(page)
var b string= (("css code") + ("blablabla"));_=b
for _, i := range iii { 
 _ = i
rt.ClickFrom(".", "left", 1, page)
var c int= 45;_=c
//call to saveOut
_out.b=b
_out.c=c
select {case <- _ctx.Done():return _ctx.Err();case _ch <- _out:}
 _err = _ctx.Err()
if _err != nil { return _err}
// call to incOut
 _out = Output_test3_async{}
}
return _err
}
