package models

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func TestGenerateVaseTrapezeTerraceBlenderPyFromStage(t *testing.T) {
	stagePath := filepath.Join("..", "cmd", "phylla", "data", "stage.go")
	if _, err := os.Stat(stagePath); os.IsNotExist(err) {
		t.Skipf("stage.go not found at %s", stagePath)
	}

	stage := NewStage("test_blender_terrace_stage")
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
		if plant.PlantType == VaseTrapeze {
			targetPlant = plant
			break
		}
	}
	if targetPlant == nil {
		t.Fatalf("expected at least one VaseTrapeze PlantAbstract in stage.go")
	}

	pyScript := GenerateVaseTrapezeTerraceBlenderPy(targetPlant)
	if len(pyScript) == 0 {
		t.Fatalf("expected non-empty Blender Python script for %q", targetPlant.Name)
	}

	// Also verify a 4-pot upper-ring stack (StackHeight = 5)
	if targetPlant.StackHeight < 5 {
		targetPlant.StackHeight = 5
		stager.enforceSemantic()
		pyScript = GenerateVaseTrapezeTerraceBlenderPy(targetPlant)
	}

	requiredSnippets := []string{
		"Mediterranean_Terrace_Planter",
		"USE_TERRACOTTA_MATERIAL = True",
		"ADD_EARTH_AND_PLANTS = True",
		"ADD_GARDEN_BACKGROUND = True",
		"UPPER_LAYERS = [",
		"RING_PROFILE = [",
		"HAS_FILL_LEVEL_LINE = True",
		"HAS_DRAINAGE_HOLE = True",
		"HAS_FOOT_RIM = True",
		"FILL_GROOVE_PROFILE = [",
		"FOOT_RIM_PROFILE = [",
		"DRAINAGE_HOLE_PROFILE = [",
		"01_Stacked_TerraCotta_Pots",
		"02_Pot_Earth_And_Mulch",
		"03_Mediterranean_Plants",
		"04_Terrace_Architecture",
		"05_Garden_Background",
		"create_terracotta_material",
		"build_pot_earth_and_plants",
		"build_terrace_and_garden",
	}
	for _, snippet := range requiredSnippets {
		if !strings.Contains(pyScript, snippet) {
			t.Errorf("expected generated Blender script to contain %q", snippet)
		}
	}

	// Verify tree menu button on VaseTrapeze plant
	var nodes []*tree.Node
	stager.treePlant(targetPlant, &nodes, targetPlant.CurrentView)
	if len(nodes) == 0 {
		t.Fatalf("expected treePlant to append a plant node")
	}
	var foundBtn *tree.Button
	for _, btn := range nodes[0].Menu.Buttons {
		if btn.Name == "export Terrace Planter Blender (.py)" {
			foundBtn = btn
			break
		}
	}
	if foundBtn == nil {
		t.Fatalf("expected 'export Terrace Planter Blender (.py)' button in VaseTrapeze menu")
	}
	if foundBtn.Icon != string(buttons.BUTTON_blender) {
		t.Errorf("expected Blender button icon %q, got %q", buttons.BUTTON_blender, foundBtn.Icon)
	}

	tmpOut := filepath.Join("/tmp", "vase_trapeze_terrace_blender.py")
	if err := os.WriteFile(tmpOut, []byte(pyScript), 0644); err != nil {
		t.Fatalf("failed to write %s: %v", tmpOut, err)
	}
	t.Logf("Generated %s (%d bytes) for %q", tmpOut, len(pyScript), targetPlant.Name)
}

func TestGenerateVaseTrapezeTerraceBlenderPyVariants(t *testing.T) {
	stage := NewStage("test_terrace_variants")

	lib := (&Library{
		Name:              "Herbarium",
		NbPixPerCharacter: 8.0,
	}).Stage(stage)

	plant := (&PlantAbstract{
		Name:               "Terrace Variant",
		PlantType:          VaseTrapeze,
		CurrentView:        VIEW_VASE_TRAPEZE_3D,
		N:                  2,
		M:                  3,
		StackHeight:        4,
		RhombusInsideAngle: 130.0,
		RhombusSideLength:  120.0,
		IsSelected:         true,
	}).Stage(stage)
	lib.Plants = append(lib.Plants, plant)

	vase := (&TubeVaseAbstract{
		Name:                       "Terrace Variant-TubeVaseAbstract",
		Z_Ribbon:                   7.5,
		RibbonVerticalScale:        6.0,
		TopPlaneHeight:             0.58,
		BottomPlaneHeight:          0.22,
		BaseRingBottomPlaneHeight:  0.25,
		ProjectionAngle:            -15.0,
		BasePlateHeight:            0.24,
		CarvedOutTopRingsParameter: 0.45,
		HasConstantThickness:       false,
		HasRingBottom:              false,
		HasCurvedWall:              false,
		HasLip:                     false,
		RelativeVerticalThickness:  0.06,
		RotationRatio:              1.0,
		RadialRepetitions:          1,
		ColorCombination:           4,
	}).Stage(stage)
	plant.TubeVaseAbstract = vase

	stager := NewStagerForTest(stage)
	stager.enforceSemantic()

	pyMulti := GenerateVaseTrapezeTerraceBlenderPy(plant)
	if !strings.Contains(pyMulti, `"ring_index": 3`) {
		t.Errorf("expected 3 upper rings (ring_index 1..3) for StackHeight=4")
	}

	// Single-ring fallback when StackHeight == 1
	plant.StackHeight = 1
	stager.enforceSemantic()
	pySingle := GenerateVaseTrapezeTerraceBlenderPy(plant)
	if len(pySingle) == 0 || !strings.Contains(pySingle, `"ring_index": 1`) {
		t.Errorf("expected non-empty script with 1 fallback upper ring when StackHeight=1")
	}
}
