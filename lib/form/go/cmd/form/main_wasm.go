//go:build js && wasm

package main

import (
	"log"

	"github.com/fullstack-lang/gong/lib/wasmregistry"
	"github.com/gin-gonic/gin"

	// insertion point for models import{{modelsImportDirective}}
	form_models "github.com/fullstack-lang/gong/lib/form/go/models"
	form_stack "github.com/fullstack-lang/gong/lib/form/go/stack"
)

func main() {
	log.SetOutput(&wasmregistry.ConsoleWriter{})
	log.SetPrefix("form: ")
	log.SetFlags(log.Lmicroseconds)

	log.Println("Initializing form WASM Backend...")

	unmarshallFromCode := ""
	marshallOnCommit := ""
	embeddedDiagrams := true

	r := gin.Default()

	// setup model stack with its probe
	stack := form_stack.NewStack(r, "form", unmarshallFromCode, marshallOnCommit, "", embeddedDiagrams, true)
	stack.Stage.SetGongMarshallingMode(form_models.GongMarshallingAppendCommit)
	stack.Stage.SetIsWithGenesisCommit(true) // the genesis commit is the first commit of the stage, it is the one that contains the initial data. It cannot be rollbacked.

	stack.Probe.Refresh()
	stack.Stage.Commit()

	

	// Expose the HTTP and Socket bridges to the Angular frontend
	wasmregistry.SetupWasmHooks(r)

	select {} // Keep the WASM instance running indefinitely
}
