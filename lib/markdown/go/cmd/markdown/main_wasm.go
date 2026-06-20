//go:build js && wasm

package main

import (
	"log"

	"github.com/fullstack-lang/gong/lib/wasmregistry"

	// insertion point for models import{{modelsImportDirective}}
	markdown_models "github.com/fullstack-lang/gong/lib/markdown/go/models"
	markdown_stack "github.com/fullstack-lang/gong/lib/markdown/go/stack"
)

func main() {
	log.SetOutput(&wasmregistry.ConsoleWriter{})
	log.SetPrefix("markdown: ")
	log.SetFlags(log.Lmicroseconds)

	log.Println("Initializing markdown WASM Backend...")

	unmarshallFromCode := ""
	marshallOnCommit := ""
	embeddedDiagrams := true

	// setup model stack with its probe
	stack := markdown_stack.NewStack(nil, "markdown", unmarshallFromCode, marshallOnCommit, "", embeddedDiagrams, true)
	stack.Stage.SetGongMarshallingMode(markdown_models.GongMarshallingAppendCommit)
	stack.Stage.SetIsWithGenesisCommit(true) // the genesis commit is the first commit of the stage, it is the one that contains the initial data. It cannot be rollbacked.

	stack.Probe.Refresh()
	stack.Stage.Commit()

	markdown_models.NewStager(nil, stack.Stage, stack.Probe)

	// Expose the HTTP and Socket bridges to the Angular frontend
	wasmregistry.SetupWasmHooks(nil)

	select {} // Keep the WASM instance running indefinitely
}
