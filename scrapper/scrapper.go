package scrapper

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/xavier268/goscrapper/config"
)

// Scrapper is the container that will be running the scrapper.
// There is one and only one browser assicited with each scrapper.
// Scrapper Name is used to save & reuse contexte (cookies, etc ... ) from previous runs.
type Scrapper struct {
	Name string

	// actual physical state
	browser *rod.Browser
	// element *rod.Element

	// configuration default
	headless       bool
	doNotLoad      []string // prevent loading from listed patterns (typically:  *.jpg *.png, etc ...)
	startTime      time.Time
	debug          int    // default to app level debug value
	rootDir        string // main directory where everything happens
	browserDataDir string // where session data is stored
}

func NewScrapper(options ...ScrapperOption) *Scrapper {
	var err error

	// apply default options
	s := new(Scrapper)
	s.Name = "noname"
	s.startTime = time.Now()
	s.rootDir, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	s.browserDataDir = filepath.Join(s.rootDir, ".browserdata-"+s.Name)
	s.headless = false
	s.doNotLoad = []string{}
	s.debug = config.DEBUG

	// apply provided options
	for _, opt := range options {
		if s.debug > 0 {
			log.Println(opt)
		}
		opt.apply(s)
	}

	// create and launch browser
	u := launcher.New().
		Headless(s.headless).
		UserDataDir(s.browserDataDir).
		Set("start-maximized").
		MustLaunch()

	s.browser = rod.New().ControlURL(u).MustConnect()

	// hijack do not load ...
	if len(s.doNotLoad) > 0 {
		router := s.browser.HijackRequests()
		for _, patt := range s.doNotLoad {
			// hijack all specified resources
			router.MustAdd(patt,
				func(ctx *rod.Hijack) {
					ctx.Response.Fail(proto.NetworkErrorReasonBlockedByClient)
				})
		}
	}

	return s
}

// Set the Scrapper to run in the background.
func (s *Scrapper) SetBackground(b bool) {
	panic("not implemented")
}

// Close Scrapper, releasing all resources.
func (s *Scrapper) Close() {
	if s == nil {
		return
	}
	if err := s.browser.Close(); err != nil {
		log.Printf("error while closing scrapper %s: %v", s.Name, err)
	}
}
