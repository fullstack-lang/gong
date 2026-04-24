package models

type Process struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	parentProcess *Process

	SubProcesses []*Process
}

var (
	_ AbstractType = (*Process)(nil)
)
