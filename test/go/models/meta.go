package models

// Triangle is a naive implementation of

type Meta struct {
	Text       string
	References []any // references to symbols in the code
}

var onTriangleImplementation = Meta{
	Text: `Astruct is the typical gongstruct and it is references Bstruct in different ways`,
	References: []any{
		Astruct{},
		Bstruct{},
		Astruct{}.AnAstruct,
		Astruct{}.Anarrayofa,
	},
}
