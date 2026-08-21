package models

type Note struct {
	Name string

	//gong:text width:300 height:300
	Description string

	Complexities []*Complexity
	Performances []*Performance
	Efforts      []*Effort

	LibraryAbstractFields
	AbstractTypeFields

	IsComplexitysNodeExpanded bool

	IsPerformancesNodeExpanded bool

	IsEffortsNodeExpanded bool
}

var _ AbstractType = (*Note)(nil)
