package main

import (
	"flag"
	"fmt"
	"runtime"

	"githib.com/xavier268/goscrapper"
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

	goscrapper.DEBUG = *flagDebug
	switch {
	case *flagVersion:
		if goscrapper.DEBUG == 0 {
			// simple version for use in build pipeline.
			fmt.Println(goscrapper.VERSION)
			return
		} else {
			// more verbose info
			fmt.Printf("goscrapper version: %s\n", goscrapper.VERSION)
			fmt.Printf("debug level: %d\n", goscrapper.DEBUG)
			fmt.Printf("githash: %s\n", goscrapper.GITHASH)
			fmt.Printf("buildtime: %s\n", goscrapper.BUILDTIME)
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
