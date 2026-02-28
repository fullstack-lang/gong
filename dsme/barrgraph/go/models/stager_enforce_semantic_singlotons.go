package models

func (stager *Stager) enforce_semantic_singlotons() (needCommit bool) {
	if len(GetGongstrucsSorted[*Desk](stager.stage)) == 0 {
		(&Desk{Name: "Desk"}).Stage(stager.stage)
		needCommit = true
	}
	stager.desk = GetGongstrucsSorted[*Desk](stager.stage)[0]

	// at least one diagram is welcome to ease the end user experience
	if len(GetGongstrucsSorted[*Diagram](stager.stage)) == 0 {
		diagram := (&Diagram{
			Name: "Default",
		}).Stage(stager.stage)
		stager.desk.SelectedDiagram = diagram
		needCommit = true
	}

	return
}
