package models

type Data struct {
	Name string

	// Acronym is a short name for the resource, used for instance in a data flow as a repalacement
	// of the full name.
	// For instance FPL For Flight Plan
	Acronym string

	//gong:text width:300 height:300
	Description string

	LibraryAbstractFields
	AbstractTypeFields

	// SVG is a raw <svg>...</svg> string rendering an icon representing the data item.
	//gong:text gong:width 600 gong:height 300
	SVG string
}

var _ AbstractType = (*Data)(nil)
