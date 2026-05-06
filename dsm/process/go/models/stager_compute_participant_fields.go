package models

func (stager *Stager) computeParticipantFields() {
	stage := stager.stage

	// reset fields
	for _, participant := range GetGongstrucsSorted[*Participant](stage) {
		participant.owningProcess = nil
	}

	// compute owningProcess
	for _, process := range GetGongstrucsSorted[*Process](stage) {
		for _, participant := range process.Participants {
			if participant != nil {
				participant.owningProcess = process
			}
		}
		for _, externalParticipant := range process.ExternalParticipants {
			if externalParticipant != nil {
				externalParticipant.owningProcess = process
			}
		}
	}
}
