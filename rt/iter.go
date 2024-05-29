package rt

import (
	"context"

	"github.com/go-rod/rod"
)

// Generic iterator interface
type Iterator[T any] interface {
	Next() (next T, ok bool)
}

// SELECT ALL

// iterator implementation
type selectAllIteratorP struct {
	source Elementer      // page or element on which we iterate
	css    string         // css selector we use
	xpath  bool           // use xpath instead of css
	elms   []*rod.Element // remaining potential elements
	// done   map[proto.RuntimeRemoteObject]bool // remote objects we have already returned - does not work, we can't recognize reliably former elements
	count int             // number of elements returned
	limit int             // limit of elements to return
	ctx   context.Context // passed on construction
}

var _ Iterator[*rod.Element] = &selectAllIteratorP{}

func NewSelectAllIterator[T Elementer](ctx context.Context, pageOrElement T, css string, limit int, xpath bool) *selectAllIteratorP {
	it := &selectAllIteratorP{
		source: pageOrElement,
		css:    css,
		xpath:  xpath,
		elms:   make([]*rod.Element, 0, 5),
		count:  0,
		limit:  limit,
		ctx:    ctx,
	}
	it.loadMore() // load initial elements
	return it
}

func (sit *selectAllIteratorP) loadMore() (err error) {
	var more rod.Elements
	if sit.xpath {
		more, err = sit.source.ElementsX(sit.css)
	} else {
		more, err = sit.source.Elements(sit.css)
	}
	if err != nil {
		return err
	}
	sit.elms = append(sit.elms, more...)
	return nil
}

func (sit *selectAllIteratorP) Next() (next *rod.Element, ok bool) {

	if sit.source == nil {
		return nil, false
	}

	// check limit - negative or zero means no limit
	if sit.limit > 0 && sit.count >= sit.limit {
		return nil, false
	}

	// try to send an element from waiting list
	if len(sit.elms) == 0 {
		return nil, false // nothing to send
	}

loop:
	el := sit.elms[0]

	// check context
	if sit.ctx.Err() != nil {
		return nil, false
	}

	if el == nil || el.Object == nil { // element is invalid
		sit.elms = sit.elms[1:] // skip and remove element from waiting list
		goto loop               // iterate ...
	}

	// found a valid element, update state and send it
	sit.count += 1
	sit.elms = sit.elms[1:] // update waiting list
	return el, true         // send valid element
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
