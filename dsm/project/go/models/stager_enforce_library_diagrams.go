package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceLibraryHasAtLeastOneDiagram() (needCommit bool) {
	for _, library := range stager.stage.GetInstancesSorted[*Library]() {
		if len(library.Diagrams) == 0 {
			hasCheckedDiagram := false
			for diagram_ := range *stager.stage.GetInstancesSet[*Diagram]() {
				if diagram_.IsChecked {
					hasCheckedDiagram = true
					break
				}
			}
			newDiagram := &Diagram{
				Name:               "Default Diagram",
				IsChecked:          !hasCheckedDiagram,
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

func (stager *Stager) enforceAtMostOneDiagramChecked() (needCommit bool) {
	hasCheckedDiagram := false
	for _, diagram := range stager.stage.GetInstancesSorted[*Diagram]() {
		if diagram.IsChecked {
			if !hasCheckedDiagram {
				hasCheckedDiagram = true
			} else {
				diagram.IsChecked = false
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Unchecked diagram '%s' to enforce at most one checked diagram", diagram.Name))
				}
				needCommit = true
			}
		}
	}
	return
}
