package parser

const (
	esc  = "\x1b"    // escape
	csi  = esc + "[" // control sequence introducer
	bell = "\x07"    // bell

	// Colors
	ColRED     = csi + "31m"
	ColGREEN   = csi + "32m"
	ColYELLOW  = csi + "33m"
	ColBLUE    = csi + "34m"
	ColMAGENTA = csi + "35m"
	ColCYAN    = csi + "36m"

	// Graphics parameters
	Bold      = csi + "1m"
	Italic    = csi + "3m"
	Underline = csi + "4m"
	Blink     = csi + "5m"
	Reverse   = csi + "7m"
	Strike    = csi + "9m"
	Dim       = csi + "2m"

	// Cursor save/restore
	CursorSave    = csi + "s"
	CursorRestore = csi + "u"

	// Clear creen
	ClearScreen             = csi + "2J"
	ClearScreenFromCursor   = csi + "0J"
	ClearScreenBeforeCursor = csi + "1J"

	// Clear line
	ClearLine             = csi + "2K"
	ClearLineFromCursor   = csi + "0K"
	ClearLineBeforeCursor = csi + "1K"

	// Reset all settings
	RESET = csi + "0m"
)
