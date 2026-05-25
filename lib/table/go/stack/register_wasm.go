//go:build wasm

// do not modify, generated file

package stack

import (
	"context"
	"encoding/json"
	"fmt"
	"syscall/js"
	
	"github.com/fullstack-lang/gong/lib/wasmregistry"
	"github.com/fullstack-lang/gong/lib/table/go/orm"
)

// This function ONLY exists in WASM builds
func registerWasmSocket(stackPath string, backRepo *orm.BackRepoStruct) {

	fmt.Println("github.com/fullstack-lang/gong/lib/table/go", "registerWasmSocket", stackPath)

	wasmregistry.Register(stackPath, func(callback js.Value) {
		pushState := func() {
			data := new(orm.BackRepoData)
			orm.CopyBackRepoToBackRepoData(backRepo, data)
			b, _ := json.Marshal(data)
			callback.Invoke(string(b))
		}	
		pushState() 
		for range backRepo.SubscribeToCommitNb(context.Background()) {
			pushState() 
		}
	})
}

//