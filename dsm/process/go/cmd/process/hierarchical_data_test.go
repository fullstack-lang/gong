package main

import (
	"math"
	"slices"
	"strings"
	"testing"

	"github.com/fullstack-lang/gong/dsm/process/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/process/go/models"
	form "github.com/fullstack-lang/gong/lib/form/go/models"
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func hdProcessFindTreeNodeByName(nodes []*tree.Node, name string) *tree.Node {
	for _, n := range nodes {
		if n.Name == name {
			return n
		}
		if found := hdProcessFindTreeNodeByName(n.Children, name); found != nil {
			return found
		}
	}
	return nil
}

func hdProcessFindRectByName(svgObj *svg.SVG, name string) *svg.Rect {
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

func hdProcessHasHierarchyConnector(r *svg.Rect) bool {
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

func TestHierarchicalDataProcessEndToEnd(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("test_hierarchical_data_process", "", "", true, true, false)
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
	// Deliberately place PBS in RootDatas too to verify strict-tree exclusivity removes it from RootDatas.
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

	rootLib.RootDatas = append(rootLib.RootDatas, semp, pbs)

	// 2. Build Process, Participants, Tasks, DataFlow carrying SEMP, and DiagramProcess
	proc := (&models.Process{Name: "Engineering Process"}).Stage(stage)
	rootLib.RootProcesses = append(rootLib.RootProcesses, proc)

	partA := (&models.Participant{Name: "Systems Eng"}).Stage(stage)
	partB := (&models.Participant{Name: "Project Mgmt"}).Stage(stage)
	proc.Participants = append(proc.Participants, partA, partB)

	taskA := (&models.Task{Name: "Define Plans"}).Stage(stage)
	taskB := (&models.Task{Name: "Execute Plans"}).Stage(stage)
	partA.Tasks = append(partA.Tasks, taskA)
	partB.Tasks = append(partB.Tasks, taskB)

	df := (&models.DataFlow{
		Name:      "Plans Flow",
		Type:      models.DataFlow_Task2Task,
		StartTask: taskA,
		EndTask:   taskB,
		Datas:     []*models.Data{semp},
	}).Stage(stage)
	rootLib.RootDataFlows = append(rootLib.RootDataFlows, df)
	proc.DataFlows = append(proc.DataFlows, df)

	diag := (&models.DiagramProcess{
		Name:        "MainDiag",
		IsChecked:   true,
		IsEditable_: true,
		Width:       1200,
		Height:      900,
	}).Stage(stage)
	proc.DiagramProcesss = append(proc.DiagramProcesss, diag)

	procShape := (&models.ProcessShape{
		Name:   "Engineering Process",
		Process: proc,
		X:      50,
		Y:      50,
		Width:  800,
		Height: 600,
	}).Stage(stage)
	diag.Process_Shapes = append(diag.Process_Shapes, procShape)

	pAShape := (&models.ParticipantShape{Name: "Systems Eng", Participant: partA}).Stage(stage)
	pBShape := (&models.ParticipantShape{Name: "Project Mgmt", Participant: partB}).Stage(stage)
	diag.Participant_Shapes = append(diag.Participant_Shapes, pAShape, pBShape)

	tAShape := (&models.TaskShape{
		Name:      "Define Plans",
		Task:      taskA,
		RectShape: models.RectShape{X: 100, Y: 250, Width: 180, Height: 60},
	}).Stage(stage)
	tBShape := (&models.TaskShape{
		Name:      "Execute Plans",
		Task:      taskB,
		RectShape: models.RectShape{X: 520, Y: 250, Width: 180, Height: 60},
	}).Stage(stage)
	diag.Task_Shapes = append(diag.Task_Shapes, tAShape, tBShape)

	dfShape := (&models.DataFlowShape{
		Name:     df.Name,
		DataFlow: df,
		LinkShape: models.LinkShape{
			StartOrientation: models.ORIENTATION_HORIZONTAL,
			EndOrientation:   models.ORIENTATION_HORIZONTAL,
			CornerOffsetRatio: 1.5,
			StartRatio:        0.5,
			EndRatio:          0.5,
		},
	}).Stage(stage)
	diag.DataFlow_Shapes = append(diag.DataFlow_Shapes, dfShape)

	stage.Commit()

	// Verify semantic enforcement:
	// - PBS was removed from RootDatas because it is a SubData of SEMP
	// - Orphan Plan was reattached to RootDatas
	if slices.Contains(rootLib.RootDatas, pbs) {
		t.Fatalf("expected sub-data PBS to be removed from RootDatas")
	}
	if !slices.Contains(rootLib.RootDatas, semp) || !slices.Contains(rootLib.RootDatas, orphanData) {
		t.Fatalf("expected SEMP and Orphan Plan in RootDatas, got %v", rootLib.RootDatas)
	}

	// Verify strict-tree enforcement: self-loop, cycle, duplicate, and multi-parent are all broken
	wbs.SubDatas = append(wbs.SubDatas, wbs)       // self-loop
	hwPBS.SubDatas = append(hwPBS.SubDatas, semp)  // cycle HW PBS -> SEMP
	semp.SubDatas = append(semp.SubDatas, pbs)     // duplicate child
	wbs.SubDatas = append(wbs.SubDatas, swPBS)     // second parent (PBS already owns SW PBS)
	stage.Commit()

	if len(wbs.SubDatas) != 0 {
		t.Fatalf("expected self-loop and second-parent edge removed from WBS.SubDatas, got %v", wbs.SubDatas)
	}
	if len(hwPBS.SubDatas) != 0 {
		t.Fatalf("expected cycle edge removed from HW PBS.SubDatas, got %v", hwPBS.SubDatas)
	}
	if len(semp.SubDatas) != 2 {
		t.Fatalf("expected duplicate PBS removed from SEMP.SubDatas (len 2), got %d", len(semp.SubDatas))
	}
	if !slices.Contains(pbs.SubDatas, swPBS) {
		t.Fatalf("expected SW PBS to remain under its first parent PBS")
	}

	// 3. Verify Library UX tree shows hierarchy and "Add SubData" button creates a child Data
	trees := stager.GetTreeStage().GetInstancesSorted[*tree.Tree]()
	if len(trees) == 0 {
		t.Fatalf("expected UX tree to be populated")
	}
	rootLibNode := hdProcessFindTreeNodeByName(trees[0].RootNodes, "RootLib")
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
	sempLibNode := hdProcessFindTreeNodeByName(dataFolder.Children, "SEMP")
	if sempLibNode == nil {
		t.Fatalf("expected SEMP node under Library -> Data")
	}
	subDatasFolder := hdProcessFindTreeNodeByName(sempLibNode.Children, "Sub Datas")
	if subDatasFolder == nil {
		t.Fatalf("expected Sub Datas folder under SEMP")
	}
	pbsLibNode := hdProcessFindTreeNodeByName(subDatasFolder.Children, "PBS")
	if pbsLibNode == nil {
		t.Fatalf("expected PBS under SEMP -> Sub Datas")
	}
	if hdProcessFindTreeNodeByName(pbsLibNode.Children, "HW PBS") == nil {
		t.Fatalf("expected HW PBS under PBS -> Sub Datas")
	}

	// Click "Add SubData" on WBS in Library tree
	wbsLibNode := hdProcessFindTreeNodeByName(subDatasFolder.Children, "WBS")
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
		t.Fatalf("expected 1 new SubData under WBS and IsSubDatasNodeExpanded=true, got len=%d expanded=%v",
			len(wbs.SubDatas), wbs.IsSubDatasNodeExpanded)
	}
	newWBSSub := wbs.SubDatas[0]
	newWBSSub.Name = "WBS Phase 1"
	newWBSSub.SetIsInRenameMode(false)
	stage.Commit()

	// 4. Diagram UX tree: check SEMP, PBS, HW PBS, WBS on the DataFlow
	checkDataInDiagram := func(dataName string, checked bool) {
		t.Helper()
		trees = stager.GetTreeStage().GetInstancesSorted[*tree.Tree]()
		diagNode := hdProcessFindTreeNodeByName(trees[0].RootNodes, "MainDiag")
		if diagNode == nil {
			t.Fatalf("expected MainDiag node in UX tree")
		}
		dfFolder := hdProcessFindTreeNodeByName(diagNode.Children, "Data Flows")
		if dfFolder == nil {
			t.Fatalf("expected Data Flows node under MainDiag")
		}
		dfNode := hdProcessFindTreeNodeByName(dfFolder.Children, df.Name)
		if dfNode == nil {
			t.Fatalf("expected DataFlow node %q under Data Flows", df.Name)
		}
		targetNode := hdProcessFindTreeNodeByName(dfNode.Children, dataName)
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
	sempRect := hdProcessFindRectByName(svgObj, "SEMP")
	pbsRect := hdProcessFindRectByName(svgObj, "PBS")
	hwRect := hdProcessFindRectByName(svgObj, "HW PBS")
	wbsRect := hdProcessFindRectByName(svgObj, "WBS")
	if sempRect == nil || pbsRect == nil || hwRect == nil || wbsRect == nil {
		t.Fatalf("expected SVG rects for SEMP, PBS, HW PBS, WBS; got %v %v %v %v", sempRect, pbsRect, hwRect, wbsRect)
	}
	if hdProcessFindRectByName(svgObj, "SW PBS") != nil {
		t.Fatalf("did not expect unchecked SW PBS in SVG")
	}

	// Verify pre-order vertical stacking: SEMP < PBS < HW PBS < WBS
	if !(sempRect.Y < pbsRect.Y && pbsRect.Y < hwRect.Y && hwRect.Y < wbsRect.Y) {
		t.Fatalf("expected vertical pre-order SEMP(%.1f) < PBS(%.1f) < HW(%.1f) < WBS(%.1f)",
			sempRect.Y, pbsRect.Y, hwRect.Y, wbsRect.Y)
	}

	// Verify indentation (16px per depth level) and right-aligned default widths
	if math.Abs((pbsRect.X-sempRect.X)-16.0) > 0.01 {
		t.Fatalf("expected PBS indented by 16px from SEMP, got %.2f", pbsRect.X-sempRect.X)
	}
	if math.Abs((hwRect.X-sempRect.X)-32.0) > 0.01 {
		t.Fatalf("expected HW PBS indented by 32px from SEMP, got %.2f", hwRect.X-sempRect.X)
	}
	if math.Abs((wbsRect.X-sempRect.X)-16.0) > 0.01 {
		t.Fatalf("expected WBS indented by 16px from SEMP, got %.2f", wbsRect.X-sempRect.X)
	}
	if math.Abs(sempRect.Width-180.0) > 0.01 || math.Abs(pbsRect.Width-164.0) > 0.01 || math.Abs(hwRect.Width-148.0) > 0.01 {
		t.Fatalf("expected right-aligned default widths 180/164/148, got %.1f/%.1f/%.1f",
			sempRect.Width, pbsRect.Width, hwRect.Width)
	}

	// Verify hierarchy connectors
	if hdProcessHasHierarchyConnector(sempRect) {
		t.Fatalf("did not expect hierarchy connector on root SEMP rect")
	}
	if !hdProcessHasHierarchyConnector(pbsRect) || !hdProcessHasHierarchyConnector(hwRect) || !hdProcessHasHierarchyConnector(wbsRect) {
		t.Fatalf("expected hierarchy connectors on PBS, HW PBS, and WBS rects")
	}

	// 5. Sparse display: uncheck SEMP; PBS & WBS become depth 0, HW PBS becomes depth 1 under PBS
	checkDataInDiagram("SEMP", false)
	svgObj = stager.GetSvgObject()
	if hdProcessFindRectByName(svgObj, "SEMP") != nil {
		t.Fatalf("expected SEMP rect removed after uncheck")
	}
	pbsRect = hdProcessFindRectByName(svgObj, "PBS")
	hwRect = hdProcessFindRectByName(svgObj, "HW PBS")
	wbsRect = hdProcessFindRectByName(svgObj, "WBS")
	if pbsRect == nil || hwRect == nil || wbsRect == nil {
		t.Fatalf("expected PBS, HW PBS, WBS still rendered when SEMP is unchecked")
	}
	if math.Abs(pbsRect.X-wbsRect.X) > 0.01 {
		t.Fatalf("expected PBS and WBS at same root X when SEMP is unchecked, got %.2f vs %.2f", pbsRect.X, wbsRect.X)
	}
	if math.Abs((hwRect.X-pbsRect.X)-16.0) > 0.01 {
		t.Fatalf("expected HW PBS indented 16px from PBS, got %.2f", hwRect.X-pbsRect.X)
	}
	if hdProcessHasHierarchyConnector(pbsRect) || hdProcessHasHierarchyConnector(wbsRect) {
		t.Fatalf("expected no connector on promoted depth-0 PBS/WBS rects")
	}
	if !hdProcessHasHierarchyConnector(hwRect) {
		t.Fatalf("expected connector on depth-1 HW PBS rect")
	}

	// 6. Bulk display & remove via menu buttons on SEMP node in Diagram UX tree
	trees = stager.GetTreeStage().GetInstancesSorted[*tree.Tree]()
	diagNode := hdProcessFindTreeNodeByName(trees[0].RootNodes, "MainDiag")
	dfFolder := hdProcessFindTreeNodeByName(diagNode.Children, "Data Flows")
	dfNode := hdProcessFindTreeNodeByName(dfFolder.Children, df.Name)
	sempDiagNode := hdProcessFindTreeNodeByName(dfNode.Children, "SEMP")
	if sempDiagNode == nil || sempDiagNode.Menu == nil {
		t.Fatalf("expected SEMP node with menu in diagram tree")
	}
	var displayAllBtn *tree.Button
	for _, b := range sempDiagNode.Menu.Buttons {
		if b.Name == "Display data & sub-datas" {
			displayAllBtn = b
			break
		}
	}
	if displayAllBtn == nil {
		t.Fatalf("expected 'Display data & sub-datas' menu button on SEMP")
	}
	displayAllBtn.OnClick()
	if len(diag.Data_Shapes) != 6 {
		t.Fatalf("expected all 6 Data nodes in SEMP hierarchy to have DataShapes, got %d", len(diag.Data_Shapes))
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

	// 7. Semantic cleanup: detaching WBS from SEMP.SubDatas removes WBS & WBS Phase 1 DataShapes and reattaches WBS to RootDatas
	semp.SubDatas = []*models.Data{pbs}
	stage.Commit()
	if !slices.Contains(rootLib.RootDatas, wbs) {
		t.Fatalf("expected detached WBS to be reattached to RootDatas")
	}
	for _, ds := range diag.Data_Shapes {
		if ds.Data == wbs || ds.Data == newWBSSub {
			t.Fatalf("expected stale DataShape for detached %s to be unstaged", ds.Data.Name)
		}
	}
	if len(diag.Data_Shapes) != 4 {
		t.Fatalf("expected 4 remaining DataShapes (SEMP, PBS, HW PBS, SW PBS), got %d", len(diag.Data_Shapes))
	}

	// 8. Removing SEMP from DataFlow.Datas cleans up all remaining DataShapes
	df.Datas = nil
	stage.Commit()
	if len(diag.Data_Shapes) != 0 {
		t.Fatalf("expected all DataShapes cleaned up when DataFlow.Datas is cleared, got %d", len(diag.Data_Shapes))
	}
}
