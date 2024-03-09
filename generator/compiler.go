package generator

import (
	"fmt"
	"time"
)

// Compiler is responsible for generating go source code from the parsed configuration.
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

		fileName = MustAbs(fileName)
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

	// evrything is ok !
	if DEBUG >= LEVEL_DEBUG {
		fmt.Println("Loaded configuration :")
		fmt.Println(PrettyJson(c.conf))
	}

	return nil
}

func (c *Compiler) Compile() (err error) {

	// prepare template data
	err = c.setTplData()
	if err != nil {
		return err
	}

	// evrything is ok !
	if DEBUG >= LEVEL_DEBUG {
		fmt.Println("Compiler configuration :")
		fmt.Println(PrettyJson(c))
	}

	err = c.generateTpl("actions", "scrapper", "job", "browser")
	if err != nil {
		return err
	}

	err = c.generateConfig()
	if err != nil {
		return err
	}

	err = c.generateBuses()
	if err != nil {
		return err
	}

	err = c.generateStates()
	if err != nil {
		return err
	}

	return nil
}
