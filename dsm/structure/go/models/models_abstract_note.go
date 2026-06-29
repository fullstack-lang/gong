package models

type Note struct {
	//gong:text width:300 height:300
	Name string

	//gong:text width:300 height:300
	Description string

	LibraryAbstractFields
	AbstractTypeFields

	IsPartsNodeExpanded bool
	Parts               []*Part

	IsPortsNodeExpanded bool
	Ports               []*Port
}

var _ AbstractType = (*Note)(nil)
