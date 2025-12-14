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

	NextVerticalDateXMargin float64

	RedColorCode string

	BackgroundGreyColorCode string

	GrayColorCode string

	BottomBoxYOffset float64

	BottomBoxWidth float64

	BottomBoxHeigth float64

	BottomBoxFontSize string

	BottomBoxFontWeigth string

	BottomBoxFontFamily string

	BottomBoxLetterSpacing string

	BottomBoxLetterColorCode string

	MovementRectAnchorType RectAnchorType

	MovementTextAnchorType TextAnchorType

	MovementDominantBaselineType DominantBaselineType

	MovementFontSize string

	MovementFontWeigth string

	MovementFontFamily string

	MovementLetterSpacing string

	ArtefactTypeFontSize string

	ArtefactTypeFontWeigth string

	ArtefactTypeFontFamily string

	ArtefactTypeLetterSpacing string

	ArtefactTypeRectAnchorType RectAnchorType

	ArtefactDominantBaselineType DominantBaselineType

	ArtefactTypeStrokeWidth float64

	ArtistRectAnchorType RectAnchorType

	ArtistTextAnchorType TextAnchorType

	ArtistDominantBaselineType DominantBaselineType

	ArtistFontSize string

	ArtistFontWeigth string

	ArtistFontFamily string

	ArtistLetterSpacing string

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
	to.NextVerticalDateXMargin = from.NextVerticalDateXMargin
	to.RedColorCode = from.RedColorCode
	to.BackgroundGreyColorCode = from.BackgroundGreyColorCode
	to.GrayColorCode = from.GrayColorCode
	to.BottomBoxYOffset = from.BottomBoxYOffset
	to.BottomBoxWidth = from.BottomBoxWidth
	to.BottomBoxHeigth = from.BottomBoxHeigth
	to.BottomBoxFontSize = from.BottomBoxFontSize
	to.BottomBoxFontWeigth = from.BottomBoxFontWeigth
	to.BottomBoxFontFamily = from.BottomBoxFontFamily
	to.BottomBoxLetterSpacing = from.BottomBoxLetterSpacing
	to.BottomBoxLetterColorCode = from.BottomBoxLetterColorCode
	to.MovementRectAnchorType = from.MovementRectAnchorType
	to.MovementTextAnchorType = from.MovementTextAnchorType
	to.MovementDominantBaselineType = from.MovementDominantBaselineType
	to.MovementFontSize = from.MovementFontSize
	to.MovementFontWeigth = from.MovementFontWeigth
	to.MovementFontFamily = from.MovementFontFamily
	to.MovementLetterSpacing = from.MovementLetterSpacing
	to.ArtefactTypeFontSize = from.ArtefactTypeFontSize
	to.ArtefactTypeFontWeigth = from.ArtefactTypeFontWeigth
	to.ArtefactTypeFontFamily = from.ArtefactTypeFontFamily
	to.ArtefactTypeLetterSpacing = from.ArtefactTypeLetterSpacing
	to.ArtefactTypeRectAnchorType = from.ArtefactTypeRectAnchorType
	to.ArtefactDominantBaselineType = from.ArtefactDominantBaselineType
	to.ArtefactTypeStrokeWidth = from.ArtefactTypeStrokeWidth
	to.ArtistRectAnchorType = from.ArtistRectAnchorType
	to.ArtistTextAnchorType = from.ArtistTextAnchorType
	to.ArtistDominantBaselineType = from.ArtistDominantBaselineType
	to.ArtistFontSize = from.ArtistFontSize
	to.ArtistFontWeigth = from.ArtistFontWeigth
	to.ArtistFontFamily = from.ArtistFontFamily
	to.ArtistLetterSpacing = from.ArtistLetterSpacing
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

