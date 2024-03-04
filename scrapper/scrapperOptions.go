package scrapper

import "fmt"

// An option that is used to define the scrapper
type ScrapperOption interface {
	String() string    // description of option
	apply(s *Scrapper) // apply to provided scrapper
}

// SOHeadless defines if the browser should be headless or not.
func SOHeadless(b bool) ScrapperOption {
	return soSetHeadless{b}
}

type soSetHeadless struct {
	b bool
}

func (s soSetHeadless) String() string {
	if s.b {
		return "ScrapperOption : browser is headless"
	} else {
		return "ScrapperOption : browser is NOT headless"
	}
}

func (s soSetHeadless) apply(scrapper *Scrapper) { scrapper.headless = s.b }

// SOIgnore defines the patterns to ignore.
func SOIgnore(pattern ...string) ScrapperOption {
	return soIgnore{pattern}
}

type soIgnore struct {
	pattern []string
}

// apply implements ScrapperOption.
func (so soIgnore) apply(s *Scrapper) {
	s.ignore = append(s.ignore, so.pattern...)
}

func (s soIgnore) String() string {
	return fmt.Sprintf("ScrapperOption : ignoring the following patterns : %v", s.pattern)
}
