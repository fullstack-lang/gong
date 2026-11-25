package main

import (
	"testing"

	"github.com/fullstack-lang/gong/dsm/process/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/process/go/models"
	form "github.com/fullstack-lang/gong/lib/form/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func renderedTreeNodeLabel(node *tree.Node) string {
	if node == nil {
		return ""
	}
	if node.IsWithPrefix {
		return node.Prefix + " " + node.Name
	}
	return node.Name
}

func TestIssue1553_DisplayDataAcronymBesideNameInTree(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("test_data_acronym_tree_1553", "", "", true, true, false)
	stager := models.NewStager(stack.R, stack.Stage, stack.Probe)
	stage := stack.Stage

	rootLib := stager.GetRootLibrary()
	rootLib.Name = "RootLib"
	rootLib.IsExpandedTmp = true
	rootLib.IsDatasNodeExpanded = true

	// Hierarchical Data:
	// - Flight Plan (Acronym: "FPL")
	//   - Repetitive Flight Plan (Acronym: "RPL")
	//   - Route Details (Acronym: "")
	rplData := (&models.Data{
		Name:    "Repetitive Flight Plan",
		Acronym: "RPL",
	}).Stage(stage)
	routeData := (&models.Data{
		Name: "Route Details",
	}).Stage(stage)
	fplData := (&models.Data{
		Name:     "Flight Plan",
		Acronym:  "FPL",
		SubDatas: []*models.Data{rplData, routeData},
	}).Stage(stage)
	rootLib.RootDatas = append(rootLib.RootDatas, fplData)

	// DataStore with acronym storing Flight Plan
	fdpsStore := (&models.DataStore{
		Name:    "Flight Data Processing Store",
		Acronym: "FDPS",
		Datas:   []*models.Data{fplData},
	}).Stage(stage)
	rootLib.RootDataStores = append(rootLib.RootDataStores, fdpsStore)

	// Process, Participants, Tasks, DataFlow carrying Flight Plan, and DiagramProcess
	proc := (&models.Process{Name: "ATC Process"}).Stage(stage)
	rootLib.RootProcesses = append(rootLib.RootProcesses, proc)

	lane1 := (&models.Participant{Name: "Pilot"}).Stage(stage)
	lane2 := (&models.Participant{Name: "Controller"}).Stage(stage)
	proc.Participants = append(proc.Participants, lane1, lane2)

	taskSubmit := (&models.Task{Name: "Submit Plan"}).Stage(stage)
	taskValidate := (&models.Task{Name: "Validate Plan"}).Stage(stage)
	lane1.Tasks = append(lane1.Tasks, taskSubmit)
	lane2.Tasks = append(lane2.Tasks, taskValidate)

	df := (&models.DataFlow{
		Name:      "Plan Submission",
		Type:      models.DataFlow_Task2Task,
		StartTask: taskSubmit,
		EndTask:   taskValidate,
		Datas:     []*models.Data{fplData},
	}).Stage(stage)
	rootLib.RootDataFlows = append(rootLib.RootDataFlows, df)
	proc.DataFlows = append(proc.DataFlows, df)

	diag := (&models.DiagramProcess{
		Name:        "ATCView",
		IsChecked:   true,
		IsEditable_: true,
		Width:       1200,
		Height:      800,
	}).Stage(stage)
	proc.DiagramProcesss = append(proc.DiagramProcesss, diag)

	pShape := (&models.ProcessShape{
		Name:      "ATC Process-ATCView",
		Process:   proc,
		RectShape: models.RectShape{X: 100, Y: 50, Width: 800, Height: 600},
	}).Stage(stage)
	diag.Process_Shapes = append(diag.Process_Shapes, pShape)

	laneShape1 := (&models.ParticipantShape{
		Name:        "Pilot-ATCView",
		Participant: lane1,
		WidthWeight: 1.0,
	}).Stage(stage)
	laneShape2 := (&models.ParticipantShape{
		Name:        "Controller-ATCView",
		Participant: lane2,
		WidthWeight: 1.0,
	}).Stage(stage)
	diag.Participant_Shapes = append(diag.Participant_Shapes, laneShape1, laneShape2)

	ts1 := (&models.TaskShape{
		Name:      "Submit Plan-ATCView",
		Task:      taskSubmit,
		RectShape: models.RectShape{X: 150, Y: 220, Width: 200, Height: 60},
	}).Stage(stage)
	ts2 := (&models.TaskShape{
		Name:      "Validate Plan-ATCView",
		Task:      taskValidate,
		RectShape: models.RectShape{X: 550, Y: 220, Width: 200, Height: 60},
	}).Stage(stage)
	diag.Task_Shapes = append(diag.Task_Shapes, ts1, ts2)

	dfShape := (&models.DataFlowShape{
		Name:     df.Name,
		DataFlow: df,
		LinkShape: models.LinkShape{
			StartOrientation:  models.ORIENTATION_HORIZONTAL,
			EndOrientation:    models.ORIENTATION_HORIZONTAL,
			CornerOffsetRatio: 1.4,
			StartRatio:        0.5,
			EndRatio:          0.5,
		},
	}).Stage(stage)
	diag.DataFlow_Shapes = append(diag.DataFlow_Shapes, dfShape)

	stage.Commit()

	trees := stager.GetTreeStage().GetInstancesSorted[*tree.Tree]()
	if len(trees) == 0 {
		t.Fatalf("expected UX tree to be populated")
	}
	roots := trees[0].RootNodes

	assertNodeDisplay := func(path []string, wantRendered string, wantWithPrefix bool) *tree.Node {
		t.Helper()
		n := findTreeNodeByPath(roots, path...)
		if n == nil {
			t.Fatalf("expected tree node at path %v", path)
		}
		if n.IsWithPrefix != wantWithPrefix {
			t.Fatalf("path %v: expected IsWithPrefix=%v, got %v (Prefix=%q Name=%q)", path, wantWithPrefix, n.IsWithPrefix, n.Prefix, n.Name)
		}
		if got := renderedTreeNodeLabel(n); got != wantRendered {
			t.Fatalf("path %v: expected rendered label %q, got %q (Prefix=%q Name=%q)", path, wantRendered, got, n.Prefix, n.Name)
		}
		return n
	}

	// 1. Library -> Data & Sub Datas
	assertNodeDisplay([]string{"RootLib", "Data", "Flight Plan"}, "FPL / Flight Plan", true)
	assertNodeDisplay([]string{"RootLib", "Data", "Flight Plan", "Sub Datas", "Repetitive Flight Plan"}, "RPL / Repetitive Flight Plan", true)
	assertNodeDisplay([]string{"RootLib", "Data", "Flight Plan", "Sub Datas", "Route Details"}, "Route Details", false)

	// 2. Library -> Data Flows -> df.Name -> Data
	assertNodeDisplay([]string{"RootLib", "Data Flows", df.Name, "Data", "Flight Plan"}, "FPL / Flight Plan", true)
	assertNodeDisplay([]string{"RootLib", "Data Flows", df.Name, "Data", "Flight Plan", "Repetitive Flight Plan"}, "RPL / Repetitive Flight Plan", true)
	assertNodeDisplay([]string{"RootLib", "Data Flows", df.Name, "Data", "Flight Plan", "Route Details"}, "Route Details", false)

	// 3. Library -> Data Stores & Diagram -> Data Stores (including DataStore acronym and stored Data)
	assertNodeDisplay([]string{"RootLib", "Data Stores", "Flight Data Processing Store"}, "FDPS / Flight Data Processing Store", true)
	assertNodeDisplay([]string{"RootLib", "Data Stores", "Flight Data Processing Store", "Data", "Flight Plan"}, "FPL / Flight Plan", true)
	assertNodeDisplay([]string{"RootLib", "Processes", "ATC Process", "ATCView", "Data Stores", "Flight Data Processing Store"}, "FDPS / Flight Data Processing Store", true)
	assertNodeDisplay([]string{"RootLib", "Processes", "ATC Process", "ATCView", "Data Stores", "Flight Data Processing Store", "Data", "Flight Plan"}, "FPL / Flight Plan", true)

	// 4. Diagram -> Data Flows -> df.Name & Task -> out data flows -> df.Name
	diagFplNode := assertNodeDisplay([]string{"RootLib", "Processes", "ATC Process", "ATCView", "Data Flows", df.Name, "Flight Plan"}, "FPL / Flight Plan", true)
	assertNodeDisplay([]string{"RootLib", "Processes", "ATC Process", "ATCView", "Data Flows", df.Name, "Flight Plan", "Repetitive Flight Plan"}, "RPL / Repetitive Flight Plan", true)
	assertNodeDisplay([]string{"RootLib", "Processes", "ATC Process", "ATCView", "Data Flows", df.Name, "Flight Plan", "Route Details"}, "Route Details", false)
	assertNodeDisplay([]string{"RootLib", "Processes", "ATC Process", "ATCView", "Participants", "Pilot", "Tasks", "Submit Plan", "out data flows", df.Name, "Flight Plan"}, "FPL / Flight Plan", true)

	// 5. Inline rename on tree node preserves Acronym without duplicating "FPL /" into Data.Name
	diagFplNode.OnNameChange("ICAO Flight Plan")
	if fplData.Name != "ICAO Flight Plan" {
		t.Fatalf("expected renamed Data.Name to be 'ICAO Flight Plan', got %q", fplData.Name)
	}
	roots = stager.GetTreeStage().GetInstancesSorted[*tree.Tree]()[0].RootNodes
	assertNodeDisplay([]string{"RootLib", "Data", "ICAO Flight Plan"}, "FPL / ICAO Flight Plan", true)
	assertNodeDisplay([]string{"RootLib", "Processes", "ATC Process", "ATCView", "Data Flows", df.Name, "ICAO Flight Plan"}, "FPL / ICAO Flight Plan", true)

	// 6. Probe form edits: adding an acronym to Route Details and clearing acronym on Repetitive Flight Plan
	stack.Probe.FillUpFormFromGongstruct(routeData, "Data")
	formStage := stack.Probe.GetFormStage()
	fg := formStage.GetInstancesSorted[*form.FormGroup]()[0]
	setFormStringField(t, fg, "Acronym", "RD")
	formStage.Commit()
	fg.OnSave.OnSave()

	stack.Probe.FillUpFormFromGongstruct(rplData, "Data")
	fg = formStage.GetInstancesSorted[*form.FormGroup]()[0]
	setFormStringField(t, fg, "Acronym", "")
	formStage.Commit()
	fg.OnSave.OnSave()

	roots = stager.GetTreeStage().GetInstancesSorted[*tree.Tree]()[0].RootNodes
	assertNodeDisplay([]string{"RootLib", "Data", "ICAO Flight Plan", "Sub Datas", "Route Details"}, "RD / Route Details", true)
	assertNodeDisplay([]string{"RootLib", "Data", "ICAO Flight Plan", "Sub Datas", "Repetitive Flight Plan"}, "Repetitive Flight Plan", false)
	assertNodeDisplay([]string{"RootLib", "Processes", "ATC Process", "ATCView", "Data Flows", df.Name, "ICAO Flight Plan", "Route Details"}, "RD / Route Details", true)
	assertNodeDisplay([]string{"RootLib", "Processes", "ATC Process", "ATCView", "Data Flows", df.Name, "ICAO Flight Plan", "Repetitive Flight Plan"}, "Repetitive Flight Plan", false)

	// 7. Diagram IsShowPrefix + ComputedPrefix composes cleanly with Acronym
	fplData.SetComputedPrefix("1.2")
	diag.SetIsShowPrefix(true)
	stage.Commit()

	roots = stager.GetTreeStage().GetInstancesSorted[*tree.Tree]()[0].RootNodes
	assertNodeDisplay([]string{"RootLib", "Processes", "ATC Process", "ATCView", "Data Flows", df.Name, "ICAO Flight Plan"}, "1.2 FPL / ICAO Flight Plan", true)
}
