package models

type TubeVase3DDiagram struct {
	Name string

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
	IsHiddenTiledFloor3DShape                          bool
	IsHiddenTopCurvePlane1Shape                        bool
	IsHiddenBottomCurvePlane1Shape                     bool
	IsHiddenTopCurvePlane2Shape                        bool
	IsHiddenBottomCurvePlane2Shape                     bool
	IsHiddenVaseTrapezeRingShape                       bool
	IsHiddenStackOfVaseTrapezeRingsShape               bool
	IsHiddenStackOfRotatedVaseTrapezeRingsShape        bool

	Rendered3DShape *Rendered3DShape

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
	TiledFloor3DShape                       *TiledFloor3DShape
	TopCurvePlane1Shape                     *TopCurvePlane1Shape
	BottomCurvePlane1Shape                  *BottomCurvePlane1Shape
	TopCurvePlane2Shape                     *TopCurvePlane2Shape
	BottomCurvePlane2Shape                  *BottomCurvePlane2Shape
	VaseTrapezeRingShape                    *VaseTrapezeRingShape
	StackOfVaseTrapezeRingsShape            *StackOfVaseTrapezeRingsShape
	StackOfRotatedVaseTrapezeRingsShape     *StackOfRotatedVaseTrapezeRingsShape

	IsChecked bool
	AbstractTypeFields
}

type VaseTrapezeRingShape struct {
	Name string
}

type StackOfVaseTrapezeRingsShape struct {
	Name string
}

type StackOfRotatedVaseTrapezeRingsShape struct {
	Name string
}

type TopCurvePlane1Shape struct {
	Name string
}

type BottomCurvePlane1Shape struct {
	Name string
}

type TopCurvePlane2Shape struct {
	Name string
}

type BottomCurvePlane2Shape struct {
	Name string
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
