package y

import (
	"testing"
)

func TestYStageBasicFunctionalities(t *testing.T) {
	stage := NewStage("test_y")

	// 1. Initially empty
	instancesMap := GetGongstructInstancesMap[Y](stage)
	if len(*instancesMap) != 0 {
		t.Fatalf("expected 0 instances, got %d", len(*instancesMap))
	}

	// 2. Stage instances
	y1 := (&Y{Name: "Y1"}).Stage(stage)
	y2 := (&Y{Name: "Y2"}).Stage(stage)

	stage.Commit()

	// 3. Test GetGongstructInstancesMap
	instancesMap = GetGongstructInstancesMap[Y](stage)
	if len(*instancesMap) != 2 {
		t.Fatalf("expected 2 instances in map, got %d", len(*instancesMap))
	}
	if (*instancesMap)["Y1"] != y1 {
		t.Fatalf("expected Y1 in map")
	}
	if (*instancesMap)["Y2"] != y2 {
		t.Fatalf("expected Y2 in map")
	}

	// 4. Test GetGongstructInstancesSet
	instancesSet := GetGongstructInstancesSet[Y](stage)
	if len(*instancesSet) != 2 {
		t.Fatalf("expected 2 instances in set, got %d", len(*instancesSet))
	}
	if _, ok := (*instancesSet)[y1]; !ok {
		t.Fatalf("expected y1 in set")
	}

	// 5. Test GetGongstrucsSorted
	sorted := GetGongstrucsSorted[*Y](stage)
	if len(sorted) != 2 {
		t.Fatalf("expected 2 sorted instances, got %d", len(sorted))
	}
	if sorted[0].Name != "Y1" || sorted[1].Name != "Y2" {
		t.Fatalf("unexpected sorted order: %s, %s", sorted[0].Name, sorted[1].Name)
	}

	// 6. Test Unstage & Commit
	y1.Unstage(stage)
	stage.Commit()

	instancesMap = GetGongstructInstancesMap[Y](stage)
	if len(*instancesMap) != 1 {
		t.Fatalf("expected 1 instance after unstage, got %d", len(*instancesMap))
	}
	if _, ok := (*instancesMap)["Y1"]; ok {
		t.Fatalf("Y1 should not be in map after unstage")
	}
}
