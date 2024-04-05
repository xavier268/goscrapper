// command line application
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/xavier268/goscrapper"
	"github.com/xavier268/goscrapper/parser"
)

var (
	flagVersion = flag.Bool("version", false, "print version and exit")
	flagInfo    = flag.Bool("info", false, "print info and exit")
	flagHelp    = flag.Bool("help", false, "print help and exit")
	flagAsync   = flag.Bool("async", false, "generate async mode api")
	flagOutdir  = flag.String("outdir", "autoscrapper", "output directory")
)

func main() {

	// define shorter flags
	flag.BoolVar(flagVersion, "v", false, "")
	flag.BoolVar(flagInfo, "i", false, "")
	flag.BoolVar(flagHelp, "h", false, "")
	flag.StringVar(flagOutdir, "o", "autoscrapper", "")
	flag.BoolVar(flagAsync, "a", false, "")

	// redefine Usage function to print default values.
	flag.Usage = func() {
		fmt.Printf("Usage of %s\n%s [ option flags ... ] inputFile1 inputFile2 inputFile3 ...\n", os.Args[0], filepath.Base(os.Args[0]))
		fmt.Println("NB : file names can be glob patterns.")
		flag.PrintDefaults()
	}

	// Parse
	flag.Parse()

	switch {
	case *flagVersion:
		// simple version for use in build pipeline.
		fmt.Println(goscrapper.VERSION)
		return

	case *flagInfo:
		// more verbose info
		fmt.Printf("goscrapper version: %s\n", goscrapper.VERSION)
		fmt.Printf("last commit: %s\n", goscrapper.GITHASH)
		fmt.Printf("buildtime: %s\n", goscrapper.BUILDDATE)
		fmt.Printf("goversion: %s\n", runtime.Version())
		fmt.Printf("goarch: %s\n", runtime.GOARCH)
		fmt.Printf("goos: %s\n", runtime.GOOS)
		return

	case *flagHelp:
		flag.Usage()
		return

	default:
		if flagOutdir == nil || *flagOutdir == "" {
			return
		}
		outDir := *flagOutdir
		if len(flag.Args()) == 0 {
			flag.Usage()
			return
		}

		for _, fn := range flag.Args() {

			if *flagAsync {
				err := parser.ParseGlobAsync(outDir, fn)
				if err != nil {
					panic(err)
				}
			} else {
				err := parser.ParseGlob(outDir, fn)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}
