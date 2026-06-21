//go:build !js

package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/fullstack-lang/gong/lib/docx/go/level1stack"
	"github.com/fullstack-lang/gong/lib/docx/go/models"
)

func executeServer(args []string) {
	if len(args) == 0 || len(args) > 1 {
		flag.Usage()
		os.Exit(1)
	}

	// setup
	// - model level1 stack with its probe
	// - unmarshall/marshall go file with stage data
	stack := level1stack.NewLevel1Stack("docx", unmarshallFromCode, marshallOnCommit, true, embeddedDiagrams)

	// refresh the probe, therefore we can see what has been unmarshalled
	stack.Probe.Refresh()

	filename := args[0]

	// initiates the UX loop
	models.NewStager(
		stack.R,
		stack.Stage,
		stack.Probe,
		filename,
	)

	log.Println("Server ready serve on localhost:" + strconv.Itoa(port))
	err := stack.R.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
