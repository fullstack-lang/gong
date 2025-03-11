package models

import (
	"log"
	"time"
)

func (command *Command) OnAfterUpdate(stage *StageStruct, stagedCommand, frontCommand *Command) {

	log.Println(time.Now().Format("2006-01-02 15:04:05.000000"), "received command update",
		frontCommand.Command.ToString())

	// force the command on the stage
	stagedCommand.Command = frontCommand.Command

	switch command.Command {
	case DECREASE_SPEED_50_PERCENTS:
		command.Engine.Speed *= 0.5
	case INCREASE_SPEED_100_PERCENTS:
		command.Engine.Speed *= 2.0
	}
}
