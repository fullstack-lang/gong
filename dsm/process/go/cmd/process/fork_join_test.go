package main

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/fullstack-lang/gong/dsm/process/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/process/go/models"
	form "github.com/fullstack-lang/gong/lib/form/go/models"
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func setFormBoolField(t *testing.T, fg *form.FormGroup, fieldName string, val bool) {
	t.Helper()
	for _, fd := range fg.FormDivs {
		if fd.Name == fieldName && len(fd.CheckBoxs) > 0 {
			fd.CheckBoxs[0].Value = val
			return
		}
	}
	t.Fatalf("boolean FormDiv %q not found on FormGroup %q", fieldName, fg.Label)
}

func setFormStringField(t *testing.T, fg *form.FormGroup, fieldName string, val string) {
	t.Helper()
	for _, fd := range fg.FormDivs {
		if fd.Name == fieldName && len(fd.FormFields) > 0 && fd.FormFields[0].FormFieldString != nil {
			fd.FormFields[0].FormFieldString.Value = val
			return
		}
	}
	t.Fatalf("string FormDiv %q not found on FormGroup %q", fieldName, fg.Label)
}

func hasFormDiv(fg *form.FormGroup, fieldName string) bool {
	for _, fd := range fg.FormDivs {
		if fd.Name == fieldName {
			return true
		}
	}
	return false
}

func findSvgRectByName(layer *svg.Layer, name string) *svg.Rect {
	for _, r := range layer.Rects {
		if r.Name == name {
			return r
		}
	}
	return nil
}

func findMenuButton(node *tree.Node, name string) *tree.Button {
	if node == nil || node.Menu == nil {
		return nil
	}
	for _, b := range node.Menu.Buttons {
		if b.Name == name {
			return b
		}
	}
	return nil
}

