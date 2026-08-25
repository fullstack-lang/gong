package models

import (
	"fmt"
	"time"
)

// enforceDefaultValues enforce default values when they are not suitable
func (stager *Stager) enforceDefaultValues() (needCommit bool) {
	const (
		defaultBoxWidth  = 250.0
		defaultBoxHeigth = 70.0
	)

	root := stager.getRootLibrary()
	if root != nil && root.NbPixPerCharacter == 0 {
		root.NbPixPerCharacter = 8
		needCommit = true

		if stager.probeForm != nil {
			stager.probeForm.AddNotification(time.Now(),
				fmt.Sprintf("Root: setting nbPixPerCharacter to %f", root.NbPixPerCharacter))
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
		if diagramEquation.FontSize == "" {
			diagramEquation.FontSize = FONT_SIZE_NORMAL
			needCommit = true
		}
	}

	return
}
