package main

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/svg/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var (
	_ time.Time
	_ = slices.Index[[]int, int]
)

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage

	__Layer__00000000_ := (&models.Layer{Name: ``}).Stage(stage)

	__Rect__00000000_ := (&models.Rect{Name: `Application Engineering`}).Stage(stage)
	__Rect__00000001_ := (&models.Rect{Name: `Domain Engineering`}).Stage(stage)

	__RectAnchoredText__00000000_ := (&models.RectAnchoredText{Name: `Application engineering`}).Stage(stage)
	__RectAnchoredText__00000001_ := (&models.RectAnchoredText{Name: `Domain Engineering`}).Stage(stage)

	__SVG__00000000_ := (&models.SVG{Name: ``}).Stage(stage)

	// insertion point for initialization of values

	__Layer__00000000_.Name = ``

	__Rect__00000000_.Name = `Application Engineering`
	__Rect__00000000_.X = 400.000000
	__Rect__00000000_.Y = 100.000000
	__Rect__00000000_.Width = 198.000000
	__Rect__00000000_.Height = 50.000000
	__Rect__00000000_.RX = 0.000000
	__Rect__00000000_.Color = ``
	__Rect__00000000_.FillOpacity = 0.000000
	__Rect__00000000_.Stroke = `black`
	__Rect__00000000_.StrokeOpacity = 1.000000
	__Rect__00000000_.StrokeWidth = 1.000000
	__Rect__00000000_.StrokeDashArray = ``
	__Rect__00000000_.StrokeDashArrayWhenSelected = ``
	__Rect__00000000_.Transform = ``
	__Rect__00000000_.IsSelectable = true
	__Rect__00000000_.IsSelected = false
	__Rect__00000000_.CanHaveLeftHandle = true
	__Rect__00000000_.HasLeftHandle = false
	__Rect__00000000_.CanHaveRightHandle = true
	__Rect__00000000_.HasRightHandle = false
	__Rect__00000000_.CanHaveTopHandle = true
	__Rect__00000000_.HasTopHandle = false
	__Rect__00000000_.IsScalingProportionally = true
	__Rect__00000000_.CanHaveBottomHandle = true
	__Rect__00000000_.HasBottomHandle = false
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
	__Rect__00000000_.URLPath = ``
	__Rect__00000000_.URLTarget = ""

	__Rect__00000001_.Name = `Domain Engineering`
	__Rect__00000001_.X = 700.000000
	__Rect__00000001_.Y = 100.000000
	__Rect__00000001_.Width = 200.000000
	__Rect__00000001_.Height = 500.000000
	__Rect__00000001_.RX = 0.000000
	__Rect__00000001_.Color = ``
	__Rect__00000001_.FillOpacity = 0.000000
	__Rect__00000001_.Stroke = `black`
	__Rect__00000001_.StrokeOpacity = 1.000000
	__Rect__00000001_.StrokeWidth = 1.000000
	__Rect__00000001_.StrokeDashArray = ``
	__Rect__00000001_.StrokeDashArrayWhenSelected = ``
	__Rect__00000001_.Transform = ``
	__Rect__00000001_.IsSelectable = false
	__Rect__00000001_.IsSelected = false
	__Rect__00000001_.CanHaveLeftHandle = false
	__Rect__00000001_.HasLeftHandle = false
	__Rect__00000001_.CanHaveRightHandle = false
	__Rect__00000001_.HasRightHandle = false
	__Rect__00000001_.CanHaveTopHandle = false
	__Rect__00000001_.HasTopHandle = false
	__Rect__00000001_.IsScalingProportionally = false
	__Rect__00000001_.CanHaveBottomHandle = false
	__Rect__00000001_.HasBottomHandle = false
	__Rect__00000001_.CanMoveHorizontaly = false
	__Rect__00000001_.CanMoveVerticaly = false
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
	__Rect__00000001_.URLPath = ``
	__Rect__00000001_.URLTarget = ""

	__RectAnchoredText__00000000_.Name = `Application engineering`
	__RectAnchoredText__00000000_.Content = `Application
engineering`
	__RectAnchoredText__00000000_.FontWeight = ``
	__RectAnchoredText__00000000_.FontSize = ``
	__RectAnchoredText__00000000_.FontStyle = ``
	__RectAnchoredText__00000000_.LetterSpacing = ``
	__RectAnchoredText__00000000_.FontFamily = ``
	__RectAnchoredText__00000000_.WhiteSpace = ""
	__RectAnchoredText__00000000_.X_Offset = 0.000000
	__RectAnchoredText__00000000_.Y_Offset = 0.000000
	__RectAnchoredText__00000000_.RectAnchorType = models.RECT_CENTER_MIDDLE
	__RectAnchoredText__00000000_.TextAnchorType = models.TEXT_ANCHOR_CENTER
	__RectAnchoredText__00000000_.DominantBaseline = ""
	__RectAnchoredText__00000000_.WritingMode = ""
	__RectAnchoredText__00000000_.Color = `black`
	__RectAnchoredText__00000000_.FillOpacity = 1.000000
	__RectAnchoredText__00000000_.Stroke = `black`
	__RectAnchoredText__00000000_.StrokeOpacity = 1.000000
	__RectAnchoredText__00000000_.StrokeWidth = 1.000000
	__RectAnchoredText__00000000_.StrokeDashArray = ``
	__RectAnchoredText__00000000_.StrokeDashArrayWhenSelected = ``
	__RectAnchoredText__00000000_.Transform = ``
	__RectAnchoredText__00000000_.URLPath = ``
	__RectAnchoredText__00000000_.URLTarget = ""

	__RectAnchoredText__00000001_.Name = `Domain Engineering`
	__RectAnchoredText__00000001_.Content = `Domain
Engineering`
	__RectAnchoredText__00000001_.FontWeight = ``
	__RectAnchoredText__00000001_.FontSize = ``
	__RectAnchoredText__00000001_.FontStyle = ``
	__RectAnchoredText__00000001_.LetterSpacing = ``
	__RectAnchoredText__00000001_.FontFamily = ``
	__RectAnchoredText__00000001_.WhiteSpace = ""
	__RectAnchoredText__00000001_.X_Offset = 0.000000
	__RectAnchoredText__00000001_.Y_Offset = 0.000000
	__RectAnchoredText__00000001_.RectAnchorType = models.RECT_CENTER_MIDDLE
	__RectAnchoredText__00000001_.TextAnchorType = models.TEXT_ANCHOR_CENTER
	__RectAnchoredText__00000001_.DominantBaseline = ""
	__RectAnchoredText__00000001_.WritingMode = ""
	__RectAnchoredText__00000001_.Color = `black`
	__RectAnchoredText__00000001_.FillOpacity = 1.000000
	__RectAnchoredText__00000001_.Stroke = `black`
	__RectAnchoredText__00000001_.StrokeOpacity = 1.000000
	__RectAnchoredText__00000001_.StrokeWidth = 1.000000
	__RectAnchoredText__00000001_.StrokeDashArray = ``
	__RectAnchoredText__00000001_.StrokeDashArrayWhenSelected = ``
	__RectAnchoredText__00000001_.Transform = ``
	__RectAnchoredText__00000001_.URLPath = ``
	__RectAnchoredText__00000001_.URLTarget = ""

	__SVG__00000000_.Name = ``
	__SVG__00000000_.DrawingState = ""
	__SVG__00000000_.IsEditable = true
	__SVG__00000000_.IsSVGFrontEndFileGenerated = false
	__SVG__00000000_.IsSVGBackEndFileGenerated = false
	__SVG__00000000_.DefaultDirectoryForGeneratedImages = ``
	__SVG__00000000_.IsControlBannerHidden = false
	__SVG__00000000_.OverrideWidth = false
	__SVG__00000000_.OverriddenWidth = 0.000000
	__SVG__00000000_.OverrideHeight = false
	__SVG__00000000_.OverriddenHeight = 0.000000

	// insertion point for setup of pointers
	__Layer__00000000_.Rects = append(__Layer__00000000_.Rects, __Rect__00000000_)
	__Layer__00000000_.Rects = append(__Layer__00000000_.Rects, __Rect__00000001_)
	__Rect__00000000_.EnclosingRect = nil
	__Rect__00000000_.AnchoredTo = nil
	__Rect__00000000_.RectAnchoredTexts = append(__Rect__00000000_.RectAnchoredTexts, __RectAnchoredText__00000000_)
	__Rect__00000001_.EnclosingRect = nil
	__Rect__00000001_.AnchoredTo = nil
	__Rect__00000001_.RectAnchoredTexts = append(__Rect__00000001_.RectAnchoredTexts, __RectAnchoredText__00000001_)
	__SVG__00000000_.Layers = append(__SVG__00000000_.Layers, __Layer__00000000_)
	__SVG__00000000_.StartRect = nil
	__SVG__00000000_.EndRect = nil
}
