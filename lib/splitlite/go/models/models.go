package models

// View is the top structuring element of the split component
//
// There can be many instancied views
// The one with [models.IsSelectedView] set to true will be displayed
// (if more than one view has this field to true, the behavior is not specified)
type View struct {
	Name             string
	ShowViewName     bool // by default, a standalone views wont display the
	RootAsSplitAreas []*AsSplitArea
	IsSelectedView   bool

	// Direction the initial direction for the split dividing of the view. Vertical if ""
	Direction Direction

	// will be displayed in a vertical manner, with the first element at the top and the last element at the bottom
	IsSecondaryView bool

	IsSizeInPixel          bool
	IsWithCustomGutterSize bool
	GutterSize             float64

	OnClick func()
}

// OnAfterUpdate is called when there is an update to the view
func (view *View) OnAfterUpdate(stage *Stage, _, frontView *View) {
	if frontView.IsSelectedView && view.OnClick != nil {
		view.OnClick()
	}
}

type Direction string

const (
	Vertical   Direction = "vertical"
	Horizontal Direction = "horizontal"
)

type AsSplit struct {
	Name string

	Direction Direction

	AsSplitAreas []*AsSplitArea

	IsSizeInPixel bool

	IsWithCustomGutterSize bool
	GutterSize             float64
}

type AsSplitArea struct {
	Name string

	ShowNameInHeader bool

	Size float64

	// IsAny makes the split area set size to "$any(*)"
	IsAny bool

	AsSplit *AsSplit

	Button *Button
	Form   *Form
	Load   *Load
	Split  *Split
	Svg    *Svg
	Table  *Table
	Tree   *Tree

	HasDiv   bool
	DivStyle string // in case of div, the div style
}

type Button struct {
	Name      string // name of the stack
	StackName string
}

type Form struct {
	Name      string // name of the stack
	StackName string
}

type Load struct {
	Name      string // name of the stack
	StackName string
}

type Split struct {
	Name      string // name of the stack
	StackName string
}

type Svg struct {
	Name      string // name of the stack
	StackName string
	Style     string
}

type Table struct {
	Name      string // name of the stack
	StackName string
}

type Tree struct {
	Name      string // name of the stack
	StackName string
}

// You can configure the browser tab title and favicon dynamically in Angular
type Title struct {
	Name string
}

// You can configure the browser tab title and favicon dynamically in Angular
// creates one instance and it will set the browser accordingly
type FavIcon struct {
	Name string

	//gong:width 600 gong:height 300
	SVG string
}

// LogoOnTheLeft displayed in the banner
type LogoOnTheLeft struct {
	Name string

	Width  int
	Height int

	//gong:width 600 gong:height 300
	SVG string
}

// LogoOnTheRight displayed in the banner
type LogoOnTheRight struct {
	Name string

	Width  int
	Height int

	//gong:width 600 gong:height 300
	SVG string
}
