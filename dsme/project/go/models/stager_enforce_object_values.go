package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceObjectValues() (needCommit bool) {
	const (
		defaultBoxWidth  = 250.0
		defaultBoxHeigth = 100.0
	)

	root := stager.root
	if root.NbPixPerCharacter == 0 {
		root.NbPixPerCharacter = 8
		needCommit = true

		stager.probeForm.AddNotification(time.Now(),
			fmt.Sprintf("Root: setting nbPixPerCharacter to %f", root.NbPixPerCharacter))
	}

	for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage) {
		if diagram.DefaultBoxHeigth == 0 {
			diagram.DefaultBoxHeigth = defaultBoxHeigth
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Diagram %s: setting default box height to %f", diagram.Name, defaultBoxHeigth))
			}
		}
		if diagram.DefaultBoxWidth == 0 {
			diagram.DefaultBoxWidth = defaultBoxWidth
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Diagram %s: setting default box width to %f", diagram.Name, defaultBoxWidth))
			}
		}
	}
	return
}
