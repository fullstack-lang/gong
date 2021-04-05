package models

// GongsimStatus is the struct of the instance that is updated by the front for issuing Statuss
// swagger:model GongsimStatus
type GongsimStatus struct {
	Name                       string
	CurrentCommand             GongsimCommandType
	CompletionDate             string
	CurrentSpeedCommand        SpeedCommandType
	SpeedCommandCompletionDate string
}

var GongsimStatusSingloton = (&GongsimStatus{
	Name:                       "Gongsim Status Singloton",
	CurrentCommand:             COMMAND_PAUSE,
	CurrentSpeedCommand: COMMAND_SPEED_STEADY,
	CompletionDate:             "",
	SpeedCommandCompletionDate: "",
}).Stage()
