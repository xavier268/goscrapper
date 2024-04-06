package parser

// add imports for the requested packages
func (m *myLexer) addImport(pkgs ...string) {
	for _, what := range pkgs {
		switch what {
		case "rt":
			m.imports["github.com/xavier268/goscrapper/rt"] = true
		case "rod":
			m.imports["github.com/go-rod/rod"] = true
		case "fmt":
			m.imports["fmt"] = true
		case "context":
			m.imports["context"] = true
		case "strings":
			m.imports["strings"] = true
		case "os":
			m.imports["os"] = true
		case "time":
			m.imports["time"] = true
		case "regexp":
			m.imports["regexp"] = true
		default:
			panic("invalid import package : " + what)
		}
	}
}
