package models

type StoolDiagram struct {
	Name string

	IsHiddenPartiallyRotatedTorusShape bool
	PartiallyRotatedTorusShape         *PartiallyRotatedTorusShape

	IsHiddenSampledPoints3DShape bool
	SampledPoints3DShape         *SampledPoints3DShape

	Rendered3DShape *Rendered3DShape
}
