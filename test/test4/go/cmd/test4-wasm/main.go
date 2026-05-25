//go:build wasm

package main

import (
	"flag"
	"fmt"

	"github.com/fullstack-lang/gong/lib/wasmregistry"

	process_level1stack "github.com/fullstack-lang/gong/dsm/process/go/level1stack"
	process_models "github.com/fullstack-lang/gong/dsm/process/go/models"
)

func main() {
	flag.Parse()

	// 1. Initialize the exact same Gong app in memory
	stackProcess := process_level1stack.NewLevel1StackDelta("process", "", "", true, true, true)

	// initiates the UX loop
	process_models.NewStager(
		stackProcess.R,
		stackProcess.Stage,
		stackProcess.Probe,
	)

	fmt.Println("From the backend. after setupApp")

	// 2. Expose the HTTP and Socket bridges to the Angular frontend
	wasmregistry.SetupWasmHooks(stackProcess.R)
	fmt.Println("From the backend. after SetupWasmHooks")

	// 3. Prevent the WASM module from exiting
	select {}
}
