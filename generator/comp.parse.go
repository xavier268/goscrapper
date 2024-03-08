package generator

import (
	"fmt"
	"os"
	"slices"

	"gopkg.in/yaml.v3"
)

// parse a single file into a configuration object.
func parseFile(fileName string) (Configuration, error) {

	conf := Configuration{}
	if DEBUG_LEVEL > 0 {
		fmt.Println("Parsing file", fileName)
	}
	data, err := os.ReadFile(fileName)
	if err != nil {
		return conf, fmt.Errorf("cannot load file %s : %v", fileName, err)
	}
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return conf, fmt.Errorf("cannot load file %s : %v", fileName, err)
	}

	if conf.Schema != SCHEMA {
		return conf, fmt.Errorf("invalid schema in %s", fileName)
	}
	if DEBUG_LEVEL >= LEVEL_DEBUG {
		fmt.Println("Loaded :")
		fmt.Println(PrettyJson(conf))
	}
	return conf, err
}

// merge a configuration object into another one.
func (c *Compiler) merge(fileName string, conf Configuration) error {

	if slices.Contains(c.sourcefiles, fileName) {
		return fmt.Errorf("file %s was already parsed", fileName)
	}
	// Source files names
	c.sourcefiles = append(c.sourcefiles, fileName)

	// Enforce no duplicate AppName
	if conf.AppName != "" && c.conf.AppName != "" && conf.AppName != c.conf.AppName {
		return fmt.Errorf("renaming application in file %s", fileName)
	}
	if conf.AppName != "" {
		c.conf.AppName = conf.AppName
	}

	// Enforce no duplicate Run state
	if conf.Run != "" && c.conf.Run != "" && conf.Run != c.conf.Run {
		return fmt.Errorf("redefining run state in file %s", fileName)
	}
	if conf.Run != "" {
		c.conf.Run = conf.Run
	}

	// Buses should not be redefined.
	for busname, bus := range conf.Buses {
		if _, ok := c.conf.Buses[busname]; ok {
			return fmt.Errorf("bus %s redefined in file %s", busname, fileName)
		} else {
			c.conf.Buses[busname] = bus
		}
	}

	// Actions inside new States should be valid.
	for statename, state := range conf.States {
		for i, ca := range state.Actions {
			_, err := configActionVerify(ca)
			if err != nil {
				return fmt.Errorf("invalid action n°%d in state %s in file %s :\n %v", i, statename, fileName, err)
			}
		}
	}

	// States should not be redefined
	for statename, state := range conf.States {
		if _, ok := c.conf.States[statename]; ok {
			return fmt.Errorf("state %s redefined in file %s", statename, fileName)
		} else {
			c.conf.States[statename] = state
		}
	}

	// Defines should not be redefined
	for defineName, defineValue := range conf.Define {
		if _, ok := c.conf.Define[defineName]; ok {
			return fmt.Errorf("define key %s is redefined in file %s", defineName, fileName)
		} else {
			c.conf.Define[defineName] = defineValue
		}
	}

	if DEBUG_LEVEL >= LEVEL_DEBUG {
		fmt.Println("Merge is :")
		fmt.Println(PrettyJson(conf))
	}
	return nil
}

func (c *Compiler) verifyConfig() error {
	if c.conf == nil {
		return fmt.Errorf("configuration is nil")
	}
	if c.conf.Schema == "" {
		return fmt.Errorf("schema is required but was not provided")
	}
	if c.conf.Run == "" {
		return fmt.Errorf("no state was specified to run")
	}

	if _, ok := c.conf.States[c.conf.Run]; !ok {
		return fmt.Errorf("run state %s is not defined", c.conf.Run)
	}

	if DEBUG_LEVEL >= LEVEL_DEBUG {
		fmt.Println("Verified configuration is :")
		fmt.Println(PrettyJson(c.conf))
	}

	return nil
}
