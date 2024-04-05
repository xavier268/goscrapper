// Autogenerated file. DO NOT EDIT.
// Version: 0.3.6
// Date: n/a
// Built : n/a
// (c) Xavier Gandillot 2024

package e2epack
// Generated from test9.sc

import (
	"context"
	"github.com/go-rod/rod"
	"github.com/xavier268/goscrapper/rt"
)


type Input_test9_async struct {
	url string
	css string
}


type Output_test9_async struct {
	url string
}


// @url string
// @css string
// p1 = PAGE url
// p2 = PAGE (url + "/login")
// SELECT FROM PAGE url ALL css + ","+ css AS r WHERE true  LIMIT 2 + 3
//     SELECT FROM p1 ONE css AS el1
//         toto = "hello" + "world"
//         titi = ( el1 == el1)
//         SELECT FROM p1 ONE css AS el2
//             p3 = 23+4
//             tutu = (el1 == el2)
//             href = el2 ATTRIBUTE "href"
//             SELECT FROM p1 ANY AS el3
//                 CASE "html" : "an html value";
//                 CASE "div" : "a div was found";
//                 DEFAULT : "none of the above";
// 
// RETURN url
func DoAsync_test9_async(_ctx context.Context,_ch chan<- Output_test9_async,  _in Input_test9_async) (_err error) {
var _out Output_test9_async
var url string = _in.url ; _ = url
var css string = _in.css ; _ = css
// call to incOut
 _out = Output_test9_async{}
var p1 *rod.Page= rt.GetPage(_ctx,url);_=p1
defer rt.ClosePage(p1)
var p2 *rod.Page= rt.GetPage(_ctx,(((url) + ("/login"))));_=p2
defer rt.ClosePage(p2)
_it_001:=rt.NewSelectAllIterator(_ctx, rt.GetPage(_ctx,url),((((css) + (","))) + (css)),((2) + (3))); 
for r, _ok_001 := _it_001.Next(); _ok_001;r, _ok_001 = _it_001.Next(){_=r;
if !(true) {continue;}// where clause checks
select{
case <- _ctx.Done():
if _err = _ctx.Err() ; _err != nil { return _err}
default: el1 := rt.SelectOne(p1,css);_=el1

var toto string= (("hello") + ("world"));_=toto
var titi bool= (((el1) == (el1)));_=titi
select{
case <- _ctx.Done():
if _err = _ctx.Err() ; _err != nil { return _err}
default: el2 := rt.SelectOne(p1,css);_=el2

var p3 int= ((23) + (4));_=p3
var tutu bool= (((el1) == (el2)));_=tutu
var href string= (rt.Attribute(el2,"href"));_=href
{
var el3 string;_=el3

for {
if _err = _ctx.Err() ; _err != nil { return _err}
if( rt.Exists(p1,"html")){ el3="an html value";_=el3;break; }

if( rt.Exists(p1,"div")){ el3="a div was found";_=el3;break; }

 el3="none of the above";_=el3;break

}//for
//call to saveOut
_out.url=url
select {case <- _ctx.Done():return _ctx.Err();case _ch <- _out:}
if _err = _ctx.Err() ; _err != nil { return _err}
// call to incOut
 _out = Output_test9_async{}
}
}
}
}
return _err
}
