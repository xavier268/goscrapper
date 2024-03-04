package scrapper

import (
	"context"
	"fmt"
	"io"
)

type Bus struct {
	s   *Scrapper
	ctx context.Context
	ch  chan string
}

// Instantiate buses from the config object.
// Should be called from the main scrapper thread.
func (s *Scrapper) MakeBuses() {

	for name, b := range s.conf.Buses {
		bus := Bus{
			s,
			s.ctx,
			make(chan string, b.Limit),
		}
		bus.s.buses[name] = &bus
	}
}

// access a given bus by name.
// Useful to retrieve data or send data from the application.
func (s *Scrapper) GetBus(name string) *Bus {
	return s.buses[name]
}

func (b *Bus) Send(data string) {
	b.ch <- data
}

func (b *Bus) Receive() string {
	return <-b.ch
}

// close Bus
func (b *Bus) Close() {
	if b != nil {
		close(b.ch)
	}
}

// Redirect the specified Bus to the specified Writer.
// Strings are writen as they are received.
// A specific thread is launch, that will stop when the context is cancelled.
func (s *Scrapper) Bus2Writer(busname string, w io.Writer) {
	bus := s.buses[busname]
	if bus == nil {
		return
	}
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		for {
			select {
			case <-s.ctx.Done():
				return
			case data := <-bus.ch:
				fmt.Print(w, data)
			}
		}
	}()
}
