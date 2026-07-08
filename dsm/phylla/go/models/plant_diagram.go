package models

type PlantDiagram struct {
	Name string

	AbstractTypeFields

	OriginX float64
	OriginY float64

	AxesShape         *AxesShape
	GrowthVectorShape *GrowthVectorShape

	IsChecked bool
}
