// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

var _ = __GONG_time_The_fool_doth_think_he_is_wise__

// insertion point
type Category1_WOP struct {
	// insertion point

	Name string
}

func (from *Category1) CopyBasicFields(to *Category1) {
	// insertion point
	to.Name = from.Name
}

type Category1Shape_WOP struct {
	// insertion point

	Name string

	X float64

	Y float64

	Width float64

	Height float64
}

func (from *Category1Shape) CopyBasicFields(to *Category1Shape) {
	// insertion point
	to.Name = from.Name
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
}

type Category2_WOP struct {
	// insertion point

	Name string
}

func (from *Category2) CopyBasicFields(to *Category2) {
	// insertion point
	to.Name = from.Name
}

type Category2Shape_WOP struct {
	// insertion point

	Name string

	X float64

	Y float64

	Width float64

	Height float64
}

func (from *Category2Shape) CopyBasicFields(to *Category2Shape) {
	// insertion point
	to.Name = from.Name
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
}

type Category3_WOP struct {
	// insertion point

	Name string
}

func (from *Category3) CopyBasicFields(to *Category3) {
	// insertion point
	to.Name = from.Name
}

type Category3Shape_WOP struct {
	// insertion point

	Name string

	X float64

	Y float64

	Width float64

	Height float64
}

func (from *Category3Shape) CopyBasicFields(to *Category3Shape) {
	// insertion point
	to.Name = from.Name
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
}

type ControlPointShape_WOP struct {
	// insertion point

	Name string

	X_Relative float64

	Y_Relative float64

	IsStartShapeTheClosestShape bool
}

func (from *ControlPointShape) CopyBasicFields(to *ControlPointShape) {
	// insertion point
	to.Name = from.Name
	to.X_Relative = from.X_Relative
	to.Y_Relative = from.Y_Relative
	to.IsStartShapeTheClosestShape = from.IsStartShapeTheClosestShape
}

type Desk_WOP struct {
	// insertion point

	Name string
}

func (from *Desk) CopyBasicFields(to *Desk) {
	// insertion point
	to.Name = from.Name
}

type Diagram_WOP struct {
	// insertion point

	Name string

	IsEditable bool

	IsNodeExpanded bool

	IsCategory1NodeExpanded bool

	IsCategory2NodeExpanded bool

	IsCategory3NodeExpanded bool

	IsInfluenceCategoryNodeExpanded bool

	IsCategory1Shown bool

	IsCategory2Shown bool

	IsCategory3Shown bool

	IsInfluenceCategoryShown bool

	XMargin float64

	YMargin float64

	Height float64

	Width float64

	RedColorCode string

	BackgroundGreyColorCode string

	GrayColorCode string

	Category1RectAnchorType RectAnchorType

	Category1TextAnchorType TextAnchorType

	Category1DominantBaselineType DominantBaselineType

	Category1FontSize string

	Category1FontWeigth string

	Category1FontFamily string

	Category1LetterSpacing string

	Category2TypeFontSize string

	Category2TypeFontWeigth string

	Category2TypeFontFamily string

	Category2TypeLetterSpacing string

	Category2TypeRectAnchorType RectAnchorType

	Category2DominantBaselineType DominantBaselineType

	Category2StrokeWidth float64

	Category3RectAnchorType RectAnchorType

	Category3TextAnchorType TextAnchorType

	Category3DominantBaselineType DominantBaselineType

	Category3FontSize string

	Category3FontWeigth string

	Category3FontFamily string

	Category3LetterSpacing string

	InfluenceArrowSize float64

	InfluenceArrowStartOffset float64

	InfluenceArrowEndOffset float64

	InfluenceCornerRadius float64

	InfluenceDashedLinePattern string
}

func (from *Diagram) CopyBasicFields(to *Diagram) {
	// insertion point
	to.Name = from.Name
	to.IsEditable = from.IsEditable
	to.IsNodeExpanded = from.IsNodeExpanded
	to.IsCategory1NodeExpanded = from.IsCategory1NodeExpanded
	to.IsCategory2NodeExpanded = from.IsCategory2NodeExpanded
	to.IsCategory3NodeExpanded = from.IsCategory3NodeExpanded
	to.IsInfluenceCategoryNodeExpanded = from.IsInfluenceCategoryNodeExpanded
	to.IsCategory1Shown = from.IsCategory1Shown
	to.IsCategory2Shown = from.IsCategory2Shown
	to.IsCategory3Shown = from.IsCategory3Shown
	to.IsInfluenceCategoryShown = from.IsInfluenceCategoryShown
	to.XMargin = from.XMargin
	to.YMargin = from.YMargin
	to.Height = from.Height
	to.Width = from.Width
	to.RedColorCode = from.RedColorCode
	to.BackgroundGreyColorCode = from.BackgroundGreyColorCode
	to.GrayColorCode = from.GrayColorCode
	to.Category1RectAnchorType = from.Category1RectAnchorType
	to.Category1TextAnchorType = from.Category1TextAnchorType
	to.Category1DominantBaselineType = from.Category1DominantBaselineType
	to.Category1FontSize = from.Category1FontSize
	to.Category1FontWeigth = from.Category1FontWeigth
	to.Category1FontFamily = from.Category1FontFamily
	to.Category1LetterSpacing = from.Category1LetterSpacing
	to.Category2TypeFontSize = from.Category2TypeFontSize
	to.Category2TypeFontWeigth = from.Category2TypeFontWeigth
	to.Category2TypeFontFamily = from.Category2TypeFontFamily
	to.Category2TypeLetterSpacing = from.Category2TypeLetterSpacing
	to.Category2TypeRectAnchorType = from.Category2TypeRectAnchorType
	to.Category2DominantBaselineType = from.Category2DominantBaselineType
	to.Category2StrokeWidth = from.Category2StrokeWidth
	to.Category3RectAnchorType = from.Category3RectAnchorType
	to.Category3TextAnchorType = from.Category3TextAnchorType
	to.Category3DominantBaselineType = from.Category3DominantBaselineType
	to.Category3FontSize = from.Category3FontSize
	to.Category3FontWeigth = from.Category3FontWeigth
	to.Category3FontFamily = from.Category3FontFamily
	to.Category3LetterSpacing = from.Category3LetterSpacing
	to.InfluenceArrowSize = from.InfluenceArrowSize
	to.InfluenceArrowStartOffset = from.InfluenceArrowStartOffset
	to.InfluenceArrowEndOffset = from.InfluenceArrowEndOffset
	to.InfluenceCornerRadius = from.InfluenceCornerRadius
	to.InfluenceDashedLinePattern = from.InfluenceDashedLinePattern
}

type Influence_WOP struct {
	// insertion point

	Name string

	IsHypothtical bool
}

func (from *Influence) CopyBasicFields(to *Influence) {
	// insertion point
	to.Name = from.Name
	to.IsHypothtical = from.IsHypothtical
}

type InfluenceShape_WOP struct {
	// insertion point

	Name string
}

func (from *InfluenceShape) CopyBasicFields(to *InfluenceShape) {
	// insertion point
	to.Name = from.Name
}

