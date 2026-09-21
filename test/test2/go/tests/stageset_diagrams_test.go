package main

import (
	"net/http"
	"strings"
	"testing"

	doc_models "github.com/fullstack-lang/gong/lib/doc/go/models"
	"github.com/fullstack-lang/gong/lib/doc/go/prepare"
	split "github.com/fullstack-lang/gong/lib/split/go/models"
	svg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
	table_models "github.com/fullstack-lang/gong/lib/table/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
	embeddedgo "github.com/fullstack-lang/gong/test/test2/go"
	"github.com/fullstack-lang/gong/test/test2/go/models"
	"github.com/fullstack-lang/gong/test/test2/go/models/probe"
	"github.com/fullstack-lang/gong/test/test2/go/models/x"
)

func TestStageSetDiagramsParseAndMarshall(t *testing.T) {
	mux := http.NewServeMux()
	docStackName := "test_stageset_diagrams"

	metaPackageImports := []*doc_models.MetaPackageImport{
		{Alias: "ref_models", Path: `"github.com/fullstack-lang/gong/test/test2/go/models"`},
		{Alias: "ref_x", Path: `"github.com/fullstack-lang/gong/test/test2/go/models/x"`},
		{Alias: "ref_y", Path: `"github.com/fullstack-lang/gong/test/test2/go/models/y"`},
		{Alias: "ref_model", Path: `"github.com/fullstack-lang/gong/test/test2/go/models/x/models"`},
	}

	receivingSplitArea := &split.AsSplitArea{
		Name: "Bottom",
		Size: 50,
	}

	// 1. Prepare StageSet with embedded diagrams
	stager := prepare.PrepareStageSet(
		mux,
		true, // embeddedDiagrams
		docStackName,
		metaPackageImports,
		embeddedgo.GoModelsDir,
		embeddedgo.GoDiagramsDir,
		receivingSplitArea,
		nil,
	)

	if stager == nil {
		t.Fatal("expected non-nil stager")
	}

	stage := stager.GetStage()
	if stage == nil {
		t.Fatal("expected non-nil doc stage")
	}

	// 2. Verify parsed diagrams
	classdiagrams := stage.GetInstancesSorted[*doc_models.Classdiagram]()
	if len(classdiagrams) == 0 {
		t.Fatal("expected at least one classdiagram parsed from diagrams_set.go")
	}

	var stageSetDiagram *doc_models.Classdiagram
	for _, cd := range classdiagrams {
		if cd.Name == "StageSet_Diagram" {
			stageSetDiagram = cd
			break
		}
	}
	if stageSetDiagram == nil {
		t.Fatal("expected to find Classdiagram named 'StageSet_Diagram'")
	}

	// 3. Verify GongStructShapes across multiple packages
	if len(stageSetDiagram.GongStructShapes) < 3 {
		t.Fatalf("expected at least 3 GongStructShapes, got %d", len(stageSetDiagram.GongStructShapes))
	}

	shapeNames := make(map[string]*doc_models.GongStructShape)
	for _, shape := range stageSetDiagram.GongStructShapes {
		pkg, name := doc_models.IdentifierMetaToPackageAndGongStructName(shape.IdentifierMeta)
		shapeNames[pkg+"."+name] = shape
	}

	for _, expected := range []string{"models.A", "x.X", "y.Y", "model.SubModel"} {
		if _, ok := shapeNames[expected]; !ok {
			t.Errorf("missing expected GongStructShape %s", expected)
		}
	}

	// 4. Verify Cross-package LinkShapes
	aShape := shapeNames["models.A"]
	if aShape == nil {
		t.Fatalf("expected GongStructShape on models.A, got nil")
	}
	var xLink *doc_models.LinkShape
	for _, l := range aShape.LinkShapes {
		if l.Name == "X" {
			xLink = l
			break
		}
	}
	if xLink == nil {
		t.Fatalf("expected link 'X' on models.A, but none found among %d links", len(aShape.LinkShapes))
	}
	expectedXLinkTarget := "ref_x.X{}"
	if xLink.FieldTypeIdentifierMeta != expectedXLinkTarget {
		t.Errorf("expected FieldTypeIdentifierMeta %s, got %v", expectedXLinkTarget, xLink.FieldTypeIdentifierMeta)
	}

	xShape := shapeNames["x.X"]
	if xShape == nil {
		t.Fatalf("expected GongStructShape on x.X, got nil")
	}
	var yLink *doc_models.LinkShape
	for _, l := range xShape.LinkShapes {
		if l.Name == "Y" {
			yLink = l
			break
		}
	}
	if yLink == nil {
		t.Fatalf("expected link 'Y' on x.X, but none found among %d links", len(xShape.LinkShapes))
	}
	expectedYLinkTarget := "ref_y.Y{}"
	if yLink.FieldTypeIdentifierMeta != expectedYLinkTarget {
		t.Errorf("expected FieldTypeIdentifierMeta %s, got %v", expectedYLinkTarget, yLink.FieldTypeIdentifierMeta)
	}

	// 5. Test Marshalling with multi-package imports
	marshalled, err := stage.MarshallToString("github.com/fullstack-lang/gong/lib/doc/go/models", "diagrams")
	if err != nil {
		t.Fatalf("MarshallToString failed: %v", err)
	}

	// Verify multi-package import declarations
	for _, expectedImport := range []string{
		`ref_models "github.com/fullstack-lang/gong/test/test2/go/models"`,
		`ref_x "github.com/fullstack-lang/gong/test/test2/go/models/x"`,
		`ref_y "github.com/fullstack-lang/gong/test/test2/go/models/y"`,
		`ref_model "github.com/fullstack-lang/gong/test/test2/go/models/x/models"`,
	} {
		if !strings.Contains(marshalled, expectedImport) {
			t.Errorf("marshalled code missing expected import: %s\nFull code:\n%s", expectedImport, marshalled)
		}
	}

	// Verify dummy declarations to avoid unused import errors
	for _, expectedDummy := range []string{
		"var _ ref_models.Stage",
		"var _ ref_x.Stage",
		"var _ ref_y.Stage",
		"var _ ref_model.Stage",
	} {
		if !strings.Contains(marshalled, expectedDummy) {
			t.Errorf("marshalled code missing expected dummy decl: %s", expectedDummy)
		}
	}

	for _, expectedRef := range []string{
		"ref_models.A{}.X",
		"ref_x.X{}",
		"ref_x.X{}.Y",
		"ref_y.Y{}",
		"ref_model.SubModel{}",
	} {
		if !strings.Contains(marshalled, expectedRef) {
			t.Errorf("marshalled code missing expected reference: %s", expectedRef)
		}
	}

	// 6. Verify doc tree sub-package hierarchy
	docTreeStage := stager.GetTreeStage()
	if len(docTreeStage.Trees) == 0 {
		t.Fatal("expected at least one tree in docTreeStage")
	}
	var docRootTree *tree.Tree
	for tr := range docTreeStage.Trees {
		docRootTree = tr
		break
	}
	if len(docRootTree.RootNodes) == 0 {
		t.Fatal("expected at least one RootNode in docRootTree")
	}

	var diagramNode *tree.Node
	for _, node := range docRootTree.RootNodes {
		if node.Name == "StageSet_Diagram" {
			diagramNode = node
			break
		}
	}
	if diagramNode == nil {
		t.Fatal("expected to find node for 'StageSet_Diagram'")
	}

	// Verify that under 'StageSet_Diagram', root package Gongstructs is directly attached (no "models" node)
	var rootGongstructsNode *tree.Node
	for _, child := range diagramNode.Children {
		if strings.HasPrefix(child.Name, "Gongstructs") {
			rootGongstructsNode = child
			break
		}
	}
	if rootGongstructsNode == nil {
		t.Fatalf("expected direct 'Gongstructs' category node under StageSet_Diagram for root package, got children: %v", diagramNode.Children)
	}

	// Verify that sub-packages ("x", "y") have their own package nodes under StageSet_Diagram
	pkgNodesByName := make(map[string]*tree.Node)
	for _, child := range diagramNode.Children {
		pkgNodesByName[child.Name] = child
	}

	if _, ok := pkgNodesByName["models"]; ok {
		t.Errorf("did not expect 'models' package node under StageSet_Diagram (root package elements should be direct children)")
	}

	for _, expectedPkg := range []string{"x", "y"} {
		pkgNode, ok := pkgNodesByName[expectedPkg]
		if !ok {
			t.Fatalf("expected sub-package node %q under StageSet_Diagram, got children: %v", expectedPkg, diagramNode.Children)
		}
		// Under each package node, verify Gongstructs node exists
		var gongstructsNode *tree.Node
		for _, child := range pkgNode.Children {
			if strings.HasPrefix(child.Name, "Gongstructs") {
				gongstructsNode = child
				break
			}
		}
		if gongstructsNode == nil {
			t.Fatalf("expected 'Gongstructs' category node under sub-package %q", expectedPkg)
		}
	}

	// Verify that 'model' package node is inside 'x' (respecting hierarchy)
	xNode := pkgNodesByName["x"]
	if xNode == nil {
		t.Fatal("expected 'x' node under StageSet_Diagram")
	}
	var modelNode *tree.Node
	for _, child := range xNode.Children {
		if child.Name == "model" {
			modelNode = child
			break
		}
	}
	if modelNode == nil {
		t.Fatalf("expected 'model' node within 'x' node, got x children: %v", xNode.Children)
	}
	var modelGongstructsNode *tree.Node
	for _, child := range modelNode.Children {
		if strings.HasPrefix(child.Name, "Gongstructs") {
			modelGongstructsNode = child
			break
		}
	}
	if modelGongstructsNode == nil {
		t.Fatalf("expected 'Gongstructs' category node under sub-package 'model'")
	}
}

