// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for CommandType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (commandtype CommandType) ToString() (res string) {

	// migration of former implementation of enum
	switch commandtype {
	// insertion code per enum code
	case COMMAND_PLAY:
		res = "PLAY"
	case COMMAND_PAUSE:
		res = "PAUSE"
	case COMMAND_FIRE_NEXT_EVENT:
		res = "FIRE_NEXT_EVENT"
	case COMMAND_FIRE_EVENT_TILL_STATES_CHANGE:
		res = "FIRE_EVENT_TILL_STATES_CHANGE"
	case COMMAND_RESET:
		res = "RESET"
	case COMMAND_ADVANCE_10_MIN:
		res = "ADVANCE_10_MIN"
	case INCREASE_SPEED_100_PERCENTS:
		res = "INCREASE_SPEED_100_PERCENTS"
	case DECREASE_SPEED_50_PERCENTS:
		res = "DECREASE_SPEED_50_PERCENTS"
	}
	return
}

func (commandtype *CommandType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "PLAY":
		*commandtype = COMMAND_PLAY
		return
	case "PAUSE":
		*commandtype = COMMAND_PAUSE
		return
	case "FIRE_NEXT_EVENT":
		*commandtype = COMMAND_FIRE_NEXT_EVENT
		return
	case "FIRE_EVENT_TILL_STATES_CHANGE":
		*commandtype = COMMAND_FIRE_EVENT_TILL_STATES_CHANGE
		return
	case "RESET":
		*commandtype = COMMAND_RESET
		return
	case "ADVANCE_10_MIN":
		*commandtype = COMMAND_ADVANCE_10_MIN
		return
	case "INCREASE_SPEED_100_PERCENTS":
		*commandtype = INCREASE_SPEED_100_PERCENTS
		return
	case "DECREASE_SPEED_50_PERCENTS":
		*commandtype = DECREASE_SPEED_50_PERCENTS
		return
	default:
		return errUnkownEnum
	}
}

func (commandtype *CommandType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "COMMAND_PLAY":
		*commandtype = COMMAND_PLAY
	case "COMMAND_PAUSE":
		*commandtype = COMMAND_PAUSE
	case "COMMAND_FIRE_NEXT_EVENT":
		*commandtype = COMMAND_FIRE_NEXT_EVENT
	case "COMMAND_FIRE_EVENT_TILL_STATES_CHANGE":
		*commandtype = COMMAND_FIRE_EVENT_TILL_STATES_CHANGE
	case "COMMAND_RESET":
		*commandtype = COMMAND_RESET
	case "COMMAND_ADVANCE_10_MIN":
		*commandtype = COMMAND_ADVANCE_10_MIN
	case "INCREASE_SPEED_100_PERCENTS":
		*commandtype = INCREASE_SPEED_100_PERCENTS
	case "DECREASE_SPEED_50_PERCENTS":
		*commandtype = DECREASE_SPEED_50_PERCENTS
	default:
		return errUnkownEnum
	}
	return
}

func (commandtype *CommandType) ToCodeString() (res string) {

	switch *commandtype {
	// insertion code per enum code
	case COMMAND_PLAY:
		res = "COMMAND_PLAY"
	case COMMAND_PAUSE:
		res = "COMMAND_PAUSE"
	case COMMAND_FIRE_NEXT_EVENT:
		res = "COMMAND_FIRE_NEXT_EVENT"
	case COMMAND_FIRE_EVENT_TILL_STATES_CHANGE:
		res = "COMMAND_FIRE_EVENT_TILL_STATES_CHANGE"
	case COMMAND_RESET:
		res = "COMMAND_RESET"
	case COMMAND_ADVANCE_10_MIN:
		res = "COMMAND_ADVANCE_10_MIN"
	case INCREASE_SPEED_100_PERCENTS:
		res = "INCREASE_SPEED_100_PERCENTS"
	case DECREASE_SPEED_50_PERCENTS:
		res = "DECREASE_SPEED_50_PERCENTS"
	}
	return
}

func (commandtype CommandType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "COMMAND_PLAY")
	res = append(res, "COMMAND_PAUSE")
	res = append(res, "COMMAND_FIRE_NEXT_EVENT")
	res = append(res, "COMMAND_FIRE_EVENT_TILL_STATES_CHANGE")
	res = append(res, "COMMAND_RESET")
	res = append(res, "COMMAND_ADVANCE_10_MIN")
	res = append(res, "INCREASE_SPEED_100_PERCENTS")
	res = append(res, "DECREASE_SPEED_50_PERCENTS")

	return
}

