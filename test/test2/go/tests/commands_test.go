package main

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/fullstack-lang/gong/test/test2/go/level1stack"
	"github.com/fullstack-lang/gong/test/test2/go/models"
	"github.com/fullstack-lang/gong/test/test2/go/models/y"
)

func TestSingleStageAndMultiStageLevel1Stack(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	cmdDataDir := filepath.Join(dir, "..", "cmd", "test2", "data")

	singleStageFile := filepath.Join(cmdDataDir, "stage.go")
	multiStageFile := filepath.Join(cmdDataDir, "stageset.go")

	// 1. Verify NewLevel1Stack (Single-stage mode used by edit)
	stackSingle := level1stack.NewLevel1Stack("test2_single", singleStageFile, "", false, false)
	if stackSingle == nil {
		t.Fatal("expected non-nil stackSingle")
	}
	aMap := stackSingle.Stage.GetInstancesMapByName[*models.A]()
	if len(aMap) != 1 {
		t.Fatalf("expected 1 A instance in single-stage, got %d", len(aMap))
	}
	if _, ok := aMap["A1"]; !ok {
		t.Fatalf("expected A1 instance in single-stage")
	}

	// 2. Verify NewLevel1StackStageSet (Multi-stage mode used by edit-stageset)
	stackMulti := level1stack.NewLevel1StackStageSet("test2_multi", multiStageFile, "", false, false)
	if stackMulti == nil {
		t.Fatal("expected non-nil stackMulti")
	}
	if stackMulti.StageSet == nil {
		t.Fatal("expected non-nil stackMulti.StageSet")
	}
	aMultiMap := stackMulti.StageSet.Stage.GetInstancesMapByName[*models.A]()
	if len(aMultiMap) != 10 {
		t.Fatalf("expected 10 A instances in multi-stage, got %d", len(aMultiMap))
	}
	yMultiMap := stackMulti.StageSet.YStage.GetInstancesMapByName[*y.Y]()
	if len(yMultiMap) != 3 {
		t.Fatalf("expected 3 Y instances in multi-stage YStage, got %d", len(yMultiMap))
	}
}
