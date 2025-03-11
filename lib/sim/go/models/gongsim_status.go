package models

// Status is the struct of the instance that is updated by the front for issuing Statuss
// swagger:model Status
type Status struct {
	Name                       string
	CurrentCommand             CommandType
	CompletionDate             string
	CurrentSpeedCommand        SpeedCommandType
	SpeedCommandCompletionDate string
}
