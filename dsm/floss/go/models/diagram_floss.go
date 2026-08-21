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

	// Complexity
	Complexity_Shapes              []*ComplexityShape
	map_Complexity_ComplexityShape map[*Complexity]*ComplexityShape
	map_Complexity_Rect            map[*Complexity]*svg.Rect
	IsComplexitysNodeExpanded      bool
	ComplexitysWhoseNodeIsExpanded []*Complexity

	// Performance
	Performance_Shapes               []*PerformanceShape
	map_Performance_PerformanceShape map[*Performance]*PerformanceShape
	map_Performance_Rect             map[*Performance]*svg.Rect
	IsPerformancesNodeExpanded       bool
	PerformancesWhoseNodeIsExpanded  []*Performance

	// Effort
	Effort_Shapes              []*EffortShape
	map_Effort_EffortShape     map[*Effort]*EffortShape
	map_Effort_Rect            map[*Effort]*svg.Rect
	IsEffortsNodeExpanded      bool
	EffortsWhoseNodeIsExpanded []*Effort

	// Note
	Note_Shapes                   []*NoteShape
	NoteComplexityShapes          []*NoteComplexityShape
	NotePerformanceShapes         []*NotePerformanceShape
	NoteEffortShapes              []*NoteEffortShape
	map_Note_NoteShape            map[*Note]*NoteShape
	map_Note_Rect                 map[*Note]*svg.Rect
	map_Note_NoteComplexityShape  map[noteComplexityKey]*NoteComplexityShape
	map_Note_NotePerformanceShape map[notePerformanceKey]*NotePerformanceShape
	map_Note_NoteEffortShape      map[noteEffortKey]*NoteEffortShape
	IsNotesNodeExpanded           bool
	NotesWhoseNodeIsExpanded      []*Note

	map_SvgRect_NoteShape        map[*svg.Rect]*NoteShape
	map_SvgRect_ComplexityShape  map[*svg.Rect]*ComplexityShape
	map_SvgRect_PerformanceShape map[*svg.Rect]*PerformanceShape
	map_SvgRect_EffortShape      map[*svg.Rect]*EffortShape

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
