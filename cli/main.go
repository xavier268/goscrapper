package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/xavier268/goscrapper/generator"
)

var (
	flagVersion = flag.Bool("version", false, "print version and exit")
	flagInfo    = flag.Bool("info", false, "print info and exit")
	flagHelp    = flag.Bool("help", false, "print help and exit")
)

func main() {

	// define shorter flags
	flag.BoolVar(flagVersion, "v", false, "print version and exit")
	flag.BoolVar(flagInfo, "i", false, "print info and exit")
	flag.BoolVar(flagHelp, "h", false, "print help and exit")

	// Parse
	flag.Parse()

	switch {
	case *flagVersion:
		// simple version for use in build pipeline.
		fmt.Println(generator.VERSION)
		return

	case *flagInfo:
		// more verbose info
		fmt.Printf("goscrapper version: %s\n", generator.VERSION)
		fmt.Printf("last commit: %s\n", generator.GITHASH)
		fmt.Printf("buildtime: %s\n", generator.BUILDDATE)
		fmt.Printf("goversion: %s\n", runtime.Version())
		fmt.Printf("goarch: %s\n", runtime.GOARCH)
		fmt.Printf("goos: %s\n", runtime.GOOS)
		return

	case *flagHelp:
		flag.PrintDefaults()
		return

	default:
		c := generator.NewCompiler()
		err := c.Parse(flag.Args()...)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = c.Compile()
		if err != nil {
			fmt.Println(err)
			return
		}

	}
}
