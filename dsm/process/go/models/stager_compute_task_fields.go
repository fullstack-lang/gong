package models

func (stager *Stager) computeTaskFields() {
	stage := stager.stage

	// reset fields
	for _, task := range stage.GetInstancesSorted[*Task]() {
		task.inControlFlows = nil
		task.outControlFlows = nil
		task.inDataFlows = nil
		task.outDataFlows = nil
		task.owningParticipant = nil
	}

	// compute owningParticipant
	for _, participant := range stage.GetInstancesSorted[*Participant]() {
		for _, task := range participant.Tasks {
			if task != nil {
				task.owningParticipant = participant
			}
		}
	}

	// compute outControlFlows and inControlFlows
	for _, controlFlow := range stage.GetInstancesSorted[*ControlFlow]() {
		if controlFlow.Start != nil {
			controlFlow.Start.outControlFlows = append(controlFlow.Start.outControlFlows, controlFlow)
		}
		if controlFlow.End != nil {
			controlFlow.End.inControlFlows = append(controlFlow.End.inControlFlows, controlFlow)
		}
	}

	// compute outDataFlows and inDataFlows
	for _, dataFlow := range stage.GetInstancesSorted[*DataFlow]() {
		if dataFlow.StartTask != nil {
			dataFlow.StartTask.outDataFlows = append(dataFlow.StartTask.outDataFlows, dataFlow)
		}
		if dataFlow.EndTask != nil {
			dataFlow.EndTask.inDataFlows = append(dataFlow.EndTask.inDataFlows, dataFlow)
		}
	}

}
