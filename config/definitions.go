package config

import "time"

// Configuration defines a specific app.
// Typically read from a yaml file, using the same key, but in lowercase.
// WARNING : !!! in the configuration file, use LOWERCASE keys only, not uppercase !!!
// Once created, this object is never modified.
type Configuration struct {
	// These fields are NOR read from file, but populated directly when parsing file.
	ErrMessages    []string  `yaml:"-"` // Error messages collected while parsing
	Files          []string  `yaml:"-"` // source file names
	ParseDate      time.Time `yaml:"-"` // Start of parsing
	RootDir        string    `yaml:"-"` // main directory where everything happens
	BrowserDataDir string    `yaml:"-"` // where session data is stored

	// These fields are read from file. Once set, they are never modified.
	Schema            string            // schema should be "1"
	Name              string            // application name. Used to separate browser data files
	Debug             int               // set app debug level
	Headless          bool              // set app headless mode
	Ignore            []string          // patterns that are never downloaded (ex : *.png) to save bandwidth
	PagePool          int               // max shared page pool - default to 10
	PagePoolIncognito int               // max shared page pool for incognito mode - default to 0
	Run               []string          // States to launch at startup. Same state can be repeated. Multiple States will run concurrently.
	Env               map[string]string // map parameter name to value. These are NEVER modified once created. Useful for passwords, etc ...

	Buses  map[string]BusDefinition   // map name to definition
	States map[string]StateDefinition // map name to definition

}

type StateDefinition struct {
	// Conditions to confirm before State can be accepted
	// If any of these is refused, this State will not be transitionned to.
	// Conditions must be idempotent, they will be called multiple times to test which state to transition to.
	Conditions []ConditionDefinition

	// Execute each action, in that order, until an error occur, or a new State is requested.
	// action can include transitions to a set of other states, selected based on their 'before' conditions.
	Actions []ActionDefinition
}

// ActionDefinition defines action to conduct.
type ActionDefinition struct {
	// === Only one of these should be set.

	// Select an element as the new base element to apply further operations
	// If empty string, will use full page content of current tab.
	Base struct {
		Selector string // selector to the new reference
		Bus      string // get selector from specified Bus
		Env      string // get selector from Env map
	}

	// Load new page
	Load struct {
		Url       string // page url to load. Leave empty and specify NewTab for blank tab.
		Bus       string // get url from specified Bus
		Env       string // get url from Env map
		NewTab    bool   // open in new tab (website MAY open newTabs even if NewTab = false )
		Fork      string // Fork newTab into a new thread, keep current thread running
		Incognito bool   // load or open new tab in incognito mode
	}

	// Click
	Click struct {
		Selector string // element selector
		Right    bool   // default is left, set to true to right click
	}

	// Type input
	Input struct {
		Content  string // use this input
		Bus      string // get input from specified Bus
		Env      string // get input from Env map
		Selector string // element selector
		Clear    bool
	}

	// Enter key
	Enter struct {
		Selector string // element selector
	}

	// Close current tab
	Close bool

	// Select in a list
	Select struct {
		Selector string   // menu <select> selector
		Choice   []string // what elements of the menu do we want to select
	}

	// Sleep for specified time
	Sleep struct {
		Duration int    // duration in millis
		Bus      string // wait until something can be retrieved from specified Bus
		Env      string // read duration from Env map
	}

	// Log a message to a Bus.
	Log struct {
		Message string // message to log
		Bus     string // log message to specified Bus
	}
	// Transition to one of these states if preconditions match.
	Next []string
}

// ConditionDefinition defines condition to be met before State can be accepted
// or after Action(s) have been done.
// Conditions will be checked multiple times, so they should be idempotent.
type ConditionDefinition struct {
	// === Only one of these will be set. Check will happen at compile time.
	Exist        string // element selector
	NotExist     string // element selector
	ContainsText struct {
		Selector string // element selector
		Regex    string // this regex should match text content
	}
}

// A Bus can send and/or receive data to/from States.
// There are implemented as go chanels.
type BusDefinition struct {
	Limit int // Max capacity of underlying channel
}
