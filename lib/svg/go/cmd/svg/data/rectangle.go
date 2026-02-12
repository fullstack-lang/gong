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

	// insertion point for declaration of instances to stage

	__Animate__00000000_ := (&models.Animate{Name: ``}).Stage(stage)

	__Circle__00000000_ := (&models.Circle{Name: `Test`}).Stage(stage)

	__ControlPoint__00000000_ := (&models.ControlPoint{Name: `Middle Link Control Point`}).Stage(stage)

	__Layer__00000000_ := (&models.Layer{Name: `Bottom Rectangle Layer`}).Stage(stage)
	__Layer__00000001_ := (&models.Layer{Name: `Layer RectLinkLink Medium to Top-Bottom`}).Stage(stage)
	__Layer__00000002_ := (&models.Layer{Name: `Line layer`}).Stage(stage)
	__Layer__00000003_ := (&models.Layer{Name: `Link layer vertical to horizontal`}).Stage(stage)
	__Layer__00000004_ := (&models.Layer{Name: `Middle Rect Layer`}).Stage(stage)
	__Layer__00000005_ := (&models.Layer{Name: `Top Rectangle layer`}).Stage(stage)
	__Layer__00000006_ := (&models.Layer{Name: `Test Automatic Layout`}).Stage(stage)

	__Link__00000000_ := (&models.Link{Name: `Arrow - Top to Bottom vertical to horizontal`}).Stage(stage)
	__Link__00000001_ := (&models.Link{Name: `Start - Middle`}).Stage(stage)
	__Link__00000002_ := (&models.Link{Name: `Top to Middle (issue #645)`}).Stage(stage)
	__Link__00000003_ := (&models.Link{Name: `Auto 1`}).Stage(stage)
	__Link__00000004_ := (&models.Link{Name: `Auto 2`}).Stage(stage)
	__Link__00000005_ := (&models.Link{Name: `Auto 3`}).Stage(stage)
	__Link__00000006_ := (&models.Link{Name: `middle to bottom`}).Stage(stage)

	__LinkAnchoredText__00000000_ := (&models.LinkAnchoredText{Name: `End Left/Top`}).Stage(stage)
	__LinkAnchoredText__00000001_ := (&models.LinkAnchoredText{Name: `End Right/Bottom`}).Stage(stage)
	__LinkAnchoredText__00000002_ := (&models.LinkAnchoredText{Name: `Liine 1
Line 2`}).Stage(stage)
	__LinkAnchoredText__00000003_ := (&models.LinkAnchoredText{Name: `Start Anchored 1`}).Stage(stage)
	__LinkAnchoredText__00000004_ := (&models.LinkAnchoredText{Name: `Start Left/Top`}).Stage(stage)
	__LinkAnchoredText__00000005_ := (&models.LinkAnchoredText{Name: `Start Right/Bottom`}).Stage(stage)
	__LinkAnchoredText__00000006_ := (&models.LinkAnchoredText{Name: `Issue #645`}).Stage(stage)
	__LinkAnchoredText__00000007_ := (&models.LinkAnchoredText{Name: `Right to the end anchored text`}).Stage(stage)
	__LinkAnchoredText__00000008_ := (&models.LinkAnchoredText{Name: `AA_RIGHT_OR_BOTTOM`}).Stage(stage)
	__LinkAnchoredText__00000009_ := (&models.LinkAnchoredText{Name: `AA_LEFT_OR_TOP`}).Stage(stage)
	__LinkAnchoredText__00000010_ := (&models.LinkAnchoredText{Name: `AA_RIGHT_OR_BOTTOM_BELOW`}).Stage(stage)
	__LinkAnchoredText__00000011_ := (&models.LinkAnchoredText{Name: `LEFT_OR_TOP_ABOVE`}).Stage(stage)

	__Rect__00000000_ := (&models.Rect{Name: `Bottom Rect`}).Stage(stage)
	__Rect__00000001_ := (&models.Rect{Name: `Middle Rect`}).Stage(stage)
	__Rect__00000002_ := (&models.Rect{Name: `Top Rect`}).Stage(stage)
	__Rect__00000003_ := (&models.Rect{Name: `Start`}).Stage(stage)
	__Rect__00000004_ := (&models.Rect{Name: `End`}).Stage(stage)

	__RectAnchoredPath__00000000_ := (&models.RectAnchoredPath{Name: `Logo to add`}).Stage(stage)
	__RectAnchoredPath__00000001_ := (&models.RectAnchoredPath{Name: `second log`}).Stage(stage)

	__RectAnchoredRect__00000000_ := (&models.RectAnchoredRect{Name: `Rect within top`}).Stage(stage)
	__RectAnchoredRect__00000001_ := (&models.RectAnchoredRect{Name: `Top on Bottom with same width`}).Stage(stage)

	__RectAnchoredText__00000000_ := (&models.RectAnchoredText{Name: `Bottom Text`}).Stage(stage)
	__RectAnchoredText__00000001_ := (&models.RectAnchoredText{Name: `Top Left`}).Stage(stage)
	__RectAnchoredText__00000002_ := (&models.RectAnchoredText{Name: `Top anchored bottom middle`}).Stage(stage)
	__RectAnchoredText__00000003_ := (&models.RectAnchoredText{Name: `Top anchored top middle`}).Stage(stage)
	__RectAnchoredText__00000004_ := (&models.RectAnchoredText{Name: `Start`}).Stage(stage)
	__RectAnchoredText__00000005_ := (&models.RectAnchoredText{Name: `End`}).Stage(stage)
	__RectAnchoredText__00000006_ := (&models.RectAnchoredText{Name: ``}).Stage(stage)

	__SVG__00000000_ := (&models.SVG{Name: `SVG`}).Stage(stage)

	__SvgText__00000000_ := (&models.SvgText{Name: `Essai`}).Stage(stage)

	__Text__00000000_ := (&models.Text{Name: `Essai`}).Stage(stage)

	// insertion point for initialization of values

	__Animate__00000000_.Name = ``
	__Animate__00000000_.AttributeName = ``
	__Animate__00000000_.Values = ``
	__Animate__00000000_.From = ``
	__Animate__00000000_.To = ``
	__Animate__00000000_.Dur = ``
	__Animate__00000000_.RepeatCount = ``

	__Circle__00000000_.Name = `Test`
	__Circle__00000000_.CX = 400.000000
	__Circle__00000000_.CY = 300.000000
	__Circle__00000000_.Radius = 100.000000
	__Circle__00000000_.Color = `lavender`
	__Circle__00000000_.FillOpacity = 50.000000
	__Circle__00000000_.Stroke = ``
	__Circle__00000000_.StrokeOpacity = 1.000000
	__Circle__00000000_.StrokeWidth = 0.000000
	__Circle__00000000_.StrokeDashArray = ``
	__Circle__00000000_.StrokeDashArrayWhenSelected = ``
	__Circle__00000000_.Transform = ``

	__ControlPoint__00000000_.Name = `Middle Link Control Point`
	__ControlPoint__00000000_.X_Relative = 0.204991
	__ControlPoint__00000000_.Y_Relative = -0.545447

	__Layer__00000000_.Name = `Bottom Rectangle Layer`

	__Layer__00000001_.Name = `Layer RectLinkLink Medium to Top-Bottom`

	__Layer__00000002_.Name = `Line layer`

	__Layer__00000003_.Name = `Link layer vertical to horizontal`

	__Layer__00000004_.Name = `Middle Rect Layer`

	__Layer__00000005_.Name = `Top Rectangle layer`

	__Layer__00000006_.Name = `Test Automatic Layout`

	__Link__00000000_.Name = `Arrow - Top to Bottom vertical to horizontal`
	__Link__00000000_.Type = models.LINK_TYPE_FLOATING_ORTHOGONAL
	__Link__00000000_.IsBezierCurve = false
	__Link__00000000_.StartAnchorType = models.ANCHOR_CENTER
	__Link__00000000_.EndAnchorType = models.ANCHOR_CENTER
	__Link__00000000_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__00000000_.StartRatio = 0.740216
	__Link__00000000_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__00000000_.EndRatio = 0.659649
	__Link__00000000_.CornerOffsetRatio = -1.556962
	__Link__00000000_.CornerRadius = 8.000000
	__Link__00000000_.HasEndArrow = true
	__Link__00000000_.EndArrowSize = 10.000000
	__Link__00000000_.EndArrowOffset = 0.000000
	__Link__00000000_.HasStartArrow = false
	__Link__00000000_.StartArrowSize = 0.000000
	__Link__00000000_.StartArrowOffset = 0.000000
	__Link__00000000_.Color = ``
	__Link__00000000_.FillOpacity = 0.000000
	__Link__00000000_.Stroke = `green`
	__Link__00000000_.StrokeOpacity = 1.000000
	__Link__00000000_.StrokeWidth = 4.000000
	__Link__00000000_.StrokeDashArray = ``
	__Link__00000000_.StrokeDashArrayWhenSelected = ``
	__Link__00000000_.Transform = ``
	__Link__00000000_.MouseX = 0.000000
	__Link__00000000_.MouseY = 0.000000
	__Link__00000000_.MouseEventKey = ""

	__Link__00000001_.Name = `Start - Middle`
	__Link__00000001_.Type = models.LINK_TYPE_FLOATING_ORTHOGONAL
	__Link__00000001_.IsBezierCurve = false
	__Link__00000001_.StartAnchorType = ""
	__Link__00000001_.EndAnchorType = ""
	__Link__00000001_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__00000001_.StartRatio = 0.242616
	__Link__00000001_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__00000001_.EndRatio = 0.000000
	__Link__00000001_.CornerOffsetRatio = -1.257384
	__Link__00000001_.CornerRadius = 20.000000
	__Link__00000001_.HasEndArrow = true
	__Link__00000001_.EndArrowSize = 10.000000
	__Link__00000001_.EndArrowOffset = 0.000000
	__Link__00000001_.HasStartArrow = true
	__Link__00000001_.StartArrowSize = 10.000000
	__Link__00000001_.StartArrowOffset = 0.000000
	__Link__00000001_.Color = ``
	__Link__00000001_.FillOpacity = 0.000000
	__Link__00000001_.Stroke = `red`
	__Link__00000001_.StrokeOpacity = 0.700000
	__Link__00000001_.StrokeWidth = 5.000000
	__Link__00000001_.StrokeDashArray = ``
	__Link__00000001_.StrokeDashArrayWhenSelected = ``
	__Link__00000001_.Transform = ``
	__Link__00000001_.MouseX = 697.000000
	__Link__00000001_.MouseY = 200.001953
	__Link__00000001_.MouseEventKey = ""

	__Link__00000002_.Name = `Top to Middle (issue #645)`
	__Link__00000002_.Type = models.LINK_TYPE_LINE_WITH_CONTROL_POINTS
	__Link__00000002_.IsBezierCurve = false
	__Link__00000002_.StartAnchorType = models.ANCHOR_CENTER
	__Link__00000002_.EndAnchorType = models.ANCHOR_CENTER
	__Link__00000002_.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__00000002_.StartRatio = 0.421941
	__Link__00000002_.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__00000002_.EndRatio = 0.575000
	__Link__00000002_.CornerOffsetRatio = 1.126591
	__Link__00000002_.CornerRadius = 10.000000
	__Link__00000002_.HasEndArrow = true
	__Link__00000002_.EndArrowSize = 11.000000
	__Link__00000002_.EndArrowOffset = 6.000000
	__Link__00000002_.HasStartArrow = true
	__Link__00000002_.StartArrowSize = 6.000000
	__Link__00000002_.StartArrowOffset = 10.000000
	__Link__00000002_.Color = ``
	__Link__00000002_.FillOpacity = 0.000000
	__Link__00000002_.Stroke = `black`
	__Link__00000002_.StrokeOpacity = 1.000000
	__Link__00000002_.StrokeWidth = 2.000000
	__Link__00000002_.StrokeDashArray = ``
	__Link__00000002_.StrokeDashArrayWhenSelected = ``
	__Link__00000002_.Transform = ``
	__Link__00000002_.MouseX = 638.000000
	__Link__00000002_.MouseY = 384.001953
	__Link__00000002_.MouseEventKey = models.MouseEventKeyMeta

	__Link__00000003_.Name = `Auto 1`
	__Link__00000003_.Type = models.LINK_TYPE_FLOATING_ORTHOGONAL
	__Link__00000003_.IsBezierCurve = false
	__Link__00000003_.StartAnchorType = ""
	__Link__00000003_.EndAnchorType = ""
	__Link__00000003_.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__00000003_.StartRatio = 0.388824
	__Link__00000003_.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__00000003_.EndRatio = 0.626972
	__Link__00000003_.CornerOffsetRatio = 1.956431
	__Link__00000003_.CornerRadius = 20.000000
	__Link__00000003_.HasEndArrow = true
	__Link__00000003_.EndArrowSize = 15.000000
	__Link__00000003_.EndArrowOffset = 0.000000
	__Link__00000003_.HasStartArrow = true
	__Link__00000003_.StartArrowSize = 10.000000
	__Link__00000003_.StartArrowOffset = 0.000000
	__Link__00000003_.Color = ``
	__Link__00000003_.FillOpacity = 0.000000
	__Link__00000003_.Stroke = `black`
	__Link__00000003_.StrokeOpacity = 1.000000
	__Link__00000003_.StrokeWidth = 1.000000
	__Link__00000003_.StrokeDashArray = ``
	__Link__00000003_.StrokeDashArrayWhenSelected = ``
	__Link__00000003_.Transform = ``
	__Link__00000003_.MouseX = 0.000000
	__Link__00000003_.MouseY = 0.000000
	__Link__00000003_.MouseEventKey = ""

	__Link__00000004_.Name = `Auto 2`
	__Link__00000004_.Type = models.LINK_TYPE_FLOATING_ORTHOGONAL
	__Link__00000004_.IsBezierCurve = false
	__Link__00000004_.StartAnchorType = ""
	__Link__00000004_.EndAnchorType = ""
	__Link__00000004_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__00000004_.StartRatio = 0.647783
	__Link__00000004_.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__00000004_.EndRatio = 0.909340
	__Link__00000004_.CornerOffsetRatio = 1.030819
	__Link__00000004_.CornerRadius = 40.000000
	__Link__00000004_.HasEndArrow = true
	__Link__00000004_.EndArrowSize = 15.000000
	__Link__00000004_.EndArrowOffset = 0.000000
	__Link__00000004_.HasStartArrow = true
	__Link__00000004_.StartArrowSize = 10.000000
	__Link__00000004_.StartArrowOffset = 0.000000
	__Link__00000004_.Color = ``
	__Link__00000004_.FillOpacity = 0.000000
	__Link__00000004_.Stroke = `black`
	__Link__00000004_.StrokeOpacity = 1.000000
	__Link__00000004_.StrokeWidth = 1.000000
	__Link__00000004_.StrokeDashArray = ``
	__Link__00000004_.StrokeDashArrayWhenSelected = ``
	__Link__00000004_.Transform = ``
	__Link__00000004_.MouseX = 0.000000
	__Link__00000004_.MouseY = 0.000000
	__Link__00000004_.MouseEventKey = ""

	__Link__00000005_.Name = `Auto 3`
	__Link__00000005_.Type = models.LINK_TYPE_FLOATING_ORTHOGONAL
	__Link__00000005_.IsBezierCurve = false
	__Link__00000005_.StartAnchorType = ""
	__Link__00000005_.EndAnchorType = ""
	__Link__00000005_.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__00000005_.StartRatio = 0.620556
	__Link__00000005_.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__00000005_.EndRatio = 0.408200
	__Link__00000005_.CornerOffsetRatio = -0.383744
	__Link__00000005_.CornerRadius = 30.000000
	__Link__00000005_.HasEndArrow = true
	__Link__00000005_.EndArrowSize = 15.000000
	__Link__00000005_.EndArrowOffset = 0.000000
	__Link__00000005_.HasStartArrow = true
	__Link__00000005_.StartArrowSize = 10.000000
	__Link__00000005_.StartArrowOffset = 0.000000
	__Link__00000005_.Color = ``
	__Link__00000005_.FillOpacity = 0.000000
	__Link__00000005_.Stroke = `black`
	__Link__00000005_.StrokeOpacity = 1.000000
	__Link__00000005_.StrokeWidth = 1.000000
	__Link__00000005_.StrokeDashArray = ``
	__Link__00000005_.StrokeDashArrayWhenSelected = ``
	__Link__00000005_.Transform = ``
	__Link__00000005_.MouseX = 0.000000
	__Link__00000005_.MouseY = 0.000000
	__Link__00000005_.MouseEventKey = ""

	__Link__00000006_.Name = `middle to bottom`
	__Link__00000006_.Type = models.LINK_TYPE_LINE_WITH_CONTROL_POINTS
	__Link__00000006_.IsBezierCurve = false
	__Link__00000006_.StartAnchorType = models.ANCHOR_BOTTOM
	__Link__00000006_.EndAnchorType = models.ANCHOR_TOP
	__Link__00000006_.StartOrientation = ""
	__Link__00000006_.StartRatio = 0.000000
	__Link__00000006_.EndOrientation = ""
	__Link__00000006_.EndRatio = 0.000000
	__Link__00000006_.CornerOffsetRatio = 0.000000
	__Link__00000006_.CornerRadius = 0.000000
	__Link__00000006_.HasEndArrow = true
	__Link__00000006_.EndArrowSize = 10.000000
	__Link__00000006_.EndArrowOffset = 3.000000
	__Link__00000006_.HasStartArrow = false
	__Link__00000006_.StartArrowSize = 0.000000
	__Link__00000006_.StartArrowOffset = 0.000000
	__Link__00000006_.Color = ``
	__Link__00000006_.FillOpacity = 0.000000
	__Link__00000006_.Stroke = `blue`
	__Link__00000006_.StrokeOpacity = 1.000000
	__Link__00000006_.StrokeWidth = 1.000000
	__Link__00000006_.StrokeDashArray = ``
	__Link__00000006_.StrokeDashArrayWhenSelected = ``
	__Link__00000006_.Transform = ``
	__Link__00000006_.MouseX = 0.000000
	__Link__00000006_.MouseY = 0.000000
	__Link__00000006_.MouseEventKey = ""

	__LinkAnchoredText__00000000_.Name = `End Left/Top`
	__LinkAnchoredText__00000000_.Content = `End Left/Top`
	__LinkAnchoredText__00000000_.AutomaticLayout = true
	__LinkAnchoredText__00000000_.LinkAnchorType = models.LINK_LEFT_OR_TOP
	__LinkAnchoredText__00000000_.X_Offset = -40.000000
	__LinkAnchoredText__00000000_.Y_Offset = 0.000000
	__LinkAnchoredText__00000000_.FontWeight = `normal`
	__LinkAnchoredText__00000000_.FontSize = ``
	__LinkAnchoredText__00000000_.FontStyle = ``
	__LinkAnchoredText__00000000_.LetterSpacing = ``
	__LinkAnchoredText__00000000_.FontFamily = ``
	__LinkAnchoredText__00000000_.WhiteSpace = ""
	__LinkAnchoredText__00000000_.Color = `blue`
	__LinkAnchoredText__00000000_.FillOpacity = 0.000000
	__LinkAnchoredText__00000000_.Stroke = `blue`
	__LinkAnchoredText__00000000_.StrokeOpacity = 1.000000
	__LinkAnchoredText__00000000_.StrokeWidth = 1.000000
	__LinkAnchoredText__00000000_.StrokeDashArray = ``
	__LinkAnchoredText__00000000_.StrokeDashArrayWhenSelected = ``
	__LinkAnchoredText__00000000_.Transform = ``

	__LinkAnchoredText__00000001_.Name = `End Right/Bottom`
	__LinkAnchoredText__00000001_.Content = `End Right/Bottom`
	__LinkAnchoredText__00000001_.AutomaticLayout = true
	__LinkAnchoredText__00000001_.LinkAnchorType = models.LINK_RIGHT_OR_BOTTOM
	__LinkAnchoredText__00000001_.X_Offset = 0.000000
	__LinkAnchoredText__00000001_.Y_Offset = 20.000000
	__LinkAnchoredText__00000001_.FontWeight = `normal`
	__LinkAnchoredText__00000001_.FontSize = ``
	__LinkAnchoredText__00000001_.FontStyle = ``
	__LinkAnchoredText__00000001_.LetterSpacing = ``
	__LinkAnchoredText__00000001_.FontFamily = ``
	__LinkAnchoredText__00000001_.WhiteSpace = ""
	__LinkAnchoredText__00000001_.Color = `black`
	__LinkAnchoredText__00000001_.FillOpacity = 1.000000
	__LinkAnchoredText__00000001_.Stroke = `black`
	__LinkAnchoredText__00000001_.StrokeOpacity = 1.000000
	__LinkAnchoredText__00000001_.StrokeWidth = 1.000000
	__LinkAnchoredText__00000001_.StrokeDashArray = ``
	__LinkAnchoredText__00000001_.StrokeDashArrayWhenSelected = ``
	__LinkAnchoredText__00000001_.Transform = ``

	__LinkAnchoredText__00000002_.Name = `Liine 1
Line 2`
	__LinkAnchoredText__00000002_.Content = `Liine 1`
	__LinkAnchoredText__00000002_.AutomaticLayout = false
	__LinkAnchoredText__00000002_.LinkAnchorType = ""
	__LinkAnchoredText__00000002_.X_Offset = -61.000000
	__LinkAnchoredText__00000002_.Y_Offset = -35.000000
	__LinkAnchoredText__00000002_.FontWeight = `100`
	__LinkAnchoredText__00000002_.FontSize = `16`
	__LinkAnchoredText__00000002_.FontStyle = ``
	__LinkAnchoredText__00000002_.LetterSpacing = `0.1em`
	__LinkAnchoredText__00000002_.FontFamily = ``
	__LinkAnchoredText__00000002_.WhiteSpace = ""
	__LinkAnchoredText__00000002_.Color = `black`
	__LinkAnchoredText__00000002_.FillOpacity = 100.000000
	__LinkAnchoredText__00000002_.Stroke = `black`
	__LinkAnchoredText__00000002_.StrokeOpacity = 1.000000
	__LinkAnchoredText__00000002_.StrokeWidth = 1.000000
	__LinkAnchoredText__00000002_.StrokeDashArray = ``
	__LinkAnchoredText__00000002_.StrokeDashArrayWhenSelected = ``
	__LinkAnchoredText__00000002_.Transform = ``

	__LinkAnchoredText__00000003_.Name = `Start Anchored 1`
	__LinkAnchoredText__00000003_.Content = `Start Anchored 1`
	__LinkAnchoredText__00000003_.AutomaticLayout = false
	__LinkAnchoredText__00000003_.LinkAnchorType = ""
	__LinkAnchoredText__00000003_.X_Offset = -129.000031
	__LinkAnchoredText__00000003_.Y_Offset = 37.012512
	__LinkAnchoredText__00000003_.FontWeight = `light`
	__LinkAnchoredText__00000003_.FontSize = ``
	__LinkAnchoredText__00000003_.FontStyle = ``
	__LinkAnchoredText__00000003_.LetterSpacing = ``
	__LinkAnchoredText__00000003_.FontFamily = ``
	__LinkAnchoredText__00000003_.WhiteSpace = ""
	__LinkAnchoredText__00000003_.Color = `cyan`
	__LinkAnchoredText__00000003_.FillOpacity = 100.000000
	__LinkAnchoredText__00000003_.Stroke = `cyan`
	__LinkAnchoredText__00000003_.StrokeOpacity = 1.000000
	__LinkAnchoredText__00000003_.StrokeWidth = 1.000000
	__LinkAnchoredText__00000003_.StrokeDashArray = ``
	__LinkAnchoredText__00000003_.StrokeDashArrayWhenSelected = ``
	__LinkAnchoredText__00000003_.Transform = ``

	__LinkAnchoredText__00000004_.Name = `Start Left/Top`
	__LinkAnchoredText__00000004_.Content = `Start Left/Top`
	__LinkAnchoredText__00000004_.AutomaticLayout = true
	__LinkAnchoredText__00000004_.LinkAnchorType = models.LINK_LEFT_OR_TOP
	__LinkAnchoredText__00000004_.X_Offset = 0.000000
	__LinkAnchoredText__00000004_.Y_Offset = 0.000000
	__LinkAnchoredText__00000004_.FontWeight = `normal`
	__LinkAnchoredText__00000004_.FontSize = ``
	__LinkAnchoredText__00000004_.FontStyle = ``
	__LinkAnchoredText__00000004_.LetterSpacing = ``
	__LinkAnchoredText__00000004_.FontFamily = ``
	__LinkAnchoredText__00000004_.WhiteSpace = ""
	__LinkAnchoredText__00000004_.Color = `black`
	__LinkAnchoredText__00000004_.FillOpacity = 0.000000
	__LinkAnchoredText__00000004_.Stroke = `black`
	__LinkAnchoredText__00000004_.StrokeOpacity = 1.000000
	__LinkAnchoredText__00000004_.StrokeWidth = 1.000000
	__LinkAnchoredText__00000004_.StrokeDashArray = ``
	__LinkAnchoredText__00000004_.StrokeDashArrayWhenSelected = ``
	__LinkAnchoredText__00000004_.Transform = ``

	__LinkAnchoredText__00000005_.Name = `Start Right/Bottom`
	__LinkAnchoredText__00000005_.Content = `Start Right/Bottom`
	__LinkAnchoredText__00000005_.AutomaticLayout = true
	__LinkAnchoredText__00000005_.LinkAnchorType = models.LINK_RIGHT_OR_BOTTOM
	__LinkAnchoredText__00000005_.X_Offset = 0.000000
	__LinkAnchoredText__00000005_.Y_Offset = 0.000000
	__LinkAnchoredText__00000005_.FontWeight = `normal`
	__LinkAnchoredText__00000005_.FontSize = ``
	__LinkAnchoredText__00000005_.FontStyle = ``
	__LinkAnchoredText__00000005_.LetterSpacing = ``
	__LinkAnchoredText__00000005_.FontFamily = ``
	__LinkAnchoredText__00000005_.WhiteSpace = ""
	__LinkAnchoredText__00000005_.Color = `black`
	__LinkAnchoredText__00000005_.FillOpacity = 1.000000
	__LinkAnchoredText__00000005_.Stroke = `black`
	__LinkAnchoredText__00000005_.StrokeOpacity = 1.000000
	__LinkAnchoredText__00000005_.StrokeWidth = 1.000000
	__LinkAnchoredText__00000005_.StrokeDashArray = ``
	__LinkAnchoredText__00000005_.StrokeDashArrayWhenSelected = ``
	__LinkAnchoredText__00000005_.Transform = ``

	__LinkAnchoredText__00000006_.Name = `Issue #645`
	__LinkAnchoredText__00000006_.Content = `varylongverylongnamexxxxxxxxxxxx`
	__LinkAnchoredText__00000006_.AutomaticLayout = true
	__LinkAnchoredText__00000006_.LinkAnchorType = models.LINK_LEFT_OR_TOP
	__LinkAnchoredText__00000006_.X_Offset = 0.000000
	__LinkAnchoredText__00000006_.Y_Offset = 0.000000
	__LinkAnchoredText__00000006_.FontWeight = ``
	__LinkAnchoredText__00000006_.FontSize = ``
	__LinkAnchoredText__00000006_.FontStyle = ``
	__LinkAnchoredText__00000006_.LetterSpacing = ``
	__LinkAnchoredText__00000006_.FontFamily = ``
	__LinkAnchoredText__00000006_.WhiteSpace = ""
	__LinkAnchoredText__00000006_.Color = `blue`
	__LinkAnchoredText__00000006_.FillOpacity = 1.000000
	__LinkAnchoredText__00000006_.Stroke = `blue`
	__LinkAnchoredText__00000006_.StrokeOpacity = 1.000000
	__LinkAnchoredText__00000006_.StrokeWidth = 1.000000
	__LinkAnchoredText__00000006_.StrokeDashArray = ``
	__LinkAnchoredText__00000006_.StrokeDashArrayWhenSelected = ``
	__LinkAnchoredText__00000006_.Transform = ``

	__LinkAnchoredText__00000007_.Name = `Right to the end anchored text`
	__LinkAnchoredText__00000007_.Content = `Right to the end anchored text`
	__LinkAnchoredText__00000007_.AutomaticLayout = true
	__LinkAnchoredText__00000007_.LinkAnchorType = models.LINK_RIGHT_OR_BOTTOM
	__LinkAnchoredText__00000007_.X_Offset = 0.000000
	__LinkAnchoredText__00000007_.Y_Offset = 0.000000
	__LinkAnchoredText__00000007_.FontWeight = ``
	__LinkAnchoredText__00000007_.FontSize = ``
	__LinkAnchoredText__00000007_.FontStyle = ``
	__LinkAnchoredText__00000007_.LetterSpacing = ``
	__LinkAnchoredText__00000007_.FontFamily = ``
	__LinkAnchoredText__00000007_.WhiteSpace = ""
	__LinkAnchoredText__00000007_.Color = `blue`
	__LinkAnchoredText__00000007_.FillOpacity = 1.000000
	__LinkAnchoredText__00000007_.Stroke = `blue`
	__LinkAnchoredText__00000007_.StrokeOpacity = 1.000000
	__LinkAnchoredText__00000007_.StrokeWidth = 1.000000
	__LinkAnchoredText__00000007_.StrokeDashArray = ``
	__LinkAnchoredText__00000007_.StrokeDashArrayWhenSelected = ``
	__LinkAnchoredText__00000007_.Transform = ``

	__LinkAnchoredText__00000008_.Name = `AA_RIGHT_OR_BOTTOM`
	__LinkAnchoredText__00000008_.Content = `RIGHT_OR_BOTTOM`
	__LinkAnchoredText__00000008_.AutomaticLayout = true
	__LinkAnchoredText__00000008_.LinkAnchorType = models.LINK_RIGHT_OR_BOTTOM
	__LinkAnchoredText__00000008_.X_Offset = 0.000000
	__LinkAnchoredText__00000008_.Y_Offset = 0.000000
	__LinkAnchoredText__00000008_.FontWeight = ``
	__LinkAnchoredText__00000008_.FontSize = `16`
	__LinkAnchoredText__00000008_.FontStyle = ``
	__LinkAnchoredText__00000008_.LetterSpacing = ``
	__LinkAnchoredText__00000008_.FontFamily = ``
	__LinkAnchoredText__00000008_.WhiteSpace = ""
	__LinkAnchoredText__00000008_.Color = `black`
	__LinkAnchoredText__00000008_.FillOpacity = 1.000000
	__LinkAnchoredText__00000008_.Stroke = `black`
	__LinkAnchoredText__00000008_.StrokeOpacity = 1.000000
	__LinkAnchoredText__00000008_.StrokeWidth = 1.000000
	__LinkAnchoredText__00000008_.StrokeDashArray = ``
	__LinkAnchoredText__00000008_.StrokeDashArrayWhenSelected = ``
	__LinkAnchoredText__00000008_.Transform = ``

	__LinkAnchoredText__00000009_.Name = `AA_LEFT_OR_TOP`
	__LinkAnchoredText__00000009_.Content = `LEFT_OR_TOP`
	__LinkAnchoredText__00000009_.AutomaticLayout = true
	__LinkAnchoredText__00000009_.LinkAnchorType = models.LINK_LEFT_OR_TOP
	__LinkAnchoredText__00000009_.X_Offset = 0.000000
	__LinkAnchoredText__00000009_.Y_Offset = 0.000000
	__LinkAnchoredText__00000009_.FontWeight = ``
	__LinkAnchoredText__00000009_.FontSize = ``
	__LinkAnchoredText__00000009_.FontStyle = ``
	__LinkAnchoredText__00000009_.LetterSpacing = ``
	__LinkAnchoredText__00000009_.FontFamily = ``
	__LinkAnchoredText__00000009_.WhiteSpace = ""
	__LinkAnchoredText__00000009_.Color = `black`
	__LinkAnchoredText__00000009_.FillOpacity = 1.000000
	__LinkAnchoredText__00000009_.Stroke = `black`
	__LinkAnchoredText__00000009_.StrokeOpacity = 1.000000
	__LinkAnchoredText__00000009_.StrokeWidth = 1.000000
	__LinkAnchoredText__00000009_.StrokeDashArray = ``
	__LinkAnchoredText__00000009_.StrokeDashArrayWhenSelected = ``
	__LinkAnchoredText__00000009_.Transform = ``

	__LinkAnchoredText__00000010_.Name = `AA_RIGHT_OR_BOTTOM_BELOW`
	__LinkAnchoredText__00000010_.Content = `AA_RIGHT_OR_BOTTOM_BELOW`
	__LinkAnchoredText__00000010_.AutomaticLayout = true
	__LinkAnchoredText__00000010_.LinkAnchorType = models.LINK_RIGHT_OR_BOTTOM
	__LinkAnchoredText__00000010_.X_Offset = 0.000000
	__LinkAnchoredText__00000010_.Y_Offset = 20.000000
	__LinkAnchoredText__00000010_.FontWeight = ``
	__LinkAnchoredText__00000010_.FontSize = ``
	__LinkAnchoredText__00000010_.FontStyle = ``
	__LinkAnchoredText__00000010_.LetterSpacing = ``
	__LinkAnchoredText__00000010_.FontFamily = ``
	__LinkAnchoredText__00000010_.WhiteSpace = ""
	__LinkAnchoredText__00000010_.Color = `red`
	__LinkAnchoredText__00000010_.FillOpacity = 1.000000
	__LinkAnchoredText__00000010_.Stroke = `red`
	__LinkAnchoredText__00000010_.StrokeOpacity = 1.000000
	__LinkAnchoredText__00000010_.StrokeWidth = 1.000000
	__LinkAnchoredText__00000010_.StrokeDashArray = ``
	__LinkAnchoredText__00000010_.StrokeDashArrayWhenSelected = ``
	__LinkAnchoredText__00000010_.Transform = ``

	__LinkAnchoredText__00000011_.Name = `LEFT_OR_TOP_ABOVE`
	__LinkAnchoredText__00000011_.Content = `LEFT_OR_TOP_ABOVE`
	__LinkAnchoredText__00000011_.AutomaticLayout = true
	__LinkAnchoredText__00000011_.LinkAnchorType = models.LINK_LEFT_OR_TOP
	__LinkAnchoredText__00000011_.X_Offset = 0.000000
	__LinkAnchoredText__00000011_.Y_Offset = -20.000000
	__LinkAnchoredText__00000011_.FontWeight = ``
	__LinkAnchoredText__00000011_.FontSize = ``
	__LinkAnchoredText__00000011_.FontStyle = ``
	__LinkAnchoredText__00000011_.LetterSpacing = ``
	__LinkAnchoredText__00000011_.FontFamily = ``
	__LinkAnchoredText__00000011_.WhiteSpace = ""
	__LinkAnchoredText__00000011_.Color = `blue`
	__LinkAnchoredText__00000011_.FillOpacity = 1.000000
	__LinkAnchoredText__00000011_.Stroke = `blue`
	__LinkAnchoredText__00000011_.StrokeOpacity = 1.000000
	__LinkAnchoredText__00000011_.StrokeWidth = 1.000000
	__LinkAnchoredText__00000011_.StrokeDashArray = ``
	__LinkAnchoredText__00000011_.StrokeDashArrayWhenSelected = ``
	__LinkAnchoredText__00000011_.Transform = ``

	__Rect__00000000_.Name = `Bottom Rect`
	__Rect__00000000_.X = 333.000000
	__Rect__00000000_.Y = 616.000000
	__Rect__00000000_.Width = 584.000000
	__Rect__00000000_.Height = 200.999985
	__Rect__00000000_.RX = 5.000000
	__Rect__00000000_.Color = `bisque`
	__Rect__00000000_.FillOpacity = 50.000000
	__Rect__00000000_.Stroke = `lightcoral`
	__Rect__00000000_.StrokeOpacity = 1.000000
	__Rect__00000000_.StrokeWidth = 3.000000
	__Rect__00000000_.StrokeDashArray = ``
	__Rect__00000000_.StrokeDashArrayWhenSelected = `5 5`
	__Rect__00000000_.Transform = ``
	__Rect__00000000_.IsSelectable = true
	__Rect__00000000_.IsSelected = true
	__Rect__00000000_.CanHaveLeftHandle = true
	__Rect__00000000_.HasLeftHandle = true
	__Rect__00000000_.CanHaveRightHandle = true
	__Rect__00000000_.HasRightHandle = true
	__Rect__00000000_.CanHaveTopHandle = true
	__Rect__00000000_.HasTopHandle = true
	__Rect__00000000_.IsScalingProportionally = false
	__Rect__00000000_.CanHaveBottomHandle = true
	__Rect__00000000_.HasBottomHandle = true
	__Rect__00000000_.CanMoveHorizontaly = true
	__Rect__00000000_.CanMoveVerticaly = true
	__Rect__00000000_.ChangeColorWhenHovered = false
	__Rect__00000000_.ColorWhenHovered = ``
	__Rect__00000000_.OriginalColor = ``
	__Rect__00000000_.FillOpacityWhenHovered = 0.000000
	__Rect__00000000_.OriginalFillOpacity = 0.000000
	__Rect__00000000_.HasToolTip = false
	__Rect__00000000_.ToolTipText = ``
	__Rect__00000000_.ToolTipPosition = ""
	__Rect__00000000_.MouseX = 0.000000
	__Rect__00000000_.MouseY = 0.000000
	__Rect__00000000_.MouseEventKey = ""

	__Rect__00000001_.Name = `Middle Rect`
	__Rect__00000001_.X = 461.000000
	__Rect__00000001_.Y = 426.000000
	__Rect__00000001_.Width = 200.000000
	__Rect__00000001_.Height = 132.000000
	__Rect__00000001_.RX = 3.000000
	__Rect__00000001_.Color = `lavender`
	__Rect__00000001_.FillOpacity = 0.000000
	__Rect__00000001_.Stroke = `turquoise`
	__Rect__00000001_.StrokeOpacity = 1.000000
	__Rect__00000001_.StrokeWidth = 1.000000
	__Rect__00000001_.StrokeDashArray = ``
	__Rect__00000001_.StrokeDashArrayWhenSelected = `5 5`
	__Rect__00000001_.Transform = ``
	__Rect__00000001_.IsSelectable = true
	__Rect__00000001_.IsSelected = false
	__Rect__00000001_.CanHaveLeftHandle = true
	__Rect__00000001_.HasLeftHandle = false
	__Rect__00000001_.CanHaveRightHandle = true
	__Rect__00000001_.HasRightHandle = false
	__Rect__00000001_.CanHaveTopHandle = true
	__Rect__00000001_.HasTopHandle = false
	__Rect__00000001_.IsScalingProportionally = true
	__Rect__00000001_.CanHaveBottomHandle = true
	__Rect__00000001_.HasBottomHandle = false
	__Rect__00000001_.CanMoveHorizontaly = true
	__Rect__00000001_.CanMoveVerticaly = true
	__Rect__00000001_.ChangeColorWhenHovered = false
	__Rect__00000001_.ColorWhenHovered = ``
	__Rect__00000001_.OriginalColor = ``
	__Rect__00000001_.FillOpacityWhenHovered = 0.000000
	__Rect__00000001_.OriginalFillOpacity = 0.000000
	__Rect__00000001_.HasToolTip = false
	__Rect__00000001_.ToolTipText = ``
	__Rect__00000001_.ToolTipPosition = ""
	__Rect__00000001_.MouseX = 0.000000
	__Rect__00000001_.MouseY = 0.000000
	__Rect__00000001_.MouseEventKey = ""

	__Rect__00000002_.Name = `Top Rect`
	__Rect__00000002_.X = 425.999969
	__Rect__00000002_.Y = 68.000000
	__Rect__00000002_.Width = 237.000000
	__Rect__00000002_.Height = 237.000000
	__Rect__00000002_.RX = 3.000000
	__Rect__00000002_.Color = `lightcyan`
	__Rect__00000002_.FillOpacity = 0.000000
	__Rect__00000002_.Stroke = `darkcyan`
	__Rect__00000002_.StrokeOpacity = 1.000000
	__Rect__00000002_.StrokeWidth = 2.000000
	__Rect__00000002_.StrokeDashArray = ``
	__Rect__00000002_.StrokeDashArrayWhenSelected = `5 5`
	__Rect__00000002_.Transform = ``
	__Rect__00000002_.IsSelectable = true
	__Rect__00000002_.IsSelected = false
	__Rect__00000002_.CanHaveLeftHandle = true
	__Rect__00000002_.HasLeftHandle = false
	__Rect__00000002_.CanHaveRightHandle = true
	__Rect__00000002_.HasRightHandle = false
	__Rect__00000002_.CanHaveTopHandle = false
	__Rect__00000002_.HasTopHandle = false
	__Rect__00000002_.IsScalingProportionally = true
	__Rect__00000002_.CanHaveBottomHandle = false
	__Rect__00000002_.HasBottomHandle = false
	__Rect__00000002_.CanMoveHorizontaly = true
	__Rect__00000002_.CanMoveVerticaly = true
	__Rect__00000002_.ChangeColorWhenHovered = false
	__Rect__00000002_.ColorWhenHovered = ``
	__Rect__00000002_.OriginalColor = ``
	__Rect__00000002_.FillOpacityWhenHovered = 0.000000
	__Rect__00000002_.OriginalFillOpacity = 0.000000
	__Rect__00000002_.HasToolTip = false
	__Rect__00000002_.ToolTipText = ``
	__Rect__00000002_.ToolTipPosition = ""
	__Rect__00000002_.MouseX = 749.000000
	__Rect__00000002_.MouseY = 289.500000
	__Rect__00000002_.MouseEventKey = models.MouseEventKeyShift

	__Rect__00000003_.Name = `Start`
	__Rect__00000003_.X = 1125.333333
	__Rect__00000003_.Y = 62.333332
	__Rect__00000003_.Width = 853.000000
	__Rect__00000003_.Height = 125.571444
	__Rect__00000003_.RX = 20.000000
	__Rect__00000003_.Color = ``
	__Rect__00000003_.FillOpacity = 0.000000
	__Rect__00000003_.Stroke = `black`
	__Rect__00000003_.StrokeOpacity = 1.000000
	__Rect__00000003_.StrokeWidth = 1.000000
	__Rect__00000003_.StrokeDashArray = ``
	__Rect__00000003_.StrokeDashArrayWhenSelected = ``
	__Rect__00000003_.Transform = ``
	__Rect__00000003_.IsSelectable = true
	__Rect__00000003_.IsSelected = false
	__Rect__00000003_.CanHaveLeftHandle = true
	__Rect__00000003_.HasLeftHandle = false
	__Rect__00000003_.CanHaveRightHandle = true
	__Rect__00000003_.HasRightHandle = false
	__Rect__00000003_.CanHaveTopHandle = true
	__Rect__00000003_.HasTopHandle = false
	__Rect__00000003_.IsScalingProportionally = false
	__Rect__00000003_.CanHaveBottomHandle = true
	__Rect__00000003_.HasBottomHandle = false
	__Rect__00000003_.CanMoveHorizontaly = true
	__Rect__00000003_.CanMoveVerticaly = true
	__Rect__00000003_.ChangeColorWhenHovered = false
	__Rect__00000003_.ColorWhenHovered = ``
	__Rect__00000003_.OriginalColor = ``
	__Rect__00000003_.FillOpacityWhenHovered = 0.000000
	__Rect__00000003_.OriginalFillOpacity = 0.000000
	__Rect__00000003_.HasToolTip = false
	__Rect__00000003_.ToolTipText = ``
	__Rect__00000003_.ToolTipPosition = ""
	__Rect__00000003_.MouseX = 0.000000
	__Rect__00000003_.MouseY = 0.000000
	__Rect__00000003_.MouseEventKey = ""

	__Rect__00000004_.Name = `End`
	__Rect__00000004_.X = 1174.769231
	__Rect__00000004_.Y = 397.256327
	__Rect__00000004_.Width = 430.000000
	__Rect__00000004_.Height = 384.000000
	__Rect__00000004_.RX = 20.000000
	__Rect__00000004_.Color = ``
	__Rect__00000004_.FillOpacity = 0.000000
	__Rect__00000004_.Stroke = `black`
	__Rect__00000004_.StrokeOpacity = 1.000000
	__Rect__00000004_.StrokeWidth = 1.000000
	__Rect__00000004_.StrokeDashArray = ``
	__Rect__00000004_.StrokeDashArrayWhenSelected = ``
	__Rect__00000004_.Transform = ``
	__Rect__00000004_.IsSelectable = true
	__Rect__00000004_.IsSelected = false
	__Rect__00000004_.CanHaveLeftHandle = true
	__Rect__00000004_.HasLeftHandle = false
	__Rect__00000004_.CanHaveRightHandle = true
	__Rect__00000004_.HasRightHandle = false
	__Rect__00000004_.CanHaveTopHandle = true
	__Rect__00000004_.HasTopHandle = false
	__Rect__00000004_.IsScalingProportionally = false
	__Rect__00000004_.CanHaveBottomHandle = true
	__Rect__00000004_.HasBottomHandle = false
	__Rect__00000004_.CanMoveHorizontaly = true
	__Rect__00000004_.CanMoveVerticaly = true
	__Rect__00000004_.ChangeColorWhenHovered = false
	__Rect__00000004_.ColorWhenHovered = ``
	__Rect__00000004_.OriginalColor = ``
	__Rect__00000004_.FillOpacityWhenHovered = 0.000000
	__Rect__00000004_.OriginalFillOpacity = 0.000000
	__Rect__00000004_.HasToolTip = false
	__Rect__00000004_.ToolTipText = ``
	__Rect__00000004_.ToolTipPosition = ""
	__Rect__00000004_.MouseX = 0.000000
	__Rect__00000004_.MouseY = 0.000000
	__Rect__00000004_.MouseEventKey = ""

	__RectAnchoredPath__00000000_.Name = `Logo to add`
	__RectAnchoredPath__00000000_.Definition = `M532-131q-6 5-12.5 8t-14.5 3q-8 0-16-3.5t-14-9.5q-41-44-60.5-90T395-320q0-37 11-78t38-106q23-57 32-87.5t9-56.5q0-34-15-63.5T423-771q-6-6-9.5-14t-3.5-16q0-8 3-14.5t8-12.5q6-6 13.5-9t15.5-3q8 0 15 3t13 8q44 41 65.5 86t21.5 95q0 35-10.5 73.5T518-474q-25 60-34 92t-9 61q0 35 14.5 67.5T534-188q5 6 8 13t3 15q0 8-3 15.5T532-131Zm195 0q-6 5-12.5 8t-14.5 3q-8 0-16-3.5t-14-9.5q-41-44-60.5-89.5T590-319q0-37 11-79t38-106q23-57 32-87t9-56q0-34-15-64.5T618-771q-6-6-9-13.5t-3-15.5q0-8 2.5-14.5T616-827q6-6 14-9.5t16-3.5q8 0 14.5 3t12.5 8q44 41 65.5 86t21.5 95q0 35-10.5 73.5T713-473q-25 60-34 92t-9 60q0 35 15 68.5t45 65.5q5 6 7.5 13t2.5 14q0 8-3 16t-10 13Zm-390 0q-6 5-12.5 8t-14.5 3q-8 0-16-3.5t-14-9.5q-41-44-60.5-89.5T200-319q0-37 11-79t38-106q23-57 32-87t9-56q0-34-15-64.5T228-771q-7-6-10-13.5t-3-15.5q0-8 3-15t8-13q6-6 13.5-9t15.5-3q8 0 15 3t13 8q44 41 65.5 85.5T370-648q0 35-10 73.5T324-474q-25 60-34 92t-9 61q0 35 14.5 68.5T340-187q5 6 7.5 13t2.5 14q0 8-3 16t-10 13Z`
	__RectAnchoredPath__00000000_.X_Offset = 0.000000
	__RectAnchoredPath__00000000_.Y_Offset = 0.000000
	__RectAnchoredPath__00000000_.RectAnchorType = models.RECT_BOTTOM_LEFT
	__RectAnchoredPath__00000000_.ScalePropotionnally = true
	__RectAnchoredPath__00000000_.AppliedScaling = 0.246875
	__RectAnchoredPath__00000000_.Color = `black`
	__RectAnchoredPath__00000000_.FillOpacity = 0.500000
	__RectAnchoredPath__00000000_.Stroke = `black`
	__RectAnchoredPath__00000000_.StrokeOpacity = 1.000000
	__RectAnchoredPath__00000000_.StrokeWidth = 1.000000
	__RectAnchoredPath__00000000_.StrokeDashArray = ``
	__RectAnchoredPath__00000000_.StrokeDashArrayWhenSelected = ``
	__RectAnchoredPath__00000000_.Transform = ``

	__RectAnchoredPath__00000001_.Name = `second log`
	__RectAnchoredPath__00000001_.Definition = `M532-131q-6 5-12.5 8t-14.5 3q-8 0-16-3.5t-14-9.5q-41-44-60.5-90T395-320q0-37 11-78t38-106q23-57 32-87.5t9-56.5q0-34-15-63.5T423-771q-6-6-9.5-14t-3.5-16q0-8 3-14.5t8-12.5q6-6 13.5-9t15.5-3q8 0 15 3t13 8q44 41 65.5 86t21.5 95q0 35-10.5 73.5T518-474q-25 60-34 92t-9 61q0 35 14.5 67.5T534-188q5 6 8 13t3 15q0 8-3 15.5T532-131Zm195 0q-6 5-12.5 8t-14.5 3q-8 0-16-3.5t-14-9.5q-41-44-60.5-89.5T590-319q0-37 11-79t38-106q23-57 32-87t9-56q0-34-15-64.5T618-771q-6-6-9-13.5t-3-15.5q0-8 2.5-14.5T616-827q6-6 14-9.5t16-3.5q8 0 14.5 3t12.5 8q44 41 65.5 86t21.5 95q0 35-10.5 73.5T713-473q-25 60-34 92t-9 60q0 35 15 68.5t45 65.5q5 6 7.5 13t2.5 14q0 8-3 16t-10 13Zm-390 0q-6 5-12.5 8t-14.5 3q-8 0-16-3.5t-14-9.5q-41-44-60.5-89.5T200-319q0-37 11-79t38-106q23-57 32-87t9-56q0-34-15-64.5T228-771q-7-6-10-13.5t-3-15.5q0-8 3-15t8-13q6-6 13.5-9t15.5-3q8 0 15 3t13 8q44 41 65.5 85.5T370-648q0 35-10 73.5T324-474q-25 60-34 92t-9 61q0 35 14.5 68.5T340-187q5 6 7.5 13t2.5 14q0 8-3 16t-10 13Z`
	__RectAnchoredPath__00000001_.X_Offset = 0.000000
	__RectAnchoredPath__00000001_.Y_Offset = 0.000000
	__RectAnchoredPath__00000001_.RectAnchorType = models.RECT_BOTTOM_LEFT
	__RectAnchoredPath__00000001_.ScalePropotionnally = true
	__RectAnchoredPath__00000001_.AppliedScaling = 0.200000
	__RectAnchoredPath__00000001_.Color = ``
	__RectAnchoredPath__00000001_.FillOpacity = 0.000000
	__RectAnchoredPath__00000001_.Stroke = `blue`
	__RectAnchoredPath__00000001_.StrokeOpacity = 1.000000
	__RectAnchoredPath__00000001_.StrokeWidth = 2.000000
	__RectAnchoredPath__00000001_.StrokeDashArray = ``
	__RectAnchoredPath__00000001_.StrokeDashArrayWhenSelected = ``
	__RectAnchoredPath__00000001_.Transform = ``

	__RectAnchoredRect__00000000_.Name = `Rect within top`
	__RectAnchoredRect__00000000_.X = 0.000000
	__RectAnchoredRect__00000000_.Y = 0.000000
	__RectAnchoredRect__00000000_.Width = 100.000000
	__RectAnchoredRect__00000000_.Height = 30.000000
	__RectAnchoredRect__00000000_.RX = 3.000000
	__RectAnchoredRect__00000000_.X_Offset = 0.000000
	__RectAnchoredRect__00000000_.Y_Offset = 40.000000
	__RectAnchoredRect__00000000_.RectAnchorType = models.RECT_TOP
	__RectAnchoredRect__00000000_.WidthFollowRect = false
	__RectAnchoredRect__00000000_.HeightFollowRect = false
	__RectAnchoredRect__00000000_.HasToolTip = false
	__RectAnchoredRect__00000000_.ToolTipText = ``
	__RectAnchoredRect__00000000_.Color = `lightgrey`
	__RectAnchoredRect__00000000_.FillOpacity = 40.000000
	__RectAnchoredRect__00000000_.Stroke = `bisque`
	__RectAnchoredRect__00000000_.StrokeOpacity = 1.000000
	__RectAnchoredRect__00000000_.StrokeWidth = 1.000000
	__RectAnchoredRect__00000000_.StrokeDashArray = ``
	__RectAnchoredRect__00000000_.StrokeDashArrayWhenSelected = ``
	__RectAnchoredRect__00000000_.Transform = ``

	__RectAnchoredRect__00000001_.Name = `Top on Bottom with same width`
	__RectAnchoredRect__00000001_.X = 0.000000
	__RectAnchoredRect__00000001_.Y = 0.000000
	__RectAnchoredRect__00000001_.Width = 100.000000
	__RectAnchoredRect__00000001_.Height = 50.000000
	__RectAnchoredRect__00000001_.RX = 0.000000
	__RectAnchoredRect__00000001_.X_Offset = 0.000000
	__RectAnchoredRect__00000001_.Y_Offset = 0.000000
	__RectAnchoredRect__00000001_.RectAnchorType = models.RECT_TOP_LEFT
	__RectAnchoredRect__00000001_.WidthFollowRect = true
	__RectAnchoredRect__00000001_.HeightFollowRect = false
	__RectAnchoredRect__00000001_.HasToolTip = false
	__RectAnchoredRect__00000001_.ToolTipText = ``
	__RectAnchoredRect__00000001_.Color = `lightblue`
	__RectAnchoredRect__00000001_.FillOpacity = 100.000000
	__RectAnchoredRect__00000001_.Stroke = ``
	__RectAnchoredRect__00000001_.StrokeOpacity = 1.000000
	__RectAnchoredRect__00000001_.StrokeWidth = 0.000000
	__RectAnchoredRect__00000001_.StrokeDashArray = ``
	__RectAnchoredRect__00000001_.StrokeDashArrayWhenSelected = ``
	__RectAnchoredRect__00000001_.Transform = ``

	__RectAnchoredText__00000000_.Name = `Bottom Text`
	__RectAnchoredText__00000000_.Content = `This      is an example of a note that
could be displayed on a diagram.

It could explain one aspect of the model
for instance, describing relations between structs

The text of a UML note refers a comment with the GONGNOTE keyword which is
a special case of go Note convention. See example
for details in the go code of the models.
`
	__RectAnchoredText__00000000_.FontWeight = `100`
	__RectAnchoredText__00000000_.FontSize = ``
	__RectAnchoredText__00000000_.FontStyle = ``
	__RectAnchoredText__00000000_.LetterSpacing = ``
	__RectAnchoredText__00000000_.FontFamily = ``
	__RectAnchoredText__00000000_.WhiteSpace = models.WhiteSpaceEnumPre
	__RectAnchoredText__00000000_.X_Offset = 0.000000
	__RectAnchoredText__00000000_.Y_Offset = 20.000000
	__RectAnchoredText__00000000_.RectAnchorType = models.RECT_TOP
	__RectAnchoredText__00000000_.TextAnchorType = models.TEXT_ANCHOR_CENTER
	__RectAnchoredText__00000000_.DominantBaseline = models.DominantBaselineCentral
	__RectAnchoredText__00000000_.WritingMode = ""
	__RectAnchoredText__00000000_.Color = `black`
	__RectAnchoredText__00000000_.FillOpacity = 100.000000
	__RectAnchoredText__00000000_.Stroke = `black`
	__RectAnchoredText__00000000_.StrokeOpacity = 1.000000
	__RectAnchoredText__00000000_.StrokeWidth = 1.000000
	__RectAnchoredText__00000000_.StrokeDashArray = ``
	__RectAnchoredText__00000000_.StrokeDashArrayWhenSelected = ``
	__RectAnchoredText__00000000_.Transform = ``

	__RectAnchoredText__00000001_.Name = `Top Left`
	__RectAnchoredText__00000001_.Content = `Top Left`
	__RectAnchoredText__00000001_.FontWeight = ``
	__RectAnchoredText__00000001_.FontSize = ``
	__RectAnchoredText__00000001_.FontStyle = ``
	__RectAnchoredText__00000001_.LetterSpacing = ``
	__RectAnchoredText__00000001_.FontFamily = ``
	__RectAnchoredText__00000001_.WhiteSpace = ""
	__RectAnchoredText__00000001_.X_Offset = 0.000000
	__RectAnchoredText__00000001_.Y_Offset = 0.000000
	__RectAnchoredText__00000001_.RectAnchorType = models.RECT_TOP_LEFT
	__RectAnchoredText__00000001_.TextAnchorType = ""
	__RectAnchoredText__00000001_.DominantBaseline = ""
	__RectAnchoredText__00000001_.WritingMode = ""
	__RectAnchoredText__00000001_.Color = ``
	__RectAnchoredText__00000001_.FillOpacity = 0.000000
	__RectAnchoredText__00000001_.Stroke = `black`
	__RectAnchoredText__00000001_.StrokeOpacity = 1.000000
	__RectAnchoredText__00000001_.StrokeWidth = 0.000000
	__RectAnchoredText__00000001_.StrokeDashArray = ``
	__RectAnchoredText__00000001_.StrokeDashArrayWhenSelected = ``
	__RectAnchoredText__00000001_.Transform = ``

	__RectAnchoredText__00000002_.Name = `Top anchored bottom middle`
	__RectAnchoredText__00000002_.Content = `Top anchored bottom middle`
	__RectAnchoredText__00000002_.FontWeight = ``
	__RectAnchoredText__00000002_.FontSize = ``
	__RectAnchoredText__00000002_.FontStyle = ``
	__RectAnchoredText__00000002_.LetterSpacing = ``
	__RectAnchoredText__00000002_.FontFamily = ``
	__RectAnchoredText__00000002_.WhiteSpace = ""
	__RectAnchoredText__00000002_.X_Offset = 0.000000
	__RectAnchoredText__00000002_.Y_Offset = 20.000000
	__RectAnchoredText__00000002_.RectAnchorType = models.RECT_BOTTOM
	__RectAnchoredText__00000002_.TextAnchorType = models.TEXT_ANCHOR_START
	__RectAnchoredText__00000002_.DominantBaseline = ""
	__RectAnchoredText__00000002_.WritingMode = ""
	__RectAnchoredText__00000002_.Color = ``
	__RectAnchoredText__00000002_.FillOpacity = 100.000000
	__RectAnchoredText__00000002_.Stroke = `blue`
	__RectAnchoredText__00000002_.StrokeOpacity = 1.000000
	__RectAnchoredText__00000002_.StrokeWidth = 2.000000
	__RectAnchoredText__00000002_.StrokeDashArray = ``
	__RectAnchoredText__00000002_.StrokeDashArrayWhenSelected = ``
	__RectAnchoredText__00000002_.Transform = ``

	__RectAnchoredText__00000003_.Name = `Top anchored top middle`
	__RectAnchoredText__00000003_.Content = `Top anchored
top middle
line 3`
	__RectAnchoredText__00000003_.FontWeight = `bold`
	__RectAnchoredText__00000003_.FontSize = ``
	__RectAnchoredText__00000003_.FontStyle = `italic`
	__RectAnchoredText__00000003_.LetterSpacing = ``
	__RectAnchoredText__00000003_.FontFamily = ``
	__RectAnchoredText__00000003_.WhiteSpace = ""
	__RectAnchoredText__00000003_.X_Offset = 0.000000
	__RectAnchoredText__00000003_.Y_Offset = 20.000000
	__RectAnchoredText__00000003_.RectAnchorType = models.RECT_TOP
	__RectAnchoredText__00000003_.TextAnchorType = models.TEXT_ANCHOR_START
	__RectAnchoredText__00000003_.DominantBaseline = ""
	__RectAnchoredText__00000003_.WritingMode = ""
	__RectAnchoredText__00000003_.Color = ``
	__RectAnchoredText__00000003_.FillOpacity = 100.000000
	__RectAnchoredText__00000003_.Stroke = `black`
	__RectAnchoredText__00000003_.StrokeOpacity = 1.000000
	__RectAnchoredText__00000003_.StrokeWidth = 1.000000
	__RectAnchoredText__00000003_.StrokeDashArray = ``
	__RectAnchoredText__00000003_.StrokeDashArrayWhenSelected = ``
	__RectAnchoredText__00000003_.Transform = ``

	__RectAnchoredText__00000004_.Name = `Start`
	__RectAnchoredText__00000004_.Content = `Start`
	__RectAnchoredText__00000004_.FontWeight = ``
	__RectAnchoredText__00000004_.FontSize = `27px`
	__RectAnchoredText__00000004_.FontStyle = ``
	__RectAnchoredText__00000004_.LetterSpacing = ``
	__RectAnchoredText__00000004_.FontFamily = ``
	__RectAnchoredText__00000004_.WhiteSpace = ""
	__RectAnchoredText__00000004_.X_Offset = 0.000000
	__RectAnchoredText__00000004_.Y_Offset = 0.000000
	__RectAnchoredText__00000004_.RectAnchorType = models.RECT_CENTER
	__RectAnchoredText__00000004_.TextAnchorType = models.TEXT_ANCHOR_CENTER
	__RectAnchoredText__00000004_.DominantBaseline = models.DominantBaselineCentral
	__RectAnchoredText__00000004_.WritingMode = ""
	__RectAnchoredText__00000004_.Color = `black`
	__RectAnchoredText__00000004_.FillOpacity = 1.000000
	__RectAnchoredText__00000004_.Stroke = `black`
	__RectAnchoredText__00000004_.StrokeOpacity = 1.000000
	__RectAnchoredText__00000004_.StrokeWidth = 1.000000
	__RectAnchoredText__00000004_.StrokeDashArray = ``
	__RectAnchoredText__00000004_.StrokeDashArrayWhenSelected = ``
	__RectAnchoredText__00000004_.Transform = ``

	__RectAnchoredText__00000005_.Name = `End`
	__RectAnchoredText__00000005_.Content = `End`
	__RectAnchoredText__00000005_.FontWeight = ``
	__RectAnchoredText__00000005_.FontSize = `27px`
	__RectAnchoredText__00000005_.FontStyle = ``
	__RectAnchoredText__00000005_.LetterSpacing = ``
	__RectAnchoredText__00000005_.FontFamily = ``
	__RectAnchoredText__00000005_.WhiteSpace = ""
	__RectAnchoredText__00000005_.X_Offset = 0.000000
	__RectAnchoredText__00000005_.Y_Offset = 0.000000
	__RectAnchoredText__00000005_.RectAnchorType = models.RECT_CENTER
	__RectAnchoredText__00000005_.TextAnchorType = models.TEXT_ANCHOR_CENTER
	__RectAnchoredText__00000005_.DominantBaseline = ""
	__RectAnchoredText__00000005_.WritingMode = ""
	__RectAnchoredText__00000005_.Color = `black`
	__RectAnchoredText__00000005_.FillOpacity = 1.000000
	__RectAnchoredText__00000005_.Stroke = `black`
	__RectAnchoredText__00000005_.StrokeOpacity = 1.000000
	__RectAnchoredText__00000005_.StrokeWidth = 1.000000
	__RectAnchoredText__00000005_.StrokeDashArray = ``
	__RectAnchoredText__00000005_.StrokeDashArrayWhenSelected = ``
	__RectAnchoredText__00000005_.Transform = ``

	__RectAnchoredText__00000006_.Name = ``
	__RectAnchoredText__00000006_.Content = ``
	__RectAnchoredText__00000006_.FontWeight = ``
	__RectAnchoredText__00000006_.FontSize = `27px`
	__RectAnchoredText__00000006_.FontStyle = ``
	__RectAnchoredText__00000006_.LetterSpacing = ``
	__RectAnchoredText__00000006_.FontFamily = ``
	__RectAnchoredText__00000006_.WhiteSpace = ""
	__RectAnchoredText__00000006_.X_Offset = 0.000000
	__RectAnchoredText__00000006_.Y_Offset = 0.000000
	__RectAnchoredText__00000006_.RectAnchorType = ""
	__RectAnchoredText__00000006_.TextAnchorType = ""
	__RectAnchoredText__00000006_.DominantBaseline = ""
	__RectAnchoredText__00000006_.WritingMode = ""
	__RectAnchoredText__00000006_.Color = ``
	__RectAnchoredText__00000006_.FillOpacity = 0.000000
	__RectAnchoredText__00000006_.Stroke = ``
	__RectAnchoredText__00000006_.StrokeOpacity = 0.000000
	__RectAnchoredText__00000006_.StrokeWidth = 0.000000
	__RectAnchoredText__00000006_.StrokeDashArray = ``
	__RectAnchoredText__00000006_.StrokeDashArrayWhenSelected = ``
	__RectAnchoredText__00000006_.Transform = ``

	__SVG__00000000_.Name = `SVG`
	__SVG__00000000_.DrawingState = models.NOT_DRAWING_LINK
	__SVG__00000000_.IsEditable = true
	__SVG__00000000_.IsSVGFrontEndFileGenerated = false
	__SVG__00000000_.IsSVGBackEndFileGenerated = false
	__SVG__00000000_.DefaultDirectoryForGeneratedImages = `../../diagrams/images`
	__SVG__00000000_.IsControlBannerHidden = false
	__SVG__00000000_.OverrideWidth = false
	__SVG__00000000_.OverriddenWidth = 0.000000
	__SVG__00000000_.OverrideHeight = false
	__SVG__00000000_.OverriddenHeight = 0.000000

	__SvgText__00000000_.Name = `Essai`
	__SvgText__00000000_.Text = `<svg xmlns="http://www.w3.org/2000/svg" width="2636.923" height="2636.923" viewBox="-10 -10 2636.923 2636.923"><style>text { font-family: Roboto, Arial, sans-serif !important; }</style>
  <g transform="translate(0 0) scale(1)">
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <rect x="1000" y="100" width="200" height="400" rx="0" fill="" fill-opacity="0" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <text class="anchored-text click-through" x="1100" y="300" fill="black" fill-opacity="1" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="" text-anchor="middle" font-weight="" font-style="" font-size="27pxpx">
      <tspan x="1100" dy="0" text-anchor="middle">Start</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <rect x="1900" y="370" width="140" height="384" rx="0" fill="" fill-opacity="0" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="">
      <!---->
    </rect>
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <text class="anchored-text click-through" x="1970" y="562" fill="black" fill-opacity="1" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="" text-anchor="middle" font-weight="" font-style="" font-size="27pxpx">
      <tspan x="1970" dy="0" text-anchor="middle">End</tspan>
      <!---->
    </text>
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <!---->
    <g>
    <line class="moveable-line" x1="1960.99996" y1="370" x2="1960.99996" y2="521.001856" fill="" fill-opacity="0" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray="" segment-orientation="vertical" segment-number="0"/>
    <line class="hit-area" x1="1960.99996" y1="370" x2="1960.99996" y2="521.001856" stroke="transparent" stroke-width="11" segment-orientation="vertical"/>
    <line class="moveable-line" x1="1960.99996" y1="521.001856" x2="1081.5384" y2="521.001856" fill="" fill-opacity="0" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray="" segment-orientation="horizontal" segment-number="1"/>
    <line class="hit-area" x1="1960.99996" y1="521.001856" x2="1081.5384" y2="521.001856" stroke="transparent" stroke-width="11" segment-orientation="horizontal"/>
    <line class="moveable-line" x1="1081.5384" y1="521.001856" x2="1081.5384" y2="500" fill="" fill-opacity="0" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray="" segment-orientation="vertical" segment-number="2"/>
    <line class="hit-area" x1="1081.5384" y1="521.001856" x2="1081.5384" y2="500" stroke="transparent" stroke-width="11" segment-orientation="vertical"/>
    <!---->
  </g>
  <!---->
  <path d="M 1960.99996 521.001856 A 0 0 0 0 1 1960.99996 521.001856" fill="" fill-opacity="0" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray=""/>
  <!---->
  <!---->
  <!---->
  <!---->
  <!---->
  <!---->
  <text class="noevents-area" x="1960.99996" y="394" font-weight="" font-size="" letter-spacing="" fill="black" fill-opacity="1" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="">
    <tspan x="1960.99996" dy="0" text-anchor="end">  LEFT_OR_TOP </tspan>
    <!---->
    <!---->
    <!---->
  </text>
  <!---->
  <!---->
  <text class="noevents-area" x="1960.99996" y="394" font-weight="" font-size="16" letter-spacing="" fill="black" fill-opacity="1" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="">
    <tspan x="1960.99996" dy="0" text-anchor="start">  RIGHT_OR_BOTTOM </tspan>
    <!---->
    <!---->
    <!---->
  </text>
  <!---->
  <!---->
  <!---->
  <path d="M 1081.5384 521.001856 A 0 0 0 0 1 1081.5384 521.001856" fill="" fill-opacity="0" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray=""/>
  <!---->
  <!---->
  <!---->
  <!---->
  <!---->
  <!---->
  <!---->
  <!---->
  <path d="M 1081.1848466094998 499.6464466095 L 1086.5384 505 M 1081.8919533905 499.6464466095 L 1076.5384 505" fill="" fill-opacity="0" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray=""/>
  <!---->
  <!---->
  <!---->
  <text class="noevents-area" x="1081.5384" y="524" font-weight="" font-size="" letter-spacing="" fill="black" fill-opacity="1" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="">
    <tspan x="1081.5384" dy="0" text-anchor="end">  LEFT_OR_TOP </tspan>
    <!---->
    <!---->
    <!---->
  </text>
  <!---->
  <!---->
  <text class="noevents-area" x="1081.5384" y="524" font-weight="" font-size="16" letter-spacing="" fill="black" fill-opacity="1" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="">
    <tspan x="1081.5384" dy="0" text-anchor="start">  RIGHT_OR_BOTTOM </tspan>
    <!---->
    <!---->
    <!---->
  </text>
  <!---->
  <!---->
  <!---->
  <!---->
  <!---->
  <!---->
  <!---->
  <!---->
  <!---->
  <g>
  <line class="moveable-line" x1="1000" y1="348.2224" x2="704.6154" y2="348.2224" fill="" fill-opacity="0" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray="" segment-orientation="horizontal" segment-number="0"/>
  <line class="hit-area" x1="1000" y1="348.2224" x2="704.6154" y2="348.2224" stroke="transparent" stroke-width="11" segment-orientation="horizontal"/>
  <line class="moveable-line" x1="704.6154" y1="348.2224" x2="704.6154" y2="596.5020159999999" fill="" fill-opacity="0" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray="" segment-orientation="vertical" segment-number="1"/>
  <line class="hit-area" x1="704.6154" y1="348.2224" x2="704.6154" y2="596.5020159999999" stroke="transparent" stroke-width="11" segment-orientation="vertical"/>
  <line class="moveable-line" x1="704.6154" y1="596.5020159999999" x2="1900" y2="596.5020159999999" fill="" fill-opacity="0" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray="" segment-orientation="horizontal" segment-number="2"/>
  <line class="hit-area" x1="704.6154" y1="596.5020159999999" x2="1900" y2="596.5020159999999" stroke="transparent" stroke-width="11" segment-orientation="horizontal"/>
  <!---->
</g>
<!---->
<path d="M 704.6154 348.2224 A 0 0 0 0 1 704.6154 348.2224" fill="" fill-opacity="0" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<text class="noevents-area" x="1000" y="335.62239999999997" font-weight="" font-size="" letter-spacing="" fill="black" fill-opacity="1" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="">
  <!---->
  <!---->
  <tspan x="1000" dy="0" text-anchor="end">  LEFT_OR_TOP </tspan>
  <!---->
  <!---->
  <!---->
</text>
<!---->
<!---->
<text class="noevents-area" x="1000" y="372.2224" font-weight="" font-size="16" letter-spacing="" fill="black" fill-opacity="1" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="">
  <!---->
  <!---->
  <tspan x="1000" dy="0" text-anchor="end">  RIGHT_OR_BOTTOM </tspan>
  <!---->
  <!---->
  <!---->
</text>
<!---->
<!---->
<!---->
<path d="M 704.6154 596.5020159999999 A 0 0 0 0 1 704.6154 596.5020159999999" fill="" fill-opacity="0" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 1900.3535533905 596.8555693904999 L 1895 591.5020159999999 M 1900.3535533905 596.1484626094999 L 1895 601.5020159999999" fill="" fill-opacity="0" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<text class="noevents-area" x="1900" y="583.9020159999999" font-weight="" font-size="" letter-spacing="" fill="black" fill-opacity="1" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="">
  <!---->
  <tspan x="1900" dy="0" text-anchor="end">  LEFT_OR_TOP </tspan>
  <!---->
  <!---->
  <!---->
  <!---->
</text>
<!---->
<!---->
<text class="noevents-area" x="1900" y="620.5020159999999" font-weight="" font-size="16" letter-spacing="" fill="black" fill-opacity="1" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="">
  <!---->
  <tspan x="1900" dy="0" text-anchor="end">  RIGHT_OR_BOTTOM </tspan>
  <!---->
  <!---->
  <!---->
  <!---->
</text>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<g>
<line class="moveable-line" x1="1200" y1="202.7452" x2="2616.923" y2="202.7452" fill="" fill-opacity="0" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray="" segment-orientation="horizontal" segment-number="0"/>
<line class="hit-area" x1="1200" y1="202.7452" x2="2616.923" y2="202.7452" stroke="transparent" stroke-width="11" segment-orientation="horizontal"/>
<line class="moveable-line" x1="2616.923" y1="202.7452" x2="2616.923" y2="428.887552" fill="" fill-opacity="0" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray="" segment-orientation="vertical" segment-number="1"/>
<line class="hit-area" x1="2616.923" y1="202.7452" x2="2616.923" y2="428.887552" stroke="transparent" stroke-width="11" segment-orientation="vertical"/>
<line class="moveable-line" x1="2616.923" y1="428.887552" x2="2040" y2="428.887552" fill="" fill-opacity="0" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray="" segment-orientation="horizontal" segment-number="2"/>
<line class="hit-area" x1="2616.923" y1="428.887552" x2="2040" y2="428.887552" stroke="transparent" stroke-width="11" segment-orientation="horizontal"/>
<!---->
</g>
<!---->
<path d="M 2616.923 202.7452 A 0 0 0 0 0 2616.923 202.7452" fill="" fill-opacity="0" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<text class="noevents-area" x="1200" y="190.14520000000002" font-weight="" font-size="" letter-spacing="" fill="black" fill-opacity="1" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="">
  <!---->
  <tspan x="1200" dy="0" text-anchor="start">  LEFT_OR_TOP </tspan>
  <!---->
  <!---->
  <!---->
  <!---->
</text>
<!---->
<!---->
<text class="noevents-area" x="1200" y="226.7452" font-weight="" font-size="16" letter-spacing="" fill="black" fill-opacity="1" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="">
  <!---->
  <tspan x="1200" dy="0" text-anchor="start">  RIGHT_OR_BOTTOM </tspan>
  <!---->
  <!---->
  <!---->
  <!---->
</text>
<!---->
<!---->
<!---->
<path d="M 2616.923 428.887552 A 0 0 0 0 1 2616.923 428.887552" fill="" fill-opacity="0" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<path d="M 2039.6464466095 429.2411053905 L 2045 423.887552 M 2039.6464466095 428.53399860950003 L 2045 433.887552" fill="" fill-opacity="0" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray=""/>
<!---->
<!---->
<!---->
<text class="noevents-area" x="2040" y="416.287552" font-weight="" font-size="" letter-spacing="" fill="black" fill-opacity="1" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="">
  <!---->
  <!---->
  <tspan x="2040" dy="0" text-anchor="start">  LEFT_OR_TOP </tspan>
  <!---->
  <!---->
  <!---->
</text>
<!---->
<!---->
<text class="noevents-area" x="2040" y="452.887552" font-weight="" font-size="16" letter-spacing="" fill="black" fill-opacity="1" stroke="black" stroke-opacity="1" stroke-width="1" stroke-dasharray="" transform="">
  <!---->
  <!---->
  <tspan x="2040" dy="0" text-anchor="start">  RIGHT_OR_BOTTOM </tspan>
  <!---->
  <!---->
  <!---->
</text>
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
<!---->
</g>
</svg>`

	__Text__00000000_.Name = `Essai`
	__Text__00000000_.X = 50.000000
	__Text__00000000_.Y = 300.000000
	__Text__00000000_.Content = `Hello`
	__Text__00000000_.Color = ``
	__Text__00000000_.FillOpacity = 0.000000
	__Text__00000000_.Stroke = `black`
	__Text__00000000_.StrokeOpacity = 1.000000
	__Text__00000000_.StrokeWidth = 2.000000
	__Text__00000000_.StrokeDashArray = ``
	__Text__00000000_.StrokeDashArrayWhenSelected = ``
	__Text__00000000_.Transform = ``
	__Text__00000000_.FontWeight = ``
	__Text__00000000_.FontSize = ``
	__Text__00000000_.FontStyle = ``
	__Text__00000000_.LetterSpacing = ``
	__Text__00000000_.FontFamily = ``
	__Text__00000000_.WhiteSpace = ""

	// insertion point for setup of pointers
	__ControlPoint__00000000_.ClosestRect = __Rect__00000001_
	__Layer__00000000_.Rects = append(__Layer__00000000_.Rects, __Rect__00000000_)
	__Layer__00000001_.Links = append(__Layer__00000001_.Links, __Link__00000001_)
	__Layer__00000003_.Links = append(__Layer__00000003_.Links, __Link__00000000_)
	__Layer__00000003_.Links = append(__Layer__00000003_.Links, __Link__00000002_)
	__Layer__00000004_.Rects = append(__Layer__00000004_.Rects, __Rect__00000001_)
	__Layer__00000004_.Links = append(__Layer__00000004_.Links, __Link__00000006_)
	__Layer__00000005_.Rects = append(__Layer__00000005_.Rects, __Rect__00000002_)
	__Layer__00000006_.Rects = append(__Layer__00000006_.Rects, __Rect__00000003_)
	__Layer__00000006_.Rects = append(__Layer__00000006_.Rects, __Rect__00000004_)
	__Layer__00000006_.Links = append(__Layer__00000006_.Links, __Link__00000003_)
	__Layer__00000006_.Links = append(__Layer__00000006_.Links, __Link__00000004_)
	__Layer__00000006_.Links = append(__Layer__00000006_.Links, __Link__00000005_)
	__Link__00000000_.Start = __Rect__00000002_
	__Link__00000000_.End = __Rect__00000000_
	__Link__00000000_.TextAtArrowStart = append(__Link__00000000_.TextAtArrowStart, __LinkAnchoredText__00000003_)
	__Link__00000001_.Start = __Rect__00000002_
	__Link__00000001_.End = __Rect__00000001_
	__Link__00000001_.TextAtArrowStart = append(__Link__00000001_.TextAtArrowStart, __LinkAnchoredText__00000004_)
	__Link__00000001_.TextAtArrowStart = append(__Link__00000001_.TextAtArrowStart, __LinkAnchoredText__00000005_)
	__Link__00000001_.TextAtArrowEnd = append(__Link__00000001_.TextAtArrowEnd, __LinkAnchoredText__00000001_)
	__Link__00000001_.TextAtArrowEnd = append(__Link__00000001_.TextAtArrowEnd, __LinkAnchoredText__00000000_)
	__Link__00000002_.Start = __Rect__00000002_
	__Link__00000002_.End = __Rect__00000001_
	__Link__00000002_.ControlPoints = append(__Link__00000002_.ControlPoints, __ControlPoint__00000000_)
	__Link__00000003_.Start = __Rect__00000003_
	__Link__00000003_.End = __Rect__00000004_
	__Link__00000003_.TextAtArrowStart = append(__Link__00000003_.TextAtArrowStart, __LinkAnchoredText__00000009_)
	__Link__00000003_.TextAtArrowStart = append(__Link__00000003_.TextAtArrowStart, __LinkAnchoredText__00000008_)
	__Link__00000003_.TextAtArrowEnd = append(__Link__00000003_.TextAtArrowEnd, __LinkAnchoredText__00000009_)
	__Link__00000003_.TextAtArrowEnd = append(__Link__00000003_.TextAtArrowEnd, __LinkAnchoredText__00000008_)
	__Link__00000004_.Start = __Rect__00000004_
	__Link__00000004_.End = __Rect__00000003_
	__Link__00000004_.TextAtArrowStart = append(__Link__00000004_.TextAtArrowStart, __LinkAnchoredText__00000009_)
	__Link__00000004_.TextAtArrowStart = append(__Link__00000004_.TextAtArrowStart, __LinkAnchoredText__00000008_)
	__Link__00000004_.TextAtArrowStart = append(__Link__00000004_.TextAtArrowStart, __LinkAnchoredText__00000010_)
	__Link__00000004_.TextAtArrowStart = append(__Link__00000004_.TextAtArrowStart, __LinkAnchoredText__00000011_)
	__Link__00000004_.TextAtArrowEnd = append(__Link__00000004_.TextAtArrowEnd, __LinkAnchoredText__00000009_)
	__Link__00000004_.TextAtArrowEnd = append(__Link__00000004_.TextAtArrowEnd, __LinkAnchoredText__00000008_)
	__Link__00000004_.TextAtArrowEnd = append(__Link__00000004_.TextAtArrowEnd, __LinkAnchoredText__00000010_)
	__Link__00000004_.TextAtArrowEnd = append(__Link__00000004_.TextAtArrowEnd, __LinkAnchoredText__00000011_)
	__Link__00000005_.Start = __Rect__00000003_
	__Link__00000005_.End = __Rect__00000004_
	__Link__00000005_.TextAtArrowStart = append(__Link__00000005_.TextAtArrowStart, __LinkAnchoredText__00000009_)
	__Link__00000005_.TextAtArrowStart = append(__Link__00000005_.TextAtArrowStart, __LinkAnchoredText__00000008_)
	__Link__00000005_.TextAtArrowEnd = append(__Link__00000005_.TextAtArrowEnd, __LinkAnchoredText__00000009_)
	__Link__00000005_.TextAtArrowEnd = append(__Link__00000005_.TextAtArrowEnd, __LinkAnchoredText__00000008_)
	__Link__00000006_.Start = __Rect__00000001_
	__Link__00000006_.End = __Rect__00000000_
	__Rect__00000000_.RectAnchoredTexts = append(__Rect__00000000_.RectAnchoredTexts, __RectAnchoredText__00000000_)
	__Rect__00000000_.RectAnchoredRects = append(__Rect__00000000_.RectAnchoredRects, __RectAnchoredRect__00000001_)
	__Rect__00000000_.RectAnchoredPaths = append(__Rect__00000000_.RectAnchoredPaths, __RectAnchoredPath__00000001_)
	__Rect__00000000_.RectAnchoredPaths = append(__Rect__00000000_.RectAnchoredPaths, __RectAnchoredPath__00000001_)
	__Rect__00000002_.RectAnchoredRects = append(__Rect__00000002_.RectAnchoredRects, __RectAnchoredRect__00000000_)
	__Rect__00000002_.RectAnchoredPaths = append(__Rect__00000002_.RectAnchoredPaths, __RectAnchoredPath__00000000_)
	__Rect__00000003_.RectAnchoredTexts = append(__Rect__00000003_.RectAnchoredTexts, __RectAnchoredText__00000004_)
	__Rect__00000004_.RectAnchoredTexts = append(__Rect__00000004_.RectAnchoredTexts, __RectAnchoredText__00000005_)
	__SVG__00000000_.Layers = append(__SVG__00000000_.Layers, __Layer__00000000_)
	__SVG__00000000_.Layers = append(__SVG__00000000_.Layers, __Layer__00000001_)
	__SVG__00000000_.Layers = append(__SVG__00000000_.Layers, __Layer__00000002_)
	__SVG__00000000_.Layers = append(__SVG__00000000_.Layers, __Layer__00000003_)
	__SVG__00000000_.Layers = append(__SVG__00000000_.Layers, __Layer__00000004_)
	__SVG__00000000_.Layers = append(__SVG__00000000_.Layers, __Layer__00000006_)
	__SVG__00000000_.Layers = append(__SVG__00000000_.Layers, __Layer__00000005_)
	__SVG__00000000_.StartRect = __Rect__00000000_
	__SVG__00000000_.EndRect = __Rect__00000002_
}
