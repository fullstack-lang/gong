package models

func (stager *Stager) computeTaskFields() {
	stage := stager.stage

	// reset fields
	for _, task := range GetGongstrucsSorted[*Task](stage) {
		task.outcontrolFlows = nil
		task.incontrolFlows = nil
		task.owningParticipant = nil
	}

	// compute owningParticipant
	for _, participant := range GetGongstrucsSorted[*Participant](stage) {
		for _, task := range participant.Tasks {
			if task != nil {
				task.owningParticipant = participant
			}
		}
	}

	// compute outcontrolFlows and incontrolFlows
	for _, controlFlow := range GetGongstrucsSorted[*ControlFlow](stage) {
		if controlFlow.Start != nil {
			controlFlow.Start.outcontrolFlows = append(controlFlow.Start.outcontrolFlows, controlFlow)
		}
		if controlFlow.End != nil {
			controlFlow.End.incontrolFlows = append(controlFlow.End.incontrolFlows, controlFlow)
		}
	}
}
