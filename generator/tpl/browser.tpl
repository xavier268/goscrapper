


import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	
)

// create a newBrowser for scrapper
func  newBrowser() (*rod.Browser) {
	u := launcher.
	New().
	UserDataDir(`{{.BrowserDataDir}}`).
	Headless({{.Headless}}).
	Set("start-maximized").
	MustLaunch()

browser := rod.New().ControlURL(u).MustConnect()

/* // TO DO - hnadle ignores by setting a hijacj router
if c.Accelerate {
	// hijack images
	router := c.browser.HijackRequests()
	router.MustAdd("*.jpg", doNotLoad)
	router.MustAdd("*.png", doNotLoad)
	//router.MustAdd("*.css", doNotLoad)
	router.MustAdd("*.woff2", doNotLoad)
	go router.Run()
}*/

return browser

}

// Get a unique browser instance.
// Creates it, if it does not exist already.
func (sc *Scrapper) Browser() *rod.Browser {

	if sc.browser == nil {
		sc.browser = newBrowser()
	}
	return sc.browser
}

// Request a new page from browser.
// Typically called upon Job creation.
// Pool pages if pool limit was set.
func (s *Scrapper) newPage() (*rod.Page) {

	// Do we need to create pool ?
	if {{.PoolSize}} >0 && s.pool == nil {
		s.pool = rod.NewPagePool({{.PoolSize}})
	}

	// func to create a page if needed
	create := func() *rod.Page {
		return s.browser.MustPage()
	}

	if s.pool != nil { // pooled
		return s.pool.Get(create)
	} else { // not pooled
		return create()
	}

}
	
	