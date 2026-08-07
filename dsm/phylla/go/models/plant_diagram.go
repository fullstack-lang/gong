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

	IsChecked bool
	AbstractTypeFields

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
