//go:build !js

package main

import (
	"strconv"

	"log"

	// Import the time package

	// Use alias 'meta' to avoid conflict
	// <-- Import the standard extension package

	// Import the html renderer package

	// YAML parsing is handled by goldmark-meta

	// insertion point for models import

	ssg_level1stack "github.com/fullstack-lang/gong/lib/ssg/go/level1stack"
	ssg_models "github.com/fullstack-lang/gong/lib/ssg/go/models"
)

func executeServer(args []string) {
	// setup model stack with its probe
	stack := ssg_level1stack.NewLevel1Stack("ssg", unmarshallFromCode, marshallOnCommit, true, embeddedDiagrams)
	stack.Probe.Refresh()

	// insertion point for call to stager
	ssg_models.NewStager(stack.R, stack.Stage)

	log.Println("Server ready serve on localhost:" + strconv.Itoa(port))
	err := stack.R.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
