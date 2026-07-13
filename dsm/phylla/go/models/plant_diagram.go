package models

type PlantDiagram struct {
	Name string

	OriginX float64
	OriginY float64

	IsRhombusNodesExpanded bool

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
	IsHiddenPerpendicularVectorGridHalfway bool
	IsHiddenBaseVectorShapeGrid            bool
	IsHiddenArcNormalVectorShapeGrid       bool
	IsHiddenStartArcShapeGrid            bool
	IsHiddenTopStartArcShapeGrid         bool
	IsHiddenEndArcShapeGrid              bool
	IsHiddenTopEndArcShapeGrid           bool
	IsHiddenBottomStartArcShapeGrid      bool
	IsHiddenBottomEndArcShapeGrid        bool
	IsHiddenGrowthCurveBezierShapeGrid     bool
	IsHiddenStackOfGrowthCurve           bool
	IsHiddenTopStackOfGrowthCurve        bool
	IsHiddenBottomStackOfGrowthCurve     bool

	IsHiddenGrowthCurve2D    bool
	IsHiddenTopGrowthCurve2D bool

	IsChecked bool
	AbstractTypeFields

	Rendered3DShape *Rendered3DShape
}
