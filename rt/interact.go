package rt

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

// Click on an element. which is one of "left", "right" or "middle"
func Click(el *rod.Element, which proto.InputMouseButton, count int) {
	if el != nil {
		switch which {
		case proto.InputMouseButtonLeft,
			proto.InputMouseButtonRight,
			proto.InputMouseButtonMiddle:
			el.Click(which, count)
		default:
			Errorf("Click: unknown button %s", which)
		}
	}
}

// Select an element and click it. which is one of "left", "right" or "middle"
// Do nothing if no element found.
func ClickFrom(css string, which proto.InputMouseButton, count int, pageOrElement Elementer) {
	if pageOrElement != nil {
		els, err := pageOrElement.Elements(css)
		if err == nil && len(els) > 0 {
			Click(els[0], which, count)
		} else {
			Errorf("Could not find a clickable element %s : %s", css, err)
		}
	}
}

// Input a txt in an lement, after selecting and focusing on it.
func Input(el *rod.Element, txt string) {
	if el != nil {
		err := el.SelectAllText()
		if err == nil {
			el.Input(txt)
		}
	}
}

// Select an element and input a txt in it.
func InputFrom(css string, txt string, pageOrElement Elementer) {
	if pageOrElement != nil {
		els, err := pageOrElement.Elements(css)
		if err == nil && len(els) > 0 {
			Input(els[0], txt)
		} else {
			Errorf("Could not find a input element %s : %s", css, err)
		}
	}
}
