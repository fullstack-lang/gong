package main

import (
	"math"
	"path/filepath"
	"strings"
	"testing"

	"github.com/fullstack-lang/gong/dsm/structure/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/structure/go/models"
	form "github.com/fullstack-lang/gong/lib/form/go/models"
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func TestDataFlowOrthogonalOptionAndDataPlacement(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("test_dataflow_orthogonal", "", "", true, true, false)
	stager := models.NewStager(
		stack.R,
		stack.Stage,
		stack.Probe,
	)
	stage := stack.Stage
	rootLib := stager.GetRootLibrary()

	sys := (&models.System{Name: "Sys"}).Stage(stage)
	rootLib.RootSystemes = append(rootLib.RootSystemes, sys)

	partA := (&models.Part{Name: "PartA"}).Stage(stage)
	partB := (&models.Part{Name: "PartB"}).Stage(stage)
	sys.Parts = append(sys.Parts, partA, partB)

	// portRightA on right border of PartA; portLeftB on left border of PartB; portTopB on top border of PartB
	portRightA := (&models.Port{Name: "OutRight"}).Stage(stage)
	portLeftB := (&models.Port{Name: "InLeft"}).Stage(stage)
	portTopB := (&models.Port{Name: "InTop"}).Stage(stage)
	partA.Ports = append(partA.Ports, portRightA)
	partB.Ports = append(partB.Ports, portLeftB, portTopB)

	payload := (&models.Data{Name: "TrackData"}).Stage(stage)
	rootLib.RootDatas = append(rootLib.RootDatas, payload)

	proto := (&models.Protocol{Name: "ASTERIX"}).Stage(stage)
	rootLib.RootProtocols = append(rootLib.RootProtocols, proto)

	dfH := (&models.DataFlow{
		Name:      "FlowHoriz",
		Type:      models.DataFlow_Port2Port,
		StartPort: portRightA,
		EndPort:   portLeftB,
		Datas:     []*models.Data{payload},
		Protocol:  proto,
		Direction: models.Forward,
	}).Stage(stage)

	dfL := (&models.DataFlow{
		Name:      "FlowCorner",
		Type:      models.DataFlow_Port2Port,
		StartPort: portRightA,
		EndPort:   portTopB,
		Direction: models.Forward,
	}).Stage(stage)

	rootLib.RootDataFlows = append(rootLib.RootDataFlows, dfH, dfL)
	sys.DataFlows = append(sys.DataFlows, dfH, dfL)

	diag := (&models.DiagramStructure{
		Name:        "MainDiagram",
		IsChecked:   true,
		IsEditable_: true,
		Width:       1200,
		Height:      900,
	}).Stage(stage)
	sys.DiagramStructures = append(sys.DiagramStructures, diag)

	sysShape := (&models.SystemShape{
		Name:      "Sys-MainDiagram",
		System:    sys,
		RectShape: models.RectShape{X: 50, Y: 50, Width: 900, Height: 700},
	}).Stage(stage)
	diag.System_Shapes = append(diag.System_Shapes, sysShape)

	partAShape := (&models.PartShape{
		Name:      "PartA-MainDiagram",
		Part:      partA,
		RectShape: models.RectShape{X: 100, Y: 200, Width: 180, Height: 120},
	}).Stage(stage)
	partBShape := (&models.PartShape{
		Name:      "PartB-MainDiagram",
		Part:      partB,
		RectShape: models.RectShape{X: 620, Y: 340, Width: 180, Height: 120},
	}).Stage(stage)
	diag.Part_Shapes = append(diag.Part_Shapes, partAShape, partBShape)

	// OutRight at right edge of PartA: X=270..290 (center 280), Y=250..270 (center 260)
	portRightAShape := (&models.PortShape{
		Name:      "OutRight-MainDiagram",
		Port:      portRightA,
		RectShape: models.RectShape{X: 270, Y: 250, Width: 20, Height: 20},
	}).Stage(stage)
	// InLeft at left edge of PartB: X=610..630 (center 620), Y=390..410 (center 400)
	portLeftBShape := (&models.PortShape{
		Name:      "InLeft-MainDiagram",
		Port:      portLeftB,
		RectShape: models.RectShape{X: 610, Y: 390, Width: 20, Height: 20},
	}).Stage(stage)
	// InTop at top edge of PartB: X=700..720 (center 710), Y=330..350 (center 340)
	portTopBShape := (&models.PortShape{
		Name:      "InTop-MainDiagram",
		Port:      portTopB,
		RectShape: models.RectShape{X: 700, Y: 330, Width: 20, Height: 20},
	}).Stage(stage)
	diag.Port_Shapes = append(diag.Port_Shapes, portRightAShape, portLeftBShape, portTopBShape)

	dfHShape := (&models.DataFlowShape{
		Name:           "FlowHoriz",
		DataFlow:       dfH,
		StartPortShape: portRightAShape,
		EndPortShape:   portLeftBShape,
	}).Stage(stage)
	dfLShape := (&models.DataFlowShape{
		Name:           "FlowCorner",
		DataFlow:       dfL,
		StartPortShape: portRightAShape,
		EndPortShape:   portTopBShape,
		IsOrthogonal:   true,
	}).Stage(stage)
	diag.DataFlow_Shapes = append(diag.DataFlow_Shapes, dfHShape, dfLShape)

	dataShape := (&models.DataShape{
		Name:          "TrackData",
		Data:          payload,
		DataFlow:      dfH,
		DataFlowShape: dfHShape,
	}).Stage(stage)
	diag.Data_Shapes = append(diag.Data_Shapes, dataShape)

	protoShape := (&models.ProtocolShape{
		Name:          "ASTERIX",
		Protocol:      proto,
		DataFlow:      dfH,
		DataFlowShape: dfHShape,
	}).Stage(stage)
	diag.Protocol_Shapes = append(diag.Protocol_Shapes, protoShape)

	stage.Commit()

	// 1. Verify FlowHoriz starts in control-point mode and FlowCorner starts in orthogonal L-mode (H -> V)
	svgObj := stager.GetSvgObject()
	horizLink := findSvgLinkByName(svgObj, "OutRight to InLeft")
	cornerLink := findSvgLinkByName(svgObj, "OutRight to InTop")
	if horizLink == nil || cornerLink == nil {
		t.Fatalf("expected both SVG links, got horiz=%v corner=%v", horizLink, cornerLink)
	}
	if horizLink.Type != svg.LINK_TYPE_LINE_WITH_CONTROL_POINTS {
		t.Fatalf("expected FlowHoriz default link type LINK_TYPE_LINE_WITH_CONTROL_POINTS, got %s", horizLink.Type)
	}
	if cornerLink.Type != svg.LINK_TYPE_FLOATING_ORTHOGONAL {
		t.Fatalf("expected FlowCorner link type LINK_TYPE_FLOATING_ORTHOGONAL, got %s", cornerLink.Type)
	}
	if dfLShape.StartOrientation != models.ORIENTATION_HORIZONTAL || dfLShape.EndOrientation != models.ORIENTATION_VERTICAL {
		t.Fatalf("expected FlowCorner port-border orientations H -> V, got %s -> %s",
			dfLShape.StartOrientation, dfLShape.EndOrientation)
	}

	// 2. Clicking FlowHoriz node in the diagram tree opens DataFlowShape appearance form with IsOrthogonal checkbox
	trees := stager.GetTreeStage().GetInstancesSorted[*tree.Tree]()
	diagNode := findTreeNodeByName(trees[0].RootNodes, "MainDiagram")
	flowHorizNode := findTreeNodeByName(diagNode.Children, "FlowHoriz")
	if flowHorizNode == nil || flowHorizNode.OnClick == nil {
		t.Fatalf("expected clickable FlowHoriz tree node under MainDiagram")
	}
	flowHorizNode.OnClick(flowHorizNode)

	shapeGroups := stager.GetShapeFormStage().GetInstancesSorted[*form.FormGroup]()
	if len(shapeGroups) != 1 || shapeGroups[0].TypeLabel != "DataFlowShape" {
		t.Fatalf("expected shape form to display DataFlowShape, got %+v", shapeGroups)
	}
	var isOrthCheckbox *form.CheckBox
	for _, fd := range shapeGroups[0].FormDivs {
		if fd.Name == "IsOrthogonal" && len(fd.CheckBoxs) > 0 {
			isOrthCheckbox = fd.CheckBoxs[0]
			break
		}
	}
	if isOrthCheckbox == nil {
		t.Fatalf("expected IsOrthogonal checkbox in DataFlowShape appearance form")
	}
	if isOrthCheckbox.Value {
		t.Fatalf("expected IsOrthogonal checkbox initially false")
	}

	// 3. Enable IsOrthogonal via the Appearance form Save callback
	isOrthCheckbox.Value = true
	stager.GetShapeFormStage().Commit()
	shapeGroups[0].OnSave.OnSave()

	if !dfHShape.IsOrthogonal {
		t.Fatalf("expected dfHShape.IsOrthogonal to be true after saving form")
	}
	if dfHShape.StartOrientation != models.ORIENTATION_HORIZONTAL || dfHShape.EndOrientation != models.ORIENTATION_HORIZONTAL {
		t.Fatalf("expected FlowHoriz orientations H -> H, got %s -> %s",
			dfHShape.StartOrientation, dfHShape.EndOrientation)
	}
	// Initial middle segment X should be centered at (280 + 620) / 2 = 450 -> TrackData.X == 450 + 4 = 454
	svgObj = stager.GetSvgObject()
	horizLink = findSvgLinkByName(svgObj, "OutRight to InLeft")
	if horizLink == nil || horizLink.Type != svg.LINK_TYPE_FLOATING_ORTHOGONAL {
		t.Fatalf("expected FlowHoriz SVG link type LINK_TYPE_FLOATING_ORTHOGONAL after toggle, got %v", horizLink)
	}
	trackRect := findSvgRectByName(svgObj, "TrackData")
	protoRect := findSvgRectByName(svgObj, "ASTERIX")
	if trackRect == nil || protoRect == nil {
		t.Fatalf("expected TrackData and ASTERIX rects on orthogonal FlowHoriz")
	}
	if math.Abs(trackRect.X-454.0) > 1.0 {
		t.Fatalf("expected initial orthogonal TrackData.X near 454.0, got %.1f", trackRect.X)
	}

	// 4. Drag the vertical middle perpendicular segment from X=450 to X=340 -> Data & Protocol placement follows X=340!
	targetMiddleX := 340.0
	frontDrag := *horizLink
	frontDrag.CornerOffsetRatio = (targetMiddleX - horizLink.Start.X) / horizLink.Start.Width
	horizLink.OnAfterUpdate(nil, horizLink, &frontDrag)

	svgObj = stager.GetSvgObject()
	trackRectMoved := findSvgRectByName(svgObj, "TrackData")
	protoRectMoved := findSvgRectByName(svgObj, "ASTERIX")
	if trackRectMoved == nil || protoRectMoved == nil {
		t.Fatalf("expected TrackData and ASTERIX rects after dragging middle perpendicular segment")
	}
	if math.Abs(trackRectMoved.X-(targetMiddleX+4.0)) > 1.0 {
		t.Fatalf("expected TrackData.X to follow dragged perpendicular segment to %.1f, got %.1f",
			targetMiddleX+4.0, trackRectMoved.X)
	}
	if math.Abs(protoRectMoved.X-(targetMiddleX+4.0)) > 1.0 {
		t.Fatalf("expected ASTERIX.X to follow dragged perpendicular segment to %.1f, got %.1f",
			targetMiddleX+4.0, protoRectMoved.X)
	}

	// Static SVG export renders the moved vertical segment at X=340 without errors
	staticSvg, _, _ := svgObj.GenerateString()
	if !strings.Contains(staticSvg, "340") {
		t.Fatalf("expected static SVG export to contain orthogonal segment at X=340")
	}

	// 5. Clicking the orthogonal link without dragging selects DataFlow + DataFlowShape without creating ControlPointShapes
	horizLink = findSvgLinkByName(svgObj, "OutRight to InLeft")
	frontClick := *horizLink
	horizLink.OnAfterUpdate(nil, horizLink, &frontClick)
	if len(dfHShape.ControlPointShapes) != 0 {
		t.Fatalf("expected 0 ControlPointShapes when clicking orthogonal link, got %d", len(dfHShape.ControlPointShapes))
	}

	// 6. AST Stage persistence round-trip preserves IsOrthogonal and orthogonal LinkShape fields
	tmpStageFile := filepath.Join(t.TempDir(), "orthogonal_stage.go")
	stage.MarshallFile(tmpStageFile, "github.com/fullstack-lang/gong/dsm/structure/go/models", "main")

	reloadedStack := level1stack.NewLevel1StackDelta("test_dataflow_orth_reload", "", "", true, true, false)
	if err := reloadedStack.Stage.ParseAstFile(tmpStageFile, false); err != nil {
		t.Fatalf("ParseAstFile failed: %v", err)
	}
	var reloadedHoriz *models.DataFlowShape
	for _, dfs := range reloadedStack.Stage.GetInstancesSorted[*models.DataFlowShape]() {
		if dfs.Name == "FlowHoriz" {
			reloadedHoriz = dfs
			break
		}
	}
	if reloadedHoriz == nil || !reloadedHoriz.IsOrthogonal {
		t.Fatalf("expected reloaded FlowHoriz DataFlowShape to have IsOrthogonal == true, got %+v", reloadedHoriz)
	}
	if math.Abs(reloadedHoriz.CornerOffsetRatio-frontDrag.CornerOffsetRatio) > 1e-6 {
		t.Fatalf("expected reloaded CornerOffsetRatio %.4f, got %.4f",
			frontDrag.CornerOffsetRatio, reloadedHoriz.CornerOffsetRatio)
	}
}
