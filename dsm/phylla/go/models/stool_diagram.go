package models

type StoolDiagram struct {
	Name string

	IsHiddenTorusStackShape bool
	TorusStackShape         *TorusStackShape

	IsHiddenSampledPoints3DShape bool
	SampledPoints3DShape         *SampledPoints3DShape

	Rendered3DShape *Rendered3DShape
}
