// n Go, if a file name ends with _$GOARCH.go (like _wasm.go), it is treated as having an implicit //go:build wasm constraint.
package wasmregistry

import (
	"fmt"
	"syscall/js"
)

// The global map you used to have in main()
var Handlers = make(map[string]func(js.Value))

// Register adds a stack's socket handler to the map dynamically
func Register(stackPath string, handler func(js.Value)) {

	fmt.Println("Registering")

	Handlers[stackPath] = handler
}
