package rt

import (
	"context"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

// Generic itérator interface
type Iterator[T any] interface {
	Next() (next T, ok bool)
}

// Can provides Elements. A page or an element.
type Elementer interface {
	Elements(string) (rod.Elements, error)
}

// iterator implementation
type selectAllIteratorP struct {
	source Elementer                            // page or element on which we iterate
	css    string                               // css selector we use
	elms   []*rod.Element                       // remaining potential elements
	done   map[proto.RuntimeRemoteObjectID]bool // ids of objects we have already returned
	count  int                                  // number of elements returned
	limit  int                                  // limit of elements to return
	ctx    context.Context                      // passed on construction
}

var _ Iterator[*rod.Element] = &selectAllIteratorP{}

func NewSelectAllIterator[T Elementer](ctx context.Context, page T, css string, limit int) *selectAllIteratorP {
	return &selectAllIteratorP{
		source: page,
		css:    css,
		elms:   make([]*rod.Element, 0, 5),
		done:   make(map[proto.RuntimeRemoteObjectID]bool, 5),
		limit:  limit,
		ctx:    ctx,
	}
}

func (it *selectAllIteratorP) Next() (next *rod.Element, ok bool) {

	for {

		// check limit - negative or zero means no limit
		if it.limit > 0 && it.count >= it.limit {
			return nil, false
		}

	loop: // try to send an element from waiting list
		for i, el := range it.elms {

			// check context
			if it.ctx.Err() != nil {
				return nil, false
			}

			if it.done[el.Object.ObjectID] {
				it.elms = it.elms[i+1:] // remove element from waiting list
				break loop
			}
			// found one, update state and send it
			it.done[el.Object.ObjectID] = true
			it.count += 1
			it.elms = it.elms[i+1:] // remove element from waiting list
			return el, true
		}

		// check context
		if it.ctx.Err() != nil {
			return nil, false
		}

		// nothing was sent - can we load more ?
		more, err := it.source.Elements(it.css)
		if err != nil {
			Errorf("error trying to retrieve elements with %s : %v", it.css, err)
			return nil, false
		}

		// check context
		if it.ctx.Err() != nil {
			return nil, false
		}

		// add to waiting list
		for _, el := range more {
			if !it.done[el.Object.ObjectID] { // only add elts not already sent
				it.elms = append(it.elms, el)
			}
		}
		// we could not find more elements not already sent
		if len(it.elms) == 0 {
			return nil, false
		}

		// check context
		if it.ctx.Err() != nil {
			return nil, false
		}

	}

}
