package models

import svg "github.com/fullstack-lang/gong/lib/svg/go/models"

// dataShapeKey is used to track the presence of a data shape for a given data flow and data.
type dataShapeKey struct {
	dataFlow *DataFlow
	data     *Data
}

type allocatedResourceShapeKey struct {
	part     *Part
	resource *Resource
}

type allocatedSystemShapeKey struct {
	part   *Part
	system *System
}

type DiagramStructure struct {
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

	// Part
	Part_Shapes             []*PartShape
	map_Part_PartShape      map[*Part]*PartShape
	map_Part_Rect           map[*Part]*svg.Rect
	IsPartsNodeExpanded     bool
	PartWhoseNodeIsExpanded []*Part

	// ExternalPart
	ExternalPart_Shapes             []*ExternalPartShape
	map_Part_ExternalPartShape      map[*Part]*ExternalPartShape
	map_ExternalPart_Rect           map[*Part]*svg.Rect
	IsExternalPartsNodeExpanded     bool
	ExternalPartWhoseNodeIsExpanded []*Part
	map_SvgRect_ExternalPartShape   map[*svg.Rect]*ExternalPartShape
	map_SvgRect_Part                map[*svg.Rect]*Part

	ExternalPartsWhoseOutDataFlowsNodeIsExpanded []*Part
	ExternalPartsWhoseInDataFlowsNodeIsExpanded  []*Part

	// Port
	PortsWhoseNodeIsExpanded []*Port
	Port_Shapes              []*PortShape
	map_Port_PortShape       map[*Port]*PortShape
	map_Port_Rect            map[*Port]*svg.Rect
	map_SvgRect_PortShape    map[*svg.Rect]*PortShape // for drawing links between port

	// ControlFlow
	ControlFlowsWhoseNodeIsExpanded  []*ControlFlow
	ControlFlow_Shapes               []*ControlFlowShape
	map_ControlFlow_ControlFlowShape map[*ControlFlow]*ControlFlowShape

	// DataFlow
	DataFlowsWhoseNodeIsExpanded []*DataFlow
	DataFlow_Shapes              []*DataFlowShape
	map_DataFlow_DataFlowShape   map[*DataFlow]*DataFlowShape

	// Data
	DatasWhoseNodeIsExpanded         []*Data
	Data_Shapes                      []*DataShape
	map_DataShapeKey_DataShape       map[dataShapeKey]*DataShape
	DataFlowsWhoseDataNodeIsExpanded []*DataFlow // within a diagram, at the data flow level, some some can be expanded or not

	// allocated resources
	AllocatedResourcesWhoseNodeIsExpanded                []*Resource
	AllocatedResourceShapes                              []*AllocatedResourceShape
	map_AllocatedResourceShapeKey_AllocatedResourceShape map[allocatedResourceShapeKey]*AllocatedResourceShape

	// alllocated systemes
	AllocatedSystemesWhoseNodeIsExpanded             []*System
	AllocatedSystemShapes                            []*AllocatedSystemShape
	map_AllocatedSystemShapeKey_AllocatedSystemShape map[allocatedSystemShapeKey]*AllocatedSystemShape

	Note_Shapes              []*NoteShape
	map_Note_NoteShape       map[*Note]*NoteShape
	NotesWhoseNodeIsExpanded []*Note
	IsNotesNodeExpanded      bool

	NotePortShapes         []*NotePortShape
	map_Note_NotePortShape map[notePortKey]*NotePortShape
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
