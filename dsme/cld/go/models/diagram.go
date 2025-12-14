package models

type Diagram struct {
	Name            string
	Category1Shapes []*Category1Shape
	Category2Shapes []*Category2Shape
	Category3Shapes []*Category3Shape

	InfluenceShapes []*InfluenceShape

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

	DiagramPresentation
}

type DiagramPresentation struct {
	XMargin float64
	YMargin float64

	Height                  float64
	NextVerticalDateXMargin float64

	RedColorCode            string
	BackgroundGreyColorCode string
	GrayColorCode           string

	BottomBoxYOffset         float64
	BottomBoxWidth           float64
	BottomBoxHeigth          float64
	BottomBoxFontSize        string
	BottomBoxFontWeigth      string
	BottomBoxFontFamily      string
	BottomBoxLetterSpacing   string
	BottomBoxLetterColorCode string

	MovementRectAnchorType       RectAnchorType
	MovementTextAnchorType       TextAnchorType
	MovementDominantBaselineType DominantBaselineType
	MovementFontSize             string
	MovementFontWeigth           string
	MovementFontFamily           string
	MovementLetterSpacing        string

	ArtefactTypeFontSize      string
	ArtefactTypeFontWeigth    string
	ArtefactTypeFontFamily    string
	ArtefactTypeLetterSpacing string

	ArtefactTypeRectAnchorType   RectAnchorType
	ArtefactDominantBaselineType DominantBaselineType
	ArtefactTypeStrokeWidth      float64

	ArtistRectAnchorType       RectAnchorType
	ArtistTextAnchorType       TextAnchorType
	ArtistDominantBaselineType DominantBaselineType
	ArtistFontSize             string
	ArtistFontWeigth           string
	ArtistFontFamily           string
	ArtistLetterSpacing        string

	InfluenceArrowSize        float64
	InfluenceArrowStartOffset float64
	InfluenceArrowEndOffset   float64
	InfluenceCornerRadius     float64

	InfluenceDashedLinePattern string
}
