package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceLibraryHasAtLeastOneDiagram() (needCommit bool) {
	for _, library := range GetGongstrucsSorted[*Library](stager.stage) {
		if len(library.Diagrams) == 0 {
			for diagramHierarchy_ := range *GetGongstructInstancesSet[DiagramHierarchy](stager.stage) {
				diagramHierarchy_.IsChecked = false
			}
			newDiagramHierarchy := &DiagramHierarchy{
				Name:        "Default Diagram",
				IsChecked:   true,
				IsEditable_: true,
				AbstractTypeFields: AbstractTypeFields{
					IsExpanded: true,
				},
			}
			newDiagramHierarchy.Stage(stager.stage)
			library.Diagrams = append(library.Diagrams, newDiagramHierarchy)

			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Created 'Default Diagram' for library '%s'", library.Name))
			}
			needCommit = true
		}
	}
	return
}
