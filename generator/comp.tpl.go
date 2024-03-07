package generator

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed tpl/*
var tplFS embed.FS

// Data used to instantiate templates
// Derived from Compiler configuration, if available.
type TplData struct {
	AppName        string
	TargetDir      string
	Package        string
	PackageDir     string
	BrowserDataDir string // saving browser state, cookies, etc ...
	// BaseName       string // BaseName of template, if any.
	ExampleDir string // example of main function
}

func (c *Compiler) setTplData() (err error) {

	// (re)set defaults values
	c.TplData = &TplData{}

	// apply AppName, if available
	if c.conf.AppName != "" {
		c.AppName = UpFirst(c.conf.AppName)
	} else {
		c.AppName = UpFirst("MyApp")
	}
	c.Package = "auto" + Normalize(c.AppName)
	c.TargetDir = MustAbs(c.AppName)
	c.PackageDir = filepath.Join(c.TargetDir, c.Package)
	c.BrowserDataDir = filepath.Join(c.TargetDir, ".browserData-"+Normalize(c.AppName))

	// ensure dir exists
	err = os.MkdirAll(c.PackageDir, 0755)
	if err != nil {
		return err
	}

	return nil
}

// Generate a go file from the template file with the same name.
func (c *Compiler) generateTpl(baseName string) error {

	baseName = Normalize(filepath.Base(strings.ToLower(baseName)))

	targetFile := filepath.Join(c.PackageDir, baseName+".go")

	tpl, err := template.ParseFS(tplFS, "tpl/"+baseName+".tpl")
	if err != nil {
		return fmt.Errorf("failed to parse template file %s: %v", baseName+".tpl", err)
	}
	w, err := c.getWriter(targetFile)
	if err != nil {
		return fmt.Errorf("failed to create target file %s: %v", targetFile, err)
	}
	defer w.Close()

	// Add a header to the file.
	err = c.writeHeader(w)
	if err != nil {
		return fmt.Errorf("failed to header to target file %s: %v", targetFile, err)
	}

	err = tpl.Execute(w, c.TplData)
	if err != nil {
		return fmt.Errorf("failed to execute template %s: %v", baseName+".tpl", err)
	}

	if DEBUG_LEVEL >= LEVEL_INFO {
		fmt.Println("Generated file: " + targetFile)
	}

	return nil
}
