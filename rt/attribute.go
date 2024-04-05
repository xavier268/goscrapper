package rt

import "github.com/go-rod/rod"

func Attribute(el *rod.Element, att string) string {
	a, err := el.Attribute(att)
	if err != nil || a == nil {
		return ""
	} else {
		return *a
	}
}
