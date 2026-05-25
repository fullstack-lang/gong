// n Go, if a file name ends with _$GOARCH.go (like _wasm.go), it is treated as having an implicit //go:build wasm constraint.
package wasmregistry

import (
	"fmt"
	"syscall/js"
)

// The global map you used to have in main()
type HandleKey struct {
	StackType string // "github.com/fullstack-lang/gong/lib/tone/go"
	StackPath string // "foo"
}

var Handlers = make(map[HandleKey]func(js.Value))

// Register adds a stack's socket handler to the map dynamically
func Register(stackType string, stackPath string, handler func(js.Value)) {

	fmt.Println("Registering")

	key := HandleKey{
		StackType: stackType,
		StackPath: stackPath,
	}

	Handlers[key] = handler
}
