package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Definitions defines a specific app.
// Typically read from a yaml file, using the same key, but in lowercase.
// WARNING : !!! in the configuration file, use LOWERCASE keys, not uppercase !!!
type Definitions struct {
	Version string
	Bus     []BusDefinition
	States  []StateDefinition
}

type StateDefinition struct {
	State string

	// Conditions to confirm before State can be accepted
	// If any of these is refused, this Ste cannot be entered into.
	OnStart []ConditionDefinition

	// Wait for conditions to confirm before moving on.
	// If any exit condition cannot be verified, State will ultimately timeout ...
	OnExit []ConditionDefinition

	// Once Ste is accepted, select a single domElement as the reference, to possibly reduce the scope
	Select string

	// Execute each action, in that order, until an error occur, or a new State is requested
	Actions []ActionDefinition
}

// ActionDefinition defines action to conduct.
type ActionDefinition struct {
	Action string // name of a valid action
	// Below is a union of all possible action parameters, never required
	// Their meaning will depend upon action
	// Syntax will be checked only at compile time.
	Next string // next state to target
	From string // Bus name
	To   string // Bus name

}

// ConditionDefinition defines condition to be met before State can be accepted
// or after Action(s) have been done.
type ConditionDefinition struct {
	Action string // name of a valid condition
	Loc    string // selector for the condition
}

type Parameters map[string]interface{}

// Parse configuration from the yaml file provided.
// Major version in configuration and in the application should match.
func ParseDefinitions(fileName string) (*Definitions, error) {

	var config Definitions

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	// fmt.Println("Read file : \n", (string)(data))

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	mc, _, _ := parseVersion(config.Version)
	mv, _, _ := ParseVersion()
	// fmt.Println(VERSION, " is version app, and file version is ", config.Version)
	if mv != mc {
		return nil, fmt.Errorf("version mismatch between %s and %s", VERSION, config.Version)
	}
	return &config, nil
}

// A Bus can send and/or receive data to/from States.
// There are implemented as go chanels.
type BusDefinition struct {
	Name   string // Unique identifier of communication bus
	Sink   bool   // Bus can receive data from State
	Source bool   // Bus can send data to State
	Limit  int    // Max capacity of underlying channel
}
