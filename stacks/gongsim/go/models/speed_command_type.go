package models

// swagger:enum GongsimCommandType
type SpeedCommandType string

const (
	COMMAND_INCREASE_SPEED_100_PERCENTS SpeedCommandType = "INCREASE_SPEED_100_PERCENTS"
	COMMAND_DECREASE_SPEED_50_PERCENTS  SpeedCommandType = "COMMAND_DECREASE_SPEED_50_PERCENTS "
	COMMAND_SPEED_STEADY                SpeedCommandType = "COMMAND_SPEED_STEADY"
)
