package models

type ArtElement interface {
	IsArtElement()
	GetName() string
}

type Place struct {
	Name string
}

type Desk struct {
	Name            string
	SelectedDiagram *Diagram
}
