//go:build !js

package main

import (
	"embed"
	"log"
	"os"
	"strconv"

	"github.com/fullstack-lang/gong/dsm/barrgraph/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/barrgraph/go/models"
)

//go:embed "data/cubism and abstract art.go"
var stageGo string

//go:embed data/*
var dataFS embed.FS

func executeServer(args []string) {
	// setup
	// - set embedded data to models
	models.DataFS = &dataFS
	// - model level1 stack with its probe
	// - unmarshall/marshall go file with stage data
	stack := level1stack.NewLevel1Stack("barrgraph", unmarshallFromCode, marshallOnCommit, true, embeddedDiagrams)

	// refresh the probe, therefore we can see what has been unmarshalled
	stack.Probe.Refresh()

	// initiates the UX loop
	models.NewStager(
		stack.R,
		stack.Stage,
		stack.Probe,
	)

	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = strconv.Itoa(port) // Fallback on flag
	}
	port, _ = strconv.Atoi(portStr)

	log.Println("Server ready serve on localhost:" + strconv.Itoa(port))
	err := stack.R.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
