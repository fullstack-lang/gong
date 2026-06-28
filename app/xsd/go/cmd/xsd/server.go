//go:build !js

package main

import (
	"log"
	"strconv"

	"github.com/fullstack-lang/gong/app/xsd/go/level1stack"
	"github.com/fullstack-lang/gong/app/xsd/go/models"
)

func executeServer(args []string) {
	// setup
	// - model level1 stack with its probe
	// - unmarshall/marshall go file with stage data
	stack := level1stack.NewLevel1Stack("xsd", "", "", true, embeddedDiagrams)

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
