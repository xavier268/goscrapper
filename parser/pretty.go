package parser

import (
	"fmt"
	"strings"
)

// Get a pretty string representation of provided object.
func Pretty(a interface{}) string {

	const step = 4 // step for indentation

	s := fmt.Sprintf("%#v", a)

	res := new(strings.Builder)

	ident := 0
	first := true
	for _, c := range s {
		switch c {
		case '{':
			ident += step
			res.WriteString("\n" + strings.Repeat(" ", ident))
			res.WriteRune(c)
			res.WriteString("\n" + strings.Repeat(" ", ident))
			first = true

		case '}':

			res.WriteString("\n" + strings.Repeat(" ", ident))
			res.WriteRune(c)
			ident -= step
			res.WriteString("\n" + strings.Repeat(" ", ident))
			first = true
		case ' ': // ignore leading blanks
			if !first {
				res.WriteRune(c)
			}
		case ';', ',':
			res.WriteRune(c)
			res.WriteString("\n" + strings.Repeat(" ", ident))
			first = true

		default:
			first = false
			res.WriteRune(c)
		}
	}

	return res.String()
}
