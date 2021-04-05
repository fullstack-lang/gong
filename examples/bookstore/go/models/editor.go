package models

// Editor describes an editor
// swagger:model Editor
type Editor struct {
	Name string

	Books []*Book
}
