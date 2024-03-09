

import (
	"context"

	"github.com/go-rod/rod"
)

// a Job maintain the context a thread is running in.
type Job struct {
	sc     *Scrapper
	state  State              // the current state of the job.
	page   *rod.Page          // the tab associated with the job. Can be nil if no page loaded yet.
	sel    *rod.Element       // the current selected element. Can be nil if no element selected yet.
	ctx    context.Context    // job context
	cancel context.CancelFunc // cancel runnng job
}

// create a new job, obtaining a new page.
// no forking is done here.
func (sc *Scrapper) newJob() (j *Job) {
	j = new(Job)
	j.sc = sc
	j.page = sc.newPage()	
	j.ctx, j.cancel = context.WithCancel(sc.ctx)
	return j
}


// access the global Browser instance
func (j *Job) Browser() *rod.Browser {
	return j.sc.Browser()
}

// close the current page, returning it to the pool.
// This does NOT necessarily ends the job.
func (j *Job) ClosePage() {
	j.sel = nil
	if j.page == nil {
		j.sel = nil
		return
	}
	if j.sc.pool == nil {
		j.page.Close()
		
	} else {
		j.sc.pool.Put(j.page)
	}
	j.page = nil
}