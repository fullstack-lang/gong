package models

type Note struct {
	Name string

	//gong:text width:300 height:300
	Description string

	LibraryAbstractFields
	AbstractTypeFields

	IsComplexitysNodeExpanded bool
	Complexities              []*Complexity

	IsPerformancesNodeExpanded bool
	Performances               []*Performance

	IsEffortsNodeExpanded bool
	Efforts               []*Effort
}

var _ AbstractType = (*Note)(nil)
