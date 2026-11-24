package main

import (
	"math"
	"slices"
	"strings"
	"testing"

	"github.com/fullstack-lang/gong/dsm/structure/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/structure/go/models"
	form "github.com/fullstack-lang/gong/lib/form/go/models"
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func hdStructFindTreeNodeByName(nodes []*tree.Node, name string) *tree.Node {
	for _, n := range nodes {
		if n.Name == name {
			return n
		}
		if found := hdStructFindTreeNodeByName(n.Children, name); found != nil {
			return found
		}
	}
	return nil
}

func hdStructFindRectByName(svgObj *svg.SVG, name string) *svg.Rect {
	if svgObj == nil || len(svgObj.Layers) == 0 {
		return nil
	}
	for _, r := range svgObj.Layers[0].Rects {
		if r.Name == name {
			return r
		}
	}
	return nil
}

func hdStructHasHierarchyConnector(r *svg.Rect) bool {
	if r == nil {
		return false
	}
	for _, p := range r.RectAnchoredPaths {
		if strings.HasSuffix(p.Name, "_hierarchy_connector") {
			if p.StrokeOpacity != 1.0 || p.Color != "none" || p.ScalePropotionnally {
				return false
			}
			return true
		}
	}
	return false
}

func TestHierarchicalDataStructureEndToEnd(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("test_hierarchical_data_structure", "", "", true, true, false)
	stager := models.NewStager(
		stack.R,
		stack.Stage,
		stack.Probe,
	)
	stage := stack.Stage
	rootLib := stager.GetRootLibrary()
	rootLib.Name = "RootLib"
	rootLib.IsDatasNodeExpanded = true

	// 1. Build hierarchical Data: SEMP -> [PBS -> [HW PBS, SW PBS], WBS], plus an orphan Data.
	hwPBS := (&models.Data{Name: "HW PBS"}).Stage(stage)
	swPBS := (&models.Data{Name: "SW PBS"}).Stage(stage)
	pbs := (&models.Data{
		Name:     "PBS",
		SubDatas: []*models.Data{hwPBS, swPBS},
	}).Stage(stage)
	wbs := (&models.Data{Name: "WBS"}).Stage(stage)
	semp := (&models.Data{
		Name:     "SEMP",
		SubDatas: []*models.Data{pbs, wbs},
	}).Stage(stage)
	orphanData := (&models.Data{Name: "Orphan Plan"}).Stage(stage)

	// Deliberately also put PBS in RootDatas to verify exclusivity cleanup
	rootLib.RootDatas = append(rootLib.RootDatas, semp, pbs)

	// 2. Build System, Parts, Ports, DataFlow carrying SEMP, and DiagramStructure
	sys := (&models.System{Name: "Engineering System"}).Stage(stage)
	rootLib.RootSystemes = append(rootLib.RootSystemes, sys)

	partA := (&models.Part{Name: "PartA"}).Stage(stage)
	partB := (&models.Part{Name: "PartB"}).Stage(stage)
	sys.Parts = append(sys.Parts, partA, partB)

	portA := (&models.Port{Name: "PA1"}).Stage(stage)
	portB := (&models.Port{Name: "PB1"}).Stage(stage)
	partA.Ports = append(partA.Ports, portA)
	partB.Ports = append(partB.Ports, portB)

	df := (&models.DataFlow{
		Name:      "Plans Flow",
		Type:      models.DataFlow_Port2Port,
		StartPort: portA,
		EndPort:   portB,
		Datas:     []*models.Data{semp},
	}).Stage(stage)
	rootLib.RootDataFlows = append(rootLib.RootDataFlows, df)
	sys.DataFlows = append(sys.DataFlows, df)

	diag := (&models.DiagramStructure{
		Name:        "MainStructDiag",
		IsChecked:   true,
		IsEditable_: true,
		Width:       1200,
		Height:      900,
	}).Stage(stage)
	sys.DiagramStructures = append(sys.DiagramStructures, diag)

	sysShape := (&models.SystemShape{
		Name:      "Engineering System-MainStructDiag",
		System:    sys,
		RectShape: models.RectShape{X: 50, Y: 50, Width: 800, Height: 600},
	}).Stage(stage)
	diag.System_Shapes = append(diag.System_Shapes, sysShape)

	partAShape := (&models.PartShape{
		Name:      "PartA-MainStructDiag",
		Part:      partA,
		RectShape: models.RectShape{X: 100, Y: 200, Width: 180, Height: 140},
	}).Stage(stage)
	partBShape := (&models.PartShape{
		Name:      "PartB-MainStructDiag",
		Part:      partB,
		RectShape: models.RectShape{X: 520, Y: 200, Width: 180, Height: 140},
	}).Stage(stage)
	diag.Part_Shapes = append(diag.Part_Shapes, partAShape, partBShape)

	portAShape := (&models.PortShape{
		Name:            "PA1-MainStructDiag",
		Port:            portA,
		ParentPartShape: partAShape,
		RectShape:       models.RectShape{X: 260, Y: 250, Width: 40, Height: 40},
	}).Stage(stage)
	portBShape := (&models.PortShape{
		Name:            "PB1-MainStructDiag",
		Port:            portB,
		ParentPartShape: partBShape,
		RectShape:       models.RectShape{X: 500, Y: 250, Width: 40, Height: 40},
	}).Stage(stage)
	diag.Port_Shapes = append(diag.Port_Shapes, portAShape, portBShape)

	dfShape := (&models.DataFlowShape{
		Name:           df.Name,
		DataFlow:       df,
		StartPortShape: portAShape,
		EndPortShape:   portBShape,
	}).Stage(stage)
	diag.DataFlow_Shapes = append(diag.DataFlow_Shapes, dfShape)

	stage.Commit()

	// Verify semantic enforcement
	if slices.Contains(rootLib.RootDatas, pbs) {
		t.Fatalf("expected sub-data PBS removed from RootDatas")
	}
	if !slices.Contains(rootLib.RootDatas, semp) || !slices.Contains(rootLib.RootDatas, orphanData) {
		t.Fatalf("expected SEMP and Orphan Plan in RootDatas, got %v", rootLib.RootDatas)
	}

	// Verify strict-tree enforcement (self-loop, cycle, duplicate, multi-parent)
	wbs.SubDatas = append(wbs.SubDatas, wbs)
	hwPBS.SubDatas = append(hwPBS.SubDatas, semp)
	semp.SubDatas = append(semp.SubDatas, pbs)
	wbs.SubDatas = append(wbs.SubDatas, swPBS)
	stage.Commit()

	if len(wbs.SubDatas) != 0 {
		t.Fatalf("expected self-loop and second-parent removed from WBS.SubDatas, got %v", wbs.SubDatas)
	}
	if len(hwPBS.SubDatas) != 0 {
		t.Fatalf("expected cycle removed from HW PBS.SubDatas, got %v", hwPBS.SubDatas)
	}
	if len(semp.SubDatas) != 2 {
		t.Fatalf("expected duplicate PBS removed from SEMP.SubDatas, got %d", len(semp.SubDatas))
	}

	// 3. Verify Library UX tree & "Add SubData" button
	trees := stager.GetTreeStage().GetInstancesSorted[*tree.Tree]()
	if len(trees) == 0 {
		t.Fatalf("expected UX tree to be populated")
	}
	rootLibNode := hdStructFindTreeNodeByName(trees[0].RootNodes, "RootLib")
	if rootLibNode == nil {
		t.Fatalf("expected RootLib node in UX tree")
	}
	var dataFolder *tree.Node
	for _, ch := range rootLibNode.Children {
		if ch.Name == "Data" {
			dataFolder = ch
			break
		}
	}
	if dataFolder == nil {
		t.Fatalf("expected Data folder directly under RootLib in Library UX tree")
	}
	sempLibNode := hdStructFindTreeNodeByName(dataFolder.Children, "SEMP")
	if sempLibNode == nil {
		t.Fatalf("expected SEMP under Library -> Data")
	}
	subDatasFolder := hdStructFindTreeNodeByName(sempLibNode.Children, "Sub Datas")
	if subDatasFolder == nil {
		t.Fatalf("expected Sub Datas folder under SEMP")
	}
	wbsLibNode := hdStructFindTreeNodeByName(subDatasFolder.Children, "WBS")
	if wbsLibNode == nil || wbsLibNode.Menu == nil {
		t.Fatalf("expected WBS node with Menu under SEMP -> Sub Datas")
	}
	var addSubBtn *tree.Button
	for _, b := range wbsLibNode.Menu.Buttons {
		if b.Name == "Add SubData" {
			addSubBtn = b
			break
		}
	}
	if addSubBtn == nil {
		t.Fatalf("expected Add SubData button on WBS menu")
	}
	addSubBtn.OnClick()
	if len(wbs.SubDatas) != 1 || !wbs.IsSubDatasNodeExpanded {
		t.Fatalf("expected 1 new SubData under WBS and IsSubDatasNodeExpanded=true")
	}
	newWBSSub := wbs.SubDatas[0]
	newWBSSub.Name = "WBS Phase 1"
	newWBSSub.SetIsInRenameMode(false)
	stage.Commit()

	// 4. Diagram UX tree: check SEMP, PBS, HW PBS, WBS
	checkDataInDiagram := func(dataName string, checked bool) {
		t.Helper()
		trees = stager.GetTreeStage().GetInstancesSorted[*tree.Tree]()
		diagNode := hdStructFindTreeNodeByName(trees[0].RootNodes, "MainStructDiag")
		if diagNode == nil {
			t.Fatalf("expected MainStructDiag node in UX tree")
		}
		dfFolder := hdStructFindTreeNodeByName(diagNode.Children, "Data Flows")
		if dfFolder == nil {
			t.Fatalf("expected Data Flows node under MainStructDiag")
		}
		dfNode := hdStructFindTreeNodeByName(dfFolder.Children, df.Name)
		if dfNode == nil {
			t.Fatalf("expected DataFlow node %q", df.Name)
		}
		targetNode := hdStructFindTreeNodeByName(dfNode.Children, dataName)
		if targetNode == nil || !targetNode.HasCheckboxButton {
			t.Fatalf("expected checkable Data node %q under DataFlow", dataName)
		}
		targetNode.OnIsCheckedChanged(checked)
	}

	checkDataInDiagram("SEMP", true)
	checkDataInDiagram("PBS", true)
	checkDataInDiagram("HW PBS", true)
	checkDataInDiagram("WBS", true)

	svgObj := stager.GetSvgObject()
	sempRect := hdStructFindRectByName(svgObj, "SEMP")
	pbsRect := hdStructFindRectByName(svgObj, "PBS")
	hwRect := hdStructFindRectByName(svgObj, "HW PBS")
	wbsRect := hdStructFindRectByName(svgObj, "WBS")
	if sempRect == nil || pbsRect == nil || hwRect == nil || wbsRect == nil {
		t.Fatalf("expected SVG rects for SEMP, PBS, HW PBS, WBS; got %v %v %v %v", sempRect, pbsRect, hwRect, wbsRect)
	}
	if hdStructFindRectByName(svgObj, "SW PBS") != nil {
		t.Fatalf("did not expect unchecked SW PBS in SVG")
	}

	// Verify vertical order, indentation, right-aligned default widths, and connectors
	if !(sempRect.Y < pbsRect.Y && pbsRect.Y < hwRect.Y && hwRect.Y < wbsRect.Y) {
		t.Fatalf("expected vertical pre-order SEMP(%.1f) < PBS(%.1f) < HW(%.1f) < WBS(%.1f)",
			sempRect.Y, pbsRect.Y, hwRect.Y, wbsRect.Y)
	}
	if math.Abs((pbsRect.X-sempRect.X)-16.0) > 0.01 ||
		math.Abs((hwRect.X-sempRect.X)-32.0) > 0.01 ||
		math.Abs((wbsRect.X-sempRect.X)-16.0) > 0.01 {
		t.Fatalf("unexpected indentation X values: SEMP=%.1f PBS=%.1f HW=%.1f WBS=%.1f",
			sempRect.X, pbsRect.X, hwRect.X, wbsRect.X)
	}
	if math.Abs(sempRect.Width-180.0) > 0.01 || math.Abs(pbsRect.Width-164.0) > 0.01 || math.Abs(hwRect.Width-148.0) > 0.01 {
		t.Fatalf("expected right-aligned default widths 180/164/148, got %.1f/%.1f/%.1f",
			sempRect.Width, pbsRect.Width, hwRect.Width)
	}
	if hdStructHasHierarchyConnector(sempRect) {
		t.Fatalf("did not expect hierarchy connector on root SEMP rect")
	}
	if !hdStructHasHierarchyConnector(pbsRect) || !hdStructHasHierarchyConnector(hwRect) || !hdStructHasHierarchyConnector(wbsRect) {
		t.Fatalf("expected hierarchy connectors on PBS, HW PBS, and WBS rects")
	}

	// 5. Sparse display: uncheck SEMP -> PBS & WBS promoted to depth 0, HW PBS at depth 1 under PBS
	checkDataInDiagram("SEMP", false)
	svgObj = stager.GetSvgObject()
	if hdStructFindRectByName(svgObj, "SEMP") != nil {
		t.Fatalf("expected SEMP rect removed after uncheck")
	}
	pbsRect = hdStructFindRectByName(svgObj, "PBS")
	hwRect = hdStructFindRectByName(svgObj, "HW PBS")
	wbsRect = hdStructFindRectByName(svgObj, "WBS")
	if pbsRect == nil || hwRect == nil || wbsRect == nil {
		t.Fatalf("expected PBS, HW PBS, WBS still rendered when SEMP is unchecked")
	}
	if math.Abs(pbsRect.X-wbsRect.X) > 0.01 || math.Abs((hwRect.X-pbsRect.X)-16.0) > 0.01 {
		t.Fatalf("unexpected sparse indentation: PBS=%.1f HW=%.1f WBS=%.1f", pbsRect.X, hwRect.X, wbsRect.X)
	}
	if hdStructHasHierarchyConnector(pbsRect) || hdStructHasHierarchyConnector(wbsRect) || !hdStructHasHierarchyConnector(hwRect) {
		t.Fatalf("unexpected connector presence in sparse mode")
	}

	// 6. Bulk display & remove via menu buttons on SEMP node in Diagram UX tree
	trees = stager.GetTreeStage().GetInstancesSorted[*tree.Tree]()
	diagNode := hdStructFindTreeNodeByName(trees[0].RootNodes, "MainStructDiag")
	dfFolder := hdStructFindTreeNodeByName(diagNode.Children, "Data Flows")
	dfNode := hdStructFindTreeNodeByName(dfFolder.Children, df.Name)
	sempDiagNode := hdStructFindTreeNodeByName(dfNode.Children, "SEMP")
	if sempDiagNode == nil || sempDiagNode.Menu == nil {
		t.Fatalf("expected SEMP node with menu in diagram tree")
	}
	var displayAllBtn, removeAllBtn *tree.Button
	for _, b := range sempDiagNode.Menu.Buttons {
		if b.Name == "Display data & sub-datas" {
			displayAllBtn = b
		}
		if b.Name == "Remove data & sub-datas" {
			removeAllBtn = b
		}
	}
	if displayAllBtn == nil || removeAllBtn == nil {
		t.Fatalf("expected bulk display/remove menu buttons on SEMP")
	}
	displayAllBtn.OnClick()
	if len(diag.Data_Shapes) != 6 {
		t.Fatalf("expected 6 DataShapes after bulk display, got %d", len(diag.Data_Shapes))
	}

	// 6b. Probe form update on DataFlow.Datas preserves existing SubData DataShapes
	// and auto-adds a DataShape only for the newly added top-level Data (Orphan Plan).
	stack.Probe.FillUpFormFromGongstruct(df, "DataFlow")
	formStage := stack.Probe.GetFormStage()
	fg := formStage.GetInstancesSorted[*form.FormGroup]()[0]
	setFormAssocSelection(t, stage, fg, "Datas", []*models.Data{semp, orphanData})
	formStage.Commit()
	fg.OnSave.OnSave()

	sempHierarchyCount := 0
	foundOrphanShape := false
	for _, ds := range diag.Data_Shapes {
		if ds.Data == semp || ds.Data == pbs || ds.Data == hwPBS || ds.Data == swPBS || ds.Data == wbs || ds.Data == newWBSSub {
			sempHierarchyCount++
		}
		if ds.Data == orphanData {
			foundOrphanShape = true
		}
	}
	if sempHierarchyCount != 6 || !foundOrphanShape {
		t.Fatalf("expected all 6 SEMP hierarchy DataShapes preserved and Orphan Plan shape created, got sempCount=%d orphan=%v",
			sempHierarchyCount, foundOrphanShape)
	}
	// Remove Orphan Plan via Probe form and verify SEMP hierarchy DataShapes stay intact
	stack.Probe.FillUpFormFromGongstruct(df, "DataFlow")
	fg = formStage.GetInstancesSorted[*form.FormGroup]()[0]
	setFormAssocSelection(t, stage, fg, "Datas", []*models.Data{semp})
	formStage.Commit()
	fg.OnSave.OnSave()
	if len(diag.Data_Shapes) != 6 {
		t.Fatalf("expected 6 SEMP hierarchy DataShapes after removing Orphan Plan via Probe, got %d", len(diag.Data_Shapes))
	}

	// 7. Semantic cleanup when detaching WBS from SEMP.SubDatas
	semp.SubDatas = []*models.Data{pbs}
	stage.Commit()
	if !slices.Contains(rootLib.RootDatas, wbs) {
		t.Fatalf("expected detached WBS reattached to RootDatas")
	}
	if len(diag.Data_Shapes) != 4 {
		t.Fatalf("expected 4 remaining DataShapes after detaching WBS, got %d", len(diag.Data_Shapes))
	}

	// 8. Bulk remove clears all remaining SEMP subtree DataShapes
	trees = stager.GetTreeStage().GetInstancesSorted[*tree.Tree]()
	diagNode = hdStructFindTreeNodeByName(trees[0].RootNodes, "MainStructDiag")
	dfFolder = hdStructFindTreeNodeByName(diagNode.Children, "Data Flows")
	dfNode = hdStructFindTreeNodeByName(dfFolder.Children, df.Name)
	sempDiagNode = hdStructFindTreeNodeByName(dfNode.Children, "SEMP")
	for _, b := range sempDiagNode.Menu.Buttons {
		if b.Name == "Remove data & sub-datas" {
			b.OnClick()
			break
		}
	}
	if len(diag.Data_Shapes) != 0 {
		t.Fatalf("expected 0 DataShapes after bulk remove, got %d", len(diag.Data_Shapes))
	}
}
