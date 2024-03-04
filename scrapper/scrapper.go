package scrapper

import (
	"context"
	"log"
	"sync"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/xavier268/goscrapper/config"
)

// Scrapper is the container that will be running the scrapper.
// There is one and only one browser assicited with each scrapper.
// Scrapper Name is used to save & reuse contexte (cookies, etc ... ) from previous runs.
type Scrapper struct {

	// actual physical state
	browser *rod.Browser
	// current page pools
	pool    rod.PagePool // shared pool
	poolinc rod.PagePool // incognito pool
	// map names to communication buses implementations
	buses map[string]*Bus

	// Sync
	wg     sync.WaitGroup     // to make sure we wait for all separate threads.
	ctx    context.Context    // global context is only cancellable.
	cancel context.CancelFunc // cancel function

	// configuration
	conf *config.Configuration
}

func NewScrapper(conf *config.Configuration) *Scrapper {

	// apply default options
	s := new(Scrapper)

	// context
	s.ctx, s.cancel = context.WithCancel(context.Background())
	// buses
	s.buses = make(map[string]*Bus)

	if conf == nil {
		panic("cannot start scrapper without a configuration")
	}
	s.conf = conf

	// create and launch browser
	u := launcher.New().
		Headless(conf.Headless).
		UserDataDir(conf.BrowserDataDir).
		Set("start-maximized").
		MustLaunch()

	s.browser = rod.New().ControlURL(u).MustConnect()

	// implement hijack do not load at browser level
	if len(conf.Ignore) > 0 {
		router := s.browser.HijackRequests()
		for _, patt := range conf.Ignore {
			// hijack all specified resources
			router.MustAdd(patt,
				func(ctx *rod.Hijack) {
					ctx.Response.Fail(proto.NetworkErrorReasonBlockedByClient)
				})
		}
	}

	// default page pool to 10 if not specified
	if conf.PagePool != 0 {
		s.pool = rod.NewPagePool(conf.PagePool)
	} else {
		panic("shared PagePool should never be 0")
	}
	if conf.PagePoolIncognito != 0 {
		s.poolinc = rod.NewPagePool(conf.PagePoolIncognito)
	}
	return s
}

// Get a new (shared) tab, loading the specified url in it.
func (s *Scrapper) GetPage(url ...string) *rod.Page {
	return s.pool.Get(func() *rod.Page { return s.browser.MustPage(url...) })
}

// Get a new incognito tab, loading the specific url in it.
func (s *Scrapper) GetIncognitoPage(url ...string) *rod.Page {
	if s.poolinc != nil {
		return s.poolinc.Get(func() *rod.Page { return s.browser.MustIncognito().MustPage(url...) })
	} else {
		// no pooling
		return s.browser.MustIncognito().MustPage(url...)
	}
}

// return a shared page not needed anymore
func (s *Scrapper) PutPage(page *rod.Page) {
	s.pool.Put(page)
}

// return an incognito page not needed anymore
func (s *Scrapper) PutIncognitoPage(page *rod.Page) {
	if s.poolinc == nil { // if no poolincognito, just close and forget.
		if page != nil {
			page.MustClose()
			return
		}
	}
	s.poolinc.Put(page) // otherwise, return to pool.
}

// Close everything, releasing all resources.
func (s *Scrapper) Close() {
	if s == nil {
		return
	}

	s.cancel()  // cancel global context
	s.wg.Wait() // block for threads to finish

	// close buses
	for _, b := range s.buses {
		if b != nil {
			close(b.ch)
		}
	}
	// close browser
	if err := s.browser.Close(); err != nil {
		log.Printf("error while closing scrapper %s: %v", s.conf.Name, err)
	}
}

// create and run a new job in a separate thread.
func (s *Scrapper) Run(startstate string) {
	j := &Job{
		s:     s,
		state: startstate,
		ctx:   s.ctx,
	}
	s.wg.Add(1)
	go j.run()
}
