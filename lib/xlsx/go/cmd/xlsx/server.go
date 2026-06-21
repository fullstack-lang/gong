//go:build !js

package main

import (
	"log"
	"strconv"

	level1stack "github.com/fullstack-lang/gong/lib/xlsx/go/level1stack"
	xlsx_models "github.com/fullstack-lang/gong/lib/xlsx/go/models"
)

func executeServer(args []string) {
	// setup stack
	stack := level1stack.NewLevel1Stack("gantt", unmarshallFromCode, marshallOnCommit, true, embeddedDiagrams)
	stage := stack.Stage

	
	if len(args) < 1 {
		log.Panic("there should be at least one file argument")
	}

	sampleXLFile := new(xlsx_models.XLFile).Stage(stage)
	sampleXLFile.Open(stage, args[0])

	if len(args) > 1 {
		sampleXLFile2 := new(xlsx_models.XLFile).Stage(stage)
		sampleXLFile2.Open(stage, args[1])
	}

	stage.Commit()

	stack.Probe.Refresh()

	NewStager(stack.R, stack.Stage)

	log.Println("Server ready serve on localhost:" + strconv.Itoa(port))
	err := stack.R.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
