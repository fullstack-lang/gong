//go:build js && wasm

package main

import (
	"log"

	"github.com/fullstack-lang/gong/lib/wasmregistry"

	// insertion point for models import{{modelsImportDirective}}
	split_models "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
)

func main() {
	log.SetOutput(&wasmregistry.ConsoleWriter{})
	log.SetPrefix("split: ")
	log.SetFlags(log.Lmicroseconds)

	log.Println("Initializing split WASM Backend...")

	unmarshallFromCode := ""
	marshallOnCommit := ""
	embeddedDiagrams := true

	// setup model stack with its probe
	stack := split_stack.NewStack(nil, "split", unmarshallFromCode, marshallOnCommit, "", embeddedDiagrams, true)
	stack.Stage.SetGongMarshallingMode(split_models.GongMarshallingAppendCommit)
	stack.Stage.SetIsWithGenesisCommit(true) // the genesis commit is the first commit of the stage, it is the one that contains the initial data. It cannot be rollbacked.

	stack.Probe.Refresh()
	stack.Stage.Commit()

	split_models.NewStager(nil, stack.Stage, stack.Probe)

	// Expose the HTTP and Socket bridges to the Angular frontend
	wasmregistry.SetupWasmHooks(nil)

	select {} // Keep the WASM instance running indefinitely
}
