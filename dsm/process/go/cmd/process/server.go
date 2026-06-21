//go:build !js

package main

import (
	"log"
	"strconv"

	"github.com/fullstack-lang/gong/dsm/process/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/process/go/models"
)

func executeServer(args []string) {
	// args contains all arguments remaining after flags are parsed
	if len(args) > 0 {
		argument := args[0] // This will correctly pick 'foo.go'
		marshallOnCommit = argument
		unmarshallFromCode = argument
	}

	// setup
	// - model level1 stack with its probe
	// - unmarshall/marshall go file with stage data
	stack := level1stack.NewLevel1StackDelta("process", unmarshallFromCode, marshallOnCommit, true, embeddedDiagrams, true)
	// stack.Stage.SetGongMarshallingMode(models.GongMarshallingAppendCommit)
	stack.Stage.SetIsWithGenesisCommit(true) // the genesis commit is the first commit of the stage, it is the one that contains the initial data. It cannot be rollbacked, therefore when the user is at the genesis commit, the backward button is disabled

	// refresh the probe, therefore we can see what has been unmarshalled
	stack.Probe.Refresh()

	// initiates the UX loop
	models.NewStager(
		stack.R,
		stack.Stage,
		stack.Probe,
	)

	log.Println("Server ready serve on localhost:" + strconv.Itoa(port))
	err := stack.R.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
