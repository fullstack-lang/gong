package models

import (
	"time"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

type DiagramHierarchy struct {
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

	Product_Shapes              []*ProductShape
	map_Product_ProductShape    map[*Product]*ProductShape
	ProductsWhoseNodeIsExpanded []*Product // to be made private once in production (no need to persist)
	IsPBSNodeExpanded           bool

	ProductComposition_Shapes           []*ProductCompositionShape
	map_Product_ProductCompositionShape map[*Product]*ProductCompositionShape

	IsWBSNodeExpanded bool

	Task_Shapes                    []*TaskShape
	map_Task_TaskShape             map[*Task]*TaskShape
	TasksWhoseNodeIsExpanded       []*Task // to be made private once in production (no need to persist)ExpandableNodeObject
	TasksWhoseInputNodeIsExpanded  []*Task
	TasksWhoseOutputNodeIsExpanded []*Task

	IsTaskGroupsNodeExpanded      bool
	TaskGroupShapes               []*TaskGroupShape
	map_TaskGroup_TaskGroupShape  map[*TaskGroup]*TaskGroupShape
	TaskGroupsWhoseNodeIsExpanded []*TaskGroup

	IsMilestonesNodeExpanded      bool
	MilestoneShapes               []*MilestoneShape
	map_Milestone_MilestoneShape  map[*Milestone]*MilestoneShape
	MilestonesWhoseNodeIsExpanded []*Milestone

	// date format for task shapes. If "", use "2006-01-02"
	// date format is go date format convention
	DateFormat string

	TaskComposition_Shapes        []*TaskCompositionShape
	map_Task_TaskCompositionShape map[*Task]*TaskCompositionShape

	TaskInputShapes         []*TaskInputShape
	map_Task_TaskInputShape map[taskProductKey]*TaskInputShape

	TaskOutputShapes         []*TaskOutputShape
	map_Task_TaskOutputShape map[taskProductKey]*TaskOutputShape

	Note_Shapes              []*NoteShape
	map_Note_NoteShape       map[*Note]*NoteShape
	NotesWhoseNodeIsExpanded []*Note
	IsNotesNodeExpanded      bool

	NoteProductShapes         []*NoteProductShape
	map_Note_NoteProductShape map[noteProductKey]*NoteProductShape

	NoteTaskShapes         []*NoteTaskShape
	map_Note_NoteTaskShape map[noteTaskKey]*NoteTaskShape

	NoteResourceShapes         []*NoteResourceShape
	map_Note_NoteResourceShape map[noteResourceKey]*NoteResourceShape

	Resource_Shapes              []*ResourceShape
	map_Resource_ResourceShape   map[*Resource]*ResourceShape
	ResourcesWhoseNodeIsExpanded []*Resource
	IsResourcesNodeExpanded      bool

	ResourceComposition_Shapes            []*ResourceCompositionShape
	map_Resource_ResourceCompositionShape map[*Resource]*ResourceCompositionShape

	ResourceTaskShapes             []*ResourceTaskShape
	map_Resource_ResourceTaskShape map[resourceTaskKey]*ResourceTaskShape

	map_Product_Rect  map[*Product]*svg.Rect
	map_Task_Rect     map[*Task]*svg.Rect
	map_Note_Rect     map[*Note]*svg.Rect
	map_Resource_Rect map[*Resource]*svg.Rect

	map_SvgRect_ProductShape  map[*svg.Rect]*ProductShape
	map_SvgRect_TaskShape     map[*svg.Rect]*TaskShape
	map_SvgRect_NoteShape     map[*svg.Rect]*NoteShape
	map_SvgRect_ResourceShape map[*svg.Rect]*ResourceShape

	// within the tree branch of one diagram, when an element is present in more than one diagram,
	// it is possible to access it via a list. Only one element have a list that is available per diagram.
	diagramListElement AbstractType

	// The following field are used when the diagram is a time diagram
	IsTimeDiagram bool

	// dates computed from tasks of the gantt
	ComputedStart    time.Time
	ComputedEnd      time.Time
	ComputedDuration time.Duration

	// start and end dates if manual setup is true
	UseManualStartAndEndDates bool
	ManualStart               time.Time
	ManualEnd                 time.Time

	// TimeStep is the step between ticks
	TimeStep int

	// TimeStepScale is the scale of the time step
	TimeStepScale TimeStepScaleEnum

	LaneHeight           float64
	RatioBarToLaneHeight float64
	YTopMargin           float64

	XLeftText  float64
	TextHeight float64

	XLeftLanes   float64
	XRightMargin float64

	ArrowLengthToTheRightOfStartBar float64
	ArrowTipLenght                  float64

	TimeLine_Color       string
	TimeLine_FillOpacity float64
	TimeLine_Stroke      string
	TimeLine_StrokeWidth float64

	DrawVerticalTimeLines bool

	Group_Stroke          string
	Group_StrokeWidth     float64
	Group_StrokeDashArray string

	DateYOffset float64

	AlignOnStartEndOnYearStart bool
}

func (d *DiagramHierarchy) IsEditable() bool {
	return d.IsEditable_
}

func (d *DiagramHierarchy) SetEditable(v bool) {
	d.IsEditable_ = v
}

func (d *DiagramHierarchy) GetDefaultBoxHeigth() float64 {
	return d.DefaultBoxHeigth
}

func (d *DiagramHierarchy) GetDefaultBoxWidth() float64 {
	return d.DefaultBoxWidth
}

func (d *DiagramHierarchy) GetDiagramListElement() AbstractType {
	return d.diagramListElement
}

func (d *DiagramHierarchy) SetDiagramListElement(v AbstractType) {
	d.diagramListElement = v
}

func (d *DiagramHierarchy) GetIsChecked() bool {
	return d.IsChecked
}

func (d *DiagramHierarchy) SetIsChecked(v bool) {
	d.IsChecked = v
}

func (d *DiagramHierarchy) GetIsShowPrefix() bool {
	return d.IsShowPrefix
}

func (d *DiagramHierarchy) SetIsShowPrefix(v bool) {
	d.IsShowPrefix = v
}
