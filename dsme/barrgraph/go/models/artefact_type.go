package models

type ArtefactType struct {
	Name string
}

func (*ArtefactType) IsArtElement() {
}

func (shape *ArtefactTypeShape) GetArtElement() *ArtefactType {
	return shape.ArtefactType
}

type ArtefactTypeShape struct {
	Name string

	ArtefactType *ArtefactType

	X, Y float64

	Width, Height float64

	IsHidden bool
}
