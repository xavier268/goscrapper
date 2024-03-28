package parser

import "fmt"

func (m *myLexer) vPage(exp value) {

	if exp.t != "string" {
		m.errorf("page expects a string argument, but got %s", exp.t)
	}
	m.imports["github.com/xavier268/goscrapper/rt"] = true
	m.imports["github.com/go-rod/rod"] = true
	m.lateDecl["var _page *rod.Page;_=_page"] = true
	m.lateDecl["var _select *rod.Element;_=_select"] = true
	li := fmt.Sprintf(`if _page == nil { 
_page,_err=rt.GetPage(%s)
defer rt.ClosePage(_page)
}else{
_err=_page.Navigate(%s)
}
if _err !=nil {
return _out, _err
}
_select=nil // reset selection within page`,
		exp.v, exp.v)
	m.addLines(li)
}
