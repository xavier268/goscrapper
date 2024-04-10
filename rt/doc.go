// browser related globals and utilities.
package rt

import "github.com/go-rod/rod"

// defining type alias to decouple from actual driver.
type (
	Page     = rod.Page
	Browser  = rod.Browser
	Element  = rod.Element
	Elements = rod.Elements
)
