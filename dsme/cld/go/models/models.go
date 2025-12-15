package models

type Category interface {
	GongstructIF
	IsCategory()
}

// Desk is the singloton organizing the diagrams
type Desk struct {
	Name            string
	SelectedDiagram *Diagram
}
