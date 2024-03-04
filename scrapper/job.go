package scrapper

import (
	"context"

	"github.com/go-rod/rod"
)

// A job represent a running process, jumping from state to state.
type Job struct {
	s         *Scrapper       // original scrapper
	page      *rod.Page       // only one job owns this page.
	elem      *rod.Element    // rod element if selected
	incognito bool            // incognito mode
	state     string          // state to process
	ctx       context.Context // job specific context
}

// run the job. Blocking.
// Will always run in a separet thread, launched by Scrapper.
func (j *Job) run() {
	// page and elem will be lazily instancited ...
	defer j.s.wg.Done() // make sure we count stopped jobs !

	for {
		select {
		case <-j.ctx.Done(): // the context was cancel. Stop everything.
			if j.page != nil {
				if j.incognito {
					j.s.PutIncognitoPage(j.page)
				} else {
					j.s.PutPage(j.page)
				}
			}
			j.page = nil
			j.elem = nil
			return

		default:
			panic("not implemented")
			// process actions
			//      including selecting next state based upon preconditions ?
		}

	}

}
