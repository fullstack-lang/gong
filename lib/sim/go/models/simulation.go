package models

import (
	"gorm.io/gorm"
)

// the simplest possible simulation
type Simulation struct {
	DB *gorm.DB
}

func (specificEngine *Simulation) EventFired(engine *Engine)              {}
func (specificEngine *Simulation) HasAnyStateChanged(engine *Engine) bool { return true }
func (specificEngine *Simulation) Reset(engine *Engine)                   {}
func (specificEngine *Simulation) CommitAgents(engine *Engine)            {}
func (specificEngine *Simulation) CheckoutAgents(engine *Engine)          {}
func (specificEngine *Simulation) GetLastCommitNb() uint                  { return 0 }
func (specificEngine *Simulation) GetLastCommitNbFromFront() uint         { return 0 }

type DummyAgent struct {
	Agent

	Name string
}

func (dummyAgent *DummyAgent) FireNextEvent() {
	event, _ := dummyAgent.GetNextEventAndRemoveIt()

	switch event.(type) {
	case *UpdateState:
		checkStateEvent := event.(*UpdateState)

		// post next event
		checkStateEvent.SetFireTime(checkStateEvent.GetFireTime().Add(checkStateEvent.Period))
		dummyAgent.QueueEvent(checkStateEvent)
	}
}
