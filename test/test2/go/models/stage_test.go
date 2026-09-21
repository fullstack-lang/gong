package models

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	model "github.com/fullstack-lang/gong/test/test2/go/models/x/models"
	"github.com/fullstack-lang/gong/test/test2/go/models/x"
	"github.com/fullstack-lang/gong/test/test2/go/models/y"
)

func TestModelsStageBasicFunctionalities(t *testing.T) {
	// 0. Stage in package models/x/models (colliding package name models, aliased to model)
	stageSub := model.NewStage("stage_sub")
	subModelInstance := (&model.SubModel{Name: "SubModel_Root"}).Stage(stageSub)
	stageSub.Commit()

	// 1. Stage in package y
	stageY := y.NewStage("stage_y")
	yInstance := (&y.Y{Name: "Y_Root", SubModel: subModelInstance}).Stage(stageY)
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
		Foo:         123,
		Bar:         45.67,
		Zorgh:       "Hello Gong",
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
	if aMap["A_Root"].X.Y.SubModel != subModelInstance {
		t.Fatalf("expected A_Root.X.Y.SubModel to link to subModelInstance")
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

	// Verify all 4 stages simultaneously maintain their respective instances
	subMap := stageSub.GetInstancesMapByName[*model.SubModel]()
	if len(subMap) != 1 || subMap["SubModel_Root"] != subModelInstance {
		t.Fatalf("expected SubModel_Root in stageSub")
	}

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

func TestStageSetMarshallAndUnmarshall(t *testing.T) {
	// 0. Stage in package models/x/models
	stageSub := model.NewStage("stage_sub")
	subModelInstance := (&model.SubModel{Name: "SubModel_Root"}).Stage(stageSub)

	// 1. Stage in package y
	stageY := y.NewStage("stage_y")
	yInstance := (&y.Y{Name: "Y_Root", SubModel: subModelInstance}).Stage(stageY)

	// 2. Stage in package x referencing y
	stageX := x.NewStage("stage_x")
	xInstance := (&x.X{Name: "X_Root", Y: yInstance}).Stage(stageX)

	// 3. Stage in package models referencing x and local B
	stageModels := NewStage("stage_models")
	bInstance := (&B{Name: "B_Root"}).Stage(stageModels)
	aInstance := (&A{
		Name:        "A_Root",
		NumberField: 42,
		B:           bInstance,
		Bs:          []*B{bInstance},
		X:           xInstance,
		Foo:         123,
		Bar:         45.67,
		Zorgh:       "Hello Gong",
	}).Stage(stageModels)

	stageSet := &StageSet{
		Stage:      stageModels,
		XStage:     stageX,
		YStage:     stageY,
		ModelStage: stageSub,
	}
	stageSet.Commit()

	// 4. Marshall to string
	marshalledCode, err := stageSet.MarshallToString("main")
	if err != nil {
		t.Fatalf("MarshallToString failed: %v", err)
	}

	t.Logf("Marshalled code:\n%s", marshalledCode)

	// Verify key elements in the marshalled code
	expectedSubstrings := []string{
		"package main",
		// Natural package imports (unaliased where possible)
		"\"github.com/fullstack-lang/gong/test/test2/go/models\"",
		"\"github.com/fullstack-lang/gong/test/test2/go/models/x\"",
		"\"github.com/fullstack-lang/gong/test/test2/go/models/y\"",
		// Aliased import for colliding package name
		"model \"github.com/fullstack-lang/gong/test/test2/go/models/x/models\"",
		// Function header
		"func _(stageSet *models.StageSet)",
		// Instance declarations with readable prefixes: __<alias>__<Struct>__<order>_
		"__model__SubModel__00000000_ := (&model.SubModel{Name: `SubModel_Root`}).Stage(stageSet.ModelStage)",
		"__y__Y__00000000_ := (&y.Y{Name: `Y_Root`}).Stage(stageSet.YStage)",
		"__x__X__00000000_ := (&x.X{Name: `X_Root`}).Stage(stageSet.XStage)",
		"__models__A__00000000_ := (&models.A{Name: `A_Root`}).Stage(stageSet.Stage)",
		"__models__B__00000000_ := (&models.B{Name: `B_Root`}).Stage(stageSet.Stage)",
		// Values
		"__model__SubModel__00000000_.Name = `SubModel_Root`",
		"__y__Y__00000000_.Name = `Y_Root`",
		"__x__X__00000000_.Name = `X_Root`",
		"__models__A__00000000_.Name = `A_Root`",
		"__models__A__00000000_.NumberField = 42",
		"__models__A__00000000_.Foo = 123",
		"__models__A__00000000_.Bar = 45.670000",
		"__models__A__00000000_.Zorgh = `Hello Gong`",
		"__models__B__00000000_.Name = `B_Root`",
		// Pointers
		"__y__Y__00000000_.SubModel = __model__SubModel__00000000_",
		"__x__X__00000000_.Y = __y__Y__00000000_",
		"__models__A__00000000_.B = __models__B__00000000_",
		"__models__A__00000000_.Bs = append(__models__A__00000000_.Bs, __models__B__00000000_)",
		"__models__A__00000000_.X = __x__X__00000000_",
	}

	for _, expected := range expectedSubstrings {
		if !strings.Contains(marshalledCode, expected) {
			t.Errorf("marshalled code missing expected string:\n%s", expected)
		}
	}

	// Verify blank line separation between stages in declarations
	declSubModelToY := "__model__SubModel__00000000_ := (&model.SubModel{Name: `SubModel_Root`}).Stage(stageSet.ModelStage)\n\n\t__y__Y__00000000_ :="
	if !strings.Contains(marshalledCode, declSubModelToY) {
		t.Errorf("expected empty line between ModelStage and YStage declarations, code was:\n%s", marshalledCode)
	}
	declYToX := "__y__Y__00000000_ := (&y.Y{Name: `Y_Root`}).Stage(stageSet.YStage)\n\n\t__x__X__00000000_ :="
	if !strings.Contains(marshalledCode, declYToX) {
		t.Errorf("expected empty line between YStage and XStage declarations")
	}
	declXToModels := "__x__X__00000000_ := (&x.X{Name: `X_Root`}).Stage(stageSet.XStage)\n\n\t__models__A__00000000_ :="
	if !strings.Contains(marshalledCode, declXToModels) {
		t.Errorf("expected empty line between XStage and ModelsStage declarations")
	}

	// 5. Unmarshall into fresh stages
	newStageModels := NewStage("new_models")
	newStageX := x.NewStage("new_x")
	newStageY := y.NewStage("new_y")
	newStageSub := model.NewStage("new_sub")

	newStageSet := &StageSet{
		Stage:      newStageModels,
		XStage:     newStageX,
		YStage:     newStageY,
		ModelStage: newStageSub,
	}

	err = newStageSet.ParseAstString(marshalledCode, true)
	if err != nil {
		t.Fatalf("ParseAstString failed: %v", err)
	}

	// Verify instances in new stages
	newSubMap := newStageSub.GetInstancesMapByName[*model.SubModel]()
	if len(newSubMap) != 1 || newSubMap["SubModel_Root"] == nil {
		t.Fatalf("expected SubModel_Root in newStageSub")
	}

	newYMap := newStageY.GetInstancesMapByName[*y.Y]()
	if len(newYMap) != 1 || newYMap["Y_Root"] == nil {
		t.Fatalf("expected Y_Root in newStageY")
	}

	newXMap := newStageX.GetInstancesMapByName[*x.X]()
	if len(newXMap) != 1 || newXMap["X_Root"] == nil {
		t.Fatalf("expected X_Root in newStageX")
	}

	newBMap := newStageModels.GetInstancesMapByName[*B]()
	if len(newBMap) != 1 || newBMap["B_Root"] == nil {
		t.Fatalf("expected B_Root in newStageModels")
	}

	newAMap := newStageModels.GetInstancesMapByName[*A]()
	if len(newAMap) != 1 || newAMap["A_Root"] == nil {
		t.Fatalf("expected A_Root in newStageModels")
	}

	newA := newAMap["A_Root"]
	newB := newBMap["B_Root"]
	newX := newXMap["X_Root"]
	newY := newYMap["Y_Root"]
	newSub := newSubMap["SubModel_Root"]

	// Verify values
	if newA.NumberField != 42 {
		t.Errorf("expected NumberField 42, got %d", newA.NumberField)
	}
	if newA.Foo != 123 {
		t.Errorf("expected Foo 123, got %d", newA.Foo)
	}
	if newA.Bar != 45.67 {
		t.Errorf("expected Bar 45.67, got %f", newA.Bar)
	}
	if newA.Zorgh != "Hello Gong" {
		t.Errorf("expected Zorgh 'Hello Gong', got %s", newA.Zorgh)
	}

	// Verify intra-stage pointers
	if newA.B != newB {
		t.Errorf("expected A.B to be newB")
	}
	if len(newA.Bs) != 1 || newA.Bs[0] != newB {
		t.Errorf("expected A.Bs to contain newB")
	}

	// Verify cross-stage pointers
	if newA.X != newX {
		t.Errorf("expected A.X to be newX")
	}
	if newX.Y != newY {
		t.Errorf("expected X.Y to be newY")
	}
	if newA.X.Y != newY {
		t.Errorf("expected A.X.Y to link to newY")
	}
	if newY.SubModel != newSub {
		t.Errorf("expected Y.SubModel to link to newSub")
	}
	if newA.X.Y.SubModel != newSub {
		t.Errorf("expected A.X.Y.SubModel to link to newSub")
	}

	// Suppress unused warnings
	_ = aInstance
}

func TestStageSetEmptyStages(t *testing.T) {
	// Test robustness to empty stages (some or all stages having 0 instances)
	stageModels := NewStage("empty_models")
	stageX := x.NewStage("empty_x")
	stageY := y.NewStage("empty_y")
	stageSub := model.NewStage("empty_sub")

	// Only stage an instance in models; x, y, and model are completely empty
	b := (&B{Name: "Lone_B"}).Stage(stageModels)
	_ = b

	stageSet := &StageSet{
		Stage:      stageModels,
		XStage:     stageX,
		YStage:     stageY,
		ModelStage: stageSub,
	}

	code, err := stageSet.MarshallToString("main")
	if err != nil {
		t.Fatalf("MarshallToString on empty stages failed: %v", err)
	}

	// Verify it still includes all package dummy declarations so unused import errors are prevented
	if !strings.Contains(code, "_ *models.Stage") ||
		!strings.Contains(code, "_ *x.Stage") ||
		!strings.Contains(code, "_ *y.Stage") ||
		!strings.Contains(code, "_ *model.Stage") {
		t.Errorf("expected dummy declarations for all stages to prevent unused imports, got:\n%s", code)
	}

	newStageSet := &StageSet{
		Stage:      NewStage("fresh_models"),
		XStage:     x.NewStage("fresh_x"),
		YStage:     y.NewStage("fresh_y"),
		ModelStage: model.NewStage("fresh_sub"),
	}

	err = newStageSet.ParseAstString(code, true)
	if err != nil {
		t.Fatalf("ParseAstString on empty stages failed: %v", err)
	}

	bMap := newStageSet.Stage.GetInstancesMapByName[*B]()
	if len(bMap) != 1 || bMap["Lone_B"] == nil {
		t.Fatalf("expected Lone_B to be parsed")
	}
}

func TestStageSetMarshallFileAndParseAstFile(t *testing.T) {
	tmpFile := filepath.Join(t.TempDir(), "stage_set_data.go")

	stageSub := model.NewStage("stage_sub")
	subInst := (&model.SubModel{Name: "File_Sub"}).Stage(stageSub)

	stageY := y.NewStage("stage_y")
	yInst := (&y.Y{Name: "File_Y", SubModel: subInst}).Stage(stageY)

	stageX := x.NewStage("stage_x")
	xInst := (&x.X{Name: "File_X", Y: yInst}).Stage(stageX)

	stageModels := NewStage("stage_models")
	aInst := (&A{Name: "File_A", X: xInst}).Stage(stageModels)
	_ = aInst

	stageSet := &StageSet{
		Stage:      stageModels,
		XStage:     stageX,
		YStage:     stageY,
		ModelStage: stageSub,
	}

	stageSet.MarshallFile(tmpFile, "main")

	// Read and verify file exists
	if _, err := os.Stat(tmpFile); err != nil {
		t.Fatalf("expected file to exist: %v", err)
	}

	newStageSet := &StageSet{
		Stage:      NewStage("new_models"),
		XStage:     x.NewStage("new_x"),
		YStage:     y.NewStage("new_y"),
		ModelStage: model.NewStage("new_sub"),
	}

	err := newStageSet.ParseAstFile(tmpFile, true)
	if err != nil {
		t.Fatalf("ParseAstFile failed: %v", err)
	}

	aMap := newStageSet.Stage.GetInstancesMapByName[*A]()
	if len(aMap) != 1 || aMap["File_A"] == nil {
		t.Fatalf("expected File_A in newStageSet")
	}
	if aMap["File_A"].X == nil || aMap["File_A"].X.Name != "File_X" {
		t.Fatalf("expected File_A.X to link to File_X")
	}
	if aMap["File_A"].X.Y == nil || aMap["File_A"].X.Y.Name != "File_Y" {
		t.Fatalf("expected File_A.X.Y to link to File_Y")
	}
	if aMap["File_A"].X.Y.SubModel == nil || aMap["File_A"].X.Y.SubModel.Name != "File_Sub" {
		t.Fatalf("expected File_A.X.Y.SubModel to link to File_Sub")
	}
}

