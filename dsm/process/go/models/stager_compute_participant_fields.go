package models

func (stager *Stager) computeParticipantFields() {
	stage := stager.stage

	// reset fields
	for _, participant := range stage.GetInstancesSorted[*Participant]() {
		participant.owningProcess = nil
		participant.inDataFlows = nil
		participant.outDataFlows = nil
	}

	// compute owningProcess
	for _, process := range stage.GetInstancesSorted[*Process]() {
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

	for _, dataFlow := range stage.GetInstancesSorted[*DataFlow]() {
		if startExternalParticipant := dataFlow.StartExternalParticipant; startExternalParticipant != nil {
			startExternalParticipant.outDataFlows = append(startExternalParticipant.outDataFlows, dataFlow)
		}
		if endExternalParticipant := dataFlow.EndExternalParticipant; endExternalParticipant != nil {
			endExternalParticipant.inDataFlows = append(endExternalParticipant.inDataFlows, dataFlow)
		}
	}
}
