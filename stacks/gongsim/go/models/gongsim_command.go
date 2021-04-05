package models

import (
	"log"
	"time"
)

// GongsimCommand is the struct of the instance that is updated by the front for issuing commands
// swagger:model GongsimCommand
type GongsimCommand struct {
	Name             string
	Command          GongsimCommandType
	CommandDate      string
	SpeedCommandType SpeedCommandType
	DateSpeedCommand string
}

//
// the states of the engine drivers
type EngineDriverState int

const (
	COMMIT_AGENT_STATES EngineDriverState = iota
	CHECKOUT_AGENT_STATES
	FIRE_ONE_EVENT
	SLEEP_100_MS
	RESET_SIMULATION
	UNKOWN
)

// the simulation can run with a relative speed to realtime or full speed
type EngineRunMode int

const (
	RELATIVE_SPEED EngineRunMode = iota
	FULL_SPEED
)

type EngineStopMode int

const (
	TEN_MINUTES EngineStopMode = iota
	STATE_CHANGED
)

var GongsimCommandSingloton = (&GongsimCommand{
	Name:        "Gongsim Command Singloton",
	Command:     COMMAND_PAUSE,
	CommandDate: "",
}).Stage().SetupGongsimThreads()

//
// SetupGongsimThreads enables GongsimCommand to periodicaly watch the GongsimCommand in the Repo
//
// It is set up with three threads:
// - The "command pooler"
// - The "commit scheduler"
// - The "engine driver"

// commandPooler" is a simple scheduled thread that manages the checkout
// of the current command from the back repo to the stage. If the front client writes a "PLAY" or "PAUSE"
// command, it is written to the back repo and the command pooler retrieves this command to the stage where it
// can be read by the "engine driver". The "command pooler" is a scheduled task
func (gongsimCommand *GongsimCommand) commandPooler() {

	// commandPoolerPeriod is the period of the "command pooler"
	//
	// The period shall not too short for performance sake but not too long because the end user needs a responsive application
	var CommandPoolerPeriod = time.NewTicker(500 * time.Millisecond)

	for {
		select {
		case t := <-CommandPoolerPeriod.C:

			gongsimCommand.Checkout()
			if GongsimStatusSingloton.CompletionDate != gongsimCommand.CommandDate {
				log.Println("commandPooler reads new command ", gongsimCommand.Command, "  timestamp ", gongsimCommand.CommandDate, " at ", t)

				GongsimStatusSingloton.CurrentCommand = gongsimCommand.Command
				GongsimStatusSingloton.CompletionDate = gongsimCommand.CommandDate
				GongsimStatusSingloton.Commit()
			}
			if GongsimStatusSingloton.SpeedCommandCompletionDate != gongsimCommand.DateSpeedCommand {
				log.Println("commandPooler reads new speed command ", gongsimCommand.SpeedCommandType, "  timestamp ", gongsimCommand.CommandDate, " at ", t)

				switch gongsimCommand.SpeedCommandType {
				case COMMAND_DECREASE_SPEED_50_PERCENTS:
					EngineSingloton.Speed *= 0.5
					EngineSingloton.Commit()
				case COMMAND_INCREASE_SPEED_100_PERCENTS:
					EngineSingloton.Speed *= 2.0
					EngineSingloton.Commit()
				}

				GongsimStatusSingloton.CurrentSpeedCommand = gongsimCommand.SpeedCommandType
				GongsimStatusSingloton.SpeedCommandCompletionDate = gongsimCommand.DateSpeedCommand
				GongsimStatusSingloton.Commit()
			}
		}
	}
}

// simulationEventForLastEngineCommit recount what was the simulation event during the last checkin
var simulationEventForLastEngineCommit int

// The "commit scheduler" is in charge of asking for a commit of the stage to the back repo.
//
// The commit is performed only if the event number of the engine has increased. Since the commit
// have to happend when the simulation is not advancing, the "commit scheduler" only schedule the "engine driver"
// to commit the simlation stage when it will be ready.
func (gongsimCommand *GongsimCommand) commitScheduler() {

	// The period shall not too short for performance sake but not too long because the end user needs a responsive application
	//
	// commitSchedulerPeriod is the period of the "commit scheduler"
	var CommitSchedulerPeriod = time.NewTicker(500 * time.Millisecond)
	simulationEventForLastEngineCommit = EngineSingloton.Fired
	for {
		select {
		case t := <-CommitSchedulerPeriod.C:

			_ = t
			// log.Println("commitScheduler  timestamp  at ", t)

			if simulationEventForLastEngineCommit != EngineSingloton.Fired {
				EngineSingloton.nextCommitDate = time.Now()
			}
		}
	}
}

