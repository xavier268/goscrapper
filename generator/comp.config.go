package generator

import (
	"fmt"
	"os"
	"path/filepath"
)

// Generate a copy of the compiler configuration and of the loaded configuration inside the package, for reference.
func (c *Compiler) generateConfig() (err error) {

	fname := "config.json"
	f, err := os.Create(filepath.Join(c.PackageDir, fname))
	if err != nil {
		return fmt.Errorf("failed to create %s: %v", fname, err)
	}
	defer f.Close()
	_, err = fmt.Fprintln(f, PrettyJson(c.conf))
	if err != nil {
		return fmt.Errorf("failed to write %s: %v", fname, err)
	}

	fname = "compconfig.json"
	fc, err := os.Create(filepath.Join(c.PackageDir, fname))
	if err != nil {
		return fmt.Errorf("failed to create %s: %v", fname, err)
	}
	defer fc.Close()
	_, err = fmt.Fprintln(fc, PrettyJson(c))
	if err != nil {
		return fmt.Errorf("failed to write %s: %v", fname, err)
	}

	return nil
}
