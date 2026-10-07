//go:build js && wasm

package runtime

import (
	"context"
	"syscall/js"

	"github.com/fullstack-lang/gong/lib/wasmregistry"
)

// RegisterWasmSocket provides a unified socket hook for Gong stacks in WASM mode
func RegisterWasmSocket(
	stackType string,
	stackPath string,
	copyDataToJSON func() ([]byte, error),
	subscribeToCommitNb func(context.Context) <-chan int,
) {
	wasmregistry.Register(stackType, stackPath, func(callback js.Value) {
		pushState := func() {
			b, err := copyDataToJSON()
			if err == nil {
				callback.Invoke(string(b))
			}
		}
		pushState()
		for range subscribeToCommitNb(context.Background()) {
			pushState()
		}
	})
}