func TestStageSetProbeEmbedsDiagramArea(t *testing.T) {
	mux := http.NewServeMux()
	stageSet := models.NewStageSet("test_stageset_diagram_probe")

	p := probe.NewStageSetProbe(mux, embeddedgo.GoModelsDir, embeddedgo.GoDiagramsDir, true, stageSet)
	if p == nil {
		t.Fatal("expected non-nil StageSetProbe")
	}

	// 1. Verify docStager is initialized
	docStager := p.GetDocStager()
	if docStager == nil {
		t.Fatal("expected non-nil DocStager in StageSetProbe")
	}

	// 2. Verify split stage contains the diagram editor
	splitStage := p.GetSplitStage()
	if splitStage == nil {
		t.Fatal("expected non-nil SplitStage in StageSetProbe")
	}

	views := splitStage.GetInstancesSorted[*split.View]()
	if len(views) == 0 {
		t.Fatal("expected at least one View in SplitStage")
	}

	mainView := views[0]
	if len(mainView.RootAsSplitAreas) != 2 {
		t.Fatalf("expected 2 RootAsSplitAreas (Top data editor + Bottom diagram editor), got %d", len(mainView.RootAsSplitAreas))
	}

	topArea := mainView.RootAsSplitAreas[0]
	bottomArea := mainView.RootAsSplitAreas[1]

	if topArea.Name != "Top" || topArea.Size != 50 {
		t.Errorf("expected Top area with Size 50, got name=%s size=%f", topArea.Name, topArea.Size)
	}

	if bottomArea.Name != "Bottom" || bottomArea.Size != 50 {
		t.Errorf("expected Bottom diagram area with Size 50, got name=%s size=%f", bottomArea.Name, bottomArea.Size)
	}
}

