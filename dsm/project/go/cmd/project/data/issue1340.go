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
	__Diagram__00000001_ := (&models.Diagram{Name: `Time Diagram`}).Stage(stage)
	__Diagram__00000002_ := (&models.Diagram{Name: `Hierarchical diagram`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: `Years diagram`}).Stage(stage)
	__Library__00000001_ := (&models.Library{Name: `Montly diagram`}).Stage(stage)

	__Task__00000003_ := (&models.Task{Name: `spring 26`}).Stage(stage)
	__Task__00000004_ := (&models.Task{Name: `summer 26`}).Stage(stage)
	__Task__00000005_ := (&models.Task{Name: `winter 26`}).Stage(stage)
	__Task__00000006_ := (&models.Task{Name: `winter 27`}).Stage(stage)
	__Task__00000007_ := (&models.Task{Name: `Seasons`}).Stage(stage)
	__Task__00000008_ := (&models.Task{Name: `Jan 26`}).Stage(stage)
	__Task__00000009_ := (&models.Task{Name: `Feb 26`}).Stage(stage)
	__Task__00000010_ := (&models.Task{Name: `Mar 26`}).Stage(stage)
	__Task__00000011_ := (&models.Task{Name: `New Year 2027`}).Stage(stage)

	__TaskCompositionShape__00000011_ := (&models.TaskCompositionShape{Name: `Time Diagram-Seasons-spring 26`}).Stage(stage)
	__TaskCompositionShape__00000012_ := (&models.TaskCompositionShape{Name: `Time Diagram-Seasons-summer 26`}).Stage(stage)
	__TaskCompositionShape__00000013_ := (&models.TaskCompositionShape{Name: `Time Diagram-Seasons-winter 27`}).Stage(stage)
	__TaskCompositionShape__00000021_ := (&models.TaskCompositionShape{Name: `Time Diagram-Seasons-New Year 2027`}).Stage(stage)
	__TaskCompositionShape__00000045_ := (&models.TaskCompositionShape{Name: `Time Diagram-Seasons-winter 26`}).Stage(stage)

	__TaskGroup__00000004_ := (&models.TaskGroup{Name: `year 2026`}).Stage(stage)
	__TaskGroup__00000005_ := (&models.TaskGroup{Name: `year 2027`}).Stage(stage)

	__TaskGroupShape__00000003_ := (&models.TaskGroupShape{Name: `Time Diagram-year 2026`}).Stage(stage)
	__TaskGroupShape__00000005_ := (&models.TaskGroupShape{Name: `Hierarchical diagram-year 2026`}).Stage(stage)
	__TaskGroupShape__00000006_ := (&models.TaskGroupShape{Name: `Hierarchical diagram-year 2027`}).Stage(stage)
	__TaskGroupShape__00000007_ := (&models.TaskGroupShape{Name: `Time Diagram-year 2027`}).Stage(stage)

	__TaskShape__00000019_ := (&models.TaskShape{Name: `Time Diagram-Seasons`}).Stage(stage)
	__TaskShape__00000021_ := (&models.TaskShape{Name: `Time Diagram-spring 26`}).Stage(stage)
	__TaskShape__00000022_ := (&models.TaskShape{Name: `Time Diagram-summer 26`}).Stage(stage)
	__TaskShape__00000023_ := (&models.TaskShape{Name: `Time Diagram-winter 27`}).Stage(stage)
	__TaskShape__00000032_ := (&models.TaskShape{Name: `Time Diagram-New Year 2027`}).Stage(stage)
	__TaskShape__00000054_ := (&models.TaskShape{Name: `Time Diagram-Jan 26`}).Stage(stage)
	__TaskShape__00000056_ := (&models.TaskShape{Name: `Time Diagram-Mar 26`}).Stage(stage)
	__TaskShape__00000057_ := (&models.TaskShape{Name: `Time Diagram-Feb 26`}).Stage(stage)
	__TaskShape__00000058_ := (&models.TaskShape{Name: `Time Diagram-winter 26`}).Stage(stage)

	// insertion point for initialization of values

	__Diagram__00000000_.Name = `Default Diagram`
	__Diagram__00000000_.ComputedPrefix = `1`
	__Diagram__00000000_.IsExpanded = true
	__Diagram__00000000_.LayoutDirection = models.Vertical
	__Diagram__00000000_.IsChecked = false
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.IsShowPrefix = false
	__Diagram__00000000_.DefaultBoxWidth = 180.000000
	__Diagram__00000000_.DefaultBoxHeigth = 50.000000
	__Diagram__00000000_.Width = 1100.000000
	__Diagram__00000000_.Height = 155.000000
	__Diagram__00000000_.IsPBSNodeExpanded = true
	__Diagram__00000000_.IsWBSNodeExpanded = true
	__Diagram__00000000_.IsTaskGroupsNodeExpanded = true
	__Diagram__00000000_.DateFormat = `2006`
	__Diagram__00000000_.IsNotesNodeExpanded = false
	__Diagram__00000000_.IsResourcesNodeExpanded = false
	__Diagram__00000000_.IsTimeDiagram = true
	__Diagram__00000000_.ComputedStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2021-01-01 00:00:00 +0000 UTC")
	__Diagram__00000000_.ComputedEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2022-01-01 00:00:00 +0000 UTC")
	__Diagram__00000000_.ComputedDuration = 31536000000000000
	__Diagram__00000000_.UseManualStartAndEndDates = false
	__Diagram__00000000_.ManualStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000000_.ManualEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000000_.TimeStep = 5
	__Diagram__00000000_.TimeStepScale = models.YEARS
	__Diagram__00000000_.LaneHeight = 100.000000
	__Diagram__00000000_.RatioBarToLaneHeight = 0.800000
	__Diagram__00000000_.YTopMargin = 40.000000
	__Diagram__00000000_.XLeftText = 10.000000
	__Diagram__00000000_.TextHeight = 15.000000
	__Diagram__00000000_.XLeftLanes = 200.000000
	__Diagram__00000000_.XRightMargin = 1000.000000
	__Diagram__00000000_.ArrowLengthToTheRightOfStartBar = 15.000000
	__Diagram__00000000_.ArrowTipLenght = 5.000000
	__Diagram__00000000_.TimeLine_Color = `grey`
	__Diagram__00000000_.TimeLine_FillOpacity = 0.100000
	__Diagram__00000000_.TimeLine_Stroke = `grey`
	__Diagram__00000000_.TimeLine_StrokeWidth = 1.000000
	__Diagram__00000000_.DrawVerticalTimeLines = false
	__Diagram__00000000_.Group_Stroke = `black`
	__Diagram__00000000_.Group_StrokeWidth = 1.000000
	__Diagram__00000000_.Group_StrokeDashArray = `2 2`
	__Diagram__00000000_.DateYOffset = 15.000000
	__Diagram__00000000_.AlignOnStartEndOnYearStart = false

	__Diagram__00000001_.Name = `Time Diagram`
	__Diagram__00000001_.ComputedPrefix = `1`
	__Diagram__00000001_.IsExpanded = true
	__Diagram__00000001_.LayoutDirection = models.Vertical
	__Diagram__00000001_.IsChecked = true
	__Diagram__00000001_.IsEditable_ = true
	__Diagram__00000001_.IsShowPrefix = false
	__Diagram__00000001_.DefaultBoxWidth = 250.000000
	__Diagram__00000001_.DefaultBoxHeigth = 70.000000
	__Diagram__00000001_.Width = 2284.246143
	__Diagram__00000001_.Height = 574.124452
	__Diagram__00000001_.IsPBSNodeExpanded = false
	__Diagram__00000001_.IsWBSNodeExpanded = true
	__Diagram__00000001_.IsTaskGroupsNodeExpanded = true
	__Diagram__00000001_.DateFormat = ``
	__Diagram__00000001_.IsNotesNodeExpanded = false
	__Diagram__00000001_.IsResourcesNodeExpanded = false
	__Diagram__00000001_.IsTimeDiagram = true
	__Diagram__00000001_.ComputedStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-01-01 00:00:00 +0000 UTC")
	__Diagram__00000001_.ComputedEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2027-04-01 00:00:00 +0000 UTC")
	__Diagram__00000001_.ComputedDuration = 39312000000000000
	__Diagram__00000001_.UseManualStartAndEndDates = false
	__Diagram__00000001_.ManualStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000001_.ManualEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000001_.TimeStep = 3
	__Diagram__00000001_.TimeStepScale = models.MONTHS
	__Diagram__00000001_.LaneHeight = 100.000000
	__Diagram__00000001_.RatioBarToLaneHeight = 0.800000
	__Diagram__00000001_.YTopMargin = 40.000000
	__Diagram__00000001_.XLeftText = 10.000000
	__Diagram__00000001_.TextHeight = 15.000000
	__Diagram__00000001_.XLeftLanes = 200.000000
	__Diagram__00000001_.XRightMargin = 1000.000000
	__Diagram__00000001_.ArrowLengthToTheRightOfStartBar = 15.000000
	__Diagram__00000001_.ArrowTipLenght = 5.000000
	__Diagram__00000001_.TimeLine_Color = `grey`
	__Diagram__00000001_.TimeLine_FillOpacity = 0.100000
	__Diagram__00000001_.TimeLine_Stroke = `grey`
	__Diagram__00000001_.TimeLine_StrokeWidth = 1.000000
	__Diagram__00000001_.DrawVerticalTimeLines = true
	__Diagram__00000001_.Group_Stroke = `black`
	__Diagram__00000001_.Group_StrokeWidth = 1.000000
	__Diagram__00000001_.Group_StrokeDashArray = `2 2`
	__Diagram__00000001_.DateYOffset = 15.000000
	__Diagram__00000001_.AlignOnStartEndOnYearStart = false

	__Diagram__00000002_.Name = `Hierarchical diagram`
	__Diagram__00000002_.ComputedPrefix = `2`
	__Diagram__00000002_.IsExpanded = true
	__Diagram__00000002_.LayoutDirection = models.Vertical
	__Diagram__00000002_.IsChecked = false
	__Diagram__00000002_.IsEditable_ = true
	__Diagram__00000002_.IsShowPrefix = false
	__Diagram__00000002_.DefaultBoxWidth = 250.000000
	__Diagram__00000002_.DefaultBoxHeigth = 70.000000
	__Diagram__00000002_.Width = 100.000000
	__Diagram__00000002_.Height = 100.000000
	__Diagram__00000002_.IsPBSNodeExpanded = false
	__Diagram__00000002_.IsWBSNodeExpanded = false
	__Diagram__00000002_.IsTaskGroupsNodeExpanded = false
	__Diagram__00000002_.DateFormat = `01/06`
	__Diagram__00000002_.IsNotesNodeExpanded = false
	__Diagram__00000002_.IsResourcesNodeExpanded = false
	__Diagram__00000002_.IsTimeDiagram = false
	__Diagram__00000002_.ComputedStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-01-01 00:00:00 +0000 UTC")
	__Diagram__00000002_.ComputedEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2027-04-01 00:00:00 +0000 UTC")
	__Diagram__00000002_.ComputedDuration = 39312000000000000
	__Diagram__00000002_.UseManualStartAndEndDates = false
	__Diagram__00000002_.ManualStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000002_.ManualEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000002_.TimeStep = 0
	__Diagram__00000002_.TimeStepScale = ""
	__Diagram__00000002_.LaneHeight = 100.000000
	__Diagram__00000002_.RatioBarToLaneHeight = 0.800000
	__Diagram__00000002_.YTopMargin = 40.000000
	__Diagram__00000002_.XLeftText = 10.000000
	__Diagram__00000002_.TextHeight = 15.000000
	__Diagram__00000002_.XLeftLanes = 200.000000
	__Diagram__00000002_.XRightMargin = 1000.000000
	__Diagram__00000002_.ArrowLengthToTheRightOfStartBar = 15.000000
	__Diagram__00000002_.ArrowTipLenght = 5.000000
	__Diagram__00000002_.TimeLine_Color = `grey`
	__Diagram__00000002_.TimeLine_FillOpacity = 0.100000
	__Diagram__00000002_.TimeLine_Stroke = `grey`
	__Diagram__00000002_.TimeLine_StrokeWidth = 1.000000
	__Diagram__00000002_.DrawVerticalTimeLines = false
	__Diagram__00000002_.Group_Stroke = `black`
	__Diagram__00000002_.Group_StrokeWidth = 1.000000
	__Diagram__00000002_.Group_StrokeDashArray = `2 2`
	__Diagram__00000002_.DateYOffset = 15.000000
	__Diagram__00000002_.AlignOnStartEndOnYearStart = false

	__Library__00000000_.Name = `Years diagram`
	__Library__00000000_.NbPixPerCharacter = 8.000000
	__Library__00000000_.LogoSVGFile = ``
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.IsExpanded = true
	__Library__00000000_.LayoutDirection = models.Vertical
	__Library__00000000_.IsRootLibrary = true

	__Library__00000001_.Name = `Montly diagram`
	__Library__00000001_.NbPixPerCharacter = 0.000000
	__Library__00000001_.LogoSVGFile = ``
	__Library__00000001_.ComputedPrefix = ``
	__Library__00000001_.IsExpanded = true
	__Library__00000001_.LayoutDirection = models.Vertical
	__Library__00000001_.IsRootLibrary = false

	__Task__00000003_.Name = `spring 26`
	__Task__00000003_.ComputedPrefix = `1.2`
	__Task__00000003_.IsExpanded = false
	__Task__00000003_.LayoutDirection = models.Vertical
	__Task__00000003_.IsImport = false
	__Task__00000003_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-04-01 00:00:00 +0000 UTC")
	__Task__00000003_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-07-01 00:00:00 +0000 UTC")
	__Task__00000003_.Duration = 0
	__Task__00000003_.IsEndDateComputedFromDuration = false
	__Task__00000003_.IsStartDateComputedFromPredecessors = false
	__Task__00000003_.IsMilestone = false
	__Task__00000003_.Description = ``
	__Task__00000003_.IsInputsNodeExpanded = false
	__Task__00000003_.IsOutputsNodeExpanded = false
	__Task__00000003_.IsWithCompletion = false
	__Task__00000003_.Completion = ""
	__Task__00000003_.DisplayVerticalBar = false
	__Task__00000003_.TextPosition = ""
	__Task__00000003_.XOffset = 0.000000
	__Task__00000003_.YOffset = 0.000000

	__Task__00000004_.Name = `summer 26`
	__Task__00000004_.ComputedPrefix = `1.3`
	__Task__00000004_.IsExpanded = false
	__Task__00000004_.LayoutDirection = models.Vertical
	__Task__00000004_.IsImport = false
	__Task__00000004_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-07-01 00:00:00 +0000 UTC")
	__Task__00000004_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-10-01 00:00:00 +0000 UTC")
	__Task__00000004_.Duration = 0
	__Task__00000004_.IsEndDateComputedFromDuration = false
	__Task__00000004_.IsStartDateComputedFromPredecessors = false
	__Task__00000004_.IsMilestone = false
	__Task__00000004_.Description = ``
	__Task__00000004_.IsInputsNodeExpanded = false
	__Task__00000004_.IsOutputsNodeExpanded = false
	__Task__00000004_.IsWithCompletion = false
	__Task__00000004_.Completion = ""
	__Task__00000004_.DisplayVerticalBar = false
	__Task__00000004_.TextPosition = ""
	__Task__00000004_.XOffset = 0.000000
	__Task__00000004_.YOffset = 0.000000

	__Task__00000005_.Name = `winter 26`
	__Task__00000005_.ComputedPrefix = `1.1`
	__Task__00000005_.IsExpanded = false
	__Task__00000005_.LayoutDirection = models.Horizontal
	__Task__00000005_.IsImport = false
	__Task__00000005_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-01-01 00:00:00 +0000 UTC")
	__Task__00000005_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-04-01 00:00:00 +0000 UTC")
	__Task__00000005_.Duration = 0
	__Task__00000005_.IsEndDateComputedFromDuration = false
	__Task__00000005_.IsStartDateComputedFromPredecessors = false
	__Task__00000005_.IsMilestone = false
	__Task__00000005_.Description = ``
	__Task__00000005_.IsInputsNodeExpanded = false
	__Task__00000005_.IsOutputsNodeExpanded = false
	__Task__00000005_.IsWithCompletion = false
	__Task__00000005_.Completion = ""
	__Task__00000005_.DisplayVerticalBar = false
	__Task__00000005_.TextPosition = ""
	__Task__00000005_.XOffset = 0.000000
	__Task__00000005_.YOffset = 0.000000

	__Task__00000006_.Name = `winter 27`
	__Task__00000006_.ComputedPrefix = `1.4`
	__Task__00000006_.IsExpanded = false
	__Task__00000006_.LayoutDirection = models.Vertical
	__Task__00000006_.IsImport = false
	__Task__00000006_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2027-01-01 00:00:00 +0000 UTC")
	__Task__00000006_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2027-04-01 00:00:00 +0000 UTC")
	__Task__00000006_.Duration = 0
	__Task__00000006_.IsEndDateComputedFromDuration = false
	__Task__00000006_.IsStartDateComputedFromPredecessors = false
	__Task__00000006_.IsMilestone = false
	__Task__00000006_.Description = ``
	__Task__00000006_.IsInputsNodeExpanded = false
	__Task__00000006_.IsOutputsNodeExpanded = false
	__Task__00000006_.IsWithCompletion = false
	__Task__00000006_.Completion = ""
	__Task__00000006_.DisplayVerticalBar = false
	__Task__00000006_.TextPosition = ""
	__Task__00000006_.XOffset = 15.000000
	__Task__00000006_.YOffset = 0.000000

	__Task__00000007_.Name = `Seasons`
	__Task__00000007_.ComputedPrefix = `1`
	__Task__00000007_.IsExpanded = false
	__Task__00000007_.LayoutDirection = models.Vertical
	__Task__00000007_.IsImport = false
	__Task__00000007_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000007_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000007_.Duration = 0
	__Task__00000007_.IsEndDateComputedFromDuration = false
	__Task__00000007_.IsStartDateComputedFromPredecessors = false
	__Task__00000007_.IsMilestone = false
	__Task__00000007_.Description = ``
	__Task__00000007_.IsInputsNodeExpanded = false
	__Task__00000007_.IsOutputsNodeExpanded = false
	__Task__00000007_.IsWithCompletion = false
	__Task__00000007_.Completion = ""
	__Task__00000007_.DisplayVerticalBar = false
	__Task__00000007_.TextPosition = ""
	__Task__00000007_.XOffset = 0.000000
	__Task__00000007_.YOffset = 0.000000

	__Task__00000008_.Name = `Jan 26`
	__Task__00000008_.ComputedPrefix = `1.1.1`
	__Task__00000008_.IsExpanded = false
	__Task__00000008_.LayoutDirection = models.Vertical
	__Task__00000008_.IsImport = false
	__Task__00000008_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-01-01 00:00:00 +0000 UTC")
	__Task__00000008_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-02-01 00:00:00 +0000 UTC")
	__Task__00000008_.Duration = 0
	__Task__00000008_.IsEndDateComputedFromDuration = false
	__Task__00000008_.IsStartDateComputedFromPredecessors = false
	__Task__00000008_.IsMilestone = false
	__Task__00000008_.Description = ``
	__Task__00000008_.IsInputsNodeExpanded = false
	__Task__00000008_.IsOutputsNodeExpanded = false
	__Task__00000008_.IsWithCompletion = false
	__Task__00000008_.Completion = ""
	__Task__00000008_.DisplayVerticalBar = false
	__Task__00000008_.TextPosition = ""
	__Task__00000008_.XOffset = 0.000000
	__Task__00000008_.YOffset = 0.000000

	__Task__00000009_.Name = `Feb 26`
	__Task__00000009_.ComputedPrefix = `1.1.2`
	__Task__00000009_.IsExpanded = false
	__Task__00000009_.LayoutDirection = models.Vertical
	__Task__00000009_.IsImport = false
	__Task__00000009_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-02-01 00:00:00 +0000 UTC")
	__Task__00000009_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-03-01 00:00:00 +0000 UTC")
	__Task__00000009_.Duration = 0
	__Task__00000009_.IsEndDateComputedFromDuration = false
	__Task__00000009_.IsStartDateComputedFromPredecessors = false
	__Task__00000009_.IsMilestone = false
	__Task__00000009_.Description = ``
	__Task__00000009_.IsInputsNodeExpanded = false
	__Task__00000009_.IsOutputsNodeExpanded = false
	__Task__00000009_.IsWithCompletion = false
	__Task__00000009_.Completion = ""
	__Task__00000009_.DisplayVerticalBar = false
	__Task__00000009_.TextPosition = ""
	__Task__00000009_.XOffset = 0.000000
	__Task__00000009_.YOffset = 0.000000

	__Task__00000010_.Name = `Mar 26`
	__Task__00000010_.ComputedPrefix = `1.1.3`
	__Task__00000010_.IsExpanded = false
	__Task__00000010_.LayoutDirection = models.Vertical
	__Task__00000010_.IsImport = false
	__Task__00000010_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-03-01 00:00:00 +0000 UTC")
	__Task__00000010_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-04-01 00:00:00 +0000 UTC")
	__Task__00000010_.Duration = 0
	__Task__00000010_.IsEndDateComputedFromDuration = false
	__Task__00000010_.IsStartDateComputedFromPredecessors = false
	__Task__00000010_.IsMilestone = false
	__Task__00000010_.Description = ``
	__Task__00000010_.IsInputsNodeExpanded = false
	__Task__00000010_.IsOutputsNodeExpanded = false
	__Task__00000010_.IsWithCompletion = false
	__Task__00000010_.Completion = ""
	__Task__00000010_.DisplayVerticalBar = false
	__Task__00000010_.TextPosition = ""
	__Task__00000010_.XOffset = 0.000000
	__Task__00000010_.YOffset = 0.000000

	__Task__00000011_.Name = `New Year 2027`
	__Task__00000011_.ComputedPrefix = `1.5`
	__Task__00000011_.IsExpanded = false
	__Task__00000011_.LayoutDirection = models.Vertical
	__Task__00000011_.IsImport = false
	__Task__00000011_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-12-24 00:00:00 +0000 UTC")
	__Task__00000011_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-12-24 00:00:00 +0000 UTC")
	__Task__00000011_.Duration = 0
	__Task__00000011_.IsEndDateComputedFromDuration = false
	__Task__00000011_.IsStartDateComputedFromPredecessors = false
	__Task__00000011_.IsMilestone = true
	__Task__00000011_.Description = ``
	__Task__00000011_.IsInputsNodeExpanded = false
	__Task__00000011_.IsOutputsNodeExpanded = false
	__Task__00000011_.IsWithCompletion = false
	__Task__00000011_.Completion = ""
	__Task__00000011_.DisplayVerticalBar = true
	__Task__00000011_.TextPosition = models.TEXT_POSITION_TOP
	__Task__00000011_.XOffset = 20.000000
	__Task__00000011_.YOffset = -17.000000

	__TaskCompositionShape__00000011_.Name = `Time Diagram-Seasons-spring 26`
	__TaskCompositionShape__00000011_.StartRatio = 0.500000
	__TaskCompositionShape__00000011_.EndRatio = 0.500000
	__TaskCompositionShape__00000011_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000011_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000011_.CornerOffsetRatio = 1.680000
	__TaskCompositionShape__00000011_.IsHidden = false

	__TaskCompositionShape__00000012_.Name = `Time Diagram-Seasons-summer 26`
	__TaskCompositionShape__00000012_.StartRatio = 0.500000
	__TaskCompositionShape__00000012_.EndRatio = 0.500000
	__TaskCompositionShape__00000012_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000012_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000012_.CornerOffsetRatio = 1.680000
	__TaskCompositionShape__00000012_.IsHidden = false

	__TaskCompositionShape__00000013_.Name = `Time Diagram-Seasons-winter 27`
	__TaskCompositionShape__00000013_.StartRatio = 0.500000
	__TaskCompositionShape__00000013_.EndRatio = 0.500000
	__TaskCompositionShape__00000013_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000013_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000013_.CornerOffsetRatio = 1.680000
	__TaskCompositionShape__00000013_.IsHidden = false

	__TaskCompositionShape__00000021_.Name = `Time Diagram-Seasons-New Year 2027`
	__TaskCompositionShape__00000021_.StartRatio = 0.500000
	__TaskCompositionShape__00000021_.EndRatio = 0.500000
	__TaskCompositionShape__00000021_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000021_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000021_.CornerOffsetRatio = 1.500000
	__TaskCompositionShape__00000021_.IsHidden = false

	__TaskCompositionShape__00000045_.Name = `Time Diagram-Seasons-winter 26`
	__TaskCompositionShape__00000045_.StartRatio = 0.500000
	__TaskCompositionShape__00000045_.EndRatio = 0.500000
	__TaskCompositionShape__00000045_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000045_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000045_.CornerOffsetRatio = 1.500000
	__TaskCompositionShape__00000045_.IsHidden = false

	__TaskGroup__00000004_.Name = `year 2026`
	__TaskGroup__00000004_.ComputedPrefix = ``
	__TaskGroup__00000004_.IsExpanded = false
	__TaskGroup__00000004_.LayoutDirection = models.Vertical

	__TaskGroup__00000005_.Name = `year 2027`
	__TaskGroup__00000005_.ComputedPrefix = ``
	__TaskGroup__00000005_.IsExpanded = false
	__TaskGroup__00000005_.LayoutDirection = models.Vertical

	__TaskGroupShape__00000003_.Name = `Time Diagram-year 2026`
	__TaskGroupShape__00000003_.X = 171.471356
	__TaskGroupShape__00000003_.Y = 186.904210
	__TaskGroupShape__00000003_.Width = 250.000000
	__TaskGroupShape__00000003_.Height = 70.000000
	__TaskGroupShape__00000003_.IsHidden = false

	__TaskGroupShape__00000005_.Name = `Hierarchical diagram-year 2026`
	__TaskGroupShape__00000005_.X = 171.471356
	__TaskGroupShape__00000005_.Y = 186.904210
	__TaskGroupShape__00000005_.Width = 250.000000
	__TaskGroupShape__00000005_.Height = 70.000000
	__TaskGroupShape__00000005_.IsHidden = false

	__TaskGroupShape__00000006_.Name = `Hierarchical diagram-year 2027`
	__TaskGroupShape__00000006_.X = 128.490669
	__TaskGroupShape__00000006_.Y = 112.566493
	__TaskGroupShape__00000006_.Width = 250.000000
	__TaskGroupShape__00000006_.Height = 70.000000
	__TaskGroupShape__00000006_.IsHidden = false

	__TaskGroupShape__00000007_.Name = `Time Diagram-year 2027`
	__TaskGroupShape__00000007_.X = 150.540984
	__TaskGroupShape__00000007_.Y = 169.888225
	__TaskGroupShape__00000007_.Width = 250.000000
	__TaskGroupShape__00000007_.Height = 70.000000
	__TaskGroupShape__00000007_.IsHidden = false

	__TaskShape__00000019_.Name = `Time Diagram-Seasons`
	__TaskShape__00000019_.IsShowDate = false
	__TaskShape__00000019_.OverideLayoutDirection = false
	__TaskShape__00000019_.LayoutDirection = models.Vertical
	__TaskShape__00000019_.X = 134.246143
	__TaskShape__00000019_.Y = 124.124452
	__TaskShape__00000019_.Width = 250.000000
	__TaskShape__00000019_.Height = 70.000000
	__TaskShape__00000019_.IsHidden = false

	__TaskShape__00000021_.Name = `Time Diagram-spring 26`
	__TaskShape__00000021_.IsShowDate = true
	__TaskShape__00000021_.OverideLayoutDirection = false
	__TaskShape__00000021_.LayoutDirection = models.Vertical
	__TaskShape__00000021_.X = 434.246143
	__TaskShape__00000021_.Y = 264.124452
	__TaskShape__00000021_.Width = 250.000000
	__TaskShape__00000021_.Height = 70.000000
	__TaskShape__00000021_.IsHidden = false

	__TaskShape__00000022_.Name = `Time Diagram-summer 26`
	__TaskShape__00000022_.IsShowDate = true
	__TaskShape__00000022_.OverideLayoutDirection = false
	__TaskShape__00000022_.LayoutDirection = models.Vertical
	__TaskShape__00000022_.X = 734.246143
	__TaskShape__00000022_.Y = 264.124452
	__TaskShape__00000022_.Width = 250.000000
	__TaskShape__00000022_.Height = 70.000000
	__TaskShape__00000022_.IsHidden = false

	__TaskShape__00000023_.Name = `Time Diagram-winter 27`
	__TaskShape__00000023_.IsShowDate = true
	__TaskShape__00000023_.OverideLayoutDirection = false
	__TaskShape__00000023_.LayoutDirection = models.Vertical
	__TaskShape__00000023_.X = 1034.246143
	__TaskShape__00000023_.Y = 264.124452
	__TaskShape__00000023_.Width = 250.000000
	__TaskShape__00000023_.Height = 70.000000
	__TaskShape__00000023_.IsHidden = false

	__TaskShape__00000032_.Name = `Time Diagram-New Year 2027`
	__TaskShape__00000032_.IsShowDate = false
	__TaskShape__00000032_.OverideLayoutDirection = false
	__TaskShape__00000032_.LayoutDirection = models.Vertical
	__TaskShape__00000032_.X = 1334.246143
	__TaskShape__00000032_.Y = 264.124452
	__TaskShape__00000032_.Width = 250.000000
	__TaskShape__00000032_.Height = 70.000000
	__TaskShape__00000032_.IsHidden = false

	__TaskShape__00000054_.Name = `Time Diagram-Jan 26`
	__TaskShape__00000054_.IsShowDate = false
	__TaskShape__00000054_.OverideLayoutDirection = false
	__TaskShape__00000054_.LayoutDirection = models.Vertical
	__TaskShape__00000054_.X = 1934.246143
	__TaskShape__00000054_.Y = 404.124452
	__TaskShape__00000054_.Width = 250.000000
	__TaskShape__00000054_.Height = 70.000000
	__TaskShape__00000054_.IsHidden = false

	__TaskShape__00000056_.Name = `Time Diagram-Mar 26`
	__TaskShape__00000056_.IsShowDate = false
	__TaskShape__00000056_.OverideLayoutDirection = false
	__TaskShape__00000056_.LayoutDirection = models.Vertical
	__TaskShape__00000056_.X = 1934.246143
	__TaskShape__00000056_.Y = 404.124452
	__TaskShape__00000056_.Width = 250.000000
	__TaskShape__00000056_.Height = 70.000000
	__TaskShape__00000056_.IsHidden = false

	__TaskShape__00000057_.Name = `Time Diagram-Feb 26`
	__TaskShape__00000057_.IsShowDate = false
	__TaskShape__00000057_.OverideLayoutDirection = false
	__TaskShape__00000057_.LayoutDirection = models.Vertical
	__TaskShape__00000057_.X = 1934.246143
	__TaskShape__00000057_.Y = 404.124452
	__TaskShape__00000057_.Width = 250.000000
	__TaskShape__00000057_.Height = 70.000000
	__TaskShape__00000057_.IsHidden = false

	__TaskShape__00000058_.Name = `Time Diagram-winter 26`
	__TaskShape__00000058_.IsShowDate = false
	__TaskShape__00000058_.OverideLayoutDirection = false
	__TaskShape__00000058_.LayoutDirection = models.Vertical
	__TaskShape__00000058_.X = 1334.246143
	__TaskShape__00000058_.Y = 264.124452
	__TaskShape__00000058_.Width = 250.000000
	__TaskShape__00000058_.Height = 70.000000
	__TaskShape__00000058_.IsHidden = false

	// insertion point for setup of pointers
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000019_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000021_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000022_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000023_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000032_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000054_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000056_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000057_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000058_)
	__Diagram__00000001_.TasksWhoseNodeIsExpanded = append(__Diagram__00000001_.TasksWhoseNodeIsExpanded, __Task__00000007_)
	__Diagram__00000001_.TaskGroupShapes = append(__Diagram__00000001_.TaskGroupShapes, __TaskGroupShape__00000003_)
	__Diagram__00000001_.TaskGroupShapes = append(__Diagram__00000001_.TaskGroupShapes, __TaskGroupShape__00000007_)
	__Diagram__00000001_.TaskGroupsWhoseNodeIsExpanded = append(__Diagram__00000001_.TaskGroupsWhoseNodeIsExpanded, __TaskGroup__00000004_)
	__Diagram__00000001_.TaskGroupsWhoseNodeIsExpanded = append(__Diagram__00000001_.TaskGroupsWhoseNodeIsExpanded, __TaskGroup__00000005_)
	__Diagram__00000001_.TaskComposition_Shapes = append(__Diagram__00000001_.TaskComposition_Shapes, __TaskCompositionShape__00000011_)
	__Diagram__00000001_.TaskComposition_Shapes = append(__Diagram__00000001_.TaskComposition_Shapes, __TaskCompositionShape__00000012_)
	__Diagram__00000001_.TaskComposition_Shapes = append(__Diagram__00000001_.TaskComposition_Shapes, __TaskCompositionShape__00000013_)
	__Diagram__00000001_.TaskComposition_Shapes = append(__Diagram__00000001_.TaskComposition_Shapes, __TaskCompositionShape__00000021_)
	__Diagram__00000001_.TaskComposition_Shapes = append(__Diagram__00000001_.TaskComposition_Shapes, __TaskCompositionShape__00000045_)
	__Diagram__00000002_.TasksWhoseNodeIsExpanded = append(__Diagram__00000002_.TasksWhoseNodeIsExpanded, __Task__00000007_)
	__Diagram__00000002_.TasksWhoseNodeIsExpanded = append(__Diagram__00000002_.TasksWhoseNodeIsExpanded, __Task__00000005_)
	__Diagram__00000002_.TaskGroupShapes = append(__Diagram__00000002_.TaskGroupShapes, __TaskGroupShape__00000005_)
	__Diagram__00000002_.TaskGroupShapes = append(__Diagram__00000002_.TaskGroupShapes, __TaskGroupShape__00000006_)
	__Library__00000000_.SubLibraries = append(__Library__00000000_.SubLibraries, __Library__00000001_)
	__Library__00000000_.Diagrams = append(__Library__00000000_.Diagrams, __Diagram__00000000_)
	__Library__00000001_.RootTasks = append(__Library__00000001_.RootTasks, __Task__00000007_)
	__Library__00000001_.RootTaskGroups = append(__Library__00000001_.RootTaskGroups, __TaskGroup__00000004_)
	__Library__00000001_.RootTaskGroups = append(__Library__00000001_.RootTaskGroups, __TaskGroup__00000005_)
	__Library__00000001_.Diagrams = append(__Library__00000001_.Diagrams, __Diagram__00000001_)
	__Library__00000001_.Diagrams = append(__Library__00000001_.Diagrams, __Diagram__00000002_)
	__Task__00000003_.ReferencedTask = nil
	__Task__00000004_.ReferencedTask = nil
	__Task__00000005_.ReferencedTask = nil
	__Task__00000005_.SubTasks = append(__Task__00000005_.SubTasks, __Task__00000008_)
	__Task__00000005_.SubTasks = append(__Task__00000005_.SubTasks, __Task__00000009_)
	__Task__00000005_.SubTasks = append(__Task__00000005_.SubTasks, __Task__00000010_)
	__Task__00000006_.ReferencedTask = nil
	__Task__00000007_.ReferencedTask = nil
	__Task__00000007_.SubTasks = append(__Task__00000007_.SubTasks, __Task__00000005_)
	__Task__00000007_.SubTasks = append(__Task__00000007_.SubTasks, __Task__00000003_)
	__Task__00000007_.SubTasks = append(__Task__00000007_.SubTasks, __Task__00000004_)
	__Task__00000007_.SubTasks = append(__Task__00000007_.SubTasks, __Task__00000006_)
	__Task__00000007_.SubTasks = append(__Task__00000007_.SubTasks, __Task__00000011_)
	__Task__00000008_.ReferencedTask = nil
	__Task__00000009_.ReferencedTask = nil
	__Task__00000010_.ReferencedTask = nil
	__Task__00000011_.ReferencedTask = nil
	__Task__00000011_.TaskGroupsToDisplay = append(__Task__00000011_.TaskGroupsToDisplay, __TaskGroup__00000004_)
	__Task__00000011_.TaskGroupsToDisplay = append(__Task__00000011_.TaskGroupsToDisplay, __TaskGroup__00000005_)
	__TaskCompositionShape__00000011_.Task = __Task__00000003_
	__TaskCompositionShape__00000012_.Task = __Task__00000004_
	__TaskCompositionShape__00000013_.Task = __Task__00000006_
	__TaskCompositionShape__00000021_.Task = __Task__00000011_
	__TaskCompositionShape__00000045_.Task = __Task__00000005_
	__TaskGroup__00000004_.Tasks = append(__TaskGroup__00000004_.Tasks, __Task__00000005_)
	__TaskGroup__00000004_.Tasks = append(__TaskGroup__00000004_.Tasks, __Task__00000003_)
	__TaskGroup__00000004_.Tasks = append(__TaskGroup__00000004_.Tasks, __Task__00000004_)
	__TaskGroup__00000004_.Tasks = append(__TaskGroup__00000004_.Tasks, __Task__00000008_)
	__TaskGroup__00000005_.Tasks = append(__TaskGroup__00000005_.Tasks, __Task__00000006_)
	__TaskGroup__00000005_.Tasks = append(__TaskGroup__00000005_.Tasks, __Task__00000008_)
	__TaskGroup__00000005_.Tasks = append(__TaskGroup__00000005_.Tasks, __Task__00000011_)
	__TaskGroupShape__00000003_.TaskGroup = __TaskGroup__00000004_
	__TaskGroupShape__00000005_.TaskGroup = __TaskGroup__00000004_
	__TaskGroupShape__00000006_.TaskGroup = __TaskGroup__00000005_
	__TaskGroupShape__00000007_.TaskGroup = __TaskGroup__00000005_
	__TaskShape__00000019_.Task = __Task__00000007_
	__TaskShape__00000021_.Task = __Task__00000003_
	__TaskShape__00000022_.Task = __Task__00000004_
	__TaskShape__00000023_.Task = __Task__00000006_
	__TaskShape__00000032_.Task = __Task__00000011_
	__TaskShape__00000054_.Task = __Task__00000008_
	__TaskShape__00000056_.Task = __Task__00000010_
	__TaskShape__00000057_.Task = __Task__00000009_
	__TaskShape__00000058_.Task = __Task__00000005_
}
