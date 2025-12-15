package main

import (
	"time"

	"github.com/fullstack-lang/gong/dsme/cld/go/models"
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

	// Declaration of instances to stage

	__Category1__000000_Concurrence := (&models.Category1{}).Stage(stage)

	__Category1Shape__000000_ := (&models.Category1Shape{}).Stage(stage)

	__Desk__000000_Desk := (&models.Desk{}).Stage(stage)

	__Diagram__000000_Default := (&models.Diagram{}).Stage(stage)

	// Setup of values

	__Category1__000000_Concurrence.Name = `Concurrence`

	__Category1Shape__000000_.Name = ``
	__Category1Shape__000000_.X = 77.000000
	__Category1Shape__000000_.Y = 50.000000
	__Category1Shape__000000_.Width = 240.000000
	__Category1Shape__000000_.Height = 80.000000

	__Desk__000000_Desk.Name = `Desk`

	__Diagram__000000_Default.Name = `Default`
	__Diagram__000000_Default.IsEditable = true
	__Diagram__000000_Default.IsNodeExpanded = true
	__Diagram__000000_Default.IsCategory1NodeExpanded = true
	__Diagram__000000_Default.IsCategory2NodeExpanded = false
	__Diagram__000000_Default.IsCategory3NodeExpanded = false
	__Diagram__000000_Default.IsInfluenceCategoryNodeExpanded = false
	__Diagram__000000_Default.IsCategory1Shown = false
	__Diagram__000000_Default.IsCategory2Shown = false
	__Diagram__000000_Default.IsCategory3Shown = false
	__Diagram__000000_Default.IsInfluenceCategoryShown = false
	__Diagram__000000_Default.XMargin = 0.000000
	__Diagram__000000_Default.YMargin = 0.000000
	__Diagram__000000_Default.Height = 700.000000
	__Diagram__000000_Default.Width = 500.000000
	__Diagram__000000_Default.RedColorCode = ``
	__Diagram__000000_Default.BackgroundGreyColorCode = `lightgray`
	__Diagram__000000_Default.GrayColorCode = `gray`
	__Diagram__000000_Default.Category1RectAnchorType = models.RECT_RIGHT
	__Diagram__000000_Default.Category1TextAnchorType = models.TEXT_ANCHOR_END
	__Diagram__000000_Default.Category1DominantBaselineType = models.DominantBaselineAuto
	__Diagram__000000_Default.Category1FontSize = `16`
	__Diagram__000000_Default.Category1FontWeigth = `normal`
	__Diagram__000000_Default.Category1FontFamily = ``
	__Diagram__000000_Default.Category1LetterSpacing = ``
	__Diagram__000000_Default.Category2TypeFontSize = ``
	__Diagram__000000_Default.Category2TypeFontWeigth = ``
	__Diagram__000000_Default.Category2TypeFontFamily = ``
	__Diagram__000000_Default.Category2TypeLetterSpacing = ``
	__Diagram__000000_Default.Category2StrokeWidth = 0.000000
	__Diagram__000000_Default.Category3FontSize = ``
	__Diagram__000000_Default.Category3FontWeigth = ``
	__Diagram__000000_Default.Category3FontFamily = ``
	__Diagram__000000_Default.Category3LetterSpacing = ``
	__Diagram__000000_Default.InfluenceArrowSize = 0.000000
	__Diagram__000000_Default.InfluenceArrowStartOffset = 0.000000
	__Diagram__000000_Default.InfluenceArrowEndOffset = 0.000000
	__Diagram__000000_Default.InfluenceCornerRadius = 0.000000
	__Diagram__000000_Default.InfluenceDashedLinePattern = ``

	// Setup of pointers
	// setup of Category1 instances pointers
	// setup of Category1Shape instances pointers
	__Category1Shape__000000_.Category1 = __Category1__000000_Concurrence
	// setup of Desk instances pointers
	__Desk__000000_Desk.SelectedDiagram = __Diagram__000000_Default
	// setup of Diagram instances pointers
	__Diagram__000000_Default.Category1Shapes = append(__Diagram__000000_Default.Category1Shapes, __Category1Shape__000000_)
}

