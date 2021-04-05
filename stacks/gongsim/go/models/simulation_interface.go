package models

// SimulationInterface is the callback support for
// events that happens on the generic engine
type SimulationInterface interface {

	// HasAnyStateChanged provides the generic engine infos about the behavior
	// of simulation
	//
	// It is used in order to have the "advance till one state changes" call
	//
	// the simulation shall implement this callback
	// it returns true if one state of the specific has changed
	HasAnyStateChanged(engine *Engine) bool

	// CommitAgents requests the simulation to commit the instances values to the back repo
	//
	// This is used when the simulation states has to be read by an outside component (the front component for instance)
	CommitAgents(engines *Engine)

	// CheckoutAgents requests the simulation to checkout the instances values from the back repo
	//
	// This is used when the simulation states has to be updated by values set by an outside component (the front component for instance)
	CheckoutAgents(engine *Engine)

	// GetLastCommitNb fetch the last commit nb from the simulation
	// this enables the engine to compute when it is necessary to update the stage following
	// updates to the back repo from the front
	GetLastCommitNb() uint

	// Reset simulation
	Reset(engine *Engine)
}
