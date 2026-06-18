//go:build js && wasm

package main

import (
	"log"

	gongreqif_level1stack "github.com/fullstack-lang/gong/dsm/reqif/go/level1stack"
	gongreqif_models "github.com/fullstack-lang/gong/dsm/reqif/go/models"
	
	"github.com/fullstack-lang/gong/dsm/reqif/go/exporter"
	"github.com/fullstack-lang/gong/dsm/reqif/go/datatypes"
	"github.com/fullstack-lang/gong/dsm/reqif/go/namer"
	"github.com/fullstack-lang/gong/dsm/reqif/go/specifications"
	"github.com/fullstack-lang/gong/dsm/reqif/go/specobjects"
	"github.com/fullstack-lang/gong/dsm/reqif/go/specrelations"
	"github.com/fullstack-lang/gong/dsm/reqif/go/spectypes"

	"github.com/fullstack-lang/gong/lib/wasmregistry"
)

func main() {
	log.SetOutput(&wasmregistry.ConsoleWriter{})
	log.SetPrefix("gongreqif: ")
	log.SetFlags(log.Lmicroseconds)

	log.Println("Initializing DSM reqif WASM Backend...")

	unmarshallFromCode := ""
	marshallOnCommit := ""
	embeddedDiagrams := true

	// setup model stack with its probe
	stack := gongreqif_level1stack.NewLevel1Stack("reqif", unmarshallFromCode, marshallOnCommit, true, embeddedDiagrams)
	stack.Probe.Refresh()

	// insertion point for call to stager
	gongreqif_models.NewStager(
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
