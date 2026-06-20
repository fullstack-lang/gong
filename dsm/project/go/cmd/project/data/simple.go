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

	__Diagram__00000000_ := (&models.Diagram{Name: `D1`}).Stage(stage)
	__Diagram__00000007_ := (&models.Diagram{Name: `Default Diagram copy`}).Stage(stage)

	__Library__00000000_ := (&models.Library{Name: ``}).Stage(stage)

	__Note__00000000_ := (&models.Note{Name: ``}).Stage(stage)

	__NoteShape__00000000_ := (&models.NoteShape{Name: `D1-`}).Stage(stage)
	__NoteShape__00000003_ := (&models.NoteShape{Name: `Default Diagram copy-`}).Stage(stage)

	__Product__00000000_ := (&models.Product{Name: `G3`}).Stage(stage)
	__Product__00000001_ := (&models.Product{Name: `G3.1`}).Stage(stage)

	__ProductCompositionShape__00000000_ := (&models.ProductCompositionShape{Name: `D1-G3-G3.1`}).Stage(stage)

	__ProductShape__00000000_ := (&models.ProductShape{Name: `D1-G3`}).Stage(stage)
	__ProductShape__00000001_ := (&models.ProductShape{Name: `D1-G3.1`}).Stage(stage)
	__ProductShape__00000023_ := (&models.ProductShape{Name: `Default Diagram copy-G3`}).Stage(stage)

	// insertion point for initialization of values

	__Diagram__00000000_.Name = `D1`
	__Diagram__00000000_.ComputedPrefix = `1`
	__Diagram__00000000_.IsExpanded = false
	__Diagram__00000000_.LayoutDirection = models.Vertical
	__Diagram__00000000_.IsChecked = false
	__Diagram__00000000_.IsEditable_ = true
	__Diagram__00000000_.IsShowPrefix = false
	__Diagram__00000000_.DefaultBoxWidth = 250.000000
	__Diagram__00000000_.DefaultBoxHeigth = 70.000000
	__Diagram__00000000_.Width = 759.078303
	__Diagram__00000000_.Height = 617.001756
	__Diagram__00000000_.IsPBSNodeExpanded = true
	__Diagram__00000000_.IsWBSNodeExpanded = false
	__Diagram__00000000_.IsTaskGroupsNodeExpanded = false
	__Diagram__00000000_.DateFormat = ``
	__Diagram__00000000_.IsNotesNodeExpanded = true
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

	__Diagram__00000007_.Name = `Default Diagram copy`
	__Diagram__00000007_.ComputedPrefix = `2`
	__Diagram__00000007_.IsExpanded = true
	__Diagram__00000007_.LayoutDirection = models.Vertical
	__Diagram__00000007_.IsChecked = true
	__Diagram__00000007_.IsEditable_ = true
	__Diagram__00000007_.IsShowPrefix = false
	__Diagram__00000007_.DefaultBoxWidth = 250.000000
	__Diagram__00000007_.DefaultBoxHeigth = 70.000000
	__Diagram__00000007_.Width = 759.078303
	__Diagram__00000007_.Height = 617.001756
	__Diagram__00000007_.IsPBSNodeExpanded = true
	__Diagram__00000007_.IsWBSNodeExpanded = false
	__Diagram__00000007_.IsTaskGroupsNodeExpanded = false
	__Diagram__00000007_.DateFormat = ``
	__Diagram__00000007_.IsNotesNodeExpanded = false
	__Diagram__00000007_.IsResourcesNodeExpanded = false
	__Diagram__00000007_.IsTimeDiagram = false
	__Diagram__00000007_.ComputedStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000007_.ComputedEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000007_.ComputedDuration = 0
	__Diagram__00000007_.UseManualStartAndEndDates = false
	__Diagram__00000007_.ManualStart, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000007_.ManualEnd, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "0001-01-01 00:00:00 +0000 UTC")
	__Diagram__00000007_.TimeStep = 0
	__Diagram__00000007_.TimeStepScale = ""
	__Diagram__00000007_.LaneHeight = 0.000000
	__Diagram__00000007_.RatioBarToLaneHeight = 0.000000
	__Diagram__00000007_.YTopMargin = 0.000000
	__Diagram__00000007_.XLeftText = 0.000000
	__Diagram__00000007_.TextHeight = 0.000000
	__Diagram__00000007_.XLeftLanes = 0.000000
	__Diagram__00000007_.XRightMargin = 0.000000
	__Diagram__00000007_.ArrowLengthToTheRightOfStartBar = 0.000000
	__Diagram__00000007_.ArrowTipLenght = 0.000000
	__Diagram__00000007_.TimeLine_Color = ``
	__Diagram__00000007_.TimeLine_FillOpacity = 0.000000
	__Diagram__00000007_.TimeLine_Stroke = ``
	__Diagram__00000007_.TimeLine_StrokeWidth = 0.000000
	__Diagram__00000007_.DrawVerticalTimeLines = false
	__Diagram__00000007_.Group_Stroke = ``
	__Diagram__00000007_.Group_StrokeWidth = 0.000000
	__Diagram__00000007_.Group_StrokeDashArray = ``
	__Diagram__00000007_.DateYOffset = 0.000000
	__Diagram__00000007_.AlignOnStartEndOnYearStart = false

	__Library__00000000_.Name = ``
	__Library__00000000_.NbPixPerCharacter = 8.000000
	__Library__00000000_.LogoSVGFile = ``
	__Library__00000000_.ComputedPrefix = ``
	__Library__00000000_.IsExpanded = true
	__Library__00000000_.LayoutDirection = models.Vertical
	__Library__00000000_.IsRootLibrary = true

	__Note__00000000_.Name = ``
	__Note__00000000_.ComputedPrefix = `1`
	__Note__00000000_.IsExpanded = false
	__Note__00000000_.LayoutDirection = models.Vertical

	__NoteShape__00000000_.Name = `D1-`
	__NoteShape__00000000_.OverideLayoutDirection = false
	__NoteShape__00000000_.LayoutDirection = models.Vertical
	__NoteShape__00000000_.X = 409.078303
	__NoteShape__00000000_.Y = 447.001756
	__NoteShape__00000000_.Width = 250.000000
	__NoteShape__00000000_.Height = 70.000000
	__NoteShape__00000000_.IsHidden = false

	__NoteShape__00000003_.Name = `Default Diagram copy-`
	__NoteShape__00000003_.OverideLayoutDirection = false
	__NoteShape__00000003_.LayoutDirection = models.Vertical
	__NoteShape__00000003_.X = 409.078303
	__NoteShape__00000003_.Y = 447.001756
	__NoteShape__00000003_.Width = 250.000000
	__NoteShape__00000003_.Height = 70.000000
	__NoteShape__00000003_.IsHidden = false

	__Product__00000000_.Name = `G3`
	__Product__00000000_.ComputedPrefix = `1`
	__Product__00000000_.IsExpanded = false
	__Product__00000000_.LayoutDirection = models.Vertical
	__Product__00000000_.IsImport = false
	__Product__00000000_.Description = ``
	__Product__00000000_.IsProducersNodeExpanded = false
	__Product__00000000_.IsConsumersNodeExpanded = false

	__Product__00000001_.Name = `G3.1`
	__Product__00000001_.ComputedPrefix = `1.1`
	__Product__00000001_.IsExpanded = false
	__Product__00000001_.LayoutDirection = models.Vertical
	__Product__00000001_.IsImport = false
	__Product__00000001_.Description = ``
	__Product__00000001_.IsProducersNodeExpanded = false
	__Product__00000001_.IsConsumersNodeExpanded = false

	__ProductCompositionShape__00000000_.Name = `D1-G3-G3.1`
	__ProductCompositionShape__00000000_.StartRatio = 0.500000
	__ProductCompositionShape__00000000_.EndRatio = 0.500000
	__ProductCompositionShape__00000000_.StartOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000000_.EndOrientation = models.ORIENTATION_VERTICAL
	__ProductCompositionShape__00000000_.CornerOffsetRatio = 1.680000
	__ProductCompositionShape__00000000_.IsHidden = false

	__ProductShape__00000000_.Name = `D1-G3`
	__ProductShape__00000000_.OverideLayoutDirection = false
	__ProductShape__00000000_.LayoutDirection = models.Vertical
	__ProductShape__00000000_.X = 159.729721
	__ProductShape__00000000_.Y = 112.660089
	__ProductShape__00000000_.Width = 250.000000
	__ProductShape__00000000_.Height = 70.000000
	__ProductShape__00000000_.IsHidden = false

	__ProductShape__00000001_.Name = `D1-G3.1`
	__ProductShape__00000001_.OverideLayoutDirection = false
	__ProductShape__00000001_.LayoutDirection = models.Vertical
	__ProductShape__00000001_.X = 159.729721
	__ProductShape__00000001_.Y = 252.660089
	__ProductShape__00000001_.Width = 250.000000
	__ProductShape__00000001_.Height = 70.000000
	__ProductShape__00000001_.IsHidden = false

	__ProductShape__00000023_.Name = `Default Diagram copy-G3`
	__ProductShape__00000023_.OverideLayoutDirection = false
	__ProductShape__00000023_.LayoutDirection = models.Vertical
	__ProductShape__00000023_.X = 159.729721
	__ProductShape__00000023_.Y = 112.660089
	__ProductShape__00000023_.Width = 250.000000
	__ProductShape__00000023_.Height = 70.000000
	__ProductShape__00000023_.IsHidden = false

	// insertion point for setup of pointers
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000000_)
	__Diagram__00000000_.Product_Shapes = append(__Diagram__00000000_.Product_Shapes, __ProductShape__00000001_)
	__Diagram__00000000_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000000_.ProductsWhoseNodeIsExpanded, __Product__00000000_)
	__Diagram__00000000_.ProductComposition_Shapes = append(__Diagram__00000000_.ProductComposition_Shapes, __ProductCompositionShape__00000000_)
	__Diagram__00000000_.Note_Shapes = append(__Diagram__00000000_.Note_Shapes, __NoteShape__00000000_)
	__Diagram__00000007_.Product_Shapes = append(__Diagram__00000007_.Product_Shapes, __ProductShape__00000023_)
	__Diagram__00000007_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000007_.ProductsWhoseNodeIsExpanded, __Product__00000000_)
	__Diagram__00000007_.ProductsWhoseNodeIsExpanded = append(__Diagram__00000007_.ProductsWhoseNodeIsExpanded, __Product__00000001_)
	__Diagram__00000007_.Note_Shapes = append(__Diagram__00000007_.Note_Shapes, __NoteShape__00000003_)
	__Library__00000000_.RootProducts = append(__Library__00000000_.RootProducts, __Product__00000000_)
	__Library__00000000_.Notes = append(__Library__00000000_.Notes, __Note__00000000_)
	__Library__00000000_.Diagrams = append(__Library__00000000_.Diagrams, __Diagram__00000000_)
	__Library__00000000_.Diagrams = append(__Library__00000000_.Diagrams, __Diagram__00000007_)
	__NoteShape__00000000_.Note = __Note__00000000_
	__NoteShape__00000003_.Note = __Note__00000000_
	__Product__00000000_.ReferencedProduct = nil
	__Product__00000000_.SubProducts = append(__Product__00000000_.SubProducts, __Product__00000001_)
	__Product__00000001_.ReferencedProduct = nil
	__ProductCompositionShape__00000000_.Product = __Product__00000001_
	__ProductShape__00000000_.Product = __Product__00000000_
	__ProductShape__00000001_.Product = __Product__00000001_
	__ProductShape__00000023_.Product = __Product__00000000_
}
