// Autogenerated file. DO NOT EDIT.
// Version: 0.3.5
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


type Input_call1 struct {
	a int
	b int
}


type Output_call1 struct {
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
// 
// t = TEXT divel
// 
// RETURN c, url, t
func Do_call1(_ctx context.Context,_in Input_call1) (_out []Output_call1, _err error) {
var a int = _in.a ; _ = a
var b int = _in.b ; _ = b
// call to incOut
 _out = append(_out, Output_call1{})
var url string= "https://www.wikipedia.fr";_=url
var p *rod.Page= rt.GetPage(url);_=p
defer rt.ClosePage(p)
select{
case <- _ctx.Done():
if _err = _ctx.Err() ; _err != nil { return _out,_err}
default: found := rt.SelectOne(p,"div");_=found

fmt.Println((("Page was correctly loaded for ") + (url)))
var c int= ((a) + (b));_=c
_it002:=rt.NewSelectAllIterator(_ctx, p,"div",5); 
for divel, _ok002 := _it002.Next(); _ok002;divel, _ok002 = _it002.Next(){_=divel;
fmt.Println("captured :")
fmt.Println(divel)
var t string= rt.GetText(divel);_=t
//call to saveOut
_out[len(_out)-1].c=c
_out[len(_out)-1].t=t
_out[len(_out)-1].url=url
if _err = _ctx.Err() ; _err != nil { return _out,_err}
// call to incOut
 _out = append(_out, Output_call1{})
}
}
return _out[:len(_out) -1], _err
}
