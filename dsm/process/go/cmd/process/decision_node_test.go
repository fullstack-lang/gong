package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/fullstack-lang/gong/dsm/process/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/process/go/models"
	form "github.com/fullstack-lang/gong/lib/form/go/models"
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func TestDecisionNodeEndToEnd(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("test_decision_node", "", "", true, true, false)
	stager := models.NewStager(
		stack.R,
		stack.Stage,
		stack.Probe,
	)
	stage := stack.Stage
	rootLib := stager.GetRootLibrary()
	rootLib.Name = "RootLib"
	rootLib.IsProcessesNodeExpanded = true

	proc := (&models.Process{Name: "CreditCheckProcess"}).Stage(stage)
	subProc := (&models.Process{Name: "SubProc"}).Stage(stage)
	rootLib.RootProcesses = append(rootLib.RootProcesses, proc, subProc)
	rootLib.ProcesssWhoseNodeIsExpanded = append(rootLib.ProcesssWhoseNodeIsExpanded, proc)

	sales := (&models.Participant{Name: "Sales Desk"}).Stage(stage)
	risk := (&models.Participant{Name: "Credit Risk"}).Stage(stage)
	proc.Participants = append(proc.Participants, sales, risk)

	validateTask := (&models.Task{Name: "Validate Order"}).Stage(stage)
	confirmTask := (&models.Task{Name: "Confirm Order"}).Stage(stage)
	rejectTask := (&models.Task{Name: "Reject Order"}).Stage(stage)
	sales.Tasks = append(sales.Tasks, validateTask, confirmTask, rejectTask)

	scoreTask := (&models.Task{Name: "Compute Credit Score"}).Stage(stage)
	risk.Tasks = append(risk.Tasks, scoreTask)

	diag := (&models.DiagramProcess{
		Name:                       "DecisionDiagram",
		IsChecked:                  true,
		IsEditable_:                true,
		DefaultBoxWidth:            180,
		DefaultBoxHeigth:           60,
		Width:                      1400,
		Height:                     1000,
		IsParticipantsNodeExpanded: true,
	}).Stage(stage)
	proc.DiagramProcesss = append(proc.DiagramProcesss, diag)
	proc.DiagramProcessWhoseNodeIsExpanded = append(proc.DiagramProcessWhoseNodeIsExpanded, diag)

	procShape := (&models.ProcessShape{
		Name:      "CreditCheckProcess-DecisionDiagram",
		Process:   proc,
		RectShape: models.RectShape{X: 50, Y: 50, Width: 960, Height: 800},
	}).Stage(stage)
	diag.Process_Shapes = append(diag.Process_Shapes, procShape)

	salesShape := (&models.ParticipantShape{
		Name:        "Sales Desk-DecisionDiagram",
		Participant: sales,
		WidthWeight: 1.6,
	}).Stage(stage)
	riskShape := (&models.ParticipantShape{
		Name:        "Credit Risk-DecisionDiagram",
		Participant: risk,
		WidthWeight: 1.0,
	}).Stage(stage)
	diag.Participant_Shapes = append(diag.Participant_Shapes, salesShape, riskShape)

	validateShape := (&models.TaskShape{
		Name:      "Validate Order-DecisionDiagram",
		Task:      validateTask,
		RectShape: models.RectShape{X: 140, Y: 200, Width: 180, Height: 60},
	}).Stage(stage)
	confirmShape := (&models.TaskShape{
		Name:      "Confirm Order-DecisionDiagram",
		Task:      confirmTask,
		RectShape: models.RectShape{X: 140, Y: 460, Width: 180, Height: 60},
	}).Stage(stage)
	rejectShape := (&models.TaskShape{
		Name:      "Reject Order-DecisionDiagram",
		Task:      rejectTask,
		RectShape: models.RectShape{X: 380, Y: 460, Width: 180, Height: 60},
	}).Stage(stage)
	scoreShape := (&models.TaskShape{
		Name:      "Compute Credit Score-DecisionDiagram",
		Task:      scoreTask,
		RectShape: models.RectShape{X: 700, Y: 320, Width: 180, Height: 60},
	}).Stage(stage)
	diag.Task_Shapes = append(diag.Task_Shapes, validateShape, confirmShape, rejectShape, scoreShape)

	stage.Commit()

	// 1. Test "Add Decision Node" button in UX Tree (with BUTTON_diamond icon)
	trees := stager.GetTreeStage().GetInstancesSorted[*tree.Tree]()
	salesNode := findTreeNodeByPath(trees[0].RootNodes, "RootLib", "Processes", "CreditCheckProcess", "DecisionDiagram", "Participants", "Sales Desk")
	if salesNode == nil {
		t.Fatalf("expected Sales Desk participant node in UX tree")
	}
	addDecBtn := findMenuButton(salesNode, "Add Decision Node")
	if addDecBtn == nil {
		t.Fatalf("expected 'Add Decision Node' button in Participant menu")
	}
	if addDecBtn.Icon != string(buttons.BUTTON_diamond) {
		t.Errorf("expected BUTTON_diamond icon on 'Add Decision Node', got %q", addDecBtn.Icon)
	}
	addDecBtn.OnClick()

	if len(sales.Tasks) != 4 {
		t.Fatalf("expected 4 tasks in Sales Desk after clicking Add Decision Node, got %d", len(sales.Tasks))
	}
	decisionTask := sales.Tasks[3]
	if !decisionTask.IsDecisionNode {
		t.Fatalf("expected newly created task to have IsDecisionNode=true")
	}
	decisionTask.Name = "Credit OK?"
	decisionTask.SetIsInRenameMode(false)

	var decisionShape *models.TaskShape
	for _, ts := range diag.Task_Shapes {
		if ts.Task == decisionTask {
			decisionShape = ts
			break
		}
	}
	if decisionShape == nil {
		t.Fatalf("expected TaskShape for Credit OK?")
	}
	if decisionShape.Width != models.DecisionDiamondSize || decisionShape.Height != models.DecisionDiamondSize || !decisionShape.IsNameDisplayed {
		t.Fatalf("expected initial 50x50 decision shape with IsNameDisplayed=true, got %vx%v IsNameDisplayed=%v",
			decisionShape.Width, decisionShape.Height, decisionShape.IsNameDisplayed)
	}
	// Center diamond horizontally with Validate Order and Confirm Order (center X = 230 -> X = 205)
	decisionShape.SetX(205)
	decisionShape.SetY(325)
	stage.Commit()

	// 2. Verify SVG 50x50 Diamond Rendering & Default Label
	layer := stager.GetSvgObject().Layers[0]
	decRect := findSvgRectByName(layer, "Credit OK?")
	if decRect == nil {
		t.Fatalf("expected svg.Rect for Credit OK?")
	}
	if decRect.Width != models.DecisionDiamondSize || decRect.Height != models.DecisionDiamondSize {
		t.Errorf("expected 50x50 svg.Rect, got %vx%v", decRect.Width, decRect.Height)
	}
	if decRect.FillOpacity != 0.0 || decRect.StrokeOpacity != 0.0 {
		t.Errorf("expected outer bounding svg.Rect to be transparent, got fillOp=%v strokeOp=%v",
			decRect.FillOpacity, decRect.StrokeOpacity)
	}
	if decRect.CanHaveLeftHandle || decRect.CanHaveRightHandle || decRect.CanHaveTopHandle || decRect.CanHaveBottomHandle {
		t.Errorf("expected Decision Node to have no resize handles")
	}
	if decRect.URLPath != "" {
		t.Errorf("expected Decision Node URLPath to be empty, got %q", decRect.URLPath)
	}
	if len(decRect.RectAnchoredPaths) != 1 {
		t.Fatalf("expected 1 diamond RectAnchoredPath, got %d", len(decRect.RectAnchoredPaths))
	}
	diamondPath := decRect.RectAnchoredPaths[0]
	if diamondPath.Color != "#E3F2FD" || diamondPath.Stroke != "#90CAF9" || diamondPath.StrokeWidth != 2.0 || diamondPath.StrokeOpacity != 1.0 {
		t.Errorf("unexpected diamond presentation: color=%q stroke=%q width=%v strokeOp=%v",
			diamondPath.Color, diamondPath.Stroke, diamondPath.StrokeWidth, diamondPath.StrokeOpacity)
	}
	if len(decRect.RectAnchoredTexts) != 1 || decRect.RectAnchoredTexts[0].Content != "Credit OK?" {
		t.Fatalf("expected 'Credit OK?' label displayed by default, got %+v", decRect.RectAnchoredTexts)
	}

	// 3. Verify Shape Appearance Form (IsNameDisplayed exposed, IsVertical NOT exposed, custom Color/Stroke overrides diamond)
	decRect.OnSelect()
	shapeFormStage := stager.GetShapeFormStage()
	decShapeFG := shapeFormStage.GetInstancesSorted[*form.FormGroup]()[0]
	if !hasFormDiv(decShapeFG, "IsNameDisplayed") {
		t.Fatalf("expected IsNameDisplayed in Decision Node Appearance form")
	}
	if hasFormDiv(decShapeFG, "IsVertical") {
		t.Errorf("did not expect IsVertical in Decision Node Appearance form")
	}

	setFormBoolField(t, decShapeFG, "IsNameDisplayed", false)
	setFormBoolField(t, decShapeFG, "IsPresentationOverridden", true)
	setFormStringField(t, decShapeFG, "Color", "#FFF9C4")
	setFormStringField(t, decShapeFG, "Stroke", "#FBC02D")
	shapeFormStage.Commit()
	decShapeFG.OnSave.OnSave()

	decRect = findSvgRectByName(stager.GetSvgObject().Layers[0], "Credit OK?")
	if len(decRect.RectAnchoredTexts) != 0 {
		t.Errorf("expected label hidden after unchecking IsNameDisplayed")
	}
	if decRect.RectAnchoredPaths[0].Color != "#FFF9C4" || decRect.RectAnchoredPaths[0].Stroke != "#FBC02D" {
		t.Errorf("expected overridden diamond color/stroke, got color=%q stroke=%q",
			decRect.RectAnchoredPaths[0].Color, decRect.RectAnchoredPaths[0].Stroke)
	}
	if decRect.FillOpacity != 0.0 || decRect.StrokeOpacity != 0.0 {
		t.Errorf("outer rect must stay invisible even when presentation is overridden")
	}

	// Re-enable label display
	decRect.OnSelect()
	decShapeFG = shapeFormStage.GetInstancesSorted[*form.FormGroup]()[0]
	setFormBoolField(t, decShapeFG, "IsNameDisplayed", true)
	shapeFormStage.Commit()
	decShapeFG.OnSave.OnSave()

	// 4. Test Interactive SVG Link Drawing & Guard Labels on Outgoing Branches
	drawLinkOnCanvas := func(startName, endName string) {
		t.Helper()
		currSvg := stager.GetSvgObject()
		currLayer := currSvg.Layers[0]
		sRect := findSvgRectByName(currLayer, startName)
		eRect := findSvgRectByName(currLayer, endName)
		if sRect == nil || eRect == nil {
			t.Fatalf("could not find rects %q or %q", startName, endName)
		}
		frontSVG := *currSvg
		frontSVG.DrawingState = svg.NOT_DRAWING_LINK
		currSvg.DrawingState = svg.DRAWING_LINK
		frontSVG.StartRect = sRect
		frontSVG.EndRect = eRect
		currSvg.OnUpdate(&frontSVG)
	}

	drawLinkOnCanvas("Validate Order", "Credit OK?")
	drawLinkOnCanvas("Credit OK?", "Confirm Order") // directly below -> ORIENTATION_VERTICAL (bottom vertex)
	drawLinkOnCanvas("Credit OK?", "Reject Order")  // offset right -> ORIENTATION_HORIZONTAL (right vertex)
	drawLinkOnCanvas("Compute Credit Score", "Credit OK?")

	if len(sales.ControlFlows) != 3 || len(diag.ControlFlow_Shapes) != 3 {
		t.Fatalf("expected 3 ControlFlows in Sales Desk, got %d", len(sales.ControlFlows))
	}
	if len(proc.DataFlows) != 1 {
		t.Fatalf("expected cross-swimlane DataFlow into Decision Node to be preserved")
	}

	cfYesShape := diag.ControlFlow_Shapes[1]
	cfNoShape := diag.ControlFlow_Shapes[2]
	if cfYesShape.GetStartOrientation() != models.ORIENTATION_VERTICAL || cfYesShape.GetStartRatio() != 0.5 {
		t.Errorf("expected downward branch to exit bottom vertex (VERTICAL, 0.5), got %v %v",
			cfYesShape.GetStartOrientation(), cfYesShape.GetStartRatio())
	}
	if cfNoShape.GetStartOrientation() != models.ORIENTATION_HORIZONTAL || cfNoShape.GetStartRatio() != 0.5 {
		t.Errorf("expected sideways branch to exit side vertex (HORIZONTAL, 0.5), got %v %v",
			cfNoShape.GetStartOrientation(), cfNoShape.GetStartRatio())
	}

	// Set Guard conditions via ControlFlow Probe Form
	formStage := stack.Probe.GetFormStage()
	stack.Probe.FillUpFormFromGongstruct(cfYesShape.ControlFlow, "ControlFlow")
	fgCFYes := formStage.GetInstancesSorted[*form.FormGroup]()[0]
	setFormStringField(t, fgCFYes, "Guard", "Yes")
	formStage.Commit()
	fgCFYes.OnSave.OnSave()

	stack.Probe.FillUpFormFromGongstruct(cfNoShape.ControlFlow, "ControlFlow")
	fgCFNo := formStage.GetInstancesSorted[*form.FormGroup]()[0]
	setFormStringField(t, fgCFNo, "Guard", "No")
	formStage.Commit()
	fgCFNo.OnSave.OnSave()

	// Verify "[Yes]" and "[No]" LinkAnchoredTexts are rendered on SVG links and written into SVG markup
	svgMarkup, _, _ := stager.GetSvgObject().GenerateString()
	if !strings.Contains(svgMarkup, "[Yes]") || !strings.Contains(svgMarkup, "[No]") {
		t.Errorf("expected '[Yes]' and '[No]' guard labels in generated SVG output")
	}

	// 5. Verify Semantic Mutual Exclusivity & Type Clearing
	decisionTask.IsForkJoin = true
	stage.Commit()
	if decisionTask.IsForkJoin && decisionTask.IsDecisionNode {
		t.Fatalf("expected IsForkJoin and IsDecisionNode to be mutually exclusive")
	}
	decisionTask.IsForkJoin = false
	decisionTask.IsDecisionNode = true
	decisionTask.Type = subProc
	stage.Commit()
	if decisionTask.Type != nil {
		t.Fatalf("expected decisionTask.Type to be cleared to nil, got %v", decisionTask.Type)
	}

	// 6. Verify Static Website Export excludes IsDecisionNode from Tasks table and includes Guard in Control Flows table
	procPage := proc.GeneratePage(stager)
	var tasksMD, cfMD string
	for _, sec := range procPage.Sections {
		if sec.Name == "Tasks" {
			tasksMD = sec.MardownContent
		}
		if sec.Name == "Control Flows" {
			cfMD = sec.MardownContent
		}
	}
	if strings.Contains(tasksMD, "Credit OK?") {
		t.Errorf("expected IsDecisionNode task to be excluded from Process Tasks table, got:\n%s", tasksMD)
	}
	if !strings.Contains(cfMD, "| Yes |") || !strings.Contains(cfMD, "| No |") {
		t.Errorf("expected Guard values 'Yes' and 'No' in Process Control Flows table, got:\n%s", cfMD)
	}

	// 7. Verify AST Marshall / Unmarshall round-trip of IsDecisionNode and ControlFlow.Guard
	outFile := filepath.Join(t.TempDir(), "decision_roundtrip.go")
	stage.ComputeReverseMaps()
	stage.ComputeInstancesNb()
	stage.ComputeReferenceAndOrders()
	stage.MarshallFile(outFile, "github.com/fullstack-lang/gong/dsm/process/go/models", "main")

	reloaded := models.NewStage("reloaded_decision")
	if err := reloaded.ParseAstFile(outFile, false); err != nil {
		t.Fatalf("ParseAstFile failed: %v", err)
	}
	var foundDecision, foundGuards bool
	guardCount := 0
	for _, rt := range reloaded.GetInstancesSorted[*models.Task]() {
		if rt.Name == "Credit OK?" && rt.IsDecisionNode {
			foundDecision = true
		}
	}
	for _, rcf := range reloaded.GetInstancesSorted[*models.ControlFlow]() {
		if rcf.Guard == "Yes" || rcf.Guard == "No" {
			guardCount++
		}
	}
	foundGuards = (guardCount == 2)
	if !foundDecision || !foundGuards {
		t.Fatalf("expected IsDecisionNode and ControlFlow.Guard to survive AST round-trip (decision=%v guards=%v)",
			foundDecision, foundGuards)
	}
}

