package scrapper

import "github.com/xavier268/goscrapper/config"

// list of Actions potentially applicables
// immutable
var RegisteredActions = []Action{BaseAction}

// if newState is different, transition and reset job to newstate
// if transition, stop processing the remaining of the action list.
// action should return fast, doing nothing, if not applicable.
type Action func(j *Job, adef config.ActionDefinition) (transition bool)

// execute actions listed in current state
func (j *Job) DoAction() {

	curState := j.state
	adefs := j.s.conf.States[curState].Actions // actions definitions for current state

	for _, adef := range adefs { // range over action definition list.
		for _, ra := range RegisteredActions { // range over available actions.
			if ra(j, adef) { // A transition occured, end for this state.
				return
			}
		}
	}
}

// Action Base
func BaseAction(j *Job, adef config.ActionDefinition) (transition bool) {
	def := adef.Base
	if def.Selector == "" && def.Bus == "" {
		// action not set, ignore
		return false
	}

	var sel = def.Selector
	if sel == "" {
		sel = j.s.GetBus(def.Bus).Receive()
	}

	// sel now contains the selector to use

	// TODO - FAIRE DE HELPER POUR MANIPULER LES PAGES ET LES CREER ...

	// ensure page exists, if not, get it
	if j.page == nil {
		if j.incognito {
			j.page = j.s.GetIncognitoPage()
		} else {
			j.page = j.s.GetPage()
		}
		j.elem = nil
	}
	return false
}
