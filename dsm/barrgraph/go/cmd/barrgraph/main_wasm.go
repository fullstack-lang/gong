//go:build js && wasm

package main

import (
	"embed"
	"go/parser"
	"go/token"
	"log"

	"github.com/fullstack-lang/gong/dsm/barrgraph/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/barrgraph/go/models"
	"github.com/fullstack-lang/gong/lib/wasmregistry"
)

//go:embed "data/cubism and abstract art.go"
var stageGo string

//go:embed data/*
var dataFS embed.FS

func main() {
	log.SetOutput(&wasmregistry.ConsoleWriter{})
	log.SetPrefix("barrgraph: ")
	log.SetFlags(log.Lmicroseconds)

	log.Println("Initializing DSM barrgraph WASM Backend...")

	unmarshallFromCode := ""
	marshallOnCommit := ""
	embeddedDiagrams := true

	// - set embedded data to models
	models.DataFS = &dataFS

	// setup
	// - model level1 stack with its probe
	// - unmarshall/marshall go file with stage data
	stack := level1stack.NewLevel1StackDelta("barrgraph", unmarshallFromCode, marshallOnCommit, true, embeddedDiagrams, true)
	stack.Stage.SetGongMarshallingMode(models.GongMarshallingAppendCommit)
	stack.Stage.SetIsWithGenesisCommit(true) // the genesis commit is the first commit of the stage, it is the one that contains the initial data. It cannot be rollbacked.

	stage := stack.Stage

	if unmarshallFromCode == "" {
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, "", stageGo, parser.ParseComments)
		if err != nil {
			log.Panic("Unable to parse stageGo " + err.Error())
		}
		err = models.ParseAstFileFromAst(stage, file, fset, true)
		// if the application is run with -unmarshallFromCode=xxx.go -marshallOnCommit
		// xxx.go might be absent the first time. However, this shall not be a show stopper.
		if err != nil {
			log.Println("no file to read " + err.Error())
		}

		stage.ComputeReverseMaps()
		stage.ComputeInstancesNb()
		stage.ComputeReferenceAndOrders()
	}

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
