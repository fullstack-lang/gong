//go:build js && wasm

package main

import (
	"log"

	"github.com/fullstack-lang/gong/dsm/reqif/go/datatypes"
	"github.com/fullstack-lang/gong/dsm/reqif/go/exporter"
	"github.com/fullstack-lang/gong/dsm/reqif/go/namer"
	"github.com/fullstack-lang/gong/dsm/reqif/go/specifications"
	"github.com/fullstack-lang/gong/dsm/reqif/go/specobjects"
	"github.com/fullstack-lang/gong/dsm/reqif/go/specrelations"
	"github.com/fullstack-lang/gong/dsm/reqif/go/spectypes"

	gongreqif_level1stack "github.com/fullstack-lang/gong/dsm/reqif/go/level1stack"
	gongreqif_models "github.com/fullstack-lang/gong/dsm/reqif/go/models"

	"github.com/fullstack-lang/gong/lib/wasmregistry"
)

func main() {
	log.SetOutput(&wasmregistry.ConsoleWriter{})
	log.SetPrefix("reqif: ")
	log.SetFlags(log.Lmicroseconds)

	log.Println("Initializing DSM reqif WASM Backend...")

	unmarshallFromCode := ""
	marshallOnCommit := ""
	embeddedDiagrams := true

	// setup model stack with its probe
	stack := gongreqif_level1stack.NewLevel1Stack("reqif", unmarshallFromCode, marshallOnCommit, true, embeddedDiagrams)
	stack.Stage.SetGongMarshallingMode(gongreqif_models.GongMarshallingAppendCommit)
	stack.Stage.SetIsWithGenesisCommit(true)

	// refresh the probe
	stack.Probe.Refresh()

	// initiates the UX loop
	gongreqif_models.NewStager(
		stack.R,
		stack.Stage,
		"", // pathToReqifFile
		"", // pathToRenderingConf
		"", // pathToOutputReqifFile
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
