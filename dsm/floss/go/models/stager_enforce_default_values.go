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

		if stager.probeForm != nil {
			stager.probeForm.AddNotification(time.Now(),
				fmt.Sprintf("Root: setting nbPixPerCharacter to %f", root.NbPixPerCharacter))
		}
	}

	for _, diagramsystem := range GetGongstrucsSorted[*DiagramFloss](stager.stage) {
		if !diagramsystem.IsEditable_ {
			diagramsystem.IsEditable_ = true
			needCommit = true
		}
		if diagramsystem.DefaultBoxHeigth == 0 {
			diagramsystem.DefaultBoxHeigth = defaultBoxHeigth
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("DiagramFloss %s: setting default box height to %f", diagramsystem.Name, defaultBoxHeigth))
			}
		}
		if diagramsystem.DefaultBoxWidth == 0 {
			diagramsystem.DefaultBoxWidth = defaultBoxWidth
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("DiagramFloss %s: setting default box width to %f", diagramsystem.Name, defaultBoxWidth))
			}
		}
	}

	for _, diagramEquation := range GetGongstrucsSorted[*DiagramFlossEquation](stager.stage) {
		if !diagramEquation.IsEditable_ {
			diagramEquation.IsEditable_ = true
			needCommit = true
		}
		if diagramEquation.Width == 0 {
			diagramEquation.Width = 1000.0
			needCommit = true
		}
		if diagramEquation.Height == 0 {
			diagramEquation.Height = 750.0
			needCommit = true
		}
		if diagramEquation.Scale == 0 {
			diagramEquation.Scale = 5.0
			needCommit = true
		}
		if diagramEquation.DefaultBoxHeigth == 0 {
			diagramEquation.DefaultBoxHeigth = defaultBoxHeigth
			needCommit = true
		}
		if diagramEquation.DefaultBoxWidth == 0 {
			diagramEquation.DefaultBoxWidth = defaultBoxWidth
			needCommit = true
		}
	}

	return
}

