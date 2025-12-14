// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

var _ = __GONG_time_The_fool_doth_think_he_is_wise__

// insertion point
type ArtefactType_WOP struct {
	// insertion point

	Name string
}

func (from *ArtefactType) CopyBasicFields(to *ArtefactType) {
	// insertion point
	to.Name = from.Name
}

type ArtefactTypeShape_WOP struct {
	// insertion point

	Name string

	X float64

	Y float64

	Width float64

	Height float64
}

func (from *ArtefactTypeShape) CopyBasicFields(to *ArtefactTypeShape) {
	// insertion point
	to.Name = from.Name
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
}

type Artist_WOP struct {
	// insertion point

	Name string
}

func (from *Artist) CopyBasicFields(to *Artist) {
	// insertion point
	to.Name = from.Name
}

type ArtistShape_WOP struct {
	// insertion point

	Name string

	X float64

	Y float64

	Width float64

	Height float64
}

func (from *ArtistShape) CopyBasicFields(to *ArtistShape) {
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

	IsMovementCategoryNodeExpanded bool

	IsArtefactTypeCategoryNodeExpanded bool

	IsArtistCategoryNodeExpanded bool

	IsInfluenceCategoryNodeExpanded bool

	IsMovementCategoryShown bool

	IsArtefactTypeCategoryShown bool

	IsArtistCategoryShown bool

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

	MajorMovementFontSize string

	MinorMovementFontSize string

	MovementFontWeigth string

	MovementFontFamily string

	MovementLetterSpacing string

	AbstractMovementFontSize string

	AbstractMovementRectAnchorType RectAnchorType

	AbstractMovementTextAnchorType TextAnchorType

	AbstractDominantBaselineType DominantBaselineType

	MovementDateRectAnchorType RectAnchorType

	MovementDateTextAnchorType TextAnchorType

	MovementDateTextDominantBaselineType DominantBaselineType

	MovementDateAndPlacesFontSize string

	MovementDateAndPlacesFontWeigth string

	MovementDateAndPlacesFontFamily string

	MovementDateAndPlacesLetterSpacing string

	MovementBelowArcY_Offset float64

	MovementBelowArcY_OffsetPerPlace float64

	MovementPlacesRectAnchorType RectAnchorType

	MovementPlacesTextAnchorType TextAnchorType

	MovementPlacesDominantBaselineType DominantBaselineType

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

	MajorArtistFontSize string

	MinorArtistFontSize string

	ArtistFontWeigth string

	ArtistFontFamily string

	ArtistLetterSpacing string

	ArtistDateRectAnchorType RectAnchorType

	ArtistDateTextAnchorType TextAnchorType

	ArtistDateDominantBaselineType DominantBaselineType

	ArtistDateAndPlacesFontSize string

	ArtistDateAndPlacesFontWeigth string

	ArtistDateAndPlacesFontFamily string

	ArtistDateAndPlacesLetterSpacing string

	ArtistPlacesRectAnchorType RectAnchorType

	ArtistPlacesTextAnchorType TextAnchorType

	ArtistPlacesDominantBaselineType DominantBaselineType

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
	to.IsMovementCategoryNodeExpanded = from.IsMovementCategoryNodeExpanded
	to.IsArtefactTypeCategoryNodeExpanded = from.IsArtefactTypeCategoryNodeExpanded
	to.IsArtistCategoryNodeExpanded = from.IsArtistCategoryNodeExpanded
	to.IsInfluenceCategoryNodeExpanded = from.IsInfluenceCategoryNodeExpanded
	to.IsMovementCategoryShown = from.IsMovementCategoryShown
	to.IsArtefactTypeCategoryShown = from.IsArtefactTypeCategoryShown
	to.IsArtistCategoryShown = from.IsArtistCategoryShown
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
	to.MajorMovementFontSize = from.MajorMovementFontSize
	to.MinorMovementFontSize = from.MinorMovementFontSize
	to.MovementFontWeigth = from.MovementFontWeigth
	to.MovementFontFamily = from.MovementFontFamily
	to.MovementLetterSpacing = from.MovementLetterSpacing
	to.AbstractMovementFontSize = from.AbstractMovementFontSize
	to.AbstractMovementRectAnchorType = from.AbstractMovementRectAnchorType
	to.AbstractMovementTextAnchorType = from.AbstractMovementTextAnchorType
	to.AbstractDominantBaselineType = from.AbstractDominantBaselineType
	to.MovementDateRectAnchorType = from.MovementDateRectAnchorType
	to.MovementDateTextAnchorType = from.MovementDateTextAnchorType
	to.MovementDateTextDominantBaselineType = from.MovementDateTextDominantBaselineType
	to.MovementDateAndPlacesFontSize = from.MovementDateAndPlacesFontSize
	to.MovementDateAndPlacesFontWeigth = from.MovementDateAndPlacesFontWeigth
	to.MovementDateAndPlacesFontFamily = from.MovementDateAndPlacesFontFamily
	to.MovementDateAndPlacesLetterSpacing = from.MovementDateAndPlacesLetterSpacing
	to.MovementBelowArcY_Offset = from.MovementBelowArcY_Offset
	to.MovementBelowArcY_OffsetPerPlace = from.MovementBelowArcY_OffsetPerPlace
	to.MovementPlacesRectAnchorType = from.MovementPlacesRectAnchorType
	to.MovementPlacesTextAnchorType = from.MovementPlacesTextAnchorType
	to.MovementPlacesDominantBaselineType = from.MovementPlacesDominantBaselineType
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
	to.MajorArtistFontSize = from.MajorArtistFontSize
	to.MinorArtistFontSize = from.MinorArtistFontSize
	to.ArtistFontWeigth = from.ArtistFontWeigth
	to.ArtistFontFamily = from.ArtistFontFamily
	to.ArtistLetterSpacing = from.ArtistLetterSpacing
	to.ArtistDateRectAnchorType = from.ArtistDateRectAnchorType
	to.ArtistDateTextAnchorType = from.ArtistDateTextAnchorType
	to.ArtistDateDominantBaselineType = from.ArtistDateDominantBaselineType
	to.ArtistDateAndPlacesFontSize = from.ArtistDateAndPlacesFontSize
	to.ArtistDateAndPlacesFontWeigth = from.ArtistDateAndPlacesFontWeigth
	to.ArtistDateAndPlacesFontFamily = from.ArtistDateAndPlacesFontFamily
	to.ArtistDateAndPlacesLetterSpacing = from.ArtistDateAndPlacesLetterSpacing
	to.ArtistPlacesRectAnchorType = from.ArtistPlacesRectAnchorType
	to.ArtistPlacesTextAnchorType = from.ArtistPlacesTextAnchorType
	to.ArtistPlacesDominantBaselineType = from.ArtistPlacesDominantBaselineType
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

type Movement_WOP struct {
	// insertion point

	Name string
}

func (from *Movement) CopyBasicFields(to *Movement) {
	// insertion point
	to.Name = from.Name
}

type MovementShape_WOP struct {
	// insertion point

	Name string

	X float64

	Y float64

	Width float64

	Height float64
}

func (from *MovementShape) CopyBasicFields(to *MovementShape) {
	// insertion point
	to.Name = from.Name
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
}

