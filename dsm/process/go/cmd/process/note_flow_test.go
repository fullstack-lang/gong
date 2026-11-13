package main

import (
	"slices"
	"testing"

	"github.com/fullstack-lang/gong/dsm/process/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/process/go/models"
	"github.com/fullstack-lang/gong/dsm/process/go/models/probe"
	form "github.com/fullstack-lang/gong/lib/form/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
	gongprobe "github.com/fullstack-lang/gong/pkg/runtime/probe"
)

func setFormAssocSelection[T models.PointerToGongstruct](
	t *testing.T,
	stage *models.Stage,
	formGroup *form.FormGroup,
	divName string,
	selected []T,
) {
	t.Helper()
	mapIDRowID := probe.GetMap_ID_RowID[T](stage)
	rowIDs := make([]uint, 0, len(selected))
	for _, item := range selected {
		id := stage.GetOrder(item)
		rowID, ok := mapIDRowID[id]
		if !ok {
			t.Fatalf("instance %v not found in rowID map", item)
		}
		rowIDs = append(rowIDs, uint(rowID))
	}
	encoded, err := gongprobe.EncodeIntSliceToString(rowIDs)
	if err != nil {
		t.Fatalf("EncodeIntSliceToString failed: %v", err)
	}
	for _, fd := range formGroup.FormDivs {
		if fd.Name == divName {
			if fd.FormEditAssocButton == nil {
				t.Fatalf("FormDiv %q has nil FormEditAssocButton", divName)
			}
			fd.FormEditAssocButton.AssociationStorage = encoded
			fd.FormEditAssocButton.HasChanged = true
			return
		}
	}
	t.Fatalf("FormDiv %q not found on FormGroup %q", divName, formGroup.Label)
}

func findTreeNodeByPath(rootNodes []*tree.Node, names ...string) *tree.Node {
	current := rootNodes
	var matched *tree.Node
	for _, name := range names {
		matched = nil
		for _, n := range current {
			if n.Name == name {
				matched = n
				break
			}
		}
		if matched == nil {
			return nil
		}
		current = matched.Children
	}
	return matched
}

