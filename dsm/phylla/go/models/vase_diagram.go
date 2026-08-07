package models

type VaseDiagram struct {
	Name string

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
}

type Angle0Shape struct {
	Name string
}

type SampledPoints3DShape struct {
	Name string
}

type OriginalPoints3DShape struct {
	Name string
}

type Rendered3DShape struct {
	Name string

	ViewX, ViewY, ViewZ       float64
	TargetX, TargetY, TargetZ float64
	Fov                       float64
}
