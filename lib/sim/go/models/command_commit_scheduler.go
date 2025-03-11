package models

import "time"

// simulationEventForLastEngineCommit recount what was the simulation event during the last checkin
var simulationEventForLastEngineCommit int

// The "commit scheduler" is in charge of asking for a commit of the stage to the back repo.
//
// The commit is performed only if the event number of the engine has increased. Since the commit
// have to happend when the simulation is not advancing, the "commit scheduler" only schedule the "engine driver"
// to commit the simlation stage when it will be ready.
func (command *Command) commitScheduler() {

	// The period shall not too short for performance sake but not too long because the end user needs a responsive application
	//
	// commitSchedulerPeriod is the period of the "commit scheduler"
	var CommitSchedulerPeriod = time.NewTicker(100 * time.Millisecond)
	simulationEventForLastEngineCommit = command.Engine.Fired
	for {
		select {
		case t := <-CommitSchedulerPeriod.C:

			_ = t
			// log.Println("commitScheduler  timestamp  at ", t)

			if simulationEventForLastEngineCommit != command.Engine.Fired {
				command.Engine.nextCommitDate = time.Now()
			}
		}
	}
}
