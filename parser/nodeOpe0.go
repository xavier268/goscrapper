package parser

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/xavier268/goscrapper"
	"github.com/xavier268/goscrapper/rt"
)

// ==== No argument operators ====
type nodeOpe0 tok

var _ Node = nodeOpe0{}

// eval implements Node.
func (n nodeOpe0) eval(it *Interpreter) (rv any, err error) {
	switch n.c {
	case NOW:
		return time.Now(), nil
	case VERSION:
		return goscrapper.VERSION, nil
	case FILE_SEPARATOR:
		return string(filepath.Separator), nil
	case NL:
		return "\n", nil
	case NIL:
		return nil, nil
	case RED:
		return ColRED, nil
	case GREEN:
		return ColGREEN, nil
	case BLUE:
		return ColBLUE, nil
	case YELLOW:
		return ColYELLOW, nil
	case CYAN:
		return ColCYAN, nil
	case MAGENTA:
		return ColMAGENTA, nil
	case NORMAL:
		return AnsiRESET, nil
	case PAGES:
		return rt.Pages()
	default:
		return nil, fmt.Errorf("unknown zero-ary operator %s", TokenAsString(n.c))
	}
}
