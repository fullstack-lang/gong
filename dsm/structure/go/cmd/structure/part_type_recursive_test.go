package main

import (
	"slices"
	"testing"

	"github.com/fullstack-lang/gong/dsm/structure/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/structure/go/models"
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func findDirectChildByName(parent *tree.Node, name string) *tree.Node {
	if parent == nil {
		return nil
	}
	for _, c := range parent.Children {
		if c.Name == name {
			return c
		}
	}
	return nil
}

func TestPartTypeRecursiveDisplayMultiInstanceAndSemantics(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("test_part_type_recursive", "", "", true, true, false)
	stager := models.NewStager(
		stack.R,
		stack.Stage,
		stack.Probe,
	)
	stage := stack.Stage
	rootLib := stager.GetRootLibrary()

	// Build 3-level System hierarchy:
	// SysRoot
	//  |- inst1 (Part, Name="Inst1", Type=SysSub) -> displays as "Inst1 : SysSub"
	//  |- inst2 (Part, Name="Inst2", Type=SysSub) -> displays as "Inst2 : SysSub"
	//  |- inst3 (Part, Name="Part 1", Type=SysLeaf) -> keeps Name="Part 1", displays as "Part 1 : SysLeaf"
	//
	// SysSub:
	//  |- subB1 (Part, Name="SubB1", Type=SysLeaf) -> displays as "SubB1 : SysLeaf"
	//  |    |- Ports: b1p1, b1p2
	//  |    |- ControlFlows: cfB1 (b1p1 -> b1p2)
	//  |- subB2 (Part, Name="SubB2")
	//  |    |- Ports: b2p1
	//  |- DataFlows: dfInternal (b1p2 -> b2p1, carrying Data "Payload" & Protocol "TCP")
	//
	// SysLeaf:
	//  |- leafC1 (Part, Name="LeafC1")

	sysLeaf := (&models.System{Name: "SysLeaf"}).Stage(stage)
	leafC1 := (&models.Part{Name: "LeafC1"}).Stage(stage)
	sysLeaf.Parts = append(sysLeaf.Parts, leafC1)

	sysSub := (&models.System{Name: "SysSub"}).Stage(stage)
	subB1 := (&models.Part{
		Name: "SubB1",
		Type: sysLeaf,
	}).Stage(stage)
	b1p1 := (&models.Port{Name: "b1p1"}).Stage(stage)
	b1p2 := (&models.Port{Name: "b1p2"}).Stage(stage)
	subB1.Ports = append(subB1.Ports, b1p1, b1p2)
	cfB1 := (&models.ControlFlow{
		Name:  "cfB1",
		Start: b1p1,
		End:   b1p2,
	}).Stage(stage)
	subB1.ControlFlows = append(subB1.ControlFlows, cfB1)

	subB2 := (&models.Part{Name: "SubB2"}).Stage(stage)
	b2p1 := (&models.Port{Name: "b2p1"}).Stage(stage)
	subB2.Ports = append(subB2.Ports, b2p1)
	sysSub.Parts = append(sysSub.Parts, subB1, subB2)

	payload := (&models.Data{Name: "Payload"}).Stage(stage)
	rootLib.RootDatas = append(rootLib.RootDatas, payload)
	protoTCP := (&models.Protocol{Name: "TCP"}).Stage(stage)
	rootLib.RootProtocols = append(rootLib.RootProtocols, protoTCP)

	dfInternal := (&models.DataFlow{
		Name:      "dfInternal",
		Type:      models.DataFlow_Port2Port,
		StartPort: b1p2,
		EndPort:   b2p1,
		Datas:     []*models.Data{payload},
		Protocol:  protoTCP,
	}).Stage(stage)
	rootLib.RootDataFlows = append(rootLib.RootDataFlows, dfInternal)
	sysSub.DataFlows = append(sysSub.DataFlows, dfInternal)

	sysRoot := (&models.System{Name: "SysRoot"}).Stage(stage)
	inst1 := (&models.Part{
		Name: "Inst1",
		Type: sysSub,
	}).Stage(stage)
	inst2 := (&models.Part{
		Name: "Inst2",
		Type: sysSub,
	}).Stage(stage)
	inst3 := (&models.Part{
		Name: "Part 1",
		Type: sysLeaf,
	}).Stage(stage)
	sysRoot.Parts = append(sysRoot.Parts, inst1, inst2, inst3)

	rootLib.RootSystemes = append(rootLib.RootSystemes, sysRoot, sysSub, sysLeaf)

	diag := (&models.DiagramStructure{
		Name:             "RootDiag",
		IsChecked:        true,
		IsEditable_:      true,
		DefaultBoxWidth:  260,
		DefaultBoxHeigth: 70,
		Width:            1400,
		Height:           1000,
	}).Stage(stage)
	sysRoot.DiagramStructures = append(sysRoot.DiagramStructures, diag)

	sysRootShape := (&models.SystemShape{
		Name:      "SysRoot-RootDiag",
		System:    sysRoot,
		RectShape: models.RectShape{X: 40, Y: 40, Width: 1200, Height: 850},
	}).Stage(stage)
	diag.System_Shapes = append(diag.System_Shapes, sysRootShape)

	stage.Commit()

	// 1. Verify Part.Name is NEVER overwritten by Part.Type.Name
	if inst1.Name != "Inst1" || inst2.Name != "Inst2" || inst3.Name != "Part 1" || subB1.Name != "SubB1" {
		t.Fatalf("expected Part names preserved (not overwritten by Type), got %q / %q / %q / %q",
			inst1.Name, inst2.Name, inst3.Name, subB1.Name)
	}

	// Helper to locate RootDiag -> Parts -> "<partName> : <typeName>" in current UX tree
	getInstNode := func(partLabel string) *tree.Node {
		t.Helper()
		trees := stager.GetTreeStage().GetInstancesSorted[*tree.Tree]()
		if len(trees) == 0 {
			t.Fatalf("expected UX tree")
		}
		diagNode := findTreeNodeByName(trees[0].RootNodes, "RootDiag")
		if diagNode == nil {
			t.Fatalf("expected RootDiag node in UX tree")
		}
		partsFolder := findDirectChildByName(diagNode, "Parts")
		if partsFolder == nil {
			t.Fatalf("expected Parts folder under RootDiag")
		}
		n := findDirectChildByName(partsFolder, partLabel)
		if n == nil {
			t.Fatalf("expected part node %q under RootDiag/Parts", partLabel)
		}
		return n
	}

	// Verify "Part 1 : SysLeaf" formatting in UX tree (and raw "Part 1" when in rename mode)
	_ = getInstNode("Part 1 : SysLeaf")
	inst3.SetIsInRenameMode(true)
	stage.Commit()
	_ = getInstNode("Part 1")
	inst3.SetIsInRenameMode(false)
	stage.Commit()
	_ = getInstNode("Part 1 : SysLeaf")

	// 2. Before Inst1 is checked, nested SubB1 checkbox must be disabled
	inst1Node := getInstNode("Inst1 : SysSub")
	inst1TypeParts := findDirectChildByName(inst1Node, "Parts")
	if inst1TypeParts == nil {
		t.Fatalf("expected recursive 'Parts' folder under typed Part Inst1")
	}
	subB1NodeUncheckedParent := findDirectChildByName(inst1TypeParts, "SubB1 : SysLeaf")
	if subB1NodeUncheckedParent == nil || !subB1NodeUncheckedParent.IsCheckboxDisabled {
		t.Fatalf("expected SubB1 node under unchecked Inst1 to have disabled checkbox, got %+v", subB1NodeUncheckedParent)
	}

	// 3. Check Inst1 in diagram -> leaf-style title "Inst1 : SysSub" centered initially
	inst1Node.OnIsCheckedChanged(true)

	svgObj := stager.GetSvgObject()
	var inst1Rect *svg.Rect
	for _, r := range svgObj.Layers[0].Rects {
		if r.Name == "Inst1-RootDiag" {
			inst1Rect = r
			break
		}
	}
	if inst1Rect == nil {
		t.Fatalf("expected SVG rect Inst1-RootDiag after checking Inst1")
	}
	if len(inst1Rect.RectAnchoredTexts) == 0 || inst1Rect.RectAnchoredTexts[0].RectAnchorType != svg.RECT_CENTER_MIDDLE {
		t.Fatalf("expected Inst1 title anchored at RECT_CENTER_MIDDLE before nested parts are checked, got %+v", inst1Rect.RectAnchoredTexts)
	}
	if inst1Rect.RectAnchoredTexts[0].Content != "Inst1 : SysSub" {
		t.Fatalf("expected SVG title %q on Inst1, got %q", "Inst1 : SysSub", inst1Rect.RectAnchoredTexts[0].Content)
	}

	// 4. Expand & check nested SubB1, SubB2, LeafC1 (3rd level!), ports, control flow, and internal data flow under Inst1
	checkSubHierarchy := func(instName string, instLabel string, xOffset float64) {
		t.Helper()
		// Position top-level instance so Inst1 and Inst2 sit side-by-side cleanly
		for _, ps := range diag.Part_Shapes {
			if ps.ParentPartShape == nil && ps.Part != nil && ps.Part.Name == instName {
				ps.X = xOffset
				ps.Y = 100
			}
		}
		stage.Commit()

		instNode := getInstNode(instLabel)
		subPartsFolder := findDirectChildByName(instNode, "Parts")
		subB1Node := findDirectChildByName(subPartsFolder, "SubB1 : SysLeaf")
		if subB1Node == nil || subB1Node.IsCheckboxDisabled {
			t.Fatalf("expected enabled SubB1 checkbox under checked %s", instName)
		}
		subB1Node.OnIsCheckedChanged(true)

		instNode = getInstNode(instLabel)
		subPartsFolder = findDirectChildByName(instNode, "Parts")
		subB2Node := findDirectChildByName(subPartsFolder, "SubB2")
		subB2Node.OnIsCheckedChanged(true)

		// Check 3rd-level LeafC1 inside SubB1
		instNode = getInstNode(instLabel)
		subB1Node = findDirectChildByName(findDirectChildByName(instNode, "Parts"), "SubB1 : SysLeaf")
		leafPartsFolder := findDirectChildByName(subB1Node, "Parts")
		if leafPartsFolder == nil {
			t.Fatalf("expected 3rd-level 'Parts' folder under SubB1")
		}
		leafC1Node := findDirectChildByName(leafPartsFolder, "LeafC1")
		if leafC1Node == nil || leafC1Node.IsCheckboxDisabled {
			t.Fatalf("expected enabled 3rd-level LeafC1 checkbox under checked SubB1")
		}
		leafC1Node.OnIsCheckedChanged(true)

		// Check Ports on SubB1 and SubB2
		instNode = getInstNode(instLabel)
		subB1Node = findDirectChildByName(findDirectChildByName(instNode, "Parts"), "SubB1 : SysLeaf")
		b1p1Node := findDirectChildByName(findDirectChildByName(subB1Node, "Ports"), "b1p1")
		b1p1Node.OnIsCheckedChanged(true)

		instNode = getInstNode(instLabel)
		subB1Node = findDirectChildByName(findDirectChildByName(instNode, "Parts"), "SubB1 : SysLeaf")
		b1p2Node := findDirectChildByName(findDirectChildByName(subB1Node, "Ports"), "b1p2")
		b1p2Node.OnIsCheckedChanged(true)

		instNode = getInstNode(instLabel)
		subB2Node = findDirectChildByName(findDirectChildByName(instNode, "Parts"), "SubB2")
		b2p1Node := findDirectChildByName(findDirectChildByName(subB2Node, "Ports"), "b2p1")
		b2p1Node.OnIsCheckedChanged(true)

		// Check ControlFlow cfB1 inside SubB1
		instNode = getInstNode(instLabel)
		subB1Node = findDirectChildByName(findDirectChildByName(instNode, "Parts"), "SubB1 : SysLeaf")
		cfNode := findDirectChildByName(findDirectChildByName(subB1Node, "ControlFlows"), cfB1.GetName())
		if cfNode == nil || cfNode.IsCheckboxDisabled {
			t.Fatalf("expected enabled ControlFlow %q checkbox under %s/SubB1, got %+v", cfB1.GetName(), instName, cfNode)
		}
		cfNode.OnIsCheckedChanged(true)

		// Check internal DataFlow dfInternal + Data Payload + Protocol TCP under <instLabel>/Data Flows
		instNode = getInstNode(instLabel)
		dfFolder := findDirectChildByName(instNode, "Data Flows")
		if dfFolder == nil {
			t.Fatalf("expected 'Data Flows' folder under typed Part %s", instName)
		}
		dfNode := findDirectChildByName(dfFolder, dfInternal.GetName())
		if dfNode == nil || dfNode.IsCheckboxDisabled {
			t.Fatalf("expected enabled internal DataFlow %q checkbox under %s", dfInternal.GetName(), instName)
		}
		dfNode.OnIsCheckedChanged(true)

		instNode = getInstNode(instLabel)
		dfNode = findDirectChildByName(findDirectChildByName(instNode, "Data Flows"), dfInternal.GetName())
		payloadNode := findDirectChildByName(dfNode, "Payload")
		payloadNode.OnIsCheckedChanged(true)

		instNode = getInstNode(instLabel)
		dfNode = findDirectChildByName(findDirectChildByName(instNode, "Data Flows"), dfInternal.GetName())
		tcpNode := findDirectChildByName(dfNode, "TCP")
		tcpNode.OnIsCheckedChanged(true)
	}

	checkSubHierarchy("Inst1", "Inst1 : SysSub", 80)

	// 5. Check Inst2 (second instance of the exact same SysSub in the same diagram!) and expand its full hierarchy too
	inst2Node := getInstNode("Inst2 : SysSub")
	inst2Node.OnIsCheckedChanged(true)
	checkSubHierarchy("Inst2", "Inst2 : SysSub", 620)

	// Verify multi-instance counts in RootDiag:
	// Part_Shapes: 2 * (Inst + SubB1 + SubB2 + LeafC1) = 8
	// Port_Shapes: 2 * (b1p1 + b1p2 + b2p1) = 6
	// ControlFlow_Shapes: 2 * (cfB1) = 2
	// DataFlow_Shapes: 2 * (dfInternal) = 2
	if len(diag.Part_Shapes) != 8 {
		t.Fatalf("expected 8 Part_Shapes across 2 recursive instances, got %d", len(diag.Part_Shapes))
	}
	if len(diag.Port_Shapes) != 6 {
		t.Fatalf("expected 6 Port_Shapes across 2 recursive instances, got %d", len(diag.Port_Shapes))
	}
	if len(diag.ControlFlow_Shapes) != 2 {
		t.Fatalf("expected 2 ControlFlow_Shapes across 2 recursive instances, got %d", len(diag.ControlFlow_Shapes))
	}
	if len(diag.DataFlow_Shapes) != 2 {
		t.Fatalf("expected 2 DataFlow_Shapes across 2 recursive instances, got %d", len(diag.DataFlow_Shapes))
	}

	// Verify SVG containment, Z-order, top titles vs leaf centered titles, Peers, and sibling-only Obstacles
	svgObj = stager.GetSvgObject()
	rectByName := make(map[string]*svg.Rect)
	rectIndex := make(map[string]int)
	for idx, r := range svgObj.Layers[0].Rects {
		rectByName[r.Name] = r
		rectIndex[r.Name] = idx
	}

	rInst1 := rectByName["Inst1-RootDiag"]
	rSubB1_1 := rectByName["Inst1-SubB1-RootDiag"]
	rSubB2_1 := rectByName["Inst1-SubB2-RootDiag"]
	rLeafC1_1 := rectByName["Inst1-SubB1-LeafC1-RootDiag"]
	rInst2 := rectByName["Inst2-RootDiag"]
	rSubB1_2 := rectByName["Inst2-SubB1-RootDiag"]

	if rInst1 == nil || rSubB1_1 == nil || rSubB2_1 == nil || rLeafC1_1 == nil || rInst2 == nil || rSubB1_2 == nil {
		t.Fatalf("missing expected nested SVG rects: got keys %v", func() []string {
			var ks []string
			for k := range rectByName {
				ks = append(ks, k)
			}
			return ks
		}())
	}

	// Z-order: ancestor before child
	if !(rectIndex["Inst1-RootDiag"] < rectIndex["Inst1-SubB1-RootDiag"] &&
		rectIndex["Inst1-SubB1-RootDiag"] < rectIndex["Inst1-SubB1-LeafC1-RootDiag"]) {
		t.Fatalf("expected parent rects before child rects in layer.Rects Z-order: idx=%v", rectIndex)
	}

	// Title anchor & "Part : Type" formatting: composite containers at RECT_TOP, leaf parts at RECT_CENTER_MIDDLE
	if rInst1.RectAnchoredTexts[0].RectAnchorType != svg.RECT_TOP || rSubB1_1.RectAnchoredTexts[0].RectAnchorType != svg.RECT_TOP {
		t.Fatalf("expected Inst1 and SubB1 titles anchored at RECT_TOP when displaying nested parts")
	}
	if rInst1.RectAnchoredTexts[0].Content != "Inst1 : SysSub" || rSubB1_1.RectAnchoredTexts[0].Content != "SubB1 : SysLeaf" {
		t.Fatalf("expected 'Part : Type' titles on Inst1/SubB1, got %q / %q",
			rInst1.RectAnchoredTexts[0].Content, rSubB1_1.RectAnchoredTexts[0].Content)
	}
	if rSubB2_1.RectAnchoredTexts[0].RectAnchorType != svg.RECT_CENTER_MIDDLE || rLeafC1_1.RectAnchoredTexts[0].RectAnchorType != svg.RECT_CENTER_MIDDLE {
		t.Fatalf("expected leaf parts SubB2 and LeafC1 titles anchored at RECT_CENTER_MIDDLE")
	}
	if rSubB2_1.RectAnchoredTexts[0].Content != "SubB2" || rLeafC1_1.RectAnchoredTexts[0].Content != "LeafC1" {
		t.Fatalf("expected untyped part titles 'SubB2'/'LeafC1', got %q / %q",
			rSubB2_1.RectAnchoredTexts[0].Content, rLeafC1_1.RectAnchoredTexts[0].Content)
	}

	// EnclosingRect & Peers
	if rSubB1_1.EnclosingRect != rInst1 || rLeafC1_1.EnclosingRect != rSubB1_1 {
		t.Fatalf("expected EnclosingRect chain LeafC1 -> SubB1 -> Inst1")
	}
	if !slices.Contains(rInst1.Peers, rSubB1_1) || !slices.Contains(rInst1.Peers, rLeafC1_1) {
		t.Fatalf("expected rInst1.Peers to include both direct child SubB1 and grandchild LeafC1")
	}

	// Sibling-only Obstacles
	if !slices.Contains(rSubB1_1.Obstacles, rSubB2_1) {
		t.Fatalf("expected siblings SubB1 and SubB2 inside Inst1 to be obstacles to each other")
	}
	if slices.Contains(rSubB1_1.Obstacles, rInst1) || slices.Contains(rInst1.Obstacles, rSubB1_1) {
		t.Fatalf("parent Inst1 and child SubB1 must NOT be obstacles to each other")
	}
	if slices.Contains(rSubB1_1.Obstacles, rSubB1_2) {
		t.Fatalf("nested parts inside different parent instances (Inst1 vs Inst2) must NOT be obstacles to each other")
	}

	// Verify 4 links rendered on SVG (2 ControlFlows + 2 internal DataFlows)
	if len(svgObj.Layers[0].Links) != 4 {
		t.Fatalf("expected 4 SVG links (2 ControlFlows + 2 DataFlows), got %d", len(svgObj.Layers[0].Links))
	}

	// 6. Semantic clamping: drag LeafC1 far outside SubB1 and verify it clamps inside SubB1
	var leafC1ShapeInInst1 *models.PartShape
	var subB1ShapeInInst1 *models.PartShape
	for _, ps := range diag.Part_Shapes {
		switch ps.Name {
		case "Inst1-SubB1-LeafC1-RootDiag":
			leafC1ShapeInInst1 = ps
		case "Inst1-SubB1-RootDiag":
			subB1ShapeInInst1 = ps
		}
	}
	leafC1ShapeInInst1.X = subB1ShapeInInst1.X + subB1ShapeInInst1.Width + 500
	leafC1ShapeInInst1.Y = subB1ShapeInInst1.Y + subB1ShapeInInst1.Height + 500
	stage.Commit()

	if leafC1ShapeInInst1.X < subB1ShapeInInst1.X ||
		leafC1ShapeInInst1.X+leafC1ShapeInInst1.Width > subB1ShapeInInst1.X+subB1ShapeInInst1.Width ||
		leafC1ShapeInInst1.Y < subB1ShapeInInst1.Y ||
		leafC1ShapeInInst1.Y+leafC1ShapeInInst1.Height > subB1ShapeInInst1.Y+subB1ShapeInInst1.Height {
		t.Fatalf("expected LeafC1 clamped inside SubB1 bounds: LeafC1=(%v,%v,%v,%v) SubB1=(%v,%v,%v,%v)",
			leafC1ShapeInInst1.X, leafC1ShapeInInst1.Y, leafC1ShapeInInst1.Width, leafC1ShapeInInst1.Height,
			subB1ShapeInInst1.X, subB1ShapeInInst1.Y, subB1ShapeInInst1.Width, subB1ShapeInInst1.Height)
	}

	// 7. Uncheck Inst1 -> cascade unstages all of Inst1's nested shapes while keeping Inst2 100% intact!
	inst1Node = getInstNode("Inst1 : SysSub")
	inst1Node.OnIsCheckedChanged(false)

	if len(diag.Part_Shapes) != 4 {
		t.Fatalf("expected 4 Part_Shapes remaining (Inst2 subtree) after unchecking Inst1, got %d", len(diag.Part_Shapes))
	}
	if len(diag.Port_Shapes) != 3 {
		t.Fatalf("expected 3 Port_Shapes remaining (Inst2 subtree) after unchecking Inst1, got %d", len(diag.Port_Shapes))
	}
	if len(diag.ControlFlow_Shapes) != 1 || len(diag.DataFlow_Shapes) != 1 {
		t.Fatalf("expected 1 ControlFlowShape and 1 DataFlowShape remaining (Inst2 subtree), got cf=%d df=%d",
			len(diag.ControlFlow_Shapes), len(diag.DataFlow_Shapes))
	}
	if len(diag.Data_Shapes) != 1 || len(diag.Protocol_Shapes) != 1 {
		t.Fatalf("expected 1 DataShape and 1 ProtocolShape remaining (Inst2 subtree), got data=%d proto=%d",
			len(diag.Data_Shapes), len(diag.Protocol_Shapes))
	}

	// 8. Cycle prevention: direct cycle (SysLeaf -> leafC1.Type = SysLeaf) and indirect cycle (SysLeaf -> leafC1.Type = SysRoot)
	leafC1.Type = sysLeaf
	stage.Commit()
	if leafC1.Type != nil {
		t.Fatalf("expected direct cyclic Part.Type (SysLeaf -> SysLeaf) to be cleared by semantic enforcement")
	}

	leafC1.Type = sysRoot
	stage.Commit()
	if leafC1.Type != nil {
		t.Fatalf("expected indirect cyclic Part.Type (SysRoot -> SysSub -> SysLeaf -> SysRoot) to be cleared by semantic enforcement")
	}
	if inst1.Type != sysSub || inst2.Type != sysSub || subB1.Type != sysLeaf {
		t.Fatalf("expected pre-existing valid Part.Type links on inst1, inst2, subB1 to remain intact after rejecting cyclic edit on leafC1")
	}

	// 9. Interactive SVG canvas link creation between nested PortShapes inside Inst2
	b2p2 := (&models.Port{Name: "b2p2"}).Stage(stage)
	subB2.Ports = append(subB2.Ports, b2p2)
	stage.Commit()

	// Check b2p2 under Inst2 -> Parts -> SubB2 -> Ports
	inst2SubB2Node := findDirectChildByName(findDirectChildByName(getInstNode("Inst2 : SysSub"), "Parts"), "SubB2")
	inst2B2P2Node := findDirectChildByName(findDirectChildByName(inst2SubB2Node, "Ports"), b2p2.GetName())
	if inst2B2P2Node == nil {
		t.Fatalf("expected b2p2 node under Inst2/SubB2/Ports")
	}
	inst2B2P2Node.OnIsCheckedChanged(true)

	findPortRectInPart := func(svgObject *svg.SVG, enclosingPartRectName, portName string) *svg.Rect {
		var partRect *svg.Rect
		for _, r := range svgObject.Layers[0].Rects {
			if r.Name == enclosingPartRectName {
				partRect = r
				break
			}
		}
		if partRect == nil {
			return nil
		}
		for _, r := range svgObject.Layers[0].Rects {
			if r.AnchoredTo == partRect && r.Name == portName {
				return r
			}
		}
		return nil
	}

	svgObj = stager.GetSvgObject()
	rInst2B2P1 := findPortRectInPart(svgObj, "Inst2-SubB2-RootDiag", b2p1.GetName())
	rInst2B2P2 := findPortRectInPart(svgObj, "Inst2-SubB2-RootDiag", b2p2.GetName())
	rInst2B1P1 := findPortRectInPart(svgObj, "Inst2-SubB1-RootDiag", b1p1.GetName())
	if rInst2B2P1 == nil || rInst2B2P2 == nil || rInst2B1P1 == nil {
		t.Fatalf("expected SVG port rects anchored inside Inst2: b2p1=%v b2p2=%v b1p1=%v", rInst2B2P1, rInst2B2P2, rInst2B1P1)
	}

	// 9a. Same nested PartShape (b2p1 -> b2p2 inside Inst2/SubB2) -> creates ControlFlow on subB2 + scoped ControlFlowShape
	prevSubB2CFCount := len(subB2.ControlFlows)
	svgObj.DrawingState = svg.DRAWING_LINK
	svgObj.OnUpdate(&svg.SVG{
		DrawingState: svg.NOT_DRAWING_LINK,
		StartRect:    rInst2B2P1,
		EndRect:      rInst2B2P2,
	})
	if len(subB2.ControlFlows) != prevSubB2CFCount+1 {
		t.Fatalf("expected interactive SVG link within Inst2/SubB2 to create ControlFlow on subB2, got %d", len(subB2.ControlFlows))
	}
	if len(diag.ControlFlow_Shapes) != 2 {
		t.Fatalf("expected 2 ControlFlow_Shapes in diagram after drawing nested ControlFlow, got %d", len(diag.ControlFlow_Shapes))
	}

	// 9b. Sibling nested PartShapes (b2p2 -> b1p1 inside Inst2) -> creates internal DataFlow on sysSub + scoped DataFlowShape
	prevSysSubDFCount := len(sysSub.DataFlows)
	prevSysRootDFCount := len(sysRoot.DataFlows)
	svgObj = stager.GetSvgObject()
	svgObj.DrawingState = svg.DRAWING_LINK
	svgObj.OnUpdate(&svg.SVG{
		DrawingState: svg.NOT_DRAWING_LINK,
		StartRect:    findPortRectInPart(svgObj, "Inst2-SubB2-RootDiag", b2p2.GetName()),
		EndRect:      findPortRectInPart(svgObj, "Inst2-SubB1-RootDiag", b1p1.GetName()),
	})
	if len(sysSub.DataFlows) != prevSysSubDFCount+1 {
		t.Fatalf("expected interactive SVG link between Inst2/SubB2 and Inst2/SubB1 to create DataFlow on sysSub, got %d (was %d)",
			len(sysSub.DataFlows), prevSysSubDFCount)
	}
	if len(sysRoot.DataFlows) != prevSysRootDFCount {
		t.Fatalf("expected sysRoot.DataFlows untouched when drawing internal DataFlow inside Inst2, got %d", len(sysRoot.DataFlows))
	}
	if len(diag.DataFlow_Shapes) != 2 {
		t.Fatalf("expected 2 DataFlow_Shapes in diagram after drawing nested DataFlow, got %d", len(diag.DataFlow_Shapes))
	}
}
