//go:build wasm

// do not modify, generated file

package fullstack

import (
	"context"
	"encoding/json"

	gong_runtime "github.com/fullstack-lang/gong/pkg/runtime"
	"github.com/fullstack-lang/gong/lib/svg/go/orm"
)

// This function ONLY exists in WASM builds
func registerWasmSocket(stackPath string, backRepo *orm.BackRepoStruct) {
	gong_runtime.RegisterWasmSocket(
		"github.com/fullstack-lang/gong/lib/svg/go",
		stackPath,
		func() ([]byte, error) {
			data := new(orm.BackRepoData)
			orm.CopyBackRepoToBackRepoData(backRepo, data)
			return json.Marshal(data)
		},
		func(ctx context.Context) <-chan int {
			return backRepo.SubscribeToCommitNb(ctx)
		},
	)
}

//