func (commandtype CommandType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "PLAY")
	res = append(res, "PAUSE")
	res = append(res, "FIRE_NEXT_EVENT")
	res = append(res, "FIRE_EVENT_TILL_STATES_CHANGE")
	res = append(res, "RESET")
	res = append(res, "ADVANCE_10_MIN")
	res = append(res, "INCREASE_SPEED_100_PERCENTS")
	res = append(res, "DECREASE_SPEED_50_PERCENTS")

	return
}

// Utility function for ControlMode
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (controlmode ControlMode) ToString() (res string) {

	// migration of former implementation of enum
	switch controlmode {
	// insertion code per enum code
	case AUTONOMOUS:
		res = "Autonomous"
	case CLIENT_CONTROL:
		res = "ClientControl"
	}
	return
}

func (controlmode *ControlMode) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Autonomous":
		*controlmode = AUTONOMOUS
		return
	case "ClientControl":
		*controlmode = CLIENT_CONTROL
		return
	default:
		return errUnkownEnum
	}
}

func (controlmode *ControlMode) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "AUTONOMOUS":
		*controlmode = AUTONOMOUS
	case "CLIENT_CONTROL":
		*controlmode = CLIENT_CONTROL
	default:
		return errUnkownEnum
	}
	return
}

func (controlmode *ControlMode) ToCodeString() (res string) {

	switch *controlmode {
	// insertion code per enum code
	case AUTONOMOUS:
		res = "AUTONOMOUS"
	case CLIENT_CONTROL:
		res = "CLIENT_CONTROL"
	}
	return
}

func (controlmode ControlMode) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "AUTONOMOUS")
	res = append(res, "CLIENT_CONTROL")

	return
}

func (controlmode ControlMode) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Autonomous")
	res = append(res, "ClientControl")

	return
}

// Utility function for EngineDriverState
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enginedriverstate EngineDriverState) ToInt() (res int) {

	// migration of former implementation of enum
	switch enginedriverstate {
	// insertion code per enum code
	case COMMIT_AGENT_STATES:
		res = 0
	case CHECKOUT_AGENT_STATES:
		res = 1
	case FIRE_ONE_EVENT:
		res = 2
	case SLEEP_100_MS:
		res = 3
	case RESET_SIMULATION:
		res = 4
	case UNKOWN:
		res = 5
	}
	return
}

func (enginedriverstate *EngineDriverState) FromInt(input int) (err error) {

	switch input {
	// insertion code per enum code
	case 0:
		*enginedriverstate = COMMIT_AGENT_STATES
		return
	case 1:
		*enginedriverstate = CHECKOUT_AGENT_STATES
		return
	case 2:
		*enginedriverstate = FIRE_ONE_EVENT
		return
	case 3:
		*enginedriverstate = SLEEP_100_MS
		return
	case 4:
		*enginedriverstate = RESET_SIMULATION
		return
	case 5:
		*enginedriverstate = UNKOWN
		return
	default:
		return errUnkownEnum
	}
}

func (enginedriverstate *EngineDriverState) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "COMMIT_AGENT_STATES":
		*enginedriverstate = COMMIT_AGENT_STATES
	case "CHECKOUT_AGENT_STATES":
		*enginedriverstate = CHECKOUT_AGENT_STATES
	case "FIRE_ONE_EVENT":
		*enginedriverstate = FIRE_ONE_EVENT
	case "SLEEP_100_MS":
		*enginedriverstate = SLEEP_100_MS
	case "RESET_SIMULATION":
		*enginedriverstate = RESET_SIMULATION
	case "UNKOWN":
		*enginedriverstate = UNKOWN
	default:
		return errUnkownEnum
	}
	return
}

func (enginedriverstate *EngineDriverState) ToCodeString() (res string) {

	switch *enginedriverstate {
	// insertion code per enum code
	case COMMIT_AGENT_STATES:
		res = "COMMIT_AGENT_STATES"
	case CHECKOUT_AGENT_STATES:
		res = "CHECKOUT_AGENT_STATES"
	case FIRE_ONE_EVENT:
		res = "FIRE_ONE_EVENT"
	case SLEEP_100_MS:
		res = "SLEEP_100_MS"
	case RESET_SIMULATION:
		res = "RESET_SIMULATION"
	case UNKOWN:
		res = "UNKOWN"
	}
	return
}

