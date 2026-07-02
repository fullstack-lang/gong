// n Go, if a file name ends with _$GOARCH.go (like _wasm.go), it is treated as having an implicit //go:build wasm constraint.
package wasmregistry

import (
	"fmt"
	"log"
	"math"
	"strings"
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
	// fmt.Println("Registering")

	key := HandleKey{
		StackType: stackType,
		StackPath: stackPath,
	}

	Handlers[key] = handler
}

// LogWasmPush logs the payload size in the same format as the backend server
func LogWasmPush(pkgPathRoot string, stackPath string, index int, jsonSize int) {
	parts := strings.Split(pkgPathRoot, "/")
	component := "unknown"
	if len(parts) > 2 {
		component = parts[len(parts)-2]
	}

	log.Printf(
		"%-12s | %-85s | Idx: %d | Size: %-9s",
		component,
		stackPath,
		index,
		formatBytes(jsonSize),
	)
}

// formatBytes converts a size in bytes to a human-readable string (KB, MB, GB).
func formatBytes(size int) string {
	if size < 1024 {
		return fmt.Sprintf("%d B", size)
	}
	sizeInKB := float64(size) / 1024.0
	if sizeInKB < 1024.0 {
		if math.Mod(sizeInKB, 1.0) == 0 {
			return fmt.Sprintf("%.0f KB", sizeInKB)
		}
		return fmt.Sprintf("%.1f KB", sizeInKB)
	}
	sizeInMB := sizeInKB / 1024.0
	return fmt.Sprintf("%.2f MB", sizeInMB)
}
