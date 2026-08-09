package models

type VaseDiagram struct {
	Name string

	IsVaseArcNodesExpanded      bool
	IsVaseClampingNodesExpanded bool

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

	IsHiddenStackOfPartiallyRotatedGrowthCurve2DRibbon bool
	IsHiddenTorusStackShape                            bool
	IsHiddenVerticalTorusStackShape                    bool
	IsHiddenPartiallyRotatedTorusShape                 bool
	IsHiddenStackOfPartiallyRotatedTorusShape          bool
	IsHiddenPointsAndLines3DShape                      bool
	IsHiddenKeyHole3DShape                             bool
	IsHiddenKey3DShape                                 bool
	IsHiddenVolumeKey3DShape                           bool
	IsHiddenTorusEdge3DShape                           bool
	IsHiddenSampledPoints3DShape                       bool
	IsHiddenOriginalPoints3DShape                      bool
	IsHiddenAngle0Shape                                bool

	Rendered3DShape                         *Rendered3DShape
	GrowthCurve2DRibbon                     *GrowthCurve2DRibbon
	ShiftedRightGrowthCurve2DRibbon         *ShiftedRightGrowthCurve2DRibbon
	ShiftedLeftGrowthCurve2DRibbon          *ShiftedLeftGrowthCurve2DRibbon
	ShiftedLeftPartiallyGrowthCurve2DRibbon *ShiftedLeftPartiallyGrowthCurve2DRibbon
	TorusStackShape                         *TorusStackShape
	VerticalTorusStackShape                 *VerticalTorusStackShape
	PartiallyRotatedTorusShape              *PartiallyRotatedTorusShape
	StackOfPartiallyRotatedTorusShape       *StackOfPartiallyRotatedTorusShape
	PointsAndLines3DShape                   *PointsAndLines3DShape
	SampledPoints3DShape                    *SampledPoints3DShape
	OriginalPoints3DShape                   *OriginalPoints3DShape
	Angle0Shape                             *Angle0Shape
	KeyHole3DShape                          *KeyHole3DShape
	Key3DShape                              *Key3DShape
	VolumeKey3DShape                        *VolumeKey3DShape
	TorusEdge3DShape                        *TorusEdge3DShape
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
