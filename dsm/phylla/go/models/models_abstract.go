package models

type Plant struct {
	Name string

	N int
	M int
	Z int // number of rhombus

	// InsideAngle is the angle in degree of the diamond at the origin 0,0
	InsideAngle float64

	// how many circle to go around for the front curve
	// the front curve goes from one circle to the nearest
	ShiftToNearestCircle int

	SideLength float64

	// Axes
	Axes *Axes

	LibraryAbstractFields
	AbstractTypeFields

	IsPlantDiagramsNodeExpanded      bool
	PlantDiagramsWhoseNodeIsExpanded []*PlantDiagram
	PlantDiagrams                    []*PlantDiagram
}

// Axes of a plant
//
// each plant has one Axes button
type Axes struct {
	Name string

	LengthX float64
	LengthY float64
}
