//go:build js && wasm

package main

import (
	"log"

	"github.com/fullstack-lang/gong/lib/wasmregistry"
	"github.com/gin-gonic/gin"

	// insertion point for models import{{modelsImportDirective}}
	slider_models "github.com/fullstack-lang/gong/lib/slider/go/models"
	slider_stack "github.com/fullstack-lang/gong/lib/slider/go/stack"
)

func main() {
	log.SetOutput(&wasmregistry.ConsoleWriter{})
	log.SetPrefix("slider: ")
	log.SetFlags(log.Lmicroseconds)

	log.Println("Initializing slider WASM Backend...")

	unmarshallFromCode := ""
	marshallOnCommit := ""
	embeddedDiagrams := true

	r := gin.Default()

	// setup model stack with its probe
	stack := slider_stack.NewStack(r, "slider", unmarshallFromCode, marshallOnCommit, "", embeddedDiagrams, true)
	stack.Stage.SetGongMarshallingMode(slider_models.GongMarshallingAppendCommit)
	stack.Stage.SetIsWithGenesisCommit(true) // the genesis commit is the first commit of the stage, it is the one that contains the initial data. It cannot be rollbacked.

	stack.Probe.Refresh()
	stack.Stage.Commit()

	NewStager(r, stack.Stage)

	// Expose the HTTP and Socket bridges to the Angular frontend
	wasmregistry.SetupWasmHooks(r)

	select {} // Keep the WASM instance running indefinitely
}
