// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type Diagram_WOP struct {
	// insertion point

	Name string

	DefaultBoxWidth float64

	DefaultBoxHeigth float64

	DateFormat string

	Width float64

	Height float64

	IsTimeDiagram bool

	ComputedStart time.Time

	ComputedEnd time.Time

	ComputedDuration time.Duration

	UseManualStartAndEndDates bool

	ManualStart time.Time

	ManualEnd time.Time

	TimeStep int

	TimeStepScale TimeStepScaleEnum

	LaneHeight float64

	RatioBarToLaneHeight float64

	YTopMargin float64

	XLeftText float64

	TextHeight float64

	XLeftLanes float64

	XRightMargin float64

	ArrowLengthToTheRightOfStartBar float64

	ArrowTipLenght float64

	TimeLine_Color string

	TimeLine_FillOpacity float64

	TimeLine_Stroke string

	TimeLine_StrokeWidth float64

	DrawVerticalTimeLines bool

	Group_Stroke string

	Group_StrokeWidth float64

	Group_StrokeDashArray string

	DateYOffset float64

	AlignOnStartEndOnYearStart bool

	ComputedPrefix string

	IsExpanded bool

	IsChecked bool

	IsEditable_ bool

	IsShowPrefix bool

	IsInAutoLayoutMode bool

	IsPBSNodeExpanded bool

	IsWBSNodeExpanded bool

	IsTaskGroupsNodeExpanded bool

	IsNotesNodeExpanded bool

	IsResourcesNodeExpanded bool
}

func (from *Diagram) GongCopyBasicFields(to *Diagram) {
	// insertion point
	to.Name = from.Name
	to.DefaultBoxWidth = from.DefaultBoxWidth
	to.DefaultBoxHeigth = from.DefaultBoxHeigth
	to.DateFormat = from.DateFormat
	to.Width = from.Width
	to.Height = from.Height
	to.IsTimeDiagram = from.IsTimeDiagram
	to.ComputedStart = from.ComputedStart
	to.ComputedEnd = from.ComputedEnd
	to.ComputedDuration = from.ComputedDuration
	to.UseManualStartAndEndDates = from.UseManualStartAndEndDates
	to.ManualStart = from.ManualStart
	to.ManualEnd = from.ManualEnd
	to.TimeStep = from.TimeStep
	to.TimeStepScale = from.TimeStepScale
	to.LaneHeight = from.LaneHeight
	to.RatioBarToLaneHeight = from.RatioBarToLaneHeight
	to.YTopMargin = from.YTopMargin
	to.XLeftText = from.XLeftText
	to.TextHeight = from.TextHeight
	to.XLeftLanes = from.XLeftLanes
	to.XRightMargin = from.XRightMargin
	to.ArrowLengthToTheRightOfStartBar = from.ArrowLengthToTheRightOfStartBar
	to.ArrowTipLenght = from.ArrowTipLenght
	to.TimeLine_Color = from.TimeLine_Color
	to.TimeLine_FillOpacity = from.TimeLine_FillOpacity
	to.TimeLine_Stroke = from.TimeLine_Stroke
	to.TimeLine_StrokeWidth = from.TimeLine_StrokeWidth
	to.DrawVerticalTimeLines = from.DrawVerticalTimeLines
	to.Group_Stroke = from.Group_Stroke
	to.Group_StrokeWidth = from.Group_StrokeWidth
	to.Group_StrokeDashArray = from.Group_StrokeDashArray
	to.DateYOffset = from.DateYOffset
	to.AlignOnStartEndOnYearStart = from.AlignOnStartEndOnYearStart
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.IsChecked = from.IsChecked
	to.IsEditable_ = from.IsEditable_
	to.IsShowPrefix = from.IsShowPrefix
	to.IsInAutoLayoutMode = from.IsInAutoLayoutMode
	to.IsPBSNodeExpanded = from.IsPBSNodeExpanded
	to.IsWBSNodeExpanded = from.IsWBSNodeExpanded
	to.IsTaskGroupsNodeExpanded = from.IsTaskGroupsNodeExpanded
	to.IsNotesNodeExpanded = from.IsNotesNodeExpanded
	to.IsResourcesNodeExpanded = from.IsResourcesNodeExpanded
}

