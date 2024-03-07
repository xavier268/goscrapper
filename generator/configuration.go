package generator

// Configuration defines a specific app.
// Typically read from a yaml file, using the same key, but in lowercase.
// WARNING : !!! in the configuration file, use LOWERCASE keys only, not uppercase !!!
// Once created, this object should never be modified.
type Configuration struct {
	// These fields are NOR read from file, but populated directly when parsing file.

	// These fields are read from file. Once set, they are never modified.
	Schema  string // schema should be "1"
	AppName string // application name. Used to separate browser data files, and for package name.
	// Debug    int             // set app debug level.
	// Headless bool            // set app headless mode
	// Ignore   map[string]bool // set of patterns that are never downloaded (ex : *.png) to save bandwidth
	Run string // State to launch at startup.

	Define ConfigParameters       // constants definitions
	Buses  map[string]ConfigBus   // map name to definition
	States map[string]ConfigState // map name to definition

}

func NewConfiguration() *Configuration {
	return &Configuration{
		Schema: "1",
		// Debug:  0,
		// Ignore: make(map[string]bool),
		Buses:  make(map[string]ConfigBus),
		States: make(map[string]ConfigState),
	}
}

type ConfigParameters map[string]string

type ConfigState struct {
	// Assert to confirm before State can be accepted
	// If any of these is refused, this State will not be transitionned to.
	// Assert must be idempotent, they will be called multiple times to test which state to transition to.
	Assert []ConfigCondition

	// Execute each action, in that order, until an error occur, or a new State is requested.
	// action can include transitions to a set of other states, selected based on their 'before' conditions.
	Actions []ConfigAction
}

// ConfigAction defines action to conduct.
// Map action name to parameters.
type ConfigAction map[string]ConfigParameters

// ConfigCondition defines condition to be met before State can be accepted
// or after Action(s) have been done.
// Conditions will be checked multiple times, so they should be idempotent.
type ConfigCondition struct {
	// === Only one of these will be set. Check will happen at compile time.
	Exist struct {
		Selector string // element selector
		Inv      bool   // inverse condition
	}
	Contains struct {
		Selector string // element selector
		Regex    string // this regex should match text content
		Inv      bool   // inverse condition
	}
}

// A Bus can send and/or receive data to/from States.
// There are implemented as go chanels.
type ConfigBus struct {
	Limit int // Max capacity of underlying channel
}

// Normalized bus name
func BusName(s string) string {
	return "Bus" + UpFirst(Normalize(s))
}

// Normalized state name
func StateName(s string) string {
	return "state" + UpFirst(Normalize(s))
}

// Normalize action name
func ActionName(state, action string) string {
	return "do" + UpFirst(Normalize(state)) + UpFirst(Normalize(action))
}
