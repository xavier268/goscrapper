package rt

func Exists(page Elementer, css string) bool {
	elts, err := page.Elements(css)
	if err != nil || len(elts) == 0 {
		return false
	}
	return false
}
