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
	"github.com/xavier268/goscrapper/rt"
)

//TODO : Add flags to provide input parameters ?

var (
	flagVersion = flag.Bool("version", false, "print version and exit")
	flagInfo    = flag.Bool("info", false, "print info and exit")
	flagHelp    = flag.Bool("help", false, "print help and exit")
	flagFormat  = flag.String("format", "gsc", "output format (gsc, json, go, raw)")
	flagParam   = flag.String("param", "", "provide map of input paramaters in gsc litteral format")
)

func main() {

	var params map[string]any

	// define shorter flags
	flag.BoolVar(flagVersion, "v", false, "")
	flag.BoolVar(flagInfo, "i", false, "")
	flag.BoolVar(flagHelp, "h", false, "")
	flag.StringVar(flagFormat, "f", "gsc", "")
	flag.StringVar(flagParam, "p", "", "")

	// redefine Usage function to print default values.
	flag.Usage = func() {
		fmt.Printf("Usage of %s\n%s [ option flags ... ] file1.gsc file2.gsc  ...\n", os.Args[0], filepath.Base(os.Args[0]))
		fmt.Println("Version", goscrapper.VERSION, " - ", goscrapper.COPYRIGHT)
		flag.PrintDefaults()
	}

	// Parse
	flag.Parse()

	switch { // handle special immediate action flags
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

	}

	// set parameters from flag
	if flagParam != nil && *flagParam != "" {
		// fmt.Printf("Read parameters : %q\n", *flagParam)
		pp, err := parser.ParseLitteral(*flagParam)
		if err != nil {
			fmt.Println(parser.ColRED, "could not parse provided input parameters :", err, parser.AnsiRESET)
			return
		}
		switch pp := pp.(type) {
		case nil:
			params = nil
		case map[string]any:
			params = pp
		default:
			fmt.Println(parser.ColRED, "could not parse provided input parameters into a map[string]any :", err, parser.AnsiRESET)
			return
		}
	}

	for _, fn := range flag.Args() {
		// fmt.Println("Using parameters : ", params)
		res, err := parser.EvalFileWithParams(fn, params)
		if err != nil {
			fmt.Printf("\n%s%s : %v%s\n", parser.ColRED, fn, parser.AnsiRESET, err)
		} else {
			switch *flagFormat {
			case "gsc":
				out, err := rt.Serialize(res)
				if err != nil {
					fmt.Println(parser.ColRED, "could not serialize result using gsc format :", err, parser.AnsiRESET)
				} else {
					fmt.Println(parser.ColGREEN, out, parser.AnsiRESET)
				}
			case "json":
				b, err := json.MarshalIndent(res, "", "  ")
				if err != nil {
					fmt.Println(parser.ColRED, "could not serialize result using json format :", err, parser.AnsiRESET)
				} else {
					fmt.Println(parser.ColGREEN, string(b), parser.AnsiRESET)
				}
			case "go":
				fmt.Println(parser.ColGREEN, res, parser.AnsiRESET)
			case "raw":
				fmt.Printf("%s%#v%s", parser.ColGREEN, res, parser.AnsiRESET)
			default:
				fmt.Println(parser.ColRED, "Unknown format", *flagFormat, parser.AnsiRESET)
			}
		}
	}
}
