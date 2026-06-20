//go:build js && wasm

package main

import (
	"log"

	"github.com/fullstack-lang/gong/lib/wasmregistry"
	"github.com/fullstack-lang/gong/dsm/capture/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/capture/go/models"
)

func main() {
	log.SetOutput(&wasmregistry.ConsoleWriter{})
	log.SetPrefix("capture: ")
	log.SetFlags(log.Lmicroseconds)

	log.Println("Initializing capture WASM Backend...")

	unmarshallFromCode := ""
	marshallOnCommit := ""
	embeddedDiagrams := true

	// setup
	// - model level1 stack with its probe
	// - unmarshall/marshall go file with stage data
	stack := level1stack.NewLevel1StackDelta("capture", unmarshallFromCode, marshallOnCommit, true, embeddedDiagrams, true)
	stack.Stage.SetGongMarshallingMode(models.GongMarshallingAppendCommit)
	stack.Stage.SetIsWithGenesisCommit(true) // the genesis commit is the first commit of the stage, it is the one that contains the initial data. It cannot be rollbacked.

	// refresh the probe, therefore we can see what has been unmarshalled
	stack.Probe.Refresh()

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
