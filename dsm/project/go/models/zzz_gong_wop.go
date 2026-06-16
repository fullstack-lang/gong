// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

var _ = __GONG_time_The_fool_doth_think_he_is_wise__

// insertion point
type Diagram_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsExpanded bool

	LayoutDirection LayoutDirection

	IsChecked bool

	IsEditable_ bool

	IsShowPrefix bool

	DefaultBoxWidth float64

	DefaultBoxHeigth float64

	Width float64

	Height float64

	IsPBSNodeExpanded bool

	IsWBSNodeExpanded bool

	IsTaskGroupsNodeExpanded bool

	DateFormat string

	IsNotesNodeExpanded bool

	IsResourcesNodeExpanded bool

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
}

func (from *Diagram) CopyBasicFields(to *Diagram) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.LayoutDirection = from.LayoutDirection
	to.IsChecked = from.IsChecked
	to.IsEditable_ = from.IsEditable_
	to.IsShowPrefix = from.IsShowPrefix
	to.DefaultBoxWidth = from.DefaultBoxWidth
	to.DefaultBoxHeigth = from.DefaultBoxHeigth
	to.Width = from.Width
	to.Height = from.Height
	to.IsPBSNodeExpanded = from.IsPBSNodeExpanded
	to.IsWBSNodeExpanded = from.IsWBSNodeExpanded
	to.IsTaskGroupsNodeExpanded = from.IsTaskGroupsNodeExpanded
	to.DateFormat = from.DateFormat
	to.IsNotesNodeExpanded = from.IsNotesNodeExpanded
	to.IsResourcesNodeExpanded = from.IsResourcesNodeExpanded
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
}

type Library_WOP struct {
	// insertion point

	Name string

	NbPixPerCharacter float64

	LogoSVGFile string

	ComputedPrefix string

	IsExpanded bool

	LayoutDirection LayoutDirection

	IsRootLibrary bool
}

func (from *Library) CopyBasicFields(to *Library) {
	// insertion point
	to.Name = from.Name
	to.NbPixPerCharacter = from.NbPixPerCharacter
	to.LogoSVGFile = from.LogoSVGFile
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.LayoutDirection = from.LayoutDirection
	to.IsRootLibrary = from.IsRootLibrary
}

type Note_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsExpanded bool

	LayoutDirection LayoutDirection
}

func (from *Note) CopyBasicFields(to *Note) {
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

func (from *NoteProductShape) CopyBasicFields(to *NoteProductShape) {
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

func (from *NoteResourceShape) CopyBasicFields(to *NoteResourceShape) {
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

func (from *NoteShape) CopyBasicFields(to *NoteShape) {
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

func (from *NoteTaskShape) CopyBasicFields(to *NoteTaskShape) {
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

	ComputedPrefix string

	IsExpanded bool

	LayoutDirection LayoutDirection

	IsImport bool

	Description string

	IsProducersNodeExpanded bool

	IsConsumersNodeExpanded bool
}

func (from *Product) CopyBasicFields(to *Product) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.LayoutDirection = from.LayoutDirection
	to.IsImport = from.IsImport
	to.Description = from.Description
	to.IsProducersNodeExpanded = from.IsProducersNodeExpanded
	to.IsConsumersNodeExpanded = from.IsConsumersNodeExpanded
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

func (from *ProductCompositionShape) CopyBasicFields(to *ProductCompositionShape) {
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

	OverideLayoutDirection bool

	LayoutDirection LayoutDirection

	X float64

	Y float64

	Width float64

	Height float64

	IsHidden bool
}

func (from *ProductShape) CopyBasicFields(to *ProductShape) {
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

type Resource_WOP struct {
	// insertion point

	Name string

	ComputedPrefix string

	IsExpanded bool

	LayoutDirection LayoutDirection

	IsImport bool

	Description string
}

func (from *Resource) CopyBasicFields(to *Resource) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.LayoutDirection = from.LayoutDirection
	to.IsImport = from.IsImport
	to.Description = from.Description
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

func (from *ResourceCompositionShape) CopyBasicFields(to *ResourceCompositionShape) {
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

func (from *ResourceShape) CopyBasicFields(to *ResourceShape) {
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

func (from *ResourceTaskShape) CopyBasicFields(to *ResourceTaskShape) {
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

	ComputedPrefix string

	IsExpanded bool

	LayoutDirection LayoutDirection

	IsImport bool

	Start time.Time

	End time.Time

	IsMilestone bool

	Description string

	IsInputsNodeExpanded bool

	IsOutputsNodeExpanded bool

	IsWithCompletion bool

	Completion CompletionEnum

	DisplayVerticalBar bool

	TextPosition TextPositionEnum

	XOffset float64

	YOffset float64
}

func (from *Task) CopyBasicFields(to *Task) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.LayoutDirection = from.LayoutDirection
	to.IsImport = from.IsImport
	to.Start = from.Start
	to.End = from.End
	to.IsMilestone = from.IsMilestone
	to.Description = from.Description
	to.IsInputsNodeExpanded = from.IsInputsNodeExpanded
	to.IsOutputsNodeExpanded = from.IsOutputsNodeExpanded
	to.IsWithCompletion = from.IsWithCompletion
	to.Completion = from.Completion
	to.DisplayVerticalBar = from.DisplayVerticalBar
	to.TextPosition = from.TextPosition
	to.XOffset = from.XOffset
	to.YOffset = from.YOffset
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

func (from *TaskCompositionShape) CopyBasicFields(to *TaskCompositionShape) {
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

	LayoutDirection LayoutDirection
}

func (from *TaskGroup) CopyBasicFields(to *TaskGroup) {
	// insertion point
	to.Name = from.Name
	to.ComputedPrefix = from.ComputedPrefix
	to.IsExpanded = from.IsExpanded
	to.LayoutDirection = from.LayoutDirection
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

func (from *TaskGroupShape) CopyBasicFields(to *TaskGroupShape) {
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

func (from *TaskInputShape) CopyBasicFields(to *TaskInputShape) {
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

func (from *TaskOutputShape) CopyBasicFields(to *TaskOutputShape) {
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

func (from *TaskShape) CopyBasicFields(to *TaskShape) {
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
