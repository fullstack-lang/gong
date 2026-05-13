package models

type Data struct {
	Name string

	//gong:text width:300 height:300
	Description string

	LibraryAbstractFields
	AbstractTypeFields

	// SVG_Path is a a SVG Path rendering an icon representing the resource.
	//gong:width 600 gong:height 300
	SVG_Path string

	// $960$ is a "magic number" because it is perfectly divisible by $24, 40, 48,$ and $60$
	//
	InverseAppliedScaling float64
}

var _ AbstractType = (*Data)(nil)
