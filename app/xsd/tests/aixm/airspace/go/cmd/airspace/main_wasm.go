//go:build js && wasm

package main

import (
	"embed"
	"log"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/level1stack"
	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models"
	"github.com/fullstack-lang/gong/lib/wasmregistry"
)

//go:embed data/stage.go data/donlon_airspace.xml
var dataFS embed.FS

func main() {
	log.SetOutput(&wasmregistry.ConsoleWriter{})
	log.SetPrefix("airspace: ")
	log.SetFlags(log.Lmicroseconds)

	log.Println("Initializing airspace WASM Backend...")

	models.DataFS = &dataFS

	unmarshallFromCode := ""
	marshallOnCommit := ""
	embeddedDiagrams := true

	// setup
	// - model level1 stack with its probe
	// - unmarshall/marshall go file with stage data
	stack := level1stack.NewLevel1StackDelta("airspace", unmarshallFromCode, marshallOnCommit, true, embeddedDiagrams, true)
	stack.Stage.SetGongMarshallingMode(models.GongMarshallingAppendCommit)
	stack.Stage.SetIsWithGenesisCommit(true) // the genesis commit is the first commit of the stage, it is the one that contains the initial data. It cannot be rollbacked.

	if unmarshallFromCode == "" {
		if stageBytes, err := dataFS.ReadFile("data/stage.go"); err == nil {
			if err := stack.Stage.ParseAstString(string(stageBytes), true); err != nil {
				log.Printf("Failed parsing embedded data/stage.go: %v", err)
			} else {
				stack.Stage.ComputeReverseMaps()
				stack.Stage.ComputeInstancesNb()
				stack.Stage.ComputeReferenceAndOrders()
			}
		}
	}

	// refresh the probe, therefore we can see what has been unmarshalled
	stack.Probe.Refresh()
	stack.Stage.Commit()

	// initiates the UX loop
	models.NewStager(
		stack.R,
		stack.Stage,
		stack.Probe,
	)

	// Expose the HTTP and Socket bridges to the Angular frontend
	wasmregistry.SetupWasmHooks(stack.R)

	select {} // Keep the WASM instance running indefinitely
}
