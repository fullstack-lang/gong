package main

import (
	"net/http"
	"testing"

	"github.com/fullstack-lang/gong/test/test2/go/models"
	"github.com/fullstack-lang/gong/test/test2/go/models/probe"
	"github.com/fullstack-lang/gong/test/test2/go/models/x"
	"github.com/fullstack-lang/gong/test/test2/go/models/y"

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

	// Find the 'A' struct node under 'models' package
	var aNode *tree.Node
	for _, pkgNode := range topNode.Children {
		if pkgNode.Name == "models" {
			for _, stNode := range pkgNode.Children {
				if stNode.ToolTipText == "Display table of all A instances" {
					aNode = stNode
					break
				}
			}
		}
	}
	if aNode == nil {
		t.Fatal("expected to find node for A instances in tree")
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
	if tableInst.Name != "A" {
		t.Fatalf("expected table name 'A', got %s", tableInst.Name)
	}
	if len(tableInst.Rows) != 1 {
		t.Fatalf("expected 1 row in table 'A', got %d", len(tableInst.Rows))
	}
	if aNode.BackgroundColor != "lightgrey" {
		t.Fatalf("expected aNode.BackgroundColor 'lightgrey', got %s", aNode.BackgroundColor)
	}
}
