package main

import (
	"net/http"
	"testing"

	"github.com/fullstack-lang/gong/test/test2/go/models"
	"github.com/fullstack-lang/gong/test/test2/go/models/probe"
	"github.com/fullstack-lang/gong/test/test2/go/models/x"
	"github.com/fullstack-lang/gong/test/test2/go/models/y"

	form "github.com/fullstack-lang/gong/lib/form/go/models"
	table "github.com/fullstack-lang/gong/lib/table/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
	embeddedgo "github.com/fullstack-lang/gong/test/test2/go"
)

func TestStageSetProbe(t *testing.T) {
	mux := http.NewServeMux()
	stageSet := models.NewStageSet("test_stageset_probe")

	// 1. Create instances in all 3 stages
	yInst := (&y.Y{Name: "Y1"}).Stage(stageSet.YStage)
	xInst := (&x.X{Name: "X1", Y: yInst}).Stage(stageSet.XStage)
	bInst := (&models.B{Name: "B1"}).Stage(stageSet.Stage)
	_ = (&models.A{
		Name:        "A1",
		NumberField: 100,
		B:           bInst,
		X:           xInst,
		Foo:         42,
		Bar:         3.14,
		Zorgh:       "Test",
	}).Stage(stageSet.Stage)

	stageSet.Commit()

	// 2. Instantiate StageSetProbe
	p := probe.NewStageSetProbe(mux, embeddedgo.GoModelsDir, embeddedgo.GoDiagramsDir, true, stageSet)
	if p == nil {
		t.Fatal("expected non-nil StageSetProbe")
	}

	// 3. Refresh probe
	p.Refresh()

	// Verify stageSet split stage name
	splitName := stageSet.GetProbeSplitStageName()
	expectedSplitName := "github.com/fullstack-lang/gong/test/test2/go/models:test_stageset_probe:probe of the probe_stageset"
	if splitName != expectedSplitName {
		t.Fatalf("unexpected split name: %s, expected %s", splitName, expectedSplitName)
	}

	// 4. Verify tree structure and clickability of struct node
	treeStage := p.GetTreeStage()
	if len(treeStage.Trees) == 0 {
		t.Fatal("expected at least one tree in treeStage")
	}
	var rootTree *tree.Tree
	for tr := range treeStage.Trees {
		rootTree = tr
		break
	}
	topNode := rootTree.RootNodes[0]
	if topNode.Name != "StageSet" {
		t.Fatalf("expected top node 'StageSet', got %s", topNode.Name)
	}

	// Verify that root package 'models' node is not present at the root,
	// and that 'A' and 'B' struct nodes are direct children of topNode.
	var aNode *tree.Node
	var xNode *tree.Node
	var yNode *tree.Node
	for _, child := range topNode.Children {
		if child.Name == "models" {
			t.Errorf("did not expect 'models' package node under topNode (root package elements should be direct children)")
		}
		if child.ToolTipText == "Display table of all A instances" {
			aNode = child
		}
		if child.Name == "x" {
			xNode = child
		}
		if child.Name == "y" {
			yNode = child
		}
	}
	if aNode == nil {
		t.Fatal("expected to find node for A instances directly under topNode")
	}
	if xNode == nil {
		t.Fatal("expected to find package node 'x' under topNode")
	}
	if yNode == nil {
		t.Fatal("expected to find package node 'y' under topNode")
	}

	// Verify that 'model' package node is inside 'x' package node
	var modelNode *tree.Node
	for _, child := range xNode.Children {
		if child.Name == "model" {
			modelNode = child
			break
		}
	}
	if modelNode == nil {
		t.Fatal("expected to find package node 'model' within package node 'x'")
	}
	if !aNode.IsNodeClickable {
		t.Fatal("expected aNode.IsNodeClickable to be true")
	}
	if aNode.OnClick == nil {
		t.Fatal("expected aNode.OnClick to be non-nil")
	}

	// 5. Click the 'A' node -> table should be created and populated
	aNode.OnClick(aNode)

	tableStage := p.GetTableStage()
	if len(tableStage.Tables) == 0 {
		t.Fatal("expected table to be populated in tableStage after clicking A node")
	}
	var tableInst *table.Table
	for tbl := range tableStage.Tables {
		tableInst = tbl
		break
	}
	if tableInst.Name != "models.A" {
		t.Fatalf("expected table name 'models.A', got %s", tableInst.Name)
	}
	if len(tableInst.Rows) != 1 {
		t.Fatalf("expected 1 row in table 'A', got %d", len(tableInst.Rows))
	}
	if aNode.BackgroundColor != "lightgrey" {
		t.Fatalf("expected aNode.BackgroundColor 'lightgrey', got %s", aNode.BackgroundColor)
	}

	// 6. Verify AddInstance (+) button on aNode
	if len(aNode.Buttons) == 0 {
		t.Fatal("expected aNode to have an addButton")
	}
	addButton := aNode.Buttons[0]
	if addButton.Icon != "add" {
		t.Fatalf("expected addButton icon 'add', got %s", addButton.Icon)
	}
	if addButton.OnClick == nil {
		t.Fatal("expected addButton OnClick to be non-nil")
	}

	// Click addButton -> initializes new instance form
	addButton.OnClick()

	formStage := p.GetFormStage()
	var formGroup *form.FormGroup
	for fg := range formStage.FormGroups {
		formGroup = fg
		break
	}
	if formGroup == nil {
		t.Fatal("expected formGroup in formStage after clicking addButton")
	}
	if formGroup.Label != "New A" {
		t.Fatalf("expected formGroup Label 'New A', got %s", formGroup.Label)
	}
	if formGroup.HasSuppressButton {
		t.Fatal("expected HasSuppressButton to be false for new instance")
	}

	// Set name in form and save
	for _, formDiv := range formGroup.FormDivs {
		if formDiv.Name == "Name" && len(formDiv.FormFields) > 0 && formDiv.FormFields[0].FormFieldString != nil {
			formDiv.FormFields[0].FormFieldString.Value = "A2"
		}
	}
	formStage.Commit()
	formGroup.OnSave.OnSave()

	// Verify new instance A2 exists in stageSet.Stage
	var foundA2 bool
	for a := range stageSet.Stage.As {
		if a.Name == "A2" {
			foundA2 = true
			break
		}
	}
	if !foundA2 {
		t.Fatal("expected new instance A2 to be staged in stageSet.Stage")
	}
}
