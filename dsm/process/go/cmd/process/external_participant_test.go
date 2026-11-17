package main

import (
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/fullstack-lang/gong/dsm/process/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/process/go/models"
	form "github.com/fullstack-lang/gong/lib/form/go/models"
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func TestIssue1541_ExternalParticipantConceptAndDataFlows(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("test_ext_participant", "", "", true, true, false)
	stager := models.NewStager(stack.R, stack.Stage, stack.Probe)
	stage := stack.Stage

	rootLib := stager.GetRootLibrary()
	rootLib.Name = "RootLib"

	proc := (&models.Process{Name: "MainProcess"}).Stage(stage)
	rootLib.RootProcesses = append(rootLib.RootProcesses, proc)

	diag := (&models.DiagramProcess{
		Name:             "MainDiagram",
		IsChecked:        true,
		IsEditable_:      true,
		DefaultBoxWidth:  200,
		DefaultBoxHeigth: 70,
		Width:            1200,
		Height:           800,
	}).Stage(stage)
	proc.DiagramProcesss = append(proc.DiagramProcesss, diag)

	procShape := (&models.ProcessShape{
		Name:      "MainProcess-MainDiagram",
		Process:   proc,
		RectShape: models.RectShape{X: 300, Y: 50, Width: 600, Height: 600},
	}).Stage(stage)
	diag.Process_Shapes = append(diag.Process_Shapes, procShape)
	stage.Commit()

	// 1. Create Participant and ExternalParticipant via Diagram UX Tree menu buttons
	trees := stager.GetTreeStage().GetInstancesSorted[*tree.Tree]()
	diagNode := findTreeNodeByPath(trees[0].RootNodes, "RootLib", "Processes", "MainProcess", "MainDiagram")
	if diagNode == nil || diagNode.Menu == nil {
		t.Fatalf("expected MainDiagram node with Menu buttons")
	}

	var addParticipantBtn, addExtParticipantBtn *tree.Button
	for _, btn := range diagNode.Menu.Buttons {
		switch btn.Name {
		case "Add Participant":
			addParticipantBtn = btn
		case "Add External Participant":
			addExtParticipantBtn = btn
		}
	}
	if addParticipantBtn == nil || addExtParticipantBtn == nil {
		t.Fatalf("expected both Add Participant and Add External Participant menu buttons, got p=%v ep=%v", addParticipantBtn, addExtParticipantBtn)
	}

	addParticipantBtn.OnClick()
	trees = stager.GetTreeStage().GetInstancesSorted[*tree.Tree]()
	diagNode = findTreeNodeByPath(trees[0].RootNodes, "RootLib", "Processes", "MainProcess", "MainDiagram")
	for _, btn := range diagNode.Menu.Buttons {
		if btn.Name == "Add External Participant" {
			addExtParticipantBtn = btn
		}
	}
	addExtParticipantBtn.OnClick()

	if len(proc.Participants) != 1 || len(diag.Participant_Shapes) != 1 {
		t.Fatalf("expected 1 internal Participant and 1 ParticipantShape, got %d / %d", len(proc.Participants), len(diag.Participant_Shapes))
	}
	if len(proc.ExternalParticipants) != 1 || len(diag.ExternalParticipant_Shapes) != 1 {
		t.Fatalf("expected 1 ExternalParticipant and 1 ExternalParticipantShape, got %d / %d", len(proc.ExternalParticipants), len(diag.ExternalParticipant_Shapes))
	}

	internalParticipant := proc.Participants[0]
	internalParticipant.Name = "Sales Lane"
	extParticipant := proc.ExternalParticipants[0]
	extParticipant.Name = "Customer"
	extShape := diag.ExternalParticipant_Shapes[0]
	extShape.Name = "Customer-MainDiagram"
	extShape.SetX(40)
	extShape.SetY(120)
	extShape.SetWidth(180)
	extShape.SetHeight(65)
	extShape.TailHeigth = 400

	taskReceive := (&models.Task{Name: "Receive Order"}).Stage(stage)
	taskReply := (&models.Task{Name: "Send Confirmation"}).Stage(stage)
	internalParticipant.Tasks = append(internalParticipant.Tasks, taskReceive, taskReply)

	tsReceive := (&models.TaskShape{
		Name:      "Receive Order-MainDiagram",
		Task:      taskReceive,
		RectShape: models.RectShape{X: 350, Y: 200, Width: 200, Height: 65},
	}).Stage(stage)
	tsReply := (&models.TaskShape{
		Name:      "Send Confirmation-MainDiagram",
		Task:      taskReply,
		RectShape: models.RectShape{X: 350, Y: 340, Width: 200, Height: 65},
	}).Stage(stage)
	diag.Task_Shapes = append(diag.Task_Shapes, tsReceive, tsReply)
	stage.Commit()

	if extParticipant.GetOwningProcess() != proc {
		t.Fatalf("expected ExternalParticipant.GetOwningProcess() == proc")
	}

	// 2. Interactive SVG link creation: ExternalParticipant -> Task AND Task -> ExternalParticipant
	svgObj := stager.GetSvgObject()
	var tailRect, receiveRect, replyRect *svg.Rect
	for _, r := range svgObj.Layers[0].Rects {
		if strings.HasPrefix(r.Name, "TailCustomer") {
			tailRect = r
		}
		if r.Name == "Receive Order" {
			receiveRect = r
		}
		if r.Name == "Send Confirmation" {
			replyRect = r
		}
	}
	if tailRect == nil || receiveRect == nil || replyRect == nil {
		t.Fatalf("expected tailRect, receiveRect, and replyRect in SVG layer, got tail=%v recv=%v reply=%v", tailRect, receiveRect, replyRect)
	}

	// Simulate ExternalParticipant -> Task link drag completion
	frontSVG := &svg.SVG{
		DrawingState: svg.NOT_DRAWING_LINK,
		StartRect:    tailRect,
		EndRect:      receiveRect,
	}
	svgObj.DrawingState = svg.DRAWING_LINK
	svgObj.OnUpdate(frontSVG)

	// Re-fetch rects from regenerated SVG and simulate Task -> ExternalParticipant link drag completion
	svgObj = stager.GetSvgObject()
	tailRect, replyRect = nil, nil
	for _, r := range svgObj.Layers[0].Rects {
		if strings.HasPrefix(r.Name, "TailCustomer") {
			tailRect = r
		}
		if r.Name == "Send Confirmation" {
			replyRect = r
		}
	}
	frontSVG2 := &svg.SVG{
		DrawingState: svg.NOT_DRAWING_LINK,
		StartRect:    replyRect,
		EndRect:      tailRect,
	}
	svgObj.DrawingState = svg.DRAWING_LINK
	svgObj.OnUpdate(frontSVG2)

	dataFlows := stage.GetInstancesSorted[*models.DataFlow]()
	if len(dataFlows) != 2 {
		t.Fatalf("expected 2 DataFlows (Ext->Task and Task->Ext) to remain staged after semantic checks, got %d", len(dataFlows))
	}
	var foundExtToTask, foundTaskToExt bool
	for _, df := range dataFlows {
		if df.Type == models.DataFlow_ExternalParticipant2Task && df.StartParticipant == extParticipant && df.EndTask == taskReceive {
			foundExtToTask = true
		}
		if df.Type == models.DataFlow_Task2ExternalParticipant && df.StartTask == taskReply && df.EndParticipant == extParticipant {
			foundTaskToExt = true
		}
	}
	if !foundExtToTask || !foundTaskToExt {
		t.Fatalf("expected both DataFlow_ExternalParticipant2Task and DataFlow_Task2ExternalParticipant, got ext2task=%v task2ext=%v", foundExtToTask, foundTaskToExt)
	}
	if len(stager.GetSvgObject().Layers[0].Links) != 2 {
		t.Fatalf("expected 2 rendered DataFlow svg.Links, got %d", len(stager.GetSvgObject().Layers[0].Links))
	}

	// 3. Orphan ExternalParticipant and orphan shapes are unstaged by semantic rules
	_ = (&models.ExternalParticipant{Name: "OrphanExt"}).Stage(stage)
	_ = (&models.ExternalParticipantShape{Name: "OrphanExtShape", ExternalParticipant: extParticipant}).Stage(stage)
	stage.Commit()

	if len(stage.GetInstancesSorted[*models.ExternalParticipant]()) != 1 {
		t.Fatalf("expected orphan ExternalParticipant to be unstaged, got %d instances", len(stage.GetInstancesSorted[*models.ExternalParticipant]()))
	}
	if len(stage.GetInstancesSorted[*models.ExternalParticipantShape]()) != 1 {
		t.Fatalf("expected orphan ExternalParticipantShape to be unstaged, got %d instances", len(stage.GetInstancesSorted[*models.ExternalParticipantShape]()))
	}
}

func TestIssue1541_ExternalParticipantProcessAndResourceAllocation(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("test_ext_alloc", "", "", true, true, false)
	stager := models.NewStager(stack.R, stack.Stage, stack.Probe)
	stage := stack.Stage

	rootLib := stager.GetRootLibrary()
	rootLib.Name = "RootLib"

	mainProc := (&models.Process{Name: "Order Fulfillment"}).Stage(stage)
	extBackingProc := (&models.Process{
		Name: "Carrier Dispatch Process",
		SVG:  `<svg viewBox="0 0 24 24"><path d="M3 3h18v18H3z"/></svg>`,
	}).Stage(stage)
	extBackingProc2 := (&models.Process{Name: " Customs Clearance Process"}).Stage(stage)
	rootLib.RootProcesses = append(rootLib.RootProcesses, mainProc, extBackingProc, extBackingProc2)

	extResource := (&models.Resource{
		Name: "EDI Gateway",
		SVG:  `<svg viewBox="0 0 24 24"><path d="M4 4h16v16H4z"/></svg>`,
	}).Stage(stage)
	rootLib.RootResources = append(rootLib.RootResources, extResource)

	extParticipant := (&models.ExternalParticipant{
		Name: "Logistics Partner",
	}).Stage(stage)
	mainProc.ExternalParticipants = append(mainProc.ExternalParticipants, extParticipant)

	activeDiag := (&models.DiagramProcess{
		Name:             "ActiveView",
		IsChecked:        true,
		IsEditable_:      true,
		DefaultBoxWidth:  220,
		DefaultBoxHeigth: 70,
		Width:            1200,
		Height:           800,
	}).Stage(stage)
	inactiveDiag := (&models.DiagramProcess{
		Name:             "InactiveView",
		IsChecked:        false,
		IsEditable_:      true,
		DefaultBoxWidth:  220,
		DefaultBoxHeigth: 70,
		Width:            1200,
		Height:           800,
	}).Stage(stage)
	mainProc.DiagramProcesss = append(mainProc.DiagramProcesss, activeDiag, inactiveDiag)

	pShape1 := (&models.ProcessShape{
		Name:      "Order Fulfillment-ActiveView",
		Process:   mainProc,
		RectShape: models.RectShape{X: 300, Y: 50, Width: 600, Height: 600},
	}).Stage(stage)
	activeDiag.Process_Shapes = append(activeDiag.Process_Shapes, pShape1)

	pShape2 := (&models.ProcessShape{
		Name:      "Order Fulfillment-InactiveView",
		Process:   mainProc,
		RectShape: models.RectShape{X: 300, Y: 50, Width: 600, Height: 600},
	}).Stage(stage)
	inactiveDiag.Process_Shapes = append(inactiveDiag.Process_Shapes, pShape2)

	epShape1 := (&models.ExternalParticipantShape{
		Name:                "Logistics Partner-ActiveView",
		ExternalParticipant: extParticipant,
		RectShape:           models.RectShape{X: 40, Y: 120, Width: 200, Height: 65},
		TailHeigth:          350,
	}).Stage(stage)
	activeDiag.ExternalParticipant_Shapes = append(activeDiag.ExternalParticipant_Shapes, epShape1)

	epShape2 := (&models.ExternalParticipantShape{
		Name:                "Logistics Partner-InactiveView",
		ExternalParticipant: extParticipant,
		RectShape:           models.RectShape{X: 40, Y: 120, Width: 200, Height: 65},
		TailHeigth:          350,
	}).Stage(stage)
	inactiveDiag.ExternalParticipant_Shapes = append(inactiveDiag.ExternalParticipant_Shapes, epShape2)
	stage.Commit()

	// 1. Allocate Process and Resource to ExternalParticipant via ExternalParticipant Probe Form
	stack.Probe.FillUpFormFromGongstruct(extParticipant, "ExternalParticipant")
	formStage := stack.Probe.GetFormStage()
	fg := formStage.GetInstancesSorted[*form.FormGroup]()[0]

	// Verify Tasks/ControlFlows are NOT exposed on ExternalParticipant Form
	for _, fd := range fg.FormDivs {
		if fd.Name == "Tasks" || fd.Name == "ControlFlows" {
			t.Fatalf("ExternalParticipant form should not contain %q FormDiv", fd.Name)
		}
	}

	setFormAssocSelection(t, stage, fg, "Processes", []*models.Process{extBackingProc})
	setFormAssocSelection(t, stage, fg, "Resources", []*models.Resource{extResource})
	formStage.Commit()
	fg.OnSave.OnSave()

	if !slices.Contains(extParticipant.Processes, extBackingProc) || !slices.Contains(extParticipant.Resources, extResource) {
		t.Fatalf("expected ExternalParticipant to reference extBackingProc and extResource")
	}
	if len(activeDiag.AllocatedProcessShapes) != 1 || activeDiag.AllocatedProcessShapes[0].ExternalParticipant != extParticipant || activeDiag.AllocatedProcessShapes[0].Process != extBackingProc {
		t.Fatalf("expected activeDiag.AllocatedProcessShapes to auto-add extBackingProc, got %+v", activeDiag.AllocatedProcessShapes)
	}
	if len(activeDiag.AllocatedResourceShapes) != 1 || activeDiag.AllocatedResourceShapes[0].ExternalParticipant != extParticipant || activeDiag.AllocatedResourceShapes[0].Resource != extResource {
		t.Fatalf("expected activeDiag.AllocatedResourceShapes to auto-add extResource, got %+v", activeDiag.AllocatedResourceShapes)
	}
	if len(inactiveDiag.AllocatedProcessShapes) != 0 || len(inactiveDiag.AllocatedResourceShapes) != 0 {
		t.Fatalf("expected inactive diagram not to auto-receive allocated shapes, got proc=%d res=%d",
			len(inactiveDiag.AllocatedProcessShapes), len(inactiveDiag.AllocatedResourceShapes))
	}

	// 2. Verify SVG renders the ExternalParticipant box with both allocated Process and Resource texts + icons
	var extRect *svg.Rect
	for _, r := range stager.GetSvgObject().Layers[0].Rects {
		if r.Name == "Logistics Partner-ActiveView" {
			extRect = r
			break
		}
	}
	if extRect == nil {
		t.Fatalf("expected Logistics Partner-ActiveView rect in SVG")
	}
	if len(extRect.RectAnchoredTexts) != 3 { // title + 1 allocated process + 1 allocated resource
		t.Fatalf("expected 3 RectAnchoredTexts (title + process + resource), got %d", len(extRect.RectAnchoredTexts))
	}
	if len(extRect.RectAnchoredRects) != 1 || extRect.RectAnchoredRects[0].StrokeOpacity != 1.0 {
		t.Fatalf("expected 1 allocated box RectAnchoredRect with StrokeOpacity=1.0, got %+v", extRect.RectAnchoredRects)
	}

	// Test IsProcessResource = true mode on ExternalParticipant
	extParticipant.IsProcessResource = true
	stage.Commit()
	extRect = nil
	for _, r := range stager.GetSvgObject().Layers[0].Rects {
		if r.Name == "Logistics Partner-ActiveView" {
			extRect = r
			break
		}
	}
	if extRect == nil || len(extRect.RectAnchoredTexts) != 2 { // no title, only allocated process + allocated resource
		t.Fatalf("expected 2 RectAnchoredTexts when IsProcessResource=true, got %+v", extRect)
	}
	extParticipant.IsProcessResource = false
	stage.Commit()

	// 3. Test reverse association save on Process Form (ExternalParticipant:Processes)
	stack.Probe.FillUpFormFromGongstruct(extBackingProc2, "Process")
	fgProc := formStage.GetInstancesSorted[*form.FormGroup]()[0]
	setFormAssocSelection(t, stage, fgProc, "ExternalParticipant:Processes", []*models.ExternalParticipant{extParticipant})
	formStage.Commit()
	fgProc.OnSave.OnSave()

	if len(activeDiag.AllocatedProcessShapes) != 2 {
		t.Fatalf("expected 2 AllocatedProcessShapes after reverse Process form save, got %d", len(activeDiag.AllocatedProcessShapes))
	}

	// 4. Verify UX Tree nodes (Library -> Resources -> EDI Gateway -> Logistics Partner, and Diagram -> External Participants -> Allocated Processes)
	trees := stager.GetTreeStage().GetInstancesSorted[*tree.Tree]()
	resExtNode := findTreeNodeByPath(trees[0].RootNodes, "RootLib", "Resources", "EDI Gateway", "Logistics Partner")
	if resExtNode == nil {
		t.Fatalf("expected Logistics Partner under Library -> Resources -> EDI Gateway")
	}
	allocProcNode := findTreeNodeByPath(trees[0].RootNodes, "RootLib", "Processes", "Order Fulfillment", "ActiveView", "External Participants", "Logistics Partner", "Allocated Processes", "Carrier Dispatch Process")
	if allocProcNode == nil || !allocProcNode.IsChecked {
		t.Fatalf("expected checked Carrier Dispatch Process node under External Participant's Allocated Processes")
	}

	// Uncheck and recheck in diagram tree
	allocProcNode.OnIsCheckedChanged(false)
	if len(activeDiag.AllocatedProcessShapes) != 1 {
		t.Fatalf("expected 1 AllocatedProcessShape after unchecking in tree, got %d", len(activeDiag.AllocatedProcessShapes))
	}
	trees = stager.GetTreeStage().GetInstancesSorted[*tree.Tree]()
	allocProcNode = findTreeNodeByPath(trees[0].RootNodes, "RootLib", "Processes", "Order Fulfillment", "ActiveView", "External Participants", "Logistics Partner", "Allocated Processes", "Carrier Dispatch Process")
	allocProcNode.OnIsCheckedChanged(true)
	if len(activeDiag.AllocatedProcessShapes) != 2 {
		t.Fatalf("expected 2 AllocatedProcessShapes after rechecking in tree, got %d", len(activeDiag.AllocatedProcessShapes))
	}

	// 5. Removing Processes and Resources from ExternalParticipant cleans up shapes without orphans
	stack.Probe.FillUpFormFromGongstruct(extParticipant, "ExternalParticipant")
	fg = formStage.GetInstancesSorted[*form.FormGroup]()[0]
	setFormAssocSelection(t, stage, fg, "Processes", []*models.Process{})
	setFormAssocSelection(t, stage, fg, "Resources", []*models.Resource{})
	formStage.Commit()
	fg.OnSave.OnSave()

	if len(stage.GetInstancesSorted[*models.AllocatedProcessShape]()) != 0 || len(stage.GetInstancesSorted[*models.AllocatedResourceShape]()) != 0 {
		t.Fatalf("expected 0 staged AllocatedProcessShape and AllocatedResourceShape after unlinking, got proc=%d res=%d",
			len(stage.GetInstancesSorted[*models.AllocatedProcessShape]()), len(stage.GetInstancesSorted[*models.AllocatedResourceShape]()))
	}

	// Orphan AllocatedProcessShape / AllocatedResourceShape not attached to any diagram are unstaged
	extParticipant.Processes = append(extParticipant.Processes, extBackingProc)
	extParticipant.Resources = append(extParticipant.Resources, extResource)
	_ = (&models.AllocatedProcessShape{Name: "orphan-proc", ExternalParticipant: extParticipant, Process: extBackingProc}).Stage(stage)
	_ = (&models.AllocatedResourceShape{Name: "orphan-res", ExternalParticipant: extParticipant, Resource: extResource}).Stage(stage)
	stage.Commit()

	if len(stage.GetInstancesSorted[*models.AllocatedProcessShape]()) != 0 || len(stage.GetInstancesSorted[*models.AllocatedResourceShape]()) != 0 {
		t.Fatalf("expected orphan AllocatedProcessShape and AllocatedResourceShape to be unstaged")
	}
}

func TestIssue1541_AllDataFilesAndEmbeddedDiagramsLoadCleanly(t *testing.T) {
	files, err := filepath.Glob("data/*.go")
	if err != nil || len(files) == 0 {
		t.Fatalf("expected data/*.go files, err=%v", err)
	}

	// 1. Parse and unmarshall every data/*.go file + merge/merged.go into a fresh Stage
	allStageFiles := append(slices.Clone(files), "merge/merged.go")
	for _, f := range allStageFiles {
		stage := models.NewStage("parse_" + filepath.Base(f))
		if err := stage.ParseAstFile(f, true); err != nil {
			t.Fatalf("ParseAstFile(%q) failed: %v", f, err)
		}
	}

	// 2. Run full Level1Stack + NewStager (including semantic & SVG loops and embedded diagrams) on the external-participant datasets
	models.DataFS = &dataFS
	for _, f := range []string{"data/stage.go", "data/order_to_cash.go", "data/icao4444.go", "data/simple.go"} {
		withEmbeddedDiagrams := filepath.Base(f) == "stage.go"
		stack := level1stack.NewLevel1StackDelta("test_"+filepath.Base(f), f, "", true, withEmbeddedDiagrams, true)
		models.NewStager(stack.R, stack.Stage, stack.Probe)
	}
}
