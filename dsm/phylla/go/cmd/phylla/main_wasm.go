//go:build js && wasm

package main

import (
	"embed"
	"log"

	"github.com/fullstack-lang/gong/dsm/phylla/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/phylla/go/models"
	threejs_clock "github.com/fullstack-lang/gong/dsm/phylla/go/visual/threejs/clock"
	threejs_plant "github.com/fullstack-lang/gong/dsm/phylla/go/visual/threejs/plant"
	threejs_stool "github.com/fullstack-lang/gong/dsm/phylla/go/visual/threejs/stool"
	threejs_vase "github.com/fullstack-lang/gong/dsm/phylla/go/visual/threejs/vase"
	"github.com/fullstack-lang/gong/lib/wasmregistry"
)

//go:embed data/*
var dataFS embed.FS

func main() {
	log.SetOutput(&wasmregistry.ConsoleWriter{})
	log.SetPrefix("phylla: ")
	log.SetFlags(log.Lmicroseconds)

	log.Println("Initializing phylla WASM Backend...")

	// setup
	models.DataFS = &dataFS

	unmarshallFromCode := ""
	marshallOnCommit := ""
	embeddedDiagrams := true

	// setup
	// - model level1 stack with its probe and stageset probe
	// - unmarshall/marshall go file with stage data
	stack := level1stack.NewLevel1StackStageSetDelta("phylla", unmarshallFromCode, marshallOnCommit, true, embeddedDiagrams, true)
	stack.Stage.SetGongMarshallingMode(models.GongMarshallingAppendCommit)
	stack.Stage.SetIsWithGenesisCommit(true) // the genesis commit is the first commit of the stage, it is the one that contains the initial data. It cannot be rollbacked.

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

	// Expose the HTTP and Socket bridges to the Angular frontend
	wasmregistry.SetupWasmHooks(stack.R)

	select {} // Keep the WASM instance running indefinitely
}
