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
	__Diagram__00000001_ := (&models.Diagram{Name: `Default Diagram`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: `Years diagram`}).Stage(stage)
	__Library__00000001_ := (&models.Library{Name: `Montly diagram`}).Stage(stage)

	__Milestone__00000000_ := (&models.Milestone{Name: `M1`}).Stage(stage)

	__MilestoneShape__00000000_ := (&models.MilestoneShape{Name: `-Default Diagram`}).Stage(stage)

	__Product__00000000_ := (&models.Product{Name: `P1`}).Stage(stage)
	__Product__00000001_ := (&models.Product{Name: `P1.1`}).Stage(stage)
	__Product__00000002_ := (&models.Product{Name: `P1.2`}).Stage(stage)
	__Product__00000003_ := (&models.Product{Name: `P1.3`}).Stage(stage)
	__Product__00000004_ := (&models.Product{Name: `P1.3.1`}).Stage(stage)
	__Product__00000005_ := (&models.Product{Name: `P1.2.1`}).Stage(stage)
	__Product__00000006_ := (&models.Product{Name: `P1.2.2`}).Stage(stage)

	__ProductCompositionShape__00000015_ := (&models.ProductCompositionShape{Name: `P1 to P1.1`}).Stage(stage)
	__ProductCompositionShape__00000016_ := (&models.ProductCompositionShape{Name: `P1 to P1.3`}).Stage(stage)
	__ProductCompositionShape__00000017_ := (&models.ProductCompositionShape{Name: `P1.3 to P1.3.1`}).Stage(stage)
	__ProductCompositionShape__00000018_ := (&models.ProductCompositionShape{Name: `P1 to P1.2`}).Stage(stage)
	__ProductCompositionShape__00000019_ := (&models.ProductCompositionShape{Name: `P1.2 to P1.2.1`}).Stage(stage)
	__ProductCompositionShape__00000020_ := (&models.ProductCompositionShape{Name: `P1.2 to P1.2.2`}).Stage(stage)

	__ProductShape__00000018_ := (&models.ProductShape{Name: `P1-Default Diagram`}).Stage(stage)
	__ProductShape__00000019_ := (&models.ProductShape{Name: `P1.1-Default Diagram`}).Stage(stage)
	__ProductShape__00000020_ := (&models.ProductShape{Name: `P1.3-Default Diagram`}).Stage(stage)
	__ProductShape__00000021_ := (&models.ProductShape{Name: `P1.3.1-Default Diagram`}).Stage(stage)
	__ProductShape__00000022_ := (&models.ProductShape{Name: `P1.2-Default Diagram`}).Stage(stage)
	__ProductShape__00000023_ := (&models.ProductShape{Name: `P1.2.1-Default Diagram`}).Stage(stage)
	__ProductShape__00000024_ := (&models.ProductShape{Name: `P1.2.2-Default Diagram`}).Stage(stage)

	__Task__00000000_ := (&models.Task{Name: `L1.W1`}).Stage(stage)
	__Task__00000001_ := (&models.Task{Name: `L1.W2`}).Stage(stage)
	__Task__00000002_ := (&models.Task{Name: `L2.W1`}).Stage(stage)
	__Task__00000003_ := (&models.Task{Name: `spring 26`}).Stage(stage)
	__Task__00000004_ := (&models.Task{Name: `summer 26`}).Stage(stage)
	__Task__00000005_ := (&models.Task{Name: `winter 26`}).Stage(stage)
	__Task__00000006_ := (&models.Task{Name: `winter 27`}).Stage(stage)
	__Task__00000007_ := (&models.Task{Name: `Seasons`}).Stage(stage)

	__TaskCompositionShape__00000010_ := (&models.TaskCompositionShape{Name: `Seasons to winter 26`}).Stage(stage)
	__TaskCompositionShape__00000011_ := (&models.TaskCompositionShape{Name: `Seasons to spring 26`}).Stage(stage)
	__TaskCompositionShape__00000012_ := (&models.TaskCompositionShape{Name: `Seasons to summer 26`}).Stage(stage)
	__TaskCompositionShape__00000013_ := (&models.TaskCompositionShape{Name: `Seasons to winter 27`}).Stage(stage)

	__TaskGroup__00000000_ := (&models.TaskGroup{Name: `L1`}).Stage(stage)
	__TaskGroup__00000002_ := (&models.TaskGroup{Name: `L2`}).Stage(stage)
	__TaskGroup__00000003_ := (&models.TaskGroup{Name: `L3`}).Stage(stage)
	__TaskGroup__00000004_ := (&models.TaskGroup{Name: `year 2026`}).Stage(stage)
	__TaskGroup__00000005_ := (&models.TaskGroup{Name: `year 2027`}).Stage(stage)

	__TaskGroupShape__00000003_ := (&models.TaskGroupShape{Name: `-Default Diagram`}).Stage(stage)
	__TaskGroupShape__00000004_ := (&models.TaskGroupShape{Name: `-Default Diagram`}).Stage(stage)

	__TaskShape__00000019_ := (&models.TaskShape{Name: `Seasons-Default Diagram`}).Stage(stage)
	__TaskShape__00000020_ := (&models.TaskShape{Name: `winter 26-Default Diagram`}).Stage(stage)
	__TaskShape__00000021_ := (&models.TaskShape{Name: `spring 26-Default Diagram`}).Stage(stage)
	__TaskShape__00000022_ := (&models.TaskShape{Name: `summer 26-Default Diagram`}).Stage(stage)
	__TaskShape__00000023_ := (&models.TaskShape{Name: `winter 27-Default Diagram`}).Stage(stage)

	// insertion point for initialization of values

	__Diagram__00000000_.Name = `Default Diagram`
	__Diagram__00000000_.ComputedPrefix = `1`
	__Diagram__00000000_.IsExpanded = true
	__Diagram__00000000_.IsChecked = true
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.IsShowPrefix = false
	__Diagram__00000000_.DefaultBoxWidth = 250.000000
	__Diagram__00000000_.DefaultBoxHeigth = 70.000000
	__Diagram__00000000_.Width = 1175.000000
	__Diagram__00000000_.Height = 600.000000
	__Diagram__00000000_.IsPBSNodeExpanded = true
	__Diagram__00000000_.IsWBSNodeExpanded = false
	__Diagram__00000000_.IsTaskGroupsNodeExpanded = true
	__Diagram__00000000_.IsMilestonesNodeExpanded = false
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

	__Diagram__00000001_.Name = `Default Diagram`
	__Diagram__00000001_.ComputedPrefix = `1`
	__Diagram__00000001_.IsExpanded = true
	__Diagram__00000001_.IsChecked = false
	__Diagram__00000001_.IsEditable_ = true
	__Diagram__00000001_.IsShowPrefix = false
	__Diagram__00000001_.DefaultBoxWidth = 250.000000
	__Diagram__00000001_.DefaultBoxHeigth = 70.000000
	__Diagram__00000001_.Width = 1384.246143
	__Diagram__00000001_.Height = 434.124452
	__Diagram__00000001_.IsPBSNodeExpanded = false
	__Diagram__00000001_.IsWBSNodeExpanded = false
	__Diagram__00000001_.IsTaskGroupsNodeExpanded = true
	__Diagram__00000001_.IsMilestonesNodeExpanded = false
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

	__Library__00000000_.Name = `Years diagram`
	__Library__00000000_.NbPixPerCharacter = 8.000000
	__Library__00000000_.LogoSVGFile = ``
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.IsExpanded = true
	__Library__00000000_.IsRootLibrary = true

	__Library__00000001_.Name = `Montly diagram`
	__Library__00000001_.NbPixPerCharacter = 0.000000
	__Library__00000001_.LogoSVGFile = ``
	__Library__00000001_.ComputedPrefix = ``
	__Library__00000001_.IsExpanded = true
	__Library__00000001_.IsRootLibrary = false

	__Milestone__00000000_.Name = `M1`
	__Milestone__00000000_.ComputedPrefix = ``
	__Milestone__00000000_.IsExpanded = false
	__Milestone__00000000_.Date, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2015-01-01 00:00:00 +0000 UTC")
	__Milestone__00000000_.DisplayVerticalBar = false

	__MilestoneShape__00000000_.Name = `-Default Diagram`
	__MilestoneShape__00000000_.X = 105.703445
	__MilestoneShape__00000000_.Y = 160.791992
	__MilestoneShape__00000000_.Width = 250.000000
	__MilestoneShape__00000000_.Height = 70.000000
	__MilestoneShape__00000000_.IsHidden = false

	__Product__00000000_.Name = `P1`
	__Product__00000000_.ComputedPrefix = `1`
	__Product__00000000_.IsExpanded = false
	__Product__00000000_.IsImport = false
	__Product__00000000_.Description = ``
	__Product__00000000_.IsProducersNodeExpanded = false
	__Product__00000000_.IsConsumersNodeExpanded = false

	__Product__00000001_.Name = `P1.1`
	__Product__00000001_.ComputedPrefix = `1.1`
	__Product__00000001_.IsExpanded = false
	__Product__00000001_.IsImport = false
	__Product__00000001_.Description = ``
	__Product__00000001_.IsProducersNodeExpanded = false
	__Product__00000001_.IsConsumersNodeExpanded = false

	__Product__00000002_.Name = `P1.2`
	__Product__00000002_.ComputedPrefix = `1.2`
	__Product__00000002_.IsExpanded = false
	__Product__00000002_.IsImport = false
	__Product__00000002_.Description = ``
	__Product__00000002_.IsProducersNodeExpanded = false
	__Product__00000002_.IsConsumersNodeExpanded = false

	__Product__00000003_.Name = `P1.3`
	__Product__00000003_.ComputedPrefix = `1.3`
	__Product__00000003_.IsExpanded = false
	__Product__00000003_.IsImport = false
	__Product__00000003_.Description = ``
	__Product__00000003_.IsProducersNodeExpanded = false
	__Product__00000003_.IsConsumersNodeExpanded = false

	__Product__00000004_.Name = `P1.3.1`
	__Product__00000004_.ComputedPrefix = `1.3.1`
	__Product__00000004_.IsExpanded = false
	__Product__00000004_.IsImport = false
	__Product__00000004_.Description = ``
	__Product__00000004_.IsProducersNodeExpanded = false
	__Product__00000004_.IsConsumersNodeExpanded = false

	__Product__00000005_.Name = `P1.2.1`
	__Product__00000005_.ComputedPrefix = `1.2.1`
	__Product__00000005_.IsExpanded = false
	__Product__00000005_.IsImport = false
	__Product__00000005_.Description = ``
	__Product__00000005_.IsProducersNodeExpanded = false
	__Product__00000005_.IsConsumersNodeExpanded = false

	__Product__00000006_.Name = `P1.2.2`
	__Product__00000006_.ComputedPrefix = `1.2.2`
	__Product__00000006_.IsExpanded = false
	__Product__00000006_.IsImport = false
	__Product__00000006_.Description = ``
	__Product__00000006_.IsProducersNodeExpanded = false
	__Product__00000006_.IsConsumersNodeExpanded = false

	__ProductCompositionShape__00000015_.Name = `P1 to P1.1`
	__ProductCompositionShape__00000015_.StartRatio = 0.450562
	__ProductCompositionShape__00000015_.EndRatio = 0.371512
	__ProductCompositionShape__00000015_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000015_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000015_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000015_.IsHidden = false

	__ProductCompositionShape__00000016_.Name = `P1 to P1.3`
	__ProductCompositionShape__00000016_.StartRatio = 0.450562
	__ProductCompositionShape__00000016_.EndRatio = 0.428655
	__ProductCompositionShape__00000016_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000016_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000016_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000016_.IsHidden = false

	__ProductCompositionShape__00000017_.Name = `P1.3 to P1.3.1`
	__ProductCompositionShape__00000017_.StartRatio = 0.500000
	__ProductCompositionShape__00000017_.EndRatio = 0.500000
	__ProductCompositionShape__00000017_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000017_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000017_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000017_.IsHidden = false

	__ProductCompositionShape__00000018_.Name = `P1 to P1.2`
	__ProductCompositionShape__00000018_.StartRatio = 0.454562
	__ProductCompositionShape__00000018_.EndRatio = 0.500084
	__ProductCompositionShape__00000018_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000018_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000018_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000018_.IsHidden = false

	__ProductCompositionShape__00000019_.Name = `P1.2 to P1.2.1`
	__ProductCompositionShape__00000019_.StartRatio = 0.500000
	__ProductCompositionShape__00000019_.EndRatio = 0.500000
	__ProductCompositionShape__00000019_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000019_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000019_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000019_.IsHidden = false

	__ProductCompositionShape__00000020_.Name = `P1.2 to P1.2.2`
	__ProductCompositionShape__00000020_.StartRatio = 0.500000
	__ProductCompositionShape__00000020_.EndRatio = 0.500000
	__ProductCompositionShape__00000020_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000020_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000020_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000020_.IsHidden = false

	__ProductShape__00000018_.Name = `P1-Default Diagram`
	__ProductShape__00000018_.LayoutDirection = models.Vertical
	__ProductShape__00000018_.X = 50.000000
	__ProductShape__00000018_.Y = 50.000000
	__ProductShape__00000018_.Width = 250.000000
	__ProductShape__00000018_.Height = 70.000000
	__ProductShape__00000018_.IsHidden = false

	__ProductShape__00000019_.Name = `P1.1-Default Diagram`
	__ProductShape__00000019_.LayoutDirection = models.Vertical
	__ProductShape__00000019_.X = 50.000000
	__ProductShape__00000019_.Y = 190.000000
	__ProductShape__00000019_.Width = 250.000000
	__ProductShape__00000019_.Height = 70.000000
	__ProductShape__00000019_.IsHidden = false

	__ProductShape__00000020_.Name = `P1.3-Default Diagram`
	__ProductShape__00000020_.LayoutDirection = models.Vertical
	__ProductShape__00000020_.X = 825.000000
	__ProductShape__00000020_.Y = 190.000000
	__ProductShape__00000020_.Width = 250.000000
	__ProductShape__00000020_.Height = 70.000000
	__ProductShape__00000020_.IsHidden = false

	__ProductShape__00000021_.Name = `P1.3.1-Default Diagram`
	__ProductShape__00000021_.LayoutDirection = models.Vertical
	__ProductShape__00000021_.X = 825.000000
	__ProductShape__00000021_.Y = 330.000000
	__ProductShape__00000021_.Width = 250.000000
	__ProductShape__00000021_.Height = 70.000000
	__ProductShape__00000021_.IsHidden = false

	__ProductShape__00000022_.Name = `P1.2-Default Diagram`
	__ProductShape__00000022_.LayoutDirection = models.Horizontal
	__ProductShape__00000022_.X = 350.000000
	__ProductShape__00000022_.Y = 190.000000
	__ProductShape__00000022_.Width = 250.000000
	__ProductShape__00000022_.Height = 70.000000
	__ProductShape__00000022_.IsHidden = false

	__ProductShape__00000023_.Name = `P1.2.1-Default Diagram`
	__ProductShape__00000023_.LayoutDirection = models.Vertical
	__ProductShape__00000023_.X = 525.000000
	__ProductShape__00000023_.Y = 310.000000
	__ProductShape__00000023_.Width = 250.000000
	__ProductShape__00000023_.Height = 70.000000
	__ProductShape__00000023_.IsHidden = false

	__ProductShape__00000024_.Name = `P1.2.2-Default Diagram`
	__ProductShape__00000024_.LayoutDirection = models.Vertical
	__ProductShape__00000024_.X = 525.000000
	__ProductShape__00000024_.Y = 430.000000
	__ProductShape__00000024_.Width = 250.000000
	__ProductShape__00000024_.Height = 70.000000
	__ProductShape__00000024_.IsHidden = false

	__Task__00000000_.Name = `L1.W1`
	__Task__00000000_.ComputedPrefix = `1`
	__Task__00000000_.IsExpanded = false
	__Task__00000000_.IsImport = false
	__Task__00000000_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2000-01-01 00:00:00 +0000 UTC")
	__Task__00000000_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2030-01-01 00:00:00 +0000 UTC")
	__Task__00000000_.Description = ``
	__Task__00000000_.IsInputsNodeExpanded = false
	__Task__00000000_.IsOutputsNodeExpanded = false
	__Task__00000000_.IsWithCompletion = false
	__Task__00000000_.Completion = ""

	__Task__00000001_.Name = `L1.W2`
	__Task__00000001_.ComputedPrefix = `2`
	__Task__00000001_.IsExpanded = false
	__Task__00000001_.IsImport = false
	__Task__00000001_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2020-01-01 00:00:00 +0000 UTC")
	__Task__00000001_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2025-01-01 00:00:00 +0000 UTC")
	__Task__00000001_.Description = ``
	__Task__00000001_.IsInputsNodeExpanded = false
	__Task__00000001_.IsOutputsNodeExpanded = false
	__Task__00000001_.IsWithCompletion = false
	__Task__00000001_.Completion = ""

	__Task__00000002_.Name = `L2.W1`
	__Task__00000002_.ComputedPrefix = `3`
	__Task__00000002_.IsExpanded = false
	__Task__00000002_.IsImport = false
	__Task__00000002_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2021-01-01 00:00:00 +0000 UTC")
	__Task__00000002_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2022-01-01 00:00:00 +0000 UTC")
	__Task__00000002_.Description = ``
	__Task__00000002_.IsInputsNodeExpanded = false
	__Task__00000002_.IsOutputsNodeExpanded = false
	__Task__00000002_.IsWithCompletion = false
	__Task__00000002_.Completion = ""

	__Task__00000003_.Name = `spring 26`
	__Task__00000003_.ComputedPrefix = `1.2`
	__Task__00000003_.IsExpanded = false
	__Task__00000003_.IsImport = false
	__Task__00000003_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-04-01 00:00:00 +0000 UTC")
	__Task__00000003_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-07-01 00:00:00 +0000 UTC")
	__Task__00000003_.Description = ``
	__Task__00000003_.IsInputsNodeExpanded = false
	__Task__00000003_.IsOutputsNodeExpanded = false
	__Task__00000003_.IsWithCompletion = false
	__Task__00000003_.Completion = ""

	__Task__00000004_.Name = `summer 26`
	__Task__00000004_.ComputedPrefix = `1.3`
	__Task__00000004_.IsExpanded = false
	__Task__00000004_.IsImport = false
	__Task__00000004_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-07-01 00:00:00 +0000 UTC")
	__Task__00000004_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-10-01 00:00:00 +0000 UTC")
	__Task__00000004_.Description = ``
	__Task__00000004_.IsInputsNodeExpanded = false
	__Task__00000004_.IsOutputsNodeExpanded = false
	__Task__00000004_.IsWithCompletion = false
	__Task__00000004_.Completion = ""

	__Task__00000005_.Name = `winter 26`
	__Task__00000005_.ComputedPrefix = `1.1`
	__Task__00000005_.IsExpanded = false
	__Task__00000005_.IsImport = false
	__Task__00000005_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-01-01 00:00:00 +0000 UTC")
	__Task__00000005_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-04-01 00:00:00 +0000 UTC")
	__Task__00000005_.Description = ``
	__Task__00000005_.IsInputsNodeExpanded = false
	__Task__00000005_.IsOutputsNodeExpanded = false
	__Task__00000005_.IsWithCompletion = false
	__Task__00000005_.Completion = ""

	__Task__00000006_.Name = `winter 27`
	__Task__00000006_.ComputedPrefix = `1.4`
	__Task__00000006_.IsExpanded = false
	__Task__00000006_.IsImport = false
	__Task__00000006_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2027-01-01 00:00:00 +0000 UTC")
	__Task__00000006_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2027-04-01 00:00:00 +0000 UTC")
	__Task__00000006_.Description = ``
	__Task__00000006_.IsInputsNodeExpanded = false
	__Task__00000006_.IsOutputsNodeExpanded = false
	__Task__00000006_.IsWithCompletion = false
	__Task__00000006_.Completion = ""

	__Task__00000007_.Name = `Seasons`
	__Task__00000007_.ComputedPrefix = `1`
	__Task__00000007_.IsExpanded = false
	__Task__00000007_.IsImport = false
	__Task__00000007_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000007_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000007_.Description = ``
	__Task__00000007_.IsInputsNodeExpanded = false
	__Task__00000007_.IsOutputsNodeExpanded = false
	__Task__00000007_.IsWithCompletion = false
	__Task__00000007_.Completion = ""

	__TaskCompositionShape__00000010_.Name = `Seasons to winter 26`
	__TaskCompositionShape__00000010_.StartRatio = 0.500000
	__TaskCompositionShape__00000010_.EndRatio = 0.500000
	__TaskCompositionShape__00000010_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000010_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000010_.CornerOffsetRatio = 1.680000
	__TaskCompositionShape__00000010_.IsHidden = false

	__TaskCompositionShape__00000011_.Name = `Seasons to spring 26`
	__TaskCompositionShape__00000011_.StartRatio = 0.500000
	__TaskCompositionShape__00000011_.EndRatio = 0.500000
	__TaskCompositionShape__00000011_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000011_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000011_.CornerOffsetRatio = 1.680000
	__TaskCompositionShape__00000011_.IsHidden = false

	__TaskCompositionShape__00000012_.Name = `Seasons to summer 26`
	__TaskCompositionShape__00000012_.StartRatio = 0.500000
	__TaskCompositionShape__00000012_.EndRatio = 0.500000
	__TaskCompositionShape__00000012_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000012_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000012_.CornerOffsetRatio = 1.680000
	__TaskCompositionShape__00000012_.IsHidden = false

	__TaskCompositionShape__00000013_.Name = `Seasons to winter 27`
	__TaskCompositionShape__00000013_.StartRatio = 0.500000
	__TaskCompositionShape__00000013_.EndRatio = 0.500000
	__TaskCompositionShape__00000013_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000013_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000013_.CornerOffsetRatio = 1.680000
	__TaskCompositionShape__00000013_.IsHidden = false

	__TaskGroup__00000000_.Name = `L1`
	__TaskGroup__00000000_.ComputedPrefix = ``
	__TaskGroup__00000000_.IsExpanded = false

	__TaskGroup__00000002_.Name = `L2`
	__TaskGroup__00000002_.ComputedPrefix = ``
	__TaskGroup__00000002_.IsExpanded = false

	__TaskGroup__00000003_.Name = `L3`
	__TaskGroup__00000003_.ComputedPrefix = ``
	__TaskGroup__00000003_.IsExpanded = false

	__TaskGroup__00000004_.Name = `year 2026`
	__TaskGroup__00000004_.ComputedPrefix = ``
	__TaskGroup__00000004_.IsExpanded = false

	__TaskGroup__00000005_.Name = `year 2027`
	__TaskGroup__00000005_.ComputedPrefix = ``
	__TaskGroup__00000005_.IsExpanded = false

	__TaskGroupShape__00000003_.Name = `-Default Diagram`
	__TaskGroupShape__00000003_.X = 171.471356
	__TaskGroupShape__00000003_.Y = 186.904210
	__TaskGroupShape__00000003_.Width = 250.000000
	__TaskGroupShape__00000003_.Height = 70.000000
	__TaskGroupShape__00000003_.IsHidden = false

	__TaskGroupShape__00000004_.Name = `-Default Diagram`
	__TaskGroupShape__00000004_.X = 128.490669
	__TaskGroupShape__00000004_.Y = 112.566493
	__TaskGroupShape__00000004_.Width = 250.000000
	__TaskGroupShape__00000004_.Height = 70.000000
	__TaskGroupShape__00000004_.IsHidden = false

	__TaskShape__00000019_.Name = `Seasons-Default Diagram`
	__TaskShape__00000019_.IsShowDate = false
	__TaskShape__00000019_.X = 134.246143
	__TaskShape__00000019_.Y = 124.124452
	__TaskShape__00000019_.Width = 250.000000
	__TaskShape__00000019_.Height = 70.000000
	__TaskShape__00000019_.IsHidden = false

	__TaskShape__00000020_.Name = `winter 26-Default Diagram`
	__TaskShape__00000020_.IsShowDate = true
	__TaskShape__00000020_.X = 134.246143
	__TaskShape__00000020_.Y = 264.124452
	__TaskShape__00000020_.Width = 250.000000
	__TaskShape__00000020_.Height = 70.000000
	__TaskShape__00000020_.IsHidden = false

	__TaskShape__00000021_.Name = `spring 26-Default Diagram`
	__TaskShape__00000021_.IsShowDate = true
	__TaskShape__00000021_.X = 434.246143
	__TaskShape__00000021_.Y = 264.124452
	__TaskShape__00000021_.Width = 250.000000
	__TaskShape__00000021_.Height = 70.000000
	__TaskShape__00000021_.IsHidden = false

	__TaskShape__00000022_.Name = `summer 26-Default Diagram`
	__TaskShape__00000022_.IsShowDate = true
	__TaskShape__00000022_.X = 734.246143
	__TaskShape__00000022_.Y = 264.124452
	__TaskShape__00000022_.Width = 250.000000
	__TaskShape__00000022_.Height = 70.000000
	__TaskShape__00000022_.IsHidden = false

	__TaskShape__00000023_.Name = `winter 27-Default Diagram`
	__TaskShape__00000023_.IsShowDate = true
	__TaskShape__00000023_.X = 1034.246143
	__TaskShape__00000023_.Y = 264.124452
	__TaskShape__00000023_.Width = 250.000000
	__TaskShape__00000023_.Height = 70.000000
	__TaskShape__00000023_.IsHidden = false

	// insertion point for setup of pointers
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000018_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000019_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000020_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000021_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000022_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000023_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000024_)
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000000_.ProductsWhoseNodeIsExpanded, __Product__00000000_)
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000000_.ProductsWhoseNodeIsExpanded, __Product__00000003_)
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000000_.ProductsWhoseNodeIsExpanded, __Product__00000002_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000015_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000016_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000017_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000018_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000019_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000020_)
	__Diagram__00000000_.MilestoneShapes = append(__Diagram__00000000_.MilestoneShapes, __MilestoneShape__00000000_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000019_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000020_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000021_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000022_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000023_)
	__Diagram__00000001_.TasksWhoseNodeIsExpanded = append(__Diagram__00000001_.TasksWhoseNodeIsExpanded, __Task__00000007_)
	__Diagram__00000001_.TaskGroupShapes = append(__Diagram__00000001_.TaskGroupShapes, __TaskGroupShape__00000003_)
	__Diagram__00000001_.TaskGroupShapes = append(__Diagram__00000001_.TaskGroupShapes, __TaskGroupShape__00000004_)
	__Diagram__00000001_.TaskGroupsWhoseNodeIsExpanded = append(__Diagram__00000001_.TaskGroupsWhoseNodeIsExpanded, __TaskGroup__00000004_)
	__Diagram__00000001_.TaskGroupsWhoseNodeIsExpanded = append(__Diagram__00000001_.TaskGroupsWhoseNodeIsExpanded, __TaskGroup__00000005_)
	__Diagram__00000001_.TaskComposition_Shapes = append(__Diagram__00000001_.TaskComposition_Shapes, __TaskCompositionShape__00000010_)
	__Diagram__00000001_.TaskComposition_Shapes = append(__Diagram__00000001_.TaskComposition_Shapes, __TaskCompositionShape__00000011_)
	__Diagram__00000001_.TaskComposition_Shapes = append(__Diagram__00000001_.TaskComposition_Shapes, __TaskCompositionShape__00000012_)
	__Diagram__00000001_.TaskComposition_Shapes = append(__Diagram__00000001_.TaskComposition_Shapes, __TaskCompositionShape__00000013_)
	__Library__00000000_.SubLibraries = append(__Library__00000000_.SubLibraries, __Library__00000001_)
	__Library__00000000_.RootProducts = append(__Library__00000000_.RootProducts, __Product__00000000_)
	__Library__00000000_.RootTasks = append(__Library__00000000_.RootTasks, __Task__00000000_)
	__Library__00000000_.RootTasks = append(__Library__00000000_.RootTasks, __Task__00000001_)
	__Library__00000000_.RootTasks = append(__Library__00000000_.RootTasks, __Task__00000002_)
	__Library__00000000_.RootTaskGroups = append(__Library__00000000_.RootTaskGroups, __TaskGroup__00000000_)
	__Library__00000000_.RootTaskGroups = append(__Library__00000000_.RootTaskGroups, __TaskGroup__00000002_)
	__Library__00000000_.RootTaskGroups = append(__Library__00000000_.RootTaskGroups, __TaskGroup__00000003_)
	__Library__00000000_.RootMilestones = append(__Library__00000000_.RootMilestones, __Milestone__00000000_)
	__Library__00000000_.Diagrams = append(__Library__00000000_.Diagrams, __Diagram__00000000_)
	__Library__00000001_.RootTasks = append(__Library__00000001_.RootTasks, __Task__00000007_)
	__Library__00000001_.RootTaskGroups = append(__Library__00000001_.RootTaskGroups, __TaskGroup__00000004_)
	__Library__00000001_.RootTaskGroups = append(__Library__00000001_.RootTaskGroups, __TaskGroup__00000005_)
	__Library__00000001_.Diagrams = append(__Library__00000001_.Diagrams, __Diagram__00000001_)
	__Milestone__00000000_.TaskGroupsToDisplay = append(__Milestone__00000000_.TaskGroupsToDisplay, __TaskGroup__00000002_)
	__Milestone__00000000_.TaskGroupsToDisplay = append(__Milestone__00000000_.TaskGroupsToDisplay, __TaskGroup__00000003_)
	__Milestone__00000000_.TaskGroupsToDisplay = append(__Milestone__00000000_.TaskGroupsToDisplay, __TaskGroup__00000000_)
	__MilestoneShape__00000000_.Milestone = __Milestone__00000000_
	__Product__00000000_.ReferencedProduct = nil
	__Product__00000000_.SubProducts = append(__Product__00000000_.SubProducts, __Product__00000001_)
	__Product__00000000_.SubProducts = append(__Product__00000000_.SubProducts, __Product__00000002_)
	__Product__00000000_.SubProducts = append(__Product__00000000_.SubProducts, __Product__00000003_)
	__Product__00000001_.ReferencedProduct = nil
	__Product__00000002_.ReferencedProduct = nil
	__Product__00000002_.SubProducts = append(__Product__00000002_.SubProducts, __Product__00000005_)
	__Product__00000002_.SubProducts = append(__Product__00000002_.SubProducts, __Product__00000006_)
	__Product__00000003_.ReferencedProduct = nil
	__Product__00000003_.SubProducts = append(__Product__00000003_.SubProducts, __Product__00000004_)
	__Product__00000004_.ReferencedProduct = nil
	__Product__00000005_.ReferencedProduct = nil
	__Product__00000006_.ReferencedProduct = nil
	__ProductCompositionShape__00000015_.Product = __Product__00000001_
	__ProductCompositionShape__00000016_.Product = __Product__00000003_
	__ProductCompositionShape__00000017_.Product = __Product__00000004_
	__ProductCompositionShape__00000018_.Product = __Product__00000002_
	__ProductCompositionShape__00000019_.Product = __Product__00000005_
	__ProductCompositionShape__00000020_.Product = __Product__00000006_
	__ProductShape__00000018_.Product = __Product__00000000_
	__ProductShape__00000019_.Product = __Product__00000001_
	__ProductShape__00000020_.Product = __Product__00000003_
	__ProductShape__00000021_.Product = __Product__00000004_
	__ProductShape__00000022_.Product = __Product__00000002_
	__ProductShape__00000023_.Product = __Product__00000005_
	__ProductShape__00000024_.Product = __Product__00000006_
	__Task__00000000_.ReferencedTask = nil
	__Task__00000001_.ReferencedTask = nil
	__Task__00000002_.ReferencedTask = nil
	__Task__00000003_.ReferencedTask = nil
	__Task__00000004_.ReferencedTask = nil
	__Task__00000005_.ReferencedTask = nil
	__Task__00000006_.ReferencedTask = nil
	__Task__00000007_.ReferencedTask = nil
	__Task__00000007_.SubTasks = append(__Task__00000007_.SubTasks, __Task__00000005_)
	__Task__00000007_.SubTasks = append(__Task__00000007_.SubTasks, __Task__00000003_)
	__Task__00000007_.SubTasks = append(__Task__00000007_.SubTasks, __Task__00000004_)
	__Task__00000007_.SubTasks = append(__Task__00000007_.SubTasks, __Task__00000006_)
	__TaskCompositionShape__00000010_.Task = __Task__00000005_
	__TaskCompositionShape__00000011_.Task = __Task__00000003_
	__TaskCompositionShape__00000012_.Task = __Task__00000004_
	__TaskCompositionShape__00000013_.Task = __Task__00000006_
	__TaskGroup__00000000_.Tasks = append(__TaskGroup__00000000_.Tasks, __Task__00000001_)
	__TaskGroup__00000000_.Tasks = append(__TaskGroup__00000000_.Tasks, __Task__00000000_)
	__TaskGroup__00000002_.Tasks = append(__TaskGroup__00000002_.Tasks, __Task__00000002_)
	__TaskGroup__00000004_.Tasks = append(__TaskGroup__00000004_.Tasks, __Task__00000005_)
	__TaskGroup__00000004_.Tasks = append(__TaskGroup__00000004_.Tasks, __Task__00000003_)
	__TaskGroup__00000004_.Tasks = append(__TaskGroup__00000004_.Tasks, __Task__00000004_)
	__TaskGroup__00000005_.Tasks = append(__TaskGroup__00000005_.Tasks, __Task__00000006_)
	__TaskGroupShape__00000003_.TaskGroup = __TaskGroup__00000004_
	__TaskGroupShape__00000004_.TaskGroup = __TaskGroup__00000005_
	__TaskShape__00000019_.Task = __Task__00000007_
	__TaskShape__00000020_.Task = __Task__00000005_
	__TaskShape__00000021_.Task = __Task__00000003_
	__TaskShape__00000022_.Task = __Task__00000004_
	__TaskShape__00000023_.Task = __Task__00000006_
}
