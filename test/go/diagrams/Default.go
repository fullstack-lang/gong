package diagrams

import (
	uml "github.com/fullstack-lang/gongdoc/go/models"

	// insertion points for import of the illustrated model
)

var Default uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
	},
	Notes: []*uml.Note{
		{
			Name: `note on the organization between line and points`,
			Body: `
github.com/fullstack-lang/test/go/models is a model that showcases gong possibilites

Astruct is associated with Bstruct via 2 kinds of association
- a N-0..1 association (directly)
- a N-M association via AstructBstructUse
`,
			X:      30.000000,
			Y:      30.000000,
			Width:  240.000000,
			Heigth: 63.000000,
		},
	},
}
