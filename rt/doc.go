// browser related globals and utilities.
package rt

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

// defining type alias to decouple from actual driver.
type (
	Page             = rod.Page
	Browser          = rod.Browser
	Element          = rod.Element
	Elements         = rod.Elements
	InputMouseButton = proto.InputMouseButton
)

const (
	MouseLeft   = InputMouseButton(proto.InputMouseButtonLeft)
	MouseRight  = InputMouseButton(proto.InputMouseButtonRight)
	MouseMiddle = InputMouseButton(proto.InputMouseButtonMiddle)
)

const VERSION = "0.1.0"

func init() {
	fmt.Println("Runtime version : " + VERSION)
}
