package rt

import (
	"fmt"

	"github.com/go-rod/rod"
)

// singleton, used to store runtime context
var rt *MyRuntime

// auto init the runtime context on first use.
func init() {
	fmt.Println("init runtime")
	rt = newRuntime()
}

// the main runtime context object
type MyRuntime struct {
	browser *rod.Browser
}

func newRuntime() *MyRuntime {
	// TODO
	return &MyRuntime{
		browser: &rod.Browser{},
	}
}

func Ignore(patterns ...string) {
	_ = rt
	// TODO, position a capture on the browser in rt ...
}
