package models

import svg "github.com/fullstack-lang/gong/lib/svg/go/models"

type DiagramProcess struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	IsChecked   bool
	IsEditable_ bool

	IsShowPrefix bool // display shapes with their prefix

	DefaultBoxWidth  float64
	DefaultBoxHeigth float64

	Width  float64
	Height float64

	elementWhoseDiagramListIsDisplayed AbstractType

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

	TasksWhoseNodeIsExpanded []*Task
	TaskShapes               []*TaskShape
	map_Task_TaskShape       map[*Task]*TaskShape
	map_Task_Rect            map[*Task]*svg.Rect

	ControlFlowShape []*ControlFlowShape
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

func (d *DiagramProcess) GetElementWhoseDiagramListIsDisplayed() AbstractType {
	return d.elementWhoseDiagramListIsDisplayed
}

func (d *DiagramProcess) SetElementWhoseDiagramListIsDisplayed(v AbstractType) {
	d.elementWhoseDiagramListIsDisplayed = v
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
