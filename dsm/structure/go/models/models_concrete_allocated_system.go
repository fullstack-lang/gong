package models

type AllocatedSystemShape struct {
	Name string

	Part *Part
	System     *System
}
