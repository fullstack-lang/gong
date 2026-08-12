package models

type ClockDiagram struct {
	Name string

	IsHiddenClockTopCurveShape bool
	ClockTopCurveShape         *ClockTopCurveShape

	IsHiddenTorus3DShape bool
	Torus3DShape         *Torus3DShape

	IsHiddenSampledPoints3DShape bool
	SampledPoints3DShape         *SampledPoints3DShape

	IsHiddenTiledFloor3DShape bool
	TiledFloor3DShape         *TiledFloor3DShape

	Rendered3DShape *Rendered3DShape
}
