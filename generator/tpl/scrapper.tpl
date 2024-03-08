
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
func (s *Scrapper) Run() error {
	panic("not implemented")
}

// Request a new page from browser.
// todo	 : implement pagePooling ...
func (s *Scrapper) newPage() (*rod.Page, error) {
panic("todo")
}

// Fork current job to a new job from Scrapper.
// Start to run the specified State with the forked thread.
// Continue running current job.
func (j *Job) fork (state State) (error) {
	j2, err := j.sc.newJob()
	if err != nil {
		return err 
	}
	j.sc.wg.Add(1)
	go func () {
		j2.Run(state)
		j.sc.wg.Done()
	}()
	return nil
}

// ================================================


// a Job maintain the context a thread is running in.
type Job struct {
	sc *Scrapper	
	state State // the current state of the job.
	page *rod.Page // the tab associated with the job. Can be nil if no page loaded yet.
	sel *rod.Element // the current selected element. Can be nil if no element selected yet.
	// TO DO
}

// create a new job, obtaining a new page.
// no forking is done here.
func (sc *Scrapper) newJob() (j *Job, err error) {
	j = new(Job)
	j.sc = sc
	j.page, err  = sc.newPage()
	if err != nil {
		return nil, err
	}
	return j, nil
}