func updateStageGoWithDecisionNode(t *testing.T) {
	t.Helper()
	stack := level1stack.NewLevel1StackDelta("update_stage_decision", "data/stage.go", "", true, true, false)
	models.NewStager(stack.R, stack.Stage, stack.Probe)
	stage := stack.Stage

	var sales *models.Participant
	for _, p := range stage.GetInstancesSorted[*models.Participant]() {
		if p.Name == "Sales & Order Desk" {
			sales = p
		}
	}
	if sales == nil {
		t.Fatalf("missing Sales & Order Desk participant")
	}

	taskByName := make(map[string]*models.Task)
	for _, task := range stage.GetInstancesSorted[*models.Task]() {
		taskByName[task.Name] = task
	}
	if taskByName["Credit Approved?"] != nil {
		return
	}

	startSales := taskByName["Start Order Desk"]
	validateOrder := taskByName["Receive & Validate Order"]
	confirmOrder := taskByName["Confirm Sales Order"]
	endSales := taskByName["End Order Desk"]

	creditDecision := (&models.Task{
		Name:           "Credit Approved?",
		IsDecisionNode: true,
	}).Stage(stage)
	rejectOrder := (&models.Task{
		Name:        "Reject Sales Order",
		Description: "Notify customer of credit hold or order rejection",
	}).Stage(stage)
	orderRejected := (&models.Task{
		Name:      "Order Rejected",
		IsEndTask: true,
	}).Stage(stage)

	sales.Tasks = []*models.Task{
		startSales,
		validateOrder,
		creditDecision,
		confirmOrder,
		endSales,
		rejectOrder,
		orderRejected,
	}

	var cfValidate *models.ControlFlow
	for _, cf := range sales.ControlFlows {
		if cf.Start == validateOrder && cf.End == confirmOrder {
			cfValidate = cf
		}
	}
	if cfValidate == nil {
		t.Fatalf("missing ControlFlow from Receive & Validate Order to Confirm Sales Order")
	}
	cfValidate.End = creditDecision
	cfValidate.Name = `"Receive & Validate Order" to "Credit Approved?"`

	cfApproved := (&models.ControlFlow{
		Name:  `"Credit Approved?" to "Confirm Sales Order"`,
		Start: creditDecision,
		End:   confirmOrder,
		Guard: "approved",
	}).Stage(stage)
	cfRejected := (&models.ControlFlow{
		Name:  `"Credit Approved?" to "Reject Sales Order"`,
		Start: creditDecision,
		End:   rejectOrder,
		Guard: "rejected",
	}).Stage(stage)
	cfRejectEnd := (&models.ControlFlow{
		Name:  `"Reject Sales Order" to "Order Rejected"`,
		Start: rejectOrder,
		End:   orderRejected,
	}).Stage(stage)
	sales.ControlFlows = append(sales.ControlFlows, cfApproved, cfRejected, cfRejectEnd)

	addTaskShape := func(diag *models.DiagramProcess, task *models.Task, x, y, w, h float64, showName bool) *models.TaskShape {
		ts := (&models.TaskShape{
			Name:            fmt.Sprintf("%s-%s", task.Name, diag.Name),
			Task:            task,
			IsNameDisplayed: showName,
			RectShape:       models.RectShape{X: x, Y: y, Width: w, Height: h},
		}).Stage(stage)
		switch {
		case task.IsDecisionNode:
			ts.Color = "#E3F2FD"
			ts.FillOpacity = 1.0
			ts.Stroke = "#90CAF9"
			ts.StrokeWidth = 2.0
			ts.StrokeOpacity = 1.0
			ts.FontSize = "14px"
			ts.FontColor = "#333333"
		case task.IsEndTask:
			ts.Stroke = "black"
			ts.StrokeWidth = 2.0
			ts.StrokeOpacity = 1.0
			ts.RX = 15.0
			ts.FontSize = "16px"
			ts.FontColor = "black"
		default:
			ts.Color = "#E3F2FD"
			ts.FillOpacity = 1.0
			ts.Stroke = "#90CAF9"
			ts.StrokeWidth = 1.5
			ts.StrokeOpacity = 1.0
			ts.RX = 5.0
			ts.FontSize = "16px"
			ts.FontColor = "#333333"
		}
		diag.Task_Shapes = append(diag.Task_Shapes, ts)
		return ts
	}

	addCFShape := func(diag *models.DiagramProcess, cf *models.ControlFlow, startOr, endOr models.OrientationType, startR, endR float64) *models.ControlFlowShape {
		cfs := (&models.ControlFlowShape{
			Name:        cf.Name,
			ControlFlow: cf,
			LinkShape: models.LinkShape{
				StartOrientation:  startOr,
				EndOrientation:    endOr,
				StartRatio:        startR,
				EndRatio:          endR,
				CornerOffsetRatio: 1.35,
			},
		}).Stage(stage)
		cfs.Stroke = "#333333"
		cfs.StrokeWidth = 1.5
		cfs.StrokeOpacity = 1.0
		diag.ControlFlow_Shapes = append(diag.ControlFlow_Shapes, cfs)
		return cfs
	}

	for _, diag := range stage.GetInstancesSorted[*models.DiagramProcess]() {
		pShape := diag.Process_Shapes[0]
		for _, ps := range diag.Participant_Shapes {
			if ps.Participant == sales {
				ps.WidthWeight = 1.4
			}
		}

		findTS := func(taskName string) *models.TaskShape {
			for _, ts := range diag.Task_Shapes {
				if ts.Task != nil && ts.Task.Name == taskName {
					return ts
				}
			}
			return nil
		}

		switch diag.Name {
		case "Order to Cash - Full Process":
			pShape.Width = 1360
			diag.Width = 2115

			// Shift Warehouse (+150), Finance (+155), Logistics Carrier (+160), and cross-lane DataShapes
			for _, ts := range diag.Task_Shapes {
				if ts.Task == nil {
					continue
				}
				switch ts.Task.Name {
				case "Start Logistics", "Check & Reserve Inventory", "Fork Fulfillment", "Pick & Pack Products", "Prepare Shipping Docs", "Join Fulfillment", "Dispatch & Ship Goods", "End Logistics":
					ts.X += 150
				case "Start Finance", "Generate Commercial Invoice", "Process Payment & Settle Account", "Close Order & Archive Record":
					ts.X += 155
				}
			}
			for _, eps := range diag.ExternalParticipant_Shapes {
				if eps.ExternalParticipant != nil && eps.ExternalParticipant.Name == "Logistics Carrier" {
					eps.X += 160
				}
			}
			for _, ns := range diag.Note_Shapes {
				ns.Y = 705
			}
			for _, ds := range diag.Data_Shapes {
				if ds.DataFlow == nil {
					continue
				}
				switch ds.DataFlow.Name {
				case "Confirm Sales Order to Customer":
					ds.X = 110
					ds.Y = 515
				case `"Confirm Sales Order" to "Check & Reserve Inventory"`:
					ds.X = 640
					ds.Y = 490
				case "Dispatch & Ship Goods to Logistics Carrier",
					`"Dispatch & Ship Goods" to "Generate Commercial Invoice"`,
					"Generate Commercial Invoice to Customer",
					"Customer to Process Payment & Settle Account":
					ds.X += 150
				}
			}
			for _, dfs := range diag.DataFlow_Shapes {
				if dfs.DataFlow != nil && dfs.DataFlow.Name == "Confirm Sales Order to Customer" {
					dfs.StartRatio = 0.5
					dfs.EndRatio = 0.381548
				}
			}

			tsStart := findTS("Start Order Desk")
			tsStart.X, tsStart.Y, tsStart.Width, tsStart.Height = 255, 185, 135, 20
			tsVal := findTS("Receive & Validate Order")
			tsVal.X, tsVal.Y, tsVal.Width, tsVal.Height = 280, 250, 220, 65
			tsConf := findTS("Confirm Sales Order")
			tsConf.X, tsConf.Y, tsConf.Width, tsConf.Height = 280, 505, 220, 65
			tsEnd := findTS("End Order Desk")
			tsEnd.X, tsEnd.Y, tsEnd.Width, tsEnd.Height = 256, 615, 152, 36

			addTaskShape(diag, creditDecision, 365, 350, 50, 50, true)
			addTaskShape(diag, rejectOrder, 535, 342.5, 180, 65, false)
			addTaskShape(diag, orderRejected, 498, 442, 145, 36, false)

			for _, cfs := range diag.ControlFlow_Shapes {
				switch cfs.ControlFlow {
				case cfValidate:
					cfs.Name = cfValidate.Name
					cfs.StartRatio = 0.5
					cfs.EndRatio = 0.5
				}
				if cfs.ControlFlow != nil && cfs.ControlFlow.Start == startSales {
					cfs.StartRatio = 1.0
					cfs.EndRatio = 0.5
				}
				if cfs.ControlFlow != nil && cfs.ControlFlow.Start == confirmOrder && cfs.ControlFlow.End == endSales {
					cfs.StartRatio = 0.5
					cfs.EndRatio = 134.0 / 152.0
				}
			}
			addCFShape(diag, cfApproved, models.ORIENTATION_VERTICAL, models.ORIENTATION_VERTICAL, 0.5, 0.5)
			addCFShape(diag, cfRejected, models.ORIENTATION_HORIZONTAL, models.ORIENTATION_HORIZONTAL, 0.5, 0.5)
			addCFShape(diag, cfRejectEnd, models.ORIENTATION_VERTICAL, models.ORIENTATION_VERTICAL, 0.5, 127.0/145.0)

		case "Order Fulfillment & Delivery":
			pShape.Width = 1030
			diag.Width = 1775

			for _, ts := range diag.Task_Shapes {
				if ts.Task == nil {
					continue
				}
				switch ts.Task.Name {
				case "Start Logistics", "Check & Reserve Inventory", "Fork Fulfillment", "Pick & Pack Products", "Prepare Shipping Docs", "Join Fulfillment", "Dispatch & Ship Goods", "End Logistics":
					ts.X += 140
				}
			}
			for _, eps := range diag.ExternalParticipant_Shapes {
				if eps.ExternalParticipant != nil && eps.ExternalParticipant.Name == "Logistics Carrier" {
					eps.X += 140
				}
			}
			for _, ds := range diag.Data_Shapes {
				if ds.DataFlow == nil {
					continue
				}
				switch ds.DataFlow.Name {
				case "Confirm Sales Order to Customer":
					ds.Y = 515
				case `"Confirm Sales Order" to "Check & Reserve Inventory"`:
					ds.X = 630
					ds.Y = 490
				case "Dispatch & Ship Goods to Logistics Carrier":
					ds.X += 140
				}
			}
			for _, dfs := range diag.DataFlow_Shapes {
				if dfs.DataFlow != nil && dfs.DataFlow.Name == "Confirm Sales Order to Customer" {
					dfs.StartRatio = 0.5
					dfs.EndRatio = 0.465
				}
			}

			tsStart := findTS("Start Order Desk")
			tsStart.X, tsStart.Y, tsStart.Width, tsStart.Height = 279, 181, 106, 20
			tsVal := findTS("Receive & Validate Order")
			tsVal.X, tsVal.Y, tsVal.Width, tsVal.Height = 275, 250, 220, 65
			tsConf := findTS("Confirm Sales Order")
			tsConf.X, tsConf.Y, tsConf.Width, tsConf.Height = 275, 505, 220, 65
			tsEnd := findTS("End Order Desk")
			tsEnd.X, tsEnd.Y, tsEnd.Width, tsEnd.Height = 292, 615, 111, 36

			addTaskShape(diag, creditDecision, 360, 350, 50, 50, true)
			addTaskShape(diag, rejectOrder, 530, 342.5, 180, 65, false)
			addTaskShape(diag, orderRejected, 493, 442, 145, 36, false)

			for _, cfs := range diag.ControlFlow_Shapes {
				if cfs.ControlFlow == cfValidate {
					cfs.Name = cfValidate.Name
					cfs.StartRatio = 0.5
					cfs.EndRatio = 0.5
				}
				if cfs.ControlFlow != nil && cfs.ControlFlow.Start == startSales {
					cfs.StartRatio = 1.0
					cfs.EndRatio = 0.5
				}
				if cfs.ControlFlow != nil && cfs.ControlFlow.Start == confirmOrder && cfs.ControlFlow.End == endSales {
					cfs.StartRatio = 0.5
					cfs.EndRatio = 93.0 / 111.0
				}
			}
			addCFShape(diag, cfApproved, models.ORIENTATION_VERTICAL, models.ORIENTATION_VERTICAL, 0.5, 0.5)
			addCFShape(diag, cfRejected, models.ORIENTATION_HORIZONTAL, models.ORIENTATION_HORIZONTAL, 0.5, 0.5)
			addCFShape(diag, cfRejectEnd, models.ORIENTATION_VERTICAL, models.ORIENTATION_VERTICAL, 0.5, 127.0/145.0)

		case "Invoicing & Payment Settlement":
			pShape.Width = 890
			diag.Width = 1444

			for _, ts := range diag.Task_Shapes {
				if ts.Task == nil {
					continue
				}
				switch ts.Task.Name {
				case "Start Finance", "Generate Commercial Invoice", "Process Payment & Settle Account", "Close Order & Archive Record":
					ts.X += 145
				}
			}
			for _, ds := range diag.Data_Shapes {
				if ds.DataFlow == nil {
					continue
				}
				switch ds.DataFlow.Name {
				case "Confirm Sales Order to Customer":
					ds.Y = 510
				case "Generate Commercial Invoice to Customer", "Customer to Process Payment & Settle Account":
					ds.X += 145
				}
			}
			for _, dfs := range diag.DataFlow_Shapes {
				if dfs.DataFlow != nil && dfs.DataFlow.Name == "Confirm Sales Order to Customer" {
					dfs.StartRatio = 0.5
					dfs.EndRatio = 0.529167
				}
			}

			tsStart := findTS("Start Order Desk")
			tsStart.X, tsStart.Y, tsStart.Width, tsStart.Height = 315, 173, 95, 20
			tsVal := findTS("Receive & Validate Order")
			tsVal.X, tsVal.Y, tsVal.Width, tsVal.Height = 300, 245, 220, 65
			tsConf := findTS("Confirm Sales Order")
			tsConf.X, tsConf.Y, tsConf.Width, tsConf.Height = 300, 500, 220, 65
			tsEnd := findTS("End Order Desk")
			tsEnd.X, tsEnd.Y, tsEnd.Width, tsEnd.Height = 267, 610, 161, 36

			addTaskShape(diag, creditDecision, 385, 345, 50, 50, true)
			addTaskShape(diag, rejectOrder, 555, 337.5, 180, 65, false)
			addTaskShape(diag, orderRejected, 518, 437, 145, 36, false)

			for _, cfs := range diag.ControlFlow_Shapes {
				if cfs.ControlFlow == cfValidate {
					cfs.Name = cfValidate.Name
					cfs.StartRatio = 0.5
					cfs.EndRatio = 0.5
				}
				if cfs.ControlFlow != nil && cfs.ControlFlow.Start == startSales {
					cfs.StartRatio = 1.0
					cfs.EndRatio = 0.5
				}
				if cfs.ControlFlow != nil && cfs.ControlFlow.Start == confirmOrder && cfs.ControlFlow.End == endSales {
					cfs.StartRatio = 0.5
					cfs.EndRatio = 143.0 / 161.0
				}
			}
			addCFShape(diag, cfApproved, models.ORIENTATION_VERTICAL, models.ORIENTATION_VERTICAL, 0.5, 0.5)
			addCFShape(diag, cfRejected, models.ORIENTATION_HORIZONTAL, models.ORIENTATION_HORIZONTAL, 0.5, 0.5)
			addCFShape(diag, cfRejectEnd, models.ORIENTATION_VERTICAL, models.ORIENTATION_VERTICAL, 0.5, 127.0/145.0)
		}
	}

	// Ensure "Order to Cash - Full Process" is the active checked diagram before marshalling
	for _, diag := range stage.GetInstancesSorted[*models.DiagramProcess]() {
		diag.IsChecked = (diag.Name == "Order to Cash - Full Process")
	}
	stage.Commit()

	stage.ComputeReverseMaps()
	stage.ComputeInstancesNb()
	stage.ComputeReferenceAndOrders()
	stage.MarshallFile("data/stage.go", "github.com/fullstack-lang/gong/dsm/process/go/models", "main")
}

