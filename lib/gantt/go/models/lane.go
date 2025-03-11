package models

type Lane struct {
	Name string

	// Lane are displayed in increasing order
	Order int
	Bars  []*Bar
}

type LaneUse struct {
	Name string

	Lane *Lane
}
