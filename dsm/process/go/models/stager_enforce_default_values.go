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

	root := stager.rootLibrary
	if root.NbPixPerCharacter == 0 {
		root.NbPixPerCharacter = 8
		needCommit = true

		stager.probeForm.AddNotification(time.Now(),
			fmt.Sprintf("Root: setting nbPixPerCharacter to %f", root.NbPixPerCharacter))
	}

	for _, diagramprocess := range GetGongstrucsSorted[*DiagramProcess](stager.stage) {
		if diagramprocess.DefaultBoxHeigth == 0 {
			diagramprocess.DefaultBoxHeigth = defaultBoxHeigth
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("DiagramProcess %s: setting default box height to %f", diagramprocess.Name, defaultBoxHeigth))
			}
		}
		if diagramprocess.DefaultBoxWidth == 0 {
			diagramprocess.DefaultBoxWidth = defaultBoxWidth
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("DiagramProcess %s: setting default box width to %f", diagramprocess.Name, defaultBoxWidth))
			}
		}
	}
	return
}
