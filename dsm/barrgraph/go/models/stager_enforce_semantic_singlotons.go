package models

import "time"

func (stager *Stager) enforce_semantic_singlotons() (needCommit bool) {
	stage := stager.stage
	if len(GetGongstrucsSorted[*Desk](stager.stage)) == 0 {
		(&Desk{Name: "Desk"}).Stage(stager.stage)
		needCommit = true
		if stage.probeIF != nil {
			stager.probeForm.AddNotification(time.Now(), "No Desk was found, creating one")
		}
	}
	stager.desk = GetGongstrucsSorted[*Desk](stager.stage)[0]

	// at least one diagram is welcome to ease the end user experience
	if len(GetGongstrucsSorted[*Diagram](stager.stage)) == 0 {
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
