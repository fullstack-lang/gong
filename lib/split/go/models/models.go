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
}

type View struct {
	Name string

	RootAsSplitAreas []*AsSplitArea
}
