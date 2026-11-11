package main

import (
	"slices"
	"testing"

	"github.com/fullstack-lang/gong/dsm/structure/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/structure/go/models"
	"github.com/fullstack-lang/gong/dsm/structure/go/models/probe"
	form "github.com/fullstack-lang/gong/lib/form/go/models"
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

func TestNoteLinksToPortPartAndDataFlowAutoAddConcreteToCurrentDiagram(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("test_note_dataflow", "", "", true, true, false)
	stager := models.NewStager(
		stack.R,
		stack.Stage,
		stack.Probe,
	)
	stage := stack.Stage
	rootLib := stager.GetRootLibrary()

	sys := (&models.System{Name: "Sys"}).Stage(stage)
	rootLib.RootSystemes = append(rootLib.RootSystemes, sys)

	part1 := (&models.Part{Name: "P1"}).Stage(stage)
	targetPart := (&models.Part{Name: "P2"}).Stage(stage)
	sys.Parts = append(sys.Parts, part1, targetPart)

	port1 := (&models.Port{Name: "Port 1"}).Stage(stage)
	targetPort := (&models.Port{Name: "Port 2"}).Stage(stage)
	part1.Ports = append(part1.Ports, port1)
	targetPart.Ports = append(targetPart.Ports, targetPort)

	targetDataFlow := (&models.DataFlow{
		Name:      "Flow 1",
		Type:      models.DataFlow_Port2Port,
		StartPort: port1,
		EndPort:   targetPort,
	}).Stage(stage)
	rootLib.RootDataFlows = append(rootLib.RootDataFlows, targetDataFlow)
	sys.DataFlows = append(sys.DataFlows, targetDataFlow)

	// Pre-existing Note already linked to part1 and port1 before currentDiagram shapes are edited
	note := (&models.Note{
		Name:  "First Note",
		Parts: []*models.Part{part1},
		Ports: []*models.Port{port1},
	}).Stage(stage)
	rootLib.RootNotes = append(rootLib.RootNotes, note)

	currentDiagram := (&models.DiagramStructure{
		Name:        "ActiveDiagram",
		IsChecked:   true,
		IsEditable_: true,
		Width:       1000,
		Height:      800,
	}).Stage(stage)
	otherDiagram := (&models.DiagramStructure{
		Name:        "UncheckedDiagram",
		IsChecked:   false,
		IsEditable_: true,
		Width:       1000,
		Height:      800,
	}).Stage(stage)
	sys.DiagramStructures = append(sys.DiagramStructures, currentDiagram, otherDiagram)

	sysShape := (&models.SystemShape{
		Name:      "Sys-ActiveDiagram",
		System:    sys,
		RectShape: models.RectShape{X: 50, Y: 50, Width: 800, Height: 600},
	}).Stage(stage)
	otherSysShape := (&models.SystemShape{
		Name:      "Sys-UncheckedDiagram",
		System:    sys,
		RectShape: models.RectShape{X: 50, Y: 50, Width: 800, Height: 600},
	}).Stage(stage)
	currentDiagram.System_Shapes = append(currentDiagram.System_Shapes, sysShape)
	otherDiagram.System_Shapes = append(otherDiagram.System_Shapes, otherSysShape)

	p1Shape := (&models.PartShape{
		Name:      "P1-ActiveDiagram",
		Part:      part1,
		RectShape: models.RectShape{X: 100, Y: 250, Width: 200, Height: 100},
	}).Stage(stage)
	p2Shape := (&models.PartShape{
		Name:      "P2-ActiveDiagram",
		Part:      targetPart,
		RectShape: models.RectShape{X: 450, Y: 250, Width: 200, Height: 100},
	}).Stage(stage)
	currentDiagram.Part_Shapes = append(currentDiagram.Part_Shapes, p1Shape, p2Shape)

	port1Shape := (&models.PortShape{
		Name:      "Port 1-ActiveDiagram",
		Port:      port1,
		RectShape: models.RectShape{X: 280, Y: 280, Width: 40, Height: 40},
	}).Stage(stage)
	port2Shape := (&models.PortShape{
		Name:      "Port 2-ActiveDiagram",
		Port:      targetPort,
		RectShape: models.RectShape{X: 430, Y: 280, Width: 40, Height: 40},
	}).Stage(stage)
	currentDiagram.Port_Shapes = append(currentDiagram.Port_Shapes, port1Shape, port2Shape)

	dfShape := (&models.DataFlowShape{
		Name:     "Flow 1",
		DataFlow: targetDataFlow,
	}).Stage(stage)
	currentDiagram.DataFlow_Shapes = append(currentDiagram.DataFlow_Shapes, dfShape)

	noteShape := (&models.NoteShape{
		Name:      "First Note shape",
		Note:      note,
		RectShape: models.RectShape{X: 120, Y: 100, Width: 250, Height: 70},
	}).Stage(stage)
	currentDiagram.Note_Shapes = append(currentDiagram.Note_Shapes, noteShape)
	stage.Commit()

	// 1. Saving an unrelated field on Note must NOT auto-add pre-existing Note.Parts / Note.Ports to currentDiagram
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

	if len(currentDiagram.NotePortShapes) != 0 || len(currentDiagram.NotePartShapes) != 0 || len(currentDiagram.NoteDataFlowShapes) != 0 {
		t.Fatalf("saving unrelated Note field should not add pre-existing links to currentDiagram, got port=%d part=%d df=%d",
			len(currentDiagram.NotePortShapes), len(currentDiagram.NotePartShapes), len(currentDiagram.NoteDataFlowShapes))
	}

	// 2. Add links from Note to targetPart, targetPort, and targetDataFlow via Note's Probe Form (forward slices)
	stack.Probe.FillUpFormFromGongstruct(note, "Note")
	fg = formStage.GetInstancesSorted[*form.FormGroup]()[0]
	setFormAssocSelection(t, stage, fg, "Parts", append(slices.Clone(note.Parts), targetPart))
	setFormAssocSelection(t, stage, fg, "Ports", append(slices.Clone(note.Ports), targetPort))
	setFormAssocSelection(t, stage, fg, "DataFlows", []*models.DataFlow{targetDataFlow})
	formStage.Commit()
	fg.OnSave.OnSave()

	if !slices.Contains(note.DataFlows, targetDataFlow) {
		t.Fatalf("expected Note.DataFlows to contain targetDataFlow")
	}
	if len(currentDiagram.NotePartShapes) != 1 || currentDiagram.NotePartShapes[0].Part != targetPart {
		t.Fatalf("expected currentDiagram.NotePartShapes to auto-add targetPart, got %+v", currentDiagram.NotePartShapes)
	}
	if len(currentDiagram.NotePortShapes) != 1 || currentDiagram.NotePortShapes[0].Port != targetPort {
		t.Fatalf("expected currentDiagram.NotePortShapes to auto-add targetPort, got %+v", currentDiagram.NotePortShapes)
	}
	if len(currentDiagram.NoteDataFlowShapes) != 1 || currentDiagram.NoteDataFlowShapes[0].DataFlow != targetDataFlow {
		t.Fatalf("expected currentDiagram.NoteDataFlowShapes to auto-add targetDataFlow, got %+v", currentDiagram.NoteDataFlowShapes)
	}
	if len(otherDiagram.NoteDataFlowShapes) != 0 {
		t.Fatalf("expected unchecked diagram not to receive NoteDataFlowShapes, got %d", len(otherDiagram.NoteDataFlowShapes))
	}

	// Verify SVG rendering of Note -> DataFlow via svg.RectLinkLink (including StrokeOpacity = 1.0!)
	svgObj := stager.GetSvgObject()
	if svgObj == nil || len(svgObj.Layers) == 0 {
		t.Fatalf("expected non-empty SVG layers")
	}
	layer := svgObj.Layers[0]
	if len(layer.RectLinkLinks) != 1 {
		t.Fatalf("expected 1 svg.RectLinkLink in layer, got %d", len(layer.RectLinkLinks))
	}
	rll := layer.RectLinkLinks[0]
	if rll.Start == nil || rll.Start.Name != note.Name {
		t.Errorf("expected RectLinkLink.Start to be Note rect %q, got %+v", note.Name, rll.Start)
	}
	if rll.End == nil {
		t.Errorf("expected RectLinkLink.End to point to DataFlow svg.Link")
	}
	if rll.StrokeOpacity != 1.0 {
		t.Errorf("expected RectLinkLink.StrokeOpacity == 1.0, got %v", rll.StrokeOpacity)
	}
	if rll.TargetAnchorPosition != 0.5 {
		t.Errorf("expected RectLinkLink.TargetAnchorPosition == 0.5, got %v", rll.TargetAnchorPosition)
	}

	// 3. Test Reverse Association in Probe Form (editing DataFlow -> Note:DataFlows)
	note2 := (&models.Note{Name: "Second Note"}).Stage(stage)
	rootLib.RootNotes = append(rootLib.RootNotes, note2)
	note2Shape := (&models.NoteShape{
		Name:      "Second Note shape",
		Note:      note2,
		RectShape: models.RectShape{X: 400, Y: 100, Width: 250, Height: 70},
	}).Stage(stage)
	currentDiagram.Note_Shapes = append(currentDiagram.Note_Shapes, note2Shape)
	stage.Commit()

	stack.Probe.FillUpFormFromGongstruct(targetDataFlow, "DataFlow")
	fgDF := formStage.GetInstancesSorted[*form.FormGroup]()[0]
	setFormAssocSelection(t, stage, fgDF, "Note:DataFlows", []*models.Note{note, note2})
	formStage.Commit()
	fgDF.OnSave.OnSave()

	if !slices.Contains(note2.DataFlows, targetDataFlow) {
		t.Fatalf("expected note2.DataFlows to contain targetDataFlow after reverse form save")
	}
	if len(currentDiagram.NoteDataFlowShapes) != 2 {
		t.Fatalf("expected 2 NoteDataFlowShapes in currentDiagram after reverse form save, got %d", len(currentDiagram.NoteDataFlowShapes))
	}
	if len(stager.GetSvgObject().Layers[0].RectLinkLinks) != 2 {
		t.Fatalf("expected 2 RectLinkLinks in SVG layer, got %d", len(stager.GetSvgObject().Layers[0].RectLinkLinks))
	}

	// 4. Test Presentation Override & Hiding on NoteDataFlowShape
	ndfs := currentDiagram.NoteDataFlowShapes[0]
	ndfs.SetIsPresentationOverridden(true)
	ndfs.SetStroke("#FF0000")
	ndfs.SetStrokeWidth(3.0)
	ndfs.SetStrokeOpacity(0.75)
	ndfs.SetStrokeDashArray("2,2")
	stage.Commit()

	rllOverridden := stager.GetSvgObject().Layers[0].RectLinkLinks[0]
	if rllOverridden.Stroke != "#FF0000" || rllOverridden.StrokeWidth != 3.0 || rllOverridden.StrokeOpacity != 0.75 || rllOverridden.StrokeDashArray != "2,2" {
		t.Errorf("expected overridden RectLinkLink presentation, got stroke=%q width=%v opacity=%v dash=%q",
			rllOverridden.Stroke, rllOverridden.StrokeWidth, rllOverridden.StrokeOpacity, rllOverridden.StrokeDashArray)
	}

	ndfs.SetIsHidden(true)
	stage.Commit()
	if len(stager.GetSvgObject().Layers[0].RectLinkLinks) != 1 {
		t.Fatalf("expected 1 visible RectLinkLink when first NoteDataFlowShape is hidden, got %d", len(stager.GetSvgObject().Layers[0].RectLinkLinks))
	}
	ndfs.SetIsHidden(false)
	stage.Commit()

	// 5. Removing link from Note2.DataFlows cleans up NoteDataFlowShape without leaving orphans
	stack.Probe.FillUpFormFromGongstruct(note2, "Note")
	fgNote2 := formStage.GetInstancesSorted[*form.FormGroup]()[0]
	setFormAssocSelection(t, stage, fgNote2, "DataFlows", []*models.DataFlow{})
	formStage.Commit()
	fgNote2.OnSave.OnSave()

	if len(currentDiagram.NoteDataFlowShapes) != 1 {
		t.Fatalf("expected 1 NoteDataFlowShape left in currentDiagram after unlinking note2, got %d", len(currentDiagram.NoteDataFlowShapes))
	}
	if len(stage.GetInstancesSorted[*models.NoteDataFlowShape]()) != 1 {
		t.Fatalf("expected 1 staged NoteDataFlowShape total (no orphans), got %d", len(stage.GetInstancesSorted[*models.NoteDataFlowShape]()))
	}

	// 6. Reverse association save on Port (Note:Ports) and Part (Note:Parts) also auto-adds concrete shapes
	stack.Probe.FillUpFormFromGongstruct(targetPort, "Port")
	fgPort := formStage.GetInstancesSorted[*form.FormGroup]()[0]
	setFormAssocSelection(t, stage, fgPort, "Note:Ports", []*models.Note{note, note2})
	formStage.Commit()
	fgPort.OnSave.OnSave()
	if len(currentDiagram.NotePortShapes) != 2 {
		t.Fatalf("expected 2 NotePortShapes after reverse Port form save, got %d", len(currentDiagram.NotePortShapes))
	}

	stack.Probe.FillUpFormFromGongstruct(targetPart, "Part")
	fgPart := formStage.GetInstancesSorted[*form.FormGroup]()[0]
	setFormAssocSelection(t, stage, fgPart, "Note:Parts", []*models.Note{note, note2})
	formStage.Commit()
	fgPart.OnSave.OnSave()
	if len(currentDiagram.NotePartShapes) != 2 {
		t.Fatalf("expected 2 NotePartShapes after reverse Part form save, got %d", len(currentDiagram.NotePartShapes))
	}

	// 7. Orphan NoteDataFlowShape not attached to any diagram is automatically unstaged by enforceShapeOrphans
	orphanShape := (&models.NoteDataFlowShape{
		Name:     "orphan",
		Note:     note,
		DataFlow: targetDataFlow,
	}).Stage(stage)
	_ = orphanShape
	stage.Commit()
	if len(stage.GetInstancesSorted[*models.NoteDataFlowShape]()) != 1 {
		t.Fatalf("expected orphan NoteDataFlowShape to be unstaged, got %d staged instances", len(stage.GetInstancesSorted[*models.NoteDataFlowShape]()))
	}
}