func TestStageGoDecisionNodeLayout(t *testing.T) {
	if os.Getenv("UPDATE_STAGE_DECISION_NODE") == "true" {
		updateStageGoWithDecisionNode(t)
	}

	stack := level1stack.NewLevel1StackDelta("verify_stage_decision", "data/stage.go", "", true, true, false)
	stager := models.NewStager(stack.R, stack.Stage, stack.Probe)
	stage := stack.Stage

	var decisionTask, rejectTask, rejectedEndTask *models.Task
	for _, task := range stage.GetInstancesSorted[*models.Task]() {
		switch task.Name {
		case "Credit Approved?":
			decisionTask = task
		case "Reject Sales Order":
			rejectTask = task
		case "Order Rejected":
			rejectedEndTask = task
		}
	}
	if decisionTask == nil || !decisionTask.IsDecisionNode || rejectTask == nil || rejectedEndTask == nil || !rejectedEndTask.IsEndTask {
		t.Fatalf("expected Credit Approved? decision node, Reject Sales Order, and Order Rejected in data/stage.go (decision=%v reject=%v end=%v)",
			decisionTask, rejectTask, rejectedEndTask)
	}

	var foundApprovedGuard, foundRejectedGuard bool
	for _, cf := range stage.GetInstancesSorted[*models.ControlFlow]() {
		if cf.Start == decisionTask && cf.Guard == "approved" {
			foundApprovedGuard = true
		}
		if cf.Start == decisionTask && cf.Guard == "rejected" {
			foundRejectedGuard = true
		}
	}
	if !foundApprovedGuard || !foundRejectedGuard {
		t.Fatalf("expected 'approved' and 'rejected' guards from Credit Approved? in data/stage.go")
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

		layer := stager.GetSvgObject().Layers[0]
		diamondRect := findSvgRectByName(layer, "Credit Approved?")
		if diamondRect == nil {
			t.Fatalf("diagram %q: expected svg.Rect for Credit Approved?", diag.Name)
		}
		if diamondRect.Width != models.DecisionDiamondSize || diamondRect.Height != models.DecisionDiamondSize || len(diamondRect.RectAnchoredPaths) != 1 {
			t.Errorf("diagram %q: expected 50x50 diamond RectAnchoredPath, got %vx%v paths=%d",
				diag.Name, diamondRect.Width, diamondRect.Height, len(diamondRect.RectAnchoredPaths))
		}
		svgOut, _, _ := stager.GetSvgObject().GenerateString()
		if !strings.Contains(svgOut, "[approved]") || !strings.Contains(svgOut, "[rejected]") {
			t.Errorf("diagram %q: expected '[approved]' and '[rejected]' guard labels in SVG", diag.Name)
		}
	}
}

