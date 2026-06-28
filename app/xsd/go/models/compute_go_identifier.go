package models

import "fmt"

func computeGoIdentifier(name string, x *WithGoIdentifier, setOfGoIdentifiers map[string]any) {
	var hasNameCollision bool
	initialGoIdentifier := xsdNameToGoIdentifier(name)

	goIdentifier := initialGoIdentifier
	index := 0
	for index == 0 || hasNameCollision {
		index++
		_, hasNameCollision = setOfGoIdentifiers[goIdentifier]

		if hasNameCollision {
			goIdentifier = initialGoIdentifier + fmt.Sprintf("_%d", index)
			x.HasNameConflict = true
		}
	}
	setOfGoIdentifiers[goIdentifier] = nil
	x.GoIdentifier = goIdentifier
}
