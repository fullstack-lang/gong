package models

func (stager *Stager) computeTaskFields() {
	stage := stager.stage

	// reset fields
	for _, task := range GetGongstrucsSorted[*Task](stage) {
		task.inControlFlows = nil
		task.outControlFlows = nil
		task.inDataFlows = nil
		task.outDataFlows = nil
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

	// compute outControlFlows and inControlFlows
	for _, controlFlow := range GetGongstrucsSorted[*ControlFlow](stage) {
		if controlFlow.Start != nil {
			controlFlow.Start.outControlFlows = append(controlFlow.Start.outControlFlows, controlFlow)
		}
		if controlFlow.End != nil {
			controlFlow.End.inControlFlows = append(controlFlow.End.inControlFlows, controlFlow)
		}
	}

	// compute outDataFlows and inDataFlows
	for _, dataFlow := range GetGongstrucsSorted[*DataFlow](stage) {
		if dataFlow.Start != nil {
			dataFlow.Start.outDataFlows = append(dataFlow.Start.outDataFlows, dataFlow)
		}
		if dataFlow.End != nil {
			dataFlow.End.inDataFlows = append(dataFlow.End.inDataFlows, dataFlow)
		}
	}

}
