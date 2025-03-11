package models

type Group struct {
	Name string

	// Lanes that need to be grouped
	GroupLanes []*Lane
}
