// Autogenerated file. DO NOT EDIT.
// Version: 0.3.4
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


type Input_test9 struct {
	url string
	css string
}


type Output_test9 struct {
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
// 
// RETURN url
func Do_test9(_ctx context.Context,_in Input_test9) (_out []Output_test9, _err error) {
var url string = _in.url ; _ = url
var css string = _in.css ; _ = css
// call to incOut
 _out = append(_out, Output_test9{})
var p1 *rod.Page= rt.GetPage(url);_=p1
defer rt.ClosePage(p1)
var p2 *rod.Page= rt.GetPage((((url) + ("/login"))));_=p2
defer rt.ClosePage(p2)
_it004:=rt.NewSelectAllIterator(_ctx, rt.GetPage(url),((((css) + (","))) + (css)),((2) + (3))); 
for r, _ok004 := _it004.Next(); _ok004;r, _ok004 = _it004.Next(){_=r;
if (true) {continue;}
select{
case <- _ctx.Done():
if _err = _ctx.Err() ; _err != nil { return _out,_err}
default: el1 := rt.SelectOne(p1,css);_=el1

var toto string= (("hello") + ("world"));_=toto
var titi bool= (((el1) == (el1)));_=titi
select{
case <- _ctx.Done():
if _err = _ctx.Err() ; _err != nil { return _out,_err}
default: el2 := rt.SelectOne(p1,css);_=el2

var p3 int= ((23) + (4));_=p3
var tutu bool= (((el1) == (el2)));_=tutu
//call to saveOut
_out[len(_out)-1].url=url
if _err = _ctx.Err() ; _err != nil { return _out,_err}
// call to incOut
 _out = append(_out, Output_test9{})
}
}
}
return _out[:len(_out) -1], _err
}