func TestForkJoinEndToEnd(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("test_fork_join", "", "", true, true, false)
	stager := models.NewStager(
		stack.R,
		stack.Stage,
		stack.Probe,
	)
	stage := stack.Stage
	rootLib := stager.GetRootLibrary()
	rootLib.Name = "RootLib"
	rootLib.IsProcessesNodeExpanded = true

	proc := (&models.Process{Name: "OrderFulfillment"}).Stage(stage)
	subProc := (&models.Process{Name: "SubProcessType"}).Stage(stage)
	rootLib.RootProcesses = append(rootLib.RootProcesses, proc, subProc)
	rootLib.ProcesssWhoseNodeIsExpanded = append(rootLib.ProcesssWhoseNodeIsExpanded, proc)

	lane1 := (&models.Participant{Name: "Warehouse"}).Stage(stage)
	lane2 := (&models.Participant{Name: "Billing"}).Stage(stage)
	extCarrier := (&models.ExternalParticipant{Name: "Carrier"}).Stage(stage)
	proc.Participants = append(proc.Participants, lane1, lane2)
	proc.ExternalParticipants = append(proc.ExternalParticipants, extCarrier)

	taskLeft := (&models.Task{Name: "Pick Dry Goods"}).Stage(stage)
	taskRight := (&models.Task{Name: "Pack Cold Chain"}).Stage(stage)
	lane1.Tasks = append(lane1.Tasks, taskLeft, taskRight)

	taskBilling := (&models.Task{Name: "Confirm Payment"}).Stage(stage)
	lane2.Tasks = append(lane2.Tasks, taskBilling)

	diag := (&models.DiagramProcess{
		Name:                               "MainDiagram",
		IsChecked:                          true,
		IsEditable_:                        true,
		DefaultBoxWidth:                    180,
		DefaultBoxHeigth:                   60,
		Width:                              1400,
		Height:                             1000,
		IsParticipantsNodeExpanded:         true,
		IsExternalParticipantsNodeExpanded: true,
	}).Stage(stage)
	proc.DiagramProcesss = append(proc.DiagramProcesss, diag)
	proc.DiagramProcessWhoseNodeIsExpanded = append(proc.DiagramProcessWhoseNodeIsExpanded, diag)

	procShape := (&models.ProcessShape{
		Name:      "OrderFulfillment-MainDiagram",
		Process:   proc,
		RectShape: models.RectShape{X: 50, Y: 50, Width: 900, Height: 800},
	}).Stage(stage)
	diag.Process_Shapes = append(diag.Process_Shapes, procShape)

	lane1Shape := (&models.ParticipantShape{
		Name:        "Warehouse-MainDiagram",
		Participant: lane1,
		WidthWeight: 1.5,
	}).Stage(stage)
	lane2Shape := (&models.ParticipantShape{
		Name:        "Billing-MainDiagram",
		Participant: lane2,
		WidthWeight: 1.0,
	}).Stage(stage)
	diag.Participant_Shapes = append(diag.Participant_Shapes, lane1Shape, lane2Shape)

	extCarrierShape := (&models.ExternalParticipantShape{
		Name:                "Carrier-MainDiagram",
		ExternalParticipant: extCarrier,
		TailHeigth:          300,
		RectShape:           models.RectShape{X: 1050, Y: 100, Width: 160, Height: 60},
	}).Stage(stage)
	diag.ExternalParticipant_Shapes = append(diag.ExternalParticipant_Shapes, extCarrierShape)

	leftShape := (&models.TaskShape{
		Name:      "Pick Dry Goods-MainDiagram",
		Task:      taskLeft,
		RectShape: models.RectShape{X: 80, Y: 320, Width: 160, Height: 60},
	}).Stage(stage)
	rightShape := (&models.TaskShape{
		Name:      "Pack Cold Chain-MainDiagram",
		Task:      taskRight,
		RectShape: models.RectShape{X: 300, Y: 320, Width: 160, Height: 60},
	}).Stage(stage)
	billingShape := (&models.TaskShape{
		Name:      "Confirm Payment-MainDiagram",
		Task:      taskBilling,
		RectShape: models.RectShape{X: 650, Y: 320, Width: 160, Height: 60},
	}).Stage(stage)
	diag.Task_Shapes = append(diag.Task_Shapes, leftShape, rightShape, billingShape)

	stage.Commit()

	// 1. Test UX Tree "Add Fork/Join" button on Participant node AND on Tasks node
	trees := stager.GetTreeStage().GetInstancesSorted[*tree.Tree]()
	if len(trees) == 0 {
		t.Fatalf("expected UX tree to be populated")
	}
	warehouseNode := findTreeNodeByPath(trees[0].RootNodes, "RootLib", "Processes", "OrderFulfillment", "MainDiagram", "Participants", "Warehouse")
	if warehouseNode == nil {
		t.Fatalf("expected Warehouse participant node in UX tree")
	}
	addForkBtn := findMenuButton(warehouseNode, "Add Fork/Join")
	if addForkBtn == nil {
		t.Fatalf("expected 'Add Fork/Join' menu button on Participant node")
	}
	addForkBtn.OnClick()

	if len(lane1.Tasks) != 3 {
		t.Fatalf("expected 3 tasks in Warehouse after clicking Add Fork/Join, got %d", len(lane1.Tasks))
	}
	forkTask := lane1.Tasks[2]
	if !forkTask.IsForkJoin {
		t.Fatalf("expected newly created task to have IsForkJoin=true")
	}
	forkTask.Name = "Fork 1"
	forkTask.SetIsInRenameMode(false)

	// Position Fork 1 horizontally spanning above both left and right tasks
	var forkShape *models.TaskShape
	for _, ts := range diag.Task_Shapes {
		if ts.Task == forkTask {
			forkShape = ts
			break
		}
	}
	if forkShape == nil {
		t.Fatalf("expected TaskShape created for Fork 1")
	}
	if forkShape.Height != models.ForkJoinBarThickness || forkShape.Width != models.ForkJoinDefaultLength {
		t.Fatalf("expected initial ForkShape dimensions %vx%v, got %vx%v",
			models.ForkJoinDefaultLength, models.ForkJoinBarThickness, forkShape.Width, forkShape.Height)
	}
	forkShape.SetX(120)
	forkShape.SetY(220)
	forkShape.SetWidth(300)
	stage.Commit()

	// Also add a second Fork/Join ("Join 1") via the Tasks category node's menu (verifies multiple Fork/Joins per Participant!)
	trees = stager.GetTreeStage().GetInstancesSorted[*tree.Tree]()
	tasksNode := findTreeNodeByPath(trees[0].RootNodes, "RootLib", "Processes", "OrderFulfillment", "MainDiagram", "Participants", "Warehouse", "Tasks")
	if tasksNode == nil {
		t.Fatalf("expected Tasks node under expanded Warehouse participant")
	}
	addJoinBtn := findMenuButton(tasksNode, "Add Fork/Join")
	if addJoinBtn == nil {
		t.Fatalf("expected 'Add Fork/Join' menu button on Tasks node")
	}
	addJoinBtn.OnClick()

	if len(lane1.Tasks) != 4 {
		t.Fatalf("expected 4 tasks in Warehouse after adding Join 1, got %d", len(lane1.Tasks))
	}
	joinTask := lane1.Tasks[3]
	if !joinTask.IsForkJoin || !forkTask.IsForkJoin {
		t.Fatalf("expected multiple IsForkJoin tasks to be allowed in the same participant")
	}
	joinTask.Name = "Join 1"
	joinTask.SetIsInRenameMode(false)

	var joinShape *models.TaskShape
	for _, ts := range diag.Task_Shapes {
		if ts.Task == joinTask {
			joinShape = ts
			break
		}
	}
	joinShape.SetX(120)
	joinShape.SetY(480)
	joinShape.SetWidth(300)
	stage.Commit()

	// 2. Verify SVG rendering of horizontal Fork/Join bar
	layer := stager.GetSvgObject().Layers[0]
	forkRect := findSvgRectByName(layer, "Fork 1")
	if forkRect == nil {
		t.Fatalf("expected svg.Rect for Fork 1")
	}
	if forkRect.Color != "#333333" || forkRect.Stroke != "#333333" || forkRect.FillOpacity != 1.0 || forkRect.StrokeOpacity != 1.0 || forkRect.RX != 2.0 {
		t.Errorf("unexpected Fork 1 rect presentation: color=%q stroke=%q fillOp=%v strokeOp=%v rx=%v",
			forkRect.Color, forkRect.Stroke, forkRect.FillOpacity, forkRect.StrokeOpacity, forkRect.RX)
	}
	if forkRect.Height != models.ForkJoinBarThickness || forkRect.Width != 300 {
		t.Errorf("expected Fork 1 rect 300x%v, got %vx%v", models.ForkJoinBarThickness, forkRect.Width, forkRect.Height)
	}
	if !forkRect.CanHaveLeftHandle || !forkRect.CanHaveRightHandle || forkRect.CanHaveTopHandle || forkRect.CanHaveBottomHandle {
		t.Errorf("expected horizontal Fork 1 to have only left/right handles, got L=%v R=%v T=%v B=%v",
			forkRect.CanHaveLeftHandle, forkRect.CanHaveRightHandle, forkRect.CanHaveTopHandle, forkRect.CanHaveBottomHandle)
	}
	if forkRect.URLPath != "" {
		t.Errorf("expected Fork 1 rect URLPath to be empty, got %q", forkRect.URLPath)
	}
	if len(forkRect.RectAnchoredTexts) != 0 {
		t.Errorf("expected Fork 1 canvas label to be hidden by default, got %d anchored texts", len(forkRect.RectAnchoredTexts))
	}

	// 3. Test Shape Appearance Form: IsNameDisplayed, IsVertical (center-preserving swap), and IsPresentationOverridden
	// First check regular task does NOT expose IsVertical / IsNameDisplayed
	leftRect := findSvgRectByName(layer, "Pick Dry Goods")
	leftRect.OnSelect()
	regShapeFG := stager.GetShapeFormStage().GetInstancesSorted[*form.FormGroup]()[0]
	if hasFormDiv(regShapeFG, "IsVertical") || hasFormDiv(regShapeFG, "IsNameDisplayed") {
		t.Errorf("regular TaskShape should not expose IsVertical or IsNameDisplayed in Appearance form")
	}

	// Now select Fork 1
	forkRect.OnSelect()
	shapeFormStage := stager.GetShapeFormStage()
	forkShapeFG := shapeFormStage.GetInstancesSorted[*form.FormGroup]()[0]
	if !hasFormDiv(forkShapeFG, "IsVertical") || !hasFormDiv(forkShapeFG, "IsNameDisplayed") {
		t.Fatalf("expected IsVertical and IsNameDisplayed in Fork 1 Appearance form")
	}

	// Enable IsNameDisplayed on horizontal Fork 1
	setFormBoolField(t, forkShapeFG, "IsNameDisplayed", true)
	shapeFormStage.Commit()
	forkShapeFG.OnSave.OnSave()

	forkRect = findSvgRectByName(stager.GetSvgObject().Layers[0], "Fork 1")
	if len(forkRect.RectAnchoredTexts) != 1 || forkRect.RectAnchoredTexts[0].Content != "Fork 1" || forkRect.RectAnchoredTexts[0].RectAnchorType != svg.RECT_RIGHT {
		t.Fatalf("expected Fork 1 label displayed at RECT_RIGHT when horizontal, got %+v", forkRect.RectAnchoredTexts)
	}

	// Toggle IsVertical = true on Join 1 and test custom Color override
	joinRect := findSvgRectByName(stager.GetSvgObject().Layers[0], "Join 1")
	joinRect.OnSelect()
	joinShapeFG := shapeFormStage.GetInstancesSorted[*form.FormGroup]()[0]
	setFormBoolField(t, joinShapeFG, "IsVertical", true)
	setFormBoolField(t, joinShapeFG, "IsNameDisplayed", true)
	setFormBoolField(t, joinShapeFG, "IsPresentationOverridden", true)
	setFormStringField(t, joinShapeFG, "Color", "#D32F2F")
	shapeFormStage.Commit()
	joinShapeFG.OnSave.OnSave()

	if !joinShape.IsVertical || joinShape.Width != models.ForkJoinBarThickness || joinShape.Height != 300 {
		t.Fatalf("expected vertical Join 1 to swap dimensions to %vx300, got %vx%v (IsVertical=%v)",
			models.ForkJoinBarThickness, joinShape.Width, joinShape.Height, joinShape.IsVertical)
	}
	joinRect = findSvgRectByName(stager.GetSvgObject().Layers[0], "Join 1")
	if joinRect.CanHaveLeftHandle || joinRect.CanHaveRightHandle || !joinRect.CanHaveTopHandle || !joinRect.CanHaveBottomHandle {
		t.Errorf("expected vertical Join 1 to have only top/bottom handles, got L=%v R=%v T=%v B=%v",
			joinRect.CanHaveLeftHandle, joinRect.CanHaveRightHandle, joinRect.CanHaveTopHandle, joinRect.CanHaveBottomHandle)
	}
	if joinRect.Color != "#D32F2F" {
		t.Errorf("expected overridden Color #D32F2F on Join 1, got %q", joinRect.Color)
	}
	if len(joinRect.RectAnchoredTexts) != 1 || joinRect.RectAnchoredTexts[0].RectAnchorType != svg.RECT_TOP {
		t.Errorf("expected vertical Join 1 label at RECT_TOP, got %+v", joinRect.RectAnchoredTexts)
	}

	// Flip Join 1 back to horizontal for parallel vertical ControlFlow & cross-swimlane DataFlow testing
	joinRect.OnSelect()
	joinShapeFG = shapeFormStage.GetInstancesSorted[*form.FormGroup]()[0]
	setFormBoolField(t, joinShapeFG, "IsVertical", false)
	shapeFormStage.Commit()
	joinShapeFG.OnSave.OnSave()
	joinShape.SetX(120)
	joinShape.SetY(480)
	joinShape.SetWidth(300)
	stage.Commit()

	// 4. Test Interactive SVG Link Creation (ControlFlows + Cross-Participant DataFlows)
	drawLinkOnCanvas := func(startName, endName string) {
		t.Helper()
		currSvg := stager.GetSvgObject()
		currLayer := currSvg.Layers[0]
		sRect := findSvgRectByName(currLayer, startName)
		eRect := findSvgRectByName(currLayer, endName)
		if sRect == nil || eRect == nil {
			t.Fatalf("could not find rects %q (%v) or %q (%v)", startName, sRect, endName, eRect)
		}
		frontSVG := *currSvg
		frontSVG.DrawingState = svg.NOT_DRAWING_LINK
		currSvg.DrawingState = svg.DRAWING_LINK
		frontSVG.StartRect = sRect
		frontSVG.EndRect = eRect
		currSvg.OnUpdate(&frontSVG)
	}

	drawLinkOnCanvas("Fork 1", "Pick Dry Goods")
	drawLinkOnCanvas("Fork 1", "Pack Cold Chain")
	drawLinkOnCanvas("Pick Dry Goods", "Join 1")
	drawLinkOnCanvas("Pack Cold Chain", "Join 1")

	if len(lane1.ControlFlows) != 4 || len(diag.ControlFlow_Shapes) != 4 {
		t.Fatalf("expected 4 intra-participant ControlFlows and shapes, got cf=%d shapes=%d",
			len(lane1.ControlFlows), len(diag.ControlFlow_Shapes))
	}

	// Verify Fork 1 outgoing ControlFlowShapes distribute StartRatio left (< 0.5) and right (> 0.5) with CornerOffsetRatio == 4.5
	cfLeftShape := diag.ControlFlow_Shapes[0]
	cfRightShape := diag.ControlFlow_Shapes[1]
	if !(cfLeftShape.GetStartRatio() < 0.5 && cfRightShape.GetStartRatio() > 0.5) {
		t.Errorf("expected parallel Fork StartRatios distributed (<0.5 and >0.5), got left=%v right=%v",
			cfLeftShape.GetStartRatio(), cfRightShape.GetStartRatio())
	}
	if cfLeftShape.GetCornerOffsetRatio() != 4.5 {
		t.Errorf("expected Fork outgoing ControlFlow CornerOffsetRatio=4.5, got %v", cfLeftShape.GetCornerOffsetRatio())
	}
	// Verify Join 1 incoming ControlFlowShapes distribute EndRatio left (< 0.5) and right (> 0.5)
	cfJoinLeftShape := diag.ControlFlow_Shapes[2]
	cfJoinRightShape := diag.ControlFlow_Shapes[3]
	if !(cfJoinLeftShape.GetEndRatio() < 0.5 && cfJoinRightShape.GetEndRatio() > 0.5) {
		t.Errorf("expected parallel Join EndRatios distributed (<0.5 and >0.5), got left=%v right=%v",
			cfJoinLeftShape.GetEndRatio(), cfJoinRightShape.GetEndRatio())
	}

	// Cross-participant DataFlow into Join 1 ("Confirm Payment" in Billing -> "Join 1" in Warehouse)
	// and from Fork 1 to ExternalParticipant ("TailCarrier-MainDiagram")
	drawLinkOnCanvas("Confirm Payment", "Join 1")
	drawLinkOnCanvas("Fork 1", "TailCarrier-MainDiagram")

	if len(proc.DataFlows) != 2 || len(diag.DataFlow_Shapes) != 2 {
		t.Fatalf("expected 2 DataFlows connected to Fork/Join to be preserved by semantic rules, got df=%d shapes=%d",
			len(proc.DataFlows), len(diag.DataFlow_Shapes))
	}

	// 5. Verify Semantic Mutual Exclusivity & Type Clearing
	forkTask.IsStartTask = true
	stage.Commit()
	// Because IsStartTask was set true alongside IsForkJoin, IsForkJoin must be cleared to false
	if forkTask.IsStartTask && forkTask.IsForkJoin {
		t.Fatalf("expected IsStartTask and IsForkJoin to be mutually exclusive")
	}

	forkTask.IsStartTask = false
	forkTask.IsEndTask = true
	forkTask.IsForkJoin = true
	stage.Commit()
	if forkTask.IsEndTask && forkTask.IsForkJoin {
		t.Fatalf("expected IsEndTask and IsForkJoin to be mutually exclusive")
	}

	// Restore Fork 1 as IsForkJoin and test Type clearing (preserving Fork 1's name)
	forkTask.IsEndTask = false
	forkTask.IsForkJoin = true
	forkTask.Type = subProc
	stage.Commit()
	if forkTask.Type != nil {
		t.Fatalf("expected forkTask.Type to be cleared to nil on IsForkJoin task, got %v", forkTask.Type)
	}
	if forkTask.Name != "Fork 1" {
		t.Fatalf("expected Fork 1 name to remain 'Fork 1', got %q", forkTask.Name)
	}

	// 6. Verify Static Website Export excludes IsForkJoin tasks from Process Tasks summary table
	procPage := proc.GeneratePage(stager)
	var tasksTableSection string
	for _, sec := range procPage.Sections {
		if sec.Name == "Tasks" {
			tasksTableSection = sec.MardownContent
			break
		}
	}
	if !strings.Contains(tasksTableSection, "Pick Dry Goods") || !strings.Contains(tasksTableSection, "Pack Cold Chain") {
		t.Errorf("expected regular tasks in Process page Tasks table, got:\n%s", tasksTableSection)
	}
	if strings.Contains(tasksTableSection, "Fork 1") || strings.Contains(tasksTableSection, "Join 1") {
		t.Errorf("expected IsForkJoin tasks to be excluded from Process page Tasks table, got:\n%s", tasksTableSection)
	}

	// 7. Verify AST Marshall / Unmarshall round-trip of IsForkJoin, IsVertical, and IsNameDisplayed
	joinShape.IsVertical = true
	joinShape.IsNameDisplayed = true
	stage.Commit()

	outFile := filepath.Join(t.TempDir(), "fork_join_roundtrip.go")
	stage.ComputeReverseMaps()
	stage.ComputeInstancesNb()
	stage.ComputeReferenceAndOrders()
	stage.MarshallFile(outFile, "github.com/fullstack-lang/gong/dsm/process/go/models", "main")

	reloaded := models.NewStage("reloaded")
	if err := reloaded.ParseAstFile(outFile, false); err != nil {
		t.Fatalf("ParseAstFile failed on round-trip file: %v", err)
	}
	var foundForkTask, foundVerticalNamedJoinShape bool
	for _, rt := range reloaded.GetInstancesSorted[*models.Task]() {
		if rt.Name == "Fork 1" && rt.IsForkJoin {
			foundForkTask = true
		}
	}
	for _, rts := range reloaded.GetInstancesSorted[*models.TaskShape]() {
		if rts.Task != nil && rts.Task.Name == "Join 1" && rts.IsVertical && rts.IsNameDisplayed {
			foundVerticalNamedJoinShape = true
		}
	}
	if !foundForkTask || !foundVerticalNamedJoinShape {
		t.Fatalf("expected IsForkJoin, IsVertical, and IsNameDisplayed to survive AST round-trip (fork=%v joinShape=%v)",
			foundForkTask, foundVerticalNamedJoinShape)
	}

	// 8. Verify backward compatibility with existing stage files in data/
	for _, dataFile := range []string{"data/stage.go", "data/icao4444.go", "data/simple.go"} {
		legacyStage := models.NewStage("legacy_" + dataFile)
		if err := legacyStage.ParseAstFile(dataFile, false); err != nil {
			t.Fatalf("failed to parse existing data file %s: %v", dataFile, err)
		}
	}
}

