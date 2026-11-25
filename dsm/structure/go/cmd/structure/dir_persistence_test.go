//go:build !js

package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	process_models "github.com/fullstack-lang/gong/dsm/process/go/models"
	structure_models "github.com/fullstack-lang/gong/dsm/structure/go/models"
)

func TestDirectoryModulePersistenceAndCoexistence(t *testing.T) {
	tmpDir := filepath.Join(t.TempDir(), "multi_dsm_data")

	// 1. Load structure/data/stage.go and persist it as a Go module directory
	structStage := structure_models.NewStage("structure")
	if err := structStage.ParseAstFile("data/stage.go", true); err != nil {
		t.Fatalf("failed parsing structure data/stage.go: %v", err)
	}
	origInstancesCount := len(structStage.GetInstances())
	if origInstancesCount == 0 {
		t.Fatalf("expected non-empty structure stage from data/stage.go")
	}

	if err := structStage.MarshallDirectory(tmpDir); err != nil {
		t.Fatalf("MarshallDirectory for structure failed: %v", err)
	}

	// Verify go.mod and stage_structure.go
	goModBytes, err := os.ReadFile(filepath.Join(tmpDir, "go.mod"))
	if err != nil {
		t.Fatalf("missing go.mod: %v", err)
	}
	if !strings.Contains(string(goModBytes), "module multi_dsm_data") {
		t.Errorf("unexpected go.mod:\n%s", string(goModBytes))
	}

	structCodeBytes, err := os.ReadFile(filepath.Join(tmpDir, "stage_structure.go"))
	if err != nil {
		t.Fatalf("missing stage_structure.go: %v", err)
	}
	structCode := string(structCodeBytes)
	if !strings.Contains(structCode, "package data") ||
		!strings.Contains(structCode, "func StageStructure(stage *structure.Stage)") ||
		!strings.Contains(structCode, "__structure__") {
		t.Errorf("unexpected stage_structure.go format:\n%.500s", structCode)
	}

	// 2. Persist process into the SAME directory to verify Multi-DSM file coexistence
	procStage := process_models.NewStage("process")
	procStageFile := filepath.Join("..", "..", "..", "..", "process", "go", "cmd", "process", "data", "stage.go")
	if err := procStage.ParseAstFile(procStageFile, true); err != nil {
		t.Fatalf("failed parsing process data/stage.go: %v", err)
	}
	origProcCount := len(procStage.GetInstances())

	if err := procStage.MarshallDirectory(tmpDir); err != nil {
		t.Fatalf("MarshallDirectory for process failed: %v", err)
	}

	procCodeBefore, err := os.ReadFile(filepath.Join(tmpDir, "stage_process.go"))
	if err != nil {
		t.Fatalf("missing stage_process.go: %v", err)
	}
	if !strings.Contains(string(procCodeBefore), "func StageProcess(stage *process.Stage)") {
		t.Errorf("expected StageProcess signature in stage_process.go")
	}

	// 3. Write a Go test inside tmpDir calling StageStructure and StageProcess directly as compiled Go code
	callerTest := `package data

import (
	"testing"

	process "github.com/fullstack-lang/gong/dsm/process/go/models"
	structure "github.com/fullstack-lang/gong/dsm/structure/go/models"
)

func TestCompiledStageFunctions(t *testing.T) {
	sStage := structure.NewStage("s")
	StageStructure(sStage)
	if len(sStage.GetInstances()) == 0 {
		t.Fatal("expected staged structure instances")
	}

	pStage := process.NewStage("p")
	StageProcess(pStage)
	if len(pStage.GetInstances()) == 0 {
		t.Fatal("expected staged process instances")
	}
}
`
	if err := os.WriteFile(filepath.Join(tmpDir, "compiled_call_test.go"), []byte(callerTest), 0644); err != nil {
		t.Fatalf("failed writing compiled_call_test.go: %v", err)
	}

	testCmd := exec.Command("go", "test", "./...")
	testCmd.Dir = tmpDir
	if out, err := testCmd.CombinedOutput(); err != nil {
		t.Fatalf("go test ./... inside generated persistence module failed: %v\n%s", err, string(out))
	}
	_ = os.Remove(filepath.Join(tmpDir, "compiled_call_test.go"))

	// 4. Verify standalone re-save of structure leaves stage_process.go untouched
	if err := structStage.MarshallDirectory(tmpDir); err != nil {
		t.Fatalf("re-marshalling structure failed: %v", err)
	}
	procCodeAfter, err := os.ReadFile(filepath.Join(tmpDir, "stage_process.go"))
	if err != nil || string(procCodeBefore) != string(procCodeAfter) {
		t.Fatalf("expected stage_process.go to remain untouched after standalone structure save")
	}

	// 5. Verify AST directory unmarshalling round-trip for both stages independently
	restoredStructStage := structure_models.NewStage("restored_structure")
	if err := restoredStructStage.ParseAstFile(tmpDir, true); err != nil {
		t.Fatalf("ParseAstFile(tmpDir) for structure failed: %v", err)
	}
	if len(restoredStructStage.GetInstances()) != origInstancesCount {
		t.Errorf("structure instance count mismatch: got %d, want %d", len(restoredStructStage.GetInstances()), origInstancesCount)
	}

	restoredProcStage := process_models.NewStage("restored_process")
	if err := restoredProcStage.ParseAstFile(tmpDir, true); err != nil {
		t.Fatalf("ParseAstFile(tmpDir) for process failed: %v", err)
	}
	if len(restoredProcStage.GetInstances()) != origProcCount {
		t.Errorf("process instance count mismatch: got %d, want %d", len(restoredProcStage.GetInstances()), origProcCount)
	}
}
