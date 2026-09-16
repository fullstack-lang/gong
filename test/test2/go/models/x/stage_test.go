package x

import (
	"testing"

	"github.com/fullstack-lang/gong/test/test2/go/models/y"
)

func TestXStageBasicFunctionalities(t *testing.T) {
	stageY := y.NewStage("test_y_for_x")
	y1 := (&y.Y{Name: "Y1"}).Stage(stageY)
	stageY.Commit()

	stageX := NewStage("test_x")

	// 1. Initially empty
	instancesMap := stageX.GetInstancesMapByName[*X]()
	if len(instancesMap) != 0 {
		t.Fatalf("expected 0 instances, got %d", len(instancesMap))
	}

	// 2. Stage instances referencing y
	a1 := (&X{Name: "A1", Y: y1}).Stage(stageX)
	a2 := (&X{Name: "A2"}).Stage(stageX)

	stageX.Commit()

	// 3. Test GetInstancesMapByName
	instancesMap = stageX.GetInstancesMapByName[*X]()
	if len(instancesMap) != 2 {
		t.Fatalf("expected 2 instances in map, got %d", len(instancesMap))
	}
	if instancesMap["A1"] != a1 {
		t.Fatalf("expected A1 in map")
	}
	if instancesMap["A2"] != a2 {
		t.Fatalf("expected A2 in map")
	}
	if instancesMap["A1"].Y != y1 {
		t.Fatalf("expected A1.Y to point to y1")
	}

	// 4. Test GetInstancesSet
	instancesSet := stageX.GetInstancesSet[*X]()
	if len(*instancesSet) != 2 {
		t.Fatalf("expected 2 instances in set, got %d", len(*instancesSet))
	}
	if _, ok := (*instancesSet)[a1]; !ok {
		t.Fatalf("expected a1 in set")
	}
	if _, ok := (*instancesSet)[a2]; !ok {
		t.Fatalf("expected a2 in set")
	}

	// 5. Test GetInstancesSorted
	sorted := stageX.GetInstancesSorted[*X]()
	if len(sorted) != 2 {
		t.Fatalf("expected 2 sorted instances, got %d", len(sorted))
	}
	if sorted[0].Name != "A1" || sorted[1].Name != "A2" {
		t.Fatalf("unexpected sorted order: %s, %s", sorted[0].Name, sorted[1].Name)
	}

	// 6. Test Unstage & Commit
	a1.Unstage(stageX)
	stageX.Commit()

	instancesMap = stageX.GetInstancesMapByName[*X]()
	if len(instancesMap) != 1 {
		t.Fatalf("expected 1 instance after unstage, got %d", len(instancesMap))
	}
}
