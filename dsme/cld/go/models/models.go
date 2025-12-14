package models

type ArtElement interface {
	IsArtElement()
	GetName() string
}

// Desk is the singloton organizing the diagrams
type Desk struct {
	Name            string
	SelectedDiagram *Diagram
}
