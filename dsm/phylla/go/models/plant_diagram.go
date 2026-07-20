package models

type PlantDiagram struct {
	Name string

	OriginX float64
	OriginY float64

	IsRhombusNodesExpanded bool
	IsArcNodesExpanded     bool

	IsHiddenAxesShape                         bool
	IsHiddenReferenceRhombus                  bool
	IsHiddenPlantCircumferenceShape           bool
	IsHiddenGridPathShape                     bool
	IsHiddenRhombusGridShape                  bool
	IsHiddenExplanationTextShape              bool
	IsHiddenRotatedReferenceRhombus           bool
	IsHiddenRotatedPlantCircumferenceShape    bool
	IsHiddenRotatedGridPathShape              bool
	IsHiddenRotatedRhombusGridShape           bool
	IsHiddenGrowthPathRhombusGridShape        bool
	IsHiddenGrowthVectorShape                 bool
	IsHiddenPerpendicularVectorGrid           bool
	IsHiddenPerpendicularVectorGridHalfway    bool
	IsHiddenBaseVectorShapeGrid               bool
	IsHiddenArcNormalVectorShapeGrid          bool
	IsHiddenStartArcShapeGrid                 bool
	IsHiddenTopStartArcShapeGrid              bool
	IsHiddenShiftedBottomTopStartArcShapeGrid bool
	IsHiddenMidArcVectorShapeGrid             bool
	IsHiddenTopMidArcVectorShapeGrid          bool
	IsHiddenStartHalfwayArcShapeGrid          bool
	IsHiddenTopStartHalfwayArcShapeGrid       bool
	IsHiddenEndHalfwayArcShapeGrid            bool
	IsHiddenTopEndHalfwayArcShapeGrid         bool
	IsHiddenEndArcShapeGrid                   bool
	IsHiddenTopEndArcShapeGrid                bool
	IsHiddenBottomStartArcShapeGrid           bool
	IsHiddenBottomEndArcShapeGrid             bool
	IsHiddenStackOfGrowthCurve                bool
	IsHiddenTopStackOfGrowthCurve             bool
	IsHiddenBottomStackOfGrowthCurve          bool
	IsHiddenShiftedLeftStackOfGrowthCurve     bool
	IsHiddenShiftedLeftStackOfNormalVector    bool

	IsHiddenGrowthCurve2D      bool
	IsHiddenTopGrowthCurve2D   bool
	IsHiddenStackOfGrowthCurve2D    bool
	IsHiddenTopStackOfGrowthCurve2D bool

	IsHiddenStackOfGrowthCurve2DRibbon bool
	IsHiddenStackOfRotatedGrowthCurve2DRibbon bool
	IsHiddenStackOfPartiallyRotatedGrowthCurve2DRibbon bool
	IsHiddenTorusStackShape bool
	IsHiddenVerticalTorusStackShape bool

	IsChecked bool
	AbstractTypeFields

	Rendered3DShape *Rendered3DShape
	TorusStackShape *TorusStackShape
	VerticalTorusStackShape *VerticalTorusStackShape
}
