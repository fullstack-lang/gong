package models

// swagger:enum GongsimCommandType
type GongsimCommandType string

// values for EnumType
const (
	COMMAND_PLAY                          GongsimCommandType = "PLAY"
	COMMAND_PAUSE                         GongsimCommandType = "PAUSE"
	COMMAND_FIRE_NEXT_EVENT               GongsimCommandType = "FIRE_NEXT_EVENT"
	COMMAND_FIRE_EVENT_TILL_STATES_CHANGE GongsimCommandType = "FIRE_EVENT_TILL_STATES_CHANGE"
	COMMAND_RESET                         GongsimCommandType = "RESET"
	COMMAND_ADVANCE_10_MIN                GongsimCommandType = "ADVANCE_10_MIN"
)
