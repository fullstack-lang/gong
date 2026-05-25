//go:build wasm

package main

import (
	"flag"
	"fmt"

	"github.com/fullstack-lang/gong/lib/wasmregistry"
)

func main() {
	flag.Parse()

	// 1. Initialize the exact same Gong app in memory
	// Capture the Gin router in a temporary variable 'r'
	r := setupApp()

	fmt.Println("From the backend. after setupApp")

	// 2. Expose the HTTP and Socket bridges to the Angular frontend
	wasmregistry.SetupWasmHooks(r)
	fmt.Println("From the backend. after SetupWasmHooks")

	// 3. Prevent the WASM module from exiting
	select {}
}
