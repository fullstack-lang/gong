package models

import svg "github.com/fullstack-lang/gong/lib/svg/go/models"

type DiagramFloss struct {
	Name string

	//gong:text width:300 height:300
	Description string

	LibraryAbstractFields
	AbstractTypeFields

	IsChecked   bool
	IsEditable_ bool

	IsShowPrefix bool // display shapes with their prefix

	DefaultBoxWidth  float64
	DefaultBoxHeigth float64

	Width  float64
	Height float64

	// within the tree branch of one diagram, when an element is present in more than one diagram,
	// it is possible to access it via a list. Only one element have a list that is available per diagram.
	diagramListElement AbstractType

	//
	//  DSM specific fields
	//

	// System
	System_Shapes              []*SystemShape
	map_System_SystemShape     map[*System]*SystemShape
	map_System_Rect            map[*System]*svg.Rect
	IsSystemsNodeExpanded      bool
	SystemsWhoseNodeIsExpanded []*System

	owningSystem *System
}

func (d *DiagramFloss) IsEditable() bool {
	return d.IsEditable_
}

func (d *DiagramFloss) SetEditable(v bool) {
	d.IsEditable_ = v
}

func (d *DiagramFloss) GetDefaultBoxHeigth() float64 {
	return d.DefaultBoxHeigth
}

func (d *DiagramFloss) GetDefaultBoxWidth() float64 {
	return d.DefaultBoxWidth
}

func (d *DiagramFloss) GetDiagramListElement() AbstractType {
	return d.diagramListElement
}

func (d *DiagramFloss) SetDiagramListElement(v AbstractType) {
	d.diagramListElement = v
}

func (d *DiagramFloss) GetIsChecked() bool {
	return d.IsChecked
}

func (d *DiagramFloss) SetIsChecked(v bool) {
	d.IsChecked = v
}

func (d *DiagramFloss) GetIsShowPrefix() bool {
	return d.IsShowPrefix
}

func (d *DiagramFloss) SetIsShowPrefix(v bool) {
	d.IsShowPrefix = v
}
