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
	__Diagram__00000002_ := (&models.Diagram{Name: `Default Diagram`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: ``}).Stage(stage)
	__Library__00000001_ := (&models.Library{Name: `L1`}).Stage(stage)
	__Library__00000002_ := (&models.Library{Name: `L1.1`}).Stage(stage)

	__Product__00000000_ := (&models.Product{Name: `A`}).Stage(stage)
	__Product__00000001_ := (&models.Product{Name: `A.1`}).Stage(stage)
	__Product__00000002_ := (&models.Product{Name: `A.1.1`}).Stage(stage)
	__Product__00000003_ := (&models.Product{Name: `A.2`}).Stage(stage)
	__Product__00000004_ := (&models.Product{Name: `A.3`}).Stage(stage)
	__Product__00000005_ := (&models.Product{Name: ``}).Stage(stage)
	__Product__00000006_ := (&models.Product{Name: `LA1`}).Stage(stage)

	__ProductCompositionShape__00000001_ := (&models.ProductCompositionShape{Name: `Default Diagram-A-A.1`}).Stage(stage)
	__ProductCompositionShape__00000002_ := (&models.ProductCompositionShape{Name: `Default Diagram-A.1-A.1.1`}).Stage(stage)
	__ProductCompositionShape__00000003_ := (&models.ProductCompositionShape{Name: `Default Diagram-A-A.2`}).Stage(stage)
	__ProductCompositionShape__00000004_ := (&models.ProductCompositionShape{Name: `Default Diagram-A-A.3`}).Stage(stage)
	__ProductCompositionShape__00000005_ := (&models.ProductCompositionShape{Name: `Default Diagram-A-`}).Stage(stage)

	__ProductShape__00000001_ := (&models.ProductShape{Name: `Default Diagram-A`}).Stage(stage)
	__ProductShape__00000002_ := (&models.ProductShape{Name: `Default Diagram-A.1`}).Stage(stage)
	__ProductShape__00000003_ := (&models.ProductShape{Name: `Default Diagram-A.1.1`}).Stage(stage)
	__ProductShape__00000004_ := (&models.ProductShape{Name: `Default Diagram-A.2`}).Stage(stage)
	__ProductShape__00000005_ := (&models.ProductShape{Name: `Default Diagram-A.3`}).Stage(stage)
	__ProductShape__00000006_ := (&models.ProductShape{Name: `Default Diagram-`}).Stage(stage)
	__ProductShape__00000007_ := (&models.ProductShape{Name: `Default Diagram-LA1`}).Stage(stage)

	__Resource__00000000_ := (&models.Resource{Name: `R1`}).Stage(stage)
	__Resource__00000001_ := (&models.Resource{Name: `R1.1`}).Stage(stage)

	__ResourceCompositionShape__00000000_ := (&models.ResourceCompositionShape{Name: `Default Diagram-R1-R1.1`}).Stage(stage)

	__ResourceShape__00000000_ := (&models.ResourceShape{Name: `Default Diagram-R1`}).Stage(stage)
	__ResourceShape__00000001_ := (&models.ResourceShape{Name: `Default Diagram-R1.1`}).Stage(stage)

	__Task__00000000_ := (&models.Task{Name: `W1`}).Stage(stage)
	__Task__00000001_ := (&models.Task{Name: `W1.1`}).Stage(stage)

	__TaskCompositionShape__00000000_ := (&models.TaskCompositionShape{Name: `Default Diagram-W1-W1.1`}).Stage(stage)

	__TaskShape__00000000_ := (&models.TaskShape{Name: `Default Diagram-W1`}).Stage(stage)
	__TaskShape__00000001_ := (&models.TaskShape{Name: `Default Diagram-W1.1`}).Stage(stage)

	// insertion point for initialization of values

	__Diagram__00000000_.Name = `Default Diagram`
	__Diagram__00000000_.DefaultBoxWidth = 250.000000
	__Diagram__00000000_.DefaultBoxHeigth = 70.000000
	__Diagram__00000000_.DateFormat = ``
	__Diagram__00000000_.Width = 1395.056982
	__Diagram__00000000_.Height = 1047.068307
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
	__Diagram__00000000_.ComputedPrefix = `1`
	__Diagram__00000000_.IsExpanded = false
	__Diagram__00000000_.IsChecked = true
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.IsShowPrefix = false
	__Diagram__00000000_.IsPBSNodeExpanded = true
	__Diagram__00000000_.IsWBSNodeExpanded = true
	__Diagram__00000000_.IsTaskGroupsNodeExpanded = false
	__Diagram__00000000_.IsNotesNodeExpanded = false
	__Diagram__00000000_.IsResourcesNodeExpanded = true

	__Diagram__00000001_.Name = `Default Diagram`
	__Diagram__00000001_.DefaultBoxWidth = 250.000000
	__Diagram__00000001_.DefaultBoxHeigth = 70.000000
	__Diagram__00000001_.DateFormat = ``
	__Diagram__00000001_.Width = 533.836804
	__Diagram__00000001_.Height = 321.095882
	__Diagram__00000001_.IsTimeDiagram = false
	__Diagram__00000001_.ComputedStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000001_.ComputedEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000001_.ComputedDuration = 0
	__Diagram__00000001_.UseManualStartAndEndDates = false
	__Diagram__00000001_.ManualStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000001_.ManualEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000001_.TimeStep = 0
	__Diagram__00000001_.TimeStepScale = ""
	__Diagram__00000001_.LaneHeight = 0.000000
	__Diagram__00000001_.RatioBarToLaneHeight = 0.000000
	__Diagram__00000001_.YTopMargin = 0.000000
	__Diagram__00000001_.XLeftText = 0.000000
	__Diagram__00000001_.TextHeight = 0.000000
	__Diagram__00000001_.XLeftLanes = 0.000000
	__Diagram__00000001_.XRightMargin = 0.000000
	__Diagram__00000001_.ArrowLengthToTheRightOfStartBar = 0.000000
	__Diagram__00000001_.ArrowTipLenght = 0.000000
	__Diagram__00000001_.TimeLine_Color = ``
	__Diagram__00000001_.TimeLine_FillOpacity = 0.000000
	__Diagram__00000001_.TimeLine_Stroke = ``
	__Diagram__00000001_.TimeLine_StrokeWidth = 0.000000
	__Diagram__00000001_.DrawVerticalTimeLines = false
	__Diagram__00000001_.Group_Stroke = ``
	__Diagram__00000001_.Group_StrokeWidth = 0.000000
	__Diagram__00000001_.Group_StrokeDashArray = ``
	__Diagram__00000001_.DateYOffset = 0.000000
	__Diagram__00000001_.AlignOnStartEndOnYearStart = false
	__Diagram__00000001_.ComputedPrefix = `1`
	__Diagram__00000001_.IsExpanded = false
	__Diagram__00000001_.IsChecked = false
	__Diagram__00000001_.IsEditable_ = true
	__Diagram__00000001_.IsShowPrefix = false
	__Diagram__00000001_.IsPBSNodeExpanded = true
	__Diagram__00000001_.IsWBSNodeExpanded = false
	__Diagram__00000001_.IsTaskGroupsNodeExpanded = false
	__Diagram__00000001_.IsNotesNodeExpanded = false
	__Diagram__00000001_.IsResourcesNodeExpanded = false

	__Diagram__00000002_.Name = `Default Diagram`
	__Diagram__00000002_.DefaultBoxWidth = 250.000000
	__Diagram__00000002_.DefaultBoxHeigth = 70.000000
	__Diagram__00000002_.DateFormat = ``
	__Diagram__00000002_.Width = 100.000000
	__Diagram__00000002_.Height = 100.000000
	__Diagram__00000002_.IsTimeDiagram = false
	__Diagram__00000002_.ComputedStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000002_.ComputedEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000002_.ComputedDuration = 0
	__Diagram__00000002_.UseManualStartAndEndDates = false
	__Diagram__00000002_.ManualStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000002_.ManualEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000002_.TimeStep = 0
	__Diagram__00000002_.TimeStepScale = ""
	__Diagram__00000002_.LaneHeight = 0.000000
	__Diagram__00000002_.RatioBarToLaneHeight = 0.000000
	__Diagram__00000002_.YTopMargin = 0.000000
	__Diagram__00000002_.XLeftText = 0.000000
	__Diagram__00000002_.TextHeight = 0.000000
	__Diagram__00000002_.XLeftLanes = 0.000000
	__Diagram__00000002_.XRightMargin = 0.000000
	__Diagram__00000002_.ArrowLengthToTheRightOfStartBar = 0.000000
	__Diagram__00000002_.ArrowTipLenght = 0.000000
	__Diagram__00000002_.TimeLine_Color = ``
	__Diagram__00000002_.TimeLine_FillOpacity = 0.000000
	__Diagram__00000002_.TimeLine_Stroke = ``
	__Diagram__00000002_.TimeLine_StrokeWidth = 0.000000
	__Diagram__00000002_.DrawVerticalTimeLines = false
	__Diagram__00000002_.Group_Stroke = ``
	__Diagram__00000002_.Group_StrokeWidth = 0.000000
	__Diagram__00000002_.Group_StrokeDashArray = ``
	__Diagram__00000002_.DateYOffset = 0.000000
	__Diagram__00000002_.AlignOnStartEndOnYearStart = false
	__Diagram__00000002_.ComputedPrefix = `1`
	__Diagram__00000002_.IsExpanded = false
	__Diagram__00000002_.IsChecked = false
	__Diagram__00000002_.IsEditable_ = true
	__Diagram__00000002_.IsShowPrefix = false
	__Diagram__00000002_.IsPBSNodeExpanded = false
	__Diagram__00000002_.IsWBSNodeExpanded = false
	__Diagram__00000002_.IsTaskGroupsNodeExpanded = false
	__Diagram__00000002_.IsNotesNodeExpanded = false
	__Diagram__00000002_.IsResourcesNodeExpanded = false

	__Library__00000000_.Name = ``
	__Library__00000000_.NbPixPerCharacter = 8.000000
	__Library__00000000_.LogoSVGFile = ``
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.IsExpanded = true
	__Library__00000000_.IsRootLibrary = true

	__Library__00000001_.Name = `L1`
	__Library__00000001_.NbPixPerCharacter = 0.000000
	__Library__00000001_.LogoSVGFile = ``
	__Library__00000001_.ComputedPrefix = ``
	__Library__00000001_.IsExpanded = false
	__Library__00000001_.IsRootLibrary = false

	__Library__00000002_.Name = `L1.1`
	__Library__00000002_.NbPixPerCharacter = 0.000000
	__Library__00000002_.LogoSVGFile = ``
	__Library__00000002_.ComputedPrefix = ``
	__Library__00000002_.IsExpanded = false
	__Library__00000002_.IsRootLibrary = false

	__Product__00000000_.Name = `A`
	__Product__00000000_.Description = ``
	__Product__00000000_.IsProducersNodeExpanded = false
	__Product__00000000_.IsConsumersNodeExpanded = false
	__Product__00000000_.IsImport = false
	__Product__00000000_.ComputedPrefix = `1`
	__Product__00000000_.IsExpanded = false
	__Product__00000000_.LayoutDirection = models.Vertical

	__Product__00000001_.Name = `A.1`
	__Product__00000001_.Description = ``
	__Product__00000001_.IsProducersNodeExpanded = false
	__Product__00000001_.IsConsumersNodeExpanded = false
	__Product__00000001_.IsImport = false
	__Product__00000001_.ComputedPrefix = `1.1`
	__Product__00000001_.IsExpanded = false
	__Product__00000001_.LayoutDirection = models.Vertical

	__Product__00000002_.Name = `A.1.1`
	__Product__00000002_.Description = ``
	__Product__00000002_.IsProducersNodeExpanded = false
	__Product__00000002_.IsConsumersNodeExpanded = false
	__Product__00000002_.IsImport = false
	__Product__00000002_.ComputedPrefix = `1.1.1`
	__Product__00000002_.IsExpanded = false
	__Product__00000002_.LayoutDirection = models.Vertical

	__Product__00000003_.Name = `A.2`
	__Product__00000003_.Description = ``
	__Product__00000003_.IsProducersNodeExpanded = false
	__Product__00000003_.IsConsumersNodeExpanded = false
	__Product__00000003_.IsImport = false
	__Product__00000003_.ComputedPrefix = `1.2`
	__Product__00000003_.IsExpanded = false
	__Product__00000003_.LayoutDirection = models.Vertical

	__Product__00000004_.Name = `A.3`
	__Product__00000004_.Description = ``
	__Product__00000004_.IsProducersNodeExpanded = false
	__Product__00000004_.IsConsumersNodeExpanded = false
	__Product__00000004_.IsImport = false
	__Product__00000004_.ComputedPrefix = `1.3`
	__Product__00000004_.IsExpanded = false
	__Product__00000004_.LayoutDirection = models.Vertical

	__Product__00000005_.Name = ``
	__Product__00000005_.Description = ``
	__Product__00000005_.IsProducersNodeExpanded = false
	__Product__00000005_.IsConsumersNodeExpanded = false
	__Product__00000005_.IsImport = false
	__Product__00000005_.ComputedPrefix = `1.4`
	__Product__00000005_.IsExpanded = false
	__Product__00000005_.LayoutDirection = models.Vertical

	__Product__00000006_.Name = `LA1`
	__Product__00000006_.Description = ``
	__Product__00000006_.IsProducersNodeExpanded = false
	__Product__00000006_.IsConsumersNodeExpanded = false
	__Product__00000006_.IsImport = false
	__Product__00000006_.ComputedPrefix = `1`
	__Product__00000006_.IsExpanded = false
	__Product__00000006_.LayoutDirection = models.Vertical

	__ProductCompositionShape__00000001_.Name = `Default Diagram-A-A.1`
	__ProductCompositionShape__00000001_.StartRatio = 0.500000
	__ProductCompositionShape__00000001_.EndRatio = 0.500000
	__ProductCompositionShape__00000001_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000001_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000001_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000001_.IsHidden = false

	__ProductCompositionShape__00000002_.Name = `Default Diagram-A.1-A.1.1`
	__ProductCompositionShape__00000002_.StartRatio = 0.500000
	__ProductCompositionShape__00000002_.EndRatio = 0.500000
	__ProductCompositionShape__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000002_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000002_.IsHidden = false

	__ProductCompositionShape__00000003_.Name = `Default Diagram-A-A.2`
	__ProductCompositionShape__00000003_.StartRatio = 0.500000
	__ProductCompositionShape__00000003_.EndRatio = 0.500000
	__ProductCompositionShape__00000003_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000003_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000003_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000003_.IsHidden = false

	__ProductCompositionShape__00000004_.Name = `Default Diagram-A-A.3`
	__ProductCompositionShape__00000004_.StartRatio = 0.500000
	__ProductCompositionShape__00000004_.EndRatio = 0.500000
	__ProductCompositionShape__00000004_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000004_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000004_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000004_.IsHidden = false

	__ProductCompositionShape__00000005_.Name = `Default Diagram-A-`
	__ProductCompositionShape__00000005_.StartRatio = 0.500000
	__ProductCompositionShape__00000005_.EndRatio = 0.500000
	__ProductCompositionShape__00000005_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000005_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000005_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000005_.IsHidden = false

	__ProductShape__00000001_.Name = `Default Diagram-A`
	__ProductShape__00000001_.OverideLayoutDirection = false
	__ProductShape__00000001_.LayoutDirection = models.Vertical
	__ProductShape__00000001_.X = 145.056982
	__ProductShape__00000001_.Y = 125.618199
	__ProductShape__00000001_.Width = 250.000000
	__ProductShape__00000001_.Height = 70.000000
	__ProductShape__00000001_.IsHidden = false

	__ProductShape__00000002_.Name = `Default Diagram-A.1`
	__ProductShape__00000002_.OverideLayoutDirection = false
	__ProductShape__00000002_.LayoutDirection = models.Vertical
	__ProductShape__00000002_.X = 185.056982
	__ProductShape__00000002_.Y = 244.618199
	__ProductShape__00000002_.Width = 250.000000
	__ProductShape__00000002_.Height = 70.000000
	__ProductShape__00000002_.IsHidden = false

	__ProductShape__00000003_.Name = `Default Diagram-A.1.1`
	__ProductShape__00000003_.OverideLayoutDirection = false
	__ProductShape__00000003_.LayoutDirection = models.Vertical
	__ProductShape__00000003_.X = 145.056982
	__ProductShape__00000003_.Y = 405.618199
	__ProductShape__00000003_.Width = 250.000000
	__ProductShape__00000003_.Height = 70.000000
	__ProductShape__00000003_.IsHidden = false

	__ProductShape__00000004_.Name = `Default Diagram-A.2`
	__ProductShape__00000004_.OverideLayoutDirection = false
	__ProductShape__00000004_.LayoutDirection = models.Vertical
	__ProductShape__00000004_.X = 445.056982
	__ProductShape__00000004_.Y = 265.618199
	__ProductShape__00000004_.Width = 250.000000
	__ProductShape__00000004_.Height = 70.000000
	__ProductShape__00000004_.IsHidden = false

	__ProductShape__00000005_.Name = `Default Diagram-A.3`
	__ProductShape__00000005_.OverideLayoutDirection = false
	__ProductShape__00000005_.LayoutDirection = models.Vertical
	__ProductShape__00000005_.X = 745.056982
	__ProductShape__00000005_.Y = 265.618199
	__ProductShape__00000005_.Width = 250.000000
	__ProductShape__00000005_.Height = 70.000000
	__ProductShape__00000005_.IsHidden = false

	__ProductShape__00000006_.Name = `Default Diagram-`
	__ProductShape__00000006_.OverideLayoutDirection = false
	__ProductShape__00000006_.LayoutDirection = models.Vertical
	__ProductShape__00000006_.X = 1045.056982
	__ProductShape__00000006_.Y = 265.618199
	__ProductShape__00000006_.Width = 250.000000
	__ProductShape__00000006_.Height = 70.000000
	__ProductShape__00000006_.IsHidden = false

	__ProductShape__00000007_.Name = `Default Diagram-LA1`
	__ProductShape__00000007_.OverideLayoutDirection = false
	__ProductShape__00000007_.LayoutDirection = models.Vertical
	__ProductShape__00000007_.X = 183.836804
	__ProductShape__00000007_.Y = 151.095882
	__ProductShape__00000007_.Width = 250.000000
	__ProductShape__00000007_.Height = 70.000000
	__ProductShape__00000007_.IsHidden = false

	__Resource__00000000_.Name = `R1`
	__Resource__00000000_.Description = ``
	__Resource__00000000_.ComputedPrefix = `1`
	__Resource__00000000_.IsExpanded = false
	__Resource__00000000_.LayoutDirection = models.Vertical
	__Resource__00000000_.IsImport = false

	__Resource__00000001_.Name = `R1.1`
	__Resource__00000001_.Description = ``
	__Resource__00000001_.ComputedPrefix = `1.1`
	__Resource__00000001_.IsExpanded = false
	__Resource__00000001_.LayoutDirection = models.Vertical
	__Resource__00000001_.IsImport = false

	__ResourceCompositionShape__00000000_.Name = `Default Diagram-R1-R1.1`
	__ResourceCompositionShape__00000000_.StartRatio = 0.500000
	__ResourceCompositionShape__00000000_.EndRatio = 0.500000
	__ResourceCompositionShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__ResourceCompositionShape__00000000_.CornerOffsetRatio = 1.680000
	__ResourceCompositionShape__00000000_.IsHidden = false

	__ResourceShape__00000000_.Name = `Default Diagram-R1`
	__ResourceShape__00000000_.OverideLayoutDirection = false
	__ResourceShape__00000000_.LayoutDirection = models.Vertical
	__ResourceShape__00000000_.X = 414.182087
	__ResourceShape__00000000_.Y = 605.740632
	__ResourceShape__00000000_.Width = 250.000000
	__ResourceShape__00000000_.Height = 70.000000
	__ResourceShape__00000000_.IsHidden = false

	__ResourceShape__00000001_.Name = `Default Diagram-R1.1`
	__ResourceShape__00000001_.OverideLayoutDirection = false
	__ResourceShape__00000001_.LayoutDirection = models.Vertical
	__ResourceShape__00000001_.X = 613.182087
	__ResourceShape__00000001_.Y = 751.740586
	__ResourceShape__00000001_.Width = 250.000000
	__ResourceShape__00000001_.Height = 70.000000
	__ResourceShape__00000001_.IsHidden = false

	__Task__00000000_.Name = `W1`
	__Task__00000000_.Description = ``
	__Task__00000000_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000000_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000000_.DurationYears = 0.000000
	__Task__00000000_.DurationMonths = 0.000000
	__Task__00000000_.DurationWeeks = 0.000000
	__Task__00000000_.DurationDays = 0.000000
	__Task__00000000_.DurationHours = 0.000000
	__Task__00000000_.IsEndDateComputedFromDuration = false
	__Task__00000000_.IsStartDateComputedFromPredecessors = false
	__Task__00000000_.IsMilestone = false
	__Task__00000000_.IsInputsNodeExpanded = false
	__Task__00000000_.IsOutputsNodeExpanded = false
	__Task__00000000_.IsWithCompletion = true
	__Task__00000000_.Completion = models.PERCENT_050
	__Task__00000000_.DisplayVerticalBar = false
	__Task__00000000_.TextPosition = ""
	__Task__00000000_.XOffset = 0.000000
	__Task__00000000_.YOffset = 0.000000
	__Task__00000000_.IsImport = false
	__Task__00000000_.ComputedPrefix = `1`
	__Task__00000000_.IsExpanded = false
	__Task__00000000_.LayoutDirection = models.Vertical

	__Task__00000001_.Name = `W1.1`
	__Task__00000001_.Description = ``
	__Task__00000001_.Start, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000001_.End, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Task__00000001_.DurationYears = 0.000000
	__Task__00000001_.DurationMonths = 0.000000
	__Task__00000001_.DurationWeeks = 0.000000
	__Task__00000001_.DurationDays = 0.000000
	__Task__00000001_.DurationHours = 0.000000
	__Task__00000001_.IsEndDateComputedFromDuration = false
	__Task__00000001_.IsStartDateComputedFromPredecessors = false
	__Task__00000001_.IsMilestone = false
	__Task__00000001_.IsInputsNodeExpanded = false
	__Task__00000001_.IsOutputsNodeExpanded = false
	__Task__00000001_.IsWithCompletion = false
	__Task__00000001_.Completion = ""
	__Task__00000001_.DisplayVerticalBar = false
	__Task__00000001_.TextPosition = ""
	__Task__00000001_.XOffset = 0.000000
	__Task__00000001_.YOffset = 0.000000
	__Task__00000001_.IsImport = false
	__Task__00000001_.ComputedPrefix = `1.1`
	__Task__00000001_.IsExpanded = false
	__Task__00000001_.LayoutDirection = models.Vertical

	__TaskCompositionShape__00000000_.Name = `Default Diagram-W1-W1.1`
	__TaskCompositionShape__00000000_.StartRatio = 0.500000
	__TaskCompositionShape__00000000_.EndRatio = 0.500000
	__TaskCompositionShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__TaskCompositionShape__00000000_.CornerOffsetRatio = 1.680000
	__TaskCompositionShape__00000000_.IsHidden = false

	__TaskShape__00000000_.Name = `Default Diagram-W1`
	__TaskShape__00000000_.IsShowDate = false
	__TaskShape__00000000_.OverideLayoutDirection = false
	__TaskShape__00000000_.LayoutDirection = models.Vertical
	__TaskShape__00000000_.X = 110.726367
	__TaskShape__00000000_.Y = 685.068292
	__TaskShape__00000000_.Width = 250.000000
	__TaskShape__00000000_.Height = 70.000000
	__TaskShape__00000000_.IsHidden = false

	__TaskShape__00000001_.Name = `Default Diagram-W1.1`
	__TaskShape__00000001_.IsShowDate = false
	__TaskShape__00000001_.OverideLayoutDirection = false
	__TaskShape__00000001_.LayoutDirection = models.Vertical
	__TaskShape__00000001_.X = 214.726367
	__TaskShape__00000001_.Y = 877.068307
	__TaskShape__00000001_.Width = 250.000000
	__TaskShape__00000001_.Height = 70.000000
	__TaskShape__00000001_.IsHidden = false

	// insertion point for setup of pointers
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000001_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000002_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000003_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000004_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000005_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000006_)
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000000_.ProductsWhoseNodeIsExpanded, __Product__00000001_)
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000000_.ProductsWhoseNodeIsExpanded, __Product__00000000_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000001_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000002_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000003_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000004_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000005_)
	__Diagram__00000000_.Task_Shapes = append(__Diagram__00000000_.Task_Shapes, __TaskShape__00000000_)
	__Diagram__00000000_.Task_Shapes = append(__Diagram__00000000_.Task_Shapes, __TaskShape__00000001_)
	__Diagram__00000000_.TasksWhoseNodeIsExpanded = append(__Diagram__00000000_.TasksWhoseNodeIsExpanded, __Task__00000000_)
	__Diagram__00000000_.TaskComposition_Shapes = append(__Diagram__00000000_.TaskComposition_Shapes, __TaskCompositionShape__00000000_)
	__Diagram__00000000_.Resource_Shapes = append(__Diagram__00000000_.Resource_Shapes, __ResourceShape__00000000_)
	__Diagram__00000000_.Resource_Shapes = append(__Diagram__00000000_.Resource_Shapes, __ResourceShape__00000001_)
	__Diagram__00000000_.ResourcesWhoseNodeIsExpanded = append(__Diagram__00000000_.ResourcesWhoseNodeIsExpanded, __Resource__00000000_)
	__Diagram__00000000_.ResourceComposition_Shapes = append(__Diagram__00000000_.ResourceComposition_Shapes, __ResourceCompositionShape__00000000_)
	__Diagram__00000001_.Product_Shapes = append(__Diagram__00000001_.Product_Shapes, __ProductShape__00000007_)
	__Library__00000000_.SubLibraries = append(__Library__00000000_.SubLibraries, __Library__00000001_)
	__Library__00000000_.RootProducts = append(__Library__00000000_.RootProducts, __Product__00000000_)
	__Library__00000000_.RootTasks = append(__Library__00000000_.RootTasks, __Task__00000000_)
	__Library__00000000_.RootResources = append(__Library__00000000_.RootResources, __Resource__00000000_)
	__Library__00000000_.Diagrams = append(__Library__00000000_.Diagrams, __Diagram__00000000_)
	__Library__00000001_.SubLibraries = append(__Library__00000001_.SubLibraries, __Library__00000002_)
	__Library__00000001_.RootProducts = append(__Library__00000001_.RootProducts, __Product__00000006_)
	__Library__00000001_.Diagrams = append(__Library__00000001_.Diagrams, __Diagram__00000001_)
	__Library__00000002_.Diagrams = append(__Library__00000002_.Diagrams, __Diagram__00000002_)
	__Product__00000000_.SubProducts = append(__Product__00000000_.SubProducts, __Product__00000001_)
	__Product__00000000_.SubProducts = append(__Product__00000000_.SubProducts, __Product__00000003_)
	__Product__00000000_.SubProducts = append(__Product__00000000_.SubProducts, __Product__00000004_)
	__Product__00000000_.SubProducts = append(__Product__00000000_.SubProducts, __Product__00000005_)
	__Product__00000000_.ReferencedProduct = nil
	__Product__00000001_.SubProducts = append(__Product__00000001_.SubProducts, __Product__00000002_)
	__Product__00000001_.ReferencedProduct = nil
	__Product__00000002_.ReferencedProduct = nil
	__Product__00000003_.ReferencedProduct = nil
	__Product__00000004_.ReferencedProduct = nil
	__Product__00000005_.ReferencedProduct = nil
	__Product__00000006_.ReferencedProduct = nil
	__ProductCompositionShape__00000001_.Product = __Product__00000001_
	__ProductCompositionShape__00000002_.Product = __Product__00000002_
	__ProductCompositionShape__00000003_.Product = __Product__00000003_
	__ProductCompositionShape__00000004_.Product = __Product__00000004_
	__ProductCompositionShape__00000005_.Product = __Product__00000005_
	__ProductShape__00000001_.Product = __Product__00000000_
	__ProductShape__00000002_.Product = __Product__00000001_
	__ProductShape__00000003_.Product = __Product__00000002_
	__ProductShape__00000004_.Product = __Product__00000003_
	__ProductShape__00000005_.Product = __Product__00000004_
	__ProductShape__00000006_.Product = __Product__00000005_
	__ProductShape__00000007_.Product = __Product__00000006_
	__Resource__00000000_.SubResources = append(__Resource__00000000_.SubResources, __Resource__00000001_)
	__Resource__00000000_.ReferencedResource = nil
	__Resource__00000001_.ReferencedResource = nil
	__ResourceCompositionShape__00000000_.Resource = __Resource__00000001_
	__ResourceShape__00000000_.Resource = __Resource__00000000_
	__ResourceShape__00000001_.Resource = __Resource__00000001_
	__Task__00000000_.ReferencedTask = nil
	__Task__00000000_.SubTasks = append(__Task__00000000_.SubTasks, __Task__00000001_)
	__Task__00000001_.ReferencedTask = nil
	__TaskCompositionShape__00000000_.Task = __Task__00000001_
	__TaskShape__00000000_.Task = __Task__00000000_
	__TaskShape__00000001_.Task = __Task__00000001_
}
