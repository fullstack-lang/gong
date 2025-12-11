package models

type A struct {
	Name string

	Bs []*B
}

type B struct {
	Name string
}
