

import (
	"context"
	"sync"
	"time"
	"fmt"
	"github.com/go-rod/rod"
)

// Scrapper defines the top level object for scrapping.
// There is typically only one Scrapper instance at a time.
type Scrapper struct {
	headless bool               // should we run headless
	ignore   map[string]bool    // set of requests patterns we hould we ignore ( eg : *.png, *.jpg, ...)
	ctx      context.Context    // global context
	cancel   context.CancelFunc // cancel to cancel everything
	wg       sync.WaitGroup     // to monitor pending threads
	browser  *rod.Browser		// the browser attached to this crapper. Unique for the entire app.
	pool 	rod.PagePool		// page pagePooling
	// TO DO
}

// Create a new Scrapper, with the provide options.
func NewScrapper(options ...ScrapperOption) *Scrapper {
	sc := Scrapper{
		headless: false,
		ignore:   make(map[string]bool),
	}
	sc.ctx, sc.cancel = context.WithCancel(context.Background())
	return &sc
}

// Kill all running goroutines and release scrapper resources
func (s *Scrapper) Close() {
	s.cancel()
	// Close all pages in pool if any
	if s.pool != nil {
		s.pool.Cleanup(func(p *rod.Page){p.Close()})
	}
	if s.browser != nil {
		s.browser.Close()
	}	
}

// Option when creating a new scrapper.
type ScrapperOption func(s *Scrapper)

// Option to go headless. Set to true for headless (default is false).
func Headless(h bool)ScrapperOption {
	return func(s *Scrapper) {		
		s.headless = h	
	}
}

// Ignore provided patterns. Can be called multiple times.
func Ignore(ss ...string) ScrapperOption {
	return func(s *Scrapper) {
		for _, is := range ss {
			s.ignore[is] = true
		}
	}
}

// Set global timeout for the entire scrapper running.
func Timeout(globalTimeout time.Duration) ScrapperOption {
	return func(s *Scrapper) {
		s.ctx, s.cancel = context.WithTimeout(s.ctx, globalTimeout)
	}
}

// Launch the Scrapper. Will block until scrapping is finished or a major error occured.
// Analysis happens in a separate Job thread.
func (sc *Scrapper) Run() {
	j  := sc.newJob()
	
	sc.wg.Add(1)
	go func (){ // run job in a separate thread.
		defer sc.wg.Done()
		err := j.Run({{.RunState}})
		if err != nil {
			fmt.Println(err)
		}
	}()

	sc.wg.Wait()
}



