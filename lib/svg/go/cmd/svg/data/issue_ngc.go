package main

import (
	"time"

	"github.com/fullstack-lang/gong/lib/svg/go/models"
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

	const __write__local_time = "2025-10-07 04:12:27.736422 CEST"
	const __write__utc_time__ = "2025-10-07 02:12:27.736422 UTC"

	const __commitId__ = "0000000034"

	// Declaration of instances to stage

	__Layer__000000_ngc := (&models.Layer{}).Stage(stage)

	__Link__000000_ngc := (&models.Link{}).Stage(stage)

	__LinkAnchoredText__000000_LAT_1 := (&models.LinkAnchoredText{}).Stage(stage)
	__LinkAnchoredText__000001_LAT_2 := (&models.LinkAnchoredText{}).Stage(stage)

	__Rect__000000_ngc_1 := (&models.Rect{}).Stage(stage)
	__Rect__000001_ngc_2 := (&models.Rect{}).Stage(stage)

	__SVG__000000_ngc := (&models.SVG{}).Stage(stage)

	// Setup of values

	__Layer__000000_ngc.Name = `ngc`

	__Link__000000_ngc.Name = `ngc`
	__Link__000000_ngc.Type = models.LINK_TYPE_FLOATING_ORTHOGONAL
	__Link__000000_ngc.IsBezierCurve = false
	__Link__000000_ngc.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_ngc.StartRatio = 0.102233
	__Link__000000_ngc.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_ngc.EndRatio = 0.735337
	__Link__000000_ngc.CornerOffsetRatio = 1.805147
	__Link__000000_ngc.CornerRadius = 0.000000
	__Link__000000_ngc.HasEndArrow = true
	__Link__000000_ngc.EndArrowSize = 20.000000
	__Link__000000_ngc.HasStartArrow = false
	__Link__000000_ngc.StartArrowSize = 0.000000
	__Link__000000_ngc.Color = ``
	__Link__000000_ngc.FillOpacity = 0.000000
	__Link__000000_ngc.Stroke = `gray`
	__Link__000000_ngc.StrokeOpacity = 1.000000
	__Link__000000_ngc.StrokeWidth = 2.000000
	__Link__000000_ngc.StrokeDashArray = ``
	__Link__000000_ngc.StrokeDashArrayWhenSelected = ``
	__Link__000000_ngc.Transform = ``

	__LinkAnchoredText__000000_LAT_1.Name = `LAT 1`
	__LinkAnchoredText__000000_LAT_1.Content = `LAT 1`
	__LinkAnchoredText__000000_LAT_1.AutomaticLayout = true
	__LinkAnchoredText__000000_LAT_1.LinkAnchorType = models.LINK_RIGHT_OR_BOTTOM
	__LinkAnchoredText__000000_LAT_1.X_Offset = 0.000000
	__LinkAnchoredText__000000_LAT_1.Y_Offset = 0.000000
	__LinkAnchoredText__000000_LAT_1.FontWeight = ``
	__LinkAnchoredText__000000_LAT_1.FontSize = ``
	__LinkAnchoredText__000000_LAT_1.FontStyle = ``
	__LinkAnchoredText__000000_LAT_1.LetterSpacing = ``
	__LinkAnchoredText__000000_LAT_1.Color = `blue`
	__LinkAnchoredText__000000_LAT_1.FillOpacity = 1.000000
	__LinkAnchoredText__000000_LAT_1.Stroke = `blue`
	__LinkAnchoredText__000000_LAT_1.StrokeOpacity = 1.000000
	__LinkAnchoredText__000000_LAT_1.StrokeWidth = 1.000000
	__LinkAnchoredText__000000_LAT_1.StrokeDashArray = ``
	__LinkAnchoredText__000000_LAT_1.StrokeDashArrayWhenSelected = ``
	__LinkAnchoredText__000000_LAT_1.Transform = ``

	__LinkAnchoredText__000001_LAT_2.Name = `LAT 2`
	__LinkAnchoredText__000001_LAT_2.Content = `LAT 2`
	__LinkAnchoredText__000001_LAT_2.AutomaticLayout = true
	__LinkAnchoredText__000001_LAT_2.LinkAnchorType = models.LINK_LEFT_OR_TOP
	__LinkAnchoredText__000001_LAT_2.X_Offset = -40.000000
	__LinkAnchoredText__000001_LAT_2.Y_Offset = 0.000000
	__LinkAnchoredText__000001_LAT_2.FontWeight = ``
	__LinkAnchoredText__000001_LAT_2.FontSize = ``
	__LinkAnchoredText__000001_LAT_2.FontStyle = ``
	__LinkAnchoredText__000001_LAT_2.LetterSpacing = ``
	__LinkAnchoredText__000001_LAT_2.Color = `red`
	__LinkAnchoredText__000001_LAT_2.FillOpacity = 1.000000
	__LinkAnchoredText__000001_LAT_2.Stroke = `red`
	__LinkAnchoredText__000001_LAT_2.StrokeOpacity = 1.000000
	__LinkAnchoredText__000001_LAT_2.StrokeWidth = 1.000000
	__LinkAnchoredText__000001_LAT_2.StrokeDashArray = ``
	__LinkAnchoredText__000001_LAT_2.StrokeDashArrayWhenSelected = ``
	__LinkAnchoredText__000001_LAT_2.Transform = ``

	__Rect__000000_ngc_1.Name = `ngc 1`
	__Rect__000000_ngc_1.X = 334.000000
	__Rect__000000_ngc_1.Y = 100.000000
	__Rect__000000_ngc_1.Width = 272.000000
	__Rect__000000_ngc_1.Height = 137.000000
	__Rect__000000_ngc_1.RX = 5.000000
	__Rect__000000_ngc_1.Color = ``
	__Rect__000000_ngc_1.FillOpacity = 0.000000
	__Rect__000000_ngc_1.Stroke = `gray`
	__Rect__000000_ngc_1.StrokeOpacity = 1.000000
	__Rect__000000_ngc_1.StrokeWidth = 2.000000
	__Rect__000000_ngc_1.StrokeDashArray = ``
	__Rect__000000_ngc_1.StrokeDashArrayWhenSelected = ``
	__Rect__000000_ngc_1.Transform = ``
	__Rect__000000_ngc_1.IsSelectable = true
	__Rect__000000_ngc_1.IsSelected = false
	__Rect__000000_ngc_1.CanHaveLeftHandle = true
	__Rect__000000_ngc_1.HasLeftHandle = false
	__Rect__000000_ngc_1.CanHaveRightHandle = true
	__Rect__000000_ngc_1.HasRightHandle = false
	__Rect__000000_ngc_1.CanHaveTopHandle = true
	__Rect__000000_ngc_1.HasTopHandle = false
	__Rect__000000_ngc_1.IsScalingProportionally = false
	__Rect__000000_ngc_1.CanHaveBottomHandle = true
	__Rect__000000_ngc_1.HasBottomHandle = false
	__Rect__000000_ngc_1.CanMoveHorizontaly = true
	__Rect__000000_ngc_1.CanMoveVerticaly = false
	__Rect__000000_ngc_1.ChangeColorWhenHovered = false
	__Rect__000000_ngc_1.ColorWhenHovered = ``
	__Rect__000000_ngc_1.OriginalColor = ``
	__Rect__000000_ngc_1.FillOpacityWhenHovered = 0.000000
	__Rect__000000_ngc_1.OriginalFillOpacity = 0.000000
	__Rect__000000_ngc_1.HasToolTip = false
	__Rect__000000_ngc_1.ToolTipText = ``

	__Rect__000001_ngc_2.Name = `ngc 2`
	__Rect__000001_ngc_2.X = 979.000000
	__Rect__000001_ngc_2.Y = 82.000000
	__Rect__000001_ngc_2.Width = 202.000000
	__Rect__000001_ngc_2.Height = 136.000000
	__Rect__000001_ngc_2.RX = 10.000000
	__Rect__000001_ngc_2.Color = ``
	__Rect__000001_ngc_2.FillOpacity = 0.000000
	__Rect__000001_ngc_2.Stroke = `black`
	__Rect__000001_ngc_2.StrokeOpacity = 1.000000
	__Rect__000001_ngc_2.StrokeWidth = 5.000000
	__Rect__000001_ngc_2.StrokeDashArray = ``
	__Rect__000001_ngc_2.StrokeDashArrayWhenSelected = ``
	__Rect__000001_ngc_2.Transform = ``
	__Rect__000001_ngc_2.IsSelectable = true
	__Rect__000001_ngc_2.IsSelected = false
	__Rect__000001_ngc_2.CanHaveLeftHandle = true
	__Rect__000001_ngc_2.HasLeftHandle = false
	__Rect__000001_ngc_2.CanHaveRightHandle = true
	__Rect__000001_ngc_2.HasRightHandle = false
	__Rect__000001_ngc_2.CanHaveTopHandle = true
	__Rect__000001_ngc_2.HasTopHandle = false
	__Rect__000001_ngc_2.IsScalingProportionally = false
	__Rect__000001_ngc_2.CanHaveBottomHandle = true
	__Rect__000001_ngc_2.HasBottomHandle = false
	__Rect__000001_ngc_2.CanMoveHorizontaly = true
	__Rect__000001_ngc_2.CanMoveVerticaly = true
	__Rect__000001_ngc_2.ChangeColorWhenHovered = false
	__Rect__000001_ngc_2.ColorWhenHovered = ``
	__Rect__000001_ngc_2.OriginalColor = ``
	__Rect__000001_ngc_2.FillOpacityWhenHovered = 0.000000
	__Rect__000001_ngc_2.OriginalFillOpacity = 0.000000
	__Rect__000001_ngc_2.HasToolTip = false
	__Rect__000001_ngc_2.ToolTipText = ``

	__SVG__000000_ngc.Name = `ngc`
	__SVG__000000_ngc.IsEditable = true
	__SVG__000000_ngc.IsSVGFrontEndFileGenerated = false
	__SVG__000000_ngc.IsSVGBackEndFileGenerated = false
	__SVG__000000_ngc.DefaultDirectoryForGeneratedImages = ``
	__SVG__000000_ngc.IsControlBannerHidden = false

	// Setup of pointers
	// setup of Layer instances pointers
	__Layer__000000_ngc.Rects = append(__Layer__000000_ngc.Rects, __Rect__000000_ngc_1)
	__Layer__000000_ngc.Rects = append(__Layer__000000_ngc.Rects, __Rect__000001_ngc_2)
	__Layer__000000_ngc.Links = append(__Layer__000000_ngc.Links, __Link__000000_ngc)
	// setup of Link instances pointers
	__Link__000000_ngc.Start = __Rect__000000_ngc_1
	__Link__000000_ngc.End = __Rect__000001_ngc_2
	__Link__000000_ngc.TextAtArrowEnd = append(__Link__000000_ngc.TextAtArrowEnd, __LinkAnchoredText__000000_LAT_1)
	__Link__000000_ngc.TextAtArrowEnd = append(__Link__000000_ngc.TextAtArrowEnd, __LinkAnchoredText__000001_LAT_2)
	// setup of LinkAnchoredText instances pointers
	// setup of Rect instances pointers
	// setup of SVG instances pointers
	__SVG__000000_ngc.Layers = append(__SVG__000000_ngc.Layers, __Layer__000000_ngc)
}

