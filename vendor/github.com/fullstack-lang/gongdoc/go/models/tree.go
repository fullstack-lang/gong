package models

// Tree is a tree of nodes for selecting items to display
type Tree struct {
	Name string

	Type TreeType

	RootNodes []*Node
}
