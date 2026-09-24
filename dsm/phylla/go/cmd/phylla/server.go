//go:build !js

package main

import (
	"embed"
	"log"
	"strconv"

	"github.com/fullstack-lang/gong/dsm/phylla/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
	threejs_clock "github.com/fullstack-lang/gong/dsm/phylla/go/visual/threejs/clock"
	threejs_plant "github.com/fullstack-lang/gong/dsm/phylla/go/visual/threejs/plant"
	threejs_stool "github.com/fullstack-lang/gong/dsm/phylla/go/visual/threejs/stool"
	threejs_vase "github.com/fullstack-lang/gong/dsm/phylla/go/visual/threejs/vase"
)

//go:embed data/*
var dataFS embed.FS

func executeServer() {

	// setup
	models.DataFS = &dataFS
	// - model level1 stack with its probe and stageset probe
	// - unmarshall/marshall go file with stage data
	stack := level1stack.NewLevel1StackStageSetDelta("phylla", unmarshallFromCode, marshallOnCommit, true, embeddedDiagrams, true)

	// refresh the probe, therefore we can see what has been unmarshalled
	stack.Probe.Refresh()
	if stack.StageSetProbe != nil {
		stack.StageSetProbe.Refresh()
	}

	// initiates the UX loop
	models.NewStagerStageSet(
		stack.R,
		stack.StageSet,
		stack.Probe,
		marshallOnCommit,
		threejs_vase.NewThreeJSStageUpdater(),
		threejs_stool.NewStool3DStageUpdater(),
		threejs_clock.NewClock3DStageUpdater(),
		threejs_plant.NewPlant3DStageUpdater(),
	)

	log.Println("Server ready serve on localhost:" + strconv.Itoa(port))
	err := stack.Run(":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
