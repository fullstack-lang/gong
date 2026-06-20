package models

import (
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

type DiagramStructure struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	IsChecked   bool
	IsEditable_ bool

	IsShowPrefix bool // display shapes with their prefix

	Width  float64
	Height float64

	DefaultBoxWidth  float64
	DefaultBoxHeigth float64

	// within the tree branch of one diagram, when an element is present in more than one diagram,
	// it is possible to access it via a list. Only one element have a list that is available per diagram.
	diagramListElement AbstractType

	//
	// DSM specific fields
	//

	owningSystem *System

	Part_Shapes              []*PartShape
	map_Part_PartShape    map[*Part]*PartShape
	IsPartsNodeExpanded      bool
	PartsWhoseNodeIsExpanded []*Part
	map_Part_Rect            map[*Part]*svg.Rect
	map_SvgRect_PartShape    map[*svg.Rect]*PartShape

	Link_Shapes []*LinkAssociationShape
	map_Link_LinkAssociationShape map[*Link]*LinkAssociationShape
	IsLinksNodeExpanded bool
	LinksWhoseNodeIsExpanded []*Link
}

func (d *DiagramStructure) IsEditable() bool {
	return d.IsEditable_
}

func (d *DiagramStructure) SetEditable(v bool) {
	d.IsEditable_ = v
}

func (d *DiagramStructure) GetDefaultBoxHeigth() float64 {
	return d.DefaultBoxHeigth
}

func (d *DiagramStructure) GetDefaultBoxWidth() float64 {
	return d.DefaultBoxWidth
}

func (d *DiagramStructure) GetDiagramListElement() AbstractType {
	return d.diagramListElement
}

func (d *DiagramStructure) SetDiagramListElement(v AbstractType) {
	d.diagramListElement = v
}

func (d *DiagramStructure) GetIsChecked() bool {
	return d.IsChecked
}

func (d *DiagramStructure) SetIsChecked(v bool) {
	d.IsChecked = v
}

func (d *DiagramStructure) GetIsShowPrefix() bool {
	return d.IsShowPrefix
}

func (d *DiagramStructure) SetIsShowPrefix(v bool) {
	d.IsShowPrefix = v
}
