package models

// Meta is a fallback solution explained to the meta keyword proposal
type Meta struct {
	Text       string
	References []any // references to symbols in the code
}

var OnTriangleImplementation = Meta{
	Text: `Astruct is the typical gongstruct and it is references Bstruct in different ways`,
	References: []any{
		Astruct{},
		Bstruct{},
		Astruct{}.AnAstruct,
		Astruct{}.Anarrayofa,
	},
}
