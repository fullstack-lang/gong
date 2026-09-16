package models

import (
	"testing"

	"github.com/fullstack-lang/gong/test/test2/go/models/x"
	"github.com/fullstack-lang/gong/test/test2/go/models/y"
)

func TestModelsStageBasicFunctionalities(t *testing.T) {
	// 1. Stage in package y
	stageY := y.NewStage("stage_y")
	yInstance := (&y.Y{Name: "Y_Root"}).Stage(stageY)
	stageY.Commit()

	// 2. Stage in package x referencing y
	stageX := x.NewStage("stage_x")
	xInstance := (&x.X{Name: "X_Root", Y: yInstance}).Stage(stageX)
	stageX.Commit()

	// 3. Stage in package models referencing x and local B
	stageModels := NewStage("stage_models")

	bInstance := (&B{Name: "B_Root"}).Stage(stageModels)
	aInstance := (&A{
		Name:        "A_Root",
		NumberField: 42,
		B:           bInstance,
		Bs:          []*B{bInstance},
		X:           xInstance,
		ToBeImported: x.ToBeImported{
			Foo:   123,
			Bar:   45.67,
			Zorgh: "Hello Gong",
		},
	}).Stage(stageModels)

	stageModels.Commit()

	// Test GetInstancesMapByName for models.A
	aMap := stageModels.GetInstancesMapByName[*A]()
	if len(aMap) != 1 {
		t.Fatalf("expected 1 A instance in models stage, got %d", len(aMap))
	}
	if aMap["A_Root"] != aInstance {
		t.Fatalf("expected A_Root in models map")
	}
	if aMap["A_Root"].X != xInstance {
		t.Fatalf("expected A_Root.X to link to xInstance")
	}
	if aMap["A_Root"].X.Y != yInstance {
		t.Fatalf("expected A_Root.X.Y to link to yInstance")
	}
	if aMap["A_Root"].Foo != 123 {
		t.Fatalf("expected Foo to be 123, got %d", aMap["A_Root"].Foo)
	}
	if aMap["A_Root"].Bar != 45.67 {
		t.Fatalf("expected Bar to be 45.67, got %f", aMap["A_Root"].Bar)
	}
	if aMap["A_Root"].Zorgh != "Hello Gong" {
		t.Fatalf("expected Zorgh to be 'Hello Gong', got %s", aMap["A_Root"].Zorgh)
	}

	// Test GetInstancesMapByName for models.B
	bMap := stageModels.GetInstancesMapByName[*B]()
	if len(bMap) != 1 {
		t.Fatalf("expected 1 B instance in models stage, got %d", len(bMap))
	}

	// Test GetInstancesSet
	aSet := stageModels.GetInstancesSet[*A]()
	if len(*aSet) != 1 {
		t.Fatalf("expected 1 A in set, got %d", len(*aSet))
	}

	// Test GetInstancesSorted
	aSorted := stageModels.GetInstancesSorted[*A]()
	if len(aSorted) != 1 || aSorted[0].Name != "A_Root" {
		t.Fatalf("expected sorted A_Root")
	}

	// Verify all 3 stages simultaneously maintain their respective instances
	yMap := stageY.GetInstancesMapByName[*y.Y]()
	if len(yMap) != 1 || yMap["Y_Root"] != yInstance {
		t.Fatalf("expected Y_Root in stageY")
	}

	xMap := stageX.GetInstancesMapByName[*x.X]()
	if len(xMap) != 1 || xMap["X_Root"] != xInstance {
		t.Fatalf("expected X_Root in stageX")
	}

	// Test Unstage & Commit
	aInstance.Unstage(stageModels)
	stageModels.Commit()

	aMap = stageModels.GetInstancesMapByName[*A]()
	if len(aMap) != 0 {
		t.Fatalf("expected 0 A instances after unstage, got %d", len(aMap))
	}
}
