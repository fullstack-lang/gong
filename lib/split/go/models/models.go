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

	Size float64

	// IsAany make the split area set size to "$any(*)"
	IsAny bool

	AsSplits []*AsSplit

	Tree  *Tree
	Table *Table
	Form  *Form
	Svg   *Svg
	Doc   *Doc
}

type View struct {
	Name string

	RootAsSplitAreas []*AsSplitArea
}

type Tree struct {
	Name string // name of the stack

	StackName string
	TreeName  string
}

type Table struct {
	Name string // name of the stack

	StackName string
	TableName string
}

type Form struct {
	Name string // name of the stack

	StackName string
	FormName  string
}

type Svg struct {
	Name string // name of the stack

	StackName string
}

type Doc struct {
	Name string // name of the stack

	StackName string
}
