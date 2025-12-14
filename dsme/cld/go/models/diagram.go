package models

type Diagram struct {
	Name               string
	MovementShapes     []*MovementShape
	ArtefactTypeShapes []*ArtefactTypeShape
	ArtistShapes       []*ArtistShape
	InfluenceShapes    []*InfluenceShape

	IsEditable bool

	IsNodeExpanded                     bool
	IsMovementCategoryNodeExpanded     bool
	IsArtefactTypeCategoryNodeExpanded bool
	IsArtistCategoryNodeExpanded       bool
	IsInfluenceCategoryNodeExpanded    bool

	IsMovementCategoryShown     bool
	IsArtefactTypeCategoryShown bool
	IsArtistCategoryShown       bool
	IsInfluenceCategoryShown    bool

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
	MajorMovementFontSize        string
	MinorMovementFontSize        string
	MovementFontWeigth           string
	MovementFontFamily           string
	MovementLetterSpacing        string

	AbstractMovementFontSize       string
	AbstractMovementRectAnchorType RectAnchorType
	AbstractMovementTextAnchorType TextAnchorType
	AbstractDominantBaselineType   DominantBaselineType

	MovementDateRectAnchorType           RectAnchorType
	MovementDateTextAnchorType           TextAnchorType
	MovementDateTextDominantBaselineType DominantBaselineType
	MovementDateAndPlacesFontSize        string
	MovementDateAndPlacesFontWeigth      string
	MovementDateAndPlacesFontFamily      string
	MovementDateAndPlacesLetterSpacing   string

	MovementBelowArcY_Offset         float64
	MovementBelowArcY_OffsetPerPlace float64

	MovementPlacesRectAnchorType       RectAnchorType
	MovementPlacesTextAnchorType       TextAnchorType
	MovementPlacesDominantBaselineType DominantBaselineType

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
	MajorArtistFontSize        string
	MinorArtistFontSize        string
	ArtistFontWeigth           string
	ArtistFontFamily           string
	ArtistLetterSpacing        string

	ArtistDateRectAnchorType         RectAnchorType
	ArtistDateTextAnchorType         TextAnchorType
	ArtistDateDominantBaselineType   DominantBaselineType
	ArtistDateAndPlacesFontSize      string
	ArtistDateAndPlacesFontWeigth    string
	ArtistDateAndPlacesFontFamily    string
	ArtistDateAndPlacesLetterSpacing string

	ArtistPlacesRectAnchorType       RectAnchorType
	ArtistPlacesTextAnchorType       TextAnchorType
	ArtistPlacesDominantBaselineType DominantBaselineType

	InfluenceArrowSize        float64
	InfluenceArrowStartOffset float64
	InfluenceArrowEndOffset   float64
	InfluenceCornerRadius     float64

	InfluenceDashedLinePattern string
}
