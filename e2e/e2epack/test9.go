// Autogenerated file. DO NOT EDIT.
// Version: 0.3.2
// Date: n/a
// Built : n/a
// (c) Xavier Gandillot 2024

package e2epack
// Generated from C:\Users\xavie\Desktop\goscrapper\e2e\test9.sc

import (
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
// 
// RETURN url
func Do_test9(_in Input_test9) (_out []Output_test9, _err error) {
var url string = _in.url ; _ = url
var css string = _in.css ; _ = css
// call to incOut
 _out = append(_out, Output_test9{})
var p1 *rod.Page= rt.GetPage( url );_=p1
defer rt.ClosePage(p1)
var p2 *rod.Page= rt.GetPage(((( url ) + ("/login"))));_=p2
defer rt.ClosePage(p2)
_it003:=rt.NewSelectAllIterator(rt.GetPage( url ),(((( css ) + (","))) + ( css )),((2) + (3))); 
for r, _ok003 := _it003.Next(); _ok003;r, _ok003 = _it003.Next(){_=r;
if (true) {continue;}
//call to saveOut
_out[len(_out)-1].url=url
// call to incOut
 _out = append(_out, Output_test9{})
}
return _out[:len(_out) -1], _err
}