func TestStageSetDiagramInstancesNb(t *testing.T) {
	mux := http.NewServeMux()
	stageSet := models.NewStageSet("test_instances_nb")

	// Stage 2 A instances and 1 X instance
	(&models.A{Name: "A_First"}).Stage(stageSet.Stage)
	(&models.A{Name: "A_Second"}).Stage(stageSet.Stage)
	(&x.X{Name: "X_First"}).Stage(stageSet.XStage)

	p := probe.NewStageSetProbe(mux, embeddedgo.GoModelsDir, embeddedgo.GoDiagramsDir, true, stageSet)
	if p == nil {
		t.Fatal("expected non-nil StageSetProbe")
	}

	instancesMap := p.ComputeInstancesNb()
	if instancesMap["A"] != 2 || instancesMap["models.A"] != 2 {
		t.Errorf("expected count 2 for A, got A=%d models.A=%d", instancesMap["A"], instancesMap["models.A"])
	}
	if instancesMap["X"] != 1 || instancesMap["x.X"] != 1 {
		t.Errorf("expected count 1 for X, got X=%d x.X=%d", instancesMap["X"], instancesMap["x.X"])
	}

	docStager := p.GetDocStager()
	if docStager == nil {
		t.Fatal("expected non-nil docStager")
	}

	// Select StageSet_Diagram and ensure ShowNbInstances is true
	stage := docStager.GetStage()
	for _, dp := range stage.GetInstancesSorted[*doc_models.DiagramPackage]() {
		for _, cd := range dp.Classdiagrams {
			if cd.Name == "StageSet_Diagram" {
				dp.SelectedClassdiagram = cd
				cd.ShowNbInstances = true
				break
			}
		}
	}

	// Refresh probe to trigger SVG generation with instance counts
	p.Refresh()

	// Verify that svgStage contains RectAnchoredText with "(2)" for A and "(1)" for X
	svgStage := docStager.GetSvgStage()
	if svgStage == nil {
		t.Fatal("expected non-nil svgStage")
	}

	foundACount := false
	foundXCount := false
	for _, rect := range svgStage.GetInstancesSorted[*svg_models.Rect]() {
		if rect.Name == "A" {
			for _, text := range rect.RectAnchoredTexts {
				if text.Content == "(2)" {
					foundACount = true
				}
			}
		}
		if rect.Name == "X" {
			for _, text := range rect.RectAnchoredTexts {
				if text.Content == "(1)" {
					foundXCount = true
				}
			}
		}
	}

	if !foundACount {
		t.Errorf("expected anchored text '(2)' on Rect 'A'")
	}
	if !foundXCount {
		t.Errorf("expected anchored text '(1)' on Rect 'X'")
	}
}

