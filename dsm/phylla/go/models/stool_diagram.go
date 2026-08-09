package models

type StoolDiagram struct {
	Name string

	IsHiddenSeatTopCurveShape bool
	SeatTopCurveShape         *SeatTopCurveShape

	IsHiddenPartiallyRotatedTorusShape bool
	PartiallyRotatedTorusShape         *PartiallyRotatedTorusShape

	IsHiddenSampledPoints3DShape bool
	SampledPoints3DShape         *SampledPoints3DShape

	Rendered3DShape *Rendered3DShape
}
