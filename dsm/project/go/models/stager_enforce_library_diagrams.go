package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceLibraryHasAtLeastOneDiagram() (needCommit bool) {
	for _, library := range stager.stage.GetInstancesSorted[*Library]() {
		if len(library.Diagrams) == 0 {
			for diagram_ := range *stager.stage.GetInstancesSet[*Diagram]() {
				diagram_.IsChecked = false
			}
			newDiagram := &Diagram{
				Name:               "Default Diagram",
				IsChecked:          true,
				IsEditable_:        true,
				IsInAutoLayoutMode: true,
				IsExpanded:         true,
			}
			newDiagram.Stage(stager.stage)
			library.Diagrams = append(library.Diagrams, newDiagram)

			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Created 'Default Diagram' for library '%s'", library.Name))
			}
			needCommit = true
		}
	}
	return
}
