package parser

import "fmt"

func (m *myLexer) Click(where value, which value, count value) {
	if (where == value{}) || where.t != "*rod.Element" {
		m.errorf("CLICK requires a *rod.Element as the first argument, but got a %s", where.t)
	}
	if (count == value{}) {
		count = value{t: "int", v: "1"}
	}
	b := "left"
	switch which.t {
	case "LEFT", "":
		b = "left"
	case "RIGHT":
		b = "right"
	case "MIDDLE":
		b = "middle"
	default:
		m.errorf("Click requires a LEFT, RIGHT or MIDDLE button, but got a %s", which.t)
	}
	m.addImport("rt")
	m.addLines(fmt.Sprintf("rt.Click(%s, %q, %s)", where.v, b, count.v))
}

func (m *myLexer) ClickFrom(css value, which value, count value, pageOrElement value) {
	if (css == value{}) || css.t != "string" {
		m.errorf("CLICK FROM  requires a string as the first argument, but got a %s", css.t)
	}
	if (count == value{}) {
		count = value{t: "int", v: "1"}
	}
	b := "left"
	switch which.t {
	case "LEFT", "":
		b = "left"
	case "RIGHT":
		b = "right"
	case "MIDDLE":
		b = "middle"
	default:
		m.errorf("ClickFrom requires a LEFT, RIGHT or MIDDLE button, but got a %s", which.t)
	}
	if (pageOrElement.t != "*rod.Page") && pageOrElement.t != "*rod.Element" {
		m.errorf("CLICK FROM requires a *rod.Page or *rod.Element as the FROM argument, but got a %s", pageOrElement.t)
	}
	m.addImport("rt")
	m.addImport("rod")
	m.addLines(fmt.Sprintf("rt.ClickFrom(%s, %q, %s, %s)", css.v, b, count.v, pageOrElement.v))
}

func (m *myLexer) input(txt value, el value) {
	if (txt == value{}) || txt.t != "string" {
		m.errorf("INPUT requires a string as the first argument, but got a %s", txt.t)
	}
	if el.t != "*rod.Element" && el.t != "*rod.Page" {
		m.errorf("INPUT requires a *rod.Element or a *rod.Page as the second argument, but got a %s", el.t)
	}
	m.addImport("rt")
	m.addImport("rod")
	m.addLines(fmt.Sprintf("rt.Input(%s, %s)", txt.v, el.v))
}

func (m *myLexer) inputFrom(txt value, css value, pageOrElement value) {
	if (txt == value{}) || txt.t != "string" {
		m.errorf("INPUT FROM requires a string as the text argument, but got a %s", txt.t)
	}
	if (css == value{}) || css.t != "string" {
		m.errorf("INPUT FROM requires a string as the css argument, but got a %s", css.t)
	}
	if pageOrElement.t != "*rod.Page" && pageOrElement.t != "*rod.Element" {
		m.errorf("INPUT FROM requires a *rod.Page or *rod.Element as the FROM argument, but got a %s", pageOrElement.t)
	}
	m.addImport("rt")
	m.addImport("rod")
	m.addLines(fmt.Sprintf("rt.InputFrom(%s, %s, %s)", txt.v, css.v, pageOrElement.v))
}
