
import (
	"context"
	"sync"
	"time"
)

// Scrapper defines the top level object for scrapping.
type Scrapper struct {
	headless bool               // should we run headless
	ignore   map[string]bool    // set of requests patterns we hould we ignore ( eg : *.png, *.jpg, ...)
	ctx      context.Context    // global context
	cancel   context.CancelFunc // cancel to cancel everything
	wg       sync.WaitGroup     // to monitor pending threads
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
	s.wg.Wait()
}

// Option when creating a new scrapper.
type ScrapperOption func(s *Scrapper)

// Option to go headless.
var Headless = func(s *Scrapper) {
	if s == nil {
		return
	}
	s.headless = true
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
func (s *Scrapper) Run() error {
	panic("not implemented")
}

