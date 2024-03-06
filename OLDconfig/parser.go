package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/xavier268/goscrapper/generator"
	"gopkg.in/yaml.v3"
)

// Parse configuration definitions from the yaml file provided.
// Major version in configuration and in the application should match.
func ParseDefinitions(files ...string) (*Configuration, error) {
	var err error

	d := &Configuration{}
	d.Schema = generator.SCHEMA
	d.ParseDate = time.Now()

	d.RootDir, err = os.Getwd()
	if err != nil {
		panic(err)
	}

	for _, fileName := range files {
		if d.Debug > 0 {
			log.Printf("loading configuration file : %s", fileName)
		}
		data, err := os.ReadFile(fileName)
		if err != nil {
			return d, err
		}
		dd := &Configuration{}
		dd.Files = []string{fileName}
		err = yaml.Unmarshal(data, dd)
		if err != nil {
			return d, err
		}
		err = d.Merge(dd)
		if err != nil {
			return d, err
		}
	}
	// Wait until the end to set browserDir with the full name
	d.BrowserDataDir = filepath.Join(d.RootDir, ".browserdata-"+d.Name)
	// set default PagePool if zero, note that default for PagePoolIncognito remains 0.
	if d.PagePool == 0 {
		d.PagePool = 10
	}
	return d, nil
}

// Merge dd into d. dd can be nil.
// Update ErrMessages in d, and return error if needed.
func (d *Configuration) Merge(dd *Configuration) error {

	// Check schema
	if d.Schema != generator.SCHEMA {
		d.ErrMessages = append(d.ErrMessages, fmt.Sprintf("schema mismatch : got %q, but expected %q", d.Schema, generator.SCHEMA))
	}

	// handle special case dd == nil
	if dd == nil {
		if len(d.ErrMessages) > 0 {
			return fmt.Errorf("configuration error : %s", strings.Join(d.ErrMessages, "\n"))
		} else {
			return nil
		}
	}

	// verify dd and d and merge into d
	d.Name += dd.Name // concatenate application names
	d.ErrMessages = append(d.ErrMessages, dd.ErrMessages...)
	d.Files = append(d.Files, dd.Files...)
	// d.ParseDate unchanged
	if dd.Schema != d.Schema {
		d.ErrMessages = append(d.ErrMessages, fmt.Sprintf("schema mismatch between merge :  %q vs. %q", d.Schema, dd.Schema))
	}
	d.Debug = max(d.Debug, dd.Debug)
	d.Headless = d.Headless && dd.Headless

	d.PagePool = d.PagePool + dd.PagePool
	d.PagePoolIncognito = d.PagePoolIncognito + dd.PagePoolIncognito

	d.Run = append(d.Run, dd.Run...)
	for k, v := range dd.Env { // overwite env values
		d.Env[k] = v
	}
	for k, v := range dd.Buses { // merge Buses, no duplication allowed, no overwite
		if _, ok := d.Buses[k]; ok {
			d.ErrMessages = append(d.ErrMessages, fmt.Sprintf("bus %q already defined", k))
		} else {
			d.Buses[k] = v
		}
	}
	for k, v := range dd.States { // merge States, no duplication allowed, no overwite
		if _, ok := d.States[k]; ok {
			d.ErrMessages = append(d.ErrMessages, fmt.Sprintf("State %q already defined", k))
		} else {
			d.States[k] = v
		}
	}

	// report errors if any
	if len(d.ErrMessages) > 0 {
		return fmt.Errorf("configuration error : %s", strings.Join(d.ErrMessages, "\n"))
	} else {
		return nil
	}
}