// simulationCommitNbAfterLastEngineCommitOrCheckout recounts what was the commit nb after the last checkin.
// It is usefull because if the commit nb has increased since the last engine checking, this
// means the front repo has performed a commit and the simulation should performed a checkout
// in order for the stage to update it state vectors
var simulationCommitNbAfterLastEngineCommitOrCheckout uint

// The "checkout scheduler" is in charge of asking for a checkout of the back repo to the stage.
//
// The checkout is performed only if both conditions are met:
//  - the event number of the engine has not increased (if it is idle for instance)
//  - the commitNb of the simulation backRepo has increased since the first condition is met
// Since the checkout
// have to happend when the simulation is not advancing, the "checkout scheduler" only schedule the "engine driver"
// to checkout the simlation stage when it will be ready.
func (gongsimCommand *GongsimCommand) checkoutScheduler() {

	// The period shall not too short for performance sake but not too long because the end user needs a responsive application
	//
	// checkoutSchedulerPeriod is the period of the "checkout scheduler"
	var CheckoutSchedulerPeriod = time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case t := <-CheckoutSchedulerPeriod.C:

			_ = t

			if simulationEventForLastEngineCommit == EngineSingloton.Fired &&
				simulationCommitNbAfterLastEngineCommitOrCheckout < EngineSingloton.GetLastCommitNb() {
				EngineSingloton.nextCheckoutDate = time.Now()
			}
		}
	}
}