func (enginedriverstate EngineDriverState) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "COMMIT_AGENT_STATES")
	res = append(res, "CHECKOUT_AGENT_STATES")
	res = append(res, "FIRE_ONE_EVENT")
	res = append(res, "SLEEP_100_MS")
	res = append(res, "RESET_SIMULATION")
	res = append(res, "UNKOWN")

	return
}

func (enginedriverstate EngineDriverState) CodeValues() (res []int) {

	res = make([]int, 0)

	// insertion code per enum code
	res = append(res, 0)
	res = append(res, 1)
	res = append(res, 2)
	res = append(res, 3)
	res = append(res, 4)
	res = append(res, 5)

	return
}

// Utility function for EngineRunMode
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enginerunmode EngineRunMode) ToInt() (res int) {

	// migration of former implementation of enum
	switch enginerunmode {
	// insertion code per enum code
	case RELATIVE_SPEED:
		res = 0
	case FULL_SPEED:
		res = 1
	}
	return
}

func (enginerunmode *EngineRunMode) FromInt(input int) (err error) {

	switch input {
	// insertion code per enum code
	case 0:
		*enginerunmode = RELATIVE_SPEED
		return
	case 1:
		*enginerunmode = FULL_SPEED
		return
	default:
		return errUnkownEnum
	}
}

func (enginerunmode *EngineRunMode) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "RELATIVE_SPEED":
		*enginerunmode = RELATIVE_SPEED
	case "FULL_SPEED":
		*enginerunmode = FULL_SPEED
	default:
		return errUnkownEnum
	}
	return
}

func (enginerunmode *EngineRunMode) ToCodeString() (res string) {

	switch *enginerunmode {
	// insertion code per enum code
	case RELATIVE_SPEED:
		res = "RELATIVE_SPEED"
	case FULL_SPEED:
		res = "FULL_SPEED"
	}
	return
}

func (enginerunmode EngineRunMode) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "RELATIVE_SPEED")
	res = append(res, "FULL_SPEED")

	return
}

func (enginerunmode EngineRunMode) CodeValues() (res []int) {

	res = make([]int, 0)

	// insertion code per enum code
	res = append(res, 0)
	res = append(res, 1)

	return
}

// Utility function for EngineState
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enginestate EngineState) ToString() (res string) {

	// migration of former implementation of enum
	switch enginestate {
	// insertion code per enum code
	case RUNNING:
		res = "RUNNING"
	case PAUSED:
		res = "PAUSED"
	case OVER:
		res = "OVER"
	}
	return
}

func (enginestate *EngineState) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "RUNNING":
		*enginestate = RUNNING
		return
	case "PAUSED":
		*enginestate = PAUSED
		return
	case "OVER":
		*enginestate = OVER
		return
	default:
		return errUnkownEnum
	}
}

func (enginestate *EngineState) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "RUNNING":
		*enginestate = RUNNING
	case "PAUSED":
		*enginestate = PAUSED
	case "OVER":
		*enginestate = OVER
	default:
		return errUnkownEnum
	}
	return
}

func (enginestate *EngineState) ToCodeString() (res string) {

	switch *enginestate {
	// insertion code per enum code
	case RUNNING:
		res = "RUNNING"
	case PAUSED:
		res = "PAUSED"
	case OVER:
		res = "OVER"
	}
	return
}

func (enginestate EngineState) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "RUNNING")
	res = append(res, "PAUSED")
	res = append(res, "OVER")

	return
}

func (enginestate EngineState) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "RUNNING")
	res = append(res, "PAUSED")
	res = append(res, "OVER")

	return
}

// Utility function for EngineStopMode
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (enginestopmode EngineStopMode) ToInt() (res int) {

	// migration of former implementation of enum
	switch enginestopmode {
	// insertion code per enum code
	case TEN_MINUTES:
		res = 0
	case STATE_CHANGED:
		res = 1
	}
	return
}

func (enginestopmode *EngineStopMode) FromInt(input int) (err error) {

	switch input {
	// insertion code per enum code
	case 0:
		*enginestopmode = TEN_MINUTES
		return
	case 1:
		*enginestopmode = STATE_CHANGED
		return
	default:
		return errUnkownEnum
	}
}

