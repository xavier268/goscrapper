package rt

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

// singleton instance browser for all calling requests.
var (
	browser        *Browser                     // browser singleton instance
	browserDataDir = mustAbs(".browserDataDir") // where browser data is kept
	headless       bool                         // headless setting
	ignorePatterns = make(map[string]bool, 5)   // patterns to ignore at browser level. These will never load for performance.
	browserLock    sync.Mutex                   // threadsafety lock
)

func mustAbs(s string) string {
	ss, err := filepath.Abs(s)
	if err != nil {
		panic(err)
	}
	return ss
}

// Elementer can provide Elements or Element. Typically, a page or an element.
type Elementer interface {
	Elements(string) (Elements, error)
	ElementsX(string) (Elements, error)
	Element(string) (*Element, error)
	ElementX(string) (*Element, error)
}

// Set browser in headless mode.
// Browser should not have started already.
func SetHeadless(h bool) error {
	browserLock.Lock()
	defer browserLock.Unlock()
	if browser == nil {
		headless = h
		return nil
	}
	return fmt.Errorf("cannot set headless mode after browser is created")
}

// Which requests patterns will always be ignored. Eg : ".jpeg", ".png", ...
func SetIgnore(patt ...string) error {
	browserLock.Lock()
	defer browserLock.Unlock()
	if browser == nil {
		for _, p := range patt {
			ignorePatterns[p] = true
		}
		return nil
	}
	return fmt.Errorf("cannot set ignore list after browser is created")
}

// The directory to save brawser internal states into (cookies, etc ...)
func SetBrowserDataDir(dir string) (err error) {
	dir, err = filepath.Abs(dir)
	if err != nil {
		return err
	}
	browserLock.Lock()
	defer browserLock.Unlock()
	if browser == nil {
		browserDataDir = dir
		err := os.MkdirAll(dir, 0755)
		return err
	}
	return fmt.Errorf("cannot set browser data directory after browser is created")
}

// Threadsafe and lazy access to the browser singleton.
func GetBrowser() *Browser {
	browserLock.Lock()
	defer browserLock.Unlock()

	// if we already have browser, return it.
	if browser != nil {
		return browser
	}

	Logf("initializing browser ...")

	// else, create it
	u := launcher.
		New().
		UserDataDir(browserDataDir).
		Headless(headless).
		Set("start-maximized").
		MustLaunch()

	browser = rod.New().ControlURL(u).MustConnect()

	// install hijack router if needed
	if len(ignorePatterns) != 0 {
		// hijack requested patterns
		browserRouter := browser.HijackRequests()
		for patt := range ignorePatterns {
			browserRouter.MustAdd(patt,
				func(ctx *rod.Hijack) {
					ctx.Response.Fail(proto.NetworkErrorReasonBlockedByClient)
				})
		}
		go browserRouter.Run()
	}
	return browser
}

// Get a new page. Use empty string for empty page.
// Browser is started if not already available.
func GetPage(ctx context.Context, url string) *Page {
	p, err := GetBrowser().Page(proto.TargetCreateTarget{URL: url})
	if err != nil {
		Errorf("Error getting page %s : %v", url, err)
		return nil
	}
	err = p.WaitLoad()
	if err != nil {
		Errorf("Error while waiting for page %s to load : %v", url, err)
	}
	p = p.Context(ctx)
	return p
}

// Get the array of opened pages.
// Note : returns a GSC array ([]any), not a []*rod.Page array.
func Pages() ([]any, error) {
	pp, err := GetBrowser().Pages()
	if err != nil {
		return nil, err
	}
	res := make([]any, 0, len(pp))
	for _, p := range pp {
		res = append(res, p)
	}
	return res, nil
}

// close page, set page pointer to nil on success.
func ClosePage(page *Page) error {
	browserLock.Lock()
	defer browserLock.Unlock()
	if page == nil {
		return nil
	}
	err := page.Close()
	if err != nil {
		return err
	}
	page = nil
	return nil
}

// Retrieve text from a *rod.Element.
// Return empty string if not found.
func GetText(el *Element) string {
	if el == nil {
		return ""
	}
	s, err := el.Text()
	if err != nil {
		Errorf("cannot retrieve text from element")
	}
	return s
}

// Retrieve attribute from element.
// Return empty string on error.
func GetAttribute(el *Element, att string) string {
	a, err := el.Attribute(att)
	if err != nil || a == nil {
		return ""
	} else {
		return *a
	}
}

// Click on an element. Use which to choose from "left", "right" or "middle" (default left).
// Use count to specify number of clicks (defaults 1).
func Click(el *Element, which InputMouseButton, count int) {
	if which == "" {
		which = proto.InputMouseButtonLeft
	}
	if count <= 0 {
		count = 1
	}
	if el != nil {
		switch which {
		case proto.InputMouseButtonLeft,
			proto.InputMouseButtonRight,
			proto.InputMouseButtonMiddle:
			err := el.Click(which, count)
			if err != nil {
				Errorf(err.Error())
			}
		default:
			Errorf("Click: unknown button %s", which)
		}
	}
}

// Input a txt in an element, after selecting and focusing on it.
func Input(el *Element, txt string) {
	if el != nil {
		err := el.SelectAllText()
		if err == nil {
			err = el.Input(txt)
			if err != nil {
				Errorf(err.Error())
			}
		} else {
			Errorf(err.Error())
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

// Return string content of the html element in page, or "" if error.
// Never abort, never wait.
func GetPageText(page *Page) string {
	if page == nil {
		return ""
	}
	els, err := page.Elements("html")
	if err != nil || len(els) < 1 {
		return ""
	}
	return GetElemText(els[0])
}

// Get text for element. Never wait, never fail. Return "" if error.
func GetElemText(elem *Element) string {
	if elem == nil {
		return ""
	}
	tt, err := elem.Text()
	if err != nil {
		return ""
	}
	return tt
}
