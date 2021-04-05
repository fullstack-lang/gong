package models

// Area describes an editor
// swagger:model Area
type Area struct {
	Name string

	Subarea *Area
}