func TestStageGoForkJoinLayout(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("verify_stage_go", "data/stage.go", "", true, true, false)
	stager := models.NewStager(stack.R, stack.Stage, stack.Probe)
	stage := stack.Stage

	var forkTask, joinTask, prepDocsTask *models.Task
	for _, task := range stage.GetInstancesSorted[*models.Task]() {
		switch task.Name {
		case "Fork Fulfillment":
			forkTask = task
		case "Join Fulfillment":
			joinTask = task
		case "Prepare Shipping Docs":
			prepDocsTask = task
		}
	}
	if forkTask == nil || !forkTask.IsForkJoin || joinTask == nil || !joinTask.IsForkJoin || prepDocsTask == nil {
		t.Fatalf("expected Fork Fulfillment, Join Fulfillment, and Prepare Shipping Docs in data/stage.go (fork=%v join=%v prep=%v)",
			forkTask, joinTask, prepDocsTask)
	}

	rectsOverlap := func(a, b *models.TaskShape) bool {
		return a.X < b.X+b.Width && a.X+a.Width > b.X &&
			a.Y < b.Y+b.Height && a.Y+a.Height > b.Y
	}

	for _, diag := range stage.GetInstancesSorted[*models.DiagramProcess]() {
		for _, other := range stage.GetInstancesSorted[*models.DiagramProcess]() {
			other.IsChecked = (other == diag)
		}
		stage.Commit()

		for i := 0; i < len(diag.Task_Shapes); i++ {
			for j := i + 1; j < len(diag.Task_Shapes); j++ {
				a, b := diag.Task_Shapes[i], diag.Task_Shapes[j]
				if rectsOverlap(a, b) {
					t.Errorf("diagram %q: TaskShapes %q and %q overlap", diag.Name, a.Name, b.Name)
				}
			}
		}

		if diag.Name == "Order to Cash - Full Process" || diag.Name == "Order Fulfillment & Delivery" {
			layer := stager.GetSvgObject().Layers[0]
			forkRect := findSvgRectByName(layer, "Fork Fulfillment")
			joinRect := findSvgRectByName(layer, "Join Fulfillment")
			if forkRect == nil || joinRect == nil {
				t.Fatalf("diagram %q: expected SVG rects for Fork Fulfillment and Join Fulfillment", diag.Name)
			}
			if forkRect.Height != models.ForkJoinBarThickness || joinRect.Height != models.ForkJoinBarThickness {
				t.Errorf("diagram %q: expected 8px bar thickness, got fork=%v join=%v",
					diag.Name, forkRect.Height, joinRect.Height)
			}
		}
	}
}


