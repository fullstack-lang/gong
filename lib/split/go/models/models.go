package models

// the following blank import will force the vendoring of all the front end code
// this allows the front end compilation of the split front end
import (
	_ "github.com/fullstack-lang/gong/lib/button/ng-github.com-fullstack-lang-gong-lib-button"
	_ "github.com/fullstack-lang/gong/lib/cursor/ng-github.com-fullstack-lang-gong-lib-cursor"
	_ "github.com/fullstack-lang/gong/lib/doc/ng-github.com-fullstack-lang-gong-lib-doc"
	_ "github.com/fullstack-lang/gong/lib/gantt/ng-github.com-fullstack-lang-gong-lib-gantt"
	_ "github.com/fullstack-lang/gong/lib/load/ng-github.com-fullstack-lang-gong-lib-load"
	_ "github.com/fullstack-lang/gong/lib/sim/ng-github.com-fullstack-lang-gong-lib-sim"
	_ "github.com/fullstack-lang/gong/lib/slider/ng-github.com-fullstack-lang-gong-lib-slider"
	_ "github.com/fullstack-lang/gong/lib/svg/ng-github.com-fullstack-lang-gong-lib-svg"
	_ "github.com/fullstack-lang/gong/lib/table/ng-github.com-fullstack-lang-gong-lib-table"
	_ "github.com/fullstack-lang/gong/lib/tone/ng-github.com-fullstack-lang-gong-lib-tone"
	_ "github.com/fullstack-lang/gong/lib/tree/ng-github.com-fullstack-lang-gong-lib-tree"
	_ "github.com/fullstack-lang/gong/lib/xlsx/ng-github.com-fullstack-lang-gong-lib-xlsx"
)

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

	Button *Button
	Cursor *Cursor
	Doc    *Doc
	Form   *Form
	Load   *Load
	Slider *Slider
	Split  *Split
	Svg    *Svg
	Table  *Table
	Tone   *Tone
	Tree   *Tree
	Xlsx   *Xlsx

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

type Doc struct {
	Name      string // name of the stack
	StackName string
}

type Form struct {
	Name      string // name of the stack
	StackName string
	FormName  string
}

type Load struct {
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
	TableName string
}

type Tone struct {
	Name      string // name of the stack
	StackName string
}

type Tree struct {
	Name      string // name of the stack
	StackName string
	TreeName  string
}

type Xlsx struct {
	Name      string // name of the stack
	StackName string
}

type View struct {
	Name             string
	ShowViewName     bool // by default, a standalone views wont display the
	RootAsSplitAreas []*AsSplitArea
}
