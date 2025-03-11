package models

import "time"

// commitFromFrontNbAfterLastEngineCommitOrCheckout recounts what was the commit nb after the last checkin.
// It is usefull because if the commit nb has increased since the last engine checking, this
// means the front repo has performed a commit and the simulation should performed a checkout
// in order for the stage to update it state vectors
var commitFromFrontNbAfterLastEngineCommitOrCheckout uint

// The "checkout scheduler" is in charge of asking for a checkout of the back repo to the stage.
//
// # This can happen when an user updates manualy update the state vector of an agent by using the GUI
//
// The checkout is performed only if both conditions are met:
//   - the event number of the engine has not increased (if it is idle for instance)
//   - the commitNb from the frontof the simulation has increased since the first condition is met
//
// Since the checkout
// have to happend when the simulation is not advancing, the "checkout scheduler" only schedule the "engine driver"
// to checkout the simlation stage when it will be ready.
func (command *Command) checkoutScheduler() {

	// The period shall not too short for performance sake but not too long because the end user needs a responsive application
	//
	// checkoutSchedulerPeriod is the period of the "checkout scheduler"
	var CheckoutSchedulerPeriod = time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case t := <-CheckoutSchedulerPeriod.C:

			_ = t

			if simulationEventForLastEngineCommit == command.Engine.Fired &&
				commitFromFrontNbAfterLastEngineCommitOrCheckout < command.Engine.GetLastCommitNbFromFront() {
				command.Engine.nextCheckoutDate = time.Now()
			}
		}
	}
}
