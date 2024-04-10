package rt

import (
	"context"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

// Generic iterator interface
type Iterator[T any] interface {
	Next() (next T, ok bool)
}

// SELECT ALL

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

func NewSelectAllIterator[T Elementer](ctx context.Context, pageOrElement T, css string, limit int) *selectAllIteratorP {
	return &selectAllIteratorP{
		source: pageOrElement,
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

// === FOR LOOP ITERATOR ===

type forLoopIterator struct {
	from int
	to   int
	cur  int // next will provide this value, if valid.
	step int
}

var _ Iterator[int] = new(forLoopIterator)

// No enforcing on dangerous values, such as step = 0. May loop forever.
// It is not a bug, but a feature ;-).
func NewForLoopIterator(from, to, step int) Iterator[int] {
	return &forLoopIterator{
		from: from,
		to:   to,
		cur:  from,
		step: step,
	}
}

// Same, but adjusts provided values to ensure no endless loops and positive direction.
func NewForLoopIteratorSafe(from, to, step int) Iterator[int] {
	from, to, step = min(from, to), max(from, to), max(step, -step, 1)
	return NewForLoopIterator(from, to, step)
}

// Next implements Iterator.
func (f *forLoopIterator) Next() (next int, ok bool) {
	if f.cur <= max(f.from, f.to) && f.cur >= min(f.from, f.to) {
		v := f.cur
		f.cur += f.step
		return v, true
	} else {
		return 0, false
	}
}
