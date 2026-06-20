package main

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/dsm/project/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var (
	_ time.Time
	_ = slices.Index[[]int, int]
)

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage

	__Diagram__00000000_ := (&models.Diagram{Name: `Default Diagram`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: ``}).Stage(stage)

	__Task__00000000_ := (&models.Task{Name: `winter 26`}).Stage(stage)
	__Task__00000002_ := (&models.Task{Name: `spring 26`}).Stage(stage)

	__TaskShape__00000000_ := (&models.TaskShape{Name: `Default Diagram-winter 26`}).Stage(stage)
	__TaskShape__00000002_ := (&models.TaskShape{Name: `Default Diagram-spring 26`}).Stage(stage)

	// insertion point for initialization of values

	__Diagram__00000000_.Name = `Default Diagram`
	__Diagram__00000000_.ComputedPrefix = `1`
	__Diagram__00000000_.IsExpanded = true
	__Diagram__00000000_.LayoutDirection = models.Vertical
	__Diagram__00000000_.IsChecked = true
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.IsShowPrefix = false
	__Diagram__00000000_.DefaultBoxWidth = 250.000000
	__Diagram__00000000_.DefaultBoxHeigth = 70.000000
	__Diagram__00000000_.Width = 807.847114
	__Diagram__00000000_.Height = 365.504572
	__Diagram__00000000_.IsPBSNodeExpanded = false
	__Diagram__00000000_.IsWBSNodeExpanded = true
	__Diagram__00000000_.IsTaskGroupsNodeExpanded = false
	__Diagram__00000000_.DateFormat = ``
	__Diagram__00000000_.IsNotesNodeExpanded = false
	__Diagram__00000000_.IsResourcesNodeExpanded = false
	__Diagram__00000000_.IsTimeDiagram = false
	__Diagram__00000000_.ComputedStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000000_.ComputedEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000000_.ComputedDuration = 0
	__Diagram__00000000_.UseManualStartAndEndDates = false
	__Diagram__00000000_.ManualStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000000_.ManualEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000000_.TimeStep = 0
	__Diagram__00000000_.TimeStepScale = ""
	__Diagram__00000000_.LaneHeight = 0.000000
	__Diagram__00000000_.RatioBarToLaneHeight = 0.000000
	__Diagram__00000000_.YTopMargin = 0.000000
	__Diagram__00000000_.XLeftText = 0.000000
	__Diagram__00000000_.TextHeight = 0.000000
	__Diagram__00000000_.XLeftLanes = 0.000000
	__Diagram__00000000_.XRightMargin = 0.000000
	__Diagram__00000000_.ArrowLengthToTheRightOfStartBar = 0.000000
	__Diagram__00000000_.ArrowTipLenght = 0.000000
	__Diagram__00000000_.TimeLine_Color = ``
	__Diagram__00000000_.TimeLine_FillOpacity = 0.000000
	__Diagram__00000000_.TimeLine_Stroke = ``
	__Diagram__00000000_.TimeLine_StrokeWidth = 0.000000
	__Diagram__00000000_.DrawVerticalTimeLines = false
	__Diagram__00000000_.Group_Stroke = ``
	__Diagram__00000000_.Group_StrokeWidth = 0.000000
	__Diagram__00000000_.Group_StrokeDashArray = ``
	__Diagram__00000000_.DateYOffset = 0.000000
	__Diagram__00000000_.AlignOnStartEndOnYearStart = false

	__Library__00000000_.Name = ``
	__Library__00000000_.NbPixPerCharacter = 8.000000
	__Library__00000000_.LogoSVGFile = ``
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.IsExpanded = true
	__Library__00000000_.LayoutDirection = models.Vertical
	__Library__00000000_.IsRootLibrary = true

	__Task__00000000_.Name = `winter 26`
	__Task__00000000_.ComputedPrefix = `1`
	__Task__00000000_.IsExpanded = false
	__Task__00000000_.LayoutDirection = models.Vertical
	__Task__00000000_.IsImport = false
	__Task__00000000_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-01-01 00:00:00 +0000 UTC")
	__Task__00000000_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-04-01 00:00:00 +0000 UTC")
	__Task__00000000_.DurationYears = 0.000000
	__Task__00000000_.DurationMonths = 3.000000
	__Task__00000000_.DurationWeeks = 0.000000
	__Task__00000000_.DurationDays = 0.000000
	__Task__00000000_.DurationHours = 0.000000
	__Task__00000000_.IsEndDateComputedFromDuration = true
	__Task__00000000_.IsStartDateComputedFromPredecessors = false
	__Task__00000000_.IsMilestone = false
	__Task__00000000_.Description = ``
	__Task__00000000_.IsInputsNodeExpanded = false
	__Task__00000000_.IsOutputsNodeExpanded = false
	__Task__00000000_.IsWithCompletion = false
	__Task__00000000_.Completion = ""
	__Task__00000000_.DisplayVerticalBar = false
	__Task__00000000_.TextPosition = ""
	__Task__00000000_.XOffset = 0.000000
	__Task__00000000_.YOffset = 0.000000

	__Task__00000002_.Name = `spring 26`
	__Task__00000002_.ComputedPrefix = `2`
	__Task__00000002_.IsExpanded = false
	__Task__00000002_.LayoutDirection = models.Vertical
	__Task__00000002_.IsImport = false
	__Task__00000002_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-04-01 00:00:00 +0000 UTC")
	__Task__00000002_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-07-01 00:00:00 +0000 UTC")
	__Task__00000002_.DurationYears = 0.000000
	__Task__00000002_.DurationMonths = 3.000000
	__Task__00000002_.DurationWeeks = 0.000000
	__Task__00000002_.DurationDays = 0.000000
	__Task__00000002_.DurationHours = 0.000000
	__Task__00000002_.IsEndDateComputedFromDuration = true
	__Task__00000002_.IsStartDateComputedFromPredecessors = true
	__Task__00000002_.IsMilestone = false
	__Task__00000002_.Description = ``
	__Task__00000002_.IsInputsNodeExpanded = false
	__Task__00000002_.IsOutputsNodeExpanded = false
	__Task__00000002_.IsWithCompletion = false
	__Task__00000002_.Completion = ""
	__Task__00000002_.DisplayVerticalBar = false
	__Task__00000002_.TextPosition = ""
	__Task__00000002_.XOffset = 0.000000
	__Task__00000002_.YOffset = 0.000000

	__TaskShape__00000000_.Name = `Default Diagram-winter 26`
	__TaskShape__00000000_.IsShowDate = true
	__TaskShape__00000000_.OverideLayoutDirection = false
	__TaskShape__00000000_.LayoutDirection = models.Vertical
	__TaskShape__00000000_.X = 139.724129
	__TaskShape__00000000_.Y = 193.067462
	__TaskShape__00000000_.Width = 250.000000
	__TaskShape__00000000_.Height = 70.000000
	__TaskShape__00000000_.IsHidden = false

	__TaskShape__00000002_.Name = `Default Diagram-spring 26`
	__TaskShape__00000002_.IsShowDate = true
	__TaskShape__00000002_.OverideLayoutDirection = false
	__TaskShape__00000002_.LayoutDirection = models.Vertical
	__TaskShape__00000002_.X = 457.847114
	__TaskShape__00000002_.Y = 195.504572
	__TaskShape__00000002_.Width = 250.000000
	__TaskShape__00000002_.Height = 70.000000
	__TaskShape__00000002_.IsHidden = false

	// insertion point for setup of pointers
	__Diagram__00000000_.Task_Shapes = append(__Diagram__00000000_.Task_Shapes, __TaskShape__00000000_)
	__Diagram__00000000_.Task_Shapes = append(__Diagram__00000000_.Task_Shapes, __TaskShape__00000002_)
	__Library__00000000_.RootTasks = append(__Library__00000000_.RootTasks, __Task__00000000_)
	__Library__00000000_.RootTasks = append(__Library__00000000_.RootTasks, __Task__00000002_)
	__Library__00000000_.Diagrams = append(__Library__00000000_.Diagrams, __Diagram__00000000_)
	__Task__00000000_.ReferencedTask = nil
	__Task__00000002_.ReferencedTask = nil
	__Task__00000002_.Predecessors = append(__Task__00000002_.Predecessors, __Task__00000000_)
	__TaskShape__00000000_.Task = __Task__00000000_
	__TaskShape__00000002_.Task = __Task__00000002_
}
