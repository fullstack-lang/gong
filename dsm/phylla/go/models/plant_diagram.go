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
	IsHiddenStartArcShapeGrid              bool
	IsHiddenStartArcShapeV2Grid            bool
	IsHiddenTopStartArcShapeV2Grid         bool
	IsHiddenEndArcShapeGrid                bool
	IsHiddenEndArcShapeV2Grid              bool
	IsHiddenTopEndArcShapeV2Grid           bool
	IsHiddenGrowthCurveBezierShapeGrid     bool
	IsHiddenStackOfGrowthCurve             bool

	IsChecked bool
	AbstractTypeFields

	Rendered3DShape *Rendered3DShape
	
}
