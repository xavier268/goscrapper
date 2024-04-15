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
	source Elementer                          // page or element on which we iterate
	css    string                             // css selector we use
	xpath  bool                               // use xpath instead of css
	elms   []*rod.Element                     // remaining potential elements
	done   map[proto.RuntimeRemoteObject]bool // remote objects we have already returned
	count  int                                // number of elements returned
	limit  int                                // limit of elements to return
	ctx    context.Context                    // passed on construction
}

var _ Iterator[*rod.Element] = &selectAllIteratorP{}

func NewSelectAllIterator[T Elementer](ctx context.Context, pageOrElement T, css string, limit int, xpath bool) *selectAllIteratorP {
	return &selectAllIteratorP{
		source: pageOrElement,
		css:    css,
		xpath:  xpath,
		elms:   make([]*rod.Element, 0, 5),
		done:   make(map[proto.RuntimeRemoteObject]bool, 5),
		count:  0,
		limit:  limit,
		ctx:    ctx,
	}
}

func (sit *selectAllIteratorP) Next() (next *rod.Element, ok bool) {
	var err error

	if sit.source == nil {
		return nil, false
	}
	for {
		// check limit - negative or zero means no limit
		if sit.limit > 0 && sit.count >= sit.limit {
			return nil, false
		}

	loop: // try to send an element from waiting list
		for i, el := range sit.elms {

			// check context
			if sit.ctx.Err() != nil {
				return nil, false
			}

			if el == nil || el.Object == nil || sit.done[*el.Object] {
				sit.elms = sit.elms[i+1:] // remove element from waiting list
				break loop
			}
			// found one, update state and send it
			sit.done[*el.Object] = true
			sit.count += 1
			sit.elms = sit.elms[i+1:] // remove element from waiting list
			return el, true
		}

		// check context
		if sit.ctx.Err() != nil {
			return nil, false
		}

		// nothing was sent - can we load more ?
		var more rod.Elements
		if sit.xpath {
			more, err = sit.source.ElementsX(sit.css)
		} else {
			more, err = sit.source.Elements(sit.css)
		}
		if err != nil {
			Errorf("error trying to retrieve elements with %s : %v", sit.css, err)
			return nil, false
		}

		// check context
		if sit.ctx.Err() != nil {
			return nil, false
		}

		// add to waiting list
		for _, el := range more {
			if el != nil && el.Object != nil && !sit.done[*el.Object] { // only add elts not already sent
				sit.elms = append(sit.elms, el)
			}
		}

		// despite all our efforts, we could not find more elements not already sent
		if len(sit.elms) == 0 {
			return nil, false
		}

		// check context
		if sit.ctx.Err() != nil {
			return nil, false
		}

		// iterate main loop ...
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
