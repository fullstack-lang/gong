package models_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/fullstack-lang/gong/lib/maplibre/go/models"
)

func TestGenerateReferenceSceneAndPopulateStageGo(t *testing.T) {
	stage := models.NewStage("maplibre")
	models.GenerateReferenceScene(stage)

	if len(*stage.GetInstancesSet[*models.MapView]()) != 1 {
		t.Fatalf("expected 1 MapView, got %d", len(*stage.GetInstancesSet[*models.MapView]()))
	}
	if len(*stage.GetInstancesSet[*models.Layer]()) == 0 {
		t.Fatalf("expected > 0 Layers")
	}
	if len(*stage.GetInstancesSet[*models.Polygon]()) == 0 {
		t.Fatalf("expected > 0 Polygons")
	}
	if len(*stage.GetInstancesSet[*models.Polyline]()) == 0 {
		t.Fatalf("expected > 0 Polylines")
	}
	if len(*stage.GetInstancesSet[*models.Marker]()) == 0 {
		t.Fatalf("expected > 0 Markers")
	}

	stage.ComputeReferenceAndOrders()
	goCode, err := stage.MarshallToString("", "main")
	if err != nil {
		t.Fatalf("MarshallToString failed: %v", err)
	}

	stageGoPath := filepath.Join("..", "cmd", "maplibre", "data", "stage.go")
	if err := os.MkdirAll(filepath.Dir(stageGoPath), 0755); err != nil {
		t.Fatalf("failed creating data dir: %v", err)
	}
	if err := os.WriteFile(stageGoPath, []byte(goCode), 0644); err != nil {
		t.Fatalf("failed writing %s: %v", stageGoPath, err)
	}

	stage2 := models.NewStage("maplibre_roundtrip")
	if err := stage2.ParseAstString(goCode, true); err != nil {
		t.Fatalf("ParseAstString failed: %v", err)
	}
	if len(*stage2.GetInstancesSet[*models.MapView]()) != 1 {
		t.Fatalf("expected 1 MapView after roundtrip, got %d", len(*stage2.GetInstancesSet[*models.MapView]()))
	}
}
