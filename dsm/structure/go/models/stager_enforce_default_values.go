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

	for _, diagramsystem := range GetGongstrucsSorted[*DiagramStructure](stager.stage) {
		if diagramsystem.DefaultBoxHeigth == 0 {
			diagramsystem.DefaultBoxHeigth = defaultBoxHeigth
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("DiagramStructure %s: setting default box height to %f", diagramsystem.Name, defaultBoxHeigth))
			}
		}
		if diagramsystem.DefaultBoxWidth == 0 {
			diagramsystem.DefaultBoxWidth = defaultBoxWidth
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("DiagramStructure %s: setting default box width to %f", diagramsystem.Name, defaultBoxWidth))
			}
		}
	}
	return
}
