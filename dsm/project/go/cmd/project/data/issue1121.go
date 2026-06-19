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
	__Diagram__00000002_ := (&models.Diagram{Name: `Default Diagram copy`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: `Years diagram`}).Stage(stage)
	__Library__00000001_ := (&models.Library{Name: `Montly diagram`}).Stage(stage)

	__Product__00000000_ := (&models.Product{Name: `P1`}).Stage(stage)
	__Product__00000001_ := (&models.Product{Name: `P1.1`}).Stage(stage)
	__Product__00000002_ := (&models.Product{Name: `P1.2`}).Stage(stage)
	__Product__00000003_ := (&models.Product{Name: `P1.3`}).Stage(stage)
	__Product__00000004_ := (&models.Product{Name: `P1.3.1`}).Stage(stage)
	__Product__00000005_ := (&models.Product{Name: `P1.2.1 PP2PP P222 FFFDSFD. FDFQSDF  QSDfqdsfsdffs FDQSFSFD`}).Stage(stage)
	__Product__00000006_ := (&models.Product{Name: `P1.2.2`}).Stage(stage)
	__Product__00000007_ := (&models.Product{Name: `P1.2.2.1`}).Stage(stage)
	__Product__00000008_ := (&models.Product{Name: `P1.2.2.2 AAAA AAAA AAAA AAAAA`}).Stage(stage)
	__Product__00000009_ := (&models.Product{Name: `P1.2.2.3`}).Stage(stage)
	__Product__00000010_ := (&models.Product{Name: `P1.2.2.4`}).Stage(stage)
	__Product__00000011_ := (&models.Product{Name: `P1.2.2.1.1`}).Stage(stage)
	__Product__00000012_ := (&models.Product{Name: `B`}).Stage(stage)

	__ProductCompositionShape__00000062_ := (&models.ProductCompositionShape{Name: `P1 to P1.1`}).Stage(stage)
	__ProductCompositionShape__00000063_ := (&models.ProductCompositionShape{Name: `P1 to P1.2`}).Stage(stage)
	__ProductCompositionShape__00000064_ := (&models.ProductCompositionShape{Name: `P1.2 to P1.2.1 PP2PP P222 FFFDSFD. FDFQSDF  QSDfqdsfsdffs FDQSFSFD`}).Stage(stage)
	__ProductCompositionShape__00000065_ := (&models.ProductCompositionShape{Name: `P1.2 to P1.2.2`}).Stage(stage)
	__ProductCompositionShape__00000066_ := (&models.ProductCompositionShape{Name: `P1.2.2 to P1.2.2.1`}).Stage(stage)
	__ProductCompositionShape__00000067_ := (&models.ProductCompositionShape{Name: `P1.2.2.1 to P1.2.2.1.1`}).Stage(stage)
	__ProductCompositionShape__00000068_ := (&models.ProductCompositionShape{Name: `P1.2.2.1 to B`}).Stage(stage)
	__ProductCompositionShape__00000069_ := (&models.ProductCompositionShape{Name: `P1.2.2 to P1.2.2.2 AAAA AAAA AAAA AAAAA`}).Stage(stage)
	__ProductCompositionShape__00000070_ := (&models.ProductCompositionShape{Name: `P1.2.2 to P1.2.2.3`}).Stage(stage)
	__ProductCompositionShape__00000071_ := (&models.ProductCompositionShape{Name: `P1.2.2 to P1.2.2.4`}).Stage(stage)
	__ProductCompositionShape__00000072_ := (&models.ProductCompositionShape{Name: `P1 to P1.3`}).Stage(stage)
	__ProductCompositionShape__00000073_ := (&models.ProductCompositionShape{Name: `P1.3 to P1.3.1`}).Stage(stage)

	__ProductShape__00000073_ := (&models.ProductShape{Name: `P1-Default Diagram`}).Stage(stage)
	__ProductShape__00000074_ := (&models.ProductShape{Name: `P1.1-Default Diagram`}).Stage(stage)
	__ProductShape__00000075_ := (&models.ProductShape{Name: `P1.2-Default Diagram`}).Stage(stage)
	__ProductShape__00000076_ := (&models.ProductShape{Name: `P1.2.1 PP2PP P222 FFFDSFD. FDFQSDF  QSDfqdsfsdffs FDQSFSFD-Default Diagram`}).Stage(stage)
	__ProductShape__00000077_ := (&models.ProductShape{Name: `P1.2.2-Default Diagram`}).Stage(stage)
	__ProductShape__00000078_ := (&models.ProductShape{Name: `P1.2.2.1-Default Diagram`}).Stage(stage)
	__ProductShape__00000079_ := (&models.ProductShape{Name: `P1.2.2.1.1-Default Diagram`}).Stage(stage)
	__ProductShape__00000080_ := (&models.ProductShape{Name: `B-Default Diagram`}).Stage(stage)
	__ProductShape__00000081_ := (&models.ProductShape{Name: `P1.2.2.2 AAAA AAAA AAAA AAAAA-Default Diagram`}).Stage(stage)
	__ProductShape__00000082_ := (&models.ProductShape{Name: `P1.2.2.3-Default Diagram`}).Stage(stage)
	__ProductShape__00000083_ := (&models.ProductShape{Name: `P1.2.2.4-Default Diagram`}).Stage(stage)
	__ProductShape__00000084_ := (&models.ProductShape{Name: `P1.3-Default Diagram`}).Stage(stage)
	__ProductShape__00000085_ := (&models.ProductShape{Name: `P1.3.1-Default Diagram`}).Stage(stage)

	__Task__00000000_ := (&models.Task{Name: `L1.W1`}).Stage(stage)
	__Task__00000001_ := (&models.Task{Name: `L1.W2`}).Stage(stage)
	__Task__00000002_ := (&models.Task{Name: `L2.W1`}).Stage(stage)
	__Task__00000003_ := (&models.Task{Name: `spring 26`}).Stage(stage)
	__Task__00000004_ := (&models.Task{Name: `summer 26`}).Stage(stage)
	__Task__00000005_ := (&models.Task{Name: `winter 26`}).Stage(stage)
	__Task__00000006_ := (&models.Task{Name: `winter 27`}).Stage(stage)
	__Task__00000007_ := (&models.Task{Name: `Seasons`}).Stage(stage)
	__Task__00000008_ := (&models.Task{Name: `Jan 26`}).Stage(stage)
	__Task__00000009_ := (&models.Task{Name: `Fev 2`}).Stage(stage)
	__Task__00000010_ := (&models.Task{Name: `Mar 3`}).Stage(stage)
	__Task__00000011_ := (&models.Task{Name: `New Year 2027`}).Stage(stage)

	__TaskCompositionShape__00000011_ := (&models.TaskCompositionShape{Name: `Seasons to spring 26`}).Stage(stage)
	__TaskCompositionShape__00000012_ := (&models.TaskCompositionShape{Name: `Seasons to summer 26`}).Stage(stage)
	__TaskCompositionShape__00000013_ := (&models.TaskCompositionShape{Name: `Seasons to winter 27`}).Stage(stage)
	__TaskCompositionShape__00000014_ := (&models.TaskCompositionShape{Name: `Seasons to winter 26`}).Stage(stage)
	__TaskCompositionShape__00000015_ := (&models.TaskCompositionShape{Name: `Seasons to spring 26`}).Stage(stage)
	__TaskCompositionShape__00000016_ := (&models.TaskCompositionShape{Name: `Seasons to summer 26`}).Stage(stage)
	__TaskCompositionShape__00000017_ := (&models.TaskCompositionShape{Name: `Seasons to winter 27`}).Stage(stage)
	__TaskCompositionShape__00000018_ := (&models.TaskCompositionShape{Name: `winter 26 to `}).Stage(stage)
	__TaskCompositionShape__00000019_ := (&models.TaskCompositionShape{Name: `winter 26 to `}).Stage(stage)
	__TaskCompositionShape__00000020_ := (&models.TaskCompositionShape{Name: `winter 26 to `}).Stage(stage)
	__TaskCompositionShape__00000021_ := (&models.TaskCompositionShape{Name: `Seasons to `}).Stage(stage)
	__TaskCompositionShape__00000029_ := (&models.TaskCompositionShape{Name: `Seasons to winter 26`}).Stage(stage)
	__TaskCompositionShape__00000032_ := (&models.TaskCompositionShape{Name: `winter 26 to Jan 26`}).Stage(stage)

	__TaskGroup__00000000_ := (&models.TaskGroup{Name: `L1`}).Stage(stage)
	__TaskGroup__00000002_ := (&models.TaskGroup{Name: `L2`}).Stage(stage)
	__TaskGroup__00000003_ := (&models.TaskGroup{Name: `L3`}).Stage(stage)
	__TaskGroup__00000004_ := (&models.TaskGroup{Name: `year 2026`}).Stage(stage)
	__TaskGroup__00000005_ := (&models.TaskGroup{Name: `year 2027`}).Stage(stage)

	__TaskGroupShape__00000003_ := (&models.TaskGroupShape{Name: `-Default Diagram`}).Stage(stage)
	__TaskGroupShape__00000004_ := (&models.TaskGroupShape{Name: `-Default Diagram`}).Stage(stage)
	__TaskGroupShape__00000005_ := (&models.TaskGroupShape{Name: `-Default Diagram`}).Stage(stage)
	__TaskGroupShape__00000006_ := (&models.TaskGroupShape{Name: `-Default Diagram`}).Stage(stage)

	__TaskShape__00000019_ := (&models.TaskShape{Name: `Seasons-Default Diagram`}).Stage(stage)
	__TaskShape__00000021_ := (&models.TaskShape{Name: `spring 26-Default Diagram`}).Stage(stage)
	__TaskShape__00000022_ := (&models.TaskShape{Name: `summer 26-Default Diagram`}).Stage(stage)
	__TaskShape__00000023_ := (&models.TaskShape{Name: `winter 27-Default Diagram`}).Stage(stage)
	__TaskShape__00000024_ := (&models.TaskShape{Name: `Seasons-Default Diagram`}).Stage(stage)
	__TaskShape__00000025_ := (&models.TaskShape{Name: `winter 26-Default Diagram`}).Stage(stage)
	__TaskShape__00000026_ := (&models.TaskShape{Name: `spring 26-Default Diagram`}).Stage(stage)
	__TaskShape__00000027_ := (&models.TaskShape{Name: `summer 26-Default Diagram`}).Stage(stage)
	__TaskShape__00000028_ := (&models.TaskShape{Name: `winter 27-Default Diagram`}).Stage(stage)
	__TaskShape__00000029_ := (&models.TaskShape{Name: `-Default Diagram copy`}).Stage(stage)
	__TaskShape__00000030_ := (&models.TaskShape{Name: `-Default Diagram copy`}).Stage(stage)
	__TaskShape__00000031_ := (&models.TaskShape{Name: `-Default Diagram copy`}).Stage(stage)
	__TaskShape__00000032_ := (&models.TaskShape{Name: `-Default Diagram`}).Stage(stage)
	__TaskShape__00000042_ := (&models.TaskShape{Name: `winter 26-Default Diagram`}).Stage(stage)
	__TaskShape__00000045_ := (&models.TaskShape{Name: `Jan 26-Default Diagram`}).Stage(stage)

	// insertion point for initialization of values

	__Diagram__00000000_.Name = `Default Diagram`
	__Diagram__00000000_.ComputedPrefix = `1`
	__Diagram__00000000_.IsExpanded = false
	__Diagram__00000000_.LayoutDirection = models.Vertical
	__Diagram__00000000_.IsChecked = false
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.IsShowPrefix = false
	__Diagram__00000000_.DefaultBoxWidth = 180.000000
	__Diagram__00000000_.DefaultBoxHeigth = 50.000000
	__Diagram__00000000_.Width = 1350.000000
	__Diagram__00000000_.Height = 955.000000
	__Diagram__00000000_.IsPBSNodeExpanded = true
	__Diagram__00000000_.IsWBSNodeExpanded = false
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

	__Diagram__00000001_.Name = `Default Diagram`
	__Diagram__00000001_.ComputedPrefix = `1`
	__Diagram__00000001_.IsExpanded = true
	__Diagram__00000001_.LayoutDirection = models.Vertical
	__Diagram__00000001_.IsChecked = true
	__Diagram__00000001_.IsEditable_ = true
	__Diagram__00000001_.IsShowPrefix = false
	__Diagram__00000001_.DefaultBoxWidth = 250.000000
	__Diagram__00000001_.DefaultBoxHeigth = 70.000000
	__Diagram__00000001_.Width = 1684.246143
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

	__Diagram__00000002_.Name = `Default Diagram copy`
	__Diagram__00000002_.ComputedPrefix = `2`
	__Diagram__00000002_.IsExpanded = false
	__Diagram__00000002_.LayoutDirection = models.Vertical
	__Diagram__00000002_.IsChecked = false
	__Diagram__00000002_.IsEditable_ = true
	__Diagram__00000002_.IsShowPrefix = false
	__Diagram__00000002_.DefaultBoxWidth = 250.000000
	__Diagram__00000002_.DefaultBoxHeigth = 70.000000
	__Diagram__00000002_.Width = 1475.000000
	__Diagram__00000002_.Height = 615.000000
	__Diagram__00000002_.IsPBSNodeExpanded = false
	__Diagram__00000002_.IsWBSNodeExpanded = true
	__Diagram__00000002_.IsTaskGroupsNodeExpanded = false
	__Diagram__00000002_.DateFormat = ``
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

	__Product__00000000_.Name = `P1`
	__Product__00000000_.ComputedPrefix = `1`
	__Product__00000000_.IsExpanded = false
	__Product__00000000_.LayoutDirection = models.Horizontal
	__Product__00000000_.IsImport = false
	__Product__00000000_.Description = ``
	__Product__00000000_.IsProducersNodeExpanded = false
	__Product__00000000_.IsConsumersNodeExpanded = false

	__Product__00000001_.Name = `P1.1`
	__Product__00000001_.ComputedPrefix = `1.1`
	__Product__00000001_.IsExpanded = false
	__Product__00000001_.LayoutDirection = models.Vertical
	__Product__00000001_.IsImport = false
	__Product__00000001_.Description = ``
	__Product__00000001_.IsProducersNodeExpanded = false
	__Product__00000001_.IsConsumersNodeExpanded = false

	__Product__00000002_.Name = `P1.2`
	__Product__00000002_.ComputedPrefix = `1.2`
	__Product__00000002_.IsExpanded = false
	__Product__00000002_.LayoutDirection = models.Vertical
	__Product__00000002_.IsImport = false
	__Product__00000002_.Description = ``
	__Product__00000002_.IsProducersNodeExpanded = false
	__Product__00000002_.IsConsumersNodeExpanded = false

	__Product__00000003_.Name = `P1.3`
	__Product__00000003_.ComputedPrefix = `1.3`
	__Product__00000003_.IsExpanded = false
	__Product__00000003_.LayoutDirection = models.Vertical
	__Product__00000003_.IsImport = false
	__Product__00000003_.Description = ``
	__Product__00000003_.IsProducersNodeExpanded = false
	__Product__00000003_.IsConsumersNodeExpanded = false

	__Product__00000004_.Name = `P1.3.1`
	__Product__00000004_.ComputedPrefix = `1.3.1`
	__Product__00000004_.IsExpanded = false
	__Product__00000004_.LayoutDirection = models.Vertical
	__Product__00000004_.IsImport = false
	__Product__00000004_.Description = ``
	__Product__00000004_.IsProducersNodeExpanded = false
	__Product__00000004_.IsConsumersNodeExpanded = false

	__Product__00000005_.Name = `P1.2.1 PP2PP P222 FFFDSFD. FDFQSDF  QSDfqdsfsdffs FDQSFSFD`
	__Product__00000005_.ComputedPrefix = `1.2.1`
	__Product__00000005_.IsExpanded = false
	__Product__00000005_.LayoutDirection = models.Vertical
	__Product__00000005_.IsImport = false
	__Product__00000005_.Description = ``
	__Product__00000005_.IsProducersNodeExpanded = false
	__Product__00000005_.IsConsumersNodeExpanded = false

	__Product__00000006_.Name = `P1.2.2`
	__Product__00000006_.ComputedPrefix = `1.2.2`
	__Product__00000006_.IsExpanded = false
	__Product__00000006_.LayoutDirection = models.Horizontal
	__Product__00000006_.IsImport = false
	__Product__00000006_.Description = ``
	__Product__00000006_.IsProducersNodeExpanded = false
	__Product__00000006_.IsConsumersNodeExpanded = false

	__Product__00000007_.Name = `P1.2.2.1`
	__Product__00000007_.ComputedPrefix = `1.2.2.1`
	__Product__00000007_.IsExpanded = false
	__Product__00000007_.LayoutDirection = models.Vertical
	__Product__00000007_.IsImport = false
	__Product__00000007_.Description = ``
	__Product__00000007_.IsProducersNodeExpanded = false
	__Product__00000007_.IsConsumersNodeExpanded = false

	__Product__00000008_.Name = `P1.2.2.2 AAAA AAAA AAAA AAAAA`
	__Product__00000008_.ComputedPrefix = `1.2.2.2`
	__Product__00000008_.IsExpanded = false
	__Product__00000008_.LayoutDirection = models.Vertical
	__Product__00000008_.IsImport = false
	__Product__00000008_.Description = ``
	__Product__00000008_.IsProducersNodeExpanded = false
	__Product__00000008_.IsConsumersNodeExpanded = false

	__Product__00000009_.Name = `P1.2.2.3`
	__Product__00000009_.ComputedPrefix = `1.2.2.3`
	__Product__00000009_.IsExpanded = false
	__Product__00000009_.LayoutDirection = models.Vertical
	__Product__00000009_.IsImport = false
	__Product__00000009_.Description = ``
	__Product__00000009_.IsProducersNodeExpanded = false
	__Product__00000009_.IsConsumersNodeExpanded = false

	__Product__00000010_.Name = `P1.2.2.4`
	__Product__00000010_.ComputedPrefix = `1.2.2.4`
	__Product__00000010_.IsExpanded = false
	__Product__00000010_.LayoutDirection = models.Vertical
	__Product__00000010_.IsImport = false
	__Product__00000010_.Description = ``
	__Product__00000010_.IsProducersNodeExpanded = false
	__Product__00000010_.IsConsumersNodeExpanded = false

	__Product__00000011_.Name = `P1.2.2.1.1`
	__Product__00000011_.ComputedPrefix = `1.2.2.1.1`
	__Product__00000011_.IsExpanded = false
	__Product__00000011_.LayoutDirection = models.Vertical
	__Product__00000011_.IsImport = false
	__Product__00000011_.Description = ``
	__Product__00000011_.IsProducersNodeExpanded = false
	__Product__00000011_.IsConsumersNodeExpanded = false

	__Product__00000012_.Name = `B`
	__Product__00000012_.ComputedPrefix = `1.2.2.1.2`
	__Product__00000012_.IsExpanded = false
	__Product__00000012_.LayoutDirection = models.Vertical
	__Product__00000012_.IsImport = false
	__Product__00000012_.Description = ``
	__Product__00000012_.IsProducersNodeExpanded = false
	__Product__00000012_.IsConsumersNodeExpanded = false

	__ProductCompositionShape__00000062_.Name = `P1 to P1.1`
	__ProductCompositionShape__00000062_.StartRatio = 0.500000
	__ProductCompositionShape__00000062_.EndRatio = 0.500000
	__ProductCompositionShape__00000062_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000062_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000062_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000062_.IsHidden = false

	__ProductCompositionShape__00000063_.Name = `P1 to P1.2`
	__ProductCompositionShape__00000063_.StartRatio = 0.500000
	__ProductCompositionShape__00000063_.EndRatio = 0.500000
	__ProductCompositionShape__00000063_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000063_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000063_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000063_.IsHidden = false

	__ProductCompositionShape__00000064_.Name = `P1.2 to P1.2.1 PP2PP P222 FFFDSFD. FDFQSDF  QSDfqdsfsdffs FDQSFSFD`
	__ProductCompositionShape__00000064_.StartRatio = 0.500000
	__ProductCompositionShape__00000064_.EndRatio = 0.500000
	__ProductCompositionShape__00000064_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000064_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000064_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000064_.IsHidden = false

	__ProductCompositionShape__00000065_.Name = `P1.2 to P1.2.2`
	__ProductCompositionShape__00000065_.StartRatio = 0.500000
	__ProductCompositionShape__00000065_.EndRatio = 0.500000
	__ProductCompositionShape__00000065_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000065_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000065_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000065_.IsHidden = false

	__ProductCompositionShape__00000066_.Name = `P1.2.2 to P1.2.2.1`
	__ProductCompositionShape__00000066_.StartRatio = 0.500000
	__ProductCompositionShape__00000066_.EndRatio = 0.500000
	__ProductCompositionShape__00000066_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000066_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000066_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000066_.IsHidden = false

	__ProductCompositionShape__00000067_.Name = `P1.2.2.1 to P1.2.2.1.1`
	__ProductCompositionShape__00000067_.StartRatio = 0.500000
	__ProductCompositionShape__00000067_.EndRatio = 0.500000
	__ProductCompositionShape__00000067_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000067_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000067_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000067_.IsHidden = false

	__ProductCompositionShape__00000068_.Name = `P1.2.2.1 to B`
	__ProductCompositionShape__00000068_.StartRatio = 0.500000
	__ProductCompositionShape__00000068_.EndRatio = 0.500000
	__ProductCompositionShape__00000068_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000068_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000068_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000068_.IsHidden = false

	__ProductCompositionShape__00000069_.Name = `P1.2.2 to P1.2.2.2 AAAA AAAA AAAA AAAAA`
	__ProductCompositionShape__00000069_.StartRatio = 0.500000
	__ProductCompositionShape__00000069_.EndRatio = 0.500000
	__ProductCompositionShape__00000069_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000069_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000069_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000069_.IsHidden = false

	__ProductCompositionShape__00000070_.Name = `P1.2.2 to P1.2.2.3`
	__ProductCompositionShape__00000070_.StartRatio = 0.500000
	__ProductCompositionShape__00000070_.EndRatio = 0.500000
	__ProductCompositionShape__00000070_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000070_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000070_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000070_.IsHidden = false

	__ProductCompositionShape__00000071_.Name = `P1.2.2 to P1.2.2.4`
	__ProductCompositionShape__00000071_.StartRatio = 0.500000
	__ProductCompositionShape__00000071_.EndRatio = 0.500000
	__ProductCompositionShape__00000071_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000071_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000071_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000071_.IsHidden = false

	__ProductCompositionShape__00000072_.Name = `P1 to P1.3`
	__ProductCompositionShape__00000072_.StartRatio = 0.500000
	__ProductCompositionShape__00000072_.EndRatio = 0.500000
	__ProductCompositionShape__00000072_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000072_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__ProductCompositionShape__00000072_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000072_.IsHidden = false

	__ProductCompositionShape__00000073_.Name = `P1.3 to P1.3.1`
	__ProductCompositionShape__00000073_.StartRatio = 0.500000
	__ProductCompositionShape__00000073_.EndRatio = 0.500000
	__ProductCompositionShape__00000073_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000073_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000073_.CornerOffsetRatio = 1.500000
	__ProductCompositionShape__00000073_.IsHidden = false

	__ProductShape__00000073_.Name = `P1-Default Diagram`
	__ProductShape__00000073_.OverideLayoutDirection = false
	__ProductShape__00000073_.LayoutDirection = models.Vertical
	__ProductShape__00000073_.X = 50.000000
	__ProductShape__00000073_.Y = 50.000000
	__ProductShape__00000073_.Width = 180.000000
	__ProductShape__00000073_.Height = 50.000000
	__ProductShape__00000073_.IsHidden = false

	__ProductShape__00000074_.Name = `P1.1-Default Diagram`
	__ProductShape__00000074_.OverideLayoutDirection = false
	__ProductShape__00000074_.LayoutDirection = models.Vertical
	__ProductShape__00000074_.X = 190.000000
	__ProductShape__00000074_.Y = 115.000000
	__ProductShape__00000074_.Width = 180.000000
	__ProductShape__00000074_.Height = 50.000000
	__ProductShape__00000074_.IsHidden = false

	__ProductShape__00000075_.Name = `P1.2-Default Diagram`
	__ProductShape__00000075_.OverideLayoutDirection = false
	__ProductShape__00000075_.LayoutDirection = models.Vertical
	__ProductShape__00000075_.X = 190.000000
	__ProductShape__00000075_.Y = 180.000000
	__ProductShape__00000075_.Width = 180.000000
	__ProductShape__00000075_.Height = 50.000000
	__ProductShape__00000075_.IsHidden = false

	__ProductShape__00000076_.Name = `P1.2.1 PP2PP P222 FFFDSFD. FDFQSDF  QSDfqdsfsdffs FDQSFSFD-Default Diagram`
	__ProductShape__00000076_.OverideLayoutDirection = false
	__ProductShape__00000076_.LayoutDirection = models.Vertical
	__ProductShape__00000076_.X = 330.000000
	__ProductShape__00000076_.Y = 280.000000
	__ProductShape__00000076_.Width = 180.000000
	__ProductShape__00000076_.Height = 50.000000
	__ProductShape__00000076_.IsHidden = false

	__ProductShape__00000077_.Name = `P1.2.2-Default Diagram`
	__ProductShape__00000077_.OverideLayoutDirection = false
	__ProductShape__00000077_.LayoutDirection = models.Vertical
	__ProductShape__00000077_.X = 560.000000
	__ProductShape__00000077_.Y = 280.000000
	__ProductShape__00000077_.Width = 180.000000
	__ProductShape__00000077_.Height = 50.000000
	__ProductShape__00000077_.IsHidden = false

	__ProductShape__00000078_.Name = `P1.2.2.1-Default Diagram`
	__ProductShape__00000078_.OverideLayoutDirection = false
	__ProductShape__00000078_.LayoutDirection = models.Vertical
	__ProductShape__00000078_.X = 700.000000
	__ProductShape__00000078_.Y = 345.000000
	__ProductShape__00000078_.Width = 180.000000
	__ProductShape__00000078_.Height = 50.000000
	__ProductShape__00000078_.IsHidden = false

	__ProductShape__00000079_.Name = `P1.2.2.1.1-Default Diagram`
	__ProductShape__00000079_.OverideLayoutDirection = false
	__ProductShape__00000079_.LayoutDirection = models.Vertical
	__ProductShape__00000079_.X = 840.000000
	__ProductShape__00000079_.Y = 445.000000
	__ProductShape__00000079_.Width = 180.000000
	__ProductShape__00000079_.Height = 50.000000
	__ProductShape__00000079_.IsHidden = false

	__ProductShape__00000080_.Name = `B-Default Diagram`
	__ProductShape__00000080_.OverideLayoutDirection = false
	__ProductShape__00000080_.LayoutDirection = models.Vertical
	__ProductShape__00000080_.X = 1070.000000
	__ProductShape__00000080_.Y = 445.000000
	__ProductShape__00000080_.Width = 180.000000
	__ProductShape__00000080_.Height = 50.000000
	__ProductShape__00000080_.IsHidden = false

	__ProductShape__00000081_.Name = `P1.2.2.2 AAAA AAAA AAAA AAAAA-Default Diagram`
	__ProductShape__00000081_.OverideLayoutDirection = false
	__ProductShape__00000081_.LayoutDirection = models.Vertical
	__ProductShape__00000081_.X = 700.000000
	__ProductShape__00000081_.Y = 510.000000
	__ProductShape__00000081_.Width = 180.000000
	__ProductShape__00000081_.Height = 50.000000
	__ProductShape__00000081_.IsHidden = false

	__ProductShape__00000082_.Name = `P1.2.2.3-Default Diagram`
	__ProductShape__00000082_.OverideLayoutDirection = false
	__ProductShape__00000082_.LayoutDirection = models.Vertical
	__ProductShape__00000082_.X = 700.000000
	__ProductShape__00000082_.Y = 575.000000
	__ProductShape__00000082_.Width = 180.000000
	__ProductShape__00000082_.Height = 50.000000
	__ProductShape__00000082_.IsHidden = false

	__ProductShape__00000083_.Name = `P1.2.2.4-Default Diagram`
	__ProductShape__00000083_.OverideLayoutDirection = false
	__ProductShape__00000083_.LayoutDirection = models.Vertical
	__ProductShape__00000083_.X = 700.000000
	__ProductShape__00000083_.Y = 640.000000
	__ProductShape__00000083_.Width = 180.000000
	__ProductShape__00000083_.Height = 50.000000
	__ProductShape__00000083_.IsHidden = false

	__ProductShape__00000084_.Name = `P1.3-Default Diagram`
	__ProductShape__00000084_.OverideLayoutDirection = false
	__ProductShape__00000084_.LayoutDirection = models.Vertical
	__ProductShape__00000084_.X = 190.000000
	__ProductShape__00000084_.Y = 705.000000
	__ProductShape__00000084_.Width = 180.000000
	__ProductShape__00000084_.Height = 50.000000
	__ProductShape__00000084_.IsHidden = false

	__ProductShape__00000085_.Name = `P1.3.1-Default Diagram`
	__ProductShape__00000085_.OverideLayoutDirection = false
	__ProductShape__00000085_.LayoutDirection = models.Vertical
	__ProductShape__00000085_.X = 330.000000
	__ProductShape__00000085_.Y = 805.000000
	__ProductShape__00000085_.Width = 180.000000
	__ProductShape__00000085_.Height = 50.000000
	__ProductShape__00000085_.IsHidden = false

	__Task__00000000_.Name = `L1.W1`
	__Task__00000000_.ComputedPrefix = `1`
	__Task__00000000_.IsExpanded = false
	__Task__00000000_.LayoutDirection = models.Vertical
	__Task__00000000_.IsImport = false
	__Task__00000000_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2000-01-01 00:00:00 +0000 UTC")
	__Task__00000000_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2030-01-01 00:00:00 +0000 UTC")
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

	__Task__00000001_.Name = `L1.W2`
	__Task__00000001_.ComputedPrefix = `2`
	__Task__00000001_.IsExpanded = false
	__Task__00000001_.LayoutDirection = models.Vertical
	__Task__00000001_.IsImport = false
	__Task__00000001_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2020-01-01 00:00:00 +0000 UTC")
	__Task__00000001_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2025-01-01 00:00:00 +0000 UTC")
	__Task__00000001_.IsMilestone = false
	__Task__00000001_.Description = ``
	__Task__00000001_.IsInputsNodeExpanded = false
	__Task__00000001_.IsOutputsNodeExpanded = false
	__Task__00000001_.IsWithCompletion = false
	__Task__00000001_.Completion = ""
	__Task__00000001_.DisplayVerticalBar = false
	__Task__00000001_.TextPosition = ""
	__Task__00000001_.XOffset = 0.000000
	__Task__00000001_.YOffset = 0.000000

	__Task__00000002_.Name = `L2.W1`
	__Task__00000002_.ComputedPrefix = `3`
	__Task__00000002_.IsExpanded = false
	__Task__00000002_.LayoutDirection = models.Vertical
	__Task__00000002_.IsImport = false
	__Task__00000002_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2021-01-01 00:00:00 +0000 UTC")
	__Task__00000002_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2022-01-01 00:00:00 +0000 UTC")
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

	__Task__00000003_.Name = `spring 26`
	__Task__00000003_.ComputedPrefix = `1.2`
	__Task__00000003_.IsExpanded = false
	__Task__00000003_.LayoutDirection = models.Vertical
	__Task__00000003_.IsImport = false
	__Task__00000003_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-04-01 00:00:00 +0000 UTC")
	__Task__00000003_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "2026-07-01 00:00:00 +0000 UTC")
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

	__Task__00000009_.Name = `Fev 2`
	__Task__00000009_.ComputedPrefix = `1.1.2`
	__Task__00000009_.IsExpanded = false
	__Task__00000009_.LayoutDirection = models.Vertical
	__Task__00000009_.IsImport = false
	__Task__00000009_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000009_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
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

	__Task__00000010_.Name = `Mar 3`
	__Task__00000010_.ComputedPrefix = `1.1.3`
	__Task__00000010_.IsExpanded = false
	__Task__00000010_.LayoutDirection = models.Vertical
	__Task__00000010_.IsImport = false
	__Task__00000010_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000010_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
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

	__TaskCompositionShape__00000014_.Name = `Seasons to winter 26`
	__TaskCompositionShape__00000014_.StartRatio = 0.500000
	__TaskCompositionShape__00000014_.EndRatio = 0.500000
	__TaskCompositionShape__00000014_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000014_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000014_.CornerOffsetRatio = 1.500000
	__TaskCompositionShape__00000014_.IsHidden = false

	__TaskCompositionShape__00000015_.Name = `Seasons to spring 26`
	__TaskCompositionShape__00000015_.StartRatio = 0.500000
	__TaskCompositionShape__00000015_.EndRatio = 0.500000
	__TaskCompositionShape__00000015_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000015_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000015_.CornerOffsetRatio = 1.500000
	__TaskCompositionShape__00000015_.IsHidden = false

	__TaskCompositionShape__00000016_.Name = `Seasons to summer 26`
	__TaskCompositionShape__00000016_.StartRatio = 0.500000
	__TaskCompositionShape__00000016_.EndRatio = 0.500000
	__TaskCompositionShape__00000016_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000016_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000016_.CornerOffsetRatio = 1.500000
	__TaskCompositionShape__00000016_.IsHidden = false

	__TaskCompositionShape__00000017_.Name = `Seasons to winter 27`
	__TaskCompositionShape__00000017_.StartRatio = 0.500000
	__TaskCompositionShape__00000017_.EndRatio = 0.500000
	__TaskCompositionShape__00000017_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000017_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000017_.CornerOffsetRatio = 1.500000
	__TaskCompositionShape__00000017_.IsHidden = false

	__TaskCompositionShape__00000018_.Name = `winter 26 to `
	__TaskCompositionShape__00000018_.StartRatio = 0.500000
	__TaskCompositionShape__00000018_.EndRatio = 0.500000
	__TaskCompositionShape__00000018_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000018_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__TaskCompositionShape__00000018_.CornerOffsetRatio = 1.500000
	__TaskCompositionShape__00000018_.IsHidden = false

	__TaskCompositionShape__00000019_.Name = `winter 26 to `
	__TaskCompositionShape__00000019_.StartRatio = 0.500000
	__TaskCompositionShape__00000019_.EndRatio = 0.500000
	__TaskCompositionShape__00000019_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000019_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__TaskCompositionShape__00000019_.CornerOffsetRatio = 1.500000
	__TaskCompositionShape__00000019_.IsHidden = false

	__TaskCompositionShape__00000020_.Name = `winter 26 to `
	__TaskCompositionShape__00000020_.StartRatio = 0.500000
	__TaskCompositionShape__00000020_.EndRatio = 0.500000
	__TaskCompositionShape__00000020_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000020_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__TaskCompositionShape__00000020_.CornerOffsetRatio = 1.500000
	__TaskCompositionShape__00000020_.IsHidden = false

	__TaskCompositionShape__00000021_.Name = `Seasons to `
	__TaskCompositionShape__00000021_.StartRatio = 0.500000
	__TaskCompositionShape__00000021_.EndRatio = 0.500000
	__TaskCompositionShape__00000021_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000021_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000021_.CornerOffsetRatio = 1.500000
	__TaskCompositionShape__00000021_.IsHidden = false

	__TaskCompositionShape__00000029_.Name = `Seasons to winter 26`
	__TaskCompositionShape__00000029_.StartRatio = 0.500000
	__TaskCompositionShape__00000029_.EndRatio = 0.500000
	__TaskCompositionShape__00000029_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000029_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000029_.CornerOffsetRatio = 1.500000
	__TaskCompositionShape__00000029_.IsHidden = false

	__TaskCompositionShape__00000032_.Name = `winter 26 to Jan 26`
	__TaskCompositionShape__00000032_.StartRatio = 0.500000
	__TaskCompositionShape__00000032_.EndRatio = 0.500000
	__TaskCompositionShape__00000032_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000032_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000032_.CornerOffsetRatio = 1.500000
	__TaskCompositionShape__00000032_.IsHidden = false

	__TaskGroup__00000000_.Name = `L1`
	__TaskGroup__00000000_.ComputedPrefix = ``
	__TaskGroup__00000000_.IsExpanded = false
	__TaskGroup__00000000_.LayoutDirection = models.Vertical

	__TaskGroup__00000002_.Name = `L2`
	__TaskGroup__00000002_.ComputedPrefix = ``
	__TaskGroup__00000002_.IsExpanded = false
	__TaskGroup__00000002_.LayoutDirection = models.Vertical

	__TaskGroup__00000003_.Name = `L3`
	__TaskGroup__00000003_.ComputedPrefix = ``
	__TaskGroup__00000003_.IsExpanded = false
	__TaskGroup__00000003_.LayoutDirection = models.Vertical

	__TaskGroup__00000004_.Name = `year 2026`
	__TaskGroup__00000004_.ComputedPrefix = ``
	__TaskGroup__00000004_.IsExpanded = false
	__TaskGroup__00000004_.LayoutDirection = models.Vertical

	__TaskGroup__00000005_.Name = `year 2027`
	__TaskGroup__00000005_.ComputedPrefix = ``
	__TaskGroup__00000005_.IsExpanded = false
	__TaskGroup__00000005_.LayoutDirection = models.Vertical

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

	__TaskGroupShape__00000005_.Name = `-Default Diagram`
	__TaskGroupShape__00000005_.X = 171.471356
	__TaskGroupShape__00000005_.Y = 186.904210
	__TaskGroupShape__00000005_.Width = 250.000000
	__TaskGroupShape__00000005_.Height = 70.000000
	__TaskGroupShape__00000005_.IsHidden = false

	__TaskGroupShape__00000006_.Name = `-Default Diagram`
	__TaskGroupShape__00000006_.X = 128.490669
	__TaskGroupShape__00000006_.Y = 112.566493
	__TaskGroupShape__00000006_.Width = 250.000000
	__TaskGroupShape__00000006_.Height = 70.000000
	__TaskGroupShape__00000006_.IsHidden = false

	__TaskShape__00000019_.Name = `Seasons-Default Diagram`
	__TaskShape__00000019_.IsShowDate = false
	__TaskShape__00000019_.OverideLayoutDirection = false
	__TaskShape__00000019_.LayoutDirection = models.Vertical
	__TaskShape__00000019_.X = 134.246143
	__TaskShape__00000019_.Y = 124.124452
	__TaskShape__00000019_.Width = 250.000000
	__TaskShape__00000019_.Height = 70.000000
	__TaskShape__00000019_.IsHidden = false

	__TaskShape__00000021_.Name = `spring 26-Default Diagram`
	__TaskShape__00000021_.IsShowDate = true
	__TaskShape__00000021_.OverideLayoutDirection = false
	__TaskShape__00000021_.LayoutDirection = models.Vertical
	__TaskShape__00000021_.X = 434.246143
	__TaskShape__00000021_.Y = 264.124452
	__TaskShape__00000021_.Width = 250.000000
	__TaskShape__00000021_.Height = 70.000000
	__TaskShape__00000021_.IsHidden = false

	__TaskShape__00000022_.Name = `summer 26-Default Diagram`
	__TaskShape__00000022_.IsShowDate = true
	__TaskShape__00000022_.OverideLayoutDirection = false
	__TaskShape__00000022_.LayoutDirection = models.Vertical
	__TaskShape__00000022_.X = 734.246143
	__TaskShape__00000022_.Y = 264.124452
	__TaskShape__00000022_.Width = 250.000000
	__TaskShape__00000022_.Height = 70.000000
	__TaskShape__00000022_.IsHidden = false

	__TaskShape__00000023_.Name = `winter 27-Default Diagram`
	__TaskShape__00000023_.IsShowDate = true
	__TaskShape__00000023_.OverideLayoutDirection = false
	__TaskShape__00000023_.LayoutDirection = models.Vertical
	__TaskShape__00000023_.X = 1034.246143
	__TaskShape__00000023_.Y = 264.124452
	__TaskShape__00000023_.Width = 250.000000
	__TaskShape__00000023_.Height = 70.000000
	__TaskShape__00000023_.IsHidden = false

	__TaskShape__00000024_.Name = `Seasons-Default Diagram`
	__TaskShape__00000024_.IsShowDate = false
	__TaskShape__00000024_.OverideLayoutDirection = false
	__TaskShape__00000024_.LayoutDirection = models.Vertical
	__TaskShape__00000024_.X = 50.000000
	__TaskShape__00000024_.Y = 50.000000
	__TaskShape__00000024_.Width = 250.000000
	__TaskShape__00000024_.Height = 70.000000
	__TaskShape__00000024_.IsHidden = false

	__TaskShape__00000025_.Name = `winter 26-Default Diagram`
	__TaskShape__00000025_.IsShowDate = true
	__TaskShape__00000025_.OverideLayoutDirection = false
	__TaskShape__00000025_.LayoutDirection = models.Vertical
	__TaskShape__00000025_.X = 50.000000
	__TaskShape__00000025_.Y = 190.000000
	__TaskShape__00000025_.Width = 250.000000
	__TaskShape__00000025_.Height = 70.000000
	__TaskShape__00000025_.IsHidden = false

	__TaskShape__00000026_.Name = `spring 26-Default Diagram`
	__TaskShape__00000026_.IsShowDate = true
	__TaskShape__00000026_.OverideLayoutDirection = false
	__TaskShape__00000026_.LayoutDirection = models.Vertical
	__TaskShape__00000026_.X = 525.000000
	__TaskShape__00000026_.Y = 190.000000
	__TaskShape__00000026_.Width = 250.000000
	__TaskShape__00000026_.Height = 70.000000
	__TaskShape__00000026_.IsHidden = false

	__TaskShape__00000027_.Name = `summer 26-Default Diagram`
	__TaskShape__00000027_.IsShowDate = true
	__TaskShape__00000027_.OverideLayoutDirection = false
	__TaskShape__00000027_.LayoutDirection = models.Vertical
	__TaskShape__00000027_.X = 825.000000
	__TaskShape__00000027_.Y = 190.000000
	__TaskShape__00000027_.Width = 250.000000
	__TaskShape__00000027_.Height = 70.000000
	__TaskShape__00000027_.IsHidden = false

	__TaskShape__00000028_.Name = `winter 27-Default Diagram`
	__TaskShape__00000028_.IsShowDate = true
	__TaskShape__00000028_.OverideLayoutDirection = false
	__TaskShape__00000028_.LayoutDirection = models.Vertical
	__TaskShape__00000028_.X = 1125.000000
	__TaskShape__00000028_.Y = 190.000000
	__TaskShape__00000028_.Width = 250.000000
	__TaskShape__00000028_.Height = 70.000000
	__TaskShape__00000028_.IsHidden = false

	__TaskShape__00000029_.Name = `-Default Diagram copy`
	__TaskShape__00000029_.IsShowDate = false
	__TaskShape__00000029_.OverideLayoutDirection = false
	__TaskShape__00000029_.LayoutDirection = models.Vertical
	__TaskShape__00000029_.X = 225.000000
	__TaskShape__00000029_.Y = 275.000000
	__TaskShape__00000029_.Width = 250.000000
	__TaskShape__00000029_.Height = 70.000000
	__TaskShape__00000029_.IsHidden = false

	__TaskShape__00000030_.Name = `-Default Diagram copy`
	__TaskShape__00000030_.IsShowDate = false
	__TaskShape__00000030_.OverideLayoutDirection = false
	__TaskShape__00000030_.LayoutDirection = models.Vertical
	__TaskShape__00000030_.X = 225.000000
	__TaskShape__00000030_.Y = 360.000000
	__TaskShape__00000030_.Width = 250.000000
	__TaskShape__00000030_.Height = 70.000000
	__TaskShape__00000030_.IsHidden = false

	__TaskShape__00000031_.Name = `-Default Diagram copy`
	__TaskShape__00000031_.IsShowDate = false
	__TaskShape__00000031_.OverideLayoutDirection = false
	__TaskShape__00000031_.LayoutDirection = models.Vertical
	__TaskShape__00000031_.X = 225.000000
	__TaskShape__00000031_.Y = 445.000000
	__TaskShape__00000031_.Width = 250.000000
	__TaskShape__00000031_.Height = 70.000000
	__TaskShape__00000031_.IsHidden = false

	__TaskShape__00000032_.Name = `-Default Diagram`
	__TaskShape__00000032_.IsShowDate = false
	__TaskShape__00000032_.OverideLayoutDirection = false
	__TaskShape__00000032_.LayoutDirection = models.Vertical
	__TaskShape__00000032_.X = 1334.246143
	__TaskShape__00000032_.Y = 264.124452
	__TaskShape__00000032_.Width = 250.000000
	__TaskShape__00000032_.Height = 70.000000
	__TaskShape__00000032_.IsHidden = false

	__TaskShape__00000042_.Name = `winter 26-Default Diagram`
	__TaskShape__00000042_.IsShowDate = false
	__TaskShape__00000042_.OverideLayoutDirection = false
	__TaskShape__00000042_.LayoutDirection = models.Vertical
	__TaskShape__00000042_.X = 1334.246143
	__TaskShape__00000042_.Y = 264.124452
	__TaskShape__00000042_.Width = 250.000000
	__TaskShape__00000042_.Height = 70.000000
	__TaskShape__00000042_.IsHidden = false

	__TaskShape__00000045_.Name = `Jan 26-Default Diagram`
	__TaskShape__00000045_.IsShowDate = false
	__TaskShape__00000045_.OverideLayoutDirection = false
	__TaskShape__00000045_.LayoutDirection = models.Vertical
	__TaskShape__00000045_.X = 1334.246143
	__TaskShape__00000045_.Y = 404.124452
	__TaskShape__00000045_.Width = 250.000000
	__TaskShape__00000045_.Height = 70.000000
	__TaskShape__00000045_.IsHidden = false

	// insertion point for setup of pointers
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000073_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000074_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000075_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000076_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000077_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000078_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000079_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000080_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000081_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000082_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000083_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000084_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000085_)
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000000_.ProductsWhoseNodeIsExpanded, __Product__00000000_)
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000000_.ProductsWhoseNodeIsExpanded, __Product__00000003_)
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000000_.ProductsWhoseNodeIsExpanded, __Product__00000002_)
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000000_.ProductsWhoseNodeIsExpanded, __Product__00000006_)
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000000_.ProductsWhoseNodeIsExpanded, __Product__00000007_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000062_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000063_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000064_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000065_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000066_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000067_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000068_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000069_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000070_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000071_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000072_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000073_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000019_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000021_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000022_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000023_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000032_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000042_)
	__Diagram__00000001_.Task_Shapes = append(__Diagram__00000001_.Task_Shapes, __TaskShape__00000045_)
	__Diagram__00000001_.TasksWhoseNodeIsExpanded = append(__Diagram__00000001_.TasksWhoseNodeIsExpanded, __Task__00000007_)
	__Diagram__00000001_.TasksWhoseNodeIsExpanded = append(__Diagram__00000001_.TasksWhoseNodeIsExpanded, __Task__00000005_)
	__Diagram__00000001_.TaskGroupShapes = append(__Diagram__00000001_.TaskGroupShapes, __TaskGroupShape__00000003_)
	__Diagram__00000001_.TaskGroupShapes = append(__Diagram__00000001_.TaskGroupShapes, __TaskGroupShape__00000004_)
	__Diagram__00000001_.TaskGroupsWhoseNodeIsExpanded = append(__Diagram__00000001_.TaskGroupsWhoseNodeIsExpanded, __TaskGroup__00000005_)
	__Diagram__00000001_.TaskGroupsWhoseNodeIsExpanded = append(__Diagram__00000001_.TaskGroupsWhoseNodeIsExpanded, __TaskGroup__00000004_)
	__Diagram__00000001_.TaskComposition_Shapes = append(__Diagram__00000001_.TaskComposition_Shapes, __TaskCompositionShape__00000011_)
	__Diagram__00000001_.TaskComposition_Shapes = append(__Diagram__00000001_.TaskComposition_Shapes, __TaskCompositionShape__00000012_)
	__Diagram__00000001_.TaskComposition_Shapes = append(__Diagram__00000001_.TaskComposition_Shapes, __TaskCompositionShape__00000013_)
	__Diagram__00000001_.TaskComposition_Shapes = append(__Diagram__00000001_.TaskComposition_Shapes, __TaskCompositionShape__00000021_)
	__Diagram__00000001_.TaskComposition_Shapes = append(__Diagram__00000001_.TaskComposition_Shapes, __TaskCompositionShape__00000029_)
	__Diagram__00000001_.TaskComposition_Shapes = append(__Diagram__00000001_.TaskComposition_Shapes, __TaskCompositionShape__00000032_)
	__Diagram__00000002_.Task_Shapes = append(__Diagram__00000002_.Task_Shapes, __TaskShape__00000024_)
	__Diagram__00000002_.Task_Shapes = append(__Diagram__00000002_.Task_Shapes, __TaskShape__00000025_)
	__Diagram__00000002_.Task_Shapes = append(__Diagram__00000002_.Task_Shapes, __TaskShape__00000026_)
	__Diagram__00000002_.Task_Shapes = append(__Diagram__00000002_.Task_Shapes, __TaskShape__00000027_)
	__Diagram__00000002_.Task_Shapes = append(__Diagram__00000002_.Task_Shapes, __TaskShape__00000028_)
	__Diagram__00000002_.Task_Shapes = append(__Diagram__00000002_.Task_Shapes, __TaskShape__00000029_)
	__Diagram__00000002_.Task_Shapes = append(__Diagram__00000002_.Task_Shapes, __TaskShape__00000030_)
	__Diagram__00000002_.Task_Shapes = append(__Diagram__00000002_.Task_Shapes, __TaskShape__00000031_)
	__Diagram__00000002_.TasksWhoseNodeIsExpanded = append(__Diagram__00000002_.TasksWhoseNodeIsExpanded, __Task__00000007_)
	__Diagram__00000002_.TasksWhoseNodeIsExpanded = append(__Diagram__00000002_.TasksWhoseNodeIsExpanded, __Task__00000005_)
	__Diagram__00000002_.TaskGroupShapes = append(__Diagram__00000002_.TaskGroupShapes, __TaskGroupShape__00000005_)
	__Diagram__00000002_.TaskGroupShapes = append(__Diagram__00000002_.TaskGroupShapes, __TaskGroupShape__00000006_)
	__Diagram__00000002_.TaskComposition_Shapes = append(__Diagram__00000002_.TaskComposition_Shapes, __TaskCompositionShape__00000014_)
	__Diagram__00000002_.TaskComposition_Shapes = append(__Diagram__00000002_.TaskComposition_Shapes, __TaskCompositionShape__00000015_)
	__Diagram__00000002_.TaskComposition_Shapes = append(__Diagram__00000002_.TaskComposition_Shapes, __TaskCompositionShape__00000016_)
	__Diagram__00000002_.TaskComposition_Shapes = append(__Diagram__00000002_.TaskComposition_Shapes, __TaskCompositionShape__00000017_)
	__Diagram__00000002_.TaskComposition_Shapes = append(__Diagram__00000002_.TaskComposition_Shapes, __TaskCompositionShape__00000018_)
	__Diagram__00000002_.TaskComposition_Shapes = append(__Diagram__00000002_.TaskComposition_Shapes, __TaskCompositionShape__00000019_)
	__Diagram__00000002_.TaskComposition_Shapes = append(__Diagram__00000002_.TaskComposition_Shapes, __TaskCompositionShape__00000020_)
	__Library__00000000_.SubLibraries = append(__Library__00000000_.SubLibraries, __Library__00000001_)
	__Library__00000000_.RootProducts = append(__Library__00000000_.RootProducts, __Product__00000000_)
	__Library__00000000_.RootTasks = append(__Library__00000000_.RootTasks, __Task__00000000_)
	__Library__00000000_.RootTasks = append(__Library__00000000_.RootTasks, __Task__00000001_)
	__Library__00000000_.RootTasks = append(__Library__00000000_.RootTasks, __Task__00000002_)
	__Library__00000000_.RootTaskGroups = append(__Library__00000000_.RootTaskGroups, __TaskGroup__00000000_)
	__Library__00000000_.RootTaskGroups = append(__Library__00000000_.RootTaskGroups, __TaskGroup__00000002_)
	__Library__00000000_.RootTaskGroups = append(__Library__00000000_.RootTaskGroups, __TaskGroup__00000003_)
	__Library__00000000_.Diagrams = append(__Library__00000000_.Diagrams, __Diagram__00000000_)
	__Library__00000001_.RootTasks = append(__Library__00000001_.RootTasks, __Task__00000007_)
	__Library__00000001_.RootTaskGroups = append(__Library__00000001_.RootTaskGroups, __TaskGroup__00000004_)
	__Library__00000001_.RootTaskGroups = append(__Library__00000001_.RootTaskGroups, __TaskGroup__00000005_)
	__Library__00000001_.Diagrams = append(__Library__00000001_.Diagrams, __Diagram__00000001_)
	__Library__00000001_.Diagrams = append(__Library__00000001_.Diagrams, __Diagram__00000002_)
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
	__Product__00000006_.SubProducts = append(__Product__00000006_.SubProducts, __Product__00000007_)
	__Product__00000006_.SubProducts = append(__Product__00000006_.SubProducts, __Product__00000008_)
	__Product__00000006_.SubProducts = append(__Product__00000006_.SubProducts, __Product__00000009_)
	__Product__00000006_.SubProducts = append(__Product__00000006_.SubProducts, __Product__00000010_)
	__Product__00000007_.ReferencedProduct = nil
	__Product__00000007_.SubProducts = append(__Product__00000007_.SubProducts, __Product__00000011_)
	__Product__00000007_.SubProducts = append(__Product__00000007_.SubProducts, __Product__00000012_)
	__Product__00000008_.ReferencedProduct = nil
	__Product__00000009_.ReferencedProduct = nil
	__Product__00000010_.ReferencedProduct = nil
	__Product__00000011_.ReferencedProduct = nil
	__Product__00000012_.ReferencedProduct = nil
	__ProductCompositionShape__00000062_.Product = __Product__00000001_
	__ProductCompositionShape__00000063_.Product = __Product__00000002_
	__ProductCompositionShape__00000064_.Product = __Product__00000005_
	__ProductCompositionShape__00000065_.Product = __Product__00000006_
	__ProductCompositionShape__00000066_.Product = __Product__00000007_
	__ProductCompositionShape__00000067_.Product = __Product__00000011_
	__ProductCompositionShape__00000068_.Product = __Product__00000012_
	__ProductCompositionShape__00000069_.Product = __Product__00000008_
	__ProductCompositionShape__00000070_.Product = __Product__00000009_
	__ProductCompositionShape__00000071_.Product = __Product__00000010_
	__ProductCompositionShape__00000072_.Product = __Product__00000003_
	__ProductCompositionShape__00000073_.Product = __Product__00000004_
	__ProductShape__00000073_.Product = __Product__00000000_
	__ProductShape__00000074_.Product = __Product__00000001_
	__ProductShape__00000075_.Product = __Product__00000002_
	__ProductShape__00000076_.Product = __Product__00000005_
	__ProductShape__00000077_.Product = __Product__00000006_
	__ProductShape__00000078_.Product = __Product__00000007_
	__ProductShape__00000079_.Product = __Product__00000011_
	__ProductShape__00000080_.Product = __Product__00000012_
	__ProductShape__00000081_.Product = __Product__00000008_
	__ProductShape__00000082_.Product = __Product__00000009_
	__ProductShape__00000083_.Product = __Product__00000010_
	__ProductShape__00000084_.Product = __Product__00000003_
	__ProductShape__00000085_.Product = __Product__00000004_
	__Task__00000000_.ReferencedTask = nil
	__Task__00000001_.ReferencedTask = nil
	__Task__00000002_.ReferencedTask = nil
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
	__TaskCompositionShape__00000014_.Task = __Task__00000005_
	__TaskCompositionShape__00000015_.Task = __Task__00000003_
	__TaskCompositionShape__00000016_.Task = __Task__00000004_
	__TaskCompositionShape__00000017_.Task = __Task__00000006_
	__TaskCompositionShape__00000018_.Task = __Task__00000008_
	__TaskCompositionShape__00000019_.Task = __Task__00000009_
	__TaskCompositionShape__00000020_.Task = __Task__00000010_
	__TaskCompositionShape__00000021_.Task = __Task__00000011_
	__TaskCompositionShape__00000029_.Task = __Task__00000005_
	__TaskCompositionShape__00000032_.Task = __Task__00000008_
	__TaskGroup__00000000_.Tasks = append(__TaskGroup__00000000_.Tasks, __Task__00000001_)
	__TaskGroup__00000000_.Tasks = append(__TaskGroup__00000000_.Tasks, __Task__00000000_)
	__TaskGroup__00000002_.Tasks = append(__TaskGroup__00000002_.Tasks, __Task__00000002_)
	__TaskGroup__00000004_.Tasks = append(__TaskGroup__00000004_.Tasks, __Task__00000005_)
	__TaskGroup__00000004_.Tasks = append(__TaskGroup__00000004_.Tasks, __Task__00000003_)
	__TaskGroup__00000004_.Tasks = append(__TaskGroup__00000004_.Tasks, __Task__00000004_)
	__TaskGroup__00000005_.Tasks = append(__TaskGroup__00000005_.Tasks, __Task__00000006_)
	__TaskGroup__00000005_.Tasks = append(__TaskGroup__00000005_.Tasks, __Task__00000008_)
	__TaskGroup__00000005_.Tasks = append(__TaskGroup__00000005_.Tasks, __Task__00000011_)
	__TaskGroupShape__00000003_.TaskGroup = __TaskGroup__00000004_
	__TaskGroupShape__00000004_.TaskGroup = __TaskGroup__00000005_
	__TaskGroupShape__00000005_.TaskGroup = __TaskGroup__00000004_
	__TaskGroupShape__00000006_.TaskGroup = __TaskGroup__00000005_
	__TaskShape__00000019_.Task = __Task__00000007_
	__TaskShape__00000021_.Task = __Task__00000003_
	__TaskShape__00000022_.Task = __Task__00000004_
	__TaskShape__00000023_.Task = __Task__00000006_
	__TaskShape__00000024_.Task = __Task__00000007_
	__TaskShape__00000025_.Task = __Task__00000005_
	__TaskShape__00000026_.Task = __Task__00000003_
	__TaskShape__00000027_.Task = __Task__00000004_
	__TaskShape__00000028_.Task = __Task__00000006_
	__TaskShape__00000029_.Task = __Task__00000008_
	__TaskShape__00000030_.Task = __Task__00000009_
	__TaskShape__00000031_.Task = __Task__00000010_
	__TaskShape__00000032_.Task = __Task__00000011_
	__TaskShape__00000042_.Task = __Task__00000005_
	__TaskShape__00000045_.Task = __Task__00000008_
}
