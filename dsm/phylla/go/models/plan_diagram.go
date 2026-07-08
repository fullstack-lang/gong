package models

type PlantDiagram struct {
	Name string

	AbstractTypeFields

	AxesShape *AxesShape

	IsChecked bool
}
