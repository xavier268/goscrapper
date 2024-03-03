package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/xavier268/goscrapper/config"
)

var (
	flagVersion = flag.Bool("version", false, "print version and exit")
	flagHelp    = flag.Bool("help", false, "print help and exit")
	flagDebug   = flag.Int("debug", 0, "set debugging level")
)

func main() {

	// define shorter flags
	flag.BoolVar(flagVersion, "v", false, "print version and exit")
	flag.BoolVar(flagHelp, "h", false, "print help and exit")
	flag.IntVar(flagDebug, "d", 0, "set debugging level")

	// Parse
	flag.Parse()

	config.DEBUG = *flagDebug
	switch {
	case *flagVersion:
		if config.DEBUG == 0 {
			// simple version for use in build pipeline.
			fmt.Println(config.VERSION)
			return
		} else {
			// more verbose info
			fmt.Printf("goscrapper version: %s\n", config.VERSION)
			fmt.Printf("debug level: %d\n", config.DEBUG)
			fmt.Printf("last commit: %s\n", config.GITHASH)
			fmt.Printf("buildtime: %s\n", config.BUILDDATE)
			fmt.Printf("goversion: %s\n", runtime.Version())
			fmt.Printf("goarch: %s\n", runtime.GOARCH)
			fmt.Printf("goos: %s\n", runtime.GOOS)
			return
		}

	case *flagHelp:
		flag.PrintDefaults()
		return
	default:

	}
}
