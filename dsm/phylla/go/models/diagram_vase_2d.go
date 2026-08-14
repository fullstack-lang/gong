package models

type Vase2DDiagram struct {
	Name string

	Zoom float64

	IsVaseArcNodesExpanded      bool
	IsVaseClampingNodesExpanded bool

	IsHiddenAxesShape                         bool
	IsHiddenBottomStartArcShapeGrid           bool
	IsHiddenBottomEndArcShapeGrid             bool
	IsHiddenBottomStackOfGrowthCurve          bool
	IsHiddenShiftedLeftStackOfGrowthCurve     bool
	IsHiddenShiftedLeftStackOfNormalVector    bool
	IsHiddenPerpendicularVectorGridHalfway    bool
	IsHiddenTopStartArcShapeGrid              bool
	IsHiddenShiftedBottomTopStartArcShapeGrid bool
	IsHiddenTopMidArcVectorShapeGrid          bool
	IsHiddenStartHalfwayArcShapeGrid          bool
	IsHiddenTopStartHalfwayArcShapeGrid       bool
	IsHiddenEndHalfwayArcShapeGrid            bool
	IsHiddenTopEndHalfwayArcShapeGrid         bool
	IsHiddenTopEndArcShapeGrid                bool
	IsHiddenStackOfGrowthCurve                bool
	IsHiddenTopStackOfGrowthCurve             bool

	IsHiddenTopGrowthCurve2D        bool
	IsHiddenStackOfGrowthCurve2D    bool
	IsHiddenTopStackOfGrowthCurve2D bool

	IsHiddenGrowthCurve2DRibbon                     bool
	IsHiddenShiftedRightGrowthCurve2DRibbon         bool
	IsHiddenShiftedLeftGrowthCurve2DRibbon          bool
	IsHiddenStackOfGrowthCurve2DRibbon              bool
	IsHiddenStackOfRotatedGrowthCurve2DRibbon       bool
	IsHiddenPartiallyGrowthCurve2DRibbon            bool
	IsHiddenShiftedLeftPartiallyGrowthCurve2DRibbon bool
	IsHiddenPartiallyGrowthCurve2DTrajectory        bool
	IsHiddenPartiallyGrowthCurve2DTrajectoryP1P2    bool
	IsHiddenPxShape                                 bool
	IsHiddenChosenP1P2PairShape                     bool
	IsHiddenKeyHoleShape                            bool

	IsChecked bool
	AbstractTypeFields
}
