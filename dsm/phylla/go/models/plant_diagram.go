package models

type PlantDiagram struct {
	Name string

	OriginX float64
	OriginY float64

	IsHiddenAxesShape                      bool
	IsHiddenReferenceRhombus               bool
	IsHiddenPlantCircumferenceShape        bool
	IsHiddenGridPathShape                  bool
	IsHiddenRhombusGridShape               bool
	IsHiddenExplanationTextShape           bool
	IsHiddenRotatedReferenceRhombus        bool
	IsHiddenRotatedPlantCircumferenceShape bool
	IsHiddenRotatedGridPathShape           bool
	IsHiddenRotatedRhombusGridShape        bool
	IsHiddenGrowthPathRhombusGridShape     bool

	IsChecked bool
	AbstractTypeFields
}
