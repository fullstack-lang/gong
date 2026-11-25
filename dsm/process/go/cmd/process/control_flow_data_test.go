package main

import (
	"math"
	"slices"
	"strings"
	"testing"

	"github.com/fullstack-lang/gong/dsm/process/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/process/go/models"
	form "github.com/fullstack-lang/gong/lib/form/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func TestIssue1552_AllowDataOnControlFlow(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("test_control_flow_data_1552", "", "", true, true, false)
	stager := models.NewStager(stack.R, stack.Stage, stack.Probe)
	stage := stack.Stage

	rootLib := stager.GetRootLibrary()
	rootLib.Name = "RootLib"
	rootLib.IsExpandedTmp = true
	rootLib.IsDatasNodeExpanded = true
	rootLib.NbPixPerCharacter = 8.5

	// 1. Hierarchical Data with Acronym + SVG icon, plus standalone Data
	scoreReport := (&models.Data{
		Name:        "Score Report",
		Description: "Detailed bureau credit score breakdown",
	}).Stage(stage)
	riskTier := (&models.Data{
		Name:        "Risk Tier",
		Description: "Computed risk classification tier",
	}).Stage(stage)
	creditData := (&models.Data{
		Name:        "Credit Assessment",
		Acronym:     "CA",
		Description: "Full credit assessment dossier",
		SVG:         `<svg viewBox="0 0 24 24"><rect x="4" y="3" width="16" height="18" rx="2"/></svg>`,
		SubDatas:    []*models.Data{scoreReport, riskTier},
	}).Stage(stage)
	auditToken := (&models.Data{
		Name:        "Approval Token Reference",
		Description: "Signed approval token passed along the approved branch",
	}).Stage(stage)
	rootLib.RootDatas = append(rootLib.RootDatas, creditData, auditToken)

	// 2. Process, Participant swimlane, Tasks (including a DecisionNode), and ControlFlows
	proc := (&models.Process{Name: "Credit Review Process"}).Stage(stage)
	rootLib.RootProcesses = append(rootLib.RootProcesses, proc)

	lane := (&models.Participant{
		Name:                       "Credit Desk",
		IsTasksNodeExpanded:        true,
		IsControlFlowsNodeExpanded: true,
	}).Stage(stage)
	proc.Participants = append(proc.Participants, lane)

	taskAssess := (&models.Task{Name: "Assess Credit"}).Stage(stage)
	taskDecision := (&models.Task{Name: "Approved?", IsDecisionNode: true}).Stage(stage)
	taskRelease := (&models.Task{Name: "Release Order"}).Stage(stage)
	lane.Tasks = append(lane.Tasks, taskAssess, taskDecision, taskRelease)

	cfVertical := (&models.ControlFlow{
		Name:  `"Assess Credit" to "Approved?"`,
		Start: taskAssess,
		End:   taskDecision,
	}).Stage(stage)
	cfGuarded := (&models.ControlFlow{
		Name:  `"Approved?" to "Release Order"`,
		Start: taskDecision,
		End:   taskRelease,
		Guard: "Yes",
	}).Stage(stage)
	lane.ControlFlows = append(lane.ControlFlows, cfVertical, cfGuarded)

	// 3. Diagram and Shapes
	diag := (&models.DiagramProcess{
		Name:             "CreditView",
		IsChecked:        true,
		IsEditable_:      true,
		DefaultBoxWidth:  200,
		DefaultBoxHeigth: 65,
		Width:            1200,
		Height:           900,
	}).Stage(stage)
	proc.DiagramProcesss = append(proc.DiagramProcesss, diag)

	pShape := (&models.ProcessShape{
		Name:      "Credit Review Process-CreditView",
		Process:   proc,
		RectShape: models.RectShape{X: 100, Y: 50, Width: 700, Height: 700},
	}).Stage(stage)
	diag.Process_Shapes = append(diag.Process_Shapes, pShape)

	laneShape := (&models.ParticipantShape{
		Name:        "Credit Desk-CreditView",
		Participant: lane,
		WidthWeight: 1.0,
	}).Stage(stage)
	diag.Participant_Shapes = append(diag.Participant_Shapes, laneShape)

	tsAssess := (&models.TaskShape{
		Name:      "Assess Credit-CreditView",
		Task:      taskAssess,
		RectShape: models.RectShape{X: 180, Y: 200, Width: 200, Height: 60},
	}).Stage(stage)
	tsDecision := (&models.TaskShape{
		Name:            "Approved?-CreditView",
		Task:            taskDecision,
		IsNameDisplayed: true,
		RectShape:       models.RectShape{X: 256, Y: 360, Width: models.DecisionDiamondSize, Height: models.DecisionDiamondSize},
	}).Stage(stage)
	tsRelease := (&models.TaskShape{
		Name:      "Release Order-CreditView",
		Task:      taskRelease,
		RectShape: models.RectShape{X: 460, Y: 480, Width: 200, Height: 60},
	}).Stage(stage)
	diag.Task_Shapes = append(diag.Task_Shapes, tsAssess, tsDecision, tsRelease)

	cfVertShape := (&models.ControlFlowShape{
		Name:        cfVertical.Name,
		ControlFlow: cfVertical,
		LinkShape: models.LinkShape{
			StartOrientation: models.ORIENTATION_VERTICAL,
			EndOrientation:   models.ORIENTATION_VERTICAL,
			StartRatio:       0.5,
			EndRatio:         0.5,
		},
	}).Stage(stage)
	cfGuardedShape := (&models.ControlFlowShape{
		Name:        cfGuarded.Name,
		ControlFlow: cfGuarded,
		LinkShape: models.LinkShape{
			StartOrientation: models.ORIENTATION_HORIZONTAL,
			EndOrientation:   models.ORIENTATION_VERTICAL,
			StartRatio:       0.5,
			EndRatio:         0.5,
		},
	}).Stage(stage)
	diag.ControlFlow_Shapes = append(diag.ControlFlow_Shapes, cfVertShape, cfGuardedShape)

	stage.Commit()

	// 4. Add Data to ControlFlow via Probe Form & verify auto-creation of DataShape on diagram
	stack.Probe.FillUpFormFromGongstruct(cfVertical, "ControlFlow")
	formStage := stack.Probe.GetFormStage()
	fg := formStage.GetInstancesSorted[*form.FormGroup]()[0]
	setFormAssocSelection(t, stage, fg, "Datas", []*models.Data{creditData})
	formStage.Commit()
	fg.OnSave.OnSave()

	if len(cfVertical.Datas) != 1 || cfVertical.Datas[0] != creditData {
		t.Fatalf("expected cfVertical.Datas to contain Credit Assessment, got %+v", cfVertical.Datas)
	}
	if len(diag.Data_Shapes) != 1 {
		t.Fatalf("expected 1 auto-created DataShape on diagram for cfVertical, got %d", len(diag.Data_Shapes))
	}
	if diag.Data_Shapes[0].Data != creditData || diag.Data_Shapes[0].ControlFlow != cfVertical || diag.Data_Shapes[0].DataFlow != nil {
		t.Fatalf("unexpected DataShape fields: %+v", diag.Data_Shapes[0])
	}

	// Also attach Audit Token to guarded ControlFlow via Probe Form
	stack.Probe.FillUpFormFromGongstruct(cfGuarded, "ControlFlow")
	fgGuarded := formStage.GetInstancesSorted[*form.FormGroup]()[0]
	setFormAssocSelection(t, stage, fgGuarded, "Datas", []*models.Data{auditToken})
	formStage.Commit()
	fgGuarded.OnSave.OnSave()

	if len(diag.Data_Shapes) != 2 {
		t.Fatalf("expected 2 DataShapes across both ControlFlows, got %d", len(diag.Data_Shapes))
	}

	// 5. Verify UX Tree in all locations:
	//    a) Diagram -> Participants -> Credit Desk -> ControlFlows -> cfVertical -> Credit Assessment (+ SubDatas)
	//    b) Diagram -> Participants -> Credit Desk -> Tasks -> Assess Credit -> out control flows -> cfVertical -> Credit Assessment
	//    c) Diagram -> Participants -> Credit Desk -> Tasks -> Approved? -> in control flows -> cfVertical -> Credit Assessment
	//    d) Library -> Data -> Credit Assessment -> cfVertical
	roots := stager.GetTreeStage().GetInstancesSorted[*tree.Tree]()[0].RootNodes

	cfDataNode := findTreeNodeByPath(roots, "RootLib", "Processes", "Credit Review Process", "CreditView", "Participants", "Credit Desk", "ControlFlows", cfVertical.Name, "Credit Assessment")
	if cfDataNode == nil || !cfDataNode.HasCheckboxButton || !cfDataNode.IsChecked {
		t.Fatalf("expected checked Credit Assessment node under Participant -> ControlFlows, got %+v", cfDataNode)
	}
	if got := renderedTreeNodeLabel(cfDataNode); got != "CA / Credit Assessment" {
		t.Fatalf("expected acronym prefix 'CA / Credit Assessment' on ControlFlow data node, got %q", got)
	}

	outTaskCfNode := findTreeNodeByPath(roots, "RootLib", "Processes", "Credit Review Process", "CreditView", "Participants", "Credit Desk", "Tasks", "Assess Credit", "out control flows", cfVertical.Name)
	if outTaskCfNode == nil || outTaskCfNode.OnIsExpandedChange == nil {
		t.Fatalf("expected expandable ControlFlow node under Task -> out control flows")
	}
	outTaskCfNode.OnIsExpandedChange(true)
	if !slices.Contains(diag.ControlFlowsWhoseDataNodeIsExpanded, cfVertical) {
		t.Fatalf("expected cfVertical in ControlFlowsWhoseDataNodeIsExpanded after expanding")
	}

	inTaskDataNode := findTreeNodeByPath(roots, "RootLib", "Processes", "Credit Review Process", "CreditView", "Participants", "Credit Desk", "Tasks", "Approved?", "in control flows", cfVertical.Name, "Credit Assessment")
	if inTaskDataNode == nil || !inTaskDataNode.IsChecked {
		t.Fatalf("expected checked Credit Assessment node under Task -> in control flows -> cfVertical")
	}

	libCfRefNode := findTreeNodeByPath(roots, "RootLib", "Data", "Credit Assessment", cfVertical.Name, "Data", "Credit Assessment")
	if libCfRefNode == nil {
		t.Fatalf("expected reverse ControlFlow navigation under Library -> Data -> Credit Assessment")
	}

	// Check Score Report sub-data individually via checkbox, then bulk display all sub-datas via menu
	scoreNode := findTreeNodeByPath(cfDataNode.Children, "Score Report")
	if scoreNode == nil || scoreNode.IsChecked {
		t.Fatalf("expected unchecked Score Report child node under Credit Assessment")
	}
	scoreNode.OnIsCheckedChanged(true)
	if len(diag.Data_Shapes) != 3 {
		t.Fatalf("expected 3 DataShapes after checking Score Report, got %d", len(diag.Data_Shapes))
	}

	roots = stager.GetTreeStage().GetInstancesSorted[*tree.Tree]()[0].RootNodes
	cfDataNode = findTreeNodeByPath(roots, "RootLib", "Processes", "Credit Review Process", "CreditView", "Participants", "Credit Desk", "ControlFlows", cfVertical.Name, "Credit Assessment")
	displayAllBtn := findMenuButton(cfDataNode, "Display data & sub-datas")
	if displayAllBtn == nil {
		t.Fatalf("expected 'Display data & sub-datas' menu button on hierarchical Data under ControlFlow")
	}
	displayAllBtn.OnClick()
	if len(diag.Data_Shapes) != 4 { // Credit Assessment + Score Report + Risk Tier + Approval Token Reference
		t.Fatalf("expected 4 DataShapes after bulk display, got %d", len(diag.Data_Shapes))
	}

	// 6. Verify SVG rendering, hierarchical stacking/connectors, vertical midpoint placement, and guard coexistence
	svgObj := stager.GetSvgObject()
	layer := svgObj.Layers[0]

	caRect := findSvgRectByName(layer, "Credit Assessment")
	scoreRect := findSvgRectByName(layer, "Score Report")
	riskRect := findSvgRectByName(layer, "Risk Tier")
	tokenRect := findSvgRectByName(layer, "Approval Token Reference")
	assessRect := findSvgRectByName(layer, "Assess Credit")
	decisionRect := findSvgRectByName(layer, "Approved?")
	if caRect == nil || scoreRect == nil || riskRect == nil || tokenRect == nil || assessRect == nil || decisionRect == nil {
		t.Fatalf("expected SVG rects for all checked ControlFlow Datas and Tasks")
	}

	// Acronym rendered on SVG card + custom SVG icon left margin (24.0)
	if len(caRect.RectAnchoredTexts) == 0 || caRect.RectAnchoredTexts[0].Content != "CA" {
		t.Fatalf("expected 'CA' acronym text on Credit Assessment SVG rect, got %+v", caRect.RectAnchoredTexts)
	}
	if caRect.RectAnchoredTexts[0].X_Offset != 24.0 {
		t.Fatalf("expected X_Offset=24.0 for icon on Credit Assessment SVG rect, got %g", caRect.RectAnchoredTexts[0].X_Offset)
	}

	// Pre-order vertical stacking & 16px indentation + L-shaped hierarchy connectors
	if !(caRect.Y < scoreRect.Y && scoreRect.Y < riskRect.Y) {
		t.Fatalf("expected vertical stacking CA(%.1f) < Score(%.1f) < Risk(%.1f)", caRect.Y, scoreRect.Y, riskRect.Y)
	}
	if math.Abs((scoreRect.X-caRect.X)-16.0) > 0.01 || math.Abs((riskRect.X-caRect.X)-16.0) > 0.01 {
		t.Fatalf("expected 16px hierarchy indent on sub-datas, got %.2f / %.2f", scoreRect.X-caRect.X, riskRect.X-caRect.X)
	}
	if !hdProcessHasHierarchyConnector(scoreRect) || !hdProcessHasHierarchyConnector(riskRect) {
		t.Fatalf("expected L-shaped hierarchy connectors with StrokeOpacity=1 on Score Report and Risk Tier")
	}

	// Vertical control flow stack centered around midpoint between Assess Credit bottom (260) and Approved? top (360) -> midpoint Y = 310
	stackCenterY := (caRect.Y + (riskRect.Y + riskRect.Height)) / 2.0
	expectedMidY := ((assessRect.Y + assessRect.Height) + decisionRect.Y) / 2.0
	if math.Abs(stackCenterY-expectedMidY) > 1.5 {
		t.Fatalf("expected vertical ControlFlow data stack centered near Y=%.1f, got %.1f", expectedMidY, stackCenterY)
	}

	// Guarded ControlFlow still renders its "[Yes]" guard at arrow start alongside Approval Token Reference at corner
	var foundGuardText bool
	for _, link := range layer.Links {
		for _, txt := range link.TextAtArrowStart {
			if txt.Content == "[Yes]" {
				foundGuardText = true
			}
		}
	}
	if !foundGuardText {
		t.Fatalf("expected '[Yes]' guard text on guarded ControlFlow link")
	}

	// Select & resize Approval Token Reference on guarded ControlFlow
	initialTokenHeight := tokenRect.Height
	tokenRect.OnSelect()
	if stack.Probe.GetFormStage().GetInstancesSorted[*form.FormGroup]()[0].Label != "Approval Token Reference" {
		t.Fatalf("expected Probe form to show Approval Token Reference on DataShape select")
	}
	tokenRect.OnResize(tokenRect.X, tokenRect.Y, 95.0, tokenRect.Height)
	resizedTokenRect := findSvgRectByName(stager.GetSvgObject().Layers[0], "Approval Token Reference")
	if resizedTokenRect == nil || resizedTokenRect.Width != 95.0 || resizedTokenRect.Height <= initialTokenHeight {
		t.Fatalf("expected resized wrapped DataShape rect on ControlFlow, got %+v", resizedTokenRect)
	}

	// 7. Website page generation for ControlFlow documents carried Data
	cfPage := cfVertical.GeneratePage(stager)
	if cfPage == nil || !strings.Contains(cfPage.MardownContent, "Credit Assessment") {
		t.Fatalf("expected ControlFlow.GeneratePage to list carried Data 'Credit Assessment', got %+v", cfPage)
	}

	// 8. Semantic rules: duplicate DataShape cleanup, invalid dual-owner cleanup, and uncheck ControlFlowShape cleanup
	dupDS := (&models.DataShape{
		Name:        "DupCreditShape",
		Data:        creditData,
		ControlFlow: cfVertical,
	}).Stage(stage)
	invalidDualOwnerDS := (&models.DataShape{
		Name:        "InvalidDualOwner",
		Data:        creditData,
		ControlFlow: cfVertical,
		DataFlow:    (&models.DataFlow{Name: "TempInvalidDF"}).Stage(stage),
	}).Stage(stage)
	diag.Data_Shapes = append(diag.Data_Shapes, dupDS, invalidDualOwnerDS)
	stage.Commit()

	if len(diag.Data_Shapes) != 4 {
		t.Fatalf("expected duplicate and invalid dual-owner DataShapes to be unstaged (leaving 4), got %d", len(diag.Data_Shapes))
	}

	// Unchecking cfVertical shape removes its 3 DataShapes while preserving cfVertical.Datas and disabling data checkboxes
	roots = stager.GetTreeStage().GetInstancesSorted[*tree.Tree]()[0].RootNodes
	cfVertNode := findTreeNodeByPath(roots, "RootLib", "Processes", "Credit Review Process", "CreditView", "Participants", "Credit Desk", "ControlFlows", cfVertical.Name)
	if cfVertNode == nil {
		t.Fatalf("expected cfVertical tree node")
	}
	cfVertNode.OnIsCheckedChanged(false)

	if len(diag.Data_Shapes) != 1 || diag.Data_Shapes[0].ControlFlow != cfGuarded {
		t.Fatalf("expected only cfGuarded DataShape to remain after unchecking cfVertical shape, got %+v", diag.Data_Shapes)
	}
	if len(cfVertical.Datas) != 1 {
		t.Fatalf("expected abstract cfVertical.Datas to remain intact when ControlFlowShape is unchecked")
	}
	roots = stager.GetTreeStage().GetInstancesSorted[*tree.Tree]()[0].RootNodes
	disabledDataNode := findTreeNodeByPath(roots, "RootLib", "Processes", "Credit Review Process", "CreditView", "Participants", "Credit Desk", "ControlFlows", cfVertical.Name, "Credit Assessment")
	if disabledDataNode == nil || !disabledDataNode.IsCheckboxDisabled {
		t.Fatalf("expected Data node checkbox to be disabled when ControlFlowShape is unchecked")
	}

	// Clearing cfGuarded.Datas cleans up its DataShape
	cfGuarded.Datas = nil
	stage.Commit()
	if len(diag.Data_Shapes) != 0 {
		t.Fatalf("expected 0 DataShapes after clearing cfGuarded.Datas, got %d", len(diag.Data_Shapes))
	}
}