type Library_WOP struct {
	// insertion point

	Name string

	NbPixPerCharacter float64

	LogoSVGFile string

	ComputedPrefix string

	IsExpanded bool

	IsRootLibrary bool
}

func (from *Library) GongCopyBasicFields(to *Library) {
	// insertion point
	to.Name = from.Name
	to.NbPixPerCharacter = from.NbPixPerCharacter
	to.LogoSVGFile = from.LogoSVGFile
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.IsRootLibrary = from.IsRootLibrary
}

type Note_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsExpanded bool

	LayoutDirection LayoutDirection
}

func (from *Note) GongCopyBasicFields(to *Note) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.LayoutDirection = from.LayoutDirection
}

type NoteProductShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *NoteProductShape) GongCopyBasicFields(to *NoteProductShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type NoteResourceShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *NoteResourceShape) GongCopyBasicFields(to *NoteResourceShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type NoteShape_WOP struct {
	// insertion point

	Name string

	OverideLayoutDirection bool

	LayoutDirection LayoutDirection

	X float64

	Y float64

	Width float64

	Height float64

	IsHidden bool
}

func (from *NoteShape) GongCopyBasicFields(to *NoteShape) {
	// insertion point
	to.Name = from.Name
	to.OverideLayoutDirection = from.OverideLayoutDirection
	to.LayoutDirection = from.LayoutDirection
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
	to.IsHidden = from.IsHidden
}

type NoteTaskShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *NoteTaskShape) GongCopyBasicFields(to *NoteTaskShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type Product_WOP struct {
	// insertion point

	Name string

	Description string

	IsProducersNodeExpanded bool

	IsConsumersNodeExpanded bool

	IsImport bool

	ComputedPrefix string

	IsExpanded bool

	LayoutDirection LayoutDirection
}

func (from *Product) GongCopyBasicFields(to *Product) {
	// insertion point
	to.Name = from.Name
	to.Description = from.Description
	to.IsProducersNodeExpanded = from.IsProducersNodeExpanded
	to.IsConsumersNodeExpanded = from.IsConsumersNodeExpanded
	to.IsImport = from.IsImport
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.LayoutDirection = from.LayoutDirection
}

type ProductCompositionShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *ProductCompositionShape) GongCopyBasicFields(to *ProductCompositionShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type ProductReferenceShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *ProductReferenceShape) GongCopyBasicFields(to *ProductReferenceShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type ProductShape_WOP struct {
	// insertion point

	Name string

	IsShowType bool

	OverideLayoutDirection bool

	LayoutDirection LayoutDirection

	X float64

	Y float64

	Width float64

	Height float64

	IsHidden bool
}

func (from *ProductShape) GongCopyBasicFields(to *ProductShape) {
	// insertion point
	to.Name = from.Name
	to.IsShowType = from.IsShowType
	to.OverideLayoutDirection = from.OverideLayoutDirection
	to.LayoutDirection = from.LayoutDirection
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
	to.IsHidden = from.IsHidden
}

type Resource_WOP struct {
	// insertion point

	Name string

	Description string

	ComputedPrefix string

	IsExpanded bool

	LayoutDirection LayoutDirection

	IsImport bool
}

func (from *Resource) GongCopyBasicFields(to *Resource) {
	// insertion point
	to.Name = from.Name
	to.Description = from.Description
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.LayoutDirection = from.LayoutDirection
	to.IsImport = from.IsImport
}

type ResourceCompositionShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *ResourceCompositionShape) GongCopyBasicFields(to *ResourceCompositionShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type ResourceShape_WOP struct {
	// insertion point

	Name string

	OverideLayoutDirection bool

	LayoutDirection LayoutDirection

	X float64

	Y float64

	Width float64

	Height float64

	IsHidden bool
}

func (from *ResourceShape) GongCopyBasicFields(to *ResourceShape) {
	// insertion point
	to.Name = from.Name
	to.OverideLayoutDirection = from.OverideLayoutDirection
	to.LayoutDirection = from.LayoutDirection
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
	to.IsHidden = from.IsHidden
}

type ResourceTaskShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *ResourceTaskShape) GongCopyBasicFields(to *ResourceTaskShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type Task_WOP struct {
	// insertion point

	Name string

	Description string

	Start time.Time

	End time.Time

	IsMilestone bool

	DependencyType DependencyTypeEnum

	DependencyDurationYears float64

	DependencyDurationMonths float64

	DependencyDurationWeeks float64

	DependencyDurationDays float64

	DependencyDurationHours float64

	DurationYears float64

	DurationMonths float64

	DurationWeeks float64

	DurationDays float64

	DurationHours float64

	IsEndDateComputedFromDuration bool

	IsWithCompletion bool

	Completion CompletionEnum

	DisplayVerticalBar bool

	TextPosition TextPositionEnum

	XOffset float64

	YOffset float64

	IsImport bool

	IsInputsNodeExpanded bool

	IsOutputsNodeExpanded bool

	ComputedPrefix string

	IsExpanded bool

	LayoutDirection LayoutDirection
}

func (from *Task) GongCopyBasicFields(to *Task) {
	// insertion point
	to.Name = from.Name
	to.Description = from.Description
	to.Start = from.Start
	to.End = from.End
	to.IsMilestone = from.IsMilestone
	to.DependencyType = from.DependencyType
	to.DependencyDurationYears = from.DependencyDurationYears
	to.DependencyDurationMonths = from.DependencyDurationMonths
	to.DependencyDurationWeeks = from.DependencyDurationWeeks
	to.DependencyDurationDays = from.DependencyDurationDays
	to.DependencyDurationHours = from.DependencyDurationHours
	to.DurationYears = from.DurationYears
	to.DurationMonths = from.DurationMonths
	to.DurationWeeks = from.DurationWeeks
	to.DurationDays = from.DurationDays
	to.DurationHours = from.DurationHours
	to.IsEndDateComputedFromDuration = from.IsEndDateComputedFromDuration
	to.IsWithCompletion = from.IsWithCompletion
	to.Completion = from.Completion
	to.DisplayVerticalBar = from.DisplayVerticalBar
	to.TextPosition = from.TextPosition
	to.XOffset = from.XOffset
	to.YOffset = from.YOffset
	to.IsImport = from.IsImport
	to.IsInputsNodeExpanded = from.IsInputsNodeExpanded
	to.IsOutputsNodeExpanded = from.IsOutputsNodeExpanded
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.LayoutDirection = from.LayoutDirection
}

type TaskCompositionShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *TaskCompositionShape) GongCopyBasicFields(to *TaskCompositionShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type TaskGroup_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsExpanded bool
}

