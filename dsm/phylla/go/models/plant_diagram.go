package models

type PlantDiagram struct {
	Name string

	OriginX float64
	OriginY float64

	AxesShape                *AxesShape
	ReferenceRhombus         *ReferenceRhombus
	GrowthVectorShape        *GrowthVectorShape
	GridPathShape            *GridPathShape
	RotatedReferenceRhombus  *ReferenceRhombus
	RotatedGrowthVectorShape *GrowthVectorShape
	RotatedGridPathShape     *GridPathShape


	IsChecked bool
	AbstractTypeFields
}
