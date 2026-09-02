//go:build !js

package main

import (
	"embed"
	"log"
	"strconv"

	"github.com/fullstack-lang/gong/dsm/phylla/go/clockstage3d"
	"github.com/fullstack-lang/gong/dsm/phylla/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
	"github.com/fullstack-lang/gong/dsm/phylla/go/plantstage3d"
	"github.com/fullstack-lang/gong/dsm/phylla/go/stoolstage3d"
	"github.com/fullstack-lang/gong/dsm/phylla/go/vasestage3d"
)

//go:embed data/*
var dataFS embed.FS

func executeServer() {

	// setup
	models.DataFS = &dataFS
	// - model level1 stack with its probe
	// - unmarshall/marshall go file with stage data
	stack := level1stack.NewLevel1Stack("phylla", unmarshallFromCode, marshallOnCommit, true, embeddedDiagrams)

	// refresh the probe, therefore we can see what has been unmarshalled
	stack.Probe.Refresh()

	// initiates the UX loop
	models.NewStager(
		stack.R,
		stack.Stage,
		stack.Probe,
		marshallOnCommit,
		vasestage3d.NewThreeJSStageUpdater(),
		stoolstage3d.NewStool3DStageUpdater(),
		clockstage3d.NewClock3DStageUpdater(),
		plantstage3d.NewPlant3DStageUpdater(),
	)

	log.Println("Server ready serve on localhost:" + strconv.Itoa(port))
	err := stack.R.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