func (from *TaskGroup) GongCopyBasicFields(to *TaskGroup) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
}

type TaskGroupShape_WOP struct {
	// insertion point

	Name string

	X float64

	Y float64

	Width float64

	Height float64

	IsHidden bool
}

func (from *TaskGroupShape) GongCopyBasicFields(to *TaskGroupShape) {
	// insertion point
	to.Name = from.Name
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
	to.IsHidden = from.IsHidden
}

type TaskInputShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *TaskInputShape) GongCopyBasicFields(to *TaskInputShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type TaskOutputShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *TaskOutputShape) GongCopyBasicFields(to *TaskOutputShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type TaskPredecessorShape_WOP struct {
	// insertion point

	Name string

	StartRatio float64

	EndRatio float64

	StartOrientation OrientationType

	EndOrientation OrientationType

	CornerOffsetRatio float64

	IsHidden bool
}

func (from *TaskPredecessorShape) GongCopyBasicFields(to *TaskPredecessorShape) {
	// insertion point
	to.Name = from.Name
	to.StartRatio = from.StartRatio
	to.EndRatio = from.EndRatio
	to.StartOrientation = from.StartOrientation
	to.EndOrientation = from.EndOrientation
	to.CornerOffsetRatio = from.CornerOffsetRatio
	to.IsHidden = from.IsHidden
}

type TaskShape_WOP struct {
	// insertion point

	Name string

	IsShowDate bool

	OverideLayoutDirection bool

	LayoutDirection LayoutDirection

	X float64

	Y float64

	Width float64

	Height float64

	IsHidden bool
}

func (from *TaskShape) GongCopyBasicFields(to *TaskShape) {
	// insertion point
	to.Name = from.Name
	to.IsShowDate = from.IsShowDate
	to.OverideLayoutDirection = from.OverideLayoutDirection
	to.LayoutDirection = from.LayoutDirection
	to.X = from.X
	to.Y = from.Y
	to.Width = from.Width
	to.Height = from.Height
	to.IsHidden = from.IsHidden
}

// end of insertion point
