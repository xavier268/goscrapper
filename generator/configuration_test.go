package generator

import (
	"fmt"
	"path/filepath"
	"reflect"
	"testing"
)

func TestParseConfiguration(t *testing.T) {

	tf1 := filepath.Join("..", "testfiles", "test1.yml")
	tf2 := filepath.Join("..", "testfiles", "test2.yml")

	c := NewCompiler()
	err := c.Parse(tf1, tf2)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("\nCaptured configuration :\n%s\n", PrettyJson(c.conf))
}

func TestHelpConfiguration(t *testing.T) {
	fmt.Println(HelpConfiguration())
}

func TestExperiment(_ *testing.T) {

	d := ConfigAction{}
	c := "Click"

	e := reflect.ValueOf(&d).Elem()
	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		varType := e.Type().Field(i).Type
		fmt.Printf("%v %v\n", varName, varType)
	}

	// Now, lets try to extract the value of d.Scope
	// We can extract the value of a field using reflect.ValueOf(&d).Elem().FieldByName(c)
	// But we need to know the type of the field, so we can use reflect.TypeOf(d).FieldByName(c)
	a := reflect.ValueOf(d).FieldByName(c)
	fmt.Printf("%#v\n", a)

	for i := 0; i < a.Type().NumField(); i++ {
		fmt.Println(" ---- field #", i)
		fmt.Printf("\t%#v\n", a.Field(i))                 // value
		fmt.Printf("\t%#v\n", a.Field(i).Type().String()) // type of value, struct{ kjh int ; jhg float64 }
		fmt.Printf("\t%#v\n", a.Field(i).Kind().String()) // kind of value, struct
		fmt.Printf("\t%#v\n", a.Type().Field(i).Name)     // param key
		fmt.Printf("\t%#v\n", a.Type().Field(i).Tag)      // param tag
		fmt.Printf("\t%#v\n", a.Type().Field(i).Type)     // type of param key

	}

}
