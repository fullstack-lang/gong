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

	// Declaration of instances to stage

	__Layer__000000_Layer_1 := (&models.Layer{}).Stage(stage)

	__Link__000000_Start_to_End := (&models.Link{}).Stage(stage)

	__LinkAnchoredText__000000_LEFT_OR_TOP := (&models.LinkAnchoredText{}).Stage(stage)
	__LinkAnchoredText__000001_RIGHT_OR_BOTTOM := (&models.LinkAnchoredText{}).Stage(stage)

	__Rect__000000_Start := (&models.Rect{}).Stage(stage)
	__Rect__000001_End := (&models.Rect{}).Stage(stage)

	__SVG__000000_simple := (&models.SVG{}).Stage(stage)

	// Setup of values

	__Layer__000000_Layer_1.Name = `Layer 1`

	__Link__000000_Start_to_End.Name = `Start to End`
	__Link__000000_Start_to_End.Type = models.LINK_TYPE_FLOATING_ORTHOGONAL
	__Link__000000_Start_to_End.IsBezierCurve = false
	__Link__000000_Start_to_End.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Start_to_End.StartRatio = 0.000000
	__Link__000000_Start_to_End.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Start_to_End.EndRatio = 0.000000
	__Link__000000_Start_to_End.CornerOffsetRatio = 2.000000
	__Link__000000_Start_to_End.CornerRadius = 0.000000
	__Link__000000_Start_to_End.HasEndArrow = true
	__Link__000000_Start_to_End.EndArrowSize = 10.000000
	__Link__000000_Start_to_End.HasStartArrow = false
	__Link__000000_Start_to_End.StartArrowSize = 0.000000
	__Link__000000_Start_to_End.Color = ``
	__Link__000000_Start_to_End.FillOpacity = 0.000000
	__Link__000000_Start_to_End.Stroke = `black`
	__Link__000000_Start_to_End.StrokeOpacity = 1.000000
	__Link__000000_Start_to_End.StrokeWidth = 1.000000
	__Link__000000_Start_to_End.StrokeDashArray = ``
	__Link__000000_Start_to_End.StrokeDashArrayWhenSelected = ``
	__Link__000000_Start_to_End.Transform = ``

	__LinkAnchoredText__000000_LEFT_OR_TOP.Name = `LEFT_OR_TOP`
	__LinkAnchoredText__000000_LEFT_OR_TOP.Content = `LEFT_OR_TOP`
	__LinkAnchoredText__000000_LEFT_OR_TOP.AutomaticLayout = true
	__LinkAnchoredText__000000_LEFT_OR_TOP.LinkAnchorType = models.LINK_LEFT_OR_TOP
	__LinkAnchoredText__000000_LEFT_OR_TOP.X_Offset = 0.000000
	__LinkAnchoredText__000000_LEFT_OR_TOP.Y_Offset = 0.000000
	__LinkAnchoredText__000000_LEFT_OR_TOP.FontWeight = ``
	__LinkAnchoredText__000000_LEFT_OR_TOP.FontSize = ``
	__LinkAnchoredText__000000_LEFT_OR_TOP.FontStyle = ``
	__LinkAnchoredText__000000_LEFT_OR_TOP.LetterSpacing = ``
	__LinkAnchoredText__000000_LEFT_OR_TOP.Color = `black`
	__LinkAnchoredText__000000_LEFT_OR_TOP.FillOpacity = 1.000000
	__LinkAnchoredText__000000_LEFT_OR_TOP.Stroke = `black`
	__LinkAnchoredText__000000_LEFT_OR_TOP.StrokeOpacity = 1.000000
	__LinkAnchoredText__000000_LEFT_OR_TOP.StrokeWidth = 1.000000
	__LinkAnchoredText__000000_LEFT_OR_TOP.StrokeDashArray = ``
	__LinkAnchoredText__000000_LEFT_OR_TOP.StrokeDashArrayWhenSelected = ``
	__LinkAnchoredText__000000_LEFT_OR_TOP.Transform = ``

	__LinkAnchoredText__000001_RIGHT_OR_BOTTOM.Name = `RIGHT_OR_BOTTOM`
	__LinkAnchoredText__000001_RIGHT_OR_BOTTOM.Content = `RIGHT_OR_BOTTOM`
	__LinkAnchoredText__000001_RIGHT_OR_BOTTOM.AutomaticLayout = true
	__LinkAnchoredText__000001_RIGHT_OR_BOTTOM.LinkAnchorType = models.LINK_RIGHT_OR_BOTTOM
	__LinkAnchoredText__000001_RIGHT_OR_BOTTOM.X_Offset = 0.000000
	__LinkAnchoredText__000001_RIGHT_OR_BOTTOM.Y_Offset = 0.000000
	__LinkAnchoredText__000001_RIGHT_OR_BOTTOM.FontWeight = ``
	__LinkAnchoredText__000001_RIGHT_OR_BOTTOM.FontSize = ``
	__LinkAnchoredText__000001_RIGHT_OR_BOTTOM.FontStyle = ``
	__LinkAnchoredText__000001_RIGHT_OR_BOTTOM.LetterSpacing = ``
	__LinkAnchoredText__000001_RIGHT_OR_BOTTOM.Color = `black`
	__LinkAnchoredText__000001_RIGHT_OR_BOTTOM.FillOpacity = 1.000000
	__LinkAnchoredText__000001_RIGHT_OR_BOTTOM.Stroke = `black`
	__LinkAnchoredText__000001_RIGHT_OR_BOTTOM.StrokeOpacity = 1.000000
	__LinkAnchoredText__000001_RIGHT_OR_BOTTOM.StrokeWidth = 1.000000
	__LinkAnchoredText__000001_RIGHT_OR_BOTTOM.StrokeDashArray = ``
	__LinkAnchoredText__000001_RIGHT_OR_BOTTOM.StrokeDashArrayWhenSelected = ``
	__LinkAnchoredText__000001_RIGHT_OR_BOTTOM.Transform = ``

	__Rect__000000_Start.Name = `Start`
	__Rect__000000_Start.X = 100.000000
	__Rect__000000_Start.Y = 100.000000
	__Rect__000000_Start.Width = 200.000000
	__Rect__000000_Start.Height = 200.000000
	__Rect__000000_Start.RX = 0.000000
	__Rect__000000_Start.Color = ``
	__Rect__000000_Start.FillOpacity = 0.000000
	__Rect__000000_Start.Stroke = `black`
	__Rect__000000_Start.StrokeOpacity = 1.000000
	__Rect__000000_Start.StrokeWidth = 1.000000
	__Rect__000000_Start.StrokeDashArray = ``
	__Rect__000000_Start.StrokeDashArrayWhenSelected = ``
	__Rect__000000_Start.Transform = ``
	__Rect__000000_Start.IsSelectable = false
	__Rect__000000_Start.IsSelected = false
	__Rect__000000_Start.CanHaveLeftHandle = false
	__Rect__000000_Start.HasLeftHandle = false
	__Rect__000000_Start.CanHaveRightHandle = false
	__Rect__000000_Start.HasRightHandle = false
	__Rect__000000_Start.CanHaveTopHandle = false
	__Rect__000000_Start.HasTopHandle = false
	__Rect__000000_Start.IsScalingProportionally = false
	__Rect__000000_Start.CanHaveBottomHandle = false
	__Rect__000000_Start.HasBottomHandle = false
	__Rect__000000_Start.CanMoveHorizontaly = false
	__Rect__000000_Start.CanMoveVerticaly = false

	__Rect__000001_End.Name = `End`
	__Rect__000001_End.X = 700.000000
	__Rect__000001_End.Y = 100.000000
	__Rect__000001_End.Width = 200.000000
	__Rect__000001_End.Height = 200.000000
	__Rect__000001_End.RX = 0.000000
	__Rect__000001_End.Color = ``
	__Rect__000001_End.FillOpacity = 0.000000
	__Rect__000001_End.Stroke = `black`
	__Rect__000001_End.StrokeOpacity = 1.000000
	__Rect__000001_End.StrokeWidth = 1.000000
	__Rect__000001_End.StrokeDashArray = ``
	__Rect__000001_End.StrokeDashArrayWhenSelected = ``
	__Rect__000001_End.Transform = ``
	__Rect__000001_End.IsSelectable = false
	__Rect__000001_End.IsSelected = false
	__Rect__000001_End.CanHaveLeftHandle = false
	__Rect__000001_End.HasLeftHandle = false
	__Rect__000001_End.CanHaveRightHandle = false
	__Rect__000001_End.HasRightHandle = false
	__Rect__000001_End.CanHaveTopHandle = false
	__Rect__000001_End.HasTopHandle = false
	__Rect__000001_End.IsScalingProportionally = false
	__Rect__000001_End.CanHaveBottomHandle = false
	__Rect__000001_End.HasBottomHandle = false
	__Rect__000001_End.CanMoveHorizontaly = false
	__Rect__000001_End.CanMoveVerticaly = false

	__SVG__000000_simple.Name = `simple`
	__SVG__000000_simple.IsEditable = false
	__SVG__000000_simple.IsSVGFrontEndFileGenerated = false
	__SVG__000000_simple.IsSVGBackEndFileGenerated = false
	__SVG__000000_simple.DefaultDirectoryForGeneratedImages = `../../diagrams/images`

	// Setup of pointers
	// setup of Layer instances pointers
	__Layer__000000_Layer_1.Rects = append(__Layer__000000_Layer_1.Rects, __Rect__000000_Start)
	__Layer__000000_Layer_1.Rects = append(__Layer__000000_Layer_1.Rects, __Rect__000001_End)
	__Layer__000000_Layer_1.Links = append(__Layer__000000_Layer_1.Links, __Link__000000_Start_to_End)
	// setup of Link instances pointers
	__Link__000000_Start_to_End.Start = __Rect__000000_Start
	__Link__000000_Start_to_End.End = __Rect__000001_End
	__Link__000000_Start_to_End.TextAtArrowStart = append(__Link__000000_Start_to_End.TextAtArrowStart, __LinkAnchoredText__000000_LEFT_OR_TOP)
	__Link__000000_Start_to_End.TextAtArrowStart = append(__Link__000000_Start_to_End.TextAtArrowStart, __LinkAnchoredText__000001_RIGHT_OR_BOTTOM)
	__Link__000000_Start_to_End.TextAtArrowEnd = append(__Link__000000_Start_to_End.TextAtArrowEnd, __LinkAnchoredText__000000_LEFT_OR_TOP)
	__Link__000000_Start_to_End.TextAtArrowEnd = append(__Link__000000_Start_to_End.TextAtArrowEnd, __LinkAnchoredText__000001_RIGHT_OR_BOTTOM)
	// setup of LinkAnchoredText instances pointers
	// setup of Rect instances pointers
	// setup of SVG instances pointers
	__SVG__000000_simple.Layers = append(__SVG__000000_simple.Layers, __Layer__000000_Layer_1)
}
