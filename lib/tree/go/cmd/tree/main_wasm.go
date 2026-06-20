//go:build js && wasm

package main

import (
	"log"

	"github.com/fullstack-lang/gong/lib/wasmregistry"

	// insertion point for models import{{modelsImportDirective}}
	tree_models "github.com/fullstack-lang/gong/lib/tree/go/models"
	tree_stack "github.com/fullstack-lang/gong/lib/tree/go/stack"
)

func main() {
	log.SetOutput(&wasmregistry.ConsoleWriter{})
	log.SetPrefix("tree: ")
	log.SetFlags(log.Lmicroseconds)

	log.Println("Initializing tree WASM Backend...")

	unmarshallFromCode := ""
	marshallOnCommit := ""
	embeddedDiagrams := true

	// setup model stack with its probe
	stack := tree_stack.NewStack(nil, "tree", unmarshallFromCode, marshallOnCommit, "", embeddedDiagrams, true)
	stack.Stage.SetGongMarshallingMode(tree_models.GongMarshallingAppendCommit)
	stack.Stage.SetIsWithGenesisCommit(true) // the genesis commit is the first commit of the stage, it is the one that contains the initial data. It cannot be rollbacked.

	stack.Probe.Refresh()
	stack.Stage.Commit()

	tree_models.NewStager(nil, stack.Stage, stack.Probe)

	// Expose the HTTP and Socket bridges to the Angular frontend
	wasmregistry.SetupWasmHooks(nil)

	select {} // Keep the WASM instance running indefinitely
}
