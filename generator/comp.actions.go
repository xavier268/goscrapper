package generator

import (
	"fmt"
	"io"
)

// Use this to generate all action function templates
func generateActionFunctionTemplates(f io.Writer) {
	for _, an := range ActionKeys {
		hdr, err := (ConfigAction{}).getActionFunctionDeclaration(an)
		if err != nil {
			fmt.Fprintln(f, err)
			panic(err)
		}
		fmt.Fprintf(f, "\n%s{panic(\"action %s is not implemented\")}", hdr, an)

	}
	fmt.Fprintln(f)
}

func (c *Compiler) generateAction(f io.Writer, actName string, confAct ConfigAction) error {

	hdr, err := confAct.getActionFunctionDeclaration(actName)
	if err != nil {
		return err
	}
	fmt.Fprintf(f, "%s {panic(\"to be implemented\")} \n ", hdr)

	// Actually call the function, with the parameters values coming form configAct
	fmt.Fprintf(f, "j.%s(", getActionFunctionName(actName))
	params, err := confAct.getActionParameterList(actName)
	if err != nil {
		fmt.Fprintln(f, err)
		return err
	}
	for i, p := range params {
		if i > 0 {
			fmt.Fprintf(f, ", ")
		}
		fmt.Fprintf(f, "%#v", p.v)
	}
	fmt.Fprintln(f, ")")

	return nil
}
