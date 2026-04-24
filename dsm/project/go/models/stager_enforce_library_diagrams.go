package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceLibraryHasAtLeastOneDiagram() (needCommit bool) {
	for _, library := range GetGongstrucsSorted[*Library](stager.stage) {
		if len(library.Diagrams) == 0 {
			for diagram_ := range *GetGongstructInstancesSet[Diagram](stager.stage) {
				diagram_.IsChecked = false
			}
			newDiagram := &Diagram{
				Name:        "Default Diagram",
				IsChecked:   true,
				IsEditable_: true,
				AbstractTypeFields: AbstractTypeFields{
					isExpanded: true,
				},
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
