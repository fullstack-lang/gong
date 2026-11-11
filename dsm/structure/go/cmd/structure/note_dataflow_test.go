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
	stack := level1stack.NewLevel1StackDelta("test_note_dataflow", "data/stage.go", "", true, true, false)
	stager := models.NewStager(
		stack.R,
		stack.Stage,
		stack.Probe,
	)

	// Locate current checked diagram ("sdqsdqsd") and unchecked diagrams
	var currentDiagram *models.DiagramStructure
	var uncheckedDiagrams []*models.DiagramStructure
	for _, d := range stack.Stage.GetInstancesSorted[*models.DiagramStructure]() {
		if d.IsChecked {
			currentDiagram = d
		} else {
			uncheckedDiagrams = append(uncheckedDiagrams, d)
		}
	}
	if currentDiagram == nil {
		t.Fatalf("expected a checked DiagramStructure in data/stage.go")
	}
	if len(currentDiagram.NotePortShapes) != 0 || len(currentDiagram.NotePartShapes) != 0 || len(currentDiagram.NoteDataFlowShapes) != 0 {
		t.Fatalf("expected checked diagram %q to start with 0 Note link shapes, got port=%d part=%d df=%d",
			currentDiagram.Name, len(currentDiagram.NotePortShapes), len(currentDiagram.NotePartShapes), len(currentDiagram.NoteDataFlowShapes))
	}

	notes := stack.Stage.GetInstancesSorted[*models.Note]()
	if len(notes) == 0 {
		t.Fatalf("expected at least 1 Note in data/stage.go")
	}
	note := notes[0]

	// Place NoteShape on currentDiagram so SVG renders Note and its links
	noteShape := (&models.NoteShape{
		Name:   note.Name + " shape",
		Note:   note,
		RectShape: models.RectShape{
			X:      120,
			Y:      120,
			Width:  250,
			Height: 70,
		},
	}).Stage(stack.Stage)
	currentDiagram.Note_Shapes = append(currentDiagram.Note_Shapes, noteShape)
	stack.Stage.Commit()

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

	// Find target Part ("P2"), Port ("Port 2" or any port on currentDiagram not yet in note.Ports), and DataFlow on currentDiagram
	var targetPart *models.Part
	for _, ps := range currentDiagram.Part_Shapes {
		if ps.Part != nil && !slices.Contains(note.Parts, ps.Part) {
			targetPart = ps.Part
			break
		}
	}
	var targetPort *models.Port
	for _, ps := range currentDiagram.Port_Shapes {
		if ps.Port != nil && !slices.Contains(note.Ports, ps.Port) {
			targetPort = ps.Port
			break
		}
	}
	if len(currentDiagram.DataFlow_Shapes) == 0 || currentDiagram.DataFlow_Shapes[0].DataFlow == nil {
		t.Fatalf("expected at least 1 DataFlowShape on currentDiagram")
	}
	targetDataFlow := currentDiagram.DataFlow_Shapes[0].DataFlow
	if targetPart == nil || targetPort == nil {
		t.Fatalf("expected unlinked Part and Port on currentDiagram, got part=%v port=%v", targetPart, targetPort)
	}

	// 2. Add links from Note to targetPart, targetPort, and targetDataFlow via Note's Probe Form (forward slices)
	stack.Probe.FillUpFormFromGongstruct(note, "Note")
	fg = formStage.GetInstancesSorted[*form.FormGroup]()[0]
	setFormAssocSelection(t, stack.Stage, fg, "Parts", append(slices.Clone(note.Parts), targetPart))
	setFormAssocSelection(t, stack.Stage, fg, "Ports", append(slices.Clone(note.Ports), targetPort))
	setFormAssocSelection(t, stack.Stage, fg, "DataFlows", []*models.DataFlow{targetDataFlow})
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
	for _, ud := range uncheckedDiagrams {
		if len(ud.NoteDataFlowShapes) != 0 {
			t.Fatalf("expected unchecked diagram %q not to receive NoteDataFlowShapes, got %d", ud.Name, len(ud.NoteDataFlowShapes))
		}
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
	rootLib := stager.GetRootLibrary()
	note2 := (&models.Note{Name: "Second Note"}).Stage(stack.Stage)
	rootLib.RootNotes = append(rootLib.RootNotes, note2)
	note2Shape := (&models.NoteShape{
		Name: "Second Note shape",
		Note: note2,
		RectShape: models.RectShape{
			X:      300,
			Y:      120,
			Width:  250,
			Height: 70,
		},
	}).Stage(stack.Stage)
	currentDiagram.Note_Shapes = append(currentDiagram.Note_Shapes, note2Shape)
	stack.Stage.Commit()

	stack.Probe.FillUpFormFromGongstruct(targetDataFlow, "DataFlow")
	fgDF := formStage.GetInstancesSorted[*form.FormGroup]()[0]
	setFormAssocSelection(t, stack.Stage, fgDF, "Note:DataFlows", []*models.Note{note, note2})
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
	stack.Stage.Commit()

	rllOverridden := stager.GetSvgObject().Layers[0].RectLinkLinks[0]
	if rllOverridden.Stroke != "#FF0000" || rllOverridden.StrokeWidth != 3.0 || rllOverridden.StrokeOpacity != 0.75 || rllOverridden.StrokeDashArray != "2,2" {
		t.Errorf("expected overridden RectLinkLink presentation, got stroke=%q width=%v opacity=%v dash=%q",
			rllOverridden.Stroke, rllOverridden.StrokeWidth, rllOverridden.StrokeOpacity, rllOverridden.StrokeDashArray)
	}

	ndfs.SetIsHidden(true)
	stack.Stage.Commit()
	if len(stager.GetSvgObject().Layers[0].RectLinkLinks) != 1 {
		t.Fatalf("expected 1 visible RectLinkLink when first NoteDataFlowShape is hidden, got %d", len(stager.GetSvgObject().Layers[0].RectLinkLinks))
	}
	ndfs.SetIsHidden(false)
	stack.Stage.Commit()

	// 5. Removing link from Note2.DataFlows cleans up NoteDataFlowShape without leaving orphans
	stack.Probe.FillUpFormFromGongstruct(note2, "Note")
	fgNote2 := formStage.GetInstancesSorted[*form.FormGroup]()[0]
	setFormAssocSelection(t, stack.Stage, fgNote2, "DataFlows", []*models.DataFlow{})
	formStage.Commit()
	fgNote2.OnSave.OnSave()

	if len(currentDiagram.NoteDataFlowShapes) != 1 {
		t.Fatalf("expected 1 NoteDataFlowShape left in currentDiagram after unlinking note2, got %d", len(currentDiagram.NoteDataFlowShapes))
	}
	if len(stack.Stage.GetInstancesSorted[*models.NoteDataFlowShape]()) != 1 {
		t.Fatalf("expected 1 staged NoteDataFlowShape total (no orphans), got %d", len(stack.Stage.GetInstancesSorted[*models.NoteDataFlowShape]()))
	}

	// 6. Reverse association save on Port (Note:Ports) and Part (Note:Parts) also auto-adds concrete shapes
	stack.Probe.FillUpFormFromGongstruct(targetPort, "Port")
	fgPort := formStage.GetInstancesSorted[*form.FormGroup]()[0]
	setFormAssocSelection(t, stack.Stage, fgPort, "Note:Ports", []*models.Note{note, note2})
	formStage.Commit()
	fgPort.OnSave.OnSave()
	if len(currentDiagram.NotePortShapes) != 2 {
		t.Fatalf("expected 2 NotePortShapes after reverse Port form save, got %d", len(currentDiagram.NotePortShapes))
	}

	stack.Probe.FillUpFormFromGongstruct(targetPart, "Part")
	fgPart := formStage.GetInstancesSorted[*form.FormGroup]()[0]
	setFormAssocSelection(t, stack.Stage, fgPart, "Note:Parts", []*models.Note{note, note2})
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
	}).Stage(stack.Stage)
	_ = orphanShape
	stack.Stage.Commit()
	if len(stack.Stage.GetInstancesSorted[*models.NoteDataFlowShape]()) != 1 {
		t.Fatalf("expected orphan NoteDataFlowShape to be unstaged, got %d staged instances", len(stack.Stage.GetInstancesSorted[*models.NoteDataFlowShape]()))
	}
}
