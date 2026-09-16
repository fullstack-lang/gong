package models

import "time"

func (stager *Stager) enforce_semantic_singlotons() (needCommit bool) {
	stage := stager.stage
	if len(stager.stage.GetInstancesSorted[*Desk]()) == 0 {
		(&Desk{Name: "Desk"}).Stage(stager.stage)
		needCommit = true
		if stage.probeIF != nil {
			stager.probeForm.AddNotification(time.Now(), "No Desk was found, creating one")
		}
	}
	stager.desk = stager.stage.GetInstancesSorted[*Desk]()[0]

	// at least one diagram is welcome to ease the end user experience
	if len(stager.stage.GetInstancesSorted[*Diagram]()) == 0 {
		diagram := (&Diagram{
			Name: "Default",
		}).Stage(stager.stage)
		stager.desk.SelectedDiagram = diagram
		needCommit = true
		if stage.probeIF != nil {
			stager.probeForm.AddNotification(time.Now(), "No Diagram was found, creating one and setting it as the selected diagram of the Desk")
		}
	}

	return
}
