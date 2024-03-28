// Autogenerated file. DO NOT EDIT.
// Version: 0.3.2
// Date: n/a
// Built : n/a
// (c) Xavier Gandillot 2024

package e2epack
// Generated from C:\Users\xavie\Desktop\goscrapper\e2e\test3.sc

import (
	"github.com/xavier268/goscrapper/rt"
	"github.com/go-rod/rod"
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
// PAGE pp
// b = "css code" + "blablabla"
// 
// FOR i  IN iii// this starts a loop ... 
//     CLICK "."
//     c = 45
//     RETURN b,c
func Do_test3(_in Input_test3) (_out []Output_test3, _err error) {
var _page *rod.Page;_=_page
var _select *rod.Element;_=_select
var iii []string = _in.iii ; _ = iii
var pp string = _in.pp ; _ = pp
// call to incOut
 _out = append(_out, Output_test3{})
if _page == nil { 
_page,_err=rt.GetPage( pp )
defer rt.ClosePage(_page)
}else{
_err=_page.Navigate( pp )
}
if _err !=nil {
return _out, _err
}
_select=nil // reset selection within page
var b string= (("css code") + ("blablabla"));_=b
for _, i := range  iii  { 
 _ = i
var c int= 45;_=c
//call to saveOut
_out[len(_out)-1].b=b
_out[len(_out)-1].c=c
// call to incOut
 _out = append(_out, Output_test3{})
}
return _out[:len(_out) -1], _err
}
