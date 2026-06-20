package models

type System struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	Parts []*Part
	IsPartsNodeExpanded bool
	PartsWhoseNodeIsExpanded []*Part

	SubSystems []*System
	IsSubSystemsNodeExpanded bool
	SubSystemsWhoseNodeIsExpanded []*System
	
	Links []*Link
	IsLinksNodeExpanded bool
	LinksWhoseNodeIsExpanded []*Link

	DiagramStructures []*DiagramStructure
	IsDiagramStructuresNodeExpanded bool
	DiagramStructuresWhoseNodeIsExpanded []*DiagramStructure
}

type Part struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields
}

type Link struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	Source *Part
	Target *Part
}
