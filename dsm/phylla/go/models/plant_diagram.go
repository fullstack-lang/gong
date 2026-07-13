package models

type PlantDiagram struct {
	Name string

	OriginX float64
	OriginY float64

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
	IsHiddenStartArcShapeV2Grid            bool
	IsHiddenTopStartArcShapeV2Grid         bool
	IsHiddenEndArcShapeV2Grid              bool
	IsHiddenTopEndArcShapeV2Grid           bool
	IsHiddenBottomStartArcShapeV2Grid      bool
	IsHiddenBottomEndArcShapeV2Grid        bool
	IsHiddenGrowthCurveBezierShapeGrid     bool
	IsHiddenStackOfGrowthCurveV2           bool
	IsHiddenTopStackOfGrowthCurveV2        bool
	IsHiddenBottomStackOfGrowthCurveV2     bool

	IsHiddenGrowthCurve2D bool
	IsHiddenTopGrowthCurve2D bool

	IsChecked bool
	AbstractTypeFields

	Rendered3DShape *Rendered3DShape
	
}
