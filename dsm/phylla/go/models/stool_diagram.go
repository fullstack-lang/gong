package models

type StoolDiagram struct {
	Name string

	IsHiddenSeatTopCurveShape bool
	SeatTopCurveShape         *SeatTopCurveShape

	IsHiddenPartiallyRotatedSeatTopCurveShape bool
	PartiallyRotatedSeatTopCurveShape         *PartiallyRotatedSeatTopCurveShape

	IsHiddenTorus3DShape bool
	Torus3DShape         *Torus3DShape

	IsHiddenPartiallyRotatedTorusShape bool
	PartiallyRotatedTorusShape         *PartiallyRotatedTorusShape

	IsHiddenSampledPoints3DShape bool
	SampledPoints3DShape         *SampledPoints3DShape

	Rendered3DShape *Rendered3DShape
}
