package models

import svg "github.com/fullstack-lang/gong/lib/svg/go/models"

// dataShapeKey is used to track the presence of a data shape for a given data flow and data.
type dataShapeKey struct {
	dataFlow *DataFlow
	data     *Data
}

type allocatedResourceShapeKey struct {
	participant *Participant
	resource    *Resource
}

type DiagramProcess struct {
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

	// Process
	Process_Shapes              []*ProcessShape
	map_Process_ProcessShape    map[*Process]*ProcessShape
	map_Process_Rect            map[*Process]*svg.Rect
	IsProcesssNodeExpanded      bool
	ProcesssWhoseNodeIsExpanded []*Process

	owningProcess *Process

	// Participant
	Participant_Shapes               []*ParticipantShape
	map_Participant_ParticipantShape map[*Participant]*ParticipantShape
	map_Participant_Rect             map[*Participant]*svg.Rect
	IsParticipantsNodeExpanded       bool
	ParticipantWhoseNodeIsExpanded   []*Participant

	// ExternalParticipant
	ExternalParticipant_Shapes               []*ExternalParticipantShape
	map_Participant_ExternalParticipantShape map[*Participant]*ExternalParticipantShape
	map_ExternalParticipant_Rect             map[*Participant]*svg.Rect
	IsExternalParticipantsNodeExpanded       bool
	ExternalParticipantWhoseNodeIsExpanded   []*Participant
	map_SvgRect_ExternalParticipantShape     map[*svg.Rect]*ExternalParticipantShape
	map_SvgRect_Participant                  map[*svg.Rect]*Participant

	ExternalParticipantsWhoseOutDataFlowsNodeIsExpanded []*Participant
	ExternalParticipantsWhoseInDataFlowsNodeIsExpanded  []*Participant

	// Task
	TasksWhoseNodeIsExpanded []*Task
	Task_Shapes              []*TaskShape
	map_Task_TaskShape       map[*Task]*TaskShape
	map_Task_Rect            map[*Task]*svg.Rect
	map_SvgRect_TaskShape    map[*svg.Rect]*TaskShape // for drawing links between task

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

	Note_Shapes              []*NoteShape
	map_Note_NoteShape       map[*Note]*NoteShape
	NotesWhoseNodeIsExpanded []*Note
	IsNotesNodeExpanded      bool

	NoteTaskShapes         []*NoteTaskShape
	map_Note_NoteTaskShape map[noteTaskKey]*NoteTaskShape
}

func (d *DiagramProcess) IsEditable() bool {
	return d.IsEditable_
}

func (d *DiagramProcess) SetEditable(v bool) {
	d.IsEditable_ = v
}

func (d *DiagramProcess) GetDefaultBoxHeigth() float64 {
	return d.DefaultBoxHeigth
}

func (d *DiagramProcess) GetDefaultBoxWidth() float64 {
	return d.DefaultBoxWidth
}

func (d *DiagramProcess) GetDiagramListElement() AbstractType {
	return d.diagramListElement
}

func (d *DiagramProcess) SetDiagramListElement(v AbstractType) {
	d.diagramListElement = v
}

func (d *DiagramProcess) GetIsChecked() bool {
	return d.IsChecked
}

func (d *DiagramProcess) SetIsChecked(v bool) {
	d.IsChecked = v
}

func (d *DiagramProcess) GetIsShowPrefix() bool {
	return d.IsShowPrefix
}

func (d *DiagramProcess) SetIsShowPrefix(v bool) {
	d.IsShowPrefix = v
}
