//go:build wasm

package main

import (
	"flag"
	"fmt"
	"log"
	"syscall/js"

	"github.com/fullstack-lang/gong/lib/wasmregistry"

	project_level1stack "github.com/fullstack-lang/gong/dsm/project/go/level1stack"
	project_models "github.com/fullstack-lang/gong/dsm/project/go/models"
)

type consoleWriter struct{}

func (consoleWriter) Write(p []byte) (n int, err error) {
	msg := string(p)
	// Remove the trailing newline added by the log package, as console.log adds its own
	if len(msg) > 0 && msg[len(msg)-1] == '\n' {
		msg = msg[:len(msg)-1]
	}
	js.Global().Get("console").Call("log", msg)
	return len(p), nil
}

func main() {
	flag.Parse()

	// Setup log to write directly to the browser console
	log.SetOutput(consoleWriter{})
	log.SetPrefix("test4-wasm: ")
	log.SetFlags(log.Lmicroseconds)

	// 1. Initialize the exact same Gong app in memory
	stackproject := project_level1stack.NewLevel1StackDelta("project", "", "", true, true, false)

	// initiates the UX loop
	project_models.NewStager(
		stackproject.R,
		stackproject.Stage,
		stackproject.Probe,
	)

	fmt.Println("From the backend. after setupApp")

	// 2. Expose the HTTP and Socket bridges to the Angular frontend
	wasmregistry.SetupWasmHooks(stackproject.R)
	fmt.Println("From the backend. after SetupWasmHooks")

	// 3. Prevent the WASM module from exiting
	select {}
}
