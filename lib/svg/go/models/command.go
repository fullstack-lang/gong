package models

type CommandType string

const (
	CommandTypeSVGInTheBack CommandType = "CommandTypeSVGInTheBack" // Ask the backend to generates command
)

// Command is a temporaty object that is created by the front
// to inform the back of a end user command
//
// once it is created, it is destroyed by the front
type Command struct {
	Name        string
	CommandType CommandType
}
