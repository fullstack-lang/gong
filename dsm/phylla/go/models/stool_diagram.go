package models

type StoolDiagram struct {
	Name string

	IsHiddenSampledPoints3DShape bool
	SampledPoints3DShape         *SampledPoints3DShape

	Rendered3DShape *Rendered3DShape
}
