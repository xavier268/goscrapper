package goscrapper

const (
	VERSION   = "0.0.2"
	COPYRIGHT = "(c) Xavier Gandillot 2024"
)

// build time variables
var (
	GITHASH   string = "n/a"
	BUILDDATE string = "n/a"
)

// runtime variables (from flags or test configuration)
var (
	DEBUG = 0
)
