package generator

import (
	"fmt"
	"io"
)

// Use this to generate all action function templates
func generateActionFunctionTemplates(f io.Writer) {

	fmt.Fprintln(f) // safety new line ...

	for _, an := range ActionKeys {
		hdr, err := (ConfigAction{}).getActionFunctionDeclaration(an)
		if err != nil {
			fmt.Fprintln(f, err)
			panic(err)
		}
		fmt.Fprintf(f, "\n// %s\n%s{\n\tpanic(\"action %s is not implemented\")\n}\n\n", getActionFunctionName(an), hdr, an)

	}
}

func (c *Compiler) generateAction(f io.Writer, actName string, confAct ConfigAction) error {

	// Actually call the function, with the parameters values coming form configAct
	fmt.Fprintf(f, "			j.%s(\n", getActionFunctionName(actName))
	params, err := confAct.getActionParameterList(actName)
	if err != nil {
		fmt.Fprintln(f, err)
		return err
	}
	for _, p := range params {
		fmt.Fprintf(f, "				%#v,  //%s %s\n", p.v, p.n, p.t)
	}
	fmt.Fprintln(f, "				)")

	return nil
}
