package models

import (
	"slices"
	"testing"
	"time"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type testMockProbe struct {
	ProbeIF
}

func (m *testMockProbe) AddNotification(date time.Time, message string)               {}
func (m *testMockProbe) CommitNotificationTable()                                     {}
func (m *testMockProbe) AddCommitNavigationNode(fn func(GongNodeIF))                  {}
func (m *testMockProbe) FillUpFormFromGongstruct(instance any, gongstructName string) {}
func (m *testMockProbe) Refresh()                                                     {}

func TestScenarioLibraryTreeAndButtons(t *testing.T) {
	stage := NewStage("test")
	treeStage := tree.NewStage("test_tree")
	probe := &testMockProbe{}

	stager := &Stager{
		stage:                stage,
		treeStage:            treeStage,
		probeForm:            probe,
		map_Element_Diagrams: make(map[AbstractType][]*Diagram),
	}

	// 1. Create an orphaned Analysis (like in sceanrio examples.go)
	analysis := (&Analysis{Name: "Test Analysis"}).Stage(stage)
	scenario := (&Scenario{Name: "Test Scenario"}).Stage(stage)
	analysis.Scenarios = append(analysis.Scenarios, scenario)

	diagram := (&Diagram{
		Name:      "Test Diagram",
		IsChecked: true,
	}).Stage(stage)
	scenario.Diagrams = append(scenario.Diagrams, diagram)

	// 2. Run enforceSemantic
	stager.enforceSemantic()

	rootLib := stager.getRootLibrary()
	if rootLib == nil {
		t.Fatal("expected root library to be created")
	}

	// Verify Analysis is attached to root library
	found := slices.Contains(rootLib.Analyses, analysis)
	if !found {
		t.Fatal("expected analysis to be attached to root library Analyses slice")
	}

	// 3. Build tree
	stager.ux_tree()

	treeInstances := treeStage.GetInstancesSorted[*tree.Tree]()
	if len(treeInstances) == 0 {
		t.Fatal("expected tree to be staged")
	}
	treeInstance := treeInstances[0]

	// Check root node is Library
	if len(treeInstance.RootNodes) == 0 {
		t.Fatal("expected tree to have root nodes")
	}
	libNode := treeInstance.RootNodes[0]
	if libNode.Name != "Root Library" && libNode.Name != rootLib.Name {
		t.Fatalf("expected root node to be Library, got %s", libNode.Name)
	}

	// Check Analyses node
	var analysesNode *tree.Node
	for _, child := range libNode.Children {
		if child.Name == "Analyses" {
			analysesNode = child
			break
		}
	}
	if analysesNode == nil {
		t.Fatal("expected library to have 'Analyses' child node")
	}

	if len(analysesNode.Buttons) == 0 {
		t.Fatal("expected Analyses node to have create item button")
	}

	// Check Analysis node
	if len(analysesNode.Children) == 0 {
		t.Fatal("expected Analyses node to have child Analysis node")
	}
	analysisNode := analysesNode.Children[0]
	if analysisNode.Name != "Test Analysis" {
		t.Fatalf("expected child to be 'Test Analysis', got %s", analysisNode.Name)
	}

	// Check Scenarios node under Analysis
	var scenariosNode *tree.Node
	for _, child := range analysisNode.Children {
		if child.Name == "Scenarios" {
			scenariosNode = child
			break
		}
	}
	if scenariosNode == nil {
		t.Fatal("expected analysis to have 'Scenarios' child node")
	}
	if len(scenariosNode.Buttons) == 0 {
		t.Fatal("expected Scenarios node to have create item button")
	}

	// Check Scenario node
	if len(scenariosNode.Children) == 0 {
		t.Fatal("expected Scenarios node to have child Scenario node")
	}
	scenarioNode := scenariosNode.Children[0]
	if scenarioNode.Name != "Test Scenario" {
		t.Fatalf("expected child to be 'Test Scenario', got %s", scenarioNode.Name)
	}

	// Check Diagrams node under Scenario
	var diagramsNode *tree.Node
	for _, child := range scenarioNode.Children {
		if child.Name == "Diagrams" {
			diagramsNode = child
			break
		}
	}
	if diagramsNode == nil {
		t.Fatal("expected scenario to have 'Diagrams' child node")
	}
	if len(diagramsNode.Buttons) == 0 {
		t.Fatal("expected Diagrams node to have create item button")
	}

	// Check Diagram node
	if len(diagramsNode.Children) == 0 {
		t.Fatal("expected Diagrams node to have child Diagram node")
	}
	diagramNode := diagramsNode.Children[0]
	if diagramNode.Name != "Test Diagram" {
		t.Fatalf("expected child to be 'Test Diagram', got %s", diagramNode.Name)
	}

	// Check diagram editability button
	var editButton *tree.Button
	for _, btn := range diagramNode.Buttons {
		if btn.Name == "Diagram Editability" {
			editButton = btn
			break
		}
	}
	if editButton == nil {
		t.Fatal("expected diagram node to have 'Diagram Editability' button")
	}
	if diagram.IsInDrawMode {
		t.Fatal("expected diagram.IsInDrawMode to initially be false")
	}
	if diagram.IsEditable() {
		t.Fatal("expected diagram.IsEditable() to initially be false")
	}

	// Click edit button to enable editability
	editButton.OnClick()
	if !diagram.IsInDrawMode {
		t.Fatal("expected diagram.IsInDrawMode to be true after clicking edit button")
	}
	if !diagram.IsEditable() {
		t.Fatal("expected diagram.IsEditable() to be true after clicking edit button")
	}

	// Click edit button again to stop editing
	editButton.OnClick()
	if diagram.IsInDrawMode {
		t.Fatal("expected diagram.IsInDrawMode to be false after clicking edit button again")
	}
	if diagram.IsEditable() {
		t.Fatal("expected diagram.IsEditable() to be false after clicking edit button again")
	}

	// Check all 5 element category nodes under Diagram
	expectedCategories := []string{
		"Parameters",
		"Actor States",
		"Evolution Directions",
		"Parameters Aggregates",
		"Actor State Transitions",
	}

	for _, catName := range expectedCategories {
		var catNode *tree.Node
		for _, child := range diagramNode.Children {
			if child.Name == catName {
				catNode = child
				break
			}
		}
		if catNode == nil {
			t.Fatalf("expected diagram to have '%s' child node", catName)
		}
		if len(catNode.Buttons) == 0 {
			t.Fatalf("expected category node '%s' to have create item button", catName)
		}
	}
}

func TestScenarioButtonClicks(t *testing.T) {
	stage := NewStage("test_click")
	treeStage := tree.NewStage("test_tree_click")
	probe := &testMockProbe{}

	stager := &Stager{
		stage:                stage,
		treeStage:            treeStage,
		probeForm:            probe,
		map_Element_Diagrams: make(map[AbstractType][]*Diagram),
	}

	analysis := (&Analysis{Name: "Analysis 1"}).Stage(stage)
	scenario := (&Scenario{Name: "Scenario 1"}).Stage(stage)
	analysis.Scenarios = append(analysis.Scenarios, scenario)

	diagram := (&Diagram{Name: "Diagram 1", IsChecked: true}).Stage(stage)
	scenario.Diagrams = append(scenario.Diagrams, diagram)

	stager.enforceSemantic()
	stager.ux_tree()

	treeInstances := treeStage.GetInstancesSorted[*tree.Tree]()
	treeInstance := treeInstances[0]
	libNode := treeInstance.RootNodes[0]

	// 1. Click Add Analysis on Analyses node
	var analysesNode *tree.Node
	for _, child := range libNode.Children {
		if child.Name == "Analyses" {
			analysesNode = child
			break
		}
	}
	if analysesNode == nil || len(analysesNode.Buttons) == 0 {
		t.Fatal("no analyses node button")
	}
	analysesCountBefore := len(stager.getRootLibrary().Analyses)
	analysesNode.Buttons[0].OnClick()
	if len(stager.getRootLibrary().Analyses) != analysesCountBefore+1 {
		t.Fatalf("expected analyses count to increase by 1, got %d", len(stager.getRootLibrary().Analyses))
	}

	// 2. Click Add Scenario on Scenarios node
	stager.ux_tree()
	libNode = treeStage.GetInstancesSorted[*tree.Tree]()[0].RootNodes[0]
	for _, child := range libNode.Children {
		if child.Name == "Analyses" {
			analysesNode = child
			break
		}
	}
	analysisNode := analysesNode.Children[0]
	var scenariosNode *tree.Node
	for _, child := range analysisNode.Children {
		if child.Name == "Scenarios" {
			scenariosNode = child
			break
		}
	}
	if scenariosNode == nil || len(scenariosNode.Buttons) == 0 {
		t.Fatal("no scenarios node button")
	}
	scenariosCountBefore := len(analysis.Scenarios)
	scenariosNode.Buttons[0].OnClick()
	if len(analysis.Scenarios) != scenariosCountBefore+1 {
		t.Fatalf("expected scenarios count to increase by 1, got %d", len(analysis.Scenarios))
	}

	// 3. Click Add Diagram on Diagrams node
	stager.ux_tree()
	libNode = treeStage.GetInstancesSorted[*tree.Tree]()[0].RootNodes[0]
	for _, child := range libNode.Children {
		if child.Name == "Analyses" {
			analysesNode = child
			break
		}
	}
	analysisNode = analysesNode.Children[0]
	for _, child := range analysisNode.Children {
		if child.Name == "Scenarios" {
			scenariosNode = child
			break
		}
	}
	var scenarioNode *tree.Node
	for _, child := range scenariosNode.Children {
		if child.Name == scenario.Name {
			scenarioNode = child
			break
		}
	}
	if scenarioNode == nil {
		t.Fatal("scenario node not found")
	}
	var diagramsNode *tree.Node
	for _, child := range scenarioNode.Children {
		if child.Name == "Diagrams" {
			diagramsNode = child
			break
		}
	}
	if diagramsNode == nil || len(diagramsNode.Buttons) == 0 {
		t.Fatal("no diagrams node button")
	}
	diagramsCountBefore := len(scenario.Diagrams)
	diagramsNode.Buttons[0].OnClick()
	if len(scenario.Diagrams) != diagramsCountBefore+1 {
		t.Fatalf("expected diagrams count to increase by 1, got %d", len(scenario.Diagrams))
	}

	// 4. Click Add Parameter on Parameters node
	stager.ux_tree()
	libNode = treeStage.GetInstancesSorted[*tree.Tree]()[0].RootNodes[0]
	for _, child := range libNode.Children {
		if child.Name == "Analyses" {
			analysesNode = child
			break
		}
	}
	analysisNode = analysesNode.Children[0]
	for _, child := range analysisNode.Children {
		if child.Name == "Scenarios" {
			scenariosNode = child
			break
		}
	}
	for _, child := range scenariosNode.Children {
		if child.Name == scenario.Name {
			scenarioNode = child
			break
		}
	}
	for _, child := range scenarioNode.Children {
		if child.Name == "Diagrams" {
			diagramsNode = child
			break
		}
	}
	diagramNode := diagramsNode.Children[0]

	var paramsNode *tree.Node
	for _, child := range diagramNode.Children {
		if child.Name == "Parameters" {
			paramsNode = child
			break
		}
	}
	if paramsNode == nil || len(paramsNode.Buttons) == 0 {
		t.Fatal("no parameters node button")
	}
	paramsCountBefore := len(scenario.Parameters)
	paramsNode.Buttons[0].OnClick()
	if len(scenario.Parameters) != paramsCountBefore+1 {
		t.Fatalf("expected parameters count to increase by 1, got %d", len(scenario.Parameters))
	}
}
