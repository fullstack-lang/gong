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

	__ControlPoint__000000_ := (&models.ControlPoint{}).Stage(stage)

	__Layer__000000_default := (&models.Layer{}).Stage(stage)

	__Link__000000_start_end := (&models.Link{}).Stage(stage)

	__LinkAnchoredText__000000_Text := (&models.LinkAnchoredText{}).Stage(stage)

	__Rect__000000_Start := (&models.Rect{}).Stage(stage)
	__Rect__000001_End := (&models.Rect{}).Stage(stage)

	__SVG__000000_svg := (&models.SVG{}).Stage(stage)

	// Setup of values

	__ControlPoint__000000_.Name = ``
	__ControlPoint__000000_.X_Relative = 2.000000
	__ControlPoint__000000_.Y_Relative = 0.100000

	__Layer__000000_default.Name = `default`

	__Link__000000_start_end.Name = `start end`
	__Link__000000_start_end.Type = models.LINK_TYPE_LINE_WITH_CONTROL_POINTS
	__Link__000000_start_end.IsBezierCurve = true
	__Link__000000_start_end.StartAnchorType = models.ANCHOR_CENTER
	__Link__000000_start_end.EndAnchorType = models.ANCHOR_CENTER
	__Link__000000_start_end.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_start_end.StartRatio = 0.000000
	__Link__000000_start_end.EndRatio = 0.000000
	__Link__000000_start_end.CornerOffsetRatio = 0.000000
	__Link__000000_start_end.CornerRadius = 20.000000
	__Link__000000_start_end.HasEndArrow = true
	__Link__000000_start_end.EndArrowSize = 10.000000
	__Link__000000_start_end.EndArrowOffset = 0.000000
	__Link__000000_start_end.HasStartArrow = false
	__Link__000000_start_end.StartArrowSize = 0.000000
	__Link__000000_start_end.StartArrowOffset = 0.000000
	__Link__000000_start_end.Color = ``
	__Link__000000_start_end.FillOpacity = 0.000000
	__Link__000000_start_end.Stroke = `black`
	__Link__000000_start_end.StrokeOpacity = 1.000000
	__Link__000000_start_end.StrokeWidth = 1.000000
	__Link__000000_start_end.StrokeDashArray = ``
	__Link__000000_start_end.StrokeDashArrayWhenSelected = ``
	__Link__000000_start_end.Transform = ``
	__Link__000000_start_end.MouseX = 0.000000
	__Link__000000_start_end.MouseY = 0.000000

	__LinkAnchoredText__000000_Text.Name = `Text`
	__LinkAnchoredText__000000_Text.Content = `Text`
	__LinkAnchoredText__000000_Text.AutomaticLayout = true
	__LinkAnchoredText__000000_Text.LinkAnchorType = models.LINK_LEFT_OR_TOP
	__LinkAnchoredText__000000_Text.X_Offset = 0.000000
	__LinkAnchoredText__000000_Text.Y_Offset = 0.000000
	__LinkAnchoredText__000000_Text.FontWeight = ``
	__LinkAnchoredText__000000_Text.FontSize = ``
	__LinkAnchoredText__000000_Text.FontStyle = ``
	__LinkAnchoredText__000000_Text.LetterSpacing = ``
	__LinkAnchoredText__000000_Text.FontFamily = ``
	__LinkAnchoredText__000000_Text.Color = ``
	__LinkAnchoredText__000000_Text.FillOpacity = 0.000000
	__LinkAnchoredText__000000_Text.Stroke = `black`
	__LinkAnchoredText__000000_Text.StrokeOpacity = 1.000000
	__LinkAnchoredText__000000_Text.StrokeWidth = 1.000000
	__LinkAnchoredText__000000_Text.StrokeDashArray = ``
	__LinkAnchoredText__000000_Text.StrokeDashArrayWhenSelected = ``
	__LinkAnchoredText__000000_Text.Transform = ``

	__Rect__000000_Start.Name = `Start`
	__Rect__000000_Start.X = 0.000000
	__Rect__000000_Start.Y = 0.000000
	__Rect__000000_Start.Width = 100.000000
	__Rect__000000_Start.Height = 100.000000
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
	__Rect__000000_Start.ChangeColorWhenHovered = false
	__Rect__000000_Start.ColorWhenHovered = ``
	__Rect__000000_Start.OriginalColor = ``
	__Rect__000000_Start.FillOpacityWhenHovered = 0.000000
	__Rect__000000_Start.OriginalFillOpacity = 0.000000
	__Rect__000000_Start.HasToolTip = false
	__Rect__000000_Start.ToolTipText = ``
	__Rect__000000_Start.MouseX = 0.000000
	__Rect__000000_Start.MouseY = 0.000000

	__Rect__000001_End.Name = `End`
	__Rect__000001_End.X = 300.000000
	__Rect__000001_End.Y = 100.000000
	__Rect__000001_End.Width = 100.000000
	__Rect__000001_End.Height = 100.000000
	__Rect__000001_End.RX = 0.000000
	__Rect__000001_End.Color = ``
	__Rect__000001_End.FillOpacity = 0.000000
	__Rect__000001_End.Stroke = `black`
	__Rect__000001_End.StrokeOpacity = 1.000000
	__Rect__000001_End.StrokeWidth = -1.000000
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
	__Rect__000001_End.ChangeColorWhenHovered = false
	__Rect__000001_End.ColorWhenHovered = ``
	__Rect__000001_End.OriginalColor = ``
	__Rect__000001_End.FillOpacityWhenHovered = 0.000000
	__Rect__000001_End.OriginalFillOpacity = 0.000000
	__Rect__000001_End.HasToolTip = false
	__Rect__000001_End.ToolTipText = ``
	__Rect__000001_End.MouseX = 0.000000
	__Rect__000001_End.MouseY = 0.000000

	__SVG__000000_svg.Name = `svg`
	__SVG__000000_svg.IsEditable = false
	__SVG__000000_svg.IsSVGFrontEndFileGenerated = false
	__SVG__000000_svg.IsSVGBackEndFileGenerated = false
	__SVG__000000_svg.DefaultDirectoryForGeneratedImages = ``
	__SVG__000000_svg.IsControlBannerHidden = false

	// Setup of pointers
	// setup of ControlPoint instances pointers
	__ControlPoint__000000_.ClosestRect = __Rect__000000_Start
	// setup of Layer instances pointers
	__Layer__000000_default.Rects = append(__Layer__000000_default.Rects, __Rect__000000_Start)
	__Layer__000000_default.Rects = append(__Layer__000000_default.Rects, __Rect__000001_End)
	__Layer__000000_default.Links = append(__Layer__000000_default.Links, __Link__000000_start_end)
	// setup of Link instances pointers
	__Link__000000_start_end.Start = __Rect__000000_Start
	__Link__000000_start_end.End = __Rect__000001_End
	__Link__000000_start_end.TextAtArrowEnd = append(__Link__000000_start_end.TextAtArrowEnd, __LinkAnchoredText__000000_Text)
	__Link__000000_start_end.ControlPoints = append(__Link__000000_start_end.ControlPoints, __ControlPoint__000000_)
	// setup of LinkAnchoredText instances pointers
	// setup of Rect instances pointers
	// setup of SVG instances pointers
	__SVG__000000_svg.Layers = append(__SVG__000000_svg.Layers, __Layer__000000_default)
}