// SetupGongsimThreads schedules gongsim threads
func (gongsimCommand *GongsimCommand) SetupGongsimThreads() *GongsimCommand {

	go gongsimCommand.commandPooler()
	go gongsimCommand.commitScheduler()
	go gongsimCommand.checkoutScheduler()

	go func() {

		var nextState EngineDriverState
		var nextMode EngineRunMode
		var engineStopMode EngineStopMode

		// when the engine receive a ADVANCE_10_MIN or FIRE_TILL_STATE_CHANGE
		// it can take into account this command only once
		// commandCompletionDate records at which date the command was completed
		var commandCompletionDate string
		for {

			//
			// compute engine driver next state
			//
			// next state is either
			/*

				RESET_SIMULATION
				COMMIT_AGENT_STATES
				CHECKOUT_AGENT_STATES
				FIRE_ONE_EVENT
				SLEEP_100_MS


			*/
			//
			//

			nextState = UNKOWN
			_, nextSimTime, _ := EngineSingloton.GetNextEvent()
			var currentTimePlus10Minute time.Time

			if EngineSingloton.nextCommitDate.After(EngineSingloton.lastCommitDate) {
				// log.Printf("Commit agent states scheduled at event # %d and commit nb # %d",
				// 	EngineSingloton.Fired,
				// 	EngineSingloton.GetLastCommitNb())
				nextState = COMMIT_AGENT_STATES
			} else {
				if EngineSingloton.nextCheckoutDate.After(EngineSingloton.lastCheckoutDate) {
					log.Printf("Checkout agent states scheduled at event # %d and commit nb # %d",
						EngineSingloton.Fired,
						EngineSingloton.GetLastCommitNb())
					nextState = CHECKOUT_AGENT_STATES
				} else if nextSimTime.Before(EngineSingloton.GetEndTime()) {
					switch gongsimCommand.Command {
					case COMMAND_PLAY:
						nextState = FIRE_ONE_EVENT
						nextMode = RELATIVE_SPEED
					case COMMAND_ADVANCE_10_MIN:
						if commandCompletionDate != gongsimCommand.CommandDate {
							nextState = FIRE_ONE_EVENT
							nextMode = FULL_SPEED
							engineStopMode = TEN_MINUTES
							currentTimePlus10Minute = nextSimTime.Add(time.Minute * 10)
							commandCompletionDate = gongsimCommand.CommandDate
						} else {
							nextState = SLEEP_100_MS
						}
					case COMMAND_FIRE_EVENT_TILL_STATES_CHANGE:
						if commandCompletionDate != gongsimCommand.CommandDate {
							nextState = FIRE_ONE_EVENT
							nextMode = FULL_SPEED
							engineStopMode = STATE_CHANGED
							commandCompletionDate = gongsimCommand.CommandDate
						} else {
							nextState = SLEEP_100_MS
						}
					case COMMAND_RESET:
						if commandCompletionDate != gongsimCommand.CommandDate {
							nextState = RESET_SIMULATION
							commandCompletionDate = gongsimCommand.CommandDate
						} else {
							nextState = SLEEP_100_MS
						}
					default:
						nextState = SLEEP_100_MS
					}
				} else {
					nextState = SLEEP_100_MS
				}
			}
			if nextState == UNKOWN {
				log.Panicln("Unkown end state for engine driver")
			}

			//
			// perform what is expected in the state
			//
			switch nextState {
			case RESET_SIMULATION:
				// reset all agents
				for _, agent := range EngineSingloton.Agents() {
					agent.Reset()
				}

				// call specific Reset function
				if EngineSingloton.Simulation != nil {
					EngineSingloton.Simulation.Reset(EngineSingloton)
				}
				// commit the engine state
				EngineSingloton.Commit()
			case COMMIT_AGENT_STATES:
				if EngineSingloton.Simulation != nil {
					EngineSingloton.Simulation.CommitAgents(EngineSingloton)
					EngineSingloton.lastCommitDate = EngineSingloton.nextCommitDate
					simulationEventForLastEngineCommit = EngineSingloton.Fired
					simulationCommitNbAfterLastEngineCommitOrCheckout = EngineSingloton.GetLastCommitNb()
				}
			case CHECKOUT_AGENT_STATES:
				if EngineSingloton.Simulation != nil {
					EngineSingloton.Simulation.CheckoutAgents(EngineSingloton)
					EngineSingloton.lastCheckoutDate = EngineSingloton.nextCheckoutDate
					simulationCommitNbAfterLastEngineCommitOrCheckout = EngineSingloton.GetLastCommitNb()
				}
			case FIRE_ONE_EVENT:
				if EngineSingloton.State != RUNNING {
					EngineSingloton.State = RUNNING
					EngineSingloton.Commit()
				}

				if nextMode == RELATIVE_SPEED {

					// log.Printf(lastSimTime.String() + " " + nextSimTime.String())
					if nextSimTime.Sub(EngineSingloton.GetCurrentTime()) > 0 {
						simTimeAdvance := nextSimTime.Sub(EngineSingloton.GetCurrentTime())

						sleepDuration := time.Duration(float64(simTimeAdvance) / EngineSingloton.Speed)
						// log.Printf("total sleep duration " + sleepDuration.String())

						// in order for the end user to see progress in the simulation time
						// the egine time is updated periodicaly
						maxSleepAtATime := time.Duration(500 * time.Millisecond)

						cumulatedSleepTime := time.Duration(0 * time.Millisecond)
						for cumulatedSleepTime < sleepDuration && gongsimCommand.Command == COMMAND_PLAY {
							sleepTime := Min(sleepDuration-cumulatedSleepTime, maxSleepAtATime)
							// log.Printf("Stepped sleep time " + sleepTime.String() + " cumulated " + cumulatedSleepTime.String())
							time.Sleep(sleepTime)
							cumulatedSleepTime += sleepTime

							// update engine current time
							progressInSimulatedTimeInMiliseconds := EngineSingloton.Speed *
								float64(sleepTime.Milliseconds())
							EngineSingloton.SetCurrentTime(EngineSingloton.GetCurrentTime().Add(
								time.Duration(progressInSimulatedTimeInMiliseconds) * time.Millisecond))
							// log.Printf("Engine current time " + EngineSingloton.CurrentTime.String())
							EngineSingloton.Commit()
						}
					}
					_, nextSimTime, _ = EngineSingloton.FireNextEvent()
				} else { // FULL SPEED
					if engineStopMode == TEN_MINUTES {
						for nextSimTime.Before(currentTimePlus10Minute) {
							_, nextSimTime, _ = EngineSingloton.FireNextEvent()
						}
					} else if engineStopMode == STATE_CHANGED {
						hasAnyStateHasChanged := false
						for nextSimTime.Before(EngineSingloton.GetEndTime()) && !hasAnyStateHasChanged {

							_, nextSimTime, _ = EngineSingloton.FireNextEvent()

							if EngineSingloton.Simulation != nil {
								hasAnyStateHasChanged = EngineSingloton.Simulation.HasAnyStateChanged(EngineSingloton)
							}
						}
					} else {
						log.Panicf("should not reach this piont")
					}

					// time has progressed, therefore an update is necessary
					EngineSingloton.Commit()
				}

			case SLEEP_100_MS:
				if EngineSingloton.State != PAUSED {
					EngineSingloton.State = PAUSED
					EngineSingloton.Commit()
				}
				time.Sleep(time.Duration(100 * time.Millisecond))
			default:
			}
		}
	}()

	return gongsimCommand
}

// Min returns the larger of x or y.
func Min(x, y time.Duration) time.Duration {
	if x > y {
		return y
	}
	return x
}
