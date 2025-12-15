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

	Height float64
	Width  float64

	NbPixPerCharacter float64

	RedColorCode            string
	BackgroundGreyColorCode string
	GrayColorCode           string

	Category1RectAnchorType       RectAnchorType
	Category1TextAnchorType       TextAnchorType
	Category1DominantBaselineType DominantBaselineType
	Category1FontSize             string
	Category1FontWeigth           string
	Category1FontFamily           string
	Category1LetterSpacing        string

	Category2TypeFontSize         string
	Category2TypeFontWeigth       string
	Category2TypeFontFamily       string
	Category2TypeLetterSpacing    string
	Category2TypeRectAnchorType   RectAnchorType
	Category2DominantBaselineType DominantBaselineType
	Category2StrokeWidth          float64

	Category3RectAnchorType       RectAnchorType
	Category3TextAnchorType       TextAnchorType
	Category3DominantBaselineType DominantBaselineType
	Category3FontSize             string
	Category3FontWeigth           string
	Category3FontFamily           string
	Category3LetterSpacing        string

	InfluenceStrokeWidth      float64
	InfluenceArrowSize        float64
	InfluenceArrowStartOffset float64
	InfluenceArrowEndOffset   float64
	InfluenceCornerRadius     float64

	InfluenceFontSize      string
	InfluenceFontWeigth    string
	InfluenceFontFamily    string
	InfluenceLetterSpacing string

	InfluenceDashedLinePattern string
}
