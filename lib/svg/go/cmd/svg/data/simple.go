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

	__RectAnchoredRect__000000_Top := (&models.RectAnchoredRect{}).Stage(stage)
	__RectAnchoredRect__000001_TopLeft := (&models.RectAnchoredRect{}).Stage(stage)
	__RectAnchoredRect__000002_TopRight := (&models.RectAnchoredRect{}).Stage(stage)
	__RectAnchoredRect__000003_Test := (&models.RectAnchoredRect{}).Stage(stage)
	__RectAnchoredRect__000004_BottomLeft := (&models.RectAnchoredRect{}).Stage(stage)
	__RectAnchoredRect__000005_BottomLeftLeft := (&models.RectAnchoredRect{}).Stage(stage)
	__RectAnchoredRect__000006_BottomBottomLeft := (&models.RectAnchoredRect{}).Stage(stage)
	__RectAnchoredRect__000007_BottomRight := (&models.RectAnchoredRect{}).Stage(stage)
	__RectAnchoredRect__000008_BottomInsideRight := (&models.RectAnchoredRect{}).Stage(stage)
	__RectAnchoredRect__000009_Left := (&models.RectAnchoredRect{}).Stage(stage)
	__RectAnchoredRect__000010_Right := (&models.RectAnchoredRect{}).Stage(stage)
	__RectAnchoredRect__000011_Center := (&models.RectAnchoredRect{}).Stage(stage)

	__RectAnchoredText__000000_Oriented := (&models.RectAnchoredText{}).Stage(stage)

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
	__Link__000000_Start_to_End.StrokeOpacity = 0.500000
	__Link__000000_Start_to_End.StrokeWidth = 2.000000
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
	__LinkAnchoredText__000000_LEFT_OR_TOP.StrokeOpacity = 0.500000
	__LinkAnchoredText__000000_LEFT_OR_TOP.StrokeWidth = 2.000000
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
	__LinkAnchoredText__000001_RIGHT_OR_BOTTOM.StrokeOpacity = 0.500000
	__LinkAnchoredText__000001_RIGHT_OR_BOTTOM.StrokeWidth = 2.000000
	__LinkAnchoredText__000001_RIGHT_OR_BOTTOM.StrokeDashArray = ``
	__LinkAnchoredText__000001_RIGHT_OR_BOTTOM.StrokeDashArrayWhenSelected = ``
	__LinkAnchoredText__000001_RIGHT_OR_BOTTOM.Transform = ``

	__Rect__000000_Start.Name = `Start`
	__Rect__000000_Start.X = 100.000000
	__Rect__000000_Start.Y = 100.000000
	__Rect__000000_Start.Width = 200.000000
	__Rect__000000_Start.Height = 200.000000
	__Rect__000000_Start.RX = 0.000000
	__Rect__000000_Start.Color = `blue`
	__Rect__000000_Start.FillOpacity = 0.000000
	__Rect__000000_Start.Stroke = `black`
	__Rect__000000_Start.StrokeOpacity = 0.500000
	__Rect__000000_Start.StrokeWidth = 2.000000
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
	__Rect__000000_Start.ChangeColorWhenHovered = true
	__Rect__000000_Start.ColorWhenHovered = `red`
	__Rect__000000_Start.OriginalColor = ``
	__Rect__000000_Start.FillOpacityWhenHovered = 0.500000
	__Rect__000000_Start.OriginalFillOpacity = 0.000000

	__Rect__000001_End.Name = `End`
	__Rect__000001_End.X = 700.000000
	__Rect__000001_End.Y = 100.000000
	__Rect__000001_End.Width = 300.000000
	__Rect__000001_End.Height = 300.000000
	__Rect__000001_End.RX = 0.000000
	__Rect__000001_End.Color = ``
	__Rect__000001_End.FillOpacity = 0.000000
	__Rect__000001_End.Stroke = `black`
	__Rect__000001_End.StrokeOpacity = 0.500000
	__Rect__000001_End.StrokeWidth = 2.000000
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

	__RectAnchoredRect__000000_Top.Name = `Top`
	__RectAnchoredRect__000000_Top.X = 0.000000
	__RectAnchoredRect__000000_Top.Y = 0.000000
	__RectAnchoredRect__000000_Top.Width = 100.000000
	__RectAnchoredRect__000000_Top.Height = 70.000000
	__RectAnchoredRect__000000_Top.RX = 5.000000
	__RectAnchoredRect__000000_Top.X_Offset = 0.000000
	__RectAnchoredRect__000000_Top.Y_Offset = 0.000000
	__RectAnchoredRect__000000_Top.RectAnchorType = models.RECT_TOP
	__RectAnchoredRect__000000_Top.WidthFollowRect = false
	__RectAnchoredRect__000000_Top.HeightFollowRect = false
	__RectAnchoredRect__000000_Top.Color = ``
	__RectAnchoredRect__000000_Top.FillOpacity = 0.000000
	__RectAnchoredRect__000000_Top.Stroke = `blue`
	__RectAnchoredRect__000000_Top.StrokeOpacity = 0.500000
	__RectAnchoredRect__000000_Top.StrokeWidth = 4.000000
	__RectAnchoredRect__000000_Top.StrokeDashArray = ``
	__RectAnchoredRect__000000_Top.StrokeDashArrayWhenSelected = ``
	__RectAnchoredRect__000000_Top.Transform = ``

	__RectAnchoredRect__000001_TopLeft.Name = `TopLeft`
	__RectAnchoredRect__000001_TopLeft.X = 0.000000
	__RectAnchoredRect__000001_TopLeft.Y = 0.000000
	__RectAnchoredRect__000001_TopLeft.Width = 100.000000
	__RectAnchoredRect__000001_TopLeft.Height = 70.000000
	__RectAnchoredRect__000001_TopLeft.RX = 5.000000
	__RectAnchoredRect__000001_TopLeft.X_Offset = 0.000000
	__RectAnchoredRect__000001_TopLeft.Y_Offset = 0.000000
	__RectAnchoredRect__000001_TopLeft.RectAnchorType = models.RECT_TOP_LEFT
	__RectAnchoredRect__000001_TopLeft.WidthFollowRect = false
	__RectAnchoredRect__000001_TopLeft.HeightFollowRect = false
	__RectAnchoredRect__000001_TopLeft.Color = ``
	__RectAnchoredRect__000001_TopLeft.FillOpacity = 0.000000
	__RectAnchoredRect__000001_TopLeft.Stroke = `red`
	__RectAnchoredRect__000001_TopLeft.StrokeOpacity = 0.500000
	__RectAnchoredRect__000001_TopLeft.StrokeWidth = 2.000000
	__RectAnchoredRect__000001_TopLeft.StrokeDashArray = ``
	__RectAnchoredRect__000001_TopLeft.StrokeDashArrayWhenSelected = ``
	__RectAnchoredRect__000001_TopLeft.Transform = ``

	__RectAnchoredRect__000002_TopRight.Name = `TopRight`
	__RectAnchoredRect__000002_TopRight.X = 0.000000
	__RectAnchoredRect__000002_TopRight.Y = 0.000000
	__RectAnchoredRect__000002_TopRight.Width = 100.000000
	__RectAnchoredRect__000002_TopRight.Height = 70.000000
	__RectAnchoredRect__000002_TopRight.RX = 5.000000
	__RectAnchoredRect__000002_TopRight.X_Offset = 0.000000
	__RectAnchoredRect__000002_TopRight.Y_Offset = 0.000000
	__RectAnchoredRect__000002_TopRight.RectAnchorType = models.RECT_TOP_RIGHT
	__RectAnchoredRect__000002_TopRight.WidthFollowRect = false
	__RectAnchoredRect__000002_TopRight.HeightFollowRect = false
	__RectAnchoredRect__000002_TopRight.Color = ``
	__RectAnchoredRect__000002_TopRight.FillOpacity = 0.000000
	__RectAnchoredRect__000002_TopRight.Stroke = `yellow`
	__RectAnchoredRect__000002_TopRight.StrokeOpacity = 0.500000
	__RectAnchoredRect__000002_TopRight.StrokeWidth = 10.000000
	__RectAnchoredRect__000002_TopRight.StrokeDashArray = ``
	__RectAnchoredRect__000002_TopRight.StrokeDashArrayWhenSelected = ``
	__RectAnchoredRect__000002_TopRight.Transform = ``

	__RectAnchoredRect__000003_Test.Name = `Test`
	__RectAnchoredRect__000003_Test.X = 0.000000
	__RectAnchoredRect__000003_Test.Y = 0.000000
	__RectAnchoredRect__000003_Test.Width = 100.000000
	__RectAnchoredRect__000003_Test.Height = 70.000000
	__RectAnchoredRect__000003_Test.RX = 5.000000
	__RectAnchoredRect__000003_Test.X_Offset = 0.000000
	__RectAnchoredRect__000003_Test.Y_Offset = 0.000000
	__RectAnchoredRect__000003_Test.RectAnchorType = models.RECT_BOTTOM
	__RectAnchoredRect__000003_Test.WidthFollowRect = false
	__RectAnchoredRect__000003_Test.HeightFollowRect = false
	__RectAnchoredRect__000003_Test.Color = ``
	__RectAnchoredRect__000003_Test.FillOpacity = 0.000000
	__RectAnchoredRect__000003_Test.Stroke = `black`
	__RectAnchoredRect__000003_Test.StrokeOpacity = 0.500000
	__RectAnchoredRect__000003_Test.StrokeWidth = 2.000000
	__RectAnchoredRect__000003_Test.StrokeDashArray = ``
	__RectAnchoredRect__000003_Test.StrokeDashArrayWhenSelected = ``
	__RectAnchoredRect__000003_Test.Transform = ``

	__RectAnchoredRect__000004_BottomLeft.Name = `BottomLeft`
	__RectAnchoredRect__000004_BottomLeft.X = 0.000000
	__RectAnchoredRect__000004_BottomLeft.Y = 0.000000
	__RectAnchoredRect__000004_BottomLeft.Width = 100.000000
	__RectAnchoredRect__000004_BottomLeft.Height = 70.000000
	__RectAnchoredRect__000004_BottomLeft.RX = 5.000000
	__RectAnchoredRect__000004_BottomLeft.X_Offset = 0.000000
	__RectAnchoredRect__000004_BottomLeft.Y_Offset = 0.000000
	__RectAnchoredRect__000004_BottomLeft.RectAnchorType = models.RECT_BOTTOM_LEFT
	__RectAnchoredRect__000004_BottomLeft.WidthFollowRect = false
	__RectAnchoredRect__000004_BottomLeft.HeightFollowRect = false
	__RectAnchoredRect__000004_BottomLeft.Color = ``
	__RectAnchoredRect__000004_BottomLeft.FillOpacity = 0.000000
	__RectAnchoredRect__000004_BottomLeft.Stroke = `green`
	__RectAnchoredRect__000004_BottomLeft.StrokeOpacity = 0.200000
	__RectAnchoredRect__000004_BottomLeft.StrokeWidth = 8.000000
	__RectAnchoredRect__000004_BottomLeft.StrokeDashArray = ``
	__RectAnchoredRect__000004_BottomLeft.StrokeDashArrayWhenSelected = ``
	__RectAnchoredRect__000004_BottomLeft.Transform = ``

	__RectAnchoredRect__000005_BottomLeftLeft.Name = `BottomLeftLeft`
	__RectAnchoredRect__000005_BottomLeftLeft.X = 0.000000
	__RectAnchoredRect__000005_BottomLeftLeft.Y = 0.000000
	__RectAnchoredRect__000005_BottomLeftLeft.Width = 100.000000
	__RectAnchoredRect__000005_BottomLeftLeft.Height = 70.000000
	__RectAnchoredRect__000005_BottomLeftLeft.RX = 5.000000
	__RectAnchoredRect__000005_BottomLeftLeft.X_Offset = 0.000000
	__RectAnchoredRect__000005_BottomLeftLeft.Y_Offset = 0.000000
	__RectAnchoredRect__000005_BottomLeftLeft.RectAnchorType = models.RECT_BOTTOM_LEFT_LEFT
	__RectAnchoredRect__000005_BottomLeftLeft.WidthFollowRect = false
	__RectAnchoredRect__000005_BottomLeftLeft.HeightFollowRect = false
	__RectAnchoredRect__000005_BottomLeftLeft.Color = ``
	__RectAnchoredRect__000005_BottomLeftLeft.FillOpacity = 0.000000
	__RectAnchoredRect__000005_BottomLeftLeft.Stroke = `violet`
	__RectAnchoredRect__000005_BottomLeftLeft.StrokeOpacity = 0.500000
	__RectAnchoredRect__000005_BottomLeftLeft.StrokeWidth = 5.000000
	__RectAnchoredRect__000005_BottomLeftLeft.StrokeDashArray = ``
	__RectAnchoredRect__000005_BottomLeftLeft.StrokeDashArrayWhenSelected = ``
	__RectAnchoredRect__000005_BottomLeftLeft.Transform = ``

	__RectAnchoredRect__000006_BottomBottomLeft.Name = `BottomBottomLeft`
	__RectAnchoredRect__000006_BottomBottomLeft.X = 0.000000
	__RectAnchoredRect__000006_BottomBottomLeft.Y = 0.000000
	__RectAnchoredRect__000006_BottomBottomLeft.Width = 100.000000
	__RectAnchoredRect__000006_BottomBottomLeft.Height = 70.000000
	__RectAnchoredRect__000006_BottomBottomLeft.RX = 5.000000
	__RectAnchoredRect__000006_BottomBottomLeft.X_Offset = 0.000000
	__RectAnchoredRect__000006_BottomBottomLeft.Y_Offset = 0.000000
	__RectAnchoredRect__000006_BottomBottomLeft.RectAnchorType = models.RECT_BOTTOM_BOTTOM_LEFT
	__RectAnchoredRect__000006_BottomBottomLeft.WidthFollowRect = false
	__RectAnchoredRect__000006_BottomBottomLeft.HeightFollowRect = false
	__RectAnchoredRect__000006_BottomBottomLeft.Color = ``
	__RectAnchoredRect__000006_BottomBottomLeft.FillOpacity = 0.000000
	__RectAnchoredRect__000006_BottomBottomLeft.Stroke = `blue`
	__RectAnchoredRect__000006_BottomBottomLeft.StrokeOpacity = 0.500000
	__RectAnchoredRect__000006_BottomBottomLeft.StrokeWidth = 4.000000
	__RectAnchoredRect__000006_BottomBottomLeft.StrokeDashArray = ``
	__RectAnchoredRect__000006_BottomBottomLeft.StrokeDashArrayWhenSelected = ``
	__RectAnchoredRect__000006_BottomBottomLeft.Transform = ``

	__RectAnchoredRect__000007_BottomRight.Name = `BottomRight`
	__RectAnchoredRect__000007_BottomRight.X = 0.000000
	__RectAnchoredRect__000007_BottomRight.Y = 0.000000
	__RectAnchoredRect__000007_BottomRight.Width = 100.000000
	__RectAnchoredRect__000007_BottomRight.Height = 70.000000
	__RectAnchoredRect__000007_BottomRight.RX = 5.000000
	__RectAnchoredRect__000007_BottomRight.X_Offset = 0.000000
	__RectAnchoredRect__000007_BottomRight.Y_Offset = 0.000000
	__RectAnchoredRect__000007_BottomRight.RectAnchorType = models.RECT_BOTTOM_RIGHT
	__RectAnchoredRect__000007_BottomRight.WidthFollowRect = false
	__RectAnchoredRect__000007_BottomRight.HeightFollowRect = false
	__RectAnchoredRect__000007_BottomRight.Color = ``
	__RectAnchoredRect__000007_BottomRight.FillOpacity = 0.000000
	__RectAnchoredRect__000007_BottomRight.Stroke = `black`
	__RectAnchoredRect__000007_BottomRight.StrokeOpacity = 0.500000
	__RectAnchoredRect__000007_BottomRight.StrokeWidth = 2.000000
	__RectAnchoredRect__000007_BottomRight.StrokeDashArray = ``
	__RectAnchoredRect__000007_BottomRight.StrokeDashArrayWhenSelected = ``
	__RectAnchoredRect__000007_BottomRight.Transform = ``

	__RectAnchoredRect__000008_BottomInsideRight.Name = `BottomInsideRight`
	__RectAnchoredRect__000008_BottomInsideRight.X = 0.000000
	__RectAnchoredRect__000008_BottomInsideRight.Y = 0.000000
	__RectAnchoredRect__000008_BottomInsideRight.Width = 100.000000
	__RectAnchoredRect__000008_BottomInsideRight.Height = 70.000000
	__RectAnchoredRect__000008_BottomInsideRight.RX = 5.000000
	__RectAnchoredRect__000008_BottomInsideRight.X_Offset = 0.000000
	__RectAnchoredRect__000008_BottomInsideRight.Y_Offset = 0.000000
	__RectAnchoredRect__000008_BottomInsideRight.RectAnchorType = models.RECT_BOTTOM_INSIDE_RIGHT
	__RectAnchoredRect__000008_BottomInsideRight.WidthFollowRect = false
	__RectAnchoredRect__000008_BottomInsideRight.HeightFollowRect = false
	__RectAnchoredRect__000008_BottomInsideRight.Color = ``
	__RectAnchoredRect__000008_BottomInsideRight.FillOpacity = 0.000000
	__RectAnchoredRect__000008_BottomInsideRight.Stroke = `red`
	__RectAnchoredRect__000008_BottomInsideRight.StrokeOpacity = 0.500000
	__RectAnchoredRect__000008_BottomInsideRight.StrokeWidth = 9.000000
	__RectAnchoredRect__000008_BottomInsideRight.StrokeDashArray = ``
	__RectAnchoredRect__000008_BottomInsideRight.StrokeDashArrayWhenSelected = ``
	__RectAnchoredRect__000008_BottomInsideRight.Transform = ``

	__RectAnchoredRect__000009_Left.Name = `Left`
	__RectAnchoredRect__000009_Left.X = 0.000000
	__RectAnchoredRect__000009_Left.Y = 0.000000
	__RectAnchoredRect__000009_Left.Width = 100.000000
	__RectAnchoredRect__000009_Left.Height = 70.000000
	__RectAnchoredRect__000009_Left.RX = 5.000000
	__RectAnchoredRect__000009_Left.X_Offset = 0.000000
	__RectAnchoredRect__000009_Left.Y_Offset = 0.000000
	__RectAnchoredRect__000009_Left.RectAnchorType = models.RECT_LEFT
	__RectAnchoredRect__000009_Left.WidthFollowRect = false
	__RectAnchoredRect__000009_Left.HeightFollowRect = false
	__RectAnchoredRect__000009_Left.Color = ``
	__RectAnchoredRect__000009_Left.FillOpacity = 0.000000
	__RectAnchoredRect__000009_Left.Stroke = `violet`
	__RectAnchoredRect__000009_Left.StrokeOpacity = 0.500000
	__RectAnchoredRect__000009_Left.StrokeWidth = 2.000000
	__RectAnchoredRect__000009_Left.StrokeDashArray = ``
	__RectAnchoredRect__000009_Left.StrokeDashArrayWhenSelected = ``
	__RectAnchoredRect__000009_Left.Transform = ``

	__RectAnchoredRect__000010_Right.Name = `Right`
	__RectAnchoredRect__000010_Right.X = 0.000000
	__RectAnchoredRect__000010_Right.Y = 0.000000
	__RectAnchoredRect__000010_Right.Width = 100.000000
	__RectAnchoredRect__000010_Right.Height = 70.000000
	__RectAnchoredRect__000010_Right.RX = 5.000000
	__RectAnchoredRect__000010_Right.X_Offset = 0.000000
	__RectAnchoredRect__000010_Right.Y_Offset = 0.000000
	__RectAnchoredRect__000010_Right.RectAnchorType = models.RECT_RIGHT
	__RectAnchoredRect__000010_Right.WidthFollowRect = false
	__RectAnchoredRect__000010_Right.HeightFollowRect = false
	__RectAnchoredRect__000010_Right.Color = ``
	__RectAnchoredRect__000010_Right.FillOpacity = 0.000000
	__RectAnchoredRect__000010_Right.Stroke = `black`
	__RectAnchoredRect__000010_Right.StrokeOpacity = 0.500000
	__RectAnchoredRect__000010_Right.StrokeWidth = 2.000000
	__RectAnchoredRect__000010_Right.StrokeDashArray = ``
	__RectAnchoredRect__000010_Right.StrokeDashArrayWhenSelected = ``
	__RectAnchoredRect__000010_Right.Transform = ``

	__RectAnchoredRect__000011_Center.Name = `Center`
	__RectAnchoredRect__000011_Center.X = 0.000000
	__RectAnchoredRect__000011_Center.Y = 0.000000
	__RectAnchoredRect__000011_Center.Width = 100.000000
	__RectAnchoredRect__000011_Center.Height = 70.000000
	__RectAnchoredRect__000011_Center.RX = 5.000000
	__RectAnchoredRect__000011_Center.X_Offset = 0.000000
	__RectAnchoredRect__000011_Center.Y_Offset = 0.000000
	__RectAnchoredRect__000011_Center.RectAnchorType = models.RECT_CENTER
	__RectAnchoredRect__000011_Center.WidthFollowRect = false
	__RectAnchoredRect__000011_Center.HeightFollowRect = false
	__RectAnchoredRect__000011_Center.Color = ``
	__RectAnchoredRect__000011_Center.FillOpacity = 0.000000
	__RectAnchoredRect__000011_Center.Stroke = `black`
	__RectAnchoredRect__000011_Center.StrokeOpacity = 0.500000
	__RectAnchoredRect__000011_Center.StrokeWidth = 2.000000
	__RectAnchoredRect__000011_Center.StrokeDashArray = ``
	__RectAnchoredRect__000011_Center.StrokeDashArrayWhenSelected = ``
	__RectAnchoredRect__000011_Center.Transform = ``

	__RectAnchoredText__000000_Oriented.Name = `Oriented`
	__RectAnchoredText__000000_Oriented.Content = `Oriented`
	__RectAnchoredText__000000_Oriented.FontWeight = ``
	__RectAnchoredText__000000_Oriented.FontSize = ``
	__RectAnchoredText__000000_Oriented.FontStyle = ``
	__RectAnchoredText__000000_Oriented.LetterSpacing = ``
	__RectAnchoredText__000000_Oriented.X_Offset = 0.000000
	__RectAnchoredText__000000_Oriented.Y_Offset = 0.000000
	__RectAnchoredText__000000_Oriented.RectAnchorType = models.RECT_CENTER
	__RectAnchoredText__000000_Oriented.TextAnchorType = models.TEXT_ANCHOR_CENTER
	__RectAnchoredText__000000_Oriented.WritingMode = models.WritingModeVerticalRL
	__RectAnchoredText__000000_Oriented.Color = `black`
	__RectAnchoredText__000000_Oriented.FillOpacity = 1.000000
	__RectAnchoredText__000000_Oriented.Stroke = `black`
	__RectAnchoredText__000000_Oriented.StrokeOpacity = 0.500000
	__RectAnchoredText__000000_Oriented.StrokeWidth = -5.000000
	__RectAnchoredText__000000_Oriented.StrokeDashArray = ``
	__RectAnchoredText__000000_Oriented.StrokeDashArrayWhenSelected = ``
	__RectAnchoredText__000000_Oriented.Transform = ``

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
	__Rect__000000_Start.RectAnchoredTexts = append(__Rect__000000_Start.RectAnchoredTexts, __RectAnchoredText__000000_Oriented)
	__Rect__000001_End.RectAnchoredRects = append(__Rect__000001_End.RectAnchoredRects, __RectAnchoredRect__000003_Test)
	__Rect__000001_End.RectAnchoredRects = append(__Rect__000001_End.RectAnchoredRects, __RectAnchoredRect__000010_Right)
	__Rect__000001_End.RectAnchoredRects = append(__Rect__000001_End.RectAnchoredRects, __RectAnchoredRect__000004_BottomLeft)
	__Rect__000001_End.RectAnchoredRects = append(__Rect__000001_End.RectAnchoredRects, __RectAnchoredRect__000009_Left)
	__Rect__000001_End.RectAnchoredRects = append(__Rect__000001_End.RectAnchoredRects, __RectAnchoredRect__000011_Center)
	__Rect__000001_End.RectAnchoredRects = append(__Rect__000001_End.RectAnchoredRects, __RectAnchoredRect__000006_BottomBottomLeft)
	__Rect__000001_End.RectAnchoredRects = append(__Rect__000001_End.RectAnchoredRects, __RectAnchoredRect__000008_BottomInsideRight)
	__Rect__000001_End.RectAnchoredRects = append(__Rect__000001_End.RectAnchoredRects, __RectAnchoredRect__000000_Top)
	__Rect__000001_End.RectAnchoredRects = append(__Rect__000001_End.RectAnchoredRects, __RectAnchoredRect__000001_TopLeft)
	__Rect__000001_End.RectAnchoredRects = append(__Rect__000001_End.RectAnchoredRects, __RectAnchoredRect__000002_TopRight)
	__Rect__000001_End.RectAnchoredRects = append(__Rect__000001_End.RectAnchoredRects, __RectAnchoredRect__000005_BottomLeftLeft)
	__Rect__000001_End.RectAnchoredRects = append(__Rect__000001_End.RectAnchoredRects, __RectAnchoredRect__000007_BottomRight)
	// setup of RectAnchoredRect instances pointers
	// setup of RectAnchoredText instances pointers
	// setup of SVG instances pointers
	__SVG__000000_simple.Layers = append(__SVG__000000_simple.Layers, __Layer__000000_Layer_1)
}
