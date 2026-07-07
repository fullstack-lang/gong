package models

import (
	"fmt"
	"time"
)

// enforceDefaultValues enforce defaut values when there are not suitable
func (stager *Stager) enforceDefaultValues() (needCommit bool) {
	const (
		defaultBoxWidth  = 250.0
		defaultBoxHeigth = 70.0
	)

	root := stager.getRootLibrary()
	if root.NbPixPerCharacter == 0 {
		root.NbPixPerCharacter = 8
		needCommit = true

		stager.probeForm.AddNotification(time.Now(),
			fmt.Sprintf("Root: setting nbPixPerCharacter to %f", root.NbPixPerCharacter))
	}

	for _, plant := range GetGongstrucsSorted[*Plant](stager.stage) {
		if plant.N == 0 {
			needCommit = true
			plant.N = 1
		}
		if plant.M == 0 {
			needCommit = true
			plant.M = 1
		}
		if plant.Z == 0 {
			needCommit = true
			plant.Z = 5
		}
		if plant.SideLength == 0.0 {
			needCommit = true
			plant.SideLength = 100.0
		}

	}

	return
}
