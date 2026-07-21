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

	IsHiddenGrowthCurve2DRibbon             bool
	IsHiddenShiftedRightGrowthCurve2DRibbon bool
	IsHiddenStackOfGrowthCurve2DRibbon      bool
	IsHiddenStackOfRotatedGrowthCurve2DRibbon bool
	IsHiddenStackOfPartiallyRotatedGrowthCurve2DRibbon bool
	IsHiddenPartiallyGrowthCurve2DRibbon bool
	IsHiddenTorusStackShape bool
	IsHiddenVerticalTorusStackShape bool
	IsHiddenPartiallyRotatedTorusShape bool
	IsHiddenStackOfPartiallyRotatedTorusShape bool

	IsChecked bool
	AbstractTypeFields

	Rendered3DShape *Rendered3DShape
	GrowthCurve2DRibbon *GrowthCurve2DRibbon
	ShiftedRightGrowthCurve2DRibbon *ShiftedRightGrowthCurve2DRibbon
	TorusStackShape *TorusStackShape
	VerticalTorusStackShape *VerticalTorusStackShape
	PartiallyRotatedTorusShape *PartiallyRotatedTorusShape
	StackOfPartiallyRotatedTorusShape *StackOfPartiallyRotatedTorusShape
}
