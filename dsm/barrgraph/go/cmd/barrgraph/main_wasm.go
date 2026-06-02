//go:build js && wasm

package main

import (
	"fmt"
	"log"

	"github.com/fullstack-lang/gong/dsm/barrgraph/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/barrgraph/go/models"
	"github.com/fullstack-lang/gong/lib/wasmregistry"
)

func main() {
	fmt.Println("Initializing DSM barrgraph WASM Backend...")

	log.SetPrefix("barrgraph: ")
	log.SetFlags(log.Lmicroseconds)

	unmarshallFromCode := ""
	marshallOnCommit := ""
	embeddedDiagrams := true

	// setup
	// - model level1 stack with its probe
	// - unmarshall/marshall go file with stage data
	stack := level1stack.NewLevel1StackDelta("barrgraph", unmarshallFromCode, marshallOnCommit, true, embeddedDiagrams, true)
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
