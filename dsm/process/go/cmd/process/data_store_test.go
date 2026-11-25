package main

import (
	"strings"
	"testing"

	"github.com/fullstack-lang/gong/dsm/process/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/process/go/models"
	form "github.com/fullstack-lang/gong/lib/form/go/models"
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func TestIssue1554_DataStoreCylinderFlowsTreeAndSemantics(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("test_data_store_1554", "", "", true, true, false)
	stager := models.NewStager(stack.R, stack.Stage, stack.Probe)
	stage := stack.Stage

	rootLib := stager.GetRootLibrary()
	rootLib.Name = "RootLib"
	rootLib.IsExpandedTmp = true

	// Create hierarchical Data to store in the DataStore and carry on DataFlows
	orderLineData := (&models.Data{Name: "Order Line"}).Stage(stage)
	orderData := (&models.Data{
		Name:     "Customer Order",
		SubDatas: []*models.Data{orderLineData},
	}).Stage(stage)
	rootLib.RootDatas = append(rootLib.RootDatas, orderData)

	// Create Process 1 and Process 2 to verify Library-level DataStore reuse across processes/diagrams
	proc1 := (&models.Process{Name: "Order Intake"}).Stage(stage)
	proc2 := (&models.Process{Name: "Order Fulfillment"}).Stage(stage)
	rootLib.RootProcesses = append(rootLib.RootProcesses, proc1, proc2)

	diag1 := (&models.DiagramProcess{
		Name:             "IntakeView",
		IsChecked:        true,
		IsEditable_:      true,
		DefaultBoxWidth:  200,
		DefaultBoxHeigth: 70,
		Width:            1200,
		Height:           800,
	}).Stage(stage)
	proc1.DiagramProcesss = append(proc1.DiagramProcesss, diag1)

	diag2 := (&models.DiagramProcess{
		Name:             "FulfillmentView",
		IsChecked:        false,
		IsEditable_:      true,
		DefaultBoxWidth:  200,
		DefaultBoxHeigth: 70,
		Width:            1200,
		Height:           800,
	}).Stage(stage)
	proc2.DiagramProcesss = append(proc2.DiagramProcesss, diag2)

	pShape1 := (&models.ProcessShape{
		Name:      "Order Intake-IntakeView",
		Process:   proc1,
		RectShape: models.RectShape{X: 250, Y: 60, Width: 600, Height: 550},
	}).Stage(stage)
	diag1.Process_Shapes = append(diag1.Process_Shapes, pShape1)

	pShape2 := (&models.ProcessShape{
		Name:      "Order Fulfillment-FulfillmentView",
		Process:   proc2,
		RectShape: models.RectShape{X: 250, Y: 60, Width: 600, Height: 550},
	}).Stage(stage)
	diag2.Process_Shapes = append(diag2.Process_Shapes, pShape2)

	lane1 := (&models.Participant{Name: "Sales"}).Stage(stage)
	proc1.Participants = append(proc1.Participants, lane1)
	laneShape1 := (&models.ParticipantShape{
		Name:        "Sales-IntakeView",
		Participant: lane1,
		WidthWeight: 1.0,
	}).Stage(stage)
	diag1.Participant_Shapes = append(diag1.Participant_Shapes, laneShape1)

	taskWrite := (&models.Task{Name: "Persist Order"}).Stage(stage)
	taskRead := (&models.Task{Name: "Validate Stored Order"}).Stage(stage)
	lane1.Tasks = append(lane1.Tasks, taskWrite, taskRead)

	tsWrite := (&models.TaskShape{
		Name:      "Persist Order-IntakeView",
		Task:      taskWrite,
		RectShape: models.RectShape{X: 310, Y: 200, Width: 200, Height: 65},
	}).Stage(stage)
	tsRead := (&models.TaskShape{
		Name:      "Validate Stored Order-IntakeView",
		Task:      taskRead,
		RectShape: models.RectShape{X: 310, Y: 360, Width: 200, Height: 65},
	}).Stage(stage)
	diag1.Task_Shapes = append(diag1.Task_Shapes, tsWrite, tsRead)
	stage.Commit()

	// 1. Add DataStore via DiagramProcess UX Tree "Add Data Store" menu button
	trees := stager.GetTreeStage().GetInstancesSorted[*tree.Tree]()
	diag1Node := findTreeNodeByPath(trees[0].RootNodes, "RootLib", "Processes", "Order Intake", "IntakeView")
	if diag1Node == nil || diag1Node.Menu == nil {
		t.Fatalf("expected IntakeView tree node with menu")
	}
	var addDataStoreBtn *tree.Button
	for _, btn := range diag1Node.Menu.Buttons {
		if btn.Name == "Add Data Store" {
			addDataStoreBtn = btn
			break
		}
	}
	if addDataStoreBtn == nil {
		t.Fatalf("expected 'Add Data Store' button on DiagramProcess menu")
	}
	addDataStoreBtn.OnClick()

	if len(rootLib.RootDataStores) != 1 || len(diag1.DataStore_Shapes) != 1 {
		t.Fatalf("expected 1 RootDataStore and 1 DataStoreShape, got %d / %d", len(rootLib.RootDataStores), len(diag1.DataStore_Shapes))
	}

	ordersDB := rootLib.RootDataStores[0]
	ordersDB.Name = "Orders Database Repository"
	ordersDB.Description = "Central persistent store for customer orders"
	ordersDB.SVG = `<svg viewBox="0 0 24 24"><path d="M4 6h16v12H4z"/></svg>`

	dsShape1 := diag1.DataStore_Shapes[0]
	if dsShape1.Width != models.DataStoreDefaultWidth || dsShape1.Height != models.DataStoreDefaultHeight {
		t.Fatalf("expected default cylinder size %gx%g, got %gx%g",
			models.DataStoreDefaultWidth, models.DataStoreDefaultHeight, dsShape1.Width, dsShape1.Height)
	}

	// Verify free positioning outside the Process pool AND diagram bounds expansion
	dsShape1.SetX(40)
	dsShape1.SetY(260)
	stage.Commit()

	if dsShape1.X != 40 || dsShape1.Y != 260 {
		t.Fatalf("expected DataStoreShape to remain freely positioned outside process pool at (40, 260), got (%g, %g)", dsShape1.X, dsShape1.Y)
	}
	if dsShape1.Name != "Orders Database Repository-IntakeView" {
		t.Fatalf("expected synced DataStoreShape name 'Orders Database Repository-IntakeView', got %q", dsShape1.Name)
	}

	// Also verify free positioning INSIDE the Participant swimlane (and far right to test enforceDiagramSize)
	dsShape1.SetX(560)
	dsShape1.SetY(280)
	stage.Commit()
	if dsShape1.X != 560 || dsShape1.Y != 280 {
		t.Fatalf("expected DataStoreShape to remain freely positioned inside swimlane at (560, 280), got (%g, %g)", dsShape1.X, dsShape1.Y)
	}

	// 2. Associate Data with DataStore via Probe Form & verify UX Tree under Library and Diagram
	stack.Probe.FillUpFormFromGongstruct(ordersDB, "DataStore")
	formStage := stack.Probe.GetFormStage()
	fg := formStage.GetInstancesSorted[*form.FormGroup]()[0]
	setFormAssocSelection(t, stage, fg, "Datas", []*models.Data{orderData})
	formStage.Commit()
	fg.OnSave.OnSave()

	if len(ordersDB.Datas) != 1 || ordersDB.Datas[0] != orderData {
		t.Fatalf("expected ordersDB.Datas to contain orderData, got %+v", ordersDB.Datas)
	}

	trees = stager.GetTreeStage().GetInstancesSorted[*tree.Tree]()
	libStoreDataNode := findTreeNodeByPath(trees[0].RootNodes, "RootLib", "Data Stores", "Orders Database Repository", "Data", "Customer Order", "Order Line")
	if libStoreDataNode == nil {
		t.Fatalf("expected Customer Order -> Order Line under RootLib -> Data Stores -> Orders Database Repository -> Data")
	}

	// Reuse the SAME DataStore in Process 2 (FulfillmentView) by checking its checkbox in diag2's tree
	dsInDiag2Node := findTreeNodeByPath(trees[0].RootNodes, "RootLib", "Processes", "Order Fulfillment", "FulfillmentView", "Data Stores", "Orders Database Repository")
	if dsInDiag2Node == nil || dsInDiag2Node.IsChecked {
		t.Fatalf("expected unchecked Orders Database Repository node under FulfillmentView -> Data Stores, got %+v", dsInDiag2Node)
	}
	dsInDiag2Node.OnIsCheckedChanged(true)
	if len(diag2.DataStore_Shapes) != 1 || diag2.DataStore_Shapes[0].DataStore != ordersDB {
		t.Fatalf("expected FulfillmentView to reuse Orders Database Repository via DataStoreShape, got %+v", diag2.DataStore_Shapes)
	}

	// 3. Verify BPMN 3D Cylinder SVG rendering (cylinder body + drum rim + wrapped label + icon + stroke opacity)
	svgObj := stager.GetSvgObject()
	var dbRect, writeRect, readRect *svg.Rect
	for _, r := range svgObj.Layers[0].Rects {
		switch r.Name {
		case "Orders Database Repository":
			dbRect = r
		case "Persist Order":
			writeRect = r
		case "Validate Stored Order":
			readRect = r
		}
	}
	if dbRect == nil || writeRect == nil || readRect == nil {
		t.Fatalf("expected dbRect, writeRect, and readRect in SVG, got db=%v write=%v read=%v", dbRect, writeRect, readRect)
	}
	if dbRect.FillOpacity != 0.0 || dbRect.StrokeOpacity != 0.0 {
		t.Fatalf("expected underlying hit-box rect to be transparent (0, 0), got fillOpacity=%g strokeOpacity=%g", dbRect.FillOpacity, dbRect.StrokeOpacity)
	}

	var foundCylinderBody, foundCylinderRim, foundIcon bool
	for _, p := range dbRect.RectAnchoredPaths {
		switch {
		case strings.HasSuffix(p.Name, "_cylinder_body"):
			foundCylinderBody = true
			if p.FillOpacity != 1.0 || p.StrokeOpacity != 1.0 {
				t.Fatalf("expected cylinder body FillOpacity=1 and StrokeOpacity=1, got %g / %g", p.FillOpacity, p.StrokeOpacity)
			}
		case strings.HasSuffix(p.Name, "_cylinder_rim"):
			foundCylinderRim = true
			if p.StrokeOpacity != 1.0 || p.Color != "none" {
				t.Fatalf("expected cylinder rim StrokeOpacity=1 and Color='none', got %g / %q", p.StrokeOpacity, p.Color)
			}
		case strings.Contains(p.Name, "_icon_"):
			foundIcon = true
		}
	}
	if !foundCylinderBody || !foundCylinderRim || !foundIcon {
		t.Fatalf("expected cylinder body, rim, and custom SVG icon paths, got body=%v rim=%v icon=%v", foundCylinderBody, foundCylinderRim, foundIcon)
	}
	if len(dbRect.RectAnchoredTexts) != 1 || !strings.Contains(dbRect.RectAnchoredTexts[0].Content, "\n") {
		t.Fatalf("expected wrapped title inside cylinder body, got %+v", dbRect.RectAnchoredTexts)
	}

	// 4. Interactive SVG Link Drawing: Task -> DataStore (write) AND DataStore -> Task (read)
	frontWrite := &svg.SVG{
		DrawingState: svg.NOT_DRAWING_LINK,
		StartRect:    writeRect,
		EndRect:      dbRect,
	}
	svgObj.DrawingState = svg.DRAWING_LINK
	svgObj.OnUpdate(frontWrite)

	svgObj = stager.GetSvgObject()
	dbRect, readRect = nil, nil
	for _, r := range svgObj.Layers[0].Rects {
		if r.Name == "Orders Database Repository" {
			dbRect = r
		}
		if r.Name == "Validate Stored Order" {
			readRect = r
		}
	}
	frontRead := &svg.SVG{
		DrawingState: svg.NOT_DRAWING_LINK,
		StartRect:    dbRect,
		EndRect:      readRect,
	}
	svgObj.DrawingState = svg.DRAWING_LINK
	svgObj.OnUpdate(frontRead)

	dataFlows := stage.GetInstancesSorted[*models.DataFlow]()
	if len(dataFlows) != 2 {
		t.Fatalf("expected 2 DataFlows (Task->DataStore and DataStore->Task), got %d", len(dataFlows))
	}
	var writeDF, readDF *models.DataFlow
	for _, df := range dataFlows {
		if df.Type == models.DataFlow_Task2DataStore && df.StartTask == taskWrite && df.EndDataStore == ordersDB {
			writeDF = df
		}
		if df.Type == models.DataFlow_DataStore2Task && df.StartDataStore == ordersDB && df.EndTask == taskRead {
			readDF = df
		}
	}
	if writeDF == nil || readDF == nil {
		t.Fatalf("expected both Task2DataStore and DataStore2Task DataFlows, got write=%v read=%v", writeDF, readDF)
	}
	if len(stager.GetSvgObject().Layers[0].Links) != 2 {
		t.Fatalf("expected 2 rendered svg.Links for DataStore flows, got %d", len(stager.GetSvgObject().Layers[0].Links))
	}

	// Attach Customer Order Data to writeDF via Probe Form and verify DataShape is auto-created on the diagram
	stack.Probe.FillUpFormFromGongstruct(writeDF, "DataFlow")
	fgDF := formStage.GetInstancesSorted[*form.FormGroup]()[0]
	setFormAssocSelection(t, stage, fgDF, "Datas", []*models.Data{orderData})
	formStage.Commit()
	fgDF.OnSave.OnSave()

	if len(diag1.Data_Shapes) != 1 || diag1.Data_Shapes[0].Data != orderData || diag1.Data_Shapes[0].DataFlow != writeDF {
		t.Fatalf("expected DataShape for Customer Order on writeDF, got %+v", diag1.Data_Shapes)
	}

	// Verify Diagram -> Data Stores -> Orders Database Repository -> in/out data flows tree nodes
	trees = stager.GetTreeStage().GetInstancesSorted[*tree.Tree]()
	inFlowNode := findTreeNodeByPath(trees[0].RootNodes, "RootLib", "Processes", "Order Intake", "IntakeView", "Data Stores", "Orders Database Repository", "in data flows", writeDF.Name)
	outFlowNode := findTreeNodeByPath(trees[0].RootNodes, "RootLib", "Processes", "Order Intake", "IntakeView", "Data Stores", "Orders Database Repository", "out data flows", readDF.Name)
	if inFlowNode == nil || !inFlowNode.IsChecked || outFlowNode == nil || !outFlowNode.IsChecked {
		t.Fatalf("expected checked in/out data flow nodes under DataStore in diagram tree, got in=%v out=%v", inFlowNode, outFlowNode)
	}

	// 5. Semantic rules: orphan DataStore reattachment, orphan/duplicate DataStoreShape cleanup, min dimensions, and website page generation
	orphanStore := (&models.DataStore{Name: "Orphan Store"}).Stage(stage)
	_ = (&models.DataStoreShape{Name: "OrphanShape", DataStore: ordersDB}).Stage(stage)
	dupShape := (&models.DataStoreShape{Name: "DupShape", DataStore: ordersDB, RectShape: models.RectShape{Width: 10, Height: 10}}).Stage(stage)
	diag1.DataStore_Shapes = append(diag1.DataStore_Shapes, dupShape)
	dsShape1.SetWidth(10)
	dsShape1.SetHeight(10)

	// Invalid DataFlow missing StartDataStore should be unstaged
	badDF := (&models.DataFlow{
		Name:    "InvalidStoreFlow",
		Type:    models.DataFlow_DataStore2Task,
		EndTask: taskRead,
	}).Stage(stage)
	rootLib.RootDataFlows = append(rootLib.RootDataFlows, badDF)

	stage.Commit()

	if orphanStore.GetOwningLibrary() != rootLib {
		t.Fatalf("expected orphan DataStore to be reattached to rootLib")
	}
	if len(diag1.DataStore_Shapes) != 1 {
		t.Fatalf("expected duplicate DataStoreShape in diag1 to be removed, got %d", len(diag1.DataStore_Shapes))
	}
	if len(stage.GetInstancesSorted[*models.DataStoreShape]()) != 2 { // 1 in diag1 + 1 in diag2
		t.Fatalf("expected orphan and duplicate DataStoreShapes to be unstaged (leaving 2), got %d", len(stage.GetInstancesSorted[*models.DataStoreShape]()))
	}
	if dsShape1.Width != models.DataStoreMinWidth || dsShape1.Height != models.DataStoreMinHeight {
		t.Fatalf("expected min dimensions %gx%g enforced, got %gx%g",
			models.DataStoreMinWidth, models.DataStoreMinHeight, dsShape1.Width, dsShape1.Height)
	}
	if len(stage.GetInstancesSorted[*models.DataFlow]()) != 2 {
		t.Fatalf("expected invalid DataFlow to be unstaged, leaving 2 valid DataFlows, got %d", len(stage.GetInstancesSorted[*models.DataFlow]()))
	}

	page := ordersDB.GeneratePage(stager)
	if page == nil || !strings.Contains(page.MardownContent, "Customer Order") {
		t.Fatalf("expected DataStore.GeneratePage to document stored Customer Order, got %+v", page)
	}

	// Unchecking DataStore in IntakeView removes its DataStoreShape AND unstages dependent DataFlowShapes on IntakeView while keeping abstract DataFlows
	trees = stager.GetTreeStage().GetInstancesSorted[*tree.Tree]()
	dsInDiag1Node := findTreeNodeByPath(trees[0].RootNodes, "RootLib", "Processes", "Order Intake", "IntakeView", "Data Stores", "Orders Database Repository")
	if dsInDiag1Node == nil {
		t.Fatalf("expected Orders Database Repository node in IntakeView")
	}
	dsInDiag1Node.OnIsCheckedChanged(false)

	if len(diag1.DataStore_Shapes) != 0 || len(diag1.DataFlow_Shapes) != 0 {
		t.Fatalf("expected DataStore_Shapes and dependent DataFlow_Shapes in diag1 to be cleared when DataStore is unchecked, got ds=%d df=%d",
			len(diag1.DataStore_Shapes), len(diag1.DataFlow_Shapes))
	}
	if len(stage.GetInstancesSorted[*models.DataFlow]()) != 2 {
		t.Fatalf("expected abstract DataFlows to remain staged when shape is unchecked from diagram")
	}
}

// Regression test: when a Task already has a DataFlowShape to ExternalParticipant #0 (GongGetOrder == 0)
// and/or Task #0 (GongGetOrder == 0), adding a DataFlow between that Task and DataStore #0 (GongGetOrder == 0)
// — either by dragging on the SVG canvas or by checking the DataFlow checkbox in the UX tree —
// must NOT collide in enforceRelationDuplicates and must create and retain the DataFlowShape.
func TestIssue1554_DataFlowShapeNoOrderCollisionWithExternalParticipantOrSameTask(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("test_ds_no_collision", "", "", true, true, false)
	stager := models.NewStager(stack.R, stack.Stage, stack.Probe)
	stage := stack.Stage

	rootLib := stager.GetRootLibrary()
	rootLib.Name = "RootLib"

	proc := (&models.Process{Name: "Order to Cash"}).Stage(stage)
	rootLib.RootProcesses = append(rootLib.RootProcesses, proc)

	// ExternalParticipant #0 -> GongGetOrder == 0
	customer := (&models.ExternalParticipant{Name: "Customer"}).Stage(stage)
	proc.ExternalParticipants = append(proc.ExternalParticipants, customer)

	// DataStore #0 -> GongGetOrder == 0
	dataStore := (&models.DataStore{Name: "Data Store"}).Stage(stage)
	rootLib.RootDataStores = append(rootLib.RootDataStores, dataStore)

	sales := (&models.Participant{Name: "Sales"}).Stage(stage)
	proc.Participants = append(proc.Participants, sales)

	// Task #0 -> GongGetOrder == 0
	confirmOrder := (&models.Task{Name: "Confirm Sales Order"}).Stage(stage)
	sales.Tasks = append(sales.Tasks, confirmOrder)

	ackData := (&models.Data{Name: "Order Confirmation"}).Stage(stage)
	rootLib.RootDatas = append(rootLib.RootDatas, ackData)

	// Pre-existing DataFlow from Confirm Sales Order (Task #0) -> Customer (ExternalParticipant #0)
	dfToCustomer := (&models.DataFlow{
		Name:           "Confirm Sales Order to Customer",
		Type:           models.DataFlow_Task2ExternalParticipant,
		StartTask:      confirmOrder,
		EndParticipant: customer,
		Datas:          []*models.Data{ackData},
	}).Stage(stage)
	rootLib.RootDataFlows = append(rootLib.RootDataFlows, dfToCustomer)
	proc.DataFlows = append(proc.DataFlows, dfToCustomer)

	diag := (&models.DiagramProcess{
		Name:             "Full Process",
		IsChecked:        true,
		IsEditable_:      true,
		DefaultBoxWidth:  220,
		DefaultBoxHeigth: 65,
		Width:            1200,
		Height:           800,
	}).Stage(stage)
	proc.DiagramProcesss = append(proc.DiagramProcesss, diag)

	diag.Process_Shapes = append(diag.Process_Shapes, (&models.ProcessShape{
		Name:      "Order to Cash-Full Process",
		Process:   proc,
		RectShape: models.RectShape{X: 250, Y: 50, Width: 600, Height: 600},
	}).Stage(stage))
	diag.Participant_Shapes = append(diag.Participant_Shapes, (&models.ParticipantShape{
		Name:        "Sales-Full Process",
		Participant: sales,
		WidthWeight: 1.0,
	}).Stage(stage))
	diag.ExternalParticipant_Shapes = append(diag.ExternalParticipant_Shapes, (&models.ExternalParticipantShape{
		Name:                "Customer-Full Process",
		ExternalParticipant: customer,
		RectShape:           models.RectShape{X: 40, Y: 150, Width: 160, Height: 65},
		TailHeigth:          400,
	}).Stage(stage))
	diag.Task_Shapes = append(diag.Task_Shapes, (&models.TaskShape{
		Name:      "Confirm Sales Order-Full Process",
		Task:      confirmOrder,
		RectShape: models.RectShape{X: 300, Y: 260, Width: 220, Height: 65},
	}).Stage(stage))
	diag.DataStore_Shapes = append(diag.DataStore_Shapes, (&models.DataStoreShape{
		Name:      "Data Store-Full Process",
		DataStore: dataStore,
		RectShape: models.RectShape{X: 500, Y: 430, Width: 140, Height: 80},
	}).Stage(stage))
	diag.DataFlow_Shapes = append(diag.DataFlow_Shapes, (&models.DataFlowShape{
		Name:     dfToCustomer.Name,
		DataFlow: dfToCustomer,
		LinkShape: models.LinkShape{
			StartOrientation: models.ORIENTATION_HORIZONTAL,
			EndOrientation:   models.ORIENTATION_HORIZONTAL,
			StartRatio:       0.5,
			EndRatio:         0.5,
		},
	}).Stage(stage))
	stage.Commit()

	// 1. Drag link from Confirm Sales Order (Task #0) -> Data Store (DataStore #0) on SVG canvas
	svgObj := stager.GetSvgObject()
	var taskRect, dsRect *svg.Rect
	for _, r := range svgObj.Layers[0].Rects {
		if r.Name == "Confirm Sales Order" {
			taskRect = r
		}
		if r.Name == "Data Store" {
			dsRect = r
		}
	}
	if taskRect == nil || dsRect == nil {
		t.Fatalf("expected taskRect and dsRect in SVG")
	}

	svgObj.DrawingState = svg.DRAWING_LINK
	svgObj.OnUpdate(&svg.SVG{
		DrawingState: svg.NOT_DRAWING_LINK,
		StartRect:    taskRect,
		EndRect:      dsRect,
	})

	if len(diag.DataFlow_Shapes) != 2 {
		t.Fatalf("expected BOTH Confirm->Customer and Confirm->DataStore DataFlowShapes to remain staged, got %d", len(diag.DataFlow_Shapes))
	}
	if len(stager.GetSvgObject().Layers[0].Links) != 2 {
		t.Fatalf("expected 2 rendered svg.Links, got %d", len(stager.GetSvgObject().Layers[0].Links))
	}

	// Also drag reverse read link Data Store (DataStore #0) -> Confirm Sales Order (Task #0)
	svgObj = stager.GetSvgObject()
	taskRect, dsRect = nil, nil
	for _, r := range svgObj.Layers[0].Rects {
		if r.Name == "Confirm Sales Order" {
			taskRect = r
		}
		if r.Name == "Data Store" {
			dsRect = r
		}
	}
	svgObj.DrawingState = svg.DRAWING_LINK
	svgObj.OnUpdate(&svg.SVG{
		DrawingState: svg.NOT_DRAWING_LINK,
		StartRect:    dsRect,
		EndRect:      taskRect,
	})
	if len(diag.DataFlow_Shapes) != 3 {
		t.Fatalf("expected all 3 DataFlowShapes (Task->Ext, Task->DS, DS->Task) to remain staged, got %d", len(diag.DataFlow_Shapes))
	}

	// 2. Uncheck "Confirm Sales Order to Data Store" under Data Store -> in data flows in the tree, then re-check it
	trees := stager.GetTreeStage().GetInstancesSorted[*tree.Tree]()
	inFlowNode := findTreeNodeByPath(trees[0].RootNodes, "RootLib", "Processes", "Order to Cash", "Full Process", "Data Stores", "Data Store", "in data flows", "Confirm Sales Order to Data Store")
	if inFlowNode == nil || !inFlowNode.IsChecked {
		t.Fatalf("expected checked 'Confirm Sales Order to Data Store' node under Data Store -> in data flows, got %+v", inFlowNode)
	}

	inFlowNode.OnIsCheckedChanged(false)
	if len(diag.DataFlow_Shapes) != 2 {
		t.Fatalf("expected 2 DataFlowShapes after unchecking, got %d", len(diag.DataFlow_Shapes))
	}

	trees = stager.GetTreeStage().GetInstancesSorted[*tree.Tree]()
	inFlowNode = findTreeNodeByPath(trees[0].RootNodes, "RootLib", "Processes", "Order to Cash", "Full Process", "Data Stores", "Data Store", "in data flows", "Confirm Sales Order to Data Store")
	if inFlowNode == nil || inFlowNode.IsChecked {
		t.Fatalf("expected unchecked 'Confirm Sales Order to Data Store' node after unchecking")
	}

	inFlowNode.OnIsCheckedChanged(true)
	if len(diag.DataFlow_Shapes) != 3 {
		t.Fatalf("expected 3 DataFlowShapes after re-checking in tree, got %d", len(diag.DataFlow_Shapes))
	}
	trees = stager.GetTreeStage().GetInstancesSorted[*tree.Tree]()
	inFlowNode = findTreeNodeByPath(trees[0].RootNodes, "RootLib", "Processes", "Order to Cash", "Full Process", "Data Stores", "Data Store", "in data flows", "Confirm Sales Order to Data Store")
	if inFlowNode == nil || !inFlowNode.IsChecked {
		t.Fatalf("expected 'Confirm Sales Order to Data Store' checkbox to stay checked after OnIsCheckedChanged(true)")
	}
}
