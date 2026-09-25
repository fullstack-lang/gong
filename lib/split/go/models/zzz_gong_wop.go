// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type AsSplit_WOP struct {
	// insertion point

	Name string

	Direction Direction

	IsSizeInPixel bool

	IsWithCustomGutterSize bool

	GutterSize float64
}

func (from *AsSplit) GongCopyBasicFields(to *AsSplit) {
	// insertion point
	to.Name = from.Name
	to.Direction = from.Direction
	to.IsSizeInPixel = from.IsSizeInPixel
	to.IsWithCustomGutterSize = from.IsWithCustomGutterSize
	to.GutterSize = from.GutterSize
}

type AsSplitArea_WOP struct {
	// insertion point

	Name string

	ShowNameInHeader bool

	Size float64

	IsAny bool

	HasDiv bool

	DivStyle string
}

func (from *AsSplitArea) GongCopyBasicFields(to *AsSplitArea) {
	// insertion point
	to.Name = from.Name
	to.ShowNameInHeader = from.ShowNameInHeader
	to.Size = from.Size
	to.IsAny = from.IsAny
	to.HasDiv = from.HasDiv
	to.DivStyle = from.DivStyle
}

type Button_WOP struct {
	// insertion point

	Name string

	StackName string
}

func (from *Button) GongCopyBasicFields(to *Button) {
	// insertion point
	*to = *from
}

type Cursor_WOP struct {
	// insertion point

	Name string

	StackName string

	Style string
}

func (from *Cursor) GongCopyBasicFields(to *Cursor) {
	// insertion point
	*to = *from
}

type FavIcon_WOP struct {
	// insertion point

	Name string

	SVG string
}

func (from *FavIcon) GongCopyBasicFields(to *FavIcon) {
	// insertion point
	*to = *from
}

type Form_WOP struct {
	// insertion point

	Name string

	StackName string
}

func (from *Form) GongCopyBasicFields(to *Form) {
	// insertion point
	*to = *from
}

type Load_WOP struct {
	// insertion point

	Name string

	StackName string
}

func (from *Load) GongCopyBasicFields(to *Load) {
	// insertion point
	*to = *from
}

type LogoOnTheLeft_WOP struct {
	// insertion point

	Name string

	Width int

	Height int

	SVG string
}

func (from *LogoOnTheLeft) GongCopyBasicFields(to *LogoOnTheLeft) {
	// insertion point
	*to = *from
}

type LogoOnTheRight_WOP struct {
	// insertion point

	Name string

	Width int

	Height int

	SVG string
}

func (from *LogoOnTheRight) GongCopyBasicFields(to *LogoOnTheRight) {
	// insertion point
	*to = *from
}

type Markdown_WOP struct {
	// insertion point

	Name string

	StackName string
}

func (from *Markdown) GongCopyBasicFields(to *Markdown) {
	// insertion point
	*to = *from
}

type Slider_WOP struct {
	// insertion point

	Name string

	StackName string
}

func (from *Slider) GongCopyBasicFields(to *Slider) {
	// insertion point
	*to = *from
}

type Split_WOP struct {
	// insertion point

	Name string

	StackName string
}

func (from *Split) GongCopyBasicFields(to *Split) {
	// insertion point
	*to = *from
}

type Svg_WOP struct {
	// insertion point

	Name string

	StackName string

	Style string
}

func (from *Svg) GongCopyBasicFields(to *Svg) {
	// insertion point
	*to = *from
}

type Table_WOP struct {
	// insertion point

	Name string

	StackName string
}

func (from *Table) GongCopyBasicFields(to *Table) {
	// insertion point
	*to = *from
}

type Threejs_WOP struct {
	// insertion point

	Name string

	StackName string
}

func (from *Threejs) GongCopyBasicFields(to *Threejs) {
	// insertion point
	*to = *from
}

type Title_WOP struct {
	// insertion point

	Name string
}

func (from *Title) GongCopyBasicFields(to *Title) {
	// insertion point
	*to = *from
}

type Tone_WOP struct {
	// insertion point

	Name string

	StackName string
}

func (from *Tone) GongCopyBasicFields(to *Tone) {
	// insertion point
	*to = *from
}

type Tree_WOP struct {
	// insertion point

	Name string

	StackName string
}

func (from *Tree) GongCopyBasicFields(to *Tree) {
	// insertion point
	*to = *from
}

type View_WOP struct {
	// insertion point

	Name string

	ShowViewName bool

	IsSelectedView bool

	Direction Direction

	IsSecondaryView bool

	IsSizeInPixel bool

	IsWithCustomGutterSize bool

	GutterSize float64
}

func (from *View) GongCopyBasicFields(to *View) {
	// insertion point
	to.Name = from.Name
	to.ShowViewName = from.ShowViewName
	to.IsSelectedView = from.IsSelectedView
	to.Direction = from.Direction
	to.IsSecondaryView = from.IsSecondaryView
	to.IsSizeInPixel = from.IsSizeInPixel
	to.IsWithCustomGutterSize = from.IsWithCustomGutterSize
	to.GutterSize = from.GutterSize
}

type Xlsx_WOP struct {
	// insertion point

	Name string

	StackName string
}

func (from *Xlsx) GongCopyBasicFields(to *Xlsx) {
	// insertion point
	*to = *from
}

// end of insertion point
