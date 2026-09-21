package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/fullstack-lang/gong/go/golang"
	"github.com/fullstack-lang/gong/test/test2/go/models"
)

func TestCLIMigrationToStageSet(t *testing.T) {
	inputStageFile := "../cmd/test2/data/stage.go"
	tempDir := t.TempDir()
	migratedStageSetFile := filepath.Join(tempDir, "stageset.go")

	// 1. Run CLI migration
	resPath, err := golang.MigrateStageFile(inputStageFile, migratedStageSetFile, "main", false)
	if err != nil {
		t.Fatalf("golang.MigrateStageFile failed: %v", err)
	}
	if resPath != migratedStageSetFile {
		t.Fatalf("expected output path %s, got %s", migratedStageSetFile, resPath)
	}

	// 2. Read and verify content format
	data, err := os.ReadFile(migratedStageSetFile)
	if err != nil {
		t.Fatalf("failed to read migrated stageset file: %v", err)
	}
	content := string(data)

	if !strings.Contains(content, `func _(stageSet *__stage_0__.StageSet)`) {
		t.Errorf("missing StageSet function signature in:\n%s", content)
	}
	if !strings.Contains(content, `__stage_0__ "github.com/fullstack-lang/gong/test/test2/go/models"`) {
		t.Errorf("missing __stage_0__ import in:\n%s", content)
	}

	// 3. Parse migrated file into a new StageSet
	stageSet := models.NewStageSet("test_migration")
	if err := stageSet.ParseAstFile(migratedStageSetFile, false); err != nil {
		t.Fatalf("failed to parse migrated StageSet file: %v", err)
	}

	// 4. Verify instances and relationships
	as := stageSet.Stage.GetInstancesSorted[*models.A]()
	bs := stageSet.Stage.GetInstancesSorted[*models.B]()

	if len(as) != 1 {
		t.Fatalf("expected 1 A instance, found %d", len(as))
	}
	if len(bs) != 1 {
		t.Fatalf("expected 1 B instance, found %d", len(bs))
	}

	a := as[0]
	b := bs[0]

	if a.Name != "A1" {
		t.Errorf("expected A name 'A1', got %s", a.Name)
	}
	if a.NumberField != 42 {
		t.Errorf("expected A NumberField 42, got %d", a.NumberField)
	}
	if a.Foo != 100 {
		t.Errorf("expected A Foo 100, got %d", a.Foo)
	}
	if a.Zorgh != "Test Single Stage" {
		t.Errorf("expected A Zorgh 'Test Single Stage', got %s", a.Zorgh)
	}
	if b.Name != "B1" {
		t.Errorf("expected B name 'B1', got %s", b.Name)
	}

	// Verify pointer wiring: A.B -> B
	if a.B == nil {
		t.Fatalf("expected A.B to be non-nil")
	}
	if a.B != b {
		t.Errorf("expected A.B to point to b (%s), got %s", b.Name, a.B.Name)
	}
}

func TestInAppMigrationRoundTrip(t *testing.T) {
	inputStageFile := "../cmd/test2/data/stage.go"

	// 1. Unmarshal legacy single-stage file into models.Stage
	stage := models.NewStage("in_app_migration")
	if err := stage.ParseAstFile(inputStageFile, true); err != nil {
		t.Fatalf("stage.ParseAstFile failed: %v", err)
	}

	// 2. Create StageSet from existing Stage
	stageSet := models.NewStageSetFromStage(stage)

	// 3. Marshall to Go code string
	marshalled, err := stageSet.MarshallToString("main")
	if err != nil {
		t.Fatalf("stageSet.MarshallToString failed: %v", err)
	}

	if !strings.Contains(marshalled, `func _(stageSet *models.StageSet)`) {
		t.Errorf("missing StageSet signature in in-app marshalled code:\n%s", marshalled)
	}
	if !strings.Contains(marshalled, `__models__A__00000000_ := (&models.A{Name: `) {
		t.Errorf("missing A declaration in in-app marshalled code:\n%s", marshalled)
	}
	if !strings.Contains(marshalled, `__models__A__00000000_.B = __models__B__00000000_`) {
		t.Errorf("missing A.B pointer setup in in-app marshalled code:\n%s", marshalled)
	}

	// 4. Parse the marshalled string into another clean StageSet and verify
	stageSet2 := models.NewStageSet("verification_stageset")
	if err := stageSet2.ParseAstString(marshalled, false); err != nil {
		t.Fatalf("stageSet2.ParseAstString failed: %v", err)
	}

	as := stageSet2.Stage.GetInstancesSorted[*models.A]()
	if len(as) != 1 {
		t.Fatalf("expected 1 A instance, got %d", len(as))
	}
	if as[0].B == nil || as[0].B.Name != "B1" {
		t.Errorf("expected A.B to point to B1 in round-trip unmarshalled StageSet")
	}
}
