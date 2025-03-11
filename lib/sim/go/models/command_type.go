package models

// swagger:enum CommandType
type CommandType string

// values for EnumType
const (
	COMMAND_PLAY                          CommandType = "PLAY"
	COMMAND_PAUSE                         CommandType = "PAUSE"
	COMMAND_FIRE_NEXT_EVENT               CommandType = "FIRE_NEXT_EVENT"
	COMMAND_FIRE_EVENT_TILL_STATES_CHANGE CommandType = "FIRE_EVENT_TILL_STATES_CHANGE"
	COMMAND_RESET                         CommandType = "RESET"
	COMMAND_ADVANCE_10_MIN                CommandType = "ADVANCE_10_MIN"
	INCREASE_SPEED_100_PERCENTS           CommandType = "INCREASE_SPEED_100_PERCENTS"
	DECREASE_SPEED_50_PERCENTS            CommandType = "DECREASE_SPEED_50_PERCENTS"
)
