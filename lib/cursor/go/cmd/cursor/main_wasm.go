//go:build js && wasm

package main

import (
	"log"

	"github.com/fullstack-lang/gong/lib/wasmregistry"

	// insertion point for models import{{modelsImportDirective}}
	cursor_models "github.com/fullstack-lang/gong/lib/cursor/go/models"
	cursor_stack "github.com/fullstack-lang/gong/lib/cursor/go/stack"
)

func main() {
	log.SetOutput(&wasmregistry.ConsoleWriter{})
	log.SetPrefix("cursor: ")
	log.SetFlags(log.Lmicroseconds)

	log.Println("Initializing cursor WASM Backend...")

	unmarshallFromCode := ""
	marshallOnCommit := ""
	embeddedDiagrams := true

	// setup model stack with its probe
	stack := cursor_stack.NewStack(nil, "cursor", unmarshallFromCode, marshallOnCommit, "", embeddedDiagrams, true)
	stack.Stage.SetGongMarshallingMode(cursor_models.GongMarshallingAppendCommit)
	stack.Stage.SetIsWithGenesisCommit(true) // the genesis commit is the first commit of the stage, it is the one that contains the initial data. It cannot be rollbacked.

	stack.Probe.Refresh()
	stack.Stage.Commit()

	cursor_models.NewStager(nil, stack.Stage, stack.Probe)

	// Expose the HTTP and Socket bridges to the Angular frontend
	wasmregistry.SetupWasmHooks(nil)

	select {} // Keep the WASM instance running indefinitely
}
