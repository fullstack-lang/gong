package models

type System struct {
	Name string
}

type Part struct {
	Name string
}

type Link struct {
	Name string

	Source *Part
	Target *Part
}
