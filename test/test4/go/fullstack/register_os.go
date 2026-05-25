//go:build !wasm

// do not modify, generated file
package fullstack
import "github.com/fullstack-lang/gong/test/test4/go/orm"

// This is a No-Op function for Linux/Mac/Windows.
// The Go compiler will inline it and completely remove it at compile time!
func registerWasmSocket(stackPath string, backRepo *orm.BackRepoStruct) {
	// Do nothing. Standard WebSockets will be handled by Gin.
}

//