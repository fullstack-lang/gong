package main

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/svg/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time
var _ = slices.Index[[]int, int]

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage

	__ControlPoint__00000000_ := (&models.ControlPoint{Name: ``}).Stage(stage)

	__Layer__00000000_ := (&models.Layer{Name: ``}).Stage(stage)

	__Link__00000000_ := (&models.Link{Name: `Link of interest`}).Stage(stage)

	__Rect__00000000_ := (&models.Rect{Name: `Start`}).Stage(stage)
	__Rect__00000001_ := (&models.Rect{Name: `End`}).Stage(stage)

	__SVG__00000000_ := (&models.SVG{Name: `issue 1216`}).Stage(stage)

	// insertion point for initialization of values

	__ControlPoint__00000000_.Name = ``
	__ControlPoint__00000000_.X_Relative = 1.100000
	__ControlPoint__00000000_.Y_Relative = 1.500000

	__Layer__00000000_.Name = ``

	__Link__00000000_.Name = `Link of interest`
	__Link__00000000_.Type = models.LINK_TYPE_LINE_WITH_CONTROL_POINTS
	__Link__00000000_.IsBezierCurve = false
	__Link__00000000_.StartAnchorType = models.ANCHOR_CENTER
	__Link__00000000_.EndAnchorType = models.ANCHOR_CENTER
	__Link__00000000_.StartOrientation = ""
	__Link__00000000_.StartRatio = 0.000000
	__Link__00000000_.EndOrientation = ""
	__Link__00000000_.EndRatio = 0.000000
	__Link__00000000_.CornerOffsetRatio = 0.000000
	__Link__00000000_.CornerRadius = 0.000000
	__Link__00000000_.HasEndArrow = true
	__Link__00000000_.EndArrowSize = 5.000000
	__Link__00000000_.EndArrowOffset = 20.000000
	__Link__00000000_.HasStartArrow = true
	__Link__00000000_.StartArrowSize = 5.000000
	__Link__00000000_.StartArrowOffset = 10.000000
	__Link__00000000_.Color = ``
	__Link__00000000_.FillOpacity = 0.000000
	__Link__00000000_.Stroke = `black`
	__Link__00000000_.StrokeOpacity = 1.000000
	__Link__00000000_.StrokeWidth = 1.000000
	__Link__00000000_.StrokeDashArray = ``
	__Link__00000000_.StrokeDashArrayWhenSelected = ``
	__Link__00000000_.Transform = ``
	__Link__00000000_.MouseX = 0.000000
	__Link__00000000_.MouseY = 0.000000
	__Link__00000000_.MouseEventKey = ""

	__Rect__00000000_.Name = `Start`
	__Rect__00000000_.X = 10.000000
	__Rect__00000000_.Y = 10.000000
	__Rect__00000000_.Width = 200.000000
	__Rect__00000000_.Height = 70.000000
	__Rect__00000000_.RX = 3.000000
	__Rect__00000000_.Color = ``
	__Rect__00000000_.FillOpacity = 0.000000
	__Rect__00000000_.Stroke = `black`
	__Rect__00000000_.StrokeOpacity = 1.000000
	__Rect__00000000_.StrokeWidth = 1.000000
	__Rect__00000000_.StrokeDashArray = ``
	__Rect__00000000_.StrokeDashArrayWhenSelected = ``
	__Rect__00000000_.Transform = ``
	__Rect__00000000_.IsSelectable = false
	__Rect__00000000_.IsSelected = false
	__Rect__00000000_.CanHaveLeftHandle = false
	__Rect__00000000_.HasLeftHandle = false
	__Rect__00000000_.CanHaveRightHandle = false
	__Rect__00000000_.HasRightHandle = false
	__Rect__00000000_.CanHaveTopHandle = false
	__Rect__00000000_.HasTopHandle = false
	__Rect__00000000_.IsScalingProportionally = false
	__Rect__00000000_.CanHaveBottomHandle = false
	__Rect__00000000_.HasBottomHandle = false
	__Rect__00000000_.CanMoveHorizontaly = false
	__Rect__00000000_.CanMoveVerticaly = false
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

	__Rect__00000001_.Name = `End`
	__Rect__00000001_.X = 200.000000
	__Rect__00000001_.Y = 170.000000
	__Rect__00000001_.Width = 200.000000
	__Rect__00000001_.Height = 70.000000
	__Rect__00000001_.RX = 4.000000
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

	__SVG__00000000_.Name = `issue 1216`
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
	__ControlPoint__00000000_.ClosestRect = __Rect__00000000_
	__Layer__00000000_.Rects = append(__Layer__00000000_.Rects, __Rect__00000000_)
	__Layer__00000000_.Rects = append(__Layer__00000000_.Rects, __Rect__00000001_)
	__Layer__00000000_.Links = append(__Layer__00000000_.Links, __Link__00000000_)
	__Link__00000000_.Start = __Rect__00000000_
	__Link__00000000_.End = __Rect__00000001_
	__Link__00000000_.ControlPoints = append(__Link__00000000_.ControlPoints, __ControlPoint__00000000_)
	__SVG__00000000_.Layers = append(__SVG__00000000_.Layers, __Layer__00000000_)
	__SVG__00000000_.StartRect = nil
	__SVG__00000000_.EndRect = nil
}
