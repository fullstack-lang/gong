package models

type A struct {
	Name string

	B *B

	Bs []*B
}

type B struct {
	Name string
}
