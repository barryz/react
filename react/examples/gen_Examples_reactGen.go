// Do not edit: file generated by reactGen

package main

import "github.com/myitcv/gopherjs/react"

func (e *ExamplesDef) ShouldComponentUpdateIntf(nextProps, nextState interface{}) bool {
	return e.ShouldComponentUpdate(nextState.(ExamplesState))
}

// SetState is an auto-generated proxy proxy to update the state for the
// Examples component.  SetState does not immediately mutate e.State()
// but creates a pending state transition.
func (e *ExamplesDef) SetState(s ExamplesState) {
	e.ComponentDef.SetState(s)
}

// State is an auto-generated proxy to return the current state in use for the
// render of the Examples component
func (e *ExamplesDef) State() ExamplesState {
	return e.ComponentDef.State().(ExamplesState)
}

// IsState is an auto-generated definition so that ExamplesState implements
// the github.com/myitcv/gopherjs/react.State interface.
func (e ExamplesState) IsState() {}

var _ react.State = ExamplesState{}

// GetInitialStateIntf is an auto-generated proxy to GetInitialState
func (e *ExamplesDef) GetInitialStateIntf() react.State {
	return e.GetInitialState()
}
