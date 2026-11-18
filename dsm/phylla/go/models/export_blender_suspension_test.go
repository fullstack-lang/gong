package models

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateTubeVaseSuspensionBlenderPySnapshot(t *testing.T) {
	stage := NewStage("test_suspension_snapshot")

	lib := (&Library{
		Name:              "Herbarium",
		NbPixPerCharacter: 8.0,
	}).Stage(stage)

	plant := (&PlantAbstract{
		Name:               "Tube Vase (2,1,104)",
		PlantType:          TubeVase,
		CurrentView:        VIEW_TUBE_VASE_3D,
		N:                  1,
		M:                  2,
		StackHeight:        7,
		RhombusInsideAngle: 162.0,
		RhombusSideLength:  245.0,
		IsSelected:         true,
	}).Stage(stage)
	lib.Plants = append(lib.Plants, plant)

	vase := (&TubeVaseAbstract{
		Name:                           "Tube Vase (2,1,104)-TubeVaseAbstract",
		RibbonVerticalScale:            1.0,
		RelativeVerticalThickness:      0.01,
		RelativeRadialThickness:        0.04,
		RelativeCuttedStackFloorHeight: 0.0,
		RelativeRotatedTorusSeparation: 0.05,
		RotationRatio:                  1.0,
		RadialRepetitions:              1,
		Transparency:                   0.0,
		ColorCombination:               7,
	}).Stage(stage)
	plant.TubeVaseAbstract = vase

	rendered := (&Rendered3DShape{
		Name:    "Plant 1-Rendered3DShape",
		ViewX:   -402.970868,
		ViewY:   145.178596,
		ViewZ:   -165.602479,
		TargetX: 30.760008,
		TargetY: 101.824254,
		TargetZ: 13.792785,
		Fov:     50.0,
	}).Stage(stage)

	diag := (&TubeVase3DDiagram{
		Name:                                      "Plant 1-TubeVase3DDiagram",
		IsChecked:                                 true,
		IsHiddenTorusStackShape:                   true,
		IsHiddenStackOfPartiallyRotatedTorusShape: false,
		Rendered3DShape:                           rendered,
	}).Stage(stage)
	plant.TubeVase3DDiagrams = append(plant.TubeVase3DDiagrams, diag)

	stager := NewStagerForTest(stage)
	stager.enforceSemantic()

	pyScript := GenerateTubeVaseSuspensionBlenderPy(plant)
	if len(pyScript) == 0 {
		t.Fatalf("expected non-empty Blender Python script")
	}

	requiredSnippets := []string{
		"01_Pure_Suspension",
		"02_Suspension_Inside_LEDs",
		"Suspension_Inside_LEDs",
		"BASE_BOTTOM = [",
		"BASE_TOP = [",
		"LAYERS = [",
		"#2b3a4a",
		"#e07a5f",
		"#e9c46a",
		"#457b63",
	}
	for _, snippet := range requiredSnippets {
		if !strings.Contains(pyScript, snippet) {
			t.Errorf("expected generated Blender script to contain %q", snippet)
		}
	}

	outPath := filepath.Join(".", "tube_vase_suspension_blender.py")
	if err := os.WriteFile(outPath, []byte(pyScript), 0644); err != nil {
		t.Fatalf("failed to write %s: %v", outPath, err)
	}
	t.Logf("Generated %s (%d bytes) for snapshot %q", outPath, len(pyScript), plant.Name)
}

func TestGenerateTubeVaseSuspensionBlenderPyFromStage(t *testing.T) {
	stagePath := filepath.Join("..", "cmd", "phylla", "data", "stage.go")
	if _, err := os.Stat(stagePath); os.IsNotExist(err) {
		t.Skipf("stage.go not found at %s", stagePath)
	}

	stage := NewStage("test_blender_suspension_stage")
	if err := stage.ParseAstFile(stagePath, true); err != nil {
		t.Fatalf("failed to parse stage.go: %v", err)
	}

	stage.ComputeReverseMaps()
	stage.ComputeInstancesNb()
	stage.ComputeReferenceAndOrders()

	stager := NewStagerForTest(stage)
	stager.enforceSemantic()

	var targetPlant *PlantAbstract
	for plant := range stage.PlantAbstracts {
		if plant.PlantType == TubeVase {
			targetPlant = plant
			break
		}
	}
	if targetPlant == nil {
		t.Fatalf("expected at least one TubeVase PlantAbstract in stage.go")
	}

	pyScript := GenerateTubeVaseSuspensionBlenderPy(targetPlant)
	if len(pyScript) == 0 {
		t.Fatalf("expected non-empty Blender Python script for %q", targetPlant.Name)
	}
	if !strings.Contains(pyScript, "01_Pure_Suspension") || !strings.Contains(pyScript, "02_Suspension_Inside_LEDs") {
		t.Errorf("expected both pure and LED scenes in generated script")
	}
}
