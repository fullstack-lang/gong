package models

type DataStore struct {
	Name string

	// Acronym is a short name for the data store, used optionally as a compact label.
	Acronym string

	//gong:text width:300 height:300
	Description string

	Datas               []*Data
	IsDatasNodeExpanded bool

	LibraryAbstractFields
	AbstractTypeFields

	// SVG is a raw <svg>...</svg> string rendering an optional icon on the data store.
	//gong:text gong:width 600 gong:height 300
	SVG string

	inDataFlows  []*DataFlow
	outDataFlows []*DataFlow
}

var _ AbstractType = (*DataStore)(nil)
