// Autogenerated file. DO NOT EDIT.
// Version: 0.3.9
// Date: n/a
// Built : n/a
// (c) Xavier Gandillot 2024

package e2epack
// Generated from call1.sc

import (
	"context"
	"fmt"
	"github.com/go-rod/rod"
	"github.com/xavier268/goscrapper/rt"
)


type Input_call1_async struct {
	a int
	b int
}


type Output_call1_async struct {
	c int
	url string
	t string
}


// @ a int
// @ b int
// url = "https://www.wikipedia.fr"
// 
// p = PAGE url
// 
// // ensure page is loaded
// SELECT FROM p ONE "div" AS found
// PRINT "Page was correctly loaded for "+ url
// 
// c = a + b
// 
// // capture five divs
// SELECT FROM p ALL "div" AS divel LIMIT 5
// PRINT "captured :"
// PRINT divel
// t = TEXT divel
// SLOW // debugging !
// 
// RETURN c, url, t
func DoAsync_call1_async(_ctx context.Context,_ch chan<- Output_call1_async,  _in Input_call1_async) (_err error) {
var _out Output_call1_async
var a int = _in.a ; _ = a
var b int = _in.b ; _ = b
// call to incOut
 _out = Output_call1_async{}
var url string= "https://www.wikipedia.fr";_=url
var p *rod.Page= rt.GetPage(_ctx,url);_=p
defer rt.ClosePage(p)
select{
case <- _ctx.Done():
 _err = _ctx.Err()
if _err != nil { return _err}
default: found := rt.SelectOne(p,"div");_=found

fmt.Println((("Page was correctly loaded for ") + (url)))
var c int= ((a) + (b));_=c
_it_001:=rt.NewSelectAllIterator(_ctx, p,"div",5); 
for divel, _ok_001 := _it_001.Next(); _ok_001;divel, _ok_001 = _it_001.Next(){_=divel;
fmt.Println("captured :")
fmt.Println(divel)
var t string= rt.GetText(divel);_=t
rt.Slow(_ctx)
//call to saveOut
_out.c=c
_out.t=t
_out.url=url
select {case <- _ctx.Done():return _ctx.Err();case _ch <- _out:}
 _err = _ctx.Err()
if _err != nil { return _err}
// call to incOut
 _out = Output_call1_async{}
}
}
return _err
}
