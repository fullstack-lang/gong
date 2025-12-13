package models

// TreeNames are the names of the tree used in the front
// using an enum allows the compile-time consistency of the tree names
// between the front and the back
type TreeNames string

const (
	Portfolio TreeNames = "portfolio"
	Model     TreeNames = "model"
)
