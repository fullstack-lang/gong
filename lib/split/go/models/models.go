package models

// the following blank import will force the vendoring of all the front end code
// this allows the front end compilation of the split front end
import (
	_ "github.com/fullstack-lang/gong/lib/button/ng-github.com-fullstack-lang-gong-lib-button"
	_ "github.com/fullstack-lang/gong/lib/cursor/ng-github.com-fullstack-lang-gong-lib-cursor"
	_ "github.com/fullstack-lang/gong/lib/load/ng-github.com-fullstack-lang-gong-lib-load"
	_ "github.com/fullstack-lang/gong/lib/markdown/ng-github.com-fullstack-lang-gong-lib-markdown"
	_ "github.com/fullstack-lang/gong/lib/sim/ng-github.com-fullstack-lang-gong-lib-sim"
	_ "github.com/fullstack-lang/gong/lib/slider/ng-github.com-fullstack-lang-gong-lib-slider"
	_ "github.com/fullstack-lang/gong/lib/svg/ng-github.com-fullstack-lang-gong-lib-svg"
	_ "github.com/fullstack-lang/gong/lib/table/ng-github.com-fullstack-lang-gong-lib-table"
	_ "github.com/fullstack-lang/gong/lib/tone/ng-github.com-fullstack-lang-gong-lib-tone"
	_ "github.com/fullstack-lang/gong/lib/tree/ng-github.com-fullstack-lang-gong-lib-tree"
)

// View is the top structuring eleement of the split component
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
	IsSecondatyView bool
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
}

type AsSplitArea struct {
	Name string

	ShowNameInHeader bool

	Size float64

	// IsAany make the split area set size to "$any(*)"
	IsAny bool

	AsSplit *AsSplit

	Button   *Button
	Cursor   *Cursor
	Form     *Form
	Load     *Load
	Markdown *Markdown
	Slider   *Slider
	Split    *Split
	Svg      *Svg
	Table    *Table
	Tone     *Tone
	Tree     *Tree
	Xlsx     *Xlsx

	HasDiv   bool
	DivStyle string // in case of div, the div style
}

type Button struct {
	Name      string // name of the stack
	StackName string
}

type Cursor struct {
	Name      string // name of the stack
	StackName string
	Style     string
}

type Form struct {
	Name      string // name of the stack
	StackName string
}

type Load struct {
	Name      string // name of the stack
	StackName string
}

type Markdown struct {
	Name      string // name of the stack
	StackName string
}

type Slider struct {
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

type Tone struct {
	Name      string // name of the stack
	StackName string
}

type Tree struct {
	Name      string // name of the stack
	StackName string
}

type Xlsx struct {
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

// LogoOnTheLeft displayed in the banner
type LogoOnTheRight struct {
	Name string

	Width  int
	Height int

	//gong:width 600 gong:height 300
	SVG string
}
