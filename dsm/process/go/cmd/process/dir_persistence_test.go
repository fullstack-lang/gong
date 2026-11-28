//go:build !js

package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/fullstack-lang/gong/dsm/process/go/level1stack"
	process_data "github.com/fullstack-lang/gong/dsm/process/go/cmd/process/data/simple"
	"github.com/fullstack-lang/gong/dsm/process/go/models"
)

func TestDirectoryPersistenceSkipGoMod(t *testing.T) {
	// 1. Verify in-repo data/simple directory compiles as a regular package without its own go.mod
	if _, err := os.Stat(filepath.Join("data", "simple", "go.mod")); !os.IsNotExist(err) {
		t.Fatalf("data/simple should not contain go.mod")
	}
	compiledStage := models.NewStage("compiled_simple")
	process_data.StageProcess(compiledStage)
	origCount := len(compiledStage.GetInstances())
	if origCount == 0 {
		t.Fatalf("expected non-empty staged instances from data/simple.StageProcess")
	}

	// 2. Verify level1stack with skipGoMod=true initializes and commits to a directory without creating go.mod
	outDir := filepath.Join(t.TempDir(), "edited_simple")
	if err := os.MkdirAll(outDir, 0755); err != nil {
		t.Fatalf("MkdirAll failed: %v", err)
	}

	stack := level1stack.NewLevel1StackDelta("process_test", "data/simple", outDir, false, false, true, true)
	if len(stack.Stage.GetInstances()) != origCount {
		t.Fatalf("instance count mismatch: got %d, want %d", len(stack.Stage.GetInstances()), origCount)
	}

	if _, err := os.Stat(filepath.Join(outDir, "go.mod")); !os.IsNotExist(err) {
		t.Fatalf("expected go.mod to NOT be generated on init when skipGoMod=true")
	}
	if _, err := os.Stat(filepath.Join(outDir, "stage_process.go")); err != nil {
		t.Fatalf("expected stage_process.go on init: %v", err)
	}

	// Mutate and commit to trigger BeforeCommit hook
	newTask := (&models.Task{Name: "AddedInDirMode"}).Stage(stack.Stage)
	_ = newTask
	stack.Stage.Commit()

	if _, err := os.Stat(filepath.Join(outDir, "go.mod")); !os.IsNotExist(err) {
		t.Fatalf("expected go.mod to NOT be generated on commit when skipGoMod=true")
	}
	if _, err := os.Stat(filepath.Join(outDir, "go.sum")); !os.IsNotExist(err) {
		t.Fatalf("expected go.sum to NOT be generated on commit when skipGoMod=true")
	}

	// Unmarshall back and verify mutation was persisted
	reloaded := models.NewStage("reloaded")
	if err := reloaded.ParseAstFile(outDir, true); err != nil {
		t.Fatalf("ParseAstFile(outDir) failed: %v", err)
	}
	tasks := reloaded.GetInstancesMapByName[*models.Task]()
	if tasks["AddedInDirMode"] == nil {
		t.Fatalf("expected AddedInDirMode task to be persisted in stage_process.go")
	}
}
