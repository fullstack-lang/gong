package models

type Plant struct {
	Name string

	// N is the number of steps to take along the first main path (up and to the right)
	// on the unrolled lattice of the plant's surface to reach the next leaf in the genetic spiral.
	N int
	// M is the number of steps to take along the second main path (up and to the left)
	// on the unrolled lattice of the plant's surface to reach the next leaf in the genetic spiral.
	M int
	Z int // number of rhombus

	// RhombusInsideAngle is set by the user. It represents the inside angle (in degrees) of the
	// fundamental diamond (rhombus) shape that makes up the grid of leaves on the plant's surface.
	// This angle determines the geometric direction of the two main lattice paths (up-right and up-left).
	RhombusInsideAngle float64

	// RhombusSideLength is set by the user. It represents the physical length of the side of the
	// fundamental diamond (rhombus) shape. It acts as the scale or distance for each step
	// taken along the lattice paths.
	RhombusSideLength float64

	// how many circle to go around for the front curve
	// the front curve goes from one circle to the nearest
	ShiftToNearestCircle int

	LibraryAbstractFields
	AbstractTypeFields

	IsPlantDiagramsNodeExpanded      bool
	PlantDiagramsWhoseNodeIsExpanded []*PlantDiagram
	PlantDiagrams                    []*PlantDiagram
}