func TestNoteLinksToTaskDataFlowAndControlFlow(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("test_note_flows", "", "", true, true, false)
	stager := models.NewStager(
		stack.R,
		stack.Stage,
		stack.Probe,
	)
	stage := stack.Stage
	rootLib := stager.GetRootLibrary()
	rootLib.Name = "RootLib"

	proc := (&models.Process{Name: "Proc"}).Stage(stage)
	rootLib.RootProcesses = append(rootLib.RootProcesses, proc)

	p1 := (&models.Participant{Name: "Lane 1"}).Stage(stage)
	p2 := (&models.Participant{Name: "Lane 2"}).Stage(stage)
	proc.Participants = append(proc.Participants, p1, p2)

	task1A := (&models.Task{Name: "Task 1A"}).Stage(stage)
	task1B := (&models.Task{Name: "Task 1B"}).Stage(stage)
	targetTask := (&models.Task{Name: "Task 2"}).Stage(stage)
	p1.Tasks = append(p1.Tasks, task1A, task1B)
	p2.Tasks = append(p2.Tasks, targetTask)

	targetControlFlow := (&models.ControlFlow{
		Name:  "CF 1A to 1B",
		Start: task1A,
		End:   task1B,
	}).Stage(stage)
	p1.ControlFlows = append(p1.ControlFlows, targetControlFlow)

	// Intentionally leave Datas empty to verify DataFlow links without DataShapes still map to svg.Link
	targetDataFlow := (&models.DataFlow{
		Name:      "DF 1B to 2",
		Type:      models.DataFlow_Task2Task,
		StartTask: task1B,
		EndTask:   targetTask,
	}).Stage(stage)
	rootLib.RootDataFlows = append(rootLib.RootDataFlows, targetDataFlow)
	proc.DataFlows = append(proc.DataFlows, targetDataFlow)

	// Pre-existing Note already linked to task1A before currentDiagram shapes are edited
	note := (&models.Note{
		Name:  "First Note",
		Tasks: []*models.Task{task1A},
	}).Stage(stage)
	rootLib.RootNotes = append(rootLib.RootNotes, note)

	currentDiagram := (&models.DiagramProcess{
		Name:             "ActiveDiagram",
		IsChecked:        true,
		IsEditable_:      true,
		DefaultBoxWidth:  200,
		DefaultBoxHeigth: 80,
		Width:            1200,
		Height:           900,
	}).Stage(stage)
	otherDiagram := (&models.DiagramProcess{
		Name:             "UncheckedDiagram",
		IsChecked:        false,
		IsEditable_:      true,
		DefaultBoxWidth:  200,
		DefaultBoxHeigth: 80,
		Width:            1200,
		Height:           900,
	}).Stage(stage)
	proc.DiagramProcesss = append(proc.DiagramProcesss, currentDiagram, otherDiagram)

	procShape := (&models.ProcessShape{
		Name:      "Proc-ActiveDiagram",
		Process:   proc,
		RectShape: models.RectShape{X: 50, Y: 50, Width: 900, Height: 700},
	}).Stage(stage)
	otherProcShape := (&models.ProcessShape{
		Name:      "Proc-UncheckedDiagram",
		Process:   proc,
		RectShape: models.RectShape{X: 50, Y: 50, Width: 900, Height: 700},
	}).Stage(stage)
	currentDiagram.Process_Shapes = append(currentDiagram.Process_Shapes, procShape)
	otherDiagram.Process_Shapes = append(otherDiagram.Process_Shapes, otherProcShape)

	p1Shape := (&models.ParticipantShape{
		Name:        "Lane 1-ActiveDiagram",
		Participant: p1,
		WidthWeight: 1.0,
	}).Stage(stage)
	p2Shape := (&models.ParticipantShape{
		Name:        "Lane 2-ActiveDiagram",
		Participant: p2,
		WidthWeight: 1.0,
	}).Stage(stage)
	currentDiagram.Participant_Shapes = append(currentDiagram.Participant_Shapes, p1Shape, p2Shape)

	t1AShape := (&models.TaskShape{
		Name:      "Task 1A-ActiveDiagram",
		Task:      task1A,
		RectShape: models.RectShape{X: 120, Y: 200, Width: 200, Height: 80},
	}).Stage(stage)
	t1BShape := (&models.TaskShape{
		Name:      "Task 1B-ActiveDiagram",
		Task:      task1B,
		RectShape: models.RectShape{X: 120, Y: 380, Width: 200, Height: 80},
	}).Stage(stage)
	t2Shape := (&models.TaskShape{
		Name:      "Task 2-ActiveDiagram",
		Task:      targetTask,
		RectShape: models.RectShape{X: 560, Y: 380, Width: 200, Height: 80},
	}).Stage(stage)
	currentDiagram.Task_Shapes = append(currentDiagram.Task_Shapes, t1AShape, t1BShape, t2Shape)

	cfShape := (&models.ControlFlowShape{
		Name:        targetControlFlow.Name,
		ControlFlow: targetControlFlow,
	}).Stage(stage)
	currentDiagram.ControlFlow_Shapes = append(currentDiagram.ControlFlow_Shapes, cfShape)

	dfShape := (&models.DataFlowShape{
		Name:     targetDataFlow.Name,
		DataFlow: targetDataFlow,
	}).Stage(stage)
	currentDiagram.DataFlow_Shapes = append(currentDiagram.DataFlow_Shapes, dfShape)

	noteShape := (&models.NoteShape{
		Name:      "First Note shape",
		Note:      note,
		RectShape: models.RectShape{X: 560, Y: 180, Width: 240, Height: 70},
	}).Stage(stage)
	currentDiagram.Note_Shapes = append(currentDiagram.Note_Shapes, noteShape)
	stage.Commit()

	// 1. Saving an unrelated field on Note must NOT auto-add pre-existing Note.Tasks to currentDiagram
	stack.Probe.FillUpFormFromGongstruct(note, "Note")
	formStage := stack.Probe.GetFormStage()
	fg := formStage.GetInstancesSorted[*form.FormGroup]()[0]
	for _, fd := range fg.FormDivs {
		if fd.Name == "Description" && len(fd.FormFields) > 0 && fd.FormFields[0].FormFieldString != nil {
			fd.FormFields[0].FormFieldString.Value = "updated description only"
		}
	}
	formStage.Commit()
	fg.OnSave.OnSave()

	if len(currentDiagram.NoteTaskShapes) != 0 || len(currentDiagram.NoteDataFlowShapes) != 0 || len(currentDiagram.NoteControlFlowShapes) != 0 {
		t.Fatalf("saving unrelated Note field should not add pre-existing links to currentDiagram, got task=%d df=%d cf=%d",
			len(currentDiagram.NoteTaskShapes), len(currentDiagram.NoteDataFlowShapes), len(currentDiagram.NoteControlFlowShapes))
	}

	// 2. Add links from Note to targetTask, targetDataFlow, and targetControlFlow via Note's Probe Form
	stack.Probe.FillUpFormFromGongstruct(note, "Note")
	fg = formStage.GetInstancesSorted[*form.FormGroup]()[0]
	setFormAssocSelection(t, stage, fg, "Tasks", append(slices.Clone(note.Tasks), targetTask))
	setFormAssocSelection(t, stage, fg, "DataFlows", []*models.DataFlow{targetDataFlow})
	setFormAssocSelection(t, stage, fg, "ControlFlows", []*models.ControlFlow{targetControlFlow})
	formStage.Commit()
	fg.OnSave.OnSave()

	if !slices.Contains(note.DataFlows, targetDataFlow) {
		t.Fatalf("expected Note.DataFlows to contain targetDataFlow")
	}
	if !slices.Contains(note.ControlFlows, targetControlFlow) {
		t.Fatalf("expected Note.ControlFlows to contain targetControlFlow")
	}
	if len(currentDiagram.NoteTaskShapes) != 1 || currentDiagram.NoteTaskShapes[0].Task != targetTask {
		t.Fatalf("expected currentDiagram.NoteTaskShapes to auto-add targetTask, got %+v", currentDiagram.NoteTaskShapes)
	}
	if len(currentDiagram.NoteDataFlowShapes) != 1 || currentDiagram.NoteDataFlowShapes[0].DataFlow != targetDataFlow {
		t.Fatalf("expected currentDiagram.NoteDataFlowShapes to auto-add targetDataFlow, got %+v", currentDiagram.NoteDataFlowShapes)
	}
	if len(currentDiagram.NoteControlFlowShapes) != 1 || currentDiagram.NoteControlFlowShapes[0].ControlFlow != targetControlFlow {
		t.Fatalf("expected currentDiagram.NoteControlFlowShapes to auto-add targetControlFlow, got %+v", currentDiagram.NoteControlFlowShapes)
	}
	if len(otherDiagram.NoteDataFlowShapes) != 0 || len(otherDiagram.NoteControlFlowShapes) != 0 {
		t.Fatalf("expected unchecked diagram not to receive Note flow shapes, got df=%d cf=%d",
			len(otherDiagram.NoteDataFlowShapes), len(otherDiagram.NoteControlFlowShapes))
	}

	// Verify SVG rendering of Note -> ControlFlow and Note -> DataFlow via svg.RectLinkLink (with StrokeOpacity = 1.0)
	svgObj := stager.GetSvgObject()
	if svgObj == nil || len(svgObj.Layers) == 0 {
		t.Fatalf("expected non-empty SVG layers")
	}
	layer := svgObj.Layers[0]
	if len(layer.RectLinkLinks) != 2 {
		t.Fatalf("expected 2 svg.RectLinkLinks in layer (1 ControlFlow + 1 DataFlow), got %d", len(layer.RectLinkLinks))
	}
	for idx, rll := range layer.RectLinkLinks {
		if rll.Start == nil || rll.Start.Name != note.Name {
			t.Errorf("RectLinkLink[%d]: expected Start to be Note rect %q, got %+v", idx, note.Name, rll.Start)
		}
		if rll.End == nil {
			t.Errorf("RectLinkLink[%d]: expected non-nil End svg.Link", idx)
		}
		if rll.StrokeOpacity != 1.0 {
			t.Errorf("RectLinkLink[%d]: expected StrokeOpacity == 1.0, got %v", idx, rll.StrokeOpacity)
		}
		if rll.TargetAnchorPosition != 0.5 {
			t.Errorf("RectLinkLink[%d]: expected TargetAnchorPosition == 0.5, got %v", idx, rll.TargetAnchorPosition)
		}
	}

	// 3. Test Reverse Associations in Probe Form (DataFlow -> Note:DataFlows, ControlFlow -> Note:ControlFlows, Task -> Note:Tasks)
	note2 := (&models.Note{Name: "Second Note"}).Stage(stage)
	rootLib.RootNotes = append(rootLib.RootNotes, note2)
	note2Shape := (&models.NoteShape{
		Name:      "Second Note shape",
		Note:      note2,
		RectShape: models.RectShape{X: 560, Y: 280, Width: 240, Height: 70},
	}).Stage(stage)
	currentDiagram.Note_Shapes = append(currentDiagram.Note_Shapes, note2Shape)
	stage.Commit()

	stack.Probe.FillUpFormFromGongstruct(targetDataFlow, "DataFlow")
	fgDF := formStage.GetInstancesSorted[*form.FormGroup]()[0]
	setFormAssocSelection(t, stage, fgDF, "Note:DataFlows", []*models.Note{note, note2})
	formStage.Commit()
	fgDF.OnSave.OnSave()

	stack.Probe.FillUpFormFromGongstruct(targetControlFlow, "ControlFlow")
	fgCF := formStage.GetInstancesSorted[*form.FormGroup]()[0]
	setFormAssocSelection(t, stage, fgCF, "Note:ControlFlows", []*models.Note{note, note2})
	formStage.Commit()
	fgCF.OnSave.OnSave()

	stack.Probe.FillUpFormFromGongstruct(targetTask, "Task")
	fgTask := formStage.GetInstancesSorted[*form.FormGroup]()[0]
	setFormAssocSelection(t, stage, fgTask, "Note:Tasks", []*models.Note{note, note2})
	formStage.Commit()
	fgTask.OnSave.OnSave()

	if !slices.Contains(note2.DataFlows, targetDataFlow) || !slices.Contains(note2.ControlFlows, targetControlFlow) || !slices.Contains(note2.Tasks, targetTask) {
		t.Fatalf("expected note2 to contain targetDataFlow, targetControlFlow, and targetTask after reverse form saves")
	}
	if len(currentDiagram.NoteDataFlowShapes) != 2 || len(currentDiagram.NoteControlFlowShapes) != 2 || len(currentDiagram.NoteTaskShapes) != 2 {
		t.Fatalf("expected 2 shapes each in currentDiagram after reverse saves, got df=%d cf=%d task=%d",
			len(currentDiagram.NoteDataFlowShapes), len(currentDiagram.NoteControlFlowShapes), len(currentDiagram.NoteTaskShapes))
	}
	if len(stager.GetSvgObject().Layers[0].RectLinkLinks) != 4 {
		t.Fatalf("expected 4 RectLinkLinks in SVG layer, got %d", len(stager.GetSvgObject().Layers[0].RectLinkLinks))
	}

	// 4. Test Presentation Override & Hiding on NoteControlFlowShape and NoteDataFlowShape
	ncfs := currentDiagram.NoteControlFlowShapes[0]
	ncfs.SetIsPresentationOverridden(true)
	ncfs.SetStroke("#0055FF")
	ncfs.SetStrokeWidth(2.5)
	ncfs.SetStrokeOpacity(0.8)
	ncfs.SetStrokeDashArray("4,2")

	ndfs := currentDiagram.NoteDataFlowShapes[0]
	ndfs.SetIsPresentationOverridden(true)
	ndfs.SetStroke("#FF0000")
	ndfs.SetStrokeWidth(3.0)
	ndfs.SetStrokeOpacity(0.75)
	ndfs.SetStrokeDashArray("2,2")
	stage.Commit()

	rllCF := stager.GetSvgObject().Layers[0].RectLinkLinks[0]
	if rllCF.Stroke != "#0055FF" || rllCF.StrokeWidth != 2.5 || rllCF.StrokeOpacity != 0.8 || rllCF.StrokeDashArray != "4,2" {
		t.Errorf("expected overridden ControlFlow RectLinkLink presentation, got stroke=%q width=%v opacity=%v dash=%q",
			rllCF.Stroke, rllCF.StrokeWidth, rllCF.StrokeOpacity, rllCF.StrokeDashArray)
	}
	rllDF := stager.GetSvgObject().Layers[0].RectLinkLinks[2]
	if rllDF.Stroke != "#FF0000" || rllDF.StrokeWidth != 3.0 || rllDF.StrokeOpacity != 0.75 || rllDF.StrokeDashArray != "2,2" {
		t.Errorf("expected overridden DataFlow RectLinkLink presentation, got stroke=%q width=%v opacity=%v dash=%q",
			rllDF.Stroke, rllDF.StrokeWidth, rllDF.StrokeOpacity, rllDF.StrokeDashArray)
	}

	ncfs.SetIsHidden(true)
	ndfs.SetIsHidden(true)
	stage.Commit()
	if len(stager.GetSvgObject().Layers[0].RectLinkLinks) != 2 {
		t.Fatalf("expected 2 visible RectLinkLinks when 2 shapes are hidden, got %d", len(stager.GetSvgObject().Layers[0].RectLinkLinks))
	}
	ncfs.SetIsHidden(false)
	ndfs.SetIsHidden(false)
	stage.Commit()

	// 5. Verify UX Tree nodes (Library tree and Diagram tree checkboxes + Appearance form on click)
	trees := stager.GetTreeStage().GetInstancesSorted[*tree.Tree]()
	if len(trees) == 0 {
		t.Fatalf("expected UX tree to be populated")
	}
	rootNodes := trees[0].RootNodes

	libCFNode := findTreeNodeByPath(rootNodes, "RootLib", "Notes", "First Note", "Control Flows", targetControlFlow.Name)
	libDFNode := findTreeNodeByPath(rootNodes, "RootLib", "Notes", "First Note", "Data Flows", targetDataFlow.Name)
	if libCFNode == nil || libDFNode == nil {
		t.Fatalf("expected Control Flows and Data Flows nodes under Library -> Notes -> First Note, got cf=%v df=%v", libCFNode, libDFNode)
	}

	diagCFNode := findTreeNodeByPath(rootNodes, "RootLib", "Processes", "Proc", "ActiveDiagram", "Notes", "First Note", "Control Flows", targetControlFlow.Name)
	diagDFNode := findTreeNodeByPath(rootNodes, "RootLib", "Processes", "Proc", "ActiveDiagram", "Notes", "First Note", "Data Flows", targetDataFlow.Name)
	if diagCFNode == nil || diagDFNode == nil {
		t.Fatalf("expected Control Flows and Data Flows nodes under ActiveDiagram -> Notes -> First Note, got cf=%v df=%v", diagCFNode, diagDFNode)
	}
	if !diagCFNode.IsChecked || !diagDFNode.IsChecked {
		t.Fatalf("expected diagram flow nodes to be checked")
	}

	diagDFNode.OnClick(diagDFNode)
	shapeFGs := stager.GetShapeFormStage().GetInstancesSorted[*form.FormGroup]()
	if len(shapeFGs) != 1 {
		t.Fatalf("expected clicking diagram DataFlow node under Note to populate shapeFormStage, got %d groups", len(shapeFGs))
	}

	// Uncheck DataFlow under First Note in diagram tree, then re-check
	diagDFNode.OnIsCheckedChanged(false)
	if len(currentDiagram.NoteDataFlowShapes) != 1 {
		t.Fatalf("expected 1 NoteDataFlowShape after unchecking First Note's DataFlow in tree, got %d", len(currentDiagram.NoteDataFlowShapes))
	}
	trees = stager.GetTreeStage().GetInstancesSorted[*tree.Tree]()
	diagDFNode = findTreeNodeByPath(trees[0].RootNodes, "RootLib", "Processes", "Proc", "ActiveDiagram", "Notes", "First Note", "Data Flows", targetDataFlow.Name)
	diagDFNode.OnIsCheckedChanged(true)
	if len(currentDiagram.NoteDataFlowShapes) != 2 {
		t.Fatalf("expected 2 NoteDataFlowShapes after re-checking First Note's DataFlow in tree, got %d", len(currentDiagram.NoteDataFlowShapes))
	}

	// 6. Removing links from Note2 cleans up NoteDataFlowShape and NoteControlFlowShape without leaving orphans
	stack.Probe.FillUpFormFromGongstruct(note2, "Note")
	fgNote2 := formStage.GetInstancesSorted[*form.FormGroup]()[0]
	setFormAssocSelection(t, stage, fgNote2, "DataFlows", []*models.DataFlow{})
	setFormAssocSelection(t, stage, fgNote2, "ControlFlows", []*models.ControlFlow{})
	formStage.Commit()
	fgNote2.OnSave.OnSave()

	if len(currentDiagram.NoteDataFlowShapes) != 1 || len(currentDiagram.NoteControlFlowShapes) != 1 {
		t.Fatalf("expected 1 NoteDataFlowShape and 1 NoteControlFlowShape left after unlinking note2, got df=%d cf=%d",
			len(currentDiagram.NoteDataFlowShapes), len(currentDiagram.NoteControlFlowShapes))
	}
	if len(stage.GetInstancesSorted[*models.NoteDataFlowShape]()) != 1 || len(stage.GetInstancesSorted[*models.NoteControlFlowShape]()) != 1 {
		t.Fatalf("expected 1 staged NoteDataFlowShape and 1 staged NoteControlFlowShape total (no orphans), got df=%d cf=%d",
			len(stage.GetInstancesSorted[*models.NoteDataFlowShape]()), len(stage.GetInstancesSorted[*models.NoteControlFlowShape]()))
	}

	// 7. Orphan NoteDataFlowShape and NoteControlFlowShape not attached to any diagram are unstaged by enforceShapeOrphans
	_ = (&models.NoteDataFlowShape{
		Name:     "orphan-df",
		Note:     note,
		DataFlow: targetDataFlow,
	}).Stage(stage)
	_ = (&models.NoteControlFlowShape{
		Name:        "orphan-cf",
		Note:        note,
		ControlFlow: targetControlFlow,
	}).Stage(stage)
	stage.Commit()

	if len(stage.GetInstancesSorted[*models.NoteDataFlowShape]()) != 1 || len(stage.GetInstancesSorted[*models.NoteControlFlowShape]()) != 1 {
		t.Fatalf("expected orphan Note flow shapes to be unstaged, got df=%d cf=%d",
			len(stage.GetInstancesSorted[*models.NoteDataFlowShape]()), len(stage.GetInstancesSorted[*models.NoteControlFlowShape]()))
	}
}
