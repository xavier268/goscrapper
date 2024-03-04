package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

// Parse configuration definitions from the yaml file provided.
// Major version in configuration and in the application should match.
func ParseDefinitions(files ...string) (*Configuration, error) {

	d := &Configuration{}
	d.Schema = SCHEMA
	d.ParseDate = time.Now()

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
	return d, nil
}

// Merge dd into d. dd can be nil.
// Update ErrMessages in d, and return error if needed.
func (d *Configuration) Merge(dd *Configuration) error {

	// Check schema
	if d.Schema != SCHEMA {
		d.ErrMessages = append(d.ErrMessages, fmt.Sprintf("schema mismatch : got %q, but expected %q", d.Schema, SCHEMA))
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

	d.ErrMessages = append(d.ErrMessages, dd.ErrMessages...)
	d.Files = append(d.Files, dd.Files...)
	// d.ParseDate unchanged
	if dd.Schema != d.Schema {
		d.ErrMessages = append(d.ErrMessages, fmt.Sprintf("schema mismatch between merge :  %q vs. %q", d.Schema, dd.Schema))
	}
	d.Debug = max(d.Debug, dd.Debug)
	d.Headless = d.Headless && dd.Headless
	d.PagePool = max(d.PagePool, dd.PagePool)
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

// Pretty print an object, using its json format
func Pretty(st interface{}) string {
	if st == nil {
		return "nil"
	}
	bb, err := json.MarshalIndent(st, "", "     ")
	if err != nil {
		return string(bb) + " **ERROR** " + err.Error()
	}
	return string(bb)
}
