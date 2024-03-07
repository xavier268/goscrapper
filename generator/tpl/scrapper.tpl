import (
	"context"
	"sync"
	"time"
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

// Request a new page from browser.
// todo	 : implement pagePooling ...
func (s *Scrapper) newPage() (*rod.Page, error) {
panic("todo")
}

// Fork a new job from Scrapper.
// Start to run the specified State with the forked thread.
func (s *Scrapper) fork (state State) (error) {
	job, err := s.newJob()
	if err != nil {
		return err 
	}
	s.wg.Add(1)
	go state(job)
	return nil
}

// ================================================

// A State is just any function that can be applied to a *job
// Everytime we enter a State, we increment the WaitGroup,
// everytime we return from a State, we decrement the WaitGroup,
// even when moving from State to State in the same thread.
type State func(*job)


// ==================================================

// a job maintain the context a thread is running in.
type job struct {
	sc *Scrapper
	page *rod.Page // the tab associated with the job
	// todo
}

// create a new job, obtaining a new page.
// no forking is done here.
func (sc *Scrapper) newJob() (j *job, err error) {
	j = new(job)
	j.sc = sc
	j.page, err  = sc.newPage()
	if err != nil {
		return nil, err
	}
	return j, nil
}