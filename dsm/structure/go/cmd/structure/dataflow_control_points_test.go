package main

import (
	"math"
	"path/filepath"
	"strings"
	"testing"

	"github.com/fullstack-lang/gong/dsm/structure/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/structure/go/models"
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func findSvgLinkByName(svgObj *svg.SVG, name string) *svg.Link {
	if svgObj == nil || len(svgObj.Layers) == 0 {
		return nil
	}
	for _, link := range svgObj.Layers[0].Links {
		if link.Name == name {
			return link
		}
	}
	return nil
}

func findSvgRectByName(svgObj *svg.SVG, name string) *svg.Rect {
	if svgObj == nil || len(svgObj.Layers) == 0 {
		return nil
	}
	for _, rect := range svgObj.Layers[0].Rects {
		if rect.Name == name {
			return rect
		}
	}
	return nil
}

func TestDataFlowControlPointsAddMoveDeleteBadgesAndOrphans(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("test_dataflow_control_points", "", "", true, true, false)
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
	extC := (&models.Part{Name: "ExtC"}).Stage(stage)
	sys.Parts = append(sys.Parts, partA, partB)
	sys.ExternalParts = append(sys.ExternalParts, extC)

	// Intentionally give both connected ports on PartA and PartB the exact same Name ("Port")
	// to verify pointer-based ClosestRect resolution instead of fragile string-name matching.
	portA := (&models.Port{Name: "Port"}).Stage(stage)
	portB := (&models.Port{Name: "Port"}).Stage(stage)
	portAExt := (&models.Port{Name: "PortExt"}).Stage(stage)
	partA.Ports = append(partA.Ports, portA, portAExt)
	partB.Ports = append(partB.Ports, portB)

	payload := (&models.Data{Name: "Telemetry"}).Stage(stage)
	rootLib.RootDatas = append(rootLib.RootDatas, payload)

	proto := (&models.Protocol{Name: "UDP"}).Stage(stage)
	rootLib.RootProtocols = append(rootLib.RootProtocols, proto)

	dfPort2Port := (&models.DataFlow{
		Name:      "FlowP2P",
		Type:      models.DataFlow_Port2Port,
		StartPort: portA,
		EndPort:   portB,
		Datas:     []*models.Data{payload},
		Protocol:  proto,
		Direction: models.Forward,
	}).Stage(stage)

	dfPort2Ext := (&models.DataFlow{
		Name:            "FlowP2Ext",
		Type:            models.DataFlow_Port2ExternalPart,
		StartPort:       portAExt,
		EndExternalPart: extC,
		Direction:       models.Forward,
	}).Stage(stage)

	rootLib.RootDataFlows = append(rootLib.RootDataFlows, dfPort2Port, dfPort2Ext)
	sys.DataFlows = append(sys.DataFlows, dfPort2Port, dfPort2Ext)

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
		RectShape: models.RectShape{X: 50, Y: 50, Width: 800, Height: 600},
	}).Stage(stage)
	diag.System_Shapes = append(diag.System_Shapes, sysShape)

	partAShape := (&models.PartShape{
		Name:      "PartA-MainDiagram",
		Part:      partA,
		RectShape: models.RectShape{X: 100, Y: 250, Width: 180, Height: 120},
	}).Stage(stage)
	partBShape := (&models.PartShape{
		Name:      "PartB-MainDiagram",
		Part:      partB,
		RectShape: models.RectShape{X: 620, Y: 250, Width: 180, Height: 120},
	}).Stage(stage)
	diag.Part_Shapes = append(diag.Part_Shapes, partAShape, partBShape)

	extCShape := (&models.ExternalPartShape{
		Name:      "ExtC-MainDiagram",
		Part:      extC,
		RectShape: models.RectShape{X: 920, Y: 120, Width: 140, Height: 60},
	}).Stage(stage)
	diag.ExternalPart_Shapes = append(diag.ExternalPart_Shapes, extCShape)

	// Place portA on right border of PartA (X=270..290, Y=290..310 -> center 280, 300)
	// and portB on left border of PartB (X=610..630, Y=290..310 -> center 620, 300)
	portAShape := (&models.PortShape{
		Name:      "Port-MainDiagram-A",
		Port:      portA,
		RectShape: models.RectShape{X: 270, Y: 290, Width: 20, Height: 20},
	}).Stage(stage)
	portBShape := (&models.PortShape{
		Name:      "Port-MainDiagram-B",
		Port:      portB,
		RectShape: models.RectShape{X: 610, Y: 290, Width: 20, Height: 20},
	}).Stage(stage)
	portAExtShape := (&models.PortShape{
		Name:      "PortExt-MainDiagram",
		Port:      portAExt,
		RectShape: models.RectShape{X: 270, Y: 340, Width: 20, Height: 20},
	}).Stage(stage)
	diag.Port_Shapes = append(diag.Port_Shapes, portAShape, portBShape, portAExtShape)

	dfP2PShape := (&models.DataFlowShape{
		Name:           "FlowP2P",
		DataFlow:       dfPort2Port,
		StartPortShape: portAShape,
		EndPortShape:   portBShape,
	}).Stage(stage)
	dfP2ExtShape := (&models.DataFlowShape{
		Name:           "FlowP2Ext",
		DataFlow:       dfPort2Ext,
		StartPortShape: portAExtShape,
	}).Stage(stage)
	diag.DataFlow_Shapes = append(diag.DataFlow_Shapes, dfP2PShape, dfP2ExtShape)

	dataShape := (&models.DataShape{
		Name:          "Telemetry",
		Data:          payload,
		DataFlow:      dfPort2Port,
		DataFlowShape: dfP2PShape,
	}).Stage(stage)
	diag.Data_Shapes = append(diag.Data_Shapes, dataShape)

	protoShape := (&models.ProtocolShape{
		Name:          "UDP",
		Protocol:      proto,
		DataFlow:      dfPort2Port,
		DataFlowShape: dfP2PShape,
	}).Stage(stage)
	diag.Protocol_Shapes = append(diag.Protocol_Shapes, protoShape)

	stage.Commit()

	// Verify baseline straight-line midpoint for DataShape & ProtocolShape (~Y = 300)
	svgObj := stager.GetSvgObject()
	telemetryRectBefore := findSvgRectByName(svgObj, "Telemetry")
	if telemetryRectBefore == nil {
		t.Fatalf("expected Telemetry DataShape rect in SVG")
	}
	initialCenterY := telemetryRectBefore.Y + telemetryRectBefore.Height/2.0
	if math.Abs(initialCenterY-300.0) > 25.0 {
		t.Fatalf("expected straight-line Telemetry badge near Y=300, got centerY=%.1f", initialCenterY)
	}

	// 1. Add first control point on FlowP2P via Link.OnAfterUpdate at symmetric upper bend (450, 120)
	p2pLink := findSvgLinkByName(svgObj, "Port to Port")
	if p2pLink == nil {
		t.Fatalf("expected SVG link 'Port to Port'")
	}
	if p2pLink.Impl == nil {
		t.Fatalf("expected editable DataFlow SVG link to have non-nil Impl")
	}
	frontClick1 := *p2pLink
	frontClick1.MouseX = 450
	frontClick1.MouseY = 120
	p2pLink.OnAfterUpdate(nil, p2pLink, &frontClick1)

	if len(dfP2PShape.ControlPointShapes) != 1 {
		t.Fatalf("expected 1 ControlPointShape after first click, got %d", len(dfP2PShape.ControlPointShapes))
	}
	cpFirst := dfP2PShape.ControlPointShapes[0]

	svgObj = stager.GetSvgObject()
	p2pLink = findSvgLinkByName(svgObj, "Port to Port")
	if p2pLink == nil || len(p2pLink.ControlPoints) != 1 {
		t.Fatalf("expected 1 svg.ControlPoint on regenerated link, got %v", p2pLink)
	}

	// Verify DataShape and ProtocolShape midpoint badges moved upward toward the bend at Y=120
	telemetryRectBent := findSvgRectByName(svgObj, "Telemetry")
	udpRectBent := findSvgRectByName(svgObj, "UDP")
	if telemetryRectBent == nil || udpRectBent == nil {
		t.Fatalf("expected Telemetry and UDP badge rects on bent DataFlow")
	}
	bentCenterY := telemetryRectBent.Y + telemetryRectBent.Height/2.0
	if math.Abs(bentCenterY-120.0) > 25.0 {
		t.Fatalf("expected Telemetry badge to follow bent polyline midpoint near Y=120, got centerY=%.1f", bentCenterY)
	}
	// Verify static SVG export includes multi-segment path through Y=120
	staticSvg, _, _ := svgObj.GenerateString()
	if !strings.Contains(staticSvg, "120") {
		t.Fatalf("expected static SVG export to contain bent path coordinate near Y=120")
	}

	// 2. Add a second control point near the FIRST segment (between startCenter ~280,300 and 450,120) at (340, 210)
	// and verify closest-segment insertion places it at index 0 (before cpFirst)!
	frontClick2 := *p2pLink
	frontClick2.MouseX = 340
	frontClick2.MouseY = 210
	p2pLink.OnAfterUpdate(nil, p2pLink, &frontClick2)

	if len(dfP2PShape.ControlPointShapes) != 2 {
		t.Fatalf("expected 2 ControlPointShapes after second click, got %d", len(dfP2PShape.ControlPointShapes))
	}
	if dfP2PShape.ControlPointShapes[1] != cpFirst {
		t.Fatalf("expected new control point on segment 0 to be inserted at index 0 before cpFirst")
	}
	if !dfP2PShape.ControlPointShapes[0].IsStartShapeTheClosestShape {
		t.Fatalf("expected control point at X=340 to anchor to startRect")
	}

	// 3. Move control point 0 closer to endRect (X=580, Y=160) even though startRect.Name == endRect.Name == "Port"
	svgObj = stager.GetSvgObject()
	p2pLink = findSvgLinkByName(svgObj, "Port to Port")
	svgCP0 := p2pLink.ControlPoints[0]
	movedCP := models.PointToControlPoint(&svg.Point{X: 580, Y: 160}, p2pLink.Start, p2pLink.End)
	if movedCP.ClosestRect != p2pLink.End {
		t.Fatalf("expected (580,160) to be closer to End rect")
	}
	svgCP0.OnAfterUpdate(nil, svgCP0, &movedCP)

	if dfP2PShape.ControlPointShapes[0].IsStartShapeTheClosestShape {
		t.Fatalf("expected moved control point to re-anchor to End shape (IsStartShapeTheClosestShape == false) despite identical port names")
	}
	reconstructed := models.ControlPointShapeToPoint(dfP2PShape.ControlPointShapes[0], p2pLink.Start, p2pLink.End)
	if math.Hypot(reconstructed.X-580, reconstructed.Y-160) > 0.5 {
		t.Fatalf("expected reconstructed point ~(580,160), got (%.2f, %.2f)", reconstructed.X, reconstructed.Y)
	}

	// 4. Port2ExternalPart with narrow 5px lifeline tailRect: add & click-to-delete with 1.5px jitter
	svgObj = stager.GetSvgObject()
	p2ExtLink := findSvgLinkByName(svgObj, "PortExt to TailExtC-MainDiagram")
	if p2ExtLink == nil {
		t.Fatalf("expected SVG link 'PortExt to TailExtC-MainDiagram'")
	}
	if p2ExtLink.End.Width != 5.0 {
		t.Fatalf("expected ExternalPart tailRect width 5.0, got %.1f", p2ExtLink.End.Width)
	}
	frontExtClick := *p2ExtLink
	frontExtClick.MouseX = 880
	frontExtClick.MouseY = 260
	p2ExtLink.OnAfterUpdate(nil, p2ExtLink, &frontExtClick)

	if len(dfP2ExtShape.ControlPointShapes) != 1 {
		t.Fatalf("expected 1 ControlPointShape on FlowP2Ext, got %d", len(dfP2ExtShape.ControlPointShapes))
	}
	extCPShape := dfP2ExtShape.ControlPointShapes[0]
	if extCPShape.IsStartShapeTheClosestShape {
		t.Fatalf("expected control point at X=880 to anchor to ExternalPart tailRect")
	}

	// Simulate clicking the handle with 1.5px mouse jitter (relative delta 1.5/5 = 0.30 > 0.05, but < 5.0px!)
	svgObj = stager.GetSvgObject()
	p2ExtLink = findSvgLinkByName(svgObj, "PortExt to TailExtC-MainDiagram")
	svgExtCP := p2ExtLink.ControlPoints[0]
	jitteredClick := models.PointToControlPoint(&svg.Point{X: 881.5, Y: 260.5}, p2ExtLink.Start, p2ExtLink.End)
	svgExtCP.OnAfterUpdate(nil, svgExtCP, &jitteredClick)

	if len(dfP2ExtShape.ControlPointShapes) != 0 {
		t.Fatalf("expected click-to-delete (< 5px) on 5px ExternalPart lifeline to delete control point, still have %d",
			len(dfP2ExtShape.ControlPointShapes))
	}
	if stage.IsStaged(extCPShape) {
		t.Fatalf("expected deleted ControlPointShape to be unstaged")
	}

	// 5. Tree "Diagram Editability" button toggles edit / non-edit mode:
	//    When non-editable, svgObj.IsEditable is false (hiding blue control-point dots in Angular)
	//    and Impl callbacks are nil while bends still render.
	findEditabilityButton := func() *tree.Button {
		t.Helper()
		trees := stager.GetTreeStage().GetInstancesSorted[*tree.Tree]()
		if len(trees) == 0 {
			t.Fatalf("expected tree in treeStage")
		}
		diagNode := findTreeNodeByName(trees[0].RootNodes, "MainDiagram")
		if diagNode == nil {
			t.Fatalf("expected MainDiagram node in tree")
		}
		for _, btn := range diagNode.Buttons {
			if btn.Name == "Diagram Editability" {
				return btn
			}
		}
		t.Fatalf("expected 'Diagram Editability' button on MainDiagram tree node")
		return nil
	}

	findEditabilityButton().OnClick()
	if diag.IsEditable_ {
		t.Fatalf("expected Diagram Editability button click to turn off IsEditable_")
	}
	svgObj = stager.GetSvgObject()
	if svgObj.IsEditable {
		t.Fatalf("expected svgObj.IsEditable to be false in non-edit mode")
	}
	p2pLink = findSvgLinkByName(svgObj, "Port to Port")
	if p2pLink == nil || len(p2pLink.ControlPoints) != 2 {
		t.Fatalf("expected 2 rendered control points when diagram is non-editable")
	}
	if p2pLink.Impl != nil || p2pLink.ControlPoints[0].Impl != nil {
		t.Fatalf("expected nil Impl callbacks when diagram is non-editable")
	}

	findEditabilityButton().OnClick()
	if !diag.IsEditable_ || !stager.GetSvgObject().IsEditable {
		t.Fatalf("expected second Diagram Editability button click to re-enable edit mode")
	}

	// 6. AST Stage persistence round-trip (MarshallFile -> ParseAstFile)
	tmpStageFile := filepath.Join(t.TempDir(), "stage_roundtrip.go")
	stage.MarshallFile(tmpStageFile, "github.com/fullstack-lang/gong/dsm/structure/go/models", "main")

	reloadedStack := level1stack.NewLevel1StackDelta("test_dataflow_cp_reload", "", "", true, true, false)
	if err := reloadedStack.Stage.ParseAstFile(tmpStageFile, false); err != nil {
		t.Fatalf("ParseAstFile failed: %v", err)
	}
	reloadedCPs := reloadedStack.Stage.GetInstancesSorted[*models.ControlPointShape]()
	if len(reloadedCPs) != 2 {
		t.Fatalf("expected 2 persisted ControlPointShapes after AST reload, got %d", len(reloadedCPs))
	}

	// 7. Orphan enforcement: stray ControlPointShape & removed DataFlowShape clean up all ControlPointShapes
	strayCP := (&models.ControlPointShape{Name: "StrayOrphan", X_Relative: 1.0, Y_Relative: 1.0}).Stage(stage)
	stage.Commit()
	if stage.IsStaged(strayCP) {
		t.Fatalf("expected enforceShapeOrphans to unstage stray ControlPointShape")
	}
	if got := len(stage.GetInstancesSorted[*models.ControlPointShape]()); got != 2 {
		t.Fatalf("expected 2 reachable ControlPointShapes remaining, got %d", got)
	}

	// Remove dfP2PShape from diagram (simulating uncheck / deletion) -> all ControlPointShapes must be unstaged
	dfP2PShape.Unstage(stage)
	diag.DataFlow_Shapes = []*models.DataFlowShape{dfP2ExtShape}
	stage.Commit()

	if got := len(stage.GetInstancesSorted[*models.ControlPointShape]()); got != 0 {
		t.Fatalf("expected 0 staged ControlPointShapes after DataFlowShape removal, got %d", got)
	}
}
