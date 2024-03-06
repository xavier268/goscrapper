package generator

import (
	"fmt"
	"os"
	"slices"
	"time"

	"gopkg.in/yaml.v3"
)

// Compiler is resposible for generating go source code from the parsed configuration.
type Compiler struct {

	// Internal compiler data
	start       time.Time // compiler creation date
	sourcefiles []string  // sources files parsed

	// data for template
	*TplData

	// parsed configuration
	conf *Configuration
}

// create a new, default compiler.
func NewCompiler() *Compiler {
	return &Compiler{
		start:       time.Now(),
		sourcefiles: []string{},
		TplData:     nil,
		conf:        NewConfiguration(),
	}
}

// Parse specification and configuration files.
func (c *Compiler) Parse(fileNames ...string) (err error) {

	for _, fileName := range fileNames {
		conf, err := parseFile(fileName)
		if err != nil {
			return err
		}
		err = c.merge(fileName, conf)
		if err != nil {
			return err
		}
	}

	// verify obvious issues
	err = c.verifyConfig()
	if err != nil {
		return err
	}

	// prepare template data
	err = c.setTplData()
	if err != nil {
		return err
	}

	// evrything is ok !
	if DEBUG_LEVEL >= LEVEL_DEBUG {
		fmt.Println("Compiler configuration :")
		fmt.Println(Pretty(c))
	}

	return nil
}

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
		fmt.Println(Pretty(conf))
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

	// Enforce no duplicates here
	if conf.AppName != "" && c.conf.AppName != "" && conf.AppName != c.conf.AppName {
		return fmt.Errorf("renaming application in file %s", fileName)
	}
	c.conf.AppName = conf.AppName

	// Allow duplicates here
	c.conf.Debug = max(c.conf.Debug, conf.Debug)
	c.conf.Headless = c.conf.Headless || conf.Headless
	for s := range conf.Ignore {
		c.conf.Ignore[s] = true
	}
	c.conf.Run = append(c.conf.Run, conf.Run...)

	// Buses should not be redefined.
	for busname, bus := range conf.Buses {
		if _, ok := c.conf.Buses[busname]; ok {
			return fmt.Errorf("bus %s redefined in file %s", busname, fileName)
		} else {
			c.conf.Buses[busname] = bus
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
		fmt.Println(Pretty(conf))
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
	if len(c.conf.Run) == 0 {
		return fmt.Errorf("at least one state must be specified to run")
	}
	for _, state := range c.conf.Run {
		if _, ok := c.conf.States[state]; !ok {
			return fmt.Errorf("run state %s is not defined", state)
		}
	}

	if DEBUG_LEVEL >= LEVEL_DEBUG {
		fmt.Println("Verified configuration is :")
		fmt.Println(Pretty(c.conf))
	}

	return nil
}
