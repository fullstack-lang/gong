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
	IsHiddenCarvedOutTopCurvePlane1Shape               bool
	IsHiddenCarvedOutBottomCurvePlane1Shape            bool
	IsHiddenVaseTrapezeRingShape                       bool
	IsHiddenCarvedOutVaseTrapezeRingShape              bool
	IsHiddenStackOfVaseTrapezeRingsShape               bool
	IsHiddenStackOfCarvedOutVaseTrapezeRingsShape      bool
	IsHiddenStackOfRotatedVaseTrapezeRingsShape        bool
	IsHiddenStackOfRotatedCarvedOutVaseTrapezeRingsShape bool
	IsHiddenVaseTrapezeBasePlateShape                  bool

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
	CarvedOutTopCurvePlane1Shape            *CarvedOutTopCurvePlane1Shape
	CarvedOutBottomCurvePlane1Shape         *CarvedOutBottomCurvePlane1Shape
	VaseTrapezeRingShape                    *VaseTrapezeRingShape
	CarvedOutVaseTrapezeRingShape           *CarvedOutVaseTrapezeRingShape
	StackOfVaseTrapezeRingsShape            *StackOfVaseTrapezeRingsShape
	StackOfCarvedOutVaseTrapezeRingsShape   *StackOfCarvedOutVaseTrapezeRingsShape
	StackOfRotatedVaseTrapezeRingsShape     *StackOfRotatedVaseTrapezeRingsShape
	StackOfRotatedCarvedOutVaseTrapezeRingsShape *StackOfRotatedCarvedOutVaseTrapezeRingsShape
	VaseTrapezeBasePlateShape               *VaseTrapezeBasePlateShape

	IsChecked bool
	AbstractTypeFields
}

type VaseTrapezeBasePlateShape struct {
	Name string
}

type VaseTrapezeRingShape struct {
	Name string
}

type CarvedOutVaseTrapezeRingShape struct {
	Name string
}

type StackOfVaseTrapezeRingsShape struct {
	Name string
}

type StackOfCarvedOutVaseTrapezeRingsShape struct {
	Name string
}

type StackOfRotatedVaseTrapezeRingsShape struct {
	Name string
}

type StackOfRotatedCarvedOutVaseTrapezeRingsShape struct {
	Name string
}

type TopCurvePlane1Shape struct {
	Name string
}

type CarvedOutTopCurvePlane1Shape struct {
	Name string
}

type CarvedOutBottomCurvePlane1Shape struct {
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
