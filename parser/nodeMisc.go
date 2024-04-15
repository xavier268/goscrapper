package parser

import (
	"fmt"
	"time"

	"github.com/xavier268/goscrapper/rt"
)

// ==== NODE SLOW ====

type nodeSlow struct {
	m Node // duration in milliseconds
}

var _ Node = nodeSlow{}

// eval Node
func (n nodeSlow) eval(it *Interpreter) (any, error) {
	// default millis duration
	dur := rt.SLOW_DELAY

	// evaluate duration
	if n.m != nil {
		d, err := n.m.eval(it)
		if err != nil {
			return nil, err
		}
		if d != nil { // if d is nil, use default duration ...
			// verify d is an int
			dd, ok := d.(int)
			if !ok {
				return nil, fmt.Errorf("expected an int millis duration, but got a %T", d)
			}
			if dd > 0 {
				dur = time.Duration(dd) * time.Millisecond // only set for valid, strictly positive values.
			}
		}
	}
	// sleep until duration has expired or context is cancelled
	timer := time.NewTimer(dur)
	select {
	case <-it.ctx.Done():
		if !timer.Stop() {
			<-timer.C // drain the channel if the timer had already expired.
			return nil, it.ctx.Err()
		}
		return nil, it.ctx.Err()
	case <-timer.C:
		return nil, nil
	}
}

// === NODE FAIL ===

type nodeFail struct {
	what Node
}

var _ Node = nodeFail{}

func (n nodeFail) eval(it *Interpreter) (any, error) {
	if n.what != nil {
		v, err := n.what.eval(it)
		if err != nil {
			return nil, err
		}
		if v != nil {
			if vv, ok := v.(string); ok {
				return nil, fmt.Errorf("fail message : %s", vv)
			}
		}
	}
	return nil, fmt.Errorf("fail requested")
}

// === NODE ASSERT ===

type nodeAssert struct {
	cond Node
}

var _ Node = nodeAssert{}

func (n nodeAssert) eval(it *Interpreter) (any, error) {
	// evaluate condition
	if n.cond != nil {

		v, err := n.cond.eval(it)
		if err != nil {
			return nil, err
		}
		if v != nil {

			// verify v is a bool
			vv, ok := v.(bool)
			if ok && vv {
				return nil, nil // assertion success
			}
		}
	}
	return nil, fmt.Errorf(fmt.Sprintf("assertion %#v failed", n.cond))
}

// ==== PRINT NODE ===

type nodePrint struct {
	nodes Nodes
}

var _ Node = nodePrint{}

func (n nodePrint) eval(it *Interpreter) (any, error) {
	if n.nodes != nil {
		// evaluate and print
		for _, p := range n.nodes {
			v, err := p.eval(it)
			if err != nil {
				return nil, err
			}
			fmt.Print(v)
		}
	}
	fmt.Println()
	return nil, nil
}
