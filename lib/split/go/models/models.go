package models

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

	AsSplits []*AsSplit

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

type View struct {
	Name             string
	RootAsSplitAreas []*AsSplitArea
}
