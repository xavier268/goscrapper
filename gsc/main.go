// command line client for running goscrapper scripts
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/xavier268/goscrapper"
	"github.com/xavier268/goscrapper/parser"
)

//TODO : Add flags to provide input parameters ?

var (
	flagVersion = flag.Bool("version", false, "print version and exit")
	flagInfo    = flag.Bool("info", false, "print info and exit")
	flagHelp    = flag.Bool("help", false, "print help and exit")
)

func main() {

	// define shorter flags
	flag.BoolVar(flagVersion, "v", false, "")
	flag.BoolVar(flagInfo, "i", false, "")
	flag.BoolVar(flagHelp, "h", false, "")

	// redefine Usage function to print default values.
	flag.Usage = func() {
		fmt.Printf("Usage of %s\n%s [ option flags ... ] file1.gsc file2.gsc  ...\n", os.Args[0], filepath.Base(os.Args[0]))
		fmt.Println("Version", goscrapper.VERSION, " - ", goscrapper.COPYRIGHT)
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
		for _, fn := range flag.Args() {
			res, err := parser.EvalFile(fn)
			if err != nil {
				fmt.Printf("\n%s%s : %v%s\n", parser.ColRED, fn, parser.AnsiRESET, err)
			} else {
				jc, err := json.MarshalIndent(res, "", "\t")
				if err != nil {
					fmt.Printf("\n%s%s : %v%s\n", parser.ColYELLOW, fn, res, parser.AnsiRESET)
				} else {
					fmt.Printf("\n%s%s : %v%s\n", parser.ColGREEN, fn, jc, parser.AnsiRESET)
				}
			}
		}
	}
}
