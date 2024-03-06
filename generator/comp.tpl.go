package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// Data used to instantiate templates
// Derived from Compiler configuration, if available.
type TplData struct {
	AppName        string
	Package        string
	TargetDir      string
	PackageDir     string
	BrowserDataDir string
	Imports        []string
	TplDir         string
	BaseName       string
}

func (c *Compiler) setTplData() (err error) {

	// (re)set defaults values
	c.TplData = &TplData{
		AppName:        "MyApp",
		Package:        "mypack",
		TargetDir:      MustAbs("MyApp"),
		PackageDir:     MustAbs(filepath.Join(MustAbs("MyApp"), "mypack")),
		BrowserDataDir: ".browserData",
		Imports:        []string{"fmt", "time"},
		TplDir:         MustAbs(filepath.Join("generator", "tpl")),
	}

	// apply configuration, if available
	if c.conf.AppName != "" {
		c.AppName = c.conf.AppName
		c.Package = strings.ToLower(c.conf.AppName)
		c.TargetDir = MustAbs(c.conf.AppName)
		c.PackageDir = filepath.Join(c.TargetDir, c.Package)
		c.BrowserDataDir = filepath.Join(c.TargetDir, ".browserData-"+c.Package)
	}

	// ensure dir exists
	err = os.MkdirAll(c.PackageDir, 0755)
	if err != nil {
		return err
	}

	return nil
}

func (c *Compiler) Compile() (err error) {

	err = c.generateTpl("util")
	if err != nil {
		return err
	}
	return nil
}

// Generate a go file from the template file with the same name.
func (c *Compiler) generateTpl(baseName string) error {

	c.BaseName = filepath.Base(strings.ToLower(baseName))

	targetFile := filepath.Join(c.PackageDir, c.BaseName+".go")
	sourceFile := filepath.Join(c.TplDir, c.BaseName+".tpl")

	tpl, err := template.ParseFiles(sourceFile)
	if err != nil {
		return fmt.Errorf("failed to parse template file %s: %v", sourceFile, err)
	}
	w, err := c.getWriter(targetFile)
	if err != nil {
		return fmt.Errorf("failed to create target file %s: %v", targetFile, err)
	}
	defer w.Close()
	err = tpl.Execute(w, c.TplData)
	if err != nil {
		return fmt.Errorf("failed to execute template %s: %v", sourceFile, err)
	}

	if DEBUG_LEVEL >= LEVEL_INFO {
		fmt.Println("Generated file: " + targetFile)
	}

	return nil
}
