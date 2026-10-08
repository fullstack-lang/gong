package main

import (
	"math"
	"strings"
	"testing"
	"time"

	"github.com/fullstack-lang/gong/dsm/project/go/level1stack"
	"github.com/fullstack-lang/gong/dsm/project/go/models"
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func TestSecondaryTimeScaleSVGIntegration(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("test_secondary_time_scale", "", "", true, false, false)
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

	task := (&models.Task{
		Name:  "Task1",
		Start: time.Date(2026, 3, 1, 0, 0, 0, 0, time.UTC),
		End:   time.Date(2026, 3, 10, 0, 0, 0, 0, time.UTC),
	}).Stage(stage)

	tg := (&models.TaskGroup{
		Name:  "TG1",
		Tasks: []*models.Task{task},
	}).Stage(stage)
	lib.RootTaskGroups = []*models.TaskGroup{tg}

	for _, d := range stage.GetInstancesSorted[*models.Diagram]() {
		d.IsChecked = false
	}

	diag := (&models.Diagram{
		Name:                           "GanttWithSecondary",
		IsTimeDiagram:                  true,
		IsChecked:                      true,
		IsEditable_:                    true,
		TextHeight:                     15.0,
		LaneHeight:                     85.0,
		YTopMargin:                     40.0,
		DateYOffset:                    15.0,
		XLeftLanes:                     200.0,
		XRightMargin:                   1200.0,
		UseManualStartAndEndDates:      true,
		ManualStart:                    time.Date(2026, 3, 1, 0, 0, 0, 0, time.UTC),
		ManualEnd:                      time.Date(2026, 3, 31, 0, 0, 0, 0, time.UTC),
		TimeStep:                       1,
		TimeStepScale:                  models.MONTHS,
		SecondaryTimeStep:              1,
		SecondaryTimeStepScale:         models.DAYS,
		DrawSecondaryVerticalTimeLines: true,
	}).Stage(stage)
	lib.Diagrams = []*models.Diagram{diag}

	tgs := (&models.TaskGroupShape{
		Name:      "GanttWithSecondary-TG1",
		TaskGroup: tg,
	}).Stage(stage)
	diag.TaskGroupShapes = []*models.TaskGroupShape{tgs}

	ts := (&models.TaskShape{
		Name: "GanttWithSecondary-Task1",
		Task: task,
	}).Stage(stage)
	diag.Task_Shapes = []*models.TaskShape{ts}

	// Trigger commit: enforces semantics, diagram size, and builds SVG
	stage.Commit()

	// 1. Verify HasSecondaryTimeScale
	if !diag.HasSecondaryTimeScale() {
		t.Fatal("expected diag.HasSecondaryTimeScale() to be true")
	}

	// 2. Verify Diagram Height
	// yTimeLine = YTopMargin(40) + LaneHeight(85) * 1 = 125.0
	// dateMargin = DateYOffset(15) + TextHeight(15) + 5 = 35.0
	// expectedHeight with secondary scale = yTimeLine + 2 * dateMargin = 125.0 + 70.0 = 195.0
	expectedHeight := 195.0
	if math.Abs(diag.Height-expectedHeight) > 0.001 {
		t.Fatalf("expected diagram height %f, got %f", expectedHeight, diag.Height)
	}

	// 3. Inspect SVG Lines
	svgLines := stager.GetSvgStage().GetInstancesSorted[*svg.Line]()
	var foundPrimaryLine, foundSecondaryLine bool
	var secondaryGridLineCount int

	for _, line := range svgLines {
		if line.Name == "Time Line" {
			foundPrimaryLine = true
			if math.Abs(line.Y1-125.0) > 0.001 || math.Abs(line.Y2-125.0) > 0.001 {
				t.Errorf("Time Line Y: got (%f, %f), expected (125.0, 125.0)", line.Y1, line.Y2)
			}
		}
		if line.Name == "Secondary Time Line" {
			foundSecondaryLine = true
			// ySecondaryTimeLine = yTimeLine + dateMargin = 125.0 + 35.0 = 160.0
			if math.Abs(line.Y1-160.0) > 0.001 || math.Abs(line.Y2-160.0) > 0.001 {
				t.Errorf("Secondary Time Line Y: got (%f, %f), expected (160.0, 160.0)", line.Y1, line.Y2)
			}
		}
		if strings.HasPrefix(line.Name, "secondary grid line for ") {
			secondaryGridLineCount++
			if math.Abs(line.Y1-40.0) > 0.001 || math.Abs(line.Y2-125.0) > 0.001 {
				t.Errorf("Secondary grid line Y: got (%f, %f), expected (40.0, 125.0)", line.Y1, line.Y2)
			}
		}
	}

	if !foundPrimaryLine {
		t.Error("Primary 'Time Line' line was not found in SVG")
	}
	if !foundSecondaryLine {
		t.Error("Secondary 'Secondary Time Line' line was not found in SVG")
	}
	if secondaryGridLineCount == 0 {
		t.Error("Expected secondary grid lines when DrawSecondaryVerticalTimeLines is true, but found none")
	}

	// 4. Inspect SVG Texts
	svgTexts := stager.GetSvgStage().GetInstancesSorted[*svg.Text]()
	var primaryTexts, secondaryTexts []*svg.Text
	expectedPrimaryY := 125.0 + 15.0   // 140.0
	expectedSecondaryY := 160.0 + 15.0 // 175.0

	for _, txt := range svgTexts {
		if math.Abs(txt.Y-expectedPrimaryY) < 0.001 {
			primaryTexts = append(primaryTexts, txt)
		}
		if math.Abs(txt.Y-expectedSecondaryY) < 0.001 {
			secondaryTexts = append(secondaryTexts, txt)
		}
	}

	if len(primaryTexts) == 0 {
		t.Error("Expected primary time scale texts at Y=140.0, found none")
	}
	if len(secondaryTexts) == 0 {
		t.Error("Expected secondary time scale texts at Y=175.0, found none")
	}

	// Verify secondary labels are day numbers like "01", "02", etc. (since primary is MONTHS)
	if len(secondaryTexts) > 0 {
		firstSecText := secondaryTexts[0].Content
		if firstSecText != "01" {
			t.Errorf("Expected first secondary label to be '01', got %q", firstSecText)
		}
	}

	// 5. Test disabling secondary time scale
	diag.SecondaryTimeStepScale = models.NONE
	stage.Commit()

	if diag.HasSecondaryTimeScale() {
		t.Error("expected HasSecondaryTimeScale() to be false when set to NONE")
	}

	expectedDisabledHeight := 160.0 // 125.0 + 35.0
	if math.Abs(diag.Height-expectedDisabledHeight) > 0.001 {
		t.Errorf("expected diagram height %f when secondary disabled, got %f", expectedDisabledHeight, diag.Height)
	}

	foundSecondaryLineAfterDisable := false
	for _, line := range stager.GetSvgStage().GetInstancesSorted[*svg.Line]() {
		if line.Name == "Secondary Time Line" {
			foundSecondaryLineAfterDisable = true
		}
	}
	if foundSecondaryLineAfterDisable {
		t.Error("Secondary Time Line should NOT be present when secondary scale is NONE")
	}
}

func TestSecondaryTimeScaleYearsMonths(t *testing.T) {
	stack := level1stack.NewLevel1StackDelta("test_sec_years_months", "", "", true, false, false)
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

	for _, d := range stage.GetInstancesSorted[*models.Diagram]() {
		d.IsChecked = false
	}

	diag := (&models.Diagram{
		Name:                      "GanttYearsMonths",
		IsTimeDiagram:             true,
		IsChecked:                 true,
		IsEditable_:               true,
		TextHeight:                15.0,
		LaneHeight:                85.0,
		YTopMargin:                40.0,
		DateYOffset:               15.0,
		XLeftLanes:                200.0,
		XRightMargin:              1200.0,
		UseManualStartAndEndDates: true,
		ManualStart:               time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		ManualEnd:                 time.Date(2026, 12, 31, 0, 0, 0, 0, time.UTC),
		TimeStep:                  1,
		TimeStepScale:             models.YEARS,
		SecondaryTimeStep:         1,
		SecondaryTimeStepScale:    models.MONTHS,
	}).Stage(stage)
	lib.Diagrams = []*models.Diagram{diag}

	stage.Commit()

	if !diag.HasSecondaryTimeScale() {
		t.Fatal("expected diag.HasSecondaryTimeScale() to be true")
	}

	// Inspect secondary text labels: since primary is YEARS, secondary MONTHS should be "Jan", "Feb", etc.
	expectedSecondaryY := 40.0 + 35.0 + 15.0 // YTopMargin(40) + nbGroups(0)*85 + dateMargin(35) + dateYOffset(15) = 90.0
	var secMonthLabels []string
	for _, txt := range stager.GetSvgStage().GetInstancesSorted[*svg.Text]() {
		if math.Abs(txt.Y-expectedSecondaryY) < 0.001 {
			secMonthLabels = append(secMonthLabels, txt.Content)
		}
	}

	var foundJan, foundFeb, foundMar bool
	for _, label := range secMonthLabels {
		if label == "Jan" {
			foundJan = true
		}
		if label == "Feb" {
			foundFeb = true
		}
		if label == "Mar" {
			foundMar = true
		}
		if strings.Contains(label, "'") {
			t.Errorf("Secondary month label should not contain year apostrophe when primary is YEARS, got %q", label)
		}
	}

	if !foundJan || !foundFeb || !foundMar {
		t.Errorf("Expected secondary month labels to contain 'Jan', 'Feb', 'Mar', got %v", secMonthLabels)
	}
}
