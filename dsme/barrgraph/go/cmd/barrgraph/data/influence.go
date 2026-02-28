package main

import (
	"time"

	"github.com/fullstack-lang/gong/dsme/barrgraph/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage

	__Category1__00000000_ := (&models.Category1{Name: `Intensité de la Concurrence`}).Stage(stage)
	__Category1__00000001_ := (&models.Category1{Name: `Niveau des prix`}).Stage(stage)
	__Category1__00000002_ := (&models.Category1{Name: `Hauteur des barrières à l'entrée`}).Stage(stage)
	__Category1__00000003_ := (&models.Category1{Name: `Probabilité d'Oligopoles`}).Stage(stage)
	__Category1__00000004_ := (&models.Category1{Name: `Intensité de la politique concurrentielle`}).Stage(stage)
	__Category1__00000005_ := (&models.Category1{Name: `Probabilité de monopole`}).Stage(stage)

	__Category1Shape__00000000_ := (&models.Category1Shape{Name: ``}).Stage(stage)
	__Category1Shape__00000001_ := (&models.Category1Shape{Name: ``}).Stage(stage)
	__Category1Shape__00000002_ := (&models.Category1Shape{Name: ``}).Stage(stage)
	__Category1Shape__00000003_ := (&models.Category1Shape{Name: ``}).Stage(stage)
	__Category1Shape__00000004_ := (&models.Category1Shape{Name: ``}).Stage(stage)
	__Category1Shape__00000005_ := (&models.Category1Shape{Name: ``}).Stage(stage)

	__Desk__00000000_ := (&models.Desk{Name: `Desk`}).Stage(stage)

	__Diagram__00000000_ := (&models.Diagram{Name: `Default`}).Stage(stage)

	__Influence__00000000_ := (&models.Influence{Name: `Intensité de la Concurrence to Niveau des prix`}).Stage(stage)
	__Influence__00000001_ := (&models.Influence{Name: `Hauteur des barrières à l'entrée to Intensité de la Concurrence`}).Stage(stage)
	__Influence__00000002_ := (&models.Influence{Name: `Intensité de la Concurrence to Probabilité d'Oligopoles`}).Stage(stage)
	__Influence__00000003_ := (&models.Influence{Name: `Hauteur des barrières à l'entrée to Probabilité d'Oligopoles`}).Stage(stage)
	__Influence__00000004_ := (&models.Influence{Name: `Intensité de la politique concurrentielle to Hauteur des barrières à l'entrée`}).Stage(stage)
	__Influence__00000005_ := (&models.Influence{Name: `Intensité de la Concurrence to Probabilité de monopole`}).Stage(stage)

	__InfluenceShape__00000000_ := (&models.InfluenceShape{Name: `Intensité de la Concurrence to Niveau des prix`}).Stage(stage)
	__InfluenceShape__00000001_ := (&models.InfluenceShape{Name: `Hauteur des barrières à l'entrée to Intensité de la Concurrence`}).Stage(stage)
	__InfluenceShape__00000002_ := (&models.InfluenceShape{Name: `Intensité de la Concurrence to Probabilité d'Oligopoles`}).Stage(stage)
	__InfluenceShape__00000003_ := (&models.InfluenceShape{Name: `Hauteur des barrières à l'entrée to Probabilité d'Oligopoles`}).Stage(stage)
	__InfluenceShape__00000004_ := (&models.InfluenceShape{Name: `Intensité de la politique concurrentielle to Hauteur des barrières à l'entrée`}).Stage(stage)
	__InfluenceShape__00000005_ := (&models.InfluenceShape{Name: `Intensité de la Concurrence to Probabilité de monopole`}).Stage(stage)

	// insertion point for initialization of values

	__Category1__00000000_.Name = `Intensité de la Concurrence`

	__Category1__00000001_.Name = `Niveau des prix`

	__Category1__00000002_.Name = `Hauteur des barrières à l'entrée`

	__Category1__00000003_.Name = `Probabilité d'Oligopoles`

	__Category1__00000004_.Name = `Intensité de la politique concurrentielle`

	__Category1__00000005_.Name = `Probabilité de monopole`

	__Category1Shape__00000000_.Name = ``
	__Category1Shape__00000000_.X = 200.000000
	__Category1Shape__00000000_.Y = 218.000000
	__Category1Shape__00000000_.Width = 131.000000
	__Category1Shape__00000000_.Height = 80.000000

	__Category1Shape__00000001_.Name = ``
	__Category1Shape__00000001_.X = 120.999938
	__Category1Shape__00000001_.Y = 454.000016
	__Category1Shape__00000001_.Width = 147.000000
	__Category1Shape__00000001_.Height = 80.000000

	__Category1Shape__00000002_.Name = ``
	__Category1Shape__00000002_.X = 398.999969
	__Category1Shape__00000002_.Y = 51.000000
	__Category1Shape__00000002_.Width = 240.000000
	__Category1Shape__00000002_.Height = 80.000000

	__Category1Shape__00000003_.Name = ``
	__Category1Shape__00000003_.X = 579.999969
	__Category1Shape__00000003_.Y = 378.999985
	__Category1Shape__00000003_.Width = 142.000000
	__Category1Shape__00000003_.Height = 80.000000

	__Category1Shape__00000004_.Name = ``
	__Category1Shape__00000004_.X = 737.999969
	__Category1Shape__00000004_.Y = 56.000000
	__Category1Shape__00000004_.Width = 194.000000
	__Category1Shape__00000004_.Height = 80.000000

	__Category1Shape__00000005_.Name = ``
	__Category1Shape__00000005_.X = 436.000000
	__Category1Shape__00000005_.Y = 529.000015
	__Category1Shape__00000005_.Width = 76.000031
	__Category1Shape__00000005_.Height = 80.000000

	__Desk__00000000_.Name = `Desk`

	__Diagram__00000000_.Name = `Default`
	__Diagram__00000000_.IsEditable = true
	__Diagram__00000000_.IsNodeExpanded = true
	__Diagram__00000000_.IsCategory1NodeExpanded = true
	__Diagram__00000000_.IsCategory2NodeExpanded = false
	__Diagram__00000000_.IsCategory3NodeExpanded = false
	__Diagram__00000000_.IsInfluenceCategoryNodeExpanded = true
	__Diagram__00000000_.IsCategory1Shown = true
	__Diagram__00000000_.IsCategory2Shown = false
	__Diagram__00000000_.IsCategory3Shown = false
	__Diagram__00000000_.IsInfluenceCategoryShown = true
	__Diagram__00000000_.XMargin = 0.000000
	__Diagram__00000000_.YMargin = 0.000000
	__Diagram__00000000_.Height = 2000.000000
	__Diagram__00000000_.Width = 2000.000000
	__Diagram__00000000_.NbPixPerCharacter = 8.000000
	__Diagram__00000000_.RedColorCode = `salmon `
	__Diagram__00000000_.BackgroundGreyColorCode = `white`
	__Diagram__00000000_.GrayColorCode = `gray`
	__Diagram__00000000_.Category1RectAnchorType = models.RECT_CENTER
	__Diagram__00000000_.Category1TextAnchorType = models.TEXT_ANCHOR_CENTER
	__Diagram__00000000_.Category1DominantBaselineType = ""
	__Diagram__00000000_.Category1FontSize = ``
	__Diagram__00000000_.Category1FontWeigth = ``
	__Diagram__00000000_.Category1FontFamily = ``
	__Diagram__00000000_.Category1LetterSpacing = ``
	__Diagram__00000000_.Category2TypeFontSize = ``
	__Diagram__00000000_.Category2TypeFontWeigth = ``
	__Diagram__00000000_.Category2TypeFontFamily = ``
	__Diagram__00000000_.Category2TypeLetterSpacing = ``
	__Diagram__00000000_.Category2TypeRectAnchorType = ""
	__Diagram__00000000_.Category2DominantBaselineType = ""
	__Diagram__00000000_.Category2StrokeWidth = 0.000000
	__Diagram__00000000_.Category3RectAnchorType = ""
	__Diagram__00000000_.Category3TextAnchorType = ""
	__Diagram__00000000_.Category3DominantBaselineType = ""
	__Diagram__00000000_.Category3FontSize = ``
	__Diagram__00000000_.Category3FontWeigth = ``
	__Diagram__00000000_.Category3FontFamily = ``
	__Diagram__00000000_.Category3LetterSpacing = ``
	__Diagram__00000000_.InfluenceStrokeWidth = 1.000000
	__Diagram__00000000_.InfluenceArrowSize = 10.000000
	__Diagram__00000000_.InfluenceArrowStartOffset = 0.000000
	__Diagram__00000000_.InfluenceArrowEndOffset = 0.000000
	__Diagram__00000000_.InfluenceCornerRadius = 19.000000
	__Diagram__00000000_.InfluenceFontSize = `28`
	__Diagram__00000000_.InfluenceFontWeigth = `bold`
	__Diagram__00000000_.InfluenceFontFamily = ``
	__Diagram__00000000_.InfluenceLetterSpacing = ``
	__Diagram__00000000_.InfluenceDashedLinePattern = ``

	__Influence__00000000_.Name = `Intensité de la Concurrence to Niveau des prix`
	__Influence__00000000_.IsHypothtical = false
	__Influence__00000000_.TextAtEndOfArrow = `-`

	__Influence__00000001_.Name = `Hauteur des barrières à l'entrée to Intensité de la Concurrence`
	__Influence__00000001_.IsHypothtical = false
	__Influence__00000001_.TextAtEndOfArrow = `-`

	__Influence__00000002_.Name = `Intensité de la Concurrence to Probabilité d'Oligopoles`
	__Influence__00000002_.IsHypothtical = false
	__Influence__00000002_.TextAtEndOfArrow = `-`

	__Influence__00000003_.Name = `Hauteur des barrières à l'entrée to Probabilité d'Oligopoles`
	__Influence__00000003_.IsHypothtical = false
	__Influence__00000003_.TextAtEndOfArrow = `+`

	__Influence__00000004_.Name = `Intensité de la politique concurrentielle to Hauteur des barrières à l'entrée`
	__Influence__00000004_.IsHypothtical = false
	__Influence__00000004_.TextAtEndOfArrow = `-`

	__Influence__00000005_.Name = `Intensité de la Concurrence to Probabilité de monopole`
	__Influence__00000005_.IsHypothtical = false
	__Influence__00000005_.TextAtEndOfArrow = `-`

	__InfluenceShape__00000000_.Name = `Intensité de la Concurrence to Niveau des prix`

	__InfluenceShape__00000001_.Name = `Hauteur des barrières à l'entrée to Intensité de la Concurrence`

	__InfluenceShape__00000002_.Name = `Intensité de la Concurrence to Probabilité d'Oligopoles`

	__InfluenceShape__00000003_.Name = `Hauteur des barrières à l'entrée to Probabilité d'Oligopoles`

	__InfluenceShape__00000004_.Name = `Intensité de la politique concurrentielle to Hauteur des barrières à l'entrée`

	__InfluenceShape__00000005_.Name = `Intensité de la Concurrence to Probabilité de monopole`

	// insertion point for setup of pointers
	__Category1Shape__00000000_.Category1 = __Category1__00000000_
	__Category1Shape__00000001_.Category1 = __Category1__00000001_
	__Category1Shape__00000002_.Category1 = __Category1__00000002_
	__Category1Shape__00000003_.Category1 = __Category1__00000003_
	__Category1Shape__00000004_.Category1 = __Category1__00000004_
	__Category1Shape__00000005_.Category1 = __Category1__00000005_
	__Desk__00000000_.SelectedDiagram = __Diagram__00000000_
	__Diagram__00000000_.Category1Shapes = append(__Diagram__00000000_.Category1Shapes, __Category1Shape__00000000_)
	__Diagram__00000000_.Category1Shapes = append(__Diagram__00000000_.Category1Shapes, __Category1Shape__00000001_)
	__Diagram__00000000_.Category1Shapes = append(__Diagram__00000000_.Category1Shapes, __Category1Shape__00000002_)
	__Diagram__00000000_.Category1Shapes = append(__Diagram__00000000_.Category1Shapes, __Category1Shape__00000003_)
	__Diagram__00000000_.Category1Shapes = append(__Diagram__00000000_.Category1Shapes, __Category1Shape__00000004_)
	__Diagram__00000000_.Category1Shapes = append(__Diagram__00000000_.Category1Shapes, __Category1Shape__00000005_)
	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __InfluenceShape__00000000_)
	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __InfluenceShape__00000001_)
	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __InfluenceShape__00000002_)
	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __InfluenceShape__00000003_)
	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __InfluenceShape__00000004_)
	__Diagram__00000000_.InfluenceShapes = append(__Diagram__00000000_.InfluenceShapes, __InfluenceShape__00000005_)
	__Influence__00000000_.SourceCategory1 = __Category1__00000000_
	__Influence__00000000_.SourceCategory2 = nil
	__Influence__00000000_.SourceCategory3 = nil
	__Influence__00000000_.TargetCategory1 = __Category1__00000001_
	__Influence__00000000_.TargetCategory2 = nil
	__Influence__00000000_.TargetCategory3 = nil
	__Influence__00000001_.SourceCategory1 = __Category1__00000002_
	__Influence__00000001_.SourceCategory2 = nil
	__Influence__00000001_.SourceCategory3 = nil
	__Influence__00000001_.TargetCategory1 = __Category1__00000000_
	__Influence__00000001_.TargetCategory2 = nil
	__Influence__00000001_.TargetCategory3 = nil
	__Influence__00000002_.SourceCategory1 = __Category1__00000000_
	__Influence__00000002_.SourceCategory2 = nil
	__Influence__00000002_.SourceCategory3 = nil
	__Influence__00000002_.TargetCategory1 = __Category1__00000003_
	__Influence__00000002_.TargetCategory2 = nil
	__Influence__00000002_.TargetCategory3 = nil
	__Influence__00000003_.SourceCategory1 = __Category1__00000002_
	__Influence__00000003_.SourceCategory2 = nil
	__Influence__00000003_.SourceCategory3 = nil
	__Influence__00000003_.TargetCategory1 = __Category1__00000003_
	__Influence__00000003_.TargetCategory2 = nil
	__Influence__00000003_.TargetCategory3 = nil
	__Influence__00000004_.SourceCategory1 = __Category1__00000004_
	__Influence__00000004_.SourceCategory2 = nil
	__Influence__00000004_.SourceCategory3 = nil
	__Influence__00000004_.TargetCategory1 = __Category1__00000002_
	__Influence__00000004_.TargetCategory2 = nil
	__Influence__00000004_.TargetCategory3 = nil
	__Influence__00000005_.SourceCategory1 = __Category1__00000000_
	__Influence__00000005_.SourceCategory2 = nil
	__Influence__00000005_.SourceCategory3 = nil
	__Influence__00000005_.TargetCategory1 = __Category1__00000005_
	__Influence__00000005_.TargetCategory2 = nil
	__Influence__00000005_.TargetCategory3 = nil
	__InfluenceShape__00000000_.Influence = __Influence__00000000_
	__InfluenceShape__00000001_.Influence = __Influence__00000001_
	__InfluenceShape__00000002_.Influence = __Influence__00000002_
	__InfluenceShape__00000003_.Influence = __Influence__00000003_
	__InfluenceShape__00000004_.Influence = __Influence__00000004_
	__InfluenceShape__00000005_.Influence = __Influence__00000005_
}
