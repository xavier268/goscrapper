package parser

import (
	"fmt"
	"testing"
	"time"
)

func TestPretty(t *testing.T) {

	ti := time.Now()
	fmt.Println(Pretty(ti))

	fmt.Println("====================")

	a := struct {
		a int
		b struct {
			c bool
			d []struct {
				e *int
				f bool
			}
		}
	}{}
	fmt.Println(Pretty(a))
}
