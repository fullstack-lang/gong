package models

type Plant2DDiagram struct {
	Name string

	OriginX float64
	OriginY float64
	Zoom    float64

	IsRhombusNodesExpanded bool
	IsArcNodesExpanded     bool

	IsHiddenAxesShape                      bool
	IsHiddenReferenceRhombus               bool
	IsHiddenPlantCircumferenceShape        bool
	IsHiddenGridPathShape                  bool
	IsHiddenRhombusGridShape               bool
	IsHiddenExplanationTextShape           bool
	IsHiddenRotatedReferenceRhombus        bool
	IsHiddenRotatedPlantCircumferenceShape bool
	IsHiddenRotatedGridPathShape           bool
	IsHiddenRotatedRhombusGridShape        bool
	IsHiddenGrowthPathRhombusGridShape     bool
	IsHiddenGrowthVectorShape              bool
	IsHiddenPerpendicularVectorGrid        bool
	IsHiddenBaseVectorShapeGrid            bool
	IsHiddenArcNormalVectorShapeGrid       bool
	IsHiddenStartArcShapeGrid              bool
	IsHiddenMidArcVectorShapeGrid          bool
	IsHiddenEndArcShapeGrid                bool

	IsHiddenGrowthCurve2D bool

	IsHiddenStackOfGrowthCurve2DByGrowthVector bool

	IsChecked bool
	AbstractTypeFields
}
