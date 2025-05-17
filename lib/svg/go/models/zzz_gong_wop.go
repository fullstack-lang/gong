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
	From string
	To string
	Dur string
	RepeatCount string
}

func (from *Animate) CopyBasicFields(to *Animate) {
	// insertion point
	to.Name = from.Name
	to.AttributeName = from.AttributeName
	to.Values = from.Values
	to.From = from.From
	to.To = from.To
	to.Dur = from.Dur
	to.RepeatCount = from.RepeatCount
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
	StrokeOpacity float64
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

func (from *Circle) CopyBasicFields(to *Circle) {
	// insertion point
	to.Name = from.Name
	to.CX = from.CX
	to.CY = from.CY
	to.Radius = from.Radius
	to.Color = from.Color
	to.FillOpacity = from.FillOpacity
	to.Stroke = from.Stroke
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
	to.StrokeDashArrayWhenSelected = from.StrokeDashArrayWhenSelected
	to.Transform = from.Transform
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
	StrokeOpacity float64
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

func (from *Ellipse) CopyBasicFields(to *Ellipse) {
	// insertion point
	to.Name = from.Name
	to.CX = from.CX
	to.CY = from.CY
	to.RX = from.RX
	to.RY = from.RY
	to.Color = from.Color
	to.FillOpacity = from.FillOpacity
	to.Stroke = from.Stroke
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
	to.StrokeDashArrayWhenSelected = from.StrokeDashArrayWhenSelected
	to.Transform = from.Transform
}

type Layer_WOP struct {
	// insertion point
	Display bool
	Name string
}

func (from *Layer) CopyBasicFields(to *Layer) {
	// insertion point
	to.Display = from.Display
	to.Name = from.Name
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
	StrokeOpacity float64
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
	MouseClickX float64
	MouseClickY float64
}

func (from *Line) CopyBasicFields(to *Line) {
	// insertion point
	to.Name = from.Name
	to.X1 = from.X1
	to.Y1 = from.Y1
	to.X2 = from.X2
	to.Y2 = from.Y2
	to.Color = from.Color
	to.FillOpacity = from.FillOpacity
	to.Stroke = from.Stroke
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
	to.StrokeDashArrayWhenSelected = from.StrokeDashArrayWhenSelected
	to.Transform = from.Transform
	to.MouseClickX = from.MouseClickX
	to.MouseClickY = from.MouseClickY
}

type Link_WOP struct {
	// insertion point
	Name string
	Type LinkType
	IsBezierCurve bool
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
	StrokeOpacity float64
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

func (from *Link) CopyBasicFields(to *Link) {
	// insertion point
	to.Name = from.Name
	to.Type = from.Type
	to.IsBezierCurve = from.IsBezierCurve
	to.StartAnchorType = from.StartAnchorType
	to.EndAnchorType = from.EndAnchorType
	to.StartOrientation = from.StartOrientation
	to.StartRatio = from.StartRatio
	to.EndOrientation = from.EndOrientation
	to.EndRatio = from.EndRatio
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.CornerRadius = from.CornerRadius
	to.HasEndArrow = from.HasEndArrow
	to.EndArrowSize = from.EndArrowSize
	to.HasStartArrow = from.HasStartArrow
	to.StartArrowSize = from.StartArrowSize
	to.Color = from.Color
	to.FillOpacity = from.FillOpacity
	to.Stroke = from.Stroke
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
	to.StrokeDashArrayWhenSelected = from.StrokeDashArrayWhenSelected
	to.Transform = from.Transform
}

type LinkAnchoredText_WOP struct {
	// insertion point
	Name string
	Content string
	AutomaticLayout bool
	LinkAnchorType LinkAnchorType
	X_Offset float64
	Y_Offset float64
	FontWeight string
	FontSize string
	LetterSpacing string
	Color string
	FillOpacity float64
	Stroke string
	StrokeOpacity float64
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

func (from *LinkAnchoredText) CopyBasicFields(to *LinkAnchoredText) {
	// insertion point
	to.Name = from.Name
	to.Content = from.Content
	to.AutomaticLayout = from.AutomaticLayout
	to.LinkAnchorType = from.LinkAnchorType
	to.X_Offset = from.X_Offset
	to.Y_Offset = from.Y_Offset
	to.FontWeight = from.FontWeight
	to.FontSize = from.FontSize
	to.LetterSpacing = from.LetterSpacing
	to.Color = from.Color
	to.FillOpacity = from.FillOpacity
	to.Stroke = from.Stroke
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
	to.StrokeDashArrayWhenSelected = from.StrokeDashArrayWhenSelected
	to.Transform = from.Transform
}

type Path_WOP struct {
	// insertion point
	Name string
	Definition string
	Color string
	FillOpacity float64
	Stroke string
	StrokeOpacity float64
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

func (from *Path) CopyBasicFields(to *Path) {
	// insertion point
	to.Name = from.Name
	to.Definition = from.Definition
	to.Color = from.Color
	to.FillOpacity = from.FillOpacity
	to.Stroke = from.Stroke
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
	to.StrokeDashArrayWhenSelected = from.StrokeDashArrayWhenSelected
	to.Transform = from.Transform
}

type Point_WOP struct {
	// insertion point
	Name string
	X float64
	Y float64
}

func (from *Point) CopyBasicFields(to *Point) {
	// insertion point
	to.Name = from.Name
	to.X = from.X
	to.Y = from.Y
}

type Polygone_WOP struct {
	// insertion point
	Name string
	Points string
	Color string
	FillOpacity float64
	Stroke string
	StrokeOpacity float64
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

func (from *Polygone) CopyBasicFields(to *Polygone) {
	// insertion point
	to.Name = from.Name
	to.Points = from.Points
	to.Color = from.Color
	to.FillOpacity = from.FillOpacity
	to.Stroke = from.Stroke
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
	to.StrokeDashArrayWhenSelected = from.StrokeDashArrayWhenSelected
	to.Transform = from.Transform
}

type Polyline_WOP struct {
	// insertion point
	Name string
	Points string
	Color string
	FillOpacity float64
	Stroke string
	StrokeOpacity float64
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

func (from *Polyline) CopyBasicFields(to *Polyline) {
	// insertion point
	to.Name = from.Name
	to.Points = from.Points
	to.Color = from.Color
	to.FillOpacity = from.FillOpacity
	to.Stroke = from.Stroke
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
	to.StrokeDashArrayWhenSelected = from.StrokeDashArrayWhenSelected
	to.Transform = from.Transform
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
	StrokeOpacity float64
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
	IsScalingProportionally bool
	CanHaveBottomHandle bool
	HasBottomHandle bool
	CanMoveHorizontaly bool
	CanMoveVerticaly bool
}

func (from *Rect) CopyBasicFields(to *Rect) {
	// insertion point
	to.Name = from.Name
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
	to.RX = from.RX
	to.Color = from.Color
	to.FillOpacity = from.FillOpacity
	to.Stroke = from.Stroke
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
	to.StrokeDashArrayWhenSelected = from.StrokeDashArrayWhenSelected
	to.Transform = from.Transform
	to.IsSelectable = from.IsSelectable
	to.IsSelected = from.IsSelected
	to.CanHaveLeftHandle = from.CanHaveLeftHandle
	to.HasLeftHandle = from.HasLeftHandle
	to.CanHaveRightHandle = from.CanHaveRightHandle
	to.HasRightHandle = from.HasRightHandle
	to.CanHaveTopHandle = from.CanHaveTopHandle
	to.HasTopHandle = from.HasTopHandle
	to.IsScalingProportionally = from.IsScalingProportionally
	to.CanHaveBottomHandle = from.CanHaveBottomHandle
	to.HasBottomHandle = from.HasBottomHandle
	to.CanMoveHorizontaly = from.CanMoveHorizontaly
	to.CanMoveVerticaly = from.CanMoveVerticaly
}

type RectAnchoredPath_WOP struct {
	// insertion point
	Name string
	Definition string
	X_Offset float64
	Y_Offset float64
	RectAnchorType RectAnchorType
	ScalePropotionnally bool
	AppliedScaling float64
	Color string
	FillOpacity float64
	Stroke string
	StrokeOpacity float64
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

func (from *RectAnchoredPath) CopyBasicFields(to *RectAnchoredPath) {
	// insertion point
	to.Name = from.Name
	to.Definition = from.Definition
	to.X_Offset = from.X_Offset
	to.Y_Offset = from.Y_Offset
	to.RectAnchorType = from.RectAnchorType
	to.ScalePropotionnally = from.ScalePropotionnally
	to.AppliedScaling = from.AppliedScaling
	to.Color = from.Color
	to.FillOpacity = from.FillOpacity
	to.Stroke = from.Stroke
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
	to.StrokeDashArrayWhenSelected = from.StrokeDashArrayWhenSelected
	to.Transform = from.Transform
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
	StrokeOpacity float64
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

func (from *RectAnchoredRect) CopyBasicFields(to *RectAnchoredRect) {
	// insertion point
	to.Name = from.Name
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
	to.RX = from.RX
	to.X_Offset = from.X_Offset
	to.Y_Offset = from.Y_Offset
	to.RectAnchorType = from.RectAnchorType
	to.WidthFollowRect = from.WidthFollowRect
	to.HeightFollowRect = from.HeightFollowRect
	to.Color = from.Color
	to.FillOpacity = from.FillOpacity
	to.Stroke = from.Stroke
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
	to.StrokeDashArrayWhenSelected = from.StrokeDashArrayWhenSelected
	to.Transform = from.Transform
}

type RectAnchoredText_WOP struct {
	// insertion point
	Name string
	Content string
	FontWeight string
	FontSize int
	FontStyle string
	X_Offset float64
	Y_Offset float64
	RectAnchorType RectAnchorType
	TextAnchorType TextAnchorType
	Color string
	FillOpacity float64
	Stroke string
	StrokeOpacity float64
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

func (from *RectAnchoredText) CopyBasicFields(to *RectAnchoredText) {
	// insertion point
	to.Name = from.Name
	to.Content = from.Content
	to.FontWeight = from.FontWeight
	to.FontSize = from.FontSize
	to.FontStyle = from.FontStyle
	to.X_Offset = from.X_Offset
	to.Y_Offset = from.Y_Offset
	to.RectAnchorType = from.RectAnchorType
	to.TextAnchorType = from.TextAnchorType
	to.Color = from.Color
	to.FillOpacity = from.FillOpacity
	to.Stroke = from.Stroke
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
	to.StrokeDashArrayWhenSelected = from.StrokeDashArrayWhenSelected
	to.Transform = from.Transform
}

type RectLinkLink_WOP struct {
	// insertion point
	Name string
	TargetAnchorPosition float64
	Color string
	FillOpacity float64
	Stroke string
	StrokeOpacity float64
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

func (from *RectLinkLink) CopyBasicFields(to *RectLinkLink) {
	// insertion point
	to.Name = from.Name
	to.TargetAnchorPosition = from.TargetAnchorPosition
	to.Color = from.Color
	to.FillOpacity = from.FillOpacity
	to.Stroke = from.Stroke
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
	to.StrokeDashArrayWhenSelected = from.StrokeDashArrayWhenSelected
	to.Transform = from.Transform
}

type SVG_WOP struct {
	// insertion point
	Name string
	DrawingState DrawingState
	IsEditable bool
	IsSVGFileGenerated bool
}

func (from *SVG) CopyBasicFields(to *SVG) {
	// insertion point
	to.Name = from.Name
	to.DrawingState = from.DrawingState
	to.IsEditable = from.IsEditable
	to.IsSVGFileGenerated = from.IsSVGFileGenerated
}

type SvgText_WOP struct {
	// insertion point
	Name string
	Text string
}

func (from *SvgText) CopyBasicFields(to *SvgText) {
	// insertion point
	to.Name = from.Name
	to.Text = from.Text
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
	StrokeOpacity float64
	StrokeWidth float64
	StrokeDashArray string
	StrokeDashArrayWhenSelected string
	Transform string
}

func (from *Text) CopyBasicFields(to *Text) {
	// insertion point
	to.Name = from.Name
	to.X = from.X
	to.Y = from.Y
	to.Content = from.Content
	to.Color = from.Color
	to.FillOpacity = from.FillOpacity
	to.Stroke = from.Stroke
	to.StrokeOpacity = from.StrokeOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
	to.StrokeDashArrayWhenSelected = from.StrokeDashArrayWhenSelected
	to.Transform = from.Transform
}

