package models

type PlantDiagram struct {
	Name string

	OriginX float64
	OriginY float64

	VaseDiagram *VaseDiagram

	IsRhombusNodesExpanded      bool
	IsArcNodesExpanded          bool
	IsVaseArcNodesExpanded      bool
	IsVaseClampingNodesExpanded bool

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
	IsHiddenBottomStartArcShapeGrid        bool
	IsHiddenBottomEndArcShapeGrid          bool
	IsHiddenBottomStackOfGrowthCurve       bool
	IsHiddenShiftedLeftStackOfGrowthCurve  bool
	IsHiddenShiftedLeftStackOfNormalVector bool

	IsHiddenGrowthCurve2D bool

	IsChecked bool
	AbstractTypeFields
}
