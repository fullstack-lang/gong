package models

const LibraryTemplate = `package models

type Library struct {
	Name string

	NbPixPerCharacter float64

	//gong:width 600 gong:height 300
	LogoSVGFile string

	LibraryAbstractFields
	AbstractTypeFields

	IsRootLibrary bool

	SubLibraries []*Library

	objects []AbstractType
}
`
