package rt

import "github.com/go-rod/rod"

func SelectOne(source Elementer, css string) *rod.Element {
	if source == nil {
		Errorf("cannot select an element from a nil source")
	}
	el, err := source.Element(css)
	if err != nil {
		Errorf("error trying to select element with %s : %v", css, err)
	}
	return el
}
