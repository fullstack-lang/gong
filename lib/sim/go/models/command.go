package models

import (
	"log"
	"time"
)

// Command is the struct of the instance that is updated by the front for issuing commands
// swagger:model Command
type Command struct {
	Name        string
	Command     CommandType
	CommandDate string
	Engine      *Engine

	stage  *Stage
	status *Status
}

func NewCommand(stage *Stage, engine *Engine) (command *Command) {
	command = &(Command{
		Name:        "Gongsim Command Singloton",
		Command:     COMMAND_PAUSE,
		CommandDate: "",
		Engine:      engine,
		stage:       stage,
	})

	command.status = (&Status{
		Name:                       "Gongsim Status Singloton",
		CurrentCommand:             COMMAND_PAUSE,
		CurrentSpeedCommand:        COMMAND_SPEED_STEADY,
		CompletionDate:             "",
		SpeedCommandCompletionDate: "",
	}).Stage(stage)

	command.Stage(stage).SetupGongsimThreads()
	return
}

//
// SetupGongsimThreads enables Command to periodicaly watch the Command in the Repo
//
// It is set up with three threads:
// - The "command pooler"
// - The "commit scheduler"
// - The "engine driver"

var Quit chan bool

// SetupGongsimThreads schedules gongsim threads
func (command *Command) SetupGongsimThreads() *Command {

	Quit = make(chan bool)

	go command.commitScheduler()
	go command.checkoutScheduler()
	go command.watcher()

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
			select {
			case <-Quit:
				return
			default:
				nextState = UNKOWN
				_, nextSimTime, _ := command.Engine.GetNextEvent()
				var currentTimePlus10Minute time.Time

				if command.Engine.nextCommitDate.After(command.Engine.lastCommitDate) {
					// log.Printf("Commit agent states scheduled at event # %d and commit nb # %d",
					// 	command.Engine.Fired,
					// 	command.Engine.GetLastCommitNb())
					nextState = COMMIT_AGENT_STATES
				} else {
					if command.Engine.nextCheckoutDate.After(command.Engine.lastCheckoutDate) {
						// log.Printf("Checkout agent states scheduled at event # %d and commit nb # %d",
						// 	command.Engine.Fired,
						// 	command.Engine.GetLastCommitNb())
						nextState = CHECKOUT_AGENT_STATES
					} else if nextSimTime.Before(command.Engine.GetEndTime()) {
						switch command.Command {
						case COMMAND_PLAY:
							nextState = FIRE_ONE_EVENT
							nextMode = RELATIVE_SPEED
						case COMMAND_ADVANCE_10_MIN:
							if commandCompletionDate != command.CommandDate {
								nextState = FIRE_ONE_EVENT
								nextMode = FULL_SPEED
								engineStopMode = TEN_MINUTES
								currentTimePlus10Minute = nextSimTime.Add(time.Minute * 10)
								commandCompletionDate = command.CommandDate
							} else {
								nextState = SLEEP_100_MS
							}
						case COMMAND_FIRE_EVENT_TILL_STATES_CHANGE:
							if commandCompletionDate != command.CommandDate {
								nextState = FIRE_ONE_EVENT
								nextMode = FULL_SPEED
								engineStopMode = STATE_CHANGED
								commandCompletionDate = command.CommandDate
							} else {
								nextState = SLEEP_100_MS
							}
						case COMMAND_RESET:
							if commandCompletionDate != command.CommandDate {
								nextState = RESET_SIMULATION
								commandCompletionDate = command.CommandDate
							} else {
								nextState = SLEEP_100_MS
							}
						default:
							nextState = SLEEP_100_MS
						}
					} else {
						switch command.Command {
						case COMMAND_RESET:
							if commandCompletionDate != command.CommandDate {
								nextState = RESET_SIMULATION
								commandCompletionDate = command.CommandDate
							} else {
								nextState = SLEEP_100_MS
							}
						default:
							nextState = SLEEP_100_MS
						}
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
					for _, agent := range command.Engine.Agents() {
						agent.Reset()
					}

					// call specific Reset function
					if command.Engine.Simulation != nil {
						command.Engine.Simulation.Reset(command.Engine)
					}
					// commit the engine state
					command.stage.Commit()
				case COMMIT_AGENT_STATES:
					if command.Engine.Simulation != nil {
						command.Engine.Simulation.CommitAgents(command.Engine)
						command.Engine.lastCommitDate = command.Engine.nextCommitDate
						simulationEventForLastEngineCommit = command.Engine.Fired
						commitFromFrontNbAfterLastEngineCommitOrCheckout = command.Engine.GetLastCommitNbFromFront()
					}
				case CHECKOUT_AGENT_STATES:
					if command.Engine.Simulation != nil {
						command.Engine.Simulation.CheckoutAgents(command.Engine)
						command.Engine.lastCheckoutDate = command.Engine.nextCheckoutDate
						commitFromFrontNbAfterLastEngineCommitOrCheckout = command.Engine.GetLastCommitNbFromFront()
					}
				case FIRE_ONE_EVENT:
					if command.Engine.State != RUNNING {
						command.Engine.State = RUNNING
						command.stage.Commit()
					}

					if nextMode == RELATIVE_SPEED {

						realtimeDurationBetweenHorizons := 500 * time.Millisecond

						command.Engine.nextRealtimeHorizon = time.Now().Add(realtimeDurationBetweenHorizons)
						command.Engine.nextSimulatedTimeHorizon =
							command.Engine.currentTime.Add(time.Duration(int64(command.Engine.Speed * float64(realtimeDurationBetweenHorizons))))

						for nextSimTime.Before(command.Engine.nextSimulatedTimeHorizon) {
							var agent AgentInterface
							agent, nextSimTime, _ = command.Engine.FireNextEvent()

							if agent == nil {
								return
							}
						}
						command.Engine.currentTime = command.Engine.nextSimulatedTimeHorizon

						sleepTime := command.Engine.nextRealtimeHorizon.Sub(time.Now())
						time.Sleep(sleepTime)

						command.stage.Commit()

						// log.Printf(lastSimTime.String() + " " + nextSimTime.String())
						// if nextSimTime.Sub(command.Engine.GetCurrentTime()) > 0 {
						// 	simTimeAdvance := nextSimTime.Sub(command.Engine.GetCurrentTime())

						// 	sleepDuration := time.Duration(float64(simTimeAdvance) / command.Engine.Speed)
						// 	// log.Printf("total sleep duration " + sleepDuration.String())

						// 	// in order for the end user to see progress in the simulation time
						// 	// the egine time is updated periodicaly
						// 	maxSleepAtATime := time.Duration(500 * time.Millisecond)

						// 	cumulatedSleepTime := time.Duration(0 * time.Millisecond)
						// 	for cumulatedSleepTime < sleepDuration && command.Command == COMMAND_PLAY {
						// 		sleepTime := Min(sleepDuration-cumulatedSleepTime, maxSleepAtATime)
						// 		// log.Printf("Stepped sleep time " + sleepTime.String() + " cumulated " + cumulatedSleepTime.String())
						// 		time.Sleep(sleepTime)
						// 		cumulatedSleepTime += sleepTime

						// 		// update engine current time
						// 		progressInSimulatedTimeInMiliseconds := command.Engine.Speed *
						// 			float64(sleepTime.Milliseconds())
						// 		command.Engine.SetCurrentTime(command.Engine.GetCurrentTime().Add(
						// 			time.Duration(progressInSimulatedTimeInMiliseconds) * time.Millisecond))
						// 		// log.Printf("Engine current time " + command.Engine.CurrentTime.String())
						// 		command.Engine.Commit(command.stage)
						// 	}
						// }
						// _, nextSimTime, _ = command.Engine.FireNextEvent()
					} else { // FULL SPEED
						if engineStopMode == TEN_MINUTES {
							for nextSimTime.Before(currentTimePlus10Minute) {
								var agent AgentInterface
								agent, nextSimTime, _ = command.Engine.FireNextEvent()

								if agent == nil {
									return
								}
							}
						} else if engineStopMode == STATE_CHANGED {
							hasAnyStateHasChanged := false
							for nextSimTime.Before(command.Engine.GetEndTime()) && !hasAnyStateHasChanged {

								var agent AgentInterface
								agent, nextSimTime, _ = command.Engine.FireNextEvent()

								if agent == nil {
									return
								}

								if command.Engine.Simulation != nil {
									hasAnyStateHasChanged = command.Engine.Simulation.HasAnyStateChanged(command.Engine)
								}
							}
						} else {
							log.Panicf("should not reach this piont")
						}

						// time has progressed, therefore an update is necessary
						command.stage.Commit()
					}

				case SLEEP_100_MS:
					if command.Engine.State != PAUSED {
						command.Engine.State = PAUSED
						command.stage.Commit()
					}
					time.Sleep(time.Duration(100 * time.Millisecond))
				default:
				}
			}
		}
	}()

	return command
}

// Min returns the larger of x or y.
func Min(x, y time.Duration) time.Duration {
	if x > y {
		return y
	}
	return x
}
