package main

import (
	"testing"
	"time"

	"github.com/fullstack-lang/gong/dsm/project/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/project/go/models"
	form "github.com/fullstack-lang/gong/lib/form/go/models"
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func TestMilestoneVerticalOffset(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("test_milestone_offset", "", "", true, false, false)
	stager := models.NewStager(
		stack.R,
		stack.Stage,
		stack.Probe,
		"",
	)
	_ = stager

	stage := stack.Stage

	lib := (&models.Library{
		Name:          "RootLib",
		IsRootLibrary: true,
	}).Stage(stage)

	taskGroup := (&models.TaskGroup{
		Name: "Milestones Group",
	}).Stage(stage)
	lib.RootTaskGroups = []*models.TaskGroup{taskGroup}

	milestoneTask := (&models.Task{
		Name:        "Alpha Milestone",
		IsMilestone: true,
		Start:       time.Date(2026, 3, 15, 0, 0, 0, 0, time.UTC),
		End:         time.Date(2026, 3, 15, 0, 0, 0, 0, time.UTC),
	}).Stage(stage)
	taskGroup.Tasks = []*models.Task{milestoneTask}
	lib.RootTasks = []*models.Task{milestoneTask}
	for _, d := range stage.GetInstancesSorted[*models.Diagram]() {
		d.IsChecked = false
	}

	diag := (&models.Diagram{
		Name:                      "Gantt",
		IsTimeDiagram:             true,
		IsChecked:                 true,
		IsEditable_:               true,
		TextHeight:                15.0,
		LaneHeight:                85.0,
		YTopMargin:                40.0,
		XLeftLanes:                240.0,
		XRightMargin:              1250.0,
		UseManualStartAndEndDates: true,
		ManualStart:               time.Date(2026, 3, 1, 0, 0, 0, 0, time.UTC),
		ManualEnd:                 time.Date(2026, 3, 31, 0, 0, 0, 0, time.UTC),
	}).Stage(stage)
	lib.Diagrams = []*models.Diagram{diag}

	taskGroupShape := (&models.TaskGroupShape{
		Name:      "Gantt-Milestones Group",
		TaskGroup: taskGroup,
	}).Stage(stage)
	diag.TaskGroupShapes = []*models.TaskGroupShape{taskGroupShape}

	taskShape := (&models.TaskShape{
		Name: "Gantt-Alpha Milestone",
		Task: milestoneTask,
	}).Stage(stage)
	diag.Task_Shapes = []*models.TaskShape{taskShape}

	stage.Commit()

	// Find the tree node for milestoneTask
	findTaskNode := func(name string) *tree.Node {
		for n := range *stager.GetTreeStage().GetInstancesSet[*tree.Node]() {
			if n.Name == name {
				return n
			}
		}
		return nil
	}

	taskNode := findTaskNode("Alpha Milestone")
	if taskNode == nil {
		t.Fatalf("taskNode for Alpha Milestone not found in tree")
	}

	getButtons := func() (up *tree.Button, down *tree.Button, reset *tree.Button) {
		node := findTaskNode("Alpha Milestone")
		if node == nil {
			return
		}
		for _, b := range node.Buttons {
			if b.Name == "Move milestone up" {
				up = b
			} else if b.Name == "Move milestone down" {
				down = b
			}
		}
		if node.Menu != nil {
			for _, b := range node.Menu.Buttons {
				if b.Name == "Reset milestone vertical offset" {
					reset = b
				}
			}
		}
		return
	}

	upButton, downButton, _ := getButtons()
	if upButton == nil || downButton == nil {
		t.Fatalf("expected up and down buttons on taskNode.Buttons")
	}

	// Helper to find the diamond Rect in svg
	findDiamond := func() *svg.Rect {
		for r := range *stager.GetSvgStage().GetInstancesSet[*svg.Rect]() {
			if r.Name == "Alpha Milestone" && r.Transform != "" {
				return r
			}
		}
		return nil
	}

	diamond := findDiamond()
	if diamond == nil {
		t.Fatalf("diamond svg.Rect not found for milestone")
	}
	initialDiamondY := diamond.Y

	// Click Up button: VerticalOffset should decrease by 15.0
	upButton.OnClick()

	if taskShape.VerticalOffset != -15.0 {
		t.Errorf("expected VerticalOffset -15.0, got %f", taskShape.VerticalOffset)
	}

	diamond = findDiamond()
	if diamond == nil {
		t.Fatalf("diamond not found after up click")
	}
	if diamond.Y != initialDiamondY-15.0 {
		t.Errorf("expected diamond.Y %f, got %f", initialDiamondY-15.0, diamond.Y)
	}

	// Click Down button twice: VerticalOffset should become +15.0
	_, downButton, _ = getButtons()
	downButton.OnClick()
	_, downButton, _ = getButtons()
	downButton.OnClick()

	if taskShape.VerticalOffset != 15.0 {
		t.Errorf("expected VerticalOffset 15.0, got %f", taskShape.VerticalOffset)
	}

	diamond = findDiamond()
	if diamond == nil {
		t.Fatalf("diamond not found after down clicks")
	}
	if diamond.Y != initialDiamondY+15.0 {
		t.Errorf("expected diamond.Y %f, got %f", initialDiamondY+15.0, diamond.Y)
	}

	// Check that reset button exists in the menu when offset != 0
	_, _, resetButton := getButtons()
	if resetButton == nil {
		t.Fatalf("expected reset button in menu when VerticalOffset != 0")
	}

	// Click reset button
	resetButton.OnClick()
	if taskShape.VerticalOffset != 0 {
		t.Errorf("expected VerticalOffset 0 after reset, got %f", taskShape.VerticalOffset)
	}
	diamond = findDiamond()
	if diamond == nil {
		t.Fatalf("diamond is nil after reset")
	}
	if diamond.Y != initialDiamondY {
		t.Errorf("expected diamond.Y back to %f, got %f", initialDiamondY, diamond.Y)
	}
}

func TestMilestoneVerticalLineAppearanceOverride(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("test_milestone_vertical_bar_override", "", "", true, false, false)
	stager := models.NewStager(
		stack.R,
		stack.Stage,
		stack.Probe,
		"",
	)

	stage := stack.Stage

	lib := (&models.Library{
		Name:          "RootLib",
		IsRootLibrary: true,
	}).Stage(stage)

	taskGroup := (&models.TaskGroup{
		Name: "Milestones Group",
	}).Stage(stage)
	lib.RootTaskGroups = []*models.TaskGroup{taskGroup}

	milestoneTask := (&models.Task{
		Name:        "Release Milestone",
		IsMilestone: true,
		Start:       time.Date(2026, 3, 15, 0, 0, 0, 0, time.UTC),
		End:         time.Date(2026, 3, 15, 0, 0, 0, 0, time.UTC),
	}).Stage(stage)
	taskGroup.Tasks = []*models.Task{milestoneTask}
	lib.RootTasks = []*models.Task{milestoneTask}

	for _, d := range stage.GetInstancesSorted[*models.Diagram]() {
		d.IsChecked = false
	}

	diag := (&models.Diagram{
		Name:                      "Gantt",
		IsTimeDiagram:             true,
		IsChecked:                 true,
		IsEditable_:               true,
		TextHeight:                15.0,
		LaneHeight:                85.0,
		YTopMargin:                40.0,
		XLeftLanes:                240.0,
		XRightMargin:              1250.0,
		UseManualStartAndEndDates: true,
		ManualStart:               time.Date(2026, 3, 1, 0, 0, 0, 0, time.UTC),
		ManualEnd:                 time.Date(2026, 3, 31, 0, 0, 0, 0, time.UTC),
	}).Stage(stage)
	lib.Diagrams = []*models.Diagram{diag}

	taskGroupShape := (&models.TaskGroupShape{
		Name:      "Gantt-Milestones Group",
		TaskGroup: taskGroup,
	}).Stage(stage)
	diag.TaskGroupShapes = []*models.TaskGroupShape{taskGroupShape}

	taskShape := (&models.TaskShape{
		Name:               "Gantt-Release Milestone",
		Task:               milestoneTask,
		DisplayVerticalBar: true,
	}).Stage(stage)
	diag.Task_Shapes = []*models.TaskShape{taskShape}

	stage.Commit()

	findMilestoneLine := func() *svg.Line {
		for l := range *stager.GetSvgStage().GetInstancesSet[*svg.Line]() {
			if l.Name == "Release Milestone" {
				return l
			}
		}
		return nil
	}

	findDiamond := func() *svg.Rect {
		for r := range *stager.GetSvgStage().GetInstancesSet[*svg.Rect]() {
			if r.Name == "Release Milestone" && r.Transform != "" {
				return r
			}
		}
		return nil
	}

	// 1. Verify default vertical bar rendering and synced TaskShape defaults
	line := findMilestoneLine()
	if line == nil {
		t.Fatalf("expected milestone vertical line to be generated")
	}
	if line.Stroke != "black" || line.StrokeWidth != 0.5 || line.StrokeOpacity != 1.0 || line.StrokeDashArray != "2 2" {
		t.Errorf("unexpected default vertical line styling: Stroke=%q, StrokeWidth=%v, StrokeOpacity=%v, StrokeDashArray=%q",
			line.Stroke, line.StrokeWidth, line.StrokeOpacity, line.StrokeDashArray)
	}
	if taskShape.VerticalBarStroke != "black" || taskShape.VerticalBarStrokeWidth != 0.5 || taskShape.VerticalBarStrokeOpacity != 1.0 || taskShape.VerticalBarStrokeDashArray != "2 2" {
		t.Errorf("expected taskShape vertical bar fields to be initialized with defaults, got Stroke=%q, Width=%v, Opacity=%v, DashArray=%q",
			taskShape.VerticalBarStroke, taskShape.VerticalBarStrokeWidth, taskShape.VerticalBarStrokeOpacity, taskShape.VerticalBarStrokeDashArray)
	}

	// 2. Select the milestone diamond and verify the Shape Appearance form contains VerticalBar* fields
	diamond := findDiamond()
	if diamond == nil || diamond.OnSelect == nil {
		t.Fatalf("expected selectable milestone diamond")
	}
	diamond.OnSelect()

	shapeFormStage := stager.GetShapeFormStage()
	var formGroup *form.FormGroup
	for fg := range *shapeFormStage.GetInstancesSet[*form.FormGroup]() {
		formGroup = fg
	}
	if formGroup == nil {
		t.Fatalf("expected FormGroup in shapeFormStage after selecting milestone")
	}

	foundFields := make(map[string]*form.FormDiv)
	for _, div := range formGroup.FormDivs {
		foundFields[div.Name] = div
	}
	for _, expectedField := range []string{"VerticalBarStroke", "VerticalBarStrokeWidth", "VerticalBarStrokeOpacity", "VerticalBarStrokeDashArray"} {
		if _, ok := foundFields[expectedField]; !ok {
			t.Errorf("expected field %q in milestone shape appearance form", expectedField)
		}
	}

	// 3. Override vertical bar styling through the form and save
	foundFields["IsPresentationOverridden"].CheckBoxs[0].Value = true
	foundFields["VerticalBarStroke"].FormFields[0].FormFieldString.Value = "crimson"
	foundFields["VerticalBarStrokeWidth"].FormFields[0].FormFieldFloat64.Value = 2.5
	foundFields["VerticalBarStrokeOpacity"].FormFields[0].FormFieldFloat64.Value = 0.75
	foundFields["VerticalBarStrokeDashArray"].FormFields[0].FormFieldString.Value = "6 3"
	shapeFormStage.Commit()

	formGroup.OnSave.OnSave()

	if !taskShape.GetIsPresentationOverridden() {
		t.Errorf("expected taskShape.IsPresentationOverridden to be true")
	}
	if taskShape.VerticalBarStroke != "crimson" || taskShape.VerticalBarStrokeWidth != 2.5 || taskShape.VerticalBarStrokeOpacity != 0.75 || taskShape.VerticalBarStrokeDashArray != "6 3" {
		t.Errorf("unexpected taskShape vertical bar fields after save: Stroke=%q, Width=%v, Opacity=%v, DashArray=%q",
			taskShape.VerticalBarStroke, taskShape.VerticalBarStrokeWidth, taskShape.VerticalBarStrokeOpacity, taskShape.VerticalBarStrokeDashArray)
	}

	line = findMilestoneLine()
	if line == nil {
		t.Fatalf("expected milestone vertical line after override")
	}
	if line.Stroke != "crimson" || line.StrokeWidth != 2.5 || line.StrokeOpacity != 0.75 || line.StrokeDashArray != "6 3" {
		t.Errorf("expected overridden vertical line styling, got Stroke=%q, StrokeWidth=%v, StrokeOpacity=%v, StrokeDashArray=%q",
			line.Stroke, line.StrokeWidth, line.StrokeOpacity, line.StrokeDashArray)
	}

	// 4. Verify setting VerticalBarStrokeDashArray to "none" produces a solid line
	taskShape.VerticalBarStrokeDashArray = "none"
	stage.Commit()

	line = findMilestoneLine()
	if line == nil {
		t.Fatalf("expected milestone vertical line")
	}
	if line.StrokeDashArray != "" {
		t.Errorf("expected empty StrokeDashArray when VerticalBarStrokeDashArray is 'none', got %q", line.StrokeDashArray)
	}
}
