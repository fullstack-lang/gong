package models

type PlantDiagram struct {
	Name string

	OriginX float64
	OriginY float64

	AxesShape                      *AxesShape
	ReferenceRhombus               *ReferenceRhombus
	PlantCircumferenceShape        *PlantCircumferenceShape
	GridPathShape                  *GridPathShape
	RhombusGridShape               *RhombusGridShape
	ExplanationTextShape           *ExplanationTextShape
	RotatedReferenceRhombus        *ReferenceRhombus
	RotatedPlantCircumferenceShape *PlantCircumferenceShape
	RotatedGridPathShape           *GridPathShape
	RotatedRhombusGridShape        *RhombusGridShape


	IsChecked bool
	AbstractTypeFields
}
