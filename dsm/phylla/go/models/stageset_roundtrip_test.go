package models

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStageSetUnmarshallStageFile(t *testing.T) {
	stagePath := filepath.Join("..", "cmd", "phylla", "data", "stage.go")
	if _, err := os.Stat(stagePath); os.IsNotExist(err) {
		t.Skipf("stage.go not found at %s", stagePath)
	}

	stageSet := NewStageSet("test_stageset")
	err := stageSet.ParseAstFile(stagePath, true)
	if err != nil {
		t.Fatalf("failed to parse stage.go with stageSet: %v", err)
	}

	stageSet.ComputeReverseMaps()
	stageSet.ComputeInstancesNb()
	stageSet.ComputeReferenceAndOrders()

	plantCount := len(stageSet.Stage.PlantAbstracts)
	if plantCount == 0 {
		t.Errorf("expected PlantAbstracts in Stage, got 0")
	}

	clockCount := len(stageSet.ClockStage.ClockAbstracts)
	if clockCount == 0 {
		t.Errorf("expected ClockAbstracts in ClockStage, got 0")
	}

	musicCount := len(stageSet.MusicStage.MusicAbstracts)
	if musicCount == 0 {
		t.Errorf("expected MusicAbstracts in MusicStage, got 0")
	}

	stoolCount := len(stageSet.StoolStage.StoolAbstracts)
	if stoolCount == 0 {
		t.Errorf("expected StoolAbstracts in StoolStage, got 0")
	}

	t.Logf("Successfully unmarshalled StageSet: %d plants, %d clocks, %d musics, %d stools",
		plantCount, clockCount, musicCount, stoolCount)

	// Verify VaseTrapezeBasePlateShape and BasePlateHeight
	basePlateCount := len(stageSet.Stage.VaseTrapezeBasePlateShapes)
	if basePlateCount == 0 {
		t.Errorf("expected VaseTrapezeBasePlateShapes in Stage, got 0")
	}

	foundBasePlateHeight := false
	for vase := range stageSet.Stage.TubeVaseAbstracts {
		if vase.BasePlateHeight > 0 {
			foundBasePlateHeight = true
			break
		}
	}
	if !foundBasePlateHeight {
		t.Errorf("expected at least one TubeVaseAbstract with BasePlateHeight > 0")
	}

	// Ensure diagrams have shapes via stager
	stager := NewStagerForTest(stageSet.Stage)
	stager.SetStageSet(stageSet)
	stager.enforceDiagramShapes()

	for vase := range stageSet.Stage.TubeVaseAbstracts {
		if vase.Name == "Vase Trapeze-TubeVaseAbstract" && vase.CarvedOutTopRingsParameter == 0 {
			vase.CarvedOutTopRingsParameter = 0.2
		}
	}

	stageSet.MarshallFile(stagePath, "main")

	// Re-parse and verify roundtrip
	stageSet2 := NewStageSet("test_stageset_2")
	err = stageSet2.ParseAstFile(stagePath, true)
	if err != nil {
		t.Fatalf("failed to re-parse stage.go: %v", err)
	}
	stageSet2.ComputeReverseMaps()
	stageSet2.ComputeInstancesNb()
	stageSet2.ComputeReferenceAndOrders()

	carvedRingCount := len(stageSet2.Stage.CarvedOutVaseTrapezeRingShapes)
	if carvedRingCount == 0 {
		t.Errorf("expected CarvedOutVaseTrapezeRingShapes in Stage, got 0")
	}

	foundCarvedParam := false
	for vase := range stageSet2.Stage.TubeVaseAbstracts {
		if vase.CarvedOutTopRingsParameter > 0 {
			foundCarvedParam = true
			break
		}
	}
	if !foundCarvedParam {
		t.Errorf("expected at least one TubeVaseAbstract with CarvedOutTopRingsParameter > 0")
	}

	foundBulbousParam := false
	for vase := range stageSet2.Stage.TubeVaseAbstracts {
		if vase.BulbousEndAngle > 0 && vase.BulbousStartTangentMagnitude > 0 {
			foundBulbousParam = true
			break
		}
	}
	if !foundBulbousParam {
		t.Errorf("expected at least one TubeVaseAbstract with BulbousEndAngle > 0 and BulbousStartTangentMagnitude > 0")
	}

	for diagram := range stageSet2.Stage.TubeVase3DDiagrams {
		if diagram.Name == "Vase Trapeze-TubeVase3DDiagram" {
			if diagram.VaseTrapezeBasePlateShape == nil {
				t.Errorf("expected VaseTrapeze-TubeVase3DDiagram to have VaseTrapezeBasePlateShape")
			}
			if diagram.IsHiddenVaseTrapezeBasePlateShape {
				t.Errorf("expected VaseTrapeze-TubeVase3DDiagram.IsHiddenVaseTrapezeBasePlateShape to be false")
			}
			if diagram.CarvedOutVaseTrapezeRingShape == nil {
				t.Errorf("expected VaseTrapeze-TubeVase3DDiagram to have CarvedOutVaseTrapezeRingShape")
			}
			if diagram.CarvedOutTopCurvePlane1Shape == nil {
				t.Errorf("expected VaseTrapeze-TubeVase3DDiagram to have CarvedOutTopCurvePlane1Shape")
			}
			if diagram.CarvedOutBottomCurvePlane1Shape == nil {
				t.Errorf("expected VaseTrapeze-TubeVase3DDiagram to have CarvedOutBottomCurvePlane1Shape")
			}
			if diagram.StackOfCarvedOutVaseTrapezeRingsShape == nil {
				t.Errorf("expected VaseTrapeze-TubeVase3DDiagram to have StackOfCarvedOutVaseTrapezeRingsShape")
			}
			if diagram.StackOfRotatedCarvedOutVaseTrapezeRingsShape == nil {
				t.Errorf("expected VaseTrapeze-TubeVase3DDiagram to have StackOfRotatedCarvedOutVaseTrapezeRingsShape")
			}
		}
	}
}
