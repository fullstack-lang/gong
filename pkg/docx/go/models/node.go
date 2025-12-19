package models

type Node struct {
	Name string

	Nodes []*Node
}
