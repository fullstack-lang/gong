// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

// insertion point
type Animate_WOP struct {
	// insertion point
	Name string
	AttributeName string
	Values string
	Dur string
	RepeatCount string
}

type Circle_WOP struct {
	// insertion point
	Name string
	CX float64
	CY float64
	Radius float64
	Color string
	FillOpacity float64
	Stroke string
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

type Ellipse_WOP struct {
	// insertion point
	Name string
	CX float64
	CY float64
	RX float64
	RY float64
	Color string
	FillOpacity float64
	Stroke string
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

type Layer_WOP struct {
	// insertion point
	Display bool
	Name string
}

type Line_WOP struct {
	// insertion point
	Name string
	X1 float64
	Y1 float64
	X2 float64
	Y2 float64
	Color string
	FillOpacity float64
	Stroke string
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
	MouseClickX float64
	MouseClickY float64
}

type Link_WOP struct {
	// insertion point
	Name string
	Type LinkType
	StartAnchorType AnchorType
	EndAnchorType AnchorType
	StartOrientation OrientationType
	StartRatio float64
	EndOrientation OrientationType
	EndRatio float64
	CornerOffsetRatio float64
	CornerRadius float64
	HasEndArrow bool
	EndArrowSize float64
	HasStartArrow bool
	StartArrowSize float64
	Color string
	FillOpacity float64
	Stroke string
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

type LinkAnchoredText_WOP struct {
	// insertion point
	Name string
	Content string
	X_Offset float64
	Y_Offset float64
	FontWeight string
	Color string
	FillOpacity float64
	Stroke string
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

type Path_WOP struct {
	// insertion point
	Name string
	Definition string
	Color string
	FillOpacity float64
	Stroke string
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

type Point_WOP struct {
	// insertion point
	Name string
	X float64
	Y float64
}

type Polygone_WOP struct {
	// insertion point
	Name string
	Points string
	Color string
	FillOpacity float64
	Stroke string
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

type Polyline_WOP struct {
	// insertion point
	Name string
	Points string
	Color string
	FillOpacity float64
	Stroke string
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

type Rect_WOP struct {
	// insertion point
	Name string
	X float64
	Y float64
	Width float64
	Height float64
	RX float64
	Color string
	FillOpacity float64
	Stroke string
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
	IsSelectable bool
	IsSelected bool
	CanHaveLeftHandle bool
	HasLeftHandle bool
	CanHaveRightHandle bool
	HasRightHandle bool
	CanHaveTopHandle bool
	HasTopHandle bool
	CanHaveBottomHandle bool
	HasBottomHandle bool
	CanMoveHorizontaly bool
	CanMoveVerticaly bool
}

type RectAnchoredRect_WOP struct {
	// insertion point
	Name string
	X float64
	Y float64
	Width float64
	Height float64
	RX float64
	X_Offset float64
	Y_Offset float64
	RectAnchorType RectAnchorType
	WidthFollowRect bool
	HeightFollowRect bool
	Color string
	FillOpacity float64
	Stroke string
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

type RectAnchoredText_WOP struct {
	// insertion point
	Name string
	Content string
	FontWeight string
	FontSize int
	X_Offset float64
	Y_Offset float64
	RectAnchorType RectAnchorType
	TextAnchorType TextAnchorType
	Color string
	FillOpacity float64
	Stroke string
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

type RectLinkLink_WOP struct {
	// insertion point
	Name string
	TargetAnchorPosition float64
	Color string
	FillOpacity float64
	Stroke string
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

type SVG_WOP struct {
	// insertion point
	Name string
	DrawingState DrawingState
	IsEditable bool
}

type Text_WOP struct {
	// insertion point
	Name string
	X float64
	Y float64
	Content string
	Color string
	FillOpacity float64
	Stroke string
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

