//go:build !(js && wasm)

package runtime

import "context"

// RegisterWasmSocket is a no-op on non-WASM platforms
func RegisterWasmSocket(
	stackType string,
	stackPath string,
	copyDataToJSON func() ([]byte, error),
	subscribeToCommitNb func(context.Context) <-chan int,
) {
}
