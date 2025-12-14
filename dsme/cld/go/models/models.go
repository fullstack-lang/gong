package models

type Category interface {
	IsCategory()
	GetName() string
}

// Desk is the singloton organizing the diagrams
type Desk struct {
	Name            string
	SelectedDiagram *Diagram
}