func TestStageSetProbeTableStableOrder(t *testing.T) {
	mux := http.NewServeMux()
	stageSet := models.NewStageSet("test_table_order")

	// Stage elements in specific order
	a1 := (&models.A{Name: "A_System1"}).Stage(stageSet.Stage)
	a2 := (&models.A{Name: "A_System2"}).Stage(stageSet.Stage)
	a3 := (&models.A{Name: "A_System3"}).Stage(stageSet.Stage)

	p := probe.NewStageSetProbe(mux, embeddedgo.GoModelsDir, embeddedgo.GoDiagramsDir, true, stageSet)
	p.Refresh()

	// Open form for a1 which also sets up probe state, then call updateStageSetTable_A_Stage
	probe.StageSetFillUpFormFromGongstruct(a1, p)
	// Refresh probe updates tree and table
	p.Refresh()

	// Verify table rows for A
	tableStage := p.GetTableStage()
	for tbl := range *tableStage.GetInstancesSet[*table_models.Table]() {
		if tbl.Name == "A" {
			if len(tbl.Rows) != 3 {
				t.Fatalf("expected 3 rows in table A, got %d", len(tbl.Rows))
			}
			if tbl.Rows[0].Cells[0].CellInt.Value != 0 || tbl.Rows[0].Name != "A_System1" {
				t.Errorf("expected row 0 to be A_System1 with ID 0, got %s with ID %d", tbl.Rows[0].Name, tbl.Rows[0].Cells[0].CellInt.Value)
			}
			if tbl.Rows[1].Cells[0].CellInt.Value != 1 || tbl.Rows[1].Name != "A_System2" {
				t.Errorf("expected row 1 to be A_System2 with ID 1, got %s with ID %d", tbl.Rows[1].Name, tbl.Rows[1].Cells[0].CellInt.Value)
			}
			if tbl.Rows[2].Cells[0].CellInt.Value != 2 || tbl.Rows[2].Name != "A_System3" {
				t.Errorf("expected row 2 to be A_System3 with ID 2, got %s with ID %d", tbl.Rows[2].Name, tbl.Rows[2].Cells[0].CellInt.Value)
			}
		}
	}

	// Let's now add a new element "A0" which comes alphabetically before "A_System1"
	a0 := (&models.A{Name: "A0"}).Stage(stageSet.Stage)

	// Verify orders in stage
	if stageSet.Stage.GetOrder(a1) != 0 {
		t.Errorf("expected a1 order 0, got %d", stageSet.Stage.GetOrder(a1))
	}
	if stageSet.Stage.GetOrder(a2) != 1 {
		t.Errorf("expected a2 order 1, got %d", stageSet.Stage.GetOrder(a2))
	}
	if stageSet.Stage.GetOrder(a3) != 2 {
		t.Errorf("expected a3 order 2, got %d", stageSet.Stage.GetOrder(a3))
	}
	if stageSet.Stage.GetOrder(a0) != 3 {
		t.Errorf("expected newly staged a0 to have order 3, got %d", stageSet.Stage.GetOrder(a0))
	}

	// Refresh probe to regenerate table
	p.Refresh()

	// Verify table rows after adding A0: existing IDs (0, 1, 2) MUST NOT be shifted!
	// A0 should be at index 3 with ID 3
	for tbl := range *tableStage.GetInstancesSet[*table_models.Table]() {
		if tbl.Name == "A" {
			if len(tbl.Rows) != 4 {
				t.Fatalf("expected 4 rows in table A, got %d", len(tbl.Rows))
			}
			if tbl.Rows[0].Cells[0].CellInt.Value != 0 || tbl.Rows[0].Name != "A_System1" {
				t.Errorf("row 0 shifted! expected A_System1 with ID 0, got %s with ID %d", tbl.Rows[0].Name, tbl.Rows[0].Cells[0].CellInt.Value)
			}
			if tbl.Rows[1].Cells[0].CellInt.Value != 1 || tbl.Rows[1].Name != "A_System2" {
				t.Errorf("row 1 shifted! expected A_System2 with ID 1, got %s with ID %d", tbl.Rows[1].Name, tbl.Rows[1].Cells[0].CellInt.Value)
			}
			if tbl.Rows[2].Cells[0].CellInt.Value != 2 || tbl.Rows[2].Name != "A_System3" {
				t.Errorf("row 2 shifted! expected A_System3 with ID 2, got %s with ID %d", tbl.Rows[2].Name, tbl.Rows[2].Cells[0].CellInt.Value)
			}
			if tbl.Rows[3].Cells[0].CellInt.Value != 3 || tbl.Rows[3].Name != "A0" {
				t.Errorf("expected newly added A0 to have ID 3 at index 3, got %s with ID %d", tbl.Rows[3].Name, tbl.Rows[3].Cells[0].CellInt.Value)
			}
		}
	}

	// Verify GetInstancesByOrder preserves order 0, 1, 2, 3
	ordered := stageSet.Stage.GetInstancesByOrder[*models.A]()
	if len(ordered) != 4 {
		t.Fatalf("expected 4 ordered instances, got %d", len(ordered))
	}
	if ordered[0] != a1 || ordered[1] != a2 || ordered[2] != a3 || ordered[3] != a0 {
		t.Errorf("instances not in staging order: %v, %v, %v, %v",
			ordered[0].GetName(), ordered[1].GetName(), ordered[2].GetName(), ordered[3].GetName())
	}
}
