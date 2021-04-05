package events

import "github.com/fullstack-lang/gong/stacks/gongsim/go/models"

// StartProgram is to start program
// machine goes from IDLE to RUNNING
type StartProgram struct {
	models.Event
}