func (enginestopmode *EngineStopMode) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "TEN_MINUTES":
		*enginestopmode = TEN_MINUTES
	case "STATE_CHANGED":
		*enginestopmode = STATE_CHANGED
	default:
		return errUnkownEnum
	}
	return
}

func (enginestopmode *EngineStopMode) ToCodeString() (res string) {

	switch *enginestopmode {
	// insertion code per enum code
	case TEN_MINUTES:
		res = "TEN_MINUTES"
	case STATE_CHANGED:
		res = "STATE_CHANGED"
	}
	return
}

func (enginestopmode EngineStopMode) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "TEN_MINUTES")
	res = append(res, "STATE_CHANGED")

	return
}

func (enginestopmode EngineStopMode) CodeValues() (res []int) {

	res = make([]int, 0)

	// insertion code per enum code
	res = append(res, 0)
	res = append(res, 1)

	return
}

// Utility function for SpeedCommandType
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (speedcommandtype SpeedCommandType) ToString() (res string) {

	// migration of former implementation of enum
	switch speedcommandtype {
	// insertion code per enum code
	case COMMAND_INCREASE_SPEED_100_PERCENTS:
		res = "INCREASE_SPEED_100_PERCENTS"
	case COMMAND_DECREASE_SPEED_50_PERCENTS:
		res = "COMMAND_DECREASE_SPEED_50_PERCENTS "
	case COMMAND_SPEED_STEADY:
		res = "COMMAND_SPEED_STEADY"
	}
	return
}

func (speedcommandtype *SpeedCommandType) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "INCREASE_SPEED_100_PERCENTS":
		*speedcommandtype = COMMAND_INCREASE_SPEED_100_PERCENTS
		return
	case "COMMAND_DECREASE_SPEED_50_PERCENTS ":
		*speedcommandtype = COMMAND_DECREASE_SPEED_50_PERCENTS
		return
	case "COMMAND_SPEED_STEADY":
		*speedcommandtype = COMMAND_SPEED_STEADY
		return
	default:
		return errUnkownEnum
	}
}

func (speedcommandtype *SpeedCommandType) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "COMMAND_INCREASE_SPEED_100_PERCENTS":
		*speedcommandtype = COMMAND_INCREASE_SPEED_100_PERCENTS
	case "COMMAND_DECREASE_SPEED_50_PERCENTS":
		*speedcommandtype = COMMAND_DECREASE_SPEED_50_PERCENTS
	case "COMMAND_SPEED_STEADY":
		*speedcommandtype = COMMAND_SPEED_STEADY
	default:
		return errUnkownEnum
	}
	return
}

func (speedcommandtype *SpeedCommandType) ToCodeString() (res string) {

	switch *speedcommandtype {
	// insertion code per enum code
	case COMMAND_INCREASE_SPEED_100_PERCENTS:
		res = "COMMAND_INCREASE_SPEED_100_PERCENTS"
	case COMMAND_DECREASE_SPEED_50_PERCENTS:
		res = "COMMAND_DECREASE_SPEED_50_PERCENTS"
	case COMMAND_SPEED_STEADY:
		res = "COMMAND_SPEED_STEADY"
	}
	return
}

func (speedcommandtype SpeedCommandType) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "COMMAND_INCREASE_SPEED_100_PERCENTS")
	res = append(res, "COMMAND_DECREASE_SPEED_50_PERCENTS")
	res = append(res, "COMMAND_SPEED_STEADY")

	return
}

func (speedcommandtype SpeedCommandType) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "INCREASE_SPEED_100_PERCENTS")
	res = append(res, "COMMAND_DECREASE_SPEED_50_PERCENTS ")
	res = append(res, "COMMAND_SPEED_STEADY")

	return
}

// end of insertion point for enum utility functions

type GongstructEnumStringField interface {
	Codes() []string
	CodeValues() []string
	ToString() string
}

type PointerToGongstructEnumStringField interface {
	FromCodeString(input string) (err error)
}

type GongstructEnumIntField interface {
	int | EngineDriverState | EngineRunMode | EngineStopMode
	Codes() []string
	CodeValues() []int
}

type PointerToGongstructEnumIntField interface {
	*EngineDriverState | *EngineRunMode | *EngineStopMode
	FromCodeString(input string) (err error)
}

// Last line of the template
