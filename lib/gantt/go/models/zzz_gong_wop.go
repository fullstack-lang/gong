// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

var _ = __GONG_time_The_fool_doth_think_he_is_wise__

// insertion point
type Arrow_WOP struct {
	// insertion point

	Name string

	OptionnalColor string

	OptionnalStroke string
}

func (from *Arrow) CopyBasicFields(to *Arrow) {
	// insertion point
	to.Name = from.Name
	to.OptionnalColor = from.OptionnalColor
	to.OptionnalStroke = from.OptionnalStroke
}

type Bar_WOP struct {
	// insertion point

	Name string

	Start time.Time

	End time.Time

	ComputedDuration time.Duration

	OptionnalColor string

	OptionnalStroke string

	FillOpacity float64

	StrokeWidth float64

	StrokeDashArray string
}

func (from *Bar) CopyBasicFields(to *Bar) {
	// insertion point
	to.Name = from.Name
	to.Start = from.Start
	to.End = from.End
	to.ComputedDuration = from.ComputedDuration
	to.OptionnalColor = from.OptionnalColor
	to.OptionnalStroke = from.OptionnalStroke
	to.FillOpacity = from.FillOpacity
	to.StrokeWidth = from.StrokeWidth
	to.StrokeDashArray = from.StrokeDashArray
}

type Gantt_WOP struct {
	// insertion point

	Name string

	ComputedStart time.Time

	ComputedEnd time.Time

	ComputedDuration time.Duration

	UseManualStartAndEndDates bool

	ManualStart time.Time

	ManualEnd time.Time

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

	Group_Stroke string

	Group_StrokeWidth float64

	Group_StrokeDashArray string

	DateYOffset float64

	AlignOnStartEndOnYearStart bool
}

func (from *Gantt) CopyBasicFields(to *Gantt) {
	// insertion point
	to.Name = from.Name
	to.ComputedStart = from.ComputedStart
	to.ComputedEnd = from.ComputedEnd
	to.ComputedDuration = from.ComputedDuration
	to.UseManualStartAndEndDates = from.UseManualStartAndEndDates
	to.ManualStart = from.ManualStart
	to.ManualEnd = from.ManualEnd
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
	to.Group_Stroke = from.Group_Stroke
	to.Group_StrokeWidth = from.Group_StrokeWidth
	to.Group_StrokeDashArray = from.Group_StrokeDashArray
	to.DateYOffset = from.DateYOffset
	to.AlignOnStartEndOnYearStart = from.AlignOnStartEndOnYearStart
}

type Group_WOP struct {
	// insertion point

	Name string
}

func (from *Group) CopyBasicFields(to *Group) {
	// insertion point
	to.Name = from.Name
}

type Lane_WOP struct {
	// insertion point

	Name string

	Order int
}

func (from *Lane) CopyBasicFields(to *Lane) {
	// insertion point
	to.Name = from.Name
	to.Order = from.Order
}

type LaneUse_WOP struct {
	// insertion point

	Name string
}

func (from *LaneUse) CopyBasicFields(to *LaneUse) {
	// insertion point
	to.Name = from.Name
}

type Milestone_WOP struct {
	// insertion point

	Name string

	Date time.Time

	DisplayVerticalBar bool
}

func (from *Milestone) CopyBasicFields(to *Milestone) {
	// insertion point
	to.Name = from.Name
	to.Date = from.Date
	to.DisplayVerticalBar = from.DisplayVerticalBar
}

