//go:build js && wasm

package main

import (
	"log"

	"github.com/fullstack-lang/gong/lib/wasmregistry"
	"github.com/fullstack-lang/gong/lib/threejs/go/models"
	threejs_stack "github.com/fullstack-lang/gong/lib/threejs/go/stack"
	threejs_static "github.com/fullstack-lang/gong/lib/threejs/go/static"
)

func main() {
	log.SetOutput(&wasmregistry.ConsoleWriter{})
	log.SetPrefix("threejs: ")
	log.SetFlags(log.Lmicroseconds)

	log.Println("Initializing threejs WASM Backend...")

	unmarshallFromCode := ""
	marshallOnCommit := ""
	embeddedDiagrams := true

	r := threejs_static.ServeStaticFiles(false)

	// setup
	// - model stack with its probe
	// - unmarshall/marshall go file with stage data
	stack := threejs_stack.NewStack(r, "threejs", unmarshallFromCode, marshallOnCommit, "", embeddedDiagrams, true)
	stack.Stage.SetGongMarshallingMode(models.GongMarshallingAppendCommit)
	stack.Stage.SetIsWithGenesisCommit(true) // the genesis commit is the first commit of the stage, it is the one that contains the initial data. It cannot be rollbacked.

	// refresh the probe, therefore we can see what has been unmarshalled
	stack.Probe.Refresh()

	// initiates the UX loop
	models.NewStager(
		r,
		stack.Stage,
		stack.Probe,
	)

	// Expose the HTTP and Socket bridges to the Angular frontend
	wasmregistry.SetupWasmHooks(r)

	select {} // Keep the WASM instance running indefinitely
}
