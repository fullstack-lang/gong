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
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__Animate__000000_ := (&models.Animate{}).Stage(stage)

	__Circle__000000_Test := (&models.Circle{}).Stage(stage)

	__Layer__000000_Bottom_Rectangle_Layer := (&models.Layer{}).Stage(stage)
	__Layer__000001_Layer_RectLinkLink_Medium_to_Top_Bottom := (&models.Layer{}).Stage(stage)
	__Layer__000002_Line_layer := (&models.Layer{}).Stage(stage)
	__Layer__000003_Link_layer_vertical_to_horizontal := (&models.Layer{}).Stage(stage)
	__Layer__000004_Middle_Rect_Layer := (&models.Layer{}).Stage(stage)
	__Layer__000005_Top_Rectangle_layer := (&models.Layer{}).Stage(stage)

	__Line__000000_Line_connecting_rect_Bottom_to_Top := (&models.Line{}).Stage(stage)

	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal := (&models.Link{}).Stage(stage)
	__Link__000001_Start_Middle := (&models.Link{}).Stage(stage)

	__LinkAnchoredText__000000_End_Left_Top := (&models.LinkAnchoredText{}).Stage(stage)
	__LinkAnchoredText__000001_End_Right_Bottom := (&models.LinkAnchoredText{}).Stage(stage)
	__LinkAnchoredText__000002_Liine_1_Line_2 := (&models.LinkAnchoredText{}).Stage(stage)
	__LinkAnchoredText__000003_Start_Anchored_1 := (&models.LinkAnchoredText{}).Stage(stage)
	__LinkAnchoredText__000004_Start_Left_Top := (&models.LinkAnchoredText{}).Stage(stage)
	__LinkAnchoredText__000005_Start_Right_Bottom := (&models.LinkAnchoredText{}).Stage(stage)

	__Rect__000000_Bottom := (&models.Rect{}).Stage(stage)
	__Rect__000001_Middle_Rect := (&models.Rect{}).Stage(stage)
	__Rect__000002_Top := (&models.Rect{}).Stage(stage)

	__RectAnchoredPath__000000_Logo_to_add := (&models.RectAnchoredPath{}).Stage(stage)
	__RectAnchoredPath__000001_second_log := (&models.RectAnchoredPath{}).Stage(stage)

	__RectAnchoredRect__000000_Rect_within_top := (&models.RectAnchoredRect{}).Stage(stage)
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width := (&models.RectAnchoredRect{}).Stage(stage)

	__RectAnchoredText__000000_Bottom_Text := (&models.RectAnchoredText{}).Stage(stage)
	__RectAnchoredText__000001_Top_Left := (&models.RectAnchoredText{}).Stage(stage)
	__RectAnchoredText__000002_Top_anchored_bottom_middle := (&models.RectAnchoredText{}).Stage(stage)
	__RectAnchoredText__000003_Top_anchored_top_middle := (&models.RectAnchoredText{}).Stage(stage)

	__RectLinkLink__000000_Test_Middle_to_Top_Bottom_Link := (&models.RectLinkLink{}).Stage(stage)

	__SVG__000000_SVG := (&models.SVG{}).Stage(stage)

	__SvgText__000000_Essai := (&models.SvgText{}).Stage(stage)

	__Text__000000_Essai := (&models.Text{}).Stage(stage)

	// Setup of values

	__Animate__000000_.Name = ``
	__Animate__000000_.AttributeName = ``
	__Animate__000000_.Values = ``
	__Animate__000000_.From = ``
	__Animate__000000_.To = ``
	__Animate__000000_.Dur = ``
	__Animate__000000_.RepeatCount = ``

	__Circle__000000_Test.Name = `Test`
	__Circle__000000_Test.CX = 400.000000
	__Circle__000000_Test.CY = 300.000000
	__Circle__000000_Test.Radius = 100.000000
	__Circle__000000_Test.Color = `lavender`
	__Circle__000000_Test.FillOpacity = 50.000000
	__Circle__000000_Test.Stroke = ``
	__Circle__000000_Test.StrokeOpacity = 1.000000
	__Circle__000000_Test.StrokeWidth = 0.000000
	__Circle__000000_Test.StrokeDashArray = ``
	__Circle__000000_Test.StrokeDashArrayWhenSelected = ``
	__Circle__000000_Test.Transform = ``

	__Layer__000000_Bottom_Rectangle_Layer.Display = false
	__Layer__000000_Bottom_Rectangle_Layer.Name = `Bottom Rectangle Layer`

	__Layer__000001_Layer_RectLinkLink_Medium_to_Top_Bottom.Display = false
	__Layer__000001_Layer_RectLinkLink_Medium_to_Top_Bottom.Name = `Layer RectLinkLink Medium to Top-Bottom`

	__Layer__000002_Line_layer.Display = false
	__Layer__000002_Line_layer.Name = `Line layer`

	__Layer__000003_Link_layer_vertical_to_horizontal.Display = false
	__Layer__000003_Link_layer_vertical_to_horizontal.Name = `Link layer vertical to horizontal`

	__Layer__000004_Middle_Rect_Layer.Display = false
	__Layer__000004_Middle_Rect_Layer.Name = `Middle Rect Layer`

	__Layer__000005_Top_Rectangle_layer.Display = false
	__Layer__000005_Top_Rectangle_layer.Name = `Top Rectangle layer`

	__Line__000000_Line_connecting_rect_Bottom_to_Top.Name = `Line connecting rect Bottom to Top`
	__Line__000000_Line_connecting_rect_Bottom_to_Top.X1 = 374.000000
	__Line__000000_Line_connecting_rect_Bottom_to_Top.Y1 = 258.000000
	__Line__000000_Line_connecting_rect_Bottom_to_Top.X2 = 555.999969
	__Line__000000_Line_connecting_rect_Bottom_to_Top.Y2 = 523.000000
	__Line__000000_Line_connecting_rect_Bottom_to_Top.Color = `olivedrab`
	__Line__000000_Line_connecting_rect_Bottom_to_Top.FillOpacity = 0.000000
	__Line__000000_Line_connecting_rect_Bottom_to_Top.Stroke = `olivedrab`
	__Line__000000_Line_connecting_rect_Bottom_to_Top.StrokeOpacity = 1.000000
	__Line__000000_Line_connecting_rect_Bottom_to_Top.StrokeWidth = 4.000000
	__Line__000000_Line_connecting_rect_Bottom_to_Top.StrokeDashArray = ``
	__Line__000000_Line_connecting_rect_Bottom_to_Top.StrokeDashArrayWhenSelected = ``
	__Line__000000_Line_connecting_rect_Bottom_to_Top.Transform = ``
	__Line__000000_Line_connecting_rect_Bottom_to_Top.MouseClickX = 0.000000
	__Line__000000_Line_connecting_rect_Bottom_to_Top.MouseClickY = 0.000000

	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.Name = `Arrow - Top to Bottom vertical to horizontal`
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.Type = models.LINK_TYPE_FLOATING_ORTHOGONAL
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.IsBezierCurve = false
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.StartAnchorType = models.ANCHOR_CENTER
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.EndAnchorType = models.ANCHOR_CENTER
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.StartRatio = 0.740216
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.EndRatio = 0.659649
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.CornerOffsetRatio = -1.556962
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.CornerRadius = 8.000000
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.HasEndArrow = true
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.EndArrowSize = 10.000000
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.HasStartArrow = false
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.StartArrowSize = 0.000000
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.Color = ``
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.FillOpacity = 0.000000
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.Stroke = `green`
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.StrokeOpacity = 1.000000
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.StrokeWidth = 4.000000
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.StrokeDashArray = ``
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.StrokeDashArrayWhenSelected = ``
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.Transform = ``

	__Link__000001_Start_Middle.Name = `Start - Middle`
	__Link__000001_Start_Middle.Type = models.LINK_TYPE_FLOATING_ORTHOGONAL
	__Link__000001_Start_Middle.IsBezierCurve = false
	__Link__000001_Start_Middle.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_Start_Middle.StartRatio = 0.218766
	__Link__000001_Start_Middle.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_Start_Middle.EndRatio = 0.478101
	__Link__000001_Start_Middle.CornerOffsetRatio = -1.105485
	__Link__000001_Start_Middle.CornerRadius = 20.000000
	__Link__000001_Start_Middle.HasEndArrow = true
	__Link__000001_Start_Middle.EndArrowSize = 10.000000
	__Link__000001_Start_Middle.HasStartArrow = true
	__Link__000001_Start_Middle.StartArrowSize = 10.000000
	__Link__000001_Start_Middle.Color = ``
	__Link__000001_Start_Middle.FillOpacity = 0.000000
	__Link__000001_Start_Middle.Stroke = `red`
	__Link__000001_Start_Middle.StrokeOpacity = 0.700000
	__Link__000001_Start_Middle.StrokeWidth = 5.000000
	__Link__000001_Start_Middle.StrokeDashArray = ``
	__Link__000001_Start_Middle.StrokeDashArrayWhenSelected = ``
	__Link__000001_Start_Middle.Transform = ``

	__LinkAnchoredText__000000_End_Left_Top.Name = `End Left/Top`
	__LinkAnchoredText__000000_End_Left_Top.Content = `End Left/Top`
	__LinkAnchoredText__000000_End_Left_Top.AutomaticLayout = true
	__LinkAnchoredText__000000_End_Left_Top.LinkAnchorType = models.LINK_LEFT_OR_TOP
	__LinkAnchoredText__000000_End_Left_Top.X_Offset = -40.000000
	__LinkAnchoredText__000000_End_Left_Top.Y_Offset = 0.000000
	__LinkAnchoredText__000000_End_Left_Top.FontWeight = `normal`
	__LinkAnchoredText__000000_End_Left_Top.FontSize = ``
	__LinkAnchoredText__000000_End_Left_Top.LetterSpacing = ``
	__LinkAnchoredText__000000_End_Left_Top.Color = `blue`
	__LinkAnchoredText__000000_End_Left_Top.FillOpacity = 0.000000
	__LinkAnchoredText__000000_End_Left_Top.Stroke = `blue`
	__LinkAnchoredText__000000_End_Left_Top.StrokeOpacity = 1.000000
	__LinkAnchoredText__000000_End_Left_Top.StrokeWidth = 1.000000
	__LinkAnchoredText__000000_End_Left_Top.StrokeDashArray = ``
	__LinkAnchoredText__000000_End_Left_Top.StrokeDashArrayWhenSelected = ``
	__LinkAnchoredText__000000_End_Left_Top.Transform = ``

	__LinkAnchoredText__000001_End_Right_Bottom.Name = `End Right/Bottom`
	__LinkAnchoredText__000001_End_Right_Bottom.Content = `End Right/Bottom`
	__LinkAnchoredText__000001_End_Right_Bottom.AutomaticLayout = true
	__LinkAnchoredText__000001_End_Right_Bottom.LinkAnchorType = models.LINK_RIGHT_OR_BOTTOM
	__LinkAnchoredText__000001_End_Right_Bottom.X_Offset = 0.000000
	__LinkAnchoredText__000001_End_Right_Bottom.Y_Offset = 20.000000
	__LinkAnchoredText__000001_End_Right_Bottom.FontWeight = `normal`
	__LinkAnchoredText__000001_End_Right_Bottom.FontSize = ``
	__LinkAnchoredText__000001_End_Right_Bottom.LetterSpacing = ``
	__LinkAnchoredText__000001_End_Right_Bottom.Color = `black`
	__LinkAnchoredText__000001_End_Right_Bottom.FillOpacity = 1.000000
	__LinkAnchoredText__000001_End_Right_Bottom.Stroke = `black`
	__LinkAnchoredText__000001_End_Right_Bottom.StrokeOpacity = 1.000000
	__LinkAnchoredText__000001_End_Right_Bottom.StrokeWidth = 1.000000
	__LinkAnchoredText__000001_End_Right_Bottom.StrokeDashArray = ``
	__LinkAnchoredText__000001_End_Right_Bottom.StrokeDashArrayWhenSelected = ``
	__LinkAnchoredText__000001_End_Right_Bottom.Transform = ``

	__LinkAnchoredText__000002_Liine_1_Line_2.Name = `Liine 1
Line 2`
	__LinkAnchoredText__000002_Liine_1_Line_2.Content = `Liine 1
Line 2`
	__LinkAnchoredText__000002_Liine_1_Line_2.AutomaticLayout = false
	__LinkAnchoredText__000002_Liine_1_Line_2.X_Offset = -61.000000
	__LinkAnchoredText__000002_Liine_1_Line_2.Y_Offset = -35.000000
	__LinkAnchoredText__000002_Liine_1_Line_2.FontWeight = `100`
	__LinkAnchoredText__000002_Liine_1_Line_2.FontSize = `16`
	__LinkAnchoredText__000002_Liine_1_Line_2.LetterSpacing = `0.1em`
	__LinkAnchoredText__000002_Liine_1_Line_2.Color = `black`
	__LinkAnchoredText__000002_Liine_1_Line_2.FillOpacity = 100.000000
	__LinkAnchoredText__000002_Liine_1_Line_2.Stroke = `black`
	__LinkAnchoredText__000002_Liine_1_Line_2.StrokeOpacity = 1.000000
	__LinkAnchoredText__000002_Liine_1_Line_2.StrokeWidth = 1.000000
	__LinkAnchoredText__000002_Liine_1_Line_2.StrokeDashArray = ``
	__LinkAnchoredText__000002_Liine_1_Line_2.StrokeDashArrayWhenSelected = ``
	__LinkAnchoredText__000002_Liine_1_Line_2.Transform = ``

	__LinkAnchoredText__000003_Start_Anchored_1.Name = `Start Anchored 1`
	__LinkAnchoredText__000003_Start_Anchored_1.Content = `Start Anchored 1
Second line
Third Line`
	__LinkAnchoredText__000003_Start_Anchored_1.AutomaticLayout = false
	__LinkAnchoredText__000003_Start_Anchored_1.X_Offset = -129.000031
	__LinkAnchoredText__000003_Start_Anchored_1.Y_Offset = 37.012512
	__LinkAnchoredText__000003_Start_Anchored_1.FontWeight = `light`
	__LinkAnchoredText__000003_Start_Anchored_1.FontSize = ``
	__LinkAnchoredText__000003_Start_Anchored_1.LetterSpacing = ``
	__LinkAnchoredText__000003_Start_Anchored_1.Color = `cyan`
	__LinkAnchoredText__000003_Start_Anchored_1.FillOpacity = 100.000000
	__LinkAnchoredText__000003_Start_Anchored_1.Stroke = `cyan`
	__LinkAnchoredText__000003_Start_Anchored_1.StrokeOpacity = 1.000000
	__LinkAnchoredText__000003_Start_Anchored_1.StrokeWidth = 1.000000
	__LinkAnchoredText__000003_Start_Anchored_1.StrokeDashArray = ``
	__LinkAnchoredText__000003_Start_Anchored_1.StrokeDashArrayWhenSelected = ``
	__LinkAnchoredText__000003_Start_Anchored_1.Transform = ``

	__LinkAnchoredText__000004_Start_Left_Top.Name = `Start Left/Top`
	__LinkAnchoredText__000004_Start_Left_Top.Content = `Start Left/Top`
	__LinkAnchoredText__000004_Start_Left_Top.AutomaticLayout = true
	__LinkAnchoredText__000004_Start_Left_Top.LinkAnchorType = models.LINK_LEFT_OR_TOP
	__LinkAnchoredText__000004_Start_Left_Top.X_Offset = 0.000000
	__LinkAnchoredText__000004_Start_Left_Top.Y_Offset = 0.000000
	__LinkAnchoredText__000004_Start_Left_Top.FontWeight = `normal`
	__LinkAnchoredText__000004_Start_Left_Top.FontSize = ``
	__LinkAnchoredText__000004_Start_Left_Top.LetterSpacing = ``
	__LinkAnchoredText__000004_Start_Left_Top.Color = `black`
	__LinkAnchoredText__000004_Start_Left_Top.FillOpacity = 0.000000
	__LinkAnchoredText__000004_Start_Left_Top.Stroke = `black`
	__LinkAnchoredText__000004_Start_Left_Top.StrokeOpacity = 1.000000
	__LinkAnchoredText__000004_Start_Left_Top.StrokeWidth = 1.000000
	__LinkAnchoredText__000004_Start_Left_Top.StrokeDashArray = ``
	__LinkAnchoredText__000004_Start_Left_Top.StrokeDashArrayWhenSelected = ``
	__LinkAnchoredText__000004_Start_Left_Top.Transform = ``

	__LinkAnchoredText__000005_Start_Right_Bottom.Name = `Start Right/Bottom`
	__LinkAnchoredText__000005_Start_Right_Bottom.Content = `Start Right/Bottom`
	__LinkAnchoredText__000005_Start_Right_Bottom.AutomaticLayout = true
	__LinkAnchoredText__000005_Start_Right_Bottom.LinkAnchorType = models.LINK_RIGHT_OR_BOTTOM
	__LinkAnchoredText__000005_Start_Right_Bottom.X_Offset = 0.000000
	__LinkAnchoredText__000005_Start_Right_Bottom.Y_Offset = 0.000000
	__LinkAnchoredText__000005_Start_Right_Bottom.FontWeight = `normal`
	__LinkAnchoredText__000005_Start_Right_Bottom.FontSize = ``
	__LinkAnchoredText__000005_Start_Right_Bottom.LetterSpacing = ``
	__LinkAnchoredText__000005_Start_Right_Bottom.Color = `black`
	__LinkAnchoredText__000005_Start_Right_Bottom.FillOpacity = 1.000000
	__LinkAnchoredText__000005_Start_Right_Bottom.Stroke = `black`
	__LinkAnchoredText__000005_Start_Right_Bottom.StrokeOpacity = 1.000000
	__LinkAnchoredText__000005_Start_Right_Bottom.StrokeWidth = 1.000000
	__LinkAnchoredText__000005_Start_Right_Bottom.StrokeDashArray = ``
	__LinkAnchoredText__000005_Start_Right_Bottom.StrokeDashArrayWhenSelected = ``
	__LinkAnchoredText__000005_Start_Right_Bottom.Transform = ``

	__Rect__000000_Bottom.Name = `Bottom`
	__Rect__000000_Bottom.X = 355.000000
	__Rect__000000_Bottom.Y = 80.000000
	__Rect__000000_Bottom.Width = 565.000000
	__Rect__000000_Bottom.Height = 266.999985
	__Rect__000000_Bottom.RX = 5.000000
	__Rect__000000_Bottom.Color = `bisque`
	__Rect__000000_Bottom.FillOpacity = 50.000000
	__Rect__000000_Bottom.Stroke = `lightcoral`
	__Rect__000000_Bottom.StrokeOpacity = 1.000000
	__Rect__000000_Bottom.StrokeWidth = 3.000000
	__Rect__000000_Bottom.StrokeDashArray = ``
	__Rect__000000_Bottom.StrokeDashArrayWhenSelected = `5 5`
	__Rect__000000_Bottom.Transform = ``
	__Rect__000000_Bottom.IsSelectable = true
	__Rect__000000_Bottom.IsSelected = false
	__Rect__000000_Bottom.CanHaveLeftHandle = true
	__Rect__000000_Bottom.HasLeftHandle = false
	__Rect__000000_Bottom.CanHaveRightHandle = true
	__Rect__000000_Bottom.HasRightHandle = false
	__Rect__000000_Bottom.CanHaveTopHandle = true
	__Rect__000000_Bottom.HasTopHandle = false
	__Rect__000000_Bottom.IsScalingProportionally = false
	__Rect__000000_Bottom.CanHaveBottomHandle = true
	__Rect__000000_Bottom.HasBottomHandle = false
	__Rect__000000_Bottom.CanMoveHorizontaly = true
	__Rect__000000_Bottom.CanMoveVerticaly = true

	__Rect__000001_Middle_Rect.Name = `Middle Rect`
	__Rect__000001_Middle_Rect.X = 521.000000
	__Rect__000001_Middle_Rect.Y = 404.000000
	__Rect__000001_Middle_Rect.Width = 200.000000
	__Rect__000001_Middle_Rect.Height = 132.000000
	__Rect__000001_Middle_Rect.RX = 3.000000
	__Rect__000001_Middle_Rect.Color = `lavender`
	__Rect__000001_Middle_Rect.FillOpacity = 50.000000
	__Rect__000001_Middle_Rect.Stroke = `turquoise`
	__Rect__000001_Middle_Rect.StrokeOpacity = 1.000000
	__Rect__000001_Middle_Rect.StrokeWidth = 1.000000
	__Rect__000001_Middle_Rect.StrokeDashArray = ``
	__Rect__000001_Middle_Rect.StrokeDashArrayWhenSelected = `5 5`
	__Rect__000001_Middle_Rect.Transform = ``
	__Rect__000001_Middle_Rect.IsSelectable = true
	__Rect__000001_Middle_Rect.IsSelected = false
	__Rect__000001_Middle_Rect.CanHaveLeftHandle = true
	__Rect__000001_Middle_Rect.HasLeftHandle = false
	__Rect__000001_Middle_Rect.CanHaveRightHandle = true
	__Rect__000001_Middle_Rect.HasRightHandle = false
	__Rect__000001_Middle_Rect.CanHaveTopHandle = true
	__Rect__000001_Middle_Rect.HasTopHandle = false
	__Rect__000001_Middle_Rect.IsScalingProportionally = true
	__Rect__000001_Middle_Rect.CanHaveBottomHandle = true
	__Rect__000001_Middle_Rect.HasBottomHandle = false
	__Rect__000001_Middle_Rect.CanMoveHorizontaly = true
	__Rect__000001_Middle_Rect.CanMoveVerticaly = true

	__Rect__000002_Top.Name = `Top`
	__Rect__000002_Top.X = 531.999969
	__Rect__000002_Top.Y = 625.000000
	__Rect__000002_Top.Width = 237.000000
	__Rect__000002_Top.Height = 237.000000
	__Rect__000002_Top.RX = 3.000000
	__Rect__000002_Top.Color = `lightcyan`
	__Rect__000002_Top.FillOpacity = 100.000000
	__Rect__000002_Top.Stroke = `darkcyan`
	__Rect__000002_Top.StrokeOpacity = 1.000000
	__Rect__000002_Top.StrokeWidth = 2.000000
	__Rect__000002_Top.StrokeDashArray = ``
	__Rect__000002_Top.StrokeDashArrayWhenSelected = `5 5`
	__Rect__000002_Top.Transform = ``
	__Rect__000002_Top.IsSelectable = true
	__Rect__000002_Top.IsSelected = false
	__Rect__000002_Top.CanHaveLeftHandle = true
	__Rect__000002_Top.HasLeftHandle = false
	__Rect__000002_Top.CanHaveRightHandle = true
	__Rect__000002_Top.HasRightHandle = false
	__Rect__000002_Top.CanHaveTopHandle = false
	__Rect__000002_Top.HasTopHandle = false
	__Rect__000002_Top.IsScalingProportionally = true
	__Rect__000002_Top.CanHaveBottomHandle = false
	__Rect__000002_Top.HasBottomHandle = false
	__Rect__000002_Top.CanMoveHorizontaly = true
	__Rect__000002_Top.CanMoveVerticaly = true

	__RectAnchoredPath__000000_Logo_to_add.Name = `Logo to add`
	__RectAnchoredPath__000000_Logo_to_add.Definition = `M532-131q-6 5-12.5 8t-14.5 3q-8 0-16-3.5t-14-9.5q-41-44-60.5-90T395-320q0-37 11-78t38-106q23-57 32-87.5t9-56.5q0-34-15-63.5T423-771q-6-6-9.5-14t-3.5-16q0-8 3-14.5t8-12.5q6-6 13.5-9t15.5-3q8 0 15 3t13 8q44 41 65.5 86t21.5 95q0 35-10.5 73.5T518-474q-25 60-34 92t-9 61q0 35 14.5 67.5T534-188q5 6 8 13t3 15q0 8-3 15.5T532-131Zm195 0q-6 5-12.5 8t-14.5 3q-8 0-16-3.5t-14-9.5q-41-44-60.5-89.5T590-319q0-37 11-79t38-106q23-57 32-87t9-56q0-34-15-64.5T618-771q-6-6-9-13.5t-3-15.5q0-8 2.5-14.5T616-827q6-6 14-9.5t16-3.5q8 0 14.5 3t12.5 8q44 41 65.5 86t21.5 95q0 35-10.5 73.5T713-473q-25 60-34 92t-9 60q0 35 15 68.5t45 65.5q5 6 7.5 13t2.5 14q0 8-3 16t-10 13Zm-390 0q-6 5-12.5 8t-14.5 3q-8 0-16-3.5t-14-9.5q-41-44-60.5-89.5T200-319q0-37 11-79t38-106q23-57 32-87t9-56q0-34-15-64.5T228-771q-7-6-10-13.5t-3-15.5q0-8 3-15t8-13q6-6 13.5-9t15.5-3q8 0 15 3t13 8q44 41 65.5 85.5T370-648q0 35-10 73.5T324-474q-25 60-34 92t-9 61q0 35 14.5 68.5T340-187q5 6 7.5 13t2.5 14q0 8-3 16t-10 13Z`
	__RectAnchoredPath__000000_Logo_to_add.X_Offset = 0.000000
	__RectAnchoredPath__000000_Logo_to_add.Y_Offset = 0.000000
	__RectAnchoredPath__000000_Logo_to_add.RectAnchorType = models.RECT_BOTTOM_LEFT
	__RectAnchoredPath__000000_Logo_to_add.ScalePropotionnally = true
	__RectAnchoredPath__000000_Logo_to_add.AppliedScaling = 0.246875
	__RectAnchoredPath__000000_Logo_to_add.Color = `black`
	__RectAnchoredPath__000000_Logo_to_add.FillOpacity = 0.500000
	__RectAnchoredPath__000000_Logo_to_add.Stroke = `black`
	__RectAnchoredPath__000000_Logo_to_add.StrokeOpacity = 1.000000
	__RectAnchoredPath__000000_Logo_to_add.StrokeWidth = 1.000000
	__RectAnchoredPath__000000_Logo_to_add.StrokeDashArray = ``
	__RectAnchoredPath__000000_Logo_to_add.StrokeDashArrayWhenSelected = ``
	__RectAnchoredPath__000000_Logo_to_add.Transform = ``

	__RectAnchoredPath__000001_second_log.Name = `second log`
	__RectAnchoredPath__000001_second_log.Definition = `M532-131q-6 5-12.5 8t-14.5 3q-8 0-16-3.5t-14-9.5q-41-44-60.5-90T395-320q0-37 11-78t38-106q23-57 32-87.5t9-56.5q0-34-15-63.5T423-771q-6-6-9.5-14t-3.5-16q0-8 3-14.5t8-12.5q6-6 13.5-9t15.5-3q8 0 15 3t13 8q44 41 65.5 86t21.5 95q0 35-10.5 73.5T518-474q-25 60-34 92t-9 61q0 35 14.5 67.5T534-188q5 6 8 13t3 15q0 8-3 15.5T532-131Zm195 0q-6 5-12.5 8t-14.5 3q-8 0-16-3.5t-14-9.5q-41-44-60.5-89.5T590-319q0-37 11-79t38-106q23-57 32-87t9-56q0-34-15-64.5T618-771q-6-6-9-13.5t-3-15.5q0-8 2.5-14.5T616-827q6-6 14-9.5t16-3.5q8 0 14.5 3t12.5 8q44 41 65.5 86t21.5 95q0 35-10.5 73.5T713-473q-25 60-34 92t-9 60q0 35 15 68.5t45 65.5q5 6 7.5 13t2.5 14q0 8-3 16t-10 13Zm-390 0q-6 5-12.5 8t-14.5 3q-8 0-16-3.5t-14-9.5q-41-44-60.5-89.5T200-319q0-37 11-79t38-106q23-57 32-87t9-56q0-34-15-64.5T228-771q-7-6-10-13.5t-3-15.5q0-8 3-15t8-13q6-6 13.5-9t15.5-3q8 0 15 3t13 8q44 41 65.5 85.5T370-648q0 35-10 73.5T324-474q-25 60-34 92t-9 61q0 35 14.5 68.5T340-187q5 6 7.5 13t2.5 14q0 8-3 16t-10 13Z`
	__RectAnchoredPath__000001_second_log.X_Offset = 0.000000
	__RectAnchoredPath__000001_second_log.Y_Offset = 0.000000
	__RectAnchoredPath__000001_second_log.RectAnchorType = models.RECT_BOTTOM_LEFT
	__RectAnchoredPath__000001_second_log.ScalePropotionnally = true
	__RectAnchoredPath__000001_second_log.AppliedScaling = 0.200000
	__RectAnchoredPath__000001_second_log.Color = ``
	__RectAnchoredPath__000001_second_log.FillOpacity = 0.000000
	__RectAnchoredPath__000001_second_log.Stroke = `blue`
	__RectAnchoredPath__000001_second_log.StrokeOpacity = 1.000000
	__RectAnchoredPath__000001_second_log.StrokeWidth = 2.000000
	__RectAnchoredPath__000001_second_log.StrokeDashArray = ``
	__RectAnchoredPath__000001_second_log.StrokeDashArrayWhenSelected = ``
	__RectAnchoredPath__000001_second_log.Transform = ``

	__RectAnchoredRect__000000_Rect_within_top.Name = `Rect within top`
	__RectAnchoredRect__000000_Rect_within_top.X = 0.000000
	__RectAnchoredRect__000000_Rect_within_top.Y = 0.000000
	__RectAnchoredRect__000000_Rect_within_top.Width = 100.000000
	__RectAnchoredRect__000000_Rect_within_top.Height = 30.000000
	__RectAnchoredRect__000000_Rect_within_top.RX = 3.000000
	__RectAnchoredRect__000000_Rect_within_top.X_Offset = 0.000000
	__RectAnchoredRect__000000_Rect_within_top.Y_Offset = 40.000000
	__RectAnchoredRect__000000_Rect_within_top.RectAnchorType = models.RECT_TOP
	__RectAnchoredRect__000000_Rect_within_top.WidthFollowRect = false
	__RectAnchoredRect__000000_Rect_within_top.HeightFollowRect = false
	__RectAnchoredRect__000000_Rect_within_top.Color = `lightgrey`
	__RectAnchoredRect__000000_Rect_within_top.FillOpacity = 40.000000
	__RectAnchoredRect__000000_Rect_within_top.Stroke = `bisque`
	__RectAnchoredRect__000000_Rect_within_top.StrokeOpacity = 1.000000
	__RectAnchoredRect__000000_Rect_within_top.StrokeWidth = 1.000000
	__RectAnchoredRect__000000_Rect_within_top.StrokeDashArray = ``
	__RectAnchoredRect__000000_Rect_within_top.StrokeDashArrayWhenSelected = ``
	__RectAnchoredRect__000000_Rect_within_top.Transform = ``

	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.Name = `Top on Bottom with same width`
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.X = 0.000000
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.Y = 0.000000
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.Width = 100.000000
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.Height = 50.000000
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.RX = 0.000000
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.X_Offset = 0.000000
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.Y_Offset = 0.000000
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.RectAnchorType = models.RECT_TOP_LEFT
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.WidthFollowRect = true
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.HeightFollowRect = false
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.Color = `lightblue`
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.FillOpacity = 100.000000
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.Stroke = ``
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.StrokeOpacity = 1.000000
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.StrokeWidth = 0.000000
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.StrokeDashArray = ``
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.StrokeDashArrayWhenSelected = ``
	__RectAnchoredRect__000001_Top_on_Bottom_with_same_width.Transform = ``

	__RectAnchoredText__000000_Bottom_Text.Name = `Bottom Text`
	__RectAnchoredText__000000_Bottom_Text.Content = `This is an example of a note that
could be displayed on a diagram.

It could explain one aspect of the model
for instance, describing relations between structs

The text of a UML note refers a comment with the GONGNOTE keyword which is
a special case of go Note convention. See example
for details in the go code of the models.
`
	__RectAnchoredText__000000_Bottom_Text.FontWeight = `100`
	__RectAnchoredText__000000_Bottom_Text.FontSize = 0
	__RectAnchoredText__000000_Bottom_Text.FontStyle = ``
	__RectAnchoredText__000000_Bottom_Text.X_Offset = 0.000000
	__RectAnchoredText__000000_Bottom_Text.Y_Offset = 20.000000
	__RectAnchoredText__000000_Bottom_Text.RectAnchorType = models.RECT_TOP
	__RectAnchoredText__000000_Bottom_Text.TextAnchorType = models.TEXT_ANCHOR_CENTER
	__RectAnchoredText__000000_Bottom_Text.Color = `black`
	__RectAnchoredText__000000_Bottom_Text.FillOpacity = 100.000000
	__RectAnchoredText__000000_Bottom_Text.Stroke = `black`
	__RectAnchoredText__000000_Bottom_Text.StrokeOpacity = 1.000000
	__RectAnchoredText__000000_Bottom_Text.StrokeWidth = 1.000000
	__RectAnchoredText__000000_Bottom_Text.StrokeDashArray = ``
	__RectAnchoredText__000000_Bottom_Text.StrokeDashArrayWhenSelected = ``
	__RectAnchoredText__000000_Bottom_Text.Transform = ``

	__RectAnchoredText__000001_Top_Left.Name = `Top Left`
	__RectAnchoredText__000001_Top_Left.Content = `Top Left`
	__RectAnchoredText__000001_Top_Left.FontWeight = ``
	__RectAnchoredText__000001_Top_Left.FontSize = 0
	__RectAnchoredText__000001_Top_Left.FontStyle = ``
	__RectAnchoredText__000001_Top_Left.X_Offset = 0.000000
	__RectAnchoredText__000001_Top_Left.Y_Offset = 0.000000
	__RectAnchoredText__000001_Top_Left.RectAnchorType = models.RECT_TOP_LEFT
	__RectAnchoredText__000001_Top_Left.Color = ``
	__RectAnchoredText__000001_Top_Left.FillOpacity = 0.000000
	__RectAnchoredText__000001_Top_Left.Stroke = `black`
	__RectAnchoredText__000001_Top_Left.StrokeOpacity = 1.000000
	__RectAnchoredText__000001_Top_Left.StrokeWidth = 0.000000
	__RectAnchoredText__000001_Top_Left.StrokeDashArray = ``
	__RectAnchoredText__000001_Top_Left.StrokeDashArrayWhenSelected = ``
	__RectAnchoredText__000001_Top_Left.Transform = ``

	__RectAnchoredText__000002_Top_anchored_bottom_middle.Name = `Top anchored bottom middle`
	__RectAnchoredText__000002_Top_anchored_bottom_middle.Content = `Top anchored bottom middle`
	__RectAnchoredText__000002_Top_anchored_bottom_middle.FontWeight = ``
	__RectAnchoredText__000002_Top_anchored_bottom_middle.FontSize = 0
	__RectAnchoredText__000002_Top_anchored_bottom_middle.FontStyle = ``
	__RectAnchoredText__000002_Top_anchored_bottom_middle.X_Offset = 0.000000
	__RectAnchoredText__000002_Top_anchored_bottom_middle.Y_Offset = 20.000000
	__RectAnchoredText__000002_Top_anchored_bottom_middle.RectAnchorType = models.RECT_BOTTOM
	__RectAnchoredText__000002_Top_anchored_bottom_middle.TextAnchorType = models.TEXT_ANCHOR_START
	__RectAnchoredText__000002_Top_anchored_bottom_middle.Color = ``
	__RectAnchoredText__000002_Top_anchored_bottom_middle.FillOpacity = 100.000000
	__RectAnchoredText__000002_Top_anchored_bottom_middle.Stroke = `blue`
	__RectAnchoredText__000002_Top_anchored_bottom_middle.StrokeOpacity = 1.000000
	__RectAnchoredText__000002_Top_anchored_bottom_middle.StrokeWidth = 2.000000
	__RectAnchoredText__000002_Top_anchored_bottom_middle.StrokeDashArray = ``
	__RectAnchoredText__000002_Top_anchored_bottom_middle.StrokeDashArrayWhenSelected = ``
	__RectAnchoredText__000002_Top_anchored_bottom_middle.Transform = ``

	__RectAnchoredText__000003_Top_anchored_top_middle.Name = `Top anchored top middle`
	__RectAnchoredText__000003_Top_anchored_top_middle.Content = `Top anchored
top middle
line 3`
	__RectAnchoredText__000003_Top_anchored_top_middle.FontWeight = `bold`
	__RectAnchoredText__000003_Top_anchored_top_middle.FontSize = 0
	__RectAnchoredText__000003_Top_anchored_top_middle.FontStyle = `italic`
	__RectAnchoredText__000003_Top_anchored_top_middle.X_Offset = 0.000000
	__RectAnchoredText__000003_Top_anchored_top_middle.Y_Offset = 20.000000
	__RectAnchoredText__000003_Top_anchored_top_middle.RectAnchorType = models.RECT_TOP
	__RectAnchoredText__000003_Top_anchored_top_middle.TextAnchorType = models.TEXT_ANCHOR_START
	__RectAnchoredText__000003_Top_anchored_top_middle.Color = ``
	__RectAnchoredText__000003_Top_anchored_top_middle.FillOpacity = 100.000000
	__RectAnchoredText__000003_Top_anchored_top_middle.Stroke = `black`
	__RectAnchoredText__000003_Top_anchored_top_middle.StrokeOpacity = 1.000000
	__RectAnchoredText__000003_Top_anchored_top_middle.StrokeWidth = 1.000000
	__RectAnchoredText__000003_Top_anchored_top_middle.StrokeDashArray = ``
	__RectAnchoredText__000003_Top_anchored_top_middle.StrokeDashArrayWhenSelected = ``
	__RectAnchoredText__000003_Top_anchored_top_middle.Transform = ``

	__RectLinkLink__000000_Test_Middle_to_Top_Bottom_Link.Name = `Test Middle to Top-Bottom Link`
	__RectLinkLink__000000_Test_Middle_to_Top_Bottom_Link.TargetAnchorPosition = 0.600000
	__RectLinkLink__000000_Test_Middle_to_Top_Bottom_Link.Color = ``
	__RectLinkLink__000000_Test_Middle_to_Top_Bottom_Link.FillOpacity = 0.000000
	__RectLinkLink__000000_Test_Middle_to_Top_Bottom_Link.Stroke = `lightgreen`
	__RectLinkLink__000000_Test_Middle_to_Top_Bottom_Link.StrokeOpacity = 1.000000
	__RectLinkLink__000000_Test_Middle_to_Top_Bottom_Link.StrokeWidth = 4.000000
	__RectLinkLink__000000_Test_Middle_to_Top_Bottom_Link.StrokeDashArray = ``
	__RectLinkLink__000000_Test_Middle_to_Top_Bottom_Link.StrokeDashArrayWhenSelected = ``
	__RectLinkLink__000000_Test_Middle_to_Top_Bottom_Link.Transform = ``

	__SVG__000000_SVG.Name = `SVG`
	__SVG__000000_SVG.DrawingState = models.NOT_DRAWING_LINK
	__SVG__000000_SVG.IsEditable = true
	__SVG__000000_SVG.IsSVGFileGenerated = false

	__SvgText__000000_Essai.Name = `Essai`
	__SvgText__000000_Essai.Text = `<svg xmlns="http://www.w3.org/2000/svg" width="1376" height="882" viewBox="-400 -10 1376 882">
  <g>
  <rect x="401" y="38" width="565" height="266.999985" rx="5" fill="bisque" fill-opacity="50" stroke="lightcoral" stroke-opacity="1" stroke-width="3" stroke-dasharray="" transform=""/>
  <rect class="click-through" x="401" y="38" width="565" height="50" rx="0" fill="lightblue" fill-opacity="100" stroke="" stroke-opacity="1" stroke-width="0"/>
  <g>
  <path class="click-through" d="M532-131q-6 5-12.5 8t-14.5 3q-8 0-16-3.5t-14-9.5q-41-44-60.5-90T395-320q0-37 11-78t38-106q23-57 32-87.5t9-56.5q0-34-15-63.5T423-771q-6-6-9.5-14t-3.5-16q0-8 3-14.5t8-12.5q6-6 13.5-9t15.5-3q8 0 15 3t13 8q44 41 65.5 86t21.5 95q0 35-10.5 73.5T518-474q-25 60-34 92t-9 61q0 35 14.5 67.5T534-188q5 6 8 13t3 15q0 8-3 15.5T532-131Zm195 0q-6 5-12.5 8t-14.5 3q-8 0-16-3.5t-14-9.5q-41-44-60.5-89.5T590-319q0-37 11-79t38-106q23-57 32-87t9-56q0-34-15-64.5T618-771q-6-6-9-13.5t-3-15.5q0-8 2.5-14.5T616-827q6-6 14-9.5t16-3.5q8 0 14.5 3t12.5 8q44 41 65.5 86t21.5 95q0 35-10.5 73.5T713-473q-25 60-34 92t-9 60q0 35 15 68.5t45 65.5q5 6 7.5 13t2.5 14q0 8-3 16t-10 13Zm-390 0q-6 5-12.5 8t-14.5 3q-8 0-16-3.5t-14-9.5q-41-44-60.5-89.5T200-319q0-37 11-79t38-106q23-57 32-87t9-56q0-34-15-64.5T228-771q-7-6-10-13.5t-3-15.5q0-8 3-15t8-13q6-6 13.5-9t15.5-3q8 0 15 3t13 8q44 41 65.5 85.5T370-648q0 35-10 73.5T324-474q-25 60-34 92t-9 61q0 35 14.5 68.5T340-187q5 6 7.5 13t2.5 14q0 8-3 16t-10 13Z" fill="" fill-opacity="0" stroke="blue" stroke-opacity="1" stroke-width="2" transform=" translate(401 304.999985) scale(0.2 0.2)"/>
</g>
<g>
<path class="click-through" d="M532-131q-6 5-12.5 8t-14.5 3q-8 0-16-3.5t-14-9.5q-41-44-60.5-90T395-320q0-37 11-78t38-106q23-57 32-87.5t9-56.5q0-34-15-63.5T423-771q-6-6-9.5-14t-3.5-16q0-8 3-14.5t8-12.5q6-6 13.5-9t15.5-3q8 0 15 3t13 8q44 41 65.5 86t21.5 95q0 35-10.5 73.5T518-474q-25 60-34 92t-9 61q0 35 14.5 67.5T534-188q5 6 8 13t3 15q0 8-3 15.5T532-131Zm195 0q-6 5-12.5 8t-14.5 3q-8 0-16-3.5t-14-9.5q-41-44-60.5-89.5T590-319q0-37 11-79t38-106q23-57 32-87t9-56q0-34-15-64.5T618-771q-6-6-9-13.5t-3-15.5q0-8 2.5-14.5T616-827q6-6 14-9.5t16-3.5q8 0 14.5 3t12.5 8q44 41 65.5 86t21.5 95q0 35-10.5 73.5T713-473q-25 60-34 92t-9 60q0 35 15 68.5t45 65.5q5 6 7.5 13t2.5 14q0 8-3 16t-10 13Zm-390 0q-6 5-12.5 8t-14.5 3q-8 0-16-3.5t-14-9.5q-41-44-60.5-89.5T200-319q0-37 11-79t38-106q23-57 32-87t9-56q0-34-15-64.5T228-771q-7-6-10-13.5t-3-15.5q0-8 3-15t8-13q6-6 13.5-9t15.5-3q8 0 15 3t13 8q44 41 65.5 85.5T370-648q0 35-10 73.5T324-474q-25 60-34 92t-9 61q0 35 14.5 68.5T340-187q5 6 7.5 13t2.5 14q0 8-3 16t-10 13Z" fill="" fill-opacity="0" stroke="blue" stroke-opacity="1" stroke-width="2" transform=" translate(401 304.999985) scale(0.2 0.2)"/>
</g>
<text class="click-through" x="683.5" y="58" fill="black" fill-opacity="100" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="" text-anchor="middle" font-weight="100" font-style="">
  <tspan x="683.5" dy="0" text-anchor="middle">This is an example of a note that</tspan>
  <tspan x="683.5" dy="1em" text-anchor="middle">could be displayed on a diagram.</tspan>
  <tspan x="683.5" dy="1em" text-anchor="middle"/>
  <tspan x="683.5" dy="1em" text-anchor="middle">It could explain one aspect of the model</tspan>
  <tspan x="683.5" dy="1em" text-anchor="middle">for instance, describing relations between structs</tspan>
  <tspan x="683.5" dy="1em" text-anchor="middle"/>
  <tspan x="683.5" dy="1em" text-anchor="middle">The text of a UML note refers a comment with the GONGNOTE keyword which is</tspan>
  <tspan x="683.5" dy="1em" text-anchor="middle">a special case of go Note convention. See example</tspan>
  <tspan x="683.5" dy="1em" text-anchor="middle">for details in the go code of the models.</tspan>
  <tspan x="683.5" dy="1em" text-anchor="middle"/>
</text>
<rect x="531.999969" y="625" width="237" height="237" rx="3" fill="lightcyan" fill-opacity="100" stroke="darkcyan" stroke-opacity="1" stroke-width="2" stroke-dasharray="" transform=""/>
<rect class="click-through" x="650.499969" y="665" width="100" height="30" rx="3" fill="lightgrey" fill-opacity="40" stroke="bisque" stroke-opacity="1" stroke-width="1"/>
<g>
<path class="click-through" d="M532-131q-6 5-12.5 8t-14.5 3q-8 0-16-3.5t-14-9.5q-41-44-60.5-90T395-320q0-37 11-78t38-106q23-57 32-87.5t9-56.5q0-34-15-63.5T423-771q-6-6-9.5-14t-3.5-16q0-8 3-14.5t8-12.5q6-6 13.5-9t15.5-3q8 0 15 3t13 8q44 41 65.5 86t21.5 95q0 35-10.5 73.5T518-474q-25 60-34 92t-9 61q0 35 14.5 67.5T534-188q5 6 8 13t3 15q0 8-3 15.5T532-131Zm195 0q-6 5-12.5 8t-14.5 3q-8 0-16-3.5t-14-9.5q-41-44-60.5-89.5T590-319q0-37 11-79t38-106q23-57 32-87t9-56q0-34-15-64.5T618-771q-6-6-9-13.5t-3-15.5q0-8 2.5-14.5T616-827q6-6 14-9.5t16-3.5q8 0 14.5 3t12.5 8q44 41 65.5 86t21.5 95q0 35-10.5 73.5T713-473q-25 60-34 92t-9 60q0 35 15 68.5t45 65.5q5 6 7.5 13t2.5 14q0 8-3 16t-10 13Zm-390 0q-6 5-12.5 8t-14.5 3q-8 0-16-3.5t-14-9.5q-41-44-60.5-89.5T200-319q0-37 11-79t38-106q23-57 32-87t9-56q0-34-15-64.5T228-771q-7-6-10-13.5t-3-15.5q0-8 3-15t8-13q6-6 13.5-9t15.5-3q8 0 15 3t13 8q44 41 65.5 85.5T370-648q0 35-10 73.5T324-474q-25 60-34 92t-9 61q0 35 14.5 68.5T340-187q5 6 7.5 13t2.5 14q0 8-3 16t-10 13Z" fill="black" fill-opacity="0.5" stroke="black" stroke-opacity="1" stroke-width="1" transform=" translate(531.999969 862) scale(0.246875 0.246875)"/>
</g>
<rect x="521" y="404" width="200" height="132" rx="3" fill="lavender" fill-opacity="50" stroke="turquoise" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform=""/>
<text class="click-through" x="621" y="424" fill="" fill-opacity="100" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="" text-anchor="start" font-weight="bold" font-style="italic">
  <tspan x="621" dy="0" text-anchor="start">Top anchored</tspan>
  <tspan x="621" dy="1em" text-anchor="start">top middle</tspan>
  <tspan x="621" dy="1em" text-anchor="start">line 3</tspan>
</text>
<g>
<line class="moveable-line" x1="531.999969" y1="800.431192" x2="170.99997499999995" y2="800.431192" fill="" fill-opacity="0" stroke="green" stroke-opacity="1" stroke-width="4" stroke-dasharray="" segment-orientation="horizontal" segment-number="0"/>
<line class="hit-area" x1="531.999969" y1="800.431192" x2="170.99997499999995" y2="800.431192" stroke="transparent" stroke-width="14" segment-orientation="horizontal"/>
</g>
<g>
<line class="moveable-line" x1="162.99997499999995" y1="792.431192" x2="162.99997499999995" y2="222.126273105265" fill="" fill-opacity="0" stroke="green" stroke-opacity="1" stroke-width="4" stroke-dasharray="" segment-orientation="vertical" segment-number="1"/>
<line class="hit-area" x1="162.99997499999995" y1="792.431192" x2="162.99997499999995" y2="222.126273105265" stroke="transparent" stroke-width="14" segment-orientation="vertical"/>
</g>
<g>
<line class="moveable-line" x1="170.99997499999995" y1="214.126273105265" x2="401" y2="214.126273105265" fill="" fill-opacity="0" stroke="green" stroke-opacity="1" stroke-width="4" stroke-dasharray="" segment-orientation="horizontal" segment-number="2"/>
<line class="hit-area" x1="170.99997499999995" y1="214.126273105265" x2="401" y2="214.126273105265" stroke="transparent" stroke-width="14" segment-orientation="horizontal"/>
</g>
<path d="M 170.99997499999995 800.431192 A 8 8 0 0 1 162.99997499999995 792.431192" fill="" fill-opacity="0" stroke="green" stroke-opacity="1" stroke-width="4" stroke-dasharray=""/>
<text class="click-through" x="402.99993799999993" y="837.443704" font-weight="light" font-size="" letter-spacing="" fill="cyan" fill-opacity="100" stroke="cyan" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="">
  <tspan x="402.99993799999993" dy="0">Start Anchored 1</tspan>
  <tspan x="402.99993799999993" dy="1em">Second line</tspan>
  <tspan x="402.99993799999993" dy="1em">Third Line</tspan>
</text>
<path d="M 162.99997499999995 222.126273105265 A 8 8 0 0 1 170.99997499999995 214.126273105265" fill="" fill-opacity="0" stroke="green" stroke-opacity="1" stroke-width="4" stroke-dasharray=""/>
<path d="M 402.414213562 215.540486667265 L 391 204.126273105265 M 402.414213562 212.712059543265 L 391 224.126273105265" fill="" fill-opacity="0" stroke="green" stroke-opacity="1" stroke-width="4" stroke-dasharray=""/>
<text class="click-through" x="340" y="179.126273105265" font-weight="100" font-size="16" letter-spacing="0.1em" fill="black" fill-opacity="100" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="">
  <tspan x="340" dy="0">Liine 1</tspan>
  <tspan x="340" dy="1em">Line 2</tspan>
</text>
<text class="click-through" x="340" y="179.126273105265" font-weight="100" font-size="16" letter-spacing="0.1em" fill="black" fill-opacity="100" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="">
  <tspan x="340" dy="0">Liine 1</tspan>
  <tspan x="340" dy="1em">Line 2</tspan>
</text>
<g>
<line class="moveable-line" x1="531.999969" y1="676.847542" x2="290.00002399999994" y2="676.847542" fill="" fill-opacity="0" stroke="red" stroke-opacity="0.7" stroke-width="5" stroke-dasharray="" segment-orientation="horizontal" segment-number="0"/>
<line class="hit-area" x1="531.999969" y1="676.847542" x2="290.00002399999994" y2="676.847542" stroke="transparent" stroke-width="15" segment-orientation="horizontal"/>
</g>
<g>
<line class="moveable-line" x1="270.00002399999994" y1="656.847542" x2="270.00002399999994" y2="487.109332" fill="" fill-opacity="0" stroke="red" stroke-opacity="0.7" stroke-width="5" stroke-dasharray="" segment-orientation="vertical" segment-number="1"/>
<line class="hit-area" x1="270.00002399999994" y1="656.847542" x2="270.00002399999994" y2="487.109332" stroke="transparent" stroke-width="15" segment-orientation="vertical"/>
</g>
<g>
<line class="moveable-line" x1="290.00002399999994" y1="467.109332" x2="521" y2="467.109332" fill="" fill-opacity="0" stroke="red" stroke-opacity="0.7" stroke-width="5" stroke-dasharray="" segment-orientation="horizontal" segment-number="2"/>
<line class="hit-area" x1="290.00002399999994" y1="467.109332" x2="521" y2="467.109332" stroke="transparent" stroke-width="15" segment-orientation="horizontal"/>
</g>
<path d="M 290.00002399999994 676.847542 A 20 20 0 0 1 270.00002399999994 656.847542" fill="" fill-opacity="0" stroke="red" stroke-opacity="0.7" stroke-width="5" stroke-dasharray=""/>
<path d="M 533.7677359525 678.6153089525 L 521.999969 666.847542 M 533.7677359525 675.0797750475 L 521.999969 686.847542" fill="" fill-opacity="0" stroke="red" stroke-opacity="0.7" stroke-width="5" stroke-dasharray=""/>
<text class="noevents-area" x="531.999969" y="659.247542" font-weight="normal" font-size="" letter-spacing="" fill="black" fill-opacity="0" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="">
  <tspan x="415.99996899999996" dy="0">Start Left/Top</tspan>
</text>
<text class="noevents-area" x="531.999969" y="705.847542" font-weight="normal" font-size="" letter-spacing="" fill="black" fill-opacity="1" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="">
  <tspan x="377.99996899999996" dy="0">Start Right/Bottom</tspan>
</text>
<path d="M 270.00002399999994 487.109332 A 20 20 0 0 1 290.00002399999994 467.109332" fill="" fill-opacity="0" stroke="red" stroke-opacity="0.7" stroke-width="5" stroke-dasharray=""/>
<path d="M 522.7677669525 468.8770989525 L 511 457.109332 M 522.7677669525 465.3415650475 L 511 477.109332" fill="" fill-opacity="0" stroke="red" stroke-opacity="0.7" stroke-width="5" stroke-dasharray=""/>
<text class="noevents-area" x="521" y="516.109332" font-weight="normal" font-size="" letter-spacing="" fill="black" fill-opacity="1" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="">
  <tspan x="373" dy="0">End Right/Bottom</tspan>
</text>
<text class="noevents-area" x="521" y="449.509332" font-weight="normal" font-size="" letter-spacing="" fill="blue" fill-opacity="0" stroke="blue" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="">
  <tspan x="371" dy="0">End Left/Top</tspan>
</text>
<line class="rect-link-link" x1="521" y1="466.3860751455546" x2="162.99997499999995" y2="453.448223263159" fill="" fill-opacity="0" stroke="lightgreen" stroke-opacity="1" stroke-width="4" stroke-dasharray="" transform=""/>
</g>
</svg>`

	__Text__000000_Essai.Name = `Essai`
	__Text__000000_Essai.X = 50.000000
	__Text__000000_Essai.Y = 300.000000
	__Text__000000_Essai.Content = `Hello`
	__Text__000000_Essai.Color = ``
	__Text__000000_Essai.FillOpacity = 0.000000
	__Text__000000_Essai.Stroke = `black`
	__Text__000000_Essai.StrokeOpacity = 1.000000
	__Text__000000_Essai.StrokeWidth = 2.000000
	__Text__000000_Essai.StrokeDashArray = ``
	__Text__000000_Essai.StrokeDashArrayWhenSelected = ``
	__Text__000000_Essai.Transform = ``

	// Setup of pointers
	// setup of Animate instances pointers
	// setup of Circle instances pointers
	// setup of Layer instances pointers
	__Layer__000000_Bottom_Rectangle_Layer.Rects = append(__Layer__000000_Bottom_Rectangle_Layer.Rects, __Rect__000000_Bottom)
	__Layer__000001_Layer_RectLinkLink_Medium_to_Top_Bottom.Links = append(__Layer__000001_Layer_RectLinkLink_Medium_to_Top_Bottom.Links, __Link__000001_Start_Middle)
	__Layer__000001_Layer_RectLinkLink_Medium_to_Top_Bottom.RectLinkLinks = append(__Layer__000001_Layer_RectLinkLink_Medium_to_Top_Bottom.RectLinkLinks, __RectLinkLink__000000_Test_Middle_to_Top_Bottom_Link)
	__Layer__000002_Line_layer.Lines = append(__Layer__000002_Line_layer.Lines, __Line__000000_Line_connecting_rect_Bottom_to_Top)
	__Layer__000003_Link_layer_vertical_to_horizontal.Links = append(__Layer__000003_Link_layer_vertical_to_horizontal.Links, __Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal)
	__Layer__000004_Middle_Rect_Layer.Rects = append(__Layer__000004_Middle_Rect_Layer.Rects, __Rect__000001_Middle_Rect)
	__Layer__000005_Top_Rectangle_layer.Rects = append(__Layer__000005_Top_Rectangle_layer.Rects, __Rect__000002_Top)
	// setup of Line instances pointers
	// setup of Link instances pointers
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.Start = __Rect__000002_Top
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.End = __Rect__000000_Bottom
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.TextAtArrowEnd = append(__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.TextAtArrowEnd, __LinkAnchoredText__000002_Liine_1_Line_2)
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.TextAtArrowEnd = append(__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.TextAtArrowEnd, __LinkAnchoredText__000002_Liine_1_Line_2)
	__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.TextAtArrowStart = append(__Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal.TextAtArrowStart, __LinkAnchoredText__000003_Start_Anchored_1)
	__Link__000001_Start_Middle.Start = __Rect__000002_Top
	__Link__000001_Start_Middle.End = __Rect__000001_Middle_Rect
	__Link__000001_Start_Middle.TextAtArrowEnd = append(__Link__000001_Start_Middle.TextAtArrowEnd, __LinkAnchoredText__000001_End_Right_Bottom)
	__Link__000001_Start_Middle.TextAtArrowEnd = append(__Link__000001_Start_Middle.TextAtArrowEnd, __LinkAnchoredText__000000_End_Left_Top)
	__Link__000001_Start_Middle.TextAtArrowStart = append(__Link__000001_Start_Middle.TextAtArrowStart, __LinkAnchoredText__000004_Start_Left_Top)
	__Link__000001_Start_Middle.TextAtArrowStart = append(__Link__000001_Start_Middle.TextAtArrowStart, __LinkAnchoredText__000005_Start_Right_Bottom)
	// setup of LinkAnchoredText instances pointers
	// setup of Rect instances pointers
	__Rect__000000_Bottom.RectAnchoredTexts = append(__Rect__000000_Bottom.RectAnchoredTexts, __RectAnchoredText__000000_Bottom_Text)
	__Rect__000000_Bottom.RectAnchoredRects = append(__Rect__000000_Bottom.RectAnchoredRects, __RectAnchoredRect__000001_Top_on_Bottom_with_same_width)
	__Rect__000000_Bottom.RectAnchoredPaths = append(__Rect__000000_Bottom.RectAnchoredPaths, __RectAnchoredPath__000001_second_log)
	__Rect__000000_Bottom.RectAnchoredPaths = append(__Rect__000000_Bottom.RectAnchoredPaths, __RectAnchoredPath__000001_second_log)
	__Rect__000001_Middle_Rect.RectAnchoredTexts = append(__Rect__000001_Middle_Rect.RectAnchoredTexts, __RectAnchoredText__000003_Top_anchored_top_middle)
	__Rect__000002_Top.RectAnchoredRects = append(__Rect__000002_Top.RectAnchoredRects, __RectAnchoredRect__000000_Rect_within_top)
	__Rect__000002_Top.RectAnchoredPaths = append(__Rect__000002_Top.RectAnchoredPaths, __RectAnchoredPath__000000_Logo_to_add)
	// setup of RectAnchoredPath instances pointers
	// setup of RectAnchoredRect instances pointers
	// setup of RectAnchoredText instances pointers
	// setup of RectLinkLink instances pointers
	__RectLinkLink__000000_Test_Middle_to_Top_Bottom_Link.Start = __Rect__000001_Middle_Rect
	__RectLinkLink__000000_Test_Middle_to_Top_Bottom_Link.End = __Link__000000_Arrow_Top_to_Bottom_vertical_to_horizontal
	// setup of SVG instances pointers
	__SVG__000000_SVG.Layers = append(__SVG__000000_SVG.Layers, __Layer__000000_Bottom_Rectangle_Layer)
	__SVG__000000_SVG.Layers = append(__SVG__000000_SVG.Layers, __Layer__000005_Top_Rectangle_layer)
	__SVG__000000_SVG.Layers = append(__SVG__000000_SVG.Layers, __Layer__000004_Middle_Rect_Layer)
	__SVG__000000_SVG.Layers = append(__SVG__000000_SVG.Layers, __Layer__000003_Link_layer_vertical_to_horizontal)
	__SVG__000000_SVG.Layers = append(__SVG__000000_SVG.Layers, __Layer__000001_Layer_RectLinkLink_Medium_to_Top_Bottom)
	__SVG__000000_SVG.StartRect = __Rect__000000_Bottom
	__SVG__000000_SVG.EndRect = __Rect__000002_Top
	// setup of SvgText instances pointers
	// setup of Text instances pointers
}
