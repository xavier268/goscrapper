package parser

import (
	"fmt"
	"strings"
)

// start a loop, saving the output and adding the opening {
func (m *myLexer) forNameInExpression(name string, expr value) {

	if !strings.HasPrefix(expr.t, "[]") {
		m.errorf("FOR name IN expression requires expression to be an array, but %s is not an array; it is a %s", expr.v, expr.t)
	}

	for _, k := range m.inparams {
		if k == name {
			m.errorf("you cannot allocate the loop variable to an input parameter %s", name)
		}
	}
	if _, ok := m.vars[name]; ok {
		m.errorf("the variable %s already exists and cannot be used as a loop variable", name)
	}

	m.vars[name] = expr.t[2:] // set type of loop variable.

	li := fmt.Sprintf("for _, %s := range %s { \n _ = %s", name, expr.v, name)
	m.addLines(li)

}

func (m *myLexer) selectAll(opt selopt) {

	// check types
	if (opt.from.t != "*rod.Page") && (opt.from.t != "*rod.Element") {
		m.errorf("cannot select all from type %s , expected *rod.Page or *rod.Element.", opt.from.t)
	}
	if opt.css.t != "string" {
		m.errorf("css selector should be a string, but got a %s", opt.css.t)
	}
	if opt.limit.v != "" && opt.limit.t != "int" {
		m.errorf("when a limit is set, an int is expected, but got a %s", opt.limit.t)
	}

	// verify loop variable correctly set
	if typ := m.vars[opt.loopv]; typ != "*rod.Element" {
		m.errorf("internal error with the loop variable %s not correctly prepared : %s", opt.loopv, typ)
	}

	// generate code
	m.imports["github.com/xavier268/goscrapper/rt"] = true
	uid := m.uid()
	lim := "0"
	if opt.limit.v != "" {
		lim = opt.limit.v
	}
	li1 := fmt.Sprintf("_it%s:=rt.NewSelectAllIterator(_ctx, %s,%s,%s); ", uid, opt.from.v, opt.css.v, lim)
	li2 := fmt.Sprintf("for %s, _ok%s := _it%s.Next(); _ok%s;%s, _ok%s = _it%s.Next(){_=%s;", opt.loopv, uid, uid, uid, opt.loopv, uid, uid, opt.loopv)
	m.addLines(li1, li2)
	for _, w := range opt.where {
		// check where clause type
		if w.t != "bool" {
			m.errorf("cannot accept where clause that is not a bool value but a %s", w.t)
			continue
		}
		li := fmt.Sprintf("if !(%s) {continue;}// where clause checks", w.v)
		m.addLines(li)
	}
}

func (m *myLexer) selectOne(source value, css value, id value) {

	// check types
	if (source.t != "*rod.Page") && (source.t != "*rod.Element") {
		m.errorf("cannot select one from type %s , expected *rod.Page or *rod.Element.", source.t)
	}
	if css.t != "string" {
		m.errorf("css selector should be a string, but got a %s", css.t)
	}
	if id.t != "IDENTIFIER" {
		m.errorf("invalid identifier %s for a select one variable", id.v)
	}

	// generate code
	m.imports["github.com/xavier268/goscrapper/rt"] = true
	m.addLines("select{\ncase <- _ctx.Done():")
	m.checkContext()
	m.addLines(fmt.Sprintf("default: %s := rt.SelectOne(%s,%s);_=%s\n", id.v, source.v, css.v, id.v))

}

func (m *myLexer) selectAny(vpage value, vid value, vcases []casopt) { // page_or_element, loopvarId, case list

	if vpage.t != "*rod.Page" && vpage.t != "*rod.Element" {
		m.errorf("cannot select any from type %s, expected *rod.Page.", vpage.t)
		return
	}
	if vid.t != "IDENTIFIER" {
		m.errorf("invalid identifier %s for a select any variable", vid.v)
		return
	}
	if _, ok := lx.vars[vid.v]; !ok {
		m.errorf("internal error : select any variable should have been defined already")
		return
	}
	if len(vcases) == 0 {
		m.errorf("no cases to select from")
		return
	}
	extyp := "" // expression type should be identical for all expressions.
	for _, c := range vcases {
		if !c.def && c.e1.t != "string" {
			m.errorf("cannot accept a case selector that is not a string, but found a %s", c.e1.t)
			return
		}
		if extyp == "" {
			extyp = c.e2.t
		} else {
			if extyp != c.e2.t {
				m.errorf("cannot accept a case result expression  that is not of the same type as the previous case result expression, but found a %s and %s", extyp, c.e2.t)
				return
			}
		}
	}
	if extyp == "" {
		m.errorf("no cases expression result specified")
		return
	}
	lx.vars[vid.v] = extyp // set type of loop variable.
	vid.t = extyp

	m.imports["github.com/xavier268/goscrapper/rt"] = true
	// declare loop variable
	m.addLines("{") // create a "pseudo-loop" block
	m.addLines(fmt.Sprintf("var %s %s;_=%s\n", vid.v, vid.t, vid.v))
	m.addLines("for {")
	m.checkContext()
	for _, c := range vcases {
		if c.def {
			m.addLines(fmt.Sprintf(" %s=%s;_=%s;break\n", vid.v, c.e2.v, vid.v))
		} else {
			m.addLines(fmt.Sprintf("if( rt.Exists(%s,%s)){ %s=%s;_=%s;break; }\n", vpage.v, c.e1.v, vid.v, c.e2.v, vid.v))
		}
	}
	m.addLines("}//for")
}
