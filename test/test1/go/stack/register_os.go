//go:build !wasm

// do not modify, generated file
package stack
import "github.com/fullstack-lang/gong/test/test1/go/orm"

// This is a No-Op function for Linux/Mac/Windows.
// The Go compiler will inline it and completely remove it at compile time!
func registerWasmSocket(stackPath string, backRepo *orm.BackRepoStruct) {
	// Do nothing. Standard WebSockets will be handled by Gin.
}

//