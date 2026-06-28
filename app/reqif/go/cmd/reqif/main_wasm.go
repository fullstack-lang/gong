//go:build js && wasm

package main

import (
	"log"

	"github.com/fullstack-lang/gong/app/reqif/go/level1stack"
	"github.com/fullstack-lang/gong/app/reqif/go/models"
	"github.com/fullstack-lang/gong/lib/wasmregistry"

	"github.com/fullstack-lang/gong/app/reqif/go/datatypes"
	"github.com/fullstack-lang/gong/app/reqif/go/exporter"
	"github.com/fullstack-lang/gong/app/reqif/go/namer"
	"github.com/fullstack-lang/gong/app/reqif/go/specifications"
	"github.com/fullstack-lang/gong/app/reqif/go/specobjects"
	"github.com/fullstack-lang/gong/app/reqif/go/specrelations"
	"github.com/fullstack-lang/gong/app/reqif/go/spectypes"
)

func main() {
	log.SetOutput(&wasmregistry.ConsoleWriter{})
	log.SetPrefix("reqif: ")
	log.SetFlags(log.Lmicroseconds)

	log.Println("Initializing reqif WASM Backend...")

	unmarshallFromCode := ""
	marshallOnCommit := ""
	embeddedDiagrams := true

	// setup
	// - model level1 stack with its probe
	// - unmarshall/marshall go file with stage data
	stack := level1stack.NewLevel1StackDelta("reqif", unmarshallFromCode, marshallOnCommit, true, embeddedDiagrams, true)
	stack.Stage.SetGongMarshallingMode(models.GongMarshallingAppendCommit)
	stack.Stage.SetIsWithGenesisCommit(true) // the genesis commit is the first commit of the stage, it is the one that contains the initial data. It cannot be rollbacked.

	// refresh the probe, therefore we can see what has been unmarshalled
	stack.Probe.Refresh()

	// initiates the UX loop
	models.NewStager(
		stack.R,
		stack.Stage,
		"",
		"",
		"",
		&datatypes.DataTypeTreeStageUpdater{},
		&spectypes.SpecTypesTreeStageUpdater{},
		&specobjects.SpecObjectsTreeStageUpdater{},
		&specrelations.SpecRelationsTreeStageUpdater{},
		&specifications.SpecificationsTreeStageUpdater{},
		&exporter.Exporter{},
		&namer.ObjectNamer{})

	// Expose the HTTP and Socket bridges to the Angular frontend
	wasmregistry.SetupWasmHooks(stack.R)

	select {} // Keep the WASM instance running indefinitely
}
