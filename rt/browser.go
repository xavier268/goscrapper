package rt

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

// singleton instance browser for the whole program.
var (
	browser        *rod.Browser                 // browser singleton instance
	browserDataDir = mustAbs(".browserDataDir") // where browser data is kept
	headless       bool                         // headless setting
	ignorePatterns = make(map[string]bool, 5)   // patterns to ignore at browser level. These will never load for performance.
	browserLock    sync.Mutex                   // threadsafety.
)

func mustAbs(s string) string {
	ss, err := filepath.Abs(s)
	if err != nil {
		panic(err)
	}
	return ss
}

func SetHeadless(h bool) error {
	browserLock.Lock()
	defer browserLock.Unlock()
	if browser == nil {
		headless = h
		return nil
	}
	return fmt.Errorf("cannot set headless mode after browser is created")
}

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

// threadsafe construction/access to  browser singleton.
// Use this, to garantee lazy, threadsafe browser initialization.
func GetBrowser() *rod.Browser {
	browserLock.Lock()
	defer browserLock.Unlock()

	// if we already have browser, return it.
	if browser != nil {
		return browser
	}

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

// return a new page. Use empty string for empty page.
// browser is started if not already available.
// todo - think about implementing PagePool ?
func GetPage(url string) *rod.Page {
	p, err := GetBrowser().Page(proto.TargetCreateTarget{URL: url})
	if err != nil {
		Errorf("Error getting page %s : %v", url, err)
		return nil
	}
	return p
}

// close page, set page pointer to nil on success.
// // todo - think about implementing PagePool ?
func ClosePage(page *rod.Page) error {
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
