package main

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/fullstack-lang/gong/test/test2/go/models"
	"github.com/fullstack-lang/gong/test/test2/go/models/x"
	"github.com/fullstack-lang/gong/test/test2/go/models/y"
)

func TestMultiPackageStageSetRoundTrip(t *testing.T) {
	// 1. Initialize StageSet with 3 stages
	stageModels := models.NewStage("stage_models")
	stageX := x.NewStage("stage_x")
	stageY := y.NewStage("stage_y")

	stageSet := &models.StageSet{
		Stage:  stageModels,
		XStage: stageX,
		YStage: stageY,
	}

	// 2. Stage multiple Y instances in stageY
	yParis := (&y.Y{Name: "Y_Paris"}).Stage(stageSet.YStage)
	yTokyo := (&y.Y{Name: "Y_Tokyo"}).Stage(stageSet.YStage)
	yNewYork := (&y.Y{Name: "Y_NewYork"}).Stage(stageSet.YStage)
	stageSet.YStage.Commit()

	// 3. Stage multiple X instances in stageX referencing stageY instances
	xAlpha := (&x.X{Name: "X_Alpha", Y: yParis}).Stage(stageSet.XStage)
	xBeta := (&x.X{Name: "X_Beta", Y: yTokyo}).Stage(stageSet.XStage)
	xGamma := (&x.X{Name: "X_Gamma", Y: yParis}).Stage(stageSet.XStage) // shares yParis
	xDelta := (&x.X{Name: "X_Delta", Y: nil}).Stage(stageSet.XStage)    // nil pointer
	stageSet.XStage.Commit()

	// 4. Stage multiple B instances in stageModels
	bComp1 := (&models.B{Name: "B_Component1"}).Stage(stageSet.Stage)
	bComp2 := (&models.B{Name: "B_Component2"}).Stage(stageSet.Stage)
	bComp3 := (&models.B{Name: "B_Component3"}).Stage(stageSet.Stage)

	// 5. Stage multiple A instances in stageModels with intra-stage and cross-stage pointers
	aSys1 := (&models.A{
		Name:        "A_System1",
		NumberField: 100,
		B:           bComp1,
		Bs:          []*models.B{bComp1, bComp2},
		X:           xAlpha,
		ToBeImported: x.ToBeImported{
			Foo:   10,
			Bar:   3.14,
			Zorgh: "First A instance",
		},
	}).Stage(stageSet.Stage)

	aSys2 := (&models.A{
		Name:        "A_System2",
		NumberField: 200,
		B:           bComp2,
		Bs:          []*models.B{bComp2, bComp3},
		X:           xBeta,
		ToBeImported: x.ToBeImported{
			Foo:   20,
			Bar:   6.28,
			Zorgh: "Second A instance",
		},
	}).Stage(stageSet.Stage)

	aSys3 := (&models.A{
		Name:        "A_System3",
		NumberField: 300,
		B:           nil,
		Bs:          []*models.B{bComp1, bComp3},
		X:           xGamma,
		ToBeImported: x.ToBeImported{
			Foo:   30,
			Bar:   9.42,
			Zorgh: "Third A instance",
		},
	}).Stage(stageSet.Stage)

	aSys4 := (&models.A{
		Name:        "A_System4_NoX",
		NumberField: 400,
		B:           bComp3,
		Bs:          []*models.B{},
		X:           nil,
		ToBeImported: x.ToBeImported{
			Foo:   40,
			Bar:   0.0,
			Zorgh: "Fourth A instance without X",
		},
	}).Stage(stageSet.Stage)

	stageSet.Stage.Commit()

	// Suppress unused variable warnings if any
	_ = aSys1
	_ = aSys2
	_ = aSys3
	_ = aSys4
	_ = yNewYork
	_ = xDelta

	// 6. Marshall to result.go in the current directory
	_, currentFile, _, _ := runtime.Caller(0)
	testDir := filepath.Dir(currentFile)
	resultFilePath := filepath.Join(testDir, "result.go")

	stageSet.MarshallFile(resultFilePath, "main")

	// 7. Unmarshall into a brand new StageSet
	newStageModels := models.NewStage("stage2_models")
	newStageX := x.NewStage("stage2_x")
	newStageY := y.NewStage("stage2_y")

	newStageSet := &models.StageSet{
		Stage:  newStageModels,
		XStage: newStageX,
		YStage: newStageY,
	}

	err := newStageSet.ParseAstFile(resultFilePath, false)
	if err != nil {
		t.Fatalf("ParseAstFile failed: %s", err)
	}

	// 8. Verifications on the unmarshalled stages

	// Verify Y instances
	if len(newStageSet.YStage.Ys) != 3 {
		t.Fatalf("expected 3 Y instances, got %d", len(newStageSet.YStage.Ys))
	}
	yMap := make(map[string]*y.Y)
	for yInst := range newStageSet.YStage.Ys {
		yMap[yInst.Name] = yInst
	}
	for _, name := range []string{"Y_Paris", "Y_Tokyo", "Y_NewYork"} {
		if _, ok := yMap[name]; !ok {
			t.Errorf("missing Y instance: %s", name)
		}
	}

	// Verify X instances
	if len(newStageSet.XStage.Xs) != 4 {
		t.Fatalf("expected 4 X instances, got %d", len(newStageSet.XStage.Xs))
	}
	xMap := make(map[string]*x.X)
	for xInst := range newStageSet.XStage.Xs {
		xMap[xInst.Name] = xInst
	}
	for _, name := range []string{"X_Alpha", "X_Beta", "X_Gamma", "X_Delta"} {
		if _, ok := xMap[name]; !ok {
			t.Errorf("missing X instance: %s", name)
		}
	}

	// Verify cross-stage pointers X -> Y
	if xMap["X_Alpha"].Y != yMap["Y_Paris"] {
		t.Errorf("X_Alpha.Y expected Y_Paris, got %v", xMap["X_Alpha"].Y)
	}
	if xMap["X_Beta"].Y != yMap["Y_Tokyo"] {
		t.Errorf("X_Beta.Y expected Y_Tokyo, got %v", xMap["X_Beta"].Y)
	}
	if xMap["X_Gamma"].Y != yMap["Y_Paris"] {
		t.Errorf("X_Gamma.Y expected Y_Paris, got %v", xMap["X_Gamma"].Y)
	}
	if xMap["X_Delta"].Y != nil {
		t.Errorf("X_Delta.Y expected nil, got %v", xMap["X_Delta"].Y)
	}

	// Verify B instances
	if len(newStageSet.Stage.Bs) != 3 {
		t.Fatalf("expected 3 B instances, got %d", len(newStageSet.Stage.Bs))
	}
	bMap := make(map[string]*models.B)
	for bInst := range newStageSet.Stage.Bs {
		bMap[bInst.Name] = bInst
	}
	for _, name := range []string{"B_Component1", "B_Component2", "B_Component3"} {
		if _, ok := bMap[name]; !ok {
			t.Errorf("missing B instance: %s", name)
		}
	}

	// Verify A instances
	if len(newStageSet.Stage.As) != 4 {
		t.Fatalf("expected 4 A instances, got %d", len(newStageSet.Stage.As))
	}
	aMap := make(map[string]*models.A)
	for aInst := range newStageSet.Stage.As {
		aMap[aInst.Name] = aInst
	}

	// Check A_System1
	a1 := aMap["A_System1"]
	if a1 == nil {
		t.Fatal("missing A_System1")
	}
	if a1.NumberField != 100 || a1.Foo != 10 || a1.Bar != 3.14 || a1.Zorgh != "First A instance" {
		t.Errorf("A_System1 scalar fields mismatch: %+v", a1)
	}
	if a1.B != bMap["B_Component1"] {
		t.Errorf("A_System1.B expected B_Component1, got %v", a1.B)
	}
	if len(a1.Bs) != 2 || a1.Bs[0] != bMap["B_Component1"] || a1.Bs[1] != bMap["B_Component2"] {
		t.Errorf("A_System1.Bs mismatch: %+v", a1.Bs)
	}
	if a1.X != xMap["X_Alpha"] {
		t.Errorf("A_System1.X expected X_Alpha, got %v", a1.X)
	}
	// Chained cross-stage check: A -> X -> Y
	if a1.X.Y != yMap["Y_Paris"] {
		t.Errorf("A_System1.X.Y expected Y_Paris, got %v", a1.X.Y)
	}

	// Check A_System2
	a2 := aMap["A_System2"]
	if a2 == nil {
		t.Fatal("missing A_System2")
	}
	if a2.NumberField != 200 || a2.Foo != 20 || a2.Bar != 6.28 || a2.Zorgh != "Second A instance" {
		t.Errorf("A_System2 scalar fields mismatch: %+v", a2)
	}
	if a2.B != bMap["B_Component2"] {
		t.Errorf("A_System2.B expected B_Component2, got %v", a2.B)
	}
	if len(a2.Bs) != 2 || a2.Bs[0] != bMap["B_Component2"] || a2.Bs[1] != bMap["B_Component3"] {
		t.Errorf("A_System2.Bs mismatch: %+v", a2.Bs)
	}
	if a2.X != xMap["X_Beta"] {
		t.Errorf("A_System2.X expected X_Beta, got %v", a2.X)
	}
	if a2.X.Y != yMap["Y_Tokyo"] {
		t.Errorf("A_System2.X.Y expected Y_Tokyo, got %v", a2.X.Y)
	}

	// Check A_System3
	a3 := aMap["A_System3"]
	if a3 == nil {
		t.Fatal("missing A_System3")
	}
	if a3.B != nil {
		t.Errorf("A_System3.B expected nil, got %v", a3.B)
	}
	if len(a3.Bs) != 2 || a3.Bs[0] != bMap["B_Component1"] || a3.Bs[1] != bMap["B_Component3"] {
		t.Errorf("A_System3.Bs mismatch: %+v", a3.Bs)
	}
	if a3.X != xMap["X_Gamma"] {
		t.Errorf("A_System3.X expected X_Gamma, got %v", a3.X)
	}
	if a3.X.Y != yMap["Y_Paris"] {
		t.Errorf("A_System3.X.Y expected Y_Paris, got %v", a3.X.Y)
	}

	// Check A_System4_NoX
	a4 := aMap["A_System4_NoX"]
	if a4 == nil {
		t.Fatal("missing A_System4_NoX")
	}
	if a4.B != bMap["B_Component3"] {
		t.Errorf("A_System4_NoX.B expected B_Component3, got %v", a4.B)
	}
	if len(a4.Bs) != 0 {
		t.Errorf("A_System4_NoX.Bs expected empty, got %+v", a4.Bs)
	}
	if a4.X != nil {
		t.Errorf("A_System4_NoX.X expected nil, got %v", a4.X)
	}
}
